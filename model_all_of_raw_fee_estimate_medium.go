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

// Estimate for medium time horizon
type AllOfRawFeeEstimateMedium struct {
	// Exponential decay (per block) for historical moving average
	Decay float64 `json:"decay,omitempty"`
	// Errors encountered during processing
	Errors []string `json:"errors,omitempty"`
	// Information about highest range of feerates to fail
	Fail *interface{} `json:"fail,omitempty"`
	// Estimate fee rate in BTC/kvB
	Feerate float64 `json:"feerate,omitempty"`
	// Information about lowest range of feerates to succeed
	Pass *interface{} `json:"pass,omitempty"`
	// Resolution of confirmation targets at this time horizon
	Scale float64 `json:"scale,omitempty"`
}
