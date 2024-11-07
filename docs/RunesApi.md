# \RunesApi

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunesGet**](RunesApi.md#RunesGet) | **Get** /runes | Get runes info list
[**RunesRuneIdGet**](RunesApi.md#RunesRuneIdGet) | **Get** /runes/{runeId} | Get rune info
[**RunesRuneIdHoldersGet**](RunesApi.md#RunesRuneIdHoldersGet) | **Get** /runes/{runeId}/holders | Get rune holders


# **RunesGet**
> InlineResponse20010 RunesGet(ctx, optional)
Get runes info list

Get information about all runes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RunesApiRunesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RunesApiRunesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Page number (default: 1) | 
 **perPage** | **optional.Int32**| Items per page (default: 10) | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunesRuneIdGet**
> interface{} RunesRuneIdGet(ctx, runeId)
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

# **RunesRuneIdHoldersGet**
> []interface{} RunesRuneIdHoldersGet(ctx, runeId)
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

