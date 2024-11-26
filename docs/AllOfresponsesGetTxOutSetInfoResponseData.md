# AllOfresponsesGetTxOutSetInfoResponseData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bestblock** | **string** | The hash of the block at which these statistics are calculated | [optional] [default to null]
**BlockInfo** | [***interface{}**](interface{}.md) | Info on amounts in the block at this height | [optional] [default to null]
**Bogosize** | **int32** | Database-independent metric indicating the UTXO set size | [optional] [default to null]
**DiskSize** | **int32** | The estimated size of the chainstate on disk | [optional] [default to null]
**HashSerialized2** | **string** | The serialized hash (only for hash_serialized_2) | [optional] [default to null]
**Height** | **int32** | The block height of the returned statistics | [optional] [default to null]
**Muhash** | **string** | The serialized hash (only for muhash) | [optional] [default to null]
**TotalAmount** | **float64** | The total amount of coins in the UTXO set | [optional] [default to null]
**TotalUnspendableAmount** | **float64** | Total amount permanently excluded from UTXO set | [optional] [default to null]
**Transactions** | **int32** | The number of transactions with unspent outputs | [optional] [default to null]
**Txouts** | **int32** | The number of unspent transaction outputs | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

