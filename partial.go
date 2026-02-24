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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
// [PartialInputSchemaMessageTypeBooleanField],
// [PartialInputSchemaMessageTypeButtonField],
// [PartialInputSchemaMessageTypeImageField],
// [PartialInputSchemaMessageTypeJsonField],
// [PartialInputSchemaMessageTypeMarkdownField],
// [PartialInputSchemaMessageTypeMultiSelectField],
// [PartialInputSchemaMessageTypeSelectField], [MessageTypeTextField],
// [PartialInputSchemaMessageTypeTextareaField],
// [PartialInputSchemaMessageTypeURLField].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PartialInputSchemaUnion struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Type  string `json:"type"`
	// This field is a union of [PartialInputSchemaMessageTypeBooleanFieldSettings],
	// [PartialInputSchemaMessageTypeButtonFieldSettings],
	// [PartialInputSchemaMessageTypeImageFieldSettings],
	// [PartialInputSchemaMessageTypeJsonFieldSettings],
	// [PartialInputSchemaMessageTypeMarkdownFieldSettings],
	// [PartialInputSchemaMessageTypeMultiSelectFieldSettings],
	// [PartialInputSchemaMessageTypeSelectFieldSettings],
	// [MessageTypeTextFieldSettings],
	// [PartialInputSchemaMessageTypeTextareaFieldSettings],
	// [PartialInputSchemaMessageTypeURLFieldSettings]
	Settings PartialInputSchemaUnionSettings `json:"settings"`
	// This field is from variant [PartialInputSchemaMessageTypeButtonField].
	Action MessageTypeTextField `json:"action"`
	// This field is from variant [PartialInputSchemaMessageTypeButtonField].
	Text MessageTypeTextField `json:"text"`
	// This field is from variant [PartialInputSchemaMessageTypeImageField].
	Alt MessageTypeTextField `json:"alt"`
	// This field is from variant [PartialInputSchemaMessageTypeImageField].
	URL  PartialInputSchemaMessageTypeImageFieldURL `json:"url"`
	JSON struct {
		Key      respjson.Field
		Label    respjson.Field
		Type     respjson.Field
		Settings respjson.Field
		Action   respjson.Field
		Text     respjson.Field
		Alt      respjson.Field
		URL      respjson.Field
		raw      string
	} `json:"-"`
}

