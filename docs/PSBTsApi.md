# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzePsbt**](PSBTsApi.md#AnalyzePsbt) | **Post** /psbt/analyze | Analyze PSBT
[**CombinePsbt**](PSBTsApi.md#CombinePsbt) | **Post** /psbt/combine | Combine PSBTs
[**CreatePsbt**](PSBTsApi.md#CreatePsbt) | **Post** /psbt/create | Create PSBT
[**DecodePsbt**](PSBTsApi.md#DecodePsbt) | **Post** /psbt/decode | Decode PSBT
[**JoinPsbts**](PSBTsApi.md#JoinPsbts) | **Post** /psbt/join | Join PSBTs

# **AnalyzePsbt**
> AnalyzePsbtResponse AnalyzePsbt(ctx, body)
Analyze PSBT

Analyzes and provides information about the current status of a PSBT and its inputs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AnalyzePsbtRequest**](AnalyzePsbtRequest.md)| PSBT to analyze | 

### Return type

[**AnalyzePsbtResponse**](AnalyzePSBTResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CombinePsbt**
> CombinePsbtResponse CombinePsbt(ctx, body)
Combine PSBTs

Combines multiple partially signed Bitcoin transactions into one transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CombinePsbtRequest**](CombinePsbtRequest.md)| Array of PSBTs to combine | 

### Return type

[**CombinePsbtResponse**](CombinePSBTResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePsbt**
> CreatePsbtResponse CreatePsbt(ctx, body)
Create PSBT

Creates a transaction in the Partially Signed Transaction format. Implements the Creator role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreatePsbtRequest**](CreatePsbtRequest.md)| Transaction parameters | 

### Return type

[**CreatePsbtResponse**](CreatePSBTResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DecodePsbt**
> DecodePsbtResponse DecodePsbt(ctx, body)
Decode PSBT

Return a JSON object representing the serialized, base64-encoded partially signed Bitcoin transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DecodePsbtRequest**](DecodePsbtRequest.md)| PSBT to decode | 

### Return type

[**DecodePsbtResponse**](DecodePSBTResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JoinPsbts**
> JoinPsbtsResponse JoinPsbts(ctx, body)
Join PSBTs

Joins multiple distinct PSBTs with different inputs and outputs into one PSBT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JoinPsbtsRequest**](JoinPsbtsRequest.md)| PSBTs to join | 

### Return type

[**JoinPsbtsResponse**](JoinPSBTsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

