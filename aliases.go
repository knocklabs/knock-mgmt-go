// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"github.com/knocklabs/knock-mgmt-go/internal/apierror"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// Attaches a goal to a workflow, guide, or broadcast for attribution tracking.
//
// This is an alias to an internal type.
type GoalAttachment = shared.GoalAttachment

// Attaches a goal to a workflow, guide, or broadcast for attribution tracking.
//
// This is an alias to an internal type.
type GoalAttachmentParam = shared.GoalAttachmentParam

// A boolean field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeBooleanField = shared.MessageTypeBooleanField

// The type of the field.
//
// This is an alias to an internal type.
type MessageTypeBooleanFieldType = shared.MessageTypeBooleanFieldType

// Equals "boolean"
const MessageTypeBooleanFieldTypeBoolean = shared.MessageTypeBooleanFieldTypeBoolean

// Settings for the boolean field.
//
// This is an alias to an internal type.
type MessageTypeBooleanFieldSettings = shared.MessageTypeBooleanFieldSettings

// A boolean field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeBooleanFieldParam = shared.MessageTypeBooleanFieldParam

// Settings for the boolean field.
//
// This is an alias to an internal type.
type MessageTypeBooleanFieldSettingsParam = shared.MessageTypeBooleanFieldSettingsParam

// A button field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeButtonField = shared.MessageTypeButtonField

// The type of the field.
//
// This is an alias to an internal type.
type MessageTypeButtonFieldType = shared.MessageTypeButtonFieldType

// Equals "button"
const MessageTypeButtonFieldTypeButton = shared.MessageTypeButtonFieldTypeButton

// Settings for the button field.
//
// This is an alias to an internal type.
type MessageTypeButtonFieldSettings = shared.MessageTypeButtonFieldSettings

// A button field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeButtonFieldParam = shared.MessageTypeButtonFieldParam

// Settings for the button field.
//
// This is an alias to an internal type.
type MessageTypeButtonFieldSettingsParam = shared.MessageTypeButtonFieldSettingsParam

// A hex color field (#RGB or #RRGGBB) used in a message type or partial input
// schema.
//
// This is an alias to an internal type.
type MessageTypeColorField = shared.MessageTypeColorField

// The type of the field.
//
// This is an alias to an internal type.
type MessageTypeColorFieldType = shared.MessageTypeColorFieldType

// Equals "color"
const MessageTypeColorFieldTypeColor = shared.MessageTypeColorFieldTypeColor

// Settings for the color field.
//
// This is an alias to an internal type.
type MessageTypeColorFieldSettings = shared.MessageTypeColorFieldSettings

// A hex color field (#RGB or #RRGGBB) used in a message type or partial input
// schema.
//
// This is an alias to an internal type.
type MessageTypeColorFieldParam = shared.MessageTypeColorFieldParam

// Settings for the color field.
//
// This is an alias to an internal type.
type MessageTypeColorFieldSettingsParam = shared.MessageTypeColorFieldSettingsParam

// An image field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeImageField = shared.MessageTypeImageField

// The type of the field.
//
// This is an alias to an internal type.
type MessageTypeImageFieldType = shared.MessageTypeImageFieldType

// Equals "image"
const MessageTypeImageFieldTypeImage = shared.MessageTypeImageFieldTypeImage

// Settings for the image field.
//
// This is an alias to an internal type.
type MessageTypeImageFieldSettings = shared.MessageTypeImageFieldSettings

// An image field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeImageFieldParam = shared.MessageTypeImageFieldParam

// Settings for the image field.
//
// This is an alias to an internal type.
type MessageTypeImageFieldSettingsParam = shared.MessageTypeImageFieldSettingsParam

// A JSON field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeJsonField = shared.MessageTypeJsonField

// The type of the field.
//
// This is an alias to an internal type.
type MessageTypeJsonFieldType = shared.MessageTypeJsonFieldType

// Equals "json"
const MessageTypeJsonFieldTypeJson = shared.MessageTypeJsonFieldTypeJson

