// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
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

// AuditLogService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuditLogService] method instead.
type AuditLogService struct {
	Options []option.RequestOption
}

// NewAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuditLogService(opts ...option.RequestOption) (r AuditLogService) {
	r = AuditLogService{}
	r.Options = opts
	return
}

// Returns a paginated list of audit log entries for the current account, across
// every environment and branch, including account-level entries that have no
// environment. Entries are ordered from most recent first.
//
// The list includes user-visible changes and environment merges. Commits are
// returned by the commits endpoint. Entries outside the account's data retention
// window are omitted.
func (r *AuditLogService) List(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[AuditLogListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/audit_logs"
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

// Returns a paginated list of audit log entries for the current account, across
// every environment and branch, including account-level entries that have no
// environment. Entries are ordered from most recent first.
//
// The list includes user-visible changes and environment merges. Commits are
// returned by the commits endpoint. Entries outside the account's data retention
// window are omitted.
func (r *AuditLogService) ListAutoPaging(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[AuditLogListResponse] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// An account audit log entry. Commit rows are returned by the commits endpoint.
// The entry does not include the stored change diff.
type AuditLogListResponse struct {
	// The unique identifier for the audit log entry.
	ID string `json:"id" api:"required" format:"uuid"`
	// The action that was performed, such as `workflow.updated` or `seat.removed`.
	Action string `json:"action" api:"required"`
	// The Knock user who performed the action.
	Actor AuditLogListResponseActor `json:"actor" api:"required"`
	// The agent session that performed the action, when the source is an agent.
	AgentSessionID string `json:"agent_session_id" api:"required" format:"uuid"`
	// The slug of the environment or branch where the action happened. Null for
	// account-level actions.
	Environment string `json:"environment" api:"required"`
	// The timestamp of when the action happened.
	InsertedAt time.Time `json:"inserted_at" api:"required" format:"date-time"`
	// The IP address of the request that performed the action.
	IPAddress string `json:"ip_address" api:"required"`
	// The location derived from the request IP address.
	Location AuditLogListResponseLocation `json:"location" api:"required"`
	// The project the action belongs to. Null for account-level actions.
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// The id of the request that performed the action.
	RequestID string `json:"request_id" api:"required"`
	// Where the action originated.
	//
	// Any of "dashboard", "mapi", "agent".
	RequestSource AuditLogListResponseRequestSource `json:"request_source" api:"required"`
	// The resource the action changed. `type` and `identifier` follow the commit
	// resource for promotable resources, including `reusable_request`. `identifier` is
	// null when the resource has no public handle or the row no longer exists.
	Resource AuditLogListResponseResource `json:"resource" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Action         respjson.Field
		Actor          respjson.Field
		AgentSessionID respjson.Field
		Environment    respjson.Field
		InsertedAt     respjson.Field
		IPAddress      respjson.Field
		Location       respjson.Field
		ProjectID      respjson.Field
		RequestID      respjson.Field
		RequestSource  respjson.Field
		Resource       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditLogListResponse) RawJSON() string { return r.JSON.raw }
func (r *AuditLogListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Knock user who performed the action.
type AuditLogListResponseActor struct {
	// The unique identifier for the user.
	ID string `json:"id" api:"required" format:"uuid"`
	// The email address of the user.
	Email string `json:"email" api:"required"`
	// The name of the user.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditLogListResponseActor) RawJSON() string { return r.JSON.raw }
func (r *AuditLogListResponseActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The location derived from the request IP address.
type AuditLogListResponseLocation struct {
	// The city.
	City string `json:"city" api:"nullable"`
	// The country code.
	CountryCode string `json:"country_code" api:"nullable"`
	// The region code.
	RegionCode string `json:"region_code" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		CountryCode respjson.Field
		RegionCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditLogListResponseLocation) RawJSON() string { return r.JSON.raw }
func (r *AuditLogListResponseLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where the action originated.
type AuditLogListResponseRequestSource string

const (
	AuditLogListResponseRequestSourceDashboard AuditLogListResponseRequestSource = "dashboard"
	AuditLogListResponseRequestSourceMapi      AuditLogListResponseRequestSource = "mapi"
	AuditLogListResponseRequestSourceAgent     AuditLogListResponseRequestSource = "agent"
)

// The resource the action changed. `type` and `identifier` follow the commit
// resource for promotable resources, including `reusable_request`. `identifier` is
// null when the resource has no public handle or the row no longer exists.
type AuditLogListResponseResource struct {
	// The stored id of the resource. External user and object ids are not UUIDs.
	ID string `json:"id" api:"required"`
	// The type of the resource.
	//
	// Any of "account", "account_user", "api_key", "asset", "audience", "broadcast",
	// "channel", "channel_group", "email_layout", "environment",
	// "environment_defaults", "event_action_mapping", "external_audience", "goal",
	// "guide", "guide_group", "invite", "member", "message_type", "object", "partial",
	// "permission_group", "preference_category", "preference_center",
	// "preference_center_domain", "reusable_request", "service_token", "signing_key",
	// "source", "subscription", "tag", "translation", "user", "variable", "view",
	// "webhook", "workflow", "unknown".
	Type string `json:"type" api:"required"`
	// The public handle for the resource, such as a key, slug, email, or
	// `locale/namespace` for a translation.
	Identifier string `json:"identifier" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		Identifier  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditLogListResponseResource) RawJSON() string { return r.JSON.raw }
func (r *AuditLogListResponseResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditLogListParams struct {
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The environment slug. When omitted, the list is account-wide and does not use
	// the account's default environment. When set, limits entries to that environment
	// or branch and omits account-level entries.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only return entries inserted at or after this RFC 3339 timestamp.
	Since param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Only return entries inserted at or before this RFC 3339 timestamp.
	Until param.Opt[time.Time] `query:"until,omitzero" format:"date-time" json:"-"`
	// Filter by action. Accepts one action or an array of actions, matched exactly.
	Action AuditLogListParamsActionUnion `query:"action,omitzero" json:"-"`
	// Filter by the id of the user who performed the action. Accepts one id or an
	// array of ids.
	ActorID AuditLogListParamsActorIDUnion `query:"actor_id,omitzero" format:"uuid" json:"-"`
	// Filter by resource type. Accepts one type or an array of types. `channel`
	// includes channel environment settings, and `source` includes source environment
	// settings.
	ResourceType AuditLogListParamsResourceTypeUnion `query:"resource_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuditLogListParams]'s query parameters as `url.Values`.
func (r AuditLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AuditLogListParamsActionUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *AuditLogListParamsActionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AuditLogListParamsActorIDUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *AuditLogListParamsActorIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AuditLogListParamsResourceTypeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfAuditLogListsResourceTypeString)
	OfAuditLogListsResourceTypeString         param.Opt[string] `query:",omitzero,inline"`
	OfAuditLogListsResourceTypeArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *AuditLogListParamsResourceTypeUnion) asAny() any {
	if !param.IsOmitted(u.OfAuditLogListsResourceTypeString) {
		return &u.OfAuditLogListsResourceTypeString
	} else if !param.IsOmitted(u.OfAuditLogListsResourceTypeArrayItemArray) {
		return &u.OfAuditLogListsResourceTypeArrayItemArray
	}
	return nil
}

type AuditLogListParamsResourceTypeString string

const (
	AuditLogListParamsResourceTypeStringAccount                AuditLogListParamsResourceTypeString = "account"
	AuditLogListParamsResourceTypeStringAccountUser            AuditLogListParamsResourceTypeString = "account_user"
	AuditLogListParamsResourceTypeStringAPIKey                 AuditLogListParamsResourceTypeString = "api_key"
	AuditLogListParamsResourceTypeStringAsset                  AuditLogListParamsResourceTypeString = "asset"
	AuditLogListParamsResourceTypeStringAudience               AuditLogListParamsResourceTypeString = "audience"
	AuditLogListParamsResourceTypeStringBroadcast              AuditLogListParamsResourceTypeString = "broadcast"
	AuditLogListParamsResourceTypeStringChannel                AuditLogListParamsResourceTypeString = "channel"
	AuditLogListParamsResourceTypeStringChannelGroup           AuditLogListParamsResourceTypeString = "channel_group"
	AuditLogListParamsResourceTypeStringEmailLayout            AuditLogListParamsResourceTypeString = "email_layout"
	AuditLogListParamsResourceTypeStringEnvironment            AuditLogListParamsResourceTypeString = "environment"
	AuditLogListParamsResourceTypeStringEnvironmentDefaults    AuditLogListParamsResourceTypeString = "environment_defaults"
	AuditLogListParamsResourceTypeStringEventActionMapping     AuditLogListParamsResourceTypeString = "event_action_mapping"
	AuditLogListParamsResourceTypeStringExternalAudience       AuditLogListParamsResourceTypeString = "external_audience"
	AuditLogListParamsResourceTypeStringGoal                   AuditLogListParamsResourceTypeString = "goal"
	AuditLogListParamsResourceTypeStringGuide                  AuditLogListParamsResourceTypeString = "guide"
	AuditLogListParamsResourceTypeStringGuideGroup             AuditLogListParamsResourceTypeString = "guide_group"
	AuditLogListParamsResourceTypeStringInvite                 AuditLogListParamsResourceTypeString = "invite"
	AuditLogListParamsResourceTypeStringMember                 AuditLogListParamsResourceTypeString = "member"
	AuditLogListParamsResourceTypeStringMessageType            AuditLogListParamsResourceTypeString = "message_type"
	AuditLogListParamsResourceTypeStringObject                 AuditLogListParamsResourceTypeString = "object"
	AuditLogListParamsResourceTypeStringPartial                AuditLogListParamsResourceTypeString = "partial"
	AuditLogListParamsResourceTypeStringPermissionGroup        AuditLogListParamsResourceTypeString = "permission_group"
	AuditLogListParamsResourceTypeStringPreferenceCategory     AuditLogListParamsResourceTypeString = "preference_category"
	AuditLogListParamsResourceTypeStringPreferenceCenter       AuditLogListParamsResourceTypeString = "preference_center"
	AuditLogListParamsResourceTypeStringPreferenceCenterDomain AuditLogListParamsResourceTypeString = "preference_center_domain"
	AuditLogListParamsResourceTypeStringReusableRequest        AuditLogListParamsResourceTypeString = "reusable_request"
	AuditLogListParamsResourceTypeStringServiceToken           AuditLogListParamsResourceTypeString = "service_token"
	AuditLogListParamsResourceTypeStringSigningKey             AuditLogListParamsResourceTypeString = "signing_key"
	AuditLogListParamsResourceTypeStringSource                 AuditLogListParamsResourceTypeString = "source"
	AuditLogListParamsResourceTypeStringSubscription           AuditLogListParamsResourceTypeString = "subscription"
	AuditLogListParamsResourceTypeStringTag                    AuditLogListParamsResourceTypeString = "tag"
	AuditLogListParamsResourceTypeStringTranslation            AuditLogListParamsResourceTypeString = "translation"
	AuditLogListParamsResourceTypeStringUser                   AuditLogListParamsResourceTypeString = "user"
	AuditLogListParamsResourceTypeStringVariable               AuditLogListParamsResourceTypeString = "variable"
	AuditLogListParamsResourceTypeStringView                   AuditLogListParamsResourceTypeString = "view"
	AuditLogListParamsResourceTypeStringWebhook                AuditLogListParamsResourceTypeString = "webhook"
	AuditLogListParamsResourceTypeStringWorkflow               AuditLogListParamsResourceTypeString = "workflow"
)
