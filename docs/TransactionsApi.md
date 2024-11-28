# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CombineRawTransaction**](TransactionsApi.md#CombineRawTransaction) | **Post** /tx/combine | Combine Raw Transactions
[**ConvertToPsbt**](TransactionsApi.md#ConvertToPsbt) | **Post** /tx/convert-to-psbt | Convert Raw Transaction to PSBT
[**CreateRawTransaction**](TransactionsApi.md#CreateRawTransaction) | **Post** /tx/create | Create Raw Transaction
[**DecodeTxInscriptions**](TransactionsApi.md#DecodeTxInscriptions) | **Get** /tx/{txid}/inscriptions | Decode transaction inscriptions
[**GetRawTransaction**](TransactionsApi.md#GetRawTransaction) | **Get** /tx/{txid}/raw/decode | Get raw transaction (verbosity 1)
[**GetRawTransactionHex**](TransactionsApi.md#GetRawTransactionHex) | **Get** /tx/{txid}/hex | Get raw transaction (verbosity 0)
[**GetRawTransactionPrevout**](TransactionsApi.md#GetRawTransactionPrevout) | **Get** /tx/{txid}/raw/prevout | Get raw transaction with prevouts (verbosity 2)
[**GetTxOut**](TransactionsApi.md#GetTxOut) | **Post** /tx/out | Get transaction output
[**GetTxOutProof**](TransactionsApi.md#GetTxOutProof) | **Post** /tx/outproof | Get transaction output proof
[**GetTxOutSetInfo**](TransactionsApi.md#GetTxOutSetInfo) | **Post** /tx/out/set/info | Get transaction output set information
[**GetTxSpendingPrevout**](TransactionsApi.md#GetTxSpendingPrevout) | **Post** /tx/spending-prevout | Get transaction spending prevout
[**SendRawTransaction**](TransactionsApi.md#SendRawTransaction) | **Post** /tx/send | Send raw transaction
[**VerifyTxOutProof**](TransactionsApi.md#VerifyTxOutProof) | **Post** /tx/outproof/verify | Verify transaction output proof

# **CombineRawTransaction**
> CombineRawTransactionResponse CombineRawTransaction(ctx, body)
Combine Raw Transactions

Combines multiple partially signed transactions into one transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionCombineRawTransactionRequest**](TransactionCombineRawTransactionRequest.md)| Array of hex-encoded raw transactions | 

### Return type

[**CombineRawTransactionResponse**](CombineRawTransactionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertToPsbt**
> ConvertToPsbtResponse ConvertToPsbt(ctx, body)
Convert Raw Transaction to PSBT

Converts a network serialized transaction to a PSBT. This should be used only with createrawtransaction and fundrawtransaction. createpsbt and walletcreatefundedpsbt should be used for new applications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionConvertToPsbtRequest**](TransactionConvertToPsbtRequest.md)| Raw transaction conversion parameters | 

### Return type

[**ConvertToPsbtResponse**](ConvertToPSBTResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRawTransaction**
> CreateRawTransactionResponse CreateRawTransaction(ctx, body)
Create Raw Transaction

Creates a raw transaction spending the given inputs and creating new outputs. Note that the transaction's inputs are not signed, and it is not stored in the wallet or transmitted to the network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionCreateRawTxRequest**](TransactionCreateRawTxRequest.md)| Transaction parameters | 

### Return type

[**CreateRawTransactionResponse**](CreateRawTransactionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DecodeTxInscriptions**
> DecodeTransactionResponse DecodeTxInscriptions(ctx, txid)
Decode transaction inscriptions

Decodes a transaction and returns its inscriptions and runestone data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**DecodeTransactionResponse**](DecodeTransactionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawTransaction**
> GetRawTransactionDecodeResponse GetRawTransaction(ctx, txid)
Get raw transaction (verbosity 1)

Get raw transaction with basic decoded information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**GetRawTransactionDecodeResponse**](GetRawTransactionDecodeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawTransactionHex**
> GetRawTransactionHexResponse GetRawTransactionHex(ctx, txid)
Get raw transaction (verbosity 0)

Get raw transaction as a raw hex string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**GetRawTransactionHexResponse**](GetRawTransactionHexResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawTransactionPrevout**
> GetRawTransactionPrevoutResponse GetRawTransactionPrevout(ctx, txid)
Get raw transaction with prevouts (verbosity 2)

Get raw transaction with prevout information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**GetRawTransactionPrevoutResponse**](GetRawTransactionPrevoutResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTxOut**
> GetTxOutResponse GetTxOut(ctx, body)
Get transaction output

Returns details about an unspent transaction output

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionGetTxOutRequest**](TransactionGetTxOutRequest.md)| Transaction output request parameters | 

### Return type

[**GetTxOutResponse**](GetTxOutResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTxOutProof**
> GetTxOutProofResponse GetTxOutProof(ctx, body)
Get transaction output proof

Returns a hex-encoded proof that one or more specified transactions were included in a block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionGetTxOutProofRequest**](TransactionGetTxOutProofRequest.md)| Transaction proof request parameters | 

### Return type

[**GetTxOutProofResponse**](GetTxOutProofResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTxOutSetInfo**
> InlineResponse2002 GetTxOutSetInfo(ctx, body)
Get transaction output set information

Returns statistics about the unspent transaction output set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionGetTxOutSetInfoRequest**](TransactionGetTxOutSetInfoRequest.md)| UTXO set info request parameters | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTxSpendingPrevout**
> GetTxSpendingPrevoutResponse GetTxSpendingPrevout(ctx, body)
Get transaction spending prevout

Scans the mempool to find transactions spending any of the given outputs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionGetTxSpendingPrevoutRequest**](TransactionGetTxSpendingPrevoutRequest.md)| Transaction spending prevout request | 

### Return type

[**GetTxSpendingPrevoutResponse**](GetTxSpendingPrevoutResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendRawTransaction**
> SendRawTransactionResponse SendRawTransaction(ctx, body)
Send raw transaction

Submits a raw transaction to local node and network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionSendRawTransactionRequest**](TransactionSendRawTransactionRequest.md)| Raw transaction to send | 

### Return type

[**SendRawTransactionResponse**](SendRawTransactionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyTxOutProof**
> VerifyTxOutProofResponse VerifyTxOutProof(ctx, body)
Verify transaction output proof

Verifies that a proof points to a transaction in a block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransactionVerifyTxOutProofRequest**](TransactionVerifyTxOutProofRequest.md)| Proof to verify | 

### Return type

[**VerifyTxOutProofResponse**](VerifyTxOutProofResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

