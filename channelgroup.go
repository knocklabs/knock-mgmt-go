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

// ChannelGroupService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChannelGroupService] method instead.
type ChannelGroupService struct {
	Options []option.RequestOption
}

// NewChannelGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChannelGroupService(opts ...option.RequestOption) (r ChannelGroupService) {
	r = ChannelGroupService{}
	r.Options = opts
	return
}

// Returns a paginated list of channel groups. Note: the list of channel groups is
// across the entire account, not scoped to an environment.
func (r *ChannelGroupService) List(ctx context.Context, query ChannelGroupListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[ChannelGroup], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/channel_groups"
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

// Returns a paginated list of channel groups. Note: the list of channel groups is
// across the entire account, not scoped to an environment.
func (r *ChannelGroupService) ListAutoPaging(ctx context.Context, query ChannelGroupListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[ChannelGroup] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// A group of channels with rules for when they are applicable.
type ChannelGroup struct {
	// Rules for determining which channels should be used.
	ChannelRules []ChannelGroupRule `json:"channel_rules,required"`
	// The type of channels contained in this group.
	//
	// Any of "email", "in_app", "in_app_feed", "in_app_guide", "sms", "push", "chat",
	// "http".
	ChannelType ChannelGroupChannelType `json:"channel_type,required"`
	// The timestamp of when the channel group was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Unique identifier for the channel group within a project.
	Key string `json:"key,required"`
	// The human-readable name of the channel group.
	Name string `json:"name,required"`
	// Determines how the channel rules are applied ('any' means any rule can match,
	// 'all' means all rules must match).
	//
	// Any of "any", "all".
	Operator ChannelGroupOperator `json:"operator,required"`
	// Whether this channel group was created by the system or a user. Only user
	// created channel groups can be modified.
	//
	// Any of "system", "user".
	Source ChannelGroupSource `json:"source,required"`
	// The timestamp of when the channel group was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelRules respjson.Field
		ChannelType  respjson.Field
		CreatedAt    respjson.Field
		Key          respjson.Field
		Name         respjson.Field
		Operator     respjson.Field
		Source       respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelGroup) RawJSON() string { return r.JSON.raw }
func (r *ChannelGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of channels contained in this group.
type ChannelGroupChannelType string

const (
	ChannelGroupChannelTypeEmail      ChannelGroupChannelType = "email"
	ChannelGroupChannelTypeInApp      ChannelGroupChannelType = "in_app"
	ChannelGroupChannelTypeInAppFeed  ChannelGroupChannelType = "in_app_feed"
	ChannelGroupChannelTypeInAppGuide ChannelGroupChannelType = "in_app_guide"
	ChannelGroupChannelTypeSMS        ChannelGroupChannelType = "sms"
	ChannelGroupChannelTypePush       ChannelGroupChannelType = "push"
	ChannelGroupChannelTypeChat       ChannelGroupChannelType = "chat"
	ChannelGroupChannelTypeHTTP       ChannelGroupChannelType = "http"
)

// Determines how the channel rules are applied ('any' means any rule can match,
// 'all' means all rules must match).
type ChannelGroupOperator string

const (
	ChannelGroupOperatorAny ChannelGroupOperator = "any"
	ChannelGroupOperatorAll ChannelGroupOperator = "all"
)

// Whether this channel group was created by the system or a user. Only user
// created channel groups can be modified.
type ChannelGroupSource string

const (
	ChannelGroupSourceSystem ChannelGroupSource = "system"
	ChannelGroupSourceUser   ChannelGroupSource = "user"
)

// A rule that determines if a channel should be executed as part of a channel
// group.
type ChannelGroupRule struct {
	// A configured channel, which is a way to route messages to a provider.
	Channel Channel `json:"channel,required"`
	// The timestamp of when the rule was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The order index of this rule within the channel group.
	Index int64 `json:"index,required"`
	// The type of rule (if = conditional, unless = negative conditional, always =
	// always apply).
	//
	// Any of "if", "unless", "always".
	RuleType ChannelGroupRuleRuleType `json:"rule_type,required"`
	// The timestamp of when the rule was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// For conditional rules, the value to compare against.
	Argument string `json:"argument,nullable"`
	// For conditional rules, the operator to apply.
	//
	// Any of "equal_to", "not_equal_to", "greater_than", "less_than",
	// "greater_than_or_equal_to", "less_than_or_equal_to", "contains", "not_contains",
	// "contains_all", "not_contains_all", "empty", "not_empty", "is_audience_member",
	// "is_not_audience_member".
	Operator ChannelGroupRuleOperator `json:"operator,nullable"`
	// For conditional rules, the variable to evaluate.
	Variable string `json:"variable,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channel     respjson.Field
		CreatedAt   respjson.Field
		Index       respjson.Field
		RuleType    respjson.Field
		UpdatedAt   respjson.Field
		Argument    respjson.Field
		Operator    respjson.Field
		Variable    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChannelGroupRule) RawJSON() string { return r.JSON.raw }
func (r *ChannelGroupRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of rule (if = conditional, unless = negative conditional, always =
// always apply).
type ChannelGroupRuleRuleType string

const (
	ChannelGroupRuleRuleTypeIf     ChannelGroupRuleRuleType = "if"
	ChannelGroupRuleRuleTypeUnless ChannelGroupRuleRuleType = "unless"
	ChannelGroupRuleRuleTypeAlways ChannelGroupRuleRuleType = "always"
)

// For conditional rules, the operator to apply.
type ChannelGroupRuleOperator string

const (
	ChannelGroupRuleOperatorEqualTo              ChannelGroupRuleOperator = "equal_to"
	ChannelGroupRuleOperatorNotEqualTo           ChannelGroupRuleOperator = "not_equal_to"
	ChannelGroupRuleOperatorGreaterThan          ChannelGroupRuleOperator = "greater_than"
	ChannelGroupRuleOperatorLessThan             ChannelGroupRuleOperator = "less_than"
	ChannelGroupRuleOperatorGreaterThanOrEqualTo ChannelGroupRuleOperator = "greater_than_or_equal_to"
	ChannelGroupRuleOperatorLessThanOrEqualTo    ChannelGroupRuleOperator = "less_than_or_equal_to"
	ChannelGroupRuleOperatorContains             ChannelGroupRuleOperator = "contains"
	ChannelGroupRuleOperatorNotContains          ChannelGroupRuleOperator = "not_contains"
	ChannelGroupRuleOperatorContainsAll          ChannelGroupRuleOperator = "contains_all"
	ChannelGroupRuleOperatorNotContainsAll       ChannelGroupRuleOperator = "not_contains_all"
	ChannelGroupRuleOperatorEmpty                ChannelGroupRuleOperator = "empty"
	ChannelGroupRuleOperatorNotEmpty             ChannelGroupRuleOperator = "not_empty"
	ChannelGroupRuleOperatorIsAudienceMember     ChannelGroupRuleOperator = "is_audience_member"
	ChannelGroupRuleOperatorIsNotAudienceMember  ChannelGroupRuleOperator = "is_not_audience_member"
)

type ChannelGroupListParams struct {
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ChannelGroupListParams]'s query parameters as `url.Values`.
func (r ChannelGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
