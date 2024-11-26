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
> InlineResponse20024 AnalyzePsbt(ctx, body)
Analyze PSBT

Analyzes and provides information about the current status of a PSBT and its inputs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsAnalyzePsbtRequest**](RequestsAnalyzePsbtRequest.md)| PSBT to analyze | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CombinePsbt**
> InlineResponse2001 CombinePsbt(ctx, body)
Combine PSBTs

Combines multiple partially signed Bitcoin transactions into one transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsCombinePsbtRequest**](RequestsCombinePsbtRequest.md)| Array of PSBTs to combine | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePsbt**
> InlineResponse2001 CreatePsbt(ctx, body)
Create PSBT

Creates a transaction in the Partially Signed Transaction format. Implements the Creator role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsCreatePsbtRequest**](RequestsCreatePsbtRequest.md)| Transaction parameters | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DecodePsbt**
> InlineResponse20025 DecodePsbt(ctx, body)
Decode PSBT

Return a JSON object representing the serialized, base64-encoded partially signed Bitcoin transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsAnalyzePsbtRequest**](RequestsAnalyzePsbtRequest.md)| PSBT to decode | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JoinPsbts**
> InlineResponse2001 JoinPsbts(ctx, body)
Join PSBTs

Joins multiple distinct PSBTs with different inputs and outputs into one PSBT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsJoinPsbtsRequest**](RequestsJoinPsbtsRequest.md)| PSBTs to join | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

