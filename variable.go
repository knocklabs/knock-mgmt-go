// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/knocklabs/knock-mgmt-go/internal/apijson"
	"github.com/knocklabs/knock-mgmt-go/internal/apiquery"
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/pagination"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// VariableService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVariableService] method instead.
type VariableService struct {
	Options []option.RequestOption
}

// NewVariableService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVariableService(opts ...option.RequestOption) (r VariableService) {
	r = VariableService{}
	r.Options = opts
	return
}

// Returns a single variable by key with per-environment value overrides.
func (r *VariableService) Get(ctx context.Context, key string, opts ...option.RequestOption) (res *Variable, err error) {
	opts = slices.Concat(r.Options, opts)
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/variables/%s", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of variables. When an environment is specified, returns
// per-environment variables. Otherwise, returns project-scoped variables with
// per-environment overrides.
func (r *VariableService) List(ctx context.Context, query VariableListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Variable], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/variables"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a list of variables. When an environment is specified, returns
// per-environment variables. Otherwise, returns project-scoped variables with
// per-environment overrides.
func (r *VariableService) ListAutoPaging(ctx context.Context, query VariableListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Variable] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// An environment variable object.
type Variable struct {
	// The timestamp of when the variable was created.
	InsertedAt time.Time `json:"inserted_at" api:"required" format:"date-time"`
	// The key of the variable.
	Key string `json:"key" api:"required"`
	// The type of the variable.
	//
	// Any of "public", "secret".
	Type VariableType `json:"type" api:"required"`
	// The timestamp of when the variable was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The description of the variable.
	Description string `json:"description" api:"nullable"`
	// A map of environment slugs to their override values. Only present for
	// project-scoped responses.
	EnvironmentValues map[string]string `json:"environment_values"`
	// The default value of the variable. For secret variables, this is obfuscated.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InsertedAt        respjson.Field
		Key               respjson.Field
		Type              respjson.Field
		UpdatedAt         respjson.Field
		Description       respjson.Field
		EnvironmentValues respjson.Field
		Value             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Variable) RawJSON() string { return r.JSON.raw }
func (r *Variable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the variable.
type VariableType string

const (
	VariableTypePublic VariableType = "public"
	VariableTypeSecret VariableType = "secret"
)

type VariableListParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter variables by type. Supports 'public' or 'secret'.
	//
	// Any of "public", "secret".
	Type VariableListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VariableListParams]'s query parameters as `url.Values`.
func (r VariableListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter variables by type. Supports 'public' or 'secret'.
type VariableListParamsType string

const (
	VariableListParamsTypePublic VariableListParamsType = "public"
	VariableListParamsTypeSecret VariableListParamsType = "secret"
)
