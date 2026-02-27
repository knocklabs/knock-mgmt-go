// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the knock mgmt API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options   []option.RequestOption
	Templates TemplateService
	// Email layouts wrap your email templates and provide a consistent look and feel.
	EmailLayouts EmailLayoutService
	// Commits are versioned changes to resources.
	Commits CommitService
	// Partials allow you to reuse content across templates.
	Partials PartialService
	// Translations are per-locale string files that can be used in your templates.
	Translations TranslationService
	// Workflows let you express your cross-channel notification logic.
	Workflows WorkflowService
	// A message type allows you to specify an in-app schema that defines the fields
	// available for your in-app notifications.
	MessageTypes MessageTypeService
	// Resources for managing your Knock account.
	Auth          AuthService
	APIKeys       APIKeyService
	ChannelGroups ChannelGroupService
	Channels      ChannelService
	Members       MemberService
	// Environments are isolated instances of your account that map to your
	// infrastructure.
	Environments EnvironmentService
	Variables    VariableService
	// Guides let you define in-app guides that can be displayed to users based on
	// priority and other conditions.
	Guides GuideService
	// Branches in Knock are a way to isolate changes to your Knock resources.
	Branches   BranchService
	Broadcasts BroadcastService
}

// DefaultClientOptions read from the environment (KNOCK_SERVICE_TOKEN,
// KNOCK_MGMT_BASE_URL). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("KNOCK_MGMT_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("KNOCK_SERVICE_TOKEN"); ok {
		defaults = append(defaults, option.WithServiceToken(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (KNOCK_SERVICE_TOKEN, KNOCK_MGMT_BASE_URL). The option passed in as
// arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.Templates = NewTemplateService(opts...)
	r.EmailLayouts = NewEmailLayoutService(opts...)
	r.Commits = NewCommitService(opts...)
	r.Partials = NewPartialService(opts...)
	r.Translations = NewTranslationService(opts...)
	r.Workflows = NewWorkflowService(opts...)
	r.MessageTypes = NewMessageTypeService(opts...)
	r.Auth = NewAuthService(opts...)
	r.APIKeys = NewAPIKeyService(opts...)
	r.ChannelGroups = NewChannelGroupService(opts...)
	r.Channels = NewChannelService(opts...)
	r.Members = NewMemberService(opts...)
	r.Environments = NewEnvironmentService(opts...)
	r.Variables = NewVariableService(opts...)
	r.Guides = NewGuideService(opts...)
	r.Branches = NewBranchService(opts...)
	r.Broadcasts = NewBroadcastService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
