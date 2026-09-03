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
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// PreferenceCenterService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPreferenceCenterService] method instead.
type PreferenceCenterService struct {
	Options []option.RequestOption
}

// NewPreferenceCenterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPreferenceCenterService(opts ...option.RequestOption) (r PreferenceCenterService) {
	r = PreferenceCenterService{}
	r.Options = opts
	return
}

// Returns the preference center configuration for the given environment.
func (r *PreferenceCenterService) Get(ctx context.Context, query PreferenceCenterGetParams, opts ...option.RequestOption) (res *PreferenceCenterGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/preference_center"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Resets the preference center configuration for the given environment to the
// built-in default content. The `enabled` flag is preserved.
func (r *PreferenceCenterService) Reset(ctx context.Context, body PreferenceCenterResetParams, opts ...option.RequestOption) (res *PreferenceCenterResetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/preference_center/reset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Creates or updates the preference center configuration for the given
// environment.
func (r *PreferenceCenterService) Upsert(ctx context.Context, params PreferenceCenterUpsertParams, opts ...option.RequestOption) (res *PreferenceCenterUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/preference_center"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// The preference center configuration for a single environment.
type PreferenceCenterGetResponse struct {
	// The preference center configuration document.
	Config any `json:"config"`
	// Whether the preference center is enabled for recipients.
	Enabled bool `json:"enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreferenceCenterGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PreferenceCenterGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The preference center configuration for a single environment.
type PreferenceCenterResetResponse struct {
	// The preference center configuration document.
	Config any `json:"config"`
	// Whether the preference center is enabled for recipients.
	Enabled bool `json:"enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreferenceCenterResetResponse) RawJSON() string { return r.JSON.raw }
func (r *PreferenceCenterResetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The preference center configuration for a single environment.
type PreferenceCenterUpsertResponse struct {
	// The preference center configuration document.
	Config any `json:"config"`
	// Whether the preference center is enabled for recipients.
	Enabled bool `json:"enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreferenceCenterUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *PreferenceCenterUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreferenceCenterGetParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [PreferenceCenterGetParams]'s query parameters as
// `url.Values`.
func (r PreferenceCenterGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PreferenceCenterResetParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [PreferenceCenterResetParams]'s query parameters as
// `url.Values`.
func (r PreferenceCenterResetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PreferenceCenterUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The preference center configuration document.
	Config any `json:"config,omitzero" api:"required"`
	// Whether the preference center is enabled for recipients.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r PreferenceCenterUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow PreferenceCenterUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreferenceCenterUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PreferenceCenterUpsertParams]'s query parameters as
// `url.Values`.
func (r PreferenceCenterUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
