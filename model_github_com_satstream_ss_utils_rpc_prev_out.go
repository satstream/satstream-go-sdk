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

type GithubComSatstreamSsUtilsRpcPrevOut struct {
	Height int32 `json:"height,omitempty"`
	N int32 `json:"n,omitempty"`
	// * The populated field is used to determine if the prevout has been populated from our code or automatically by the RPC on fetch
	Populated bool `json:"populated,omitempty"`
	RuneHoldings []GithubComSatstreamSsUtilsRpcUtxoRune `json:"rune_holdings,omitempty"`
	ScriptPubKey *GithubComSatstreamSsUtilsRpcScriptPubKey `json:"scriptPubKey,omitempty"`
	Value float64 `json:"value,omitempty"`
}
