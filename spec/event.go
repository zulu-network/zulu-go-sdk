package spec

const (
	DepositEvent  = "BRIDGE_DEPOSIT"
	WithdrawEvent = "BRIDGE_WITHDRAW"
)

type Event struct {
	Address  string `json:"address"`
	Identify string `json:"identify"`
	TxHash   string `json:"txHash"`
	Amount   string `json:"amount"`
}
