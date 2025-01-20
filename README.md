# Go API client for satstream_go_sdk

Satstream API

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.44
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
*AddressesApi* | [**GetAddressBalance**](docs/AddressesApi.md#getaddressbalance) | **Get** /address/{address}/balance | Get address balance
*AddressesApi* | [**GetAddressDeltas**](docs/AddressesApi.md#getaddressdeltas) | **Get** /address/{address}/deltas | Get address deltas
*AddressesApi* | [**GetAddressRuneDeltas**](docs/AddressesApi.md#getaddressrunedeltas) | **Get** /address/{address}/deltas/runes | Get address rune deltas
*AddressesApi* | [**GetAddressUtxos**](docs/AddressesApi.md#getaddressutxos) | **Get** /address/{address}/outputs | Get UTXOs for an address
*AddressesApi* | [**ValidateAddress**](docs/AddressesApi.md#validateaddress) | **Get** /address/{address}/validate | Validate address
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
*InscriptionsApi* | [**GetBlockInscriptionsPage**](docs/InscriptionsApi.md#getblockinscriptionspage) | **Get** /inscriptions/block/{block_height}/{page} | Get paginated inscriptions in a specific block
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
*RunesApi* | [**GetRune**](docs/RunesApi.md#getrune) | **Get** /rune/{identifier} | Get rune info
*SatoshisApi* | [**GetSatoshi**](docs/SatoshisApi.md#getsatoshi) | **Get** /sat/{number} | Get satoshi info
*ScriptsApi* | [**DecodeScript**](docs/ScriptsApi.md#decodescript) | **Post** /script/decode | Decode Script
*StatusApi* | [**GetStatus**](docs/StatusApi.md#getstatus) | **Get** /status | Get server status
*TransactionsApi* | [**CombineRawTransaction**](docs/TransactionsApi.md#combinerawtransaction) | **Post** /tx/combine | Combine Raw Transactions
*TransactionsApi* | [**ConvertToPsbt**](docs/TransactionsApi.md#converttopsbt) | **Post** /tx/convert-to-psbt | Convert Raw Transaction to PSBT
*TransactionsApi* | [**CreateRawTransaction**](docs/TransactionsApi.md#createrawtransaction) | **Post** /tx/create | Create Raw Transaction
*TransactionsApi* | [**DecodeTxInscriptions**](docs/TransactionsApi.md#decodetxinscriptions) | **Get** /tx/{txid}/inscriptions | Decode transaction inscriptions
*TransactionsApi* | [**GetRawTransaction**](docs/TransactionsApi.md#getrawtransaction) | **Get** /tx/{txid}/raw/decode | Get raw transaction (verbosity 1)
*TransactionsApi* | [**GetRawTransactionHex**](docs/TransactionsApi.md#getrawtransactionhex) | **Get** /tx/{txid}/hex | Get raw transaction (verbosity 0)
*TransactionsApi* | [**GetRawTransactionPrevout**](docs/TransactionsApi.md#getrawtransactionprevout) | **Get** /tx/{txid}/raw/prevout | Get raw transaction with prevouts (verbosity 2)
*TransactionsApi* | [**GetTxOut**](docs/TransactionsApi.md#gettxout) | **Post** /tx/out | Get transaction output
*TransactionsApi* | [**GetTxOutProof**](docs/TransactionsApi.md#gettxoutproof) | **Post** /tx/outproof | Get transaction output proof
*TransactionsApi* | [**GetTxOutSetInfo**](docs/TransactionsApi.md#gettxoutsetinfo) | **Post** /tx/out/set/info | Get transaction output set information
*TransactionsApi* | [**GetTxSpendingPrevout**](docs/TransactionsApi.md#gettxspendingprevout) | **Post** /tx/spending-prevout | Get transaction spending prevout
*TransactionsApi* | [**SendRawTransaction**](docs/TransactionsApi.md#sendrawtransaction) | **Post** /tx/send | Send raw transaction
*TransactionsApi* | [**VerifyTxOutProof**](docs/TransactionsApi.md#verifytxoutproof) | **Post** /tx/outproof/verify | Verify transaction output proof

## Documentation For Models

 - [AddressResponse](docs/AddressResponse.md)
 - [AddressRuneDelta](docs/AddressRuneDelta.md)
 - [AllOfBlockVin2ScriptSig](docs/AllOfBlockVin2ScriptSig.md)
 - [AllOfBlockVin3Prevout](docs/AllOfBlockVin3Prevout.md)
 - [AllOfDecodedPsbtInputFinalScriptsig](docs/AllOfDecodedPsbtInputFinalScriptsig.md)
 - [AllOfDecodedPsbtInputNonWitnessUtxo](docs/AllOfDecodedPsbtInputNonWitnessUtxo.md)
 - [AllOfDecodedPsbtInputRedeemScript](docs/AllOfDecodedPsbtInputRedeemScript.md)
 - [AllOfDecodedPsbtInputUnknown](docs/AllOfDecodedPsbtInputUnknown.md)
 - [AllOfDecodedPsbtInputWitnessScript](docs/AllOfDecodedPsbtInputWitnessScript.md)
 - [AllOfDecodedPsbtInputWitnessUtxo](docs/AllOfDecodedPsbtInputWitnessUtxo.md)
 - [AllOfDecodedPsbtOutputRedeemScript](docs/AllOfDecodedPsbtOutputRedeemScript.md)
 - [AllOfDecodedPsbtOutputUnknown](docs/AllOfDecodedPsbtOutputUnknown.md)
 - [AllOfDecodedPsbtOutputWitnessScript](docs/AllOfDecodedPsbtOutputWitnessScript.md)
 - [AllOfDecodedPsbtUnknown](docs/AllOfDecodedPsbtUnknown.md)
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
 - [AllOfTestMempoolAcceptResultFees](docs/AllOfTestMempoolAcceptResultFees.md)
 - [AllOfTxOutScriptPubKey](docs/AllOfTxOutScriptPubKey.md)
 - [AllOfUtxoBlockInfoUnspendables](docs/AllOfUtxoBlockInfoUnspendables.md)
 - [AllOfUtxoSetInfoBlockInfo](docs/AllOfUtxoSetInfoBlockInfo.md)
 - [AnalyzePsbtRequest](docs/AnalyzePsbtRequest.md)
 - [AnalyzePsbtResponse](docs/AnalyzePsbtResponse.md)
 - [BigInt](docs/BigInt.md)
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
 - [CombinePsbtRequest](docs/CombinePsbtRequest.md)
 - [CombinePsbtResponse](docs/CombinePsbtResponse.md)
 - [CombineRawTransactionResponse](docs/CombineRawTransactionResponse.md)
 - [ConvertToPsbtResponse](docs/ConvertToPsbtResponse.md)
 - [CreatePsbtInput](docs/CreatePsbtInput.md)
 - [CreatePsbtRequest](docs/CreatePsbtRequest.md)
 - [CreatePsbtResponse](docs/CreatePsbtResponse.md)
 - [CreateRawTransactionResponse](docs/CreateRawTransactionResponse.md)
 - [DecodePsbtRequest](docs/DecodePsbtRequest.md)
 - [DecodePsbtResponse](docs/DecodePsbtResponse.md)
 - [DecodeResponse](docs/DecodeResponse.md)
 - [DecodeScriptRequest](docs/DecodeScriptRequest.md)
 - [DecodeScriptResponse](docs/DecodeScriptResponse.md)
 - [DecodeTransactionResponse](docs/DecodeTransactionResponse.md)
 - [DecodedInscription](docs/DecodedInscription.md)
 - [DecodedPsbt](docs/DecodedPsbt.md)
 - [DecodedPsbtInput](docs/DecodedPsbtInput.md)
 - [DecodedPsbtOutput](docs/DecodedPsbtOutput.md)
 - [DecodedScript](docs/DecodedScript.md)
 - [EstimateRawFeeRequest](docs/EstimateRawFeeRequest.md)
 - [EstimateRawFeeResponse](docs/EstimateRawFeeResponse.md)
 - [EstimateSmartFeeRequest](docs/EstimateSmartFeeRequest.md)
 - [EstimateSmartFeeResponse](docs/EstimateSmartFeeResponse.md)
 - [FeeHorizonEstimate](docs/FeeHorizonEstimate.md)
 - [FeeRange](docs/FeeRange.md)
 - [FetchInscriptionsResponse](docs/FetchInscriptionsResponse.md)
 - [GetAddressBalanceResponse](docs/GetAddressBalanceResponse.md)
 - [GetAddressBalanceResponseData](docs/GetAddressBalanceResponseData.md)
 - [GetAddressDeltasResponse](docs/GetAddressDeltasResponse.md)
 - [GetAddressResponse](docs/GetAddressResponse.md)
 - [GetAddressRuneDeltasResponse](docs/GetAddressRuneDeltasResponse.md)
 - [GetAddressUtxosResponse](docs/GetAddressUtxosResponse.md)
 - [GetBlockCountResponse](docs/GetBlockCountResponse.md)
 - [GetBlockDecodedResponse](docs/GetBlockDecodedResponse.md)
 - [GetBlockHashByHeightResponse](docs/GetBlockHashByHeightResponse.md)
 - [GetBlockHexResponse](docs/GetBlockHexResponse.md)
 - [GetBlockInscriptionsResponse](docs/GetBlockInscriptionsResponse.md)
 - [GetBlockPrevoutResponse](docs/GetBlockPrevoutResponse.md)
 - [GetBlockResponse](docs/GetBlockResponse.md)
 - [GetBlockStatsRequest](docs/GetBlockStatsRequest.md)
 - [GetBlockStatsResponse](docs/GetBlockStatsResponse.md)
 - [GetBlockSummaryResponse](docs/GetBlockSummaryResponse.md)
 - [GetBlockchainInfoResponse](docs/GetBlockchainInfoResponse.md)
 - [GetBlocksResponse](docs/GetBlocksResponse.md)
 - [GetChainTxStatsRequest](docs/GetChainTxStatsRequest.md)
 - [GetChainTxStatsResponse](docs/GetChainTxStatsResponse.md)
 - [GetDifficultyResponse](docs/GetDifficultyResponse.md)
 - [GetInscriptionChildResponse](docs/GetInscriptionChildResponse.md)
 - [GetInscriptionResponse](docs/GetInscriptionResponse.md)
 - [GetLatestBlockHashResponse](docs/GetLatestBlockHashResponse.md)
 - [GetLatestBlockHeightResponse](docs/GetLatestBlockHeightResponse.md)
 - [GetLatestBlockTimeResponse](docs/GetLatestBlockTimeResponse.md)
 - [GetLatestInscriptionsResponse](docs/GetLatestInscriptionsResponse.md)
 - [GetLatestRunesResponse](docs/GetLatestRunesResponse.md)
 - [GetMempoolAncestorsRequest](docs/GetMempoolAncestorsRequest.md)
 - [GetMempoolAncestorsResponse](docs/GetMempoolAncestorsResponse.md)
 - [GetMempoolDescendantsRequest](docs/GetMempoolDescendantsRequest.md)
 - [GetMempoolDescendantsResponse](docs/GetMempoolDescendantsResponse.md)
 - [GetMempoolInfoResponse](docs/GetMempoolInfoResponse.md)
 - [GetMiningInfoResponse](docs/GetMiningInfoResponse.md)
 - [GetNetworkHashPsRequest](docs/GetNetworkHashPsRequest.md)
 - [GetNetworkHashPsResponse](docs/GetNetworkHashPsResponse.md)
 - [GetRawMempoolRequest](docs/GetRawMempoolRequest.md)
 - [GetRawMempoolResponse](docs/GetRawMempoolResponse.md)
 - [GetRawTransactionDecodeResponse](docs/GetRawTransactionDecodeResponse.md)
 - [GetRawTransactionHexResponse](docs/GetRawTransactionHexResponse.md)
 - [GetRawTransactionPrevoutResponse](docs/GetRawTransactionPrevoutResponse.md)
 - [GetRuneResponse](docs/GetRuneResponse.md)
 - [GetSatoshiResponse](docs/GetSatoshiResponse.md)
 - [GetStatusResponse](docs/GetStatusResponse.md)
 - [GetTxOutProofResponse](docs/GetTxOutProofResponse.md)
 - [GetTxOutResponse](docs/GetTxOutResponse.md)
 - [GetTxSpendingPrevoutResponse](docs/GetTxSpendingPrevoutResponse.md)
 - [GithubComSatstreamSsUtilsDatabaseAddressDelta](docs/GithubComSatstreamSsUtilsDatabaseAddressDelta.md)
 - [GithubComSatstreamSsUtilsOrdServerResponsesRuneDetails](docs/GithubComSatstreamSsUtilsOrdServerResponsesRuneDetails.md)
 - [GithubComSatstreamSsUtilsOrdServerResponsesRuneListEntry](docs/GithubComSatstreamSsUtilsOrdServerResponsesRuneListEntry.md)
 - [GithubComSatstreamSsUtilsOrdServerResponsesRunesListResponse](docs/GithubComSatstreamSsUtilsOrdServerResponsesRunesListResponse.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [Input](docs/Input.md)
 - [InscriptionData](docs/InscriptionData.md)
 - [InscriptionResponse](docs/InscriptionResponse.md)
 - [JoinPsbtsRequest](docs/JoinPsbtsRequest.md)
 - [JoinPsbtsResponse](docs/JoinPsbtsResponse.md)
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
 - [RuneResponse](docs/RuneResponse.md)
 - [RuneTerms](docs/RuneTerms.md)
 - [RunesBalance](docs/RunesBalance.md)
 - [RunestoneData](docs/RunestoneData.md)
 - [SatoshiResponse](docs/SatoshiResponse.md)
 - [Script](docs/Script.md)
 - [ScriptPubKey](docs/ScriptPubKey.md)
 - [ScriptSig](docs/ScriptSig.md)
 - [SegwitDetails](docs/SegwitDetails.md)
 - [SendRawTransactionResponse](docs/SendRawTransactionResponse.md)
 - [SmartFeeEstimate](docs/SmartFeeEstimate.md)
 - [StatusResponse](docs/StatusResponse.md)
 - [TestMempoolAcceptRequest](docs/TestMempoolAcceptRequest.md)
 - [TestMempoolAcceptResponse](docs/TestMempoolAcceptResponse.md)
 - [TestMempoolAcceptResult](docs/TestMempoolAcceptResult.md)
 - [TestMempoolFees](docs/TestMempoolFees.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionCombineRawTransactionRequest](docs/TransactionCombineRawTransactionRequest.md)
 - [TransactionConvertToPsbtRequest](docs/TransactionConvertToPsbtRequest.md)
 - [TransactionCreateRawTxInput](docs/TransactionCreateRawTxInput.md)
 - [TransactionCreateRawTxRequest](docs/TransactionCreateRawTxRequest.md)
 - [TransactionGetTxOutProofRequest](docs/TransactionGetTxOutProofRequest.md)
 - [TransactionGetTxOutRequest](docs/TransactionGetTxOutRequest.md)
 - [TransactionGetTxOutSetInfoRequest](docs/TransactionGetTxOutSetInfoRequest.md)
 - [TransactionGetTxSpendingPrevoutRequest](docs/TransactionGetTxSpendingPrevoutRequest.md)
 - [TransactionSendRawTransactionRequest](docs/TransactionSendRawTransactionRequest.md)
 - [TransactionVerifyTxOutProofRequest](docs/TransactionVerifyTxOutProofRequest.md)
 - [TxOut](docs/TxOut.md)
 - [TxSpendingPrevoutInput](docs/TxSpendingPrevoutInput.md)
 - [TxSpendingPrevoutResult](docs/TxSpendingPrevoutResult.md)
 - [TxVin1](docs/TxVin1.md)
 - [TxVin2](docs/TxVin2.md)
 - [UnknownFields](docs/UnknownFields.md)
 - [UtilsResponseEnvelope](docs/UtilsResponseEnvelope.md)
 - [UtxoBlockInfo](docs/UtxoBlockInfo.md)
 - [UtxoSetInfo](docs/UtxoSetInfo.md)
 - [UtxoUnspendables](docs/UtxoUnspendables.md)
 - [ValidateAddressResponse](docs/ValidateAddressResponse.md)
 - [ValidateAddressResult](docs/ValidateAddressResult.md)
 - [VerifyTxOutProofResponse](docs/VerifyTxOutProofResponse.md)
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
