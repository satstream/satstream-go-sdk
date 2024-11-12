# \BlocksApi

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockByHash**](BlocksApi.md#GetBlockByHash) | **Get** /blocks/hash/{hash} | Get block by hash
[**GetBlockInfo**](BlocksApi.md#GetBlockInfo) | **Get** /blocks/{height} | Get block info
[**GetBlockTransactions**](BlocksApi.md#GetBlockTransactions) | **Get** /blocks/{height}/transactions | Get block transactions
[**GetCurrentBlockHeight**](BlocksApi.md#GetCurrentBlockHeight) | **Get** /blocks/current-height | Get current block height


# **GetBlockByHash**
> InlineResponse2005 GetBlockByHash(ctx, hash)
Get block by hash

Get information about a specific block by its hash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| Block hash | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockInfo**
> InlineResponse2005 GetBlockInfo(ctx, height)
Get block info

Get information about a specific block by height

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **height** | **int32**| Block height | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockTransactions**
> []interface{} GetBlockTransactions(ctx, height)
Get block transactions

Get transactions for a specific block height

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **height** | **int32**| Block height | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentBlockHeight**
> ResponsesBlockHeight GetCurrentBlockHeight(ctx, )
Get current block height

Get the current block height of the Bitcoin blockchain

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponsesBlockHeight**](responses.BlockHeight.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
