// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/knocklabs/knock-mgmt-go/internal/apijson"
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type EntriesCursorPageInfo struct {
	After string `json:"after"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		After       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntriesCursorPageInfo) RawJSON() string { return r.JSON.raw }
func (r *EntriesCursorPageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntriesCursor[T any] struct {
	Entries  []T                   `json:"entries"`
	PageInfo EntriesCursorPageInfo `json:"page_info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r EntriesCursor[T]) RawJSON() string { return r.JSON.raw }
func (r *EntriesCursor[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *EntriesCursor[T]) GetNextPage() (res *EntriesCursor[T], err error) {
	if len(r.Entries) == 0 {
		return nil, nil
	}
	next := r.PageInfo.After
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("after", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *EntriesCursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &EntriesCursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type EntriesCursorAutoPager[T any] struct {
	page *EntriesCursor[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewEntriesCursorAutoPager[T any](page *EntriesCursor[T], err error) *EntriesCursorAutoPager[T] {
	return &EntriesCursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *EntriesCursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Entries) == 0 {
		return false
	}
	if r.idx >= len(r.page.Entries) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Entries) == 0 {
			return false
		}
	}
	r.cur = r.page.Entries[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *EntriesCursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *EntriesCursorAutoPager[T]) Err() error {
	return r.err
}

func (r *EntriesCursorAutoPager[T]) Index() int {
	return r.run
}
