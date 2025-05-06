// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/knock-mapi-go/internal/apijson"
	"github.com/stainless-sdks/knock-mapi-go/internal/apiquery"
	"github.com/stainless-sdks/knock-mapi-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-mapi-go/option"
	"github.com/stainless-sdks/knock-mapi-go/packages/pagination"
	"github.com/stainless-sdks/knock-mapi-go/packages/param"
	"github.com/stainless-sdks/knock-mapi-go/packages/respjson"
)

// EnvironmentService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentService] method instead.
type EnvironmentService struct {
	Options []option.RequestOption
}

// NewEnvironmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnvironmentService(opts ...option.RequestOption) (r EnvironmentService) {
	r = EnvironmentService{}
	r.Options = opts
	return
}

// Returns a single environment by the `environment_slug`.
func (r *EnvironmentService) Get(ctx context.Context, environmentSlug string, opts ...option.RequestOption) (res *Environment, err error) {
	opts = append(r.Options[:], opts...)
	if environmentSlug == "" {
		err = errors.New("missing required environment_slug parameter")
		return
	}
	path := fmt.Sprintf("v1/environments/%s", environmentSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a paginated list of environments. The environments will be returned in
// order of their index, with the `development` environment first.
func (r *EnvironmentService) List(ctx context.Context, query EnvironmentListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Environment], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/environments"
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

// Returns a paginated list of environments. The environments will be returned in
// order of their index, with the `development` environment first.
func (r *EnvironmentService) ListAutoPaging(ctx context.Context, query EnvironmentListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Environment] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// An environment object.
type Environment struct {
	// The timestamp of when the environment was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A human-readable name for the environment. Cannot exceed 255 characters.
	Name string `json:"name,required"`
	// The order of the environment. The lowest number is the first environment, the
	// highest number is the last environment. The order will not always be sequential.
	Order int64 `json:"order,required"`
	// The owner of the environment.
	//
	// Any of "system", "user".
	Owner EnvironmentOwner `json:"owner,required"`
	// A unique slug for the environment. Cannot exceed 255 characters.
	Slug string `json:"slug,required"`
	// The timestamp of when the environment was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The timestamp of when the environment was deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Whether PII data is hidden from the environment. Read more in the
	// [data obfuscation docs](https://docs.knock.app/manage-your-account/data-obfuscation).
	HidePiiData bool `json:"hide_pii_data"`
	// The color of the environment label to display in the dashboard.
	LabelColor string `json:"label_color,nullable"`
	// The timestamp of the most-recent commit in the environment.
	LastCommitAt time.Time `json:"last_commit_at,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		Name         respjson.Field
		Order        respjson.Field
		Owner        respjson.Field
		Slug         respjson.Field
		UpdatedAt    respjson.Field
		DeletedAt    respjson.Field
		HidePiiData  respjson.Field
		LabelColor   respjson.Field
		LastCommitAt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Environment) RawJSON() string { return r.JSON.raw }
func (r *Environment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The owner of the environment.
type EnvironmentOwner string

const (
	EnvironmentOwnerSystem EnvironmentOwner = "system"
	EnvironmentOwnerUser   EnvironmentOwner = "user"
)

type EnvironmentListParams struct {
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EnvironmentListParams]'s query parameters as `url.Values`.
func (r EnvironmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
