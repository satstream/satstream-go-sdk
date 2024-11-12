# \AddressesApi

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddressBalance**](AddressesApi.md#GetAddressBalance) | **Get** /addresses/{address}/balance | Get address balance
[**GetAddressNonInscriptionUtxos**](AddressesApi.md#GetAddressNonInscriptionUtxos) | **Get** /addresses/{address}/utxos | Get address non-inscription UTXOs
[**GetAddressRuneBalance**](AddressesApi.md#GetAddressRuneBalance) | **Get** /addresses/{address}/runes/{runeid} | Get address rune balance
[**GetAddressRunesBalanceList**](AddressesApi.md#GetAddressRunesBalanceList) | **Get** /addresses/{address}/runes | Get address runes balance list
[**GetAddressTimeframeBalance**](AddressesApi.md#GetAddressTimeframeBalance) | **Get** /addresses/{address}/balance/timeframe | Get address timeframe balance


# **GetAddressBalance**
> InlineResponse200 GetAddressBalance(ctx, address)
Get address balance

Get the current balance of a Bitcoin address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Bitcoin address | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressNonInscriptionUtxos**
> InlineResponse2004 GetAddressNonInscriptionUtxos(ctx, address)
Get address non-inscription UTXOs

Get all non-inscription UTXOs for a Bitcoin address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Bitcoin address | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressRuneBalance**
> InlineResponse2003 GetAddressRuneBalance(ctx, address, runeid)
Get address rune balance

Get the balance of a specific rune for a Bitcoin address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Bitcoin address | 
  **runeid** | **string**| Rune ID | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressRunesBalanceList**
> InlineResponse2002 GetAddressRunesBalanceList(ctx, address)
Get address runes balance list

Get the balance of all runes for a Bitcoin address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Bitcoin address | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressTimeframeBalance**
> InlineResponse2001 GetAddressTimeframeBalance(ctx, address, timeframe, optional)
Get address timeframe balance

Get the balance of a Bitcoin address for a specific timeframe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Bitcoin address | 
  **timeframe** | **string**| Timeframe | 
 **optional** | ***AddressesApiGetAddressTimeframeBalanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressesApiGetAddressTimeframeBalanceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **token** | **optional.String**| Token | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
