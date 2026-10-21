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
)

// Reusable requests are HTTP steps that workflows can link to and share across an
// account.
//
// ReusableRequestService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReusableRequestService] method instead.
type ReusableRequestService struct {
	Options []option.RequestOption
}

// NewReusableRequestService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReusableRequestService(opts ...option.RequestOption) (r ReusableRequestService) {
	r = ReusableRequestService{}
	r.Options = opts
	return
}

// Retrieves a reusable request by its key.
func (r *ReusableRequestService) Get(ctx context.Context, reusableRequestKey string, query ReusableRequestGetParams, opts ...option.RequestOption) (res *ReusableRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if reusableRequestKey == "" {
		err = errors.New("missing required reusable_request_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/reusable_requests/%s", reusableRequestKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of reusable requests for an environment.
func (r *ReusableRequestService) List(ctx context.Context, query ReusableRequestListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[ReusableRequest], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/reusable_requests"
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

// Returns a paginated list of reusable requests for an environment.
func (r *ReusableRequestService) ListAutoPaging(ctx context.Context, query ReusableRequestListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[ReusableRequest] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Archives a reusable request by key across all environments.
func (r *ReusableRequestService) Archive(ctx context.Context, reusableRequestKey string, body ReusableRequestArchiveParams, opts ...option.RequestOption) (res *ReusableRequestArchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if reusableRequestKey == "" {
		err = errors.New("missing required reusable_request_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/reusable_requests/%s", reusableRequestKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Unarchives a reusable request by key, with every copy archived alongside it in
// other environments.
//
// Returns a 404 when the key isn't archived, including when an earlier or
// concurrent request already restored it. Returns a 422 with the code `key_in_use`
// when a live reusable request uses the key.
func (r *ReusableRequestService) UpdateUnarchive(ctx context.Context, reusableRequestKey string, body ReusableRequestUpdateUnarchiveParams, opts ...option.RequestOption) (res *ReusableRequestUpdateUnarchiveResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if reusableRequestKey == "" {
		err = errors.New("missing required reusable_request_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/reusable_requests/%s/unarchive", reusableRequestKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Updates a reusable request of a given key, or creates a new one if it does not
// yet exist.
//
// The write targets the environment in the request. Accounts that allow writes in
// every environment can create and update in that environment. Pass force to write
// outside the request's origin environment. When writes are limited to
// development, this endpoint accepts the development environment unless force is
// set.
func (r *ReusableRequestService) Upsert(ctx context.Context, reusableRequestKey string, params ReusableRequestUpsertParams, opts ...option.RequestOption) (res *ReusableRequestUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if reusableRequestKey == "" {
		err = errors.New("missing required reusable_request_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/reusable_requests/%s", reusableRequestKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Validates a reusable request payload without persisting it.
//
// Validation uses the environment in the request. Accounts that allow writes in
// every environment can validate in that environment. When writes are limited to
// development, this endpoint accepts the development environment.
func (r *ReusableRequestService) Validate(ctx context.Context, reusableRequestKey string, params ReusableRequestValidateParams, opts ...option.RequestOption) (res *ReusableRequestValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if reusableRequestKey == "" {
		err = errors.New("missing required reusable_request_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/reusable_requests/%s/validate", reusableRequestKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A reusable HTTP request that workflows can link to.
type ReusableRequest struct {
	// The timestamp of when the reusable request was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The slug of the environment in which the reusable request exists.
	Environment string `json:"environment" api:"required"`
	// The unique key for the reusable request. Must be 3 to 255 characters and match
	// ^[a-z0-9_-]+$.
	Key string `json:"key" api:"required"`
	// The HTTP request a reusable request runs.
	Settings ReusableRequestSettings `json:"settings" api:"required"`
	// The reusable request type. Only http_fetch is supported.
	//
	// Any of "http_fetch".
	Type ReusableRequestType `json:"type" api:"required"`
	// The timestamp of when the reusable request was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnion `json:"conditions" api:"nullable"`
	// An internal note about the reusable request.
	Description string `json:"description" api:"nullable"`
	// A name for the reusable request.
	Name string `json:"name" api:"nullable"`
	// The SHA hash of the reusable request.
	Sha string `json:"sha" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Environment respjson.Field
		Key         respjson.Field
		Settings    respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		Conditions  respjson.Field
		Description respjson.Field
		Name        respjson.Field
		Sha         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReusableRequest) RawJSON() string { return r.JSON.raw }
func (r *ReusableRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reusable request type. Only http_fetch is supported.
type ReusableRequestType string

const (
	ReusableRequestTypeHTTPFetch ReusableRequestType = "http_fetch"
)

// A reusable request object with attributes to update or create a reusable
// request.
//
// The properties Settings, Type are required.
type ReusableRequestRequestParam struct {
	// The HTTP request a reusable request runs.
	Settings ReusableRequestSettingsParam `json:"settings,omitzero" api:"required"`
	// The reusable request type. Only http_fetch is supported.
	//
	// Any of "http_fetch".
	Type ReusableRequestRequestType `json:"type,omitzero" api:"required"`
	// An internal note about the reusable request.
	Description param.Opt[string] `json:"description,omitzero"`
	// A name for the reusable request.
	Name param.Opt[string] `json:"name,omitzero"`
	// A group of conditions to be evaluated.
	Conditions ConditionGroupUnionParam `json:"conditions,omitzero"`
	paramObj
}

func (r ReusableRequestRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ReusableRequestRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReusableRequestRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The reusable request type. Only http_fetch is supported.
type ReusableRequestRequestType string

const (
	ReusableRequestRequestTypeHTTPFetch ReusableRequestRequestType = "http_fetch"
)

// The HTTP request a reusable request runs.
type ReusableRequestSettings struct {
	// The HTTP method of the request.
	//
	// Any of "get", "post", "put", "delete", "patch".
	Method ReusableRequestSettingsMethod `json:"method" api:"required"`
	// The body of the request. Only used for POST or PUT requests.
	Body string `json:"body" api:"nullable"`
	// A JSON schema for inputs the request accepts.
	DataInputSchema map[string]any `json:"data_input_schema" api:"nullable"`
	// A JSON string of input values for the request.
	DataInputValuesStr string `json:"data_input_values_str" api:"nullable"`
	// The headers of the request. Can be a template string or a list of key-value
	// pairs.
	Headers ReusableRequestSettingsHeadersUnion `json:"headers"`
	// The query params of the request. Can be a template string or a list of key-value
	// pairs.
	QueryParams ReusableRequestSettingsQueryParamsUnion `json:"query_params"`
	// A JSON path into the response that is merged into the workflow data.
	ResponsePath string `json:"response_path" api:"nullable"`
	// The signing key used to sign the request.
	SigningKey string `json:"signing_key" api:"nullable"`
	// The URL of the request.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method             respjson.Field
		Body               respjson.Field
		DataInputSchema    respjson.Field
		DataInputValuesStr respjson.Field
		Headers            respjson.Field
		QueryParams        respjson.Field
		ResponsePath       respjson.Field
		SigningKey         respjson.Field
		URL                respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReusableRequestSettings) RawJSON() string { return r.JSON.raw }
func (r *ReusableRequestSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ReusableRequestSettings to a ReusableRequestSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ReusableRequestSettingsParam.Overrides()
func (r ReusableRequestSettings) ToParam() ReusableRequestSettingsParam {
	return param.Override[ReusableRequestSettingsParam](json.RawMessage(r.RawJSON()))
}

// The HTTP method of the request.
type ReusableRequestSettingsMethod string

const (
	ReusableRequestSettingsMethodGet    ReusableRequestSettingsMethod = "get"
	ReusableRequestSettingsMethodPost   ReusableRequestSettingsMethod = "post"
	ReusableRequestSettingsMethodPut    ReusableRequestSettingsMethod = "put"
	ReusableRequestSettingsMethodDelete ReusableRequestSettingsMethod = "delete"
	ReusableRequestSettingsMethodPatch  ReusableRequestSettingsMethod = "patch"
)

// ReusableRequestSettingsHeadersUnion contains all possible properties and values
// from [string],
// [[]ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfRequestTemplateHeadersArray]
type ReusableRequestSettingsHeadersUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItem] instead of an
	// object.
	OfRequestTemplateHeadersArray []ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItem `json:",inline"`
	JSON                          struct {
		OfString                      respjson.Field
		OfRequestTemplateHeadersArray respjson.Field
		raw                           string
	} `json:"-"`
}

func (u ReusableRequestSettingsHeadersUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ReusableRequestSettingsHeadersUnion) AsRequestTemplateHeadersArray() (v []ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ReusableRequestSettingsHeadersUnion) RawJSON() string { return u.JSON.raw }

func (r *ReusableRequestSettingsHeadersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItem struct {
	// The key of the header.
	Key string `json:"key" api:"required"`
	// The value of the header.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ReusableRequestSettingsQueryParamsUnion contains all possible properties and
// values from [string],
// [[]ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfRequestTemplateQuerysArray]
type ReusableRequestSettingsQueryParamsUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItem]
	// instead of an object.
	OfRequestTemplateQuerysArray []ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItem `json:",inline"`
	JSON                         struct {
		OfString                     respjson.Field
		OfRequestTemplateQuerysArray respjson.Field
		raw                          string
	} `json:"-"`
}

func (u ReusableRequestSettingsQueryParamsUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ReusableRequestSettingsQueryParamsUnion) AsRequestTemplateQueryParamsArray() (v []ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ReusableRequestSettingsQueryParamsUnion) RawJSON() string { return u.JSON.raw }

func (r *ReusableRequestSettingsQueryParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItem struct {
	// The key of the query param.
	Key string `json:"key" api:"required"`
	// The value of the query param.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItem) RawJSON() string {
	return r.JSON.raw
}
func (r *ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP request a reusable request runs.
//
// The property Method is required.
type ReusableRequestSettingsParam struct {
	// The HTTP method of the request.
	//
	// Any of "get", "post", "put", "delete", "patch".
	Method ReusableRequestSettingsMethod `json:"method,omitzero" api:"required"`
	// The body of the request. Only used for POST or PUT requests.
	Body param.Opt[string] `json:"body,omitzero"`
	// A JSON string of input values for the request.
	DataInputValuesStr param.Opt[string] `json:"data_input_values_str,omitzero"`
	// A JSON path into the response that is merged into the workflow data.
	ResponsePath param.Opt[string] `json:"response_path,omitzero"`
	// The signing key used to sign the request.
	SigningKey param.Opt[string] `json:"signing_key,omitzero"`
	// The URL of the request.
	URL param.Opt[string] `json:"url,omitzero"`
	// A JSON schema for inputs the request accepts.
	DataInputSchema map[string]any `json:"data_input_schema,omitzero"`
	// The headers of the request. Can be a template string or a list of key-value
	// pairs.
	Headers ReusableRequestSettingsHeadersUnionParam `json:"headers,omitzero"`
	// The query params of the request. Can be a template string or a list of key-value
	// pairs.
	QueryParams ReusableRequestSettingsQueryParamsUnionParam `json:"query_params,omitzero"`
	paramObj
}

func (r ReusableRequestSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow ReusableRequestSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReusableRequestSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ReusableRequestSettingsHeadersUnionParam struct {
	OfString                      param.Opt[string]                                                    `json:",omitzero,inline"`
	OfRequestTemplateHeadersArray []ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItemParam `json:",omitzero,inline"`
	paramUnion
}

func (u ReusableRequestSettingsHeadersUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfRequestTemplateHeadersArray)
}
func (u *ReusableRequestSettingsHeadersUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ReusableRequestSettingsHeadersUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfRequestTemplateHeadersArray) {
		return &u.OfRequestTemplateHeadersArray
	}
	return nil
}

// The properties Key, Value are required.
type ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItemParam struct {
	// The key of the header.
	Key string `json:"key" api:"required"`
	// The value of the header.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItemParam) MarshalJSON() (data []byte, err error) {
	type shadow ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ReusableRequestSettingsQueryParamsUnionParam struct {
	OfString                     param.Opt[string]                                                            `json:",omitzero,inline"`
	OfRequestTemplateQuerysArray []ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItemParam `json:",omitzero,inline"`
	paramUnion
}

func (u ReusableRequestSettingsQueryParamsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfRequestTemplateQuerysArray)
}
func (u *ReusableRequestSettingsQueryParamsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ReusableRequestSettingsQueryParamsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfRequestTemplateQuerysArray) {
		return &u.OfRequestTemplateQuerysArray
	}
	return nil
}

// The properties Key, Value are required.
type ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItemParam struct {
	// The key of the query param.
	Key string `json:"key" api:"required"`
	// The value of the query param.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItemParam) MarshalJSON() (data []byte, err error) {
	type shadow ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response from archiving a reusable request.
type ReusableRequestArchiveResponse struct {
	// The result of the archive operation.
	Result string `json:"result" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReusableRequestArchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *ReusableRequestArchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the ReusableRequest response under the `reusable_request` key.
type ReusableRequestUpdateUnarchiveResponse struct {
	// A reusable HTTP request that workflows can link to.
	ReusableRequest ReusableRequest `json:"reusable_request" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReusableRequest respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReusableRequestUpdateUnarchiveResponse) RawJSON() string { return r.JSON.raw }
func (r *ReusableRequestUpdateUnarchiveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the ReusableRequest response under the `reusable_request` key.
type ReusableRequestUpsertResponse struct {
	// A reusable HTTP request that workflows can link to.
	ReusableRequest ReusableRequest `json:"reusable_request" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReusableRequest respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReusableRequestUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *ReusableRequestUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the ReusableRequest response under the `reusable_request` key.
type ReusableRequestValidateResponse struct {
	// A reusable HTTP request that workflows can link to.
	ReusableRequest ReusableRequest `json:"reusable_request" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReusableRequest respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReusableRequestValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *ReusableRequestValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReusableRequestGetParams struct {
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// The slug of a branch to use. When `environment` is omitted, the branch is
	// resolved from Development after the account default is injected. When
	// `environment` is supplied, it must be `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	// Whether to hide uncommitted changes. When true, only committed changes will be
	// returned. When false, both committed and uncommitted changes will be returned.
	HideUncommittedChanges param.Opt[bool] `query:"hide_uncommitted_changes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReusableRequestGetParams]'s query parameters as
// `url.Values`.
func (r ReusableRequestGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReusableRequestListParams struct {
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The slug of a branch to use. When `environment` is omitted, the branch is
	// resolved from Development after the account default is injected. When
	// `environment` is supplied, it must be `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	// Whether to hide uncommitted changes. When true, only committed changes will be
	// returned. When false, both committed and uncommitted changes will be returned.
	HideUncommittedChanges param.Opt[bool] `query:"hide_uncommitted_changes,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReusableRequestListParams]'s query parameters as
// `url.Values`.
func (r ReusableRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReusableRequestArchiveParams struct {
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReusableRequestArchiveParams]'s query parameters as
// `url.Values`.
func (r ReusableRequestArchiveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReusableRequestUpdateUnarchiveParams struct {
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReusableRequestUpdateUnarchiveParams]'s query parameters as
// `url.Values`.
func (r ReusableRequestUpdateUnarchiveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReusableRequestUpsertParams struct {
	// A reusable request object with attributes to update or create a reusable
	// request.
	ReusableRequest ReusableRequestRequestParam `json:"reusable_request,omitzero" api:"required"`
	// When used with commit, creates a new version with identical content and commits
	// it if there are no unpublished changes.
	AllowEmpty param.Opt[bool] `query:"allow_empty,omitzero" json:"-"`
	// Whether to annotate the resource. Only used in the Knock CLI.
	Annotate param.Opt[bool] `query:"annotate,omitzero" json:"-"`
	// The slug of a branch to use. When `environment` is omitted, the branch is
	// resolved from Development after the account default is injected. When
	// `environment` is supplied, it must be `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// Whether to commit the resource at the same time as modifying it.
	Commit param.Opt[bool] `query:"commit,omitzero" json:"-"`
	// The message to commit the resource with, only used if `commit` is `true`.
	CommitMessage param.Opt[string] `query:"commit_message,omitzero" json:"-"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	// When set to true, forces the upsert to override existing content regardless of
	// environment restrictions. This bypasses the development-only environment check
	// and origin environment checks.
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
	paramObj
}

func (r ReusableRequestUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow ReusableRequestUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReusableRequestUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ReusableRequestUpsertParams]'s query parameters as
// `url.Values`.
func (r ReusableRequestUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReusableRequestValidateParams struct {
	// A reusable request object with attributes to update or create a reusable
	// request.
	ReusableRequest ReusableRequestRequestParam `json:"reusable_request,omitzero" api:"required"`
	// The slug of a branch to use. When `environment` is omitted, the branch is
	// resolved from Development after the account default is injected. When
	// `environment` is supplied, it must be `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

func (r ReusableRequestValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow ReusableRequestValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReusableRequestValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ReusableRequestValidateParams]'s query parameters as
// `url.Values`.
func (r ReusableRequestValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
