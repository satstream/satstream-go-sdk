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
> InlineResponse20020 GetMempoolAncestors(ctx, body)
Get mempool ancestors

Returns all in-mempool ancestors for a transaction in the mempool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetMempoolAncestorsRequest**](RequestsGetMempoolAncestorsRequest.md)| Mempool ancestors request parameters | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMempoolDescendants**
> InlineResponse20021 GetMempoolDescendants(ctx, body)
Get mempool descendants

Returns all in-mempool descendants for a transaction in the mempool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetMempoolDescendantsRequest**](RequestsGetMempoolDescendantsRequest.md)| Mempool descendants request parameters | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMempoolInfo**
> InlineResponse20022 GetMempoolInfo(ctx, )
Get mempool information

Returns details on the active state of the TX memory pool

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawMempool**
> InlineResponse20023 GetRawMempool(ctx, body)
Get raw mempool

Returns all transaction ids in memory pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetRawMempoolRequest**](RequestsGetRawMempoolRequest.md)| Raw mempool request parameters | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestMempoolAccept**
> InlineResponse20024 TestMempoolAccept(ctx, body)
Test mempool accept

Tests whether raw transactions would be accepted by mempool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsTestMempoolAcceptRequest**](RequestsTestMempoolAcceptRequest.md)| Raw transactions to test | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

