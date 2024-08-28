// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package satstream

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/satstream-go/internal/requestconfig"
	"github.com/stainless-sdks/satstream-go/option"
)

// FeeService contains methods and other services that help with interacting with
// the satstream API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeeService] method instead.
type FeeService struct {
	Options []option.RequestOption
}

// NewFeeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFeeService(opts ...option.RequestOption) (r *FeeService) {
	r = &FeeService{}
	r.Options = opts
	return
}

// Get recommended fees
func (r *FeeService) List(ctx context.Context, opts ...option.RequestOption) (res *[]int64, err error) {
	opts = append(r.Options[:], opts...)
	path := "fees"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
