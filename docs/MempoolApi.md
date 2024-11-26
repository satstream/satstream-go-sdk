# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMempoolAncestors**](MempoolApi.md#GetMempoolAncestors) | **Post** /mempool/ancestors | Get mempool ancestors
[**GetMempoolDescendants**](MempoolApi.md#GetMempoolDescendants) | **Post** /mempool/descendants | Get mempool descendants
[**GetMempoolInfo**](MempoolApi.md#GetMempoolInfo) | **Get** /mempool/info | Get mempool information
[**GetRawMempool**](MempoolApi.md#GetRawMempool) | **Post** /mempool/raw | Get raw mempool
[**TestMempoolAccept**](MempoolApi.md#TestMempoolAccept) | **Post** /mempool/test-accept | Test mempool accept

# **GetMempoolAncestors**
> GetMempoolAncestorsResponse GetMempoolAncestors(ctx, body)
Get mempool ancestors

Returns all in-mempool ancestors for a transaction in the mempool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetMempoolAncestorsRequest**](GetMempoolAncestorsRequest.md)| Mempool ancestors request parameters | 

### Return type

[**GetMempoolAncestorsResponse**](GetMempoolAncestorsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMempoolDescendants**
> GetMempoolDescendantsResponse GetMempoolDescendants(ctx, body)
Get mempool descendants

Returns all in-mempool descendants for a transaction in the mempool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetMempoolDescendantsRequest**](GetMempoolDescendantsRequest.md)| Mempool descendants request parameters | 

### Return type

[**GetMempoolDescendantsResponse**](GetMempoolDescendantsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMempoolInfo**
> GetMempoolInfoResponse GetMempoolInfo(ctx, )
Get mempool information

Returns details on the active state of the TX memory pool

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetMempoolInfoResponse**](GetMempoolInfoResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawMempool**
> GetRawMempoolResponse GetRawMempool(ctx, body)
Get raw mempool

Returns all transaction ids in memory pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetRawMempoolRequest**](GetRawMempoolRequest.md)| Raw mempool request parameters | 

### Return type

[**GetRawMempoolResponse**](GetRawMempoolResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestMempoolAccept**
> TestMempoolAcceptResponse TestMempoolAccept(ctx, body)
Test mempool accept

Tests whether raw transactions would be accepted by mempool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestMempoolAcceptRequest**](TestMempoolAcceptRequest.md)| Raw transactions to test | 

### Return type

[**TestMempoolAcceptResponse**](TestMempoolAcceptResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

