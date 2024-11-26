# ValidateAddressResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The bitcoin address validated | [optional] [default to null]
**Error_** | **string** | Error message, if any | [optional] [default to null]
**ErrorLocations** | **[]int32** | Indices of likely error locations | [optional] [default to null]
**Isscript** | **bool** | If the key is a script | [optional] [default to null]
**Isvalid** | **bool** | If the address is valid or not | [optional] [default to null]
**Iswitness** | **bool** | If the address is a witness address | [optional] [default to null]
**ScriptPubKey** | **string** | The hex-encoded scriptPubKey | [optional] [default to null]
**WitnessProgram** | **string** | The hex value of the witness program | [optional] [default to null]
**WitnessVersion** | **int32** | The version number of the witness program | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

