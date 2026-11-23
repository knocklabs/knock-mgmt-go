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

// Environments are isolated instances of your account that map to your
// infrastructure.
//
// EnvironmentService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentService] method instead.
type EnvironmentService struct {
	Options []option.RequestOption
	// Environments are isolated instances of your account that map to your
	// infrastructure.
	Defaults EnvironmentDefaultService
}

// NewEnvironmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnvironmentService(opts ...option.RequestOption) (r EnvironmentService) {
	r = EnvironmentService{}
	r.Options = opts
	r.Defaults = NewEnvironmentDefaultService(opts...)
	return
}

// Creates an environment. The slug is derived from the `name`. Without
// `preceding_environment`, the environment is placed just before Production.
// Creating an environment also creates its public and secret API keys and copies
// the project's variables, channel and source settings, and preference center into
// it. Everything committed in the preceding environment is then promoted into it
// asynchronously, so its content may not be available as soon as this request
// returns. Requires environments to be enabled on the account.
func (r *EnvironmentService) New(ctx context.Context, body EnvironmentNewParams, opts ...option.RequestOption) (res *Environment, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/environments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single environment by the `environment_slug`.
func (r *EnvironmentService) Get(ctx context.Context, environmentSlug string, opts ...option.RequestOption) (res *Environment, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentSlug == "" {
		err = errors.New("missing required environment_slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/environments/%s", environmentSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an environment. Omitted fields are left unchanged, and at least one
// field is required. What can change depends on the environment's `owner`.
// Environments you create (`user`) can change `name`, `label_color` and
// `hide_pii_data`. Renaming one also changes its slug, and the response holds the
// new slug. Later requests must use the new slug, and permission groups that
// reference the environment are updated to match. Production can change its
// `name`, but its slug stays `production`. Other Knock-defined environments, such
// as Development, can only change `label_color` and `hide_pii_data`. Sending them
// a `name` other than their current one returns an error.
func (r *EnvironmentService) Update(ctx context.Context, environmentSlug string, body EnvironmentUpdateParams, opts ...option.RequestOption) (res *Environment, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentSlug == "" {
		err = errors.New("missing required environment_slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/environments/%s", environmentSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of visible environments in catalog order.
func (r *EnvironmentService) List(ctx context.Context, query EnvironmentListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Environment], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/environments"
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

// Returns a paginated list of visible environments in catalog order.
func (r *EnvironmentService) ListAutoPaging(ctx context.Context, query EnvironmentListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Environment] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Deletes an environment you created (`owner` is `user`). Every API key in the
// environment is revoked first, so any integration still using those keys stops
// working. A deleted environment can't be restored from the API or the dashboard.
// Knock-defined environments, such as Development and Production, can't be
// deleted. To delete a branch, use the delete branch endpoint.
func (r *EnvironmentService) Delete(ctx context.Context, environmentSlug string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if environmentSlug == "" {
		err = errors.New("missing required environment_slug parameter")
		return err
	}
	path := fmt.Sprintf("v1/environments/%s", environmentSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// An environment object.
type Environment struct {
	// The timestamp of when the environment was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A human-readable name for the environment. Cannot exceed 255 characters.
	Name string `json:"name" api:"required"`
	// The order of the environment. The lowest number is the first environment, the
	// highest number is the last environment. The order will not always be sequential.
	Order int64 `json:"order" api:"required"`
	// Who created the environment. `system` environments, such as Development and
	// Production, are defined by Knock: they can't be deleted, and only Production can
	// be renamed. `user` environments were created by you and can be renamed or
	// deleted.
	//
	// Any of "system", "user".
	Owner EnvironmentOwner `json:"owner" api:"required"`
	// A unique slug for the environment. Cannot exceed 255 characters.
	Slug string `json:"slug" api:"required"`
	// The timestamp of when the environment was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Whether the environment is part of the account catalog.
	//
	// Any of "visible", "hidden".
	Visibility EnvironmentVisibility `json:"visibility" api:"required"`
	// The timestamp of when the environment was deleted.
	DeletedAt time.Time `json:"deleted_at" api:"nullable" format:"date-time"`
	// Whether PII data is hidden from the environment. Read more in the
	// [data obfuscation docs](https://docs.knock.app/manage-your-account/data-obfuscation).
	HidePiiData bool `json:"hide_pii_data"`
	// The color of the environment label to display in the dashboard.
	LabelColor string `json:"label_color" api:"nullable"`
	// The timestamp of the most-recent commit in the environment.
	LastCommitAt time.Time `json:"last_commit_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		Name         respjson.Field
		Order        respjson.Field
		Owner        respjson.Field
		Slug         respjson.Field
		UpdatedAt    respjson.Field
		Visibility   respjson.Field
		DeletedAt    respjson.Field
		HidePiiData  respjson.Field
		LabelColor   respjson.Field
		LastCommitAt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Environment) RawJSON() string { return r.JSON.raw }
func (r *Environment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who created the environment. `system` environments, such as Development and
// Production, are defined by Knock: they can't be deleted, and only Production can
// be renamed. `user` environments were created by you and can be renamed or
// deleted.
type EnvironmentOwner string

const (
	EnvironmentOwnerSystem EnvironmentOwner = "system"
	EnvironmentOwnerUser   EnvironmentOwner = "user"
)

// Whether the environment is part of the account catalog.
type EnvironmentVisibility string

const (
	EnvironmentVisibilityVisible EnvironmentVisibility = "visible"
	EnvironmentVisibilityHidden  EnvironmentVisibility = "hidden"
)

type EnvironmentNewParams struct {
	// A human-readable name for the environment. The environment's slug is derived
	// from it and must be unique in the project.
	Name string `json:"name" api:"required"`
	// A 3- or 6-digit hex color code for the environment label in the dashboard. A
	// color is picked at random when omitted.
	LabelColor param.Opt[string] `json:"label_color,omitzero"`
	// The slug of the environment the new environment comes directly after. Cannot be
	// the last environment (Production). Defaults to the environment just before
	// Production.
	PrecedingEnvironment param.Opt[string] `json:"preceding_environment,omitzero"`
	// Whether to hide PII data in the dashboard for the environment. Read more in the
	// [data obfuscation docs](https://docs.knock.app/manage-your-account/data-obfuscation).
	HidePiiData param.Opt[bool] `json:"hide_pii_data,omitzero"`
	paramObj
}

func (r EnvironmentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EnvironmentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnvironmentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnvironmentUpdateParams struct {
	// Whether to hide PII data in the dashboard for the environment. Read more in the
	// [data obfuscation docs](https://docs.knock.app/manage-your-account/data-obfuscation).
	HidePiiData param.Opt[bool] `json:"hide_pii_data,omitzero"`
	// A 3- or 6-digit hex color code for the environment label in the dashboard.
	LabelColor param.Opt[string] `json:"label_color,omitzero"`
	// A new name for the environment. Renaming an environment you created also changes
	// its slug. Production keeps its slug. Other Knock-defined environments can't be
	// renamed, and a name other than their current one returns an error.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r EnvironmentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EnvironmentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnvironmentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnvironmentListParams struct {
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EnvironmentListParams]'s query parameters as `url.Values`.
func (r EnvironmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
