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

type GithubComSatstreamSsUtilsOrdServerResponsesTransaction struct {
	Input []GithubComSatstreamSsUtilsOrdServerResponsesInput `json:"input,omitempty"`
	LockTime int32 `json:"lock_time,omitempty"`
	Output []GithubComSatstreamSsUtilsOrdServerResponsesOutput `json:"output,omitempty"`
	Version int32 `json:"version,omitempty"`
}