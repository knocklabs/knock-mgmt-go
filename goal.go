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
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/pagination"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// Goals define event conditions that are tracked and attributed to messaging
// resources.
//
// GoalService contains methods and other services that help with interacting with
// the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGoalService] method instead.
type GoalService struct {
	Options []option.RequestOption
}

// NewGoalService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGoalService(opts ...option.RequestOption) (r GoalService) {
	r = GoalService{}
	r.Options = opts
	return
}

// Retrieve a goal by its key in a given environment.
func (r *GoalService) Get(ctx context.Context, goalKey string, query GoalGetParams, opts ...option.RequestOption) (res *Goal, err error) {
	opts = slices.Concat(r.Options, opts)
	if goalKey == "" {
		err = errors.New("missing required goal_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/goals/%s", goalKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of goals for the given environment.
func (r *GoalService) List(ctx context.Context, query GoalListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Goal], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/goals"
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

// Returns a paginated list of goals for the given environment.
func (r *GoalService) ListAutoPaging(ctx context.Context, query GoalListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Goal] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Archives a given goal across all environments. Refuses if any workflow, guide,
// or broadcast is attached to the goal.
func (r *GoalService) Archive(ctx context.Context, goalKey string, body GoalArchiveParams, opts ...option.RequestOption) (res *GoalArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if goalKey == "" {
		err = errors.New("missing required goal_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/goals/%s", goalKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Clones a goal into a destination environment.
func (r *GoalService) Clone(ctx context.Context, goalKey string, params GoalCloneParams, opts ...option.RequestOption) (res *GoalCloneResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if goalKey == "" {
		err = errors.New("missing required goal_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/goals/%s/clone", goalKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Updates a goal of a given key, or creates a new one if it does not yet exist.
// The goal is published immediately; this endpoint does not accept a commit
// parameter.
func (r *GoalService) Upsert(ctx context.Context, goalKey string, params GoalUpsertParams, opts ...option.RequestOption) (res *GoalUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if goalKey == "" {
		err = errors.New("missing required goal_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/goals/%s", goalKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Validates a goal payload without persisting it.
func (r *GoalService) Validate(ctx context.Context, goalKey string, params GoalValidateParams, opts ...option.RequestOption) (res *GoalValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if goalKey == "" {
		err = errors.New("missing required goal_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/goals/%s/validate", goalKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A goal defines an event condition that is tracked and attributed to messaging
// resources.
type Goal struct {
	// A goal condition consisting of a polymorphic event and optional match
	// conditions.
	Condition GoalCondition `json:"condition" api:"required"`
	// The timestamp of when the goal was created. (read-only).
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The slug of the environment in which the goal exists. (read-only).
	Environment string `json:"environment" api:"required"`
	// The unique key string for the goal. Must be at minimum 1 character and at
	// maximum 255 characters in length.
	Key string `json:"key" api:"required"`
	// A name for the goal. Must be at minimum 1 character and at maximum 255
	// characters in length.
	Name string `json:"name" api:"required"`
	// The SHA hash of the goal data. (read-only).
	Sha string `json:"sha" api:"required"`
	// The timestamp of when the goal was last updated. (read-only).
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional description for the goal. Maximum of 280 characters allowed.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Condition   respjson.Field
		CreatedAt   respjson.Field
		Environment respjson.Field
		Key         respjson.Field
		Name        respjson.Field
		Sha         respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Goal) RawJSON() string { return r.JSON.raw }
func (r *Goal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A goal condition consisting of a polymorphic event and optional match
// conditions.
type GoalCondition struct {
	// The event to track. Supports recipient, integration_source, and audience event
	// types.
	Event GoalConditionEventUnion `json:"event" api:"required"`
	// A list of condition groups. Required for recipient events; each group uses an
	// operator (and/or) with nested conditions.
	MatchConditions []ConditionGroupUnion `json:"match_conditions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event           respjson.Field
		MatchConditions respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalCondition) RawJSON() string { return r.JSON.raw }
func (r *GoalCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GoalCondition to a GoalConditionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GoalConditionParam.Overrides()
func (r GoalCondition) ToParam() GoalConditionParam {
	return param.Override[GoalConditionParam](json.RawMessage(r.RawJSON()))
}

// GoalConditionEventUnion contains all possible properties and values from
// [GoalConditionEventWorkflowWaitForEventRecipientEvent],
// [GoalConditionEventWorkflowWaitForEventIntegrationSourceEvent],
// [GoalConditionEventWorkflowWaitForEventAudienceEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type GoalConditionEventUnion struct {
	EventType string `json:"event_type"`
	EventKey  string `json:"event_key"`
	// This field is from variant
	// [GoalConditionEventWorkflowWaitForEventIntegrationSourceEvent].
	IntegrationSourceKey string `json:"integration_source_key"`
	// This field is from variant
	// [GoalConditionEventWorkflowWaitForEventAudienceEvent].
	AudienceKey string `json:"audience_key"`
	JSON        struct {
		EventType            respjson.Field
		EventKey             respjson.Field
		IntegrationSourceKey respjson.Field
		AudienceKey          respjson.Field
		raw                  string
	} `json:"-"`
}

func (u GoalConditionEventUnion) AsWorkflowWaitForEventRecipientEvent() (v GoalConditionEventWorkflowWaitForEventRecipientEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GoalConditionEventUnion) AsWorkflowWaitForEventIntegrationSourceEvent() (v GoalConditionEventWorkflowWaitForEventIntegrationSourceEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GoalConditionEventUnion) AsWorkflowWaitForEventAudienceEvent() (v GoalConditionEventWorkflowWaitForEventAudienceEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GoalConditionEventUnion) RawJSON() string { return u.JSON.raw }

func (r *GoalConditionEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A recipient updated event to wait for from the workflow recipient.
type GoalConditionEventWorkflowWaitForEventRecipientEvent struct {
	// The type of event to wait for.
	//
	// Any of "recipient".
	EventType string `json:"event_type" api:"required"`
	// Recipient lifecycle event to wait for. Always "updated" today.
	//
	// Any of "updated".
	EventKey string `json:"event_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventType   respjson.Field
		EventKey    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalConditionEventWorkflowWaitForEventRecipientEvent) RawJSON() string { return r.JSON.raw }
func (r *GoalConditionEventWorkflowWaitForEventRecipientEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An integration source event to wait for.
type GoalConditionEventWorkflowWaitForEventIntegrationSourceEvent struct {
	// The name of the event to wait for.
	EventKey string `json:"event_key" api:"required"`
	// The type of event to wait for.
	//
	// Any of "integration_source".
	EventType string `json:"event_type" api:"required"`
	// The key of the integration source that emits the event to wait for.
	IntegrationSourceKey string `json:"integration_source_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventKey             respjson.Field
		EventType            respjson.Field
		IntegrationSourceKey respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalConditionEventWorkflowWaitForEventIntegrationSourceEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *GoalConditionEventWorkflowWaitForEventIntegrationSourceEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An audience membership event to wait for when a recipient enters or exits an
// audience.
type GoalConditionEventWorkflowWaitForEventAudienceEvent struct {
	// The key of the audience to wait for membership changes.
	AudienceKey string `json:"audience_key" api:"required"`
	// The audience membership transition to wait for.
	//
	// Any of "enter", "exit".
	EventKey string `json:"event_key" api:"required"`
	// The type of event to wait for.
	//
	// Any of "audience".
	EventType string `json:"event_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudienceKey respjson.Field
		EventKey    respjson.Field
		EventType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalConditionEventWorkflowWaitForEventAudienceEvent) RawJSON() string { return r.JSON.raw }
func (r *GoalConditionEventWorkflowWaitForEventAudienceEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A goal condition consisting of a polymorphic event and optional match
// conditions.
//
// The property Event is required.
type GoalConditionParam struct {
	// The event to track. Supports recipient, integration_source, and audience event
	// types.
	Event GoalConditionEventUnionParam `json:"event,omitzero" api:"required"`
	// A list of condition groups. Required for recipient events; each group uses an
	// operator (and/or) with nested conditions.
	MatchConditions []ConditionGroupUnionParam `json:"match_conditions,omitzero"`
	paramObj
}

func (r GoalConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow GoalConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoalConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GoalConditionEventUnionParam struct {
	OfWorkflowWaitForEventRecipientEvent         *GoalConditionEventWorkflowWaitForEventRecipientEventParam         `json:",omitzero,inline"`
	OfWorkflowWaitForEventIntegrationSourceEvent *GoalConditionEventWorkflowWaitForEventIntegrationSourceEventParam `json:",omitzero,inline"`
	OfWorkflowWaitForEventAudienceEvent          *GoalConditionEventWorkflowWaitForEventAudienceEventParam          `json:",omitzero,inline"`
	paramUnion
}

func (u GoalConditionEventUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWorkflowWaitForEventRecipientEvent, u.OfWorkflowWaitForEventIntegrationSourceEvent, u.OfWorkflowWaitForEventAudienceEvent)
}
func (u *GoalConditionEventUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GoalConditionEventUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWorkflowWaitForEventRecipientEvent) {
		return u.OfWorkflowWaitForEventRecipientEvent
	} else if !param.IsOmitted(u.OfWorkflowWaitForEventIntegrationSourceEvent) {
		return u.OfWorkflowWaitForEventIntegrationSourceEvent
	} else if !param.IsOmitted(u.OfWorkflowWaitForEventAudienceEvent) {
		return u.OfWorkflowWaitForEventAudienceEvent
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GoalConditionEventUnionParam) GetIntegrationSourceKey() *string {
	if vt := u.OfWorkflowWaitForEventIntegrationSourceEvent; vt != nil {
		return &vt.IntegrationSourceKey
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GoalConditionEventUnionParam) GetAudienceKey() *string {
	if vt := u.OfWorkflowWaitForEventAudienceEvent; vt != nil {
		return &vt.AudienceKey
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GoalConditionEventUnionParam) GetEventType() *string {
	if vt := u.OfWorkflowWaitForEventRecipientEvent; vt != nil {
		return (*string)(&vt.EventType)
	} else if vt := u.OfWorkflowWaitForEventIntegrationSourceEvent; vt != nil {
		return (*string)(&vt.EventType)
	} else if vt := u.OfWorkflowWaitForEventAudienceEvent; vt != nil {
		return (*string)(&vt.EventType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u GoalConditionEventUnionParam) GetEventKey() *string {
	if vt := u.OfWorkflowWaitForEventRecipientEvent; vt != nil {
		return (*string)(&vt.EventKey)
	} else if vt := u.OfWorkflowWaitForEventIntegrationSourceEvent; vt != nil {
		return (*string)(&vt.EventKey)
	} else if vt := u.OfWorkflowWaitForEventAudienceEvent; vt != nil {
		return (*string)(&vt.EventKey)
	}
	return nil
}

// A recipient updated event to wait for from the workflow recipient.
//
// The property EventType is required.
type GoalConditionEventWorkflowWaitForEventRecipientEventParam struct {
	// The type of event to wait for.
	//
	// Any of "recipient".
	EventType string `json:"event_type,omitzero" api:"required"`
	// Recipient lifecycle event to wait for. Always "updated" today.
	//
	// Any of "updated".
	EventKey string `json:"event_key,omitzero"`
	paramObj
}

func (r GoalConditionEventWorkflowWaitForEventRecipientEventParam) MarshalJSON() (data []byte, err error) {
	type shadow GoalConditionEventWorkflowWaitForEventRecipientEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoalConditionEventWorkflowWaitForEventRecipientEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GoalConditionEventWorkflowWaitForEventRecipientEventParam](
		"event_type", "recipient",
	)
	apijson.RegisterFieldValidator[GoalConditionEventWorkflowWaitForEventRecipientEventParam](
		"event_key", "updated",
	)
}

// An integration source event to wait for.
//
// The properties EventKey, EventType, IntegrationSourceKey are required.
type GoalConditionEventWorkflowWaitForEventIntegrationSourceEventParam struct {
	// The name of the event to wait for.
	EventKey string `json:"event_key" api:"required"`
	// The type of event to wait for.
	//
	// Any of "integration_source".
	EventType string `json:"event_type,omitzero" api:"required"`
	// The key of the integration source that emits the event to wait for.
	IntegrationSourceKey string `json:"integration_source_key" api:"required"`
	paramObj
}

func (r GoalConditionEventWorkflowWaitForEventIntegrationSourceEventParam) MarshalJSON() (data []byte, err error) {
	type shadow GoalConditionEventWorkflowWaitForEventIntegrationSourceEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoalConditionEventWorkflowWaitForEventIntegrationSourceEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GoalConditionEventWorkflowWaitForEventIntegrationSourceEventParam](
		"event_type", "integration_source",
	)
}

// An audience membership event to wait for when a recipient enters or exits an
// audience.
//
// The properties AudienceKey, EventKey, EventType are required.
type GoalConditionEventWorkflowWaitForEventAudienceEventParam struct {
	// The key of the audience to wait for membership changes.
	AudienceKey string `json:"audience_key" api:"required"`
	// The audience membership transition to wait for.
	//
	// Any of "enter", "exit".
	EventKey string `json:"event_key,omitzero" api:"required"`
	// The type of event to wait for.
	//
	// Any of "audience".
	EventType string `json:"event_type,omitzero" api:"required"`
	paramObj
}

func (r GoalConditionEventWorkflowWaitForEventAudienceEventParam) MarshalJSON() (data []byte, err error) {
	type shadow GoalConditionEventWorkflowWaitForEventAudienceEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoalConditionEventWorkflowWaitForEventAudienceEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GoalConditionEventWorkflowWaitForEventAudienceEventParam](
		"event_key", "enter", "exit",
	)
	apijson.RegisterFieldValidator[GoalConditionEventWorkflowWaitForEventAudienceEventParam](
		"event_type", "audience",
	)
}

// A goal payload for upsert or validate.
//
// The properties Condition, Name are required.
type GoalRequestParam struct {
	// A goal condition consisting of a polymorphic event and optional match
	// conditions.
	Condition GoalConditionParam `json:"condition,omitzero" api:"required"`
	// A name for the goal. Must be at minimum 1 character and at maximum 255
	// characters in length.
	Name string `json:"name" api:"required"`
	// An optional description for the goal. Maximum of 280 characters allowed.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r GoalRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow GoalRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoalRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response from archiving a goal.
type GoalArchiveResponse struct {
	// The result of the archive operation.
	Result string `json:"result" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *GoalArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Goal response under the `goal` key.
type GoalCloneResponse struct {
	// A goal defines an event condition that is tracked and attributed to messaging
	// resources.
	Goal Goal `json:"goal" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Goal        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalCloneResponse) RawJSON() string { return r.JSON.raw }
func (r *GoalCloneResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Goal response under the `goal` key.
type GoalUpsertResponse struct {
	// A goal defines an event condition that is tracked and attributed to messaging
	// resources.
	Goal Goal `json:"goal" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Goal        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *GoalUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Goal response under the `goal` key.
type GoalValidateResponse struct {
	// A goal defines an event condition that is tracked and attributed to messaging
	// resources.
	Goal Goal `json:"goal" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Goal        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *GoalValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalGetParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GoalGetParams]'s query parameters as `url.Values`.
func (r GoalGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GoalListParams struct {
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
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GoalListParams]'s query parameters as `url.Values`.
func (r GoalListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GoalArchiveParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [GoalArchiveParams]'s query parameters as `url.Values`.
func (r GoalArchiveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GoalCloneParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The destination key, name, and environment for the cloned goal.
	Clone GoalCloneParamsClone `json:"clone,omitzero" api:"required"`
	paramObj
}

func (r GoalCloneParams) MarshalJSON() (data []byte, err error) {
	type shadow GoalCloneParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoalCloneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [GoalCloneParams]'s query parameters as `url.Values`.
func (r GoalCloneParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The destination key, name, and environment for the cloned goal.
//
// The properties Environment, Key, Name are required.
type GoalCloneParamsClone struct {
	// The destination environment slug.
	Environment string `json:"environment" api:"required"`
	// The key for the cloned goal.
	Key string `json:"key" api:"required"`
	// The name for the cloned goal.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r GoalCloneParamsClone) MarshalJSON() (data []byte, err error) {
	type shadow GoalCloneParamsClone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoalCloneParamsClone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A goal payload for upsert or validate.
	Goal GoalRequestParam `json:"goal,omitzero" api:"required"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	paramObj
}

func (r GoalUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow GoalUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoalUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [GoalUpsertParams]'s query parameters as `url.Values`.
func (r GoalUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GoalValidateParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A goal payload for upsert or validate.
	Goal GoalRequestParam `json:"goal,omitzero" api:"required"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	paramObj
}

func (r GoalValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow GoalValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoalValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [GoalValidateParams]'s query parameters as `url.Values`.
func (r GoalValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
