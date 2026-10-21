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

// WebhookService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

// Creates a new outbound webhook in the given environment. The response includes
// the `secret_key` which is only revealed at creation time.
func (r *WebhookService) New(ctx context.Context, params WebhookNewParams, opts ...option.RequestOption) (res *WebhookNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a single outbound webhook by its ID.
func (r *WebhookService) Get(ctx context.Context, id string, query WebhookGetParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Updates an existing outbound webhook.
func (r *WebhookService) Update(ctx context.Context, id string, params WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of outbound webhooks for a given environment.
func (r *WebhookService) List(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Webhook], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/webhooks"
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

// Returns a paginated list of outbound webhooks for a given environment.
func (r *WebhookService) ListAutoPaging(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Webhook] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Archives an outbound webhook.
func (r *WebhookService) Delete(ctx context.Context, id string, body WebhookDeleteParams, opts ...option.RequestOption) (res *WebhookDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// An outbound webhook that receives event notifications.
type Webhook struct {
	// The unique identifier for the webhook.
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether the webhook is currently active.
	Active bool `json:"active" api:"required"`
	// The timestamp of when the webhook was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The list of event types that this webhook is subscribed to, in dotted format
	// (e.g. `message.sent`).
	EventTypes []string `json:"event_types" api:"required"`
	// The timestamp of when the webhook was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The URL that the webhook will send events to.
	URL string `json:"url" api:"required"`
	// An optional description of the webhook.
	Description string `json:"description" api:"nullable"`
	// The secret key used to sign webhook payloads. Only returned when a webhook is
	// first created.
	SecretKey string `json:"secret_key" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Active      respjson.Field
		CreatedAt   respjson.Field
		EventTypes  respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		Description respjson.Field
		SecretKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Webhook) RawJSON() string { return r.JSON.raw }
func (r *Webhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The request body for creating or updating an outbound webhook.
//
// The properties EventTypes, URL are required.
type WebhookRequestParam struct {
	// The list of event types that this webhook should be subscribed to, in dotted
	// format (e.g. `message.sent`).
	EventTypes []string `json:"event_types,omitzero" api:"required"`
	// The URL that the webhook will send events to.
	URL string `json:"url" api:"required"`
	// Whether the webhook should be active.
	Active param.Opt[bool] `json:"active,omitzero"`
	// An optional description of the webhook.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r WebhookRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Webhook response under the `webhook` key.
type WebhookNewResponse struct {
	// An outbound webhook that receives event notifications.
	Webhook Webhook `json:"webhook" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Webhook     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Webhook response under the `webhook` key.
type WebhookUpdateResponse struct {
	// An outbound webhook that receives event notifications.
	Webhook Webhook `json:"webhook" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Webhook     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response from archiving a webhook.
type WebhookDeleteResponse struct {
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
func (r WebhookDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {
	// The request body for creating or updating an outbound webhook.
	Webhook WebhookRequestParam `json:"webhook,omitzero" api:"required"`
	// The slug of a branch to use. When `environment` is omitted, the branch is
	// resolved from Development after the account default is injected. When
	// `environment` is supplied, it must be `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WebhookNewParams]'s query parameters as `url.Values`.
func (r WebhookNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookGetParams struct {
	// The slug of a branch to use. When `environment` is omitted, the branch is
	// resolved from Development after the account default is injected. When
	// `environment` is supplied, it must be `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookGetParams]'s query parameters as `url.Values`.
func (r WebhookGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookUpdateParams struct {
	// The request body for creating or updating an outbound webhook.
	Webhook WebhookRequestParam `json:"webhook,omitzero" api:"required"`
	// The slug of a branch to use. When `environment` is omitted, the branch is
	// resolved from Development after the account default is injected. When
	// `environment` is supplied, it must be `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WebhookUpdateParams]'s query parameters as `url.Values`.
func (r WebhookUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookListParams struct {
	// Filter webhooks by active status.
	Active param.Opt[bool] `query:"active,omitzero" json:"-"`
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The slug of a branch to use. When `environment` is omitted, the branch is
	// resolved from Development after the account default is injected. When
	// `environment` is supplied, it must be `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	// Filter webhooks by event type in dotted format (e.g. `message.sent`).
	EventType param.Opt[string] `query:"event_type,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookListParams]'s query parameters as `url.Values`.
func (r WebhookListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookDeleteParams struct {
	// The slug of a branch to use. When `environment` is omitted, the branch is
	// resolved from Development after the account default is injected. When
	// `environment` is supplied, it must be `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookDeleteParams]'s query parameters as `url.Values`.
func (r WebhookDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
