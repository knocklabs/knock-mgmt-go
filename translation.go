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

	"github.com/stainless-sdks/knock-mapi-go/internal/apijson"
	"github.com/stainless-sdks/knock-mapi-go/internal/apiquery"
	"github.com/stainless-sdks/knock-mapi-go/internal/requestconfig"
	"github.com/stainless-sdks/knock-mapi-go/option"
	"github.com/stainless-sdks/knock-mapi-go/packages/pagination"
	"github.com/stainless-sdks/knock-mapi-go/packages/param"
	"github.com/stainless-sdks/knock-mapi-go/packages/respjson"
)

// TranslationService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTranslationService] method instead.
type TranslationService struct {
	Options []option.RequestOption
}

// NewTranslationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTranslationService(opts ...option.RequestOption) (r TranslationService) {
	r = TranslationService{}
	r.Options = opts
	return
}

// Retrieve a translation by its locale and namespace, in a given environment.
func (r *TranslationService) Get(ctx context.Context, localeCode string, query TranslationGetParams, opts ...option.RequestOption) (res *TranslationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if localeCode == "" {
		err = errors.New("missing required locale_code parameter")
		return
	}
	path := fmt.Sprintf("v1/translations/%s", localeCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a paginated list of translations available in a given environment. The
// translations are returned in alphabetical order by locale code.
func (r *TranslationService) List(ctx context.Context, query TranslationListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Translation], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/translations"
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

// Returns a paginated list of translations available in a given environment. The
// translations are returned in alphabetical order by locale code.
func (r *TranslationService) ListAutoPaging(ctx context.Context, query TranslationListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Translation] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Updates a translation of a given locale code + namespace, or creates a new one
// if it does not yet exist.
//
// Note: this endpoint only operates on translations in the "development"
// environment.
func (r *TranslationService) Upsert(ctx context.Context, localeCode string, params TranslationUpsertParams, opts ...option.RequestOption) (res *TranslationUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if localeCode == "" {
		err = errors.New("missing required locale_code parameter")
		return
	}
	path := fmt.Sprintf("v1/translations/%s", localeCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Validates a translation payload without persisting it.
//
// Note: this endpoint only operates on translations in the "development"
// environment.
func (r *TranslationService) Validate(ctx context.Context, localeCode string, params TranslationValidateParams, opts ...option.RequestOption) (res *TranslationValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if localeCode == "" {
		err = errors.New("missing required locale_code parameter")
		return
	}
	path := fmt.Sprintf("v1/translations/%s/validate", localeCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// A translation object.
type Translation struct {
	// A JSON encoded string containing the key-value pairs of translation references
	// and translation strings.
	Content string `json:"content,required"`
	// Indicates whether content is a JSON encoded object string or a string in the PO
	// format.
	//
	// Any of "json", "po".
	Format TranslationFormat `json:"format,required"`
	// The timestamp of when the translation was created.
	InsertedAt time.Time `json:"inserted_at,required" format:"date-time"`
	// The locale code for the translation object.
	LocaleCode string `json:"locale_code,required"`
	// An optional namespace for the translation to help categorize your translations.
	Namespace string `json:"namespace,required"`
	// The timestamp of when the translation was last updated.
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Format      respjson.Field
		InsertedAt  respjson.Field
		LocaleCode  respjson.Field
		Namespace   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Translation) RawJSON() string { return r.JSON.raw }
func (r *Translation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether content is a JSON encoded object string or a string in the PO
// format.
type TranslationFormat string

const (
	TranslationFormatJson TranslationFormat = "json"
	TranslationFormatPo   TranslationFormat = "po"
)

// Wraps the Translation response under the `translation` key.
type TranslationGetResponse struct {
	// A translation object.
	Translation Translation `json:"translation,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Translation respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranslationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TranslationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Translation response under the `translation` key.
type TranslationUpsertResponse struct {
	// A translation object.
	Translation Translation `json:"translation,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Translation respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranslationUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *TranslationUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Translation response under the `translation` key.
type TranslationValidateResponse struct {
	// A translation object.
	Translation Translation `json:"translation,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Translation respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranslationValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *TranslationValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranslationGetParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// Whether to hide uncommitted changes. When true, only committed changes will be
	// returned. When false, both committed and uncommitted changes will be returned.
	HideUncommittedChanges param.Opt[bool] `query:"hide_uncommitted_changes,omitzero" json:"-"`
	// A specific namespace to filter translations for.
	Namespace param.Opt[string] `query:"namespace,omitzero" json:"-"`
	// Optionally specify the returned content format. Supports 'json' and 'po'.
	// Defaults to 'json'.
	//
	// Any of "json", "po".
	Format TranslationGetParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TranslationGetParams]'s query parameters as `url.Values`.
func (r TranslationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optionally specify the returned content format. Supports 'json' and 'po'.
// Defaults to 'json'.
type TranslationGetParamsFormat string

const (
	TranslationGetParamsFormatJson TranslationGetParamsFormat = "json"
	TranslationGetParamsFormatPo   TranslationGetParamsFormat = "po"
)

type TranslationListParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Whether to hide uncommitted changes. When true, only committed changes will be
	// returned. When false, both committed and uncommitted changes will be returned.
	HideUncommittedChanges param.Opt[bool] `query:"hide_uncommitted_changes,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// A specific locale code to filter translations for.
	LocaleCode param.Opt[string] `query:"locale_code,omitzero" json:"-"`
	// A specific namespace to filter translations for.
	Namespace param.Opt[string] `query:"namespace,omitzero" json:"-"`
	// Optionally specify the returned content format. Supports 'json' and 'po'.
	// Defaults to 'json'.
	//
	// Any of "json", "po".
	Format TranslationListParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TranslationListParams]'s query parameters as `url.Values`.
func (r TranslationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optionally specify the returned content format. Supports 'json' and 'po'.
// Defaults to 'json'.
type TranslationListParamsFormat string

const (
	TranslationListParamsFormatJson TranslationListParamsFormat = "json"
	TranslationListParamsFormatPo   TranslationListParamsFormat = "po"
)

type TranslationUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// An optional namespace that identifies the translation.
	Namespace string `query:"namespace,required" json:"-"`
	// A translation object with a content attribute used to update or create a
	// translation.
	Translation TranslationUpsertParamsTranslation `json:"translation,omitzero,required"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// Whether to commit the resource at the same time as modifying it.
	Commit param.Opt[bool] `query:"commit,omitzero" json:"-"`
	// The message to commit the resource with, only used if `commit` is `true`.
	CommitMessage param.Opt[string] `query:"commit_message,omitzero" json:"-"`
	// Optionally specify the returned content format. Supports 'json' and 'po'.
	// Defaults to 'json'.
	//
	// Any of "json", "po".
	Format TranslationUpsertParamsFormat `query:"format,omitzero" json:"-"`
	paramObj
}

func (r TranslationUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow TranslationUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranslationUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TranslationUpsertParams]'s query parameters as
// `url.Values`.
func (r TranslationUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A translation object with a content attribute used to update or create a
// translation.
//
// The properties Content, Format are required.
type TranslationUpsertParamsTranslation struct {
	// A JSON encoded string containing the key-value pairs of translation references
	// and translation strings.
	Content string `json:"content,required"`
	// Indicates whether content is a JSON encoded object string or a string in the PO
	// format.
	//
	// Any of "json", "po".
	Format string `json:"format,omitzero,required"`
	paramObj
}

func (r TranslationUpsertParamsTranslation) MarshalJSON() (data []byte, err error) {
	type shadow TranslationUpsertParamsTranslation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranslationUpsertParamsTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TranslationUpsertParamsTranslation](
		"format", "json", "po",
	)
}

// Optionally specify the returned content format. Supports 'json' and 'po'.
// Defaults to 'json'.
type TranslationUpsertParamsFormat string

const (
	TranslationUpsertParamsFormatJson TranslationUpsertParamsFormat = "json"
	TranslationUpsertParamsFormatPo   TranslationUpsertParamsFormat = "po"
)

type TranslationValidateParams struct {
	// The environment slug.
	Environment string `query:"environment,required" json:"-"`
	// A translation object with a content attribute used to update or create a
	// translation.
	Translation TranslationValidateParamsTranslation `json:"translation,omitzero,required"`
	paramObj
}

func (r TranslationValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow TranslationValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranslationValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TranslationValidateParams]'s query parameters as
// `url.Values`.
func (r TranslationValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A translation object with a content attribute used to update or create a
// translation.
//
// The properties Content, Format are required.
type TranslationValidateParamsTranslation struct {
	// A JSON encoded string containing the key-value pairs of translation references
	// and translation strings.
	Content string `json:"content,required"`
	// Indicates whether content is a JSON encoded object string or a string in the PO
	// format.
	//
	// Any of "json", "po".
	Format string `json:"format,omitzero,required"`
	paramObj
}

func (r TranslationValidateParamsTranslation) MarshalJSON() (data []byte, err error) {
	type shadow TranslationValidateParamsTranslation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranslationValidateParamsTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TranslationValidateParamsTranslation](
		"format", "json", "po",
	)
}
