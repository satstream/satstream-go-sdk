# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOutputByOutpoint**](OutputsApi.md#GetOutputByOutpoint) | **Get** /output/{outpoint} | Get output info by outpoint
[**GetOutputs**](OutputsApi.md#GetOutputs) | **Post** /outputs | Get multiple outputs

# **GetOutputByOutpoint**
> InlineResponse20022 GetOutputByOutpoint(ctx, outpoint)
Get output info by outpoint

Retrieve information about a specific UTXO using outpoint string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **outpoint** | **string**| Outpoint | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOutputs**
> InlineResponse20023 GetOutputs(ctx, body)
Get multiple outputs

Retrieve information about multiple UTXOs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)| Outpoints | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

