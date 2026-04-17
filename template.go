// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

	"github.com/knocklabs/knock-mgmt-go/internal/apijson"
	"github.com/knocklabs/knock-mgmt-go/internal/apiquery"
	"github.com/knocklabs/knock-mgmt-go/internal/paramutil"
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// TemplateService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTemplateService] method instead.
type TemplateService struct {
	Options []option.RequestOption
}

// NewTemplateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTemplateService(opts ...option.RequestOption) (r TemplateService) {
	r = TemplateService{}
	r.Options = opts
	return
}

// Renders a template preview, without requiring a template to be persisted within
// Knock. This is useful for previewing templates in isolation, without the need to
// use a workflow.
//
// For email templates, you can optionally specify a layout by key or provide
// inline layout content.
func (r *TemplateService) Preview(ctx context.Context, params TemplatePreviewParams, opts ...option.RequestOption) (res *TemplatePreviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/templates/preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// A chat template.
type ChatTemplate struct {
	// The markdown body of the chat template.
	MarkdownBody string `json:"markdown_body" api:"required"`
	// A JSON template for the chat notification message payload. Only present if not
	// using the markdown body.
	JsonBody string `json:"json_body" api:"nullable"`
	// The summary of the chat template. Used by some chat apps in their push
	// notifications.
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MarkdownBody respjson.Field
		JsonBody     respjson.Field
		Summary      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatTemplate) RawJSON() string { return r.JSON.raw }
func (r *ChatTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ChatTemplate to a ChatTemplateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChatTemplateParam.Overrides()
func (r ChatTemplate) ToParam() ChatTemplateParam {
	return param.Override[ChatTemplateParam](json.RawMessage(r.RawJSON()))
}

// A chat template.
//
// The property MarkdownBody is required.
type ChatTemplateParam struct {
	// The markdown body of the chat template.
	MarkdownBody string `json:"markdown_body" api:"required"`
	// A JSON template for the chat notification message payload. Only present if not
	// using the markdown body.
	JsonBody param.Opt[string] `json:"json_body,omitzero"`
	// The summary of the chat template. Used by some chat apps in their push
	// notifications.
	Summary param.Opt[string] `json:"summary,omitzero"`
	paramObj
}

func (r ChatTemplateParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatTemplateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatTemplateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An email message template.
type EmailTemplate struct {
	// The [settings](https://docs.knock.app/integrations/email/settings) for the email
	// template. Must be supplied with at least `layout_key`.
	Settings EmailTemplateSettings `json:"settings" api:"required"`
	// The subject of the email. Supports Liquid templating with variables like
	// `{{ recipient.name }}`, `{{ actor.name }}`, `{{ vars.app_name }}`,
	// `{{ data.custom_field }}`, and `{{ tenant.name }}`. See the
	// [template variables reference](https://docs.knock.app/designing-workflows/template-editor/variables)
	// for available variables.
	Subject string `json:"subject" api:"required"`
	// An HTML or MJML template for the email body. **Required** if `visual_blocks` is
	// not provided. Only one of `html_body` or `visual_blocks` should be set. When
	// `is_mjml` is true, this must contain MJML components. Supports Liquid templating
	// with variables like `{{ recipient.name }}`, `{{ actor.name }}`,
	// `{{ vars.app_name }}`, `{{ data.custom_field }}`, and `{{ tenant.name }}`. See
	// the
	// [template variables reference](https://docs.knock.app/designing-workflows/template-editor/variables)
	// for available variables.
	HTMLBody string `json:"html_body" api:"nullable"`
	// Whether this template uses MJML format. When true, the template content will be
	// compiled from MJML to HTML. Only valid when the selected layout is also MJML or
	// when no layout is selected.
	IsMjml bool `json:"is_mjml" api:"nullable"`
	// A text template for the email body. When omitted, the email template will be
	// autogenerated from the `html_body` or `visual_blocks`.
	TextBody string `json:"text_body" api:"nullable"`
	// The visual blocks that make up the email template.
	VisualBlocks []EmailTemplateVisualBlockUnion `json:"visual_blocks" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Settings     respjson.Field
		Subject      respjson.Field
		HTMLBody     respjson.Field
		IsMjml       respjson.Field
		TextBody     respjson.Field
		VisualBlocks respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplate) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmailTemplate to a EmailTemplateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmailTemplateParam.Overrides()
func (r EmailTemplate) ToParam() EmailTemplateParam {
	return param.Override[EmailTemplateParam](json.RawMessage(r.RawJSON()))
}

// The [settings](https://docs.knock.app/integrations/email/settings) for the email
// template. Must be supplied with at least `layout_key`.
type EmailTemplateSettings struct {
	// The object path in the workflow trigger's `data` payload to resolve
	// attachments.Defaults to `attachments`.
	AttachmentKey string `json:"attachment_key" api:"nullable"`
	// The `key` of the
	// [email layout](https://docs.knock.app/integrations/email/layouts) that wraps the
	// email template. When omitted, the email template will need to define the
	// `<html>` structure.
	LayoutKey string `json:"layout_key" api:"nullable"`
	// A liquid template that will be injected into the email layout above the message
	// template content. Useful for setting variables that should be available to the
	// email layout.
	PreContent string `json:"pre_content" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachmentKey respjson.Field
		LayoutKey     respjson.Field
		PreContent    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateSettings) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EmailTemplateVisualBlockUnion contains all possible properties and values from
// [EmailTemplateVisualBlockEmailButtonSetBlock],
// [EmailTemplateVisualBlockEmailDividerBlock],
// [EmailTemplateVisualBlockEmailHTMLBlock],
// [EmailTemplateVisualBlockEmailImageBlock],
// [EmailTemplateVisualBlockEmailMarkdownBlock],
// [EmailTemplateVisualBlockEmailPartialBlock].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EmailTemplateVisualBlockUnion struct {
	// This field is from variant [EmailTemplateVisualBlockEmailButtonSetBlock].
	Buttons []EmailTemplateVisualBlockEmailButtonSetBlockButton `json:"buttons"`
	Type    string                                              `json:"type"`
	ID      string                                              `json:"id"`
	// This field is a union of
	// [EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrs],
	// [EmailTemplateVisualBlockEmailDividerBlockLayoutAttrs],
	// [EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrs],
	// [EmailTemplateVisualBlockEmailImageBlockLayoutAttrs],
	// [EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrs],
	// [EmailTemplateVisualBlockEmailPartialBlockLayoutAttrs]
	LayoutAttrs EmailTemplateVisualBlockUnionLayoutAttrs `json:"layout_attrs"`
	Version     int64                                    `json:"version"`
	Content     string                                   `json:"content"`
	// This field is from variant [EmailTemplateVisualBlockEmailImageBlock].
	URL string `json:"url"`
	// This field is from variant [EmailTemplateVisualBlockEmailImageBlock].
	Action string `json:"action"`
	// This field is from variant [EmailTemplateVisualBlockEmailImageBlock].
	Alt string `json:"alt"`
	// This field is from variant [EmailTemplateVisualBlockEmailImageBlock].
	StyleAttrs EmailTemplateVisualBlockEmailImageBlockStyleAttrs `json:"style_attrs"`
	// This field is from variant [EmailTemplateVisualBlockEmailMarkdownBlock].
	Variant string `json:"variant"`
	// This field is from variant [EmailTemplateVisualBlockEmailPartialBlock].
	Attrs map[string]any `json:"attrs"`
	// This field is from variant [EmailTemplateVisualBlockEmailPartialBlock].
	Key string `json:"key"`
	// This field is from variant [EmailTemplateVisualBlockEmailPartialBlock].
	Name string `json:"name"`
	JSON struct {
		Buttons     respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		LayoutAttrs respjson.Field
		Version     respjson.Field
		Content     respjson.Field
		URL         respjson.Field
		Action      respjson.Field
		Alt         respjson.Field
		StyleAttrs  respjson.Field
		Variant     respjson.Field
		Attrs       respjson.Field
		Key         respjson.Field
		Name        respjson.Field
		raw         string
	} `json:"-"`
}

func (u EmailTemplateVisualBlockUnion) AsEmailButtonSetBlock() (v EmailTemplateVisualBlockEmailButtonSetBlock) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EmailTemplateVisualBlockUnion) AsEmailDividerBlock() (v EmailTemplateVisualBlockEmailDividerBlock) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EmailTemplateVisualBlockUnion) AsEmailHTMLBlock() (v EmailTemplateVisualBlockEmailHTMLBlock) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EmailTemplateVisualBlockUnion) AsEmailImageBlock() (v EmailTemplateVisualBlockEmailImageBlock) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EmailTemplateVisualBlockUnion) AsEmailMarkdownBlock() (v EmailTemplateVisualBlockEmailMarkdownBlock) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EmailTemplateVisualBlockUnion) AsEmailPartialBlock() (v EmailTemplateVisualBlockEmailPartialBlock) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EmailTemplateVisualBlockUnion) RawJSON() string { return u.JSON.raw }

func (r *EmailTemplateVisualBlockUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EmailTemplateVisualBlockUnionLayoutAttrs is an implicit subunion of
// [EmailTemplateVisualBlockUnion]. EmailTemplateVisualBlockUnionLayoutAttrs
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [EmailTemplateVisualBlockUnion].
type EmailTemplateVisualBlockUnionLayoutAttrs struct {
	// This field is from variant
	// [EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrs].
	ColumnGap       int64  `json:"column_gap"`
	HorizontalAlign string `json:"horizontal_align"`
	PaddingBottom   int64  `json:"padding_bottom"`
	PaddingLeft     int64  `json:"padding_left"`
	PaddingRight    int64  `json:"padding_right"`
	PaddingTop      int64  `json:"padding_top"`
	JSON            struct {
		ColumnGap       respjson.Field
		HorizontalAlign respjson.Field
		PaddingBottom   respjson.Field
		PaddingLeft     respjson.Field
		PaddingRight    respjson.Field
		PaddingTop      respjson.Field
		raw             string
	} `json:"-"`
}

func (r *EmailTemplateVisualBlockUnionLayoutAttrs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A button set block in an email template.
type EmailTemplateVisualBlockEmailButtonSetBlock struct {
	// A list of buttons in the button set.
	Buttons []EmailTemplateVisualBlockEmailButtonSetBlockButton `json:"buttons" api:"required"`
	// The type of the block.
	//
	// Any of "button_set".
	Type string `json:"type" api:"required"`
	// The ID of the block.
	ID string `json:"id" format:"uuid"`
	// The layout attributes of the block.
	LayoutAttrs EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrs `json:"layout_attrs"`
	// The version of the block schema. This is automatically managed by Knock and
	// should not be set manually. Currently all blocks are at version 1.
	Version int64 `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buttons     respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		LayoutAttrs respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailButtonSetBlock) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailButtonSetBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A button in a button set block.
type EmailTemplateVisualBlockEmailButtonSetBlockButton struct {
	// The action of the button.
	Action string `json:"action" api:"required"`
	// The label of the button.
	Label string `json:"label" api:"required"`
	// The variant of the button.
	//
	// Any of "solid", "outline".
	Variant string `json:"variant" api:"required"`
	// The size attributes of the button.
	SizeAttrs EmailTemplateVisualBlockEmailButtonSetBlockButtonSizeAttrs `json:"size_attrs"`
	// The style attributes of the button.
	StyleAttrs EmailTemplateVisualBlockEmailButtonSetBlockButtonStyleAttrs `json:"style_attrs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Label       respjson.Field
		Variant     respjson.Field
		SizeAttrs   respjson.Field
		StyleAttrs  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailButtonSetBlockButton) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailButtonSetBlockButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The size attributes of the button.
type EmailTemplateVisualBlockEmailButtonSetBlockButtonSizeAttrs struct {
	// Whether the button is full width.
	IsFullwidth bool `json:"is_fullwidth"`
	// The size of the button.
	//
	// Any of "sm", "md", "lg".
	Size string `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsFullwidth respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailButtonSetBlockButtonSizeAttrs) RawJSON() string {
	return r.JSON.raw
}
func (r *EmailTemplateVisualBlockEmailButtonSetBlockButtonSizeAttrs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The style attributes of the button.
type EmailTemplateVisualBlockEmailButtonSetBlockButtonStyleAttrs struct {
	// The background color of the button.
	BackgroundColor string `json:"background_color"`
	// The border color of the button.
	BorderColor string `json:"border_color"`
	// The border radius of the button.
	BorderRadius int64 `json:"border_radius"`
	// The border width of the button.
	BorderWidth int64 `json:"border_width"`
	// The text color of the button.
	TextColor string `json:"text_color"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BackgroundColor respjson.Field
		BorderColor     respjson.Field
		BorderRadius    respjson.Field
		BorderWidth     respjson.Field
		TextColor       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailButtonSetBlockButtonStyleAttrs) RawJSON() string {
	return r.JSON.raw
}
func (r *EmailTemplateVisualBlockEmailButtonSetBlockButtonStyleAttrs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The layout attributes of the block.
type EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrs struct {
	// The column_gap layout attribute of the block.
	ColumnGap int64 `json:"column_gap" api:"required"`
	// The horizontal alignment of the block.
	//
	// Any of "left", "center", "right".
	HorizontalAlign string `json:"horizontal_align" api:"required"`
	// The padding_bottom layout attribute of the block.
	PaddingBottom int64 `json:"padding_bottom" api:"required"`
	// The padding_left layout attribute of the block.
	PaddingLeft int64 `json:"padding_left" api:"required"`
	// The padding_right layout attribute of the block.
	PaddingRight int64 `json:"padding_right" api:"required"`
	// The padding_top layout attribute of the block.
	PaddingTop int64 `json:"padding_top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ColumnGap       respjson.Field
		HorizontalAlign respjson.Field
		PaddingBottom   respjson.Field
		PaddingLeft     respjson.Field
		PaddingRight    respjson.Field
		PaddingTop      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrs) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A divider block in an email template.
type EmailTemplateVisualBlockEmailDividerBlock struct {
	// The type of the block.
	//
	// Any of "divider".
	Type string `json:"type" api:"required"`
	// The ID of the block.
	ID string `json:"id" format:"uuid"`
	// The layout attributes of the block.
	LayoutAttrs EmailTemplateVisualBlockEmailDividerBlockLayoutAttrs `json:"layout_attrs"`
	// The version of the block schema. This is automatically managed by Knock and
	// should not be set manually. Currently all blocks are at version 1.
	Version int64 `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ID          respjson.Field
		LayoutAttrs respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailDividerBlock) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailDividerBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The layout attributes of the block.
type EmailTemplateVisualBlockEmailDividerBlockLayoutAttrs struct {
	// The padding_bottom layout attribute of the block.
	PaddingBottom int64 `json:"padding_bottom" api:"required"`
	// The padding_left layout attribute of the block.
	PaddingLeft int64 `json:"padding_left" api:"required"`
	// The padding_right layout attribute of the block.
	PaddingRight int64 `json:"padding_right" api:"required"`
	// The padding_top layout attribute of the block.
	PaddingTop int64 `json:"padding_top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaddingBottom respjson.Field
		PaddingLeft   respjson.Field
		PaddingRight  respjson.Field
		PaddingTop    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailDividerBlockLayoutAttrs) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailDividerBlockLayoutAttrs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An HTML block in an email template.
type EmailTemplateVisualBlockEmailHTMLBlock struct {
	// The HTML content of the block. Supports Liquid templating with variables like
	// `{{ recipient.name }}`, `{{ actor.name }}`, `{{ vars.app_name }}`,
	// `{{ data.custom_field }}`, and `{{ tenant.name }}`. See the
	// [template variables reference](https://docs.knock.app/designing-workflows/template-editor/variables)
	// for available variables.
	Content string `json:"content" api:"required"`
	// The type of the block.
	//
	// Any of "html".
	Type string `json:"type" api:"required"`
	// The ID of the block.
	ID string `json:"id" format:"uuid"`
	// The layout attributes of the block.
	LayoutAttrs EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrs `json:"layout_attrs"`
	// The version of the block schema. This is automatically managed by Knock and
	// should not be set manually. Currently all blocks are at version 1.
	Version int64 `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		LayoutAttrs respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailHTMLBlock) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailHTMLBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The layout attributes of the block.
type EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrs struct {
	// The padding_bottom layout attribute of the block.
	PaddingBottom int64 `json:"padding_bottom" api:"required"`
	// The padding_left layout attribute of the block.
	PaddingLeft int64 `json:"padding_left" api:"required"`
	// The padding_right layout attribute of the block.
	PaddingRight int64 `json:"padding_right" api:"required"`
	// The padding_top layout attribute of the block.
	PaddingTop int64 `json:"padding_top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaddingBottom respjson.Field
		PaddingLeft   respjson.Field
		PaddingRight  respjson.Field
		PaddingTop    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrs) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image block in an email template.
type EmailTemplateVisualBlockEmailImageBlock struct {
	// The type of the block.
	//
	// Any of "image".
	Type string `json:"type" api:"required"`
	// The URL of the image to display.
	URL string `json:"url" api:"required" format:"uri"`
	// The ID of the block.
	ID string `json:"id" format:"uuid"`
	// Optional action URL for the image.
	Action string `json:"action" api:"nullable"`
	// Alt text for the image.
	Alt string `json:"alt" api:"nullable"`
	// The layout attributes of the block.
	LayoutAttrs EmailTemplateVisualBlockEmailImageBlockLayoutAttrs `json:"layout_attrs"`
	// The style attributes of the image.
	StyleAttrs EmailTemplateVisualBlockEmailImageBlockStyleAttrs `json:"style_attrs"`
	// The version of the block schema. This is automatically managed by Knock and
	// should not be set manually. Currently all blocks are at version 1.
	Version int64 `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ID          respjson.Field
		Action      respjson.Field
		Alt         respjson.Field
		LayoutAttrs respjson.Field
		StyleAttrs  respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailImageBlock) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailImageBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The layout attributes of the block.
type EmailTemplateVisualBlockEmailImageBlockLayoutAttrs struct {
	// The horizontal alignment of the block.
	//
	// Any of "left", "center", "right".
	HorizontalAlign string `json:"horizontal_align" api:"required"`
	// The padding_bottom layout attribute of the block.
	PaddingBottom int64 `json:"padding_bottom" api:"required"`
	// The padding_left layout attribute of the block.
	PaddingLeft int64 `json:"padding_left" api:"required"`
	// The padding_right layout attribute of the block.
	PaddingRight int64 `json:"padding_right" api:"required"`
	// The padding_top layout attribute of the block.
	PaddingTop int64 `json:"padding_top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HorizontalAlign respjson.Field
		PaddingBottom   respjson.Field
		PaddingLeft     respjson.Field
		PaddingRight    respjson.Field
		PaddingTop      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailImageBlockLayoutAttrs) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailImageBlockLayoutAttrs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The style attributes of the image.
type EmailTemplateVisualBlockEmailImageBlockStyleAttrs struct {
	// The width of the image.
	Width string `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailImageBlockStyleAttrs) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailImageBlockStyleAttrs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A markdown block in an email template.
type EmailTemplateVisualBlockEmailMarkdownBlock struct {
	// The markdown content of the block. Supports Liquid templating with variables
	// like `{{ recipient.name }}`, `{{ actor.name }}`, `{{ vars.app_name }}`,
	// `{{ data.custom_field }}`, and `{{ tenant.name }}`. See the
	// [template variables reference](https://docs.knock.app/designing-workflows/template-editor/variables)
	// for available variables.
	Content string `json:"content" api:"required"`
	// The type of the block.
	//
	// Any of "markdown".
	Type string `json:"type" api:"required"`
	// The ID of the block.
	ID string `json:"id" format:"uuid"`
	// The layout attributes of the block.
	LayoutAttrs EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrs `json:"layout_attrs"`
	// The flavor of markdown to use for the block.
	//
	// Any of "default".
	Variant string `json:"variant"`
	// The version of the block schema. This is automatically managed by Knock and
	// should not be set manually. Currently all blocks are at version 1.
	Version int64 `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		LayoutAttrs respjson.Field
		Variant     respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailMarkdownBlock) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailMarkdownBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The layout attributes of the block.
type EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrs struct {
	// The padding_bottom layout attribute of the block.
	PaddingBottom int64 `json:"padding_bottom" api:"required"`
	// The padding_left layout attribute of the block.
	PaddingLeft int64 `json:"padding_left" api:"required"`
	// The padding_right layout attribute of the block.
	PaddingRight int64 `json:"padding_right" api:"required"`
	// The padding_top layout attribute of the block.
	PaddingTop int64 `json:"padding_top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaddingBottom respjson.Field
		PaddingLeft   respjson.Field
		PaddingRight  respjson.Field
		PaddingTop    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrs) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A partial block in an email template, used to render a reusable partial
// component.
type EmailTemplateVisualBlockEmailPartialBlock struct {
	// The attributes to pass to the partial block.
	Attrs map[string]any `json:"attrs" api:"required"`
	// The key of the partial block to invoke.
	Key string `json:"key" api:"required"`
	// The name of the partial block.
	Name string `json:"name" api:"required"`
	// The type of the block.
	//
	// Any of "partial".
	Type string `json:"type" api:"required"`
	// The ID of the block.
	ID string `json:"id" format:"uuid"`
	// The layout attributes of the block.
	LayoutAttrs EmailTemplateVisualBlockEmailPartialBlockLayoutAttrs `json:"layout_attrs"`
	// The version of the block schema. This is automatically managed by Knock and
	// should not be set manually. Currently all blocks are at version 1.
	Version int64 `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attrs       respjson.Field
		Key         respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		LayoutAttrs respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailPartialBlock) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailPartialBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The layout attributes of the block.
type EmailTemplateVisualBlockEmailPartialBlockLayoutAttrs struct {
	// The padding_bottom layout attribute of the block.
	PaddingBottom int64 `json:"padding_bottom" api:"required"`
	// The padding_left layout attribute of the block.
	PaddingLeft int64 `json:"padding_left" api:"required"`
	// The padding_right layout attribute of the block.
	PaddingRight int64 `json:"padding_right" api:"required"`
	// The padding_top layout attribute of the block.
	PaddingTop int64 `json:"padding_top" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaddingBottom respjson.Field
		PaddingLeft   respjson.Field
		PaddingRight  respjson.Field
		PaddingTop    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailTemplateVisualBlockEmailPartialBlockLayoutAttrs) RawJSON() string { return r.JSON.raw }
func (r *EmailTemplateVisualBlockEmailPartialBlockLayoutAttrs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An email message template.
//
// The properties Settings, Subject are required.
type EmailTemplateParam struct {
	// The [settings](https://docs.knock.app/integrations/email/settings) for the email
	// template. Must be supplied with at least `layout_key`.
	Settings EmailTemplateSettingsParam `json:"settings,omitzero" api:"required"`
	// The subject of the email. Supports Liquid templating with variables like
	// `{{ recipient.name }}`, `{{ actor.name }}`, `{{ vars.app_name }}`,
	// `{{ data.custom_field }}`, and `{{ tenant.name }}`. See the
	// [template variables reference](https://docs.knock.app/designing-workflows/template-editor/variables)
	// for available variables.
	Subject string `json:"subject" api:"required"`
	// An HTML or MJML template for the email body. **Required** if `visual_blocks` is
	// not provided. Only one of `html_body` or `visual_blocks` should be set. When
	// `is_mjml` is true, this must contain MJML components. Supports Liquid templating
	// with variables like `{{ recipient.name }}`, `{{ actor.name }}`,
	// `{{ vars.app_name }}`, `{{ data.custom_field }}`, and `{{ tenant.name }}`. See
	// the
	// [template variables reference](https://docs.knock.app/designing-workflows/template-editor/variables)
	// for available variables.
	HTMLBody param.Opt[string] `json:"html_body,omitzero"`
	// Whether this template uses MJML format. When true, the template content will be
	// compiled from MJML to HTML. Only valid when the selected layout is also MJML or
	// when no layout is selected.
	IsMjml param.Opt[bool] `json:"is_mjml,omitzero"`
	// A text template for the email body. When omitted, the email template will be
	// autogenerated from the `html_body` or `visual_blocks`.
	TextBody param.Opt[string] `json:"text_body,omitzero"`
	// The visual blocks that make up the email template.
	VisualBlocks []EmailTemplateVisualBlockUnionParam `json:"visual_blocks,omitzero"`
	paramObj
}

func (r EmailTemplateParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The [settings](https://docs.knock.app/integrations/email/settings) for the email
// template. Must be supplied with at least `layout_key`.
type EmailTemplateSettingsParam struct {
	// The object path in the workflow trigger's `data` payload to resolve
	// attachments.Defaults to `attachments`.
	AttachmentKey param.Opt[string] `json:"attachment_key,omitzero"`
	// The `key` of the
	// [email layout](https://docs.knock.app/integrations/email/layouts) that wraps the
	// email template. When omitted, the email template will need to define the
	// `<html>` structure.
	LayoutKey param.Opt[string] `json:"layout_key,omitzero"`
	// A liquid template that will be injected into the email layout above the message
	// template content. Useful for setting variables that should be available to the
	// email layout.
	PreContent param.Opt[string] `json:"pre_content,omitzero"`
	paramObj
}

func (r EmailTemplateSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EmailTemplateVisualBlockUnionParam struct {
	OfEmailButtonSetBlock *EmailTemplateVisualBlockEmailButtonSetBlockParam `json:",omitzero,inline"`
	OfEmailDividerBlock   *EmailTemplateVisualBlockEmailDividerBlockParam   `json:",omitzero,inline"`
	OfEmailHTMLBlock      *EmailTemplateVisualBlockEmailHTMLBlockParam      `json:",omitzero,inline"`
	OfEmailImageBlock     *EmailTemplateVisualBlockEmailImageBlockParam     `json:",omitzero,inline"`
	OfEmailMarkdownBlock  *EmailTemplateVisualBlockEmailMarkdownBlockParam  `json:",omitzero,inline"`
	OfEmailPartialBlock   *EmailTemplateVisualBlockEmailPartialBlockParam   `json:",omitzero,inline"`
	paramUnion
}

func (u EmailTemplateVisualBlockUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEmailButtonSetBlock,
		u.OfEmailDividerBlock,
		u.OfEmailHTMLBlock,
		u.OfEmailImageBlock,
		u.OfEmailMarkdownBlock,
		u.OfEmailPartialBlock)
}
func (u *EmailTemplateVisualBlockUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EmailTemplateVisualBlockUnionParam) asAny() any {
	if !param.IsOmitted(u.OfEmailButtonSetBlock) {
		return u.OfEmailButtonSetBlock
	} else if !param.IsOmitted(u.OfEmailDividerBlock) {
		return u.OfEmailDividerBlock
	} else if !param.IsOmitted(u.OfEmailHTMLBlock) {
		return u.OfEmailHTMLBlock
	} else if !param.IsOmitted(u.OfEmailImageBlock) {
		return u.OfEmailImageBlock
	} else if !param.IsOmitted(u.OfEmailMarkdownBlock) {
		return u.OfEmailMarkdownBlock
	} else if !param.IsOmitted(u.OfEmailPartialBlock) {
		return u.OfEmailPartialBlock
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EmailTemplateVisualBlockUnionParam) GetButtons() []EmailTemplateVisualBlockEmailButtonSetBlockButtonParam {
	if vt := u.OfEmailButtonSetBlock; vt != nil {
		return vt.Buttons
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EmailTemplateVisualBlockUnionParam) GetURL() *string {
	if vt := u.OfEmailImageBlock; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EmailTemplateVisualBlockUnionParam) GetAction() *string {
	if vt := u.OfEmailImageBlock; vt != nil && vt.Action.Valid() {
		return &vt.Action.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EmailTemplateVisualBlockUnionParam) GetAlt() *string {
	if vt := u.OfEmailImageBlock; vt != nil && vt.Alt.Valid() {
		return &vt.Alt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EmailTemplateVisualBlockUnionParam) GetStyleAttrs() *EmailTemplateVisualBlockEmailImageBlockStyleAttrsParam {
	if vt := u.OfEmailImageBlock; vt != nil {
		return &vt.StyleAttrs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EmailTemplateVisualBlockUnionParam) GetVariant() *string {
	if vt := u.OfEmailMarkdownBlock; vt != nil {
		return &vt.Variant
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EmailTemplateVisualBlockUnionParam) GetAttrs() map[string]any {
	if vt := u.OfEmailPartialBlock; vt != nil {
		return vt.Attrs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EmailTemplateVisualBlockUnionParam) GetKey() *string {
	if vt := u.OfEmailPartialBlock; vt != nil {
		return &vt.Key
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EmailTemplateVisualBlockUnionParam) GetName() *string {
	if vt := u.OfEmailPartialBlock; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EmailTemplateVisualBlockUnionParam) GetType() *string {
	if vt := u.OfEmailButtonSetBlock; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEmailDividerBlock; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEmailHTMLBlock; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEmailImageBlock; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEmailMarkdownBlock; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfEmailPartialBlock; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EmailTemplateVisualBlockUnionParam) GetID() *string {
	if vt := u.OfEmailButtonSetBlock; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfEmailDividerBlock; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfEmailHTMLBlock; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfEmailImageBlock; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfEmailMarkdownBlock; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfEmailPartialBlock; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EmailTemplateVisualBlockUnionParam) GetVersion() *int64 {
	if vt := u.OfEmailButtonSetBlock; vt != nil && vt.Version.Valid() {
		return &vt.Version.Value
	} else if vt := u.OfEmailDividerBlock; vt != nil && vt.Version.Valid() {
		return &vt.Version.Value
	} else if vt := u.OfEmailHTMLBlock; vt != nil && vt.Version.Valid() {
		return &vt.Version.Value
	} else if vt := u.OfEmailImageBlock; vt != nil && vt.Version.Valid() {
		return &vt.Version.Value
	} else if vt := u.OfEmailMarkdownBlock; vt != nil && vt.Version.Valid() {
		return &vt.Version.Value
	} else if vt := u.OfEmailPartialBlock; vt != nil && vt.Version.Valid() {
		return &vt.Version.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EmailTemplateVisualBlockUnionParam) GetContent() *string {
	if vt := u.OfEmailHTMLBlock; vt != nil {
		return (*string)(&vt.Content)
	} else if vt := u.OfEmailMarkdownBlock; vt != nil {
		return (*string)(&vt.Content)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u EmailTemplateVisualBlockUnionParam) GetLayoutAttrs() (res emailTemplateVisualBlockUnionParamLayoutAttrs) {
	if vt := u.OfEmailButtonSetBlock; vt != nil {
		res.any = &vt.LayoutAttrs
	} else if vt := u.OfEmailDividerBlock; vt != nil {
		res.any = &vt.LayoutAttrs
	} else if vt := u.OfEmailHTMLBlock; vt != nil {
		res.any = &vt.LayoutAttrs
	} else if vt := u.OfEmailImageBlock; vt != nil {
		res.any = &vt.LayoutAttrs
	} else if vt := u.OfEmailMarkdownBlock; vt != nil {
		res.any = &vt.LayoutAttrs
	} else if vt := u.OfEmailPartialBlock; vt != nil {
		res.any = &vt.LayoutAttrs
	}
	return
}

// Can have the runtime types
// [*EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam],
// [*EmailTemplateVisualBlockEmailDividerBlockLayoutAttrsParam],
// [*EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrsParam],
// [*EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam],
// [*EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrsParam],
// [*EmailTemplateVisualBlockEmailPartialBlockLayoutAttrsParam]
type emailTemplateVisualBlockUnionParamLayoutAttrs struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam:
//	case *knockmapi.EmailTemplateVisualBlockEmailDividerBlockLayoutAttrsParam:
//	case *knockmapi.EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrsParam:
//	case *knockmapi.EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam:
//	case *knockmapi.EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrsParam:
//	case *knockmapi.EmailTemplateVisualBlockEmailPartialBlockLayoutAttrsParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u emailTemplateVisualBlockUnionParamLayoutAttrs) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u emailTemplateVisualBlockUnionParamLayoutAttrs) GetColumnGap() *int64 {
	switch vt := u.any.(type) {
	case *EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam:
		return &vt.ColumnGap
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u emailTemplateVisualBlockUnionParamLayoutAttrs) GetHorizontalAlign() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam:
		return (*string)(&vt.HorizontalAlign)
	case *EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam:
		return (*string)(&vt.HorizontalAlign)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u emailTemplateVisualBlockUnionParamLayoutAttrs) GetPaddingBottom() *int64 {
	switch vt := u.any.(type) {
	case *EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingBottom)
	case *EmailTemplateVisualBlockEmailDividerBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingBottom)
	case *EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingBottom)
	case *EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingBottom)
	case *EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingBottom)
	case *EmailTemplateVisualBlockEmailPartialBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingBottom)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u emailTemplateVisualBlockUnionParamLayoutAttrs) GetPaddingLeft() *int64 {
	switch vt := u.any.(type) {
	case *EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingLeft)
	case *EmailTemplateVisualBlockEmailDividerBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingLeft)
	case *EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingLeft)
	case *EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingLeft)
	case *EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingLeft)
	case *EmailTemplateVisualBlockEmailPartialBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingLeft)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u emailTemplateVisualBlockUnionParamLayoutAttrs) GetPaddingRight() *int64 {
	switch vt := u.any.(type) {
	case *EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingRight)
	case *EmailTemplateVisualBlockEmailDividerBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingRight)
	case *EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingRight)
	case *EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingRight)
	case *EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingRight)
	case *EmailTemplateVisualBlockEmailPartialBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingRight)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u emailTemplateVisualBlockUnionParamLayoutAttrs) GetPaddingTop() *int64 {
	switch vt := u.any.(type) {
	case *EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingTop)
	case *EmailTemplateVisualBlockEmailDividerBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingTop)
	case *EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingTop)
	case *EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingTop)
	case *EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingTop)
	case *EmailTemplateVisualBlockEmailPartialBlockLayoutAttrsParam:
		return (*int64)(&vt.PaddingTop)
	}
	return nil
}

