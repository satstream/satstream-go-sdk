# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BroadcastTransaction**](TransactionsApi.md#BroadcastTransaction) | **Post** /transactions/broadcast | Broadcast transaction
[**GetTransaction**](TransactionsApi.md#GetTransaction) | **Get** /indexer/tx/{hash} | Get transaction
[**GetTransactionInfo**](TransactionsApi.md#GetTransactionInfo) | **Get** /transactions/{txid} | Get transaction info

# **BroadcastTransaction**
> ResponsesSendRawTransaction BroadcastTransaction(ctx, body)
Broadcast transaction

Broadcast a raw transaction to the Bitcoin network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| Raw transaction hex | 

### Return type

[**ResponsesSendRawTransaction**](responses.SendRawTransaction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransaction**
> ResponsesGetTransaction GetTransaction(ctx, hash)
Get transaction

Get a transaction by its hash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| Transaction hash | 

### Return type

[**ResponsesGetTransaction**](responses.GetTransaction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionInfo**
> ResponsesGetTxInfo GetTransactionInfo(ctx, txid)
Get transaction info

Get detailed information about a specific transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**ResponsesGetTxInfo**](responses.GetTxInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

