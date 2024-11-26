# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DecodeScript**](ScriptsApi.md#DecodeScript) | **Post** /script/decode | Decode Script

# **DecodeScript**
> DecodeScriptResponse DecodeScript(ctx, body)
Decode Script

Decode a hex-encoded script and return detailed information about it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DecodeScriptRequest**](DecodeScriptRequest.md)| Script to decode | 

### Return type

[**DecodeScriptResponse**](DecodeScriptResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

