# DecodedPsbtInput

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bip32Derivs** | [**map[string]Bip32Deriv**](Bip32Deriv.md) | The BIP32 derivation paths | [optional] [default to null]
**FinalScriptsig** | [***AllOfDecodedPsbtInputFinalScriptsig**](AllOfDecodedPsbtInputFinalScriptsig.md) | The final script sig | [optional] [default to null]
**FinalScriptwitness** | **[]string** | The final script witness | [optional] [default to null]
**NonWitnessUtxo** | [***AllOfDecodedPsbtInputNonWitnessUtxo**](AllOfDecodedPsbtInputNonWitnessUtxo.md) | Decoded network transaction for non-witness UTXOs | [optional] [default to null]
**PartialSignatures** | **map[string]string** | The public key and signature pairs | [optional] [default to null]
**RedeemScript** | [***AllOfDecodedPsbtInputRedeemScript**](AllOfDecodedPsbtInputRedeemScript.md) | The redeem script | [optional] [default to null]
**Sighash** | **string** | The sighash type to be used | [optional] [default to null]
**Unknown** | [***AllOfDecodedPsbtInputUnknown**](AllOfDecodedPsbtInputUnknown.md) | Unknown fields | [optional] [default to null]
**WitnessScript** | [***AllOfDecodedPsbtInputWitnessScript**](AllOfDecodedPsbtInputWitnessScript.md) | The witness script | [optional] [default to null]
**WitnessUtxo** | [***AllOfDecodedPsbtInputWitnessUtxo**](AllOfDecodedPsbtInputWitnessUtxo.md) | Transaction output for witness UTXOs | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

