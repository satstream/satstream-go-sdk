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

// Only if mempool_sequence is true
type AllOfRawMempoolDataSequence struct {
	// The mempool sequence value
	MempoolSequence int32 `json:"mempool_sequence,omitempty"`
	// The transaction ids
	Txids []string `json:"txids,omitempty"`
}
