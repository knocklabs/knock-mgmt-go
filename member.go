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

// MemberService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemberService] method instead.
type MemberService struct {
	Options []option.RequestOption
}

// NewMemberService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMemberService(opts ...option.RequestOption) (r MemberService) {
	r = MemberService{}
	r.Options = opts
	return
}

// Returns a single member by their ID.
func (r *MemberService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Member, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/members/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a paginated list of members for the current account. Optionally filter
// by role.
func (r *MemberService) List(ctx context.Context, query MemberListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Member], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/members"
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

// Returns a paginated list of members for the current account. Optionally filter
// by role.
func (r *MemberService) ListAutoPaging(ctx context.Context, query MemberListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Member] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Removes a member from the account.
func (r *MemberService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/members/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A member of the account.
type Member struct {
	// The unique identifier of the member.
	ID string `json:"id" api:"required" format:"uuid"`
	// The timestamp of when the member joined the account.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The member's role in the account.
	//
	// Any of "owner", "admin", "member", "production_only_member", "billing",
	// "support".
	Role MemberRole `json:"role" api:"required"`
	// The timestamp of when the member was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Information about a user within the Knock dashboard. Not to be confused with an
	// external user (recipient) of a workflow.
	User MemberUser `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Role        respjson.Field
		UpdatedAt   respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Member) RawJSON() string { return r.JSON.raw }
func (r *Member) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The member's role in the account.
type MemberRole string

const (
	MemberRoleOwner                MemberRole = "owner"
	MemberRoleAdmin                MemberRole = "admin"
	MemberRoleMember               MemberRole = "member"
	MemberRoleProductionOnlyMember MemberRole = "production_only_member"
	MemberRoleBilling              MemberRole = "billing"
	MemberRoleSupport              MemberRole = "support"
)

// Information about a user within the Knock dashboard. Not to be confused with an
// external user (recipient) of a workflow.
type MemberUser struct {
	// The user's unique identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// The timestamp of when the user was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The user's email address.
	Email string `json:"email" api:"required" format:"email"`
	// The timestamp of when the user was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The URL of the user's avatar image.
	AvatarURL string `json:"avatar_url" api:"nullable"`
	// The user's display name.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Email       respjson.Field
		UpdatedAt   respjson.Field
		AvatarURL   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemberUser) RawJSON() string { return r.JSON.raw }
func (r *MemberUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemberListParams struct {
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter members by email address (exact match).
	Email param.Opt[string] `query:"email,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter members by role. One of: owner, admin, member, production_only_member,
	// billing, support.
	Role param.Opt[string] `query:"role,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MemberListParams]'s query parameters as `url.Values`.
func (r MemberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
