# \BlocksApi

All URIs are relative to *https://localhost:8085/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlocksCurrentHeightGet**](BlocksApi.md#BlocksCurrentHeightGet) | **Get** /blocks/current-height | Get current block height
[**BlocksHashHashGet**](BlocksApi.md#BlocksHashHashGet) | **Get** /blocks/hash/{hash} | Get block by hash
[**BlocksHeightGet**](BlocksApi.md#BlocksHeightGet) | **Get** /blocks/{height} | Get block info
[**BlocksHeightTransactionsGet**](BlocksApi.md#BlocksHeightTransactionsGet) | **Get** /blocks/{height}/transactions | Get block transactions


# **BlocksCurrentHeightGet**
> int32 BlocksCurrentHeightGet(ctx, )
Get current block height

Get the current block height of the Bitcoin blockchain

### Required Parameters
This endpoint does not need any parameter.

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlocksHashHashGet**
> RpcBlock BlocksHashHashGet(ctx, hash)
Get block by hash

Get information about a specific block by its hash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| Block hash | 

### Return type

[**RpcBlock**](rpc.Block.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlocksHeightGet**
> RpcBlock BlocksHeightGet(ctx, height)
Get block info

Get information about a specific block by height

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **height** | **int32**| Block height | 

### Return type

[**RpcBlock**](rpc.Block.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlocksHeightTransactionsGet**
> []StoreTransactionDocument BlocksHeightTransactionsGet(ctx, height)
Get block transactions

Get transactions for a specific block height

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **height** | **int32**| Block height | 

### Return type

[**[]StoreTransactionDocument**](store.TransactionDocument.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

