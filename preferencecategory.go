// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/knocklabs/knock-mgmt-go/internal/apijson"
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// Preference categories are a project-level catalog of categories that can be
// applied to workflows and broadcasts.
//
// PreferenceCategoryService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPreferenceCategoryService] method instead.
type PreferenceCategoryService struct {
	Options []option.RequestOption
}

// NewPreferenceCategoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPreferenceCategoryService(opts ...option.RequestOption) (r PreferenceCategoryService) {
	r = PreferenceCategoryService{}
	r.Options = opts
	return
}

// Returns all preference categories in the project's catalog, ordered by name.
// Preference categories are project-scoped and not tied to an environment.
func (r *PreferenceCategoryService) List(ctx context.Context, opts ...option.RequestOption) (res *PreferenceCategoryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/preference_categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Archives a preference category by name.
func (r *PreferenceCategoryService) Delete(ctx context.Context, name string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if name == "" {
		err = errors.New("missing required name parameter")
		return err
	}
	path := fmt.Sprintf("v1/preference_categories/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Creates a preference category by name. If a non-archived category with the same
// name already exists, returns the existing category.
func (r *PreferenceCategoryService) Upsert(ctx context.Context, name string, opts ...option.RequestOption) (res *PreferenceCategoryUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/preference_categories/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// A named preference category in a project's catalog.
type PreferenceCategory struct {
	// The timestamp of when the preference category was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The unique name of the preference category within a project.
	Name string `json:"name" api:"required"`
	// The timestamp of when the preference category was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The timestamp of when the preference category was archived.
	ArchivedAt time.Time `json:"archived_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ArchivedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreferenceCategory) RawJSON() string { return r.JSON.raw }
func (r *PreferenceCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A list of preference categories in the project's catalog.
type PreferenceCategoryListResponse struct {
	// Preference categories, ordered by name.
	Entries []PreferenceCategory `json:"entries" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreferenceCategoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *PreferenceCategoryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the PreferenceCategory response under the `preference_category` key.
type PreferenceCategoryUpsertResponse struct {
	// A named preference category in a project's catalog.
	PreferenceCategory PreferenceCategory `json:"preference_category" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PreferenceCategory respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreferenceCategoryUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *PreferenceCategoryUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
