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
	"github.com/knocklabs/knock-mgmt-go/shared/constant"
)

// Audiences define sets of users that can be targeted for notifications.
//
// AudienceService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudienceService] method instead.
type AudienceService struct {
	Options []option.RequestOption
}

// NewAudienceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAudienceService(opts ...option.RequestOption) (r AudienceService) {
	r = AudienceService{}
	r.Options = opts
	return
}

// Retrieve an audience by its key in a given environment.
func (r *AudienceService) Get(ctx context.Context, audienceKey string, query AudienceGetParams, opts ...option.RequestOption) (res *AudienceUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if audienceKey == "" {
		err = errors.New("missing required audience_key parameter")
		return
	}
	path := fmt.Sprintf("v1/audiences/%s", audienceKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a paginated list of audiences for the given environment.
func (r *AudienceService) List(ctx context.Context, query AudienceListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[AudienceUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/audiences"
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

// Returns a paginated list of audiences for the given environment.
func (r *AudienceService) ListAutoPaging(ctx context.Context, query AudienceListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[AudienceUnion] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Archives a given audience across all environments.
func (r *AudienceService) Archive(ctx context.Context, audienceKey string, body AudienceArchiveParams, opts ...option.RequestOption) (res *AudienceArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if audienceKey == "" {
		err = errors.New("missing required audience_key parameter")
		return
	}
	path := fmt.Sprintf("v1/audiences/%s", audienceKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Updates an audience of a given key, or creates a new one if it does not yet
// exist.
func (r *AudienceService) Upsert(ctx context.Context, audienceKey string, params AudienceUpsertParams, opts ...option.RequestOption) (res *AudienceUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if audienceKey == "" {
		err = errors.New("missing required audience_key parameter")
		return
	}
	path := fmt.Sprintf("v1/audiences/%s", audienceKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Validates an audience payload without persisting it.
func (r *AudienceService) Validate(ctx context.Context, audienceKey string, params AudienceValidateParams, opts ...option.RequestOption) (res *AudienceValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if audienceKey == "" {
		err = errors.New("missing required audience_key parameter")
		return
	}
	path := fmt.Sprintf("v1/audiences/%s/validate", audienceKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// AudienceUnion contains all possible properties and values from [StaticAudience],
// [DynamicAudience].
//
// Use the [AudienceUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AudienceUnion struct {
	CreatedAt   time.Time `json:"created_at"`
	Environment string    `json:"environment"`
	Key         string    `json:"key"`
	Name        string    `json:"name"`
	// Any of "static", "dynamic".
	Type        string    `json:"type"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	Sha         string    `json:"sha"`
	// This field is from variant [DynamicAudience].
	Segments []DynamicAudienceSegment `json:"segments"`
	JSON     struct {
		CreatedAt   respjson.Field
		Environment respjson.Field
		Key         respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		Sha         respjson.Field
		Segments    respjson.Field
		raw         string
	} `json:"-"`
}

// anyAudience is implemented by each variant of [AudienceUnion] to add type safety
// for the return type of [AudienceUnion.AsAny]
type anyAudience interface {
	implAudienceUnion()
}

func (StaticAudience) implAudienceUnion()  {}
func (DynamicAudience) implAudienceUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AudienceUnion.AsAny().(type) {
//	case knockmapi.StaticAudience:
//	case knockmapi.DynamicAudience:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AudienceUnion) AsAny() anyAudience {
	switch u.Type {
	case "static":
		return u.AsStatic()
	case "dynamic":
		return u.AsDynamic()
	}
	return nil
}

func (u AudienceUnion) AsStatic() (v StaticAudience) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AudienceUnion) AsDynamic() (v DynamicAudience) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AudienceUnion) RawJSON() string { return u.JSON.raw }

func (r *AudienceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A condition to evaluate for audience membership.
type AudienceCondition struct {
	// The operator to use when evaluating the condition.
	//
	// Any of "equal_to", "not_equal_to", "greater_than", "less_than",
	// "greater_than_or_equal_to", "less_than_or_equal_to", "contains", "not_contains",
	// "contains_all", "not_contains_all", "is_timestamp_before",
	// "is_timestamp_on_or_after", "is_timestamp_between", "empty", "not_empty",
	// "exists", "not_exists", "is_timestamp", "is_audience_member",
	// "is_not_audience_member".
	Operator AudienceConditionOperator `json:"operator" api:"required"`
	// The property to be evaluated. Properties are dynamic values using path
	// expressions like `recipient.plan` or `recipient.created_at`.
	Property string `json:"property" api:"required"`
	// The argument to compare against. Can be a static value (string, number, boolean)
	// or a dynamic path expression.
	Argument string `json:"argument" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Property    respjson.Field
		Argument    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceCondition) RawJSON() string { return r.JSON.raw }
func (r *AudienceCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AudienceCondition to a AudienceConditionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AudienceConditionParam.Overrides()
func (r AudienceCondition) ToParam() AudienceConditionParam {
	return param.Override[AudienceConditionParam](json.RawMessage(r.RawJSON()))
}

// The operator to use when evaluating the condition.
type AudienceConditionOperator string

const (
	AudienceConditionOperatorEqualTo              AudienceConditionOperator = "equal_to"
	AudienceConditionOperatorNotEqualTo           AudienceConditionOperator = "not_equal_to"
	AudienceConditionOperatorGreaterThan          AudienceConditionOperator = "greater_than"
	AudienceConditionOperatorLessThan             AudienceConditionOperator = "less_than"
	AudienceConditionOperatorGreaterThanOrEqualTo AudienceConditionOperator = "greater_than_or_equal_to"
	AudienceConditionOperatorLessThanOrEqualTo    AudienceConditionOperator = "less_than_or_equal_to"
	AudienceConditionOperatorContains             AudienceConditionOperator = "contains"
	AudienceConditionOperatorNotContains          AudienceConditionOperator = "not_contains"
	AudienceConditionOperatorContainsAll          AudienceConditionOperator = "contains_all"
	AudienceConditionOperatorNotContainsAll       AudienceConditionOperator = "not_contains_all"
	AudienceConditionOperatorIsTimestampBefore    AudienceConditionOperator = "is_timestamp_before"
	AudienceConditionOperatorIsTimestampOnOrAfter AudienceConditionOperator = "is_timestamp_on_or_after"
	AudienceConditionOperatorIsTimestampBetween   AudienceConditionOperator = "is_timestamp_between"
	AudienceConditionOperatorEmpty                AudienceConditionOperator = "empty"
	AudienceConditionOperatorNotEmpty             AudienceConditionOperator = "not_empty"
	AudienceConditionOperatorExists               AudienceConditionOperator = "exists"
	AudienceConditionOperatorNotExists            AudienceConditionOperator = "not_exists"
	AudienceConditionOperatorIsTimestamp          AudienceConditionOperator = "is_timestamp"
	AudienceConditionOperatorIsAudienceMember     AudienceConditionOperator = "is_audience_member"
	AudienceConditionOperatorIsNotAudienceMember  AudienceConditionOperator = "is_not_audience_member"
)

// A condition to evaluate for audience membership.
//
// The properties Operator, Property are required.
type AudienceConditionParam struct {
	// The operator to use when evaluating the condition.
	//
	// Any of "equal_to", "not_equal_to", "greater_than", "less_than",
	// "greater_than_or_equal_to", "less_than_or_equal_to", "contains", "not_contains",
	// "contains_all", "not_contains_all", "is_timestamp_before",
	// "is_timestamp_on_or_after", "is_timestamp_between", "empty", "not_empty",
	// "exists", "not_exists", "is_timestamp", "is_audience_member",
	// "is_not_audience_member".
	Operator AudienceConditionOperator `json:"operator,omitzero" api:"required"`
	// The property to be evaluated. Properties are dynamic values using path
	// expressions like `recipient.plan` or `recipient.created_at`.
	Property string `json:"property" api:"required"`
	// The argument to compare against. Can be a static value (string, number, boolean)
	// or a dynamic path expression.
	Argument param.Opt[string] `json:"argument,omitzero"`
	paramObj
}

func (r AudienceConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow AudienceConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dynamic audience where membership is determined by segment conditions
// evaluated at runtime.
type DynamicAudience struct {
	// The timestamp of when the audience was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The slug of the environment in which the audience exists.
	Environment string `json:"environment" api:"required"`
	// The unique key of the audience.
	Key string `json:"key" api:"required"`
	// The name of the audience.
	Name string `json:"name" api:"required"`
	// A list of segments that define the dynamic audience membership criteria. Each
	// segment contains one or more conditions joined by AND. Multiple segments are
	// joined by OR.
	Segments []DynamicAudienceSegment `json:"segments" api:"required"`
	// The type of audience. Always `dynamic` for dynamic audiences.
	//
	// Any of "dynamic".
	Type DynamicAudienceType `json:"type" api:"required"`
	// The timestamp of when the audience was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// A description of the audience.
	Description string `json:"description" api:"nullable"`
	// The SHA hash of the audience data.
	Sha string `json:"sha" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Environment respjson.Field
		Key         respjson.Field
		Name        respjson.Field
		Segments    respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		Sha         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicAudience) RawJSON() string { return r.JSON.raw }
func (r *DynamicAudience) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DynamicAudienceSegment struct {
	// A list of conditions within this segment, joined by AND.
	Conditions []AudienceCondition `json:"conditions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conditions  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicAudienceSegment) RawJSON() string { return r.JSON.raw }
func (r *DynamicAudienceSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of audience. Always `dynamic` for dynamic audiences.
type DynamicAudienceType string

const (
	DynamicAudienceTypeDynamic DynamicAudienceType = "dynamic"
)

// A static audience where members are explicitly added or removed via the API.
type StaticAudience struct {
	// The timestamp of when the audience was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The slug of the environment in which the audience exists.
	Environment string `json:"environment" api:"required"`
	// The unique key of the audience.
	Key string `json:"key" api:"required"`
	// The name of the audience.
	Name string `json:"name" api:"required"`
	// The type of audience. Always `static` for static audiences.
	//
	// Any of "static".
	Type StaticAudienceType `json:"type" api:"required"`
	// The timestamp of when the audience was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// A description of the audience.
	Description string `json:"description" api:"nullable"`
	// The SHA hash of the audience data.
	Sha string `json:"sha" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Environment respjson.Field
		Key         respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		Sha         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StaticAudience) RawJSON() string { return r.JSON.raw }
func (r *StaticAudience) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of audience. Always `static` for static audiences.
type StaticAudienceType string

const (
	StaticAudienceTypeStatic StaticAudienceType = "static"
)

// The response from archiving an audience.
type AudienceArchiveResponse struct {
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
func (r AudienceArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *AudienceArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Audience response under the `audience` key.
type AudienceUpsertResponse struct {
	// An audience defines a set of users that can be targeted for notifications. Can
	// be either a `StaticAudience` (members explicitly added/removed) or a
	// `DynamicAudience` (members determined by segment conditions).
	Audience AudienceUnion `json:"audience" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audience    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *AudienceUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Audience response under the `audience` key.
type AudienceValidateResponse struct {
	// An audience defines a set of users that can be targeted for notifications. Can
	// be either a `StaticAudience` (members explicitly added/removed) or a
	// `DynamicAudience` (members determined by segment conditions).
	Audience AudienceUnion `json:"audience" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audience    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *AudienceValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceGetParams struct {
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

// URLQuery serializes [AudienceGetParams]'s query parameters as `url.Values`.
func (r AudienceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AudienceListParams struct {
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

// URLQuery serializes [AudienceListParams]'s query parameters as `url.Values`.
func (r AudienceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AudienceArchiveParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [AudienceArchiveParams]'s query parameters as `url.Values`.
func (r AudienceArchiveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AudienceUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// An audience object with attributes to create or update an audience. Use
	// `type: static` for audiences with explicitly managed members, or `type: dynamic`
	// for audiences with segment-based membership.
	Audience AudienceUpsertParamsAudienceUnion `json:"audience,omitzero" api:"required"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// Whether to commit the resource at the same time as modifying it.
	Commit param.Opt[bool] `query:"commit,omitzero" json:"-"`
	// The message to commit the resource with, only used if `commit` is `true`.
	CommitMessage param.Opt[string] `query:"commit_message,omitzero" json:"-"`
	paramObj
}

func (r AudienceUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow AudienceUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [AudienceUpsertParams]'s query parameters as `url.Values`.
func (r AudienceUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AudienceUpsertParamsAudienceUnion struct {
	OfStaticAudienceRequest  *AudienceUpsertParamsAudienceStaticAudienceRequest  `json:",omitzero,inline"`
	OfDynamicAudienceRequest *AudienceUpsertParamsAudienceDynamicAudienceRequest `json:",omitzero,inline"`
	paramUnion
}

func (u AudienceUpsertParamsAudienceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStaticAudienceRequest, u.OfDynamicAudienceRequest)
}
func (u *AudienceUpsertParamsAudienceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AudienceUpsertParamsAudienceUnion) asAny() any {
	if !param.IsOmitted(u.OfStaticAudienceRequest) {
		return u.OfStaticAudienceRequest
	} else if !param.IsOmitted(u.OfDynamicAudienceRequest) {
		return u.OfDynamicAudienceRequest
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AudienceUpsertParamsAudienceUnion) GetSegments() []AudienceUpsertParamsAudienceDynamicAudienceRequestSegment {
	if vt := u.OfDynamicAudienceRequest; vt != nil {
		return vt.Segments
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AudienceUpsertParamsAudienceUnion) GetName() *string {
	if vt := u.OfStaticAudienceRequest; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfDynamicAudienceRequest; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AudienceUpsertParamsAudienceUnion) GetType() *string {
	if vt := u.OfStaticAudienceRequest; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDynamicAudienceRequest; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AudienceUpsertParamsAudienceUnion) GetDescription() *string {
	if vt := u.OfStaticAudienceRequest; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfDynamicAudienceRequest; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AudienceUpsertParamsAudienceUnion](
		"type",
		apijson.Discriminator[AudienceUpsertParamsAudienceStaticAudienceRequest]("static"),
		apijson.Discriminator[AudienceUpsertParamsAudienceDynamicAudienceRequest]("dynamic"),
	)
}

// Request body for creating/updating a static audience.
//
// The properties Name, Type are required.
type AudienceUpsertParamsAudienceStaticAudienceRequest struct {
	// The name of the audience.
	Name string `json:"name" api:"required"`
	// A description of the audience.
	Description param.Opt[string] `json:"description,omitzero"`
	// The type of audience. Set to `static` for static audiences.
	//
	// This field can be elided, and will marshal its zero value as "static".
	Type constant.Static `json:"type" api:"required"`
	paramObj
}

func (r AudienceUpsertParamsAudienceStaticAudienceRequest) MarshalJSON() (data []byte, err error) {
	type shadow AudienceUpsertParamsAudienceStaticAudienceRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceUpsertParamsAudienceStaticAudienceRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for creating/updating a dynamic audience.
//
// The properties Name, Type are required.
type AudienceUpsertParamsAudienceDynamicAudienceRequest struct {
	// The name of the audience.
	Name string `json:"name" api:"required"`
	// A description of the audience.
	Description param.Opt[string] `json:"description,omitzero"`
	// A list of segments that define the dynamic audience membership criteria. Each
	// segment contains one or more conditions joined by AND. Multiple segments are
	// joined by OR.
	Segments []AudienceUpsertParamsAudienceDynamicAudienceRequestSegment `json:"segments,omitzero"`
	// The type of audience. Set to `dynamic` for dynamic audiences.
	//
	// This field can be elided, and will marshal its zero value as "dynamic".
	Type constant.Dynamic `json:"type" api:"required"`
	paramObj
}

func (r AudienceUpsertParamsAudienceDynamicAudienceRequest) MarshalJSON() (data []byte, err error) {
	type shadow AudienceUpsertParamsAudienceDynamicAudienceRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceUpsertParamsAudienceDynamicAudienceRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Conditions is required.
type AudienceUpsertParamsAudienceDynamicAudienceRequestSegment struct {
	// A list of conditions within this segment, joined by AND.
	Conditions []AudienceConditionParam `json:"conditions,omitzero" api:"required"`
	paramObj
}

func (r AudienceUpsertParamsAudienceDynamicAudienceRequestSegment) MarshalJSON() (data []byte, err error) {
	type shadow AudienceUpsertParamsAudienceDynamicAudienceRequestSegment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceUpsertParamsAudienceDynamicAudienceRequestSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceValidateParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// An audience object with attributes to create or update an audience. Use
	// `type: static` for audiences with explicitly managed members, or `type: dynamic`
	// for audiences with segment-based membership.
	Audience AudienceValidateParamsAudienceUnion `json:"audience,omitzero" api:"required"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	paramObj
}

func (r AudienceValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow AudienceValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [AudienceValidateParams]'s query parameters as `url.Values`.
func (r AudienceValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AudienceValidateParamsAudienceUnion struct {
	OfStaticAudienceRequest  *AudienceValidateParamsAudienceStaticAudienceRequest  `json:",omitzero,inline"`
	OfDynamicAudienceRequest *AudienceValidateParamsAudienceDynamicAudienceRequest `json:",omitzero,inline"`
	paramUnion
}

func (u AudienceValidateParamsAudienceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStaticAudienceRequest, u.OfDynamicAudienceRequest)
}
func (u *AudienceValidateParamsAudienceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AudienceValidateParamsAudienceUnion) asAny() any {
	if !param.IsOmitted(u.OfStaticAudienceRequest) {
		return u.OfStaticAudienceRequest
	} else if !param.IsOmitted(u.OfDynamicAudienceRequest) {
		return u.OfDynamicAudienceRequest
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AudienceValidateParamsAudienceUnion) GetSegments() []AudienceValidateParamsAudienceDynamicAudienceRequestSegment {
	if vt := u.OfDynamicAudienceRequest; vt != nil {
		return vt.Segments
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AudienceValidateParamsAudienceUnion) GetName() *string {
	if vt := u.OfStaticAudienceRequest; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfDynamicAudienceRequest; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AudienceValidateParamsAudienceUnion) GetType() *string {
	if vt := u.OfStaticAudienceRequest; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDynamicAudienceRequest; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AudienceValidateParamsAudienceUnion) GetDescription() *string {
	if vt := u.OfStaticAudienceRequest; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfDynamicAudienceRequest; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AudienceValidateParamsAudienceUnion](
		"type",
		apijson.Discriminator[AudienceValidateParamsAudienceStaticAudienceRequest]("static"),
		apijson.Discriminator[AudienceValidateParamsAudienceDynamicAudienceRequest]("dynamic"),
	)
}

// Request body for creating/updating a static audience.
//
// The properties Name, Type are required.
type AudienceValidateParamsAudienceStaticAudienceRequest struct {
	// The name of the audience.
	Name string `json:"name" api:"required"`
	// A description of the audience.
	Description param.Opt[string] `json:"description,omitzero"`
	// The type of audience. Set to `static` for static audiences.
	//
	// This field can be elided, and will marshal its zero value as "static".
	Type constant.Static `json:"type" api:"required"`
	paramObj
}

func (r AudienceValidateParamsAudienceStaticAudienceRequest) MarshalJSON() (data []byte, err error) {
	type shadow AudienceValidateParamsAudienceStaticAudienceRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceValidateParamsAudienceStaticAudienceRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for creating/updating a dynamic audience.
//
// The properties Name, Type are required.
type AudienceValidateParamsAudienceDynamicAudienceRequest struct {
	// The name of the audience.
	Name string `json:"name" api:"required"`
	// A description of the audience.
	Description param.Opt[string] `json:"description,omitzero"`
	// A list of segments that define the dynamic audience membership criteria. Each
	// segment contains one or more conditions joined by AND. Multiple segments are
	// joined by OR.
	Segments []AudienceValidateParamsAudienceDynamicAudienceRequestSegment `json:"segments,omitzero"`
	// The type of audience. Set to `dynamic` for dynamic audiences.
	//
	// This field can be elided, and will marshal its zero value as "dynamic".
	Type constant.Dynamic `json:"type" api:"required"`
	paramObj
}

func (r AudienceValidateParamsAudienceDynamicAudienceRequest) MarshalJSON() (data []byte, err error) {
	type shadow AudienceValidateParamsAudienceDynamicAudienceRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceValidateParamsAudienceDynamicAudienceRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Conditions is required.
type AudienceValidateParamsAudienceDynamicAudienceRequestSegment struct {
	// A list of conditions within this segment, joined by AND.
	Conditions []AudienceConditionParam `json:"conditions,omitzero" api:"required"`
	paramObj
}

func (r AudienceValidateParamsAudienceDynamicAudienceRequestSegment) MarshalJSON() (data []byte, err error) {
	type shadow AudienceValidateParamsAudienceDynamicAudienceRequestSegment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceValidateParamsAudienceDynamicAudienceRequestSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
