// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/knocklabs/knock-mgmt-go/internal/apijson"
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// Headless accounts let an agent create a Knock account on its own and hand it to
// a person to claim.
//
// HeadlessAccountService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHeadlessAccountService] method instead.
type HeadlessAccountService struct {
	Options []option.RequestOption
}

// NewHeadlessAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHeadlessAccountService(opts ...option.RequestOption) (r HeadlessAccountService) {
	r = HeadlessAccountService{}
	r.Options = opts
	return
}

// Creates an account that has no owner yet, and returns a service token the agent
// uses to work in it. The account is deleted after 24 hours unless a person claims
// it. This endpoint needs no authentication and is rate limited.
func (r *HeadlessAccountService) New(ctx context.Context, opts ...option.RequestOption) (res *HeadlessAccountNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/headless_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Starts handing the account to a person. Returns a link and a six-digit code for
// the agent to show that person. They open the link, sign in with the email given
// here, and type the code to become the account's owner. The service token keeps
// working afterwards. Starting a new claim replaces the previous link and code.
func (r *HeadlessAccountService) Claim(ctx context.Context, body HeadlessAccountClaimParams, opts ...option.RequestOption) (res *HeadlessAccountClaimResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/headless_accounts/claim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns whether the headless account is still being set up, waiting to be
// claimed, claimed by a person, or past its claim window. Agents poll this
// endpoint to find out when a person has claimed the account.
func (r *HeadlessAccountService) GetStatus(ctx context.Context, opts ...option.RequestOption) (res *HeadlessAccountGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/headless_accounts/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A headless account an agent created, with the service token the agent uses to
// work in it.
type HeadlessAccountNewResponse struct {
	// A service token for the Management API. It is only returned once and keeps
	// working after a person claims the account.
	Token string `json:"token" api:"required"`
	// When the account and its token are deleted unless a person claims the account
	// first.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// The account is still being set up. Poll the status endpoint.
	//
	// Any of "provisioning".
	Status HeadlessAccountNewResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExpiresAt   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeadlessAccountNewResponse) RawJSON() string { return r.JSON.raw }
func (r *HeadlessAccountNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The account is still being set up. Poll the status endpoint.
type HeadlessAccountNewResponseStatus string

const (
	HeadlessAccountNewResponseStatusProvisioning HeadlessAccountNewResponseStatus = "provisioning"
)

// A claim an agent started for a person, with the link and code to show them.
type HeadlessAccountClaimResponse struct {
	// The dashboard link the person opens and signs in with the email given when the
	// claim started.
	ClaimURL string `json:"claim_url" api:"required"`
	// When the link and code stop working.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// A six-digit code to show the person, who types it on the claim page. Show it
	// separately from the link.
	UserCode string `json:"user_code" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClaimURL    respjson.Field
		ExpiresAt   respjson.Field
		UserCode    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeadlessAccountClaimResponse) RawJSON() string { return r.JSON.raw }
func (r *HeadlessAccountClaimResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where a headless account stands, from the point of view of the agent.
type HeadlessAccountGetStatusResponse struct {
	// When an unclaimed account is deleted, or null once it has been claimed.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Whether the account is still being set up, waiting to be claimed, claimed by a
	// person, or expired because its claim window closed or its setup never finished.
	//
	// Any of "provisioning", "unclaimed", "claimed", "expired".
	Status HeadlessAccountGetStatusResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeadlessAccountGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *HeadlessAccountGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the account is still being set up, waiting to be claimed, claimed by a
// person, or expired because its claim window closed or its setup never finished.
type HeadlessAccountGetStatusResponseStatus string

const (
	HeadlessAccountGetStatusResponseStatusProvisioning HeadlessAccountGetStatusResponseStatus = "provisioning"
	HeadlessAccountGetStatusResponseStatusUnclaimed    HeadlessAccountGetStatusResponseStatus = "unclaimed"
	HeadlessAccountGetStatusResponseStatusClaimed      HeadlessAccountGetStatusResponseStatus = "claimed"
	HeadlessAccountGetStatusResponseStatusExpired      HeadlessAccountGetStatusResponseStatus = "expired"
)

type HeadlessAccountClaimParams struct {
	// The email of the person who will own the account. Only they can finish the
	// claim.
	Email string `json:"email" api:"required" format:"email"`
	paramObj
}

func (r HeadlessAccountClaimParams) MarshalJSON() (data []byte, err error) {
	type shadow HeadlessAccountClaimParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HeadlessAccountClaimParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
