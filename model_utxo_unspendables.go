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

type UtxoUnspendables struct {
	// Transactions overridden by duplicates
	Bip30 float64 `json:"bip30,omitempty"`
	// Unspendable amount of Genesis block subsidy
	GenesisBlock float64 `json:"genesis_block,omitempty"`
	// Amounts sent to unspendable scripts
	Scripts float64 `json:"scripts,omitempty"`
	// Fee rewards unclaimed by miners
	UnclaimedRewards float64 `json:"unclaimed_rewards,omitempty"`
}
