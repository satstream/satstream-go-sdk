# \TransactionsApi

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IndexerTxHashGet**](TransactionsApi.md#IndexerTxHashGet) | **Get** /indexer/tx/{hash} | Get transaction
[**TransactionsBroadcastPost**](TransactionsApi.md#TransactionsBroadcastPost) | **Post** /transactions/broadcast | Broadcast transaction
[**TransactionsTxidGet**](TransactionsApi.md#TransactionsTxidGet) | **Get** /transactions/{txid} | Get transaction info
[**TransactionsTxidInputsGet**](TransactionsApi.md#TransactionsTxidInputsGet) | **Get** /transactions/{txid}/inputs | Get transaction inputs


# **IndexerTxHashGet**
> RpcBtcTx IndexerTxHashGet(ctx, hash)
Get transaction

Get a transaction by its hash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| Transaction hash | 

### Return type

[**RpcBtcTx**](rpc.BtcTx.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsBroadcastPost**
> ResponsesSendRawTransaction TransactionsBroadcastPost(ctx, transaction)
Broadcast transaction

Broadcast a raw transaction to the Bitcoin network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transaction** | **string**| Raw transaction hex | 

### Return type

[**ResponsesSendRawTransaction**](responses.SendRawTransaction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsTxidGet**
> RpcBtcTx TransactionsTxidGet(ctx, txid)
Get transaction info

Get detailed information about a specific transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**RpcBtcTx**](rpc.BtcTx.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsTxidInputsGet**
> []RpcVin TransactionsTxidInputsGet(ctx, txid)
Get transaction inputs

Get the inputs of a specific transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**[]RpcVin**](rpc.Vin.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

