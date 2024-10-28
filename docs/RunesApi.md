# \RunesApi

All URIs are relative to *https://localhost:8085/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunesGet**](RunesApi.md#RunesGet) | **Get** /runes | Get runes info list
[**RunesRuneIdGet**](RunesApi.md#RunesRuneIdGet) | **Get** /runes/{runeId} | Get rune info
[**RunesRuneIdHoldersGet**](RunesApi.md#RunesRuneIdHoldersGet) | **Get** /runes/{runeId}/holders | Get rune holders
[**RunesStatusGet**](RunesApi.md#RunesStatusGet) | **Get** /runes/status | Get runes status


# **RunesGet**
> []interface{} RunesGet(ctx, )
Get runes info list

Get information about all runes

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunesRuneIdGet**
> interface{} RunesRuneIdGet(ctx, runeId)
Get rune info

Get detailed information about a specific rune

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **runeId** | **string**| Rune ID | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunesRuneIdHoldersGet**
> []interface{} RunesRuneIdHoldersGet(ctx, runeId)
Get rune holders

Get a list of addresses holding a specific rune

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **runeId** | **string**| Rune ID | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunesStatusGet**
> interface{} RunesStatusGet(ctx, )
Get runes status

Get the current status of the runes system

### Required Parameters
This endpoint does not need any parameter.

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

