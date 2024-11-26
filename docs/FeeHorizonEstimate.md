# FeeHorizonEstimate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decay** | **float64** | Exponential decay (per block) for historical moving average | [optional] [default to null]
**Errors** | **[]string** | Errors encountered during processing | [optional] [default to null]
**Fail** | [***AllOfFeeHorizonEstimateFail**](AllOfFeeHorizonEstimateFail.md) | Information about highest range of feerates to fail | [optional] [default to null]
**Feerate** | **float64** | Estimate fee rate in BTC/kvB | [optional] [default to null]
**Pass** | [***AllOfFeeHorizonEstimatePass**](AllOfFeeHorizonEstimatePass.md) | Information about lowest range of feerates to succeed | [optional] [default to null]
**Scale** | **float64** | Resolution of confirmation targets at this time horizon | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

