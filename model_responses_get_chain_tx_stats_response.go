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

type ResponsesGetChainTxStatsResponse struct {
	Code int32 `json:"code,omitempty"`
	// Chain transaction statistics
	Data *AllOfresponsesGetChainTxStatsResponseData `json:"data,omitempty"`
	Msg string `json:"msg,omitempty"`
}