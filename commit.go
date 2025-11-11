// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
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

// CommitService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommitService] method instead.
type CommitService struct {
	Options []option.RequestOption
}

// NewCommitService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCommitService(opts ...option.RequestOption) (r CommitService) {
	r = CommitService{}
	r.Options = opts
	return
}

// Retrieve a single commit by its ID.
func (r *CommitService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Commit, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/commits/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a paginated list of commits in a given environment. The commits are
// ordered from most recent first.
func (r *CommitService) List(ctx context.Context, query CommitListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Commit], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/commits"
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

// Returns a paginated list of commits in a given environment. The commits are
// ordered from most recent first.
func (r *CommitService) ListAutoPaging(ctx context.Context, query CommitListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Commit] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Commit all changes across all resources in the development environment.
func (r *CommitService) CommitAll(ctx context.Context, body CommitCommitAllParams, opts ...option.RequestOption) (res *CommitCommitAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/commits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Promote all changes across all resources to the target environment from its
// preceding environment.
func (r *CommitService) PromoteAll(ctx context.Context, body CommitPromoteAllParams, opts ...option.RequestOption) (res *CommitPromoteAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/commits/promote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Promotes one change to the subsequent environment.
func (r *CommitService) PromoteOne(ctx context.Context, id string, opts ...option.RequestOption) (res *CommitPromoteOneResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/commits/%s/promote", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// A commit is a change to a resource within an environment, made by an author.
type Commit struct {
	// The unique identifier for the commit.
	ID string `json:"id,required" format:"uuid"`
	// The author of the commit.
	Author CommitAuthor `json:"author,required"`
	// The timestamp of when the commit was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The environment of the commit.
	Environment string `json:"environment,required"`
	// The resource object associated with the commit.
	Resource CommitResource `json:"resource,required"`
	// The optional message about the commit.
	CommitMessage string `json:"commit_message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Author        respjson.Field
		CreatedAt     respjson.Field
		Environment   respjson.Field
		Resource      respjson.Field
		CommitMessage respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Commit) RawJSON() string { return r.JSON.raw }
func (r *Commit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The author of the commit.
type CommitAuthor struct {
	// The email address of the commit author.
	Email string `json:"email,required"`
	// The name of the commit author.
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitAuthor) RawJSON() string { return r.JSON.raw }
func (r *CommitAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The resource object associated with the commit.
type CommitResource struct {
	// The unique identifier for the resource.
	Identifier string `json:"identifier,required"`
	// The type of the resource object.
	//
	// Any of "audience", "email_layout", "guide", "message_type", "partial",
	// "translation", "workflow".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitResource) RawJSON() string { return r.JSON.raw }
func (r *CommitResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response from committing all changes.
type CommitCommitAllResponse struct {
	// The result of the commit operation.
	Result string `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitCommitAllResponse) RawJSON() string { return r.JSON.raw }
func (r *CommitCommitAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response from promoting all changes.
type CommitPromoteAllResponse struct {
	// The result of the promote operation.
	Result string `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitPromoteAllResponse) RawJSON() string { return r.JSON.raw }
func (r *CommitPromoteAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Commit response under the `commit` key.
type CommitPromoteOneResponse struct {
	// A commit is a change to a resource within an environment, made by an author.
	Commit Commit `json:"commit,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commit      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitPromoteOneResponse) RawJSON() string { return r.JSON.raw }
func (r *CommitPromoteOneResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommitListParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Whether to show commits in the given environment that have not been promoted to
	// the subsequent environment (false) or commits which have been promoted (true).
	Promoted param.Opt[bool] `query:"promoted,omitzero" json:"-"`
	// Filter commits by resource identifier. Must be used together with resource_type.
	// For most resources, this will be the resource key. In the case of translations,
	// this will be the locale code and namespace, separated by a `/`. For example,
	// `en/courses` or `en`.
	ResourceID param.Opt[string] `query:"resource_id,omitzero" json:"-"`
	// Filter commits by resource type(s). Accepts a single type or array of types. Can
	// be combined with resource_id to filter for specific resources.
	ResourceType CommitListParamsResourceTypeUnion `query:"resource_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CommitListParams]'s query parameters as `url.Values`.
func (r CommitListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CommitListParamsResourceTypeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCommitListsResourceTypeString)
	OfCommitListsResourceTypeString         param.Opt[string] `query:",omitzero,inline"`
	OfCommitListsResourceTypeArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *CommitListParamsResourceTypeUnion) asAny() any {
	if !param.IsOmitted(u.OfCommitListsResourceTypeString) {
		return &u.OfCommitListsResourceTypeString
	} else if !param.IsOmitted(u.OfCommitListsResourceTypeArrayItemArray) {
		return &u.OfCommitListsResourceTypeArrayItemArray
	}
	return nil
}

type CommitListParamsResourceTypeString string

const (
	CommitListParamsResourceTypeStringAudience    CommitListParamsResourceTypeString = "audience"
	CommitListParamsResourceTypeStringEmailLayout CommitListParamsResourceTypeString = "email_layout"
	CommitListParamsResourceTypeStringGuide       CommitListParamsResourceTypeString = "guide"
	CommitListParamsResourceTypeStringMessageType CommitListParamsResourceTypeString = "message_type"
	CommitListParamsResourceTypeStringPartial     CommitListParamsResourceTypeString = "partial"
	CommitListParamsResourceTypeStringTranslation CommitListParamsResourceTypeString = "translation"
	CommitListParamsResourceTypeStringWorkflow    CommitListParamsResourceTypeString = "workflow"
)

type CommitCommitAllParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// An optional message to include in a commit.
	CommitMessage param.Opt[string] `query:"commit_message,omitzero" json:"-"`
	// Filter changes to commit by resource identifier. Must be used together with
	// resource_type.
	ResourceID param.Opt[string] `query:"resource_id,omitzero" json:"-"`
	// Filter changes to commit by resource type(s). Accepts a single type or array of
	// types. Can be combined with resource_id to filter for specific resources.
	ResourceType CommitCommitAllParamsResourceTypeUnion `query:"resource_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CommitCommitAllParams]'s query parameters as `url.Values`.
func (r CommitCommitAllParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CommitCommitAllParamsResourceTypeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCommitCommitAllsResourceTypeString)
	OfCommitCommitAllsResourceTypeString         param.Opt[string] `query:",omitzero,inline"`
	OfCommitCommitAllsResourceTypeArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *CommitCommitAllParamsResourceTypeUnion) asAny() any {
	if !param.IsOmitted(u.OfCommitCommitAllsResourceTypeString) {
		return &u.OfCommitCommitAllsResourceTypeString
	} else if !param.IsOmitted(u.OfCommitCommitAllsResourceTypeArrayItemArray) {
		return &u.OfCommitCommitAllsResourceTypeArrayItemArray
	}
	return nil
}

type CommitCommitAllParamsResourceTypeString string

const (
	CommitCommitAllParamsResourceTypeStringAudience    CommitCommitAllParamsResourceTypeString = "audience"
	CommitCommitAllParamsResourceTypeStringEmailLayout CommitCommitAllParamsResourceTypeString = "email_layout"
	CommitCommitAllParamsResourceTypeStringGuide       CommitCommitAllParamsResourceTypeString = "guide"
	CommitCommitAllParamsResourceTypeStringMessageType CommitCommitAllParamsResourceTypeString = "message_type"
	CommitCommitAllParamsResourceTypeStringPartial     CommitCommitAllParamsResourceTypeString = "partial"
	CommitCommitAllParamsResourceTypeStringTranslation CommitCommitAllParamsResourceTypeString = "translation"
	CommitCommitAllParamsResourceTypeStringWorkflow    CommitCommitAllParamsResourceTypeString = "workflow"
)

type CommitPromoteAllParams struct {
	// A slug of the target environment to which you want to promote all changes from
	// its directly preceding environment.
	//
	// For example, if you have three environments “development”, “staging”, and
	// “production” (in that order), setting this param to “production” will promote
	// all commits not currently in production from staging.
	//
	// Note: This must be a non-development environment.
	ToEnvironment string `query:"to_environment,required" json:"-"`
	// Filter commits to promote by resource identifier. Must be used together with
	// resource_type.
	ResourceID param.Opt[string] `query:"resource_id,omitzero" json:"-"`
	// Filter commits to promote by resource type(s). Accepts a single type or array of
	// types. Can be combined with resource_id to filter for specific resources.
	ResourceType CommitPromoteAllParamsResourceTypeUnion `query:"resource_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CommitPromoteAllParams]'s query parameters as `url.Values`.
func (r CommitPromoteAllParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CommitPromoteAllParamsResourceTypeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCommitPromoteAllsResourceTypeString)
	OfCommitPromoteAllsResourceTypeString         param.Opt[string] `query:",omitzero,inline"`
	OfCommitPromoteAllsResourceTypeArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *CommitPromoteAllParamsResourceTypeUnion) asAny() any {
	if !param.IsOmitted(u.OfCommitPromoteAllsResourceTypeString) {
		return &u.OfCommitPromoteAllsResourceTypeString
	} else if !param.IsOmitted(u.OfCommitPromoteAllsResourceTypeArrayItemArray) {
		return &u.OfCommitPromoteAllsResourceTypeArrayItemArray
	}
	return nil
}

type CommitPromoteAllParamsResourceTypeString string

const (
	CommitPromoteAllParamsResourceTypeStringAudience    CommitPromoteAllParamsResourceTypeString = "audience"
	CommitPromoteAllParamsResourceTypeStringEmailLayout CommitPromoteAllParamsResourceTypeString = "email_layout"
	CommitPromoteAllParamsResourceTypeStringGuide       CommitPromoteAllParamsResourceTypeString = "guide"
	CommitPromoteAllParamsResourceTypeStringMessageType CommitPromoteAllParamsResourceTypeString = "message_type"
	CommitPromoteAllParamsResourceTypeStringPartial     CommitPromoteAllParamsResourceTypeString = "partial"
	CommitPromoteAllParamsResourceTypeStringTranslation CommitPromoteAllParamsResourceTypeString = "translation"
	CommitPromoteAllParamsResourceTypeStringWorkflow    CommitPromoteAllParamsResourceTypeString = "workflow"
)
