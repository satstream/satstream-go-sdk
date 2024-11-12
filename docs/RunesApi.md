# \RunesApi

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRunesHolders**](RunesApi.md#GetRunesHolders) | **Get** /runes/{runeId}/holders | Get rune holders
[**GetRunesInfo**](RunesApi.md#GetRunesInfo) | **Get** /runes/{runeId} | Get rune info
[**GetRunesInfoList**](RunesApi.md#GetRunesInfoList) | **Get** /runes | Get runes info list


# **GetRunesHolders**
> []interface{} GetRunesHolders(ctx, runeId)
Get rune holders

Get a list of addresses holding a specific rune

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **runeId** | **string**| Rune ID | 

### Return type

**[]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRunesInfo**
> interface{} GetRunesInfo(ctx, runeId)
Get rune info

Get detailed information about a specific rune

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **runeId** | **string**| Rune ID | 

### Return type

**interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRunesInfoList**
> InlineResponse2009 GetRunesInfoList(ctx, optional)
Get runes info list

Get information about all runes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RunesApiGetRunesInfoListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunesApiGetRunesInfoListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number (default: 1) | 
 **perPage** | **optional.Int32**| Items per page (default: 10) | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

