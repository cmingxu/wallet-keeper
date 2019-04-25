package omnijson

/*
Result:
{
  "txid": "84504b62edb18d6b9fa7089c5cba09fabaa7f1f46034ad9d49fb5781d7cf1bc6",
  "fee": "0.00100000",
  "sendingaddress": "1Po1oWkD2LmodfkBYiAktwh76vkF93LKnh",
  "referenceaddress": "1PowyXXycpvSEbfY7cFcuDzpVAh6sTNejo",
  "ismine": false,
  "version": 0,
  "type_int": 0,
  "type": "Simple Send",
  "propertyid": 31,
  "divisible": true,
  "amount": "97.96000000",
  "valid": true,
  "blockhash": "000000000000000000190b41bf8c7b5e1c49275ab71b5fe1aef57864bedf2b5b",
  "blocktime": 1549012217,
  "positioninblock": 38,
  "block": 561034,
  "confirmations": 2
}
*/

type OmniGettransactionResult struct {
	ID              string `json:"txid"`
	Fee             string `json:"fee"`
	From            string `json:"sendingaddress"`
	To              string `json:"referenceaddress"`
	Type            string `json:"type"`
	Amount          string `json:"amount"`
	BlockHash       string `json:"blockhash"`
	InvalidReason   string `json:"invalidreason"`
	Version         int32  `json:"version"`
	TypeInt         int32  `json:"type_int"`
	PropertyID      int32  `json:"propertyid"`
	BlockTimestamp  int32  `json:"blocktime"`
	PositionInBlock int32  `json:"positioninblock"`
	BlockHeight     int32  `json:"block"`
	Confirmations   uint32 `json:"confirmations"`
	Mine            bool   `json:"ismine"`
	Divisible       bool   `json:"divisible"`
	Valid           bool   `json:"valid"`
}

type OmniGetTransactionCommand struct {
	Hash string
}

func (OmniGetTransactionCommand) Method() string {
	return "omni_gettransaction"
}

func (OmniGetTransactionCommand) ID() string {
	return "1"
}

func (cmd OmniGetTransactionCommand) Params() []interface{} {
	return []interface{}{cmd.Hash}
}
