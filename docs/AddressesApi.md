# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddress**](AddressesApi.md#GetAddress) | **Get** /address/{address} | Get address info
[**GetAddressBalance**](AddressesApi.md#GetAddressBalance) | **Get** /address/{address}/balance | Get address balance
[**GetAddressDeltas**](AddressesApi.md#GetAddressDeltas) | **Get** /address/{address}/deltas | Get address deltas
[**GetAddressRuneDeltas**](AddressesApi.md#GetAddressRuneDeltas) | **Get** /address/{address}/deltas/runes | Get address rune deltas
[**GetAddressUtxos**](AddressesApi.md#GetAddressUtxos) | **Get** /address/{address}/outputs | Get UTXOs for an address
[**ValidateAddress**](AddressesApi.md#ValidateAddress) | **Get** /address/{address}/validate | Validate address

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

# **GetAddressBalance**
> GetAddressBalanceResponse GetAddressBalance(ctx, address)
Get address balance

Get the total BTC balance in satoshis of an address by summing all its deltas

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address | 

### Return type

[**GetAddressBalanceResponse**](GetAddressBalanceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressDeltas**
> GetAddressDeltasResponse GetAddressDeltas(ctx, address, optional)
Get address deltas

Get deltas for a specific address with pagination

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address | 
 **optional** | ***AddressesApiGetAddressDeltasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressesApiGetAddressDeltasOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| Number of results per page (default: 100, max: 1000) | 
 **startHeight** | **optional.Int32**| Start block height | 
 **endHeight** | **optional.Int32**| End block height | 
 **cursor** | **optional.String**| Base64 encoded cursor for pagination | 

### Return type

[**GetAddressDeltasResponse**](GetAddressDeltasResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressRuneDeltas**
> GetAddressRuneDeltasResponse GetAddressRuneDeltas(ctx, address, optional)
Get address rune deltas

Get rune deltas for a specific address with pagination

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address | 
 **optional** | ***AddressesApiGetAddressRuneDeltasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddressesApiGetAddressRuneDeltasOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| Number of results per page (default: 100, max: 1000) | 
 **startHeight** | **optional.Int32**| Start block height | 
 **endHeight** | **optional.Int32**| End block height | 
 **cursor** | **optional.String**| Cursor for pagination | 

### Return type

[**GetAddressRuneDeltasResponse**](GetAddressRuneDeltasResponse.md)

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

