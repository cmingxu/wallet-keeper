package notifier

type EthBalanceChangeEvent struct {
	meta map[string]string
}

func NewEthBalanceChangeEvent(meta map[string]string) *EthBalanceChangeEvent {
	return &EthBalanceChangeEvent{
		meta: meta,
	}
}

func (ethEvent *EthBalanceChangeEvent) GetEvent() map[string]string {
	return ethEvent.meta
}

func (ethEvent *EthBalanceChangeEvent) Type() string {
	return "eth_balance_change_event"
}
