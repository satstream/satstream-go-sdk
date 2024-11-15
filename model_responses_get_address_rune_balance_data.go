/*
 * Satstream API
 *
 * Satstream API
 *
 * API version: 1.0
 * Contact: team@satstream.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package satstream_go_sdk

type ResponsesGetAddressRuneBalanceData struct {
	Amount string `json:"amount,omitempty"`
	Divisibility int32 `json:"divisibility,omitempty"`
	Rune_ string `json:"rune,omitempty"`
	Runeid string `json:"runeid,omitempty"`
	SpacedAmount string `json:"spacedAmount,omitempty"`
	SpacedRune string `json:"spacedRune,omitempty"`
	Symbol string `json:"symbol,omitempty"`
}
