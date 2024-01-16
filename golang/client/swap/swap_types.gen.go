// Package swap provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package swap

// Defines values for QuoteRequestErrorStatusCode.
const (
	QuoteRequestErrorStatusCodeN400 QuoteRequestErrorStatusCode = 400
	QuoteRequestErrorStatusCodeN500 QuoteRequestErrorStatusCode = 500
)

// Defines values for SwapRequestErrorStatusCode.
const (
	SwapRequestErrorStatusCodeN400 SwapRequestErrorStatusCode = 400
	SwapRequestErrorStatusCodeN500 SwapRequestErrorStatusCode = 500
)

// AllowanceResponse defines model for AllowanceResponse.
type AllowanceResponse struct {
	// Allowance Allowance amount
	Allowance string `json:"allowance"`
}

// ApproveCallDataResponse defines model for ApproveCallDataResponse.
type ApproveCallDataResponse struct {
	// Data The encoded data to call the approve method on the swapped token contract
	Data string `json:"data"`

	// GasPrice Network high gas price in wei
	GasPrice string `json:"gasPrice"`

	// To Token address that will be allowed to exchange through 1inch Router
	To string `json:"to"`

	// Value Native token value in wei (for approve is always 0)
	Value string `json:"value"`
}

// HttpExceptionMeta defines model for HttpExceptionMeta.
type HttpExceptionMeta struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// ProtocolImage defines model for ProtocolImage.
type ProtocolImage struct {
	// Id Protocol id
	Id string `json:"id"`

	// Img Protocol logo image
	Img string `json:"img"`

	// ImgColor Protocol logo image in color
	ImgColor string `json:"img_color"`

	// Title Protocol title
	Title string `json:"title"`
}

// ProtocolsResponse defines model for ProtocolsResponse.
type ProtocolsResponse struct {
	// Protocols List of protocols that are available for routing in the 1inch Aggregation Protocol
	Protocols []ProtocolImage `json:"protocols"`
}

// QuoteRequestError defines model for QuoteRequestError.
type QuoteRequestError struct {
	// Description Error description
	Description string `json:"description"`

	// Error Error code description
	Error string `json:"error"`

	// Meta Meta information
	Meta []HttpExceptionMeta `json:"meta"`

	// RequestId Request id
	RequestId string `json:"requestId"`

	// StatusCode HTTP code
	StatusCode QuoteRequestErrorStatusCode `json:"statusCode"`
}

// QuoteRequestErrorStatusCode HTTP code
type QuoteRequestErrorStatusCode float32

// QuoteResponse defines model for QuoteResponse.
type QuoteResponse struct {
	// FromToken Source token info
	FromToken *TokenInfo `json:"fromToken,omitempty"`

	// Gas Estimated gas
	Gas *float32 `json:"gas,omitempty"`

	// Protocols Selected protocols in a path
	Protocols *[][][]SelectedProtocol `json:"protocols,omitempty"`

	// ToAmount Expected amount of destination token
	ToAmount string `json:"toAmount"`

	// ToToken Destination token info
	ToToken *TokenInfo `json:"toToken,omitempty"`
}

// SelectedProtocol defines model for SelectedProtocol.
type SelectedProtocol struct {
	// FromTokenAddress Source token to swap on protocol
	FromTokenAddress string `json:"fromTokenAddress"`

	// Name Protocol id
	Name string `json:"name"`

	// Part Protocol share
	Part float32 `json:"part"`

	// ToTokenAddress Destination token to swap on protocol
	ToTokenAddress string `json:"toTokenAddress"`
}

// SpenderResponse defines model for SpenderResponse.
type SpenderResponse struct {
	// Address Address of the 1inch Router that is trusted to spend funds for the swap
	Address string `json:"address"`
}

// SwapRequestError defines model for SwapRequestError.
type SwapRequestError struct {
	// Description Error description
	Description string `json:"description"`

	// Error Error code description
	Error string `json:"error"`

	// Meta Meta information
	Meta []HttpExceptionMeta `json:"meta"`

	// RequestId Request id
	RequestId string `json:"requestId"`

	// StatusCode HTTP code
	StatusCode SwapRequestErrorStatusCode `json:"statusCode"`
}

// SwapRequestErrorStatusCode HTTP code
type SwapRequestErrorStatusCode float32

// SwapResponse defines model for SwapResponse.
type SwapResponse struct {
	// FromToken Source token info
	FromToken *TokenInfo `json:"fromToken,omitempty"`

	// Protocols Selected protocols in a path
	Protocols *[][][]SelectedProtocol `json:"protocols,omitempty"`

	// ToAmount Expected amount of destination token
	ToAmount string `json:"toAmount"`

	// ToToken Destination token info
	ToToken *TokenInfo `json:"toToken,omitempty"`

	// Tx Transaction object
	Tx TransactionData `json:"tx"`
}

// TokenInfo defines model for TokenInfo.
type TokenInfo struct {
	Address       string    `json:"address"`
	Decimals      float32   `json:"decimals"`
	DomainVersion *string   `json:"domainVersion,omitempty"`
	Eip2612       *bool     `json:"eip2612,omitempty"`
	IsFoT         *bool     `json:"isFoT,omitempty"`
	LogoURI       string    `json:"logoURI"`
	Name          string    `json:"name"`
	Symbol        string    `json:"symbol"`
	Tags          *[]string `json:"tags,omitempty"`
}

// TokensResponse defines model for TokensResponse.
type TokensResponse struct {
	Tokens map[string]TokenInfo `json:"tokens"`
}

