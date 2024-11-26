# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchInscriptionChild**](InscriptionsApi.md#FetchInscriptionChild) | **Get** /inscription/{inscription_id}/{child_index} | Get inscription child info
[**FetchInscriptions**](InscriptionsApi.md#FetchInscriptions) | **Post** /inscriptions | Fetch multiple inscriptions
[**GetBlockInscriptions**](InscriptionsApi.md#GetBlockInscriptions) | **Get** /inscriptions/block/{block_height} | Get inscriptions in a specific block
[**GetInscription**](InscriptionsApi.md#GetInscription) | **Get** /inscription/{inscription_id} | Get inscription info
[**GetLatestInscriptions**](InscriptionsApi.md#GetLatestInscriptions) | **Get** /inscriptions | Get latest inscriptions
[**GetLatestInscriptionsPage**](InscriptionsApi.md#GetLatestInscriptionsPage) | **Get** /inscriptions/{page} | Get latest inscriptions page

# **FetchInscriptionChild**
> InlineResponse20013 FetchInscriptionChild(ctx, inscriptionId, childIndex)
Get inscription child info

Retrieve information about a specific child of an inscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inscriptionId** | **string**| Inscription ID | 
  **childIndex** | **int32**| Child Index | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchInscriptions**
> InlineResponse20015 FetchInscriptions(ctx, body)
Fetch multiple inscriptions

Retrieve information about multiple inscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)| Inscription IDs | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockInscriptions**
> InlineResponse20014 GetBlockInscriptions(ctx, blockHeight)
Get inscriptions in a specific block

Retrieve all inscriptions in a specific block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockHeight** | **int32**| Block Height | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInscription**
> InlineResponse20013 GetInscription(ctx, inscriptionId)
Get inscription info

Get information about a specific inscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inscriptionId** | **string**| Inscription ID | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestInscriptions**
> InlineResponse20014 GetLatestInscriptions(ctx, )
Get latest inscriptions

Retrieve the latest 100 inscriptions (first page)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestInscriptionsPage**
> InlineResponse20014 GetLatestInscriptionsPage(ctx, page)
Get latest inscriptions page

Retrieve a specific page of 100 inscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page number | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

