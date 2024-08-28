// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package satstream

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/satstream-go/internal/apijson"
	"github.com/stainless-sdks/satstream-go/internal/requestconfig"
	"github.com/stainless-sdks/satstream-go/option"
)

// BlockService contains methods and other services that help with interacting with
// the satstream API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockService] method instead.
type BlockService struct {
	Options      []option.RequestOption
	Transactions *BlockTransactionService
}

// NewBlockService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBlockService(opts ...option.RequestOption) (r *BlockService) {
	r = &BlockService{}
	r.Options = opts
	r.Transactions = NewBlockTransactionService(opts...)
	return
}

// Get block info
func (r *BlockService) Get(ctx context.Context, height int64, opts ...option.RequestOption) (res *RpcBlock, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("blocks/%v", height)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RpcBlock struct {
	Bits              string  `json:"bits"`
	Chainwork         string  `json:"chainwork"`
	Confirmations     int64   `json:"confirmations"`
	Difficulty        float64 `json:"difficulty"`
	Hash              string  `json:"hash"`
	Height            int64   `json:"height"`
	Mediantime        int64   `json:"mediantime"`
	Merkleroot        string  `json:"merkleroot"`
	Nonce             int64   `json:"nonce"`
	NTx               int64   `json:"nTx"`
	Previousblockhash string  `json:"previousblockhash"`
	Size              int64   `json:"size"`
	StrippedSize      int64   `json:"strippedSize"`
	Time              int64   `json:"time"`
	// Txns will be stored in a separate collection in the DB
	Tx         []RpcBlockTx `json:"tx"`
	Txns       []string     `json:"txns"`
	Version    int64        `json:"version"`
	VersionHex string       `json:"versionHex"`
	Weight     int64        `json:"weight"`
	JSON       rpcBlockJSON `json:"-"`
}

// rpcBlockJSON contains the JSON metadata for the struct [RpcBlock]
type rpcBlockJSON struct {
	Bits              apijson.Field
	Chainwork         apijson.Field
	Confirmations     apijson.Field
	Difficulty        apijson.Field
	Hash              apijson.Field
	Height            apijson.Field
	Mediantime        apijson.Field
	Merkleroot        apijson.Field
	Nonce             apijson.Field
	NTx               apijson.Field
	Previousblockhash apijson.Field
	Size              apijson.Field
	StrippedSize      apijson.Field
	Time              apijson.Field
	Tx                apijson.Field
	Txns              apijson.Field
	Version           apijson.Field
	VersionHex        apijson.Field
	Weight            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RpcBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rpcBlockJSON) RawJSON() string {
	return r.raw
}

type RpcBlockTx struct {
	Blockheight int64   `json:"blockheight"`
	Fee         float64 `json:"fee"`
	Hash        string  `json:"hash"`
	Hex         string  `json:"hex"`
	Index       int64   `json:"index"`
	Locktime    int64   `json:"locktime"`
	Size        int64   `json:"size"`
	Txid        string  `json:"txid"`
	Version     int64   `json:"version"`
	// Vin will be stored in a separate collection in the DB
	Vin []RpcBlockTxVin `json:"vin"`
	// Vout will be stored in a separate collection in the DB
	Vout   []RpcBlockTxVout `json:"vout"`
	Vsize  int64            `json:"vsize"`
	Weight int64            `json:"weight"`
	JSON   rpcBlockTxJSON   `json:"-"`
}

// rpcBlockTxJSON contains the JSON metadata for the struct [RpcBlockTx]
type rpcBlockTxJSON struct {
	Blockheight apijson.Field
	Fee         apijson.Field
	Hash        apijson.Field
	Hex         apijson.Field
	Index       apijson.Field
	Locktime    apijson.Field
	Size        apijson.Field
	Txid        apijson.Field
	Version     apijson.Field
	Vin         apijson.Field
	Vout        apijson.Field
	Vsize       apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RpcBlockTx) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rpcBlockTxJSON) RawJSON() string {
	return r.raw
}

type RpcBlockTxVin struct {
	Coinbase    string                 `json:"coinbase"`
	Prevout     RpcBlockTxVinPrevout   `json:"prevout"`
	ScriptSig   RpcBlockTxVinScriptSig `json:"scriptSig"`
	Sequence    int64                  `json:"sequence"`
	Txid        string                 `json:"txid"`
	Txinwitness []string               `json:"txinwitness"`
	Vout        int64                  `json:"vout"`
	JSON        rpcBlockTxVinJSON      `json:"-"`
}

// rpcBlockTxVinJSON contains the JSON metadata for the struct [RpcBlockTxVin]
type rpcBlockTxVinJSON struct {
	Coinbase    apijson.Field
	Prevout     apijson.Field
	ScriptSig   apijson.Field
	Sequence    apijson.Field
	Txid        apijson.Field
	Txinwitness apijson.Field
	Vout        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RpcBlockTxVin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rpcBlockTxVinJSON) RawJSON() string {
	return r.raw
}

