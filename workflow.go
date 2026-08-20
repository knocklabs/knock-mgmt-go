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

// Workflows let you express your cross-channel notification logic.
//
// WorkflowService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	Options []option.RequestOption
	// Workflows let you express your cross-channel notification logic.
	Steps WorkflowStepService
}

// NewWorkflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowService(opts ...option.RequestOption) (r WorkflowService) {
	r = WorkflowService{}
	r.Options = opts
	r.Steps = NewWorkflowStepService(opts...)
	return
}

// Retrieve a workflow by its key in a given environment.
func (r *WorkflowService) Get(ctx context.Context, workflowKey string, query WorkflowGetParams, opts ...option.RequestOption) (res *WorkflowGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workflowKey == "" {
		err = errors.New("missing required workflow_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workflows/%s", workflowKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of workflows available in a given environment. The
// workflows are returned alphabetically by `key`.
func (r *WorkflowService) List(ctx context.Context, query WorkflowListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Workflow], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/workflows"
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

// Returns a paginated list of workflows available in a given environment. The
// workflows are returned alphabetically by `key`.
func (r *WorkflowService) ListAutoPaging(ctx context.Context, query WorkflowListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Workflow] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Activates (or deactivates) a workflow in a given environment. Read more in the
// [docs](https://docs.knock.app/concepts/workflows#workflow-status).
//
// Note: This immediately enables or disables a workflow in a given environment
// without needing to go through environment promotion.
func (r *WorkflowService) Activate(ctx context.Context, workflowKey string, params WorkflowActivateParams, opts ...option.RequestOption) (res *WorkflowActivateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workflowKey == "" {
		err = errors.New("missing required workflow_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workflows/%s/activate", workflowKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Runs the latest version of a committed workflow in a given environment using the
// params provided.
func (r *WorkflowService) Run(ctx context.Context, workflowKey string, params WorkflowRunParams, opts ...option.RequestOption) (res *WorkflowRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workflowKey == "" {
		err = errors.New("missing required workflow_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workflows/%s/run", workflowKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Updates a workflow of a given key, or creates a new one if it does not yet
// exist.
//
// Note: this endpoint only operates on workflows in the `development` environment.
func (r *WorkflowService) Upsert(ctx context.Context, workflowKey string, params WorkflowUpsertParams, opts ...option.RequestOption) (res *WorkflowUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workflowKey == "" {
		err = errors.New("missing required workflow_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workflows/%s", workflowKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Validates a workflow payload without persisting it. Some read-only fields will
// be empty as they are generated by the system when persisted.
//
// Note: Validating a workflow is only done in the development environment context.
func (r *WorkflowService) Validate(ctx context.Context, workflowKey string, params WorkflowValidateParams, opts ...option.RequestOption) (res *WorkflowValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workflowKey == "" {
		err = errors.New("missing required workflow_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workflows/%s/validate", workflowKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A condition to be evaluated.
type Condition struct {
	// The operator to use in the evaluation of the condition.
	//
	// Any of "equal_to", "not_equal_to", "greater_than", "less_than",
	// "greater_than_or_equal_to", "less_than_or_equal_to", "contains", "not_contains",
	// "contains_all", "not_contains_all", "is_timestamp_before",
	// "is_timestamp_on_or_after", "is_timestamp_between", "is_between", "empty",
	// "not_empty", "exists", "not_exists", "is_timestamp", "is_timestamp_before_now",
	// "is_timestamp_on_or_after_now", "is_audience_member", "is_not_audience_member".
	Operator ConditionOperator `json:"operator" api:"required"`
	// The variable to be evaluated. Variables can be either static values or dynamic
	// properties. Static values will always be JSON decoded so will support strings,
	// lists, objects, numbers, and booleans. Dynamic values should be path
	// expressions.
	Variable string `json:"variable" api:"required"`
	// The argument to be evaluated. Arguments can be either static values or dynamic
	// properties. Static values will always be JSON decoded so will support strings,
	// lists, objects, numbers, and booleans. Dynamic values should be path
	// expressions.
	Argument string `json:"argument" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Variable    respjson.Field
		Argument    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Condition) RawJSON() string { return r.JSON.raw }
func (r *Condition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Condition to a ConditionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConditionParam.Overrides()
func (r Condition) ToParam() ConditionParam {
	return param.Override[ConditionParam](json.RawMessage(r.RawJSON()))
}

// The operator to use in the evaluation of the condition.
type ConditionOperator string

const (
	ConditionOperatorEqualTo                 ConditionOperator = "equal_to"
	ConditionOperatorNotEqualTo              ConditionOperator = "not_equal_to"
	ConditionOperatorGreaterThan             ConditionOperator = "greater_than"
	ConditionOperatorLessThan                ConditionOperator = "less_than"
	ConditionOperatorGreaterThanOrEqualTo    ConditionOperator = "greater_than_or_equal_to"
	ConditionOperatorLessThanOrEqualTo       ConditionOperator = "less_than_or_equal_to"
	ConditionOperatorContains                ConditionOperator = "contains"
	ConditionOperatorNotContains             ConditionOperator = "not_contains"
	ConditionOperatorContainsAll             ConditionOperator = "contains_all"
	ConditionOperatorNotContainsAll          ConditionOperator = "not_contains_all"
	ConditionOperatorIsTimestampBefore       ConditionOperator = "is_timestamp_before"
	ConditionOperatorIsTimestampOnOrAfter    ConditionOperator = "is_timestamp_on_or_after"
	ConditionOperatorIsTimestampBetween      ConditionOperator = "is_timestamp_between"
	ConditionOperatorIsBetween               ConditionOperator = "is_between"
	ConditionOperatorEmpty                   ConditionOperator = "empty"
	ConditionOperatorNotEmpty                ConditionOperator = "not_empty"
	ConditionOperatorExists                  ConditionOperator = "exists"
	ConditionOperatorNotExists               ConditionOperator = "not_exists"
	ConditionOperatorIsTimestamp             ConditionOperator = "is_timestamp"
	ConditionOperatorIsTimestampBeforeNow    ConditionOperator = "is_timestamp_before_now"
	ConditionOperatorIsTimestampOnOrAfterNow ConditionOperator = "is_timestamp_on_or_after_now"
	ConditionOperatorIsAudienceMember        ConditionOperator = "is_audience_member"
	ConditionOperatorIsNotAudienceMember     ConditionOperator = "is_not_audience_member"
)

// A condition to be evaluated.
//
// The properties Operator, Variable are required.
type ConditionParam struct {
	// The operator to use in the evaluation of the condition.
	//
	// Any of "equal_to", "not_equal_to", "greater_than", "less_than",
	// "greater_than_or_equal_to", "less_than_or_equal_to", "contains", "not_contains",
	// "contains_all", "not_contains_all", "is_timestamp_before",
	// "is_timestamp_on_or_after", "is_timestamp_between", "is_between", "empty",
	// "not_empty", "exists", "not_exists", "is_timestamp", "is_timestamp_before_now",
	// "is_timestamp_on_or_after_now", "is_audience_member", "is_not_audience_member".
	Operator ConditionOperator `json:"operator,omitzero" api:"required"`
	// The variable to be evaluated. Variables can be either static values or dynamic
	// properties. Static values will always be JSON decoded so will support strings,
	// lists, objects, numbers, and booleans. Dynamic values should be path
	// expressions.
	Variable string `json:"variable" api:"required"`
	// The argument to be evaluated. Arguments can be either static values or dynamic
	// properties. Static values will always be JSON decoded so will support strings,
	// lists, objects, numbers, and booleans. Dynamic values should be path
	// expressions.
	Argument param.Opt[string] `json:"argument,omitzero"`
	paramObj
}

func (r ConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow ConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConditionGroupUnion contains all possible properties and values from
// [ConditionGroupConditionGroupAllMatch], [ConditionGroupConditionGroupAnyMatch].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConditionGroupUnion struct {
	// This field is from variant [ConditionGroupConditionGroupAllMatch].
	All []Condition `json:"all"`
	// This field is from variant [ConditionGroupConditionGroupAnyMatch].
	Any  []ConditionGroupConditionGroupAnyMatchAnyUnion `json:"any"`
	JSON struct {
		All respjson.Field
		Any respjson.Field
		raw string
	} `json:"-"`
}

func (u ConditionGroupUnion) AsConditionGroupAllMatch() (v ConditionGroupConditionGroupAllMatch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConditionGroupUnion) AsConditionGroupAnyMatch() (v ConditionGroupConditionGroupAnyMatch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConditionGroupUnion) RawJSON() string { return u.JSON.raw }

func (r *ConditionGroupUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConditionGroupUnion to a ConditionGroupUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConditionGroupUnionParam.Overrides()
func (r ConditionGroupUnion) ToParam() ConditionGroupUnionParam {
	return param.Override[ConditionGroupUnionParam](json.RawMessage(r.RawJSON()))
}

// A group of conditions that must all be met.
type ConditionGroupConditionGroupAllMatch struct {
	// A list of conditions.
	All []Condition `json:"all"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		All         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConditionGroupConditionGroupAllMatch) RawJSON() string { return r.JSON.raw }
func (r *ConditionGroupConditionGroupAllMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A group of conditions that any must be met. Can contain nested alls.
type ConditionGroupConditionGroupAnyMatch struct {
	// An array of conditions or nested condition groups to evaluate.
	Any []ConditionGroupConditionGroupAnyMatchAnyUnion `json:"any"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Any         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConditionGroupConditionGroupAnyMatch) RawJSON() string { return r.JSON.raw }
func (r *ConditionGroupConditionGroupAnyMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConditionGroupConditionGroupAnyMatchAnyUnion contains all possible properties
// and values from [Condition],
// [ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatch].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConditionGroupConditionGroupAnyMatchAnyUnion struct {
	// This field is from variant [Condition].
	Operator ConditionOperator `json:"operator"`
	// This field is from variant [Condition].
	Variable string `json:"variable"`
	// This field is from variant [Condition].
	Argument string `json:"argument"`
	// This field is from variant
	// [ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatch].
	All  []Condition `json:"all"`
	JSON struct {
		Operator respjson.Field
		Variable respjson.Field
		Argument respjson.Field
		All      respjson.Field
		raw      string
	} `json:"-"`
}

func (u ConditionGroupConditionGroupAnyMatchAnyUnion) AsCondition() (v Condition) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConditionGroupConditionGroupAnyMatchAnyUnion) AsConditionGroupAllMatch() (v ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConditionGroupConditionGroupAnyMatchAnyUnion) RawJSON() string { return u.JSON.raw }

func (r *ConditionGroupConditionGroupAnyMatchAnyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A group of conditions that must all be met.
type ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatch struct {
	// A list of conditions.
	All []Condition `json:"all"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		All         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatch) RawJSON() string {
	return r.JSON.raw
}
func (r *ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConditionGroupUnionParam struct {
	OfConditionGroupAllMatch *ConditionGroupConditionGroupAllMatchParam `json:",omitzero,inline"`
	OfConditionGroupAnyMatch *ConditionGroupConditionGroupAnyMatchParam `json:",omitzero,inline"`
	paramUnion
}

func (u ConditionGroupUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfConditionGroupAllMatch, u.OfConditionGroupAnyMatch)
}
func (u *ConditionGroupUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConditionGroupUnionParam) asAny() any {
	if !param.IsOmitted(u.OfConditionGroupAllMatch) {
		return u.OfConditionGroupAllMatch
	} else if !param.IsOmitted(u.OfConditionGroupAnyMatch) {
		return u.OfConditionGroupAnyMatch
	}
	return nil
}

// A group of conditions that must all be met.
type ConditionGroupConditionGroupAllMatchParam struct {
	// A list of conditions.
	All []ConditionParam `json:"all,omitzero"`
	paramObj
}

func (r ConditionGroupConditionGroupAllMatchParam) MarshalJSON() (data []byte, err error) {
	type shadow ConditionGroupConditionGroupAllMatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConditionGroupConditionGroupAllMatchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A group of conditions that any must be met. Can contain nested alls.
type ConditionGroupConditionGroupAnyMatchParam struct {
	// An array of conditions or nested condition groups to evaluate.
	Any []ConditionGroupConditionGroupAnyMatchAnyUnionParam `json:"any,omitzero"`
	paramObj
}

func (r ConditionGroupConditionGroupAnyMatchParam) MarshalJSON() (data []byte, err error) {
	type shadow ConditionGroupConditionGroupAnyMatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConditionGroupConditionGroupAnyMatchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConditionGroupConditionGroupAnyMatchAnyUnionParam struct {
	OfCondition              *ConditionParam                                                     `json:",omitzero,inline"`
	OfConditionGroupAllMatch *ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatchParam `json:",omitzero,inline"`
	paramUnion
}

func (u ConditionGroupConditionGroupAnyMatchAnyUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCondition, u.OfConditionGroupAllMatch)
}
func (u *ConditionGroupConditionGroupAnyMatchAnyUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConditionGroupConditionGroupAnyMatchAnyUnionParam) asAny() any {
	if !param.IsOmitted(u.OfCondition) {
		return u.OfCondition
	} else if !param.IsOmitted(u.OfConditionGroupAllMatch) {
		return u.OfConditionGroupAllMatch
	}
	return nil
}

// A group of conditions that must all be met.
type ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatchParam struct {
	// A list of conditions.
	All []ConditionParam `json:"all,omitzero"`
	paramObj
}

func (r ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatchParam) MarshalJSON() (data []byte, err error) {
	type shadow ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A duration of time, represented as a unit and a value.
type Duration struct {
	// The unit of time.
	//
	// Any of "minutes", "hours", "days", "weeks", "months".
	Unit DurationUnit `json:"unit" api:"required"`
	// The value of the duration.
	Value int64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Duration) RawJSON() string { return r.JSON.raw }
func (r *Duration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Duration to a DurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DurationParam.Overrides()
func (r Duration) ToParam() DurationParam {
	return param.Override[DurationParam](json.RawMessage(r.RawJSON()))
}

// The unit of time.
type DurationUnit string

const (
	DurationUnitMinutes DurationUnit = "minutes"
	DurationUnitHours   DurationUnit = "hours"
	DurationUnitDays    DurationUnit = "days"
	DurationUnitWeeks   DurationUnit = "weeks"
	DurationUnitMonths  DurationUnit = "months"
)

// A duration of time, represented as a unit and a value.
//
// The properties Unit, Value are required.
type DurationParam struct {
	// The unit of time.
	//
	// Any of "minutes", "hours", "days", "weeks", "months".
	Unit DurationUnit `json:"unit,omitzero" api:"required"`
	// The value of the duration.
	Value int64 `json:"value" api:"required"`
	paramObj
}

func (r DurationParam) MarshalJSON() (data []byte, err error) {
	type shadow DurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A send window time for a notification. Describes a single day.
type SendWindow struct {
	// The day of the week.
	//
	// Any of "monday", "tuesday", "wednesday", "thursday", "friday", "saturday",
	// "sunday".
	Day SendWindowDay `json:"day" api:"required"`
	// The type of send window.
	//
	// Any of "send", "do_not_send".
	Type SendWindowType `json:"type" api:"required"`
	// The start time of the send window.
	From string `json:"from" api:"nullable"`
	// The end time of the send window.
	Until string `json:"until" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day         respjson.Field
		Type        respjson.Field
		From        respjson.Field
		Until       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SendWindow) RawJSON() string { return r.JSON.raw }
func (r *SendWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SendWindow to a SendWindowParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SendWindowParam.Overrides()
func (r SendWindow) ToParam() SendWindowParam {
	return param.Override[SendWindowParam](json.RawMessage(r.RawJSON()))
}

// The day of the week.
type SendWindowDay string

const (
	SendWindowDayMonday    SendWindowDay = "monday"
	SendWindowDayTuesday   SendWindowDay = "tuesday"
	SendWindowDayWednesday SendWindowDay = "wednesday"
	SendWindowDayThursday  SendWindowDay = "thursday"
	SendWindowDayFriday    SendWindowDay = "friday"
	SendWindowDaySaturday  SendWindowDay = "saturday"
	SendWindowDaySunday    SendWindowDay = "sunday"
)

// The type of send window.
type SendWindowType string

const (
	SendWindowTypeSend      SendWindowType = "send"
	SendWindowTypeDoNotSend SendWindowType = "do_not_send"
)

// A send window time for a notification. Describes a single day.
//
// The properties Day, Type are required.
type SendWindowParam struct {
	// The day of the week.
	//
	// Any of "monday", "tuesday", "wednesday", "thursday", "friday", "saturday",
	// "sunday".
	Day SendWindowDay `json:"day,omitzero" api:"required"`
	// The type of send window.
	//
	// Any of "send", "do_not_send".
	Type SendWindowType `json:"type,omitzero" api:"required"`
	// The start time of the send window.
	From param.Opt[string] `json:"from,omitzero"`
	// The end time of the send window.
	Until param.Opt[string] `json:"until,omitzero"`
	paramObj
}

func (r SendWindowParam) MarshalJSON() (data []byte, err error) {
	type shadow SendWindowParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendWindowParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A workflow object. Read more in the
// [docs](https://docs.knock.app/concepts/workflows).
type Workflow struct {
	// Whether the workflow is
	// [active](https://docs.knock.app/concepts/workflows#workflow-status) in the
	// current environment. (read-only).
	Active bool `json:"active" api:"required"`
	// The timestamp of when the workflow was created. (read-only).
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The slug of the environment in which the workflow exists. (read-only).
	Environment string `json:"environment" api:"required"`
	// The unique key string for the workflow object. Must be at minimum 3 characters
	// and at maximum 255 characters in length. Must be in the format of ^[a-z0-9_-]+$.
	Key string `json:"key" api:"required"`
	// A name for the workflow. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// The SHA hash of the workflow data. (read-only).
	Sha string `json:"sha" api:"required"`
	// A list of workflow step objects in the workflow.
	Steps []WorkflowStepUnion `json:"steps" api:"required"`
	// The timestamp of when the workflow was last updated. (read-only).
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Whether the workflow and its steps are in a valid state. (read-only).
	Valid bool `json:"valid" api:"required"`
	// A list of
	// [categories](https://docs.knock.app/concepts/workflows#workflow-categories) that
	// the workflow belongs to.
	Categories []string `json:"categories"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// The timestamp of when the workflow was deleted. (read-only).
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// An arbitrary string attached to a workflow object. Useful for adding notes about
	// the workflow for internal purposes. Maximum of 280 characters allowed.
	Description string `json:"description"`
	// Attaches a goal to a workflow, guide, or broadcast for attribution tracking.
	GoalAttachment WorkflowGoalAttachment `json:"goal_attachment" api:"nullable"`
	// A map of workflow settings.
	Settings WorkflowSettings `json:"settings"`
	// Use tags to organize resources internally within your account. For example, by
	// team or product area.
	Tags []string `json:"tags"`
	// A JSON schema for the expected structure of the workflow trigger's `data`
	// payload (available in templates as `{{ data.field_name }}`). Used to validate
	// trigger requests. Read more in the
	// [docs](https://docs.knock.app/developer-tools/validating-trigger-data).
	TriggerDataJsonSchema map[string]any `json:"trigger_data_json_schema"`
	// The frequency at which the workflow should be triggered. One of:
	// `once_per_recipient`, `once_per_recipient_per_tenant`, `every_trigger`. Defaults
	// to `every_trigger`. Read more in
	// [docs](https://docs.knock.app/send-notifications/triggering-workflows/overview#controlling-workflow-trigger-frequency).
	//
	// Any of "every_trigger", "once_per_recipient", "once_per_recipient_per_tenant".
	TriggerFrequency WorkflowTriggerFrequency `json:"trigger_frequency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active                respjson.Field
		CreatedAt             respjson.Field
		Environment           respjson.Field
		Key                   respjson.Field
		Name                  respjson.Field
		Sha                   respjson.Field
		Steps                 respjson.Field
		UpdatedAt             respjson.Field
		Valid                 respjson.Field
		Categories            respjson.Field
		Conditions            respjson.Field
		DeletedAt             respjson.Field
		Description           respjson.Field
		GoalAttachment        respjson.Field
		Settings              respjson.Field
		Tags                  respjson.Field
		TriggerDataJsonSchema respjson.Field
		TriggerFrequency      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Workflow) RawJSON() string { return r.JSON.raw }
func (r *Workflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attaches a goal to a workflow, guide, or broadcast for attribution tracking.
type WorkflowGoalAttachment struct {
	// The key of the goal to attach.
	GoalKey string `json:"goal_key" api:"required"`
	// The number of days to attribute conversions after the notification is sent. Must
	// be between 1 and 30. Defaults to 7.
	AttributionWindowDays int64 `json:"attribution_window_days"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GoalKey               respjson.Field
		AttributionWindowDays respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGoalAttachment) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGoalAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A map of workflow settings.
type WorkflowSettings struct {
	// Whether the workflow is commercial. Defaults to false.
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
func (r WorkflowSettings) RawJSON() string { return r.JSON.raw }
func (r *WorkflowSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which the workflow should be triggered. One of:
// `once_per_recipient`, `once_per_recipient_per_tenant`, `every_trigger`. Defaults
// to `every_trigger`. Read more in
// [docs](https://docs.knock.app/send-notifications/triggering-workflows/overview#controlling-workflow-trigger-frequency).
type WorkflowTriggerFrequency string

const (
	WorkflowTriggerFrequencyEveryTrigger              WorkflowTriggerFrequency = "every_trigger"
	WorkflowTriggerFrequencyOncePerRecipient          WorkflowTriggerFrequency = "once_per_recipient"
	WorkflowTriggerFrequencyOncePerRecipientPerTenant WorkflowTriggerFrequency = "once_per_recipient_per_tenant"
)

// An AI agent function step. Fetches data from an AI model and merges it into the
// workflow's `data` scope for use in later steps. Supports Liquid templating in
// the prompt. Read more in the
// [docs](https://docs.knock.app/designing-workflows/ai-agent-function).
type WorkflowAIAgentStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the AI agent step.
	Settings WorkflowAIAgentStepSettings `json:"settings" api:"required"`
	// The type of the workflow step.
	//
	// Any of "ai_agent".
	Type WorkflowAIAgentStepType `json:"type" api:"required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref         respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		Conditions  respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowAIAgentStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowAIAgentStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowAIAgentStep to a WorkflowAIAgentStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowAIAgentStepParam.Overrides()
func (r WorkflowAIAgentStep) ToParam() WorkflowAIAgentStepParam {
	return param.Override[WorkflowAIAgentStepParam](json.RawMessage(r.RawJSON()))
}

// The settings for the AI agent step.
type WorkflowAIAgentStepSettings struct {
	// The AI model to use in `provider:model` format (e.g.
	// `anthropic:claude-haiku-4-5`, `openai:gpt-5.2-chat-latest`). See the
	// documentation for a list of supported models.
	Model string `json:"model" api:"required"`
	// The prompt template for the AI request. Supports Liquid templating.
	RequestPrompt string `json:"request_prompt" api:"required"`
	// The type of response to expect from the AI model.
	//
	// Any of "text", "json".
	ResponseType string `json:"response_type" api:"required"`
	// Whether to halt the workflow if the AI fetch fails.
	HaltOnError bool `json:"halt_on_error" api:"nullable"`
	// A JSON schema string for structured output. Required when `response_type` is
	// `json`. Must not be set when `response_type` is `text`.
	ResponseSchema string `json:"response_schema" api:"nullable"`
	// Whether to enable web search for the AI request.
	WebSearchEnabled bool `json:"web_search_enabled" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Model            respjson.Field
		RequestPrompt    respjson.Field
		ResponseType     respjson.Field
		HaltOnError      respjson.Field
		ResponseSchema   respjson.Field
		WebSearchEnabled respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowAIAgentStepSettings) RawJSON() string { return r.JSON.raw }
func (r *WorkflowAIAgentStepSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the workflow step.
type WorkflowAIAgentStepType string

const (
	WorkflowAIAgentStepTypeAIAgent WorkflowAIAgentStepType = "ai_agent"
)

// An AI agent function step. Fetches data from an AI model and merges it into the
// workflow's `data` scope for use in later steps. Supports Liquid templating in
// the prompt. Read more in the
// [docs](https://docs.knock.app/designing-workflows/ai-agent-function).
//
// The properties Ref, Settings, Type are required.
type WorkflowAIAgentStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the AI agent step.
	Settings WorkflowAIAgentStepSettingsParam `json:"settings,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "ai_agent".
	Type WorkflowAIAgentStepType `json:"type,omitzero" api:"required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowAIAgentStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowAIAgentStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowAIAgentStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The settings for the AI agent step.
//
// The properties Model, RequestPrompt, ResponseType are required.
type WorkflowAIAgentStepSettingsParam struct {
	// The AI model to use in `provider:model` format (e.g.
	// `anthropic:claude-haiku-4-5`, `openai:gpt-5.2-chat-latest`). See the
	// documentation for a list of supported models.
	Model string `json:"model" api:"required"`
	// The prompt template for the AI request. Supports Liquid templating.
	RequestPrompt string `json:"request_prompt" api:"required"`
	// The type of response to expect from the AI model.
	//
	// Any of "text", "json".
	ResponseType string `json:"response_type,omitzero" api:"required"`
	// Whether to halt the workflow if the AI fetch fails.
	HaltOnError param.Opt[bool] `json:"halt_on_error,omitzero"`
	// A JSON schema string for structured output. Required when `response_type` is
	// `json`. Must not be set when `response_type` is `text`.
	ResponseSchema param.Opt[string] `json:"response_schema,omitzero"`
	// Whether to enable web search for the AI request.
	WebSearchEnabled param.Opt[bool] `json:"web_search_enabled,omitzero"`
	paramObj
}

func (r WorkflowAIAgentStepSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowAIAgentStepSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowAIAgentStepSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowAIAgentStepSettingsParam](
		"response_type", "text", "json",
	)
}

// A batch function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/batch-function).
type WorkflowBatchStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the batch step.
	Settings WorkflowBatchStepSettings `json:"settings" api:"required"`
	// The type of the workflow step.
	//
	// Any of "batch".
	Type WorkflowBatchStepType `json:"type" api:"required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref         respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowBatchStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowBatchStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowBatchStep to a WorkflowBatchStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowBatchStepParam.Overrides()
func (r WorkflowBatchStep) ToParam() WorkflowBatchStepParam {
	return param.Override[WorkflowBatchStepParam](json.RawMessage(r.RawJSON()))
}

// The settings for the batch step.
type WorkflowBatchStepSettings struct {
	// The execution mode of the batch step. One of: `accumulate` or `flush_leading`.
	// When set to `flush_leading`, the first item in the batch will be executed
	// immediately, and the rest will be batched. See
	// [these docs](https://docs.knock.app/designing-workflows/batch-function#immediately-flushing-the-first-item-in-a-batch)
	// for more information.
	//
	// Any of "accumulate", "flush_leading".
	BatchExecutionMode string `json:"batch_execution_mode" api:"nullable"`
	// The maximum number of batch items allowed in a batch. Between: 2 and 1000.
	BatchItemsMaxLimit int64 `json:"batch_items_max_limit" api:"nullable"`
	// The maximum number of batch items allowed to be rendered into a template.
	// Between: 1 and 100. Defaults to 10.
	BatchItemsRenderLimit int64 `json:"batch_items_render_limit" api:"nullable"`
	// The data property to use to batch notifications per recipient.
	BatchKey string `json:"batch_key" api:"nullable"`
	// The order describing whether to return the first or last ten batch items in the
	// activities variable. One of: `asc` or `desc`.
	//
	// Any of "asc", "desc".
	BatchOrder string `json:"batch_order" api:"nullable"`
	// The data path to resolve the batch window. The resolved value must be an
	// ISO-8601 timestamp.
	BatchUntilFieldPath string `json:"batch_until_field_path" api:"nullable"`
	// A duration of time, represented as a unit and a value.
	BatchWindow Duration `json:"batch_window" api:"nullable"`
	// A duration of time, represented as a unit and a value.
	BatchWindowExtensionLimit Duration `json:"batch_window_extension_limit" api:"nullable"`
	// The type of the batch window used. One of: `fixed` or `sliding`.
	//
	// Any of "fixed", "sliding".
	BatchWindowType string `json:"batch_window_type" api:"nullable"`
	// Whether the batch is pinned to the opening workflow version or continues on the
	// latest compatible version. One of: `pinned` or `latest`. New batch steps default
	// to `latest`. Configs that omit the field hydrate as `pinned`. When set to
	// `latest`, compatible triggers share a cross-version batch and resume on the
	// latest published workflow after close.
	//
	// Any of "pinned", "latest".
	WorkflowVersionMode string `json:"workflow_version_mode" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchExecutionMode        respjson.Field
		BatchItemsMaxLimit        respjson.Field
		BatchItemsRenderLimit     respjson.Field
		BatchKey                  respjson.Field
		BatchOrder                respjson.Field
		BatchUntilFieldPath       respjson.Field
		BatchWindow               respjson.Field
		BatchWindowExtensionLimit respjson.Field
		BatchWindowType           respjson.Field
		WorkflowVersionMode       respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowBatchStepSettings) RawJSON() string { return r.JSON.raw }
func (r *WorkflowBatchStepSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the workflow step.
type WorkflowBatchStepType string

const (
	WorkflowBatchStepTypeBatch WorkflowBatchStepType = "batch"
)

// A batch function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/batch-function).
//
// The properties Ref, Settings, Type are required.
type WorkflowBatchStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the batch step.
	Settings WorkflowBatchStepSettingsParam `json:"settings,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "batch".
	Type WorkflowBatchStepType `json:"type,omitzero" api:"required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r WorkflowBatchStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowBatchStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowBatchStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The settings for the batch step.
type WorkflowBatchStepSettingsParam struct {
	// The maximum number of batch items allowed in a batch. Between: 2 and 1000.
	BatchItemsMaxLimit param.Opt[int64] `json:"batch_items_max_limit,omitzero"`
	// The maximum number of batch items allowed to be rendered into a template.
	// Between: 1 and 100. Defaults to 10.
	BatchItemsRenderLimit param.Opt[int64] `json:"batch_items_render_limit,omitzero"`
	// The data property to use to batch notifications per recipient.
	BatchKey param.Opt[string] `json:"batch_key,omitzero"`
	// The data path to resolve the batch window. The resolved value must be an
	// ISO-8601 timestamp.
	BatchUntilFieldPath param.Opt[string] `json:"batch_until_field_path,omitzero"`
	// The execution mode of the batch step. One of: `accumulate` or `flush_leading`.
	// When set to `flush_leading`, the first item in the batch will be executed
	// immediately, and the rest will be batched. See
	// [these docs](https://docs.knock.app/designing-workflows/batch-function#immediately-flushing-the-first-item-in-a-batch)
	// for more information.
	//
	// Any of "accumulate", "flush_leading".
	BatchExecutionMode string `json:"batch_execution_mode,omitzero"`
	// The order describing whether to return the first or last ten batch items in the
	// activities variable. One of: `asc` or `desc`.
	//
	// Any of "asc", "desc".
	BatchOrder string `json:"batch_order,omitzero"`
	// The type of the batch window used. One of: `fixed` or `sliding`.
	//
	// Any of "fixed", "sliding".
	BatchWindowType string `json:"batch_window_type,omitzero"`
	// Whether the batch is pinned to the opening workflow version or continues on the
	// latest compatible version. One of: `pinned` or `latest`. New batch steps default
	// to `latest`. Configs that omit the field hydrate as `pinned`. When set to
	// `latest`, compatible triggers share a cross-version batch and resume on the
	// latest published workflow after close.
	//
	// Any of "pinned", "latest".
	WorkflowVersionMode string `json:"workflow_version_mode,omitzero"`
	// A duration of time, represented as a unit and a value.
	BatchWindow DurationParam `json:"batch_window,omitzero"`
	// A duration of time, represented as a unit and a value.
	BatchWindowExtensionLimit DurationParam `json:"batch_window_extension_limit,omitzero"`
	paramObj
}

func (r WorkflowBatchStepSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowBatchStepSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowBatchStepSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowBatchStepSettingsParam](
		"batch_execution_mode", "accumulate", "flush_leading",
	)
	apijson.RegisterFieldValidator[WorkflowBatchStepSettingsParam](
		"batch_order", "asc", "desc",
	)
	apijson.RegisterFieldValidator[WorkflowBatchStepSettingsParam](
		"batch_window_type", "fixed", "sliding",
	)
	apijson.RegisterFieldValidator[WorkflowBatchStepSettingsParam](
		"workflow_version_mode", "pinned", "latest",
	)
}

// A branch function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/branch-function).
type WorkflowBranchStep struct {
	// A list of workflow branches to be evaluated.
	Branches []WorkflowBranchStepBranch `json:"branches" api:"required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The type of step.
	//
	// Any of "branch".
	Type WorkflowBranchStepType `json:"type" api:"required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Branches    respjson.Field
		Ref         respjson.Field
		Type        respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowBranchStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowBranchStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowBranchStep to a WorkflowBranchStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowBranchStepParam.Overrides()
func (r WorkflowBranchStep) ToParam() WorkflowBranchStepParam {
	return param.Override[WorkflowBranchStepParam](json.RawMessage(r.RawJSON()))
}

// A branch in a branch step.
type WorkflowBranchStepBranch struct {
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// The name of the branch.
	Name string `json:"name"`
	// A list of steps that will be executed if the branch is chosen.
	Steps []WorkflowStepUnion `json:"steps"`
	// If the workflow should halt at the end of the branch. Defaults to false if not
	// provided.
	Terminates bool `json:"terminates"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conditions  respjson.Field
		Name        respjson.Field
		Steps       respjson.Field
		Terminates  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowBranchStepBranch) RawJSON() string { return r.JSON.raw }
func (r *WorkflowBranchStepBranch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of step.
type WorkflowBranchStepType string

const (
	WorkflowBranchStepTypeBranch WorkflowBranchStepType = "branch"
)

// A branch function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/branch-function).
//
// The properties Branches, Ref, Type are required.
type WorkflowBranchStepParam struct {
	// A list of workflow branches to be evaluated.
	Branches []WorkflowBranchStepBranchParam `json:"branches,omitzero" api:"required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The type of step.
	//
	// Any of "branch".
	Type WorkflowBranchStepType `json:"type,omitzero" api:"required"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r WorkflowBranchStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowBranchStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowBranchStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A branch in a branch step.
type WorkflowBranchStepBranchParam struct {
	// The name of the branch.
	Name param.Opt[string] `json:"name,omitzero"`
	// If the workflow should halt at the end of the branch. Defaults to false if not
	// provided.
	Terminates param.Opt[bool] `json:"terminates,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	// A list of steps that will be executed if the branch is chosen.
	Steps []WorkflowStepUnionParam `json:"steps,omitzero"`
	paramObj
}

func (r WorkflowBranchStepBranchParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowBranchStepBranchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowBranchStepBranchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A chat step within a workflow. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
type WorkflowChatStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// A chat template.
	Template ChatTemplate `json:"template" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowChatStepType `json:"type" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey string `json:"channel_group_key" api:"nullable"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey string `json:"channel_key" api:"nullable"`
	// Chat channel settings. Only used as configuration as part of a workflow channel
	// step.
	ChannelOverrides ChatChannelSettings `json:"channel_overrides" api:"nullable"`
	// The type of the channel step. Always `chat` for chat steps.
	//
	// Any of "chat".
	ChannelType WorkflowChatStepChannelType `json:"channel_type"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindow `json:"send_windows" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref              respjson.Field
		Template         respjson.Field
		Type             respjson.Field
		ChannelGroupKey  respjson.Field
		ChannelKey       respjson.Field
		ChannelOverrides respjson.Field
		ChannelType      respjson.Field
		Conditions       respjson.Field
		Description      respjson.Field
		Name             respjson.Field
		SendWindows      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowChatStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowChatStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowChatStep to a WorkflowChatStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowChatStepParam.Overrides()
func (r WorkflowChatStep) ToParam() WorkflowChatStepParam {
	return param.Override[WorkflowChatStepParam](json.RawMessage(r.RawJSON()))
}

// The type of the workflow step.
type WorkflowChatStepType string

const (
	WorkflowChatStepTypeChannel WorkflowChatStepType = "channel"
)

// The type of the channel step. Always `chat` for chat steps.
type WorkflowChatStepChannelType string

const (
	WorkflowChatStepChannelTypeChat WorkflowChatStepChannelType = "chat"
)

// A chat step within a workflow. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
//
// The properties Ref, Template, Type are required.
type WorkflowChatStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// A chat template.
	Template ChatTemplateParam `json:"template,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowChatStepType `json:"type,omitzero" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey param.Opt[string] `json:"channel_group_key,omitzero"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey param.Opt[string] `json:"channel_key,omitzero"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindowParam `json:"send_windows,omitzero"`
	// Chat channel settings. Only used as configuration as part of a workflow channel
	// step.
	ChannelOverrides ChatChannelSettingsParam `json:"channel_overrides,omitzero"`
	// The type of the channel step. Always `chat` for chat steps.
	//
	// Any of "chat".
	ChannelType WorkflowChatStepChannelType `json:"channel_type,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowChatStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowChatStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowChatStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A delay function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/delay-function).
type WorkflowDelayStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the delay step. Both fields can be set to compute a delay where
	// `delay_for` is an offset from the `delay_until_field_path`.
	Settings WorkflowDelayStepSettings `json:"settings" api:"required"`
	// The type of the workflow step.
	//
	// Any of "delay".
	Type WorkflowDelayStepType `json:"type" api:"required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref         respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		Conditions  respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowDelayStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowDelayStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowDelayStep to a WorkflowDelayStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowDelayStepParam.Overrides()
func (r WorkflowDelayStep) ToParam() WorkflowDelayStepParam {
	return param.Override[WorkflowDelayStepParam](json.RawMessage(r.RawJSON()))
}

// The settings for the delay step. Both fields can be set to compute a delay where
// `delay_for` is an offset from the `delay_until_field_path`.
type WorkflowDelayStepSettings struct {
	// A duration of time, represented as a unit and a value.
	DelayFor Duration `json:"delay_for" api:"nullable"`
	// When set will use the path to resolve the delay into a timestamp from the
	// property referenced
	DelayUntilFieldPath string `json:"delay_until_field_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DelayFor            respjson.Field
		DelayUntilFieldPath respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowDelayStepSettings) RawJSON() string { return r.JSON.raw }
func (r *WorkflowDelayStepSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the workflow step.
type WorkflowDelayStepType string

const (
	WorkflowDelayStepTypeDelay WorkflowDelayStepType = "delay"
)

// A delay function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/delay-function).
//
// The properties Ref, Settings, Type are required.
type WorkflowDelayStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the delay step. Both fields can be set to compute a delay where
	// `delay_for` is an offset from the `delay_until_field_path`.
	Settings WorkflowDelayStepSettingsParam `json:"settings,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "delay".
	Type WorkflowDelayStepType `json:"type,omitzero" api:"required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowDelayStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowDelayStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowDelayStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The settings for the delay step. Both fields can be set to compute a delay where
// `delay_for` is an offset from the `delay_until_field_path`.
type WorkflowDelayStepSettingsParam struct {
	// When set will use the path to resolve the delay into a timestamp from the
	// property referenced
	DelayUntilFieldPath param.Opt[string] `json:"delay_until_field_path,omitzero"`
	// A duration of time, represented as a unit and a value.
	DelayFor DurationParam `json:"delay_for,omitzero"`
	paramObj
}

func (r WorkflowDelayStepSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowDelayStepSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowDelayStepSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An email step within a workflow. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
type WorkflowEmailStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// An email message template.
	Template EmailTemplate `json:"template" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowEmailStepType `json:"type" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey string `json:"channel_group_key" api:"nullable"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey string `json:"channel_key" api:"nullable"`
	// Email channel settings. Only used as configuration as part of a workflow channel
	// step.
	ChannelOverrides EmailChannelSettings `json:"channel_overrides" api:"nullable"`
	// The category of channel for this step. Always `email` for email steps. This
	// identifies the type of notification (email, sms, push, etc.) while `channel_key`
	// specifies which configured provider instance to use.
	//
	// Any of "email".
	ChannelType WorkflowEmailStepChannelType `json:"channel_type"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindow `json:"send_windows" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref              respjson.Field
		Template         respjson.Field
		Type             respjson.Field
		ChannelGroupKey  respjson.Field
		ChannelKey       respjson.Field
		ChannelOverrides respjson.Field
		ChannelType      respjson.Field
		Conditions       respjson.Field
		Description      respjson.Field
		Name             respjson.Field
		SendWindows      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEmailStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowEmailStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowEmailStep to a WorkflowEmailStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowEmailStepParam.Overrides()
func (r WorkflowEmailStep) ToParam() WorkflowEmailStepParam {
	return param.Override[WorkflowEmailStepParam](json.RawMessage(r.RawJSON()))
}

// The type of the workflow step.
type WorkflowEmailStepType string

const (
	WorkflowEmailStepTypeChannel WorkflowEmailStepType = "channel"
)

// The category of channel for this step. Always `email` for email steps. This
// identifies the type of notification (email, sms, push, etc.) while `channel_key`
// specifies which configured provider instance to use.
type WorkflowEmailStepChannelType string

const (
	WorkflowEmailStepChannelTypeEmail WorkflowEmailStepChannelType = "email"
)

// An email step within a workflow. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
//
// The properties Ref, Template, Type are required.
type WorkflowEmailStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// An email message template.
	Template EmailTemplateParam `json:"template,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowEmailStepType `json:"type,omitzero" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey param.Opt[string] `json:"channel_group_key,omitzero"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey param.Opt[string] `json:"channel_key,omitzero"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindowParam `json:"send_windows,omitzero"`
	// Email channel settings. Only used as configuration as part of a workflow channel
	// step.
	ChannelOverrides EmailChannelSettingsParam `json:"channel_overrides,omitzero"`
	// The category of channel for this step. Always `email` for email steps. This
	// identifies the type of notification (email, sms, push, etc.) while `channel_key`
	// specifies which configured provider instance to use.
	//
	// Any of "email".
	ChannelType WorkflowEmailStepChannelType `json:"channel_type,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowEmailStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowEmailStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowEmailStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fetch function step. Retrieves data from an external source and merges it into
// the workflow's `data` scope for use in later steps. Read more in the
// [docs](https://docs.knock.app/designing-workflows/fetch-function).
type WorkflowFetchStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// A request template for a fetch function step.
	Settings RequestTemplate `json:"settings" api:"required"`
	// The type of the workflow step.
	//
	// Any of "http_fetch".
	Type WorkflowFetchStepType `json:"type" api:"required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref         respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		Conditions  respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowFetchStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowFetchStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowFetchStep to a WorkflowFetchStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowFetchStepParam.Overrides()
func (r WorkflowFetchStep) ToParam() WorkflowFetchStepParam {
	return param.Override[WorkflowFetchStepParam](json.RawMessage(r.RawJSON()))
}

// The type of the workflow step.
type WorkflowFetchStepType string

const (
	WorkflowFetchStepTypeHTTPFetch WorkflowFetchStepType = "http_fetch"
)

// A fetch function step. Retrieves data from an external source and merges it into
// the workflow's `data` scope for use in later steps. Read more in the
// [docs](https://docs.knock.app/designing-workflows/fetch-function).
//
// The properties Ref, Settings, Type are required.
type WorkflowFetchStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// A request template for a fetch function step.
	Settings RequestTemplateParam `json:"settings,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "http_fetch".
	Type WorkflowFetchStepType `json:"type,omitzero" api:"required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowFetchStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFetchStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowFetchStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An in-app feed step within a workflow. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
type WorkflowInAppFeedStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// An in-app feed template.
	Template InAppFeedTemplate `json:"template" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowInAppFeedStepType `json:"type" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey string `json:"channel_group_key" api:"nullable"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey string `json:"channel_key" api:"nullable"`
	// In-app feed channel settings. Only used as configuration as part of a workflow
	// channel step.
	ChannelOverrides InAppFeedChannelSettings `json:"channel_overrides" api:"nullable"`
	// The type of the channel step. Always `in_app_feed` for in-app feed steps.
	//
	// Any of "in_app_feed".
	ChannelType WorkflowInAppFeedStepChannelType `json:"channel_type"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindow `json:"send_windows" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref              respjson.Field
		Template         respjson.Field
		Type             respjson.Field
		ChannelGroupKey  respjson.Field
		ChannelKey       respjson.Field
		ChannelOverrides respjson.Field
		ChannelType      respjson.Field
		Conditions       respjson.Field
		Description      respjson.Field
		Name             respjson.Field
		SendWindows      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowInAppFeedStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowInAppFeedStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowInAppFeedStep to a WorkflowInAppFeedStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowInAppFeedStepParam.Overrides()
func (r WorkflowInAppFeedStep) ToParam() WorkflowInAppFeedStepParam {
	return param.Override[WorkflowInAppFeedStepParam](json.RawMessage(r.RawJSON()))
}

// The type of the workflow step.
type WorkflowInAppFeedStepType string

const (
	WorkflowInAppFeedStepTypeChannel WorkflowInAppFeedStepType = "channel"
)

// The type of the channel step. Always `in_app_feed` for in-app feed steps.
type WorkflowInAppFeedStepChannelType string

const (
	WorkflowInAppFeedStepChannelTypeInAppFeed WorkflowInAppFeedStepChannelType = "in_app_feed"
)

// An in-app feed step within a workflow. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
//
// The properties Ref, Template, Type are required.
type WorkflowInAppFeedStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// An in-app feed template.
	Template InAppFeedTemplateParam `json:"template,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowInAppFeedStepType `json:"type,omitzero" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey param.Opt[string] `json:"channel_group_key,omitzero"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey param.Opt[string] `json:"channel_key,omitzero"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindowParam `json:"send_windows,omitzero"`
	// In-app feed channel settings. Only used as configuration as part of a workflow
	// channel step.
	ChannelOverrides InAppFeedChannelSettingsParam `json:"channel_overrides,omitzero"`
	// The type of the channel step. Always `in_app_feed` for in-app feed steps.
	//
	// Any of "in_app_feed".
	ChannelType WorkflowInAppFeedStepChannelType `json:"channel_type,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowInAppFeedStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowInAppFeedStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowInAppFeedStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A push step within a workflow. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
type WorkflowPushStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// A push notification template.
	Template PushTemplate `json:"template" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowPushStepType `json:"type" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey string `json:"channel_group_key" api:"nullable"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey string `json:"channel_key" api:"nullable"`
	// Push channel settings. Only used as configuration as part of a workflow channel
	// step.
	ChannelOverrides PushChannelSettings `json:"channel_overrides" api:"nullable"`
	// The type of the channel step. Always `push` for push steps.
	//
	// Any of "push".
	ChannelType WorkflowPushStepChannelType `json:"channel_type"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindow `json:"send_windows" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref              respjson.Field
		Template         respjson.Field
		Type             respjson.Field
		ChannelGroupKey  respjson.Field
		ChannelKey       respjson.Field
		ChannelOverrides respjson.Field
		ChannelType      respjson.Field
		Conditions       respjson.Field
		Description      respjson.Field
		Name             respjson.Field
		SendWindows      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowPushStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowPushStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowPushStep to a WorkflowPushStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowPushStepParam.Overrides()
func (r WorkflowPushStep) ToParam() WorkflowPushStepParam {
	return param.Override[WorkflowPushStepParam](json.RawMessage(r.RawJSON()))
}

// The type of the workflow step.
type WorkflowPushStepType string

const (
	WorkflowPushStepTypeChannel WorkflowPushStepType = "channel"
)

// The type of the channel step. Always `push` for push steps.
type WorkflowPushStepChannelType string

const (
	WorkflowPushStepChannelTypePush WorkflowPushStepChannelType = "push"
)

// A push step within a workflow. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
//
// The properties Ref, Template, Type are required.
type WorkflowPushStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// A push notification template.
	Template PushTemplateParam `json:"template,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowPushStepType `json:"type,omitzero" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey param.Opt[string] `json:"channel_group_key,omitzero"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey param.Opt[string] `json:"channel_key,omitzero"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindowParam `json:"send_windows,omitzero"`
	// Push channel settings. Only used as configuration as part of a workflow channel
	// step.
	ChannelOverrides PushChannelSettingsParam `json:"channel_overrides,omitzero"`
	// The type of the channel step. Always `push` for push steps.
	//
	// Any of "push".
	ChannelType WorkflowPushStepChannelType `json:"channel_type,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowPushStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowPushStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowPushStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An experiment step. Deterministically assigns recipients to percentage-based
// cohorts for A/B testing and experimentation.
type WorkflowRandomCohortStep struct {
	// A list of cohort branches. Must have between 2 and 10 branches, and percentages
	// must sum to 100.
	CohortBranches []any `json:"cohort_branches" api:"required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The type of step.
	//
	// Any of "random_cohort".
	Type WorkflowRandomCohortStepType `json:"type" api:"required"`
	// The key used to deterministically assign recipients to cohorts. Defaults to the
	// recipient ID if not provided.
	CohortKey string `json:"cohort_key" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CohortBranches respjson.Field
		Ref            respjson.Field
		Type           respjson.Field
		CohortKey      respjson.Field
		Description    respjson.Field
		Name           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRandomCohortStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRandomCohortStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowRandomCohortStep to a
// WorkflowRandomCohortStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowRandomCohortStepParam.Overrides()
func (r WorkflowRandomCohortStep) ToParam() WorkflowRandomCohortStepParam {
	return param.Override[WorkflowRandomCohortStepParam](json.RawMessage(r.RawJSON()))
}

// The type of step.
type WorkflowRandomCohortStepType string

const (
	WorkflowRandomCohortStepTypeRandomCohort WorkflowRandomCohortStepType = "random_cohort"
)

// An experiment step. Deterministically assigns recipients to percentage-based
// cohorts for A/B testing and experimentation.
//
// The properties CohortBranches, Ref, Type are required.
type WorkflowRandomCohortStepParam struct {
	// A list of cohort branches. Must have between 2 and 10 branches, and percentages
	// must sum to 100.
	CohortBranches []any `json:"cohort_branches,omitzero" api:"required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The type of step.
	//
	// Any of "random_cohort".
	Type WorkflowRandomCohortStepType `json:"type,omitzero" api:"required"`
	// The key used to deterministically assign recipients to cohorts. Defaults to the
	// recipient ID if not provided.
	CohortKey param.Opt[string] `json:"cohort_key,omitzero"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r WorkflowRandomCohortStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRandomCohortStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowRandomCohortStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A SMS step within a workflow. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
type WorkflowSMSStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// An SMS template.
	Template SMSTemplate `json:"template" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowSMSStepType `json:"type" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey string `json:"channel_group_key" api:"nullable"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey string `json:"channel_key" api:"nullable"`
	// SMS channel settings. Only used as configuration as part of a workflow channel
	// step.
	ChannelOverrides SMSChannelSettings `json:"channel_overrides" api:"nullable"`
	// The type of the channel step. Always `sms` for SMS steps.
	//
	// Any of "sms".
	ChannelType WorkflowSMSStepChannelType `json:"channel_type"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindow `json:"send_windows" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref              respjson.Field
		Template         respjson.Field
		Type             respjson.Field
		ChannelGroupKey  respjson.Field
		ChannelKey       respjson.Field
		ChannelOverrides respjson.Field
		ChannelType      respjson.Field
		Conditions       respjson.Field
		Description      respjson.Field
		Name             respjson.Field
		SendWindows      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowSMSStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowSMSStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowSMSStep to a WorkflowSMSStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowSMSStepParam.Overrides()
func (r WorkflowSMSStep) ToParam() WorkflowSMSStepParam {
	return param.Override[WorkflowSMSStepParam](json.RawMessage(r.RawJSON()))
}

// The type of the workflow step.
type WorkflowSMSStepType string

const (
	WorkflowSMSStepTypeChannel WorkflowSMSStepType = "channel"
)

// The type of the channel step. Always `sms` for SMS steps.
type WorkflowSMSStepChannelType string

const (
	WorkflowSMSStepChannelTypeSMS WorkflowSMSStepChannelType = "sms"
)

// A SMS step within a workflow. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
//
// The properties Ref, Template, Type are required.
type WorkflowSMSStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// An SMS template.
	Template SMSTemplateParam `json:"template,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowSMSStepType `json:"type,omitzero" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey param.Opt[string] `json:"channel_group_key,omitzero"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey param.Opt[string] `json:"channel_key,omitzero"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindowParam `json:"send_windows,omitzero"`
	// SMS channel settings. Only used as configuration as part of a workflow channel
	// step.
	ChannelOverrides SMSChannelSettingsParam `json:"channel_overrides,omitzero"`
	// The type of the channel step. Always `sms` for SMS steps.
	//
	// Any of "sms".
	ChannelType WorkflowSMSStepChannelType `json:"channel_type,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowSMSStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowSMSStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowSMSStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowStepUnion contains all possible properties and values from
// [WorkflowWebhookStep], [WorkflowInAppFeedStep],
// [WorkflowStepWorkflowInAppGuideStep], [WorkflowChatStep], [WorkflowSMSStep],
// [WorkflowPushStep], [WorkflowEmailStep], [WorkflowAIAgentStep],
// [WorkflowDelayStep], [WorkflowStepWorkflowWaitForEventStep],
// [WorkflowBatchStep], [WorkflowFetchStep], [WorkflowUpdateDataStep],
// [WorkflowUpdateObjectStep], [WorkflowUpdateTenantStep],
// [WorkflowUpdateUserStep], [WorkflowThrottleStep], [WorkflowBranchStep],
// [WorkflowRandomCohortStep], [WorkflowTriggerWorkflowStep].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkflowStepUnion struct {
	Ref string `json:"ref"`
	// This field is a union of [WebhookTemplate], [InAppFeedTemplate], [ChatTemplate],
	// [SMSTemplate], [PushTemplate], [EmailTemplate]
	Template        WorkflowStepUnionTemplate `json:"template"`
	Type            string                    `json:"type"`
	ChannelGroupKey string                    `json:"channel_group_key"`
	ChannelKey      string                    `json:"channel_key"`
	ChannelType     string                    `json:"channel_type"`
	// This field is from variant [WorkflowWebhookStep].
	Conditions  ConditionGroupUnion `json:"conditions"`
	Description string              `json:"description"`
	Name        string              `json:"name"`
	SendWindows []SendWindow        `json:"send_windows"`
	// This field is a union of [InAppFeedChannelSettings], [ChatChannelSettings],
	// [SMSChannelSettings], [PushChannelSettings], [EmailChannelSettings]
	ChannelOverrides WorkflowStepUnionChannelOverrides `json:"channel_overrides"`
	// This field is from variant [WorkflowStepWorkflowInAppGuideStep].
	GuideKey string `json:"guide_key"`
	// This field is a union of [WorkflowAIAgentStepSettings],
	// [WorkflowDelayStepSettings],
	// [WorkflowStepWorkflowWaitForEventStepSettingsUnion],
	// [WorkflowBatchStepSettings], [RequestTemplate],
	// [WorkflowUpdateDataStepSettings], [WorkflowUpdateObjectStepSettings],
	// [WorkflowUpdateTenantStepSettings], [WorkflowUpdateUserStepSettings],
	// [WorkflowThrottleStepSettings], [WorkflowTriggerWorkflowStepSettings]
	Settings WorkflowStepUnionSettings `json:"settings"`
	// This field is from variant [WorkflowBranchStep].
	Branches []WorkflowBranchStepBranch `json:"branches"`
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
		GuideKey         respjson.Field
		Settings         respjson.Field
		Branches         respjson.Field
		CohortBranches   respjson.Field
		CohortKey        respjson.Field
		raw              string
	} `json:"-"`
}

func (u WorkflowStepUnion) AsWorkflowWebhookStep() (v WorkflowWebhookStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowInAppFeedStep() (v WorkflowInAppFeedStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowInAppGuideStep() (v WorkflowStepWorkflowInAppGuideStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowChatStep() (v WorkflowChatStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowSMSStep() (v WorkflowSMSStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowPushStep() (v WorkflowPushStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowEmailStep() (v WorkflowEmailStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowAIAgentStep() (v WorkflowAIAgentStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowDelayStep() (v WorkflowDelayStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowWaitForEventStep() (v WorkflowStepWorkflowWaitForEventStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowBatchStep() (v WorkflowBatchStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowFetchStep() (v WorkflowFetchStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowUpdateDataStep() (v WorkflowUpdateDataStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowUpdateObjectStep() (v WorkflowUpdateObjectStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowUpdateTenantStep() (v WorkflowUpdateTenantStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowUpdateUserStep() (v WorkflowUpdateUserStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowThrottleStep() (v WorkflowThrottleStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowBranchStep() (v WorkflowBranchStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowRandomCohortStep() (v WorkflowRandomCohortStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowTriggerWorkflowStep() (v WorkflowTriggerWorkflowStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowStepUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkflowStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowStepUnionTemplate is an implicit subunion of [WorkflowStepUnion].
// WorkflowStepUnionTemplate provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowStepUnion].
type WorkflowStepUnionTemplate struct {
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
	Settings WorkflowStepUnionTemplateSettings `json:"settings"`
	// This field is from variant [PushTemplate].
	Title string `json:"title"`
	// This field is from variant [EmailTemplate].
	Subject string `json:"subject"`
	// This field is from variant [EmailTemplate].
	HTMLBody string `json:"html_body"`
	// This field is from variant [EmailTemplate].
	IsMjml bool `json:"is_mjml"`
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
		IsMjml        respjson.Field
		VisualBlocks  respjson.Field
		raw           string
	} `json:"-"`
}

func (r *WorkflowStepUnionTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowStepUnionTemplateSettings is an implicit subunion of
// [WorkflowStepUnion]. WorkflowStepUnionTemplateSettings provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowStepUnion].
type WorkflowStepUnionTemplateSettings struct {
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

func (r *WorkflowStepUnionTemplateSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowStepUnionChannelOverrides is an implicit subunion of
// [WorkflowStepUnion]. WorkflowStepUnionChannelOverrides provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowStepUnion].
type WorkflowStepUnionChannelOverrides struct {
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

func (r *WorkflowStepUnionChannelOverrides) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowStepUnionSettings is an implicit subunion of [WorkflowStepUnion].
// WorkflowStepUnionSettings provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowStepUnion].
type WorkflowStepUnionSettings struct {
	// This field is from variant [WorkflowAIAgentStepSettings].
	Model string `json:"model"`
	// This field is from variant [WorkflowAIAgentStepSettings].
	RequestPrompt string `json:"request_prompt"`
	// This field is from variant [WorkflowAIAgentStepSettings].
	ResponseType string `json:"response_type"`
	// This field is from variant [WorkflowAIAgentStepSettings].
	HaltOnError bool `json:"halt_on_error"`
	// This field is from variant [WorkflowAIAgentStepSettings].
	ResponseSchema string `json:"response_schema"`
	// This field is from variant [WorkflowAIAgentStepSettings].
	WebSearchEnabled bool `json:"web_search_enabled"`
	// This field is from variant [WorkflowDelayStepSettings].
	DelayFor Duration `json:"delay_for"`
	// This field is from variant [WorkflowDelayStepSettings].
	DelayUntilFieldPath string `json:"delay_until_field_path"`
	// This field is a union of
	// [WorkflowStepWorkflowWaitForEventStepSettingsObjectEvent],
	// [WorkflowStepWorkflowWaitForEventStepSettingsObject2Event],
	// [WorkflowStepWorkflowWaitForEventStepSettingsObject3Event],
	// [WorkflowStepWorkflowWaitForEventStepSettingsObject4Event],
	// [WorkflowStepWorkflowWaitForEventStepSettingsObject5Event]
	Event WorkflowStepUnionSettingsEvent `json:"event"`
	// This field is from variant [WorkflowStepWorkflowWaitForEventStepSettingsUnion].
	ExpiresAfter Duration `json:"expires_after"`
	// This field is a union of
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchCondition],
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchCondition],
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchCondition],
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchCondition],
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchCondition]
	MatchConditions WorkflowStepUnionSettingsMatchConditions `json:"match_conditions"`
	OnMatch         string                                   `json:"on_match"`
	OnTimeout       string                                   `json:"on_timeout"`
	// This field is from variant [WorkflowBatchStepSettings].
	BatchExecutionMode string `json:"batch_execution_mode"`
	// This field is from variant [WorkflowBatchStepSettings].
	BatchItemsMaxLimit int64 `json:"batch_items_max_limit"`
	// This field is from variant [WorkflowBatchStepSettings].
	BatchItemsRenderLimit int64 `json:"batch_items_render_limit"`
	// This field is from variant [WorkflowBatchStepSettings].
	BatchKey string `json:"batch_key"`
	// This field is from variant [WorkflowBatchStepSettings].
	BatchOrder string `json:"batch_order"`
	// This field is from variant [WorkflowBatchStepSettings].
	BatchUntilFieldPath string `json:"batch_until_field_path"`
	// This field is from variant [WorkflowBatchStepSettings].
	BatchWindow Duration `json:"batch_window"`
	// This field is from variant [WorkflowBatchStepSettings].
	BatchWindowExtensionLimit Duration `json:"batch_window_extension_limit"`
	// This field is from variant [WorkflowBatchStepSettings].
	BatchWindowType string `json:"batch_window_type"`
	// This field is from variant [WorkflowBatchStepSettings].
	WorkflowVersionMode string `json:"workflow_version_mode"`
	// This field is from variant [RequestTemplate].
	Method RequestTemplateMethod `json:"method"`
	// This field is from variant [RequestTemplate].
	URL string `json:"url"`
	// This field is from variant [RequestTemplate].
	Body string `json:"body"`
	// This field is from variant [RequestTemplate].
	Headers RequestTemplateHeadersUnion `json:"headers"`
	// This field is from variant [RequestTemplate].
	QueryParams      RequestTemplateQueryParamsUnion `json:"query_params"`
	Data             string                          `json:"data"`
	RecipientGid     string                          `json:"recipient_gid"`
	UpdateProperties string                          `json:"update_properties"`
	RecipientMode    string                          `json:"recipient_mode"`
	// This field is from variant [WorkflowThrottleStepSettings].
	ThrottleKey string `json:"throttle_key"`
	// This field is from variant [WorkflowThrottleStepSettings].
	ThrottleLimit int64 `json:"throttle_limit"`
	// This field is from variant [WorkflowThrottleStepSettings].
	ThrottleWindow Duration `json:"throttle_window"`
	// This field is from variant [WorkflowThrottleStepSettings].
	ThrottleWindowFieldPath string `json:"throttle_window_field_path"`
	// This field is from variant [WorkflowTriggerWorkflowStepSettings].
	Actor string `json:"actor"`
	// This field is from variant [WorkflowTriggerWorkflowStepSettings].
	CancellationKey string `json:"cancellation_key"`
	// This field is from variant [WorkflowTriggerWorkflowStepSettings].
	Recipients string `json:"recipients"`
	// This field is from variant [WorkflowTriggerWorkflowStepSettings].
	Tenant string `json:"tenant"`
	// This field is from variant [WorkflowTriggerWorkflowStepSettings].
	WorkflowKey string `json:"workflow_key"`
	JSON        struct {
		Model                     respjson.Field
		RequestPrompt             respjson.Field
		ResponseType              respjson.Field
		HaltOnError               respjson.Field
		ResponseSchema            respjson.Field
		WebSearchEnabled          respjson.Field
		DelayFor                  respjson.Field
		DelayUntilFieldPath       respjson.Field
		Event                     respjson.Field
		ExpiresAfter              respjson.Field
		MatchConditions           respjson.Field
		OnMatch                   respjson.Field
		OnTimeout                 respjson.Field
		BatchExecutionMode        respjson.Field
		BatchItemsMaxLimit        respjson.Field
		BatchItemsRenderLimit     respjson.Field
		BatchKey                  respjson.Field
		BatchOrder                respjson.Field
		BatchUntilFieldPath       respjson.Field
		BatchWindow               respjson.Field
		BatchWindowExtensionLimit respjson.Field
		BatchWindowType           respjson.Field
		WorkflowVersionMode       respjson.Field
		Method                    respjson.Field
		URL                       respjson.Field
		Body                      respjson.Field
		Headers                   respjson.Field
		QueryParams               respjson.Field
		Data                      respjson.Field
		RecipientGid              respjson.Field
		UpdateProperties          respjson.Field
		RecipientMode             respjson.Field
		ThrottleKey               respjson.Field
		ThrottleLimit             respjson.Field
		ThrottleWindow            respjson.Field
		ThrottleWindowFieldPath   respjson.Field
		Actor                     respjson.Field
		CancellationKey           respjson.Field
		Recipients                respjson.Field
		Tenant                    respjson.Field
		WorkflowKey               respjson.Field
		raw                       string
	} `json:"-"`
}

func (r *WorkflowStepUnionSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowStepUnionSettingsEvent is an implicit subunion of [WorkflowStepUnion].
// WorkflowStepUnionSettingsEvent provides convenient access to the sub-properties
// of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowStepUnion].
type WorkflowStepUnionSettingsEvent struct {
	EventKey  string `json:"event_key"`
	EventType string `json:"event_type"`
	// This field is from variant
	// [WorkflowStepWorkflowWaitForEventStepSettingsObjectEvent].
	IntegrationSourceKey string `json:"integration_source_key"`
	SourceKey            string `json:"source_key"`
	// This field is from variant
	// [WorkflowStepWorkflowWaitForEventStepSettingsObject2Event].
	SourceType string `json:"source_type"`
	// This field is from variant
	// [WorkflowStepWorkflowWaitForEventStepSettingsObject4Event].
	AudienceKey string `json:"audience_key"`
	JSON        struct {
		EventKey             respjson.Field
		EventType            respjson.Field
		IntegrationSourceKey respjson.Field
		SourceKey            respjson.Field
		SourceType           respjson.Field
		AudienceKey          respjson.Field
		raw                  string
	} `json:"-"`
}

func (r *WorkflowStepUnionSettingsEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowStepUnionSettingsMatchConditions is an implicit subunion of
// [WorkflowStepUnion]. WorkflowStepUnionSettingsMatchConditions provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowStepUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfWorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditions
// OfWorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditions
// OfWorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditions
// OfWorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditions
// OfWorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditions]
type WorkflowStepUnionSettingsMatchConditions struct {
	// This field will be present if the value is a
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchCondition] instead of
	// an object.
	OfWorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchCondition `json:",inline"`
	// This field will be present if the value is a
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchCondition] instead of
	// an object.
	OfWorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchCondition `json:",inline"`
	// This field will be present if the value is a
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchCondition] instead of
	// an object.
	OfWorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchCondition `json:",inline"`
	// This field will be present if the value is a
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchCondition] instead of
	// an object.
	OfWorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchCondition `json:",inline"`
	// This field will be present if the value is a
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchCondition] instead of
	// an object.
	OfWorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchCondition `json:",inline"`
	JSON                                                                 struct {
		OfWorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditions  respjson.Field
		OfWorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditions respjson.Field
		OfWorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditions respjson.Field
		OfWorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditions respjson.Field
		OfWorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditions respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (r *WorkflowStepUnionSettingsMatchConditions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowStepUnion to a WorkflowStepUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowStepUnionParam.Overrides()
func (r WorkflowStepUnion) ToParam() WorkflowStepUnionParam {
	return param.Override[WorkflowStepUnionParam](json.RawMessage(r.RawJSON()))
}

// An in-app guide step within a workflow. References a guide that will be shown to
// recipients who execute this step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
type WorkflowStepWorkflowInAppGuideStep struct {
	// The type of the channel step. Always `in_app_guide` for in-app guide steps.
	//
	// Any of "in_app_guide".
	ChannelType string `json:"channel_type" api:"required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type string `json:"type" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey string `json:"channel_group_key" api:"nullable"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey string `json:"channel_key" api:"nullable"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// The key of the guide to reference. When a recipient executes this step they are
	// added to the managed audience that backs the guide's workflow-derived targeting.
	GuideKey string `json:"guide_key" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindow `json:"send_windows" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelType     respjson.Field
		Ref             respjson.Field
		Type            respjson.Field
		ChannelGroupKey respjson.Field
		ChannelKey      respjson.Field
		Conditions      respjson.Field
		Description     respjson.Field
		GuideKey        respjson.Field
		Name            respjson.Field
		SendWindows     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowInAppGuideStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStepWorkflowInAppGuideStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A wait for event function step that pauses a workflow until a matching event is
// received.
type WorkflowStepWorkflowWaitForEventStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the wait for event step. When `event.event_type` is `message`
	// or `workflow`, `match_conditions` is required.
	Settings WorkflowStepWorkflowWaitForEventStepSettingsUnion `json:"settings" api:"required"`
	// The type of the workflow step.
	//
	// Any of "wait_for_event".
	Type string `json:"type" api:"required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref         respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		Conditions  respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowWaitForEventStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStepWorkflowWaitForEventStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowStepWorkflowWaitForEventStepSettingsUnion contains all possible
// properties and values from [WorkflowStepWorkflowWaitForEventStepSettingsObject],
// [WorkflowStepWorkflowWaitForEventStepSettingsObject2],
// [WorkflowStepWorkflowWaitForEventStepSettingsObject3],
// [WorkflowStepWorkflowWaitForEventStepSettingsObject4],
// [WorkflowStepWorkflowWaitForEventStepSettingsObject5].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkflowStepWorkflowWaitForEventStepSettingsUnion struct {
	// This field is a union of
	// [WorkflowStepWorkflowWaitForEventStepSettingsObjectEvent],
	// [WorkflowStepWorkflowWaitForEventStepSettingsObject2Event],
	// [WorkflowStepWorkflowWaitForEventStepSettingsObject3Event],
	// [WorkflowStepWorkflowWaitForEventStepSettingsObject4Event],
	// [WorkflowStepWorkflowWaitForEventStepSettingsObject5Event]
	Event WorkflowStepWorkflowWaitForEventStepSettingsUnionEvent `json:"event"`
	// This field is from variant [WorkflowStepWorkflowWaitForEventStepSettingsObject].
	ExpiresAfter Duration `json:"expires_after"`
	// This field is a union of
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchCondition],
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchCondition],
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchCondition],
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchCondition],
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchCondition]
	MatchConditions WorkflowStepWorkflowWaitForEventStepSettingsUnionMatchConditions `json:"match_conditions"`
	OnMatch         string                                                           `json:"on_match"`
	OnTimeout       string                                                           `json:"on_timeout"`
	JSON            struct {
		Event           respjson.Field
		ExpiresAfter    respjson.Field
		MatchConditions respjson.Field
		OnMatch         respjson.Field
		OnTimeout       respjson.Field
		raw             string
	} `json:"-"`
}

func (u WorkflowStepWorkflowWaitForEventStepSettingsUnion) AsWorkflowStepWorkflowWaitForEventStepSettingsObject() (v WorkflowStepWorkflowWaitForEventStepSettingsObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepWorkflowWaitForEventStepSettingsUnion) AsWorkflowStepWorkflowWaitForEventStepSettingsObject2() (v WorkflowStepWorkflowWaitForEventStepSettingsObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepWorkflowWaitForEventStepSettingsUnion) AsWorkflowStepWorkflowWaitForEventStepSettingsObject3() (v WorkflowStepWorkflowWaitForEventStepSettingsObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepWorkflowWaitForEventStepSettingsUnion) AsWorkflowStepWorkflowWaitForEventStepSettingsObject4() (v WorkflowStepWorkflowWaitForEventStepSettingsObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepWorkflowWaitForEventStepSettingsUnion) AsWorkflowStepWorkflowWaitForEventStepSettingsObject5() (v WorkflowStepWorkflowWaitForEventStepSettingsObject5) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowStepWorkflowWaitForEventStepSettingsUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkflowStepWorkflowWaitForEventStepSettingsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowStepWorkflowWaitForEventStepSettingsUnionEvent is an implicit subunion
// of [WorkflowStepWorkflowWaitForEventStepSettingsUnion].
// WorkflowStepWorkflowWaitForEventStepSettingsUnionEvent provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowStepWorkflowWaitForEventStepSettingsUnion].
type WorkflowStepWorkflowWaitForEventStepSettingsUnionEvent struct {
	EventKey  string `json:"event_key"`
	EventType string `json:"event_type"`
	// This field is from variant
	// [WorkflowStepWorkflowWaitForEventStepSettingsObjectEvent].
	IntegrationSourceKey string `json:"integration_source_key"`
	SourceKey            string `json:"source_key"`
	// This field is from variant
	// [WorkflowStepWorkflowWaitForEventStepSettingsObject2Event].
	SourceType string `json:"source_type"`
	// This field is from variant
	// [WorkflowStepWorkflowWaitForEventStepSettingsObject4Event].
	AudienceKey string `json:"audience_key"`
	JSON        struct {
		EventKey             respjson.Field
		EventType            respjson.Field
		IntegrationSourceKey respjson.Field
		SourceKey            respjson.Field
		SourceType           respjson.Field
		AudienceKey          respjson.Field
		raw                  string
	} `json:"-"`
}

func (r *WorkflowStepWorkflowWaitForEventStepSettingsUnionEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowStepWorkflowWaitForEventStepSettingsUnionMatchConditions is an implicit
// subunion of [WorkflowStepWorkflowWaitForEventStepSettingsUnion].
// WorkflowStepWorkflowWaitForEventStepSettingsUnionMatchConditions provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowStepWorkflowWaitForEventStepSettingsUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfWorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditions
// OfWorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditions
// OfWorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditions
// OfWorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditions
// OfWorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditions]
type WorkflowStepWorkflowWaitForEventStepSettingsUnionMatchConditions struct {
	// This field will be present if the value is a
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchCondition] instead of
	// an object.
	OfWorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchCondition `json:",inline"`
	// This field will be present if the value is a
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchCondition] instead of
	// an object.
	OfWorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchCondition `json:",inline"`
	// This field will be present if the value is a
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchCondition] instead of
	// an object.
	OfWorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchCondition `json:",inline"`
	// This field will be present if the value is a
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchCondition] instead of
	// an object.
	OfWorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchCondition `json:",inline"`
	// This field will be present if the value is a
	// [[]WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchCondition] instead of
	// an object.
	OfWorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchCondition `json:",inline"`
	JSON                                                                 struct {
		OfWorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditions  respjson.Field
		OfWorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditions respjson.Field
		OfWorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditions respjson.Field
		OfWorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditions respjson.Field
		OfWorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditions respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (r *WorkflowStepWorkflowWaitForEventStepSettingsUnionMatchConditions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for waiting on an integration source event.
type WorkflowStepWorkflowWaitForEventStepSettingsObject struct {
	// An integration source event to wait for.
	Event WorkflowStepWorkflowWaitForEventStepSettingsObjectEvent `json:"event" api:"required"`
	// A duration of time, represented as a unit and a value.
	ExpiresAfter Duration `json:"expires_after" api:"nullable"`
	// A list of condition groups the incoming event must match to resolve the wait.
	MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchCondition `json:"match_conditions"`
	// The action to take when a matching event is received.
	//
	// Any of "continue", "halt".
	OnMatch string `json:"on_match"`
	// The action to take when the wait expires before a match.
	//
	// Any of "continue", "halt".
	OnTimeout string `json:"on_timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event           respjson.Field
		ExpiresAfter    respjson.Field
		MatchConditions respjson.Field
		OnMatch         respjson.Field
		OnTimeout       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowWaitForEventStepSettingsObject) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An integration source event to wait for.
type WorkflowStepWorkflowWaitForEventStepSettingsObjectEvent struct {
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
func (r WorkflowStepWorkflowWaitForEventStepSettingsObjectEvent) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObjectEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchCondition struct {
	// A list of conditions.
	Conditions []Condition `json:"conditions"`
	// The operator used to join the conditions in the group.
	//
	// Any of "and".
	Operator string `json:"operator"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conditions  respjson.Field
		Operator    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchCondition) RawJSON() string {
	return r.JSON.raw
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for waiting on a message event.
type WorkflowStepWorkflowWaitForEventStepSettingsObject2 struct {
	// A message event to wait for from a message source.
	Event WorkflowStepWorkflowWaitForEventStepSettingsObject2Event `json:"event" api:"required"`
	// Required when waiting for a message event. A list of condition groups the
	// incoming event must match to resolve the wait.
	MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchCondition `json:"match_conditions" api:"required"`
	// A duration of time, represented as a unit and a value.
	ExpiresAfter Duration `json:"expires_after" api:"nullable"`
	// The action to take when a matching event is received.
	//
	// Any of "continue", "halt".
	OnMatch string `json:"on_match"`
	// The action to take when the wait expires before a match.
	//
	// Any of "continue", "halt".
	OnTimeout string `json:"on_timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event           respjson.Field
		MatchConditions respjson.Field
		ExpiresAfter    respjson.Field
		OnMatch         respjson.Field
		OnTimeout       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowWaitForEventStepSettingsObject2) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message event to wait for from a message source.
type WorkflowStepWorkflowWaitForEventStepSettingsObject2Event struct {
	// The message lifecycle event to wait for.
	//
	// Any of "created", "queued", "sent", "not_sent", "delivered",
	// "delivery_attempted", "undelivered", "bounced", "read", "unread", "seen",
	// "unseen", "archived", "unarchived", "interacted", "link_clicked".
	EventKey string `json:"event_key" api:"required"`
	// The type of event to wait for.
	//
	// Any of "message".
	EventType string `json:"event_type" api:"required"`
	// The key of the message source to scope the wait to.
	SourceKey string `json:"source_key" api:"required"`
	// The type of message source to scope the wait to.
	//
	// Any of "workflow", "broadcast", "guide".
	SourceType string `json:"source_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventKey    respjson.Field
		EventType   respjson.Field
		SourceKey   respjson.Field
		SourceType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowWaitForEventStepSettingsObject2Event) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject2Event) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchCondition struct {
	// A list of conditions.
	Conditions []Condition `json:"conditions"`
	// The operator used to join the conditions in the group.
	//
	// Any of "and".
	Operator string `json:"operator"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conditions  respjson.Field
		Operator    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchCondition) RawJSON() string {
	return r.JSON.raw
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for waiting on a workflow event.
type WorkflowStepWorkflowWaitForEventStepSettingsObject3 struct {
	// A workflow lifecycle event to wait for from a child workflow run for the same
	// recipient.
	Event WorkflowStepWorkflowWaitForEventStepSettingsObject3Event `json:"event" api:"required"`
	// Required when waiting for a workflow event. A list of condition groups the
	// incoming event must match to resolve the wait.
	MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchCondition `json:"match_conditions" api:"required"`
	// A duration of time, represented as a unit and a value.
	ExpiresAfter Duration `json:"expires_after" api:"nullable"`
	// The action to take when a matching event is received.
	//
	// Any of "continue", "halt".
	OnMatch string `json:"on_match"`
	// The action to take when the wait expires before a match.
	//
	// Any of "continue", "halt".
	OnTimeout string `json:"on_timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event           respjson.Field
		MatchConditions respjson.Field
		ExpiresAfter    respjson.Field
		OnMatch         respjson.Field
		OnTimeout       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowWaitForEventStepSettingsObject3) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A workflow lifecycle event to wait for from a child workflow run for the same
// recipient.
type WorkflowStepWorkflowWaitForEventStepSettingsObject3Event struct {
	// The workflow lifecycle event to wait for.
	//
	// Any of "started", "completed".
	EventKey string `json:"event_key" api:"required"`
	// The type of event to wait for.
	//
	// Any of "workflow".
	EventType string `json:"event_type" api:"required"`
	// The key of the workflow whose lifecycle event should match this wait.
	SourceKey string `json:"source_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventKey    respjson.Field
		EventType   respjson.Field
		SourceKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowWaitForEventStepSettingsObject3Event) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject3Event) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchCondition struct {
	// A list of conditions.
	Conditions []Condition `json:"conditions"`
	// The operator used to join the conditions in the group.
	//
	// Any of "and".
	Operator string `json:"operator"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conditions  respjson.Field
		Operator    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchCondition) RawJSON() string {
	return r.JSON.raw
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for waiting on an audience membership event.
type WorkflowStepWorkflowWaitForEventStepSettingsObject4 struct {
	// An audience membership event to wait for when a recipient enters or exits an
	// audience.
	Event WorkflowStepWorkflowWaitForEventStepSettingsObject4Event `json:"event" api:"required"`
	// A duration of time, represented as a unit and a value.
	ExpiresAfter Duration `json:"expires_after" api:"nullable"`
	// A list of condition groups the incoming event must match to resolve the wait.
	MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchCondition `json:"match_conditions"`
	// The action to take when a matching event is received.
	//
	// Any of "continue", "halt".
	OnMatch string `json:"on_match"`
	// The action to take when the wait expires before a match.
	//
	// Any of "continue", "halt".
	OnTimeout string `json:"on_timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event           respjson.Field
		ExpiresAfter    respjson.Field
		MatchConditions respjson.Field
		OnMatch         respjson.Field
		OnTimeout       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowWaitForEventStepSettingsObject4) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An audience membership event to wait for when a recipient enters or exits an
// audience.
type WorkflowStepWorkflowWaitForEventStepSettingsObject4Event struct {
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
func (r WorkflowStepWorkflowWaitForEventStepSettingsObject4Event) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject4Event) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchCondition struct {
	// A list of conditions.
	Conditions []Condition `json:"conditions"`
	// The operator used to join the conditions in the group.
	//
	// Any of "and".
	Operator string `json:"operator"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conditions  respjson.Field
		Operator    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchCondition) RawJSON() string {
	return r.JSON.raw
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for waiting on a recipient change event.
type WorkflowStepWorkflowWaitForEventStepSettingsObject5 struct {
	// A recipient updated event to wait for from the workflow recipient.
	Event WorkflowStepWorkflowWaitForEventStepSettingsObject5Event `json:"event" api:"required"`
	// A duration of time, represented as a unit and a value.
	ExpiresAfter Duration `json:"expires_after" api:"nullable"`
	// A list of condition groups the incoming event must match to resolve the wait.
	MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchCondition `json:"match_conditions"`
	// The action to take when a matching event is received.
	//
	// Any of "continue", "halt".
	OnMatch string `json:"on_match"`
	// The action to take when the wait expires before a match.
	//
	// Any of "continue", "halt".
	OnTimeout string `json:"on_timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event           respjson.Field
		ExpiresAfter    respjson.Field
		MatchConditions respjson.Field
		OnMatch         respjson.Field
		OnTimeout       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowWaitForEventStepSettingsObject5) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A recipient updated event to wait for from the workflow recipient.
type WorkflowStepWorkflowWaitForEventStepSettingsObject5Event struct {
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
func (r WorkflowStepWorkflowWaitForEventStepSettingsObject5Event) RawJSON() string { return r.JSON.raw }
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject5Event) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchCondition struct {
	// A list of conditions.
	Conditions []Condition `json:"conditions"`
	// The operator used to join the conditions in the group.
	//
	// Any of "and".
	Operator string `json:"operator"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conditions  respjson.Field
		Operator    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchCondition) RawJSON() string {
	return r.JSON.raw
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func WorkflowStepParamOfWorkflowWebhookStep(ref string, template WebhookTemplateParam, type_ WorkflowWebhookStepType) WorkflowStepUnionParam {
	var variant WorkflowWebhookStepParam
	variant.Ref = ref
	variant.Template = template
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowWebhookStep: &variant}
}

func WorkflowStepParamOfWorkflowInAppFeedStep(ref string, template InAppFeedTemplateParam, type_ WorkflowInAppFeedStepType) WorkflowStepUnionParam {
	var variant WorkflowInAppFeedStepParam
	variant.Ref = ref
	variant.Template = template
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowInAppFeedStep: &variant}
}

func WorkflowStepParamOfWorkflowInAppGuideStep(channelType string, ref string, type_ string) WorkflowStepUnionParam {
	var variant WorkflowStepWorkflowInAppGuideStepParam
	variant.ChannelType = channelType
	variant.Ref = ref
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowInAppGuideStep: &variant}
}

func WorkflowStepParamOfWorkflowChatStep(ref string, template ChatTemplateParam, type_ WorkflowChatStepType) WorkflowStepUnionParam {
	var variant WorkflowChatStepParam
	variant.Ref = ref
	variant.Template = template
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowChatStep: &variant}
}

func WorkflowStepParamOfWorkflowSMSStep(ref string, template SMSTemplateParam, type_ WorkflowSMSStepType) WorkflowStepUnionParam {
	var variant WorkflowSMSStepParam
	variant.Ref = ref
	variant.Template = template
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowSMSStep: &variant}
}

func WorkflowStepParamOfWorkflowPushStep(ref string, template PushTemplateParam, type_ WorkflowPushStepType) WorkflowStepUnionParam {
	var variant WorkflowPushStepParam
	variant.Ref = ref
	variant.Template = template
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowPushStep: &variant}
}

func WorkflowStepParamOfWorkflowEmailStep(ref string, template EmailTemplateParam, type_ WorkflowEmailStepType) WorkflowStepUnionParam {
	var variant WorkflowEmailStepParam
	variant.Ref = ref
	variant.Template = template
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowEmailStep: &variant}
}

func WorkflowStepParamOfWorkflowAIAgentStep(ref string, settings WorkflowAIAgentStepSettingsParam, type_ WorkflowAIAgentStepType) WorkflowStepUnionParam {
	var variant WorkflowAIAgentStepParam
	variant.Ref = ref
	variant.Settings = settings
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowAIAgentStep: &variant}
}

func WorkflowStepParamOfWorkflowDelayStep(ref string, settings WorkflowDelayStepSettingsParam, type_ WorkflowDelayStepType) WorkflowStepUnionParam {
	var variant WorkflowDelayStepParam
	variant.Ref = ref
	variant.Settings = settings
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowDelayStep: &variant}
}

func WorkflowStepParamOfWorkflowWaitForEventStep[
	T WorkflowStepWorkflowWaitForEventStepSettingsObjectParam | WorkflowStepWorkflowWaitForEventStepSettingsObject2Param | WorkflowStepWorkflowWaitForEventStepSettingsObject3Param | WorkflowStepWorkflowWaitForEventStepSettingsObject4Param | WorkflowStepWorkflowWaitForEventStepSettingsObject5Param,
](ref string, settings T, type_ string) WorkflowStepUnionParam {
	var variant WorkflowStepWorkflowWaitForEventStepParam
	variant.Ref = ref
	switch v := any(settings).(type) {
	case WorkflowStepWorkflowWaitForEventStepSettingsObjectParam:
		variant.Settings.OfWorkflowStepWorkflowWaitForEventStepSettingsObject = &v
	case WorkflowStepWorkflowWaitForEventStepSettingsObject2Param:
		variant.Settings.OfWorkflowStepWorkflowWaitForEventStepSettingsObject2 = &v
	case WorkflowStepWorkflowWaitForEventStepSettingsObject3Param:
		variant.Settings.OfWorkflowStepWorkflowWaitForEventStepSettingsObject3 = &v
	case WorkflowStepWorkflowWaitForEventStepSettingsObject4Param:
		variant.Settings.OfWorkflowStepWorkflowWaitForEventStepSettingsObject4 = &v
	case WorkflowStepWorkflowWaitForEventStepSettingsObject5Param:
		variant.Settings.OfWorkflowStepWorkflowWaitForEventStepSettingsObject5 = &v
	}
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowWaitForEventStep: &variant}
}

func WorkflowStepParamOfWorkflowBatchStep(ref string, settings WorkflowBatchStepSettingsParam, type_ WorkflowBatchStepType) WorkflowStepUnionParam {
	var variant WorkflowBatchStepParam
	variant.Ref = ref
	variant.Settings = settings
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowBatchStep: &variant}
}

func WorkflowStepParamOfWorkflowFetchStep(ref string, settings RequestTemplateParam, type_ WorkflowFetchStepType) WorkflowStepUnionParam {
	var variant WorkflowFetchStepParam
	variant.Ref = ref
	variant.Settings = settings
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowFetchStep: &variant}
}

func WorkflowStepParamOfWorkflowUpdateDataStep(ref string, settings WorkflowUpdateDataStepSettingsParam, type_ WorkflowUpdateDataStepType) WorkflowStepUnionParam {
	var variant WorkflowUpdateDataStepParam
	variant.Ref = ref
	variant.Settings = settings
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowUpdateDataStep: &variant}
}

func WorkflowStepParamOfWorkflowUpdateObjectStep(ref string, settings WorkflowUpdateObjectStepSettingsParam, type_ WorkflowUpdateObjectStepType) WorkflowStepUnionParam {
	var variant WorkflowUpdateObjectStepParam
	variant.Ref = ref
	variant.Settings = settings
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowUpdateObjectStep: &variant}
}

func WorkflowStepParamOfWorkflowUpdateTenantStep(ref string, settings WorkflowUpdateTenantStepSettingsParam, type_ WorkflowUpdateTenantStepType) WorkflowStepUnionParam {
	var variant WorkflowUpdateTenantStepParam
	variant.Ref = ref
	variant.Settings = settings
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowUpdateTenantStep: &variant}
}

func WorkflowStepParamOfWorkflowUpdateUserStep(ref string, settings WorkflowUpdateUserStepSettingsParam, type_ WorkflowUpdateUserStepType) WorkflowStepUnionParam {
	var variant WorkflowUpdateUserStepParam
	variant.Ref = ref
	variant.Settings = settings
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowUpdateUserStep: &variant}
}

func WorkflowStepParamOfWorkflowThrottleStep(ref string, settings WorkflowThrottleStepSettingsParam, type_ WorkflowThrottleStepType) WorkflowStepUnionParam {
	var variant WorkflowThrottleStepParam
	variant.Ref = ref
	variant.Settings = settings
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowThrottleStep: &variant}
}

func WorkflowStepParamOfWorkflowBranchStep(branches []WorkflowBranchStepBranchParam, ref string, type_ WorkflowBranchStepType) WorkflowStepUnionParam {
	var variant WorkflowBranchStepParam
	variant.Branches = branches
	variant.Ref = ref
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowBranchStep: &variant}
}

func WorkflowStepParamOfWorkflowRandomCohortStep(cohortBranches []any, ref string, type_ WorkflowRandomCohortStepType) WorkflowStepUnionParam {
	var variant WorkflowRandomCohortStepParam
	variant.CohortBranches = cohortBranches
	variant.Ref = ref
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowRandomCohortStep: &variant}
}

func WorkflowStepParamOfWorkflowTriggerWorkflowStep(ref string, settings WorkflowTriggerWorkflowStepSettingsParam, type_ WorkflowTriggerWorkflowStepType) WorkflowStepUnionParam {
	var variant WorkflowTriggerWorkflowStepParam
	variant.Ref = ref
	variant.Settings = settings
	variant.Type = type_
	return WorkflowStepUnionParam{OfWorkflowTriggerWorkflowStep: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowStepUnionParam struct {
	OfWorkflowWebhookStep         *WorkflowWebhookStepParam                  `json:",omitzero,inline"`
	OfWorkflowInAppFeedStep       *WorkflowInAppFeedStepParam                `json:",omitzero,inline"`
	OfWorkflowInAppGuideStep      *WorkflowStepWorkflowInAppGuideStepParam   `json:",omitzero,inline"`
	OfWorkflowChatStep            *WorkflowChatStepParam                     `json:",omitzero,inline"`
	OfWorkflowSMSStep             *WorkflowSMSStepParam                      `json:",omitzero,inline"`
	OfWorkflowPushStep            *WorkflowPushStepParam                     `json:",omitzero,inline"`
	OfWorkflowEmailStep           *WorkflowEmailStepParam                    `json:",omitzero,inline"`
	OfWorkflowAIAgentStep         *WorkflowAIAgentStepParam                  `json:",omitzero,inline"`
	OfWorkflowDelayStep           *WorkflowDelayStepParam                    `json:",omitzero,inline"`
	OfWorkflowWaitForEventStep    *WorkflowStepWorkflowWaitForEventStepParam `json:",omitzero,inline"`
	OfWorkflowBatchStep           *WorkflowBatchStepParam                    `json:",omitzero,inline"`
	OfWorkflowFetchStep           *WorkflowFetchStepParam                    `json:",omitzero,inline"`
	OfWorkflowUpdateDataStep      *WorkflowUpdateDataStepParam               `json:",omitzero,inline"`
	OfWorkflowUpdateObjectStep    *WorkflowUpdateObjectStepParam             `json:",omitzero,inline"`
	OfWorkflowUpdateTenantStep    *WorkflowUpdateTenantStepParam             `json:",omitzero,inline"`
	OfWorkflowUpdateUserStep      *WorkflowUpdateUserStepParam               `json:",omitzero,inline"`
	OfWorkflowThrottleStep        *WorkflowThrottleStepParam                 `json:",omitzero,inline"`
	OfWorkflowBranchStep          *WorkflowBranchStepParam                   `json:",omitzero,inline"`
	OfWorkflowRandomCohortStep    *WorkflowRandomCohortStepParam             `json:",omitzero,inline"`
	OfWorkflowTriggerWorkflowStep *WorkflowTriggerWorkflowStepParam          `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowStepUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWorkflowWebhookStep,
		u.OfWorkflowInAppFeedStep,
		u.OfWorkflowInAppGuideStep,
		u.OfWorkflowChatStep,
		u.OfWorkflowSMSStep,
		u.OfWorkflowPushStep,
		u.OfWorkflowEmailStep,
		u.OfWorkflowAIAgentStep,
		u.OfWorkflowDelayStep,
		u.OfWorkflowWaitForEventStep,
		u.OfWorkflowBatchStep,
		u.OfWorkflowFetchStep,
		u.OfWorkflowUpdateDataStep,
		u.OfWorkflowUpdateObjectStep,
		u.OfWorkflowUpdateTenantStep,
		u.OfWorkflowUpdateUserStep,
		u.OfWorkflowThrottleStep,
		u.OfWorkflowBranchStep,
		u.OfWorkflowRandomCohortStep,
		u.OfWorkflowTriggerWorkflowStep)
}
func (u *WorkflowStepUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowStepUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWorkflowWebhookStep) {
		return u.OfWorkflowWebhookStep
	} else if !param.IsOmitted(u.OfWorkflowInAppFeedStep) {
		return u.OfWorkflowInAppFeedStep
	} else if !param.IsOmitted(u.OfWorkflowInAppGuideStep) {
		return u.OfWorkflowInAppGuideStep
	} else if !param.IsOmitted(u.OfWorkflowChatStep) {
		return u.OfWorkflowChatStep
	} else if !param.IsOmitted(u.OfWorkflowSMSStep) {
		return u.OfWorkflowSMSStep
	} else if !param.IsOmitted(u.OfWorkflowPushStep) {
		return u.OfWorkflowPushStep
	} else if !param.IsOmitted(u.OfWorkflowEmailStep) {
		return u.OfWorkflowEmailStep
	} else if !param.IsOmitted(u.OfWorkflowAIAgentStep) {
		return u.OfWorkflowAIAgentStep
	} else if !param.IsOmitted(u.OfWorkflowDelayStep) {
		return u.OfWorkflowDelayStep
	} else if !param.IsOmitted(u.OfWorkflowWaitForEventStep) {
		return u.OfWorkflowWaitForEventStep
	} else if !param.IsOmitted(u.OfWorkflowBatchStep) {
		return u.OfWorkflowBatchStep
	} else if !param.IsOmitted(u.OfWorkflowFetchStep) {
		return u.OfWorkflowFetchStep
	} else if !param.IsOmitted(u.OfWorkflowUpdateDataStep) {
		return u.OfWorkflowUpdateDataStep
	} else if !param.IsOmitted(u.OfWorkflowUpdateObjectStep) {
		return u.OfWorkflowUpdateObjectStep
	} else if !param.IsOmitted(u.OfWorkflowUpdateTenantStep) {
		return u.OfWorkflowUpdateTenantStep
	} else if !param.IsOmitted(u.OfWorkflowUpdateUserStep) {
		return u.OfWorkflowUpdateUserStep
	} else if !param.IsOmitted(u.OfWorkflowThrottleStep) {
		return u.OfWorkflowThrottleStep
	} else if !param.IsOmitted(u.OfWorkflowBranchStep) {
		return u.OfWorkflowBranchStep
	} else if !param.IsOmitted(u.OfWorkflowRandomCohortStep) {
		return u.OfWorkflowRandomCohortStep
	} else if !param.IsOmitted(u.OfWorkflowTriggerWorkflowStep) {
		return u.OfWorkflowTriggerWorkflowStep
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetGuideKey() *string {
	if vt := u.OfWorkflowInAppGuideStep; vt != nil && vt.GuideKey.Valid() {
		return &vt.GuideKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetBranches() []WorkflowBranchStepBranchParam {
	if vt := u.OfWorkflowBranchStep; vt != nil {
		return vt.Branches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetCohortBranches() []any {
	if vt := u.OfWorkflowRandomCohortStep; vt != nil {
		return vt.CohortBranches
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetCohortKey() *string {
	if vt := u.OfWorkflowRandomCohortStep; vt != nil && vt.CohortKey.Valid() {
		return &vt.CohortKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetRef() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowInAppGuideStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowChatStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowSMSStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowPushStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowEmailStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowAIAgentStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowDelayStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowWaitForEventStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowBatchStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowFetchStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowUpdateDataStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowUpdateObjectStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowUpdateTenantStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowUpdateUserStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowThrottleStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowBranchStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowRandomCohortStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowTriggerWorkflowStep; vt != nil {
		return (*string)(&vt.Ref)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetType() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowInAppGuideStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowChatStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowSMSStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowPushStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowEmailStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowAIAgentStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowDelayStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowWaitForEventStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowBatchStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowFetchStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowUpdateDataStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowUpdateObjectStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowUpdateTenantStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowUpdateUserStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowThrottleStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowBranchStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowRandomCohortStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowTriggerWorkflowStep; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetChannelGroupKey() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil && vt.ChannelGroupKey.Valid() {
		return &vt.ChannelGroupKey.Value
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil && vt.ChannelGroupKey.Valid() {
		return &vt.ChannelGroupKey.Value
	} else if vt := u.OfWorkflowInAppGuideStep; vt != nil && vt.ChannelGroupKey.Valid() {
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
func (u WorkflowStepUnionParam) GetChannelKey() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil && vt.ChannelKey.Valid() {
		return &vt.ChannelKey.Value
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil && vt.ChannelKey.Valid() {
		return &vt.ChannelKey.Value
	} else if vt := u.OfWorkflowInAppGuideStep; vt != nil && vt.ChannelKey.Valid() {
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
func (u WorkflowStepUnionParam) GetChannelType() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil {
		return (*string)(&vt.ChannelType)
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil {
		return (*string)(&vt.ChannelType)
	} else if vt := u.OfWorkflowInAppGuideStep; vt != nil {
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
func (u WorkflowStepUnionParam) GetDescription() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowInAppGuideStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowChatStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowSMSStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowPushStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowEmailStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowAIAgentStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowDelayStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowWaitForEventStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowBatchStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowFetchStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowUpdateDataStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowUpdateObjectStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowUpdateTenantStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowUpdateUserStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowThrottleStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowBranchStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowRandomCohortStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowTriggerWorkflowStep; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetName() *string {
	if vt := u.OfWorkflowWebhookStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowInAppGuideStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowChatStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowSMSStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowPushStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowEmailStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowAIAgentStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowDelayStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowWaitForEventStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowBatchStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowFetchStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowUpdateDataStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowUpdateObjectStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowUpdateTenantStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowUpdateUserStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowThrottleStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowBranchStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowRandomCohortStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfWorkflowTriggerWorkflowStep; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u WorkflowStepUnionParam) GetTemplate() (res workflowStepUnionParamTemplate) {
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
type workflowStepUnionParamTemplate struct{ any }

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
func (u workflowStepUnionParamTemplate) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetMethod() *string {
	switch vt := u.any.(type) {
	case *WebhookTemplateParam:
		return (*string)(&vt.Method)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetURL() *string {
	switch vt := u.any.(type) {
	case *WebhookTemplateParam:
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetBody() *string {
	switch vt := u.any.(type) {
	case *WebhookTemplateParam:
		return paramutil.AddrIfPresent(vt.Body)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetHeaders() []WebhookTemplateHeaderParam {
	switch vt := u.any.(type) {
	case *WebhookTemplateParam:
		return vt.Headers
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetQueryParams() []WebhookTemplateQueryParamParam {
	switch vt := u.any.(type) {
	case *WebhookTemplateParam:
		return vt.QueryParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetActionButtons() []InAppFeedTemplateActionButtonParam {
	switch vt := u.any.(type) {
	case *InAppFeedTemplateParam:
		return vt.ActionButtons
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetActionURL() *string {
	switch vt := u.any.(type) {
	case *InAppFeedTemplateParam:
		return paramutil.AddrIfPresent(vt.ActionURL)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetJsonBody() *string {
	switch vt := u.any.(type) {
	case *ChatTemplateParam:
		return paramutil.AddrIfPresent(vt.JsonBody)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetSummary() *string {
	switch vt := u.any.(type) {
	case *ChatTemplateParam:
		return paramutil.AddrIfPresent(vt.Summary)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetTitle() *string {
	switch vt := u.any.(type) {
	case *PushTemplateParam:
		return &vt.Title
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetSubject() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateParam:
		return &vt.Subject
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetHTMLBody() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateParam:
		return paramutil.AddrIfPresent(vt.HTMLBody)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetIsMjml() *bool {
	switch vt := u.any.(type) {
	case *EmailTemplateParam:
		return paramutil.AddrIfPresent(vt.IsMjml)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetVisualBlocks() []EmailTemplateVisualBlockUnionParam {
	switch vt := u.any.(type) {
	case *EmailTemplateParam:
		return vt.VisualBlocks
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetMarkdownBody() *string {
	switch vt := u.any.(type) {
	case *InAppFeedTemplateParam:
		return (*string)(&vt.MarkdownBody)
	case *ChatTemplateParam:
		return (*string)(&vt.MarkdownBody)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplate) GetTextBody() *string {
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
func (u workflowStepUnionParamTemplate) GetSettings() (res workflowStepUnionParamTemplateSettings) {
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
type workflowStepUnionParamTemplateSettings struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.SMSTemplateSettingsParam:
//	case *knockmapi.PushTemplateSettingsParam:
//	case *knockmapi.EmailTemplateSettingsParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowStepUnionParamTemplateSettings) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplateSettings) GetToNumber() *string {
	switch vt := u.any.(type) {
	case *SMSTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.ToNumber)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplateSettings) GetDeliveryType() *string {
	switch vt := u.any.(type) {
	case *PushTemplateSettingsParam:
		return &vt.DeliveryType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplateSettings) GetAttachmentKey() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.AttachmentKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplateSettings) GetLayoutKey() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.LayoutKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplateSettings) GetPreContent() *string {
	switch vt := u.any.(type) {
	case *EmailTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.PreContent)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamTemplateSettings) GetPayloadOverrides() *string {
	switch vt := u.any.(type) {
	case *SMSTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.PayloadOverrides)
	case *PushTemplateSettingsParam:
		return paramutil.AddrIfPresent(vt.PayloadOverrides)
	}
	return nil
}

// Returns a pointer to the underlying variant's Conditions property, if present.
func (u WorkflowStepUnionParam) GetConditions() *ConditionGroupUnionParam {
	if vt := u.OfWorkflowWebhookStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowInAppGuideStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowChatStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowSMSStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowPushStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowEmailStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowAIAgentStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowDelayStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowWaitForEventStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowFetchStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowUpdateDataStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowUpdateObjectStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowUpdateTenantStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowUpdateUserStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowThrottleStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowTriggerWorkflowStep; vt != nil {
		return &vt.Conditions
	}
	return nil
}

// Returns a pointer to the underlying variant's SendWindows property, if present.
func (u WorkflowStepUnionParam) GetSendWindows() []SendWindowParam {
	if vt := u.OfWorkflowWebhookStep; vt != nil {
		return vt.SendWindows
	} else if vt := u.OfWorkflowInAppFeedStep; vt != nil {
		return vt.SendWindows
	} else if vt := u.OfWorkflowInAppGuideStep; vt != nil {
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
func (u WorkflowStepUnionParam) GetChannelOverrides() (res workflowStepUnionParamChannelOverrides) {
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
type workflowStepUnionParamChannelOverrides struct{ any }

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
func (u workflowStepUnionParamChannelOverrides) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamChannelOverrides) GetEmailBasedUserIDResolution() *bool {
	switch vt := u.any.(type) {
	case *ChatChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.EmailBasedUserIDResolution)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamChannelOverrides) GetTokenDeregistration() *bool {
	switch vt := u.any.(type) {
	case *PushChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.TokenDeregistration)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamChannelOverrides) GetBccAddress() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.BccAddress)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamChannelOverrides) GetCcAddress() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.CcAddress)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamChannelOverrides) GetFromAddress() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.FromAddress)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamChannelOverrides) GetFromName() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.FromName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamChannelOverrides) GetJsonOverrides() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.JsonOverrides)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamChannelOverrides) GetOpenTracking() *bool {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.OpenTracking)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamChannelOverrides) GetReplyToAddress() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.ReplyToAddress)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamChannelOverrides) GetToAddress() *string {
	switch vt := u.any.(type) {
	case *EmailChannelSettingsParam:
		return paramutil.AddrIfPresent(vt.ToAddress)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamChannelOverrides) GetLinkTracking() *bool {
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

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u WorkflowStepUnionParam) GetSettings() (res workflowStepUnionParamSettings) {
	if vt := u.OfWorkflowAIAgentStep; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfWorkflowDelayStep; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfWorkflowWaitForEventStep; vt != nil {
		res.any = vt.Settings.asAny()
	} else if vt := u.OfWorkflowBatchStep; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfWorkflowFetchStep; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfWorkflowUpdateDataStep; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfWorkflowUpdateObjectStep; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfWorkflowUpdateTenantStep; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfWorkflowUpdateUserStep; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfWorkflowThrottleStep; vt != nil {
		res.any = &vt.Settings
	} else if vt := u.OfWorkflowTriggerWorkflowStep; vt != nil {
		res.any = &vt.Settings
	}
	return
}

// Can have the runtime types [*WorkflowAIAgentStepSettingsParam],
// [*WorkflowDelayStepSettingsParam],
// [*WorkflowStepWorkflowWaitForEventStepSettingsObjectParam],
// [*WorkflowStepWorkflowWaitForEventStepSettingsObject2Param],
// [*WorkflowStepWorkflowWaitForEventStepSettingsObject3Param],
// [*WorkflowStepWorkflowWaitForEventStepSettingsObject4Param],
// [*WorkflowStepWorkflowWaitForEventStepSettingsObject5Param],
// [*WorkflowBatchStepSettingsParam], [*RequestTemplateParam],
// [*WorkflowUpdateDataStepSettingsParam],
// [*WorkflowUpdateObjectStepSettingsParam],
// [*WorkflowUpdateTenantStepSettingsParam],
// [*WorkflowUpdateUserStepSettingsParam], [*WorkflowThrottleStepSettingsParam],
// [*WorkflowTriggerWorkflowStepSettingsParam]
type workflowStepUnionParamSettings struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.WorkflowAIAgentStepSettingsParam:
//	case *knockmapi.WorkflowDelayStepSettingsParam:
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObjectParam:
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject2Param:
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject3Param:
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject4Param:
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject5Param:
//	case *knockmapi.WorkflowBatchStepSettingsParam:
//	case *knockmapi.RequestTemplateParam:
//	case *knockmapi.WorkflowUpdateDataStepSettingsParam:
//	case *knockmapi.WorkflowUpdateObjectStepSettingsParam:
//	case *knockmapi.WorkflowUpdateTenantStepSettingsParam:
//	case *knockmapi.WorkflowUpdateUserStepSettingsParam:
//	case *knockmapi.WorkflowThrottleStepSettingsParam:
//	case *knockmapi.WorkflowTriggerWorkflowStepSettingsParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowStepUnionParamSettings) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetModel() *string {
	switch vt := u.any.(type) {
	case *WorkflowAIAgentStepSettingsParam:
		return &vt.Model
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetRequestPrompt() *string {
	switch vt := u.any.(type) {
	case *WorkflowAIAgentStepSettingsParam:
		return &vt.RequestPrompt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetResponseType() *string {
	switch vt := u.any.(type) {
	case *WorkflowAIAgentStepSettingsParam:
		return &vt.ResponseType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetHaltOnError() *bool {
	switch vt := u.any.(type) {
	case *WorkflowAIAgentStepSettingsParam:
		return paramutil.AddrIfPresent(vt.HaltOnError)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetResponseSchema() *string {
	switch vt := u.any.(type) {
	case *WorkflowAIAgentStepSettingsParam:
		return paramutil.AddrIfPresent(vt.ResponseSchema)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetWebSearchEnabled() *bool {
	switch vt := u.any.(type) {
	case *WorkflowAIAgentStepSettingsParam:
		return paramutil.AddrIfPresent(vt.WebSearchEnabled)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetDelayFor() *DurationParam {
	switch vt := u.any.(type) {
	case *WorkflowDelayStepSettingsParam:
		return &vt.DelayFor
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetDelayUntilFieldPath() *string {
	switch vt := u.any.(type) {
	case *WorkflowDelayStepSettingsParam:
		return paramutil.AddrIfPresent(vt.DelayUntilFieldPath)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchExecutionMode() *string {
	switch vt := u.any.(type) {
	case *WorkflowBatchStepSettingsParam:
		return &vt.BatchExecutionMode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchItemsMaxLimit() *int64 {
	switch vt := u.any.(type) {
	case *WorkflowBatchStepSettingsParam:
		return paramutil.AddrIfPresent(vt.BatchItemsMaxLimit)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchItemsRenderLimit() *int64 {
	switch vt := u.any.(type) {
	case *WorkflowBatchStepSettingsParam:
		return paramutil.AddrIfPresent(vt.BatchItemsRenderLimit)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchKey() *string {
	switch vt := u.any.(type) {
	case *WorkflowBatchStepSettingsParam:
		return paramutil.AddrIfPresent(vt.BatchKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchOrder() *string {
	switch vt := u.any.(type) {
	case *WorkflowBatchStepSettingsParam:
		return &vt.BatchOrder
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchUntilFieldPath() *string {
	switch vt := u.any.(type) {
	case *WorkflowBatchStepSettingsParam:
		return paramutil.AddrIfPresent(vt.BatchUntilFieldPath)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchWindow() *DurationParam {
	switch vt := u.any.(type) {
	case *WorkflowBatchStepSettingsParam:
		return &vt.BatchWindow
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchWindowExtensionLimit() *DurationParam {
	switch vt := u.any.(type) {
	case *WorkflowBatchStepSettingsParam:
		return &vt.BatchWindowExtensionLimit
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchWindowType() *string {
	switch vt := u.any.(type) {
	case *WorkflowBatchStepSettingsParam:
		return &vt.BatchWindowType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetWorkflowVersionMode() *string {
	switch vt := u.any.(type) {
	case *WorkflowBatchStepSettingsParam:
		return &vt.WorkflowVersionMode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetMethod() *string {
	switch vt := u.any.(type) {
	case *RequestTemplateParam:
		return (*string)(&vt.Method)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetURL() *string {
	switch vt := u.any.(type) {
	case *RequestTemplateParam:
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBody() *string {
	switch vt := u.any.(type) {
	case *RequestTemplateParam:
		return paramutil.AddrIfPresent(vt.Body)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetHeaders() *RequestTemplateHeadersUnionParam {
	switch vt := u.any.(type) {
	case *RequestTemplateParam:
		return &vt.Headers
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetQueryParams() *RequestTemplateQueryParamsUnionParam {
	switch vt := u.any.(type) {
	case *RequestTemplateParam:
		return &vt.QueryParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetThrottleKey() *string {
	switch vt := u.any.(type) {
	case *WorkflowThrottleStepSettingsParam:
		return paramutil.AddrIfPresent(vt.ThrottleKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetThrottleLimit() *int64 {
	switch vt := u.any.(type) {
	case *WorkflowThrottleStepSettingsParam:
		return paramutil.AddrIfPresent(vt.ThrottleLimit)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetThrottleWindow() *DurationParam {
	switch vt := u.any.(type) {
	case *WorkflowThrottleStepSettingsParam:
		return &vt.ThrottleWindow
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetThrottleWindowFieldPath() *string {
	switch vt := u.any.(type) {
	case *WorkflowThrottleStepSettingsParam:
		return paramutil.AddrIfPresent(vt.ThrottleWindowFieldPath)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetActor() *string {
	switch vt := u.any.(type) {
	case *WorkflowTriggerWorkflowStepSettingsParam:
		return paramutil.AddrIfPresent(vt.Actor)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetCancellationKey() *string {
	switch vt := u.any.(type) {
	case *WorkflowTriggerWorkflowStepSettingsParam:
		return paramutil.AddrIfPresent(vt.CancellationKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetRecipients() *string {
	switch vt := u.any.(type) {
	case *WorkflowTriggerWorkflowStepSettingsParam:
		return paramutil.AddrIfPresent(vt.Recipients)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetTenant() *string {
	switch vt := u.any.(type) {
	case *WorkflowTriggerWorkflowStepSettingsParam:
		return paramutil.AddrIfPresent(vt.Tenant)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetWorkflowKey() *string {
	switch vt := u.any.(type) {
	case *WorkflowTriggerWorkflowStepSettingsParam:
		return paramutil.AddrIfPresent(vt.WorkflowKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetOnMatch() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsUnionParam:
		return vt.GetOnMatch()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetOnTimeout() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsUnionParam:
		return vt.GetOnTimeout()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetData() *string {
	switch vt := u.any.(type) {
	case *WorkflowUpdateDataStepSettingsParam:
		return (*string)(&vt.Data)
	case *WorkflowTriggerWorkflowStepSettingsParam:
		return paramutil.AddrIfPresent(vt.Data)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetRecipientGid() *string {
	switch vt := u.any.(type) {
	case *WorkflowUpdateObjectStepSettingsParam:
		return (*string)(&vt.RecipientGid)
	case *WorkflowUpdateTenantStepSettingsParam:
		return paramutil.AddrIfPresent(vt.RecipientGid)
	case *WorkflowUpdateUserStepSettingsParam:
		return paramutil.AddrIfPresent(vt.RecipientGid)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetUpdateProperties() *string {
	switch vt := u.any.(type) {
	case *WorkflowUpdateObjectStepSettingsParam:
		return (*string)(&vt.UpdateProperties)
	case *WorkflowUpdateTenantStepSettingsParam:
		return (*string)(&vt.UpdateProperties)
	case *WorkflowUpdateUserStepSettingsParam:
		return (*string)(&vt.UpdateProperties)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetRecipientMode() *string {
	switch vt := u.any.(type) {
	case *WorkflowUpdateTenantStepSettingsParam:
		return (*string)(&vt.RecipientMode)
	case *WorkflowUpdateUserStepSettingsParam:
		return (*string)(&vt.RecipientMode)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u workflowStepUnionParamSettings) GetEvent() (res workflowStepUnionParamSettingsEvent) {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsUnionParam:
		res.any = vt.GetEvent()
	}
	return res
}

// Can have the runtime types
// [*WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam],
// [*WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam],
// [*WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam],
// [*WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam],
// [*WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam]
type workflowStepUnionParamSettingsEvent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam:
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam:
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam:
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam:
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowStepUnionParamSettingsEvent) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettingsEvent) GetIntegrationSourceKey() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam:
		return &vt.IntegrationSourceKey
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettingsEvent) GetSourceType() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam:
		return &vt.SourceType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettingsEvent) GetAudienceKey() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam:
		return &vt.AudienceKey
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettingsEvent) GetEventKey() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam:
		return (*string)(&vt.EventKey)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam:
		return (*string)(&vt.EventKey)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam:
		return (*string)(&vt.EventKey)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam:
		return (*string)(&vt.EventKey)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam:
		return (*string)(&vt.EventKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettingsEvent) GetEventType() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam:
		return (*string)(&vt.EventType)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam:
		return (*string)(&vt.EventType)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam:
		return (*string)(&vt.EventType)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam:
		return (*string)(&vt.EventType)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam:
		return (*string)(&vt.EventType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettingsEvent) GetSourceKey() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam:
		return (*string)(&vt.SourceKey)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam:
		return (*string)(&vt.SourceKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's ExpiresAfter property, if present.
func (u workflowStepUnionParamSettings) GetExpiresAfter() *DurationParam {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsUnionParam:
		return vt.GetExpiresAfter()
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u workflowStepUnionParamSettings) GetMatchConditions() (res workflowStepUnionParamSettingsMatchConditions) {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsUnionParam:
		res.any = vt.GetMatchConditions()
	}
	return res
}

// Can have the runtime types
// [_[]WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditionParam],
// [_[]WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditionParam],
// [_[]WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditionParam],
// [_[]WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditionParam],
// [\*[]WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditionParam]
type workflowStepUnionParamSettingsMatchConditions struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditionParam:
//	case *[]knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditionParam:
//	case *[]knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditionParam:
//	case *[]knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditionParam:
//	case *[]knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowStepUnionParamSettingsMatchConditions) AsAny() any { return u.any }

// An in-app guide step within a workflow. References a guide that will be shown to
// recipients who execute this step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
//
// The properties ChannelType, Ref, Type are required.
type WorkflowStepWorkflowInAppGuideStepParam struct {
	// The type of the channel step. Always `in_app_guide` for in-app guide steps.
	//
	// Any of "in_app_guide".
	ChannelType string `json:"channel_type,omitzero" api:"required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type string `json:"type,omitzero" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey param.Opt[string] `json:"channel_group_key,omitzero"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey param.Opt[string] `json:"channel_key,omitzero"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// The key of the guide to reference. When a recipient executes this step they are
	// added to the managed audience that backs the guide's workflow-derived targeting.
	GuideKey param.Opt[string] `json:"guide_key,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindowParam `json:"send_windows,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowStepWorkflowInAppGuideStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowInAppGuideStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowInAppGuideStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowInAppGuideStepParam](
		"channel_type", "in_app_guide",
	)
	apijson.RegisterFieldValidator[WorkflowStepWorkflowInAppGuideStepParam](
		"type", "channel",
	)
}

// A wait for event function step that pauses a workflow until a matching event is
// received.
//
// The properties Ref, Settings, Type are required.
type WorkflowStepWorkflowWaitForEventStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the wait for event step. When `event.event_type` is `message`
	// or `workflow`, `match_conditions` is required.
	Settings WorkflowStepWorkflowWaitForEventStepSettingsUnionParam `json:"settings,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "wait_for_event".
	Type string `json:"type,omitzero" api:"required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowStepWorkflowWaitForEventStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepParam](
		"type", "wait_for_event",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowStepWorkflowWaitForEventStepSettingsUnionParam struct {
	OfWorkflowStepWorkflowWaitForEventStepSettingsObject  *WorkflowStepWorkflowWaitForEventStepSettingsObjectParam  `json:",omitzero,inline"`
	OfWorkflowStepWorkflowWaitForEventStepSettingsObject2 *WorkflowStepWorkflowWaitForEventStepSettingsObject2Param `json:",omitzero,inline"`
	OfWorkflowStepWorkflowWaitForEventStepSettingsObject3 *WorkflowStepWorkflowWaitForEventStepSettingsObject3Param `json:",omitzero,inline"`
	OfWorkflowStepWorkflowWaitForEventStepSettingsObject4 *WorkflowStepWorkflowWaitForEventStepSettingsObject4Param `json:",omitzero,inline"`
	OfWorkflowStepWorkflowWaitForEventStepSettingsObject5 *WorkflowStepWorkflowWaitForEventStepSettingsObject5Param `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowStepWorkflowWaitForEventStepSettingsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject,
		u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject2,
		u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject3,
		u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject4,
		u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject5)
}
func (u *WorkflowStepWorkflowWaitForEventStepSettingsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowStepWorkflowWaitForEventStepSettingsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject) {
		return u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject
	} else if !param.IsOmitted(u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject2) {
		return u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject2
	} else if !param.IsOmitted(u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject3) {
		return u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject3
	} else if !param.IsOmitted(u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject4) {
		return u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject4
	} else if !param.IsOmitted(u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject5) {
		return u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject5
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepWorkflowWaitForEventStepSettingsUnionParam) GetOnMatch() *string {
	if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject; vt != nil {
		return (*string)(&vt.OnMatch)
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject2; vt != nil {
		return (*string)(&vt.OnMatch)
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject3; vt != nil {
		return (*string)(&vt.OnMatch)
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject4; vt != nil {
		return (*string)(&vt.OnMatch)
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject5; vt != nil {
		return (*string)(&vt.OnMatch)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepWorkflowWaitForEventStepSettingsUnionParam) GetOnTimeout() *string {
	if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject; vt != nil {
		return (*string)(&vt.OnTimeout)
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject2; vt != nil {
		return (*string)(&vt.OnTimeout)
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject3; vt != nil {
		return (*string)(&vt.OnTimeout)
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject4; vt != nil {
		return (*string)(&vt.OnTimeout)
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject5; vt != nil {
		return (*string)(&vt.OnTimeout)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u WorkflowStepWorkflowWaitForEventStepSettingsUnionParam) GetEvent() (res workflowStepWorkflowWaitForEventStepSettingsUnionParamEvent) {
	if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject; vt != nil {
		res.any = &vt.Event
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject2; vt != nil {
		res.any = &vt.Event
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject3; vt != nil {
		res.any = &vt.Event
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject4; vt != nil {
		res.any = &vt.Event
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject5; vt != nil {
		res.any = &vt.Event
	}
	return
}

// Can have the runtime types
// [*WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam],
// [*WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam],
// [*WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam],
// [*WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam],
// [*WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam]
type workflowStepWorkflowWaitForEventStepSettingsUnionParamEvent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam:
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam:
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam:
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam:
//	case *knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowStepWorkflowWaitForEventStepSettingsUnionParamEvent) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepWorkflowWaitForEventStepSettingsUnionParamEvent) GetIntegrationSourceKey() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam:
		return &vt.IntegrationSourceKey
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepWorkflowWaitForEventStepSettingsUnionParamEvent) GetSourceType() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam:
		return &vt.SourceType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepWorkflowWaitForEventStepSettingsUnionParamEvent) GetAudienceKey() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam:
		return &vt.AudienceKey
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepWorkflowWaitForEventStepSettingsUnionParamEvent) GetEventKey() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam:
		return (*string)(&vt.EventKey)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam:
		return (*string)(&vt.EventKey)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam:
		return (*string)(&vt.EventKey)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam:
		return (*string)(&vt.EventKey)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam:
		return (*string)(&vt.EventKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepWorkflowWaitForEventStepSettingsUnionParamEvent) GetEventType() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam:
		return (*string)(&vt.EventType)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam:
		return (*string)(&vt.EventType)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam:
		return (*string)(&vt.EventType)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam:
		return (*string)(&vt.EventType)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam:
		return (*string)(&vt.EventType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepWorkflowWaitForEventStepSettingsUnionParamEvent) GetSourceKey() *string {
	switch vt := u.any.(type) {
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam:
		return (*string)(&vt.SourceKey)
	case *WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam:
		return (*string)(&vt.SourceKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's ExpiresAfter property, if present.
func (u WorkflowStepWorkflowWaitForEventStepSettingsUnionParam) GetExpiresAfter() *DurationParam {
	if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject; vt != nil {
		return &vt.ExpiresAfter
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject2; vt != nil {
		return &vt.ExpiresAfter
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject3; vt != nil {
		return &vt.ExpiresAfter
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject4; vt != nil {
		return &vt.ExpiresAfter
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject5; vt != nil {
		return &vt.ExpiresAfter
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u WorkflowStepWorkflowWaitForEventStepSettingsUnionParam) GetMatchConditions() (res workflowStepWorkflowWaitForEventStepSettingsUnionParamMatchConditions) {
	if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject; vt != nil {
		res.any = &vt.MatchConditions
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject2; vt != nil {
		res.any = &vt.MatchConditions
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject3; vt != nil {
		res.any = &vt.MatchConditions
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject4; vt != nil {
		res.any = &vt.MatchConditions
	} else if vt := u.OfWorkflowStepWorkflowWaitForEventStepSettingsObject5; vt != nil {
		res.any = &vt.MatchConditions
	}
	return
}

// Can have the runtime types
// [_[]WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditionParam],
// [_[]WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditionParam],
// [_[]WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditionParam],
// [_[]WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditionParam],
// [\*[]WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditionParam]
type workflowStepWorkflowWaitForEventStepSettingsUnionParamMatchConditions struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditionParam:
//	case *[]knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditionParam:
//	case *[]knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditionParam:
//	case *[]knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditionParam:
//	case *[]knockmapi.WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowStepWorkflowWaitForEventStepSettingsUnionParamMatchConditions) AsAny() any {
	return u.any
}

// Settings for waiting on an integration source event.
//
// The property Event is required.
type WorkflowStepWorkflowWaitForEventStepSettingsObjectParam struct {
	// An integration source event to wait for.
	Event WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam `json:"event,omitzero" api:"required"`
	// A duration of time, represented as a unit and a value.
	ExpiresAfter DurationParam `json:"expires_after,omitzero"`
	// A list of condition groups the incoming event must match to resolve the wait.
	MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditionParam `json:"match_conditions,omitzero"`
	// The action to take when a matching event is received.
	//
	// Any of "continue", "halt".
	OnMatch string `json:"on_match,omitzero"`
	// The action to take when the wait expires before a match.
	//
	// Any of "continue", "halt".
	OnTimeout string `json:"on_timeout,omitzero"`
	paramObj
}

func (r WorkflowStepWorkflowWaitForEventStepSettingsObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObjectParam](
		"on_match", "continue", "halt",
	)
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObjectParam](
		"on_timeout", "continue", "halt",
	)
}

// An integration source event to wait for.
//
// The properties EventKey, EventType, IntegrationSourceKey are required.
type WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam struct {
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

func (r WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObjectEventParam](
		"event_type", "integration_source",
	)
}

type WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditionParam struct {
	// A list of conditions.
	Conditions []ConditionParam `json:"conditions,omitzero"`
	// The operator used to join the conditions in the group.
	//
	// Any of "and".
	Operator string `json:"operator,omitzero"`
	paramObj
}

func (r WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObjectMatchConditionParam](
		"operator", "and",
	)
}

// Settings for waiting on a message event.
//
// The properties Event, MatchConditions are required.
type WorkflowStepWorkflowWaitForEventStepSettingsObject2Param struct {
	// A message event to wait for from a message source.
	Event WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam `json:"event,omitzero" api:"required"`
	// Required when waiting for a message event. A list of condition groups the
	// incoming event must match to resolve the wait.
	MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditionParam `json:"match_conditions,omitzero" api:"required"`
	// A duration of time, represented as a unit and a value.
	ExpiresAfter DurationParam `json:"expires_after,omitzero"`
	// The action to take when a matching event is received.
	//
	// Any of "continue", "halt".
	OnMatch string `json:"on_match,omitzero"`
	// The action to take when the wait expires before a match.
	//
	// Any of "continue", "halt".
	OnTimeout string `json:"on_timeout,omitzero"`
	paramObj
}

func (r WorkflowStepWorkflowWaitForEventStepSettingsObject2Param) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObject2Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject2Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject2Param](
		"on_match", "continue", "halt",
	)
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject2Param](
		"on_timeout", "continue", "halt",
	)
}

// A message event to wait for from a message source.
//
// The properties EventKey, EventType, SourceKey, SourceType are required.
type WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam struct {
	// The message lifecycle event to wait for.
	//
	// Any of "created", "queued", "sent", "not_sent", "delivered",
	// "delivery_attempted", "undelivered", "bounced", "read", "unread", "seen",
	// "unseen", "archived", "unarchived", "interacted", "link_clicked".
	EventKey string `json:"event_key,omitzero" api:"required"`
	// The type of event to wait for.
	//
	// Any of "message".
	EventType string `json:"event_type,omitzero" api:"required"`
	// The key of the message source to scope the wait to.
	SourceKey string `json:"source_key" api:"required"`
	// The type of message source to scope the wait to.
	//
	// Any of "workflow", "broadcast", "guide".
	SourceType string `json:"source_type,omitzero" api:"required"`
	paramObj
}

func (r WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam](
		"event_key", "created", "queued", "sent", "not_sent", "delivered", "delivery_attempted", "undelivered", "bounced", "read", "unread", "seen", "unseen", "archived", "unarchived", "interacted", "link_clicked",
	)
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam](
		"event_type", "message",
	)
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject2EventParam](
		"source_type", "workflow", "broadcast", "guide",
	)
}

type WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditionParam struct {
	// A list of conditions.
	Conditions []ConditionParam `json:"conditions,omitzero"`
	// The operator used to join the conditions in the group.
	//
	// Any of "and".
	Operator string `json:"operator,omitzero"`
	paramObj
}

func (r WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject2MatchConditionParam](
		"operator", "and",
	)
}

// Settings for waiting on a workflow event.
//
// The properties Event, MatchConditions are required.
type WorkflowStepWorkflowWaitForEventStepSettingsObject3Param struct {
	// A workflow lifecycle event to wait for from a child workflow run for the same
	// recipient.
	Event WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam `json:"event,omitzero" api:"required"`
	// Required when waiting for a workflow event. A list of condition groups the
	// incoming event must match to resolve the wait.
	MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditionParam `json:"match_conditions,omitzero" api:"required"`
	// A duration of time, represented as a unit and a value.
	ExpiresAfter DurationParam `json:"expires_after,omitzero"`
	// The action to take when a matching event is received.
	//
	// Any of "continue", "halt".
	OnMatch string `json:"on_match,omitzero"`
	// The action to take when the wait expires before a match.
	//
	// Any of "continue", "halt".
	OnTimeout string `json:"on_timeout,omitzero"`
	paramObj
}

func (r WorkflowStepWorkflowWaitForEventStepSettingsObject3Param) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObject3Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject3Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject3Param](
		"on_match", "continue", "halt",
	)
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject3Param](
		"on_timeout", "continue", "halt",
	)
}

// A workflow lifecycle event to wait for from a child workflow run for the same
// recipient.
//
// The properties EventKey, EventType, SourceKey are required.
type WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam struct {
	// The workflow lifecycle event to wait for.
	//
	// Any of "started", "completed".
	EventKey string `json:"event_key,omitzero" api:"required"`
	// The type of event to wait for.
	//
	// Any of "workflow".
	EventType string `json:"event_type,omitzero" api:"required"`
	// The key of the workflow whose lifecycle event should match this wait.
	SourceKey string `json:"source_key" api:"required"`
	paramObj
}

func (r WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam](
		"event_key", "started", "completed",
	)
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject3EventParam](
		"event_type", "workflow",
	)
}

type WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditionParam struct {
	// A list of conditions.
	Conditions []ConditionParam `json:"conditions,omitzero"`
	// The operator used to join the conditions in the group.
	//
	// Any of "and".
	Operator string `json:"operator,omitzero"`
	paramObj
}

func (r WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject3MatchConditionParam](
		"operator", "and",
	)
}

// Settings for waiting on an audience membership event.
//
// The property Event is required.
type WorkflowStepWorkflowWaitForEventStepSettingsObject4Param struct {
	// An audience membership event to wait for when a recipient enters or exits an
	// audience.
	Event WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam `json:"event,omitzero" api:"required"`
	// A duration of time, represented as a unit and a value.
	ExpiresAfter DurationParam `json:"expires_after,omitzero"`
	// A list of condition groups the incoming event must match to resolve the wait.
	MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditionParam `json:"match_conditions,omitzero"`
	// The action to take when a matching event is received.
	//
	// Any of "continue", "halt".
	OnMatch string `json:"on_match,omitzero"`
	// The action to take when the wait expires before a match.
	//
	// Any of "continue", "halt".
	OnTimeout string `json:"on_timeout,omitzero"`
	paramObj
}

func (r WorkflowStepWorkflowWaitForEventStepSettingsObject4Param) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObject4Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject4Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject4Param](
		"on_match", "continue", "halt",
	)
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject4Param](
		"on_timeout", "continue", "halt",
	)
}

// An audience membership event to wait for when a recipient enters or exits an
// audience.
//
// The properties AudienceKey, EventKey, EventType are required.
type WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam struct {
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

func (r WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam](
		"event_key", "enter", "exit",
	)
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject4EventParam](
		"event_type", "audience",
	)
}

type WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditionParam struct {
	// A list of conditions.
	Conditions []ConditionParam `json:"conditions,omitzero"`
	// The operator used to join the conditions in the group.
	//
	// Any of "and".
	Operator string `json:"operator,omitzero"`
	paramObj
}

func (r WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject4MatchConditionParam](
		"operator", "and",
	)
}

// Settings for waiting on a recipient change event.
//
// The property Event is required.
type WorkflowStepWorkflowWaitForEventStepSettingsObject5Param struct {
	// A recipient updated event to wait for from the workflow recipient.
	Event WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam `json:"event,omitzero" api:"required"`
	// A duration of time, represented as a unit and a value.
	ExpiresAfter DurationParam `json:"expires_after,omitzero"`
	// A list of condition groups the incoming event must match to resolve the wait.
	MatchConditions []WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditionParam `json:"match_conditions,omitzero"`
	// The action to take when a matching event is received.
	//
	// Any of "continue", "halt".
	OnMatch string `json:"on_match,omitzero"`
	// The action to take when the wait expires before a match.
	//
	// Any of "continue", "halt".
	OnTimeout string `json:"on_timeout,omitzero"`
	paramObj
}

func (r WorkflowStepWorkflowWaitForEventStepSettingsObject5Param) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObject5Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject5Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject5Param](
		"on_match", "continue", "halt",
	)
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject5Param](
		"on_timeout", "continue", "halt",
	)
}

// A recipient updated event to wait for from the workflow recipient.
//
// The property EventType is required.
type WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam struct {
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

func (r WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam](
		"event_type", "recipient",
	)
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject5EventParam](
		"event_key", "updated",
	)
}

type WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditionParam struct {
	// A list of conditions.
	Conditions []ConditionParam `json:"conditions,omitzero"`
	// The operator used to join the conditions in the group.
	//
	// Any of "and".
	Operator string `json:"operator,omitzero"`
	paramObj
}

func (r WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowStepWorkflowWaitForEventStepSettingsObject5MatchConditionParam](
		"operator", "and",
	)
}

// A throttle function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/throttle-function).
type WorkflowThrottleStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the throttle step.
	Settings WorkflowThrottleStepSettings `json:"settings" api:"required"`
	// The type of the workflow step.
	//
	// Any of "throttle".
	Type WorkflowThrottleStepType `json:"type" api:"required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref         respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		Conditions  respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowThrottleStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowThrottleStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowThrottleStep to a WorkflowThrottleStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowThrottleStepParam.Overrides()
func (r WorkflowThrottleStep) ToParam() WorkflowThrottleStepParam {
	return param.Override[WorkflowThrottleStepParam](json.RawMessage(r.RawJSON()))
}

// The settings for the throttle step.
type WorkflowThrottleStepSettings struct {
	// The data property to use to throttle notifications per recipient.
	ThrottleKey string `json:"throttle_key" api:"nullable"`
	// The maximum number of workflows to allow within the duration window. Defaults
	// to 1.
	ThrottleLimit int64 `json:"throttle_limit" api:"nullable"`
	// A duration of time, represented as a unit and a value.
	ThrottleWindow Duration `json:"throttle_window" api:"nullable"`
	// The data path to resolve a dynamic throttle window. The resolved value must be
	// an ISO-8601 timestamp. See more in the
	// [docs](https://docs.knock.app/designing-workflows/throttle-function#set-a-dynamic-throttle-window).
	ThrottleWindowFieldPath string `json:"throttle_window_field_path" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ThrottleKey             respjson.Field
		ThrottleLimit           respjson.Field
		ThrottleWindow          respjson.Field
		ThrottleWindowFieldPath respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowThrottleStepSettings) RawJSON() string { return r.JSON.raw }
func (r *WorkflowThrottleStepSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the workflow step.
type WorkflowThrottleStepType string

const (
	WorkflowThrottleStepTypeThrottle WorkflowThrottleStepType = "throttle"
)

// A throttle function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/throttle-function).
//
// The properties Ref, Settings, Type are required.
type WorkflowThrottleStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the throttle step.
	Settings WorkflowThrottleStepSettingsParam `json:"settings,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "throttle".
	Type WorkflowThrottleStepType `json:"type,omitzero" api:"required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowThrottleStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowThrottleStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowThrottleStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The settings for the throttle step.
type WorkflowThrottleStepSettingsParam struct {
	// The data property to use to throttle notifications per recipient.
	ThrottleKey param.Opt[string] `json:"throttle_key,omitzero"`
	// The maximum number of workflows to allow within the duration window. Defaults
	// to 1.
	ThrottleLimit param.Opt[int64] `json:"throttle_limit,omitzero"`
	// The data path to resolve a dynamic throttle window. The resolved value must be
	// an ISO-8601 timestamp. See more in the
	// [docs](https://docs.knock.app/designing-workflows/throttle-function#set-a-dynamic-throttle-window).
	ThrottleWindowFieldPath param.Opt[string] `json:"throttle_window_field_path,omitzero"`
	// A duration of time, represented as a unit and a value.
	ThrottleWindow DurationParam `json:"throttle_window,omitzero"`
	paramObj
}

func (r WorkflowThrottleStepSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowThrottleStepSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowThrottleStepSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A workflow trigger function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/trigger-workflow-function).
type WorkflowTriggerWorkflowStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the workflow trigger workflow step.
	Settings WorkflowTriggerWorkflowStepSettings `json:"settings" api:"required"`
	// The type of the workflow step.
	//
	// Any of "trigger_workflow".
	Type WorkflowTriggerWorkflowStepType `json:"type" api:"required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// A description for the workflow step.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref         respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		Conditions  respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerWorkflowStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerWorkflowStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowTriggerWorkflowStep to a
// WorkflowTriggerWorkflowStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowTriggerWorkflowStepParam.Overrides()
func (r WorkflowTriggerWorkflowStep) ToParam() WorkflowTriggerWorkflowStepParam {
	return param.Override[WorkflowTriggerWorkflowStepParam](json.RawMessage(r.RawJSON()))
}

// The settings for the workflow trigger workflow step.
type WorkflowTriggerWorkflowStepSettings struct {
	// The actor to trigger the workflow with. Supports liquid.
	Actor string `json:"actor"`
	// The cancellation key to trigger the workflow with. Supports liquid.
	CancellationKey string `json:"cancellation_key"`
	// The data to be supplied to the workflow. Supports liquid.
	Data string `json:"data"`
	// The recipients or recipient to trigger the workflow for. Supports liquid.
	Recipients string `json:"recipients"`
	// The tenant to trigger the workflow with. Supports liquid.
	Tenant string `json:"tenant"`
	// The key of the workflow to trigger. Supports liquid.
	WorkflowKey string `json:"workflow_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actor           respjson.Field
		CancellationKey respjson.Field
		Data            respjson.Field
		Recipients      respjson.Field
		Tenant          respjson.Field
		WorkflowKey     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTriggerWorkflowStepSettings) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerWorkflowStepSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the workflow step.
type WorkflowTriggerWorkflowStepType string

const (
	WorkflowTriggerWorkflowStepTypeTriggerWorkflow WorkflowTriggerWorkflowStepType = "trigger_workflow"
)

// A workflow trigger function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/trigger-workflow-function).
//
// The properties Ref, Settings, Type are required.
type WorkflowTriggerWorkflowStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the workflow trigger workflow step.
	Settings WorkflowTriggerWorkflowStepSettingsParam `json:"settings,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "trigger_workflow".
	Type WorkflowTriggerWorkflowStepType `json:"type,omitzero" api:"required"`
	// A description for the workflow step.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowTriggerWorkflowStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerWorkflowStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTriggerWorkflowStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The settings for the workflow trigger workflow step.
type WorkflowTriggerWorkflowStepSettingsParam struct {
	// The actor to trigger the workflow with. Supports liquid.
	Actor param.Opt[string] `json:"actor,omitzero"`
	// The cancellation key to trigger the workflow with. Supports liquid.
	CancellationKey param.Opt[string] `json:"cancellation_key,omitzero"`
	// The data to be supplied to the workflow. Supports liquid.
	Data param.Opt[string] `json:"data,omitzero"`
	// The recipients or recipient to trigger the workflow for. Supports liquid.
	Recipients param.Opt[string] `json:"recipients,omitzero"`
	// The tenant to trigger the workflow with. Supports liquid.
	Tenant param.Opt[string] `json:"tenant,omitzero"`
	// The key of the workflow to trigger. Supports liquid.
	WorkflowKey param.Opt[string] `json:"workflow_key,omitzero"`
	paramObj
}

func (r WorkflowTriggerWorkflowStepSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerWorkflowStepSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTriggerWorkflowStepSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An update data function step. Merges data into the workflow's `data` scope for
// use in subsequent steps.
type WorkflowUpdateDataStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the update data step.
	Settings WorkflowUpdateDataStepSettings `json:"settings" api:"required"`
	// The type of the workflow step.
	//
	// Any of "update_data".
	Type WorkflowUpdateDataStepType `json:"type" api:"required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref         respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		Conditions  respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowUpdateDataStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowUpdateDataStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowUpdateDataStep to a WorkflowUpdateDataStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowUpdateDataStepParam.Overrides()
func (r WorkflowUpdateDataStep) ToParam() WorkflowUpdateDataStepParam {
	return param.Override[WorkflowUpdateDataStepParam](json.RawMessage(r.RawJSON()))
}

// The settings for the update data step.
type WorkflowUpdateDataStepSettings struct {
	// A JSON string or Liquid template that evaluates to the data to merge into the
	// workflow's data scope.
	Data string `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowUpdateDataStepSettings) RawJSON() string { return r.JSON.raw }
func (r *WorkflowUpdateDataStepSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the workflow step.
type WorkflowUpdateDataStepType string

const (
	WorkflowUpdateDataStepTypeUpdateData WorkflowUpdateDataStepType = "update_data"
)

// An update data function step. Merges data into the workflow's `data` scope for
// use in subsequent steps.
//
// The properties Ref, Settings, Type are required.
type WorkflowUpdateDataStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the update data step.
	Settings WorkflowUpdateDataStepSettingsParam `json:"settings,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "update_data".
	Type WorkflowUpdateDataStepType `json:"type,omitzero" api:"required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowUpdateDataStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateDataStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateDataStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The settings for the update data step.
//
// The property Data is required.
type WorkflowUpdateDataStepSettingsParam struct {
	// A JSON string or Liquid template that evaluates to the data to merge into the
	// workflow's data scope.
	Data string `json:"data" api:"required"`
	paramObj
}

func (r WorkflowUpdateDataStepSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateDataStepSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateDataStepSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An update object step. Updates properties of a specific object referenced in the
// workflow.
type WorkflowUpdateObjectStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the update object step.
	Settings WorkflowUpdateObjectStepSettings `json:"settings" api:"required"`
	// The type of the workflow step.
	//
	// Any of "update_object".
	Type WorkflowUpdateObjectStepType `json:"type" api:"required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref         respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		Conditions  respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowUpdateObjectStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowUpdateObjectStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowUpdateObjectStep to a
// WorkflowUpdateObjectStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowUpdateObjectStepParam.Overrides()
func (r WorkflowUpdateObjectStep) ToParam() WorkflowUpdateObjectStepParam {
	return param.Override[WorkflowUpdateObjectStepParam](json.RawMessage(r.RawJSON()))
}

// The settings for the update object step.
type WorkflowUpdateObjectStepSettings struct {
	// The global identifier (GID) of the object to update. Format:
	// gid://Object/{collection}/{id}
	RecipientGid string `json:"recipient_gid" api:"required"`
	// A JSON string or Liquid template that evaluates to the properties to update on
	// the object.
	UpdateProperties string `json:"update_properties" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecipientGid     respjson.Field
		UpdateProperties respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowUpdateObjectStepSettings) RawJSON() string { return r.JSON.raw }
func (r *WorkflowUpdateObjectStepSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the workflow step.
type WorkflowUpdateObjectStepType string

const (
	WorkflowUpdateObjectStepTypeUpdateObject WorkflowUpdateObjectStepType = "update_object"
)

// An update object step. Updates properties of a specific object referenced in the
// workflow.
//
// The properties Ref, Settings, Type are required.
type WorkflowUpdateObjectStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the update object step.
	Settings WorkflowUpdateObjectStepSettingsParam `json:"settings,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "update_object".
	Type WorkflowUpdateObjectStepType `json:"type,omitzero" api:"required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowUpdateObjectStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateObjectStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateObjectStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The settings for the update object step.
//
// The properties RecipientGid, UpdateProperties are required.
type WorkflowUpdateObjectStepSettingsParam struct {
	// The global identifier (GID) of the object to update. Format:
	// gid://Object/{collection}/{id}
	RecipientGid string `json:"recipient_gid" api:"required"`
	// A JSON string or Liquid template that evaluates to the properties to update on
	// the object.
	UpdateProperties string `json:"update_properties" api:"required"`
	paramObj
}

func (r WorkflowUpdateObjectStepSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateObjectStepSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateObjectStepSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An update tenant step. Updates properties of a specific tenant referenced in the
// workflow.
type WorkflowUpdateTenantStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the update tenant step.
	Settings WorkflowUpdateTenantStepSettings `json:"settings" api:"required"`
	// The type of the workflow step.
	//
	// Any of "update_tenant".
	Type WorkflowUpdateTenantStepType `json:"type" api:"required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref         respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		Conditions  respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowUpdateTenantStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowUpdateTenantStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowUpdateTenantStep to a
// WorkflowUpdateTenantStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowUpdateTenantStepParam.Overrides()
func (r WorkflowUpdateTenantStep) ToParam() WorkflowUpdateTenantStepParam {
	return param.Override[WorkflowUpdateTenantStepParam](json.RawMessage(r.RawJSON()))
}

// The settings for the update tenant step.
type WorkflowUpdateTenantStepSettings struct {
	// The recipient mode determining how the tenant is selected. 'current' uses the
	// workflow's current tenant. 'reference' uses a specific tenant ID.
	//
	// Any of "current", "reference".
	RecipientMode string `json:"recipient_mode" api:"required"`
	// A JSON string or Liquid template that evaluates to the properties to update on
	// the tenant.
	UpdateProperties string `json:"update_properties" api:"required"`
	// The global identifier (GID) of the tenant to update. Required when
	// recipient_mode is 'reference'. Format: gid://Object/$tenants/{id}
	RecipientGid string `json:"recipient_gid" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecipientMode    respjson.Field
		UpdateProperties respjson.Field
		RecipientGid     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowUpdateTenantStepSettings) RawJSON() string { return r.JSON.raw }
func (r *WorkflowUpdateTenantStepSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the workflow step.
type WorkflowUpdateTenantStepType string

const (
	WorkflowUpdateTenantStepTypeUpdateTenant WorkflowUpdateTenantStepType = "update_tenant"
)

// An update tenant step. Updates properties of a specific tenant referenced in the
// workflow.
//
// The properties Ref, Settings, Type are required.
type WorkflowUpdateTenantStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the update tenant step.
	Settings WorkflowUpdateTenantStepSettingsParam `json:"settings,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "update_tenant".
	Type WorkflowUpdateTenantStepType `json:"type,omitzero" api:"required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowUpdateTenantStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateTenantStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateTenantStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The settings for the update tenant step.
//
// The properties RecipientMode, UpdateProperties are required.
type WorkflowUpdateTenantStepSettingsParam struct {
	// The recipient mode determining how the tenant is selected. 'current' uses the
	// workflow's current tenant. 'reference' uses a specific tenant ID.
	//
	// Any of "current", "reference".
	RecipientMode string `json:"recipient_mode,omitzero" api:"required"`
	// A JSON string or Liquid template that evaluates to the properties to update on
	// the tenant.
	UpdateProperties string `json:"update_properties" api:"required"`
	// The global identifier (GID) of the tenant to update. Required when
	// recipient_mode is 'reference'. Format: gid://Object/$tenants/{id}
	RecipientGid param.Opt[string] `json:"recipient_gid,omitzero"`
	paramObj
}

func (r WorkflowUpdateTenantStepSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateTenantStepSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateTenantStepSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowUpdateTenantStepSettingsParam](
		"recipient_mode", "current", "reference",
	)
}

// An update user step. Updates properties of a specific user referenced in the
// workflow.
type WorkflowUpdateUserStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the update user step.
	Settings WorkflowUpdateUserStepSettings `json:"settings" api:"required"`
	// The type of the workflow step.
	//
	// Any of "update_user".
	Type WorkflowUpdateUserStepType `json:"type" api:"required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref         respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		Conditions  respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowUpdateUserStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowUpdateUserStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowUpdateUserStep to a WorkflowUpdateUserStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowUpdateUserStepParam.Overrides()
func (r WorkflowUpdateUserStep) ToParam() WorkflowUpdateUserStepParam {
	return param.Override[WorkflowUpdateUserStepParam](json.RawMessage(r.RawJSON()))
}

// The settings for the update user step.
type WorkflowUpdateUserStepSettings struct {
	// The recipient mode determining how the user is selected. 'current' uses the
	// workflow's current user. 'reference' uses a specific user ID.
	//
	// Any of "current", "reference".
	RecipientMode string `json:"recipient_mode" api:"required"`
	// A JSON string or Liquid template that evaluates to the properties to update on
	// the user.
	UpdateProperties string `json:"update_properties" api:"required"`
	// The global identifier (GID) of the user to update. Required when recipient_mode
	// is 'reference'. Format: gid://Object/$users/{id}
	RecipientGid string `json:"recipient_gid" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecipientMode    respjson.Field
		UpdateProperties respjson.Field
		RecipientGid     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowUpdateUserStepSettings) RawJSON() string { return r.JSON.raw }
func (r *WorkflowUpdateUserStepSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the workflow step.
type WorkflowUpdateUserStepType string

const (
	WorkflowUpdateUserStepTypeUpdateUser WorkflowUpdateUserStepType = "update_user"
)

// An update user step. Updates properties of a specific user referenced in the
// workflow.
//
// The properties Ref, Settings, Type are required.
type WorkflowUpdateUserStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// The settings for the update user step.
	Settings WorkflowUpdateUserStepSettingsParam `json:"settings,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "update_user".
	Type WorkflowUpdateUserStepType `json:"type,omitzero" api:"required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowUpdateUserStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateUserStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateUserStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The settings for the update user step.
//
// The properties RecipientMode, UpdateProperties are required.
type WorkflowUpdateUserStepSettingsParam struct {
	// The recipient mode determining how the user is selected. 'current' uses the
	// workflow's current user. 'reference' uses a specific user ID.
	//
	// Any of "current", "reference".
	RecipientMode string `json:"recipient_mode,omitzero" api:"required"`
	// A JSON string or Liquid template that evaluates to the properties to update on
	// the user.
	UpdateProperties string `json:"update_properties" api:"required"`
	// The global identifier (GID) of the user to update. Required when recipient_mode
	// is 'reference'. Format: gid://Object/$users/{id}
	RecipientGid param.Opt[string] `json:"recipient_gid,omitzero"`
	paramObj
}

func (r WorkflowUpdateUserStepSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateUserStepSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateUserStepSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowUpdateUserStepSettingsParam](
		"recipient_mode", "current", "reference",
	)
}

// A webhook step within a workflow to send an HTTP request to a generic channel.
// Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
type WorkflowWebhookStep struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// A webhook template. By default, a webhook step will use the request settings you
	// configured in your webhook channel. You can override this as you see fit on a
	// per-step basis.
	Template WebhookTemplate `json:"template" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowWebhookStepType `json:"type" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey string `json:"channel_group_key" api:"nullable"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey string `json:"channel_key" api:"nullable"`
	// The type of the channel step. Always `http` for webhook steps.
	//
	// Any of "http".
	ChannelType WorkflowWebhookStepChannelType `json:"channel_type"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description" api:"nullable"`
	// A name for the workflow step.
	Name string `json:"name" api:"nullable"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindow `json:"send_windows" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref             respjson.Field
		Template        respjson.Field
		Type            respjson.Field
		ChannelGroupKey respjson.Field
		ChannelKey      respjson.Field
		ChannelType     respjson.Field
		Conditions      respjson.Field
		Description     respjson.Field
		Name            respjson.Field
		SendWindows     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowWebhookStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowWebhookStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowWebhookStep to a WorkflowWebhookStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowWebhookStepParam.Overrides()
func (r WorkflowWebhookStep) ToParam() WorkflowWebhookStepParam {
	return param.Override[WorkflowWebhookStepParam](json.RawMessage(r.RawJSON()))
}

// The type of the workflow step.
type WorkflowWebhookStepType string

const (
	WorkflowWebhookStepTypeChannel WorkflowWebhookStepType = "channel"
)

// The type of the channel step. Always `http` for webhook steps.
type WorkflowWebhookStepChannelType string

const (
	WorkflowWebhookStepChannelTypeHTTP WorkflowWebhookStepChannelType = "http"
)

// A webhook step within a workflow to send an HTTP request to a generic channel.
// Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
//
// The properties Ref, Template, Type are required.
type WorkflowWebhookStepParam struct {
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref" api:"required"`
	// A webhook template. By default, a webhook step will use the request settings you
	// configured in your webhook channel. You can override this as you see fit on a
	// per-step basis.
	Template WebhookTemplateParam `json:"template,omitzero" api:"required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowWebhookStepType `json:"type,omitzero" api:"required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. Either `channel_key` or `channel_group_key` must be provided, but
	// not both.
	ChannelGroupKey param.Opt[string] `json:"channel_group_key,omitzero"`
	// The key of a specific configured channel instance (e.g., 'knock-email',
	// 'postmark', 'sendgrid-marketing') to send the notification through. Either
	// `channel_key` or `channel_group_key` must be provided, but not both.
	ChannelKey param.Opt[string] `json:"channel_key,omitzero"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the workflow step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindowParam `json:"send_windows,omitzero"`
	// The type of the channel step. Always `http` for webhook steps.
	//
	// Any of "http".
	ChannelType WorkflowWebhookStepChannelType `json:"channel_type,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r WorkflowWebhookStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowWebhookStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowWebhookStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A workflow object.
type WorkflowGetResponse struct {
	// Whether the workflow is
	// [active](https://docs.knock.app/concepts/workflows#workflow-status) in the
	// current environment. (read-only).
	Active bool `json:"active" api:"required"`
	// The timestamp of when the workflow was created. (read-only).
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The slug of the environment in which the workflow exists. (read-only).
	Environment string `json:"environment" api:"required"`
	// The unique key string for the workflow object. Must be at minimum 3 characters
	// and at maximum 255 characters in length. Must be in the format of ^[a-z0-9_-]+$.
	Key string `json:"key" api:"required"`
	// A name for the workflow. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// The SHA hash of the workflow data. (read-only).
	Sha string `json:"sha" api:"required"`
	// A list of workflow step objects in the workflow.
	Steps []WorkflowStepUnion `json:"steps" api:"required"`
	// The timestamp of when the workflow was last updated. (read-only).
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Whether the workflow and its steps are in a valid state. (read-only).
	Valid bool `json:"valid" api:"required"`
	// A list of
	// [categories](https://docs.knock.app/concepts/workflows#workflow-categories) that
	// the workflow belongs to.
	Categories []string `json:"categories"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// Information about a user within the Knock dashboard. Not to be confused with an
	// external user (recipient) of a workflow.
	CreatedBy MemberUser `json:"created_by" api:"nullable"`
	// The timestamp of when the workflow was deleted. (read-only).
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// An arbitrary string attached to a workflow object. Useful for adding notes about
	// the workflow for internal purposes. Maximum of 280 characters allowed.
	Description string `json:"description"`
	// Attaches a goal to a workflow, guide, or broadcast for attribution tracking.
	GoalAttachment WorkflowGetResponseGoalAttachment `json:"goal_attachment" api:"nullable"`
	// A map of workflow settings.
	Settings WorkflowGetResponseSettings `json:"settings"`
	// Use tags to organize resources internally within your account. For example, by
	// team or product area.
	Tags []string `json:"tags"`
	// A JSON schema for the expected structure of the workflow trigger's `data`
	// payload (available in templates as `{{ data.field_name }}`). Used to validate
	// trigger requests. Read more in the
	// [docs](https://docs.knock.app/developer-tools/validating-trigger-data).
	TriggerDataJsonSchema map[string]any `json:"trigger_data_json_schema"`
	// The frequency at which the workflow should be triggered. One of:
	// `once_per_recipient`, `once_per_recipient_per_tenant`, `every_trigger`. Defaults
	// to `every_trigger`. Read more in
	// [docs](https://docs.knock.app/send-notifications/triggering-workflows/overview#controlling-workflow-trigger-frequency).
	//
	// Any of "every_trigger", "once_per_recipient", "once_per_recipient_per_tenant".
	TriggerFrequency WorkflowGetResponseTriggerFrequency `json:"trigger_frequency"`
	// Information about a user within the Knock dashboard. Not to be confused with an
	// external user (recipient) of a workflow.
	UpdatedBy MemberUser `json:"updated_by" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active                respjson.Field
		CreatedAt             respjson.Field
		Environment           respjson.Field
		Key                   respjson.Field
		Name                  respjson.Field
		Sha                   respjson.Field
		Steps                 respjson.Field
		UpdatedAt             respjson.Field
		Valid                 respjson.Field
		Categories            respjson.Field
		Conditions            respjson.Field
		CreatedBy             respjson.Field
		DeletedAt             respjson.Field
		Description           respjson.Field
		GoalAttachment        respjson.Field
		Settings              respjson.Field
		Tags                  respjson.Field
		TriggerDataJsonSchema respjson.Field
		TriggerFrequency      respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attaches a goal to a workflow, guide, or broadcast for attribution tracking.
type WorkflowGetResponseGoalAttachment struct {
	// The key of the goal to attach.
	GoalKey string `json:"goal_key" api:"required"`
	// The number of days to attribute conversions after the notification is sent. Must
	// be between 1 and 30. Defaults to 7.
	AttributionWindowDays int64 `json:"attribution_window_days"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GoalKey               respjson.Field
		AttributionWindowDays respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetResponseGoalAttachment) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetResponseGoalAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A map of workflow settings.
type WorkflowGetResponseSettings struct {
	// Whether the workflow is commercial. Defaults to false.
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
func (r WorkflowGetResponseSettings) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetResponseSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which the workflow should be triggered. One of:
// `once_per_recipient`, `once_per_recipient_per_tenant`, `every_trigger`. Defaults
// to `every_trigger`. Read more in
// [docs](https://docs.knock.app/send-notifications/triggering-workflows/overview#controlling-workflow-trigger-frequency).
type WorkflowGetResponseTriggerFrequency string

const (
	WorkflowGetResponseTriggerFrequencyEveryTrigger              WorkflowGetResponseTriggerFrequency = "every_trigger"
	WorkflowGetResponseTriggerFrequencyOncePerRecipient          WorkflowGetResponseTriggerFrequency = "once_per_recipient"
	WorkflowGetResponseTriggerFrequencyOncePerRecipientPerTenant WorkflowGetResponseTriggerFrequency = "once_per_recipient_per_tenant"
)

// Wraps the Workflow response under the `workflow` key.
type WorkflowActivateResponse struct {
	// A workflow object. Read more in the
	// [docs](https://docs.knock.app/concepts/workflows).
	Workflow Workflow `json:"workflow" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Workflow    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActivateResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to a run workflow request.
type WorkflowRunResponse struct {
	// The ID of the workflow run.
	WorkflowRunID string `json:"workflow_run_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WorkflowRunID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRunResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Workflow response under the `workflow` key.
type WorkflowUpsertResponse struct {
	// A workflow object. Read more in the
	// [docs](https://docs.knock.app/concepts/workflows).
	Workflow Workflow `json:"workflow" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Workflow    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Workflow response under the `workflow` key.
type WorkflowValidateResponse struct {
	// A workflow object. Read more in the
	// [docs](https://docs.knock.app/concepts/workflows).
	Workflow Workflow `json:"workflow" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Workflow    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowGetParams struct {
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

// URLQuery serializes [WorkflowGetParams]'s query parameters as `url.Values`.
func (r WorkflowGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowListParams struct {
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

// URLQuery serializes [WorkflowListParams]'s query parameters as `url.Values`.
func (r WorkflowListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowActivateParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// Whether to activate or deactivate the workflow. Set to `true` by default, which
	// will activate the workflow.
	Status bool `json:"status" api:"required"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	paramObj
}

func (r WorkflowActivateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowActivateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowActivateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WorkflowActivateParams]'s query parameters as `url.Values`.
func (r WorkflowActivateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowRunParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A list of recipients to run the workflow for. Supports user IDs, object
	// references, or inline identify user objects (id + optional email/name).
	Recipients []WorkflowRunParamsRecipientUnion `json:"recipients,omitzero" api:"required"`
	// A key to cancel the workflow run.
	CancellationKey param.Opt[string] `json:"cancellation_key,omitzero"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The tenant to associate the workflow run with. Must not contain whitespace.
	Tenant param.Opt[string] `json:"tenant,omitzero"`
	// The actor to reference in the the workflow run.
	Actor WorkflowRunParamsActorUnion `json:"actor,omitzero"`
	// A map of data to be used in the workflow run. The structure should conform to
	// the workflow's `trigger_data_json_schema` if one is defined. Available in
	// templates as `{{ data.field_name }}`. See
	// [trigger data validation docs](https://docs.knock.app/developer-tools/validating-trigger-data).
	Data map[string]any `json:"data,omitzero"`
	paramObj
}

func (r WorkflowRunParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WorkflowRunParams]'s query parameters as `url.Values`.
func (r WorkflowRunParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowRunParamsRecipientUnion struct {
	OfString                    param.Opt[string]                                    `json:",omitzero,inline"`
	OfObjectRecipientReference  *WorkflowRunParamsRecipientObjectRecipientReference  `json:",omitzero,inline"`
	OfInlineIdentifyUserRequest *WorkflowRunParamsRecipientInlineIdentifyUserRequest `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowRunParamsRecipientUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfObjectRecipientReference, u.OfInlineIdentifyUserRequest)
}
func (u *WorkflowRunParamsRecipientUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowRunParamsRecipientUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfObjectRecipientReference) {
		return u.OfObjectRecipientReference
	} else if !param.IsOmitted(u.OfInlineIdentifyUserRequest) {
		return u.OfInlineIdentifyUserRequest
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowRunParamsRecipientUnion) GetCollection() *string {
	if vt := u.OfObjectRecipientReference; vt != nil {
		return &vt.Collection
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowRunParamsRecipientUnion) GetEmail() *string {
	if vt := u.OfInlineIdentifyUserRequest; vt != nil && vt.Email.Valid() {
		return &vt.Email.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowRunParamsRecipientUnion) GetName() *string {
	if vt := u.OfInlineIdentifyUserRequest; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowRunParamsRecipientUnion) GetID() *string {
	if vt := u.OfObjectRecipientReference; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfInlineIdentifyUserRequest; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// An object reference.
//
// The properties ID, Collection are required.
type WorkflowRunParamsRecipientObjectRecipientReference struct {
	// The ID of the object.
	ID string `json:"id" api:"required"`
	// The collection of the object.
	Collection string `json:"collection" api:"required"`
	paramObj
}

func (r WorkflowRunParamsRecipientObjectRecipientReference) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRunParamsRecipientObjectRecipientReference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowRunParamsRecipientObjectRecipientReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A user recipient with optional identify properties. When email or name are
// provided, the user is created or updated as part of the workflow run. The
// collection is always `$users` and should not be sent.
//
// The property ID is required.
type WorkflowRunParamsRecipientInlineIdentifyUserRequest struct {
	// The ID of the user.
	ID string `json:"id" api:"required"`
	// The email address to set on the user.
	Email param.Opt[string] `json:"email,omitzero"`
	// The display name to set on the user.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r WorkflowRunParamsRecipientInlineIdentifyUserRequest) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRunParamsRecipientInlineIdentifyUserRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowRunParamsRecipientInlineIdentifyUserRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowRunParamsActorUnion struct {
	OfString                    param.Opt[string]                                `json:",omitzero,inline"`
	OfObjectRecipientReference  *WorkflowRunParamsActorObjectRecipientReference  `json:",omitzero,inline"`
	OfInlineIdentifyUserRequest *WorkflowRunParamsActorInlineIdentifyUserRequest `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowRunParamsActorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfObjectRecipientReference, u.OfInlineIdentifyUserRequest)
}
func (u *WorkflowRunParamsActorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowRunParamsActorUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfObjectRecipientReference) {
		return u.OfObjectRecipientReference
	} else if !param.IsOmitted(u.OfInlineIdentifyUserRequest) {
		return u.OfInlineIdentifyUserRequest
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowRunParamsActorUnion) GetCollection() *string {
	if vt := u.OfObjectRecipientReference; vt != nil {
		return &vt.Collection
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowRunParamsActorUnion) GetEmail() *string {
	if vt := u.OfInlineIdentifyUserRequest; vt != nil && vt.Email.Valid() {
		return &vt.Email.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowRunParamsActorUnion) GetName() *string {
	if vt := u.OfInlineIdentifyUserRequest; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowRunParamsActorUnion) GetID() *string {
	if vt := u.OfObjectRecipientReference; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfInlineIdentifyUserRequest; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// An object reference.
//
// The properties ID, Collection are required.
type WorkflowRunParamsActorObjectRecipientReference struct {
	// The ID of the object.
	ID string `json:"id" api:"required"`
	// The collection of the object.
	Collection string `json:"collection" api:"required"`
	paramObj
}

func (r WorkflowRunParamsActorObjectRecipientReference) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRunParamsActorObjectRecipientReference
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowRunParamsActorObjectRecipientReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A user recipient with optional identify properties. When email or name are
// provided, the user is created or updated as part of the workflow run. The
// collection is always `$users` and should not be sent.
//
// The property ID is required.
type WorkflowRunParamsActorInlineIdentifyUserRequest struct {
	// The ID of the user.
	ID string `json:"id" api:"required"`
	// The email address to set on the user.
	Email param.Opt[string] `json:"email,omitzero"`
	// The display name to set on the user.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r WorkflowRunParamsActorInlineIdentifyUserRequest) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRunParamsActorInlineIdentifyUserRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowRunParamsActorInlineIdentifyUserRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A workflow request for upserting a workflow.
	Workflow WorkflowUpsertParamsWorkflow `json:"workflow,omitzero" api:"required"`
	// When used with commit, creates a new version with identical content and commits
	// it if there are no unpublished changes.
	AllowEmpty param.Opt[bool] `query:"allow_empty,omitzero" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// Whether to commit the resource at the same time as modifying it.
	Commit param.Opt[bool] `query:"commit,omitzero" json:"-"`
	// The message to commit the resource with, only used if `commit` is `true`.
	CommitMessage param.Opt[string] `query:"commit_message,omitzero" json:"-"`
	// When set to true, forces the upsert to override existing content regardless of
	// environment restrictions. This bypasses the development-only environment check
	// and origin environment checks.
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
	paramObj
}

func (r WorkflowUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WorkflowUpsertParams]'s query parameters as `url.Values`.
func (r WorkflowUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A workflow request for upserting a workflow.
//
// The properties Name, Steps are required.
type WorkflowUpsertParamsWorkflow struct {
	// A name for the workflow. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// A list of workflow step objects in the workflow.
	Steps []WorkflowStepUnionParam `json:"steps,omitzero" api:"required"`
	// An arbitrary string attached to a workflow object. Useful for adding notes about
	// the workflow for internal purposes. Maximum of 280 characters allowed.
	Description param.Opt[string] `json:"description,omitzero"`
	// Attaches a goal to a workflow, guide, or broadcast for attribution tracking.
	GoalAttachment WorkflowUpsertParamsWorkflowGoalAttachment `json:"goal_attachment,omitzero"`
	// A list of
	// [categories](https://docs.knock.app/concepts/workflows#workflow-categories) that
	// the workflow belongs to.
	Categories []string `json:"categories,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	// A map of workflow settings.
	Settings WorkflowUpsertParamsWorkflowSettings `json:"settings,omitzero"`
	// Use tags to organize resources internally within your account. For example, by
	// team or product area.
	Tags []string `json:"tags,omitzero"`
	// A JSON schema for the expected structure of the workflow trigger's `data`
	// payload (available in templates as `{{ data.field_name }}`). Used to validate
	// trigger requests. Read more in the
	// [docs](https://docs.knock.app/developer-tools/validating-trigger-data).
	TriggerDataJsonSchema map[string]any `json:"trigger_data_json_schema,omitzero"`
	// The frequency at which the workflow should be triggered. One of:
	// `once_per_recipient`, `once_per_recipient_per_tenant`, `every_trigger`. Defaults
	// to `every_trigger`. Read more in
	// [docs](https://docs.knock.app/send-notifications/triggering-workflows/overview#controlling-workflow-trigger-frequency).
	//
	// Any of "every_trigger", "once_per_recipient", "once_per_recipient_per_tenant".
	TriggerFrequency string `json:"trigger_frequency,omitzero"`
	paramObj
}

func (r WorkflowUpsertParamsWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpsertParamsWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpsertParamsWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowUpsertParamsWorkflow](
		"trigger_frequency", "every_trigger", "once_per_recipient", "once_per_recipient_per_tenant",
	)
}

// Attaches a goal to a workflow, guide, or broadcast for attribution tracking.
//
// The property GoalKey is required.
type WorkflowUpsertParamsWorkflowGoalAttachment struct {
	// The key of the goal to attach.
	GoalKey string `json:"goal_key" api:"required"`
	// The number of days to attribute conversions after the notification is sent. Must
	// be between 1 and 30. Defaults to 7.
	AttributionWindowDays param.Opt[int64] `json:"attribution_window_days,omitzero"`
	paramObj
}

func (r WorkflowUpsertParamsWorkflowGoalAttachment) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpsertParamsWorkflowGoalAttachment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpsertParamsWorkflowGoalAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A map of workflow settings.
type WorkflowUpsertParamsWorkflowSettings struct {
	// Whether the workflow is commercial. Defaults to false.
	IsCommercial param.Opt[bool] `json:"is_commercial,omitzero"`
	// Whether to ignore recipient preferences for a given type of notification. If
	// true, will send for every channel in the workflow even if the recipient has
	// opted out of a certain kind. Defaults to false.
	OverridePreferences param.Opt[bool] `json:"override_preferences,omitzero"`
	paramObj
}

func (r WorkflowUpsertParamsWorkflowSettings) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpsertParamsWorkflowSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpsertParamsWorkflowSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowValidateParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A workflow request for upserting a workflow.
	Workflow WorkflowValidateParamsWorkflow `json:"workflow,omitzero" api:"required"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	paramObj
}

func (r WorkflowValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WorkflowValidateParams]'s query parameters as `url.Values`.
func (r WorkflowValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A workflow request for upserting a workflow.
//
// The properties Name, Steps are required.
type WorkflowValidateParamsWorkflow struct {
	// A name for the workflow. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// A list of workflow step objects in the workflow.
	Steps []WorkflowStepUnionParam `json:"steps,omitzero" api:"required"`
	// An arbitrary string attached to a workflow object. Useful for adding notes about
	// the workflow for internal purposes. Maximum of 280 characters allowed.
	Description param.Opt[string] `json:"description,omitzero"`
	// Attaches a goal to a workflow, guide, or broadcast for attribution tracking.
	GoalAttachment WorkflowValidateParamsWorkflowGoalAttachment `json:"goal_attachment,omitzero"`
	// A list of
	// [categories](https://docs.knock.app/concepts/workflows#workflow-categories) that
	// the workflow belongs to.
	Categories []string `json:"categories,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	// A map of workflow settings.
	Settings WorkflowValidateParamsWorkflowSettings `json:"settings,omitzero"`
	// Use tags to organize resources internally within your account. For example, by
	// team or product area.
	Tags []string `json:"tags,omitzero"`
	// A JSON schema for the expected structure of the workflow trigger's `data`
	// payload (available in templates as `{{ data.field_name }}`). Used to validate
	// trigger requests. Read more in the
	// [docs](https://docs.knock.app/developer-tools/validating-trigger-data).
	TriggerDataJsonSchema map[string]any `json:"trigger_data_json_schema,omitzero"`
	// The frequency at which the workflow should be triggered. One of:
	// `once_per_recipient`, `once_per_recipient_per_tenant`, `every_trigger`. Defaults
	// to `every_trigger`. Read more in
	// [docs](https://docs.knock.app/send-notifications/triggering-workflows/overview#controlling-workflow-trigger-frequency).
	//
	// Any of "every_trigger", "once_per_recipient", "once_per_recipient_per_tenant".
	TriggerFrequency string `json:"trigger_frequency,omitzero"`
	paramObj
}

func (r WorkflowValidateParamsWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowValidateParamsWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowValidateParamsWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowValidateParamsWorkflow](
		"trigger_frequency", "every_trigger", "once_per_recipient", "once_per_recipient_per_tenant",
	)
}

// Attaches a goal to a workflow, guide, or broadcast for attribution tracking.
//
// The property GoalKey is required.
type WorkflowValidateParamsWorkflowGoalAttachment struct {
	// The key of the goal to attach.
	GoalKey string `json:"goal_key" api:"required"`
	// The number of days to attribute conversions after the notification is sent. Must
	// be between 1 and 30. Defaults to 7.
	AttributionWindowDays param.Opt[int64] `json:"attribution_window_days,omitzero"`
	paramObj
}

func (r WorkflowValidateParamsWorkflowGoalAttachment) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowValidateParamsWorkflowGoalAttachment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowValidateParamsWorkflowGoalAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A map of workflow settings.
type WorkflowValidateParamsWorkflowSettings struct {
	// Whether the workflow is commercial. Defaults to false.
	IsCommercial param.Opt[bool] `json:"is_commercial,omitzero"`
	// Whether to ignore recipient preferences for a given type of notification. If
	// true, will send for every channel in the workflow even if the recipient has
	// opted out of a certain kind. Defaults to false.
	OverridePreferences param.Opt[bool] `json:"override_preferences,omitzero"`
	paramObj
}

func (r WorkflowValidateParamsWorkflowSettings) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowValidateParamsWorkflowSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowValidateParamsWorkflowSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
