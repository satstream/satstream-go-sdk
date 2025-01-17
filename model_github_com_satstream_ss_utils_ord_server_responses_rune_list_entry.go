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

type GithubComSatstreamSsUtilsOrdServerResponsesRuneListEntry struct {
	Block int32 `json:"block,omitempty"`
	Burned int32 `json:"burned,omitempty"`
	Divisibility int32 `json:"divisibility,omitempty"`
	Etching string `json:"etching,omitempty"`
	Id string `json:"id,omitempty"`
	Mints int32 `json:"mints,omitempty"`
	Number int32 `json:"number,omitempty"`
	Premine *BigInt `json:"premine,omitempty"`
	SpacedRune string `json:"spaced_rune,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	Terms *RuneTerms `json:"terms,omitempty"`
	Timestamp int32 `json:"timestamp,omitempty"`
	Turbo bool `json:"turbo,omitempty"`
}
