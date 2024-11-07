# \AddressesApi

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressesAddressBalanceGet**](AddressesApi.md#AddressesAddressBalanceGet) | **Get** /addresses/{address}/balance | Get address balance
[**AddressesAddressBalanceTimeframeGet**](AddressesApi.md#AddressesAddressBalanceTimeframeGet) | **Get** /addresses/{address}/balance/timeframe | Get address timeframe balance
[**AddressesAddressRunesGet**](AddressesApi.md#AddressesAddressRunesGet) | **Get** /addresses/{address}/runes | Get address runes balance list
[**AddressesAddressRunesRuneidGet**](AddressesApi.md#AddressesAddressRunesRuneidGet) | **Get** /addresses/{address}/runes/{runeid} | Get address rune balance
[**AddressesAddressUtxosGet**](AddressesApi.md#AddressesAddressUtxosGet) | **Get** /addresses/{address}/utxos | Get address non-inscription UTXOs


# **AddressesAddressBalanceGet**
> InlineResponse200 AddressesAddressBalanceGet(ctx, address)
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

# **AddressesAddressBalanceTimeframeGet**
> InlineResponse2001 AddressesAddressBalanceTimeframeGet(ctx, address, timeframe, optional)
Get address timeframe balance

Get the balance of a Bitcoin address for a specific timeframe

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Bitcoin address | 
  **timeframe** | **string**| Timeframe | 
 **optional** | ***AddressesApiAddressesAddressBalanceTimeframeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressesApiAddressesAddressBalanceTimeframeGetOpts struct

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

# **AddressesAddressRunesGet**
> InlineResponse2002 AddressesAddressRunesGet(ctx, address)
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

# **AddressesAddressRunesRuneidGet**
> InlineResponse2003 AddressesAddressRunesRuneidGet(ctx, address, runeid)
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

# **AddressesAddressUtxosGet**
> InlineResponse2004 AddressesAddressUtxosGet(ctx, address)
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

