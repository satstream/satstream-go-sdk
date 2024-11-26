# RequestsCreateRawTxRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | [**[]RequestsCreateRawTxInput**](requests.CreateRawTxInput.md) | The inputs for the transaction | [default to null]
**Locktime** | **int32** | Raw locktime. Non-0 value also locktime-activates inputs Optional, defaults to 0 | [optional] [default to null]
**Outputs** | [**[]RequestsCreateRawTxOutput**](requests.CreateRawTxOutput.md) | The outputs for the transaction Each address can only appear once and there can only be one &#x27;data&#x27; object | [default to null]
**Replaceable** | **bool** | Marks this transaction as BIP125-replaceable Allows this transaction to be replaced by a transaction with higher fees If provided, it is an error if explicit sequence numbers are incompatible Optional, defaults to true | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

