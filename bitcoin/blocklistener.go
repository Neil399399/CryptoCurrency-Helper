package bitcoin

import (
	"errors"
	"fmt"
)

// ListenBitcoinChain get input address transcation from block
func (t *BtcTx) ListenBitcoinChain(addressBook []string, endBlockHeight int64) (map[string][]VoutDetail, error) {
	result := map[string][]VoutDetail{}
	currentBlockHeight, err := t.client.GetBlockCount()
	if err != nil {
		return nil, err
	}
	fmt.Println("currentBlockHeight", currentBlockHeight)
	fmt.Println("endBlockHeight", endBlockHeight)

	// The genesis block coinbase is not considered an ordinary transaction
	// and cannot be retrieved. So we should stop at 1st block.
	if endBlockHeight < 1 {
		endBlockHeight = 1
	}

	if endBlockHeight > currentBlockHeight {
		return nil, errors.New("Block doesn't exist")
	}

	currBlockHash, err := t.client.GetBlockHash(currentBlockHeight)
	if err != nil {
		return nil, err
	}

	currBlock, err := t.client.GetBlock(currBlockHash)
	if err != nil {
		return nil, err
	}

	endBlockHash, err := t.client.GetBlockHash(endBlockHeight)
	if err != nil {
		return nil, err
	}

	for {
		for _, tx := range currBlock.Transactions {
			txHash := tx.TxHash()
			txRawResult, err := t.client.GetRawTransactionVerbose(&txHash)
			if err != nil {
				return nil, err
			}

			for _, out := range txRawResult.Vout {
				for _, addr := range addressBook {
					if len(out.ScriptPubKey.Addresses) > 0 && addr == out.ScriptPubKey.Addresses[0] {
						voutDetail := VoutDetail{
							BlockHash:     txRawResult.BlockHash,
							Txid:          txRawResult.Txid,
							Address:       []string{addr},
							Category:      "receive",
							Amount:        out.Value,
							Vout:          out.N,
							Confirmations: txRawResult.Confirmations,
							Time:          txRawResult.Time,
							LockTime:      txRawResult.LockTime,
							Blocktime:     txRawResult.Blocktime,
						}
						result[addr] = append(result[addr], voutDetail)
					}
				}
			}
		}

		if currBlockHash.String() == endBlockHash.String() {
			break
		}

		currBlockHash = &currBlock.Header.PrevBlock
		currBlock, err = t.client.GetBlock(currBlockHash)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
