// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"encoding/json"
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

// Environments are isolated instances of your account that map to your
// infrastructure.
//
// EnvironmentDefaultService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentDefaultService] method instead.
type EnvironmentDefaultService struct {
	Options []option.RequestOption
}

// NewEnvironmentDefaultService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEnvironmentDefaultService(opts ...option.RequestOption) (r EnvironmentDefaultService) {
	r = EnvironmentDefaultService{}
	r.Options = opts
	return
}

// Creates or replaces the default settings for an environment. The change applies
// to the environment immediately, including production. It is not committed or
// promoted, and concurrent writes are last-write-wins. The `preference_set`
// replaces any existing set in full, and `null` clears it.
func (r *EnvironmentDefaultService) Update(ctx context.Context, environmentSlug string, body EnvironmentDefaultUpdateParams, opts ...option.RequestOption) (res *EnvironmentDefaults, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentSlug == "" {
		err = errors.New("missing required environment_slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/environments/%s/defaults", environmentSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns the default settings for an environment. When no defaults have been set,
// `preference_set` and both timestamps are `null`.
func (r *EnvironmentDefaultService) List(ctx context.Context, environmentSlug string, opts ...option.RequestOption) (res *EnvironmentDefaults, err error) {
	opts = slices.Concat(r.Options, opts)
	if environmentSlug == "" {
		err = errors.New("missing required environment_slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/environments/%s/defaults", environmentSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The default settings for an environment.
type EnvironmentDefaults struct {
	// The timestamp of when the defaults were first set. Null when no defaults have
	// been set.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A preference set applied to every recipient in the environment unless the
	// recipient overrides it. The `__persistence_strategy__` and `__strategy__` keys
	// are not supported, because defaults are always replaced in full.
	PreferenceSet EnvironmentDefaultsPreferenceSet `json:"preference_set" api:"required"`
	// The timestamp of when the defaults were last updated. Null when no defaults have
	// been set.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		PreferenceSet respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnvironmentDefaults) RawJSON() string { return r.JSON.raw }
func (r *EnvironmentDefaults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A preference set applied to every recipient in the environment unless the
// recipient overrides it. The `__persistence_strategy__` and `__strategy__` keys
// are not supported, because defaults are always replaced in full.
type EnvironmentDefaultsPreferenceSet struct {
	// A map of category key to a boolean or an object containing `channel_types`,
	// `channels`, and `conditions`.
	Categories any `json:"categories"`
	// A map of channel type to a boolean or an object containing `conditions`. Valid
	// channel types are `email`, `sms`, `push`, `chat`, `in_app_feed`, and `http`.
	ChannelTypes any `json:"channel_types"`
	// A map of channel ID to a boolean or an object containing `conditions`.
	Channels any `json:"channels"`
	// Whether the recipient is subscribed to commercial messages, as a boolean or a
	// settings object.
	CommercialSubscribed any `json:"commercial_subscribed"`
	// A map of workflow key to a boolean or an object containing `channel_types`,
	// `channels`, and `conditions`.
	Workflows any `json:"workflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories           respjson.Field
		ChannelTypes         respjson.Field
		Channels             respjson.Field
		CommercialSubscribed respjson.Field
		Workflows            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnvironmentDefaultsPreferenceSet) RawJSON() string { return r.JSON.raw }
func (r *EnvironmentDefaultsPreferenceSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EnvironmentDefaultsPreferenceSet to a
// EnvironmentDefaultsPreferenceSetParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EnvironmentDefaultsPreferenceSetParam.Overrides()
func (r EnvironmentDefaultsPreferenceSet) ToParam() EnvironmentDefaultsPreferenceSetParam {
	return param.Override[EnvironmentDefaultsPreferenceSetParam](json.RawMessage(r.RawJSON()))
}

// A preference set applied to every recipient in the environment unless the
// recipient overrides it. The `__persistence_strategy__` and `__strategy__` keys
// are not supported, because defaults are always replaced in full.
type EnvironmentDefaultsPreferenceSetParam struct {
	// A map of category key to a boolean or an object containing `channel_types`,
	// `channels`, and `conditions`.
	Categories any `json:"categories,omitzero"`
	// A map of channel type to a boolean or an object containing `conditions`. Valid
	// channel types are `email`, `sms`, `push`, `chat`, `in_app_feed`, and `http`.
	ChannelTypes any `json:"channel_types,omitzero"`
	// A map of channel ID to a boolean or an object containing `conditions`.
	Channels any `json:"channels,omitzero"`
	// Whether the recipient is subscribed to commercial messages, as a boolean or a
	// settings object.
	CommercialSubscribed any `json:"commercial_subscribed,omitzero"`
	// A map of workflow key to a boolean or an object containing `channel_types`,
	// `channels`, and `conditions`.
	Workflows any `json:"workflows,omitzero"`
	paramObj
}

func (r EnvironmentDefaultsPreferenceSetParam) MarshalJSON() (data []byte, err error) {
	type shadow EnvironmentDefaultsPreferenceSetParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnvironmentDefaultsPreferenceSetParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnvironmentDefaultUpdateParams struct {
	// A preference set applied to every recipient in the environment unless the
	// recipient overrides it. The `__persistence_strategy__` and `__strategy__` keys
	// are not supported, because defaults are always replaced in full.
	PreferenceSet EnvironmentDefaultsPreferenceSetParam `json:"preference_set,omitzero" api:"required"`
	paramObj
}

func (r EnvironmentDefaultUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EnvironmentDefaultUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnvironmentDefaultUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
