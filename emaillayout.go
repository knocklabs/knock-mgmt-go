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
	"github.com/knocklabs/knock-mgmt-go/shared"
)

// Email layouts wrap your email templates and provide a consistent look and feel.
//
// EmailLayoutService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailLayoutService] method instead.
type EmailLayoutService struct {
	Options []option.RequestOption
}

// NewEmailLayoutService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailLayoutService(opts ...option.RequestOption) (r EmailLayoutService) {
	r = EmailLayoutService{}
	r.Options = opts
	return
}

// Retrieve an email layout by its key, in a given environment.
func (r *EmailLayoutService) Get(ctx context.Context, emailLayoutKey string, query EmailLayoutGetParams, opts ...option.RequestOption) (res *EmailLayout, err error) {
	opts = slices.Concat(r.Options, opts)
	if emailLayoutKey == "" {
		err = errors.New("missing required email_layout_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/email_layouts/%s", emailLayoutKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of email layouts available in a given environment.
func (r *EmailLayoutService) List(ctx context.Context, query EmailLayoutListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[EmailLayout], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/email_layouts"
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

// Returns a paginated list of email layouts available in a given environment.
func (r *EmailLayoutService) ListAutoPaging(ctx context.Context, query EmailLayoutListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[EmailLayout] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Renders an email layout preview, without requiring a layout to be persisted
// within Knock. This is useful for previewing layouts in isolation, before saving
// them.
func (r *EmailLayoutService) Preview(ctx context.Context, params EmailLayoutPreviewParams, opts ...option.RequestOption) (res *EmailLayoutPreviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/email_layouts/preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Updates an email layout, or creates a new one if it does not yet exist.
//
// Note: this endpoint only operates in the "development" environment.
func (r *EmailLayoutService) Upsert(ctx context.Context, emailLayoutKey string, params EmailLayoutUpsertParams, opts ...option.RequestOption) (res *EmailLayoutUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if emailLayoutKey == "" {
		err = errors.New("missing required email_layout_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/email_layouts/%s", emailLayoutKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Validates an email layout payload without persisting it.
//
// Note: this endpoint only operates in the "development" environment.
func (r *EmailLayoutService) Validate(ctx context.Context, emailLayoutKey string, params EmailLayoutValidateParams, opts ...option.RequestOption) (res *EmailLayoutValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if emailLayoutKey == "" {
		err = errors.New("missing required email_layout_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/email_layouts/%s/validate", emailLayoutKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Overrides to apply against account branding variables in an email layout,
// including dark mode-specific values.
type BrandingOverrides struct {
	// A URL for a dark mode icon override.
	DarkIconURL string `json:"dark_icon_url" api:"nullable"`
	// A URL for a dark mode logo override.
	DarkLogoURL string `json:"dark_logo_url" api:"nullable"`
	// The dark mode primary brand color in hex format.
	DarkPrimaryColor string `json:"dark_primary_color" api:"nullable"`
	// The dark mode contrast color for the primary brand color in hex format.
	DarkPrimaryColorContrast string `json:"dark_primary_color_contrast" api:"nullable"`
	// A URL for a light mode icon override.
	IconURL string `json:"icon_url" api:"nullable"`
	// A URL for a light mode logo override.
	LogoURL string `json:"logo_url" api:"nullable"`
	// The light mode primary brand color in hex format.
	PrimaryColor string `json:"primary_color" api:"nullable"`
	// The light mode contrast color for the primary brand color in hex format.
	PrimaryColorContrast string `json:"primary_color_contrast" api:"nullable"`
	// The light mode primary text color in hex format.
	PrimaryTextColor string `json:"primary_text_color" api:"nullable"`
	// The light mode secondary text color in hex format.
	SecondaryTextColor string `json:"secondary_text_color" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DarkIconURL              respjson.Field
		DarkLogoURL              respjson.Field
		DarkPrimaryColor         respjson.Field
		DarkPrimaryColorContrast respjson.Field
		IconURL                  respjson.Field
		LogoURL                  respjson.Field
		PrimaryColor             respjson.Field
		PrimaryColorContrast     respjson.Field
		PrimaryTextColor         respjson.Field
		SecondaryTextColor       respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandingOverrides) RawJSON() string { return r.JSON.raw }
func (r *BrandingOverrides) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandingOverrides to a BrandingOverridesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandingOverridesParam.Overrides()
func (r BrandingOverrides) ToParam() BrandingOverridesParam {
	return param.Override[BrandingOverridesParam](json.RawMessage(r.RawJSON()))
}

// Overrides to apply against account branding variables in an email layout,
// including dark mode-specific values.
type BrandingOverridesParam struct {
	// A URL for a dark mode icon override.
	DarkIconURL param.Opt[string] `json:"dark_icon_url,omitzero"`
	// A URL for a dark mode logo override.
	DarkLogoURL param.Opt[string] `json:"dark_logo_url,omitzero"`
	// The dark mode primary brand color in hex format.
	DarkPrimaryColor param.Opt[string] `json:"dark_primary_color,omitzero"`
	// The dark mode contrast color for the primary brand color in hex format.
	DarkPrimaryColorContrast param.Opt[string] `json:"dark_primary_color_contrast,omitzero"`
	// A URL for a light mode icon override.
	IconURL param.Opt[string] `json:"icon_url,omitzero"`
	// A URL for a light mode logo override.
	LogoURL param.Opt[string] `json:"logo_url,omitzero"`
	// The light mode primary brand color in hex format.
	PrimaryColor param.Opt[string] `json:"primary_color,omitzero"`
	// The light mode contrast color for the primary brand color in hex format.
	PrimaryColorContrast param.Opt[string] `json:"primary_color_contrast,omitzero"`
	// The light mode primary text color in hex format.
	PrimaryTextColor param.Opt[string] `json:"primary_text_color,omitzero"`
	// The light mode secondary text color in hex format.
	SecondaryTextColor param.Opt[string] `json:"secondary_text_color,omitzero"`
	paramObj
}

func (r BrandingOverridesParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandingOverridesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandingOverridesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A versioned email layout used within an environment.
type EmailLayout struct {
	// The timestamp of when the email layout was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The complete HTML or MJML content of the email layout.
	HTMLLayout string `json:"html_layout" api:"required"`
	// The unique key for this email layout.
	Key string `json:"key" api:"required"`
	// The human-readable name of this email layout.
	Name string `json:"name" api:"required"`
	// The SHA of the email layout.
	Sha string `json:"sha" api:"required"`
	// The complete plaintext content of the email layout.
	TextLayout string `json:"text_layout" api:"required"`
	// Overrides to apply against account branding variables in an email layout,
	// including dark mode-specific values.
	BrandingOverrides BrandingOverrides `json:"branding_overrides" api:"nullable"`
	// The environment of the email layout.
	Environment string `json:"environment"`
	// A list of one or more items to show in the footer of the email layout.
	FooterLinks []EmailLayoutFooterLink `json:"footer_links"`
	// Whether this layout uses MJML format. When true, html_layout must contain <mjml>
	// tags.
	IsMjml bool `json:"is_mjml"`
	// The timestamp of when the email layout was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt         respjson.Field
		HTMLLayout        respjson.Field
		Key               respjson.Field
		Name              respjson.Field
		Sha               respjson.Field
		TextLayout        respjson.Field
		BrandingOverrides respjson.Field
		Environment       respjson.Field
		FooterLinks       respjson.Field
		IsMjml            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailLayout) RawJSON() string { return r.JSON.raw }
func (r *EmailLayout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailLayoutFooterLink struct {
	// The text to display as the link.
	Text string `json:"text" api:"required"`
	// The URL to link to.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailLayoutFooterLink) RawJSON() string { return r.JSON.raw }
func (r *EmailLayoutFooterLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to update or create an email layout.
//
// The properties HTMLLayout, Name, TextLayout are required.
type EmailLayoutRequestParam struct {
	// The complete HTML or MJML content of the email layout.
	HTMLLayout string `json:"html_layout" api:"required"`
	// The friendly name of this email layout.
	Name string `json:"name" api:"required"`
	// The complete plain text content of the email layout.
	TextLayout string `json:"text_layout" api:"required"`
	// Whether this layout uses MJML format. When true, html_layout must contain <mjml>
	// tags.
	IsMjml param.Opt[bool] `json:"is_mjml,omitzero"`
	// Overrides to apply against account branding variables in an email layout,
	// including dark mode-specific values.
	BrandingOverrides BrandingOverridesParam `json:"branding_overrides,omitzero"`
	// A list of one or more items to show in the footer of the email layout.
	FooterLinks []EmailLayoutRequestFooterLinkParam `json:"footer_links,omitzero"`
	paramObj
}

func (r EmailLayoutRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailLayoutRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailLayoutRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Text, URL are required.
type EmailLayoutRequestFooterLinkParam struct {
	// The text to display as the link.
	Text string `json:"text" api:"required"`
	// The URL to link to.
	URL string `json:"url" api:"required"`
	paramObj
}

func (r EmailLayoutRequestFooterLinkParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailLayoutRequestFooterLinkParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailLayoutRequestFooterLinkParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A response to an email layout preview request.
type EmailLayoutPreviewResponse struct {
	// The result of the preview.
	//
	// Any of "success", "error".
	Result EmailLayoutPreviewResponseResult `json:"result" api:"required"`
	// A list of errors encountered during rendering. Present when result is "error".
	Errors []EmailLayoutPreviewResponseError `json:"errors" api:"nullable"`
	// The rendered email layout, ready to be previewed.
	Layout EmailLayoutPreviewResponseLayout `json:"layout" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		Errors      respjson.Field
		Layout      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailLayoutPreviewResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailLayoutPreviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result of the preview.
type EmailLayoutPreviewResponseResult string

const (
	EmailLayoutPreviewResponseResultSuccess EmailLayoutPreviewResponseResult = "success"
	EmailLayoutPreviewResponseResultError   EmailLayoutPreviewResponseResult = "error"
)

// A rendering error with optional location information.
type EmailLayoutPreviewResponseError struct {
	// A human-readable description of the error.
	Message string `json:"message" api:"required"`
	// The layout field that caused the error, if available.
	Field string `json:"field" api:"nullable"`
	// The line number where the error occurred, if available.
	Line int64 `json:"line" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Field       respjson.Field
		Line        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailLayoutPreviewResponseError) RawJSON() string { return r.JSON.raw }
func (r *EmailLayoutPreviewResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The rendered email layout, ready to be previewed.
type EmailLayoutPreviewResponseLayout struct {
	// The fully rendered HTML body of the email layout.
	HTMLBody string `json:"html_body" api:"nullable"`
	// The fully rendered plain text body of the email layout.
	TextBody string `json:"text_body" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTMLBody    respjson.Field
		TextBody    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailLayoutPreviewResponseLayout) RawJSON() string { return r.JSON.raw }
func (r *EmailLayoutPreviewResponseLayout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the EmailLayout response under the `email_layout` key.
type EmailLayoutUpsertResponse struct {
	// A versioned email layout used within an environment.
	EmailLayout EmailLayout `json:"email_layout" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmailLayout respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailLayoutUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailLayoutUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the EmailLayout response under the `email_layout` key.
type EmailLayoutValidateResponse struct {
	// A versioned email layout used within an environment.
	EmailLayout EmailLayout `json:"email_layout" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmailLayout respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailLayoutValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailLayoutValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailLayoutGetParams struct {
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

// URLQuery serializes [EmailLayoutGetParams]'s query parameters as `url.Values`.
func (r EmailLayoutGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailLayoutListParams struct {
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

// URLQuery serializes [EmailLayoutListParams]'s query parameters as `url.Values`.
func (r EmailLayoutListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailLayoutPreviewParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A request to update or create an email layout.
	EmailLayout EmailLayoutRequestParam `json:"email_layout,omitzero" api:"required"`
	// A recipient reference, used when referencing a recipient by either their ID (for
	// a user), or by a reference for an object.
	Recipient shared.RecipientReferenceUnionParam `json:"recipient,omitzero" api:"required"`
	// The tenant to associate with the preview. Must not contain whitespace.
	Tenant param.Opt[string] `json:"tenant,omitzero"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// Optional workflow context for variable hydration. When provided,
	// recipient/actor/tenant are resolved via Knock.
	Workflow EmailLayoutPreviewParamsWorkflow `json:"workflow,omitzero"`
	// A recipient reference, used when referencing a recipient by either their ID (for
	// a user), or by a reference for an object.
	Actor shared.RecipientReferenceUnionParam `json:"actor,omitzero"`
	// The data to pass to the layout for rendering.
	Data map[string]any `json:"data,omitzero"`
	paramObj
}

func (r EmailLayoutPreviewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailLayoutPreviewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailLayoutPreviewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [EmailLayoutPreviewParams]'s query parameters as
// `url.Values`.
func (r EmailLayoutPreviewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional workflow context for variable hydration. When provided,
// recipient/actor/tenant are resolved via Knock.
//
// The property Key is required.
type EmailLayoutPreviewParamsWorkflow struct {
	// The workflow key.
	Key string `json:"key" api:"required"`
	// Workflow categories.
	Categories []string `json:"categories,omitzero"`
	paramObj
}

func (r EmailLayoutPreviewParamsWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow EmailLayoutPreviewParamsWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailLayoutPreviewParamsWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailLayoutUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A request to update or create an email layout.
	EmailLayout EmailLayoutRequestParam `json:"email_layout,omitzero" api:"required"`
	// When used with commit, creates a new version with identical content and commits
	// it if there are no unpublished changes.
	AllowEmpty param.Opt[bool] `query:"allow_empty,omitzero" json:"-"`
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

func (r EmailLayoutUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailLayoutUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailLayoutUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [EmailLayoutUpsertParams]'s query parameters as
// `url.Values`.
func (r EmailLayoutUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailLayoutValidateParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A request to update or create an email layout.
	EmailLayout EmailLayoutRequestParam `json:"email_layout,omitzero" api:"required"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	paramObj
}

func (r EmailLayoutValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailLayoutValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailLayoutValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [EmailLayoutValidateParams]'s query parameters as
// `url.Values`.
func (r EmailLayoutValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
