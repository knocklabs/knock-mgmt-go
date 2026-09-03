// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/knocklabs/knock-mgmt-go/internal/apijson"
	"github.com/knocklabs/knock-mgmt-go/internal/apiquery"
	shimjson "github.com/knocklabs/knock-mgmt-go/internal/encoding/json"
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
	"github.com/knocklabs/knock-mgmt-go/shared"
)

// SchemaService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSchemaService] method instead.
type SchemaService struct {
	Options []option.RequestOption
}

// NewSchemaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSchemaService(opts ...option.RequestOption) (r SchemaService) {
	r = SchemaService{}
	r.Options = opts
	return
}

// Retrieve the configuration for an item schema (`user`, `tenant`, or `object`) in
// a given environment, including all of its configured properties.
func (r *SchemaService) Get(ctx context.Context, itemType string, query SchemaGetParams, opts ...option.RequestOption) (res *SchemaGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if itemType == "" {
		err = errors.New("missing required item_type parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/schemas/%s", itemType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the configuration for all managed item schemas (`user`, `tenant`, and
// `object`) in a given environment. Branch-qualified reads return the schemas
// inherited from the parent environment.
func (r *SchemaService) List(ctx context.Context, query SchemaListParams, opts ...option.RequestOption) (res *SchemaListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/schemas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Applies changes for the item schema properties in the request. Omitted
// properties are left unchanged; hide a property with `visible: false` rather than
// removing it. The required permissions depend on what changes: changing a
// property's display settings (`visible`/`description`) requires
// `item_schemas:manage`; changing a property's type or example, or adding a
// property, requires `item_schemas:edit`. Adding a property that is already hidden
// or already has a description requires both.
func (r *SchemaService) Upsert(ctx context.Context, itemType string, params SchemaUpsertParams, opts ...option.RequestOption) (res *SchemaUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if itemType == "" {
		err = errors.New("missing required item_type parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/schemas/%s", itemType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Checks an item schema configuration payload and reports which permissions it
// would require, without saving any changes.
func (r *SchemaService) Validate(ctx context.Context, itemType string, params SchemaValidateParams, opts ...option.RequestOption) (res *SchemaValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if itemType == "" {
		err = errors.New("missing required item_type parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/schemas/%s/validate", itemType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A managed schema configuration for users, tenants, or objects.
type ItemSchema struct {
	// The item type the schema applies to.
	//
	// Any of "user", "tenant", "object".
	ItemType ItemSchemaItemType `json:"item_type" api:"required"`
	// The managed properties for the schema.
	Properties []ItemSchemaProperty `json:"properties" api:"required"`
	// The object collection key. Only present for object schemas.
	ItemID string `json:"item_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemType    respjson.Field
		Properties  respjson.Field
		ItemID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemSchema) RawJSON() string { return r.JSON.raw }
func (r *ItemSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The item type the schema applies to.
type ItemSchemaItemType string

const (
	ItemSchemaItemTypeUser   ItemSchemaItemType = "user"
	ItemSchemaItemTypeTenant ItemSchemaItemType = "tenant"
	ItemSchemaItemTypeObject ItemSchemaItemType = "object"
)

// A property definition within an item schema.
type ItemSchemaProperty struct {
	// The property key.
	Key string `json:"key" api:"required"`
	// The description of the property.
	Description string `json:"description" api:"nullable"`
	// The referenced item type when the property stores an item reference.
	ItemType string `json:"item_type" api:"nullable"`
	// The property preview text.
	PreviewText string `json:"preview_text" api:"nullable"`
	// The primitive or referenced item type for the property.
	Type string `json:"type" api:"nullable"`
	// Whether the property is visible in the schema management UI.
	Visible bool `json:"visible"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Description respjson.Field
		ItemType    respjson.Field
		PreviewText respjson.Field
		Type        respjson.Field
		Visible     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemSchemaProperty) RawJSON() string { return r.JSON.raw }
func (r *ItemSchemaProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SchemaGetResponse = any

// A paginated list of ItemSchema. Contains a list of entries and page information.
type SchemaListResponse struct {
	// A list of entries.
	Entries []ItemSchema `json:"entries" api:"required"`
	// The information about a paginated result.
	PageInfo shared.PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchemaListResponse) RawJSON() string { return r.JSON.raw }
func (r *SchemaListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SchemaUpsertResponse = any

type SchemaValidateResponse = any

type SchemaGetParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The object collection, required when `item_type` is `object`.
	Collection param.Opt[string] `query:"collection,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SchemaGetParams]'s query parameters as `url.Values`.
func (r SchemaGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SchemaListParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// Filter schemas by item type (`user`, `tenant`, or `object`).
	ItemType param.Opt[string] `query:"item_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SchemaListParams]'s query parameters as `url.Values`.
func (r SchemaListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SchemaUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The object collection, required when `item_type` is `object`.
	Collection param.Opt[string] `query:"collection,omitzero" json:"-"`
	Body       any
	paramObj
}

func (r SchemaUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SchemaUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SchemaUpsertParams]'s query parameters as `url.Values`.
func (r SchemaUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SchemaValidateParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The slug of a branch to use. This option can only be used when `environment` is
	// `"development"`.
	Branch param.Opt[string] `query:"branch,omitzero" json:"-"`
	// The object collection, required when `item_type` is `object`.
	Collection param.Opt[string] `query:"collection,omitzero" json:"-"`
	Body       any
	paramObj
}

func (r SchemaValidateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *SchemaValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SchemaValidateParams]'s query parameters as `url.Values`.
func (r SchemaValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
