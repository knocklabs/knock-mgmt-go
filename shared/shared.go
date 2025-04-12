// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-sdks/knock-mapi-go/internal/apijson"
	"github.com/stainless-sdks/knock-mapi-go/packages/param"
	"github.com/stainless-sdks/knock-mapi-go/packages/resp"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// The information about a paginated result.
type PageInfo struct {
	// The number of entries to fetch per-page.
	PageSize int64 `json:"page_size,required"`
	// The cursor to fetch entries after. Will only be present if there are more
	// entries to fetch.
	After string `json:"after,nullable"`
	// The cursor to fetch entries before. Will only be present if there are more
	// entries to fetch before the current page.
	Before string `json:"before,nullable"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		PageSize    resp.Field
		After       resp.Field
		Before      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageInfo) RawJSON() string { return r.JSON.raw }
func (r *PageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
