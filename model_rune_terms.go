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

type RuneTerms struct {
	Amount *BigInt `json:"amount,omitempty"`
	Cap *BigInt `json:"cap,omitempty"`
	Height []int32 `json:"height,omitempty"`
	Offset []int32 `json:"offset,omitempty"`
}
