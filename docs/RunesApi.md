# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestRunes**](RunesApi.md#GetLatestRunes) | **Get** /runes | Get latest runes
[**GetLatestRunesPage**](RunesApi.md#GetLatestRunesPage) | **Get** /runes/{page} | Get latest runes page
[**GetRune**](RunesApi.md#GetRune) | **Get** /rune/{identifier} | Get rune info

# **GetLatestRunes**
> GetLatestRunesResponse GetLatestRunes(ctx, )
Get latest runes

Retrieve information about the last 100 inscribed runes (first page)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetLatestRunesResponse**](GetLatestRunesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestRunesPage**
> GetLatestRunesResponse GetLatestRunesPage(ctx, page)
Get latest runes page

Retrieve a specific page of 100 inscribed runes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page number | 

### Return type

[**GetLatestRunesResponse**](GetLatestRunesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRune**
> GetRuneResponse GetRune(ctx, identifier)
Get rune info

Retrieve information about a specific rune by name or ID (e.g., \"UNCOMMON•GOODS\" or \"1:0\")

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Rune Name or ID | 

### Return type

[**GetRuneResponse**](GetRuneResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

