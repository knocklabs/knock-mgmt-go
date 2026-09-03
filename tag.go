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
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// Tags are a project-level catalog of labels that can be applied to workflows,
// partials, guides, and broadcasts.
//
// TagService contains methods and other services that help with interacting with
// the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagService] method instead.
type TagService struct {
	Options []option.RequestOption
}

// NewTagService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTagService(opts ...option.RequestOption) (r TagService) {
	r = TagService{}
	r.Options = opts
	return
}

// Returns all tags in the project's catalog, ordered by name. Tags are
// project-scoped and not tied to an environment.
func (r *TagService) List(ctx context.Context, opts ...option.RequestOption) (res *TagListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deletes a tag by name.
func (r *TagService) Delete(ctx context.Context, name string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if name == "" {
		err = errors.New("missing required name parameter")
		return err
	}
	path := fmt.Sprintf("v1/tags/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Creates a tag by name. If a tag with the same name already exists, updates its
// description and color. Omitted description and color fields are set to null.
func (r *TagService) Upsert(ctx context.Context, name string, body TagUpsertParams, opts ...option.RequestOption) (res *TagUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tags/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// A named tag in a project's resource-tag catalog.
type Tag struct {
	// The timestamp of when the tag was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The unique name of the tag within a project.
	Name string `json:"name" api:"required"`
	// The timestamp of when the tag was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// An optional hex color for the tag (e.g. #3B82F6).
	Color string `json:"color" api:"nullable"`
	// An optional description of the tag.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		Color       respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tag) RawJSON() string { return r.JSON.raw }
func (r *Tag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A list of tags in the project's catalog.
type TagListResponse struct {
	// Tags, ordered by name.
	Entries []Tag `json:"entries" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagListResponse) RawJSON() string { return r.JSON.raw }
func (r *TagListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wraps the Tag response under the `tag` key.
type TagUpsertResponse struct {
	// A named tag in a project's resource-tag catalog.
	Tag Tag `json:"tag" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tag         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *TagUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagUpsertParams struct {
	// A request to create or update a tag. The tag name is taken from the path. On
	// conflict, omitted description and color fields are set to null.
	Tag TagUpsertParamsTag `json:"tag,omitzero" api:"required"`
	paramObj
}

func (r TagUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow TagUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to create or update a tag. The tag name is taken from the path. On
// conflict, omitted description and color fields are set to null.
type TagUpsertParamsTag struct {
	// An optional hex color for the tag (e.g. #3B82F6).
	Color param.Opt[string] `json:"color,omitzero"`
	// An optional description of the tag.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r TagUpsertParamsTag) MarshalJSON() (data []byte, err error) {
	type shadow TagUpsertParamsTag
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagUpsertParamsTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
