package omnijson

type OmniCreateRawTxOpReturnResult = string

type OmniCreateRawTxOpReturnCommand struct {
	Raw     string
	Payload string
}

func (OmniCreateRawTxOpReturnCommand) Method() string {
	return "omni_createrawtx_opreturn"
}

func (OmniCreateRawTxOpReturnCommand) ID() string {
	return "1"
}

func (cmd OmniCreateRawTxOpReturnCommand) Params() []interface{} {
	return []interface{}{cmd.Raw, cmd.Payload}
}
