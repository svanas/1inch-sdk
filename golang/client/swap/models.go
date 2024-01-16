package swap

type ExecuteSwapConfig struct {
	FromToken          string
	ToToken            string
	Amount             string
	Slippage           float32
	EstimatedAmountOut string
	TransactionData    string
	IsPermitSwap       bool
	SkipWarnings       bool
}