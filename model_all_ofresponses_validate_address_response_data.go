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

// Address validation result
type AllOfresponsesValidateAddressResponseData struct {
	// The bitcoin address validated
	Address string `json:"address,omitempty"`
	// Error message, if any
	Error_ string `json:"error,omitempty"`
	// Indices of likely error locations
	ErrorLocations []int32 `json:"error_locations,omitempty"`
	// If the key is a script
	Isscript bool `json:"isscript,omitempty"`
	// If the address is valid or not
	Isvalid bool `json:"isvalid,omitempty"`
	// If the address is a witness address
	Iswitness bool `json:"iswitness,omitempty"`
	// The hex-encoded scriptPubKey
	ScriptPubKey string `json:"scriptPubKey,omitempty"`
	// The hex value of the witness program
	WitnessProgram string `json:"witness_program,omitempty"`
	// The version number of the witness program
	WitnessVersion int32 `json:"witness_version,omitempty"`
}