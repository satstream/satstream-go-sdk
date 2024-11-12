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

// Detailed information about a transaction
type GithubComSatstreamSsUtilsStoreTransactionDocument struct {
	Address string `json:"address,omitempty"`
	Blockheight int32 `json:"blockheight,omitempty"`
	Fee float32 `json:"fee,omitempty"`
	Hash string `json:"hash,omitempty"`
	Hex string `json:"hex,omitempty"`
	Index int32 `json:"index,omitempty"`
	Locktime int32 `json:"locktime,omitempty"`
	Size int32 `json:"size,omitempty"`
	Txid string `json:"txid,omitempty"`
	Version int32 `json:"version,omitempty"`
	Vsize int32 `json:"vsize,omitempty"`
	Weight int32 `json:"weight,omitempty"`
}