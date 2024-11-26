# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockCount**](BlocksApi.md#GetBlockCount) | **Get** /blockcount | Get the height of the latest block
[**GetBlockDecoded**](BlocksApi.md#GetBlockDecoded) | **Get** /block/raw/{identifier}/decoded | Get block by hash or height (verbosity 2)
[**GetBlockHashByHeight**](BlocksApi.md#GetBlockHashByHeight) | **Get** /blockhash/{block_height} | Returns blockhash of specified block.
[**GetBlockHex**](BlocksApi.md#GetBlockHex) | **Get** /block/raw/{identifier}/hex | Get block by hash or height (verbosity 0)
[**GetBlockInfo**](BlocksApi.md#GetBlockInfo) | **Get** /block/{identifier} | Get block info by hash or height
[**GetBlockPrevout**](BlocksApi.md#GetBlockPrevout) | **Get** /block/raw/{identifier}/prevout | Get block by hash or height (verbosity 3)
[**GetBlockStats**](BlocksApi.md#GetBlockStats) | **Post** /block/stats | Get block stats
[**GetBlockSummary**](BlocksApi.md#GetBlockSummary) | **Get** /block/raw/{identifier}/summary | Get block by hash or height (verbosity 1)
[**GetBlockchainInfo**](BlocksApi.md#GetBlockchainInfo) | **Get** /blockchain/info | Get blockchain information
[**GetBlocks**](BlocksApi.md#GetBlocks) | **Get** /blocks | Returns the latest block height, last 100 block hashes, and featured inscriptions
[**GetLatestBlockHeight**](BlocksApi.md#GetLatestBlockHeight) | **Get** /blockheight | Returns the height of the latest block.
[**GetLatestBlockhash**](BlocksApi.md#GetLatestBlockhash) | **Get** /blockhash | Returns blockhash for the latest block.
[**GetLatestBlocktime**](BlocksApi.md#GetLatestBlocktime) | **Get** /blocktime | Get the timestamp of the latest block

# **GetBlockCount**
> InlineResponse20011 GetBlockCount(ctx, )
Get the height of the latest block

Returns the height of the latest block

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockDecoded**
> InlineResponse2004 GetBlockDecoded(ctx, identifier)
Get block by hash or height (verbosity 2)

Get block by hash or height as a decoded object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Block hash or height | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockHashByHeight**
> InlineResponse2005 GetBlockHashByHeight(ctx, blockHeight)
Returns blockhash of specified block.

Returns blockhash of specified block.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockHeight** | **string**| Block Height | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockHex**
> InlineResponse2005 GetBlockHex(ctx, identifier)
Get block by hash or height (verbosity 0)

Get block by hash or height as a raw hex string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Block hash or height | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockInfo**
> InlineResponse2009 GetBlockInfo(ctx, identifier)
Get block info by hash or height

Get detailed information about a specific block by hash or height

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Block hash or height | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockPrevout**
> InlineResponse2006 GetBlockPrevout(ctx, identifier)
Get block by hash or height (verbosity 3)

Get block by hash or height with prevout information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Block hash or height | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockStats**
> InlineResponse2008 GetBlockStats(ctx, body)
Get block stats

Computes per block statistics for a given window

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetBlockStatsRequest**](RequestsGetBlockStatsRequest.md)| Block stats request parameters | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockSummary**
> InlineResponse2007 GetBlockSummary(ctx, identifier)
Get block by hash or height (verbosity 1)

Get block by hash or height as a summary object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Block hash or height | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockchainInfo**
> InlineResponse20010 GetBlockchainInfo(ctx, )
Get blockchain information

Returns an object containing various state info regarding blockchain processing

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlocks**
> InlineResponse20012 GetBlocks(ctx, )
Returns the latest block height, last 100 block hashes, and featured inscriptions

Returns the latest block height, last 100 block hashes, and featured inscriptions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestBlockHeight**
> InlineResponse20011 GetLatestBlockHeight(ctx, )
Returns the height of the latest block.

Returns the height of the latest block.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestBlockhash**
> InlineResponse2005 GetLatestBlockhash(ctx, )
Returns blockhash for the latest block.

Returns blockhash for the latest block.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestBlocktime**
> InlineResponse20011 GetLatestBlocktime(ctx, )
Get the timestamp of the latest block

Returns the UNIX timestamp of when the latest block was mined

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

