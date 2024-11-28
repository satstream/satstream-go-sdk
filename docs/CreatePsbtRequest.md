# CreatePsbtRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | [**[]CreatePsbtInput**](CreatePSBTInput.md) | The inputs for the transaction | [default to null]
**Locktime** | **int32** | Raw locktime. Non-0 value also locktime-activates inputs | [optional] [default to null]
**Outputs** | [**[]ModelMap**](map.md) | The outputs (address:amount pairs or {\&quot;data\&quot;:\&quot;hex\&quot;}) | [default to null]
**Replaceable** | **bool** | Marks this transaction as BIP125-replaceable | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

