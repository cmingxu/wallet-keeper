package omnilayer

import "github.com/cmingxu/wallet-keeper/omnilayer/omnijson"

func (c *Client) OmniFoundedSend(from, to string, propertyid int64, amount, fee string) (omnijson.OmniFoundedSendResult, error) {
	return futureOmniFoundedSend(c.do(omnijson.OmniFoundedSendCommand{
		From:     from,
		To:       to,
		ProperID: propertyid,
		Amount:   amount,
		Fee:      fee,
	})).Receive()
}

func (c *Client) ListAccounts(minConf int64) (omnijson.ListAccountsResult, error) {
	return futureListAccounts(c.do(omnijson.ListAccountsCommand{
		MinConf: minConf,
	})).Receive()
}

func (c *Client) GetAddressesByAccount(account string) (omnijson.GetAddressesByAccountResult, error) {
	return futureGetAddressesByAccount(c.do(omnijson.GetAddressesByAccountCommand{
		Account: account,
	})).Receive()
}

func (c *Client) GetAccountAddress(account string) (omnijson.GetAccountAddressResult, error) {
	return futureGetAccountAddress(c.do(omnijson.GetAccountAddressCommand{
		Account: account,
	})).Receive()
}

func (c *Client) GetNewAddress(account string) (omnijson.GetNewAddressResult, error) {
	return futureGetNewAddress(c.do(omnijson.GetNewAddressCommand{
		Account: account,
	})).Receive()
}

func (c *Client) GetBlockChainInfo() (omnijson.GetBlockChainInfoResult, error) {
	return futureGetBlockChainInfo(c.do(omnijson.GetBlockChainInfoCommand{})).Receive()
}

func (c *Client) OmniListBlockTransactions(block int64) (omnijson.OmniListBlockTransactionsResult, error) {
	return futureOmniListBlockTransactions(c.do(omnijson.OmniListBlockTransactionsCommand{
		Block: block,
	})).Receive()
}

func (c *Client) GetInfo() (omnijson.OmniGetInfoResult, error) {
	return futureGetInfo(c.do(omnijson.OmniGetInfoCommand{})).Receive()
}

func (c *Client) OmniGetTransaction(hash string) (omnijson.OmniGettransactionResult, error) {
	return futureOmniGetTransaction(c.do(omnijson.OmniGetTransactionCommand{
		Hash: hash,
	})).Receive()
}

func (c *Client) ListUnspent(cmd omnijson.ListUnspentCommand) (omnijson.ListUnspentResult, error) {
	return futureListUnspent(c.do(cmd)).Receive()
}

func (c *Client) OmniCreatePayloadSimpleSend(cmd omnijson.OmniCreatePayloadSimpleSendCommand) (omnijson.OmniCreatePayloadSimpleSendResult, error) {
	return futureOmniCreatePayloadSimpleSend(c.do(cmd)).Receive()
}

func (c *Client) CreateRawTransaction(cmd omnijson.CreateRawTransactionCommand) (omnijson.CreateRawTransactionResult, error) {
	return futureCreateRawTransaction(c.do(cmd)).Receive()
}

func (c *Client) OmniCreateRawTxOpReturn(cmd omnijson.OmniCreateRawTxOpReturnCommand) (omnijson.OmniCreateRawTxOpReturnResult, error) {
	return futureOmniCreateRawTxOpReturn(c.do(cmd)).Receive()
}

func (c *Client) OmniCreateRawTxReference(cmd omnijson.OmniCreateRawTxReferenceCommand) (omnijson.OmniCreateRawTxReferenceResult, error) {
	return futureOmniCreateRawTxReference(c.do(cmd)).Receive()
}

func (c *Client) OmniCreateRawTxChange(cmd omnijson.OmniCreateRawTxChangeCommand) (omnijson.OmniCreateRawTxChangeResult, error) {
	return futureOmniCreateRawTxChange(c.do(cmd)).Receive()
}

func (c *Client) ImportAddress(address string, rescan bool) error {
	return futureImportAddress(c.do(omnijson.ImportAddressCommand{Adress: address, Rescan: rescan})).Receive()
}

func (c *Client) SendRawTransaction(cmd omnijson.SendRawTransactionCommand) (omnijson.SendRawTransactionResult, error) {
	return futureSendRawTransaction(c.do(cmd)).Receive()
}

func (c *Client) SignRawTransaction(cmd omnijson.SignRawTransactionCommand) (omnijson.SignRawTransactionResult, error) {
	return futureSignRawTransaction(c.do(cmd)).Receive()
}

func (c *Client) SignRawTransactionWithKey(cmd omnijson.SignRawTransactionWithKeyCommand) (omnijson.SignRawTransactionWithKeyResult, error) {
	return futureSignRawTransactionWithKey(c.do(cmd)).Receive()
}

func (c *Client) OmniGetBalance(cmd omnijson.OmniGetBalanceCommand) (omnijson.OmniGetBalanceResult, error) {
	return futureOmniGetBalance(c.do(cmd)).Receive()
}
