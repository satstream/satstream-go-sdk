# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestRunes**](RunesApi.md#GetLatestRunes) | **Get** /runes | Get latest runes
[**GetLatestRunesPage**](RunesApi.md#GetLatestRunesPage) | **Get** /runes/{page} | Get latest runes page
[**GetRune**](RunesApi.md#GetRune) | **Get** /rune/{rune_name} | Get rune info

# **GetLatestRunes**
> InlineResponse20030 GetLatestRunes(ctx, )
Get latest runes

Retrieve information about the last 100 inscribed runes (first page)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestRunesPage**
> InlineResponse20030 GetLatestRunesPage(ctx, page)
Get latest runes page

Retrieve a specific page of 100 inscribed runes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page number | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRune**
> InlineResponse20029 GetRune(ctx, runeName)
Get rune info

Retrieve information about a specific rune

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **runeName** | **string**| Rune Name | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

