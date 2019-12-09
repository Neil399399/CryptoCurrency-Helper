package ethereum

import (
	"context"
	"errors"
	"log"
	"math/big"
	"strconv"
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
							events[address] = append(events[address], event)
						}
					}
				}
			}
		}
	}
	return events, nil
}
