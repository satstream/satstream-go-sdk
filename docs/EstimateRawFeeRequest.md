# EstimateRawFeeRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfTarget** | **int32** | Confirmation target in blocks (1 - 1008) | [default to null]
**Threshold** | **float64** | The proportion of transactions in a given feerate range that must have been confirmed within conf_target in order to consider those feerates as high enough and proceed to check lower buckets. Optional, defaults to 0.95 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