// A button set block in an email template.
//
// The properties Buttons, Type are required.
type EmailTemplateVisualBlockEmailButtonSetBlockParam struct {
	// A list of buttons in the button set.
	Buttons []EmailTemplateVisualBlockEmailButtonSetBlockButtonParam `json:"buttons,omitzero" api:"required"`
	// The type of the block.
	//
	// Any of "button_set".
	Type string `json:"type,omitzero" api:"required"`
	// The ID of the block.
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// The version of the block schema. This is automatically managed by Knock and
	// should not be set manually. Currently all blocks are at version 1.
	Version param.Opt[int64] `json:"version,omitzero"`
	// The layout attributes of the block.
	LayoutAttrs EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam `json:"layout_attrs,omitzero"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailButtonSetBlockParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailButtonSetBlockParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailButtonSetBlockParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EmailTemplateVisualBlockEmailButtonSetBlockParam](
		"type", "button_set",
	)
}

// A button in a button set block.
//
// The properties Action, Label, Variant are required.
type EmailTemplateVisualBlockEmailButtonSetBlockButtonParam struct {
	// The action of the button.
	Action string `json:"action" api:"required"`
	// The label of the button.
	Label string `json:"label" api:"required"`
	// The variant of the button.
	//
	// Any of "solid", "outline".
	Variant string `json:"variant,omitzero" api:"required"`
	// The size attributes of the button.
	SizeAttrs EmailTemplateVisualBlockEmailButtonSetBlockButtonSizeAttrsParam `json:"size_attrs,omitzero"`
	// The style attributes of the button.
	StyleAttrs EmailTemplateVisualBlockEmailButtonSetBlockButtonStyleAttrsParam `json:"style_attrs,omitzero"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailButtonSetBlockButtonParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailButtonSetBlockButtonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailButtonSetBlockButtonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EmailTemplateVisualBlockEmailButtonSetBlockButtonParam](
		"variant", "solid", "outline",
	)
}

