package omni

import (
	"encoding/json"
	"log"

	"github.com/Neil399399/bitcoin-helper/omni/omnijson"
)

func (c *Client) GetInfo() (omnijson.OmniGetInfoResult, error) {
	var result omnijson.OmniGetInfoResult
	data, err := receive(c.do(omnijson.OmniGetInfoCommand{}))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) GetBalance(address string, propertyId int32) (omnijson.OmniGetBalanceResult, error) {
	var result omnijson.OmniGetBalanceResult
	data, err := receive(c.do(omnijson.OmniGetBalanceCommand{
		Address:    address,
		PropertyID: propertyId,
	}))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) GetAddressTxids(addresses []string) ([]omnijson.OmniAddressTxidsResult, error) {
	var result []omnijson.OmniAddressTxidsResult
	data, err := receive(c.do(omnijson.OmniAddressTxidsCommand{
		Addresses: addresses,
	}))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) GetAllBalancesForAddress(address string) (omnijson.OmniGetAllBalancesForAddressResult, error) {
	var result omnijson.OmniGetAllBalancesForAddressResult
	data, err := receive(c.do(omnijson.OmniGetAllBalancesForAddressCommand{
		Address: address,
	}))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) ListUnSpent(addresses []string) (omnijson.OmniListUnSpentResult, error) {
	var result omnijson.OmniListUnSpentResult
	data, err := receive(c.do(omnijson.OmniListUnspentCommand{
		Addresses: addresses,
	}))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) GetTransaction(txid string) (omnijson.OmniGetTransactionResult, error) {
	var result omnijson.OmniGetTransactionResult
	data, err := receive(c.do(omnijson.OmniGetTransactionCommand{
		TxID: txid,
	}))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) CreatePayload(propertyId int32, amount string) (omnijson.OmniCreatePayloadResult, error) {
	var result omnijson.OmniCreatePayloadResult
	data, err := receive(c.do(omnijson.OmniCreatePayloadCommand{
		PropertyId: propertyId,
		Amount:     amount,
	}))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) CreateOpReturn(txid, payload string) (omnijson.OmniCreateOpReturnResult, error) {
	var result omnijson.OmniCreateOpReturnResult
	data, err := receive(c.do(omnijson.OmniCreateOpReturnCommand{
		TxID:    txid,
		Payload: payload,
	}))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) CreateRawTxInput(rawTx string, txid string, number int32) (omnijson.OmniCreateInputResult, error) {
	var result omnijson.OmniCreateInputResult
	data, err := receive(c.do(omnijson.OmniCreateInputCommand{
		RawTx: rawTx,
		TxID:  txid,
		N:     number,
	}))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) SendTransaction(from string, to string, propertyId int32, amount string) (omnijson.OmniSendTransactionResult, error) {
	var result omnijson.OmniSendTransactionResult
	data, err := receive(c.do(omnijson.OmniSendTransactionCommand{
		FromAddress: from,
		ToAddress:   to,
		PropertyID:  propertyId,
		Amount:      amount,
	}))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) SendRawTx(fromAddr, toAddr, tx, redeemAddr, amount string) (omnijson.OmniSendRawTransactionResult, error) {
	var result omnijson.OmniSendRawTransactionResult
	data, err := receive(c.do(omnijson.OmniSendRawTransactionCommand{
		FromAddress:      fromAddr,
		RawTx:            tx,
		ReferenceAddress: toAddr,
		RedeemAddress:    redeemAddr,
		ReferenceAmount:  amount,
	}))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = json.Unmarshal(data, &result)
	return result, err
}
