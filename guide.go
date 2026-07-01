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

// Guides let you define in-app guides that can be displayed to users based on
// priority and other conditions.
//
// GuideService contains methods and other services that help with interacting with
// the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGuideService] method instead.
type GuideService struct {
	Options []option.RequestOption
}

// NewGuideService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGuideService(opts ...option.RequestOption) (r GuideService) {
	r = GuideService{}
	r.Options = opts
	return
}

// Get a guide by its key.
func (r *GuideService) Get(ctx context.Context, guideKey string, query GuideGetParams, opts ...option.RequestOption) (res *Guide, err error) {
	opts = slices.Concat(r.Options, opts)
	if guideKey == "" {
		err = errors.New("missing required guide_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/guides/%s", guideKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of guides available in a given environment.
func (r *GuideService) List(ctx context.Context, query GuideListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Guide], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/guides"
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

// Returns a paginated list of guides available in a given environment.
func (r *GuideService) ListAutoPaging(ctx context.Context, query GuideListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Guide] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Activates (or deactivates) a guide in a given environment. You can either set
// the active status immediately or schedule it.
//
// Note: This immediately enables or disables a guide in a given environment
// without needing to go through environment promotion.
func (r *GuideService) Activate(ctx context.Context, guideKey string, params GuideActivateParams, opts ...option.RequestOption) (res *GuideActivateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if guideKey == "" {
		err = errors.New("missing required guide_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/guides/%s/activate", guideKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Archives a given guide across all environments.
func (r *GuideService) Archive(ctx context.Context, guideKey string, opts ...option.RequestOption) (res *GuideArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if guideKey == "" {
		err = errors.New("missing required guide_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/guides/%s", guideKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Updates a guide of a given key, or creates a new one if it does not yet exist.
//
// Note: this endpoint only operates on guides in the "development" environment.
func (r *GuideService) Upsert(ctx context.Context, guideKey string, params GuideUpsertParams, opts ...option.RequestOption) (res *GuideUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if guideKey == "" {
		err = errors.New("missing required guide_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/guides/%s", guideKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Validates a guide payload without persisting it.
//
// Note: Validating a guide is only done in the development environment context.
func (r *GuideService) Validate(ctx context.Context, guideKey string, params GuideValidateParams, opts ...option.RequestOption) (res *GuideValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if guideKey == "" {
		err = errors.New("missing required guide_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/guides/%s/validate", guideKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A guide defines an in-app guide that can be displayed to users based on priority
// and other conditions.
type Guide struct {
	// Whether the guide is active.
	Active bool `json:"active" api:"required"`
	// The timestamp of when the guide was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The slug of the environment in which the guide exists.
	Environment string `json:"environment" api:"required"`
	// The unique key string for the guide object. Must be at minimum 3 characters and
	// at maximum 255 characters in length. Must be in the format of ^[a-z0-9_-]+$.
	Key string `json:"key" api:"required"`
	// A name for the guide. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// The SHA hash of the guide.
	Sha string `json:"sha" api:"required"`
	// The timestamp of when the guide was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// A list of activation url patterns that describe when the guide should be shown.
	ActivationURLPatterns []GuideActivationURLPattern `json:"activation_url_patterns"`
	// The timestamp of when the guide was archived.
	ArchivedAt time.Time `json:"archived_at" api:"nullable" format:"date-time"`
	// The key of the channel in which the guide exists.
	ChannelKey string `json:"channel_key"`
	// The timestamp of when the guide was deleted.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// An arbitrary string attached to a guide object. Useful for adding notes about
	// the guide for internal purposes. Maximum of 280 characters allowed.
	Description string `json:"description" api:"nullable"`
	// The semver of the guide.
	Semver string `json:"semver"`
	// A list of guide step objects in the guide.
	Steps []GuideStep `json:"steps"`
	// Use tags to organize resources internally within your account. For example, by
	// team or product area.
	Tags []string `json:"tags"`
	// The key of the target audience for the guide. When not set, will default to
	// targeting all users.
	TargetAudienceKey string `json:"target_audience_key" api:"nullable"`
	// A group of conditions to be evaluated.
	TargetPropertyConditions ConditionGroupUnion `json:"target_property_conditions" api:"nullable"`
	// The type of the guide. This is derived from the message type of the guide steps.
	Type string `json:"type"`
	// Whether the guide is valid.
	Valid bool `json:"valid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active                   respjson.Field
		CreatedAt                respjson.Field
		Environment              respjson.Field
		Key                      respjson.Field
		Name                     respjson.Field
		Sha                      respjson.Field
		UpdatedAt                respjson.Field
		ActivationURLPatterns    respjson.Field
		ArchivedAt               respjson.Field
		ChannelKey               respjson.Field
		DeletedAt                respjson.Field
		Description              respjson.Field
		Semver                   respjson.Field
		Steps                    respjson.Field
		Tags                     respjson.Field
		TargetAudienceKey        respjson.Field
		TargetPropertyConditions respjson.Field
		Type                     respjson.Field
		Valid                    respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Guide) RawJSON() string { return r.JSON.raw }
func (r *Guide) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A rule that controls when a guide should be shown based on the user's location
// in the application. At least one of `pathname` or `search` must be provided.
type GuideActivationURLPattern struct {
	// Whether to allow or block the guide at the specified location.
	//
	// Any of "allow", "block".
	Directive GuideActivationURLPatternDirective `json:"directive" api:"required"`
	// The URL pathname pattern to match against. Must be a valid URI path.
	Pathname string `json:"pathname"`
	// The URL query string pattern to match against (without the leading '?').
	// Supports URLPattern API syntax.
	Search string `json:"search"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Directive   respjson.Field
		Pathname    respjson.Field
		Search      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuideActivationURLPattern) RawJSON() string { return r.JSON.raw }
func (r *GuideActivationURLPattern) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GuideActivationURLPattern to a
// GuideActivationURLPatternParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GuideActivationURLPatternParam.Overrides()
func (r GuideActivationURLPattern) ToParam() GuideActivationURLPatternParam {
	return param.Override[GuideActivationURLPatternParam](json.RawMessage(r.RawJSON()))
}

// Whether to allow or block the guide at the specified location.
type GuideActivationURLPatternDirective string

const (
	GuideActivationURLPatternDirectiveAllow GuideActivationURLPatternDirective = "allow"
	GuideActivationURLPatternDirectiveBlock GuideActivationURLPatternDirective = "block"
)

// A rule that controls when a guide should be shown based on the user's location
// in the application. At least one of `pathname` or `search` must be provided.
//
// The property Directive is required.
type GuideActivationURLPatternParam struct {
	// Whether to allow or block the guide at the specified location.
	//
	// Any of "allow", "block".
	Directive GuideActivationURLPatternDirective `json:"directive,omitzero" api:"required"`
	// The URL pathname pattern to match against. Must be a valid URI path.
	Pathname param.Opt[string] `json:"pathname,omitzero"`
	// The URL query string pattern to match against (without the leading '?').
	// Supports URLPattern API syntax.
	Search param.Opt[string] `json:"search,omitzero"`
	paramObj
}

func (r GuideActivationURLPatternParam) MarshalJSON() (data []byte, err error) {
	type shadow GuideActivationURLPatternParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GuideActivationURLPatternParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A step in a guide that corresponds to a piece of UI and its content.
type GuideStep struct {
	// The unique reference string for the step. Must be at minimum 3 characters and at
	// maximum 255 characters in length. Must be in the format of ^[a-z0-9_-]+$.
	Ref string `json:"ref" api:"required"`
	// The key of the template schema to use for this step.
	SchemaKey string `json:"schema_key" api:"required"`
	// The semantic version of the template schema to use.
	SchemaSemver string `json:"schema_semver" api:"required"`
	// The key of the template schema variant to use.
	SchemaVariantKey string `json:"schema_variant_key" api:"required"`
	// A name for the step.
	Name string `json:"name"`
	// A map of values that make up the step's content. Each value must conform to its
	// respective template schema field settings.
	Values map[string]any `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ref              respjson.Field
		SchemaKey        respjson.Field
		SchemaSemver     respjson.Field
		SchemaVariantKey respjson.Field
		Name             respjson.Field
		Values           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuideStep) RawJSON() string { return r.JSON.raw }
func (r *GuideStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GuideStep to a GuideStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GuideStepParam.Overrides()
func (r GuideStep) ToParam() GuideStepParam {
	return param.Override[GuideStepParam](json.RawMessage(r.RawJSON()))
}

// A step in a guide that corresponds to a piece of UI and its content.
//
// The properties Ref, SchemaKey, SchemaSemver, SchemaVariantKey are required.
type GuideStepParam struct {
	// The unique reference string for the step. Must be at minimum 3 characters and at
	// maximum 255 characters in length. Must be in the format of ^[a-z0-9_-]+$.
	Ref string `json:"ref" api:"required"`
	// The key of the template schema to use for this step.
	SchemaKey string `json:"schema_key" api:"required"`
	// The semantic version of the template schema to use.
	SchemaSemver string `json:"schema_semver" api:"required"`
	// The key of the template schema variant to use.
	SchemaVariantKey string `json:"schema_variant_key" api:"required"`
	// A name for the step.
	Name param.Opt[string] `json:"name,omitzero"`
	// A map of values that make up the step's content. Each value must conform to its
	// respective template schema field settings.
	Values map[string]any `json:"values,omitzero"`
	paramObj
}

func (r GuideStepParam) MarshalJSON() (data []byte, err error) {
	type shadow GuideStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GuideStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Guide response under the `guide` key.
type GuideActivateResponse struct {
	// A guide defines an in-app guide that can be displayed to users based on priority
	// and other conditions.
	Guide Guide `json:"guide" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Guide       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuideActivateResponse) RawJSON() string { return r.JSON.raw }
func (r *GuideActivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response from archiving a guide.
type GuideArchiveResponse struct {
	// The result of the promote operation.
	Result string `json:"result" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuideArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *GuideArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Guide response under the `guide` key.
type GuideUpsertResponse struct {
	// A guide defines an in-app guide that can be displayed to users based on priority
	// and other conditions.
	Guide Guide `json:"guide" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Guide       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuideUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *GuideUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Guide response under the `guide` key.
type GuideValidateResponse struct {
	// A guide defines an in-app guide that can be displayed to users based on priority
	// and other conditions.
	Guide Guide `json:"guide" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Guide       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuideValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *GuideValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuideGetParams struct {
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

// URLQuery serializes [GuideGetParams]'s query parameters as `url.Values`.
func (r GuideGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GuideListParams struct {
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

// URLQuery serializes [GuideListParams]'s query parameters as `url.Values`.
func (r GuideListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GuideActivateParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. A
	// request to activate or deactivate a guide.
	OfGuideBooleanActivations *GuideActivateParamsBodyGuideBooleanActivationParams `json:",inline"`
	// This field is a request body variant, only one variant field can be set. A
	// request to schedule the activation of a guide. At least one of from or until
	// must be provided.
	OfGuideScheduledActivations *GuideActivateParamsBodyGuideScheduledActivationParams `json:",inline"`

	paramObj
}

func (u GuideActivateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfGuideBooleanActivations, u.OfGuideScheduledActivations)
}
func (r *GuideActivateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [GuideActivateParams]'s query parameters as `url.Values`.
func (r GuideActivateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A request to activate or deactivate a guide.
//
// The property Status is required.
type GuideActivateParamsBodyGuideBooleanActivationParams struct {
	// Whether to activate or deactivate the guide.
	Status bool `json:"status" api:"required"`
	paramObj
}

func (r GuideActivateParamsBodyGuideBooleanActivationParams) MarshalJSON() (data []byte, err error) {
	type shadow GuideActivateParamsBodyGuideBooleanActivationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GuideActivateParamsBodyGuideBooleanActivationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to schedule the activation of a guide. At least one of from or until
// must be provided.
type GuideActivateParamsBodyGuideScheduledActivationParams struct {
	// When to activate the guide. If provided, the guide will be scheduled to activate
	// at this time. Must be in ISO 8601 UTC format.
	From param.Opt[time.Time] `json:"from,omitzero" format:"date-time"`
	// When to deactivate the guide. If provided, the guide will be scheduled to
	// deactivate at this time. Must be in ISO 8601 UTC format.
	Until param.Opt[time.Time] `json:"until,omitzero" format:"date-time"`
	paramObj
}

func (r GuideActivateParamsBodyGuideScheduledActivationParams) MarshalJSON() (data []byte, err error) {
	type shadow GuideActivateParamsBodyGuideScheduledActivationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GuideActivateParamsBodyGuideScheduledActivationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuideUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A request to create or update a guide.
	Guide GuideUpsertParamsGuide `json:"guide,omitzero" api:"required"`
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

func (r GuideUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow GuideUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GuideUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [GuideUpsertParams]'s query parameters as `url.Values`.
func (r GuideUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A request to create or update a guide.
//
// The properties ChannelKey, Name, Steps are required.
type GuideUpsertParamsGuide struct {
	// The key of the channel in which the guide exists.
	ChannelKey string `json:"channel_key" api:"required"`
	// A name for the guide. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// A list of guide step objects in the guide.
	Steps []GuideStepParam `json:"steps,omitzero" api:"required"`
	// The timestamp of when the guide was archived.
	ArchivedAt param.Opt[time.Time] `json:"archived_at,omitzero" format:"date-time"`
	// The timestamp of when the guide was deleted.
	DeletedAt param.Opt[time.Time] `json:"deleted_at,omitzero" format:"date-time"`
	// An arbitrary string attached to a guide object. Useful for adding notes about
	// the guide for internal purposes. Maximum of 280 characters allowed.
	Description param.Opt[string] `json:"description,omitzero"`
	// The key of the target audience for the guide. When not set, will default to
	// targeting all users.
	TargetAudienceKey param.Opt[string] `json:"target_audience_key,omitzero"`
	// A list of activation url patterns that describe when the guide should be shown.
	ActivationURLPatterns []GuideActivationURLPatternParam `json:"activation_url_patterns,omitzero"`
	// Use tags to organize resources internally within your account. For example, by
	// team or product area.
	Tags []string `json:"tags,omitzero"`
	// A group of conditions to be evaluated.
	TargetPropertyConditions ConditionGroupUnionParam `json:"target_property_conditions,omitzero"`
	paramObj
}

func (r GuideUpsertParamsGuide) MarshalJSON() (data []byte, err error) {
	type shadow GuideUpsertParamsGuide
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GuideUpsertParamsGuide) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuideValidateParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A request to create or update a guide.
	Guide GuideValidateParamsGuide `json:"guide,omitzero" api:"required"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	paramObj
}

func (r GuideValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow GuideValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GuideValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [GuideValidateParams]'s query parameters as `url.Values`.
func (r GuideValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A request to create or update a guide.
//
// The properties ChannelKey, Name, Steps are required.
type GuideValidateParamsGuide struct {
	// The key of the channel in which the guide exists.
	ChannelKey string `json:"channel_key" api:"required"`
	// A name for the guide. Must be at maximum 255 characters in length.
	Name string `json:"name" api:"required"`
	// A list of guide step objects in the guide.
	Steps []GuideStepParam `json:"steps,omitzero" api:"required"`
	// The timestamp of when the guide was archived.
	ArchivedAt param.Opt[time.Time] `json:"archived_at,omitzero" format:"date-time"`
	// The timestamp of when the guide was deleted.
	DeletedAt param.Opt[time.Time] `json:"deleted_at,omitzero" format:"date-time"`
	// An arbitrary string attached to a guide object. Useful for adding notes about
	// the guide for internal purposes. Maximum of 280 characters allowed.
	Description param.Opt[string] `json:"description,omitzero"`
	// The key of the target audience for the guide. When not set, will default to
	// targeting all users.
	TargetAudienceKey param.Opt[string] `json:"target_audience_key,omitzero"`
	// A list of activation url patterns that describe when the guide should be shown.
	ActivationURLPatterns []GuideActivationURLPatternParam `json:"activation_url_patterns,omitzero"`
	// Use tags to organize resources internally within your account. For example, by
	// team or product area.
	Tags []string `json:"tags,omitzero"`
	// A group of conditions to be evaluated.
	TargetPropertyConditions ConditionGroupUnionParam `json:"target_property_conditions,omitzero"`
	paramObj
}

func (r GuideValidateParamsGuide) MarshalJSON() (data []byte, err error) {
	type shadow GuideValidateParamsGuide
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GuideValidateParamsGuide) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
