package omnijson

type OmniCreateRawTxChangeResult = string

type OmniCreateRawTxChangeCommand struct {
	Raw         string
	Previous    []OmniCreateRawTxChangeParameter
	Destination string
	Fee         float64
}

type OmniCreateRawTxChangeParameter struct {
	Tx           string  `json:"txid"`
	Vout         uint32  `json:"vout"`
	ScriptPubKey string  `json:"scriptPubKey"`
	Value        float64 `json:"value"`
}

func (OmniCreateRawTxChangeCommand) Method() string {
	return "omni_createrawtx_change"
}

func (OmniCreateRawTxChangeCommand) ID() string {
	return "1"
}

func (cmd OmniCreateRawTxChangeCommand) Params() []interface{} {
	return []interface{}{cmd.Raw, cmd.Previous, cmd.Destination, cmd.Fee}
}
