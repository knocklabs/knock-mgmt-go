// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
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
	"github.com/stainless-sdks/knock-mapi-go/packages/respjson"
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	path := "v1/commits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Promote all changes across all resources to the target environment from its
// preceding environment.
func (r *CommitService) PromoteAll(ctx context.Context, body CommitPromoteAllParams, opts ...option.RequestOption) (res *CommitPromoteAllResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/commits/promote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Promotes one change to the subsequent environment.
func (r *CommitService) PromoteOne(ctx context.Context, id string, opts ...option.RequestOption) (res *CommitPromoteOneResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	CommitAuthor CommitCommitAuthor `json:"commit_author,required"`
	// The optional message about the commit.
	CommitMessage string `json:"commit_message,required"`
	// The timestamp of when the commit was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The environment of the commit.
	Environment string `json:"environment,required"`
	// The resource object associated with the commit.
	Resource CommitResource `json:"resource,required"`
	// The timestamp of when the commit was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CommitAuthor  respjson.Field
		CommitMessage respjson.Field
		CreatedAt     respjson.Field
		Environment   respjson.Field
		Resource      respjson.Field
		UpdatedAt     respjson.Field
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
type CommitCommitAuthor struct {
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
func (r CommitCommitAuthor) RawJSON() string { return r.JSON.raw }
func (r *CommitCommitAuthor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The resource object associated with the commit.
type CommitResource struct {
	// The unique identifier for the resource.
	Identifier string `json:"identifier,required"`
	// The type of the resource object.
	//
	// Any of "email_layout", "workflow", "translation", "partial", "message_type".
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
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Whether to show commits in the given environment that have not been promoted to
	// the subsequent environment (false) or commits which have been promoted (true).
	Promoted param.Opt[bool] `query:"promoted,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CommitListParams]'s query parameters as `url.Values`.
func (r CommitListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CommitCommitAllParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// An optional message to include in a commit.
	CommitMessage param.Opt[string] `query:"commit_message,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CommitCommitAllParams]'s query parameters as `url.Values`.
func (r CommitCommitAllParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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
	paramObj
}

// URLQuery serializes [CommitPromoteAllParams]'s query parameters as `url.Values`.
func (r CommitPromoteAllParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
