# Go API client for satstream_go_sdk

Satstream API

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.21
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://satstream.io](https://satstream.io)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./satstream_go_sdk"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.satstream.io/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AddressesApi* | [**GetAddress**](docs/AddressesApi.md#getaddress) | **Get** /address/{address} | Get address info
*AddressesApi* | [**GetAddressUtxos**](docs/AddressesApi.md#getaddressutxos) | **Get** /address/{address}/outputs | Get UTXOs for an address
*AddressesApi* | [**ValidateAddress**](docs/AddressesApi.md#validateaddress) | **Get** /address/{address}/validate | Validate address
*AddressesApi* | [**VerifyMessage**](docs/AddressesApi.md#verifymessage) | **Post** /address/verify-message | Verify message
*BlocksApi* | [**GetBlockCount**](docs/BlocksApi.md#getblockcount) | **Get** /blockcount | Get the height of the latest block
*BlocksApi* | [**GetBlockDecoded**](docs/BlocksApi.md#getblockdecoded) | **Get** /block/raw/{identifier}/decoded | Get block by hash or height (verbosity 2)
*BlocksApi* | [**GetBlockHashByHeight**](docs/BlocksApi.md#getblockhashbyheight) | **Get** /blockhash/{block_height} | Returns blockhash of specified block.
*BlocksApi* | [**GetBlockHex**](docs/BlocksApi.md#getblockhex) | **Get** /block/raw/{identifier}/hex | Get block by hash or height (verbosity 0)
*BlocksApi* | [**GetBlockInfo**](docs/BlocksApi.md#getblockinfo) | **Get** /block/{identifier} | Get block info by hash or height
*BlocksApi* | [**GetBlockPrevout**](docs/BlocksApi.md#getblockprevout) | **Get** /block/raw/{identifier}/prevout | Get block by hash or height (verbosity 3)
*BlocksApi* | [**GetBlockStats**](docs/BlocksApi.md#getblockstats) | **Post** /block/stats | Get block stats
*BlocksApi* | [**GetBlockSummary**](docs/BlocksApi.md#getblocksummary) | **Get** /block/raw/{identifier}/summary | Get block by hash or height (verbosity 1)
*BlocksApi* | [**GetBlockchainInfo**](docs/BlocksApi.md#getblockchaininfo) | **Get** /blockchain/info | Get blockchain information
*BlocksApi* | [**GetBlocks**](docs/BlocksApi.md#getblocks) | **Get** /blocks | Returns the latest block height, last 100 block hashes, and featured inscriptions
*BlocksApi* | [**GetLatestBlockHeight**](docs/BlocksApi.md#getlatestblockheight) | **Get** /blockheight | Returns the height of the latest block.
*BlocksApi* | [**GetLatestBlockhash**](docs/BlocksApi.md#getlatestblockhash) | **Get** /blockhash | Returns blockhash for the latest block.
*BlocksApi* | [**GetLatestBlocktime**](docs/BlocksApi.md#getlatestblocktime) | **Get** /blocktime | Get the timestamp of the latest block
*FeesApi* | [**EstimateRawFee**](docs/FeesApi.md#estimaterawfee) | **Post** /fee/estimate-raw | Estimate Raw Fee
*FeesApi* | [**EstimateSmartFee**](docs/FeesApi.md#estimatesmartfee) | **Post** /fee/estimate-smart | Estimate smart fee
*InscriptionsApi* | [**FetchInscriptionChild**](docs/InscriptionsApi.md#fetchinscriptionchild) | **Get** /inscription/{inscription_id}/{child_index} | Get inscription child info
*InscriptionsApi* | [**FetchInscriptions**](docs/InscriptionsApi.md#fetchinscriptions) | **Post** /inscriptions | Fetch multiple inscriptions
*InscriptionsApi* | [**GetBlockInscriptions**](docs/InscriptionsApi.md#getblockinscriptions) | **Get** /inscriptions/block/{block_height} | Get inscriptions in a specific block
*InscriptionsApi* | [**GetInscription**](docs/InscriptionsApi.md#getinscription) | **Get** /inscription/{inscription_id} | Get inscription info
*InscriptionsApi* | [**GetLatestInscriptions**](docs/InscriptionsApi.md#getlatestinscriptions) | **Get** /inscriptions | Get latest inscriptions
*InscriptionsApi* | [**GetLatestInscriptionsPage**](docs/InscriptionsApi.md#getlatestinscriptionspage) | **Get** /inscriptions/{page} | Get latest inscriptions page
*MempoolApi* | [**GetMempoolAncestors**](docs/MempoolApi.md#getmempoolancestors) | **Post** /mempool/ancestors | Get mempool ancestors
*MempoolApi* | [**GetMempoolDescendants**](docs/MempoolApi.md#getmempooldescendants) | **Post** /mempool/descendants | Get mempool descendants
*MempoolApi* | [**GetMempoolInfo**](docs/MempoolApi.md#getmempoolinfo) | **Get** /mempool/info | Get mempool information
*MempoolApi* | [**GetRawMempool**](docs/MempoolApi.md#getrawmempool) | **Post** /mempool/raw | Get raw mempool
*MempoolApi* | [**TestMempoolAccept**](docs/MempoolApi.md#testmempoolaccept) | **Post** /mempool/test-accept | Test mempool accept
*MiningApi* | [**GetMiningInfo**](docs/MiningApi.md#getmininginfo) | **Get** /mining/info | Get mining information
*MiningApi* | [**GetNetworkHashps**](docs/MiningApi.md#getnetworkhashps) | **Post** /mining/networkhashps | Get network hash per second
*NetworkApi* | [**GetChainTxStats**](docs/NetworkApi.md#getchaintxstats) | **Post** /chain/txstats | Get chain tx stats
*NetworkApi* | [**GetDifficulty**](docs/NetworkApi.md#getdifficulty) | **Get** /chain/difficulty | Get difficulty
*OutputsApi* | [**GetOutputByOutpoint**](docs/OutputsApi.md#getoutputbyoutpoint) | **Get** /output/{outpoint} | Get output info by outpoint
*OutputsApi* | [**GetOutputs**](docs/OutputsApi.md#getoutputs) | **Post** /outputs | Get multiple outputs
*PSBTsApi* | [**AnalyzePsbt**](docs/PSBTsApi.md#analyzepsbt) | **Post** /psbt/analyze | Analyze PSBT
*PSBTsApi* | [**CombinePsbt**](docs/PSBTsApi.md#combinepsbt) | **Post** /psbt/combine | Combine PSBTs
*PSBTsApi* | [**CreatePsbt**](docs/PSBTsApi.md#createpsbt) | **Post** /psbt/create | Create PSBT
*PSBTsApi* | [**DecodePsbt**](docs/PSBTsApi.md#decodepsbt) | **Post** /psbt/decode | Decode PSBT
*PSBTsApi* | [**JoinPsbts**](docs/PSBTsApi.md#joinpsbts) | **Post** /psbt/join | Join PSBTs
*RunesApi* | [**GetLatestRunes**](docs/RunesApi.md#getlatestrunes) | **Get** /runes | Get latest runes
*RunesApi* | [**GetLatestRunesPage**](docs/RunesApi.md#getlatestrunespage) | **Get** /runes/{page} | Get latest runes page
*RunesApi* | [**GetRune**](docs/RunesApi.md#getrune) | **Get** /rune/{rune_name} | Get rune info
*SatoshisApi* | [**GetSatoshi**](docs/SatoshisApi.md#getsatoshi) | **Get** /sat/{number} | Get satoshi info
*ScriptsApi* | [**DecodeScript**](docs/ScriptsApi.md#decodescript) | **Post** /script/decode | Decode Script
*StatusApi* | [**GetStatus**](docs/StatusApi.md#getstatus) | **Get** /status | Get server status
*TransactionsApi* | [**CombineRawTransaction**](docs/TransactionsApi.md#combinerawtransaction) | **Post** /tx/combine | Combine Raw Transactions
*TransactionsApi* | [**ConvertToPsbt**](docs/TransactionsApi.md#converttopsbt) | **Post** /tx/convert-to-psbt | Convert Raw Transaction to PSBT
*TransactionsApi* | [**CreateRawTransaction**](docs/TransactionsApi.md#createrawtransaction) | **Post** /tx/create | Create Raw Transaction
*TransactionsApi* | [**DecodeTx**](docs/TransactionsApi.md#decodetx) | **Get** /tx/{txid}/decode | Decode a transaction
*TransactionsApi* | [**GetRawTransactionDecoded**](docs/TransactionsApi.md#getrawtransactiondecoded) | **Get** /tx/{txid}/decoded | Get raw transaction (verbosity 1)
*TransactionsApi* | [**GetRawTransactionHex**](docs/TransactionsApi.md#getrawtransactionhex) | **Get** /tx/{txid}/hex | Get raw transaction (verbosity 0)
*TransactionsApi* | [**GetRawTransactionPrevout**](docs/TransactionsApi.md#getrawtransactionprevout) | **Get** /tx/{txid}/prevout | Get raw transaction (verbosity 2)
*TransactionsApi* | [**GetTransaction**](docs/TransactionsApi.md#gettransaction) | **Get** /tx/{txid} | Get transaction info
*TransactionsApi* | [**GetTxOut**](docs/TransactionsApi.md#gettxout) | **Post** /tx/out | Get transaction output
*TransactionsApi* | [**GetTxOutProof**](docs/TransactionsApi.md#gettxoutproof) | **Post** /tx/outproof | Get transaction output proof
*TransactionsApi* | [**GetTxOutSetInfo**](docs/TransactionsApi.md#gettxoutsetinfo) | **Post** /tx/out/set/info | Get transaction output set information
*TransactionsApi* | [**GetTxSpendingPrevout**](docs/TransactionsApi.md#gettxspendingprevout) | **Post** /tx/spending-prevout | Get transaction spending prevout
*TransactionsApi* | [**SendRawTransaction**](docs/TransactionsApi.md#sendrawtransaction) | **Post** /tx/send | Send raw transaction
*TransactionsApi* | [**VerifyTxOutProof**](docs/TransactionsApi.md#verifytxoutproof) | **Post** /tx/outproof/verify | Verify transaction output proof

## Documentation For Models

 - [AddressResponse](docs/AddressResponse.md)
 - [AllOfBlockVin2ScriptSig](docs/AllOfBlockVin2ScriptSig.md)
 - [AllOfBlockVin3Prevout](docs/AllOfBlockVin3Prevout.md)
 - [AllOfDecodedPsbtInputFinalScriptsig](docs/AllOfDecodedPsbtInputFinalScriptsig.md)
 - [AllOfDecodedPsbtInputNonWitnessUtxo](docs/AllOfDecodedPsbtInputNonWitnessUtxo.md)
 - [AllOfDecodedPsbtInputRedeemScript](docs/AllOfDecodedPsbtInputRedeemScript.md)
 - [AllOfDecodedPsbtInputWitnessScript](docs/AllOfDecodedPsbtInputWitnessScript.md)
 - [AllOfDecodedPsbtInputWitnessUtxo](docs/AllOfDecodedPsbtInputWitnessUtxo.md)
 - [AllOfDecodedPsbtOutputRedeemScript](docs/AllOfDecodedPsbtOutputRedeemScript.md)
 - [AllOfDecodedPsbtOutputWitnessScript](docs/AllOfDecodedPsbtOutputWitnessScript.md)
 - [AllOfDecodedPsbttx](docs/AllOfDecodedPsbttx.md)
 - [AllOfDecodedScriptSegwit](docs/AllOfDecodedScriptSegwit.md)
 - [AllOfFeeHorizonEstimateFail](docs/AllOfFeeHorizonEstimateFail.md)
 - [AllOfFeeHorizonEstimatePass](docs/AllOfFeeHorizonEstimatePass.md)
 - [AllOfMempoolEntryFees](docs/AllOfMempoolEntryFees.md)
 - [AllOfPsbtWitnessUtxoScriptPubKey](docs/AllOfPsbtWitnessUtxoScriptPubKey.md)
 - [AllOfRawFeeEstimateLong](docs/AllOfRawFeeEstimateLong.md)
 - [AllOfRawFeeEstimateMedium](docs/AllOfRawFeeEstimateMedium.md)
 - [AllOfRawFeeEstimateShort](docs/AllOfRawFeeEstimateShort.md)
 - [AllOfRawMempoolDataSequence](docs/AllOfRawMempoolDataSequence.md)
 - [AllOfRuneListEntryVar1](docs/AllOfRuneListEntryVar1.md)
 - [AllOfTestMempoolAcceptResultFees](docs/AllOfTestMempoolAcceptResultFees.md)
 - [AllOfTxOutScriptPubKey](docs/AllOfTxOutScriptPubKey.md)
 - [AllOfUtxoBlockInfoUnspendables](docs/AllOfUtxoBlockInfoUnspendables.md)
 - [AllOfUtxoSetInfoBlockInfo](docs/AllOfUtxoSetInfoBlockInfo.md)
 - [Bip32Deriv](docs/Bip32Deriv.md)
 - [Block1](docs/Block1.md)
 - [Block2](docs/Block2.md)
 - [Block3](docs/Block3.md)
 - [BlockResponse](docs/BlockResponse.md)
 - [BlockStats](docs/BlockStats.md)
 - [BlockVin2](docs/BlockVin2.md)
 - [BlockVin3](docs/BlockVin3.md)
 - [BlockchainInfo](docs/BlockchainInfo.md)
 - [BlocksResponse](docs/BlocksResponse.md)
 - [BtcTx2](docs/BtcTx2.md)
 - [BtcTx3](docs/BtcTx3.md)
 - [ChainTxStats](docs/ChainTxStats.md)
 - [DecodeResponse](docs/DecodeResponse.md)
 - [DecodedInscription](docs/DecodedInscription.md)
 - [DecodedPsbt](docs/DecodedPsbt.md)
 - [DecodedPsbtInput](docs/DecodedPsbtInput.md)
 - [DecodedPsbtOutput](docs/DecodedPsbtOutput.md)
 - [DecodedScript](docs/DecodedScript.md)
 - [Duration](docs/Duration.md)
 - [FeeHorizonEstimate](docs/FeeHorizonEstimate.md)
 - [FeeRange](docs/FeeRange.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse20013](docs/InlineResponse20013.md)
 - [InlineResponse20014](docs/InlineResponse20014.md)
 - [InlineResponse20015](docs/InlineResponse20015.md)
 - [InlineResponse20016](docs/InlineResponse20016.md)
 - [InlineResponse20017](docs/InlineResponse20017.md)
 - [InlineResponse20018](docs/InlineResponse20018.md)
 - [InlineResponse20019](docs/InlineResponse20019.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse20020](docs/InlineResponse20020.md)
 - [InlineResponse20021](docs/InlineResponse20021.md)
 - [InlineResponse20022](docs/InlineResponse20022.md)
 - [InlineResponse20023](docs/InlineResponse20023.md)
 - [InlineResponse20024](docs/InlineResponse20024.md)
 - [InlineResponse20025](docs/InlineResponse20025.md)
 - [InlineResponse20026](docs/InlineResponse20026.md)
 - [InlineResponse20027](docs/InlineResponse20027.md)
 - [InlineResponse20028](docs/InlineResponse20028.md)
 - [InlineResponse20029](docs/InlineResponse20029.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse20030](docs/InlineResponse20030.md)
 - [InlineResponse20031](docs/InlineResponse20031.md)
 - [InlineResponse20032](docs/InlineResponse20032.md)
 - [InlineResponse20033](docs/InlineResponse20033.md)
 - [InlineResponse20034](docs/InlineResponse20034.md)
 - [InlineResponse20035](docs/InlineResponse20035.md)
 - [InlineResponse20036](docs/InlineResponse20036.md)
 - [InlineResponse20037](docs/InlineResponse20037.md)
 - [InlineResponse20038](docs/InlineResponse20038.md)
 - [InlineResponse20039](docs/InlineResponse20039.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse20040](docs/InlineResponse20040.md)
 - [InlineResponse20041](docs/InlineResponse20041.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [Input](docs/Input.md)
 - [InscriptionData](docs/InscriptionData.md)
 - [InscriptionResponse](docs/InscriptionResponse.md)
 - [LatestInscriptionsResponse](docs/LatestInscriptionsResponse.md)
 - [MempoolAncestorsData](docs/MempoolAncestorsData.md)
 - [MempoolDescendantsData](docs/MempoolDescendantsData.md)
 - [MempoolEntry](docs/MempoolEntry.md)
 - [MempoolFees](docs/MempoolFees.md)
 - [MempoolInfo](docs/MempoolInfo.md)
 - [MempoolSequence](docs/MempoolSequence.md)
 - [MiningInfo](docs/MiningInfo.md)
 - [Output](docs/Output.md)
 - [OutputResponse](docs/OutputResponse.md)
 - [PrevOut](docs/PrevOut.md)
 - [PsbtAnalysis](docs/PsbtAnalysis.md)
 - [PsbtBip32Deriv](docs/PsbtBip32Deriv.md)
 - [PsbtInputAnalysis](docs/PsbtInputAnalysis.md)
 - [PsbtMissingData](docs/PsbtMissingData.md)
 - [PsbtWitnessUtxo](docs/PsbtWitnessUtxo.md)
 - [RawFeeEstimate](docs/RawFeeEstimate.md)
 - [RawMempoolData](docs/RawMempoolData.md)
 - [RawTx1](docs/RawTx1.md)
 - [RawTx2](docs/RawTx2.md)
 - [RequestsAnalyzePsbtRequest](docs/RequestsAnalyzePsbtRequest.md)
 - [RequestsCombinePsbtRequest](docs/RequestsCombinePsbtRequest.md)
 - [RequestsCombineRawTransactionRequest](docs/RequestsCombineRawTransactionRequest.md)
 - [RequestsConvertToPsbtRequest](docs/RequestsConvertToPsbtRequest.md)
 - [RequestsCreatePsbtInput](docs/RequestsCreatePsbtInput.md)
 - [RequestsCreatePsbtOutput](docs/RequestsCreatePsbtOutput.md)
 - [RequestsCreatePsbtRequest](docs/RequestsCreatePsbtRequest.md)
 - [RequestsCreateRawTxInput](docs/RequestsCreateRawTxInput.md)
 - [RequestsCreateRawTxOutput](docs/RequestsCreateRawTxOutput.md)
 - [RequestsCreateRawTxRequest](docs/RequestsCreateRawTxRequest.md)
 - [RequestsDecodeScriptRequest](docs/RequestsDecodeScriptRequest.md)
 - [RequestsEstimateRawFeeRequest](docs/RequestsEstimateRawFeeRequest.md)
 - [RequestsEstimateSmartFeeRequest](docs/RequestsEstimateSmartFeeRequest.md)
 - [RequestsGetBlockStatsRequest](docs/RequestsGetBlockStatsRequest.md)
 - [RequestsGetChainTxStatsRequest](docs/RequestsGetChainTxStatsRequest.md)
 - [RequestsGetMempoolAncestorsRequest](docs/RequestsGetMempoolAncestorsRequest.md)
 - [RequestsGetMempoolDescendantsRequest](docs/RequestsGetMempoolDescendantsRequest.md)
 - [RequestsGetNetworkHashPsRequest](docs/RequestsGetNetworkHashPsRequest.md)
 - [RequestsGetRawMempoolRequest](docs/RequestsGetRawMempoolRequest.md)
 - [RequestsGetTxOutProofRequest](docs/RequestsGetTxOutProofRequest.md)
 - [RequestsGetTxOutRequest](docs/RequestsGetTxOutRequest.md)
 - [RequestsGetTxOutSetInfoRequest](docs/RequestsGetTxOutSetInfoRequest.md)
 - [RequestsGetTxSpendingPrevoutRequest](docs/RequestsGetTxSpendingPrevoutRequest.md)
 - [RequestsJoinPsbtsRequest](docs/RequestsJoinPsbtsRequest.md)
 - [RequestsSendRawTransactionRequest](docs/RequestsSendRawTransactionRequest.md)
 - [RequestsTestMempoolAcceptRequest](docs/RequestsTestMempoolAcceptRequest.md)
 - [RequestsVerifyMessageRequest](docs/RequestsVerifyMessageRequest.md)
 - [RequestsVerifyTxOutProofRequest](docs/RequestsVerifyTxOutProofRequest.md)
 - [RuneEntry](docs/RuneEntry.md)
 - [RuneListEntry](docs/RuneListEntry.md)
 - [RuneResponse](docs/RuneResponse.md)
 - [RuneTerms](docs/RuneTerms.md)
 - [RunesBalance](docs/RunesBalance.md)
 - [RunesListResponse](docs/RunesListResponse.md)
 - [RunestoneData](docs/RunestoneData.md)
 - [SatoshiResponse](docs/SatoshiResponse.md)
 - [Script](docs/Script.md)
 - [ScriptPubKey](docs/ScriptPubKey.md)
 - [ScriptSig](docs/ScriptSig.md)
 - [SegwitDetails](docs/SegwitDetails.md)
 - [SmartFeeEstimate](docs/SmartFeeEstimate.md)
 - [StatusResponse](docs/StatusResponse.md)
 - [TestMempoolAcceptResult](docs/TestMempoolAcceptResult.md)
 - [TestMempoolFees](docs/TestMempoolFees.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionDetails](docs/TransactionDetails.md)
 - [TransactionInput](docs/TransactionInput.md)
 - [TransactionOutput](docs/TransactionOutput.md)
 - [TransactionResponse](docs/TransactionResponse.md)
 - [TxOut](docs/TxOut.md)
 - [TxSpendingPrevoutInput](docs/TxSpendingPrevoutInput.md)
 - [TxSpendingPrevoutResult](docs/TxSpendingPrevoutResult.md)
 - [TxVin1](docs/TxVin1.md)
 - [TxVin2](docs/TxVin2.md)
 - [UtilsResponseEnvelope](docs/UtilsResponseEnvelope.md)
 - [UtxoBlockInfo](docs/UtxoBlockInfo.md)
 - [UtxoSetInfo](docs/UtxoSetInfo.md)
 - [UtxoUnspendables](docs/UtxoUnspendables.md)
 - [ValidateAddressResult](docs/ValidateAddressResult.md)
 - [Vout](docs/Vout.md)

## Documentation For Authorization

## ApiKeyAuth
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

team@satstream.io