// TransactionData defines model for TransactionData.
type TransactionData struct {
	Data     string  `json:"data"`
	From     string  `json:"from"`
	Gas      float32 `json:"gas"`
	GasPrice string  `json:"gasPrice"`
	To       string  `json:"to"`
	Value    string  `json:"value"`
}

// ApproveControllerGetAllowanceParams defines parameters for ApproveControllerGetAllowance.
type ApproveControllerGetAllowanceParams struct {
	// TokenAddress Token address you want to swap
	TokenAddress string `url:"tokenAddress" json:"tokenAddress"`

	// WalletAddress Wallet address for which you want to check
	WalletAddress string `url:"walletAddress" json:"walletAddress"`
}

// ApproveControllerGetCallDataParams defines parameters for ApproveControllerGetCallData.
type ApproveControllerGetCallDataParams struct {
	// TokenAddress Token address you want to swap
	TokenAddress string `url:"tokenAddress" json:"tokenAddress"`

	// Amount The number of tokens that the 1inch Router is allowed to swap.If not specified, it will be allowed to spend an infinite amount of tokens.
	Amount *string `url:"amount,omitempty" json:"amount,omitempty"`
}

// AggregationControllerGetQuoteParams defines parameters for AggregationControllerGetQuote.
type AggregationControllerGetQuoteParams struct {
	Src    string `url:"src" json:"src"`
	Dst    string `url:"dst" json:"dst"`
	Amount string `url:"amount" json:"amount"`

	// Protocols All supported liquidity sources by default
	Protocols *string `url:"protocols,omitempty" json:"protocols,omitempty"`

	// Fee Partner fee. min: 0; max: 3 Should be the same for /quote and /swap
	Fee *float32 `url:"fee,omitempty" json:"fee,omitempty"`

	// GasPrice Network price per gas in wei. By default fast network gas price
	GasPrice        *string  `url:"gasPrice,omitempty" json:"gasPrice,omitempty"`
	ComplexityLevel *float32 `url:"complexityLevel,omitempty" json:"complexityLevel,omitempty"`
	Parts           *float32 `url:"parts,omitempty" json:"parts,omitempty"`
	MainRouteParts  *float32 `url:"mainRouteParts,omitempty" json:"mainRouteParts,omitempty"`
	GasLimit        *float32 `url:"gasLimit,omitempty" json:"gasLimit,omitempty"`

	// IncludeTokensInfo Return fromToken and toToken info in response
	IncludeTokensInfo *bool `url:"includeTokensInfo,omitempty" json:"includeTokensInfo,omitempty"`

	// IncludeProtocols Return used swap protocols in response
	IncludeProtocols *bool `url:"includeProtocols,omitempty" json:"includeProtocols,omitempty"`

	// IncludeGas Return approximated gas in response
	IncludeGas      *bool   `url:"includeGas,omitempty" json:"includeGas,omitempty"`
	ConnectorTokens *string `url:"connectorTokens,omitempty" json:"connectorTokens,omitempty"`
}

// AggregationControllerGetSwapParams defines parameters for AggregationControllerGetSwap.
type AggregationControllerGetSwapParams struct {
	Src    string `url:"src" json:"src"`
	Dst    string `url:"dst" json:"dst"`
	Amount string `url:"amount" json:"amount"`

	// From The address that calls the 1inch contract
	From string `url:"from" json:"from"`

	// Slippage min: 0; max: 50
	Slippage float32 `url:"slippage" json:"slippage"`

	// Protocols All supported liquidity sources by default
	Protocols *string `url:"protocols,omitempty" json:"protocols,omitempty"`

	// Fee Partner fee. min: 0; max: 3 Should be the same for /quote and /swap
	Fee *float32 `url:"fee,omitempty" json:"fee,omitempty"`

	// GasPrice Network price per gas in wei. By default fast network gas price
	GasPrice        *string  `url:"gasPrice,omitempty" json:"gasPrice,omitempty"`
	ComplexityLevel *float32 `url:"complexityLevel,omitempty" json:"complexityLevel,omitempty"`
	Parts           *float32 `url:"parts,omitempty" json:"parts,omitempty"`
	MainRouteParts  *float32 `url:"mainRouteParts,omitempty" json:"mainRouteParts,omitempty"`
	GasLimit        *float32 `url:"gasLimit,omitempty" json:"gasLimit,omitempty"`

	// IncludeTokensInfo Return fromToken and toToken info in response
	IncludeTokensInfo *bool `url:"includeTokensInfo,omitempty" json:"includeTokensInfo,omitempty"`

	// IncludeProtocols Return used swap protocols in response
	IncludeProtocols *bool `url:"includeProtocols,omitempty" json:"includeProtocols,omitempty"`

	// IncludeGas Return approximated gas in response
	IncludeGas      *bool   `url:"includeGas,omitempty" json:"includeGas,omitempty"`
	ConnectorTokens *string `url:"connectorTokens,omitempty" json:"connectorTokens,omitempty"`

	// Permit https://eips.ethereum.org/EIPS/eip-2612
	Permit *string `url:"permit,omitempty" json:"permit,omitempty"`

	// Receiver This address will receive funds after the swap. By default same address as "from" param
	Receiver *string `url:"receiver,omitempty" json:"receiver,omitempty"`
	Referrer *string `url:"referrer,omitempty" json:"referrer,omitempty"`

	// AllowPartialFill By default set to false
	AllowPartialFill *bool `url:"allowPartialFill,omitempty" json:"allowPartialFill,omitempty"`

	// DisableEstimate Enable this flag to disable onchain simulation
	DisableEstimate *bool `url:"disableEstimate,omitempty" json:"disableEstimate,omitempty"`
}
