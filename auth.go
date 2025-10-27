// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/knock-mapi-go/internal/apijson"
	"github.com/stainless-sdks/knock-mapi-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-mapi-go/option"
	"github.com/stainless-sdks/knock-mapi-go/packages/respjson"
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

// Return information about the current calling scope. Will either be a service
// token or from an OAuth context.
func (r *AuthService) Verify(ctx context.Context, opts ...option.RequestOption) (res *AuthVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whoami"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Information about the current calling scope.
type AuthVerifyResponse struct {
	// Account plan features and limits.
	AccountFeatures AuthVerifyResponseAccountFeatures `json:"account_features,required"`
	// The display name of the account.
	AccountName string `json:"account_name,required"`
	// The unique slug identifier for the account.
	AccountSlug string `json:"account_slug,required"`
	// The type of authentication context - either a service token or OAuth user
	// context.
	//
	// Any of "service_token", "oauth_context".
	Type AuthVerifyResponseType `json:"type,required"`
	// The name of the service token if authenticated via service token, null for OAuth
	// contexts.
	ServiceTokenName string `json:"service_token_name,nullable"`
	// The ID of the authenticated user if in OAuth context, null for service token
	// contexts.
	UserID string `json:"user_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountFeatures  respjson.Field
		AccountName      respjson.Field
		AccountSlug      respjson.Field
		Type             respjson.Field
		ServiceTokenName respjson.Field
		UserID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account plan features and limits.
type AuthVerifyResponseAccountFeatures struct {
	// Whether batch rendering limits can be configured.
	BatchItemsRenderLimitAllowed bool `json:"batch_items_render_limit_allowed"`
	// Whether custom branding can be applied to notifications.
	CustomBrandingAllowed bool `json:"custom_branding_allowed"`
	// Number of days data is retained, null for unlimited retention.
	DataRetentionDays int64 `json:"data_retention_days,nullable"`
	// Whether data warehouse integration extensions are available.
	DataWarehouseExtensionAllowed bool `json:"data_warehouse_extension_allowed"`
	// Whether Datadog integration extension is available.
	DatadogExtensionAllowed bool `json:"datadog_extension_allowed"`
	// Whether directory sync functionality is available.
	DsyncAllowed bool `json:"dsync_allowed"`
	// Monthly limit for guide notification recipients, null for unlimited.
	GuidesMonthlyNotifiedRecipientsLimit int64 `json:"guides_monthly_notified_recipients_limit,nullable"`
	// Whether Heap integration extension is available.
	HeapExtensionAllowed bool `json:"heap_extension_allowed"`
	// Whether Knock branding is required to be displayed.
	KnockBrandingRequired bool `json:"knock_branding_required"`
	// Whether Litmus email preview integration is available.
	LitmusEmailPreviewAllowed bool `json:"litmus_email_preview_allowed"`
	// Monthly limit for messages sent, null for unlimited.
	MessageSentLimit int64 `json:"message_sent_limit,nullable"`
	// Whether New Relic integration extension is available.
	NewRelicExtensionAllowed bool `json:"new_relic_extension_allowed"`
	// Whether Segment integration extension is available.
	SegmentExtensionAllowed bool `json:"segment_extension_allowed"`
	// Whether self-service account management features are available.
	SelfServeAllowed bool `json:"self_serve_allowed"`
	// Whether single sign-on (SSO) is enabled for the account.
	SSOAllowed bool `json:"sso_allowed"`
	// Whether tenant-level preferences are supported.
	TenantPreferencesAllowed bool `json:"tenant_preferences_allowed"`
	// Whether multi-language translations are supported.
	TranslationsAllowed bool `json:"translations_allowed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchItemsRenderLimitAllowed         respjson.Field
		CustomBrandingAllowed                respjson.Field
		DataRetentionDays                    respjson.Field
		DataWarehouseExtensionAllowed        respjson.Field
		DatadogExtensionAllowed              respjson.Field
		DsyncAllowed                         respjson.Field
		GuidesMonthlyNotifiedRecipientsLimit respjson.Field
		HeapExtensionAllowed                 respjson.Field
		KnockBrandingRequired                respjson.Field
		LitmusEmailPreviewAllowed            respjson.Field
		MessageSentLimit                     respjson.Field
		NewRelicExtensionAllowed             respjson.Field
		SegmentExtensionAllowed              respjson.Field
		SelfServeAllowed                     respjson.Field
		SSOAllowed                           respjson.Field
		TenantPreferencesAllowed             respjson.Field
		TranslationsAllowed                  respjson.Field
		ExtraFields                          map[string]respjson.Field
		raw                                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthVerifyResponseAccountFeatures) RawJSON() string { return r.JSON.raw }
func (r *AuthVerifyResponseAccountFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of authentication context - either a service token or OAuth user
// context.
type AuthVerifyResponseType string

const (
	AuthVerifyResponseTypeServiceToken AuthVerifyResponseType = "service_token"
	AuthVerifyResponseTypeOAuthContext AuthVerifyResponseType = "oauth_context"
)
