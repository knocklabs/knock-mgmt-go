// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"encoding/json"
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

// ChannelService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelService] method instead.
type ChannelService struct {
	Options []option.RequestOption
}

// NewChannelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChannelService(opts ...option.RequestOption) (r ChannelService) {
	r = ChannelService{}
	r.Options = opts
	return
}

// Returns a paginated list of channels. Note: the list of channels is across the
// entire account, not scoped to an environment.
func (r *ChannelService) List(ctx context.Context, query ChannelListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Channel], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/channels"
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

// Returns a paginated list of channels. Note: the list of channels is across the
// entire account, not scoped to an environment.
func (r *ChannelService) ListAutoPaging(ctx context.Context, query ChannelListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Channel] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// A configured channel, which is a way to route messages to a provider.
type Channel struct {
	// The unique identifier for the channel.
	ID string `json:"id" api:"required"`
	// The timestamp of when the channel was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unique identifier for the channel within a project (immutable once created).
	Key string `json:"key" api:"required"`
	// The human-readable name of the channel.
	Name string `json:"name" api:"required"`
	// The ID of the provider that this channel uses to deliver messages. Learn more
	// about the providers available
	// [in our documentation](https://docs.knock.app/integrations/overview).
	Provider string `json:"provider" api:"required"`
	// The type of channel, determining what kind of messages it can send.
	//
	// Any of "email", "in_app", "in_app_feed", "in_app_guide", "sms", "push", "chat",
	// "http".
	Type ChannelType `json:"type" api:"required"`
	// The timestamp of when the channel was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The timestamp of when the channel was deleted.
	ArchivedAt time.Time `json:"archived_at" api:"nullable" format:"date-time"`
	// Optional URL to a custom icon for the channel. Only used for display purposes in
	// the dashboard.
	CustomIconURL string `json:"custom_icon_url" api:"nullable"`
	// Optional description of the channel's purpose or usage.
	Description string `json:"description" api:"nullable"`
	// Per-environment settings for this channel, keyed by environment slug (e.g.,
	// 'development', 'production'). Only included when requested via the `include`
	// parameter or when retrieving a single channel.
	EnvironmentSettings map[string]ChannelEnvironmentSetting `json:"environment_settings" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		Key                 respjson.Field
		Name                respjson.Field
		Provider            respjson.Field
		Type                respjson.Field
		UpdatedAt           respjson.Field
		ArchivedAt          respjson.Field
		CustomIconURL       respjson.Field
		Description         respjson.Field
		EnvironmentSettings respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Channel) RawJSON() string { return r.JSON.raw }
func (r *Channel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of channel, determining what kind of messages it can send.
type ChannelType string

const (
	ChannelTypeEmail      ChannelType = "email"
	ChannelTypeInApp      ChannelType = "in_app"
	ChannelTypeInAppFeed  ChannelType = "in_app_feed"
	ChannelTypeInAppGuide ChannelType = "in_app_guide"
	ChannelTypeSMS        ChannelType = "sms"
	ChannelTypePush       ChannelType = "push"
	ChannelTypeChat       ChannelType = "chat"
	ChannelTypeHTTP       ChannelType = "http"
)

// Environment-specific settings for a channel.
type ChannelEnvironmentSetting struct {
	// The unique identifier for these environment settings.
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether the channel is in sandbox mode for this environment. Sandbox mode may
	// prevent actual message delivery.
	IsSandbox bool `json:"is_sandbox" api:"required"`
	// Whether the channel configuration is valid and ready to send messages in this
	// environment.
	IsValid bool `json:"is_valid" api:"required"`
	// Channel-type-specific settings (e.g., from_address for email). Structure varies
	// by channel type.
	ChannelSettings map[string]any `json:"channel_settings" api:"nullable"`
	// Provider-specific settings (e.g., API keys, credentials). Structure varies by
	// provider. Secret values are obfuscated unless they are Liquid templates.
	ProviderSettings map[string]any `json:"provider_settings" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		IsSandbox        respjson.Field
		IsValid          respjson.Field
		ChannelSettings  respjson.Field
		ProviderSettings respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelEnvironmentSetting) RawJSON() string { return r.JSON.raw }
