# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSatoshi**](SatoshisApi.md#GetSatoshi) | **Get** /sat/{number} | Get satoshi info

# **GetSatoshi**
> GetSatoshiResponse GetSatoshi(ctx, number)
Get satoshi info

Retrieve information about a specific satoshi

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **number** | **int32**| Satoshi Number | 

### Return type

[**GetSatoshiResponse**](GetSatoshiResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

