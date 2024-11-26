# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DecodeTx**](InscriptionsApi.md#DecodeTx) | **Get** /decode/{txid} | Decode a transaction
[**FetchInscriptionChild**](InscriptionsApi.md#FetchInscriptionChild) | **Get** /inscription/{inscription_id}/child/{child_index} | Get inscription child info
[**FetchInscriptions**](InscriptionsApi.md#FetchInscriptions) | **Post** /inscriptions | Fetch multiple inscriptions
[**GetAddress**](InscriptionsApi.md#GetAddress) | **Get** /address/{address} | Get address info
[**GetAddressUtxos**](InscriptionsApi.md#GetAddressUtxos) | **Get** /address/{address}/outputs | Get UTXOs for an address
[**GetBlockByHeight**](InscriptionsApi.md#GetBlockByHeight) | **Get** /block/height/{block_height} | Get block info by height
[**GetBlockCount**](InscriptionsApi.md#GetBlockCount) | **Get** /blockcount | Get the height of the latest block
[**GetBlockHashByHeight**](InscriptionsApi.md#GetBlockHashByHeight) | **Get** /blockhash/{block_height} | Returns blockhash of specified block.
[**GetBlockInscriptions**](InscriptionsApi.md#GetBlockInscriptions) | **Get** /inscriptions/block/{block_height} | Get inscriptions in a specific block
[**GetBlocks**](InscriptionsApi.md#GetBlocks) | **Get** /blocks | Returns the latest block height, last 100 block hashes, and featured inscriptions
[**GetInscription**](InscriptionsApi.md#GetInscription) | **Get** /inscription/{inscription_id} | Get inscription info
[**GetLatestBlockHeight**](InscriptionsApi.md#GetLatestBlockHeight) | **Get** /latestblockheight | Returns the height of the latest block.
[**GetLatestBlockhash**](InscriptionsApi.md#GetLatestBlockhash) | **Get** /latestblockhash | Returns blockhash for the latest block.
[**GetLatestBlocktime**](InscriptionsApi.md#GetLatestBlocktime) | **Get** /blocktime | Get the timestamp of the latest block
[**GetLatestInscriptions**](InscriptionsApi.md#GetLatestInscriptions) | **Get** /inscriptions/latest | Get latest inscriptions
[**GetLatestInscriptionsPage**](InscriptionsApi.md#GetLatestInscriptionsPage) | **Get** /inscriptions/page/{page} | Get latest inscriptions page
[**GetLatestRunes**](InscriptionsApi.md#GetLatestRunes) | **Get** /runes/latest | Get latest runes
[**GetLatestRunesPage**](InscriptionsApi.md#GetLatestRunesPage) | **Get** /runes/page/{page} | Get latest runes page
[**GetOutputByOutpoint**](InscriptionsApi.md#GetOutputByOutpoint) | **Get** /output/outpoint/{outpoint} | Get output info by outpoint
[**GetOutputs**](InscriptionsApi.md#GetOutputs) | **Post** /outputs | Get multiple outputs
[**GetRune**](InscriptionsApi.md#GetRune) | **Get** /rune/{rune_name} | Get rune info
[**GetSatoshi**](InscriptionsApi.md#GetSatoshi) | **Get** /sat/{number} | Get satoshi info
[**GetStatus**](InscriptionsApi.md#GetStatus) | **Get** /status | Get server status
[**GetTransaction**](InscriptionsApi.md#GetTransaction) | **Get** /tx/{txid} | Get transaction info

# **DecodeTx**
> GithubComSatstreamSsApiServerApiTransactionResponsesDecodeResponse DecodeTx(ctx, txid)
Decode a transaction

Decodes a transaction and returns its inscriptions and runestone data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**GithubComSatstreamSsApiServerApiTransactionResponsesDecodeResponse**](github_com_satstream_ss-api_server_api_transaction_responses.DecodeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchInscriptionChild**
> GithubComSatstreamSsApiServerApiInscriptionResponsesInscriptionResponse FetchInscriptionChild(ctx, inscriptionId, childIndex)
Get inscription child info

Retrieve information about a specific child of an inscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inscriptionId** | **string**| Inscription ID | 
  **childIndex** | **int32**| Child Index | 

### Return type

[**GithubComSatstreamSsApiServerApiInscriptionResponsesInscriptionResponse**](github_com_satstream_ss-api_server_api_inscription_responses.InscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchInscriptions**
> []GithubComSatstreamSsApiServerApiInscriptionResponsesInscriptionResponse FetchInscriptions(ctx, body)
Fetch multiple inscriptions

Retrieve information about multiple inscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)| Inscription IDs | 

### Return type

[**[]GithubComSatstreamSsApiServerApiInscriptionResponsesInscriptionResponse**](github_com_satstream_ss-api_server_api_inscription_responses.InscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddress**
> GithubComSatstreamSsApiServerApiAddressResponsesAddressResponse GetAddress(ctx, address)
Get address info

Get detailed information about a specific address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address | 

### Return type

[**GithubComSatstreamSsApiServerApiAddressResponsesAddressResponse**](github_com_satstream_ss-api_server_api_address_responses.AddressResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddressUtxos**
> []GithubComSatstreamSsApiServerApiAddressResponsesOutputResponse GetAddressUtxos(ctx, address, optional)
Get UTXOs for an address

Retrieve UTXOs held by a specific address with optional type filtering

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| Address | 
 **optional** | ***InscriptionsApiGetAddressUtxosOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InscriptionsApiGetAddressUtxosOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| UTXO Type | 

### Return type

[**[]GithubComSatstreamSsApiServerApiAddressResponsesOutputResponse**](github_com_satstream_ss-api_server_api_address_responses.OutputResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockByHeight**
> GithubComSatstreamSsApiServerApiBlockResponsesBlockResponse GetBlockByHeight(ctx, blockHeight)
Get block info by height

Get detailed information about a specific block by height

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockHeight** | **string**| Block Height | 

### Return type

[**GithubComSatstreamSsApiServerApiBlockResponsesBlockResponse**](github_com_satstream_ss-api_server_api_block_responses.BlockResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockCount**
> GithubComSatstreamSsApiServerApiBlockResponsesBlockCountResponse GetBlockCount(ctx, )
Get the height of the latest block

Returns the height of the latest block

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GithubComSatstreamSsApiServerApiBlockResponsesBlockCountResponse**](github_com_satstream_ss-api_server_api_block_responses.BlockCountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockHashByHeight**
> GithubComSatstreamSsApiServerApiBlockResponsesBlockHashResponse GetBlockHashByHeight(ctx, blockHeight)
Returns blockhash of specified block.

Returns blockhash of specified block.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockHeight** | **string**| Block Height | 

### Return type

[**GithubComSatstreamSsApiServerApiBlockResponsesBlockHashResponse**](github_com_satstream_ss-api_server_api_block_responses.BlockHashResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockInscriptions**
> GithubComSatstreamSsApiServerApiInscriptionResponsesLatestInscriptionsResponse GetBlockInscriptions(ctx, blockHeight)
Get inscriptions in a specific block

Retrieve all inscriptions in a specific block

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockHeight** | **int32**| Block Height | 

### Return type

[**GithubComSatstreamSsApiServerApiInscriptionResponsesLatestInscriptionsResponse**](github_com_satstream_ss-api_server_api_inscription_responses.LatestInscriptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlocks**
> GithubComSatstreamSsApiServerApiBlockResponsesBlocksResponse GetBlocks(ctx, )
Returns the latest block height, last 100 block hashes, and featured inscriptions

Returns the latest block height, last 100 block hashes, and featured inscriptions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GithubComSatstreamSsApiServerApiBlockResponsesBlocksResponse**](github_com_satstream_ss-api_server_api_block_responses.BlocksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInscription**
> GithubComSatstreamSsApiServerApiInscriptionResponsesInscriptionResponse GetInscription(ctx, inscriptionId)
Get inscription info

Get information about a specific inscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inscriptionId** | **string**| Inscription ID | 

### Return type

[**GithubComSatstreamSsApiServerApiInscriptionResponsesInscriptionResponse**](github_com_satstream_ss-api_server_api_inscription_responses.InscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestBlockHeight**
> ResponsesLatestBlockHeightResponse GetLatestBlockHeight(ctx, )
Returns the height of the latest block.

Returns the height of the latest block.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponsesLatestBlockHeightResponse**](responses.LatestBlockHeightResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestBlockhash**
> ResponsesLatestBlockHashResponse GetLatestBlockhash(ctx, )
Returns blockhash for the latest block.

Returns blockhash for the latest block.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponsesLatestBlockHashResponse**](responses.LatestBlockHashResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestBlocktime**
> ResponsesLatestBlockTimeResponse GetLatestBlocktime(ctx, )
Get the timestamp of the latest block

Returns the UNIX timestamp of when the latest block was mined

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ResponsesLatestBlockTimeResponse**](responses.LatestBlockTimeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestInscriptions**
> GithubComSatstreamSsApiServerApiInscriptionResponsesLatestInscriptionsResponse GetLatestInscriptions(ctx, )
Get latest inscriptions

Retrieve the latest 100 inscriptions (first page)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GithubComSatstreamSsApiServerApiInscriptionResponsesLatestInscriptionsResponse**](github_com_satstream_ss-api_server_api_inscription_responses.LatestInscriptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestInscriptionsPage**
> GithubComSatstreamSsApiServerApiInscriptionResponsesLatestInscriptionsResponse GetLatestInscriptionsPage(ctx, page)
Get latest inscriptions page

Retrieve a specific page of 100 inscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page number | 

### Return type

[**GithubComSatstreamSsApiServerApiInscriptionResponsesLatestInscriptionsResponse**](github_com_satstream_ss-api_server_api_inscription_responses.LatestInscriptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestRunes**
> GithubComSatstreamSsApiServerApiRuneResponsesRunesListResponse GetLatestRunes(ctx, )
Get latest runes

Retrieve information about the last 100 inscribed runes (first page)

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GithubComSatstreamSsApiServerApiRuneResponsesRunesListResponse**](github_com_satstream_ss-api_server_api_rune_responses.RunesListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestRunesPage**
> GithubComSatstreamSsApiServerApiRuneResponsesRunesListResponse GetLatestRunesPage(ctx, page)
Get latest runes page

Retrieve a specific page of 100 inscribed runes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page number | 

### Return type

[**GithubComSatstreamSsApiServerApiRuneResponsesRunesListResponse**](github_com_satstream_ss-api_server_api_rune_responses.RunesListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOutputByOutpoint**
> ResponsesGetOutputByOutpointResponse GetOutputByOutpoint(ctx, outpoint)
Get output info by outpoint

Retrieve information about a specific UTXO using outpoint string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **outpoint** | **string**| Outpoint | 

### Return type

[**ResponsesGetOutputByOutpointResponse**](responses.GetOutputByOutpointResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOutputs**
> []ResponsesGetOutputsResponse GetOutputs(ctx, body)
Get multiple outputs

Retrieve information about multiple UTXOs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)| Outpoints | 

### Return type

[**[]ResponsesGetOutputsResponse**](responses.GetOutputsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRune**
> GithubComSatstreamSsApiServerApiRuneResponsesRuneResponse GetRune(ctx, runeName)
Get rune info

Retrieve information about a specific rune

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **runeName** | **string**| Rune Name | 

### Return type

[**GithubComSatstreamSsApiServerApiRuneResponsesRuneResponse**](github_com_satstream_ss-api_server_api_rune_responses.RuneResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSatoshi**
> GithubComSatstreamSsApiServerApiSatoshiResponsesSatoshiResponse GetSatoshi(ctx, number)
Get satoshi info

Retrieve information about a specific satoshi

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **number** | **int32**| Satoshi Number | 

### Return type

[**GithubComSatstreamSsApiServerApiSatoshiResponsesSatoshiResponse**](github_com_satstream_ss-api_server_api_satoshi_responses.SatoshiResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatus**
> GithubComSatstreamSsApiServerApiStatusResponsesStatusResponse GetStatus(ctx, )
Get server status

Retrieve information about the server installation and index

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GithubComSatstreamSsApiServerApiStatusResponsesStatusResponse**](github_com_satstream_ss-api_server_api_status_responses.StatusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransaction**
> GithubComSatstreamSsApiServerApiTransactionResponsesTransactionResponse GetTransaction(ctx, txid)
Get transaction info

Retrieve information about a specific transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **txid** | **string**| Transaction ID | 

### Return type

[**GithubComSatstreamSsApiServerApiTransactionResponsesTransactionResponse**](github_com_satstream_ss-api_server_api_transaction_responses.TransactionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

