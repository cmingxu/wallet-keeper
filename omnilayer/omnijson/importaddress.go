package omnijson

type ImportAddressCommand struct {
	Adress string
	Tag    string
	Rescan bool
}

func (ImportAddressCommand) Method() string {
	return "importaddress"
}

func (ImportAddressCommand) ID() string {
	return "1"
}

func (cmd ImportAddressCommand) Params() []interface{} {
	return []interface{}{cmd.Adress, cmd.Tag, cmd.Rescan}
}
