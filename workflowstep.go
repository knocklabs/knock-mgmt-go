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

	"github.com/knocklabs/knock-mgmt-go/internal/apijson"
	"github.com/knocklabs/knock-mgmt-go/internal/apiquery"
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// WorkflowStepService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowStepService] method instead.
type WorkflowStepService struct {
	Options []option.RequestOption
}

// NewWorkflowStepService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowStepService(opts ...option.RequestOption) (r WorkflowStepService) {
	r = WorkflowStepService{}
	r.Options = opts
	return
}

// Generates a rendered template for a given channel step in a workflow.
func (r *WorkflowStepService) PreviewTemplate(ctx context.Context, stepRef string, params WorkflowStepPreviewTemplateParams, opts ...option.RequestOption) (res *WorkflowStepPreviewTemplateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.WorkflowKey == "" {
		err = errors.New("missing required workflow_key parameter")
		return
	}
	if stepRef == "" {
		err = errors.New("missing required step_ref parameter")
		return
	}
	path := fmt.Sprintf("v1/workflows/%s/steps/%s/preview_template", params.WorkflowKey, stepRef)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// A response to a preview workflow template request.
type WorkflowStepPreviewTemplateResponse struct {
	// The content type of the preview.
	//
	// Any of "email", "in_app_feed", "push", "chat", "sms", "http".
	ContentType WorkflowStepPreviewTemplateResponseContentType `json:"content_type,required"`
	// The result of the preview.
	//
	// Any of "success", "error".
	Result WorkflowStepPreviewTemplateResponseResult `json:"result,required"`
	// The rendered template, ready to be previewed.
	Template WorkflowStepPreviewTemplateResponseTemplateUnion `json:"template,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		Result      respjson.Field
		Template    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepPreviewTemplateResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStepPreviewTemplateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The content type of the preview.
type WorkflowStepPreviewTemplateResponseContentType string

const (
	WorkflowStepPreviewTemplateResponseContentTypeEmail     WorkflowStepPreviewTemplateResponseContentType = "email"
	WorkflowStepPreviewTemplateResponseContentTypeInAppFeed WorkflowStepPreviewTemplateResponseContentType = "in_app_feed"
	WorkflowStepPreviewTemplateResponseContentTypePush      WorkflowStepPreviewTemplateResponseContentType = "push"
	WorkflowStepPreviewTemplateResponseContentTypeChat      WorkflowStepPreviewTemplateResponseContentType = "chat"
	WorkflowStepPreviewTemplateResponseContentTypeSMS       WorkflowStepPreviewTemplateResponseContentType = "sms"
	WorkflowStepPreviewTemplateResponseContentTypeHTTP      WorkflowStepPreviewTemplateResponseContentType = "http"
)

// The result of the preview.
type WorkflowStepPreviewTemplateResponseResult string

const (
	WorkflowStepPreviewTemplateResponseResultSuccess WorkflowStepPreviewTemplateResponseResult = "success"
	WorkflowStepPreviewTemplateResponseResultError   WorkflowStepPreviewTemplateResponseResult = "error"
)

// WorkflowStepPreviewTemplateResponseTemplateUnion contains all possible
// properties and values from [EmailTemplate], [InAppFeedTemplate], [PushTemplate],
// [ChatTemplate], [SMSTemplate], [RequestTemplate].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkflowStepPreviewTemplateResponseTemplateUnion struct {
	// This field is a union of [EmailTemplateSettings], [PushTemplateSettings],
	// [SMSTemplateSettings]
	Settings WorkflowStepPreviewTemplateResponseTemplateUnionSettings `json:"settings"`
	// This field is from variant [EmailTemplate].
	Subject string `json:"subject"`
	// This field is from variant [EmailTemplate].
	HTMLBody string `json:"html_body"`
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
	// This field is from variant [RequestTemplate].
	Method RequestTemplateMethod `json:"method"`
	// This field is from variant [RequestTemplate].
	URL string `json:"url"`
	// This field is from variant [RequestTemplate].
	Body string `json:"body"`
	// This field is from variant [RequestTemplate].
	Headers RequestTemplateHeadersUnion `json:"headers"`
	// This field is from variant [RequestTemplate].
	QueryParams RequestTemplateQueryParamsUnion `json:"query_params"`
	JSON        struct {
		Settings      respjson.Field
		Subject       respjson.Field
		HTMLBody      respjson.Field
		TextBody      respjson.Field
		VisualBlocks  respjson.Field
		MarkdownBody  respjson.Field
		ActionButtons respjson.Field
		ActionURL     respjson.Field
		Title         respjson.Field
		JsonBody      respjson.Field
		Summary       respjson.Field
		Method        respjson.Field
		URL           respjson.Field
		Body          respjson.Field
		Headers       respjson.Field
		QueryParams   respjson.Field
		raw           string
	} `json:"-"`
}

func (u WorkflowStepPreviewTemplateResponseTemplateUnion) AsEmailTemplate() (v EmailTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepPreviewTemplateResponseTemplateUnion) AsInAppFeedTemplate() (v InAppFeedTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepPreviewTemplateResponseTemplateUnion) AsPushTemplate() (v PushTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepPreviewTemplateResponseTemplateUnion) AsChatTemplate() (v ChatTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepPreviewTemplateResponseTemplateUnion) AsSMSTemplate() (v SMSTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepPreviewTemplateResponseTemplateUnion) AsRequestTemplate() (v RequestTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowStepPreviewTemplateResponseTemplateUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkflowStepPreviewTemplateResponseTemplateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowStepPreviewTemplateResponseTemplateUnionSettings is an implicit subunion
// of [WorkflowStepPreviewTemplateResponseTemplateUnion].
// WorkflowStepPreviewTemplateResponseTemplateUnionSettings provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowStepPreviewTemplateResponseTemplateUnion].
type WorkflowStepPreviewTemplateResponseTemplateUnionSettings struct {
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

func (r *WorkflowStepPreviewTemplateResponseTemplateUnionSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowStepPreviewTemplateParams struct {
	WorkflowKey string `path:"workflow_key,required" json:"-"`
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// A recipient reference, used when referencing a recipient by either their ID (for
	// a user), or by a reference for an object.
	Recipient WorkflowStepPreviewTemplateParamsRecipientUnion `json:"recipient,omitzero,required"`
	// The tenant to associate the workflow with. Must not contain whitespace.
	Tenant param.Opt[string] `json:"tenant,omitzero"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// A recipient reference, used when referencing a recipient by either their ID (for
	// a user), or by a reference for an object.
	Actor WorkflowStepPreviewTemplateParamsActorUnion `json:"actor,omitzero"`
	// The data to pass to the workflow template for rendering.
	Data map[string]any `json:"data,omitzero"`
	paramObj
}

func (r WorkflowStepPreviewTemplateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepPreviewTemplateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepPreviewTemplateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WorkflowStepPreviewTemplateParams]'s query parameters as
// `url.Values`.
func (r WorkflowStepPreviewTemplateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowStepPreviewTemplateParamsRecipientUnion struct {
	OfString                   param.Opt[string]                                                   `json:",omitzero,inline"`
	OfObjectRecipientReference *WorkflowStepPreviewTemplateParamsRecipientObjectRecipientReference `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowStepPreviewTemplateParamsRecipientUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfObjectRecipientReference)
}
func (u *WorkflowStepPreviewTemplateParamsRecipientUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowStepPreviewTemplateParamsRecipientUnion) asAny() any {
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
type WorkflowStepPreviewTemplateParamsRecipientObjectRecipientReference struct {
	// The ID of the object.
	ID string `json:"id,required"`
	// The collection of the object.
	Collection string `json:"collection,required"`
	paramObj
}

func (r WorkflowStepPreviewTemplateParamsRecipientObjectRecipientReference) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepPreviewTemplateParamsRecipientObjectRecipientReference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepPreviewTemplateParamsRecipientObjectRecipientReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowStepPreviewTemplateParamsActorUnion struct {
	OfString                   param.Opt[string]                                               `json:",omitzero,inline"`
	OfObjectRecipientReference *WorkflowStepPreviewTemplateParamsActorObjectRecipientReference `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowStepPreviewTemplateParamsActorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfObjectRecipientReference)
}
func (u *WorkflowStepPreviewTemplateParamsActorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowStepPreviewTemplateParamsActorUnion) asAny() any {
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
type WorkflowStepPreviewTemplateParamsActorObjectRecipientReference struct {
	// The ID of the object.
	ID string `json:"id,required"`
	// The collection of the object.
	Collection string `json:"collection,required"`
	paramObj
}

func (r WorkflowStepPreviewTemplateParamsActorObjectRecipientReference) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepPreviewTemplateParamsActorObjectRecipientReference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepPreviewTemplateParamsActorObjectRecipientReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
