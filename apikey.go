// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/knocklabs/knock-mgmt-go/internal/apijson"
	"github.com/knocklabs/knock-mgmt-go/internal/apiquery"
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// APIKeyService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyService] method instead.
type APIKeyService struct {
	Options []option.RequestOption
}

// NewAPIKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIKeyService(opts ...option.RequestOption) (r APIKeyService) {
	r = APIKeyService{}
	r.Options = opts
	return
}

// Given an authenticated service token and an environment, will exchange the
// service token for a secret API key that can be used to make requests to the
// public API.
func (r *APIKeyService) Exchange(ctx context.Context, body APIKeyExchangeParams, opts ...option.RequestOption) (res *APIKeyExchangeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/api_keys/exchange"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns an API key that can be used to make requests to the public API.
type APIKeyExchangeResponse struct {
	// The secret API key exchanged from the service token.
	APIKey string `json:"api_key,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyExchangeResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyExchangeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyExchangeParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	paramObj
}

// URLQuery serializes [APIKeyExchangeParams]'s query parameters as `url.Values`.
func (r APIKeyExchangeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
