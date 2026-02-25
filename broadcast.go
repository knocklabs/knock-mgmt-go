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

// BroadcastService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBroadcastService] method instead.
type BroadcastService struct {
	Options []option.RequestOption
}

// NewBroadcastService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBroadcastService(opts ...option.RequestOption) (r BroadcastService) {
	r = BroadcastService{}
	r.Options = opts
	return
}

// Get a broadcast by its key in a given environment.
func (r *BroadcastService) Get(ctx context.Context, broadcastKey string, query BroadcastGetParams, opts ...option.RequestOption) (res *Broadcast, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastKey == "" {
		err = errors.New("missing required broadcast_key parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s", broadcastKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a paginated list of broadcasts available in a given environment. The
// broadcasts are returned ordered by creation time (newest first).
func (r *BroadcastService) List(ctx context.Context, query BroadcastListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Broadcast], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/broadcasts"
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

// Returns a paginated list of broadcasts available in a given environment. The
// broadcasts are returned ordered by creation time (newest first).
func (r *BroadcastService) ListAutoPaging(ctx context.Context, query BroadcastListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Broadcast] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Cancels sending a scheduled broadcast. The broadcast will return to draft
// status.
func (r *BroadcastService) Cancel(ctx context.Context, broadcastKey string, body BroadcastCancelParams, opts ...option.RequestOption) (res *BroadcastCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastKey == "" {
		err = errors.New("missing required broadcast_key parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s/cancel", broadcastKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Sends a broadcast immediately or schedules it to send at a future time.
func (r *BroadcastService) Send(ctx context.Context, broadcastKey string, params BroadcastSendParams, opts ...option.RequestOption) (res *BroadcastSendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastKey == "" {
		err = errors.New("missing required broadcast_key parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s/send", broadcastKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Updates a broadcast of a given key, or creates a new one if it does not yet
// exist.
func (r *BroadcastService) Upsert(ctx context.Context, broadcastKey string, params BroadcastUpsertParams, opts ...option.RequestOption) (res *BroadcastUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastKey == "" {
		err = errors.New("missing required broadcast_key parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s", broadcastKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Validates a broadcast payload without persisting it.
func (r *BroadcastService) Validate(ctx context.Context, broadcastKey string, params BroadcastValidateParams, opts ...option.RequestOption) (res *BroadcastValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if broadcastKey == "" {
		err = errors.New("missing required broadcast_key parameter")
		return
	}
	path := fmt.Sprintf("v1/broadcasts/%s/validate", broadcastKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// A broadcast object.
type Broadcast struct {
	// The timestamp of when the broadcast was created. (read-only).
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The slug of the environment in which the broadcast exists. (read-only).
	Environment string `json:"environment" api:"required"`
	// The unique key string for the broadcast object. Must be at minimum 3 characters
	// and at maximum 255 characters in length. Must be in the format of ^[a-z0-9_-]+$.
	Key string `json:"key" api:"required"`
	// A name for the broadcast. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// The SHA hash of the workflow data. (read-only).
	Sha string `json:"sha" api:"required"`
	// The current status of the broadcast. One of: `draft`, `scheduled`, `sent`.
	//
	// Any of "draft", "scheduled", "sent".
	Status BroadcastStatus `json:"status" api:"required"`
	// A list of broadcast step objects in the broadcast. Broadcasts only support
	// channel, branch, and delay steps.
	Steps []BroadcastStepUnion `json:"steps" api:"required"`
	// The timestamp of when the broadcast was last updated. (read-only).
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Whether the broadcast and its steps are in a valid state. (read-only).
	Valid bool `json:"valid" api:"required"`
	// The timestamp of when the broadcast was archived.
	ArchivedAt time.Time `json:"archived_at" api:"nullable" format:"date-time"`
	// A list of categories that the broadcast belongs to.
	Categories []string `json:"categories"`
	// An arbitrary string attached to a broadcast object. Useful for adding notes
	// about the broadcast for internal purposes. Maximum of 280 characters allowed.
	Description string `json:"description"`
	// The timestamp of when the broadcast is scheduled to be sent.
	ScheduledAt time.Time `json:"scheduled_at" api:"nullable" format:"date-time"`
	// The timestamp of when the broadcast was sent. (read-only).
	SentAt time.Time `json:"sent_at" api:"nullable" format:"date-time"`
	// A map of broadcast settings.
	Settings BroadcastSettings `json:"settings"`
	// The key of the audience to target for this broadcast.
	TargetAudienceKey string `json:"target_audience_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt         respjson.Field
		Environment       respjson.Field
		Key               respjson.Field
		Name              respjson.Field
		Sha               respjson.Field
		Status            respjson.Field
		Steps             respjson.Field
		UpdatedAt         respjson.Field
		Valid             respjson.Field
		ArchivedAt        respjson.Field
		Categories        respjson.Field
		Description       respjson.Field
		ScheduledAt       respjson.Field
		SentAt            respjson.Field
		Settings          respjson.Field
		TargetAudienceKey respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Broadcast) RawJSON() string { return r.JSON.raw }
func (r *Broadcast) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the broadcast. One of: `draft`, `scheduled`, `sent`.
type BroadcastStatus string

const (
	BroadcastStatusDraft     BroadcastStatus = "draft"
	BroadcastStatusScheduled BroadcastStatus = "scheduled"
	BroadcastStatusSent      BroadcastStatus = "sent"
)

// BroadcastStepUnion contains all possible properties and values from
// [WorkflowWebhookStep], [WorkflowInAppFeedStep], [WorkflowChatStep],
// [WorkflowSMSStep], [WorkflowPushStep], [WorkflowEmailStep],
// [WorkflowBranchStep], [WorkflowDelayStep], [WorkflowRandomCohortStep].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BroadcastStepUnion struct {
	Ref string `json:"ref"`
	// This field is a union of [WebhookTemplate], [InAppFeedTemplate], [ChatTemplate],
	// [SMSTemplate], [PushTemplate], [EmailTemplate]
	Template        BroadcastStepUnionTemplate `json:"template"`
	Type            string                     `json:"type"`
	ChannelGroupKey string                     `json:"channel_group_key"`
	ChannelKey      string                     `json:"channel_key"`
	ChannelType     string                     `json:"channel_type"`
	// This field is from variant [WorkflowWebhookStep].
	Conditions  ConditionGroupUnion `json:"conditions"`
	Description string              `json:"description"`
	Name        string              `json:"name"`
	SendWindows []SendWindow        `json:"send_windows"`
	// This field is a union of [InAppFeedChannelSettings], [ChatChannelSettings],
	// [SMSChannelSettings], [PushChannelSettings], [EmailChannelSettings]
	ChannelOverrides BroadcastStepUnionChannelOverrides `json:"channel_overrides"`
	// This field is from variant [WorkflowBranchStep].
	Branches []WorkflowBranchStepBranch `json:"branches"`
	// This field is from variant [WorkflowDelayStep].
	Settings WorkflowDelayStepSettings `json:"settings"`
	// This field is from variant [WorkflowRandomCohortStep].
	CohortBranches []any `json:"cohort_branches"`
	// This field is from variant [WorkflowRandomCohortStep].
	CohortKey string `json:"cohort_key"`
	JSON      struct {
		Ref              respjson.Field
		Template         respjson.Field
		Type             respjson.Field
		ChannelGroupKey  respjson.Field
		ChannelKey       respjson.Field
		ChannelType      respjson.Field
		Conditions       respjson.Field
		Description      respjson.Field
		Name             respjson.Field
		SendWindows      respjson.Field
		ChannelOverrides respjson.Field
		Branches         respjson.Field
		Settings         respjson.Field
		CohortBranches   respjson.Field
		CohortKey        respjson.Field
		raw              string
	} `json:"-"`
}

func (u BroadcastStepUnion) AsWorkflowWebhookStep() (v WorkflowWebhookStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BroadcastStepUnion) AsWorkflowInAppFeedStep() (v WorkflowInAppFeedStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BroadcastStepUnion) AsWorkflowChatStep() (v WorkflowChatStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BroadcastStepUnion) AsWorkflowSMSStep() (v WorkflowSMSStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BroadcastStepUnion) AsWorkflowPushStep() (v WorkflowPushStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BroadcastStepUnion) AsWorkflowEmailStep() (v WorkflowEmailStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BroadcastStepUnion) AsWorkflowBranchStep() (v WorkflowBranchStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BroadcastStepUnion) AsWorkflowDelayStep() (v WorkflowDelayStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BroadcastStepUnion) AsWorkflowRandomCohortStep() (v WorkflowRandomCohortStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BroadcastStepUnion) RawJSON() string { return u.JSON.raw }

func (r *BroadcastStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BroadcastStepUnionTemplate is an implicit subunion of [BroadcastStepUnion].
// BroadcastStepUnionTemplate provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [BroadcastStepUnion].
type BroadcastStepUnionTemplate struct {
	// This field is from variant [WebhookTemplate].
	Method WebhookTemplateMethod `json:"method"`
	// This field is from variant [WebhookTemplate].
	URL string `json:"url"`
	// This field is from variant [WebhookTemplate].
	Body string `json:"body"`
	// This field is from variant [WebhookTemplate].
	Headers []WebhookTemplateHeader `json:"headers"`
	// This field is from variant [WebhookTemplate].
	QueryParams  []WebhookTemplateQueryParam `json:"query_params"`
	MarkdownBody string                      `json:"markdown_body"`
	// This field is from variant [InAppFeedTemplate].
	ActionButtons []InAppFeedTemplateActionButton `json:"action_buttons"`
	// This field is from variant [InAppFeedTemplate].
	ActionURL string `json:"action_url"`
	// This field is from variant [ChatTemplate].
	JsonBody string `json:"json_body"`
	// This field is from variant [ChatTemplate].
	Summary  string `json:"summary"`
	TextBody string `json:"text_body"`
	// This field is a union of [SMSTemplateSettings], [PushTemplateSettings],
	// [EmailTemplateSettings]
	Settings BroadcastStepUnionTemplateSettings `json:"settings"`
	// This field is from variant [PushTemplate].
	Title string `json:"title"`
	// This field is from variant [EmailTemplate].
	Subject string `json:"subject"`
	// This field is from variant [EmailTemplate].
	HTMLBody string `json:"html_body"`
	// This field is from variant [EmailTemplate].
	VisualBlocks []EmailTemplateVisualBlockUnion `json:"visual_blocks"`
	JSON         struct {
		Method        respjson.Field
		URL           respjson.Field
		Body          respjson.Field
		Headers       respjson.Field
		QueryParams   respjson.Field
		MarkdownBody  respjson.Field
		ActionButtons respjson.Field
		ActionURL     respjson.Field
		JsonBody      respjson.Field
		Summary       respjson.Field
		TextBody      respjson.Field
		Settings      respjson.Field
		Title         respjson.Field
		Subject       respjson.Field
		HTMLBody      respjson.Field
		VisualBlocks  respjson.Field
		raw           string
	} `json:"-"`
}

func (r *BroadcastStepUnionTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BroadcastStepUnionTemplateSettings is an implicit subunion of
// [BroadcastStepUnion]. BroadcastStepUnionTemplateSettings provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BroadcastStepUnion].
type BroadcastStepUnionTemplateSettings struct {
	PayloadOverrides string `json:"payload_overrides"`
	// This field is from variant [SMSTemplateSettings].
	ToNumber string `json:"to_number"`
	// This field is from variant [PushTemplateSettings].
	DeliveryType string `json:"delivery_type"`
	// This field is from variant [EmailTemplateSettings].
	AttachmentKey string `json:"attachment_key"`
	// This field is from variant [EmailTemplateSettings].
	LayoutKey string `json:"layout_key"`
	// This field is from variant [EmailTemplateSettings].
	PreContent string `json:"pre_content"`
	JSON       struct {
		PayloadOverrides respjson.Field
		ToNumber         respjson.Field
		DeliveryType     respjson.Field
		AttachmentKey    respjson.Field
		LayoutKey        respjson.Field
		PreContent       respjson.Field
		raw              string
	} `json:"-"`
}

func (r *BroadcastStepUnionTemplateSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BroadcastStepUnionChannelOverrides is an implicit subunion of
// [BroadcastStepUnion]. BroadcastStepUnionChannelOverrides provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [BroadcastStepUnion].
type BroadcastStepUnionChannelOverrides struct {
	LinkTracking bool `json:"link_tracking"`
	// This field is from variant [ChatChannelSettings].
	EmailBasedUserIDResolution bool `json:"email_based_user_id_resolution"`
	// This field is from variant [PushChannelSettings].
	TokenDeregistration bool `json:"token_deregistration"`
	// This field is from variant [EmailChannelSettings].
	BccAddress string `json:"bcc_address"`
	// This field is from variant [EmailChannelSettings].
	CcAddress string `json:"cc_address"`
	// This field is from variant [EmailChannelSettings].
	FromAddress string `json:"from_address"`
	// This field is from variant [EmailChannelSettings].
	FromName string `json:"from_name"`
	// This field is from variant [EmailChannelSettings].
	JsonOverrides string `json:"json_overrides"`
	// This field is from variant [EmailChannelSettings].
	OpenTracking bool `json:"open_tracking"`
	// This field is from variant [EmailChannelSettings].
	ReplyToAddress string `json:"reply_to_address"`
	// This field is from variant [EmailChannelSettings].
	ToAddress string `json:"to_address"`
	JSON      struct {
		LinkTracking               respjson.Field
		EmailBasedUserIDResolution respjson.Field
		TokenDeregistration        respjson.Field
		BccAddress                 respjson.Field
		CcAddress                  respjson.Field
		FromAddress                respjson.Field
		FromName                   respjson.Field
		JsonOverrides              respjson.Field
		OpenTracking               respjson.Field
		ReplyToAddress             respjson.Field
		ToAddress                  respjson.Field
		raw                        string
	} `json:"-"`
}

func (r *BroadcastStepUnionChannelOverrides) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A map of broadcast settings.
type BroadcastSettings struct {
	// Whether the broadcast is commercial. Defaults to true.
	IsCommercial bool `json:"is_commercial"`
	// Whether to ignore recipient preferences for a given type of notification. If
	// true, will send for every channel in the workflow even if the recipient has
	// opted out of a certain kind. Defaults to false.
	OverridePreferences bool `json:"override_preferences"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsCommercial        respjson.Field
		OverridePreferences respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastSettings) RawJSON() string { return r.JSON.raw }
func (r *BroadcastSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A broadcast request for upserting a broadcast.
//
// The properties Name, Steps are required.
type BroadcastRequestParam struct {
	// A name for the broadcast. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// A list of broadcast step objects in the broadcast. Broadcasts only support
	// channel, branch, and delay steps.
	Steps []BroadcastRequestStepUnionParam `json:"steps,omitzero" api:"required"`
	// The timestamp of when the broadcast is scheduled to be sent.
	ScheduledAt param.Opt[time.Time] `json:"scheduled_at,omitzero" format:"date-time"`
	// An arbitrary string attached to a broadcast object. Useful for adding notes
	// about the broadcast for internal purposes. Maximum of 280 characters allowed.
	Description param.Opt[string] `json:"description,omitzero"`
	// The key of the audience to target for this broadcast.
	TargetAudienceKey param.Opt[string] `json:"target_audience_key,omitzero"`
	// A list of categories that the broadcast belongs to.
	Categories []string `json:"categories,omitzero"`
	// A map of broadcast settings.
	Settings BroadcastRequestSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r BroadcastRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BroadcastRequestStepUnionParam struct {
	OfWorkflowWebhookStep      *WorkflowWebhookStepParam      `json:",omitzero,inline"`
	OfWorkflowInAppFeedStep    *WorkflowInAppFeedStepParam    `json:",omitzero,inline"`
	OfWorkflowChatStep         *WorkflowChatStepParam         `json:",omitzero,inline"`
	OfWorkflowSMSStep          *WorkflowSMSStepParam          `json:",omitzero,inline"`
	OfWorkflowPushStep         *WorkflowPushStepParam         `json:",omitzero,inline"`
	OfWorkflowEmailStep        *WorkflowEmailStepParam        `json:",omitzero,inline"`
	OfWorkflowBranchStep       *WorkflowBranchStepParam       `json:",omitzero,inline"`
	OfWorkflowDelayStep        *WorkflowDelayStepParam        `json:",omitzero,inline"`
	OfWorkflowRandomCohortStep *WorkflowRandomCohortStepParam `json:",omitzero,inline"`
	paramUnion
}

func (u BroadcastRequestStepUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWorkflowWebhookStep,
		u.OfWorkflowInAppFeedStep,
		u.OfWorkflowChatStep,
		u.OfWorkflowSMSStep,
		u.OfWorkflowPushStep,
		u.OfWorkflowEmailStep,
		u.OfWorkflowBranchStep,
		u.OfWorkflowDelayStep,
		u.OfWorkflowRandomCohortStep)
}
func (u *BroadcastRequestStepUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BroadcastRequestStepUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWorkflowWebhookStep) {
		return u.OfWorkflowWebhookStep
	} else if !param.IsOmitted(u.OfWorkflowInAppFeedStep) {
		return u.OfWorkflowInAppFeedStep
	} else if !param.IsOmitted(u.OfWorkflowChatStep) {
		return u.OfWorkflowChatStep
	} else if !param.IsOmitted(u.OfWorkflowSMSStep) {
		return u.OfWorkflowSMSStep
	} else if !param.IsOmitted(u.OfWorkflowPushStep) {
		return u.OfWorkflowPushStep
	} else if !param.IsOmitted(u.OfWorkflowEmailStep) {
		return u.OfWorkflowEmailStep
	} else if !param.IsOmitted(u.OfWorkflowBranchStep) {
		return u.OfWorkflowBranchStep
	} else if !param.IsOmitted(u.OfWorkflowDelayStep) {
		return u.OfWorkflowDelayStep
	} else if !param.IsOmitted(u.OfWorkflowRandomCohortStep) {
		return u.OfWorkflowRandomCohortStep
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BroadcastRequestStepUnionParam) GetBranches() []WorkflowBranchStepBranchParam {
	if vt := u.OfWorkflowBranchStep; vt != nil {
		return vt.Branches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BroadcastRequestStepUnionParam) GetSettings() *WorkflowDelayStepSettingsParam {
	if vt := u.OfWorkflowDelayStep; vt != nil {
		return &vt.Settings
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BroadcastRequestStepUnionParam) GetCohortBranches() []any {
	if vt := u.OfWorkflowRandomCohortStep; vt != nil {
		return vt.CohortBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BroadcastRequestStepUnionParam) GetCohortKey() *string {
	if vt := u.OfWorkflowRandomCohortStep; vt != nil && vt.CohortKey.Valid() {
		return &vt.CohortKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BroadcastRequestStepUnionParam) GetRef() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowChatStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowSMSStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowPushStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowEmailStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowBranchStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowDelayStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowRandomCohortStep; vt != nil {
		return (*string)(&vt.Ref)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BroadcastRequestStepUnionParam) GetType() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowChatStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowSMSStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowPushStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowEmailStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowBranchStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowDelayStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowRandomCohortStep; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BroadcastRequestStepUnionParam) GetChannelGroupKey() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil && vt.ChannelGroupKey.Valid() {
		return &vt.ChannelGroupKey.Value
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil && vt.ChannelGroupKey.Valid() {
		return &vt.ChannelGroupKey.Value
	} else if vt := u.OfWorkflowChatStep; vt != nil && vt.ChannelGroupKey.Valid() {
		return &vt.ChannelGroupKey.Value
	} else if vt := u.OfWorkflowSMSStep; vt != nil && vt.ChannelGroupKey.Valid() {
		return &vt.ChannelGroupKey.Value
	} else if vt := u.OfWorkflowPushStep; vt != nil && vt.ChannelGroupKey.Valid() {
		return &vt.ChannelGroupKey.Value
	} else if vt := u.OfWorkflowEmailStep; vt != nil && vt.ChannelGroupKey.Valid() {
		return &vt.ChannelGroupKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BroadcastRequestStepUnionParam) GetChannelKey() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil && vt.ChannelKey.Valid() {
		return &vt.ChannelKey.Value
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil && vt.ChannelKey.Valid() {
		return &vt.ChannelKey.Value
	} else if vt := u.OfWorkflowChatStep; vt != nil && vt.ChannelKey.Valid() {
		return &vt.ChannelKey.Value
	} else if vt := u.OfWorkflowSMSStep; vt != nil && vt.ChannelKey.Valid() {
		return &vt.ChannelKey.Value
	} else if vt := u.OfWorkflowPushStep; vt != nil && vt.ChannelKey.Valid() {
		return &vt.ChannelKey.Value
	} else if vt := u.OfWorkflowEmailStep; vt != nil && vt.ChannelKey.Valid() {
		return &vt.ChannelKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BroadcastRequestStepUnionParam) GetChannelType() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil {
		return (*string)(&vt.ChannelType)
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil {
		return (*string)(&vt.ChannelType)
	} else if vt := u.OfWorkflowChatStep; vt != nil {
		return (*string)(&vt.ChannelType)
	} else if vt := u.OfWorkflowSMSStep; vt != nil {
		return (*string)(&vt.ChannelType)
	} else if vt := u.OfWorkflowPushStep; vt != nil {
		return (*string)(&vt.ChannelType)
	} else if vt := u.OfWorkflowEmailStep; vt != nil {
		return (*string)(&vt.ChannelType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BroadcastRequestStepUnionParam) GetDescription() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowChatStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowSMSStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowPushStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowEmailStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowBranchStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowDelayStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowRandomCohortStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BroadcastRequestStepUnionParam) GetName() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowChatStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowSMSStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowPushStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowEmailStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowBranchStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowDelayStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowRandomCohortStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u BroadcastRequestStepUnionParam) GetTemplate() (res broadcastRequestStepUnionParamTemplate) {
	if vt := u.OfWorkflowWebhookStep; vt != nil {
		res.any = &vt.Template
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil {
		res.any = &vt.Template
	} else if vt := u.OfWorkflowChatStep; vt != nil {
		res.any = &vt.Template
	} else if vt := u.OfWorkflowSMSStep; vt != nil {
		res.any = &vt.Template
	} else if vt := u.OfWorkflowPushStep; vt != nil {
		res.any = &vt.Template
	} else if vt := u.OfWorkflowEmailStep; vt != nil {
		res.any = &vt.Template
	}
	return
}

// Can have the runtime types [*WebhookTemplateParam], [*InAppFeedTemplateParam],
// [*ChatTemplateParam], [*SMSTemplateParam], [*PushTemplateParam],
// [*EmailTemplateParam]
type broadcastRequestStepUnionParamTemplate struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.WebhookTemplateParam:
//	case *knockmapi.InAppFeedTemplateParam:
//	case *knockmapi.ChatTemplateParam:
//	case *knockmapi.SMSTemplateParam:
//	case *knockmapi.PushTemplateParam:
//	case *knockmapi.EmailTemplateParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u broadcastRequestStepUnionParamTemplate) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetMethod() *string {
	switch vt := u.any.(type) {
	case *WebhookTemplateParam:
		return (*string)(&vt.Method)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetURL() *string {
	switch vt := u.any.(type) {
	case *WebhookTemplateParam:
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetBody() *string {
	switch vt := u.any.(type) {
	case *WebhookTemplateParam:
		return paramutil.AddrIfPresent(vt.Body)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetHeaders() []WebhookTemplateHeaderParam {
	switch vt := u.any.(type) {
	case *WebhookTemplateParam:
		return vt.Headers
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetQueryParams() []WebhookTemplateQueryParamParam {
	switch vt := u.any.(type) {
	case *WebhookTemplateParam:
		return vt.QueryParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetActionButtons() []InAppFeedTemplateActionButtonParam {
	switch vt := u.any.(type) {
	case *InAppFeedTemplateParam:
		return vt.ActionButtons
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetActionURL() *string {
	switch vt := u.any.(type) {
	case *InAppFeedTemplateParam:
		return paramutil.AddrIfPresent(vt.ActionURL)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetJsonBody() *string {
	switch vt := u.any.(type) {
	case *ChatTemplateParam:
		return paramutil.AddrIfPresent(vt.JsonBody)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetSummary() *string {
	switch vt := u.any.(type) {
	case *ChatTemplateParam:
		return paramutil.AddrIfPresent(vt.Summary)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetTitle() *string {
	switch vt := u.any.(type) {
	case *PushTemplateParam:
		return &vt.Title
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetSubject() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateParam:
		return &vt.Subject
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetHTMLBody() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateParam:
		return paramutil.AddrIfPresent(vt.HTMLBody)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetVisualBlocks() []EmailTemplateVisualBlockUnionParam {
	switch vt := u.any.(type) {
	case *EmailTemplateParam:
		return vt.VisualBlocks
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetMarkdownBody() *string {
	switch vt := u.any.(type) {
	case *InAppFeedTemplateParam:
		return (*string)(&vt.MarkdownBody)
	case *ChatTemplateParam:
		return (*string)(&vt.MarkdownBody)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplate) GetTextBody() *string {
	switch vt := u.any.(type) {
	case *SMSTemplateParam:
		return (*string)(&vt.TextBody)
	case *PushTemplateParam:
		return (*string)(&vt.TextBody)
	case *EmailTemplateParam:
		return paramutil.AddrIfPresent(vt.TextBody)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u broadcastRequestStepUnionParamTemplate) GetSettings() (res broadcastRequestStepUnionParamTemplateSettings) {
	switch vt := u.any.(type) {
	case *SMSTemplateParam:
		res.any = &vt.Settings
	case *PushTemplateParam:
		res.any = &vt.Settings
	case *EmailTemplateParam:
		res.any = &vt.Settings
	}
	return res
}

// Can have the runtime types [*SMSTemplateSettingsParam],
// [*PushTemplateSettingsParam], [*EmailTemplateSettingsParam]
type broadcastRequestStepUnionParamTemplateSettings struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.SMSTemplateSettingsParam:
//	case *knockmapi.PushTemplateSettingsParam:
//	case *knockmapi.EmailTemplateSettingsParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u broadcastRequestStepUnionParamTemplateSettings) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplateSettings) GetToNumber() *string {
	switch vt := u.any.(type) {
	case *SMSTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.ToNumber)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplateSettings) GetDeliveryType() *string {
	switch vt := u.any.(type) {
	case *PushTemplateSettingsParam:
		return &vt.DeliveryType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplateSettings) GetAttachmentKey() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.AttachmentKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplateSettings) GetLayoutKey() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.LayoutKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplateSettings) GetPreContent() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.PreContent)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamTemplateSettings) GetPayloadOverrides() *string {
	switch vt := u.any.(type) {
	case *SMSTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.PayloadOverrides)
	case *PushTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.PayloadOverrides)
	}
	return nil
}

// Returns a pointer to the underlying variant's Conditions property, if present.
func (u BroadcastRequestStepUnionParam) GetConditions() *ConditionGroupUnionParam {
	if vt := u.OfWorkflowWebhookStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowChatStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowSMSStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowPushStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowEmailStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowDelayStep; vt != nil {
		return &vt.Conditions
	}
	return nil
}

// Returns a pointer to the underlying variant's SendWindows property, if present.
func (u BroadcastRequestStepUnionParam) GetSendWindows() []SendWindowParam {
	if vt := u.OfWorkflowWebhookStep; vt != nil {
		return vt.SendWindows
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil {
		return vt.SendWindows
	} else if vt := u.OfWorkflowChatStep; vt != nil {
		return vt.SendWindows
	} else if vt := u.OfWorkflowSMSStep; vt != nil {
		return vt.SendWindows
	} else if vt := u.OfWorkflowPushStep; vt != nil {
		return vt.SendWindows
	} else if vt := u.OfWorkflowEmailStep; vt != nil {
		return vt.SendWindows
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u BroadcastRequestStepUnionParam) GetChannelOverrides() (res broadcastRequestStepUnionParamChannelOverrides) {
	if vt := u.OfWorkflowInAppFeedStep; vt != nil {
		res.any = &vt.ChannelOverrides
	} else if vt := u.OfWorkflowChatStep; vt != nil {
		res.any = &vt.ChannelOverrides
	} else if vt := u.OfWorkflowSMSStep; vt != nil {
		res.any = &vt.ChannelOverrides
	} else if vt := u.OfWorkflowPushStep; vt != nil {
		res.any = &vt.ChannelOverrides
	} else if vt := u.OfWorkflowEmailStep; vt != nil {
		res.any = &vt.ChannelOverrides
	}
	return
}

// Can have the runtime types [*InAppFeedChannelSettingsParam],
// [*ChatChannelSettingsParam], [*SMSChannelSettingsParam],
// [*PushChannelSettingsParam], [*EmailChannelSettingsParam]
type broadcastRequestStepUnionParamChannelOverrides struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.InAppFeedChannelSettingsParam:
//	case *knockmapi.ChatChannelSettingsParam:
//	case *knockmapi.SMSChannelSettingsParam:
//	case *knockmapi.PushChannelSettingsParam:
//	case *knockmapi.EmailChannelSettingsParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u broadcastRequestStepUnionParamChannelOverrides) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamChannelOverrides) GetEmailBasedUserIDResolution() *bool {
	switch vt := u.any.(type) {
	case *ChatChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.EmailBasedUserIDResolution)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamChannelOverrides) GetTokenDeregistration() *bool {
	switch vt := u.any.(type) {
	case *PushChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.TokenDeregistration)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamChannelOverrides) GetBccAddress() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.BccAddress)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamChannelOverrides) GetCcAddress() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.CcAddress)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamChannelOverrides) GetFromAddress() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.FromAddress)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamChannelOverrides) GetFromName() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.FromName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamChannelOverrides) GetJsonOverrides() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.JsonOverrides)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamChannelOverrides) GetOpenTracking() *bool {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.OpenTracking)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamChannelOverrides) GetReplyToAddress() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.ReplyToAddress)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamChannelOverrides) GetToAddress() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.ToAddress)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u broadcastRequestStepUnionParamChannelOverrides) GetLinkTracking() *bool {
	switch vt := u.any.(type) {
	case *InAppFeedChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.LinkTracking)
	case *ChatChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.LinkTracking)
	case *SMSChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.LinkTracking)
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.LinkTracking)
	}
	return nil
}

// A map of broadcast settings.
type BroadcastRequestSettingsParam struct {
	// Whether the broadcast is commercial. Defaults to true.
	IsCommercial param.Opt[bool] `json:"is_commercial,omitzero"`
	// Whether to ignore recipient preferences for a given type of notification. If
	// true, will send for every channel in the workflow even if the recipient has
	// opted out of a certain kind. Defaults to false.
	OverridePreferences param.Opt[bool] `json:"override_preferences,omitzero"`
	paramObj
}

func (r BroadcastRequestSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastRequestSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastRequestSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Broadcast response under the `broadcast` key.
type BroadcastCancelResponse struct {
	// A broadcast object.
	Broadcast Broadcast `json:"broadcast" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Broadcast   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *BroadcastCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Broadcast response under the `broadcast` key.
type BroadcastSendResponse struct {
	// A broadcast object.
	Broadcast Broadcast `json:"broadcast" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Broadcast   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastSendResponse) RawJSON() string { return r.JSON.raw }
func (r *BroadcastSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Broadcast response under the `broadcast` key.
type BroadcastUpsertResponse struct {
	// A broadcast object.
	Broadcast Broadcast `json:"broadcast" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Broadcast   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *BroadcastUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Broadcast response under the `broadcast` key.
type BroadcastValidateResponse struct {
	// A broadcast object.
	Broadcast Broadcast `json:"broadcast" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Broadcast   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BroadcastValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *BroadcastValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BroadcastGetParams struct {
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

// URLQuery serializes [BroadcastGetParams]'s query parameters as `url.Values`.
func (r BroadcastGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BroadcastListParams struct {
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

// URLQuery serializes [BroadcastListParams]'s query parameters as `url.Values`.
func (r BroadcastListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BroadcastCancelParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BroadcastCancelParams]'s query parameters as `url.Values`.
func (r BroadcastCancelParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BroadcastSendParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// When to send the broadcast. If provided, the broadcast will be scheduled to send
	// at this time. Must be in ISO 8601 UTC format. If not provided, the broadcast
	// will be sent immediately.
	SendAt param.Opt[time.Time] `json:"send_at,omitzero" format:"date-time"`
	paramObj
}

func (r BroadcastSendParams) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BroadcastSendParams]'s query parameters as `url.Values`.
func (r BroadcastSendParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BroadcastUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A broadcast request for upserting a broadcast.
	Broadcast BroadcastRequestParam `json:"broadcast,omitzero" api:"required"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	paramObj
}

func (r BroadcastUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BroadcastUpsertParams]'s query parameters as `url.Values`.
func (r BroadcastUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BroadcastValidateParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A broadcast request for upserting a broadcast.
	Broadcast BroadcastRequestParam `json:"broadcast,omitzero" api:"required"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	paramObj
}

func (r BroadcastValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow BroadcastValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BroadcastValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BroadcastValidateParams]'s query parameters as
// `url.Values`.
func (r BroadcastValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