// Settings for the json field.
//
// This is an alias to an internal type.
type MessageTypeJsonFieldSettings = shared.MessageTypeJsonFieldSettings

// A JSON field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeJsonFieldParam = shared.MessageTypeJsonFieldParam

// Settings for the json field.
//
// This is an alias to an internal type.
type MessageTypeJsonFieldSettingsParam = shared.MessageTypeJsonFieldSettingsParam

// A list field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeListField = shared.MessageTypeListField

// The type of the field.
//
// This is an alias to an internal type.
type MessageTypeListFieldType = shared.MessageTypeListFieldType

// Equals "list"
const MessageTypeListFieldTypeList = shared.MessageTypeListFieldTypeList

// Settings for the list field.
//
// This is an alias to an internal type.
type MessageTypeListFieldSettings = shared.MessageTypeListFieldSettings

// A list field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeListFieldParam = shared.MessageTypeListFieldParam

// Settings for the list field.
//
// This is an alias to an internal type.
type MessageTypeListFieldSettingsParam = shared.MessageTypeListFieldSettingsParam

// A markdown field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeMarkdownField = shared.MessageTypeMarkdownField

// The type of the field.
//
// This is an alias to an internal type.
type MessageTypeMarkdownFieldType = shared.MessageTypeMarkdownFieldType

// Equals "markdown"
const MessageTypeMarkdownFieldTypeMarkdown = shared.MessageTypeMarkdownFieldTypeMarkdown

// Settings for the markdown field.
//
// This is an alias to an internal type.
type MessageTypeMarkdownFieldSettings = shared.MessageTypeMarkdownFieldSettings

// A markdown field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeMarkdownFieldParam = shared.MessageTypeMarkdownFieldParam

// Settings for the markdown field.
//
// This is an alias to an internal type.
type MessageTypeMarkdownFieldSettingsParam = shared.MessageTypeMarkdownFieldSettingsParam

// A multi-select field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeMultiSelectField = shared.MessageTypeMultiSelectField

// Settings for the multi_select field.
//
// This is an alias to an internal type.
type MessageTypeMultiSelectFieldSettings = shared.MessageTypeMultiSelectFieldSettings

// This is an alias to an internal type.
type MessageTypeMultiSelectFieldSettingsOption = shared.MessageTypeMultiSelectFieldSettingsOption

// The type of the field.
//
// This is an alias to an internal type.
type MessageTypeMultiSelectFieldType = shared.MessageTypeMultiSelectFieldType

// Equals "multi_select"
const MessageTypeMultiSelectFieldTypeMultiSelect = shared.MessageTypeMultiSelectFieldTypeMultiSelect

// A multi-select field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeMultiSelectFieldParam = shared.MessageTypeMultiSelectFieldParam

// Settings for the multi_select field.
//
// This is an alias to an internal type.
type MessageTypeMultiSelectFieldSettingsParam = shared.MessageTypeMultiSelectFieldSettingsParam

// This is an alias to an internal type.
type MessageTypeMultiSelectFieldSettingsOptionParam = shared.MessageTypeMultiSelectFieldSettingsOptionParam

// A numeric field used in a message type or partial input schema, with optional
// min/max bounds and a unit label for display.
//
// This is an alias to an internal type.
type MessageTypeNumberField = shared.MessageTypeNumberField

// The type of the field.
//
// This is an alias to an internal type.
type MessageTypeNumberFieldType = shared.MessageTypeNumberFieldType

// Equals "number"
const MessageTypeNumberFieldTypeNumber = shared.MessageTypeNumberFieldTypeNumber

// Settings for the number field.
//
// This is an alias to an internal type.
type MessageTypeNumberFieldSettings = shared.MessageTypeNumberFieldSettings

// A numeric field used in a message type or partial input schema, with optional
// min/max bounds and a unit label for display.
//
// This is an alias to an internal type.
type MessageTypeNumberFieldParam = shared.MessageTypeNumberFieldParam

