# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EstimateRawFee**](FeesApi.md#EstimateRawFee) | **Post** /fee/estimate-raw | Estimate Raw Fee
[**EstimateSmartFee**](FeesApi.md#EstimateSmartFee) | **Post** /fee/estimate-smart | Estimate smart fee

# **EstimateRawFee**
> InlineResponse20011 EstimateRawFee(ctx, body)
Estimate Raw Fee

Estimates the approximate fee per kilobyte needed for a transaction to begin confirmation within conf_target blocks if possible.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsEstimateRawFeeRequest**](RequestsEstimateRawFeeRequest.md)| Fee estimation parameters | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EstimateSmartFee**
> InlineResponse20012 EstimateSmartFee(ctx, body)
Estimate smart fee

Estimates the approximate fee per kilobyte needed for a transaction to begin confirmation within conf_target blocks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsEstimateSmartFeeRequest**](RequestsEstimateSmartFeeRequest.md)| Fee estimation parameters | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

