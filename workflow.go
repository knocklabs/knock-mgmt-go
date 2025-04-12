// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

// WorkflowService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	Options []option.RequestOption
	Steps   WorkflowStepService
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
func (r *WorkflowService) Get(ctx context.Context, workflowKey string, query WorkflowGetParams, opts ...option.RequestOption) (res *Workflow, err error) {
	opts = append(r.Options[:], opts...)
	if workflowKey == "" {
		err = errors.New("missing required workflow_key parameter")
		return
	}
	path := fmt.Sprintf("v1/workflows/%s", workflowKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a paginated list of workflows available in a given environment. The
// workflows are returned alphabetically by `key`.
func (r *WorkflowService) List(ctx context.Context, query WorkflowListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Workflow], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	if workflowKey == "" {
		err = errors.New("missing required workflow_key parameter")
		return
	}
	path := fmt.Sprintf("v1/workflows/%s/activate", workflowKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Runs the latest version of a committed workflow in a given environment using the
// params provided.
func (r *WorkflowService) Run(ctx context.Context, workflowKey string, params WorkflowRunParams, opts ...option.RequestOption) (res *WorkflowRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	if workflowKey == "" {
		err = errors.New("missing required workflow_key parameter")
		return
	}
	path := fmt.Sprintf("v1/workflows/%s/run", workflowKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Updates a workflow of a given key, or creates a new one if it does not yet
// exist.
//
// Note: this endpoint only operates on workflows in the `development` environment.
func (r *WorkflowService) Upsert(ctx context.Context, workflowKey string, params WorkflowUpsertParams, opts ...option.RequestOption) (res *WorkflowUpsertResponse, err error) {
	opts = append(r.Options[:], opts...)
	if workflowKey == "" {
		err = errors.New("missing required workflow_key parameter")
		return
	}
	path := fmt.Sprintf("v1/workflows/%s", workflowKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Validates a workflow payload without persisting it. Some read-only fields will
// be empty as they are generated by the system when persisted.
//
// Note: Validating a workflow is only done in the development environment context.
func (r *WorkflowService) Validate(ctx context.Context, workflowKey string, params WorkflowValidateParams, opts ...option.RequestOption) (res *WorkflowValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if workflowKey == "" {
		err = errors.New("missing required workflow_key parameter")
		return
	}
	path := fmt.Sprintf("v1/workflows/%s/validate", workflowKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// A condition to be evaluated.
type Condition struct {
	// The operator to use in the evaluation of the condition.
	//
	// Any of "equal_to", "not_equal_to", "greater_than", "less_than",
	// "greater_than_or_equal_to", "less_than_or_equal_to", "contains", "not_contains",
	// "contains_all", "empty", "not_empty", "is_audience_member",
	// "is_not_audience_member".
	Operator ConditionOperator `json:"operator,required"`
	// The variable to be evaluated. Variables can be either static values or dynamic
	// properties. Static values will always be JSON decoded so will support strings,
	// lists, objects, numbers, and booleans. Dynamic values should be path
	// expressions.
	Variable string `json:"variable,required"`
	// The argument to be evaluated. Arguments can be either static values or dynamic
	// properties. Static values will always be JSON decoded so will support strings,
	// lists, objects, numbers, and booleans. Dynamic values should be path
	// expressions.
	Argument string `json:"argument,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Operator    resp.Field
		Variable    resp.Field
		Argument    resp.Field
		ExtraFields map[string]resp.Field
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
// ConditionParam.IsOverridden()
func (r Condition) ToParam() ConditionParam {
	return param.OverrideObj[ConditionParam](r.RawJSON())
}

// The operator to use in the evaluation of the condition.
type ConditionOperator string

const (
	ConditionOperatorEqualTo              ConditionOperator = "equal_to"
	ConditionOperatorNotEqualTo           ConditionOperator = "not_equal_to"
	ConditionOperatorGreaterThan          ConditionOperator = "greater_than"
	ConditionOperatorLessThan             ConditionOperator = "less_than"
	ConditionOperatorGreaterThanOrEqualTo ConditionOperator = "greater_than_or_equal_to"
	ConditionOperatorLessThanOrEqualTo    ConditionOperator = "less_than_or_equal_to"
	ConditionOperatorContains             ConditionOperator = "contains"
	ConditionOperatorNotContains          ConditionOperator = "not_contains"
	ConditionOperatorContainsAll          ConditionOperator = "contains_all"
	ConditionOperatorEmpty                ConditionOperator = "empty"
	ConditionOperatorNotEmpty             ConditionOperator = "not_empty"
	ConditionOperatorIsAudienceMember     ConditionOperator = "is_audience_member"
	ConditionOperatorIsNotAudienceMember  ConditionOperator = "is_not_audience_member"
)

// A condition to be evaluated.
//
// The properties Operator, Variable are required.
type ConditionParam struct {
	// The operator to use in the evaluation of the condition.
	//
	// Any of "equal_to", "not_equal_to", "greater_than", "less_than",
	// "greater_than_or_equal_to", "less_than_or_equal_to", "contains", "not_contains",
	// "contains_all", "empty", "not_empty", "is_audience_member",
	// "is_not_audience_member".
	Operator ConditionOperator `json:"operator,omitzero,required"`
	// The variable to be evaluated. Variables can be either static values or dynamic
	// properties. Static values will always be JSON decoded so will support strings,
	// lists, objects, numbers, and booleans. Dynamic values should be path
	// expressions.
	Variable string `json:"variable,required"`
	// The argument to be evaluated. Arguments can be either static values or dynamic
	// properties. Static values will always be JSON decoded so will support strings,
	// lists, objects, numbers, and booleans. Dynamic values should be path
	// expressions.
	Argument param.Opt[string] `json:"argument,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ConditionParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r ConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow ConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
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
		All resp.Field
		Any resp.Field
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
// ConditionGroupUnionParam.IsOverridden()
func (r ConditionGroupUnion) ToParam() ConditionGroupUnionParam {
	return param.OverrideObj[ConditionGroupUnionParam](r.RawJSON())
}

// A group of conditions that must all be met.
type ConditionGroupConditionGroupAllMatch struct {
	// A list of conditions.
	All []Condition `json:"all"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		All         resp.Field
		ExtraFields map[string]resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Any         resp.Field
		ExtraFields map[string]resp.Field
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
		Operator resp.Field
		Variable resp.Field
		Argument resp.Field
		All      resp.Field
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		All         resp.Field
		ExtraFields map[string]resp.Field
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u ConditionGroupUnionParam) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u ConditionGroupUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ConditionGroupUnionParam](u.OfConditionGroupAllMatch, u.OfConditionGroupAnyMatch)
}

func (u *ConditionGroupUnionParam) asAny() any {
	if !param.IsOmitted(u.OfConditionGroupAllMatch) {
		return u.OfConditionGroupAllMatch
	} else if !param.IsOmitted(u.OfConditionGroupAnyMatch) {
		return u.OfConditionGroupAnyMatch
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConditionGroupUnionParam) GetAll() []ConditionParam {
	if vt := u.OfConditionGroupAllMatch; vt != nil {
		return vt.All
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConditionGroupUnionParam) GetAny() []ConditionGroupConditionGroupAnyMatchAnyUnionParam {
	if vt := u.OfConditionGroupAnyMatch; vt != nil {
		return vt.Any
	}
	return nil
}

// A group of conditions that must all be met.
type ConditionGroupConditionGroupAllMatchParam struct {
	// A list of conditions.
	All []ConditionParam `json:"all,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ConditionGroupConditionGroupAllMatchParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ConditionGroupConditionGroupAllMatchParam) MarshalJSON() (data []byte, err error) {
	type shadow ConditionGroupConditionGroupAllMatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A group of conditions that any must be met. Can contain nested alls.
type ConditionGroupConditionGroupAnyMatchParam struct {
	// An array of conditions or nested condition groups to evaluate.
	Any []ConditionGroupConditionGroupAnyMatchAnyUnionParam `json:"any,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ConditionGroupConditionGroupAnyMatchParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ConditionGroupConditionGroupAnyMatchParam) MarshalJSON() (data []byte, err error) {
	type shadow ConditionGroupConditionGroupAnyMatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConditionGroupConditionGroupAnyMatchAnyUnionParam struct {
	OfCondition              *ConditionParam                                                     `json:",omitzero,inline"`
	OfConditionGroupAllMatch *ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatchParam `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u ConditionGroupConditionGroupAnyMatchAnyUnionParam) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u ConditionGroupConditionGroupAnyMatchAnyUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ConditionGroupConditionGroupAnyMatchAnyUnionParam](u.OfCondition, u.OfConditionGroupAllMatch)
}

func (u *ConditionGroupConditionGroupAnyMatchAnyUnionParam) asAny() any {
	if !param.IsOmitted(u.OfCondition) {
		return u.OfCondition
	} else if !param.IsOmitted(u.OfConditionGroupAllMatch) {
		return u.OfConditionGroupAllMatch
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConditionGroupConditionGroupAnyMatchAnyUnionParam) GetOperator() *string {
	if vt := u.OfCondition; vt != nil {
		return (*string)(&vt.Operator)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConditionGroupConditionGroupAnyMatchAnyUnionParam) GetVariable() *string {
	if vt := u.OfCondition; vt != nil {
		return &vt.Variable
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConditionGroupConditionGroupAnyMatchAnyUnionParam) GetArgument() *string {
	if vt := u.OfCondition; vt != nil && vt.Argument.IsPresent() {
		return &vt.Argument.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConditionGroupConditionGroupAnyMatchAnyUnionParam) GetAll() []ConditionParam {
	if vt := u.OfConditionGroupAllMatch; vt != nil {
		return vt.All
	}
	return nil
}

// A group of conditions that must all be met.
type ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatchParam struct {
	// A list of conditions.
	All []ConditionParam `json:"all,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatchParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatchParam) MarshalJSON() (data []byte, err error) {
	type shadow ConditionGroupConditionGroupAnyMatchAnyConditionGroupAllMatchParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A duration of time, represented as a unit and a value.
type Duration struct {
	// The unit of time.
	//
	// Any of "minutes", "hours", "days", "weeks", "months".
	Unit DurationUnit `json:"unit,required"`
	// The value of the duration.
	Value int64 `json:"value,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Unit        resp.Field
		Value       resp.Field
		ExtraFields map[string]resp.Field
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
// DurationParam.IsOverridden()
func (r Duration) ToParam() DurationParam {
	return param.OverrideObj[DurationParam](r.RawJSON())
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
	Unit DurationUnit `json:"unit,omitzero,required"`
	// The value of the duration.
	Value int64 `json:"value,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f DurationParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r DurationParam) MarshalJSON() (data []byte, err error) {
	type shadow DurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A send window time for a notification. Describes a single day.
type SendWindow struct {
	// The day of the week.
	//
	// Any of "monday", "tuesday", "wednesday", "thursday", "friday", "saturday",
	// "sunday".
	Day SendWindowDay `json:"day,required"`
	// The type of send window.
	//
	// Any of "send", "do_not_send".
	Type SendWindowType `json:"type,required"`
	// The start time of the send window.
	From string `json:"from,nullable" format:"time"`
	// The end time of the send window.
	Until string `json:"until,nullable" format:"time"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Day         resp.Field
		Type        resp.Field
		From        resp.Field
		Until       resp.Field
		ExtraFields map[string]resp.Field
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
// SendWindowParam.IsOverridden()
func (r SendWindow) ToParam() SendWindowParam {
	return param.OverrideObj[SendWindowParam](r.RawJSON())
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
	Day SendWindowDay `json:"day,omitzero,required"`
	// The type of send window.
	//
	// Any of "send", "do_not_send".
	Type SendWindowType `json:"type,omitzero,required"`
	// The start time of the send window.
	From param.Opt[string] `json:"from,omitzero" format:"time"`
	// The end time of the send window.
	Until param.Opt[string] `json:"until,omitzero" format:"time"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f SendWindowParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r SendWindowParam) MarshalJSON() (data []byte, err error) {
	type shadow SendWindowParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A workflow object.
type Workflow struct {
	// Whether the workflow is
	// [active](https://docs.knock.app/concepts/workflows#workflow-status) in the
	// current environment. (read-only).
	Active bool `json:"active,required"`
	// The timestamp of when the workflow was created. (read-only).
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The slug of the environment in which the workflow exists. (read-only).
	Environment string `json:"environment,required"`
	// The unique key string for the workflow object. Must be at minimum 3 characters
	// and at maximum 255 characters in length. Must be in the format of ^[a-z0-9_-]+$.
	Key string `json:"key,required"`
	// A name for the workflow. Must be at maximum 255 characters in length.
	Name string `json:"name,required"`
	// The SHA hash of the workflow data. (read-only).
	Sha string `json:"sha,required"`
	// A list of workflow step objects in the workflow.
	Steps []WorkflowStepUnion `json:"steps,required"`
	// The timestamp of when the workflow was last updated. (read-only).
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Whether the workflow and its steps are in a valid state. (read-only).
	Valid bool `json:"valid,required"`
	// A list of
	// [categories](https://docs.knock.app/concepts/workflows#workflow-categories) that
	// the workflow belongs to.
	Categories []string `json:"categories"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions,nullable"`
	// The timestamp of when the workflow was deleted. (read-only).
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// An arbitrary string attached to a workflow object. Useful for adding notes about
	// the workflow for internal purposes. Maximum of 280 characters allowed.
	Description string `json:"description"`
	// A map of workflow settings.
	Settings WorkflowSettings `json:"settings"`
	// A JSON schema for the expected structure of the workflow trigger's data payload.
	// Used to validate trigger requests. Read more in the
	// [docs](https://docs.knock.app/developer-tools/validating-trigger-data).
	TriggerDataJsonSchema map[string]interface{} `json:"trigger_data_json_schema"`
	// The frequency at which the workflow should be triggered. One of:
	// `once_per_recipient`, `once_per_recipient_per_tenant`, `every_trigger`. Defaults
	// to `every_trigger`. Read more in
	// [docs](https://docs.knock.app/send-notifications/triggering-workflows/overview#controlling-workflow-trigger-frequency).
	//
	// Any of "every_trigger", "once_per_recipient", "once_per_recipient_per_tenant".
	TriggerFrequency WorkflowTriggerFrequency `json:"trigger_frequency"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Active                resp.Field
		CreatedAt             resp.Field
		Environment           resp.Field
		Key                   resp.Field
		Name                  resp.Field
		Sha                   resp.Field
		Steps                 resp.Field
		UpdatedAt             resp.Field
		Valid                 resp.Field
		Categories            resp.Field
		Conditions            resp.Field
		DeletedAt             resp.Field
		Description           resp.Field
		Settings              resp.Field
		TriggerDataJsonSchema resp.Field
		TriggerFrequency      resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Workflow) RawJSON() string { return r.JSON.raw }
func (r *Workflow) UnmarshalJSON(data []byte) error {
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IsCommercial        resp.Field
		OverridePreferences resp.Field
		ExtraFields         map[string]resp.Field
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

// A batch function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/batch-function).
type WorkflowBatchStep struct {
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description,required"`
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// The settings for the batch step.
	Settings WorkflowBatchStepSettings `json:"settings,required"`
	// The type of the workflow step.
	//
	// Any of "batch".
	Type WorkflowBatchStepType `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Description resp.Field
		Name        resp.Field
		Ref         resp.Field
		Settings    resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
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
// WorkflowBatchStepParam.IsOverridden()
func (r WorkflowBatchStep) ToParam() WorkflowBatchStepParam {
	return param.OverrideObj[WorkflowBatchStepParam](r.RawJSON())
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
	BatchExecutionMode string `json:"batch_execution_mode,nullable"`
	// The maximum number of batch items allowed in a batch. Between: 2 and 1000.
	BatchItemsMaxLimit int64 `json:"batch_items_max_limit,nullable"`
	// The maximum number of batch items allowed to be rendered into a template.
	// Between: 1 and 100. Defaults to 10.
	BatchItemsRenderLimit int64 `json:"batch_items_render_limit,nullable"`
	// The data property to use to batch notifications per recipient.
	BatchKey string `json:"batch_key,nullable"`
	// The order describing whether to return the first or last ten batch items in the
	// activities variable. One of: `asc` or `desc`.
	//
	// Any of "asc", "desc".
	BatchOrder string `json:"batch_order,nullable"`
	// The data path to resolve the batch window. The resolved value must be an
	// ISO-8601 timestamp.
	BatchUntilFieldPath string `json:"batch_until_field_path,nullable"`
	// A duration of time, represented as a unit and a value.
	BatchWindow Duration `json:"batch_window,nullable"`
	// A duration of time, represented as a unit and a value.
	BatchWindowExtensionLimit Duration `json:"batch_window_extension_limit,nullable"`
	// The type of the batch window used. One of: `fixed` or `sliding`.
	//
	// Any of "fixed", "sliding".
	BatchWindowType string `json:"batch_window_type,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		BatchExecutionMode        resp.Field
		BatchItemsMaxLimit        resp.Field
		BatchItemsRenderLimit     resp.Field
		BatchKey                  resp.Field
		BatchOrder                resp.Field
		BatchUntilFieldPath       resp.Field
		BatchWindow               resp.Field
		BatchWindowExtensionLimit resp.Field
		BatchWindowType           resp.Field
		ExtraFields               map[string]resp.Field
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
// The properties Description, Name, Ref, Settings, Type are required.
type WorkflowBatchStepParam struct {
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero,required"`
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// The settings for the batch step.
	Settings WorkflowBatchStepSettingsParam `json:"settings,omitzero,required"`
	// The type of the workflow step.
	//
	// Any of "batch".
	Type WorkflowBatchStepType `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowBatchStepParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowBatchStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowBatchStepParam
	return param.MarshalObject(r, (*shadow)(&r))
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
	// A duration of time, represented as a unit and a value.
	BatchWindow DurationParam `json:"batch_window,omitzero"`
	// A duration of time, represented as a unit and a value.
	BatchWindowExtensionLimit DurationParam `json:"batch_window_extension_limit,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowBatchStepSettingsParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowBatchStepSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowBatchStepSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[WorkflowBatchStepSettingsParam](
		"BatchExecutionMode", true, "accumulate", "flush_leading",
	)
	apijson.RegisterFieldValidator[WorkflowBatchStepSettingsParam](
		"BatchOrder", true, "asc", "desc",
	)
	apijson.RegisterFieldValidator[WorkflowBatchStepSettingsParam](
		"BatchWindowType", true, "fixed", "sliding",
	)
}

// A branch function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/branch-function).
type WorkflowBranchStep struct {
	// A list of workflow branches to be evaluated.
	Branches []WorkflowBranchStepBranch `json:"branches,required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description,required"`
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// The type of step.
	//
	// Any of "branch".
	Type WorkflowBranchStepType `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Branches    resp.Field
		Description resp.Field
		Name        resp.Field
		Ref         resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
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
// WorkflowBranchStepParam.IsOverridden()
func (r WorkflowBranchStep) ToParam() WorkflowBranchStepParam {
	return param.OverrideObj[WorkflowBranchStepParam](r.RawJSON())
}

// A branch in a branch step.
type WorkflowBranchStepBranch struct {
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions,nullable"`
	// The name of the branch.
	Name string `json:"name"`
	// A list of steps that will be executed if the branch is chosen.
	Steps []WorkflowStepUnion `json:"steps"`
	// If the workflow should halt at the end of the branch.
	Terminates bool `json:"terminates"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Conditions  resp.Field
		Name        resp.Field
		Steps       resp.Field
		Terminates  resp.Field
		ExtraFields map[string]resp.Field
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
// The properties Branches, Description, Name, Ref, Type are required.
type WorkflowBranchStepParam struct {
	// A list of workflow branches to be evaluated.
	Branches []WorkflowBranchStepBranchParam `json:"branches,omitzero,required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description,required"`
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// The type of step.
	//
	// Any of "branch".
	Type WorkflowBranchStepType `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowBranchStepParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowBranchStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowBranchStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A branch in a branch step.
type WorkflowBranchStepBranchParam struct {
	// The name of the branch.
	Name param.Opt[string] `json:"name,omitzero"`
	// If the workflow should halt at the end of the branch.
	Terminates param.Opt[bool] `json:"terminates,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	// A list of steps that will be executed if the branch is chosen.
	Steps []WorkflowStepUnionParam `json:"steps,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowBranchStepBranchParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowBranchStepBranchParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowBranchStepBranchParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A channel step within a workflow. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
type WorkflowChannelStep struct {
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// The message template for the channel step. The shape of the template depends on
	// the type of the channel you'll be sending to. See below for definitions of each
	// channel type template: email, in-app, SMS, push, chat, and webhook.
	Template WorkflowChannelStepTemplateUnion `json:"template,required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowChannelStepType `json:"type,required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. A channel step can have either a channel key or a channel group
	// key, but not both.
	ChannelGroupKey string `json:"channel_group_key,nullable"`
	// The key of the channel to which the channel step will be sending a notification.
	// A channel step can have either a channel key or a channel group key, but not
	// both.
	ChannelKey string `json:"channel_key,nullable"`
	// A map of channel overrides for the channel step.
	ChannelOverrides WorkflowChannelStepChannelOverridesUnion `json:"channel_overrides,nullable"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions,nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description,nullable"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindow `json:"send_windows,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name             resp.Field
		Ref              resp.Field
		Template         resp.Field
		Type             resp.Field
		ChannelGroupKey  resp.Field
		ChannelKey       resp.Field
		ChannelOverrides resp.Field
		Conditions       resp.Field
		Description      resp.Field
		SendWindows      resp.Field
		ExtraFields      map[string]resp.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowChannelStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowChannelStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowChannelStep to a WorkflowChannelStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowChannelStepParam.IsOverridden()
func (r WorkflowChannelStep) ToParam() WorkflowChannelStepParam {
	return param.OverrideObj[WorkflowChannelStepParam](r.RawJSON())
}

// WorkflowChannelStepTemplateUnion contains all possible properties and values
// from [EmailTemplate], [InAppFeedTemplate], [SMSTemplate], [PushTemplate],
// [ChatTemplate], [WebhookTemplate].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkflowChannelStepTemplateUnion struct {
	// This field is from variant [EmailTemplate].
	Subject string `json:"subject"`
	// This field is from variant [EmailTemplate].
	HTMLBody string `json:"html_body"`
	// This field is a union of [EmailTemplateSettings], [SMSTemplateSettings],
	// [PushTemplateSettings]
	Settings WorkflowChannelStepTemplateUnionSettings `json:"settings"`
	TextBody string                                   `json:"text_body"`
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
	// This field is from variant [WebhookTemplate].
	Method WebhookTemplateMethod `json:"method"`
	// This field is from variant [WebhookTemplate].
	URL string `json:"url"`
	// This field is from variant [WebhookTemplate].
	Body string `json:"body"`
	// This field is from variant [WebhookTemplate].
	Headers []WebhookTemplateHeader `json:"headers"`
	// This field is from variant [WebhookTemplate].
	QueryParams []WebhookTemplateQueryParam `json:"query_params"`
	JSON        struct {
		Subject       resp.Field
		HTMLBody      resp.Field
		Settings      resp.Field
		TextBody      resp.Field
		VisualBlocks  resp.Field
		MarkdownBody  resp.Field
		ActionButtons resp.Field
		ActionURL     resp.Field
		Title         resp.Field
		JsonBody      resp.Field
		Summary       resp.Field
		Method        resp.Field
		URL           resp.Field
		Body          resp.Field
		Headers       resp.Field
		QueryParams   resp.Field
		raw           string
	} `json:"-"`
}

func (u WorkflowChannelStepTemplateUnion) AsEmailTemplate() (v EmailTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowChannelStepTemplateUnion) AsInAppFeedTemplate() (v InAppFeedTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowChannelStepTemplateUnion) AsSMSTemplate() (v SMSTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowChannelStepTemplateUnion) AsPushTemplate() (v PushTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowChannelStepTemplateUnion) AsChatTemplate() (v ChatTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowChannelStepTemplateUnion) AsWebhookTemplate() (v WebhookTemplate) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowChannelStepTemplateUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkflowChannelStepTemplateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowChannelStepTemplateUnionSettings is an implicit subunion of
// [WorkflowChannelStepTemplateUnion]. WorkflowChannelStepTemplateUnionSettings
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowChannelStepTemplateUnion].
type WorkflowChannelStepTemplateUnionSettings struct {
	// This field is from variant [EmailTemplateSettings].
	AttachmentKey string `json:"attachment_key"`
	// This field is from variant [EmailTemplateSettings].
	LayoutKey string `json:"layout_key"`
	// This field is from variant [EmailTemplateSettings].
	PreContent       string `json:"pre_content"`
	PayloadOverrides string `json:"payload_overrides"`
	// This field is from variant [SMSTemplateSettings].
	ToNumber string `json:"to_number"`
	// This field is from variant [PushTemplateSettings].
	DeliveryType string `json:"delivery_type"`
	JSON         struct {
		AttachmentKey    resp.Field
		LayoutKey        resp.Field
		PreContent       resp.Field
		PayloadOverrides resp.Field
		ToNumber         resp.Field
		DeliveryType     resp.Field
		raw              string
	} `json:"-"`
}

func (r *WorkflowChannelStepTemplateUnionSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the workflow step.
type WorkflowChannelStepType string

const (
	WorkflowChannelStepTypeChannel WorkflowChannelStepType = "channel"
)

// WorkflowChannelStepChannelOverridesUnion contains all possible properties and
// values from [EmailChannelSettings], [InAppFeedChannelSettings],
// [SMSChannelSettings], [PushChannelSettings], [ChatChannelSettings].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkflowChannelStepChannelOverridesUnion struct {
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
	LinkTracking  bool   `json:"link_tracking"`
	// This field is from variant [EmailChannelSettings].
	OpenTracking bool `json:"open_tracking"`
	// This field is from variant [EmailChannelSettings].
	ReplyToAddress string `json:"reply_to_address"`
	// This field is from variant [EmailChannelSettings].
	ToAddress string `json:"to_address"`
	// This field is from variant [PushChannelSettings].
	TokenDeregistration bool `json:"token_deregistration"`
	// This field is from variant [ChatChannelSettings].
	EmailBasedUserIDResolution bool `json:"email_based_user_id_resolution"`
	JSON                       struct {
		BccAddress                 resp.Field
		CcAddress                  resp.Field
		FromAddress                resp.Field
		FromName                   resp.Field
		JsonOverrides              resp.Field
		LinkTracking               resp.Field
		OpenTracking               resp.Field
		ReplyToAddress             resp.Field
		ToAddress                  resp.Field
		TokenDeregistration        resp.Field
		EmailBasedUserIDResolution resp.Field
		raw                        string
	} `json:"-"`
}

func (u WorkflowChannelStepChannelOverridesUnion) AsEmailChannelSettings() (v EmailChannelSettings) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowChannelStepChannelOverridesUnion) AsInAppFeedChannelSettings() (v InAppFeedChannelSettings) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowChannelStepChannelOverridesUnion) AsSMSChannelSettings() (v SMSChannelSettings) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowChannelStepChannelOverridesUnion) AsPushChannelSettings() (v PushChannelSettings) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowChannelStepChannelOverridesUnion) AsChatChannelSettings() (v ChatChannelSettings) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowChannelStepChannelOverridesUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkflowChannelStepChannelOverridesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A channel step within a workflow. Read more in the
// [docs](https://docs.knock.app/designing-workflows/channel-step).
//
// The properties Name, Ref, Template, Type are required.
type WorkflowChannelStepParam struct {
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// The message template for the channel step. The shape of the template depends on
	// the type of the channel you'll be sending to. See below for definitions of each
	// channel type template: email, in-app, SMS, push, chat, and webhook.
	Template WorkflowChannelStepTemplateUnionParam `json:"template,omitzero,required"`
	// The type of the workflow step.
	//
	// Any of "channel".
	Type WorkflowChannelStepType `json:"type,omitzero,required"`
	// The key of the channel group to which the channel step will be sending a
	// notification. A channel step can have either a channel key or a channel group
	// key, but not both.
	ChannelGroupKey param.Opt[string] `json:"channel_group_key,omitzero"`
	// The key of the channel to which the channel step will be sending a notification.
	// A channel step can have either a channel key or a channel group key, but not
	// both.
	ChannelKey param.Opt[string] `json:"channel_key,omitzero"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A map of channel overrides for the channel step.
	ChannelOverrides WorkflowChannelStepChannelOverridesUnionParam `json:"channel_overrides,omitzero"`
	// A list of send window objects. Must include one send window object per day of
	// the week.
	SendWindows []SendWindowParam `json:"send_windows,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowChannelStepParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowChannelStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowChannelStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowChannelStepTemplateUnionParam struct {
	OfEmailTemplate     *EmailTemplateParam     `json:",omitzero,inline"`
	OfInAppFeedTemplate *InAppFeedTemplateParam `json:",omitzero,inline"`
	OfSMSTemplate       *SMSTemplateParam       `json:",omitzero,inline"`
	OfPushTemplate      *PushTemplateParam      `json:",omitzero,inline"`
	OfChatTemplate      *ChatTemplateParam      `json:",omitzero,inline"`
	OfWebhookTemplate   *WebhookTemplateParam   `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u WorkflowChannelStepTemplateUnionParam) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u WorkflowChannelStepTemplateUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[WorkflowChannelStepTemplateUnionParam](u.OfEmailTemplate,
		u.OfInAppFeedTemplate,
		u.OfSMSTemplate,
		u.OfPushTemplate,
		u.OfChatTemplate,
		u.OfWebhookTemplate)
}

func (u *WorkflowChannelStepTemplateUnionParam) asAny() any {
	if !param.IsOmitted(u.OfEmailTemplate) {
		return u.OfEmailTemplate
	} else if !param.IsOmitted(u.OfInAppFeedTemplate) {
		return u.OfInAppFeedTemplate
	} else if !param.IsOmitted(u.OfSMSTemplate) {
		return u.OfSMSTemplate
	} else if !param.IsOmitted(u.OfPushTemplate) {
		return u.OfPushTemplate
	} else if !param.IsOmitted(u.OfChatTemplate) {
		return u.OfChatTemplate
	} else if !param.IsOmitted(u.OfWebhookTemplate) {
		return u.OfWebhookTemplate
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetSubject() *string {
	if vt := u.OfEmailTemplate; vt != nil {
		return &vt.Subject
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetHTMLBody() *string {
	if vt := u.OfEmailTemplate; vt != nil && vt.HTMLBody.IsPresent() {
		return &vt.HTMLBody.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetVisualBlocks() []EmailTemplateVisualBlockUnionParam {
	if vt := u.OfEmailTemplate; vt != nil {
		return vt.VisualBlocks
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetActionButtons() []InAppFeedTemplateActionButtonParam {
	if vt := u.OfInAppFeedTemplate; vt != nil {
		return vt.ActionButtons
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetActionURL() *string {
	if vt := u.OfInAppFeedTemplate; vt != nil && vt.ActionURL.IsPresent() {
		return &vt.ActionURL.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetTitle() *string {
	if vt := u.OfPushTemplate; vt != nil {
		return &vt.Title
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetJsonBody() *string {
	if vt := u.OfChatTemplate; vt != nil && vt.JsonBody.IsPresent() {
		return &vt.JsonBody.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetSummary() *string {
	if vt := u.OfChatTemplate; vt != nil && vt.Summary.IsPresent() {
		return &vt.Summary.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetMethod() *string {
	if vt := u.OfWebhookTemplate; vt != nil {
		return (*string)(&vt.Method)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetURL() *string {
	if vt := u.OfWebhookTemplate; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetBody() *string {
	if vt := u.OfWebhookTemplate; vt != nil && vt.Body.IsPresent() {
		return &vt.Body.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetHeaders() []WebhookTemplateHeaderParam {
	if vt := u.OfWebhookTemplate; vt != nil {
		return vt.Headers
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetQueryParams() []WebhookTemplateQueryParamParam {
	if vt := u.OfWebhookTemplate; vt != nil {
		return vt.QueryParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetTextBody() *string {
	if vt := u.OfEmailTemplate; vt != nil && vt.TextBody.IsPresent() {
		return &vt.TextBody.Value
	} else if vt := u.OfSMSTemplate; vt != nil {
		return (*string)(&vt.TextBody)
	} else if vt := u.OfPushTemplate; vt != nil {
		return (*string)(&vt.TextBody)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepTemplateUnionParam) GetMarkdownBody() *string {
	if vt := u.OfInAppFeedTemplate; vt != nil {
		return (*string)(&vt.MarkdownBody)
	} else if vt := u.OfChatTemplate; vt != nil {
		return (*string)(&vt.MarkdownBody)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u WorkflowChannelStepTemplateUnionParam) GetSettings() (res workflowChannelStepTemplateUnionParamSettings) {
	if vt := u.OfEmailTemplate; vt != nil {
		res.ofEmailTemplateSettings = &vt.Settings
	} else if vt := u.OfSMSTemplate; vt != nil {
		res.ofSMSTemplateSettings = &vt.Settings
	} else if vt := u.OfPushTemplate; vt != nil {
		res.ofPushTemplateSettings = &vt.Settings
	}
	return
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type workflowChannelStepTemplateUnionParamSettings struct {
	ofEmailTemplateSettings *EmailTemplateSettingsParam
	ofSMSTemplateSettings   *SMSTemplateSettingsParam
	ofPushTemplateSettings  *PushTemplateSettingsParam
}

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.EmailTemplateSettingsParam:
//	case *knockmapi.SMSTemplateSettingsParam:
//	case *knockmapi.PushTemplateSettingsParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowChannelStepTemplateUnionParamSettings) AsAny() any {
	if !param.IsOmitted(u.ofEmailTemplateSettings) {
		return u.ofEmailTemplateSettings
	} else if !param.IsOmitted(u.ofSMSTemplateSettings) {
		return u.ofSMSTemplateSettings
	} else if !param.IsOmitted(u.ofPushTemplateSettings) {
		return u.ofPushTemplateSettings
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowChannelStepTemplateUnionParamSettings) GetAttachmentKey() *string {
	if vt := u.ofEmailTemplateSettings; vt != nil && vt.AttachmentKey.IsPresent() {
		return &vt.AttachmentKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowChannelStepTemplateUnionParamSettings) GetLayoutKey() *string {
	if vt := u.ofEmailTemplateSettings; vt != nil && vt.LayoutKey.IsPresent() {
		return &vt.LayoutKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowChannelStepTemplateUnionParamSettings) GetPreContent() *string {
	if vt := u.ofEmailTemplateSettings; vt != nil && vt.PreContent.IsPresent() {
		return &vt.PreContent.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowChannelStepTemplateUnionParamSettings) GetToNumber() *string {
	if vt := u.ofSMSTemplateSettings; vt != nil && vt.ToNumber.IsPresent() {
		return &vt.ToNumber.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowChannelStepTemplateUnionParamSettings) GetDeliveryType() *string {
	if vt := u.ofPushTemplateSettings; vt != nil {
		return &vt.DeliveryType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowChannelStepTemplateUnionParamSettings) GetPayloadOverrides() *string {
	if vt := u.ofSMSTemplateSettings; vt != nil && vt.PayloadOverrides.IsPresent() {
		return &vt.PayloadOverrides.Value
	} else if vt := u.ofPushTemplateSettings; vt != nil && vt.PayloadOverrides.IsPresent() {
		return &vt.PayloadOverrides.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowChannelStepChannelOverridesUnionParam struct {
	OfEmailChannelSettings     *EmailChannelSettingsParam     `json:",omitzero,inline"`
	OfInAppFeedChannelSettings *InAppFeedChannelSettingsParam `json:",omitzero,inline"`
	OfSMSChannelSettings       *SMSChannelSettingsParam       `json:",omitzero,inline"`
	OfPushChannelSettings      *PushChannelSettingsParam      `json:",omitzero,inline"`
	OfChatChannelSettings      *ChatChannelSettingsParam      `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u WorkflowChannelStepChannelOverridesUnionParam) IsPresent() bool {
	return !param.IsOmitted(u) && !u.IsNull()
}
func (u WorkflowChannelStepChannelOverridesUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[WorkflowChannelStepChannelOverridesUnionParam](u.OfEmailChannelSettings,
		u.OfInAppFeedChannelSettings,
		u.OfSMSChannelSettings,
		u.OfPushChannelSettings,
		u.OfChatChannelSettings)
}

func (u *WorkflowChannelStepChannelOverridesUnionParam) asAny() any {
	if !param.IsOmitted(u.OfEmailChannelSettings) {
		return u.OfEmailChannelSettings
	} else if !param.IsOmitted(u.OfInAppFeedChannelSettings) {
		return u.OfInAppFeedChannelSettings
	} else if !param.IsOmitted(u.OfSMSChannelSettings) {
		return u.OfSMSChannelSettings
	} else if !param.IsOmitted(u.OfPushChannelSettings) {
		return u.OfPushChannelSettings
	} else if !param.IsOmitted(u.OfChatChannelSettings) {
		return u.OfChatChannelSettings
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepChannelOverridesUnionParam) GetBccAddress() *string {
	if vt := u.OfEmailChannelSettings; vt != nil && vt.BccAddress.IsPresent() {
		return &vt.BccAddress.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepChannelOverridesUnionParam) GetCcAddress() *string {
	if vt := u.OfEmailChannelSettings; vt != nil && vt.CcAddress.IsPresent() {
		return &vt.CcAddress.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepChannelOverridesUnionParam) GetFromAddress() *string {
	if vt := u.OfEmailChannelSettings; vt != nil && vt.FromAddress.IsPresent() {
		return &vt.FromAddress.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepChannelOverridesUnionParam) GetFromName() *string {
	if vt := u.OfEmailChannelSettings; vt != nil && vt.FromName.IsPresent() {
		return &vt.FromName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepChannelOverridesUnionParam) GetJsonOverrides() *string {
	if vt := u.OfEmailChannelSettings; vt != nil && vt.JsonOverrides.IsPresent() {
		return &vt.JsonOverrides.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepChannelOverridesUnionParam) GetOpenTracking() *bool {
	if vt := u.OfEmailChannelSettings; vt != nil && vt.OpenTracking.IsPresent() {
		return &vt.OpenTracking.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepChannelOverridesUnionParam) GetReplyToAddress() *string {
	if vt := u.OfEmailChannelSettings; vt != nil && vt.ReplyToAddress.IsPresent() {
		return &vt.ReplyToAddress.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepChannelOverridesUnionParam) GetToAddress() *string {
	if vt := u.OfEmailChannelSettings; vt != nil && vt.ToAddress.IsPresent() {
		return &vt.ToAddress.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepChannelOverridesUnionParam) GetTokenDeregistration() *bool {
	if vt := u.OfPushChannelSettings; vt != nil && vt.TokenDeregistration.IsPresent() {
		return &vt.TokenDeregistration.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepChannelOverridesUnionParam) GetEmailBasedUserIDResolution() *bool {
	if vt := u.OfChatChannelSettings; vt != nil && vt.EmailBasedUserIDResolution.IsPresent() {
		return &vt.EmailBasedUserIDResolution.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowChannelStepChannelOverridesUnionParam) GetLinkTracking() *bool {
	if vt := u.OfEmailChannelSettings; vt != nil && vt.LinkTracking.IsPresent() {
		return &vt.LinkTracking.Value
	} else if vt := u.OfInAppFeedChannelSettings; vt != nil && vt.LinkTracking.IsPresent() {
		return &vt.LinkTracking.Value
	} else if vt := u.OfSMSChannelSettings; vt != nil && vt.LinkTracking.IsPresent() {
		return &vt.LinkTracking.Value
	} else if vt := u.OfChatChannelSettings; vt != nil && vt.LinkTracking.IsPresent() {
		return &vt.LinkTracking.Value
	}
	return nil
}

// A delay function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/delay-function).
type WorkflowDelayStep struct {
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions,required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description,required"`
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// The settings for the delay step. Both fields can be set to compute a delay where
	// `delay_for` is an offset from the `delay_until_field_path`.
	Settings WorkflowDelayStepSettings `json:"settings,required"`
	// The type of the workflow step.
	//
	// Any of "delay".
	Type WorkflowDelayStepType `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Conditions  resp.Field
		Description resp.Field
		Name        resp.Field
		Ref         resp.Field
		Settings    resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
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
// WorkflowDelayStepParam.IsOverridden()
func (r WorkflowDelayStep) ToParam() WorkflowDelayStepParam {
	return param.OverrideObj[WorkflowDelayStepParam](r.RawJSON())
}

// The settings for the delay step. Both fields can be set to compute a delay where
// `delay_for` is an offset from the `delay_until_field_path`.
type WorkflowDelayStepSettings struct {
	// A duration of time, represented as a unit and a value.
	DelayFor Duration `json:"delay_for,nullable"`
	// When set will use the path to resolve the delay into a timestamp from the
	// property referenced
	DelayUntilFieldPath string `json:"delay_until_field_path"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		DelayFor            resp.Field
		DelayUntilFieldPath resp.Field
		ExtraFields         map[string]resp.Field
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
// The properties Conditions, Description, Name, Ref, Settings, Type are required.
type WorkflowDelayStepParam struct {
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero,required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero,required"`
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// The settings for the delay step. Both fields can be set to compute a delay where
	// `delay_for` is an offset from the `delay_until_field_path`.
	Settings WorkflowDelayStepSettingsParam `json:"settings,omitzero,required"`
	// The type of the workflow step.
	//
	// Any of "delay".
	Type WorkflowDelayStepType `json:"type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowDelayStepParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowDelayStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowDelayStepParam
	return param.MarshalObject(r, (*shadow)(&r))
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowDelayStepSettingsParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowDelayStepSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowDelayStepSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A fetch function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/fetch-function).
type WorkflowFetchStep struct {
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// A request template for a fetch function step.
	Settings RequestTemplate `json:"settings,required"`
	// The type of the workflow step.
	//
	// Any of "fetch".
	Type WorkflowFetchStepType `json:"type,required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions,nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		Ref         resp.Field
		Settings    resp.Field
		Type        resp.Field
		Conditions  resp.Field
		Description resp.Field
		ExtraFields map[string]resp.Field
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
// WorkflowFetchStepParam.IsOverridden()
func (r WorkflowFetchStep) ToParam() WorkflowFetchStepParam {
	return param.OverrideObj[WorkflowFetchStepParam](r.RawJSON())
}

// The type of the workflow step.
type WorkflowFetchStepType string

const (
	WorkflowFetchStepTypeFetch WorkflowFetchStepType = "fetch"
)

// A fetch function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/fetch-function).
//
// The properties Name, Ref, Settings, Type are required.
type WorkflowFetchStepParam struct {
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// A request template for a fetch function step.
	Settings RequestTemplateParam `json:"settings,omitzero,required"`
	// The type of the workflow step.
	//
	// Any of "fetch".
	Type WorkflowFetchStepType `json:"type,omitzero,required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowFetchStepParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowFetchStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowFetchStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// WorkflowStepUnion contains all possible properties and values from
// [WorkflowChannelStep], [WorkflowDelayStep], [WorkflowBatchStep],
// [WorkflowFetchStep], [WorkflowThrottleStep], [WorkflowBranchStep],
// [WorkflowTriggerWorkflowStep].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkflowStepUnion struct {
	Name string `json:"name"`
	Ref  string `json:"ref"`
	// This field is from variant [WorkflowChannelStep].
	Template WorkflowChannelStepTemplateUnion `json:"template"`
	Type     string                           `json:"type"`
	// This field is from variant [WorkflowChannelStep].
	ChannelGroupKey string `json:"channel_group_key"`
	// This field is from variant [WorkflowChannelStep].
	ChannelKey string `json:"channel_key"`
	// This field is from variant [WorkflowChannelStep].
	ChannelOverrides WorkflowChannelStepChannelOverridesUnion `json:"channel_overrides"`
	// This field is from variant [WorkflowChannelStep].
	Conditions  ConditionGroupUnion `json:"conditions"`
	Description string              `json:"description"`
	// This field is from variant [WorkflowChannelStep].
	SendWindows []SendWindow `json:"send_windows"`
	// This field is a union of [WorkflowDelayStepSettings],
	// [WorkflowBatchStepSettings], [RequestTemplate], [WorkflowThrottleStepSettings],
	// [WorkflowTriggerWorkflowStepSettings]
	Settings WorkflowStepUnionSettings `json:"settings"`
	// This field is from variant [WorkflowBranchStep].
	Branches []WorkflowBranchStepBranch `json:"branches"`
	JSON     struct {
		Name             resp.Field
		Ref              resp.Field
		Template         resp.Field
		Type             resp.Field
		ChannelGroupKey  resp.Field
		ChannelKey       resp.Field
		ChannelOverrides resp.Field
		Conditions       resp.Field
		Description      resp.Field
		SendWindows      resp.Field
		Settings         resp.Field
		Branches         resp.Field
		raw              string
	} `json:"-"`
}

func (u WorkflowStepUnion) AsWorkflowChannelStep() (v WorkflowChannelStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowDelayStep() (v WorkflowDelayStep) {
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

func (u WorkflowStepUnion) AsWorkflowThrottleStep() (v WorkflowThrottleStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowStepUnion) AsWorkflowBranchStep() (v WorkflowBranchStep) {
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

// WorkflowStepUnionSettings is an implicit subunion of [WorkflowStepUnion].
// WorkflowStepUnionSettings provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowStepUnion].
type WorkflowStepUnionSettings struct {
	// This field is from variant [WorkflowDelayStepSettings].
	DelayFor Duration `json:"delay_for"`
	// This field is from variant [WorkflowDelayStepSettings].
	DelayUntilFieldPath string `json:"delay_until_field_path"`
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
	// This field is from variant [RequestTemplate].
	Method RequestTemplateMethod `json:"method"`
	// This field is from variant [RequestTemplate].
	URL string `json:"url"`
	// This field is from variant [RequestTemplate].
	Body string `json:"body"`
	// This field is from variant [RequestTemplate].
	Headers []RequestTemplateHeader `json:"headers"`
	// This field is from variant [RequestTemplate].
	QueryParams []RequestTemplateQueryParam `json:"query_params"`
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
	Data string `json:"data"`
	// This field is from variant [WorkflowTriggerWorkflowStepSettings].
	Recipients string `json:"recipients"`
	// This field is from variant [WorkflowTriggerWorkflowStepSettings].
	Tenant string `json:"tenant"`
	// This field is from variant [WorkflowTriggerWorkflowStepSettings].
	WorkflowKey string `json:"workflow_key"`
	JSON        struct {
		DelayFor                  resp.Field
		DelayUntilFieldPath       resp.Field
		BatchExecutionMode        resp.Field
		BatchItemsMaxLimit        resp.Field
		BatchItemsRenderLimit     resp.Field
		BatchKey                  resp.Field
		BatchOrder                resp.Field
		BatchUntilFieldPath       resp.Field
		BatchWindow               resp.Field
		BatchWindowExtensionLimit resp.Field
		BatchWindowType           resp.Field
		Method                    resp.Field
		URL                       resp.Field
		Body                      resp.Field
		Headers                   resp.Field
		QueryParams               resp.Field
		ThrottleKey               resp.Field
		ThrottleLimit             resp.Field
		ThrottleWindow            resp.Field
		ThrottleWindowFieldPath   resp.Field
		Actor                     resp.Field
		CancellationKey           resp.Field
		Data                      resp.Field
		Recipients                resp.Field
		Tenant                    resp.Field
		WorkflowKey               resp.Field
		raw                       string
	} `json:"-"`
}

func (r *WorkflowStepUnionSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkflowStepUnion to a WorkflowStepUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkflowStepUnionParam.IsOverridden()
func (r WorkflowStepUnion) ToParam() WorkflowStepUnionParam {
	return param.OverrideObj[WorkflowStepUnionParam](r.RawJSON())
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowStepUnionParam struct {
	OfWorkflowChannelStep         *WorkflowChannelStepParam         `json:",omitzero,inline"`
	OfWorkflowDelayStep           *WorkflowDelayStepParam           `json:",omitzero,inline"`
	OfWorkflowBatchStep           *WorkflowBatchStepParam           `json:",omitzero,inline"`
	OfWorkflowFetchStep           *WorkflowFetchStepParam           `json:",omitzero,inline"`
	OfWorkflowThrottleStep        *WorkflowThrottleStepParam        `json:",omitzero,inline"`
	OfWorkflowBranchStep          *WorkflowBranchStepParam          `json:",omitzero,inline"`
	OfWorkflowTriggerWorkflowStep *WorkflowTriggerWorkflowStepParam `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u WorkflowStepUnionParam) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u WorkflowStepUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[WorkflowStepUnionParam](u.OfWorkflowChannelStep,
		u.OfWorkflowDelayStep,
		u.OfWorkflowBatchStep,
		u.OfWorkflowFetchStep,
		u.OfWorkflowThrottleStep,
		u.OfWorkflowBranchStep,
		u.OfWorkflowTriggerWorkflowStep)
}

func (u *WorkflowStepUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWorkflowChannelStep) {
		return u.OfWorkflowChannelStep
	} else if !param.IsOmitted(u.OfWorkflowDelayStep) {
		return u.OfWorkflowDelayStep
	} else if !param.IsOmitted(u.OfWorkflowBatchStep) {
		return u.OfWorkflowBatchStep
	} else if !param.IsOmitted(u.OfWorkflowFetchStep) {
		return u.OfWorkflowFetchStep
	} else if !param.IsOmitted(u.OfWorkflowThrottleStep) {
		return u.OfWorkflowThrottleStep
	} else if !param.IsOmitted(u.OfWorkflowBranchStep) {
		return u.OfWorkflowBranchStep
	} else if !param.IsOmitted(u.OfWorkflowTriggerWorkflowStep) {
		return u.OfWorkflowTriggerWorkflowStep
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetTemplate() *WorkflowChannelStepTemplateUnionParam {
	if vt := u.OfWorkflowChannelStep; vt != nil {
		return &vt.Template
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetChannelGroupKey() *string {
	if vt := u.OfWorkflowChannelStep; vt != nil && vt.ChannelGroupKey.IsPresent() {
		return &vt.ChannelGroupKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetChannelKey() *string {
	if vt := u.OfWorkflowChannelStep; vt != nil && vt.ChannelKey.IsPresent() {
		return &vt.ChannelKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetChannelOverrides() *WorkflowChannelStepChannelOverridesUnionParam {
	if vt := u.OfWorkflowChannelStep; vt != nil {
		return &vt.ChannelOverrides
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetSendWindows() []SendWindowParam {
	if vt := u.OfWorkflowChannelStep; vt != nil {
		return vt.SendWindows
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
func (u WorkflowStepUnionParam) GetName() *string {
	if vt := u.OfWorkflowChannelStep; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfWorkflowDelayStep; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfWorkflowBatchStep; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfWorkflowFetchStep; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfWorkflowThrottleStep; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfWorkflowBranchStep; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfWorkflowTriggerWorkflowStep; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetRef() *string {
	if vt := u.OfWorkflowChannelStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowDelayStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowBatchStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowFetchStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowThrottleStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowBranchStep; vt != nil {
		return (*string)(&vt.Ref)
	} else if vt := u.OfWorkflowTriggerWorkflowStep; vt != nil {
		return (*string)(&vt.Ref)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetType() *string {
	if vt := u.OfWorkflowChannelStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowDelayStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowBatchStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowFetchStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowThrottleStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowBranchStep; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWorkflowTriggerWorkflowStep; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowStepUnionParam) GetDescription() *string {
	if vt := u.OfWorkflowChannelStep; vt != nil && vt.Description.IsPresent() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowDelayStep; vt != nil && vt.Description.IsPresent() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowBatchStep; vt != nil && vt.Description.IsPresent() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowFetchStep; vt != nil && vt.Description.IsPresent() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowThrottleStep; vt != nil && vt.Description.IsPresent() {
		return &vt.Description.Value
	} else if vt := u.OfWorkflowBranchStep; vt != nil {
		return (*string)(&vt.Description)
	} else if vt := u.OfWorkflowTriggerWorkflowStep; vt != nil && vt.Description.IsPresent() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's Conditions property, if present.
func (u WorkflowStepUnionParam) GetConditions() *ConditionGroupUnionParam {
	if vt := u.OfWorkflowChannelStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowDelayStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowFetchStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowThrottleStep; vt != nil {
		return &vt.Conditions
	} else if vt := u.OfWorkflowTriggerWorkflowStep; vt != nil {
		return &vt.Conditions
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u WorkflowStepUnionParam) GetSettings() (res workflowStepUnionParamSettings) {
	if vt := u.OfWorkflowDelayStep; vt != nil {
		res.ofWorkflowDelayStepSettings = &vt.Settings
	} else if vt := u.OfWorkflowBatchStep; vt != nil {
		res.ofWorkflowBatchStepSettings = &vt.Settings
	} else if vt := u.OfWorkflowFetchStep; vt != nil {
		res.ofRequestTemplate = &vt.Settings
	} else if vt := u.OfWorkflowThrottleStep; vt != nil {
		res.ofWorkflowThrottleStepSettings = &vt.Settings
	} else if vt := u.OfWorkflowTriggerWorkflowStep; vt != nil {
		res.ofWorkflowTriggerWorkflowStepSettings = &vt.Settings
	}
	return
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type workflowStepUnionParamSettings struct {
	ofWorkflowDelayStepSettings           *WorkflowDelayStepSettingsParam
	ofWorkflowBatchStepSettings           *WorkflowBatchStepSettingsParam
	ofRequestTemplate                     *RequestTemplateParam
	ofWorkflowThrottleStepSettings        *WorkflowThrottleStepSettingsParam
	ofWorkflowTriggerWorkflowStepSettings *WorkflowTriggerWorkflowStepSettingsParam
}

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *knockmapi.WorkflowDelayStepSettingsParam:
//	case *knockmapi.WorkflowBatchStepSettingsParam:
//	case *knockmapi.RequestTemplateParam:
//	case *knockmapi.WorkflowThrottleStepSettingsParam:
//	case *knockmapi.WorkflowTriggerWorkflowStepSettingsParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowStepUnionParamSettings) AsAny() any {
	if !param.IsOmitted(u.ofWorkflowDelayStepSettings) {
		return u.ofWorkflowDelayStepSettings
	} else if !param.IsOmitted(u.ofWorkflowBatchStepSettings) {
		return u.ofWorkflowBatchStepSettings
	} else if !param.IsOmitted(u.ofRequestTemplate) {
		return u.ofRequestTemplate
	} else if !param.IsOmitted(u.ofWorkflowThrottleStepSettings) {
		return u.ofWorkflowThrottleStepSettings
	} else if !param.IsOmitted(u.ofWorkflowTriggerWorkflowStepSettings) {
		return u.ofWorkflowTriggerWorkflowStepSettings
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetDelayFor() *DurationParam {
	if vt := u.ofWorkflowDelayStepSettings; vt != nil {
		return &vt.DelayFor
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetDelayUntilFieldPath() *string {
	if vt := u.ofWorkflowDelayStepSettings; vt != nil && vt.DelayUntilFieldPath.IsPresent() {
		return &vt.DelayUntilFieldPath.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchExecutionMode() *string {
	if vt := u.ofWorkflowBatchStepSettings; vt != nil {
		return &vt.BatchExecutionMode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchItemsMaxLimit() *int64 {
	if vt := u.ofWorkflowBatchStepSettings; vt != nil && vt.BatchItemsMaxLimit.IsPresent() {
		return &vt.BatchItemsMaxLimit.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchItemsRenderLimit() *int64 {
	if vt := u.ofWorkflowBatchStepSettings; vt != nil && vt.BatchItemsRenderLimit.IsPresent() {
		return &vt.BatchItemsRenderLimit.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchKey() *string {
	if vt := u.ofWorkflowBatchStepSettings; vt != nil && vt.BatchKey.IsPresent() {
		return &vt.BatchKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchOrder() *string {
	if vt := u.ofWorkflowBatchStepSettings; vt != nil {
		return &vt.BatchOrder
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchUntilFieldPath() *string {
	if vt := u.ofWorkflowBatchStepSettings; vt != nil && vt.BatchUntilFieldPath.IsPresent() {
		return &vt.BatchUntilFieldPath.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchWindow() *DurationParam {
	if vt := u.ofWorkflowBatchStepSettings; vt != nil {
		return &vt.BatchWindow
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchWindowExtensionLimit() *DurationParam {
	if vt := u.ofWorkflowBatchStepSettings; vt != nil {
		return &vt.BatchWindowExtensionLimit
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBatchWindowType() *string {
	if vt := u.ofWorkflowBatchStepSettings; vt != nil {
		return &vt.BatchWindowType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetMethod() *string {
	if vt := u.ofRequestTemplate; vt != nil {
		return (*string)(&vt.Method)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetURL() *string {
	if vt := u.ofRequestTemplate; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetBody() *string {
	if vt := u.ofRequestTemplate; vt != nil && vt.Body.IsPresent() {
		return &vt.Body.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetHeaders() []RequestTemplateHeaderParam {
	if vt := u.ofRequestTemplate; vt != nil {
		return vt.Headers
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetQueryParams() []RequestTemplateQueryParamParam {
	if vt := u.ofRequestTemplate; vt != nil {
		return vt.QueryParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetThrottleKey() *string {
	if vt := u.ofWorkflowThrottleStepSettings; vt != nil && vt.ThrottleKey.IsPresent() {
		return &vt.ThrottleKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetThrottleLimit() *int64 {
	if vt := u.ofWorkflowThrottleStepSettings; vt != nil && vt.ThrottleLimit.IsPresent() {
		return &vt.ThrottleLimit.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetThrottleWindow() *DurationParam {
	if vt := u.ofWorkflowThrottleStepSettings; vt != nil {
		return &vt.ThrottleWindow
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetThrottleWindowFieldPath() *string {
	if vt := u.ofWorkflowThrottleStepSettings; vt != nil && vt.ThrottleWindowFieldPath.IsPresent() {
		return &vt.ThrottleWindowFieldPath.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetActor() *string {
	if vt := u.ofWorkflowTriggerWorkflowStepSettings; vt != nil && vt.Actor.IsPresent() {
		return &vt.Actor.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetCancellationKey() *string {
	if vt := u.ofWorkflowTriggerWorkflowStepSettings; vt != nil && vt.CancellationKey.IsPresent() {
		return &vt.CancellationKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetData() *string {
	if vt := u.ofWorkflowTriggerWorkflowStepSettings; vt != nil && vt.Data.IsPresent() {
		return &vt.Data.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetRecipients() *string {
	if vt := u.ofWorkflowTriggerWorkflowStepSettings; vt != nil && vt.Recipients.IsPresent() {
		return &vt.Recipients.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetTenant() *string {
	if vt := u.ofWorkflowTriggerWorkflowStepSettings; vt != nil && vt.Tenant.IsPresent() {
		return &vt.Tenant.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowStepUnionParamSettings) GetWorkflowKey() *string {
	if vt := u.ofWorkflowTriggerWorkflowStepSettings; vt != nil && vt.WorkflowKey.IsPresent() {
		return &vt.WorkflowKey.Value
	}
	return nil
}

// A throttle function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/throttle-function).
type WorkflowThrottleStep struct {
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// The settings for the throttle step.
	Settings WorkflowThrottleStepSettings `json:"settings,required"`
	// The type of the workflow step.
	//
	// Any of "throttle".
	Type WorkflowThrottleStepType `json:"type,required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions,nullable"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description string `json:"description,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		Ref         resp.Field
		Settings    resp.Field
		Type        resp.Field
		Conditions  resp.Field
		Description resp.Field
		ExtraFields map[string]resp.Field
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
// WorkflowThrottleStepParam.IsOverridden()
func (r WorkflowThrottleStep) ToParam() WorkflowThrottleStepParam {
	return param.OverrideObj[WorkflowThrottleStepParam](r.RawJSON())
}

// The settings for the throttle step.
type WorkflowThrottleStepSettings struct {
	// The data property to use to throttle notifications per recipient.
	ThrottleKey string `json:"throttle_key,nullable"`
	// The maximum number of workflows to allow within the duration window. Defaults
	// to 1.
	ThrottleLimit int64 `json:"throttle_limit,nullable"`
	// A duration of time, represented as a unit and a value.
	ThrottleWindow Duration `json:"throttle_window,nullable"`
	// The data path to resolve a dynamic throttle window. The resolved value must be
	// an ISO-8601 timestamp. See more in the
	// [docs](https://docs.knock.app/designing-workflows/throttle-function#set-a-dynamic-throttle-window).
	ThrottleWindowFieldPath string `json:"throttle_window_field_path,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ThrottleKey             resp.Field
		ThrottleLimit           resp.Field
		ThrottleWindow          resp.Field
		ThrottleWindowFieldPath resp.Field
		ExtraFields             map[string]resp.Field
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
// The properties Name, Ref, Settings, Type are required.
type WorkflowThrottleStepParam struct {
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// The settings for the throttle step.
	Settings WorkflowThrottleStepSettingsParam `json:"settings,omitzero,required"`
	// The type of the workflow step.
	//
	// Any of "throttle".
	Type WorkflowThrottleStepType `json:"type,omitzero,required"`
	// An arbitrary string attached to a workflow step. Useful for adding notes about
	// the workflow for internal purposes.
	Description param.Opt[string] `json:"description,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowThrottleStepParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowThrottleStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowThrottleStepParam
	return param.MarshalObject(r, (*shadow)(&r))
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowThrottleStepSettingsParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r WorkflowThrottleStepSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowThrottleStepSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// A workflow trigger function step. Read more in the
// [docs](https://docs.knock.app/designing-workflows/trigger-workflow-function).
type WorkflowTriggerWorkflowStep struct {
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// The settings for the workflow trigger workflow step.
	Settings WorkflowTriggerWorkflowStepSettings `json:"settings,required"`
	// The type of the workflow step.
	//
	// Any of "trigger_workflow".
	Type WorkflowTriggerWorkflowStepType `json:"type,required"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions,nullable"`
	// A description for the workflow step.
	Description string `json:"description"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Name        resp.Field
		Ref         resp.Field
		Settings    resp.Field
		Type        resp.Field
		Conditions  resp.Field
		Description resp.Field
		ExtraFields map[string]resp.Field
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
// WorkflowTriggerWorkflowStepParam.IsOverridden()
func (r WorkflowTriggerWorkflowStep) ToParam() WorkflowTriggerWorkflowStepParam {
	return param.OverrideObj[WorkflowTriggerWorkflowStepParam](r.RawJSON())
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
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Actor           resp.Field
		CancellationKey resp.Field
		Data            resp.Field
		Recipients      resp.Field
		Tenant          resp.Field
		WorkflowKey     resp.Field
		ExtraFields     map[string]resp.Field
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
// The properties Name, Ref, Settings, Type are required.
type WorkflowTriggerWorkflowStepParam struct {
	// A name for the workflow step.
	Name string `json:"name,required"`
	// The reference key of the workflow step. Must be unique per workflow.
	Ref string `json:"ref,required"`
	// The settings for the workflow trigger workflow step.
	Settings WorkflowTriggerWorkflowStepSettingsParam `json:"settings,omitzero,required"`
	// The type of the workflow step.
	//
	// Any of "trigger_workflow".
	Type WorkflowTriggerWorkflowStepType `json:"type,omitzero,required"`
	// A description for the workflow step.
	Description param.Opt[string] `json:"description,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowTriggerWorkflowStepParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowTriggerWorkflowStepParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerWorkflowStepParam
	return param.MarshalObject(r, (*shadow)(&r))
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowTriggerWorkflowStepSettingsParam) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r WorkflowTriggerWorkflowStepSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerWorkflowStepSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Wraps the Workflow response under the `workflow` key.
type WorkflowActivateResponse struct {
	// A workflow object.
	Workflow Workflow `json:"workflow,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Workflow    resp.Field
		ExtraFields map[string]resp.Field
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
	WorkflowRunID string `json:"workflow_run_id,required" format:"uuid"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		WorkflowRunID resp.Field
		ExtraFields   map[string]resp.Field
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
	// A workflow object.
	Workflow Workflow `json:"workflow,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Workflow    resp.Field
		ExtraFields map[string]resp.Field
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
	// A workflow object.
	Workflow Workflow `json:"workflow,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Workflow    resp.Field
		ExtraFields map[string]resp.Field
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
	Environment string `query:"environment,required" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// Whether to hide uncommitted changes. When true, only committed changes will be
	// returned. When false, both committed and uncommitted changes will be returned.
	HideUncommittedChanges param.Opt[bool] `query:"hide_uncommitted_changes,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowGetParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [WorkflowGetParams]'s query parameters as `url.Values`.
func (r WorkflowGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowListParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Whether to hide uncommitted changes. When true, only committed changes will be
	// returned. When false, both committed and uncommitted changes will be returned.
	HideUncommittedChanges param.Opt[bool] `query:"hide_uncommitted_changes,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [WorkflowListParams]'s query parameters as `url.Values`.
func (r WorkflowListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowActivateParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// Whether to activate or deactivate the workflow. Set to `true` by default, which
	// will activate the workflow.
	Status bool `json:"status,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowActivateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r WorkflowActivateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowActivateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// URLQuery serializes [WorkflowActivateParams]'s query parameters as `url.Values`.
func (r WorkflowActivateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowRunParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// A list of recipients to run the workflow for.
	Recipients []WorkflowRunParamsRecipientUnion `json:"recipients,omitzero,required"`
	// A key to cancel the workflow run.
	CancellationKey param.Opt[string] `json:"cancellation_key,omitzero"`
	// The tenant to associate the workflow run with.
	Tenant param.Opt[string] `json:"tenant,omitzero"`
	// A recipient reference, used when referencing a recipient by either their ID (for
	// a user), or by a reference for an object.
	Actor WorkflowRunParamsActorUnion `json:"actor,omitzero"`
	// A map of data to be used in the workflow run.
	Data map[string]interface{} `json:"data,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowRunParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r WorkflowRunParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// URLQuery serializes [WorkflowRunParams]'s query parameters as `url.Values`.
func (r WorkflowRunParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowRunParamsRecipientUnion struct {
	OfString                      param.Opt[string]                 `json:",omitzero,inline"`
	OfWorkflowRunsRecipientObject *WorkflowRunParamsRecipientObject `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u WorkflowRunParamsRecipientUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u WorkflowRunParamsRecipientUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[WorkflowRunParamsRecipientUnion](u.OfString, u.OfWorkflowRunsRecipientObject)
}

func (u *WorkflowRunParamsRecipientUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfWorkflowRunsRecipientObject) {
		return u.OfWorkflowRunsRecipientObject
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowRunParamsRecipientUnion) GetID() *string {
	if vt := u.OfWorkflowRunsRecipientObject; vt != nil {
		return &vt.ID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowRunParamsRecipientUnion) GetCollection() *string {
	if vt := u.OfWorkflowRunsRecipientObject; vt != nil {
		return &vt.Collection
	}
	return nil
}

// An object reference.
//
// The properties ID, Collection are required.
type WorkflowRunParamsRecipientObject struct {
	ID         string `json:"id,required"`
	Collection string `json:"collection,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowRunParamsRecipientObject) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowRunParamsRecipientObject) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRunParamsRecipientObject
	return param.MarshalObject(r, (*shadow)(&r))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowRunParamsActorUnion struct {
	OfString                  param.Opt[string]             `json:",omitzero,inline"`
	OfWorkflowRunsActorObject *WorkflowRunParamsActorObject `json:",omitzero,inline"`
	paramUnion
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (u WorkflowRunParamsActorUnion) IsPresent() bool { return !param.IsOmitted(u) && !u.IsNull() }
func (u WorkflowRunParamsActorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[WorkflowRunParamsActorUnion](u.OfString, u.OfWorkflowRunsActorObject)
}

func (u *WorkflowRunParamsActorUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfWorkflowRunsActorObject) {
		return u.OfWorkflowRunsActorObject
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowRunParamsActorUnion) GetID() *string {
	if vt := u.OfWorkflowRunsActorObject; vt != nil {
		return &vt.ID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowRunParamsActorUnion) GetCollection() *string {
	if vt := u.OfWorkflowRunsActorObject; vt != nil {
		return &vt.Collection
	}
	return nil
}

// An object reference.
//
// The properties ID, Collection are required.
type WorkflowRunParamsActorObject struct {
	ID         string `json:"id,required"`
	Collection string `json:"collection,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowRunParamsActorObject) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowRunParamsActorObject) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRunParamsActorObject
	return param.MarshalObject(r, (*shadow)(&r))
}

type WorkflowUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// A workflow request for upserting a workflow.
	Workflow WorkflowUpsertParamsWorkflow `json:"workflow,omitzero,required"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// Whether to commit the resource at the same time as modifying it.
	Commit param.Opt[bool] `query:"commit,omitzero" json:"-"`
	// The message to commit the resource with, only used if `commit` is `true`.
	CommitMessage param.Opt[string] `query:"commit_message,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowUpsertParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r WorkflowUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// URLQuery serializes [WorkflowUpsertParams]'s query parameters as `url.Values`.
func (r WorkflowUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A workflow request for upserting a workflow.
//
// The properties Name, Steps are required.
type WorkflowUpsertParamsWorkflow struct {
	// A name for the workflow. Must be at maximum 255 characters in length.
	Name string `json:"name,required"`
	// A list of workflow step objects in the workflow.
	Steps []WorkflowStepUnionParam `json:"steps,omitzero,required"`
	// An arbitrary string attached to a workflow object. Useful for adding notes about
	// the workflow for internal purposes. Maximum of 280 characters allowed.
	Description param.Opt[string] `json:"description,omitzero"`
	// A list of
	// [categories](https://docs.knock.app/concepts/workflows#workflow-categories) that
	// the workflow belongs to.
	Categories []string `json:"categories,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	// A map of workflow settings.
	Settings WorkflowUpsertParamsWorkflowSettings `json:"settings,omitzero"`
	// A JSON schema for the expected structure of the workflow trigger's data payload.
	// Used to validate trigger requests. Read more in the
	// [docs](https://docs.knock.app/developer-tools/validating-trigger-data).
	TriggerDataJsonSchema map[string]interface{} `json:"trigger_data_json_schema,omitzero"`
	// The frequency at which the workflow should be triggered. One of:
	// `once_per_recipient`, `once_per_recipient_per_tenant`, `every_trigger`. Defaults
	// to `every_trigger`. Read more in
	// [docs](https://docs.knock.app/send-notifications/triggering-workflows/overview#controlling-workflow-trigger-frequency).
	//
	// Any of "every_trigger", "once_per_recipient", "once_per_recipient_per_tenant".
	TriggerFrequency string `json:"trigger_frequency,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowUpsertParamsWorkflow) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowUpsertParamsWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpsertParamsWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[WorkflowUpsertParamsWorkflow](
		"TriggerFrequency", false, "every_trigger", "once_per_recipient", "once_per_recipient_per_tenant",
	)
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowUpsertParamsWorkflowSettings) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r WorkflowUpsertParamsWorkflowSettings) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpsertParamsWorkflowSettings
	return param.MarshalObject(r, (*shadow)(&r))
}

type WorkflowValidateParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// A workflow request for upserting a workflow.
	Workflow WorkflowValidateParamsWorkflow `json:"workflow,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowValidateParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r WorkflowValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}

// URLQuery serializes [WorkflowValidateParams]'s query parameters as `url.Values`.
func (r WorkflowValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A workflow request for upserting a workflow.
//
// The properties Name, Steps are required.
type WorkflowValidateParamsWorkflow struct {
	// A name for the workflow. Must be at maximum 255 characters in length.
	Name string `json:"name,required"`
	// A list of workflow step objects in the workflow.
	Steps []WorkflowStepUnionParam `json:"steps,omitzero,required"`
	// An arbitrary string attached to a workflow object. Useful for adding notes about
	// the workflow for internal purposes. Maximum of 280 characters allowed.
	Description param.Opt[string] `json:"description,omitzero"`
	// A list of
	// [categories](https://docs.knock.app/concepts/workflows#workflow-categories) that
	// the workflow belongs to.
	Categories []string `json:"categories,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	// A map of workflow settings.
	Settings WorkflowValidateParamsWorkflowSettings `json:"settings,omitzero"`
	// A JSON schema for the expected structure of the workflow trigger's data payload.
	// Used to validate trigger requests. Read more in the
	// [docs](https://docs.knock.app/developer-tools/validating-trigger-data).
	TriggerDataJsonSchema map[string]interface{} `json:"trigger_data_json_schema,omitzero"`
	// The frequency at which the workflow should be triggered. One of:
	// `once_per_recipient`, `once_per_recipient_per_tenant`, `every_trigger`. Defaults
	// to `every_trigger`. Read more in
	// [docs](https://docs.knock.app/send-notifications/triggering-workflows/overview#controlling-workflow-trigger-frequency).
	//
	// Any of "every_trigger", "once_per_recipient", "once_per_recipient_per_tenant".
	TriggerFrequency string `json:"trigger_frequency,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowValidateParamsWorkflow) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r WorkflowValidateParamsWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowValidateParamsWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}

func init() {
	apijson.RegisterFieldValidator[WorkflowValidateParamsWorkflow](
		"TriggerFrequency", false, "every_trigger", "once_per_recipient", "once_per_recipient_per_tenant",
	)
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

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WorkflowValidateParamsWorkflowSettings) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}
func (r WorkflowValidateParamsWorkflowSettings) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowValidateParamsWorkflowSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
