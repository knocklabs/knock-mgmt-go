// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/knock-mapi-go/internal/apijson"
	"github.com/stainless-sdks/knock-mapi-go/internal/apiquery"
	"github.com/stainless-sdks/knock-mapi-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-mapi-go/option"
	"github.com/stainless-sdks/knock-mapi-go/packages/pagination"
	"github.com/stainless-sdks/knock-mapi-go/packages/param"
	"github.com/stainless-sdks/knock-mapi-go/packages/resp"
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
	opts = append(r.Options[:], opts...)
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
	// The timestamp of when the channel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Unique identifier for the channel within a project (immutable once created).
	Key string `json:"key,required"`
	// The human-readable name of the channel.
	Name string `json:"name,required"`
	// The ID of the provider that this channel uses to deliver messages. Learn more
	// about the providers available
	// [in our documentation](https://docs.knock.app/integrations/overview).
	Provider string `json:"provider,required"`
	// The type of channel, determining what kind of messages it can send.
	//
	// Any of "email", "in_app", "in_app_feed", "in_app_guide", "sms", "push", "chat",
	// "http".
	Type ChannelType `json:"type,required"`
	// The timestamp of when the channel was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The timestamp of when the channel was deleted.
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	// Optional URL to a custom icon for the channel. Only used for display purposes in
	// the dashboard.
	CustomIconURL string `json:"custom_icon_url,nullable"`
	// Optional description of the channel's purpose or usage.
	Description string `json:"description,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CreatedAt     resp.Field
		Key           resp.Field
		Name          resp.Field
		Provider      resp.Field
		Type          resp.Field
		UpdatedAt     resp.Field
		ArchivedAt    resp.Field
		CustomIconURL resp.Field
		Description   resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
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

// Chat channel settings. Only used as configuration as part of a workflow channel
// step.
type ChatChannelSettings struct {
	// Whether to resolve chat provider user IDs using a Knock user's email address.
	// Only relevant for Slack channels for the time being.
	EmailBasedUserIDResolution bool `json:"email_based_user_id_resolution"`
	// Whether to track link clicks on chat notifications.
	LinkTracking bool `json:"link_tracking"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		EmailBasedUserIDResolution resp.Field
		LinkTracking               resp.Field
		ExtraFields                map[string]resp.Field
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
// ChatChannelSettingsParam.IsOverridden()
func (r ChatChannelSettings) ToParam() ChatChannelSettingsParam {
	return param.OverrideObj[ChatChannelSettingsParam](r.RawJSON())
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ChatChannelSettingsParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ChatChannelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatChannelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Email channel settings. Only used as configuration as part of a workflow channel
// step.
type EmailChannelSettings struct {
	// The BCC address on email notifications. Supports liquid.
	BccAddress string `json:"bcc_address,nullable"`
	// The CC address on email notifications. Supports liquid.
	CcAddress string `json:"cc_address,nullable"`
	// The email address from which this channel will send. Supports liquid.
	FromAddress string `json:"from_address,nullable"`
	// The name from which this channel will send. Supports liquid.
	FromName string `json:"from_name,nullable"`
	// A JSON template for any custom overrides to merge into the API payload that is
	// sent to the email provider. Supports liquid.
	JsonOverrides string `json:"json_overrides,nullable"`
	// Whether to track link clicks on email notifications.
	LinkTracking bool `json:"link_tracking"`
	// Whether to track opens on email notifications.
	OpenTracking bool `json:"open_tracking"`
	// The Reply-to address on email notifications. Supports liquid.
	ReplyToAddress string `json:"reply_to_address,nullable"`
	// The email address to which this channel will send. Defaults to
	// `recipient.email`. Supports liquid.
	ToAddress string `json:"to_address"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		BccAddress     resp.Field
		CcAddress      resp.Field
		FromAddress    resp.Field
		FromName       resp.Field
		JsonOverrides  resp.Field
		LinkTracking   resp.Field
		OpenTracking   resp.Field
		ReplyToAddress resp.Field
		ToAddress      resp.Field
		ExtraFields    map[string]resp.Field
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
// EmailChannelSettingsParam.IsOverridden()
func (r EmailChannelSettings) ToParam() EmailChannelSettingsParam {
	return param.OverrideObj[EmailChannelSettingsParam](r.RawJSON())
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f EmailChannelSettingsParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r EmailChannelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailChannelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// In-app feed channel settings. Only used as configuration as part of a workflow
// channel step.
type InAppFeedChannelSettings struct {
	// Whether to track link clicks on in-app feed notifications.
	LinkTracking bool `json:"link_tracking"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		LinkTracking resp.Field
		ExtraFields  map[string]resp.Field
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
// InAppFeedChannelSettingsParam.IsOverridden()
func (r InAppFeedChannelSettings) ToParam() InAppFeedChannelSettingsParam {
	return param.OverrideObj[InAppFeedChannelSettingsParam](r.RawJSON())
}

// In-app feed channel settings. Only used as configuration as part of a workflow
// channel step.
type InAppFeedChannelSettingsParam struct {
	// Whether to track link clicks on in-app feed notifications.
	LinkTracking param.Opt[bool] `json:"link_tracking,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f InAppFeedChannelSettingsParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r InAppFeedChannelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow InAppFeedChannelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Push channel settings. Only used as configuration as part of a workflow channel
// step.
type PushChannelSettings struct {
	// Whether to deregister a push-token when a push send hard bounces. This is to
	// prevent the same token from being used for future pushes.
	TokenDeregistration bool `json:"token_deregistration"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		TokenDeregistration resp.Field
		ExtraFields         map[string]resp.Field
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
// PushChannelSettingsParam.IsOverridden()
func (r PushChannelSettings) ToParam() PushChannelSettingsParam {
	return param.OverrideObj[PushChannelSettingsParam](r.RawJSON())
}

// Push channel settings. Only used as configuration as part of a workflow channel
// step.
type PushChannelSettingsParam struct {
	// Whether to deregister a push-token when a push send hard bounces. This is to
	// prevent the same token from being used for future pushes.
	TokenDeregistration param.Opt[bool] `json:"token_deregistration,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f PushChannelSettingsParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r PushChannelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow PushChannelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// SMS channel settings. Only used as configuration as part of a workflow channel
// step.
type SMSChannelSettings struct {
	// Whether to track link clicks on SMS notifications.
	LinkTracking bool `json:"link_tracking"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		LinkTracking resp.Field
		ExtraFields  map[string]resp.Field
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
// SMSChannelSettingsParam.IsOverridden()
func (r SMSChannelSettings) ToParam() SMSChannelSettingsParam {
	return param.OverrideObj[SMSChannelSettingsParam](r.RawJSON())
}

// SMS channel settings. Only used as configuration as part of a workflow channel
// step.
type SMSChannelSettingsParam struct {
	// Whether to track link clicks on SMS notifications.
	LinkTracking param.Opt[bool] `json:"link_tracking,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SMSChannelSettingsParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r SMSChannelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow SMSChannelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type ChannelListParams struct {
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ChannelListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [ChannelListParams]'s query parameters as `url.Values`.
func (r ChannelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
