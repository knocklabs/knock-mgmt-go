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

// Assets are uploaded files available to your Knock account.
//
// AssetService contains methods and other services that help with interacting with
// the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetService] method instead.
type AssetService struct {
	Options []option.RequestOption
}

// NewAssetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAssetService(opts ...option.RequestOption) (r AssetService) {
	r = AssetService{}
	r.Options = opts
	return
}

// Records a library asset after the file has been uploaded to the key from an
// upload request. The asset is always a library asset. Assets are account-wide and
// do not require an environment parameter.
func (r *AssetService) New(ctx context.Context, body AssetNewParams, opts ...option.RequestOption) (res *Asset, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/assets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of active account assets. Assets are account-wide and
// do not require an environment parameter.
func (r *AssetService) List(ctx context.Context, query AssetListParams, opts ...option.RequestOption) (res *pagination.EntriesCursor[Asset], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/assets"
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

// Returns a paginated list of active account assets. Assets are account-wide and
// do not require an environment parameter.
func (r *AssetService) ListAutoPaging(ctx context.Context, query AssetListParams, opts ...option.RequestOption) *pagination.EntriesCursorAutoPager[Asset] {
	return pagination.NewEntriesCursorAutoPager(r.List(ctx, query, opts...))
}

// Archives a library or template image asset by id. The stored file remains at its
// URL. Other asset types, unknown ids, and already archived assets are not found.
// Assets are account-wide and do not require an environment parameter.
func (r *AssetService) Delete(ctx context.Context, assetID string, opts ...option.RequestOption) (res *AssetDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assetID == "" {
		err = errors.New("missing required asset_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/assets/%s", assetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns a time-limited URL for uploading an asset file. PUT the file bytes to
// upload_url with Content-Type and Content-Length from this response, then create
// the asset with the returned key. This request does not store an asset. Assets
// are account-wide and do not require an environment parameter.
func (r *AssetService) Uploads(ctx context.Context, body AssetUploadsParams, opts ...option.RequestOption) (res *AssetUploadsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/assets/uploads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// An uploaded asset file for an account.
type Asset struct {
	// The unique ID for this asset.
	ID string `json:"id" api:"required" format:"uuid"`
	// The type of asset.
	AssetType string `json:"asset_type" api:"required"`
	// The timestamp of when this asset was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The MIME type for this asset.
	MimeType string `json:"mime_type" api:"required"`
	// The timestamp of when this asset was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The public URL for this asset.
	URL string `json:"url" api:"required"`
	// The human-readable filename for this asset.
	Filename string `json:"filename" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AssetType   respjson.Field
		CreatedAt   respjson.Field
		MimeType    respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Asset) RawJSON() string { return r.JSON.raw }
func (r *Asset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response from archiving an asset.
type AssetDeleteResponse struct {
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
func (r AssetDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A time-limited permission to upload an asset file directly to storage.
type AssetUploadsResponse struct {
	// The byte length the PUT must send as Content-Length.
	ContentLength int64 `json:"content_length" api:"required"`
	// Seconds until the upload URL expires.
	ExpiresIn int64 `json:"expires_in" api:"required"`
	// The storage key to send when creating the asset.
	Key string `json:"key" api:"required"`
	// The MIME type the PUT must send as Content-Type.
	MimeType string `json:"mime_type" api:"required"`
	// The URL to PUT the file bytes to. The signature covers Content-Type and
	// Content-Length, and the PUT must send both.
	UploadURL string `json:"upload_url" api:"required"`
	// The filename from the request, when one was sent.
	Filename string `json:"filename" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentLength respjson.Field
		ExpiresIn     respjson.Field
		Key           respjson.Field
		MimeType      respjson.Field
		UploadURL     respjson.Field
		Filename      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssetUploadsResponse) RawJSON() string { return r.JSON.raw }
func (r *AssetUploadsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetNewParams struct {
	// The display name for this asset.
	Filename string `json:"filename" api:"required"`
	// The storage key from the upload request. The file must already be uploaded.
	Key string `json:"key" api:"required"`
	paramObj
}

func (r AssetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AssetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetListParams struct {
	// The cursor to fetch entries after.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// The cursor to fetch entries before.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// The number of entries to fetch per-page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssetListParams]'s query parameters as `url.Values`.
func (r AssetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AssetUploadsParams struct {
	// The file size in bytes. Maximum is 5 MB. The upload URL is signed for this
	// length.
	ContentLength int64 `json:"content_length" api:"required"`
	// The MIME type of the file. One of image/jpeg, image/png, or image/gif.
	MimeType string `json:"mime_type" api:"required"`
	// The display name to echo back. The upload request does not store it. Send it
	// again when creating the asset.
	Filename param.Opt[string] `json:"filename,omitzero"`
	paramObj
}

func (r AssetUploadsParams) MarshalJSON() (data []byte, err error) {
	type shadow AssetUploadsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssetUploadsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
