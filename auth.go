// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/knock-mapi-go/internal/apijson"
	"github.com/stainless-sdks/knock-mapi-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-mapi-go/option"
	"github.com/stainless-sdks/knock-mapi-go/packages/resp"
)

// AuthService contains methods and other services that help with interacting with
// the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	Options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r AuthService) {
	r = AuthService{}
	r.Options = opts
	return
}

// Return information about the current service token.
func (r *AuthService) Verify(ctx context.Context, opts ...option.RequestOption) (res *AuthVerifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/whoami"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Information about the current service token.
type AuthVerifyResponse struct {
	AccountName      string `json:"account_name,required"`
	AccountSlug      string `json:"account_slug,required"`
	ServiceTokenName string `json:"service_token_name,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AccountName      resp.Field
		AccountSlug      resp.Field
		ServiceTokenName resp.Field
		ExtraFields      map[string]resp.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
