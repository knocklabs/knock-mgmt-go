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

// Warehouse syncs read rows from a warehouse on a schedule and write them to
// Knock.
//
// WarehouseSyncService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWarehouseSyncService] method instead.
type WarehouseSyncService struct {
	Options []option.RequestOption
}

// NewWarehouseSyncService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWarehouseSyncService(opts ...option.RequestOption) (r WarehouseSyncService) {
	r = WarehouseSyncService{}
	r.Options = opts
	return
}

// Retrieve a warehouse sync by its key in a given environment.
func (r *WarehouseSyncService) Get(ctx context.Context, warehouseSyncKey string, query WarehouseSyncGetParams, opts ...option.RequestOption) (res *WarehouseSync, err error) {
	opts = slices.Concat(r.Options, opts)
	if warehouseSyncKey == "" {
		err = errors.New("missing required warehouse_sync_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/warehouse_syncs/%s", warehouseSyncKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Creates a warehouse sync, or saves a new draft of it if it exists. Fields left
// out keep their current values. Pass `commit=true` to publish the draft. A draft
// that fails to publish stays saved.
func (r *WarehouseSyncService) Update(ctx context.Context, warehouseSyncKey string, params WarehouseSyncUpdateParams, opts ...option.RequestOption) (res *WarehouseSyncUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if warehouseSyncKey == "" {
		err = errors.New("missing required warehouse_sync_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/warehouse_syncs/%s", warehouseSyncKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of warehouse syncs for the given environment.
func (r *WarehouseSyncService) List(ctx context.Context, query WarehouseSyncListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[WarehouseSync], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/warehouse_syncs"
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

// Returns a paginated list of warehouse syncs for the given environment.
func (r *WarehouseSyncService) ListAutoPaging(ctx context.Context, query WarehouseSyncListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[WarehouseSync] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Returns the most recent runs of a warehouse sync, newest first.
func (r *WarehouseSyncService) GetRuns(ctx context.Context, warehouseSyncKey string, query WarehouseSyncGetRunsParams, opts ...option.RequestOption) (res *WarehouseSyncGetRunsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if warehouseSyncKey == "" {
		err = errors.New("missing required warehouse_sync_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/warehouse_syncs/%s/runs", warehouseSyncKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Starts a run of the published warehouse sync now, instead of waiting for its
// schedule. Refused while a run is in progress.
func (r *WarehouseSyncService) UpdateRun(ctx context.Context, warehouseSyncKey string, body WarehouseSyncUpdateRunParams, opts ...option.RequestOption) (res *WarehouseSyncUpdateRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if warehouseSyncKey == "" {
		err = errors.New("missing required warehouse_sync_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/warehouse_syncs/%s/run", warehouseSyncKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Pauses or resumes a warehouse sync, or changes how often it runs. Settings take
// effect right away and don't need a publish.
func (r *WarehouseSyncService) UpdateSettings(ctx context.Context, warehouseSyncKey string, params WarehouseSyncUpdateSettingsParams, opts ...option.RequestOption) (res *WarehouseSyncUpdateSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if warehouseSyncKey == "" {
		err = errors.New("missing required warehouse_sync_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/warehouse_syncs/%s/settings", warehouseSyncKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Checks a warehouse sync payload the way an upsert would, without saving it.
// Returns the payload as sent.
func (r *WarehouseSyncService) UpdateValidate(ctx context.Context, warehouseSyncKey string, params WarehouseSyncUpdateValidateParams, opts ...option.RequestOption) (res *WarehouseSyncUpdateValidateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if warehouseSyncKey == "" {
		err = errors.New("missing required warehouse_sync_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/warehouse_syncs/%s/validate", warehouseSyncKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A warehouse sync reads rows from a warehouse on a schedule and writes them to
// Knock as users, tenants, objects, audience members, or subscriptions.
type WarehouseSync struct {
	// When the sync was created. (read-only).
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether the sync runs on its schedule.
	Enabled bool `json:"enabled" api:"required"`
	// The slug of the environment the sync belongs to. (read-only).
	Environment string `json:"environment" api:"required"`
	// How often the sync runs, in minutes.
	//
	// Any of 10, 15, 60, 360, 720, 1440.
	FrequencyMinutes int64 `json:"frequency_minutes" api:"required"`
	// Whether the sync has a draft that hasn't been published.
	HasUnpublishedChanges bool `json:"has_unpublished_changes" api:"required"`
	// The unique key of the sync in its environment.
	Key string `json:"key" api:"required"`
	// How columns map to Knock fields.
	Mappings []WarehouseSyncMapping `json:"mappings" api:"required"`
	// The name of the sync. Must be at minimum 1 character and at maximum 255
	// characters in length.
	Name string `json:"name" api:"required"`
	// The SHA hash of the sync's definition. (read-only).
	Sha string `json:"sha" api:"required"`
	// The status of the sync.
	//
	// Any of "draft", "publishing", "delivery_failed", "paused", "needs_attention",
	// "live".
	Status WarehouseSyncStatus `json:"status" api:"required"`
	// What each row becomes in Knock. Can't change after creation.
	//
	// Any of "user", "tenant", "object", "static_audience_membership", "subscription".
	TargetType WarehouseSyncTargetType `json:"target_type" api:"required"`
	// When the returned version was saved. (read-only).
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The ID of the version returned. Send it back when saving to make sure nobody
	// saved since.
	VersionID string `json:"version_id" api:"required" format:"uuid"`
	// The key of the warehouse the sync reads from. Can't change after creation.
	WarehouseKey string `json:"warehouse_key" api:"required"`
	// How the sync detects changed and deleted rows.
	ChangeTracking WarehouseSyncChangeTracking `json:"change_tracking"`
	// An optional description of the sync. Maximum of 280 characters allowed.
	Description string `json:"description" api:"nullable"`
	// The last run failure, cleared when a run succeeds.
	LastFailure WarehouseSyncLastFailure `json:"last_failure" api:"nullable"`
	// When the sync was last published.
	LastPublishedAt time.Time `json:"last_published_at" api:"nullable" format:"date-time"`
	// The column that uniquely identifies each row.
	PrimaryKeyColumn string `json:"primary_key_column" api:"nullable"`
	// Why the last publish didn't go live, when it failed.
	PublishError string `json:"publish_error" api:"nullable"`
	// The rows a warehouse sync reads: a table or view (`relation`), or the result of
	// a SQL query (`sql`).
	Source WarehouseSyncSource `json:"source"`
	// The collection or audience the sync writes to. `object` and `subscription` syncs
	// set `collection`, and `static_audience_membership` syncs set `audience_key`.
	// Null for `user` and `tenant` syncs.
	Target WarehouseSyncTarget `json:"target" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt             respjson.Field
		Enabled               respjson.Field
		Environment           respjson.Field
		FrequencyMinutes      respjson.Field
		HasUnpublishedChanges respjson.Field
		Key                   respjson.Field
		Mappings              respjson.Field
		Name                  respjson.Field
		Sha                   respjson.Field
		Status                respjson.Field
		TargetType            respjson.Field
		UpdatedAt             respjson.Field
		VersionID             respjson.Field
		WarehouseKey          respjson.Field
		ChangeTracking        respjson.Field
		Description           respjson.Field
		LastFailure           respjson.Field
		LastPublishedAt       respjson.Field
		PrimaryKeyColumn      respjson.Field
		PublishError          respjson.Field
		Source                respjson.Field
		Target                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WarehouseSync) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSync) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the sync.
type WarehouseSyncStatus string

const (
	WarehouseSyncStatusDraft          WarehouseSyncStatus = "draft"
	WarehouseSyncStatusPublishing     WarehouseSyncStatus = "publishing"
	WarehouseSyncStatusDeliveryFailed WarehouseSyncStatus = "delivery_failed"
	WarehouseSyncStatusPaused         WarehouseSyncStatus = "paused"
	WarehouseSyncStatusNeedsAttention WarehouseSyncStatus = "needs_attention"
	WarehouseSyncStatusLive           WarehouseSyncStatus = "live"
)

// What each row becomes in Knock. Can't change after creation.
type WarehouseSyncTargetType string

const (
	WarehouseSyncTargetTypeUser                     WarehouseSyncTargetType = "user"
	WarehouseSyncTargetTypeTenant                   WarehouseSyncTargetType = "tenant"
	WarehouseSyncTargetTypeObject                   WarehouseSyncTargetType = "object"
	WarehouseSyncTargetTypeStaticAudienceMembership WarehouseSyncTargetType = "static_audience_membership"
	WarehouseSyncTargetTypeSubscription             WarehouseSyncTargetType = "subscription"
)

// The last run failure, cleared when a run succeeds.
type WarehouseSyncLastFailure struct {
	// What went wrong.
	Message string `json:"message"`
	// Where the run failed: reading the warehouse (`extract`), sending rows to Knock
	// (`load`), or applying them (`apply`).
	//
	// Any of "extract", "load", "apply".
	Stage string `json:"stage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Stage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WarehouseSyncLastFailure) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSyncLastFailure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the sync detects changed and deleted rows.
type WarehouseSyncChangeTracking struct {
	// A boolean column that marks a row as deleted. Requires `last_modified_column`.
	IsDeletedColumn string `json:"is_deleted_column" api:"nullable"`
	// A timestamp column that changes whenever a row changes. With it, each run reads
	// only rows changed since the last run.
	LastModifiedColumn string `json:"last_modified_column" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsDeletedColumn    respjson.Field
		LastModifiedColumn respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WarehouseSyncChangeTracking) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSyncChangeTracking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WarehouseSyncChangeTracking to a
// WarehouseSyncChangeTrackingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WarehouseSyncChangeTrackingParam.Overrides()
func (r WarehouseSyncChangeTracking) ToParam() WarehouseSyncChangeTrackingParam {
	return param.Override[WarehouseSyncChangeTrackingParam](json.RawMessage(r.RawJSON()))
}

// How the sync detects changed and deleted rows.
type WarehouseSyncChangeTrackingParam struct {
	// A boolean column that marks a row as deleted. Requires `last_modified_column`.
	IsDeletedColumn param.Opt[string] `json:"is_deleted_column,omitzero"`
	// A timestamp column that changes whenever a row changes. With it, each run reads
	// only rows changed since the last run.
	LastModifiedColumn param.Opt[string] `json:"last_modified_column,omitzero"`
	paramObj
}

func (r WarehouseSyncChangeTrackingParam) MarshalJSON() (data []byte, err error) {
	type shadow WarehouseSyncChangeTrackingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WarehouseSyncChangeTrackingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maps one column from the warehouse to one field in Knock.
type WarehouseSyncMapping struct {
	// The column to read. Can be a dotted path into an object column, such as
	// `PROPERTIES.plan`.
	SourceColumn string `json:"source_column" api:"required"`
	// The Knock field to write the value to.
	TargetField string `json:"target_field" api:"required"`
	// What to do when the column is `NULL`. `set_null` clears the Knock field, and
	// `preserve` keeps its current value.
	//
	// Any of "set_null", "preserve".
	OnNull WarehouseSyncMappingOnNull `json:"on_null"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SourceColumn respjson.Field
		TargetField  respjson.Field
		OnNull       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WarehouseSyncMapping) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSyncMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WarehouseSyncMapping to a WarehouseSyncMappingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WarehouseSyncMappingParam.Overrides()
func (r WarehouseSyncMapping) ToParam() WarehouseSyncMappingParam {
	return param.Override[WarehouseSyncMappingParam](json.RawMessage(r.RawJSON()))
}

// What to do when the column is `NULL`. `set_null` clears the Knock field, and
// `preserve` keeps its current value.
type WarehouseSyncMappingOnNull string

const (
	WarehouseSyncMappingOnNullSetNull  WarehouseSyncMappingOnNull = "set_null"
	WarehouseSyncMappingOnNullPreserve WarehouseSyncMappingOnNull = "preserve"
)

// Maps one column from the warehouse to one field in Knock.
//
// The properties SourceColumn, TargetField are required.
type WarehouseSyncMappingParam struct {
	// The column to read. Can be a dotted path into an object column, such as
	// `PROPERTIES.plan`.
	SourceColumn string `json:"source_column" api:"required"`
	// The Knock field to write the value to.
	TargetField string `json:"target_field" api:"required"`
	// What to do when the column is `NULL`. `set_null` clears the Knock field, and
	// `preserve` keeps its current value.
	//
	// Any of "set_null", "preserve".
	OnNull WarehouseSyncMappingOnNull `json:"on_null,omitzero"`
	paramObj
}

func (r WarehouseSyncMappingParam) MarshalJSON() (data []byte, err error) {
	type shadow WarehouseSyncMappingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WarehouseSyncMappingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A warehouse sync payload for upsert or validate. Fields left out keep their
// current values. `warehouse_key`, `target_type`, and `target` are required to
// create a sync and can't change after. Change `enabled` and `frequency_minutes`
// with the settings endpoint.
type WarehouseSyncRequest struct {
	// How the sync detects changed and deleted rows.
	ChangeTracking WarehouseSyncChangeTracking `json:"change_tracking"`
	// An optional description of the sync. Maximum of 280 characters allowed.
	Description string `json:"description" api:"nullable"`
	// How columns map to Knock fields.
	Mappings []WarehouseSyncMapping `json:"mappings"`
	// The name of the sync. Must be at minimum 1 character and at maximum 255
	// characters in length.
	Name string `json:"name"`
	// The column that uniquely identifies each row.
	PrimaryKeyColumn string `json:"primary_key_column" api:"nullable"`
	// The rows a warehouse sync reads: a table or view (`relation`), or the result of
	// a SQL query (`sql`).
	Source WarehouseSyncSource `json:"source"`
	// The collection or audience the sync writes to. `object` and `subscription` syncs
	// set `collection`, and `static_audience_membership` syncs set `audience_key`.
	// Null for `user` and `tenant` syncs.
	Target WarehouseSyncTarget `json:"target" api:"nullable"`
	// What each row becomes in Knock. Can't change after creation.
	//
	// Any of "user", "tenant", "object", "static_audience_membership", "subscription".
	TargetType WarehouseSyncRequestTargetType `json:"target_type"`
	// The `version_id` of the sync you last read. When it isn't the current version,
	// someone saved since, and the request is refused.
	VersionID string `json:"version_id" format:"uuid"`
	// The key of the warehouse the sync reads from. Can't change after creation.
	WarehouseKey string `json:"warehouse_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChangeTracking   respjson.Field
		Description      respjson.Field
		Mappings         respjson.Field
		Name             respjson.Field
		PrimaryKeyColumn respjson.Field
		Source           respjson.Field
		Target           respjson.Field
		TargetType       respjson.Field
		VersionID        respjson.Field
		WarehouseKey     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WarehouseSyncRequest) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSyncRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WarehouseSyncRequest to a WarehouseSyncRequestParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WarehouseSyncRequestParam.Overrides()
func (r WarehouseSyncRequest) ToParam() WarehouseSyncRequestParam {
	return param.Override[WarehouseSyncRequestParam](json.RawMessage(r.RawJSON()))
}

// What each row becomes in Knock. Can't change after creation.
type WarehouseSyncRequestTargetType string

const (
	WarehouseSyncRequestTargetTypeUser                     WarehouseSyncRequestTargetType = "user"
	WarehouseSyncRequestTargetTypeTenant                   WarehouseSyncRequestTargetType = "tenant"
	WarehouseSyncRequestTargetTypeObject                   WarehouseSyncRequestTargetType = "object"
	WarehouseSyncRequestTargetTypeStaticAudienceMembership WarehouseSyncRequestTargetType = "static_audience_membership"
	WarehouseSyncRequestTargetTypeSubscription             WarehouseSyncRequestTargetType = "subscription"
)

// A warehouse sync payload for upsert or validate. Fields left out keep their
// current values. `warehouse_key`, `target_type`, and `target` are required to
// create a sync and can't change after. Change `enabled` and `frequency_minutes`
// with the settings endpoint.
type WarehouseSyncRequestParam struct {
	// An optional description of the sync. Maximum of 280 characters allowed.
	Description param.Opt[string] `json:"description,omitzero"`
	// The column that uniquely identifies each row.
	PrimaryKeyColumn param.Opt[string] `json:"primary_key_column,omitzero"`
	// The name of the sync. Must be at minimum 1 character and at maximum 255
	// characters in length.
	Name param.Opt[string] `json:"name,omitzero"`
	// The `version_id` of the sync you last read. When it isn't the current version,
	// someone saved since, and the request is refused.
	VersionID param.Opt[string] `json:"version_id,omitzero" format:"uuid"`
	// The key of the warehouse the sync reads from. Can't change after creation.
	WarehouseKey param.Opt[string] `json:"warehouse_key,omitzero"`
	// The collection or audience the sync writes to. `object` and `subscription` syncs
	// set `collection`, and `static_audience_membership` syncs set `audience_key`.
	// Null for `user` and `tenant` syncs.
	Target WarehouseSyncTargetParam `json:"target,omitzero"`
	// How the sync detects changed and deleted rows.
	ChangeTracking WarehouseSyncChangeTrackingParam `json:"change_tracking,omitzero"`
	// How columns map to Knock fields.
	Mappings []WarehouseSyncMappingParam `json:"mappings,omitzero"`
	// The rows a warehouse sync reads: a table or view (`relation`), or the result of
	// a SQL query (`sql`).
	Source WarehouseSyncSourceParam `json:"source,omitzero"`
	// What each row becomes in Knock. Can't change after creation.
	//
	// Any of "user", "tenant", "object", "static_audience_membership", "subscription".
	TargetType WarehouseSyncRequestTargetType `json:"target_type,omitzero"`
	paramObj
}

func (r WarehouseSyncRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow WarehouseSyncRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WarehouseSyncRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The rows a warehouse sync reads: a table or view (`relation`), or the result of
// a SQL query (`sql`).
type WarehouseSyncSource struct {
	// Whether the sync reads a table or view, or a SQL query.
	//
	// Any of "relation", "sql".
	Mode WarehouseSyncSourceMode `json:"mode" api:"required"`
	// The name of the table or view. Required when `mode` is `relation`.
	Name string `json:"name" api:"nullable"`
	// The SQL query to read. Required when `mode` is `sql`.
	Query string `json:"query" api:"nullable"`
	// The schema of the table or view. Required when `mode` is `relation`.
	Schema string `json:"schema" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		Name        respjson.Field
		Query       respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WarehouseSyncSource) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSyncSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WarehouseSyncSource to a WarehouseSyncSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WarehouseSyncSourceParam.Overrides()
func (r WarehouseSyncSource) ToParam() WarehouseSyncSourceParam {
	return param.Override[WarehouseSyncSourceParam](json.RawMessage(r.RawJSON()))
}

// Whether the sync reads a table or view, or a SQL query.
type WarehouseSyncSourceMode string

const (
	WarehouseSyncSourceModeRelation WarehouseSyncSourceMode = "relation"
	WarehouseSyncSourceModeSql      WarehouseSyncSourceMode = "sql"
)

// The rows a warehouse sync reads: a table or view (`relation`), or the result of
// a SQL query (`sql`).
//
// The property Mode is required.
type WarehouseSyncSourceParam struct {
	// Whether the sync reads a table or view, or a SQL query.
	//
	// Any of "relation", "sql".
	Mode WarehouseSyncSourceMode `json:"mode,omitzero" api:"required"`
	// The name of the table or view. Required when `mode` is `relation`.
	Name param.Opt[string] `json:"name,omitzero"`
	// The SQL query to read. Required when `mode` is `sql`.
	Query param.Opt[string] `json:"query,omitzero"`
	// The schema of the table or view. Required when `mode` is `relation`.
	Schema param.Opt[string] `json:"schema,omitzero"`
	paramObj
}

func (r WarehouseSyncSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow WarehouseSyncSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WarehouseSyncSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The collection or audience the sync writes to. `object` and `subscription` syncs
// set `collection`, and `static_audience_membership` syncs set `audience_key`.
// Null for `user` and `tenant` syncs.
type WarehouseSyncTarget struct {
	// The key of the static audience the rows become members of.
	AudienceKey string `json:"audience_key" api:"nullable"`
	// The object collection the rows are written to.
	Collection string `json:"collection" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudienceKey respjson.Field
		Collection  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WarehouseSyncTarget) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSyncTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WarehouseSyncTarget to a WarehouseSyncTargetParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WarehouseSyncTargetParam.Overrides()
func (r WarehouseSyncTarget) ToParam() WarehouseSyncTargetParam {
	return param.Override[WarehouseSyncTargetParam](json.RawMessage(r.RawJSON()))
}

// The collection or audience the sync writes to. `object` and `subscription` syncs
// set `collection`, and `static_audience_membership` syncs set `audience_key`.
// Null for `user` and `tenant` syncs.
type WarehouseSyncTargetParam struct {
	// The key of the static audience the rows become members of.
	AudienceKey param.Opt[string] `json:"audience_key,omitzero"`
	// The object collection the rows are written to.
	Collection param.Opt[string] `json:"collection,omitzero"`
	paramObj
}

func (r WarehouseSyncTargetParam) MarshalJSON() (data []byte, err error) {
	type shadow WarehouseSyncTargetParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WarehouseSyncTargetParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the WarehouseSync response under the `warehouse_sync` key.
type WarehouseSyncUpdateResponse struct {
	// A warehouse sync reads rows from a warehouse on a schedule and writes them to
	// Knock as users, tenants, objects, audience members, or subscriptions.
	WarehouseSync WarehouseSync `json:"warehouse_sync" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WarehouseSync respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WarehouseSyncUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSyncUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A paginated list of WarehouseSyncRun. Contains a list of entries and page
// information.
type WarehouseSyncGetRunsResponse struct {
	// A list of entries.
	Entries []WarehouseSyncGetRunsResponseEntry `json:"entries" api:"required"`
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
func (r WarehouseSyncGetRunsResponse) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSyncGetRunsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One run of a warehouse sync: reading changed rows from the warehouse, then
// applying them in Knock.
type WarehouseSyncGetRunsResponseEntry struct {
	// The ID of the run.
	ID string `json:"id" api:"required"`
	// The status of the run.
	//
	// Any of "running", "completed", "completed_with_errors", "failed".
	Status string `json:"status" api:"required"`
	// When the run ended.
	EndedAt time.Time `json:"ended_at" api:"nullable" format:"date-time"`
	// Why the run failed.
	Error WarehouseSyncGetRunsResponseEntryError `json:"error" api:"nullable"`
	// The ID of the delivery of rows to Knock. Null when the run hasn't sent rows yet,
	// failed before sending them, or found no changes.
	LoadID string `json:"load_id" api:"nullable"`
	// How many rows failed before reaching Knock.
	RecordsErrored int64 `json:"records_errored" api:"nullable"`
	// How many changed rows were read from the warehouse.
	RecordsProcessed int64 `json:"records_processed" api:"nullable"`
	// How many rows were applied in Knock.
	RowsApplied int64 `json:"rows_applied" api:"nullable"`
	// How many rows failed to apply in Knock.
	RowsFailed int64 `json:"rows_failed" api:"nullable"`
	// When the run started.
	StartedAt time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// When the run was queued.
	SubmittedAt time.Time `json:"submitted_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Status           respjson.Field
		EndedAt          respjson.Field
		Error            respjson.Field
		LoadID           respjson.Field
		RecordsErrored   respjson.Field
		RecordsProcessed respjson.Field
		RowsApplied      respjson.Field
		RowsFailed       respjson.Field
		StartedAt        respjson.Field
		SubmittedAt      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WarehouseSyncGetRunsResponseEntry) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSyncGetRunsResponseEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why the run failed.
type WarehouseSyncGetRunsResponseEntryError struct {
	// What went wrong.
	Message string `json:"message"`
	// Where the run failed.
	//
	// Any of "extract", "load", "apply".
	Stage string `json:"stage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Stage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WarehouseSyncGetRunsResponseEntryError) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSyncGetRunsResponseEntryError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the WarehouseSync response under the `warehouse_sync` key.
type WarehouseSyncUpdateRunResponse struct {
	// A warehouse sync reads rows from a warehouse on a schedule and writes them to
	// Knock as users, tenants, objects, audience members, or subscriptions.
	WarehouseSync WarehouseSync `json:"warehouse_sync" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WarehouseSync respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WarehouseSyncUpdateRunResponse) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSyncUpdateRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the WarehouseSync response under the `warehouse_sync` key.
type WarehouseSyncUpdateSettingsResponse struct {
	// A warehouse sync reads rows from a warehouse on a schedule and writes them to
	// Knock as users, tenants, objects, audience members, or subscriptions.
	WarehouseSync WarehouseSync `json:"warehouse_sync" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WarehouseSync respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WarehouseSyncUpdateSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSyncUpdateSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the WarehouseSyncRequest response under the `warehouse_sync` key.
type WarehouseSyncUpdateValidateResponse struct {
	// A warehouse sync payload for upsert or validate. Fields left out keep their
	// current values. `warehouse_key`, `target_type`, and `target` are required to
	// create a sync and can't change after. Change `enabled` and `frequency_minutes`
	// with the settings endpoint.
	WarehouseSync WarehouseSyncRequest `json:"warehouse_sync" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WarehouseSync respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WarehouseSyncUpdateValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *WarehouseSyncUpdateValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WarehouseSyncGetParams struct {
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	// Whether to hide uncommitted changes. When true, only committed changes will be
	// returned. When false, both committed and uncommitted changes will be returned.
	HideUncommittedChanges param.Opt[bool] `query:"hide_uncommitted_changes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WarehouseSyncGetParams]'s query parameters as `url.Values`.
func (r WarehouseSyncGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WarehouseSyncUpdateParams struct {
	// A warehouse sync payload for upsert or validate. Fields left out keep their
	// current values. `warehouse_key`, `target_type`, and `target` are required to
	// create a sync and can't change after. Change `enabled` and `frequency_minutes`
	// with the settings endpoint.
	WarehouseSync WarehouseSyncRequestParam `json:"warehouse_sync,omitzero" api:"required"`
	// Whether to commit the resource at the same time as modifying it.
	Commit param.Opt[bool] `query:"commit,omitzero" json:"-"`
	// The message to commit the resource with, only used if `commit` is `true`.
	CommitMessage param.Opt[string] `query:"commit_message,omitzero" json:"-"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

func (r WarehouseSyncUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WarehouseSyncUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WarehouseSyncUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WarehouseSyncUpdateParams]'s query parameters as
// `url.Values`.
func (r WarehouseSyncUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WarehouseSyncListParams struct {
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	// Whether to hide uncommitted changes. When true, only committed changes will be
	// returned. When false, both committed and uncommitted changes will be returned.
	HideUncommittedChanges param.Opt[bool] `query:"hide_uncommitted_changes,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WarehouseSyncListParams]'s query parameters as
// `url.Values`.
func (r WarehouseSyncListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WarehouseSyncGetRunsParams struct {
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WarehouseSyncGetRunsParams]'s query parameters as
// `url.Values`.
func (r WarehouseSyncGetRunsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WarehouseSyncUpdateRunParams struct {
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WarehouseSyncUpdateRunParams]'s query parameters as
// `url.Values`.
func (r WarehouseSyncUpdateRunParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WarehouseSyncUpdateSettingsParams struct {
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	// Whether the sync runs on its schedule.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// How often the sync runs, in minutes.
	//
	// Any of 10, 15, 60, 360, 720, 1440.
	FrequencyMinutes int64 `json:"frequency_minutes,omitzero"`
	paramObj
}

func (r WarehouseSyncUpdateSettingsParams) MarshalJSON() (data []byte, err error) {
	type shadow WarehouseSyncUpdateSettingsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WarehouseSyncUpdateSettingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WarehouseSyncUpdateSettingsParams]'s query parameters as
// `url.Values`.
func (r WarehouseSyncUpdateSettingsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WarehouseSyncUpdateValidateParams struct {
	// A warehouse sync payload for upsert or validate. Fields left out keep their
	// current values. `warehouse_key`, `target_type`, and `target` are required to
	// create a sync and can't change after. Change `enabled` and `frequency_minutes`
	// with the settings endpoint.
	WarehouseSync WarehouseSyncRequestParam `json:"warehouse_sync,omitzero" api:"required"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

func (r WarehouseSyncUpdateValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow WarehouseSyncUpdateValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WarehouseSyncUpdateValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WarehouseSyncUpdateValidateParams]'s query parameters as
// `url.Values`.
func (r WarehouseSyncUpdateValidateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