func (r *ChannelEnvironmentSetting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chat channel settings. Only used as configuration as part of a workflow channel
// step.
type ChatChannelSettings struct {
	// Whether to resolve chat provider user IDs using a Knock user's email address.
	// Only relevant for Slack channels for the time being.
	EmailBasedUserIDResolution bool `json:"email_based_user_id_resolution"`
	// Whether to track link clicks on chat notifications.
	LinkTracking bool `json:"link_tracking"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmailBasedUserIDResolution respjson.Field
		LinkTracking               respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatChannelSettings) RawJSON() string { return r.JSON.raw }
func (r *ChatChannelSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ChatChannelSettings to a ChatChannelSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChatChannelSettingsParam.Overrides()
func (r ChatChannelSettings) ToParam() ChatChannelSettingsParam {
	return param.Override[ChatChannelSettingsParam](json.RawMessage(r.RawJSON()))
}

// Chat channel settings. Only used as configuration as part of a workflow channel
// step.
type ChatChannelSettingsParam struct {
	// Whether to resolve chat provider user IDs using a Knock user's email address.
	// Only relevant for Slack channels for the time being.
	EmailBasedUserIDResolution param.Opt[bool] `json:"email_based_user_id_resolution,omitzero"`
	// Whether to track link clicks on chat notifications.
	LinkTracking param.Opt[bool] `json:"link_tracking,omitzero"`
	paramObj
}

func (r ChatChannelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatChannelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatChannelSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Email channel settings. Only used as configuration as part of a workflow channel
// step.
type EmailChannelSettings struct {
	// The BCC address on email notifications. Supports liquid.
	BccAddress string `json:"bcc_address" api:"nullable"`
	// The CC address on email notifications. Supports liquid.
	CcAddress string `json:"cc_address" api:"nullable"`
	// The email address from which this channel will send. Supports liquid.
	FromAddress string `json:"from_address" api:"nullable"`
	// The name from which this channel will send. Supports liquid.
	FromName string `json:"from_name" api:"nullable"`
	// A JSON template for any custom overrides to merge into the API payload that is
	// sent to the email provider. Supports liquid.
	JsonOverrides string `json:"json_overrides" api:"nullable"`
	// Whether to track link clicks on email notifications.
	LinkTracking bool `json:"link_tracking"`
	// Whether to track opens on email notifications.
	OpenTracking bool `json:"open_tracking"`
	// The Reply-to address on email notifications. Supports liquid.
	ReplyToAddress string `json:"reply_to_address" api:"nullable"`
	// The email address to which this channel will send. Defaults to
	// `recipient.email`. Supports liquid.
	ToAddress string `json:"to_address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BccAddress     respjson.Field
		CcAddress      respjson.Field
		FromAddress    respjson.Field
		FromName       respjson.Field
		JsonOverrides  respjson.Field
		LinkTracking   respjson.Field
		OpenTracking   respjson.Field
		ReplyToAddress respjson.Field
		ToAddress      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailChannelSettings) RawJSON() string { return r.JSON.raw }
func (r *EmailChannelSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EmailChannelSettings to a EmailChannelSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EmailChannelSettingsParam.Overrides()
func (r EmailChannelSettings) ToParam() EmailChannelSettingsParam {
	return param.Override[EmailChannelSettingsParam](json.RawMessage(r.RawJSON()))
}

// Email channel settings. Only used as configuration as part of a workflow channel
// step.
type EmailChannelSettingsParam struct {
	// The BCC address on email notifications. Supports liquid.
	BccAddress param.Opt[string] `json:"bcc_address,omitzero"`
	// The CC address on email notifications. Supports liquid.
	CcAddress param.Opt[string] `json:"cc_address,omitzero"`
	// The email address from which this channel will send. Supports liquid.
	FromAddress param.Opt[string] `json:"from_address,omitzero"`
	// The name from which this channel will send. Supports liquid.
	FromName param.Opt[string] `json:"from_name,omitzero"`
	// A JSON template for any custom overrides to merge into the API payload that is
	// sent to the email provider. Supports liquid.
	JsonOverrides param.Opt[string] `json:"json_overrides,omitzero"`
	// The Reply-to address on email notifications. Supports liquid.
	ReplyToAddress param.Opt[string] `json:"reply_to_address,omitzero"`
	// Whether to track link clicks on email notifications.
	LinkTracking param.Opt[bool] `json:"link_tracking,omitzero"`
	// Whether to track opens on email notifications.
	OpenTracking param.Opt[bool] `json:"open_tracking,omitzero"`
	// The email address to which this channel will send. Defaults to
	// `recipient.email`. Supports liquid.
	ToAddress param.Opt[string] `json:"to_address,omitzero"`
	paramObj
}

func (r EmailChannelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailChannelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailChannelSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// In-app feed channel settings. Only used as configuration as part of a workflow
// channel step.
type InAppFeedChannelSettings struct {
	// Whether to track link clicks on in-app feed notifications.
	LinkTracking bool `json:"link_tracking"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkTracking respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InAppFeedChannelSettings) RawJSON() string { return r.JSON.raw }
func (r *InAppFeedChannelSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InAppFeedChannelSettings to a
// InAppFeedChannelSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InAppFeedChannelSettingsParam.Overrides()
func (r InAppFeedChannelSettings) ToParam() InAppFeedChannelSettingsParam {
	return param.Override[InAppFeedChannelSettingsParam](json.RawMessage(r.RawJSON()))
}

// In-app feed channel settings. Only used as configuration as part of a workflow
// channel step.
type InAppFeedChannelSettingsParam struct {
	// Whether to track link clicks on in-app feed notifications.
	LinkTracking param.Opt[bool] `json:"link_tracking,omitzero"`
	paramObj
}

func (r InAppFeedChannelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow InAppFeedChannelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InAppFeedChannelSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Push channel settings. Only used as configuration as part of a workflow channel
// step.
type PushChannelSettings struct {
	// Whether to deregister a push-token when a push send hard bounces. This is to
	// prevent the same token from being used for future pushes.
	TokenDeregistration bool `json:"token_deregistration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TokenDeregistration respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PushChannelSettings) RawJSON() string { return r.JSON.raw }
func (r *PushChannelSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PushChannelSettings to a PushChannelSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PushChannelSettingsParam.Overrides()
func (r PushChannelSettings) ToParam() PushChannelSettingsParam {
	return param.Override[PushChannelSettingsParam](json.RawMessage(r.RawJSON()))
}

// Push channel settings. Only used as configuration as part of a workflow channel
// step.
type PushChannelSettingsParam struct {
	// Whether to deregister a push-token when a push send hard bounces. This is to
	// prevent the same token from being used for future pushes.
	TokenDeregistration param.Opt[bool] `json:"token_deregistration,omitzero"`
	paramObj
}

func (r PushChannelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow PushChannelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PushChannelSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SMS channel settings. Only used as configuration as part of a workflow channel
// step.
type SMSChannelSettings struct {
	// Whether to track link clicks on SMS notifications.
	LinkTracking bool `json:"link_tracking"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkTracking respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SMSChannelSettings) RawJSON() string { return r.JSON.raw }
func (r *SMSChannelSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SMSChannelSettings to a SMSChannelSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SMSChannelSettingsParam.Overrides()
func (r SMSChannelSettings) ToParam() SMSChannelSettingsParam {
	return param.Override[SMSChannelSettingsParam](json.RawMessage(r.RawJSON()))
}

// SMS channel settings. Only used as configuration as part of a workflow channel
// step.
type SMSChannelSettingsParam struct {
	// Whether to track link clicks on SMS notifications.
	LinkTracking param.Opt[bool] `json:"link_tracking,omitzero"`
	paramObj
}

func (r SMSChannelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow SMSChannelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SMSChannelSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelListParams struct {
	// A channel id to filter the results by.
	ID param.Opt[string] `query:"id,omitzero" format:"uuid" json:"-"`
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Associated resources to include in the response. Accepts `environment_settings`
	// to include per-environment channel configuration.
	//
	// Any of "environment_settings".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChannelListParams]'s query parameters as `url.Values`.
func (r ChannelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
