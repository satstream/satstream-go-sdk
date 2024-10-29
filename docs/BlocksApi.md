# \BlocksApi

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlocksCurrentHeightGet**](BlocksApi.md#BlocksCurrentHeightGet) | **Get** /blocks/current-height | Get current block height
[**BlocksHashHashGet**](BlocksApi.md#BlocksHashHashGet) | **Get** /blocks/hash/{hash} | Get block by hash
[**BlocksHeightGet**](BlocksApi.md#BlocksHeightGet) | **Get** /blocks/{height} | Get block info
[**BlocksHeightTransactionsGet**](BlocksApi.md#BlocksHeightTransactionsGet) | **Get** /blocks/{height}/transactions | Get block transactions


# **BlocksCurrentHeightGet**
> InlineResponse2005 BlocksCurrentHeightGet(ctx, )
Get current block height

Get the current block height of the Bitcoin blockchain

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlocksHashHashGet**
> InlineResponse2006 BlocksHashHashGet(ctx, hash)
Get block by hash

Get information about a specific block by its hash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| Block hash | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlocksHeightGet**
> InlineResponse2006 BlocksHeightGet(ctx, height)
Get block info

Get information about a specific block by height

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **height** | **int32**| Block height | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlocksHeightTransactionsGet**
> []interface{} BlocksHeightTransactionsGet(ctx, height)
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

