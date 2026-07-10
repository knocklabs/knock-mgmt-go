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

// Branches in Knock are a way to isolate changes to your Knock resources.
//
// BranchService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBranchService] method instead.
type BranchService struct {
	Options []option.RequestOption
}

// NewBranchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBranchService(opts ...option.RequestOption) (r BranchService) {
	r = BranchService{}
	r.Options = opts
	return
}

// Creates a new branch off of the development environment with the given slug.
func (r *BranchService) New(ctx context.Context, branchSlug string, body BranchNewParams, opts ...option.RequestOption) (res *Branch, err error) {
	opts = slices.Concat(r.Options, opts)
	if branchSlug == "" {
		err = errors.New("missing required branch_slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/branches/%s", branchSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single branch by the `branch_slug`.
func (r *BranchService) Get(ctx context.Context, branchSlug string, query BranchGetParams, opts ...option.RequestOption) (res *Branch, err error) {
	opts = slices.Concat(r.Options, opts)
	if branchSlug == "" {
		err = errors.New("missing required branch_slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/branches/%s", branchSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of branches. The branches will be returned in order of
// their last commit time (newest first).
func (r *BranchService) List(ctx context.Context, query BranchListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Branch], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/branches"
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

// Returns a paginated list of branches. The branches will be returned in order of
// their last commit time (newest first).
func (r *BranchService) ListAutoPaging(ctx context.Context, query BranchListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Branch] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Deletes a branch by the `branch_slug`.
func (r *BranchService) Delete(ctx context.Context, branchSlug string, body BranchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if branchSlug == "" {
		err = errors.New("missing required branch_slug parameter")
		return err
	}
	path := fmt.Sprintf("v1/branches/%s", branchSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// A branch object.
type Branch struct {
	// The timestamp of when the branch was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A unique slug for the branch. Cannot exceed 255 characters.
	Slug string `json:"slug" api:"required"`
	// The timestamp of when the branch was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The timestamp of when the branch was deleted.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// The timestamp of the most-recent commit in the branch.
	LastCommitAt time.Time `json:"last_commit_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		Slug         respjson.Field
		UpdatedAt    respjson.Field
		DeletedAt    respjson.Field
		LastCommitAt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Branch) RawJSON() string { return r.JSON.raw }
func (r *Branch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BranchNewParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [BranchNewParams]'s query parameters as `url.Values`.
func (r BranchNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BranchGetParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [BranchGetParams]'s query parameters as `url.Values`.
func (r BranchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BranchListParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BranchListParams]'s query parameters as `url.Values`.
func (r BranchListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BranchDeleteParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [BranchDeleteParams]'s query parameters as `url.Values`.
func (r BranchDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
