# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddress**](AddressesApi.md#GetAddress) | **Get** /address/{address} | Get address info
[**GetAddressUtxos**](AddressesApi.md#GetAddressUtxos) | **Get** /address/{address}/outputs | Get UTXOs for an address
[**ValidateAddress**](AddressesApi.md#ValidateAddress) | **Get** /address/{address}/validate | Validate address
[**VerifyMessage**](AddressesApi.md#VerifyMessage) | **Post** /address/verify-message | Verify message

# **GetAddress**
> GetAddressResponse GetAddress(ctx, address)
Get address info

Get detailed information about a specific address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address | 

### Return type

[**GetAddressResponse**](GetAddressResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressUtxos**
> GetAddressUtxosResponse GetAddressUtxos(ctx, address, optional)
Get UTXOs for an address

Retrieve UTXOs held by a specific address with optional type filtering

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address | 
 **optional** | ***AddressesApiGetAddressUtxosOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressesApiGetAddressUtxosOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| UTXO Type | 

### Return type

[**GetAddressUtxosResponse**](GetAddressUTXOsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateAddress**
> ValidateAddressResponse ValidateAddress(ctx, address)
Validate address

Returns information about the given Bitcoin address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Bitcoin address to validate | 

### Return type

[**ValidateAddressResponse**](ValidateAddressResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyMessage**
> VerifyMessageResponse VerifyMessage(ctx, body)
Verify message

Verifies a signed message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VerifyMessageRequest**](VerifyMessageRequest.md)| Message verification parameters | 

### Return type

[**VerifyMessageResponse**](VerifyMessageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

