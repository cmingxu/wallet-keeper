package notifier

type EthBalanceChangeEvent struct {
	meta map[string]interface{}
}

func NewEthBalanceChangeEvent(meta map[string]interface{}) *EthBalanceChangeEvent {
	return &EthBalanceChangeEvent{
		meta: meta,
	}
}

func (ethEvent *EthBalanceChangeEvent) GetEvent() map[string]interface{} {
	return ethEvent.meta
}

func (ethEvent *EthBalanceChangeEvent) Type() string {
	return "eth_balance_change_event"
}
