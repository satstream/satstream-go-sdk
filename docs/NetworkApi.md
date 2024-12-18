# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChainTxStats**](NetworkApi.md#GetChainTxStats) | **Post** /chain/txstats | Get chain tx stats
[**GetDifficulty**](NetworkApi.md#GetDifficulty) | **Get** /chain/difficulty | Get difficulty

# **GetChainTxStats**
> GetChainTxStatsResponse GetChainTxStats(ctx, body)
Get chain tx stats

Computes statistics about the total number and rate of transactions in the chain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetChainTxStatsRequest**](GetChainTxStatsRequest.md)| Chain tx stats request parameters | 

### Return type

[**GetChainTxStatsResponse**](GetChainTxStatsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDifficulty**
> GetDifficultyResponse GetDifficulty(ctx, )
Get difficulty

Returns the proof-of-work difficulty as a multiple of the minimum difficulty

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetDifficultyResponse**](GetDifficultyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

