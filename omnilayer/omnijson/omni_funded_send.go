package omnijson

//OmniFoundedSendResult ...
type OmniFoundedSendResult = string

//OmniFoundedSendCommand ...
type OmniFoundedSendCommand struct {
	From     string `json:"fromaddress"`
	To       string `json:"toaddress"`
	ProperID int64  `json:"propertyid"`
	Amount   string `json:"amount"`
	Fee      string `json:"feeaddress"`
}

//Method ...
func (OmniFoundedSendCommand) Method() string {
	return "omni_funded_send"
}

//ID ...
func (OmniFoundedSendCommand) ID() string {
	return "1"
}

//Params ...
func (cmd OmniFoundedSendCommand) Params() []interface{} {
	return []interface{}{cmd.From, cmd.To, cmd.ProperID, cmd.Amount, cmd.Fee}
}
