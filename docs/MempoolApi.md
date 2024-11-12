# \MempoolApi

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddressMempoolTransactions**](MempoolApi.md#GetAddressMempoolTransactions) | **Get** /mempool/addresses/{address}/transactions | Get address mempool transactions
[**GetMempoolTransactionInfo**](MempoolApi.md#GetMempoolTransactionInfo) | **Get** /mempool/transactions/{txid} | Get mempool transaction info
[**GetMempoolTransactions**](MempoolApi.md#GetMempoolTransactions) | **Get** /mempool/transactions | Get mempool transactions


# **GetAddressMempoolTransactions**
> []interface{} GetAddressMempoolTransactions(ctx, address)
Get address mempool transactions

Get all mempool transactions for a specific address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Bitcoin address | 

### Return type

**[]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMempoolTransactionInfo**
> InlineResponse2009 GetMempoolTransactionInfo(ctx, txid)
Get mempool transaction info

Get information about a specific transaction in the mempool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMempoolTransactions**
> []interface{} GetMempoolTransactions(ctx, )
Get mempool transactions

Get all transactions currently in the mempool

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