// The size attributes of the button.
type EmailTemplateVisualBlockEmailButtonSetBlockButtonSizeAttrsParam struct {
	// Whether the button is full width.
	IsFullwidth param.Opt[bool] `json:"is_fullwidth,omitzero"`
	// The size of the button.
	//
	// Any of "sm", "md", "lg".
	Size string `json:"size,omitzero"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailButtonSetBlockButtonSizeAttrsParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailButtonSetBlockButtonSizeAttrsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailButtonSetBlockButtonSizeAttrsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EmailTemplateVisualBlockEmailButtonSetBlockButtonSizeAttrsParam](
		"size", "sm", "md", "lg",
	)
}

// The style attributes of the button.
type EmailTemplateVisualBlockEmailButtonSetBlockButtonStyleAttrsParam struct {
	// The background color of the button.
	BackgroundColor param.Opt[string] `json:"background_color,omitzero"`
	// The border color of the button.
	BorderColor param.Opt[string] `json:"border_color,omitzero"`
	// The border radius of the button.
	BorderRadius param.Opt[int64] `json:"border_radius,omitzero"`
	// The border width of the button.
	BorderWidth param.Opt[int64] `json:"border_width,omitzero"`
	// The text color of the button.
	TextColor param.Opt[string] `json:"text_color,omitzero"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailButtonSetBlockButtonStyleAttrsParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailButtonSetBlockButtonStyleAttrsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailButtonSetBlockButtonStyleAttrsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The layout attributes of the block.
//
// The properties ColumnGap, HorizontalAlign, PaddingBottom, PaddingLeft,
// PaddingRight, PaddingTop are required.
type EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam struct {
	// The column_gap layout attribute of the block.
	ColumnGap int64 `json:"column_gap" api:"required"`
	// The horizontal alignment of the block.
	//
	// Any of "left", "center", "right".
	HorizontalAlign string `json:"horizontal_align,omitzero" api:"required"`
	// The padding_bottom layout attribute of the block.
	PaddingBottom int64 `json:"padding_bottom" api:"required"`
	// The padding_left layout attribute of the block.
	PaddingLeft int64 `json:"padding_left" api:"required"`
	// The padding_right layout attribute of the block.
	PaddingRight int64 `json:"padding_right" api:"required"`
	// The padding_top layout attribute of the block.
	PaddingTop int64 `json:"padding_top" api:"required"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam](
		"horizontal_align", "left", "center", "right",
	)
}

// A divider block in an email template.
//
// The property Type is required.
type EmailTemplateVisualBlockEmailDividerBlockParam struct {
	// The type of the block.
	//
	// Any of "divider".
	Type string `json:"type,omitzero" api:"required"`
	// The ID of the block.
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// The version of the block schema. This is automatically managed by Knock and
	// should not be set manually. Currently all blocks are at version 1.
	Version param.Opt[int64] `json:"version,omitzero"`
	// The layout attributes of the block.
	LayoutAttrs EmailTemplateVisualBlockEmailDividerBlockLayoutAttrsParam `json:"layout_attrs,omitzero"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailDividerBlockParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailDividerBlockParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailDividerBlockParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EmailTemplateVisualBlockEmailDividerBlockParam](
		"type", "divider",
	)
}

// The layout attributes of the block.
//
// The properties PaddingBottom, PaddingLeft, PaddingRight, PaddingTop are
// required.
type EmailTemplateVisualBlockEmailDividerBlockLayoutAttrsParam struct {
	// The padding_bottom layout attribute of the block.
	PaddingBottom int64 `json:"padding_bottom" api:"required"`
	// The padding_left layout attribute of the block.
	PaddingLeft int64 `json:"padding_left" api:"required"`
	// The padding_right layout attribute of the block.
	PaddingRight int64 `json:"padding_right" api:"required"`
	// The padding_top layout attribute of the block.
	PaddingTop int64 `json:"padding_top" api:"required"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailDividerBlockLayoutAttrsParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailDividerBlockLayoutAttrsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailDividerBlockLayoutAttrsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An HTML block in an email template.
//
// The properties Content, Type are required.
type EmailTemplateVisualBlockEmailHTMLBlockParam struct {
	// The HTML content of the block. Supports Liquid templating with variables like
	// `{{ recipient.name }}`, `{{ actor.name }}`, `{{ vars.app_name }}`,
	// `{{ data.custom_field }}`, and `{{ tenant.name }}`. See the
	// [template variables reference](https://docs.knock.app/designing-workflows/template-editor/variables)
	// for available variables.
	Content string `json:"content" api:"required"`
	// The type of the block.
	//
	// Any of "html".
	Type string `json:"type,omitzero" api:"required"`
	// The ID of the block.
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// The version of the block schema. This is automatically managed by Knock and
	// should not be set manually. Currently all blocks are at version 1.
	Version param.Opt[int64] `json:"version,omitzero"`
	// The layout attributes of the block.
	LayoutAttrs EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrsParam `json:"layout_attrs,omitzero"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailHTMLBlockParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailHTMLBlockParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailHTMLBlockParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EmailTemplateVisualBlockEmailHTMLBlockParam](
		"type", "html",
	)
}