type RpcBlockTxVinPrevout struct {
	Generated    bool                              `json:"generated"`
	Height       int64                             `json:"height"`
	N            int64                             `json:"n"`
	RuneHoldings []RpcBlockTxVinPrevoutRuneHolding `json:"rune_holdings"`
	ScriptPubKey RpcBlockTxVinPrevoutScriptPubKey  `json:"scriptPubKey"`
	Value        float64                           `json:"value"`
	JSON         rpcBlockTxVinPrevoutJSON          `json:"-"`
}

// rpcBlockTxVinPrevoutJSON contains the JSON metadata for the struct
// [RpcBlockTxVinPrevout]
type rpcBlockTxVinPrevoutJSON struct {
	Generated    apijson.Field
	Height       apijson.Field
	N            apijson.Field
	RuneHoldings apijson.Field
	ScriptPubKey apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RpcBlockTxVinPrevout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rpcBlockTxVinPrevoutJSON) RawJSON() string {
	return r.raw
}

type RpcBlockTxVinPrevoutRuneHolding struct {
	Amount string                              `json:"amount"`
	RuneID string                              `json:"rune_id"`
	JSON   rpcBlockTxVinPrevoutRuneHoldingJSON `json:"-"`
}

// rpcBlockTxVinPrevoutRuneHoldingJSON contains the JSON metadata for the struct
// [RpcBlockTxVinPrevoutRuneHolding]
type rpcBlockTxVinPrevoutRuneHoldingJSON struct {
	Amount      apijson.Field
	RuneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RpcBlockTxVinPrevoutRuneHolding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rpcBlockTxVinPrevoutRuneHoldingJSON) RawJSON() string {
	return r.raw
}

type RpcBlockTxVinPrevoutScriptPubKey struct {
	Address string                               `json:"address"`
	Asm     string                               `json:"asm"`
	Desc    string                               `json:"desc"`
	Hex     string                               `json:"hex"`
	Type    string                               `json:"type"`
	JSON    rpcBlockTxVinPrevoutScriptPubKeyJSON `json:"-"`
}

// rpcBlockTxVinPrevoutScriptPubKeyJSON contains the JSON metadata for the struct
// [RpcBlockTxVinPrevoutScriptPubKey]
type rpcBlockTxVinPrevoutScriptPubKeyJSON struct {
	Address     apijson.Field
	Asm         apijson.Field
	Desc        apijson.Field
	Hex         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RpcBlockTxVinPrevoutScriptPubKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rpcBlockTxVinPrevoutScriptPubKeyJSON) RawJSON() string {
	return r.raw
}

type RpcBlockTxVinScriptSig struct {
	Asm  string                     `json:"asm"`
	Hex  string                     `json:"hex"`
	JSON rpcBlockTxVinScriptSigJSON `json:"-"`
}

// rpcBlockTxVinScriptSigJSON contains the JSON metadata for the struct
// [RpcBlockTxVinScriptSig]
type rpcBlockTxVinScriptSigJSON struct {
	Asm         apijson.Field
	Hex         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RpcBlockTxVinScriptSig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rpcBlockTxVinScriptSigJSON) RawJSON() string {
	return r.raw
}

type RpcBlockTxVout struct {
	N            int64                       `json:"n"`
	RuneHoldings []RpcBlockTxVoutRuneHolding `json:"rune_holdings"`
	ScriptPubKey RpcBlockTxVoutScriptPubKey  `json:"scriptPubKey"`
	Value        float64                     `json:"value"`
	JSON         rpcBlockTxVoutJSON          `json:"-"`
}

// rpcBlockTxVoutJSON contains the JSON metadata for the struct [RpcBlockTxVout]
type rpcBlockTxVoutJSON struct {
	N            apijson.Field
	RuneHoldings apijson.Field
	ScriptPubKey apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RpcBlockTxVout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rpcBlockTxVoutJSON) RawJSON() string {
	return r.raw
}

type RpcBlockTxVoutRuneHolding struct {
	Amount string                        `json:"amount"`
	RuneID string                        `json:"rune_id"`
	JSON   rpcBlockTxVoutRuneHoldingJSON `json:"-"`
}

// rpcBlockTxVoutRuneHoldingJSON contains the JSON metadata for the struct
// [RpcBlockTxVoutRuneHolding]
type rpcBlockTxVoutRuneHoldingJSON struct {
	Amount      apijson.Field
	RuneID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RpcBlockTxVoutRuneHolding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rpcBlockTxVoutRuneHoldingJSON) RawJSON() string {
	return r.raw
}

type RpcBlockTxVoutScriptPubKey struct {
	Address string                         `json:"address"`
	Asm     string                         `json:"asm"`
	Desc    string                         `json:"desc"`
	Hex     string                         `json:"hex"`
	Type    string                         `json:"type"`
	JSON    rpcBlockTxVoutScriptPubKeyJSON `json:"-"`
}

// rpcBlockTxVoutScriptPubKeyJSON contains the JSON metadata for the struct
// [RpcBlockTxVoutScriptPubKey]
type rpcBlockTxVoutScriptPubKeyJSON struct {
	Address     apijson.Field
	Asm         apijson.Field
	Desc        apijson.Field
	Hex         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RpcBlockTxVoutScriptPubKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rpcBlockTxVoutScriptPubKeyJSON) RawJSON() string {
	return r.raw
}
