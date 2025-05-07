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

// PartialService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPartialService] method instead.
type PartialService struct {
	Options []option.RequestOption
}

// NewPartialService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPartialService(opts ...option.RequestOption) (r PartialService) {
	r = PartialService{}
	r.Options = opts
	return
}

// Get a partial by its key.
func (r *PartialService) Get(ctx context.Context, partialKey string, query PartialGetParams, opts ...option.RequestOption) (res *Partial, err error) {
	opts = append(r.Options[:], opts...)
	if partialKey == "" {
		err = errors.New("missing required partial_key parameter")
		return
	}
	path := fmt.Sprintf("v1/partials/%s", partialKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List all partials for a given environment.
func (r *PartialService) List(ctx context.Context, query PartialListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Partial], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/partials"
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

// List all partials for a given environment.
func (r *PartialService) ListAutoPaging(ctx context.Context, query PartialListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Partial] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Updates a partial of a given key, or creates a new one if it does not yet exist.
//
// Note: this endpoint only operates on partials in the “development” environment.
func (r *PartialService) Upsert(ctx context.Context, partialKey string, params PartialUpsertParams, opts ...option.RequestOption) (res *PartialUpsertResponse, err error) {
	opts = append(r.Options[:], opts...)
	if partialKey == "" {
		err = errors.New("missing required partial_key parameter")
		return
	}
	path := fmt.Sprintf("v1/partials/%s", partialKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Validates a partial payload without persisting it.
//
// Note: this endpoint only operates on partials in the “development” environment.
func (r *PartialService) Validate(ctx context.Context, partialKey string, params PartialValidateParams, opts ...option.RequestOption) (res *PartialValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if partialKey == "" {
		err = errors.New("missing required partial_key parameter")
		return
	}
	path := fmt.Sprintf("v1/partials/%s/validate", partialKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// A partial is a reusable piece of content that can be used in a template.
type Partial struct {
	// The partial content.
	Content string `json:"content,required"`
	// The timestamp of when the partial was created.
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// The unique key string for the partial object. Must be at minimum 3 characters
	// and at maximum 255 characters in length. Must be in the format of ^[a-z0-9_-]+$.
	Key string `json:"key,required"`
	// A name for the partial. Must be at maximum 255 characters in length.
	Name string `json:"name,required"`
	// The partial type. One of 'html', 'json', 'markdown', 'text'.
	//
	// Any of "html", "text", "json", "markdown".
	Type PartialType `json:"type,required"`
	// The timestamp of when the partial was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Whether the partial and its content are in a valid state.
	Valid bool `json:"valid,required"`
	// An arbitrary string attached to a partial object. Useful for adding notes about
	// the partial for internal purposes. Maximum of 280 characters allowed.
	Description string `json:"description"`
	// The slug of the environment in which the partial exists.
	Environment string `json:"environment"`
	// The name of the icon to be used in the visual editor.
	IconName string `json:"icon_name"`
	// Indicates whether the partial can be used in the visual editor. Only applies to
	// HTML partials.
	VisualBlockEnabled bool `json:"visual_block_enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content            respjson.Field
		InsertedAt         respjson.Field
		Key                respjson.Field
		Name               respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		Valid              respjson.Field
		Description        respjson.Field
		Environment        respjson.Field
		IconName           respjson.Field
		VisualBlockEnabled respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Partial) RawJSON() string { return r.JSON.raw }
func (r *Partial) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The partial type. One of 'html', 'json', 'markdown', 'text'.
type PartialType string

const (
	PartialTypeHTML     PartialType = "html"
	PartialTypeText     PartialType = "text"
	PartialTypeJson     PartialType = "json"
	PartialTypeMarkdown PartialType = "markdown"
)

// Wraps the Partial response under the `partial` key.
type PartialUpsertResponse struct {
	// A partial is a reusable piece of content that can be used in a template.
	Partial Partial `json:"partial,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Partial     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *PartialUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Partial response under the `partial` key.
type PartialValidateResponse struct {
	// A partial is a reusable piece of content that can be used in a template.
	Partial Partial `json:"partial,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Partial     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *PartialValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartialGetParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// Whether to hide uncommitted changes. When true, only committed changes will be
	// returned. When false, both committed and uncommitted changes will be returned.
	HideUncommittedChanges param.Opt[bool] `query:"hide_uncommitted_changes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PartialGetParams]'s query parameters as `url.Values`.
func (r PartialGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PartialListParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Whether to hide uncommitted changes. When true, only committed changes will be
	// returned. When false, both committed and uncommitted changes will be returned.
	HideUncommittedChanges param.Opt[bool] `query:"hide_uncommitted_changes,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PartialListParams]'s query parameters as `url.Values`.
func (r PartialListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PartialUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// A partial object with attributes to update or create a partial.
	Partial PartialUpsertParamsPartial `json:"partial,omitzero,required"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// Whether to commit the resource at the same time as modifying it.
	Commit param.Opt[bool] `query:"commit,omitzero" json:"-"`
	// The message to commit the resource with, only used if `commit` is `true`.
	CommitMessage param.Opt[string] `query:"commit_message,omitzero" json:"-"`
	paramObj
}

func (r PartialUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PartialUpsertParams]'s query parameters as `url.Values`.
func (r PartialUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A partial object with attributes to update or create a partial.
//
// The properties Content, Name, Type are required.
type PartialUpsertParamsPartial struct {
	// The content of the partial.
	Content string `json:"content,required"`
	// The name of the partial.
	Name string `json:"name,required"`
	// The type of the partial.
	//
	// Any of "html", "text", "json", "markdown".
	Type string `json:"type,omitzero,required"`
	// The description of the partial.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the icon to be used in the visual editor. Only relevant when
	// `visual_block_enabled` is `true`.
	IconName param.Opt[string] `json:"icon_name,omitzero"`
	// Indicates whether the partial can be used in the visual editor. Only applies to
	// HTML partials.
	VisualBlockEnabled param.Opt[bool] `json:"visual_block_enabled,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartial) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartial
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartial) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartial](
		"type", "html", "text", "json", "markdown",
	)
}

type PartialValidateParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// A partial object with attributes to update or create a partial.
	Partial PartialValidateParamsPartial `json:"partial,omitzero,required"`
	paramObj
}

func (r PartialValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PartialValidateParams]'s query parameters as `url.Values`.
func (r PartialValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A partial object with attributes to update or create a partial.
//
// The properties Content, Name, Type are required.
type PartialValidateParamsPartial struct {
	// The content of the partial.
	Content string `json:"content,required"`
	// The name of the partial.
	Name string `json:"name,required"`
	// The type of the partial.
	//
	// Any of "html", "text", "json", "markdown".
	Type string `json:"type,omitzero,required"`
	// The description of the partial.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the icon to be used in the visual editor. Only relevant when
	// `visual_block_enabled` is `true`.
	IconName param.Opt[string] `json:"icon_name,omitzero"`
	// Indicates whether the partial can be used in the visual editor. Only applies to
	// HTML partials.
	VisualBlockEnabled param.Opt[bool] `json:"visual_block_enabled,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartial) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartial
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartial) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartial](
		"type", "html", "text", "json", "markdown",
	)
}
