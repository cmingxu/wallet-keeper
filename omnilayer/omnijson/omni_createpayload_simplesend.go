package omnijson

/*
Result:
"payload"             (string) the hex-encoded payload
*/

type OmniCreatePayloadSimpleSendResult = string

type OmniCreatePayloadSimpleSendCommand struct {
	Property int32
	Amount   string
}

func (OmniCreatePayloadSimpleSendCommand) Method() string {
	return "omni_createpayload_simplesend"
}

func (OmniCreatePayloadSimpleSendCommand) ID() string {
	return "1"
}

func (cmd OmniCreatePayloadSimpleSendCommand) Params() []interface{} {
	return []interface{}{cmd.Property, cmd.Amount}
}
