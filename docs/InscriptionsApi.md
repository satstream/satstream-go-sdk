# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchInscriptionChild**](InscriptionsApi.md#FetchInscriptionChild) | **Get** /inscription/{inscription_id}/{child_index} | Get inscription child info
[**FetchInscriptions**](InscriptionsApi.md#FetchInscriptions) | **Post** /inscriptions | Fetch multiple inscriptions
[**GetBlockInscriptions**](InscriptionsApi.md#GetBlockInscriptions) | **Get** /inscriptions/block/{block_height} | Get inscriptions in a specific block
[**GetBlockInscriptionsPage**](InscriptionsApi.md#GetBlockInscriptionsPage) | **Get** /inscriptions/block/{block_height}/{page} | Get paginated inscriptions in a specific block
[**GetInscription**](InscriptionsApi.md#GetInscription) | **Get** /inscription/{inscription_id} | Get inscription info
[**GetLatestInscriptions**](InscriptionsApi.md#GetLatestInscriptions) | **Get** /inscriptions | Get latest inscriptions
[**GetLatestInscriptionsPage**](InscriptionsApi.md#GetLatestInscriptionsPage) | **Get** /inscriptions/{page} | Get latest inscriptions page

# **FetchInscriptionChild**
> GetInscriptionChildResponse FetchInscriptionChild(ctx, inscriptionId, childIndex)
Get inscription child info

Retrieve information about a specific child of an inscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inscriptionId** | **string**| Inscription ID | 
  **childIndex** | **int32**| Child Index | 

### Return type

[**GetInscriptionChildResponse**](GetInscriptionChildResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchInscriptions**
> FetchInscriptionsResponse FetchInscriptions(ctx, body)
Fetch multiple inscriptions

Retrieve information about multiple inscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)| Inscription IDs | 

### Return type

[**FetchInscriptionsResponse**](FetchInscriptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockInscriptions**
> GetBlockInscriptionsResponse GetBlockInscriptions(ctx, blockHeight)
Get inscriptions in a specific block

Retrieve all inscriptions in a specific block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockHeight** | **int32**| Block Height | 

### Return type

[**GetBlockInscriptionsResponse**](GetBlockInscriptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockInscriptionsPage**
> GetBlockInscriptionsResponse GetBlockInscriptionsPage(ctx, blockHeight, page)
Get paginated inscriptions in a specific block

Retrieve paginated inscriptions in a specific block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockHeight** | **int32**| Block Height | 
  **page** | **int32**| Page Number | 

### Return type

[**GetBlockInscriptionsResponse**](GetBlockInscriptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInscription**
> GetInscriptionResponse GetInscription(ctx, inscriptionId)
Get inscription info

Get information about a specific inscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inscriptionId** | **string**| Inscription ID | 

### Return type

[**GetInscriptionResponse**](GetInscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestInscriptions**
> GetLatestInscriptionsResponse GetLatestInscriptions(ctx, )
Get latest inscriptions

Retrieve the latest 100 inscriptions (first page)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetLatestInscriptionsResponse**](GetLatestInscriptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestInscriptionsPage**
> GetLatestInscriptionsResponse GetLatestInscriptionsPage(ctx, page)
Get latest inscriptions page

Retrieve a specific page of 100 inscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page number | 

### Return type

[**GetLatestInscriptionsResponse**](GetLatestInscriptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

