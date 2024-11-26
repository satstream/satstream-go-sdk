# DecodedPsbt

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | **float64** | The transaction fee paid if all UTXOs slots are filled | [optional] [default to null]
**Inputs** | [**[]DecodedPsbtInput**](DecodedPSBTInput.md) | Array of inputs | [optional] [default to null]
**Outputs** | [**[]DecodedPsbtOutput**](DecodedPSBTOutput.md) | Array of outputs | [optional] [default to null]
**Tx** | [***AllOfDecodedPsbttx**](AllOfDecodedPsbttx.md) | The decoded network-serialized unsigned transaction | [optional] [default to null]
**Unknown** | [**ModelMap**](interface{}.md) | The unknown global fields | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