// Settings for the number field.
//
// This is an alias to an internal type.
type MessageTypeNumberFieldSettingsParam = shared.MessageTypeNumberFieldSettingsParam

// A select field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeSelectField = shared.MessageTypeSelectField

// Settings for the select field.
//
// This is an alias to an internal type.
type MessageTypeSelectFieldSettings = shared.MessageTypeSelectFieldSettings

// This is an alias to an internal type.
type MessageTypeSelectFieldSettingsOption = shared.MessageTypeSelectFieldSettingsOption

// The type of the field.
//
// This is an alias to an internal type.
type MessageTypeSelectFieldType = shared.MessageTypeSelectFieldType

// Equals "select"
const MessageTypeSelectFieldTypeSelect = shared.MessageTypeSelectFieldTypeSelect

// A select field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeSelectFieldParam = shared.MessageTypeSelectFieldParam

// Settings for the select field.
//
// This is an alias to an internal type.
type MessageTypeSelectFieldSettingsParam = shared.MessageTypeSelectFieldSettingsParam

// This is an alias to an internal type.
type MessageTypeSelectFieldSettingsOptionParam = shared.MessageTypeSelectFieldSettingsOptionParam

// A text field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeTextField = shared.MessageTypeTextField

// The type of the field.
//
// This is an alias to an internal type.
type MessageTypeTextFieldType = shared.MessageTypeTextFieldType

// Equals "text"
const MessageTypeTextFieldTypeText = shared.MessageTypeTextFieldTypeText

// Settings for the text field.
//
// This is an alias to an internal type.
type MessageTypeTextFieldSettings = shared.MessageTypeTextFieldSettings

// A text field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeTextFieldParam = shared.MessageTypeTextFieldParam

// Settings for the text field.
//
// This is an alias to an internal type.
type MessageTypeTextFieldSettingsParam = shared.MessageTypeTextFieldSettingsParam

// A textarea field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeTextareaField = shared.MessageTypeTextareaField

// The type of the field.
//
// This is an alias to an internal type.
type MessageTypeTextareaFieldType = shared.MessageTypeTextareaFieldType

// Equals "textarea"
const MessageTypeTextareaFieldTypeTextarea = shared.MessageTypeTextareaFieldTypeTextarea

// Settings for the textarea field.
//
// This is an alias to an internal type.
type MessageTypeTextareaFieldSettings = shared.MessageTypeTextareaFieldSettings

// A textarea field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeTextareaFieldParam = shared.MessageTypeTextareaFieldParam

// Settings for the textarea field.
//
// This is an alias to an internal type.
type MessageTypeTextareaFieldSettingsParam = shared.MessageTypeTextareaFieldSettingsParam

// A URL field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeURLField = shared.MessageTypeURLField

// The type of the field.
//
// This is an alias to an internal type.
type MessageTypeURLFieldType = shared.MessageTypeURLFieldType

// Equals "url"
const MessageTypeURLFieldTypeURL = shared.MessageTypeURLFieldTypeURL

// Settings for the url field.
//
// This is an alias to an internal type.
type MessageTypeURLFieldSettings = shared.MessageTypeURLFieldSettings

// A URL field used in a message type.
//
// This is an alias to an internal type.
type MessageTypeURLFieldParam = shared.MessageTypeURLFieldParam

// Settings for the url field.
//
// This is an alias to an internal type.
type MessageTypeURLFieldSettingsParam = shared.MessageTypeURLFieldSettingsParam

// The information about a paginated result.
//
// This is an alias to an internal type.
type PageInfo = shared.PageInfo

// A recipient reference, used when referencing a recipient by either their ID (for
// a user), or by a reference for an object.
//
// This is an alias to an internal type.
type RecipientReferenceUnionParam = shared.RecipientReferenceUnionParam

// An object reference.
//
// This is an alias to an internal type.
type RecipientReferenceObjectRecipientReferenceParam = shared.RecipientReferenceObjectRecipientReferenceParam
