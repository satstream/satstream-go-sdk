# MempoolEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ancestorcount** | **int32** | Number of in-mempool ancestor transactions | [optional] [default to null]
**Ancestorsize** | **int32** | Virtual size of in-mempool ancestors | [optional] [default to null]
**Bip125Replaceable** | **bool** | Whether this transaction is replaceable | [optional] [default to null]
**Depends** | **[]string** | Parent transaction IDs | [optional] [default to null]
**Descendantcount** | **int32** | Number of in-mempool descendant transactions | [optional] [default to null]
**Descendantsize** | **int32** | Virtual size of in-mempool descendants | [optional] [default to null]
**Fees** | [***AllOfMempoolEntryFees**](AllOfMempoolEntryFees.md) | Fee information | [optional] [default to null]
**Height** | **int32** | Block height when transaction entered pool | [optional] [default to null]
**Spentby** | **[]string** | Child transaction IDs | [optional] [default to null]
**Time** | **int32** | Time transaction entered pool | [optional] [default to null]
**Unbroadcast** | **bool** | Whether this transaction is currently unbroadcast | [optional] [default to null]
**Vsize** | **int32** | Virtual transaction size | [optional] [default to null]
**Weight** | **int32** | Transaction weight | [optional] [default to null]
**Wtxid** | **string** | Hash of serialized transaction with witness data | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

