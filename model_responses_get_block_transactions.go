/*
 * Satstream API
 *
 * Satstream API
 *
 * API version: 1.0
 * Contact: team@satstream.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package satstream-go-sdk

type ResponsesGetBlockTransactions struct {
	Code int32 `json:"code,omitempty"`
	Data []GithubComSatstreamSsUtilsStoreTransactionDocument `json:"data,omitempty"`
	Msg string `json:"msg,omitempty"`
}
