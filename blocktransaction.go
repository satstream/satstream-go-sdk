// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package satstream

import (
	"context"
	"fmt"
	"net/http"

	"github.com/satstream/satstream-go-sdk/internal/apijson"
	"github.com/satstream/satstream-go-sdk/internal/requestconfig"
	"github.com/satstream/satstream-go-sdk/option"
)

// BlockTransactionService contains methods and other services that help with
// interacting with the satstream API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBlockTransactionService] method instead.
type BlockTransactionService struct {
	Options []option.RequestOption
}

// NewBlockTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBlockTransactionService(opts ...option.RequestOption) (r *BlockTransactionService) {
	r = &BlockTransactionService{}
	r.Options = opts
	return
}

// Get block transactions
func (r *BlockTransactionService) List(ctx context.Context, height int64, opts ...option.RequestOption) (res *[][]StoreTransactionDocument, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("blocks/%v/transactions", height)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StoreTransactionDocument struct {
	Address     string                       `json:"address"`
	Blockheight int64                        `json:"blockheight"`
	Fee         float64                      `json:"fee"`
	Hash        string                       `json:"hash"`
	Hex         string                       `json:"hex"`
	Index       int64                        `json:"index"`
	Locktime    int64                        `json:"locktime"`
	Size        int64                        `json:"size"`
	Txid        string                       `json:"txid"`
	Version     int64                        `json:"version"`
	Vsize       int64                        `json:"vsize"`
	Weight      int64                        `json:"weight"`
	JSON        storeTransactionDocumentJSON `json:"-"`
}

// storeTransactionDocumentJSON contains the JSON metadata for the struct
// [StoreTransactionDocument]
type storeTransactionDocumentJSON struct {
	Address     apijson.Field
	Blockheight apijson.Field
	Fee         apijson.Field
	Hash        apijson.Field
	Hex         apijson.Field
	Index       apijson.Field
	Locktime    apijson.Field
	Size        apijson.Field
	Txid        apijson.Field
	Version     apijson.Field
	Vsize       apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StoreTransactionDocument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r storeTransactionDocumentJSON) RawJSON() string {
	return r.raw
}
