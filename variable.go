// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/knock-mapi-go/internal/apijson"
	"github.com/stainless-sdks/knock-mapi-go/internal/apiquery"
	"github.com/stainless-sdks/knock-mapi-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-mapi-go/option"
	"github.com/stainless-sdks/knock-mapi-go/packages/pagination"
	"github.com/stainless-sdks/knock-mapi-go/packages/param"
	"github.com/stainless-sdks/knock-mapi-go/packages/resp"
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

// Returns a paginated list of variables for a given environment.
func (r *VariableService) List(ctx context.Context, query VariableListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Variable], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

// Returns a paginated list of variables for a given environment.
func (r *VariableService) ListAutoPaging(ctx context.Context, query VariableListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Variable] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// An environment variable object.
type Variable struct {
	// The timestamp of when the variable was created.
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// The key of the variable.
	Key string `json:"key,required"`
	// The type of the variable.
	//
	// Any of "public", "secret".
	Type VariableType `json:"type,required"`
	// The timestamp of when the variable was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The value of the variable.
	Value string `json:"value,required"`
	// The description of the variable.
	Description string `json:"description,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		InsertedAt  resp.Field
		Key         resp.Field
		Type        resp.Field
		UpdatedAt   resp.Field
		Value       resp.Field
		Description resp.Field
		ExtraFields map[string]resp.Field
		raw         string
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
	Environment string `query:"environment,required" json:"-"`
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f VariableListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [VariableListParams]'s query parameters as `url.Values`.
func (r VariableListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
