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

// A message type allows you to specify an in-app schema that defines the fields
// available for your in-app notifications.
//
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
	opts = slices.Concat(r.Options, opts)
	if messageTypeKey == "" {
		err = errors.New("missing required message_type_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/message_types/%s", messageTypeKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of message types available in a given environment.
func (r *MessageTypeService) List(ctx context.Context, query MessageTypeListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[MessageType], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	if messageTypeKey == "" {
		err = errors.New("missing required message_type_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/message_types/%s", messageTypeKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Validates a message type payload without persisting it.
//
// Note: this endpoint only operates on message types in the `development`
// environment.
func (r *MessageTypeService) Validate(ctx context.Context, messageTypeKey string, params MessageTypeValidateParams, opts ...option.RequestOption) (res *MessageTypeValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageTypeKey == "" {
		err = errors.New("missing required message_type_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/message_types/%s/validate", messageTypeKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A message type is a schema for a message that maps to a UI component or element
// within your application.
type MessageType struct {
	// The timestamp of when the message type was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The environment of the message type.
	Environment string `json:"environment" api:"required"`
	// The unique key string for the message type object. Must be at minimum 3
	// characters and at maximum 255 characters in length. Must be in the format of
	// ^[a-z0-9_-]+$.
	Key string `json:"key" api:"required"`
	// A name for the message type. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// The owner of the message type.
	//
	// Any of "system", "user".
	Owner MessageTypeOwner `json:"owner" api:"required"`
	// An HTML/liquid template for the message type preview.
	Preview string `json:"preview" api:"required"`
	// The semantic version of the message type.
	Semver string `json:"semver" api:"required"`
	// The SHA hash of the message type.
	Sha string `json:"sha" api:"required"`
	// The timestamp of when the message type was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Whether the message type is valid.
	Valid bool `json:"valid" api:"required"`
	// The variants of the message type.
	Variants []MessageTypeVariant `json:"variants" api:"required"`
	// The timestamp of when the message type was deleted.
	ArchivedAt time.Time `json:"archived_at" api:"nullable" format:"date-time"`
	// The timestamp of when the message type was deleted.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// An arbitrary string attached to a message type object. Useful for adding notes
	// about the message type for internal purposes. Maximum of 280 characters allowed.
	Description string `json:"description" api:"nullable"`
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
	Key string `json:"key" api:"required"`
	// The label of the field.
	Label string `json:"label" api:"required"`
	// The type of the field.
	//
	// Any of "text".
	Type MessageTypeTextFieldType `json:"type" api:"required"`
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
	Default     string `json:"default" api:"nullable"`
	Description string `json:"description" api:"nullable"`
	MaxLength   int64  `json:"max_length"`
	MinLength   int64  `json:"min_length"`
	Placeholder string `json:"placeholder" api:"nullable"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		MaxLength   respjson.Field
		MinLength   respjson.Field
		Placeholder respjson.Field
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
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The type of the field.
	//
	// Any of "text".
	Type MessageTypeTextFieldType `json:"type,omitzero" api:"required"`
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
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
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
	Fields []MessageTypeVariantFieldUnion `json:"fields" api:"required"`
	// The unique key string for the variant. Must be at minimum 3 characters and at
	// maximum 255 characters in length. Must be in the format of ^[a-z0-9_-]+$.
	Key string `json:"key" api:"required"`
	// A name for the variant. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
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
// [shared.MessageTypeBooleanField], [shared.MessageTypeButtonField],
// [shared.MessageTypeImageField], [shared.MessageTypeJsonField],
// [shared.MessageTypeMarkdownField], [shared.MessageTypeMultiSelectField],
// [shared.MessageTypeSelectField], [MessageTypeTextField],
// [shared.MessageTypeTextareaField], [shared.MessageTypeURLField].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessageTypeVariantFieldUnion struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Type  string `json:"type"`
	// This field is a union of [shared.MessageTypeBooleanFieldSettings],
	// [shared.MessageTypeButtonFieldSettings], [shared.MessageTypeImageFieldSettings],
	// [shared.MessageTypeJsonFieldSettings],
	// [shared.MessageTypeMarkdownFieldSettings],
	// [shared.MessageTypeMultiSelectFieldSettings],
	// [shared.MessageTypeSelectFieldSettings], [MessageTypeTextFieldSettings],
	// [shared.MessageTypeTextareaFieldSettings], [shared.MessageTypeURLFieldSettings]
	Settings MessageTypeVariantFieldUnionSettings `json:"settings"`
	// This field is from variant [shared.MessageTypeButtonField].
	Action MessageTypeTextField `json:"action"`
	// This field is from variant [shared.MessageTypeButtonField].
	Text MessageTypeTextField `json:"text"`
	// This field is from variant [shared.MessageTypeImageField].
	Alt MessageTypeTextField `json:"alt"`
	// This field is from variant [shared.MessageTypeImageField].
	URL  shared.MessageTypeURLField `json:"url"`
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

func (u MessageTypeVariantFieldUnion) AsMessageTypeBooleanField() (v shared.MessageTypeBooleanField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeButtonField() (v shared.MessageTypeButtonField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeImageField() (v shared.MessageTypeImageField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeJsonField() (v shared.MessageTypeJsonField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeMarkdownField() (v shared.MessageTypeMarkdownField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeMultiSelectField() (v shared.MessageTypeMultiSelectField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeSelectField() (v shared.MessageTypeSelectField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeTextField() (v MessageTypeTextField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeTextareaField() (v shared.MessageTypeTextareaField) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageTypeVariantFieldUnion) AsMessageTypeURLField() (v shared.MessageTypeURLField) {
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
	// This field is a union of [bool], [any], [string], [[]string], [string],
	// [string], [string], [string]
	Default     MessageTypeVariantFieldUnionSettingsDefault `json:"default"`
	Description string                                      `json:"description"`
	Placeholder string                                      `json:"placeholder"`
	Required    bool                                        `json:"required"`
	// This field is from variant [shared.MessageTypeJsonFieldSettings].
	Schema any `json:"schema"`
	// This field is a union of [[]shared.MessageTypeMultiSelectFieldSettingsOption],
	// [[]shared.MessageTypeSelectFieldSettingsOption]
	Options   MessageTypeVariantFieldUnionSettingsOptions `json:"options"`
	MaxLength int64                                       `json:"max_length"`
	MinLength int64                                       `json:"min_length"`
	JSON      struct {
		Default     respjson.Field
		Description respjson.Field
		Placeholder respjson.Field
		Required    respjson.Field
		Schema      respjson.Field
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
// will be valid: OfBool OfMessageTypeJsonFieldSettingsDefault OfString
// OfStringArray]
type MessageTypeVariantFieldUnionSettingsDefault struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMessageTypeJsonFieldSettingsDefault any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfBool                                respjson.Field
		OfMessageTypeJsonFieldSettingsDefault respjson.Field
		OfString                              respjson.Field
		OfStringArray                         respjson.Field
		raw                                   string
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
// will be valid: OfMessageTypeMultiSelectFieldSettingsOptions
// OfMessageTypeSelectFieldSettingsOptions]
type MessageTypeVariantFieldUnionSettingsOptions struct {
	// This field will be present if the value is a
	// [[]shared.MessageTypeMultiSelectFieldSettingsOption] instead of an object.
	OfMessageTypeMultiSelectFieldSettingsOptions []shared.MessageTypeMultiSelectFieldSettingsOption `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.MessageTypeSelectFieldSettingsOption] instead of an object.
	OfMessageTypeSelectFieldSettingsOptions []shared.MessageTypeSelectFieldSettingsOption `json:",inline"`
	JSON                                    struct {
		OfMessageTypeMultiSelectFieldSettingsOptions respjson.Field
		OfMessageTypeSelectFieldSettingsOptions      respjson.Field
		raw                                          string
	} `json:"-"`
}

func (r *MessageTypeVariantFieldUnionSettingsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A variant of a message type.
//
// The properties Fields, Key, Name are required.
type MessageTypeVariantParam struct {
	// The field types available for the variant.
	Fields []MessageTypeVariantFieldUnionParam `json:"fields,omitzero" api:"required"`
	// The unique key string for the variant. Must be at minimum 3 characters and at
	// maximum 255 characters in length. Must be in the format of ^[a-z0-9_-]+$.
	Key string `json:"key" api:"required"`
	// A name for the variant. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
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
	OfMessageTypeBooleanField     *shared.MessageTypeBooleanFieldParam     `json:",omitzero,inline"`
	OfMessageTypeButtonField      *shared.MessageTypeButtonFieldParam      `json:",omitzero,inline"`
	OfMessageTypeImageField       *shared.MessageTypeImageFieldParam       `json:",omitzero,inline"`
	OfMessageTypeJsonField        *shared.MessageTypeJsonFieldParam        `json:",omitzero,inline"`
	OfMessageTypeMarkdownField    *shared.MessageTypeMarkdownFieldParam    `json:",omitzero,inline"`
	OfMessageTypeMultiSelectField *shared.MessageTypeMultiSelectFieldParam `json:",omitzero,inline"`
	OfMessageTypeSelectField      *shared.MessageTypeSelectFieldParam      `json:",omitzero,inline"`
	OfMessageTypeTextField        *MessageTypeTextFieldParam               `json:",omitzero,inline"`
	OfMessageTypeTextareaField    *shared.MessageTypeTextareaFieldParam    `json:",omitzero,inline"`
	OfMessageTypeURLField         *shared.MessageTypeURLFieldParam         `json:",omitzero,inline"`
	paramUnion
}

func (u MessageTypeVariantFieldUnionParam) MarshalJSON() ([]byte, error) {
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
func (u MessageTypeVariantFieldUnionParam) GetURL() *shared.MessageTypeURLFieldParam {
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
func (u MessageTypeVariantFieldUnionParam) GetLabel() *string {
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
func (u MessageTypeVariantFieldUnionParam) GetType() *string {
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
func (u MessageTypeVariantFieldUnionParam) GetSettings() (res messageTypeVariantFieldUnionParamSettings) {
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

// Can have the runtime types [*shared.MessageTypeBooleanFieldSettingsParam],
// [*shared.MessageTypeButtonFieldSettingsParam],
// [*shared.MessageTypeImageFieldSettingsParam],
// [*shared.MessageTypeJsonFieldSettingsParam],
// [*shared.MessageTypeMarkdownFieldSettingsParam],
// [*shared.MessageTypeMultiSelectFieldSettingsParam],
// [*shared.MessageTypeSelectFieldSettingsParam],
// [*MessageTypeTextFieldSettingsParam],
// [*shared.MessageTypeTextareaFieldSettingsParam],
// [*shared.MessageTypeURLFieldSettingsParam]
type messageTypeVariantFieldUnionParamSettings struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *shared.MessageTypeBooleanFieldSettingsParam:
//	case *shared.MessageTypeButtonFieldSettingsParam:
//	case *shared.MessageTypeImageFieldSettingsParam:
//	case *shared.MessageTypeJsonFieldSettingsParam:
//	case *shared.MessageTypeMarkdownFieldSettingsParam:
//	case *shared.MessageTypeMultiSelectFieldSettingsParam:
//	case *shared.MessageTypeSelectFieldSettingsParam:
//	case *knockmapi.MessageTypeTextFieldSettingsParam:
//	case *shared.MessageTypeTextareaFieldSettingsParam:
//	case *shared.MessageTypeURLFieldSettingsParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u messageTypeVariantFieldUnionParamSettings) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u messageTypeVariantFieldUnionParamSettings) GetSchema() *any {
	switch vt := u.any.(type) {
	case *shared.MessageTypeJsonFieldSettingsParam:
		return &vt.Schema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u messageTypeVariantFieldUnionParamSettings) GetDescription() *string {
	switch vt := u.any.(type) {
	case *shared.MessageTypeBooleanFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeButtonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeImageFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeJsonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeMarkdownFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	case *shared.MessageTypeURLFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Description)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u messageTypeVariantFieldUnionParamSettings) GetPlaceholder() *string {
	switch vt := u.any.(type) {
	case *shared.MessageTypeBooleanFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeButtonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeImageFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeJsonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeMarkdownFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	case *shared.MessageTypeURLFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Placeholder)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u messageTypeVariantFieldUnionParamSettings) GetRequired() *bool {
	switch vt := u.any.(type) {
	case *shared.MessageTypeBooleanFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeButtonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeImageFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeJsonFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeMarkdownFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeSelectFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	case *shared.MessageTypeURLFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.Required)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u messageTypeVariantFieldUnionParamSettings) GetMaxLength() *int64 {
	switch vt := u.any.(type) {
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MaxLength)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MaxLength)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u messageTypeVariantFieldUnionParamSettings) GetMinLength() *int64 {
	switch vt := u.any.(type) {
	case *MessageTypeTextFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MinLength)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		return paramutil.AddrIfPresent(vt.MinLength)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u messageTypeVariantFieldUnionParamSettings) GetDefault() (res messageTypeVariantFieldUnionParamSettingsDefault) {
	switch vt := u.any.(type) {
	case *shared.MessageTypeBooleanFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeJsonFieldSettingsParam:
		res.any = &vt.Default
	case *shared.MessageTypeMarkdownFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		res.any = &vt.Default
	case *shared.MessageTypeSelectFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *MessageTypeTextFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeTextareaFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	case *shared.MessageTypeURLFieldSettingsParam:
		res.any = paramutil.AddrIfPresent(vt.Default)
	}
	return res
}

// Can have the runtime types [*bool], [*any], [*string], [\*[]string]
type messageTypeVariantFieldUnionParamSettingsDefault struct{ any }

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
func (u messageTypeVariantFieldUnionParamSettingsDefault) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u messageTypeVariantFieldUnionParamSettings) GetOptions() (res messageTypeVariantFieldUnionParamSettingsOptions) {
	switch vt := u.any.(type) {
	case *shared.MessageTypeMultiSelectFieldSettingsParam:
		res.any = &vt.Options
	case *shared.MessageTypeSelectFieldSettingsParam:
		res.any = &vt.Options
	}
	return res
}

// Can have the runtime types
// [_[]shared.MessageTypeMultiSelectFieldSettingsOptionParam],
// [_[]shared.MessageTypeSelectFieldSettingsOptionParam]
type messageTypeVariantFieldUnionParamSettingsOptions struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.MessageTypeMultiSelectFieldSettingsOptionParam:
//	case *[]shared.MessageTypeSelectFieldSettingsOptionParam:
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

// Wraps the MessageType response under the `message_type` key.
type MessageTypeUpsertResponse struct {
	// A message type is a schema for a message that maps to a UI component or element
	// within your application.
	MessageType MessageType `json:"message_type" api:"required"`
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
	MessageType MessageType `json:"message_type" api:"required"`
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

// URLQuery serializes [MessageTypeGetParams]'s query parameters as `url.Values`.
func (r MessageTypeGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageTypeListParams struct {
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

// URLQuery serializes [MessageTypeListParams]'s query parameters as `url.Values`.
func (r MessageTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessageTypeUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A request to create a message type.
	MessageType MessageTypeUpsertParamsMessageType `json:"message_type,omitzero" api:"required"`
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
	Description param.Opt[string] `json:"description,omitzero" api:"required"`
	// A name for the message type. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// An HTML/liquid template for the message type preview.
	Preview string `json:"preview" api:"required"`
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
	Environment string `query:"environment" api:"required" json:"-"`
	// A request to create a message type.
	MessageType MessageTypeValidateParamsMessageType `json:"message_type,omitzero" api:"required"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
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
	Description param.Opt[string] `json:"description,omitzero" api:"required"`
	// A name for the message type. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// An HTML/liquid template for the message type preview.
	Preview string `json:"preview" api:"required"`
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