func (u PartialInputSchemaUnion) AsMessageTypeBooleanField() (v PartialInputSchemaMessageTypeBooleanField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeButtonField() (v PartialInputSchemaMessageTypeButtonField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeImageField() (v PartialInputSchemaMessageTypeImageField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeJsonField() (v PartialInputSchemaMessageTypeJsonField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeMarkdownField() (v PartialInputSchemaMessageTypeMarkdownField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeMultiSelectField() (v PartialInputSchemaMessageTypeMultiSelectField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeSelectField() (v PartialInputSchemaMessageTypeSelectField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeTextField() (v MessageTypeTextField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeTextareaField() (v PartialInputSchemaMessageTypeTextareaField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PartialInputSchemaUnion) AsMessageTypeURLField() (v PartialInputSchemaMessageTypeURLField) {
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
	// This field is a union of [bool], [any], [string], [[]string], [string],
	// [string], [string], [string]
	Default     PartialInputSchemaUnionSettingsDefault `json:"default"`
	Description string                                 `json:"description"`
	Required    bool                                   `json:"required"`
	// This field is from variant [PartialInputSchemaMessageTypeJsonFieldSettings].
	Schema any `json:"schema"`
	// This field is a union of
	// [[]PartialInputSchemaMessageTypeMultiSelectFieldSettingsOption],
	// [[]PartialInputSchemaMessageTypeSelectFieldSettingsOption]
	Options   PartialInputSchemaUnionSettingsOptions `json:"options"`
	MaxLength int64                                  `json:"max_length"`
	MinLength int64                                  `json:"min_length"`
	JSON      struct {
		Default     respjson.Field
		Description respjson.Field
		Required    respjson.Field
		Schema      respjson.Field
		Options     respjson.Field
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
// will be valid: OfBool OfPartialInputSchemaMessageTypeJsonFieldSettingsDefault
// OfString OfStringArray]
type PartialInputSchemaUnionSettingsDefault struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfPartialInputSchemaMessageTypeJsonFieldSettingsDefault any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfBool                                                  respjson.Field
		OfPartialInputSchemaMessageTypeJsonFieldSettingsDefault respjson.Field
		OfString                                                respjson.Field
		OfStringArray                                           respjson.Field
		raw                                                     string
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
// will be valid: OfPartialInputSchemaMessageTypeMultiSelectFieldSettingsOptions
// OfPartialInputSchemaMessageTypeSelectFieldSettingsOptions]
type PartialInputSchemaUnionSettingsOptions struct {
	// This field will be present if the value is a
	// [[]PartialInputSchemaMessageTypeMultiSelectFieldSettingsOption] instead of an
	// object.
	OfPartialInputSchemaMessageTypeMultiSelectFieldSettingsOptions []PartialInputSchemaMessageTypeMultiSelectFieldSettingsOption `json:",inline"`
	// This field will be present if the value is a
	// [[]PartialInputSchemaMessageTypeSelectFieldSettingsOption] instead of an object.
	OfPartialInputSchemaMessageTypeSelectFieldSettingsOptions []PartialInputSchemaMessageTypeSelectFieldSettingsOption `json:",inline"`
	JSON                                                      struct {
		OfPartialInputSchemaMessageTypeMultiSelectFieldSettingsOptions respjson.Field
		OfPartialInputSchemaMessageTypeSelectFieldSettingsOptions      respjson.Field
		raw                                                            string
	} `json:"-"`
}

func (r *PartialInputSchemaUnionSettingsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A boolean field used in a message type.
type PartialInputSchemaMessageTypeBooleanField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// The type of the field.
	//
	// Any of "boolean".
	Type string `json:"type,required"`
	// Settings for the boolean field.
	Settings PartialInputSchemaMessageTypeBooleanFieldSettings `json:"settings"`
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
func (r PartialInputSchemaMessageTypeBooleanField) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeBooleanField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the boolean field.
type PartialInputSchemaMessageTypeBooleanFieldSettings struct {
	// The default value of the boolean field.
	Default     bool   `json:"default"`
	Description string `json:"description"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeBooleanFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeBooleanFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A button field used in a message type.
type PartialInputSchemaMessageTypeButtonField struct {
	// A text field used in a message type.
	Action MessageTypeTextField `json:"action,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// A text field used in a message type.
	Text MessageTypeTextField `json:"text,required"`
	// The type of the field.
	//
	// Any of "button".
	Type string `json:"type,required"`
	// Settings for the button field.
	Settings PartialInputSchemaMessageTypeButtonFieldSettings `json:"settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Key         respjson.Field
		Label       respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeButtonField) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeButtonField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the button field.
type PartialInputSchemaMessageTypeButtonFieldSettings struct {
	Description string `json:"description"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeButtonFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeButtonFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image field used in a message type.
type PartialInputSchemaMessageTypeImageField struct {
	// A text field used in a message type.
	Action MessageTypeTextField `json:"action,required"`
	// A text field used in a message type.
	Alt MessageTypeTextField `json:"alt,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// The type of the field.
	//
	// Any of "image".
	Type string `json:"type,required"`
	// A URL field used in a message type.
	URL PartialInputSchemaMessageTypeImageFieldURL `json:"url,required"`
	// Settings for the image field.
	Settings PartialInputSchemaMessageTypeImageFieldSettings `json:"settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Alt         respjson.Field
		Key         respjson.Field
		Label       respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeImageField) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeImageField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL field used in a message type.
type PartialInputSchemaMessageTypeImageFieldURL struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// The type of the field.
	//
	// Any of "url".
	Type string `json:"type,required"`
	// Settings for the url field.
	Settings PartialInputSchemaMessageTypeImageFieldURLSettings `json:"settings"`
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
func (r PartialInputSchemaMessageTypeImageFieldURL) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeImageFieldURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the url field.
type PartialInputSchemaMessageTypeImageFieldURLSettings struct {
	// The default value of the URL field.
	Default     string `json:"default,nullable"`
	Description string `json:"description"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeImageFieldURLSettings) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeImageFieldURLSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the image field.
type PartialInputSchemaMessageTypeImageFieldSettings struct {
	Description string `json:"description"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeImageFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeImageFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A JSON field used in a message type.
type PartialInputSchemaMessageTypeJsonField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// The type of the field.
	//
	// Any of "json".
	Type string `json:"type,required"`
	// Settings for the json field.
	Settings PartialInputSchemaMessageTypeJsonFieldSettings `json:"settings"`
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
func (r PartialInputSchemaMessageTypeJsonField) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeJsonField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the json field.
type PartialInputSchemaMessageTypeJsonFieldSettings struct {
	// The default value of the JSON field.
	Default     any    `json:"default,nullable"`
	Description string `json:"description"`
	// Whether the field is required.
	Required bool `json:"required"`
	// A JSON schema used to validate the structure of the JSON provided. Must be a
	// valid JSON schema.
	Schema any `json:"schema,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Required    respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeJsonFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeJsonFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A markdown field used in a message type.
type PartialInputSchemaMessageTypeMarkdownField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// The type of the field.
	//
	// Any of "markdown".
	Type string `json:"type,required"`
	// Settings for the markdown field.
	Settings PartialInputSchemaMessageTypeMarkdownFieldSettings `json:"settings"`
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
func (r PartialInputSchemaMessageTypeMarkdownField) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeMarkdownField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the markdown field.
type PartialInputSchemaMessageTypeMarkdownFieldSettings struct {
	// The default value of the markdown field.
	Default     string `json:"default"`
	Description string `json:"description"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeMarkdownFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeMarkdownFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A multi-select field used in a message type.
type PartialInputSchemaMessageTypeMultiSelectField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// Settings for the multi_select field.
	Settings PartialInputSchemaMessageTypeMultiSelectFieldSettings `json:"settings,required"`
	// The type of the field.
	//
	// Any of "multi_select".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Label       respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeMultiSelectField) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeMultiSelectField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the multi_select field.
type PartialInputSchemaMessageTypeMultiSelectFieldSettings struct {
	// The default values for the multi-select field.
	Default     []string `json:"default,nullable"`
	Description string   `json:"description"`
	// The available options for the multi-select field.
	Options []PartialInputSchemaMessageTypeMultiSelectFieldSettingsOption `json:"options"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Options     respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeMultiSelectFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeMultiSelectFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartialInputSchemaMessageTypeMultiSelectFieldSettingsOption struct {
	// The value for the option.
	Value string `json:"value,required"`
	// The display label for the option.
	Label string `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeMultiSelectFieldSettingsOption) RawJSON() string {
	return r.JSON.raw
}
func (r *PartialInputSchemaMessageTypeMultiSelectFieldSettingsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A select field used in a message type.
type PartialInputSchemaMessageTypeSelectField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// Settings for the select field.
	Settings PartialInputSchemaMessageTypeSelectFieldSettings `json:"settings,required"`
	// The type of the field.
	//
	// Any of "select".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Label       respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeSelectField) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeSelectField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the select field.
type PartialInputSchemaMessageTypeSelectFieldSettings struct {
	// The default value for the select field.
	Default     string `json:"default,nullable"`
	Description string `json:"description"`
	// The available options for the select field.
	Options []PartialInputSchemaMessageTypeSelectFieldSettingsOption `json:"options"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Options     respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeSelectFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeSelectFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartialInputSchemaMessageTypeSelectFieldSettingsOption struct {
	// The value for the option.
	Value string `json:"value,required"`
	// The display label for the option.
	Label string `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeSelectFieldSettingsOption) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeSelectFieldSettingsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A textarea field used in a message type.
type PartialInputSchemaMessageTypeTextareaField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// The type of the field.
	//
	// Any of "textarea".
	Type string `json:"type,required"`
	// Settings for the textarea field.
	Settings PartialInputSchemaMessageTypeTextareaFieldSettings `json:"settings"`
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
func (r PartialInputSchemaMessageTypeTextareaField) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeTextareaField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the textarea field.
type PartialInputSchemaMessageTypeTextareaFieldSettings struct {
	// The default value of the textarea field.
	Default     string `json:"default,nullable"`
	Description string `json:"description"`
	MaxLength   int64  `json:"max_length"`
	MinLength   int64  `json:"min_length"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		MaxLength   respjson.Field
		MinLength   respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeTextareaFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeTextareaFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL field used in a message type.
type PartialInputSchemaMessageTypeURLField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// The type of the field.
	//
	// Any of "url".
	Type string `json:"type,required"`
	// Settings for the url field.
	Settings PartialInputSchemaMessageTypeURLFieldSettings `json:"settings"`
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
func (r PartialInputSchemaMessageTypeURLField) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeURLField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the url field.
type PartialInputSchemaMessageTypeURLFieldSettings struct {
	// The default value of the URL field.
	Default     string `json:"default,nullable"`
	Description string `json:"description"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PartialInputSchemaMessageTypeURLFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *PartialInputSchemaMessageTypeURLFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
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
	// The partial content.
	Content string `json:"content,required"`
	// A name for the partial. Must be at maximum 255 characters in length.
	Name string `json:"name,required"`
	// The partial type. One of 'html', 'json', 'markdown', 'text'.
	//
	// Any of "html", "text", "json", "markdown".
	Type string `json:"type,omitzero,required"`
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
	OfMessageTypeBooleanField     *PartialUpsertParamsPartialInputSchemaMessageTypeBooleanField     `json:",omitzero,inline"`
	OfMessageTypeButtonField      *PartialUpsertParamsPartialInputSchemaMessageTypeButtonField      `json:",omitzero,inline"`
	OfMessageTypeImageField       *PartialUpsertParamsPartialInputSchemaMessageTypeImageField       `json:",omitzero,inline"`
	OfMessageTypeJsonField        *PartialUpsertParamsPartialInputSchemaMessageTypeJsonField        `json:",omitzero,inline"`
	OfMessageTypeMarkdownField    *PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownField    `json:",omitzero,inline"`
	OfMessageTypeMultiSelectField *PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectField `json:",omitzero,inline"`
	OfMessageTypeSelectField      *PartialUpsertParamsPartialInputSchemaMessageTypeSelectField      `json:",omitzero,inline"`
	OfMessageTypeTextField        *MessageTypeTextFieldParam                                        `json:",omitzero,inline"`
	OfMessageTypeTextareaField    *PartialUpsertParamsPartialInputSchemaMessageTypeTextareaField    `json:",omitzero,inline"`
	OfMessageTypeURLField         *PartialUpsertParamsPartialInputSchemaMessageTypeURLField         `json:",omitzero,inline"`
	paramUnion
}

func (u PartialUpsertParamsPartialInputSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMessageTypeBooleanField,
		u.OfMessageTypeButtonField,
		u.OfMessageTypeImageField,
		u.OfMessageTypeJsonField,
		u.OfMessageTypeMarkdownField,
		u.OfMessageTypeMultiSelectField,
		u.OfMessageTypeSelectField,
		u.OfMessageTypeTextField,
		u.OfMessageTypeTextareaField,
		u.OfMessageTypeURLField)
}
func (u *PartialUpsertParamsPartialInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PartialUpsertParamsPartialInputSchemaUnion) asAny() any {
	if !param.IsOmitted(u.OfMessageTypeBooleanField) {
		return u.OfMessageTypeBooleanField
	} else if !param.IsOmitted(u.OfMessageTypeButtonField) {
		return u.OfMessageTypeButtonField
	} else if !param.IsOmitted(u.OfMessageTypeImageField) {
		return u.OfMessageTypeImageField
	} else if !param.IsOmitted(u.OfMessageTypeJsonField) {
		return u.OfMessageTypeJsonField
	} else if !param.IsOmitted(u.OfMessageTypeMarkdownField) {
		return u.OfMessageTypeMarkdownField
	} else if !param.IsOmitted(u.OfMessageTypeMultiSelectField) {
		return u.OfMessageTypeMultiSelectField
	} else if !param.IsOmitted(u.OfMessageTypeSelectField) {
		return u.OfMessageTypeSelectField
	} else if !param.IsOmitted(u.OfMessageTypeTextField) {
		return u.OfMessageTypeTextField
	} else if !param.IsOmitted(u.OfMessageTypeTextareaField) {
		return u.OfMessageTypeTextareaField
	} else if !param.IsOmitted(u.OfMessageTypeURLField) {
		return u.OfMessageTypeURLField
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetText() *MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeButtonField; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetAlt() *MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.Alt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetURL() *PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldURL {
	if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetKey() *string {
	if vt := u.OfMessageTypeBooleanField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeJsonField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeSelectField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeTextField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeTextareaField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeURLField; vt != nil {
		return (*string)(&vt.Key)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetLabel() *string {
	if vt := u.OfMessageTypeBooleanField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeButtonField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeImageField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeJsonField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeSelectField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeTextField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeTextareaField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeURLField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetType() *string {
	if vt := u.OfMessageTypeBooleanField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeJsonField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeSelectField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeTextField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeTextareaField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeURLField; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u PartialUpsertParamsPartialInputSchemaUnion) GetSettings() (res partialUpsertParamsPartialInputSchemaUnionSettings) {
	if vt := u.OfMessageTypeBooleanField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeJsonField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeSelectField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeTextField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeTextareaField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeURLField; vt != nil {
		res.any = &vt.Settings
	}
	return
}

// Can have the runtime types
// [*PartialUpsertParamsPartialInputSchemaMessageTypeBooleanFieldSettings],
// [*PartialUpsertParamsPartialInputSchemaMessageTypeButtonFieldSettings],
// [*PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldSettings],
// [*PartialUpsertParamsPartialInputSchemaMessageTypeJsonFieldSettings],
// [*PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownFieldSettings],
// [*PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings],
// [*PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettings],
// [*MessageTypeTextFieldSettingsParam],
// [*PartialUpsertParamsPartialInputSchemaMessageTypeTextareaFieldSettings],
// [*PartialUpsertParamsPartialInputSchemaMessageTypeURLFieldSettings]
type partialUpsertParamsPartialInputSchemaUnionSettings struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeBooleanFieldSettings:
//	case *knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeButtonFieldSettings:
//	case *knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldSettings:
//	case *knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeJsonFieldSettings:
//	case *knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownFieldSettings:
//	case *knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings:
//	case *knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettings:
//	case *knockmapi.MessageTypeTextFieldSettingsParam:
//	case *knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeTextareaFieldSettings:
//	case *knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeURLFieldSettings:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u partialUpsertParamsPartialInputSchemaUnionSettings) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetSchema() *any {
	switch vt := u.any.(type) {
	case *PartialUpsertParamsPartialInputSchemaMessageTypeJsonFieldSettings:
		return &vt.Schema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetDescription() *string {
	switch vt := u.any.(type) {
	case *PartialUpsertParamsPartialInputSchemaMessageTypeBooleanFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeButtonFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeJsonFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeTextareaFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeURLFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetRequired() *bool {
	switch vt := u.any.(type) {
	case *PartialUpsertParamsPartialInputSchemaMessageTypeBooleanFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeButtonFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeJsonFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeTextareaFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeURLFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetMaxLength() *int64 {
	switch vt := u.any.(type) {
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MaxLength)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeTextareaFieldSettings:
		return paramutil.AddrIfPresent(vt.MaxLength)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetMinLength() *int64 {
	switch vt := u.any.(type) {
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MinLength)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeTextareaFieldSettings:
		return paramutil.AddrIfPresent(vt.MinLength)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u partialUpsertParamsPartialInputSchemaUnionSettings) GetDefault() (res partialUpsertParamsPartialInputSchemaUnionSettingsDefault) {
	switch vt := u.any.(type) {
	case *PartialUpsertParamsPartialInputSchemaMessageTypeBooleanFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeJsonFieldSettings:
		res.any = &vt.Default
	case *PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings:
		res.any = &vt.Default
	case *PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *MessageTypeTextFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeTextareaFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *PartialUpsertParamsPartialInputSchemaMessageTypeURLFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	}
	return res
}

// Can have the runtime types [*bool], [*any], [*string], [\*[]string]
type partialUpsertParamsPartialInputSchemaUnionSettingsDefault struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *bool:
//	case *any:
//	case *string:
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
	case *PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings:
		res.any = &vt.Options
	case *PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettings:
		res.any = &vt.Options
	}
	return res
}

// Can have the runtime types
// [_[]PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption],
// [_[]PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption]
type partialUpsertParamsPartialInputSchemaUnionSettingsOptions struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption:
//	case *[]knockmapi.PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u partialUpsertParamsPartialInputSchemaUnionSettingsOptions) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's Action property, if present.
func (u PartialUpsertParamsPartialInputSchemaUnion) GetAction() *MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeButtonField; vt != nil {
		return &vt.Action
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.Action
	}
	return nil
}

// A boolean field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialUpsertParamsPartialInputSchemaMessageTypeBooleanField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "boolean".
	Type string `json:"type,omitzero,required"`
	// Settings for the boolean field.
	Settings PartialUpsertParamsPartialInputSchemaMessageTypeBooleanFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeBooleanField) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeBooleanField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeBooleanField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartialInputSchemaMessageTypeBooleanField](
		"type", "boolean",
	)
}

// Settings for the boolean field.
type PartialUpsertParamsPartialInputSchemaMessageTypeBooleanFieldSettings struct {
	// The default value of the boolean field.
	Default     param.Opt[bool]   `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeBooleanFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeBooleanFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeBooleanFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A button field used in a message type.
//
// The properties Action, Key, Label, Text, Type are required.
type PartialUpsertParamsPartialInputSchemaMessageTypeButtonField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// A text field used in a message type.
	Action MessageTypeTextFieldParam `json:"action,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// A text field used in a message type.
	Text MessageTypeTextFieldParam `json:"text,omitzero,required"`
	// The type of the field.
	//
	// Any of "button".
	Type string `json:"type,omitzero,required"`
	// Settings for the button field.
	Settings PartialUpsertParamsPartialInputSchemaMessageTypeButtonFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeButtonField) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeButtonField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeButtonField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartialInputSchemaMessageTypeButtonField](
		"type", "button",
	)
}

// Settings for the button field.
type PartialUpsertParamsPartialInputSchemaMessageTypeButtonFieldSettings struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeButtonFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeButtonFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeButtonFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image field used in a message type.
//
// The properties Action, Alt, Key, Label, Type, URL are required.
type PartialUpsertParamsPartialInputSchemaMessageTypeImageField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// A text field used in a message type.
	Action MessageTypeTextFieldParam `json:"action,omitzero,required"`
	// A text field used in a message type.
	Alt MessageTypeTextFieldParam `json:"alt,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "image".
	Type string `json:"type,omitzero,required"`
	// A URL field used in a message type.
	URL PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldURL `json:"url,omitzero,required"`
	// Settings for the image field.
	Settings PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeImageField) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeImageField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeImageField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartialInputSchemaMessageTypeImageField](
		"type", "image",
	)
}

// A URL field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldURL struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "url".
	Type string `json:"type,omitzero,required"`
	// Settings for the url field.
	Settings PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldURLSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldURL) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldURL](
		"type", "url",
	)
}

// Settings for the url field.
type PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldURLSettings struct {
	// The default value of the URL field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldURLSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldURLSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldURLSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the image field.
type PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldSettings struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeImageFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A JSON field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialUpsertParamsPartialInputSchemaMessageTypeJsonField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "json".
	Type string `json:"type,omitzero,required"`
	// Settings for the json field.
	Settings PartialUpsertParamsPartialInputSchemaMessageTypeJsonFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeJsonField) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeJsonField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeJsonField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartialInputSchemaMessageTypeJsonField](
		"type", "json",
	)
}

// Settings for the json field.
type PartialUpsertParamsPartialInputSchemaMessageTypeJsonFieldSettings struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	// The default value of the JSON field.
	Default any `json:"default,omitzero"`
	// A JSON schema used to validate the structure of the JSON provided. Must be a
	// valid JSON schema.
	Schema any `json:"schema,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeJsonFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeJsonFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeJsonFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A markdown field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "markdown".
	Type string `json:"type,omitzero,required"`
	// Settings for the markdown field.
	Settings PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownField) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownField](
		"type", "markdown",
	)
}

// Settings for the markdown field.
type PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownFieldSettings struct {
	// The default value of the markdown field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeMarkdownFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A multi-select field used in a message type.
//
// The properties Key, Label, Settings, Type are required.
type PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// Settings for the multi_select field.
	Settings PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings `json:"settings,omitzero,required"`
	// The type of the field.
	//
	// Any of "multi_select".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectField) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectField](
		"type", "multi_select",
	)
}

// Settings for the multi_select field.
type PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	// The default values for the multi-select field.
	Default []string `json:"default,omitzero"`
	// The available options for the multi-select field.
	Options []PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption `json:"options,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Value is required.
type PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption struct {
	// The value for the option.
	Value string `json:"value,required"`
	// The display label for the option.
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A select field used in a message type.
//
// The properties Key, Label, Settings, Type are required.
type PartialUpsertParamsPartialInputSchemaMessageTypeSelectField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// Settings for the select field.
	Settings PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettings `json:"settings,omitzero,required"`
	// The type of the field.
	//
	// Any of "select".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeSelectField) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeSelectField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeSelectField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartialInputSchemaMessageTypeSelectField](
		"type", "select",
	)
}

// Settings for the select field.
type PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettings struct {
	// The default value for the select field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	// The available options for the select field.
	Options []PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption `json:"options,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Value is required.
type PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption struct {
	// The value for the option.
	Value string `json:"value,required"`
	// The display label for the option.
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A textarea field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialUpsertParamsPartialInputSchemaMessageTypeTextareaField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "textarea".
	Type string `json:"type,omitzero,required"`
	// Settings for the textarea field.
	Settings PartialUpsertParamsPartialInputSchemaMessageTypeTextareaFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeTextareaField) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeTextareaField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeTextareaField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartialInputSchemaMessageTypeTextareaField](
		"type", "textarea",
	)
}

// Settings for the textarea field.
type PartialUpsertParamsPartialInputSchemaMessageTypeTextareaFieldSettings struct {
	// The default value of the textarea field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	MaxLength   param.Opt[int64]  `json:"max_length,omitzero"`
	MinLength   param.Opt[int64]  `json:"min_length,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeTextareaFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeTextareaFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeTextareaFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialUpsertParamsPartialInputSchemaMessageTypeURLField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "url".
	Type string `json:"type,omitzero,required"`
	// Settings for the url field.
	Settings PartialUpsertParamsPartialInputSchemaMessageTypeURLFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeURLField) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeURLField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeURLField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialUpsertParamsPartialInputSchemaMessageTypeURLField](
		"type", "url",
	)
}

// Settings for the url field.
type PartialUpsertParamsPartialInputSchemaMessageTypeURLFieldSettings struct {
	// The default value of the URL field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialUpsertParamsPartialInputSchemaMessageTypeURLFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialUpsertParamsPartialInputSchemaMessageTypeURLFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialUpsertParamsPartialInputSchemaMessageTypeURLFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartialValidateParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// A partial object with attributes to update or create a partial.
	Partial PartialValidateParamsPartial `json:"partial,omitzero,required"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A partial object with attributes to update or create a partial.
//
// The properties Content, Name, Type are required.
type PartialValidateParamsPartial struct {
	// The partial content.
	Content string `json:"content,required"`
	// A name for the partial. Must be at maximum 255 characters in length.
	Name string `json:"name,required"`
	// The partial type. One of 'html', 'json', 'markdown', 'text'.
	//
	// Any of "html", "text", "json", "markdown".
	Type string `json:"type,omitzero,required"`
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
	OfMessageTypeBooleanField     *PartialValidateParamsPartialInputSchemaMessageTypeBooleanField     `json:",omitzero,inline"`
	OfMessageTypeButtonField      *PartialValidateParamsPartialInputSchemaMessageTypeButtonField      `json:",omitzero,inline"`
	OfMessageTypeImageField       *PartialValidateParamsPartialInputSchemaMessageTypeImageField       `json:",omitzero,inline"`
	OfMessageTypeJsonField        *PartialValidateParamsPartialInputSchemaMessageTypeJsonField        `json:",omitzero,inline"`
	OfMessageTypeMarkdownField    *PartialValidateParamsPartialInputSchemaMessageTypeMarkdownField    `json:",omitzero,inline"`
	OfMessageTypeMultiSelectField *PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectField `json:",omitzero,inline"`
	OfMessageTypeSelectField      *PartialValidateParamsPartialInputSchemaMessageTypeSelectField      `json:",omitzero,inline"`
	OfMessageTypeTextField        *MessageTypeTextFieldParam                                          `json:",omitzero,inline"`
	OfMessageTypeTextareaField    *PartialValidateParamsPartialInputSchemaMessageTypeTextareaField    `json:",omitzero,inline"`
	OfMessageTypeURLField         *PartialValidateParamsPartialInputSchemaMessageTypeURLField         `json:",omitzero,inline"`
	paramUnion
}

func (u PartialValidateParamsPartialInputSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMessageTypeBooleanField,
		u.OfMessageTypeButtonField,
		u.OfMessageTypeImageField,
		u.OfMessageTypeJsonField,
		u.OfMessageTypeMarkdownField,
		u.OfMessageTypeMultiSelectField,
		u.OfMessageTypeSelectField,
		u.OfMessageTypeTextField,
		u.OfMessageTypeTextareaField,
		u.OfMessageTypeURLField)
}
func (u *PartialValidateParamsPartialInputSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PartialValidateParamsPartialInputSchemaUnion) asAny() any {
	if !param.IsOmitted(u.OfMessageTypeBooleanField) {
		return u.OfMessageTypeBooleanField
	} else if !param.IsOmitted(u.OfMessageTypeButtonField) {
		return u.OfMessageTypeButtonField
	} else if !param.IsOmitted(u.OfMessageTypeImageField) {
		return u.OfMessageTypeImageField
	} else if !param.IsOmitted(u.OfMessageTypeJsonField) {
		return u.OfMessageTypeJsonField
	} else if !param.IsOmitted(u.OfMessageTypeMarkdownField) {
		return u.OfMessageTypeMarkdownField
	} else if !param.IsOmitted(u.OfMessageTypeMultiSelectField) {
		return u.OfMessageTypeMultiSelectField
	} else if !param.IsOmitted(u.OfMessageTypeSelectField) {
		return u.OfMessageTypeSelectField
	} else if !param.IsOmitted(u.OfMessageTypeTextField) {
		return u.OfMessageTypeTextField
	} else if !param.IsOmitted(u.OfMessageTypeTextareaField) {
		return u.OfMessageTypeTextareaField
	} else if !param.IsOmitted(u.OfMessageTypeURLField) {
		return u.OfMessageTypeURLField
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetText() *MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeButtonField; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetAlt() *MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.Alt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetURL() *PartialValidateParamsPartialInputSchemaMessageTypeImageFieldURL {
	if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetKey() *string {
	if vt := u.OfMessageTypeBooleanField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeJsonField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeSelectField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeTextField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeTextareaField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeURLField; vt != nil {
		return (*string)(&vt.Key)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetLabel() *string {
	if vt := u.OfMessageTypeBooleanField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeButtonField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeImageField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeJsonField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeSelectField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeTextField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeTextareaField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeURLField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetType() *string {
	if vt := u.OfMessageTypeBooleanField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeJsonField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeSelectField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeTextField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeTextareaField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeURLField; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u PartialValidateParamsPartialInputSchemaUnion) GetSettings() (res partialValidateParamsPartialInputSchemaUnionSettings) {
	if vt := u.OfMessageTypeBooleanField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeJsonField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeMarkdownField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeMultiSelectField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeSelectField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeTextField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeTextareaField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeURLField; vt != nil {
		res.any = &vt.Settings
	}
	return
}

// Can have the runtime types
// [*PartialValidateParamsPartialInputSchemaMessageTypeBooleanFieldSettings],
// [*PartialValidateParamsPartialInputSchemaMessageTypeButtonFieldSettings],
// [*PartialValidateParamsPartialInputSchemaMessageTypeImageFieldSettings],
// [*PartialValidateParamsPartialInputSchemaMessageTypeJsonFieldSettings],
// [*PartialValidateParamsPartialInputSchemaMessageTypeMarkdownFieldSettings],
// [*PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings],
// [*PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettings],
// [*MessageTypeTextFieldSettingsParam],
// [*PartialValidateParamsPartialInputSchemaMessageTypeTextareaFieldSettings],
// [*PartialValidateParamsPartialInputSchemaMessageTypeURLFieldSettings]
type partialValidateParamsPartialInputSchemaUnionSettings struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeBooleanFieldSettings:
//	case *knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeButtonFieldSettings:
//	case *knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeImageFieldSettings:
//	case *knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeJsonFieldSettings:
//	case *knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeMarkdownFieldSettings:
//	case *knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings:
//	case *knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettings:
//	case *knockmapi.MessageTypeTextFieldSettingsParam:
//	case *knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeTextareaFieldSettings:
//	case *knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeURLFieldSettings:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u partialValidateParamsPartialInputSchemaUnionSettings) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetSchema() *any {
	switch vt := u.any.(type) {
	case *PartialValidateParamsPartialInputSchemaMessageTypeJsonFieldSettings:
		return &vt.Schema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetDescription() *string {
	switch vt := u.any.(type) {
	case *PartialValidateParamsPartialInputSchemaMessageTypeBooleanFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialValidateParamsPartialInputSchemaMessageTypeButtonFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialValidateParamsPartialInputSchemaMessageTypeImageFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialValidateParamsPartialInputSchemaMessageTypeJsonFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialValidateParamsPartialInputSchemaMessageTypeMarkdownFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialValidateParamsPartialInputSchemaMessageTypeTextareaFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	case *PartialValidateParamsPartialInputSchemaMessageTypeURLFieldSettings:
		return paramutil.AddrIfPresent(vt.Description)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetRequired() *bool {
	switch vt := u.any.(type) {
	case *PartialValidateParamsPartialInputSchemaMessageTypeBooleanFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialValidateParamsPartialInputSchemaMessageTypeButtonFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialValidateParamsPartialInputSchemaMessageTypeImageFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialValidateParamsPartialInputSchemaMessageTypeJsonFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialValidateParamsPartialInputSchemaMessageTypeMarkdownFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialValidateParamsPartialInputSchemaMessageTypeTextareaFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	case *PartialValidateParamsPartialInputSchemaMessageTypeURLFieldSettings:
		return paramutil.AddrIfPresent(vt.Required)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetMaxLength() *int64 {
	switch vt := u.any.(type) {
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MaxLength)
	case *PartialValidateParamsPartialInputSchemaMessageTypeTextareaFieldSettings:
		return paramutil.AddrIfPresent(vt.MaxLength)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetMinLength() *int64 {
	switch vt := u.any.(type) {
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MinLength)
	case *PartialValidateParamsPartialInputSchemaMessageTypeTextareaFieldSettings:
		return paramutil.AddrIfPresent(vt.MinLength)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u partialValidateParamsPartialInputSchemaUnionSettings) GetDefault() (res partialValidateParamsPartialInputSchemaUnionSettingsDefault) {
	switch vt := u.any.(type) {
	case *PartialValidateParamsPartialInputSchemaMessageTypeBooleanFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *PartialValidateParamsPartialInputSchemaMessageTypeJsonFieldSettings:
		res.any = &vt.Default
	case *PartialValidateParamsPartialInputSchemaMessageTypeMarkdownFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings:
		res.any = &vt.Default
	case *PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *MessageTypeTextFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *PartialValidateParamsPartialInputSchemaMessageTypeTextareaFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *PartialValidateParamsPartialInputSchemaMessageTypeURLFieldSettings:
		res.any = paramutil.AddrIfPresent(vt.Default)
	}
	return res
}

// Can have the runtime types [*bool], [*any], [*string], [\*[]string]
type partialValidateParamsPartialInputSchemaUnionSettingsDefault struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *bool:
//	case *any:
//	case *string:
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
	case *PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings:
		res.any = &vt.Options
	case *PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettings:
		res.any = &vt.Options
	}
	return res
}

// Can have the runtime types
// [_[]PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption],
// [_[]PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption]
type partialValidateParamsPartialInputSchemaUnionSettingsOptions struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption:
//	case *[]knockmapi.PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u partialValidateParamsPartialInputSchemaUnionSettingsOptions) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's Action property, if present.
func (u PartialValidateParamsPartialInputSchemaUnion) GetAction() *MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeButtonField; vt != nil {
		return &vt.Action
	} else if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.Action
	}
	return nil
}

// A boolean field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialValidateParamsPartialInputSchemaMessageTypeBooleanField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "boolean".
	Type string `json:"type,omitzero,required"`
	// Settings for the boolean field.
	Settings PartialValidateParamsPartialInputSchemaMessageTypeBooleanFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeBooleanField) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeBooleanField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeBooleanField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartialInputSchemaMessageTypeBooleanField](
		"type", "boolean",
	)
}

// Settings for the boolean field.
type PartialValidateParamsPartialInputSchemaMessageTypeBooleanFieldSettings struct {
	// The default value of the boolean field.
	Default     param.Opt[bool]   `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeBooleanFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeBooleanFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeBooleanFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A button field used in a message type.
//
// The properties Action, Key, Label, Text, Type are required.
type PartialValidateParamsPartialInputSchemaMessageTypeButtonField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// A text field used in a message type.
	Action MessageTypeTextFieldParam `json:"action,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// A text field used in a message type.
	Text MessageTypeTextFieldParam `json:"text,omitzero,required"`
	// The type of the field.
	//
	// Any of "button".
	Type string `json:"type,omitzero,required"`
	// Settings for the button field.
	Settings PartialValidateParamsPartialInputSchemaMessageTypeButtonFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeButtonField) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeButtonField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeButtonField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartialInputSchemaMessageTypeButtonField](
		"type", "button",
	)
}

// Settings for the button field.
type PartialValidateParamsPartialInputSchemaMessageTypeButtonFieldSettings struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeButtonFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeButtonFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeButtonFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image field used in a message type.
//
// The properties Action, Alt, Key, Label, Type, URL are required.
type PartialValidateParamsPartialInputSchemaMessageTypeImageField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// A text field used in a message type.
	Action MessageTypeTextFieldParam `json:"action,omitzero,required"`
	// A text field used in a message type.
	Alt MessageTypeTextFieldParam `json:"alt,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "image".
	Type string `json:"type,omitzero,required"`
	// A URL field used in a message type.
	URL PartialValidateParamsPartialInputSchemaMessageTypeImageFieldURL `json:"url,omitzero,required"`
	// Settings for the image field.
	Settings PartialValidateParamsPartialInputSchemaMessageTypeImageFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeImageField) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeImageField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeImageField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartialInputSchemaMessageTypeImageField](
		"type", "image",
	)
}

// A URL field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialValidateParamsPartialInputSchemaMessageTypeImageFieldURL struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "url".
	Type string `json:"type,omitzero,required"`
	// Settings for the url field.
	Settings PartialValidateParamsPartialInputSchemaMessageTypeImageFieldURLSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeImageFieldURL) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeImageFieldURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeImageFieldURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartialInputSchemaMessageTypeImageFieldURL](
		"type", "url",
	)
}

// Settings for the url field.
type PartialValidateParamsPartialInputSchemaMessageTypeImageFieldURLSettings struct {
	// The default value of the URL field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeImageFieldURLSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeImageFieldURLSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeImageFieldURLSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the image field.
type PartialValidateParamsPartialInputSchemaMessageTypeImageFieldSettings struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeImageFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeImageFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeImageFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A JSON field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialValidateParamsPartialInputSchemaMessageTypeJsonField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "json".
	Type string `json:"type,omitzero,required"`
	// Settings for the json field.
	Settings PartialValidateParamsPartialInputSchemaMessageTypeJsonFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeJsonField) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeJsonField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeJsonField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartialInputSchemaMessageTypeJsonField](
		"type", "json",
	)
}

// Settings for the json field.
type PartialValidateParamsPartialInputSchemaMessageTypeJsonFieldSettings struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	// The default value of the JSON field.
	Default any `json:"default,omitzero"`
	// A JSON schema used to validate the structure of the JSON provided. Must be a
	// valid JSON schema.
	Schema any `json:"schema,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeJsonFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeJsonFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeJsonFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A markdown field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialValidateParamsPartialInputSchemaMessageTypeMarkdownField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "markdown".
	Type string `json:"type,omitzero,required"`
	// Settings for the markdown field.
	Settings PartialValidateParamsPartialInputSchemaMessageTypeMarkdownFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeMarkdownField) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeMarkdownField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeMarkdownField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartialInputSchemaMessageTypeMarkdownField](
		"type", "markdown",
	)
}

// Settings for the markdown field.
type PartialValidateParamsPartialInputSchemaMessageTypeMarkdownFieldSettings struct {
	// The default value of the markdown field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeMarkdownFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeMarkdownFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeMarkdownFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A multi-select field used in a message type.
//
// The properties Key, Label, Settings, Type are required.
type PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// Settings for the multi_select field.
	Settings PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings `json:"settings,omitzero,required"`
	// The type of the field.
	//
	// Any of "multi_select".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectField) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectField](
		"type", "multi_select",
	)
}

// Settings for the multi_select field.
type PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	// The default values for the multi-select field.
	Default []string `json:"default,omitzero"`
	// The available options for the multi-select field.
	Options []PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption `json:"options,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Value is required.
type PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption struct {
	// The value for the option.
	Value string `json:"value,required"`
	// The display label for the option.
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeMultiSelectFieldSettingsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A select field used in a message type.
//
// The properties Key, Label, Settings, Type are required.
type PartialValidateParamsPartialInputSchemaMessageTypeSelectField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// Settings for the select field.
	Settings PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettings `json:"settings,omitzero,required"`
	// The type of the field.
	//
	// Any of "select".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeSelectField) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeSelectField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeSelectField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartialInputSchemaMessageTypeSelectField](
		"type", "select",
	)
}

// Settings for the select field.
type PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettings struct {
	// The default value for the select field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	// The available options for the select field.
	Options []PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption `json:"options,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Value is required.
type PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption struct {
	// The value for the option.
	Value string `json:"value,required"`
	// The display label for the option.
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeSelectFieldSettingsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A textarea field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialValidateParamsPartialInputSchemaMessageTypeTextareaField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "textarea".
	Type string `json:"type,omitzero,required"`
	// Settings for the textarea field.
	Settings PartialValidateParamsPartialInputSchemaMessageTypeTextareaFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeTextareaField) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeTextareaField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeTextareaField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartialInputSchemaMessageTypeTextareaField](
		"type", "textarea",
	)
}

// Settings for the textarea field.
type PartialValidateParamsPartialInputSchemaMessageTypeTextareaFieldSettings struct {
	// The default value of the textarea field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	MaxLength   param.Opt[int64]  `json:"max_length,omitzero"`
	MinLength   param.Opt[int64]  `json:"min_length,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeTextareaFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeTextareaFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeTextareaFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL field used in a message type.
//
// The properties Key, Label, Type are required.
type PartialValidateParamsPartialInputSchemaMessageTypeURLField struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "url".
	Type string `json:"type,omitzero,required"`
	// Settings for the url field.
	Settings PartialValidateParamsPartialInputSchemaMessageTypeURLFieldSettings `json:"settings,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeURLField) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeURLField
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeURLField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PartialValidateParamsPartialInputSchemaMessageTypeURLField](
		"type", "url",
	)
}

// Settings for the url field.
type PartialValidateParamsPartialInputSchemaMessageTypeURLFieldSettings struct {
	// The default value of the URL field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r PartialValidateParamsPartialInputSchemaMessageTypeURLFieldSettings) MarshalJSON() (data []byte, err error) {
	type shadow PartialValidateParamsPartialInputSchemaMessageTypeURLFieldSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartialValidateParamsPartialInputSchemaMessageTypeURLFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
