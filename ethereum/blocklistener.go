package ethereum

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (e *EthTx) BlockListener(addressBook []string, endBlockHeight int64) (map[string][]TxDetail, error) {
	events := map[string][]TxDetail{}
	ctx := context.Background()

	// get latest block.
	latestBlock, err := e.client.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	// compares latestBlock and endBlock
	lastestBlockNumber := int64(latestBlock.NumberU64())
	if lastestBlockNumber < endBlockHeight {
		return nil, errors.New("Block doesn't exist")
	}

	if lastestBlockNumber < e.blockRange {
		return nil, errors.New("Block doesn't exist")
	}

	// get next block.
	for num := 0; int64(num) < e.blockRange; num++ {
		startBlock := endBlockHeight + int64(num)
		if startBlock > lastestBlockNumber {
			log.Println("WAITTING NEW BLOCK ...")
			break
		}
		block, err := e.client.BlockByNumber(ctx, big.NewInt(startBlock))
		if err != nil {
			log.Println("GET BLOCK BY BLOCKNUMBER FAILED !", num)
			continue
		}
		if block.Transactions().Len() > 0 {
			for i, txn := range block.Transactions() {
				sender, err := e.client.TransactionSender(ctx, txn, block.Hash(), uint(i))
				if err != nil {
					return nil, err
				}
				for _, address := range addressBook {
					if to := txn.To(); to != nil {
						if address == txn.To().Hex() {
							// check target not the contract address.
							event := TxDetail{}
							event.from = sender.Hex()
							event.to = to.Hex()
							event.txnHash = txn.Hash().Hex()
							event.symbol = "wei"
							event.value = txn.Value().Uint64()
							event.blockNumber = block.Number().Uint64()
							event.timeStamp = strconv.FormatUint(block.Time(), 10)
							event.erc = false
							event.memo = string(txn.Data()) // comment
							events[address] = append(events[address], event)
						}
					}
				}
			}
		}
	}
	return events, nil
}

func (e *EthTx) GetContractRecord(contractAddress string, addressBook []string, endBlockHeight int64) (events []TxDetail, err error) {
	ctx := context.Background()
	events = []TxDetail{}

	latestBlock, err := e.client.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	lastestBlockNumber := int64(latestBlock.NumberU64())
	if lastestBlockNumber < endBlockHeight {
		return nil, errors.New("Block doesn't exist")
	}

	if lastestBlockNumber < e.blockRange {
		return nil, errors.New("Block doesn't exist")
	}

	// set contract query
	contractAddr := common.HexToAddress(contractAddress)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(endBlockHeight)),
		ToBlock:   latestBlock.Number(),
		Addresses: []common.Address{
			contractAddr,
		},
	}
	logs, err := e.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	for _, vLog := range logs {
		event := TxDetail{}
		var temp int64
		value := big.NewInt(temp)
		// found transfer event
		if vLog.Topics[0].Hex() == logTransferSigHash.Hex() {
			// match the address
			sender := common.HexToAddress(vLog.Topics[1].Hex()).Hex()
			receiver := common.HexToAddress(vLog.Topics[2].Hex()).Hex()
			fmt.Println("sender", sender)
			fmt.Println("receiver", receiver)

			for _, address := range addressBook {
				if receiver == address {
					value.SetBytes(vLog.Data)
					// set event
					event.from = sender
					event.to = receiver
					event.txnHash = vLog.TxHash.Hex()
					event.symbol = "token"
					event.value = value.Uint64()
					event.blockNumber = vLog.BlockNumber
					// event.timeStamp = strconv.FormatUint(block.Time(), 10)
					event.erc = true
					events = append(events, event)
				}
				continue
			}
		}
	}

	fmt.Println(events)
	return events, nil
}
