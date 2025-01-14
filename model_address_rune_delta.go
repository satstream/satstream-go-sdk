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

type AddressRuneDelta struct {
	BlockHeight int32 `json:"block_height,omitempty"`
	Delta string `json:"delta,omitempty"`
	Incoming string `json:"incoming,omitempty"`
	Outgoing string `json:"outgoing,omitempty"`
	RuneId string `json:"rune_id,omitempty"`
}
