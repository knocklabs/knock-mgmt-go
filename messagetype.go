// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/knock-mapi-go/internal/apijson"
	"github.com/stainless-sdks/knock-mapi-go/internal/apiquery"
	"github.com/stainless-sdks/knock-mapi-go/internal/paramutil"
	"github.com/stainless-sdks/knock-mapi-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-mapi-go/option"
	"github.com/stainless-sdks/knock-mapi-go/packages/pagination"
	"github.com/stainless-sdks/knock-mapi-go/packages/param"
	"github.com/stainless-sdks/knock-mapi-go/packages/respjson"
)

// MessageTypeService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageTypeService] method instead.
type MessageTypeService struct {
	Options []option.RequestOption
}

// NewMessageTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessageTypeService(opts ...option.RequestOption) (r MessageTypeService) {
	r = MessageTypeService{}
	r.Options = opts
	return
}

// Retrieve a message type by its key, in a given environment.
func (r *MessageTypeService) Get(ctx context.Context, messageTypeKey string, query MessageTypeGetParams, opts ...option.RequestOption) (res *MessageType, err error) {
	opts = append(r.Options[:], opts...)
	if messageTypeKey == "" {
		err = errors.New("missing required message_type_key parameter")
		return
	}
	path := fmt.Sprintf("v1/message_types/%s", messageTypeKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a paginated list of message types available in a given environment.
func (r *MessageTypeService) List(ctx context.Context, query MessageTypeListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[MessageType], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/message_types"
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

// Returns a paginated list of message types available in a given environment.
func (r *MessageTypeService) ListAutoPaging(ctx context.Context, query MessageTypeListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[MessageType] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Updates a message type, or creates a new one if it does not yet exist.
//
// Note: this endpoint only operates in the `development` environment.
func (r *MessageTypeService) Upsert(ctx context.Context, messageTypeKey string, params MessageTypeUpsertParams, opts ...option.RequestOption) (res *MessageTypeUpsertResponse, err error) {
	opts = append(r.Options[:], opts...)
	if messageTypeKey == "" {
		err = errors.New("missing required message_type_key parameter")
		return
	}
	path := fmt.Sprintf("v1/message_types/%s", messageTypeKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Validates a message type payload without persisting it.
//
// Note: this endpoint only operates on message types in the `development`
// environment.
func (r *MessageTypeService) Validate(ctx context.Context, messageTypeKey string, params MessageTypeValidateParams, opts ...option.RequestOption) (res *MessageTypeValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if messageTypeKey == "" {
		err = errors.New("missing required message_type_key parameter")
		return
	}
	path := fmt.Sprintf("v1/message_types/%s/validate", messageTypeKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// A message type is a schema for a message that maps to a UI component or element
// within your application.
type MessageType struct {
	// The timestamp of when the message type was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The environment of the message type.
	Environment string `json:"environment,required"`
	// The unique key string for the message type object. Must be at minimum 3
	// characters and at maximum 255 characters in length. Must be in the format of
	// ^[a-z0-9_-]+$.
	Key string `json:"key,required"`
	// A name for the message type. Must be at maximum 255 characters in length.
	Name string `json:"name,required"`
	// The owner of the message type.
	//
	// Any of "system", "user".
	Owner MessageTypeOwner `json:"owner,required"`
	// An HTML/liquid template for the message type preview.
	Preview string `json:"preview,required"`
	// The semantic version of the message type.
	Semver string `json:"semver,required"`
	// The SHA hash of the message type.
	Sha string `json:"sha,required"`
	// The timestamp of when the message type was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Whether the message type is valid.
	Valid bool `json:"valid,required"`
	// The variants of the message type.
	Variants []MessageTypeVariant `json:"variants,required"`
	// The timestamp of when the message type was deleted.
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// The timestamp of when the message type was deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// An arbitrary string attached to a message type object. Useful for adding notes
	// about the message type for internal purposes. Maximum of 280 characters allowed.
	Description string `json:"description,nullable"`
	// The icon name of the message type.
	IconName string `json:"icon_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Environment respjson.Field
		Key         respjson.Field
		Name        respjson.Field
		Owner       respjson.Field
		Preview     respjson.Field
		Semver      respjson.Field
		Sha         respjson.Field
		UpdatedAt   respjson.Field
		Valid       respjson.Field
		Variants    respjson.Field
		ArchivedAt  respjson.Field
		DeletedAt   respjson.Field
		Description respjson.Field
		IconName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageType) RawJSON() string { return r.JSON.raw }
func (r *MessageType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The owner of the message type.
type MessageTypeOwner string

const (
	MessageTypeOwnerSystem MessageTypeOwner = "system"
	MessageTypeOwnerUser   MessageTypeOwner = "user"
)

// A text field used in a message type.
type MessageTypeTextField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// The type of the field.
	//
	// Any of "text".
	Type MessageTypeTextFieldType `json:"type,required"`
	// Settings for the text field.
	Settings MessageTypeTextFieldSettings `json:"settings"`
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
func (r MessageTypeTextField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeTextField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageTypeTextField to a MessageTypeTextFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageTypeTextFieldParam.Overrides()
func (r MessageTypeTextField) ToParam() MessageTypeTextFieldParam {
	return param.Override[MessageTypeTextFieldParam](json.RawMessage(r.RawJSON()))
}

// The type of the field.
type MessageTypeTextFieldType string

const (
	MessageTypeTextFieldTypeText MessageTypeTextFieldType = "text"
)

// Settings for the text field.
type MessageTypeTextFieldSettings struct {
	// The default value of the text field.
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
func (r MessageTypeTextFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeTextFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text field used in a message type.
//
// The properties Key, Label, Type are required.
type MessageTypeTextFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "text".
	Type MessageTypeTextFieldType `json:"type,omitzero,required"`
	// Settings for the text field.
	Settings MessageTypeTextFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeTextFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeTextFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeTextFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the text field.
type MessageTypeTextFieldSettingsParam struct {
	// The default value of the text field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	MaxLength   param.Opt[int64]  `json:"max_length,omitzero"`
	MinLength   param.Opt[int64]  `json:"min_length,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeTextFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeTextFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeTextFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A variant of a message type.
type MessageTypeVariant struct {
	// The field types available for the variant.
	Fields []MessageTypeVariantFieldUnion `json:"fields,required"`
	// The unique key string for the variant. Must be at minimum 3 characters and at
	// maximum 255 characters in length. Must be in the format of ^[a-z0-9_-]+$.
	Key string `json:"key,required"`
	// A name for the variant. Must be at maximum 255 characters in length.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fields      respjson.Field
		Key         respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageTypeVariant) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageTypeVariant to a MessageTypeVariantParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageTypeVariantParam.Overrides()
func (r MessageTypeVariant) ToParam() MessageTypeVariantParam {
	return param.Override[MessageTypeVariantParam](json.RawMessage(r.RawJSON()))
}

// MessageTypeVariantFieldUnion contains all possible properties and values from
// [MessageTypeVariantFieldMessageTypeBooleanField],
// [MessageTypeVariantFieldMessageTypeButtonField],
// [MessageTypeVariantFieldMessageTypeImageField],
// [MessageTypeVariantFieldMessageTypeMarkdownField],
// [MessageTypeVariantFieldMessageTypeMultiSelectField],
// [MessageTypeVariantFieldMessageTypeSelectField], [MessageTypeTextField],
// [MessageTypeVariantFieldMessageTypeTextareaField],
// [MessageTypeVariantFieldMessageTypeURLField].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessageTypeVariantFieldUnion struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Type  string `json:"type"`
	// This field is a union of
	// [MessageTypeVariantFieldMessageTypeBooleanFieldSettings],
	// [MessageTypeVariantFieldMessageTypeButtonFieldSettings],
	// [MessageTypeVariantFieldMessageTypeImageFieldSettings],
	// [MessageTypeVariantFieldMessageTypeMarkdownFieldSettings],
	// [MessageTypeVariantFieldMessageTypeMultiSelectFieldSettings],
	// [MessageTypeVariantFieldMessageTypeSelectFieldSettings],
	// [MessageTypeTextFieldSettings],
	// [MessageTypeVariantFieldMessageTypeTextareaFieldSettings],
	// [MessageTypeVariantFieldMessageTypeURLFieldSettings]
	Settings MessageTypeVariantFieldUnionSettings `json:"settings"`
	// This field is from variant [MessageTypeVariantFieldMessageTypeButtonField].
	Action MessageTypeTextField `json:"action"`
	// This field is from variant [MessageTypeVariantFieldMessageTypeButtonField].
	Text MessageTypeTextField `json:"text"`
	// This field is from variant [MessageTypeVariantFieldMessageTypeImageField].
	Alt MessageTypeTextField `json:"alt"`
	// This field is from variant [MessageTypeVariantFieldMessageTypeImageField].
	URL  MessageTypeVariantFieldMessageTypeImageFieldURL `json:"url"`
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

func (u MessageTypeVariantFieldUnion) AsMessageTypeBooleanField() (v MessageTypeVariantFieldMessageTypeBooleanField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeButtonField() (v MessageTypeVariantFieldMessageTypeButtonField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeImageField() (v MessageTypeVariantFieldMessageTypeImageField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeMarkdownField() (v MessageTypeVariantFieldMessageTypeMarkdownField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeMultiSelectField() (v MessageTypeVariantFieldMessageTypeMultiSelectField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeSelectField() (v MessageTypeVariantFieldMessageTypeSelectField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeTextField() (v MessageTypeTextField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeTextareaField() (v MessageTypeVariantFieldMessageTypeTextareaField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeURLField() (v MessageTypeVariantFieldMessageTypeURLField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageTypeVariantFieldUnion) RawJSON() string { return u.JSON.raw }

func (r *MessageTypeVariantFieldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageTypeVariantFieldUnionSettings is an implicit subunion of
// [MessageTypeVariantFieldUnion]. MessageTypeVariantFieldUnionSettings provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageTypeVariantFieldUnion].
type MessageTypeVariantFieldUnionSettings struct {
	// This field is a union of [bool], [string], [[]string], [string], [string],
	// [string], [string]
	Default     MessageTypeVariantFieldUnionSettingsDefault `json:"default"`
	Description string                                      `json:"description"`
	Required    bool                                        `json:"required"`
	// This field is a union of
	// [[]MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOption],
	// [[]MessageTypeVariantFieldMessageTypeSelectFieldSettingsOption]
	Options   MessageTypeVariantFieldUnionSettingsOptions `json:"options"`
	MaxLength int64                                       `json:"max_length"`
	MinLength int64                                       `json:"min_length"`
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

func (r *MessageTypeVariantFieldUnionSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageTypeVariantFieldUnionSettingsDefault is an implicit subunion of
// [MessageTypeVariantFieldUnion]. MessageTypeVariantFieldUnionSettingsDefault
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageTypeVariantFieldUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfString OfStringArray]
type MessageTypeVariantFieldUnionSettingsDefault struct {
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

func (r *MessageTypeVariantFieldUnionSettingsDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageTypeVariantFieldUnionSettingsOptions is an implicit subunion of
// [MessageTypeVariantFieldUnion]. MessageTypeVariantFieldUnionSettingsOptions
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageTypeVariantFieldUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfMessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOptions
// OfMessageTypeVariantFieldMessageTypeSelectFieldSettingsOptions]
type MessageTypeVariantFieldUnionSettingsOptions struct {
	// This field will be present if the value is a
	// [[]MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOption] instead of
	// an object.
	OfMessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOptions []MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOption `json:",inline"`
	// This field will be present if the value is a
	// [[]MessageTypeVariantFieldMessageTypeSelectFieldSettingsOption] instead of an
	// object.
	OfMessageTypeVariantFieldMessageTypeSelectFieldSettingsOptions []MessageTypeVariantFieldMessageTypeSelectFieldSettingsOption `json:",inline"`
	JSON                                                           struct {
		OfMessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOptions respjson.Field
		OfMessageTypeVariantFieldMessageTypeSelectFieldSettingsOptions      respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (r *MessageTypeVariantFieldUnionSettingsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A boolean field used in a message type.
type MessageTypeVariantFieldMessageTypeBooleanField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// The type of the field.
	//
	// Any of "boolean".
	Type string `json:"type,required"`
	// Settings for the boolean field.
	Settings MessageTypeVariantFieldMessageTypeBooleanFieldSettings `json:"settings"`
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
func (r MessageTypeVariantFieldMessageTypeBooleanField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeBooleanField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the boolean field.
type MessageTypeVariantFieldMessageTypeBooleanFieldSettings struct {
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
func (r MessageTypeVariantFieldMessageTypeBooleanFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeBooleanFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A button field used in a message type.
type MessageTypeVariantFieldMessageTypeButtonField struct {
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
	Settings MessageTypeVariantFieldMessageTypeButtonFieldSettings `json:"settings"`
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
func (r MessageTypeVariantFieldMessageTypeButtonField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeButtonField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the button field.
type MessageTypeVariantFieldMessageTypeButtonFieldSettings struct {
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
func (r MessageTypeVariantFieldMessageTypeButtonFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeButtonFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image field used in a message type.
type MessageTypeVariantFieldMessageTypeImageField struct {
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
	URL MessageTypeVariantFieldMessageTypeImageFieldURL `json:"url,required"`
	// Settings for the image field.
	Settings MessageTypeVariantFieldMessageTypeImageFieldSettings `json:"settings"`
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
func (r MessageTypeVariantFieldMessageTypeImageField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeImageField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL field used in a message type.
type MessageTypeVariantFieldMessageTypeImageFieldURL struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// The type of the field.
	//
	// Any of "url".
	Type string `json:"type,required"`
	// Settings for the url field.
	Settings MessageTypeVariantFieldMessageTypeImageFieldURLSettings `json:"settings"`
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
func (r MessageTypeVariantFieldMessageTypeImageFieldURL) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeImageFieldURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the url field.
type MessageTypeVariantFieldMessageTypeImageFieldURLSettings struct {
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
func (r MessageTypeVariantFieldMessageTypeImageFieldURLSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeImageFieldURLSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the image field.
type MessageTypeVariantFieldMessageTypeImageFieldSettings struct {
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
func (r MessageTypeVariantFieldMessageTypeImageFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeImageFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A markdown field used in a message type.
type MessageTypeVariantFieldMessageTypeMarkdownField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// The type of the field.
	//
	// Any of "markdown".
	Type string `json:"type,required"`
	// Settings for the markdown field.
	Settings MessageTypeVariantFieldMessageTypeMarkdownFieldSettings `json:"settings"`
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
func (r MessageTypeVariantFieldMessageTypeMarkdownField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeMarkdownField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the markdown field.
type MessageTypeVariantFieldMessageTypeMarkdownFieldSettings struct {
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
func (r MessageTypeVariantFieldMessageTypeMarkdownFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeMarkdownFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A multi-select field used in a message type.
type MessageTypeVariantFieldMessageTypeMultiSelectField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// Settings for the multi_select field.
	Settings MessageTypeVariantFieldMessageTypeMultiSelectFieldSettings `json:"settings,required"`
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
func (r MessageTypeVariantFieldMessageTypeMultiSelectField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeMultiSelectField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the multi_select field.
type MessageTypeVariantFieldMessageTypeMultiSelectFieldSettings struct {
	// The default values for the multi-select field.
	Default     []string `json:"default,nullable"`
	Description string   `json:"description"`
	// The available options for the multi-select field.
	Options []MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOption `json:"options"`
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
func (r MessageTypeVariantFieldMessageTypeMultiSelectFieldSettings) RawJSON() string {
	return r.JSON.raw
}
func (r *MessageTypeVariantFieldMessageTypeMultiSelectFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOption struct {
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
func (r MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOption) RawJSON() string {
	return r.JSON.raw
}
func (r *MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A select field used in a message type.
type MessageTypeVariantFieldMessageTypeSelectField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// Settings for the select field.
	Settings MessageTypeVariantFieldMessageTypeSelectFieldSettings `json:"settings,required"`
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
func (r MessageTypeVariantFieldMessageTypeSelectField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeSelectField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the select field.
type MessageTypeVariantFieldMessageTypeSelectFieldSettings struct {
	// The default value for the select field.
	Default     string `json:"default,nullable"`
	Description string `json:"description"`
	// The available options for the select field.
	Options []MessageTypeVariantFieldMessageTypeSelectFieldSettingsOption `json:"options"`
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
func (r MessageTypeVariantFieldMessageTypeSelectFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeSelectFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageTypeVariantFieldMessageTypeSelectFieldSettingsOption struct {
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
func (r MessageTypeVariantFieldMessageTypeSelectFieldSettingsOption) RawJSON() string {
	return r.JSON.raw
}
func (r *MessageTypeVariantFieldMessageTypeSelectFieldSettingsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A textarea field used in a message type.
type MessageTypeVariantFieldMessageTypeTextareaField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// The type of the field.
	//
	// Any of "textarea".
	Type string `json:"type,required"`
	// Settings for the textarea field.
	Settings MessageTypeVariantFieldMessageTypeTextareaFieldSettings `json:"settings"`
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
func (r MessageTypeVariantFieldMessageTypeTextareaField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeTextareaField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the textarea field.
type MessageTypeVariantFieldMessageTypeTextareaFieldSettings struct {
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
func (r MessageTypeVariantFieldMessageTypeTextareaFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeTextareaFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL field used in a message type.
type MessageTypeVariantFieldMessageTypeURLField struct {
	// The unique key of the field.
	Key string `json:"key,required"`
	// The label of the field.
	Label string `json:"label,required"`
	// The type of the field.
	//
	// Any of "url".
	Type string `json:"type,required"`
	// Settings for the url field.
	Settings MessageTypeVariantFieldMessageTypeURLFieldSettings `json:"settings"`
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
func (r MessageTypeVariantFieldMessageTypeURLField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeURLField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the url field.
type MessageTypeVariantFieldMessageTypeURLFieldSettings struct {
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
func (r MessageTypeVariantFieldMessageTypeURLFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeVariantFieldMessageTypeURLFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A variant of a message type.
//
// The properties Fields, Key, Name are required.
type MessageTypeVariantParam struct {
	// The field types available for the variant.
	Fields []MessageTypeVariantFieldUnionParam `json:"fields,omitzero,required"`
	// The unique key string for the variant. Must be at minimum 3 characters and at
	// maximum 255 characters in length. Must be in the format of ^[a-z0-9_-]+$.
	Key string `json:"key,required"`
	// A name for the variant. Must be at maximum 255 characters in length.
	Name string `json:"name,required"`
	paramObj
}

func (r MessageTypeVariantParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageTypeVariantFieldUnionParam struct {
	OfMessageTypeBooleanField     *MessageTypeVariantFieldMessageTypeBooleanFieldParam     `json:",omitzero,inline"`
	OfMessageTypeButtonField      *MessageTypeVariantFieldMessageTypeButtonFieldParam      `json:",omitzero,inline"`
	OfMessageTypeImageField       *MessageTypeVariantFieldMessageTypeImageFieldParam       `json:",omitzero,inline"`
	OfMessageTypeMarkdownField    *MessageTypeVariantFieldMessageTypeMarkdownFieldParam    `json:",omitzero,inline"`
	OfMessageTypeMultiSelectField *MessageTypeVariantFieldMessageTypeMultiSelectFieldParam `json:",omitzero,inline"`
	OfMessageTypeSelectField      *MessageTypeVariantFieldMessageTypeSelectFieldParam      `json:",omitzero,inline"`
	OfMessageTypeTextField        *MessageTypeTextFieldParam                               `json:",omitzero,inline"`
	OfMessageTypeTextareaField    *MessageTypeVariantFieldMessageTypeTextareaFieldParam    `json:",omitzero,inline"`
	OfMessageTypeURLField         *MessageTypeVariantFieldMessageTypeURLFieldParam         `json:",omitzero,inline"`
	paramUnion
}

func (u MessageTypeVariantFieldUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMessageTypeBooleanField,
		u.OfMessageTypeButtonField,
		u.OfMessageTypeImageField,
		u.OfMessageTypeMarkdownField,
		u.OfMessageTypeMultiSelectField,
		u.OfMessageTypeSelectField,
		u.OfMessageTypeTextField,
		u.OfMessageTypeTextareaField,
		u.OfMessageTypeURLField)
}
func (u *MessageTypeVariantFieldUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageTypeVariantFieldUnionParam) asAny() any {
	if !param.IsOmitted(u.OfMessageTypeBooleanField) {
		return u.OfMessageTypeBooleanField
	} else if !param.IsOmitted(u.OfMessageTypeButtonField) {
		return u.OfMessageTypeButtonField
	} else if !param.IsOmitted(u.OfMessageTypeImageField) {
		return u.OfMessageTypeImageField
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
func (u MessageTypeVariantFieldUnionParam) GetText() *MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeButtonField; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageTypeVariantFieldUnionParam) GetAlt() *MessageTypeTextFieldParam {
	if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.Alt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageTypeVariantFieldUnionParam) GetURL() *MessageTypeVariantFieldMessageTypeImageFieldURLParam {
	if vt := u.OfMessageTypeImageField; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageTypeVariantFieldUnionParam) GetKey() *string {
	if vt := u.OfMessageTypeBooleanField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfMessageTypeImageField; vt != nil {
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
func (u MessageTypeVariantFieldUnionParam) GetLabel() *string {
	if vt := u.OfMessageTypeBooleanField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeButtonField; vt != nil && vt.Label.Valid() {
		return &vt.Label.Value
	} else if vt := u.OfMessageTypeImageField; vt != nil && vt.Label.Valid() {
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
func (u MessageTypeVariantFieldUnionParam) GetType() *string {
	if vt := u.OfMessageTypeBooleanField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMessageTypeImageField; vt != nil {
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
func (u MessageTypeVariantFieldUnionParam) GetSettings() (res messageTypeVariantFieldUnionParamSettings) {
	if vt := u.OfMessageTypeBooleanField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeButtonField; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfMessageTypeImageField; vt != nil {
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
// [*MessageTypeVariantFieldMessageTypeBooleanFieldSettingsParam],
// [*MessageTypeVariantFieldMessageTypeButtonFieldSettingsParam],
// [*MessageTypeVariantFieldMessageTypeImageFieldSettingsParam],
// [*MessageTypeVariantFieldMessageTypeMarkdownFieldSettingsParam],
// [*MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsParam],
// [*MessageTypeVariantFieldMessageTypeSelectFieldSettingsParam],
// [*MessageTypeTextFieldSettingsParam],
// [*MessageTypeVariantFieldMessageTypeTextareaFieldSettingsParam],
// [*MessageTypeVariantFieldMessageTypeURLFieldSettingsParam]
type messageTypeVariantFieldUnionParamSettings struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.MessageTypeVariantFieldMessageTypeBooleanFieldSettingsParam:
//	case *knockmapi.MessageTypeVariantFieldMessageTypeButtonFieldSettingsParam:
//	case *knockmapi.MessageTypeVariantFieldMessageTypeImageFieldSettingsParam:
//	case *knockmapi.MessageTypeVariantFieldMessageTypeMarkdownFieldSettingsParam:
//	case *knockmapi.MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsParam:
//	case *knockmapi.MessageTypeVariantFieldMessageTypeSelectFieldSettingsParam:
//	case *knockmapi.MessageTypeTextFieldSettingsParam:
//	case *knockmapi.MessageTypeVariantFieldMessageTypeTextareaFieldSettingsParam:
//	case *knockmapi.MessageTypeVariantFieldMessageTypeURLFieldSettingsParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u messageTypeVariantFieldUnionParamSettings) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u messageTypeVariantFieldUnionParamSettings) GetDescription() *string {
	switch vt := u.any.(type) {
	case *MessageTypeVariantFieldMessageTypeBooleanFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *MessageTypeVariantFieldMessageTypeButtonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *MessageTypeVariantFieldMessageTypeImageFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *MessageTypeVariantFieldMessageTypeMarkdownFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *MessageTypeVariantFieldMessageTypeSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *MessageTypeVariantFieldMessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *MessageTypeVariantFieldMessageTypeURLFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u messageTypeVariantFieldUnionParamSettings) GetRequired() *bool {
	switch vt := u.any.(type) {
	case *MessageTypeVariantFieldMessageTypeBooleanFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *MessageTypeVariantFieldMessageTypeButtonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *MessageTypeVariantFieldMessageTypeImageFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *MessageTypeVariantFieldMessageTypeMarkdownFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *MessageTypeVariantFieldMessageTypeSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *MessageTypeVariantFieldMessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *MessageTypeVariantFieldMessageTypeURLFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u messageTypeVariantFieldUnionParamSettings) GetMaxLength() *int64 {
	switch vt := u.any.(type) {
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MaxLength)
	case *MessageTypeVariantFieldMessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MaxLength)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u messageTypeVariantFieldUnionParamSettings) GetMinLength() *int64 {
	switch vt := u.any.(type) {
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MinLength)
	case *MessageTypeVariantFieldMessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MinLength)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u messageTypeVariantFieldUnionParamSettings) GetDefault() (res messageTypeVariantFieldUnionParamSettingsDefault) {
	switch vt := u.any.(type) {
	case *MessageTypeVariantFieldMessageTypeBooleanFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *MessageTypeVariantFieldMessageTypeMarkdownFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsParam:
		res.any = &vt.Default
	case *MessageTypeVariantFieldMessageTypeSelectFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *MessageTypeTextFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *MessageTypeVariantFieldMessageTypeTextareaFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *MessageTypeVariantFieldMessageTypeURLFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	}
	return res
}

// Can have the runtime types [*bool], [*string], [\*[]string]
type messageTypeVariantFieldUnionParamSettingsDefault struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *bool:
//	case *string:
//	case *[]string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u messageTypeVariantFieldUnionParamSettingsDefault) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u messageTypeVariantFieldUnionParamSettings) GetOptions() (res messageTypeVariantFieldUnionParamSettingsOptions) {
	switch vt := u.any.(type) {
	case *MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsParam:
		res.any = &vt.Options
	case *MessageTypeVariantFieldMessageTypeSelectFieldSettingsParam:
		res.any = &vt.Options
	}
	return res
}

// Can have the runtime types
// [_[]MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOptionParam],
// [_[]MessageTypeVariantFieldMessageTypeSelectFieldSettingsOptionParam]
type messageTypeVariantFieldUnionParamSettingsOptions struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]knockmapi.MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOptionParam:
//	case *[]knockmapi.MessageTypeVariantFieldMessageTypeSelectFieldSettingsOptionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u messageTypeVariantFieldUnionParamSettingsOptions) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's Action property, if present.
func (u MessageTypeVariantFieldUnionParam) GetAction() *MessageTypeTextFieldParam {
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
type MessageTypeVariantFieldMessageTypeBooleanFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "boolean".
	Type string `json:"type,omitzero,required"`
	// Settings for the boolean field.
	Settings MessageTypeVariantFieldMessageTypeBooleanFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeBooleanFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeBooleanFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeBooleanFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageTypeVariantFieldMessageTypeBooleanFieldParam](
		"type", "boolean",
	)
}

// Settings for the boolean field.
type MessageTypeVariantFieldMessageTypeBooleanFieldSettingsParam struct {
	// The default value of the boolean field.
	Default     param.Opt[bool]   `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeBooleanFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeBooleanFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeBooleanFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A button field used in a message type.
//
// The properties Action, Key, Label, Text, Type are required.
type MessageTypeVariantFieldMessageTypeButtonFieldParam struct {
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
	Settings MessageTypeVariantFieldMessageTypeButtonFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeButtonFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeButtonFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeButtonFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageTypeVariantFieldMessageTypeButtonFieldParam](
		"type", "button",
	)
}

// Settings for the button field.
type MessageTypeVariantFieldMessageTypeButtonFieldSettingsParam struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeButtonFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeButtonFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeButtonFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image field used in a message type.
//
// The properties Action, Alt, Key, Label, Type, URL are required.
type MessageTypeVariantFieldMessageTypeImageFieldParam struct {
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
	URL MessageTypeVariantFieldMessageTypeImageFieldURLParam `json:"url,omitzero,required"`
	// Settings for the image field.
	Settings MessageTypeVariantFieldMessageTypeImageFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeImageFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeImageFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeImageFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageTypeVariantFieldMessageTypeImageFieldParam](
		"type", "image",
	)
}

// A URL field used in a message type.
//
// The properties Key, Label, Type are required.
type MessageTypeVariantFieldMessageTypeImageFieldURLParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "url".
	Type string `json:"type,omitzero,required"`
	// Settings for the url field.
	Settings MessageTypeVariantFieldMessageTypeImageFieldURLSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeImageFieldURLParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeImageFieldURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeImageFieldURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageTypeVariantFieldMessageTypeImageFieldURLParam](
		"type", "url",
	)
}

// Settings for the url field.
type MessageTypeVariantFieldMessageTypeImageFieldURLSettingsParam struct {
	// The default value of the URL field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeImageFieldURLSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeImageFieldURLSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeImageFieldURLSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the image field.
type MessageTypeVariantFieldMessageTypeImageFieldSettingsParam struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeImageFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeImageFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeImageFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A markdown field used in a message type.
//
// The properties Key, Label, Type are required.
type MessageTypeVariantFieldMessageTypeMarkdownFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "markdown".
	Type string `json:"type,omitzero,required"`
	// Settings for the markdown field.
	Settings MessageTypeVariantFieldMessageTypeMarkdownFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeMarkdownFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeMarkdownFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeMarkdownFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageTypeVariantFieldMessageTypeMarkdownFieldParam](
		"type", "markdown",
	)
}

// Settings for the markdown field.
type MessageTypeVariantFieldMessageTypeMarkdownFieldSettingsParam struct {
	// The default value of the markdown field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeMarkdownFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeMarkdownFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeMarkdownFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A multi-select field used in a message type.
//
// The properties Key, Label, Settings, Type are required.
type MessageTypeVariantFieldMessageTypeMultiSelectFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// Settings for the multi_select field.
	Settings MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsParam `json:"settings,omitzero,required"`
	// The type of the field.
	//
	// Any of "multi_select".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeMultiSelectFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeMultiSelectFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeMultiSelectFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageTypeVariantFieldMessageTypeMultiSelectFieldParam](
		"type", "multi_select",
	)
}

// Settings for the multi_select field.
type MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsParam struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	// The default values for the multi-select field.
	Default []string `json:"default,omitzero"`
	// The available options for the multi-select field.
	Options []MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOptionParam `json:"options,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Value is required.
type MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOptionParam struct {
	// The value for the option.
	Value string `json:"value,required"`
	// The display label for the option.
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOptionParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeMultiSelectFieldSettingsOptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A select field used in a message type.
//
// The properties Key, Label, Settings, Type are required.
type MessageTypeVariantFieldMessageTypeSelectFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// Settings for the select field.
	Settings MessageTypeVariantFieldMessageTypeSelectFieldSettingsParam `json:"settings,omitzero,required"`
	// The type of the field.
	//
	// Any of "select".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeSelectFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeSelectFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeSelectFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageTypeVariantFieldMessageTypeSelectFieldParam](
		"type", "select",
	)
}

// Settings for the select field.
type MessageTypeVariantFieldMessageTypeSelectFieldSettingsParam struct {
	// The default value for the select field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	// The available options for the select field.
	Options []MessageTypeVariantFieldMessageTypeSelectFieldSettingsOptionParam `json:"options,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeSelectFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeSelectFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeSelectFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Value is required.
type MessageTypeVariantFieldMessageTypeSelectFieldSettingsOptionParam struct {
	// The value for the option.
	Value string `json:"value,required"`
	// The display label for the option.
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeSelectFieldSettingsOptionParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeSelectFieldSettingsOptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeSelectFieldSettingsOptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A textarea field used in a message type.
//
// The properties Key, Label, Type are required.
type MessageTypeVariantFieldMessageTypeTextareaFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "textarea".
	Type string `json:"type,omitzero,required"`
	// Settings for the textarea field.
	Settings MessageTypeVariantFieldMessageTypeTextareaFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeTextareaFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeTextareaFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeTextareaFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageTypeVariantFieldMessageTypeTextareaFieldParam](
		"type", "textarea",
	)
}

// Settings for the textarea field.
type MessageTypeVariantFieldMessageTypeTextareaFieldSettingsParam struct {
	// The default value of the textarea field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	MaxLength   param.Opt[int64]  `json:"max_length,omitzero"`
	MinLength   param.Opt[int64]  `json:"min_length,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeTextareaFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeTextareaFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeTextareaFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL field used in a message type.
//
// The properties Key, Label, Type are required.
type MessageTypeVariantFieldMessageTypeURLFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero,required"`
	// The unique key of the field.
	Key string `json:"key,required"`
	// The type of the field.
	//
	// Any of "url".
	Type string `json:"type,omitzero,required"`
	// Settings for the url field.
	Settings MessageTypeVariantFieldMessageTypeURLFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeURLFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeURLFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeURLFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageTypeVariantFieldMessageTypeURLFieldParam](
		"type", "url",
	)
}

// Settings for the url field.
type MessageTypeVariantFieldMessageTypeURLFieldSettingsParam struct {
	// The default value of the URL field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeVariantFieldMessageTypeURLFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeVariantFieldMessageTypeURLFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeVariantFieldMessageTypeURLFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the MessageType response under the `message_type` key.
type MessageTypeUpsertResponse struct {
	// A message type is a schema for a message that maps to a UI component or element
	// within your application.
	MessageType MessageType `json:"message_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageTypeUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the MessageType response under the `message_type` key.
type MessageTypeValidateResponse struct {
	// A message type is a schema for a message that maps to a UI component or element
	// within your application.
	MessageType MessageType `json:"message_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageTypeValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageTypeGetParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// Whether to hide uncommitted changes. When true, only committed changes will be
	// returned. When false, both committed and uncommitted changes will be returned.
	HideUncommittedChanges param.Opt[bool] `query:"hide_uncommitted_changes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessageTypeGetParams]'s query parameters as `url.Values`.
func (r MessageTypeGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageTypeListParams struct {
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

// URLQuery serializes [MessageTypeListParams]'s query parameters as `url.Values`.
func (r MessageTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageTypeUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// A request to create a message type.
	MessageType MessageTypeUpsertParamsMessageType `json:"message_type,omitzero,required"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// Whether to commit the resource at the same time as modifying it.
	Commit param.Opt[bool] `query:"commit,omitzero" json:"-"`
	// The message to commit the resource with, only used if `commit` is `true`.
	CommitMessage param.Opt[string] `query:"commit_message,omitzero" json:"-"`
	paramObj
}

func (r MessageTypeUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessageTypeUpsertParams]'s query parameters as
// `url.Values`.
func (r MessageTypeUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A request to create a message type.
//
// The properties Description, Name, Preview are required.
type MessageTypeUpsertParamsMessageType struct {
	// An arbitrary string attached to a message type object. Useful for adding notes
	// about the message type for internal purposes. Maximum of 280 characters allowed.
	Description param.Opt[string] `json:"description,omitzero,required"`
	// A name for the message type. Must be at maximum 255 characters in length.
	Name string `json:"name,required"`
	// An HTML/liquid template for the message type preview.
	Preview string `json:"preview,required"`
	// The icon name of the message type.
	IconName param.Opt[string] `json:"icon_name,omitzero"`
	// The semantic version of the message type.
	Semver param.Opt[string] `json:"semver,omitzero"`
	// The variants of the message type.
	Variants []MessageTypeVariantParam `json:"variants,omitzero"`
	paramObj
}

func (r MessageTypeUpsertParamsMessageType) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeUpsertParamsMessageType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeUpsertParamsMessageType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageTypeValidateParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// A request to create a message type.
	MessageType MessageTypeValidateParamsMessageType `json:"message_type,omitzero,required"`
	paramObj
}

func (r MessageTypeValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessageTypeValidateParams]'s query parameters as
// `url.Values`.
func (r MessageTypeValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A request to create a message type.
//
// The properties Description, Name, Preview are required.
type MessageTypeValidateParamsMessageType struct {
	// An arbitrary string attached to a message type object. Useful for adding notes
	// about the message type for internal purposes. Maximum of 280 characters allowed.
	Description param.Opt[string] `json:"description,omitzero,required"`
	// A name for the message type. Must be at maximum 255 characters in length.
	Name string `json:"name,required"`
	// An HTML/liquid template for the message type preview.
	Preview string `json:"preview,required"`
	// The icon name of the message type.
	IconName param.Opt[string] `json:"icon_name,omitzero"`
	// The semantic version of the message type.
	Semver param.Opt[string] `json:"semver,omitzero"`
	// The variants of the message type.
	Variants []MessageTypeVariantParam `json:"variants,omitzero"`
	paramObj
}

func (r MessageTypeValidateParamsMessageType) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeValidateParamsMessageType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeValidateParamsMessageType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
