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
	// This field is a union of [bool], [string], [[]string], [string], [string],
	// [string], [string]
	Default     PartialInputSchemaUnionSettingsDefault `json:"default"`
	Description string                                 `json:"description"`
	Required    bool                                   `json:"required"`
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
// will be valid: OfBool OfString OfStringArray]
type PartialInputSchemaUnionSettingsDefault struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfBool        respjson.Field
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
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
