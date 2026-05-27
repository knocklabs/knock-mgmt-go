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
	shimjson "github.com/knocklabs/knock-mgmt-go/internal/encoding/json"
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/pagination"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
	"github.com/knocklabs/knock-mgmt-go/shared"
)

// Sources receive external events that can trigger Knock actions.
//
// DataSourceService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataSourceService] method instead.
type DataSourceService struct {
	Options []option.RequestOption
}

// NewDataSourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDataSourceService(opts ...option.RequestOption) (r DataSourceService) {
	r = DataSourceService{}
	r.Options = opts
	return
}

// Returns a source with environment-specific settings, preprocess scripts, and
// event mappings.
func (r *DataSourceService) Get(ctx context.Context, key string, query DataSourceGetParams, opts ...option.RequestOption) (res *Source, err error) {
	opts = slices.Concat(r.Options, opts)
	if key == "" {
		err = errors.New("missing required key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sources/%s", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns known unique events received by a source in the requested environment.
func (r *DataSourceService) ListEvents(ctx context.Context, key string, query DataSourceListEventsParams, opts ...option.RequestOption) (res *SourceEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if key == "" {
		err = errors.New("missing required key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sources/%s/events", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns event logs received by a source in the requested environment. Supports
// filtering by date/time, event, and log ID.
func (r *DataSourceService) ListLogs(ctx context.Context, key string, query DataSourceListLogsParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[SourceLog], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if key == "" {
		err = errors.New("missing required key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sources/%s/logs", key)
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

// Returns event logs received by a source in the requested environment. Supports
// filtering by date/time, event, and log ID.
func (r *DataSourceService) ListLogsAutoPaging(ctx context.Context, key string, query DataSourceListLogsParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[SourceLog] {
	return pagination.NewEntriesCursorAutoPager(r.ListLogs(ctx, key, query, opts...))
}

// Returns the source provider templates available for creating sources.
func (r *DataSourceService) ListProviders(ctx context.Context, opts ...option.RequestOption) (res *SourceProvidersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/source_providers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns connected sources for the project.
func (r *DataSourceService) ListSources(ctx context.Context, query DataSourceListSourcesParams, opts ...option.RequestOption) (res *SourcesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Sends an arbitrary payload through the source's parse, preprocess, and mapping
// pipeline in the requested environment. This endpoint cannot be run in production
// environments.
func (r *DataSourceService) Rehearse(ctx context.Context, key string, params DataSourceRehearseParams, opts ...option.RequestOption) (res *SourceRehearseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if key == "" {
		err = errors.New("missing required key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sources/%s/rehearse", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a source provider template available for creating sources.
func (r *DataSourceService) GetProvider(ctx context.Context, key string, query DataSourceGetProviderParams, opts ...option.RequestOption) (res *SourceProviderResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if key == "" {
		err = errors.New("missing required key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/source_providers/%s", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns source activity and workflow-trigger mappings that need action in the
// requested environment.
func (r *DataSourceService) GetStatus(ctx context.Context, key string, query DataSourceGetStatusParams, opts ...option.RequestOption) (res *SourceStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if key == "" {
		err = errors.New("missing required key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sources/%s/status", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Creates or updates a source with environment-specific settings, preprocess
// scripts, and event mappings.
func (r *DataSourceService) Upsert(ctx context.Context, key string, params DataSourceUpsertParams, opts ...option.RequestOption) (res *DataSourceUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if key == "" {
		err = errors.New("missing required key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sources/%s", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// A source that receives external events and maps them to Knock actions.
type Source struct {
	// The timestamp of when the source was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Per-environment settings keyed by environment slug.
	EnvironmentSettings map[string]SourceEnvironmentSettings `json:"environment_settings" api:"required"`
	// The unique key for the source within the project.
	Key string `json:"key" api:"required"`
	// The human-readable name of the source.
	Name string `json:"name" api:"required"`
	// The timestamp of when the source was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional URL for a custom image representing the source.
	CustomImageURL string `json:"custom_image_url" api:"nullable"`
	// An optional description of the source.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt           respjson.Field
		EnvironmentSettings respjson.Field
		Key                 respjson.Field
		Name                respjson.Field
		UpdatedAt           respjson.Field
		CustomImageURL      respjson.Field
		Description         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Source) RawJSON() string { return r.JSON.raw }
func (r *Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Environment-specific settings for a source.
type SourceEnvironmentSettings struct {
	// The timestamp of when these environment settings were created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Event action mappings configured for this source in the environment.
	Mappings []SourceEventActionMapping `json:"mappings" api:"required"`
	// Source settings for this environment, including endpoint, verification behavior,
	// and preprocess script.
	Settings SourceEnvironmentSettingsSettings `json:"settings" api:"required"`
	// The timestamp of when these environment settings were last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Mappings    respjson.Field
		Settings    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceEnvironmentSettings) RawJSON() string { return r.JSON.raw }
func (r *SourceEnvironmentSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source settings for this environment, including endpoint, verification behavior,
// and preprocess script.
type SourceEnvironmentSettingsSettings struct {
	// The public endpoint that receives source events for this environment.
	Endpoint            string `json:"endpoint"`
	EnforceIdempotency  bool   `json:"enforce_idempotency" api:"nullable"`
	EnforceVerification bool   `json:"enforce_verification"`
	EventTypePath       string `json:"event_type_path" api:"nullable"`
	HandleIdentifies    bool   `json:"handle_identifies" api:"nullable"`
	IdempotencyKeyPath  string `json:"idempotency_key_path" api:"nullable"`
	// A script that runs before source events are mapped.
	PreprocessScript SourcePreprocessScript `json:"preprocess_script" api:"nullable"`
	TimestampPath    string                 `json:"timestamp_path" api:"nullable"`
	ExtraFields      map[string]any         `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Endpoint            respjson.Field
		EnforceIdempotency  respjson.Field
		EnforceVerification respjson.Field
		EventTypePath       respjson.Field
		HandleIdentifies    respjson.Field
		IdempotencyKeyPath  respjson.Field
		PreprocessScript    respjson.Field
		TimestampPath       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceEnvironmentSettingsSettings) RawJSON() string { return r.JSON.raw }
func (r *SourceEnvironmentSettingsSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A known unique event received by a source.
type SourceEvent struct {
	// The decoded event name.
	Event string `json:"event" api:"required"`
	// The timestamp of when this event was last received.
	LastReceivedAt time.Time `json:"last_received_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event          respjson.Field
		LastReceivedAt respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceEvent) RawJSON() string { return r.JSON.raw }
func (r *SourceEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An action mapping attached to a source event.
type SourceEventActionMapping struct {
	// The action that is performed when this mapping matches a source event.
	//
	// Any of "workflows_trigger", "users_identify", "users_delete", "objects_set",
	// "objects_delete", "objects_subscribe", "objects_unsubscribe", "tenants_set",
	// "tenants_delete", "audiences_add_member", "audiences_remove_member".
	ActionType SourceEventActionMappingActionType `json:"action_type" api:"required"`
	// The timestamp of when the mapping was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The decoded event type that triggers the action.
	EventType string `json:"event_type" api:"required"`
	// Whether the mapping has been deleted. When true, this indicates the trigger is
	// present in the workflow's published version and may be active until the workflow
	// is committed and published.
	IsDeleted bool `json:"is_deleted" api:"required"`
	// The timestamp of when the mapping was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The action-specific parameters for the mapping.
	ActionParameters map[string]any `json:"action_parameters" api:"nullable"`
	// The timestamp of when the mapping was deactivated.
	InactiveAt time.Time `json:"inactive_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType       respjson.Field
		CreatedAt        respjson.Field
		EventType        respjson.Field
		IsDeleted        respjson.Field
		UpdatedAt        respjson.Field
		ActionParameters respjson.Field
		InactiveAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceEventActionMapping) RawJSON() string { return r.JSON.raw }
func (r *SourceEventActionMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The action that is performed when this mapping matches a source event.
type SourceEventActionMappingActionType string

const (
	SourceEventActionMappingActionTypeWorkflowsTrigger      SourceEventActionMappingActionType = "workflows_trigger"
	SourceEventActionMappingActionTypeUsersIdentify         SourceEventActionMappingActionType = "users_identify"
	SourceEventActionMappingActionTypeUsersDelete           SourceEventActionMappingActionType = "users_delete"
	SourceEventActionMappingActionTypeObjectsSet            SourceEventActionMappingActionType = "objects_set"
	SourceEventActionMappingActionTypeObjectsDelete         SourceEventActionMappingActionType = "objects_delete"
	SourceEventActionMappingActionTypeObjectsSubscribe      SourceEventActionMappingActionType = "objects_subscribe"
	SourceEventActionMappingActionTypeObjectsUnsubscribe    SourceEventActionMappingActionType = "objects_unsubscribe"
	SourceEventActionMappingActionTypeTenantsSet            SourceEventActionMappingActionType = "tenants_set"
	SourceEventActionMappingActionTypeTenantsDelete         SourceEventActionMappingActionType = "tenants_delete"
	SourceEventActionMappingActionTypeAudiencesAddMember    SourceEventActionMappingActionType = "audiences_add_member"
	SourceEventActionMappingActionTypeAudiencesRemoveMember SourceEventActionMappingActionType = "audiences_remove_member"
)

// A list of known unique source events.
type SourceEventsResponse struct {
	// The known unique events for the requested source and environment.
	Entries []SourceEvent `json:"entries" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A log entry for an event received by a source.
type SourceLog struct {
	// The source log ID.
	ID string `json:"id" api:"required"`
	// The decoded source event name.
	Event string `json:"event" api:"required"`
	// The actions executed after receiving the source event. Only present when
	// `includes` contains `actions`.
	Actions []SourceLogAction `json:"actions"`
	// The data payload parsed by the source.
	Data map[string]any `json:"data" api:"nullable"`
	// The timestamp of when the source log was created.
	InsertedAt time.Time `json:"inserted_at" api:"nullable" format:"date-time"`
	// The output returned by the preprocess script.
	PreprocessOutput map[string]any `json:"preprocess_output" api:"nullable"`
	// Indicates the origin of the log; if the log is a product of a test event. This
	// is typically null for regular source events received from the data source.
	Source string `json:"source" api:"nullable"`
	// The verification status for the received event.
	VerificationStatus string `json:"verification_status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Event              respjson.Field
		Actions            respjson.Field
		Data               respjson.Field
		InsertedAt         respjson.Field
		PreprocessOutput   respjson.Field
		Source             respjson.Field
		VerificationStatus respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceLog) RawJSON() string { return r.JSON.raw }
func (r *SourceLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An action executed after receiving a source event.
type SourceLogAction struct {
	// The action log ID.
	ID string `json:"id" api:"required"`
	// The configured mapping parameters used to derive the action payload.
	ActionParameters map[string]any `json:"action_parameters" api:"nullable"`
	// The parsed values generated from the mapping parameters for this action.
	ActionPayload map[string]any `json:"action_payload" api:"nullable"`
	// The result returned by the action.
	ActionResult map[string]any `json:"action_result" api:"nullable"`
	// The status of the action.
	ActionStatus string `json:"action_status" api:"nullable"`
	// The type of action that was executed.
	ActionType string `json:"action_type"`
	// The timestamp of when the action log was created.
	InsertedAt time.Time `json:"inserted_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ActionParameters respjson.Field
		ActionPayload    respjson.Field
		ActionResult     respjson.Field
		ActionStatus     respjson.Field
		ActionType       respjson.Field
		InsertedAt       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceLogAction) RawJSON() string { return r.JSON.raw }
func (r *SourceLogAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A paginated list of source logs. Include `actions` in the `includes` query
// parameter to return action details for each log.
type SourceLogsResponse struct {
	// The source logs for the requested source and environment.
	Entries []SourceLog `json:"entries" api:"required"`
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
func (r SourceLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A script that runs before source events are mapped.
type SourcePreprocessScript struct {
	// The programming language used by the preprocess script.
	//
	// Any of "javascript".
	Language SourcePreprocessScriptLanguage `json:"language" api:"required"`
	// The source code for the preprocess script.
	Source string `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourcePreprocessScript) RawJSON() string { return r.JSON.raw }
func (r *SourcePreprocessScript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SourcePreprocessScript to a SourcePreprocessScriptParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SourcePreprocessScriptParam.Overrides()
func (r SourcePreprocessScript) ToParam() SourcePreprocessScriptParam {
	return param.Override[SourcePreprocessScriptParam](json.RawMessage(r.RawJSON()))
}

// The programming language used by the preprocess script.
type SourcePreprocessScriptLanguage string

const (
	SourcePreprocessScriptLanguageJavascript SourcePreprocessScriptLanguage = "javascript"
)

// A script that runs before source events are mapped.
//
// The properties Language, Source are required.
type SourcePreprocessScriptParam struct {
	// The programming language used by the preprocess script.
	//
	// Any of "javascript".
	Language SourcePreprocessScriptLanguage `json:"language,omitzero" api:"required"`
	// The source code for the preprocess script.
	Source string `json:"source" api:"required"`
	paramObj
}

func (r SourcePreprocessScriptParam) MarshalJSON() (data []byte, err error) {
	type shadow SourcePreprocessScriptParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SourcePreprocessScriptParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A source provider available for creating sources.
type SourceProviderResponse struct {
	// Default source settings for this provider.
	DefaultSettings SourceProviderResponseDefaultSettings `json:"default_settings" api:"required"`
	// Provider key.
	Key string `json:"key" api:"required"`
	// Provider display metadata.
	Provider SourceProviderResponseProvider `json:"provider" api:"required"`
	// Provider version.
	Version string `json:"version" api:"required"`
	// Default event action mappings for the provider. Only present when `includes`
	// contains `default_action_mappings`.
	DefaultActionMappings []SourceProviderResponseDefaultActionMapping `json:"default_action_mappings"`
	// Example payloads keyed by event type.
	ExamplePayloads map[string][]SourceProviderResponseExamplePayload `json:"example_payloads" api:"nullable"`
	// JSON Schema fields needed to configure the source. Only present when `includes`
	// contains `static_fields`.
	StaticFields map[string]any `json:"static_fields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultSettings       respjson.Field
		Key                   respjson.Field
		Provider              respjson.Field
		Version               respjson.Field
		DefaultActionMappings respjson.Field
		ExamplePayloads       respjson.Field
		StaticFields          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceProviderResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceProviderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default source settings for this provider.
type SourceProviderResponseDefaultSettings struct {
	// Whether the source should enforce webhook verification.
	EnforceVerification bool `json:"enforce_verification" api:"required"`
	// Path to find the event type from the data.
	EventTypePath string `json:"event_type_path" api:"required"`
	// Path to find the idempotency key from the data.
	IdempotencyKeyPath string `json:"idempotency_key_path" api:"nullable"`
	// Verification script source code. Only present when `includes` contains
	// `preprocessing_script`.
	PreprocessingScript SourceProviderResponseDefaultSettingsPreprocessingScript `json:"preprocessing_script" api:"nullable"`
	// Path to find the timestamp from the data.
	TimestampPath string `json:"timestamp_path" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnforceVerification respjson.Field
		EventTypePath       respjson.Field
		IdempotencyKeyPath  respjson.Field
		PreprocessingScript respjson.Field
		TimestampPath       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceProviderResponseDefaultSettings) RawJSON() string { return r.JSON.raw }
func (r *SourceProviderResponseDefaultSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification script source code. Only present when `includes` contains
// `preprocessing_script`.
type SourceProviderResponseDefaultSettingsPreprocessingScript struct {
	// Script language.
	//
	// Any of "javascript".
	Language string `json:"language" api:"required"`
	// Verification script source code.
	Source string `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceProviderResponseDefaultSettingsPreprocessingScript) RawJSON() string { return r.JSON.raw }
func (r *SourceProviderResponseDefaultSettingsPreprocessingScript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provider display metadata.
type SourceProviderResponseProvider struct {
	// Provider categories for filtering and grouping.
	//
	// Any of "Billing", "Infrastructure", "Analytics", "CRM", "Ecommerce",
	// "Communications", "Identity".
	Categories []string `json:"categories" api:"required"`
	// Provider display description.
	Description string `json:"description" api:"required"`
	// Provider display name.
	Name string `json:"name" api:"required"`
	// Provider webhook documentation URL.
	WebhookDocsURL string `json:"webhook_docs_url" api:"required"`
	// Provider website URL.
	WebsiteURL string `json:"website_url" api:"required"`
	// Provider branding assets. Only present when `includes` contains `branding`.
	Branding SourceProviderResponseProviderBranding `json:"branding" api:"nullable"`
	// Knock tutorial URL for setting up the provider.
	KnockTutorialURL string `json:"knock_tutorial_url" api:"nullable"`
	// Provider webhook configuration URL.
	WebhookConfigDeepLink string `json:"webhook_config_deep_link" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories            respjson.Field
		Description           respjson.Field
		Name                  respjson.Field
		WebhookDocsURL        respjson.Field
		WebsiteURL            respjson.Field
		Branding              respjson.Field
		KnockTutorialURL      respjson.Field
		WebhookConfigDeepLink respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceProviderResponseProvider) RawJSON() string { return r.JSON.raw }
func (r *SourceProviderResponseProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provider branding assets. Only present when `includes` contains `branding`.
type SourceProviderResponseProviderBranding struct {
	// Provider icon image URL or path.
	IconImage      string                                               `json:"icon_image" api:"required"`
	WordmarkImages SourceProviderResponseProviderBrandingWordmarkImages `json:"wordmark_images" api:"required"`
	Colors         SourceProviderResponseProviderBrandingColors         `json:"colors" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IconImage      respjson.Field
		WordmarkImages respjson.Field
		Colors         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceProviderResponseProviderBranding) RawJSON() string { return r.JSON.raw }
func (r *SourceProviderResponseProviderBranding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceProviderResponseProviderBrandingWordmarkImages struct {
	// Wordmark image URL or path for dark backgrounds.
	Dark string `json:"dark" api:"required"`
	// Wordmark image URL or path for light backgrounds.
	Light string `json:"light" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dark        respjson.Field
		Light       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceProviderResponseProviderBrandingWordmarkImages) RawJSON() string { return r.JSON.raw }
func (r *SourceProviderResponseProviderBrandingWordmarkImages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceProviderResponseProviderBrandingColors struct {
	// Primary brand color.
	Primary string `json:"primary" api:"required"`
	// Secondary brand color.
	Secondary string `json:"secondary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		Secondary   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceProviderResponseProviderBrandingColors) RawJSON() string { return r.JSON.raw }
func (r *SourceProviderResponseProviderBrandingColors) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceProviderResponseDefaultActionMapping struct {
	// Action-specific data paths and options.
	ActionParameters map[string]any `json:"action_parameters" api:"required"`
	// The action performed when the mapping matches a source event.
	//
	// Any of "workflows_trigger", "users_identify", "users_delete", "objects_set",
	// "objects_delete", "objects_subscribe", "objects_unsubscribe", "tenants_set",
	// "tenants_delete", "audiences_add_member", "audiences_remove_member".
	ActionType string `json:"action_type" api:"required"`
	// Event type to match.
	EventType string `json:"event_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionParameters respjson.Field
		ActionType       respjson.Field
		EventType        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceProviderResponseDefaultActionMapping) RawJSON() string { return r.JSON.raw }
func (r *SourceProviderResponseDefaultActionMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceProviderResponseExamplePayload struct {
	// Example payload body.
	Body map[string]any `json:"body" api:"nullable"`
	// Example payload headers.
	Headers map[string]any `json:"headers" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceProviderResponseExamplePayload) RawJSON() string { return r.JSON.raw }
func (r *SourceProviderResponseExamplePayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source providers available for creating sources.
type SourceProvidersResponse struct {
	// Source providers.
	Entries []SourceProvidersResponseEntry `json:"entries" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceProvidersResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceProvidersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A source provider summary.
type SourceProvidersResponseEntry struct {
	// Provider key.
	Key string `json:"key" api:"required"`
	// Provider display metadata.
	Provider SourceProvidersResponseEntryProvider `json:"provider" api:"required"`
	// Provider version.
	Version string `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Provider    respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceProvidersResponseEntry) RawJSON() string { return r.JSON.raw }
func (r *SourceProvidersResponseEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provider display metadata.
type SourceProvidersResponseEntryProvider struct {
	// Provider display description.
	Description string `json:"description" api:"required"`
	// Provider display name.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceProvidersResponseEntryProvider) RawJSON() string { return r.JSON.raw }
func (r *SourceProvidersResponseEntryProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for rehearsing a source event.
//
// The property Payload is required.
type SourceRehearseRequestParam struct {
	// An arbitrary payload to send through the source's parse, preprocess, and mapping
	// pipeline.
	Payload map[string]any `json:"payload,omitzero" api:"required"`
	paramObj
}

func (r SourceRehearseRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SourceRehearseRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SourceRehearseRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result of a simulated source event rehearsal.
type SourceRehearseResponse struct {
	// The total number of action logs produced by the rehearsal.
	ActionLogsCount int64 `json:"action_logs_count" api:"required"`
	// The number of failed action logs produced by the rehearsal.
	FailedActionLogsCount int64 `json:"failed_action_logs_count" api:"required"`
	// The ID of the source event log created by the rehearsal.
	LogID string `json:"log_id" api:"required"`
	// Whether the rehearsal completed without action errors.
	//
	// Any of "ok", "error".
	Status SourceRehearseResponseStatus `json:"status" api:"required"`
	// The number of successful action logs produced by the rehearsal.
	SuccessfulActionLogsCount int64 `json:"successful_action_logs_count" api:"required"`
	// Errors returned while rehearsing the source event.
	Errors []SourceRehearseResponseError `json:"errors" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionLogsCount           respjson.Field
		FailedActionLogsCount     respjson.Field
		LogID                     respjson.Field
		Status                    respjson.Field
		SuccessfulActionLogsCount respjson.Field
		Errors                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceRehearseResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceRehearseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the rehearsal completed without action errors.
type SourceRehearseResponseStatus string

const (
	SourceRehearseResponseStatusOk    SourceRehearseResponseStatus = "ok"
	SourceRehearseResponseStatusError SourceRehearseResponseStatus = "error"
)

type SourceRehearseResponseError struct {
	Field   string `json:"field" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceRehearseResponseError) RawJSON() string { return r.JSON.raw }
func (r *SourceRehearseResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A source request for setting a source and its environment-specific
// configuration.
//
// The property Name is required.
type SourceRequestParam struct {
	// The human-readable name of the source.
	Name string `json:"name" api:"required"`
	// An optional URL for a custom image representing the source.
	CustomImageURL param.Opt[string] `json:"custom_image_url,omitzero"`
	// An optional description of the source.
	Description param.Opt[string] `json:"description,omitzero"`
	// When creating a source, bootstraps configuration from a preconfigured provider
	// template. Ignored when updating an existing source.
	PreconfiguredProvider param.Opt[string] `json:"preconfigured_provider,omitzero"`
	// Per-environment settings keyed by environment slug.
	EnvironmentSettings map[string]SourceRequestEnvironmentSettingParam `json:"environment_settings,omitzero"`
	paramObj
}

func (r SourceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SourceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SourceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Environment-specific source settings to configure.
type SourceRequestEnvironmentSettingParam struct {
	// Event action mappings to configure for this source in the environment.
	Mappings []SourceRequestEnvironmentSettingMappingParam `json:"mappings,omitzero"`
	// Writable source settings for this environment.
	Settings SourceRequestEnvironmentSettingSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r SourceRequestEnvironmentSettingParam) MarshalJSON() (data []byte, err error) {
	type shadow SourceRequestEnvironmentSettingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SourceRequestEnvironmentSettingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An action mapping to configure for a source event.
//
// The properties ActionType, EventType are required.
type SourceRequestEnvironmentSettingMappingParam struct {
	// The action that is performed when this mapping matches a source event.
	//
	// Any of "workflows_trigger", "users_identify", "users_delete", "objects_set",
	// "objects_delete", "objects_subscribe", "objects_unsubscribe", "tenants_set",
	// "tenants_delete", "audiences_add_member", "audiences_remove_member".
	ActionType string `json:"action_type,omitzero" api:"required"`
	// The decoded event type that triggers the action.
	EventType string `json:"event_type" api:"required"`
	// The timestamp to deactivate the mapping.
	InactiveAt param.Opt[time.Time] `json:"inactive_at,omitzero" format:"date-time"`
	// Whether to delete the mapping. Workflow trigger mappings must be marked deleted
	// before they can be removed.
	IsDeleted param.Opt[bool] `json:"is_deleted,omitzero"`
	// The action-specific parameters for the mapping.
	ActionParameters map[string]any `json:"action_parameters,omitzero"`
	paramObj
}

func (r SourceRequestEnvironmentSettingMappingParam) MarshalJSON() (data []byte, err error) {
	type shadow SourceRequestEnvironmentSettingMappingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SourceRequestEnvironmentSettingMappingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SourceRequestEnvironmentSettingMappingParam](
		"action_type", "workflows_trigger", "users_identify", "users_delete", "objects_set", "objects_delete", "objects_subscribe", "objects_unsubscribe", "tenants_set", "tenants_delete", "audiences_add_member", "audiences_remove_member",
	)
}

// Writable source settings for this environment.
type SourceRequestEnvironmentSettingSettingsParam struct {
	EventTypePath       param.Opt[string] `json:"event_type_path,omitzero"`
	IdempotencyKeyPath  param.Opt[string] `json:"idempotency_key_path,omitzero"`
	TimestampPath       param.Opt[string] `json:"timestamp_path,omitzero"`
	EnforceVerification param.Opt[bool]   `json:"enforce_verification,omitzero"`
	// A script that runs before source events are mapped.
	PreprocessScript SourcePreprocessScriptParam `json:"preprocess_script,omitzero"`
	ExtraFields      map[string]any              `json:"-"`
	paramObj
}

func (r SourceRequestEnvironmentSettingSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow SourceRequestEnvironmentSettingSettingsParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *SourceRequestEnvironmentSettingSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status information for a source in an environment.
type SourceStatusResponse struct {
	// The number of source events received in the last 30 days.
	EventsCount int64 `json:"events_count" api:"required"`
	// The timestamp of the most recently received source event.
	LastEventReceived string `json:"last_event_received" api:"required"`
	// The total number of event action mappings for the source environment.
	MappingsCount int64 `json:"mappings_count" api:"required"`
	// Workflow trigger event action mappings that need a workflow commit before their
	// changes are applied.
	MappingsRequiringCommit []SourceStatusResponseMappingsRequiringCommit `json:"mappings_requiring_commit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventsCount             respjson.Field
		LastEventReceived       respjson.Field
		MappingsCount           respjson.Field
		MappingsRequiringCommit respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceStatusResponseMappingsRequiringCommit struct {
	// The action that is performed when this mapping matches a source event.
	//
	// Any of "workflows_trigger".
	ActionType string `json:"action_type" api:"required"`
	// The decoded event type that triggers the action.
	EventType string `json:"event_type" api:"required"`
	// Whether the mapping is pending deletion.
	IsDeleted bool `json:"is_deleted" api:"required"`
	// The key of the workflow resource referenced by the mapping.
	ResourceKey string `json:"resource_key" api:"required"`
	// Whether the mapping is pending deletion or update.
	//
	// Any of "deleted", "updated".
	Status string `json:"status" api:"required"`
	// The timestamp of when the mapping was deactivated.
	InactiveAt time.Time `json:"inactive_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		EventType   respjson.Field
		IsDeleted   respjson.Field
		ResourceKey respjson.Field
		Status      respjson.Field
		InactiveAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceStatusResponseMappingsRequiringCommit) RawJSON() string { return r.JSON.raw }
func (r *SourceStatusResponseMappingsRequiringCommit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sources connected to the project.
type SourcesResponse struct {
	// Sources.
	Entries []SourcesResponseEntry `json:"entries" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourcesResponse) RawJSON() string { return r.JSON.raw }
func (r *SourcesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A source summary.
type SourcesResponseEntry struct {
	// Source key.
	Key string `json:"key" api:"required"`
	// Source display name.
	Name string `json:"name" api:"required"`
	// Custom image URL for the source.
	CustomImageURL string `json:"custom_image_url" api:"nullable"`
	// Source description.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key            respjson.Field
		Name           respjson.Field
		CustomImageURL respjson.Field
		Description    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourcesResponseEntry) RawJSON() string { return r.JSON.raw }
func (r *SourcesResponseEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Source response under the `source` key.
type DataSourceUpsertResponse struct {
	// A source that receives external events and maps them to Knock actions.
	Source Source `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataSourceUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *DataSourceUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataSourceGetParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [DataSourceGetParams]'s query parameters as `url.Values`.
func (r DataSourceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataSourceListEventsParams struct {
	// The environment slug.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataSourceListEventsParams]'s query parameters as
// `url.Values`.
func (r DataSourceListEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataSourceListLogsParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// The log ID to filter by.
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Returns event logs that were produced on this date.
	Date param.Opt[string] `query:"date,omitzero" json:"-"`
	// Only return source logs at or before this timestamp.
	EndingAt param.Opt[time.Time] `query:"ending_at,omitzero" format:"date-time" json:"-"`
	// The event name to filter by.
	Event param.Opt[string] `query:"event,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only return source logs at or after this timestamp.
	StartingAt param.Opt[time.Time] `query:"starting_at,omitzero" format:"date-time" json:"-"`
	// Associated resources to include in the response. Accepts `actions` to include
	// the actions executed after receiving each source event.
	//
	// Any of "actions".
	Includes []string `query:"includes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataSourceListLogsParams]'s query parameters as
// `url.Values`.
func (r DataSourceListLogsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataSourceListSourcesParams struct {
	// The environment slug.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataSourceListSourcesParams]'s query parameters as
// `url.Values`.
func (r DataSourceListSourcesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataSourceRehearseParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// Request body for rehearsing a source event.
	SourceRehearseRequest SourceRehearseRequestParam
	paramObj
}

func (r DataSourceRehearseParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SourceRehearseRequest)
}
func (r *DataSourceRehearseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [DataSourceRehearseParams]'s query parameters as
// `url.Values`.
func (r DataSourceRehearseParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataSourceGetProviderParams struct {
	// Associated resources to include in the response. Accepts `branding`,
	// `default_action_mappings`, `example_payloads`, `preprocessing_script`,
	// `static_fields`.
	//
	// Any of "branding", "default_action_mappings", "example_payloads",
	// "preprocessing_script", "static_fields".
	Includes []string `query:"includes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataSourceGetProviderParams]'s query parameters as
// `url.Values`.
func (r DataSourceGetProviderParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataSourceGetStatusParams struct {
	// The environment slug.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataSourceGetStatusParams]'s query parameters as
// `url.Values`.
func (r DataSourceGetStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataSourceUpsertParams struct {
	// The environment slug.
	Environment string `query:"environment" api:"required" json:"-"`
	// A source request for setting a source and its environment-specific
	// configuration.
	Source SourceRequestParam `json:"source,omitzero" api:"required"`
	paramObj
}

func (r DataSourceUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow DataSourceUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DataSourceUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [DataSourceUpsertParams]'s query parameters as `url.Values`.
func (r DataSourceUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
