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
		return
	}
	path := fmt.Sprintf("v1/email_layouts/%s", emailLayoutKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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

// Updates an email layout, or creates a new one if it does not yet exist.
//
// Note: this endpoint only operates in the "development" environment.
func (r *EmailLayoutService) Upsert(ctx context.Context, emailLayoutKey string, params EmailLayoutUpsertParams, opts ...option.RequestOption) (res *EmailLayoutUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if emailLayoutKey == "" {
		err = errors.New("missing required email_layout_key parameter")
		return
	}
	path := fmt.Sprintf("v1/email_layouts/%s", emailLayoutKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Validates an email layout payload without persisting it.
//
// Note: this endpoint only operates in the "development" environment.
func (r *EmailLayoutService) Validate(ctx context.Context, emailLayoutKey string, params EmailLayoutValidateParams, opts ...option.RequestOption) (res *EmailLayoutValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if emailLayoutKey == "" {
		err = errors.New("missing required email_layout_key parameter")
		return
	}
	path := fmt.Sprintf("v1/email_layouts/%s/validate", emailLayoutKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// A versioned email layout used within an environment.
type EmailLayout struct {
	// The timestamp of when the email layout was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The complete HTML content of the email layout.
	HTMLLayout string `json:"html_layout" api:"required"`
	// The unique key for this email layout.
	Key string `json:"key" api:"required"`
	// The human-readable name of this email layout.
	Name string `json:"name" api:"required"`
	// The SHA of the email layout.
	Sha string `json:"sha" api:"required"`
	// The complete plaintext content of the email layout.
	TextLayout string `json:"text_layout" api:"required"`
	// The environment of the email layout.
	Environment string `json:"environment"`
	// A list of one or more items to show in the footer of the email layout.
	FooterLinks []EmailLayoutFooterLink `json:"footer_links"`
	// The timestamp of when the email layout was last updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		HTMLLayout  respjson.Field
		Key         respjson.Field
		Name        respjson.Field
		Sha         respjson.Field
		TextLayout  respjson.Field
		Environment respjson.Field
		FooterLinks respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmailLayoutUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A request to update or create an email layout.
	EmailLayout EmailLayoutUpsertParamsEmailLayout `json:"email_layout,omitzero" api:"required"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A request to update or create an email layout.
//
// The properties HTMLLayout, Name, TextLayout are required.
type EmailLayoutUpsertParamsEmailLayout struct {
	// The complete HTML content of the email layout.
	HTMLLayout string `json:"html_layout" api:"required"`
	// The friendly name of this email layout.
	Name string `json:"name" api:"required"`
	// The complete plain text content of the email layout.
	TextLayout string `json:"text_layout" api:"required"`
	// A list of one or more items to show in the footer of the email layout.
	FooterLinks []EmailLayoutUpsertParamsEmailLayoutFooterLink `json:"footer_links,omitzero"`
	paramObj
}

func (r EmailLayoutUpsertParamsEmailLayout) MarshalJSON() (data []byte, err error) {
	type shadow EmailLayoutUpsertParamsEmailLayout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailLayoutUpsertParamsEmailLayout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Text, URL are required.
type EmailLayoutUpsertParamsEmailLayoutFooterLink struct {
	// The text to display as the link.
	Text string `json:"text" api:"required"`
	// The URL to link to.
	URL string `json:"url" api:"required"`
	paramObj
}

func (r EmailLayoutUpsertParamsEmailLayoutFooterLink) MarshalJSON() (data []byte, err error) {
	type shadow EmailLayoutUpsertParamsEmailLayoutFooterLink
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailLayoutUpsertParamsEmailLayoutFooterLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailLayoutValidateParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A request to update or create an email layout.
	EmailLayout EmailLayoutValidateParamsEmailLayout `json:"email_layout,omitzero" api:"required"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A request to update or create an email layout.
//
// The properties HTMLLayout, Name, TextLayout are required.
type EmailLayoutValidateParamsEmailLayout struct {
	// The complete HTML content of the email layout.
	HTMLLayout string `json:"html_layout" api:"required"`
	// The friendly name of this email layout.
	Name string `json:"name" api:"required"`
	// The complete plain text content of the email layout.
	TextLayout string `json:"text_layout" api:"required"`
	// A list of one or more items to show in the footer of the email layout.
	FooterLinks []EmailLayoutValidateParamsEmailLayoutFooterLink `json:"footer_links,omitzero"`
	paramObj
}

func (r EmailLayoutValidateParamsEmailLayout) MarshalJSON() (data []byte, err error) {
	type shadow EmailLayoutValidateParamsEmailLayout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailLayoutValidateParamsEmailLayout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Text, URL are required.
type EmailLayoutValidateParamsEmailLayoutFooterLink struct {
	// The text to display as the link.
	Text string `json:"text" api:"required"`
	// The URL to link to.
	URL string `json:"url" api:"required"`
	paramObj
}

func (r EmailLayoutValidateParamsEmailLayoutFooterLink) MarshalJSON() (data []byte, err error) {
	type shadow EmailLayoutValidateParamsEmailLayoutFooterLink
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailLayoutValidateParamsEmailLayoutFooterLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