// The layout attributes of the block.
//
// The properties PaddingBottom, PaddingLeft, PaddingRight, PaddingTop are
// required.
type EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrsParam struct {
	// The padding_bottom layout attribute of the block.
	PaddingBottom int64 `json:"padding_bottom" api:"required"`
	// The padding_left layout attribute of the block.
	PaddingLeft int64 `json:"padding_left" api:"required"`
	// The padding_right layout attribute of the block.
	PaddingRight int64 `json:"padding_right" api:"required"`
	// The padding_top layout attribute of the block.
	PaddingTop int64 `json:"padding_top" api:"required"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrsParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image block in an email template.
//
// The properties Type, URL are required.
type EmailTemplateVisualBlockEmailImageBlockParam struct {
	// The type of the block.
	//
	// Any of "image".
	Type string `json:"type,omitzero" api:"required"`
	// The URL of the image to display.
	URL string `json:"url" api:"required" format:"uri"`
	// Optional action URL for the image.
	Action param.Opt[string] `json:"action,omitzero"`
	// Alt text for the image.
	Alt param.Opt[string] `json:"alt,omitzero"`
	// The ID of the block.
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// The version of the block schema. This is automatically managed by Knock and
	// should not be set manually. Currently all blocks are at version 1.
	Version param.Opt[int64] `json:"version,omitzero"`
	// The layout attributes of the block.
	LayoutAttrs EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam `json:"layout_attrs,omitzero"`
	// The style attributes of the image.
	StyleAttrs EmailTemplateVisualBlockEmailImageBlockStyleAttrsParam `json:"style_attrs,omitzero"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailImageBlockParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailImageBlockParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailImageBlockParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EmailTemplateVisualBlockEmailImageBlockParam](
		"type", "image",
	)
}

