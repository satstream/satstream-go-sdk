/*
 * Satstream API
 *
 * Satstream API
 *
 * API version: 1.0
 * Contact: team@satstream.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GithubComSatstreamSsUtilsOrdinalsTerms struct {
	Amount *BigInt `json:"amount,omitempty"`
	Cap *BigInt `json:"cap,omitempty"`
	Height *GithubComSatstreamSsUtilsOrdinalsTermsRange `json:"height,omitempty"`
	Offset *GithubComSatstreamSsUtilsOrdinalsTermsRange `json:"offset,omitempty"`
}