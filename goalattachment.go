// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/knocklabs/knock-mgmt-go/internal/apijson"
	"github.com/knocklabs/knock-mgmt-go/internal/apiquery"
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// Goals define event conditions that are tracked and attributed to messaging
// resources.
//
// GoalAttachmentService contains methods and other services that help with
// interacting with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGoalAttachmentService] method instead.
type GoalAttachmentService struct {
	Options []option.RequestOption
}

// NewGoalAttachmentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGoalAttachmentService(opts ...option.RequestOption) (r GoalAttachmentService) {
	r = GoalAttachmentService{}
	r.Options = opts
	return
}

// Attaches the goal to a workflow, guide, or broadcast in the given environment,
// or updates its attribution window. The attachment goes live without a commit:
// the resource is published in the background, so the response doesn't wait on it.
// Workflows and guides must have been committed in the environment, and broadcasts
// must not have been sent. A resource holds one goal, so this replaces any other
// goal attached to it.
func (r *GoalAttachmentService) Update(ctx context.Context, resourceKey string, params GoalAttachmentUpdateParams, opts ...option.RequestOption) (res *GoalAttachmentUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.GoalKey == "" {
		err = errors.New("missing required goal_key parameter")
		return nil, err
	}
	if resourceKey == "" {
		err = errors.New("missing required resource_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/goals/%s/attachments/%v/%s", params.GoalKey, params.ResourceType, resourceKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns the workflows, guides, and broadcasts the goal is attached to in the
// given environment.
func (r *GoalAttachmentService) List(ctx context.Context, goalKey string, query GoalAttachmentListParams, opts ...option.RequestOption) (res *GoalAttachmentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if goalKey == "" {
		err = errors.New("missing required goal_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/goals/%s/attachments", goalKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Detaches the goal from a workflow, guide, or broadcast in the given environment.
// The change goes live without a commit: the resource is published in the
// background, so the response doesn't wait on it. Returns 404 when the goal isn't
// attached to the resource.
func (r *GoalAttachmentService) Delete(ctx context.Context, resourceKey string, params GoalAttachmentDeleteParams, opts ...option.RequestOption) (res *GoalAttachmentDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.GoalKey == "" {
		err = errors.New("missing required goal_key parameter")
		return nil, err
	}
	if resourceKey == "" {
		err = errors.New("missing required resource_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/goals/%s/attachments/%v/%s", params.GoalKey, params.ResourceType, resourceKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// A goal attached to a workflow, guide, or broadcast in one environment. An
// attachment goes live without a commit.
type GoalResourceAttachment struct {
	// The number of days after a message is sent in which a goal completion is
	// attributed to it.
	AttributionWindowDays int64 `json:"attribution_window_days" api:"required"`
	// The key of the attached goal.
	GoalKey string `json:"goal_key" api:"required"`
	// The key of the workflow, guide, or broadcast.
	ResourceKey string `json:"resource_key" api:"required"`
	// The type of resource the goal is attached to.
	//
	// Any of "workflow", "guide", "broadcast".
	ResourceType GoalResourceAttachmentResourceType `json:"resource_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttributionWindowDays respjson.Field
		GoalKey               respjson.Field
		ResourceKey           respjson.Field
		ResourceType          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalResourceAttachment) RawJSON() string { return r.JSON.raw }
func (r *GoalResourceAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of resource the goal is attached to.
type GoalResourceAttachmentResourceType string

const (
	GoalResourceAttachmentResourceTypeWorkflow  GoalResourceAttachmentResourceType = "workflow"
	GoalResourceAttachmentResourceTypeGuide     GoalResourceAttachmentResourceType = "guide"
	GoalResourceAttachmentResourceTypeBroadcast GoalResourceAttachmentResourceType = "broadcast"
)

// Wraps the GoalResourceAttachment response under the `attachment` key.
type GoalAttachmentUpdateResponse struct {
	// A goal attached to a workflow, guide, or broadcast in one environment. An
	// attachment goes live without a commit.
	Attachment GoalResourceAttachment `json:"attachment" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attachment  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalAttachmentUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *GoalAttachmentUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The goal's attachments in the environment.
type GoalAttachmentListResponse struct {
	// The goal's attachments in the environment.
	Attachments []GoalResourceAttachment `json:"attachments" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attachments respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalAttachmentListResponse) RawJSON() string { return r.JSON.raw }
func (r *GoalAttachmentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response from detaching a goal.
type GoalAttachmentDeleteResponse struct {
	// The result of the detach operation.
	Result string `json:"result" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GoalAttachmentDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *GoalAttachmentDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalAttachmentUpdateParams struct {
	GoalKey string `path:"goal_key" api:"required" json:"-"`
	// Any of "workflow", "guide", "broadcast".
	ResourceType GoalAttachmentUpdateParamsResourceType `path:"resource_type,omitzero" api:"required" json:"-"`
	// Settings for the goal attachment.
	Attachment GoalAttachmentUpdateParamsAttachment `json:"attachment,omitzero" api:"required"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

func (r GoalAttachmentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow GoalAttachmentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoalAttachmentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [GoalAttachmentUpdateParams]'s query parameters as
// `url.Values`.
func (r GoalAttachmentUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GoalAttachmentUpdateParamsResourceType string

const (
	GoalAttachmentUpdateParamsResourceTypeWorkflow  GoalAttachmentUpdateParamsResourceType = "workflow"
	GoalAttachmentUpdateParamsResourceTypeGuide     GoalAttachmentUpdateParamsResourceType = "guide"
	GoalAttachmentUpdateParamsResourceTypeBroadcast GoalAttachmentUpdateParamsResourceType = "broadcast"
)

// Settings for the goal attachment.
type GoalAttachmentUpdateParamsAttachment struct {
	// The number of days after a message is sent in which a goal completion is
	// attributed to it. Must be between 1 and 30. When omitted, re-saving the attached
	// goal keeps its window; otherwise 7.
	AttributionWindowDays param.Opt[int64] `json:"attribution_window_days,omitzero"`
	paramObj
}

func (r GoalAttachmentUpdateParamsAttachment) MarshalJSON() (data []byte, err error) {
	type shadow GoalAttachmentUpdateParamsAttachment
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GoalAttachmentUpdateParamsAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GoalAttachmentListParams struct {
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GoalAttachmentListParams]'s query parameters as
// `url.Values`.
func (r GoalAttachmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GoalAttachmentDeleteParams struct {
	GoalKey string `path:"goal_key" api:"required" json:"-"`
	// Any of "workflow", "guide", "broadcast".
	ResourceType GoalAttachmentDeleteParamsResourceType `path:"resource_type,omitzero" api:"required" json:"-"`
	// The environment slug. When omitted, the account's default environment is used.
	Environment param.Opt[string] `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GoalAttachmentDeleteParams]'s query parameters as
// `url.Values`.
func (r GoalAttachmentDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GoalAttachmentDeleteParamsResourceType string

const (
	GoalAttachmentDeleteParamsResourceTypeWorkflow  GoalAttachmentDeleteParamsResourceType = "workflow"
	GoalAttachmentDeleteParamsResourceTypeGuide     GoalAttachmentDeleteParamsResourceType = "guide"
	GoalAttachmentDeleteParamsResourceTypeBroadcast GoalAttachmentDeleteParamsResourceType = "broadcast"
)
