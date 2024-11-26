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

type GithubComSatstreamSsUtilsBitcoinCliPsbtWitnessUtxo struct {
	// The value in BTC
	Amount float64 `json:"amount,omitempty"`
	// The script pub key
	ScriptPubKey *AllOfgithubComSatstreamSsUtilsBitcoinCliPsbtWitnessUtxoScriptPubKey `json:"scriptPubKey,omitempty"`
}