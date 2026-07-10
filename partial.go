// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/knocklabs/knock-mgmt-go/internal/apijson"
	"github.com/knocklabs/knock-mgmt-go/internal/apiquery"
	"github.com/knocklabs/knock-mgmt-go/internal/paramutil"
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/pagination"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
	"github.com/knocklabs/knock-mgmt-go/shared"
)

// Partials allow you to reuse content across templates.
//
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
	opts = slices.Concat(r.Options, opts)
	if partialKey == "" {
		err = errors.New("missing required partial_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/partials/%s", partialKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List all partials for a given environment.
func (r *PartialService) List(ctx context.Context, query PartialListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Partial], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	if partialKey == "" {
		err = errors.New("missing required partial_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/partials/%s", partialKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Validates a partial payload without persisting it.
//
// Note: this endpoint only operates on partials in the “development” environment.
func (r *PartialService) Validate(ctx context.Context, partialKey string, params PartialValidateParams, opts ...option.RequestOption) (res *PartialValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if partialKey == "" {
		err = errors.New("missing required partial_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/partials/%s/validate", partialKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A partial is a reusable piece of content that can be used in a template.
type Partial struct {
	// The partial content.
	Content string `json:"content" api:"required"`
	// The timestamp of when the partial was created.
	InsertedAt time.Time `json:"inserted_at" api:"required" format:"date-time"`
	// The unique key string for the partial object. Must be at minimum 3 characters
	// and at maximum 255 characters in length. Must be in the format of ^[a-z0-9_-]+$.
	Key string `json:"key" api:"required"`
	// A name for the partial. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// The partial type. One of 'html', 'json', 'markdown', 'text'.
	//
	// Any of "html", "text", "json", "markdown".
	Type PartialType `json:"type" api:"required"`
	// The timestamp of when the partial was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Whether the partial and its content are in a valid state.
	Valid bool `json:"valid" api:"required"`
	// An arbitrary string attached to a partial object. Useful for adding notes about
	// the partial for internal purposes. Maximum of 280 characters allowed.
	Description string `json:"description"`
	// The slug of the environment in which the partial exists.
	Environment string `json:"environment"`
	// The name of the icon to be used in the visual editor.
	IconName string `json:"icon_name"`
	// The field types available for the partial.
	InputSchema []PartialInputSchemaUnion `json:"input_schema"`
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
		InputSchema        respjson.Field
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

// PartialInputSchemaUnion contains all possible properties and values from
// [PartialInputSchemaMessageTypeListField], [shared.MessageTypeSelectField],
// [shared.MessageTypeBooleanField], [shared.MessageTypeJsonField],
// [PartialInputSchemaMessageTypeNumberField], [shared.MessageTypeTextField],
// [shared.MessageTypeImageField], [PartialInputSchemaMessageTypeColorField],
// [shared.MessageTypeURLField], [shared.MessageTypeMarkdownField],
// [shared.MessageTypeMultiSelectField], [shared.MessageTypeButtonField],
// [shared.MessageTypeTextareaField].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PartialInputSchemaUnion struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Type  string `json:"type"`
	// This field is a union of [PartialInputSchemaMessageTypeListFieldSettings],
	// [shared.MessageTypeSelectFieldSettings],
	// [shared.MessageTypeBooleanFieldSettings], [shared.MessageTypeJsonFieldSettings],
	// [PartialInputSchemaMessageTypeNumberFieldSettings],
	// [shared.MessageTypeTextFieldSettings], [shared.MessageTypeImageFieldSettings],
	// [PartialInputSchemaMessageTypeColorFieldSettings],
	// [shared.MessageTypeURLFieldSettings], [shared.MessageTypeMarkdownFieldSettings],
	// [shared.MessageTypeMultiSelectFieldSettings],
	// [shared.MessageTypeButtonFieldSettings],
	// [shared.MessageTypeTextareaFieldSettings]
	Settings PartialInputSchemaUnionSettings `json:"settings"`
	// This field is from variant [shared.MessageTypeImageField].
	Action shared.MessageTypeTextField `json:"action"`
	// This field is from variant [shared.MessageTypeImageField].
	Alt shared.MessageTypeTextField `json:"alt"`
	// This field is from variant [shared.MessageTypeImageField].
	URL shared.MessageTypeURLField `json:"url"`
	// This field is from variant [shared.MessageTypeButtonField].
	Text shared.MessageTypeTextField `json:"text"`
	JSON struct {
		Key      respjson.Field
		Label    respjson.Field
		Type     respjson.Field
		Settings respjson.Field
		Action   respjson.Field
		Alt      respjson.Field
		URL      respjson.Field
		Text     respjson.Field
		raw      string
	} `json:"-"`
}

func (u PartialInputSchemaUnion) AsMessageTypeListField() (v PartialInputSchemaMessageTypeListField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeSelectField() (v shared.MessageTypeSelectField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeBooleanField() (v shared.MessageTypeBooleanField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeJsonField() (v shared.MessageTypeJsonField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeNumberField() (v PartialInputSchemaMessageTypeNumberField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeTextField() (v shared.MessageTypeTextField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeImageField() (v shared.MessageTypeImageField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeColorField() (v PartialInputSchemaMessageTypeColorField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeURLField() (v shared.MessageTypeURLField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeMarkdownField() (v shared.MessageTypeMarkdownField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeMultiSelectField() (v shared.MessageTypeMultiSelectField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeButtonField() (v shared.MessageTypeButtonField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeTextareaField() (v shared.MessageTypeTextareaField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PartialInputSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *PartialInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PartialInputSchemaUnionSettings is an implicit subunion of
// [PartialInputSchemaUnion]. PartialInputSchemaUnionSettings provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PartialInputSchemaUnion].
type PartialInputSchemaUnionSettings struct {
	// This field is a union of [[]any], [string], [bool], [any], [float64], [string],
	// [string], [string], [string], [[]string], [string]
	Default     PartialInputSchemaUnionSettingsDefault `json:"default"`
	Description string                                 `json:"description"`
	// This field is from variant [PartialInputSchemaMessageTypeListFieldSettings].
	ItemSchema  any    `json:"item_schema"`
	Placeholder string `json:"placeholder"`
	Required    bool   `json:"required"`
	// This field is a union of [[]shared.MessageTypeSelectFieldSettingsOption],
	// [[]shared.MessageTypeMultiSelectFieldSettingsOption]
	Options PartialInputSchemaUnionSettingsOptions `json:"options"`
	// This field is from variant [shared.MessageTypeJsonFieldSettings].
	Schema any `json:"schema"`
	// This field is from variant [PartialInputSchemaMessageTypeNumberFieldSettings].
	Max float64 `json:"max"`
	// This field is from variant [PartialInputSchemaMessageTypeNumberFieldSettings].
	Min float64 `json:"min"`
	// This field is from variant [PartialInputSchemaMessageTypeNumberFieldSettings].
	UnitLabel string `json:"unit_label"`
	MaxLength int64  `json:"max_length"`
	MinLength int64  `json:"min_length"`
	JSON      struct {
		Default     respjson.Field
		Description respjson.Field
		ItemSchema  respjson.Field
		Placeholder respjson.Field
		Required    respjson.Field
		Options     respjson.Field
		Schema      respjson.Field
		Max         respjson.Field
		Min         respjson.Field
		UnitLabel   respjson.Field
		MaxLength   respjson.Field
		MinLength   respjson.Field
		raw         string
	} `json:"-"`
}

func (r *PartialInputSchemaUnionSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PartialInputSchemaUnionSettingsDefault is an implicit subunion of
// [PartialInputSchemaUnion]. PartialInputSchemaUnionSettingsDefault provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PartialInputSchemaUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfBool OfMessageTypeJsonFieldSettingsDefault
// OfFloat OfStringArray]
type PartialInputSchemaUnionSettingsDefault struct {
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMessageTypeJsonFieldSettingsDefault any `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfAnyArray                            respjson.Field
		OfString                              respjson.Field
		OfBool                                respjson.Field
		OfMessageTypeJsonFieldSettingsDefault respjson.Field
		OfFloat                               respjson.Field
		OfStringArray                         respjson.Field
		raw                                   string
	} `json:"-"`
}

func (r *PartialInputSchemaUnionSettingsDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PartialInputSchemaUnionSettingsOptions is an implicit subunion of
// [PartialInputSchemaUnion]. PartialInputSchemaUnionSettingsOptions provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PartialInputSchemaUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfMessageTypeSelectFieldSettingsOptions
// OfMessageTypeMultiSelectFieldSettingsOptions]
type PartialInputSchemaUnionSettingsOptions struct {
	// This field will be present if the value is a
	// [[]shared.MessageTypeSelectFieldSettingsOption] instead of an object.
	OfMessageTypeSelectFieldSettingsOptions []shared.MessageTypeSelectFieldSettingsOption `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.MessageTypeMultiSelectFieldSettingsOption] instead of an object.
	OfMessageTypeMultiSelectFieldSettingsOptions []shared.MessageTypeMultiSelectFieldSettingsOption `json:",inline"`
	JSON                                         struct {
		OfMessageTypeSelectFieldSettingsOptions      respjson.Field
		OfMessageTypeMultiSelectFieldSettingsOptions respjson.Field
		raw                                          string
	} `json:"-"`
}

func (r *PartialInputSchemaUnionSettingsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A list field used in a message type.
type PartialInputSchemaMessageTypeListField struct {
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The label of the field.
	Label string `json:"label" api:"required"`
	// The type of the field.
	//
	// Any of "list".
	Type string `json:"type" api:"required"`
	// Settings for the list field.
	Settings PartialInputSchemaMessageTypeListFieldSettings `json:"settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Label       respjson.Field
		Type        respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeListField) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeListField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the list field.
type PartialInputSchemaMessageTypeListFieldSettings struct {
	// The default value of the list field.
	Default     []any  `json:"default" api:"nullable"`
	Description string `json:"description" api:"nullable"`
	// A JSON schema used to validate the structure of each item in the list. Must be a
	// valid JSON schema.
	ItemSchema  any    `json:"item_schema" api:"nullable"`
	Placeholder string `json:"placeholder" api:"nullable"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		ItemSchema  respjson.Field
		Placeholder respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeListFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeListFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A numeric field used in a message type or partial input schema, with optional
// min/max bounds and a unit label for display.
type PartialInputSchemaMessageTypeNumberField struct {
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The label of the field.
	Label string `json:"label" api:"required"`
	// The type of the field.
	//
	// Any of "number".
	Type string `json:"type" api:"required"`
	// Settings for the number field.
	Settings PartialInputSchemaMessageTypeNumberFieldSettings `json:"settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Label       respjson.Field
		Type        respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeNumberField) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeNumberField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the number field.
type PartialInputSchemaMessageTypeNumberFieldSettings struct {
	// The default numeric value.
	Default     float64 `json:"default" api:"nullable"`
	Description string  `json:"description" api:"nullable"`
	// Optional inclusive maximum allowed value.
	Max float64 `json:"max" api:"nullable"`
	// Optional inclusive minimum allowed value.
	Min         float64 `json:"min" api:"nullable"`
	Placeholder string  `json:"placeholder" api:"nullable"`
	// Whether the field is required.
	Required bool `json:"required"`
	// Optional short label shown after the input (e.g. px, kg).
	UnitLabel string `json:"unit_label" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Max         respjson.Field
		Min         respjson.Field
		Placeholder respjson.Field
		Required    respjson.Field
		UnitLabel   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeNumberFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeNumberFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A hex color field (#RGB or #RRGGBB) used in a message type or partial input
// schema.
type PartialInputSchemaMessageTypeColorField struct {
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The label of the field.
	Label string `json:"label" api:"required"`
	// The type of the field.
	//
	// Any of "color".
	Type string `json:"type" api:"required"`
	// Settings for the color field.
	Settings PartialInputSchemaMessageTypeColorFieldSettings `json:"settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Label       respjson.Field
		Type        respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeColorField) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeColorField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the color field.
type PartialInputSchemaMessageTypeColorFieldSettings struct {
	// The default hex color value.
	Default     string `json:"default" api:"nullable"`
	Description string `json:"description" api:"nullable"`
	Placeholder string `json:"placeholder" api:"nullable"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Placeholder respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeColorFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeColorFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Partial response under the `partial` key.
type PartialUpsertResponse struct {
	// A partial is a reusable piece of content that can be used in a template.
	Partial Partial `json:"partial" api:"required"`
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
	Partial Partial `json:"partial" api:"required"`
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
	Environment string `query:"environment" api:"required" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// Whether to hide uncommitted changes. When true, only committed changes will be
	// returned. When false, both committed and uncommitted changes will be returned.
	HideUncommittedChanges param.Opt[bool] `query:"hide_uncommitted_changes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PartialGetParams]'s query parameters as `url.Values`.
func (r PartialGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PartialListParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PartialUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A partial object with attributes to update or create a partial.
	Partial PartialUpsertParamsPartial `json:"partial,omitzero" api:"required"`
	// When used with commit, creates a new version with identical content and commits
	// it if there are no unpublished changes.
	AllowEmpty param.Opt[bool] `query:"allow_empty,omitzero" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// Whether to commit the resource at the same time as modifying it.
	Commit param.Opt[bool] `query:"commit,omitzero" json:"-"`
	// The message to commit the resource with, only used if `commit` is `true`.
	CommitMessage param.Opt[string] `query:"commit_message,omitzero" json:"-"`
	// When set to true, forces the upsert to override existing content regardless of
	// environment restrictions. This bypasses the development-only environment check
	// and origin environment checks.
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A partial object with attributes to update or create a partial.
//
// The properties Content, Name, Type are required.
type PartialUpsertParamsPartial struct {
	// The partial content.
	Content string `json:"content" api:"required"`
	// A name for the partial. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// The partial type. One of 'html', 'json', 'markdown', 'text'.
	//
	// Any of "html", "text", "json", "markdown".
	Type string `json:"type,omitzero" api:"required"`
	// An arbitrary string attached to a partial object. Useful for adding notes about
	// the partial for internal purposes. Maximum of 280 characters allowed.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the icon to be used in the visual editor.
	IconName param.Opt[string] `json:"icon_name,omitzero"`
	// Indicates whether the partial can be used in the visual editor. Only applies to
	// HTML partials.
	VisualBlockEnabled param.Opt[bool] `json:"visual_block_enabled,omitzero"`
	// The field types available for the partial.
	InputSchema []PartialUpsertParamsPartialInputSchemaUnion `json:"input_schema,omitzero"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PartialUpsertParamsPartialInputSchemaUnion struct {
	OfMessageTypeListField        *PartialUpsertParamsPartialInputSchemaMessageTypeListField   `json:",omitzero,inline"`
	OfMessageTypeSelectField      *shared.MessageTypeSelectFieldParam                          `json:",omitzero,inline"`
	OfMessageTypeBooleanField     *shared.MessageTypeBooleanFieldParam                         `json:",omitzero,inline"`
	OfMessageTypeJsonField        *shared.MessageTypeJsonFieldParam                            `json:",omitzero,inline"`
	OfMessageTypeNumberField      *PartialUpsertParamsPartialInputSchemaMessageTypeNumberField `json:",omitzero,inline"`
	OfMessageTypeTextField        *shared.MessageTypeTextFieldParam                            `json:",omitzero,inline"`
	OfMessageTypeImageField       *shared.MessageTypeImageFieldParam                           `json:",omitzero,inline"`
	OfMessageTypeColorField       *PartialUpsertParamsPartialInputSchemaMessageTypeColorField  `json:",omitzero,inline"`
	OfMessageTypeURLField         *shared.MessageTypeURLFieldParam                             `json:",omitzero,inline"`
	OfMessageTypeMarkdownField    *shared.MessageTypeMarkdownFieldParam                        `json:",omitzero,inline"`
	OfMessageTypeMultiSelectField *shared.MessageTypeMultiSelectFieldParam                     `json:",omitzero,inline"`
	OfMessageTypeButtonField      *shared.MessageTypeButtonFieldParam                          `json:",omitzero,inline"`
	OfMessageTypeTextareaField    *shared.MessageTypeTextareaFieldParam                        `json:",omitzero,inline"`
	paramUnion
}

func (u PartialUpsertParamsPartialInputSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMessageTypeListField,
		u.OfMessageTypeSelectField,
		u.OfMessageTypeBooleanField,
		u.OfMessageTypeJsonField,
		u.OfMessageTypeNumberField,
		u.OfMessageTypeTextField,
		u.OfMessageTypeImageField,
		u.OfMessageTypeColorField,
		u.OfMessageTypeURLField,
		u.OfMessageTypeMarkdownField,
		u.OfMessageTypeMultiSelectField,
		u.OfMessageTypeButtonField,
		u.OfMessageTypeTextareaField)
}
func (u *PartialUpsertParamsPartialInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PartialUpsertParamsPartialInputSchemaUnion) asAny() any {
	if !param.IsOmitted(u.OfMessageTypeListField) {
		return u.OfMessageTypeListField
	} else if !param.IsOmitted(u.OfMessageTypeSelectField) {
		return u.OfMessageTypeSelectField
	} else if !param.IsOmitted(u.OfMessageTypeBooleanField) {
		return u.OfMessageTypeBooleanField
	} else if !param.IsOmitted(u.OfMessageTypeJsonField) {
		return u.OfMessageTypeJsonField
	} else if !param.IsOmitted(u.OfMessageTypeNumberField) {
		return u.OfMessageTypeNumberField
	} else if !param.IsOmitted(u.OfMessageTypeTextField) {
		return u.OfMessageTypeTextField
	} else if !param.IsOmitted(u.OfMessageTypeImageField) {
		return u.OfMessageTypeImageField
	} else if !param.IsOmitted(u.OfMessageTypeColorField) {
		return u.OfMessageTypeColorField
	} else if !param.IsOmitted(u.OfMessageTypeURLField) {
		return u.OfMessageTypeURLField
	} else if !param.IsOmitted(u.OfMessageTypeMarkdownField) {
		return u.OfMessageTypeMarkdownField
	} else if !param.IsOmitted(u.OfMessageTypeMultiSelectField) {
		return u.OfMessageTypeMultiSelectField
	} else if !param.IsOmitted(u.OfMessageTypeButtonField) {
		return u.OfMessageTypeButtonField
	} else if !param.IsOmitted(u.OfMessageTypeTextareaField) {
		return u.OfMessageTypeTextareaField
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetAlt() *shared.MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.Alt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetURL() *shared.MessageTypeURLFieldParam {
	if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetText() *shared.MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeButtonField; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetKey() *string {
	if vt := u.OfMessageTypeListField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeSelectField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeBooleanField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeJsonField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeNumberField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeTextField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeColorField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeURLField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeTextareaField; vt != nil {
		return (*string)(&vt.Key)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetLabel() *string {
	if vt := u.OfMessageTypeListField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeSelectField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeBooleanField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeJsonField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeNumberField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeTextField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeImageField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeColorField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeURLField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeButtonField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeTextareaField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetType() *string {
	if vt := u.OfMessageTypeListField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeSelectField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeBooleanField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeJsonField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeNumberField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeTextField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeColorField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeURLField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeTextareaField; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u PartialUpsertParamsPartialInputSchemaUnion) GetSettings() (res partialUpsertParamsPartialInputSchemaUnionSettings) {
	if vt := u.OfMessageTypeListField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeSelectField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeBooleanField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeJsonField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeNumberField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeTextField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeColorField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeURLField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeTextareaField; vt != nil {
		res.any = &vt.Settings
	}
	return
}

// Can have the runtime types
// [*PartialUpsertParamsPartialInputSchemaMessageTypeListFieldSettings],
// [*shared.MessageTypeSelectFieldSettingsParam],
// [*shared.MessageTypeBooleanFieldSettingsParam],
// [*shared.MessageTypeJsonFieldSettingsParam],
// [*PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings],
// [*shared.MessageTypeTextFieldSettingsParam],
// [*shared.MessageTypeImageFieldSettingsParam],
// [*PartialUpsertParamsPartialInputSchemaMessageTypeColorFieldSettings],
// [*shared.MessageTypeURLFieldSettingsParam],
// [*shared.MessageTypeMarkdownFieldSettingsParam],
// [*shared.MessageTypeMultiSelectFieldSettingsParam],
// [*shared.MessageTypeButtonFieldSettingsParam],
// [*shared.MessageTypeTextareaFieldSettingsParam]
type partialUpsertParamsPartialInputSchemaUnionSettings struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeListFieldSettings:
//	case *shared.MessageTypeSelectFieldSettingsParam:
//	case *shared.MessageTypeBooleanFieldSettingsParam:
//	case *shared.MessageTypeJsonFieldSettingsParam:
//	case *knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings:
//	case *shared.MessageTypeTextFieldSettingsParam:
//	case *shared.MessageTypeImageFieldSettingsParam:
//	case *knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeColorFieldSettings:
//	case *shared.MessageTypeURLFieldSettingsParam:
//	case *shared.MessageTypeMarkdownFieldSettingsParam:
//	case *shared.MessageTypeMultiSelectFieldSettingsParam:
//	case *shared.MessageTypeButtonFieldSettingsParam:
//	case *shared.MessageTypeTextareaFieldSettingsParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u partialUpsertParamsPartialInputSchemaUnionSettings) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetItemSchema() *any {
	switch vt := u.any.(type) {
	case *PartialUpsertParamsPartialInputSchemaMessageTypeListFieldSettings:
		return &vt.ItemSchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetSchema() *any {
	switch vt := u.any.(type) {
	case *shared.MessageTypeJsonFieldSettingsParam:
		return &vt.Schema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetMax() *float64 {
	switch vt := u.any.(type) {
	case *PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		return paramutil.AddrIfPresent(vt.Max)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetMin() *float64 {
	switch vt := u.any.(type) {
	case *PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		return paramutil.AddrIfPresent(vt.Min)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetUnitLabel() *string {
	switch vt := u.any.(type) {
	case *PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		return paramutil.AddrIfPresent(vt.UnitLabel)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetDescription() *string {
	switch vt := u.any.(type) {
	case *PartialUpsertParamsPartialInputSchemaMessageTypeListFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeBooleanFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeJsonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeImageFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeColorFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeURLFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeMarkdownFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeButtonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetPlaceholder() *string {
	switch vt := u.any.(type) {
	case *PartialUpsertParamsPartialInputSchemaMessageTypeListFieldSettings:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeBooleanFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeJsonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeImageFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeColorFieldSettings:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeURLFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeMarkdownFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeButtonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetRequired() *bool {
	switch vt := u.any.(type) {
	case *PartialUpsertParamsPartialInputSchemaMessageTypeListFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeBooleanFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeJsonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeImageFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeColorFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeURLFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeMarkdownFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeButtonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetMaxLength() *int64 {
	switch vt := u.any.(type) {
	case *shared.MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MaxLength)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MaxLength)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetMinLength() *int64 {
	switch vt := u.any.(type) {
	case *shared.MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MinLength)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MinLength)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetDefault() (res partialUpsertParamsPartialInputSchemaUnionSettingsDefault) {
	switch vt := u.any.(type) {
	case *PartialUpsertParamsPartialInputSchemaMessageTypeListFieldSettings:
		res.any = &vt.Default
	case *shared.MessageTypeSelectFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeBooleanFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeJsonFieldSettingsParam:
		res.any = &vt.Default
	case *PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeTextFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeColorFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeURLFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeMarkdownFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		res.any = &vt.Default
	case *shared.MessageTypeTextareaFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	}
	return res
}

// Can have the runtime types [*[]any], [*string], [*bool], [*any], [*float64],
// [\*[]string]
type partialUpsertParamsPartialInputSchemaUnionSettingsDefault struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]any:
//	case *string:
//	case *bool:
//	case *any:
//	case *float64:
//	case *[]string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u partialUpsertParamsPartialInputSchemaUnionSettingsDefault) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetOptions() (res partialUpsertParamsPartialInputSchemaUnionSettingsOptions) {
	switch vt := u.any.(type) {
	case *shared.MessageTypeSelectFieldSettingsParam:
		res.any = &vt.Options
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		res.any = &vt.Options
	}
	return res
}

// Can have the runtime types
// [_[]shared.MessageTypeSelectFieldSettingsOptionParam],
// [_[]shared.MessageTypeMultiSelectFieldSettingsOptionParam]
type partialUpsertParamsPartialInputSchemaUnionSettingsOptions struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.MessageTypeSelectFieldSettingsOptionParam:
//	case *[]shared.MessageTypeMultiSelectFieldSettingsOptionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u partialUpsertParamsPartialInputSchemaUnionSettingsOptions) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's Action property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetAction() *shared.MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.Action
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		return &vt.Action
	}
	return nil
}

// A list field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialUpsertParamsPartialInputSchemaMessageTypeListField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The type of the field.
	//
	// Any of "list".
	Type string `json:"type,omitzero" api:"required"`
	// Settings for the list field.
	Settings PartialUpsertParamsPartialInputSchemaMessageTypeListFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeListField) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeListField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeListField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartialInputSchemaMessageTypeListField](
		"type", "list",
	)
}

// Settings for the list field.
type PartialUpsertParamsPartialInputSchemaMessageTypeListFieldSettings struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	// The default value of the list field.
	Default []any `json:"default,omitzero"`
	// A JSON schema used to validate the structure of each item in the list. Must be a
	// valid JSON schema.
	ItemSchema any `json:"item_schema,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeListFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeListFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeListFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A numeric field used in a message type or partial input schema, with optional
// min/max bounds and a unit label for display.
//
// The properties Key, Label, Type are required.
type PartialUpsertParamsPartialInputSchemaMessageTypeNumberField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The type of the field.
	//
	// Any of "number".
	Type string `json:"type,omitzero" api:"required"`
	// Settings for the number field.
	Settings PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeNumberField) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeNumberField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeNumberField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartialInputSchemaMessageTypeNumberField](
		"type", "number",
	)
}

// Settings for the number field.
type PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings struct {
	// The default numeric value.
	Default     param.Opt[float64] `json:"default,omitzero"`
	Description param.Opt[string]  `json:"description,omitzero"`
	// Optional inclusive maximum allowed value.
	Max param.Opt[float64] `json:"max,omitzero"`
	// Optional inclusive minimum allowed value.
	Min         param.Opt[float64] `json:"min,omitzero"`
	Placeholder param.Opt[string]  `json:"placeholder,omitzero"`
	// Optional short label shown after the input (e.g. px, kg).
	UnitLabel param.Opt[string] `json:"unit_label,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeNumberFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A hex color field (#RGB or #RRGGBB) used in a message type or partial input
// schema.
//
// The properties Key, Label, Type are required.
type PartialUpsertParamsPartialInputSchemaMessageTypeColorField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The type of the field.
	//
	// Any of "color".
	Type string `json:"type,omitzero" api:"required"`
	// Settings for the color field.
	Settings PartialUpsertParamsPartialInputSchemaMessageTypeColorFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeColorField) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeColorField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeColorField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartialInputSchemaMessageTypeColorField](
		"type", "color",
	)
}

// Settings for the color field.
type PartialUpsertParamsPartialInputSchemaMessageTypeColorFieldSettings struct {
	// The default hex color value.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeColorFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeColorFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeColorFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartialValidateParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A partial object with attributes to update or create a partial.
	Partial PartialValidateParamsPartial `json:"partial,omitzero" api:"required"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A partial object with attributes to update or create a partial.
//
// The properties Content, Name, Type are required.
type PartialValidateParamsPartial struct {
	// The partial content.
	Content string `json:"content" api:"required"`
	// A name for the partial. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// The partial type. One of 'html', 'json', 'markdown', 'text'.
	//
	// Any of "html", "text", "json", "markdown".
	Type string `json:"type,omitzero" api:"required"`
	// An arbitrary string attached to a partial object. Useful for adding notes about
	// the partial for internal purposes. Maximum of 280 characters allowed.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the icon to be used in the visual editor.
	IconName param.Opt[string] `json:"icon_name,omitzero"`
	// Indicates whether the partial can be used in the visual editor. Only applies to
	// HTML partials.
	VisualBlockEnabled param.Opt[bool] `json:"visual_block_enabled,omitzero"`
	// The field types available for the partial.
	InputSchema []PartialValidateParamsPartialInputSchemaUnion `json:"input_schema,omitzero"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PartialValidateParamsPartialInputSchemaUnion struct {
	OfMessageTypeListField        *PartialValidateParamsPartialInputSchemaMessageTypeListField   `json:",omitzero,inline"`
	OfMessageTypeSelectField      *shared.MessageTypeSelectFieldParam                            `json:",omitzero,inline"`
	OfMessageTypeBooleanField     *shared.MessageTypeBooleanFieldParam                           `json:",omitzero,inline"`
	OfMessageTypeJsonField        *shared.MessageTypeJsonFieldParam                              `json:",omitzero,inline"`
	OfMessageTypeNumberField      *PartialValidateParamsPartialInputSchemaMessageTypeNumberField `json:",omitzero,inline"`
	OfMessageTypeTextField        *shared.MessageTypeTextFieldParam                              `json:",omitzero,inline"`
	OfMessageTypeImageField       *shared.MessageTypeImageFieldParam                             `json:",omitzero,inline"`
	OfMessageTypeColorField       *PartialValidateParamsPartialInputSchemaMessageTypeColorField  `json:",omitzero,inline"`
	OfMessageTypeURLField         *shared.MessageTypeURLFieldParam                               `json:",omitzero,inline"`
	OfMessageTypeMarkdownField    *shared.MessageTypeMarkdownFieldParam                          `json:",omitzero,inline"`
	OfMessageTypeMultiSelectField *shared.MessageTypeMultiSelectFieldParam                       `json:",omitzero,inline"`
	OfMessageTypeButtonField      *shared.MessageTypeButtonFieldParam                            `json:",omitzero,inline"`
	OfMessageTypeTextareaField    *shared.MessageTypeTextareaFieldParam                          `json:",omitzero,inline"`
	paramUnion
}

func (u PartialValidateParamsPartialInputSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMessageTypeListField,
		u.OfMessageTypeSelectField,
		u.OfMessageTypeBooleanField,
		u.OfMessageTypeJsonField,
		u.OfMessageTypeNumberField,
		u.OfMessageTypeTextField,
		u.OfMessageTypeImageField,
		u.OfMessageTypeColorField,
		u.OfMessageTypeURLField,
		u.OfMessageTypeMarkdownField,
		u.OfMessageTypeMultiSelectField,
		u.OfMessageTypeButtonField,
		u.OfMessageTypeTextareaField)
}
func (u *PartialValidateParamsPartialInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PartialValidateParamsPartialInputSchemaUnion) asAny() any {
	if !param.IsOmitted(u.OfMessageTypeListField) {
		return u.OfMessageTypeListField
	} else if !param.IsOmitted(u.OfMessageTypeSelectField) {
		return u.OfMessageTypeSelectField
	} else if !param.IsOmitted(u.OfMessageTypeBooleanField) {
		return u.OfMessageTypeBooleanField
	} else if !param.IsOmitted(u.OfMessageTypeJsonField) {
		return u.OfMessageTypeJsonField
	} else if !param.IsOmitted(u.OfMessageTypeNumberField) {
		return u.OfMessageTypeNumberField
	} else if !param.IsOmitted(u.OfMessageTypeTextField) {
		return u.OfMessageTypeTextField
	} else if !param.IsOmitted(u.OfMessageTypeImageField) {
		return u.OfMessageTypeImageField
	} else if !param.IsOmitted(u.OfMessageTypeColorField) {
		return u.OfMessageTypeColorField
	} else if !param.IsOmitted(u.OfMessageTypeURLField) {
		return u.OfMessageTypeURLField
	} else if !param.IsOmitted(u.OfMessageTypeMarkdownField) {
		return u.OfMessageTypeMarkdownField
	} else if !param.IsOmitted(u.OfMessageTypeMultiSelectField) {
		return u.OfMessageTypeMultiSelectField
	} else if !param.IsOmitted(u.OfMessageTypeButtonField) {
		return u.OfMessageTypeButtonField
	} else if !param.IsOmitted(u.OfMessageTypeTextareaField) {
		return u.OfMessageTypeTextareaField
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetAlt() *shared.MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.Alt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetURL() *shared.MessageTypeURLFieldParam {
	if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetText() *shared.MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeButtonField; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetKey() *string {
	if vt := u.OfMessageTypeListField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeSelectField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeBooleanField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeJsonField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeNumberField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeTextField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeColorField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeURLField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeTextareaField; vt != nil {
		return (*string)(&vt.Key)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetLabel() *string {
	if vt := u.OfMessageTypeListField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeSelectField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeBooleanField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeJsonField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeNumberField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeTextField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeImageField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeColorField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeURLField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeButtonField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeTextareaField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetType() *string {
	if vt := u.OfMessageTypeListField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeSelectField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeBooleanField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeJsonField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeNumberField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeTextField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeColorField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeURLField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeTextareaField; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u PartialValidateParamsPartialInputSchemaUnion) GetSettings() (res partialValidateParamsPartialInputSchemaUnionSettings) {
	if vt := u.OfMessageTypeListField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeSelectField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeBooleanField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeJsonField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeNumberField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeTextField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeColorField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeURLField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeTextareaField; vt != nil {
		res.any = &vt.Settings
	}
	return
}

// Can have the runtime types
// [*PartialValidateParamsPartialInputSchemaMessageTypeListFieldSettings],
// [*shared.MessageTypeSelectFieldSettingsParam],
// [*shared.MessageTypeBooleanFieldSettingsParam],
// [*shared.MessageTypeJsonFieldSettingsParam],
// [*PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings],
// [*shared.MessageTypeTextFieldSettingsParam],
// [*shared.MessageTypeImageFieldSettingsParam],
// [*PartialValidateParamsPartialInputSchemaMessageTypeColorFieldSettings],
// [*shared.MessageTypeURLFieldSettingsParam],
// [*shared.MessageTypeMarkdownFieldSettingsParam],
// [*shared.MessageTypeMultiSelectFieldSettingsParam],
// [*shared.MessageTypeButtonFieldSettingsParam],
// [*shared.MessageTypeTextareaFieldSettingsParam]
type partialValidateParamsPartialInputSchemaUnionSettings struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeListFieldSettings:
//	case *shared.MessageTypeSelectFieldSettingsParam:
//	case *shared.MessageTypeBooleanFieldSettingsParam:
//	case *shared.MessageTypeJsonFieldSettingsParam:
//	case *knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings:
//	case *shared.MessageTypeTextFieldSettingsParam:
//	case *shared.MessageTypeImageFieldSettingsParam:
//	case *knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeColorFieldSettings:
//	case *shared.MessageTypeURLFieldSettingsParam:
//	case *shared.MessageTypeMarkdownFieldSettingsParam:
//	case *shared.MessageTypeMultiSelectFieldSettingsParam:
//	case *shared.MessageTypeButtonFieldSettingsParam:
//	case *shared.MessageTypeTextareaFieldSettingsParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u partialValidateParamsPartialInputSchemaUnionSettings) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetItemSchema() *any {
	switch vt := u.any.(type) {
	case *PartialValidateParamsPartialInputSchemaMessageTypeListFieldSettings:
		return &vt.ItemSchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetSchema() *any {
	switch vt := u.any.(type) {
	case *shared.MessageTypeJsonFieldSettingsParam:
		return &vt.Schema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetMax() *float64 {
	switch vt := u.any.(type) {
	case *PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		return paramutil.AddrIfPresent(vt.Max)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetMin() *float64 {
	switch vt := u.any.(type) {
	case *PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		return paramutil.AddrIfPresent(vt.Min)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetUnitLabel() *string {
	switch vt := u.any.(type) {
	case *PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		return paramutil.AddrIfPresent(vt.UnitLabel)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetDescription() *string {
	switch vt := u.any.(type) {
	case *PartialValidateParamsPartialInputSchemaMessageTypeListFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeBooleanFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeJsonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeImageFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialValidateParamsPartialInputSchemaMessageTypeColorFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeURLFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeMarkdownFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeButtonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetPlaceholder() *string {
	switch vt := u.any.(type) {
	case *PartialValidateParamsPartialInputSchemaMessageTypeListFieldSettings:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeBooleanFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeJsonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeImageFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *PartialValidateParamsPartialInputSchemaMessageTypeColorFieldSettings:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeURLFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeMarkdownFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeButtonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetRequired() *bool {
	switch vt := u.any.(type) {
	case *PartialValidateParamsPartialInputSchemaMessageTypeListFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeBooleanFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeJsonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeImageFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialValidateParamsPartialInputSchemaMessageTypeColorFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeURLFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeMarkdownFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeButtonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetMaxLength() *int64 {
	switch vt := u.any.(type) {
	case *shared.MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MaxLength)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MaxLength)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetMinLength() *int64 {
	switch vt := u.any.(type) {
	case *shared.MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MinLength)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MinLength)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetDefault() (res partialValidateParamsPartialInputSchemaUnionSettingsDefault) {
	switch vt := u.any.(type) {
	case *PartialValidateParamsPartialInputSchemaMessageTypeListFieldSettings:
		res.any = &vt.Default
	case *shared.MessageTypeSelectFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeBooleanFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeJsonFieldSettingsParam:
		res.any = &vt.Default
	case *PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeTextFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *PartialValidateParamsPartialInputSchemaMessageTypeColorFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeURLFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeMarkdownFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		res.any = &vt.Default
	case *shared.MessageTypeTextareaFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	}
	return res
}

// Can have the runtime types [*[]any], [*string], [*bool], [*any], [*float64],
// [\*[]string]
type partialValidateParamsPartialInputSchemaUnionSettingsDefault struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]any:
//	case *string:
//	case *bool:
//	case *any:
//	case *float64:
//	case *[]string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u partialValidateParamsPartialInputSchemaUnionSettingsDefault) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetOptions() (res partialValidateParamsPartialInputSchemaUnionSettingsOptions) {
	switch vt := u.any.(type) {
	case *shared.MessageTypeSelectFieldSettingsParam:
		res.any = &vt.Options
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		res.any = &vt.Options
	}
	return res
}

// Can have the runtime types
// [_[]shared.MessageTypeSelectFieldSettingsOptionParam],
// [_[]shared.MessageTypeMultiSelectFieldSettingsOptionParam]
type partialValidateParamsPartialInputSchemaUnionSettingsOptions struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.MessageTypeSelectFieldSettingsOptionParam:
//	case *[]shared.MessageTypeMultiSelectFieldSettingsOptionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u partialValidateParamsPartialInputSchemaUnionSettingsOptions) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's Action property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetAction() *shared.MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.Action
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		return &vt.Action
	}
	return nil
}

// A list field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialValidateParamsPartialInputSchemaMessageTypeListField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The type of the field.
	//
	// Any of "list".
	Type string `json:"type,omitzero" api:"required"`
	// Settings for the list field.
	Settings PartialValidateParamsPartialInputSchemaMessageTypeListFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeListField) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeListField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeListField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartialInputSchemaMessageTypeListField](
		"type", "list",
	)
}

// Settings for the list field.
type PartialValidateParamsPartialInputSchemaMessageTypeListFieldSettings struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	// The default value of the list field.
	Default []any `json:"default,omitzero"`
	// A JSON schema used to validate the structure of each item in the list. Must be a
	// valid JSON schema.
	ItemSchema any `json:"item_schema,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeListFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeListFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeListFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A numeric field used in a message type or partial input schema, with optional
// min/max bounds and a unit label for display.
//
// The properties Key, Label, Type are required.
type PartialValidateParamsPartialInputSchemaMessageTypeNumberField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The type of the field.
	//
	// Any of "number".
	Type string `json:"type,omitzero" api:"required"`
	// Settings for the number field.
	Settings PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeNumberField) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeNumberField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeNumberField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartialInputSchemaMessageTypeNumberField](
		"type", "number",
	)
}

// Settings for the number field.
type PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings struct {
	// The default numeric value.
	Default     param.Opt[float64] `json:"default,omitzero"`
	Description param.Opt[string]  `json:"description,omitzero"`
	// Optional inclusive maximum allowed value.
	Max param.Opt[float64] `json:"max,omitzero"`
	// Optional inclusive minimum allowed value.
	Min         param.Opt[float64] `json:"min,omitzero"`
	Placeholder param.Opt[string]  `json:"placeholder,omitzero"`
	// Optional short label shown after the input (e.g. px, kg).
	UnitLabel param.Opt[string] `json:"unit_label,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeNumberFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A hex color field (#RGB or #RRGGBB) used in a message type or partial input
// schema.
//
// The properties Key, Label, Type are required.
type PartialValidateParamsPartialInputSchemaMessageTypeColorField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The type of the field.
	//
	// Any of "color".
	Type string `json:"type,omitzero" api:"required"`
	// Settings for the color field.
	Settings PartialValidateParamsPartialInputSchemaMessageTypeColorFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeColorField) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeColorField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeColorField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartialInputSchemaMessageTypeColorField](
		"type", "color",
	)
}

// Settings for the color field.
type PartialValidateParamsPartialInputSchemaMessageTypeColorFieldSettings struct {
	// The default hex color value.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeColorFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeColorFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeColorFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
