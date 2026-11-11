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

// Warehouses are the databases in your data warehouse that Knock reads from.
//
// WarehouseService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWarehouseService] method instead.
type WarehouseService struct {
	Options []option.RequestOption
}

// NewWarehouseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWarehouseService(opts ...option.RequestOption) (r WarehouseService) {
	r = WarehouseService{}
	r.Options = opts
	return
}

// Retrieve a warehouse by its key.
func (r *WarehouseService) Get(ctx context.Context, warehouseKey string, query WarehouseGetParams, opts ...option.RequestOption) (res *Warehouse, err error) {
	opts = slices.Concat(r.Options, opts)
	if warehouseKey == "" {
		err = errors.New("missing required warehouse_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/warehouses/%s", warehouseKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of the project's warehouses. Warehouses are shared by
// every environment.
func (r *WarehouseService) List(ctx context.Context, query WarehouseListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Warehouse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/warehouses"
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

// Returns a paginated list of the project's warehouses. Warehouses are shared by
// every environment.
func (r *WarehouseService) ListAutoPaging(ctx context.Context, query WarehouseListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Warehouse] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// A warehouse is one database in your data warehouse that Knock reads from.
// Warehouses belong to the project and are shared by every environment.
type Warehouse struct {
	// The status of the warehouse's connection.
	//
	// Any of "not_configured", "key_not_installed", "failed", "connecting",
	// "untested", "connected".
	ConnectionStatus WarehouseConnectionStatus `json:"connection_status" api:"required"`
	// When the warehouse was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The unique key of the warehouse.
	Key string `json:"key" api:"required"`
	// The name of the warehouse.
	Name string `json:"name" api:"required"`
	// The type of data warehouse.
	//
	// Any of "redshift", "snowflake", "postgres", "bigquery".
	Type WarehouseType `json:"type" api:"required"`
	// When the warehouse was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The error from the last connection test, with a `code` and a `message`. Null
	// when the last test passed.
	ConnectionError map[string]any `json:"connection_error" api:"nullable"`
	// How Knock connects to the warehouse. The fields depend on the warehouse type,
	// such as `host`, `port`, `database`, and `username`, or `project_id`, `region`,
	// and `service_account_email` for BigQuery. Credentials such as passwords are
	// never returned.
	ConnectionSettings map[string]any `json:"connection_settings" api:"nullable"`
	// When the connection was last tested.
	LastTestedAt time.Time `json:"last_tested_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectionStatus   respjson.Field
		CreatedAt          respjson.Field
		Key                respjson.Field
		Name               respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		ConnectionError    respjson.Field
		ConnectionSettings respjson.Field
		LastTestedAt       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Warehouse) RawJSON() string { return r.JSON.raw }
func (r *Warehouse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the warehouse's connection.
type WarehouseConnectionStatus string

const (
	WarehouseConnectionStatusNotConfigured   WarehouseConnectionStatus = "not_configured"
	WarehouseConnectionStatusKeyNotInstalled WarehouseConnectionStatus = "key_not_installed"
	WarehouseConnectionStatusFailed          WarehouseConnectionStatus = "failed"
	WarehouseConnectionStatusConnecting      WarehouseConnectionStatus = "connecting"
	WarehouseConnectionStatusUntested        WarehouseConnectionStatus = "untested"
	WarehouseConnectionStatusConnected       WarehouseConnectionStatus = "connected"
)

// The type of data warehouse.
type WarehouseType string

const (
	WarehouseTypeRedshift  WarehouseType = "redshift"
	WarehouseTypeSnowflake WarehouseType = "snowflake"
	WarehouseTypePostgres  WarehouseType = "postgres"
	WarehouseTypeBigquery  WarehouseType = "bigquery"
)

type WarehouseGetParams struct {
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WarehouseGetParams]'s query parameters as `url.Values`.
func (r WarehouseGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WarehouseListParams struct {
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WarehouseListParams]'s query parameters as `url.Values`.
func (r WarehouseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