// The layout attributes of the block.
//
// The properties HorizontalAlign, PaddingBottom, PaddingLeft, PaddingRight,
// PaddingTop are required.
type EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam struct {
	// The horizontal alignment of the block.
	//
	// Any of "left", "center", "right".
	HorizontalAlign string `json:"horizontal_align,omitzero" api:"required"`
	// The padding_bottom layout attribute of the block.
	PaddingBottom int64 `json:"padding_bottom" api:"required"`
	// The padding_left layout attribute of the block.
	PaddingLeft int64 `json:"padding_left" api:"required"`
	// The padding_right layout attribute of the block.
	PaddingRight int64 `json:"padding_right" api:"required"`
	// The padding_top layout attribute of the block.
	PaddingTop int64 `json:"padding_top" api:"required"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam](
		"horizontal_align", "left", "center", "right",
	)
}

// The style attributes of the image.
type EmailTemplateVisualBlockEmailImageBlockStyleAttrsParam struct {
	// The width of the image.
	Width param.Opt[string] `json:"width,omitzero"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailImageBlockStyleAttrsParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailImageBlockStyleAttrsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailImageBlockStyleAttrsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A markdown block in an email template.
//
// The properties Content, Type are required.
type EmailTemplateVisualBlockEmailMarkdownBlockParam struct {
	// The markdown content of the block. Supports Liquid templating with variables
	// like `{{ recipient.name }}`, `{{ actor.name }}`, `{{ vars.app_name }}`,
	// `{{ data.custom_field }}`, and `{{ tenant.name }}`. See the
	// [template variables reference](https://docs.knock.app/designing-workflows/template-editor/variables)
	// for available variables.
	Content string `json:"content" api:"required"`
	// The type of the block.
	//
	// Any of "markdown".
	Type string `json:"type,omitzero" api:"required"`
	// The ID of the block.
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// The version of the block schema. This is automatically managed by Knock and
	// should not be set manually. Currently all blocks are at version 1.
	Version param.Opt[int64] `json:"version,omitzero"`
	// The layout attributes of the block.
	LayoutAttrs EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrsParam `json:"layout_attrs,omitzero"`
	// The flavor of markdown to use for the block.
	//
	// Any of "default".
	Variant string `json:"variant,omitzero"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailMarkdownBlockParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailMarkdownBlockParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailMarkdownBlockParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EmailTemplateVisualBlockEmailMarkdownBlockParam](
		"type", "markdown",
	)
	apijson.RegisterFieldValidator[EmailTemplateVisualBlockEmailMarkdownBlockParam](
		"variant", "default",
	)
}

// The layout attributes of the block.
//
// The properties PaddingBottom, PaddingLeft, PaddingRight, PaddingTop are
// required.
type EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrsParam struct {
	// The padding_bottom layout attribute of the block.
	PaddingBottom int64 `json:"padding_bottom" api:"required"`
	// The padding_left layout attribute of the block.
	PaddingLeft int64 `json:"padding_left" api:"required"`
	// The padding_right layout attribute of the block.
	PaddingRight int64 `json:"padding_right" api:"required"`
	// The padding_top layout attribute of the block.
	PaddingTop int64 `json:"padding_top" api:"required"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrsParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A partial block in an email template, used to render a reusable partial
// component.
//
// The properties Attrs, Key, Name, Type are required.
type EmailTemplateVisualBlockEmailPartialBlockParam struct {
	// The attributes to pass to the partial block.
	Attrs map[string]any `json:"attrs,omitzero" api:"required"`
	// The key of the partial block to invoke.
	Key string `json:"key" api:"required"`
	// The name of the partial block.
	Name string `json:"name" api:"required"`
	// The type of the block.
	//
	// Any of "partial".
	Type string `json:"type,omitzero" api:"required"`
	// The ID of the block.
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// The version of the block schema. This is automatically managed by Knock and
	// should not be set manually. Currently all blocks are at version 1.
	Version param.Opt[int64] `json:"version,omitzero"`
	// The layout attributes of the block.
	LayoutAttrs EmailTemplateVisualBlockEmailPartialBlockLayoutAttrsParam `json:"layout_attrs,omitzero"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailPartialBlockParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailPartialBlockParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailPartialBlockParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EmailTemplateVisualBlockEmailPartialBlockParam](
		"type", "partial",
	)
}

// The layout attributes of the block.
//
// The properties PaddingBottom, PaddingLeft, PaddingRight, PaddingTop are
// required.
type EmailTemplateVisualBlockEmailPartialBlockLayoutAttrsParam struct {
	// The padding_bottom layout attribute of the block.
	PaddingBottom int64 `json:"padding_bottom" api:"required"`
	// The padding_left layout attribute of the block.
	PaddingLeft int64 `json:"padding_left" api:"required"`
	// The padding_right layout attribute of the block.
	PaddingRight int64 `json:"padding_right" api:"required"`
	// The padding_top layout attribute of the block.
	PaddingTop int64 `json:"padding_top" api:"required"`
	paramObj
}

func (r EmailTemplateVisualBlockEmailPartialBlockLayoutAttrsParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailTemplateVisualBlockEmailPartialBlockLayoutAttrsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailTemplateVisualBlockEmailPartialBlockLayoutAttrsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An in-app feed template.
type InAppFeedTemplate struct {
	// The markdown body of the in-app feed.
	MarkdownBody string `json:"markdown_body" api:"required"`
	// The action buttons of the in-app feed message.
	ActionButtons []InAppFeedTemplateActionButton `json:"action_buttons"`
	// The URL to navigate to when the in-app feed is tapped. Can be omitted for
	// multi-action templates, where the action buttons will be used instead.
	ActionURL string `json:"action_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MarkdownBody  respjson.Field
		ActionButtons respjson.Field
		ActionURL     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InAppFeedTemplate) RawJSON() string { return r.JSON.raw }
func (r *InAppFeedTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InAppFeedTemplate to a InAppFeedTemplateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InAppFeedTemplateParam.Overrides()
func (r InAppFeedTemplate) ToParam() InAppFeedTemplateParam {
	return param.Override[InAppFeedTemplateParam](json.RawMessage(r.RawJSON()))
}

// A single-action button to be rendered in an in-app feed cell.
type InAppFeedTemplateActionButton struct {
	// The URI for this action.
	Action string `json:"action" api:"required"`
	// The label of the action button.
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InAppFeedTemplateActionButton) RawJSON() string { return r.JSON.raw }
func (r *InAppFeedTemplateActionButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An in-app feed template.
//
// The property MarkdownBody is required.
type InAppFeedTemplateParam struct {
	// The markdown body of the in-app feed.
	MarkdownBody string `json:"markdown_body" api:"required"`
	// The URL to navigate to when the in-app feed is tapped. Can be omitted for
	// multi-action templates, where the action buttons will be used instead.
	ActionURL param.Opt[string] `json:"action_url,omitzero"`
	// The action buttons of the in-app feed message.
	ActionButtons []InAppFeedTemplateActionButtonParam `json:"action_buttons,omitzero"`
	paramObj
}

func (r InAppFeedTemplateParam) MarshalJSON() (data []byte, err error) {
	type shadow InAppFeedTemplateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InAppFeedTemplateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single-action button to be rendered in an in-app feed cell.
//
// The properties Action, Label are required.
type InAppFeedTemplateActionButtonParam struct {
	// The URI for this action.
	Action string `json:"action" api:"required"`
	// The label of the action button.
	Label string `json:"label" api:"required"`
	paramObj
}

func (r InAppFeedTemplateActionButtonParam) MarshalJSON() (data []byte, err error) {
	type shadow InAppFeedTemplateActionButtonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InAppFeedTemplateActionButtonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A push notification template.
type PushTemplate struct {
	// The [settings](https://docs.knock.app/integrations/sms/settings-and-overrides)
	// for the push template.
	Settings PushTemplateSettings `json:"settings" api:"required"`
	// The body of the push notification.
	TextBody string `json:"text_body" api:"required"`
	// The title of the push notification.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Settings    respjson.Field
		TextBody    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PushTemplate) RawJSON() string { return r.JSON.raw }
func (r *PushTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PushTemplate to a PushTemplateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PushTemplateParam.Overrides()
func (r PushTemplate) ToParam() PushTemplateParam {
	return param.Override[PushTemplateParam](json.RawMessage(r.RawJSON()))
}

// The [settings](https://docs.knock.app/integrations/sms/settings-and-overrides)
// for the push template.
type PushTemplateSettings struct {
	// The delivery type of the push notification. Set as silent to send a data-only
	// notification. When set to `silent`, no body will be sent.
	//
	// Any of "silent", "content".
	DeliveryType string `json:"delivery_type" api:"required"`
	// A JSON object that overrides the payload sent to the push provider.
	PayloadOverrides string `json:"payload_overrides"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeliveryType     respjson.Field
		PayloadOverrides respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PushTemplateSettings) RawJSON() string { return r.JSON.raw }
func (r *PushTemplateSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A push notification template.
//
// The properties Settings, TextBody, Title are required.
type PushTemplateParam struct {
	// The [settings](https://docs.knock.app/integrations/sms/settings-and-overrides)
	// for the push template.
	Settings PushTemplateSettingsParam `json:"settings,omitzero" api:"required"`
	// The body of the push notification.
	TextBody string `json:"text_body" api:"required"`
	// The title of the push notification.
	Title string `json:"title" api:"required"`
	paramObj
}

func (r PushTemplateParam) MarshalJSON() (data []byte, err error) {
	type shadow PushTemplateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PushTemplateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The [settings](https://docs.knock.app/integrations/sms/settings-and-overrides)
// for the push template.
//
// The property DeliveryType is required.
type PushTemplateSettingsParam struct {
	// The delivery type of the push notification. Set as silent to send a data-only
	// notification. When set to `silent`, no body will be sent.
	//
	// Any of "silent", "content".
	DeliveryType string `json:"delivery_type,omitzero" api:"required"`
	// A JSON object that overrides the payload sent to the push provider.
	PayloadOverrides param.Opt[string] `json:"payload_overrides,omitzero"`
	paramObj
}

func (r PushTemplateSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow PushTemplateSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PushTemplateSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PushTemplateSettingsParam](
		"delivery_type", "silent", "content",
	)
}

// A request template for a fetch function step.
type RequestTemplate struct {
	// The HTTP method of the request.
	//
	// Any of "get", "post", "put", "delete", "patch".
	Method RequestTemplateMethod `json:"method" api:"required"`
	// The URL of the request.
	URL string `json:"url" api:"required"`
	// The body of the request. Only used for POST or PUT requests.
	Body string `json:"body" api:"nullable"`
	// The headers of the request. Can be a template string or a list of key-value
	// pairs.
	Headers RequestTemplateHeadersUnion `json:"headers"`
	// The query params of the request. Can be a template string or a list of key-value
	// pairs.
	QueryParams RequestTemplateQueryParamsUnion `json:"query_params"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		URL         respjson.Field
		Body        respjson.Field
		Headers     respjson.Field
		QueryParams respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequestTemplate) RawJSON() string { return r.JSON.raw }
func (r *RequestTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RequestTemplate to a RequestTemplateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RequestTemplateParam.Overrides()
func (r RequestTemplate) ToParam() RequestTemplateParam {
	return param.Override[RequestTemplateParam](json.RawMessage(r.RawJSON()))
}

// The HTTP method of the request.
type RequestTemplateMethod string

const (
	RequestTemplateMethodGet    RequestTemplateMethod = "get"
	RequestTemplateMethodPost   RequestTemplateMethod = "post"
	RequestTemplateMethodPut    RequestTemplateMethod = "put"
	RequestTemplateMethodDelete RequestTemplateMethod = "delete"
	RequestTemplateMethodPatch  RequestTemplateMethod = "patch"
)

// RequestTemplateHeadersUnion contains all possible properties and values from
// [string], [[]RequestTemplateHeadersRequestTemplateHeadersArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfRequestTemplateHeadersArray]
type RequestTemplateHeadersUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]RequestTemplateHeadersRequestTemplateHeadersArrayItem] instead of an object.
	OfRequestTemplateHeadersArray []RequestTemplateHeadersRequestTemplateHeadersArrayItem `json:",inline"`
	JSON                          struct {
		OfString                      respjson.Field
		OfRequestTemplateHeadersArray respjson.Field
		raw                           string
	} `json:"-"`
}

func (u RequestTemplateHeadersUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RequestTemplateHeadersUnion) AsRequestTemplateHeadersArray() (v []RequestTemplateHeadersRequestTemplateHeadersArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RequestTemplateHeadersUnion) RawJSON() string { return u.JSON.raw }

func (r *RequestTemplateHeadersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequestTemplateHeadersRequestTemplateHeadersArrayItem struct {
	// The key of the header.
	Key string `json:"key" api:"required"`
	// The value of the header.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequestTemplateHeadersRequestTemplateHeadersArrayItem) RawJSON() string { return r.JSON.raw }
func (r *RequestTemplateHeadersRequestTemplateHeadersArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RequestTemplateQueryParamsUnion contains all possible properties and values from
// [string], [[]RequestTemplateQueryParamsRequestTemplateQueryParamsArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfRequestTemplateQuerysArray]
type RequestTemplateQueryParamsUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]RequestTemplateQueryParamsRequestTemplateQueryParamsArrayItem] instead of an
	// object.
	OfRequestTemplateQuerysArray []RequestTemplateQueryParamsRequestTemplateQueryParamsArrayItem `json:",inline"`
	JSON                         struct {
		OfString                     respjson.Field
		OfRequestTemplateQuerysArray respjson.Field
		raw                          string
	} `json:"-"`
}

func (u RequestTemplateQueryParamsUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u RequestTemplateQueryParamsUnion) AsRequestTemplateQueryParamsArray() (v []RequestTemplateQueryParamsRequestTemplateQueryParamsArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u RequestTemplateQueryParamsUnion) RawJSON() string { return u.JSON.raw }

func (r *RequestTemplateQueryParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequestTemplateQueryParamsRequestTemplateQueryParamsArrayItem struct {
	// The key of the query param.
	Key string `json:"key" api:"required"`
	// The value of the query param.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequestTemplateQueryParamsRequestTemplateQueryParamsArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *RequestTemplateQueryParamsRequestTemplateQueryParamsArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request template for a fetch function step.
//
// The properties Method, URL are required.
type RequestTemplateParam struct {
	// The HTTP method of the request.
	//
	// Any of "get", "post", "put", "delete", "patch".
	Method RequestTemplateMethod `json:"method,omitzero" api:"required"`
	// The URL of the request.
	URL string `json:"url" api:"required"`
	// The body of the request. Only used for POST or PUT requests.
	Body param.Opt[string] `json:"body,omitzero"`
	// The headers of the request. Can be a template string or a list of key-value
	// pairs.
	Headers RequestTemplateHeadersUnionParam `json:"headers,omitzero"`
	// The query params of the request. Can be a template string or a list of key-value
	// pairs.
	QueryParams RequestTemplateQueryParamsUnionParam `json:"query_params,omitzero"`
	paramObj
}

func (r RequestTemplateParam) MarshalJSON() (data []byte, err error) {
	type shadow RequestTemplateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RequestTemplateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RequestTemplateHeadersUnionParam struct {
	OfString                      param.Opt[string]                                            `json:",omitzero,inline"`
	OfRequestTemplateHeadersArray []RequestTemplateHeadersRequestTemplateHeadersArrayItemParam `json:",omitzero,inline"`
	paramUnion
}

func (u RequestTemplateHeadersUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfRequestTemplateHeadersArray)
}
func (u *RequestTemplateHeadersUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *RequestTemplateHeadersUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfRequestTemplateHeadersArray) {
		return &u.OfRequestTemplateHeadersArray
	}
	return nil
}

// The properties Key, Value are required.
type RequestTemplateHeadersRequestTemplateHeadersArrayItemParam struct {
	// The key of the header.
	Key string `json:"key" api:"required"`
	// The value of the header.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r RequestTemplateHeadersRequestTemplateHeadersArrayItemParam) MarshalJSON() (data []byte, err error) {
	type shadow RequestTemplateHeadersRequestTemplateHeadersArrayItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RequestTemplateHeadersRequestTemplateHeadersArrayItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RequestTemplateQueryParamsUnionParam struct {
	OfString                     param.Opt[string]                                                    `json:",omitzero,inline"`
	OfRequestTemplateQuerysArray []RequestTemplateQueryParamsRequestTemplateQueryParamsArrayItemParam `json:",omitzero,inline"`
	paramUnion
}

func (u RequestTemplateQueryParamsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfRequestTemplateQuerysArray)
}
func (u *RequestTemplateQueryParamsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *RequestTemplateQueryParamsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfRequestTemplateQuerysArray) {
		return &u.OfRequestTemplateQuerysArray
	}
	return nil
}

// The properties Key, Value are required.
type RequestTemplateQueryParamsRequestTemplateQueryParamsArrayItemParam struct {
	// The key of the query param.
	Key string `json:"key" api:"required"`
	// The value of the query param.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r RequestTemplateQueryParamsRequestTemplateQueryParamsArrayItemParam) MarshalJSON() (data []byte, err error) {
	type shadow RequestTemplateQueryParamsRequestTemplateQueryParamsArrayItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RequestTemplateQueryParamsRequestTemplateQueryParamsArrayItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An SMS template.
type SMSTemplate struct {
	// The message of the SMS.
	TextBody string `json:"text_body" api:"required"`
	// The [settings](https://docs.knock.app/integrations/sms/settings-and-overrides)
	// for the SMS template.
	Settings SMSTemplateSettings `json:"settings" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TextBody    respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SMSTemplate) RawJSON() string { return r.JSON.raw }
func (r *SMSTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SMSTemplate to a SMSTemplateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SMSTemplateParam.Overrides()
func (r SMSTemplate) ToParam() SMSTemplateParam {
	return param.Override[SMSTemplateParam](json.RawMessage(r.RawJSON()))
}

// The [settings](https://docs.knock.app/integrations/sms/settings-and-overrides)
// for the SMS template.
type SMSTemplateSettings struct {
	// A JSON object that overrides the payload sent to the SMS provider.
	PayloadOverrides string `json:"payload_overrides" api:"nullable"`
	// An override for the phone number to send the SMS to. When not set, defaults to
	// `recipient.phone_number`.
	ToNumber string `json:"to_number" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PayloadOverrides respjson.Field
		ToNumber         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SMSTemplateSettings) RawJSON() string { return r.JSON.raw }
func (r *SMSTemplateSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An SMS template.
//
// The property TextBody is required.
type SMSTemplateParam struct {
	// The message of the SMS.
	TextBody string `json:"text_body" api:"required"`
	// The [settings](https://docs.knock.app/integrations/sms/settings-and-overrides)
	// for the SMS template.
	Settings SMSTemplateSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r SMSTemplateParam) MarshalJSON() (data []byte, err error) {
	type shadow SMSTemplateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SMSTemplateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The [settings](https://docs.knock.app/integrations/sms/settings-and-overrides)
// for the SMS template.
type SMSTemplateSettingsParam struct {
	// A JSON object that overrides the payload sent to the SMS provider.
	PayloadOverrides param.Opt[string] `json:"payload_overrides,omitzero"`
	// An override for the phone number to send the SMS to. When not set, defaults to
	// `recipient.phone_number`.
	ToNumber param.Opt[string] `json:"to_number,omitzero"`
	paramObj
}

func (r SMSTemplateSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow SMSTemplateSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SMSTemplateSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A webhook template. By default, a webhook step will use the request settings you
// configured in your webhook channel. You can override this as you see fit on a
// per-step basis.
type WebhookTemplate struct {
	// The HTTP method of the webhook.
	//
	// Any of "get", "post", "put", "delete", "patch".
	Method WebhookTemplateMethod `json:"method" api:"required"`
	// The URL of the webhook.
	URL string `json:"url" api:"required"`
	// The body of the request. Only used for POST or PUT requests.
	Body string `json:"body" api:"nullable"`
	// A list of key-value pairs for the request headers. Each object should contain
	// key and value fields with string values.
	Headers []WebhookTemplateHeader `json:"headers"`
	// A list of key-value pairs for the request query params. Each object should
	// contain key and value fields with string values.
	QueryParams []WebhookTemplateQueryParam `json:"query_params"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		URL         respjson.Field
		Body        respjson.Field
		Headers     respjson.Field
		QueryParams respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTemplate) RawJSON() string { return r.JSON.raw }
func (r *WebhookTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WebhookTemplate to a WebhookTemplateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WebhookTemplateParam.Overrides()
func (r WebhookTemplate) ToParam() WebhookTemplateParam {
	return param.Override[WebhookTemplateParam](json.RawMessage(r.RawJSON()))
}

// The HTTP method of the webhook.
type WebhookTemplateMethod string

const (
	WebhookTemplateMethodGet    WebhookTemplateMethod = "get"
	WebhookTemplateMethodPost   WebhookTemplateMethod = "post"
	WebhookTemplateMethodPut    WebhookTemplateMethod = "put"
	WebhookTemplateMethodDelete WebhookTemplateMethod = "delete"
	WebhookTemplateMethodPatch  WebhookTemplateMethod = "patch"
)

type WebhookTemplateHeader struct {
	// The key of the header.
	Key string `json:"key" api:"required"`
	// The value of the header.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTemplateHeader) RawJSON() string { return r.JSON.raw }
func (r *WebhookTemplateHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookTemplateQueryParam struct {
	// The key of the query param.
	Key string `json:"key" api:"required"`
	// The value of the query param.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTemplateQueryParam) RawJSON() string { return r.JSON.raw }
func (r *WebhookTemplateQueryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A webhook template. By default, a webhook step will use the request settings you
// configured in your webhook channel. You can override this as you see fit on a
// per-step basis.
//
// The properties Method, URL are required.
type WebhookTemplateParam struct {
	// The HTTP method of the webhook.
	//
	// Any of "get", "post", "put", "delete", "patch".
	Method WebhookTemplateMethod `json:"method,omitzero" api:"required"`
	// The URL of the webhook.
	URL string `json:"url" api:"required"`
	// The body of the request. Only used for POST or PUT requests.
	Body param.Opt[string] `json:"body,omitzero"`
	// A list of key-value pairs for the request headers. Each object should contain
	// key and value fields with string values.
	Headers []WebhookTemplateHeaderParam `json:"headers,omitzero"`
	// A list of key-value pairs for the request query params. Each object should
	// contain key and value fields with string values.
	QueryParams []WebhookTemplateQueryParamParam `json:"query_params,omitzero"`
	paramObj
}

func (r WebhookTemplateParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookTemplateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookTemplateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Key, Value are required.
type WebhookTemplateHeaderParam struct {
	// The key of the header.
	Key string `json:"key" api:"required"`
	// The value of the header.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r WebhookTemplateHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookTemplateHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookTemplateHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Key, Value are required.
type WebhookTemplateQueryParamParam struct {
	// The key of the query param.
	Key string `json:"key" api:"required"`
	// The value of the query param.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r WebhookTemplateQueryParamParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookTemplateQueryParamParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookTemplateQueryParamParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to a template preview request.
type TemplatePreviewResponse struct {
	// The content type of the preview.
	//
	// Any of "email", "in_app_feed", "push", "chat", "sms".
	ContentType TemplatePreviewResponseContentType `json:"content_type" api:"required"`
	// The result of the preview.
	//
	// Any of "success", "error".
	Result TemplatePreviewResponseResult `json:"result" api:"required"`
	// A list of errors encountered during rendering. Present when result is "error".
	Errors []TemplatePreviewResponseError `json:"errors" api:"nullable"`
	// The rendered template, ready to be previewed.
	Template TemplatePreviewResponseTemplateUnion `json:"template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		Result      respjson.Field
		Errors      respjson.Field
		Template    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplatePreviewResponse) RawJSON() string { return r.JSON.raw }
func (r *TemplatePreviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The content type of the preview.
type TemplatePreviewResponseContentType string

const (
	TemplatePreviewResponseContentTypeEmail     TemplatePreviewResponseContentType = "email"
	TemplatePreviewResponseContentTypeInAppFeed TemplatePreviewResponseContentType = "in_app_feed"
	TemplatePreviewResponseContentTypePush      TemplatePreviewResponseContentType = "push"
	TemplatePreviewResponseContentTypeChat      TemplatePreviewResponseContentType = "chat"
	TemplatePreviewResponseContentTypeSMS       TemplatePreviewResponseContentType = "sms"
)

// The result of the preview.
type TemplatePreviewResponseResult string

const (
	TemplatePreviewResponseResultSuccess TemplatePreviewResponseResult = "success"
	TemplatePreviewResponseResultError   TemplatePreviewResponseResult = "error"
)

// A rendering error with optional location information.
type TemplatePreviewResponseError struct {
	// A human-readable description of the error.
	Message string `json:"message" api:"required"`
	// The template field that caused the error, if available.
	Field string `json:"field" api:"nullable"`
	// The line number where the error occurred, if available.
	Line int64 `json:"line" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Field       respjson.Field
		Line        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplatePreviewResponseError) RawJSON() string { return r.JSON.raw }
func (r *TemplatePreviewResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TemplatePreviewResponseTemplateUnion contains all possible properties and values
// from [EmailTemplate], [InAppFeedTemplate], [PushTemplate], [ChatTemplate],
// [SMSTemplate].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TemplatePreviewResponseTemplateUnion struct {
	// This field is a union of [EmailTemplateSettings], [PushTemplateSettings],
	// [SMSTemplateSettings]
	Settings TemplatePreviewResponseTemplateUnionSettings `json:"settings"`
	// This field is from variant [EmailTemplate].
	Subject string `json:"subject"`
	// This field is from variant [EmailTemplate].
	HTMLBody string `json:"html_body"`
	// This field is from variant [EmailTemplate].
	IsMjml   bool   `json:"is_mjml"`
	TextBody string `json:"text_body"`
	// This field is from variant [EmailTemplate].
	VisualBlocks []EmailTemplateVisualBlockUnion `json:"visual_blocks"`
	MarkdownBody string                          `json:"markdown_body"`
	// This field is from variant [InAppFeedTemplate].
	ActionButtons []InAppFeedTemplateActionButton `json:"action_buttons"`
	// This field is from variant [InAppFeedTemplate].
	ActionURL string `json:"action_url"`
	// This field is from variant [PushTemplate].
	Title string `json:"title"`
	// This field is from variant [ChatTemplate].
	JsonBody string `json:"json_body"`
	// This field is from variant [ChatTemplate].
	Summary string `json:"summary"`
	JSON    struct {
		Settings      respjson.Field
		Subject       respjson.Field
		HTMLBody      respjson.Field
		IsMjml        respjson.Field
		TextBody      respjson.Field
		VisualBlocks  respjson.Field
		MarkdownBody  respjson.Field
		ActionButtons respjson.Field
		ActionURL     respjson.Field
		Title         respjson.Field
		JsonBody      respjson.Field
		Summary       respjson.Field
		raw           string
	} `json:"-"`
}

func (u TemplatePreviewResponseTemplateUnion) AsEmailTemplate() (v EmailTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TemplatePreviewResponseTemplateUnion) AsInAppFeedTemplate() (v InAppFeedTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TemplatePreviewResponseTemplateUnion) AsPushTemplate() (v PushTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TemplatePreviewResponseTemplateUnion) AsChatTemplate() (v ChatTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TemplatePreviewResponseTemplateUnion) AsSMSTemplate() (v SMSTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TemplatePreviewResponseTemplateUnion) RawJSON() string { return u.JSON.raw }

func (r *TemplatePreviewResponseTemplateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TemplatePreviewResponseTemplateUnionSettings is an implicit subunion of
// [TemplatePreviewResponseTemplateUnion].
// TemplatePreviewResponseTemplateUnionSettings provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TemplatePreviewResponseTemplateUnion].
type TemplatePreviewResponseTemplateUnionSettings struct {
	// This field is from variant [EmailTemplateSettings].
	AttachmentKey string `json:"attachment_key"`
	// This field is from variant [EmailTemplateSettings].
	LayoutKey string `json:"layout_key"`
	// This field is from variant [EmailTemplateSettings].
	PreContent string `json:"pre_content"`
	// This field is from variant [PushTemplateSettings].
	DeliveryType     string `json:"delivery_type"`
	PayloadOverrides string `json:"payload_overrides"`
	// This field is from variant [SMSTemplateSettings].
	ToNumber string `json:"to_number"`
	JSON     struct {
		AttachmentKey    respjson.Field
		LayoutKey        respjson.Field
		PreContent       respjson.Field
		DeliveryType     respjson.Field
		PayloadOverrides respjson.Field
		ToNumber         respjson.Field
		raw              string
	} `json:"-"`
}

func (r *TemplatePreviewResponseTemplateUnionSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TemplatePreviewParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The channel type of the template to preview.
	//
	// Any of "email", "sms", "push", "chat", "in_app_feed".
	ChannelType TemplatePreviewParamsChannelType `json:"channel_type,omitzero" api:"required"`
	// A recipient reference, used when referencing a recipient by either their ID (for
	// a user), or by a reference for an object.
	Recipient TemplatePreviewParamsRecipientUnion `json:"recipient,omitzero" api:"required"`
	// The template content to preview. Structure depends on channel_type.
	Template TemplatePreviewParamsTemplateUnion `json:"template,omitzero" api:"required"`
	// The tenant to associate with the preview. Must not contain whitespace.
	Tenant param.Opt[string] `json:"tenant,omitzero"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// A recipient reference, used when referencing a recipient by either their ID (for
	// a user), or by a reference for an object.
	Actor TemplatePreviewParamsActorUnion `json:"actor,omitzero"`
	// Email layout configuration. Only applicable for email channel type. Falls back
	// to environment default if not provided.
	Layout TemplatePreviewParamsLayout `json:"layout,omitzero"`
	// Optional workflow context for variable hydration. When provided,
	// recipient/actor/tenant are resolved via Knock.
	Workflow TemplatePreviewParamsWorkflow `json:"workflow,omitzero"`
	// The data to pass to the template for rendering.
	Data map[string]any `json:"data,omitzero"`
	paramObj
}

func (r TemplatePreviewParams) MarshalJSON() (data []byte, err error) {
	type shadow TemplatePreviewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplatePreviewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TemplatePreviewParams]'s query parameters as `url.Values`.
func (r TemplatePreviewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The channel type of the template to preview.
type TemplatePreviewParamsChannelType string

const (
	TemplatePreviewParamsChannelTypeEmail     TemplatePreviewParamsChannelType = "email"
	TemplatePreviewParamsChannelTypeSMS       TemplatePreviewParamsChannelType = "sms"
	TemplatePreviewParamsChannelTypePush      TemplatePreviewParamsChannelType = "push"
	TemplatePreviewParamsChannelTypeChat      TemplatePreviewParamsChannelType = "chat"
	TemplatePreviewParamsChannelTypeInAppFeed TemplatePreviewParamsChannelType = "in_app_feed"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TemplatePreviewParamsRecipientUnion struct {
	OfString                   param.Opt[string]                                       `json:",omitzero,inline"`
	OfObjectRecipientReference *TemplatePreviewParamsRecipientObjectRecipientReference `json:",omitzero,inline"`
	paramUnion
}

func (u TemplatePreviewParamsRecipientUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfObjectRecipientReference)
}
func (u *TemplatePreviewParamsRecipientUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TemplatePreviewParamsRecipientUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfObjectRecipientReference) {
		return u.OfObjectRecipientReference
	}
	return nil
}

// An object reference.
//
// The properties ID, Collection are required.
type TemplatePreviewParamsRecipientObjectRecipientReference struct {
	// The ID of the object.
	ID string `json:"id" api:"required"`
	// The collection of the object.
	Collection string `json:"collection" api:"required"`
	paramObj
}

func (r TemplatePreviewParamsRecipientObjectRecipientReference) MarshalJSON() (data []byte, err error) {
	type shadow TemplatePreviewParamsRecipientObjectRecipientReference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplatePreviewParamsRecipientObjectRecipientReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TemplatePreviewParamsTemplateUnion struct {
	OfEmailTemplate     *EmailTemplateParam     `json:",omitzero,inline"`
	OfSMSTemplate       *SMSTemplateParam       `json:",omitzero,inline"`
	OfPushTemplate      *PushTemplateParam      `json:",omitzero,inline"`
	OfChatTemplate      *ChatTemplateParam      `json:",omitzero,inline"`
	OfInAppFeedTemplate *InAppFeedTemplateParam `json:",omitzero,inline"`
	paramUnion
}

func (u TemplatePreviewParamsTemplateUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEmailTemplate,
		u.OfSMSTemplate,
		u.OfPushTemplate,
		u.OfChatTemplate,
		u.OfInAppFeedTemplate)
}
func (u *TemplatePreviewParamsTemplateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TemplatePreviewParamsTemplateUnion) asAny() any {
	if !param.IsOmitted(u.OfEmailTemplate) {
		return u.OfEmailTemplate
	} else if !param.IsOmitted(u.OfSMSTemplate) {
		return u.OfSMSTemplate
	} else if !param.IsOmitted(u.OfPushTemplate) {
		return u.OfPushTemplate
	} else if !param.IsOmitted(u.OfChatTemplate) {
		return u.OfChatTemplate
	} else if !param.IsOmitted(u.OfInAppFeedTemplate) {
		return u.OfInAppFeedTemplate
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TemplatePreviewParamsTemplateUnion) GetSubject() *string {
	if vt := u.OfEmailTemplate; vt != nil {
		return &vt.Subject
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TemplatePreviewParamsTemplateUnion) GetHTMLBody() *string {
	if vt := u.OfEmailTemplate; vt != nil && vt.HTMLBody.Valid() {
		return &vt.HTMLBody.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TemplatePreviewParamsTemplateUnion) GetIsMjml() *bool {
	if vt := u.OfEmailTemplate; vt != nil && vt.IsMjml.Valid() {
		return &vt.IsMjml.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TemplatePreviewParamsTemplateUnion) GetVisualBlocks() []EmailTemplateVisualBlockUnionParam {
	if vt := u.OfEmailTemplate; vt != nil {
		return vt.VisualBlocks
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TemplatePreviewParamsTemplateUnion) GetTitle() *string {
	if vt := u.OfPushTemplate; vt != nil {
		return &vt.Title
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TemplatePreviewParamsTemplateUnion) GetJsonBody() *string {
	if vt := u.OfChatTemplate; vt != nil && vt.JsonBody.Valid() {
		return &vt.JsonBody.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TemplatePreviewParamsTemplateUnion) GetSummary() *string {
	if vt := u.OfChatTemplate; vt != nil && vt.Summary.Valid() {
		return &vt.Summary.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TemplatePreviewParamsTemplateUnion) GetActionButtons() []InAppFeedTemplateActionButtonParam {
	if vt := u.OfInAppFeedTemplate; vt != nil {
		return vt.ActionButtons
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TemplatePreviewParamsTemplateUnion) GetActionURL() *string {
	if vt := u.OfInAppFeedTemplate; vt != nil && vt.ActionURL.Valid() {
		return &vt.ActionURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TemplatePreviewParamsTemplateUnion) GetTextBody() *string {
	if vt := u.OfEmailTemplate; vt != nil && vt.TextBody.Valid() {
		return &vt.TextBody.Value
	} else if vt := u.OfSMSTemplate; vt != nil {
		return (*string)(&vt.TextBody)
	} else if vt := u.OfPushTemplate; vt != nil {
		return (*string)(&vt.TextBody)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TemplatePreviewParamsTemplateUnion) GetMarkdownBody() *string {
	if vt := u.OfChatTemplate; vt != nil {
		return (*string)(&vt.MarkdownBody)
	} else if vt := u.OfInAppFeedTemplate; vt != nil {
		return (*string)(&vt.MarkdownBody)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u TemplatePreviewParamsTemplateUnion) GetSettings() (res templatePreviewParamsTemplateUnionSettings) {
	if vt := u.OfEmailTemplate; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfSMSTemplate; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfPushTemplate; vt != nil {
		res.any = &vt.Settings
	}
	return
}

// Can have the runtime types [*EmailTemplateSettingsParam],
// [*SMSTemplateSettingsParam], [*PushTemplateSettingsParam]
type templatePreviewParamsTemplateUnionSettings struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.EmailTemplateSettingsParam:
//	case *knockmapi.SMSTemplateSettingsParam:
//	case *knockmapi.PushTemplateSettingsParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u templatePreviewParamsTemplateUnionSettings) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u templatePreviewParamsTemplateUnionSettings) GetAttachmentKey() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.AttachmentKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u templatePreviewParamsTemplateUnionSettings) GetLayoutKey() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.LayoutKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u templatePreviewParamsTemplateUnionSettings) GetPreContent() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.PreContent)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u templatePreviewParamsTemplateUnionSettings) GetToNumber() *string {
	switch vt := u.any.(type) {
	case *SMSTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.ToNumber)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u templatePreviewParamsTemplateUnionSettings) GetDeliveryType() *string {
	switch vt := u.any.(type) {
	case *PushTemplateSettingsParam:
		return &vt.DeliveryType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u templatePreviewParamsTemplateUnionSettings) GetPayloadOverrides() *string {
	switch vt := u.any.(type) {
	case *SMSTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.PayloadOverrides)
	case *PushTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.PayloadOverrides)
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TemplatePreviewParamsActorUnion struct {
	OfString                   param.Opt[string]                                   `json:",omitzero,inline"`
	OfObjectRecipientReference *TemplatePreviewParamsActorObjectRecipientReference `json:",omitzero,inline"`
	paramUnion
}

func (u TemplatePreviewParamsActorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfObjectRecipientReference)
}
func (u *TemplatePreviewParamsActorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TemplatePreviewParamsActorUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfObjectRecipientReference) {
		return u.OfObjectRecipientReference
	}
	return nil
}

// An object reference.
//
// The properties ID, Collection are required.
type TemplatePreviewParamsActorObjectRecipientReference struct {
	// The ID of the object.
	ID string `json:"id" api:"required"`
	// The collection of the object.
	Collection string `json:"collection" api:"required"`
	paramObj
}

func (r TemplatePreviewParamsActorObjectRecipientReference) MarshalJSON() (data []byte, err error) {
	type shadow TemplatePreviewParamsActorObjectRecipientReference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplatePreviewParamsActorObjectRecipientReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email layout configuration. Only applicable for email channel type. Falls back
// to environment default if not provided.
type TemplatePreviewParamsLayout struct {
	// Inline HTML content for the layout. Must include `{{ content }}` placeholder.
	HTMLContent param.Opt[string] `json:"html_content,omitzero"`
	// The key of an existing email layout to use.
	Key param.Opt[string] `json:"key,omitzero"`
	// Inline text content for the layout.
	TextContent param.Opt[string] `json:"text_content,omitzero"`
	paramObj
}

func (r TemplatePreviewParamsLayout) MarshalJSON() (data []byte, err error) {
	type shadow TemplatePreviewParamsLayout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplatePreviewParamsLayout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional workflow context for variable hydration. When provided,
// recipient/actor/tenant are resolved via Knock.
//
// The property Key is required.
type TemplatePreviewParamsWorkflow struct {
	// The workflow key.
	Key string `json:"key" api:"required"`
	// Workflow categories.
	Categories []string `json:"categories,omitzero"`
	paramObj
}

func (r TemplatePreviewParamsWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow TemplatePreviewParamsWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TemplatePreviewParamsWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
