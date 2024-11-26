# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzePsbt**](BitcoinApi.md#AnalyzePsbt) | **Post** /psbt/analyze | Analyze PSBT
[**CombinePsbt**](BitcoinApi.md#CombinePsbt) | **Post** /psbt/combine | Combine PSBTs
[**CombineRawTransaction**](BitcoinApi.md#CombineRawTransaction) | **Post** /tx/combine | Combine Raw Transactions
[**ConvertToPsbt**](BitcoinApi.md#ConvertToPsbt) | **Post** /tx/convert-to-psbt | Convert Raw Transaction to PSBT
[**CreatePsbt**](BitcoinApi.md#CreatePsbt) | **Post** /psbt/create | Create PSBT
[**CreateRawTransaction**](BitcoinApi.md#CreateRawTransaction) | **Post** /tx/create | Create Raw Transaction
[**DecodePsbt**](BitcoinApi.md#DecodePsbt) | **Post** /psbt/decode | Decode PSBT
[**DecodeScript**](BitcoinApi.md#DecodeScript) | **Post** /script/decode | Decode Script
[**EstimateRawFee**](BitcoinApi.md#EstimateRawFee) | **Post** /fee/estimate-raw | Estimate Raw Fee
[**EstimateSmartFee**](BitcoinApi.md#EstimateSmartFee) | **Post** /fee/estimate-smart | Estimate smart fee
[**GetBlockByHashDecoded**](BitcoinApi.md#GetBlockByHashDecoded) | **Get** /block/hash/{hash}/decoded | Get block by hash (verbosity 2)
[**GetBlockByHashHex**](BitcoinApi.md#GetBlockByHashHex) | **Get** /block/hash/{hash}/hex | Get block by hash (verbosity 0)
[**GetBlockByHashPrevout**](BitcoinApi.md#GetBlockByHashPrevout) | **Get** /block/hash/{hash}/prevout | Get block by hash (verbosity 3)
[**GetBlockByHashSummary**](BitcoinApi.md#GetBlockByHashSummary) | **Get** /block/hash/{hash}/summary | Get block by hash (verbosity 1)
[**GetBlockByHeightDecoded**](BitcoinApi.md#GetBlockByHeightDecoded) | **Get** /block/height/{height}/decoded | Get block by height (verbosity 2)
[**GetBlockByHeightHex**](BitcoinApi.md#GetBlockByHeightHex) | **Get** /block/height/{height}/hex | Get block by height (verbosity 0)
[**GetBlockByHeightPrevout**](BitcoinApi.md#GetBlockByHeightPrevout) | **Get** /block/height/{height}/prevout | Get block by height (verbosity 3)
[**GetBlockByHeightSummary**](BitcoinApi.md#GetBlockByHeightSummary) | **Get** /block/height/{height}/summary | Get block by height (verbosity 1)
[**GetBlockStats**](BitcoinApi.md#GetBlockStats) | **Post** /block/stats | Get block stats
[**GetBlockchainInfo**](BitcoinApi.md#GetBlockchainInfo) | **Get** /blockchain/info | Get blockchain information
[**GetChainTxStats**](BitcoinApi.md#GetChainTxStats) | **Post** /chain/txstats | Get chain tx stats
[**GetDifficulty**](BitcoinApi.md#GetDifficulty) | **Get** /difficulty | Get difficulty
[**GetMempoolAncestors**](BitcoinApi.md#GetMempoolAncestors) | **Post** /mempool/ancestors | Get mempool ancestors
[**GetMempoolDescendants**](BitcoinApi.md#GetMempoolDescendants) | **Post** /mempool/descendants | Get mempool descendants
[**GetMempoolInfo**](BitcoinApi.md#GetMempoolInfo) | **Get** /mempool/info | Get mempool information
[**GetMiningInfo**](BitcoinApi.md#GetMiningInfo) | **Get** /mining/info | Get mining information
[**GetNetworkHashps**](BitcoinApi.md#GetNetworkHashps) | **Post** /network/hashps | Get network hash per second
[**GetRawMempool**](BitcoinApi.md#GetRawMempool) | **Post** /mempool/raw | Get raw mempool
[**GetRawTransactionDecoded**](BitcoinApi.md#GetRawTransactionDecoded) | **Get** /tx/{txid}/decoded | Get raw transaction (verbosity 1)
[**GetRawTransactionHex**](BitcoinApi.md#GetRawTransactionHex) | **Get** /tx/{txid}/hex | Get raw transaction (verbosity 0)
[**GetRawTransactionPrevout**](BitcoinApi.md#GetRawTransactionPrevout) | **Get** /tx/{txid}/prevout | Get raw transaction (verbosity 2)
[**GetTxOut**](BitcoinApi.md#GetTxOut) | **Post** /tx/out | Get transaction output
[**GetTxOutProof**](BitcoinApi.md#GetTxOutProof) | **Post** /tx/out/proof | Get transaction output proof
[**GetTxOutSetInfo**](BitcoinApi.md#GetTxOutSetInfo) | **Post** /tx/out/set/info | Get transaction output set information
[**GetTxSpendingPrevout**](BitcoinApi.md#GetTxSpendingPrevout) | **Post** /tx/spending/prevout | Get transaction spending prevout
[**JoinPsbts**](BitcoinApi.md#JoinPsbts) | **Post** /psbt/join | Join PSBTs
[**SendRawTransaction**](BitcoinApi.md#SendRawTransaction) | **Post** /tx/send | Send raw transaction
[**TestMempoolAccept**](BitcoinApi.md#TestMempoolAccept) | **Post** /mempool/test-accept | Test mempool accept
[**ValidateAddress**](BitcoinApi.md#ValidateAddress) | **Get** /address/{address}/validate | Validate address
[**VerifyMessage**](BitcoinApi.md#VerifyMessage) | **Post** /address/verify-message | Verify message
[**VerifyTxOutProof**](BitcoinApi.md#VerifyTxOutProof) | **Post** /tx/out/proof/verify | Verify transaction output proof

# **AnalyzePsbt**
> ResponsesAnalyzePsbtResponse AnalyzePsbt(ctx, body)
Analyze PSBT

Analyzes and provides information about the current status of a PSBT and its inputs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsAnalyzePsbtRequest**](RequestsAnalyzePsbtRequest.md)| PSBT to analyze | 

### Return type

[**ResponsesAnalyzePsbtResponse**](responses.AnalyzePSBTResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CombinePsbt**
> ResponsesCombinePsbtResponse CombinePsbt(ctx, body)
Combine PSBTs

Combines multiple partially signed Bitcoin transactions into one transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsCombinePsbtRequest**](RequestsCombinePsbtRequest.md)| Array of PSBTs to combine | 

### Return type

[**ResponsesCombinePsbtResponse**](responses.CombinePSBTResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CombineRawTransaction**
> ResponsesCombineRawTransactionResponse CombineRawTransaction(ctx, body)
Combine Raw Transactions

Combines multiple partially signed transactions into one transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsCombineRawTransactionRequest**](RequestsCombineRawTransactionRequest.md)| Array of hex-encoded raw transactions | 

### Return type

[**ResponsesCombineRawTransactionResponse**](responses.CombineRawTransactionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConvertToPsbt**
> ResponsesConvertToPsbtResponse ConvertToPsbt(ctx, body)
Convert Raw Transaction to PSBT

Converts a network serialized transaction to a PSBT. This should be used only with createrawtransaction and fundrawtransaction. createpsbt and walletcreatefundedpsbt should be used for new applications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsConvertToPsbtRequest**](RequestsConvertToPsbtRequest.md)| Raw transaction conversion parameters | 

### Return type

[**ResponsesConvertToPsbtResponse**](responses.ConvertToPSBTResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePsbt**
> ResponsesCreatePsbtResponse CreatePsbt(ctx, body)
Create PSBT

Creates a transaction in the Partially Signed Transaction format. Implements the Creator role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GithubComSatstreamSsApiServerApiPsbtRequestsCreatePsbtRequest**](GithubComSatstreamSsApiServerApiPsbtRequestsCreatePsbtRequest.md)| Transaction parameters | 

### Return type

[**ResponsesCreatePsbtResponse**](responses.CreatePSBTResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRawTransaction**
> ResponsesCreateRawTransactionResponse CreateRawTransaction(ctx, body)
Create Raw Transaction

Creates a raw transaction spending the given inputs and creating new outputs. Note that the transaction's inputs are not signed, and it is not stored in the wallet or transmitted to the network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GithubComSatstreamSsApiServerApiTransactionRequestsCreatePsbtRequest**](GithubComSatstreamSsApiServerApiTransactionRequestsCreatePsbtRequest.md)| Transaction parameters | 

### Return type

[**ResponsesCreateRawTransactionResponse**](responses.CreateRawTransactionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DecodePsbt**
> ResponsesDecodePsbtResponse DecodePsbt(ctx, body)
Decode PSBT

Return a JSON object representing the serialized, base64-encoded partially signed Bitcoin transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsAnalyzePsbtRequest**](RequestsAnalyzePsbtRequest.md)| PSBT to decode | 

### Return type

[**ResponsesDecodePsbtResponse**](responses.DecodePSBTResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DecodeScript**
> ResponsesDecodeScriptResponse DecodeScript(ctx, body)
Decode Script

Decode a hex-encoded script and return detailed information about it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsDecodeScriptRequest**](RequestsDecodeScriptRequest.md)| Script to decode | 

### Return type

[**ResponsesDecodeScriptResponse**](responses.DecodeScriptResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EstimateRawFee**
> ResponsesEstimateRawFeeResponse EstimateRawFee(ctx, body)
Estimate Raw Fee

Estimates the approximate fee per kilobyte needed for a transaction to begin confirmation within conf_target blocks if possible.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsEstimateRawFeeRequest**](RequestsEstimateRawFeeRequest.md)| Fee estimation parameters | 

### Return type

[**ResponsesEstimateRawFeeResponse**](responses.EstimateRawFeeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EstimateSmartFee**
> ResponsesEstimateSmartFeeResponse EstimateSmartFee(ctx, body)
Estimate smart fee

Estimates the approximate fee per kilobyte needed for a transaction to begin confirmation within conf_target blocks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsEstimateSmartFeeRequest**](RequestsEstimateSmartFeeRequest.md)| Fee estimation parameters | 

### Return type

[**ResponsesEstimateSmartFeeResponse**](responses.EstimateSmartFeeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockByHashDecoded**
> ResponsesGetBlockDecodedResponse GetBlockByHashDecoded(ctx, hash)
Get block by hash (verbosity 2)

Get block by hash as a decoded object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| Block hash | 

### Return type

[**ResponsesGetBlockDecodedResponse**](responses.GetBlockDecodedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockByHashHex**
> ResponsesGetBlockHexResponse GetBlockByHashHex(ctx, hash)
Get block by hash (verbosity 0)

Get block by hash as a raw hex string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| Block hash | 

### Return type

[**ResponsesGetBlockHexResponse**](responses.GetBlockHexResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockByHashPrevout**
> ResponsesGetBlockPrevoutResponse GetBlockByHashPrevout(ctx, hash)
Get block by hash (verbosity 3)

Get block by hash with prevout information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| Block hash | 

### Return type

[**ResponsesGetBlockPrevoutResponse**](responses.GetBlockPrevoutResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockByHashSummary**
> ResponsesGetBlockSummaryResponse GetBlockByHashSummary(ctx, hash)
Get block by hash (verbosity 1)

Get block by hash as a summary object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| Block hash | 

### Return type

[**ResponsesGetBlockSummaryResponse**](responses.GetBlockSummaryResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockByHeightDecoded**
> ResponsesGetBlockDecodedResponse GetBlockByHeightDecoded(ctx, height)
Get block by height (verbosity 2)

Get block by height as a decoded object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **height** | **int32**| Block height | 

### Return type

[**ResponsesGetBlockDecodedResponse**](responses.GetBlockDecodedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockByHeightHex**
> ResponsesGetBlockHexResponse GetBlockByHeightHex(ctx, height)
Get block by height (verbosity 0)

Get block by height as a raw hex string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **height** | **int32**| Block height | 

### Return type

[**ResponsesGetBlockHexResponse**](responses.GetBlockHexResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockByHeightPrevout**
> ResponsesGetBlockPrevoutResponse GetBlockByHeightPrevout(ctx, height)
Get block by height (verbosity 3)

Get block by height with prevout information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **height** | **int32**| Block height | 

### Return type

[**ResponsesGetBlockPrevoutResponse**](responses.GetBlockPrevoutResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockByHeightSummary**
> ResponsesGetBlockSummaryResponse GetBlockByHeightSummary(ctx, height)
Get block by height (verbosity 1)

Get block by height as a summary object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **height** | **int32**| Block height | 

### Return type

[**ResponsesGetBlockSummaryResponse**](responses.GetBlockSummaryResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockStats**
> ResponsesGetBlockStatsResponse GetBlockStats(ctx, body)
Get block stats

Computes per block statistics for a given window

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetBlockStatsRequest**](RequestsGetBlockStatsRequest.md)| Block stats request parameters | 

### Return type

[**ResponsesGetBlockStatsResponse**](responses.GetBlockStatsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockchainInfo**
> ResponsesGetBlockchainInfoResponse GetBlockchainInfo(ctx, )
Get blockchain information

Returns an object containing various state info regarding blockchain processing

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponsesGetBlockchainInfoResponse**](responses.GetBlockchainInfoResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChainTxStats**
> ResponsesGetChainTxStatsResponse GetChainTxStats(ctx, body)
Get chain tx stats

Computes statistics about the total number and rate of transactions in the chain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetChainTxStatsRequest**](RequestsGetChainTxStatsRequest.md)| Chain tx stats request parameters | 

### Return type

[**ResponsesGetChainTxStatsResponse**](responses.GetChainTxStatsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDifficulty**
> ResponsesGetDifficultyResponse GetDifficulty(ctx, )
Get difficulty

Returns the proof-of-work difficulty as a multiple of the minimum difficulty

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponsesGetDifficultyResponse**](responses.GetDifficultyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMempoolAncestors**
> ResponsesGetMempoolAncestorsResponse GetMempoolAncestors(ctx, body)
Get mempool ancestors

Returns all in-mempool ancestors for a transaction in the mempool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetMempoolAncestorsRequest**](RequestsGetMempoolAncestorsRequest.md)| Mempool ancestors request parameters | 

### Return type

[**ResponsesGetMempoolAncestorsResponse**](responses.GetMempoolAncestorsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMempoolDescendants**
> ResponsesGetMempoolDescendantsResponse GetMempoolDescendants(ctx, body)
Get mempool descendants

Returns all in-mempool descendants for a transaction in the mempool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetMempoolDescendantsRequest**](RequestsGetMempoolDescendantsRequest.md)| Mempool descendants request parameters | 

### Return type

[**ResponsesGetMempoolDescendantsResponse**](responses.GetMempoolDescendantsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMempoolInfo**
> ResponsesGetMempoolInfoResponse GetMempoolInfo(ctx, )
Get mempool information

Returns details on the active state of the TX memory pool

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponsesGetMempoolInfoResponse**](responses.GetMempoolInfoResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMiningInfo**
> ResponsesGetMiningInfoResponse GetMiningInfo(ctx, )
Get mining information

Returns a json object containing mining-related information

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponsesGetMiningInfoResponse**](responses.GetMiningInfoResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkHashps**
> ResponsesGetNetworkHashPsResponse GetNetworkHashps(ctx, body)
Get network hash per second

Returns the estimated network hashes per second based on the last n blocks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetNetworkHashPsRequest**](RequestsGetNetworkHashPsRequest.md)| Network hash rate parameters | 

### Return type

[**ResponsesGetNetworkHashPsResponse**](responses.GetNetworkHashPSResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawMempool**
> ResponsesGetRawMempoolResponse GetRawMempool(ctx, body)
Get raw mempool

Returns all transaction ids in memory pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetRawMempoolRequest**](RequestsGetRawMempoolRequest.md)| Raw mempool request parameters | 

### Return type

[**ResponsesGetRawMempoolResponse**](responses.GetRawMempoolResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawTransactionDecoded**
> ResponsesGetRawTransactionDecodedResponse GetRawTransactionDecoded(ctx, txid)
Get raw transaction (verbosity 1)

Get raw transaction as a decoded object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**ResponsesGetRawTransactionDecodedResponse**](responses.GetRawTransactionDecodedResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawTransactionHex**
> ResponsesGetRawTransactionHexResponse GetRawTransactionHex(ctx, txid)
Get raw transaction (verbosity 0)

Get raw transaction as a raw hex string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**ResponsesGetRawTransactionHexResponse**](responses.GetRawTransactionHexResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawTransactionPrevout**
> ResponsesGetRawTransactionPrevoutResponse GetRawTransactionPrevout(ctx, txid)
Get raw transaction (verbosity 2)

Get raw transaction with prevout information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**ResponsesGetRawTransactionPrevoutResponse**](responses.GetRawTransactionPrevoutResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTxOut**
> ResponsesGetTxOutResponse GetTxOut(ctx, body)
Get transaction output

Returns details about an unspent transaction output

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetTxOutRequest**](RequestsGetTxOutRequest.md)| Transaction output request parameters | 

### Return type

[**ResponsesGetTxOutResponse**](responses.GetTxOutResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTxOutProof**
> ResponsesGetTxOutProofResponse GetTxOutProof(ctx, body)
Get transaction output proof

Returns a hex-encoded proof that one or more specified transactions were included in a block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetTxOutProofRequest**](RequestsGetTxOutProofRequest.md)| Transaction proof request parameters | 

### Return type

[**ResponsesGetTxOutProofResponse**](responses.GetTxOutProofResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTxOutSetInfo**
> ResponsesGetTxOutSetInfoResponse GetTxOutSetInfo(ctx, body)
Get transaction output set information

Returns statistics about the unspent transaction output set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetTxOutSetInfoRequest**](RequestsGetTxOutSetInfoRequest.md)| UTXO set info request parameters | 

### Return type

[**ResponsesGetTxOutSetInfoResponse**](responses.GetTxOutSetInfoResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTxSpendingPrevout**
> ResponsesGetTxSpendingPrevoutResponse GetTxSpendingPrevout(ctx, body)
Get transaction spending prevout

Scans the mempool to find transactions spending any of the given outputs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetTxSpendingPrevoutRequest**](RequestsGetTxSpendingPrevoutRequest.md)| Transaction spending prevout request | 

### Return type

[**ResponsesGetTxSpendingPrevoutResponse**](responses.GetTxSpendingPrevoutResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JoinPsbts**
> ResponsesJoinPsbtsResponse JoinPsbts(ctx, body)
Join PSBTs

Joins multiple distinct PSBTs with different inputs and outputs into one PSBT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsJoinPsbtsRequest**](RequestsJoinPsbtsRequest.md)| PSBTs to join | 

### Return type

[**ResponsesJoinPsbtsResponse**](responses.JoinPSBTsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendRawTransaction**
> ResponsesSendRawTransactionResponse SendRawTransaction(ctx, body)
Send raw transaction

Submits a raw transaction to local node and network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsSendRawTransactionRequest**](RequestsSendRawTransactionRequest.md)| Raw transaction to send | 

### Return type

[**ResponsesSendRawTransactionResponse**](responses.SendRawTransactionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestMempoolAccept**
> ResponsesTestMempoolAcceptResponse TestMempoolAccept(ctx, body)
Test mempool accept

Tests whether raw transactions would be accepted by mempool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsTestMempoolAcceptRequest**](RequestsTestMempoolAcceptRequest.md)| Raw transactions to test | 

### Return type

[**ResponsesTestMempoolAcceptResponse**](responses.TestMempoolAcceptResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateAddress**
> ResponsesValidateAddressResponse ValidateAddress(ctx, address)
Validate address

Returns information about the given Bitcoin address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Bitcoin address to validate | 

### Return type

[**ResponsesValidateAddressResponse**](responses.ValidateAddressResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyMessage**
> ResponsesVerifyMessageResponse VerifyMessage(ctx, body)
Verify message

Verifies a signed message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsVerifyMessageRequest**](RequestsVerifyMessageRequest.md)| Message verification parameters | 

### Return type

[**ResponsesVerifyMessageResponse**](responses.VerifyMessageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyTxOutProof**
> ResponsesVerifyTxOutProofResponse VerifyTxOutProof(ctx, body)
Verify transaction output proof

Verifies that a proof points to a transaction in a block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsVerifyTxOutProofRequest**](RequestsVerifyTxOutProofRequest.md)| Proof to verify | 

### Return type

[**ResponsesVerifyTxOutProofResponse**](responses.VerifyTxOutProofResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

