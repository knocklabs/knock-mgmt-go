// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/knocklabs/knock-mgmt-go/internal/apijson"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// A boolean field used in a message type.
type MessageTypeBooleanField struct {
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The label of the field.
	Label string `json:"label" api:"required"`
	// The type of the field.
	//
	// Any of "boolean".
	Type MessageTypeBooleanFieldType `json:"type" api:"required"`
	// Settings for the boolean field.
	Settings MessageTypeBooleanFieldSettings `json:"settings"`
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
func (r MessageTypeBooleanField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeBooleanField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageTypeBooleanField to a MessageTypeBooleanFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageTypeBooleanFieldParam.Overrides()
func (r MessageTypeBooleanField) ToParam() MessageTypeBooleanFieldParam {
	return param.Override[MessageTypeBooleanFieldParam](json.RawMessage(r.RawJSON()))
}

// The type of the field.
type MessageTypeBooleanFieldType string

const (
	MessageTypeBooleanFieldTypeBoolean MessageTypeBooleanFieldType = "boolean"
)

// Settings for the boolean field.
type MessageTypeBooleanFieldSettings struct {
	// The default value of the boolean field.
	Default     bool   `json:"default"`
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
func (r MessageTypeBooleanFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeBooleanFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A boolean field used in a message type.
//
// The properties Key, Label, Type are required.
type MessageTypeBooleanFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The type of the field.
	//
	// Any of "boolean".
	Type MessageTypeBooleanFieldType `json:"type,omitzero" api:"required"`
	// Settings for the boolean field.
	Settings MessageTypeBooleanFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeBooleanFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeBooleanFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeBooleanFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the boolean field.
type MessageTypeBooleanFieldSettingsParam struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	// The default value of the boolean field.
	Default param.Opt[bool] `json:"default,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeBooleanFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeBooleanFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeBooleanFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A button field used in a message type.
type MessageTypeButtonField struct {
	// A text field used in a message type.
	Action MessageTypeTextField `json:"action" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The label of the field.
	Label string `json:"label" api:"required"`
	// A text field used in a message type.
	Text MessageTypeTextField `json:"text" api:"required"`
	// The type of the field.
	//
	// Any of "button".
	Type MessageTypeButtonFieldType `json:"type" api:"required"`
	// Settings for the button field.
	Settings MessageTypeButtonFieldSettings `json:"settings"`
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
func (r MessageTypeButtonField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeButtonField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageTypeButtonField to a MessageTypeButtonFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageTypeButtonFieldParam.Overrides()
func (r MessageTypeButtonField) ToParam() MessageTypeButtonFieldParam {
	return param.Override[MessageTypeButtonFieldParam](json.RawMessage(r.RawJSON()))
}

// The type of the field.
type MessageTypeButtonFieldType string

const (
	MessageTypeButtonFieldTypeButton MessageTypeButtonFieldType = "button"
)

// Settings for the button field.
type MessageTypeButtonFieldSettings struct {
	Description string `json:"description" api:"nullable"`
	Placeholder string `json:"placeholder" api:"nullable"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Placeholder respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageTypeButtonFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeButtonFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A button field used in a message type.
//
// The properties Action, Key, Label, Text, Type are required.
type MessageTypeButtonFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// A text field used in a message type.
	Action MessageTypeTextFieldParam `json:"action,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// A text field used in a message type.
	Text MessageTypeTextFieldParam `json:"text,omitzero" api:"required"`
	// The type of the field.
	//
	// Any of "button".
	Type MessageTypeButtonFieldType `json:"type,omitzero" api:"required"`
	// Settings for the button field.
	Settings MessageTypeButtonFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeButtonFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeButtonFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeButtonFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the button field.
type MessageTypeButtonFieldSettingsParam struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeButtonFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeButtonFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeButtonFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image field used in a message type.
type MessageTypeImageField struct {
	// A text field used in a message type.
	Action MessageTypeTextField `json:"action" api:"required"`
	// A text field used in a message type.
	Alt MessageTypeTextField `json:"alt" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The label of the field.
	Label string `json:"label" api:"required"`
	// The type of the field.
	//
	// Any of "image".
	Type MessageTypeImageFieldType `json:"type" api:"required"`
	// A URL field used in a message type.
	URL MessageTypeURLField `json:"url" api:"required"`
	// Settings for the image field.
	Settings MessageTypeImageFieldSettings `json:"settings"`
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
func (r MessageTypeImageField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeImageField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageTypeImageField to a MessageTypeImageFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageTypeImageFieldParam.Overrides()
func (r MessageTypeImageField) ToParam() MessageTypeImageFieldParam {
	return param.Override[MessageTypeImageFieldParam](json.RawMessage(r.RawJSON()))
}

// The type of the field.
type MessageTypeImageFieldType string

const (
	MessageTypeImageFieldTypeImage MessageTypeImageFieldType = "image"
)

// Settings for the image field.
type MessageTypeImageFieldSettings struct {
	Description string `json:"description" api:"nullable"`
	Placeholder string `json:"placeholder" api:"nullable"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Placeholder respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageTypeImageFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeImageFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image field used in a message type.
//
// The properties Action, Alt, Key, Label, Type, URL are required.
type MessageTypeImageFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// A text field used in a message type.
	Action MessageTypeTextFieldParam `json:"action,omitzero" api:"required"`
	// A text field used in a message type.
	Alt MessageTypeTextFieldParam `json:"alt,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The type of the field.
	//
	// Any of "image".
	Type MessageTypeImageFieldType `json:"type,omitzero" api:"required"`
	// A URL field used in a message type.
	URL MessageTypeURLFieldParam `json:"url,omitzero" api:"required"`
	// Settings for the image field.
	Settings MessageTypeImageFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeImageFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeImageFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeImageFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the image field.
type MessageTypeImageFieldSettingsParam struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeImageFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeImageFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeImageFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A JSON field used in a message type.
type MessageTypeJsonField struct {
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The label of the field.
	Label string `json:"label" api:"required"`
	// The type of the field.
	//
	// Any of "json".
	Type MessageTypeJsonFieldType `json:"type" api:"required"`
	// Settings for the json field.
	Settings MessageTypeJsonFieldSettings `json:"settings"`
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
func (r MessageTypeJsonField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeJsonField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageTypeJsonField to a MessageTypeJsonFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageTypeJsonFieldParam.Overrides()
func (r MessageTypeJsonField) ToParam() MessageTypeJsonFieldParam {
	return param.Override[MessageTypeJsonFieldParam](json.RawMessage(r.RawJSON()))
}

// The type of the field.
type MessageTypeJsonFieldType string

const (
	MessageTypeJsonFieldTypeJson MessageTypeJsonFieldType = "json"
)

// Settings for the json field.
type MessageTypeJsonFieldSettings struct {
	// The default value of the JSON field.
	Default     any    `json:"default" api:"nullable"`
	Description string `json:"description" api:"nullable"`
	Placeholder string `json:"placeholder" api:"nullable"`
	// Whether the field is required.
	Required bool `json:"required"`
	// A JSON schema used to validate the structure of the JSON provided. Must be a
	// valid JSON schema.
	Schema any `json:"schema" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Placeholder respjson.Field
		Required    respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageTypeJsonFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeJsonFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A JSON field used in a message type.
//
// The properties Key, Label, Type are required.
type MessageTypeJsonFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The type of the field.
	//
	// Any of "json".
	Type MessageTypeJsonFieldType `json:"type,omitzero" api:"required"`
	// Settings for the json field.
	Settings MessageTypeJsonFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeJsonFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeJsonFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeJsonFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the json field.
type MessageTypeJsonFieldSettingsParam struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	// The default value of the JSON field.
	Default any `json:"default,omitzero"`
	// A JSON schema used to validate the structure of the JSON provided. Must be a
	// valid JSON schema.
	Schema any `json:"schema,omitzero"`
	paramObj
}

func (r MessageTypeJsonFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeJsonFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeJsonFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A markdown field used in a message type.
type MessageTypeMarkdownField struct {
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The label of the field.
	Label string `json:"label" api:"required"`
	// The type of the field.
	//
	// Any of "markdown".
	Type MessageTypeMarkdownFieldType `json:"type" api:"required"`
	// Settings for the markdown field.
	Settings MessageTypeMarkdownFieldSettings `json:"settings"`
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
func (r MessageTypeMarkdownField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeMarkdownField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageTypeMarkdownField to a
// MessageTypeMarkdownFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageTypeMarkdownFieldParam.Overrides()
func (r MessageTypeMarkdownField) ToParam() MessageTypeMarkdownFieldParam {
	return param.Override[MessageTypeMarkdownFieldParam](json.RawMessage(r.RawJSON()))
}

// The type of the field.
type MessageTypeMarkdownFieldType string

const (
	MessageTypeMarkdownFieldTypeMarkdown MessageTypeMarkdownFieldType = "markdown"
)

// Settings for the markdown field.
type MessageTypeMarkdownFieldSettings struct {
	// The default value of the markdown field.
	Default     string `json:"default"`
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
func (r MessageTypeMarkdownFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeMarkdownFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A markdown field used in a message type.
//
// The properties Key, Label, Type are required.
type MessageTypeMarkdownFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The type of the field.
	//
	// Any of "markdown".
	Type MessageTypeMarkdownFieldType `json:"type,omitzero" api:"required"`
	// Settings for the markdown field.
	Settings MessageTypeMarkdownFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeMarkdownFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeMarkdownFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeMarkdownFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the markdown field.
type MessageTypeMarkdownFieldSettingsParam struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	// The default value of the markdown field.
	Default param.Opt[string] `json:"default,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeMarkdownFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeMarkdownFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeMarkdownFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A multi-select field used in a message type.
type MessageTypeMultiSelectField struct {
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The label of the field.
	Label string `json:"label" api:"required"`
	// Settings for the multi_select field.
	Settings MessageTypeMultiSelectFieldSettings `json:"settings" api:"required"`
	// The type of the field.
	//
	// Any of "multi_select".
	Type MessageTypeMultiSelectFieldType `json:"type" api:"required"`
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
func (r MessageTypeMultiSelectField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeMultiSelectField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageTypeMultiSelectField to a
// MessageTypeMultiSelectFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageTypeMultiSelectFieldParam.Overrides()
func (r MessageTypeMultiSelectField) ToParam() MessageTypeMultiSelectFieldParam {
	return param.Override[MessageTypeMultiSelectFieldParam](json.RawMessage(r.RawJSON()))
}

// Settings for the multi_select field.
type MessageTypeMultiSelectFieldSettings struct {
	// The default values for the multi-select field.
	Default     []string `json:"default" api:"nullable"`
	Description string   `json:"description" api:"nullable"`
	// The available options for the multi-select field.
	Options     []MessageTypeMultiSelectFieldSettingsOption `json:"options"`
	Placeholder string                                      `json:"placeholder" api:"nullable"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Options     respjson.Field
		Placeholder respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageTypeMultiSelectFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeMultiSelectFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageTypeMultiSelectFieldSettingsOption struct {
	// The value for the option.
	Value string `json:"value" api:"required"`
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
func (r MessageTypeMultiSelectFieldSettingsOption) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeMultiSelectFieldSettingsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the field.
type MessageTypeMultiSelectFieldType string

const (
	MessageTypeMultiSelectFieldTypeMultiSelect MessageTypeMultiSelectFieldType = "multi_select"
)

// A multi-select field used in a message type.
//
// The properties Key, Label, Settings, Type are required.
type MessageTypeMultiSelectFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// Settings for the multi_select field.
	Settings MessageTypeMultiSelectFieldSettingsParam `json:"settings,omitzero" api:"required"`
	// The type of the field.
	//
	// Any of "multi_select".
	Type MessageTypeMultiSelectFieldType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r MessageTypeMultiSelectFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeMultiSelectFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeMultiSelectFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the multi_select field.
type MessageTypeMultiSelectFieldSettingsParam struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	// The default values for the multi-select field.
	Default []string `json:"default,omitzero"`
	// The available options for the multi-select field.
	Options []MessageTypeMultiSelectFieldSettingsOptionParam `json:"options,omitzero"`
	paramObj
}

func (r MessageTypeMultiSelectFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeMultiSelectFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeMultiSelectFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Value is required.
type MessageTypeMultiSelectFieldSettingsOptionParam struct {
	// The value for the option.
	Value string `json:"value" api:"required"`
	// The display label for the option.
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r MessageTypeMultiSelectFieldSettingsOptionParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeMultiSelectFieldSettingsOptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeMultiSelectFieldSettingsOptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A select field used in a message type.
type MessageTypeSelectField struct {
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The label of the field.
	Label string `json:"label" api:"required"`
	// Settings for the select field.
	Settings MessageTypeSelectFieldSettings `json:"settings" api:"required"`
	// The type of the field.
	//
	// Any of "select".
	Type MessageTypeSelectFieldType `json:"type" api:"required"`
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
func (r MessageTypeSelectField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeSelectField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageTypeSelectField to a MessageTypeSelectFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageTypeSelectFieldParam.Overrides()
func (r MessageTypeSelectField) ToParam() MessageTypeSelectFieldParam {
	return param.Override[MessageTypeSelectFieldParam](json.RawMessage(r.RawJSON()))
}

// Settings for the select field.
type MessageTypeSelectFieldSettings struct {
	// The default value for the select field.
	Default     string `json:"default" api:"nullable"`
	Description string `json:"description" api:"nullable"`
	// The available options for the select field.
	Options     []MessageTypeSelectFieldSettingsOption `json:"options"`
	Placeholder string                                 `json:"placeholder" api:"nullable"`
	// Whether the field is required.
	Required bool `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Options     respjson.Field
		Placeholder respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageTypeSelectFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeSelectFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageTypeSelectFieldSettingsOption struct {
	// The value for the option.
	Value string `json:"value" api:"required"`
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
func (r MessageTypeSelectFieldSettingsOption) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeSelectFieldSettingsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the field.
type MessageTypeSelectFieldType string

const (
	MessageTypeSelectFieldTypeSelect MessageTypeSelectFieldType = "select"
)

// A select field used in a message type.
//
// The properties Key, Label, Settings, Type are required.
type MessageTypeSelectFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// Settings for the select field.
	Settings MessageTypeSelectFieldSettingsParam `json:"settings,omitzero" api:"required"`
	// The type of the field.
	//
	// Any of "select".
	Type MessageTypeSelectFieldType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r MessageTypeSelectFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeSelectFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeSelectFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the select field.
type MessageTypeSelectFieldSettingsParam struct {
	// The default value for the select field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	// The available options for the select field.
	Options []MessageTypeSelectFieldSettingsOptionParam `json:"options,omitzero"`
	paramObj
}

func (r MessageTypeSelectFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeSelectFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeSelectFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Value is required.
type MessageTypeSelectFieldSettingsOptionParam struct {
	// The value for the option.
	Value string `json:"value" api:"required"`
	// The display label for the option.
	Label param.Opt[string] `json:"label,omitzero"`
	paramObj
}

func (r MessageTypeSelectFieldSettingsOptionParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeSelectFieldSettingsOptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeSelectFieldSettingsOptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

// A textarea field used in a message type.
type MessageTypeTextareaField struct {
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The label of the field.
	Label string `json:"label" api:"required"`
	// The type of the field.
	//
	// Any of "textarea".
	Type MessageTypeTextareaFieldType `json:"type" api:"required"`
	// Settings for the textarea field.
	Settings MessageTypeTextareaFieldSettings `json:"settings"`
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
func (r MessageTypeTextareaField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeTextareaField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageTypeTextareaField to a
// MessageTypeTextareaFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageTypeTextareaFieldParam.Overrides()
func (r MessageTypeTextareaField) ToParam() MessageTypeTextareaFieldParam {
	return param.Override[MessageTypeTextareaFieldParam](json.RawMessage(r.RawJSON()))
}

// The type of the field.
type MessageTypeTextareaFieldType string

const (
	MessageTypeTextareaFieldTypeTextarea MessageTypeTextareaFieldType = "textarea"
)

// Settings for the textarea field.
type MessageTypeTextareaFieldSettings struct {
	// The default value of the textarea field.
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
func (r MessageTypeTextareaFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeTextareaFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A textarea field used in a message type.
//
// The properties Key, Label, Type are required.
type MessageTypeTextareaFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The type of the field.
	//
	// Any of "textarea".
	Type MessageTypeTextareaFieldType `json:"type,omitzero" api:"required"`
	// Settings for the textarea field.
	Settings MessageTypeTextareaFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeTextareaFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeTextareaFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeTextareaFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the textarea field.
type MessageTypeTextareaFieldSettingsParam struct {
	// The default value of the textarea field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	MaxLength   param.Opt[int64]  `json:"max_length,omitzero"`
	MinLength   param.Opt[int64]  `json:"min_length,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeTextareaFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeTextareaFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeTextareaFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL field used in a message type.
type MessageTypeURLField struct {
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The label of the field.
	Label string `json:"label" api:"required"`
	// The type of the field.
	//
	// Any of "url".
	Type MessageTypeURLFieldType `json:"type" api:"required"`
	// Settings for the url field.
	Settings MessageTypeURLFieldSettings `json:"settings"`
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
func (r MessageTypeURLField) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeURLField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessageTypeURLField to a MessageTypeURLFieldParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessageTypeURLFieldParam.Overrides()
func (r MessageTypeURLField) ToParam() MessageTypeURLFieldParam {
	return param.Override[MessageTypeURLFieldParam](json.RawMessage(r.RawJSON()))
}

// The type of the field.
type MessageTypeURLFieldType string

const (
	MessageTypeURLFieldTypeURL MessageTypeURLFieldType = "url"
)

// Settings for the url field.
type MessageTypeURLFieldSettings struct {
	// The default value of the URL field.
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
func (r MessageTypeURLFieldSettings) RawJSON() string { return r.JSON.raw }
func (r *MessageTypeURLFieldSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL field used in a message type.
//
// The properties Key, Label, Type are required.
type MessageTypeURLFieldParam struct {
	// The label of the field.
	Label param.Opt[string] `json:"label,omitzero" api:"required"`
	// The unique key of the field.
	Key string `json:"key" api:"required"`
	// The type of the field.
	//
	// Any of "url".
	Type MessageTypeURLFieldType `json:"type,omitzero" api:"required"`
	// Settings for the url field.
	Settings MessageTypeURLFieldSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r MessageTypeURLFieldParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeURLFieldParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeURLFieldParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for the url field.
type MessageTypeURLFieldSettingsParam struct {
	// The default value of the URL field.
	Default     param.Opt[string] `json:"default,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	Placeholder param.Opt[string] `json:"placeholder,omitzero"`
	// Whether the field is required.
	Required param.Opt[bool] `json:"required,omitzero"`
	paramObj
}

func (r MessageTypeURLFieldSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageTypeURLFieldSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageTypeURLFieldSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The information about a paginated result.
type PageInfo struct {
	// The number of entries to fetch per-page.
	PageSize int64 `json:"page_size" api:"required"`
	// The cursor to fetch entries after. Will only be present if there are more
	// entries to fetch.
	After string `json:"after" api:"nullable"`
	// The cursor to fetch entries before. Will only be present if there are more
	// entries to fetch before the current page.
	Before string `json:"before" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageSize    respjson.Field
		After       respjson.Field
		Before      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageInfo) RawJSON() string { return r.JSON.raw }
func (r *PageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
