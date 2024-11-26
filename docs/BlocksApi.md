# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockByHash**](BlocksApi.md#GetBlockByHash) | **Get** /block/hash/{block_hash} | Get block info by hash

# **GetBlockByHash**
> GithubComSatstreamSsApiServerApiBlockResponsesBlockResponse GetBlockByHash(ctx, blockHash)
Get block info by hash

Get detailed information about a specific block by hash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockHash** | **string**| Block Hash | 

### Return type

[**GithubComSatstreamSsApiServerApiBlockResponsesBlockResponse**](github_com_satstream_ss-api_server_api_block_responses.BlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

