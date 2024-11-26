# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CombineRawTransaction**](TransactionsApi.md#CombineRawTransaction) | **Post** /tx/combine | Combine Raw Transactions
[**ConvertToPsbt**](TransactionsApi.md#ConvertToPsbt) | **Post** /tx/convert-to-psbt | Convert Raw Transaction to PSBT
[**CreateRawTransaction**](TransactionsApi.md#CreateRawTransaction) | **Post** /tx/create | Create Raw Transaction
[**DecodeTx**](TransactionsApi.md#DecodeTx) | **Get** /tx/{txid}/decode | Decode a transaction
[**GetRawTransactionDecoded**](TransactionsApi.md#GetRawTransactionDecoded) | **Get** /tx/{txid}/decoded | Get raw transaction (verbosity 1)
[**GetRawTransactionHex**](TransactionsApi.md#GetRawTransactionHex) | **Get** /tx/{txid}/hex | Get raw transaction (verbosity 0)
[**GetRawTransactionPrevout**](TransactionsApi.md#GetRawTransactionPrevout) | **Get** /tx/{txid}/prevout | Get raw transaction (verbosity 2)
[**GetTransaction**](TransactionsApi.md#GetTransaction) | **Get** /tx/{txid} | Get transaction info
[**GetTxOut**](TransactionsApi.md#GetTxOut) | **Post** /tx/out | Get transaction output
[**GetTxOutProof**](TransactionsApi.md#GetTxOutProof) | **Post** /tx/outproof | Get transaction output proof
[**GetTxOutSetInfo**](TransactionsApi.md#GetTxOutSetInfo) | **Post** /tx/out/set/info | Get transaction output set information
[**GetTxSpendingPrevout**](TransactionsApi.md#GetTxSpendingPrevout) | **Post** /tx/spending-prevout | Get transaction spending prevout
[**SendRawTransaction**](TransactionsApi.md#SendRawTransaction) | **Post** /tx/send | Send raw transaction
[**VerifyTxOutProof**](TransactionsApi.md#VerifyTxOutProof) | **Post** /tx/outproof/verify | Verify transaction output proof

# **CombineRawTransaction**
> InlineResponse2005 CombineRawTransaction(ctx, body)
Combine Raw Transactions

Combines multiple partially signed transactions into one transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsCombineRawTransactionRequest**](RequestsCombineRawTransactionRequest.md)| Array of hex-encoded raw transactions | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertToPsbt**
> InlineResponse2005 ConvertToPsbt(ctx, body)
Convert Raw Transaction to PSBT

Converts a network serialized transaction to a PSBT. This should be used only with createrawtransaction and fundrawtransaction. createpsbt and walletcreatefundedpsbt should be used for new applications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsConvertToPsbtRequest**](RequestsConvertToPsbtRequest.md)| Raw transaction conversion parameters | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRawTransaction**
> InlineResponse2005 CreateRawTransaction(ctx, body)
Create Raw Transaction

Creates a raw transaction spending the given inputs and creating new outputs. Note that the transaction's inputs are not signed, and it is not stored in the wallet or transmitted to the network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsCreateRawTxRequest**](RequestsCreateRawTxRequest.md)| Transaction parameters | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DecodeTx**
> InlineResponse20038 DecodeTx(ctx, txid)
Decode a transaction

Decodes a transaction and returns its inscriptions and runestone data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawTransactionDecoded**
> InlineResponse20039 GetRawTransactionDecoded(ctx, txid)
Get raw transaction (verbosity 1)

Get raw transaction as a decoded object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawTransactionHex**
> InlineResponse2005 GetRawTransactionHex(ctx, txid)
Get raw transaction (verbosity 0)

Get raw transaction as a raw hex string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawTransactionPrevout**
> InlineResponse20040 GetRawTransactionPrevout(ctx, txid)
Get raw transaction (verbosity 2)

Get raw transaction with prevout information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransaction**
> InlineResponse20037 GetTransaction(ctx, txid)
Get transaction info

Retrieve information about a specific transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTxOut**
> InlineResponse20033 GetTxOut(ctx, body)
Get transaction output

Returns details about an unspent transaction output

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetTxOutRequest**](RequestsGetTxOutRequest.md)| Transaction output request parameters | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTxOutProof**
> InlineResponse2005 GetTxOutProof(ctx, body)
Get transaction output proof

Returns a hex-encoded proof that one or more specified transactions were included in a block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetTxOutProofRequest**](RequestsGetTxOutProofRequest.md)| Transaction proof request parameters | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTxOutSetInfo**
> InlineResponse20034 GetTxOutSetInfo(ctx, body)
Get transaction output set information

Returns statistics about the unspent transaction output set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetTxOutSetInfoRequest**](RequestsGetTxOutSetInfoRequest.md)| UTXO set info request parameters | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTxSpendingPrevout**
> InlineResponse20036 GetTxSpendingPrevout(ctx, body)
Get transaction spending prevout

Scans the mempool to find transactions spending any of the given outputs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetTxSpendingPrevoutRequest**](RequestsGetTxSpendingPrevoutRequest.md)| Transaction spending prevout request | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendRawTransaction**
> InlineResponse2005 SendRawTransaction(ctx, body)
Send raw transaction

Submits a raw transaction to local node and network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsSendRawTransactionRequest**](RequestsSendRawTransactionRequest.md)| Raw transaction to send | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyTxOutProof**
> InlineResponse20035 VerifyTxOutProof(ctx, body)
Verify transaction output proof

Verifies that a proof points to a transaction in a block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsVerifyTxOutProofRequest**](RequestsVerifyTxOutProofRequest.md)| Proof to verify | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

