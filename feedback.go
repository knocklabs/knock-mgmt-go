// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/knocklabs/knock-mgmt-go/internal/apijson"
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// Feedback lets you report bugs, friction, and feature requests to Knock.
//
// FeedbackService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeedbackService] method instead.
type FeedbackService struct {
	Options []option.RequestOption
}

// NewFeedbackService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFeedbackService(opts ...option.RequestOption) (r FeedbackService) {
	r = FeedbackService{}
	r.Options = opts
	return
}

// Sends bugs, friction, and feature requests to the Knock team. Agents can use
// this to report issues they encounter while completing a task. Do not include
// secrets or personally identifiable information.
func (r *FeedbackService) New(ctx context.Context, body FeedbackNewParams, opts ...option.RequestOption) (res *FeedbackNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/feedback"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Wraps the Feedback response under the `feedback` key.
type FeedbackNewResponse struct {
	// Feedback received by Knock.
	Feedback FeedbackNewResponseFeedback `json:"feedback" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Feedback    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeedbackNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FeedbackNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feedback received by Knock.
type FeedbackNewResponseFeedback struct {
	// The unique identifier of the feedback.
	ID string `json:"id" api:"required" format:"uuid"`
	// The timestamp of when the feedback was received.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeedbackNewResponseFeedback) RawJSON() string { return r.JSON.raw }
func (r *FeedbackNewResponseFeedback) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeedbackNewParams struct {
	// Your feedback. Must contain non-whitespace text.
	Message string `json:"message" api:"required"`
	// Relevant details about the feedback, up to 64 KiB when JSON encoded. The listed
	// keys are recommended and any other keys are accepted.
	Context FeedbackNewParamsContext `json:"context,omitzero"`
	paramObj
}

func (r FeedbackNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FeedbackNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FeedbackNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Relevant details about the feedback, up to 64 KiB when JSON encoded. The listed
// keys are recommended and any other keys are accepted.
type FeedbackNewParamsContext struct {
	// What happened, including any error text.
	Actual param.Opt[string] `json:"actual,omitzero"`
	// What you expected to happen.
	Expected param.Opt[string] `json:"expected,omitzero"`
	// The time spent, retries, or workarounds used.
	Impact param.Opt[string] `json:"impact,omitzero"`
	// The name of the model used for the task.
	Model param.Opt[string] `json:"model,omitzero"`
	// The Knock product involved, such as workflows or guides.
	Product param.Opt[string] `json:"product,omitzero"`
	// What you were trying to do.
	Task param.Opt[string] `json:"task,omitzero"`
	// The tool, command, or endpoint affected.
	Tool param.Opt[string] `json:"tool,omitzero"`
	// A category such as bug, friction, or feature.
	Type        param.Opt[string] `json:"type,omitzero"`
	ExtraFields map[string]any    `json:"-"`
	paramObj
}

func (r FeedbackNewParamsContext) MarshalJSON() (data []byte, err error) {
	type shadow FeedbackNewParamsContext
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *FeedbackNewParamsContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
