// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/knocklabs/knock-mgmt-go"
	"github.com/knocklabs/knock-mgmt-go/internal/testutil"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
)

func TestReusableRequestGetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knockmapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithServiceToken("My Service Token"),
	)
	_, err := client.ReusableRequests.Get(
		context.TODO(),
		"fetch-user",
		knockmapi.ReusableRequestGetParams{
			Annotate:               knockmapi.Bool(true),
			Branch:                 knockmapi.String("feature-branch"),
			Environment:            knockmapi.String("development"),
			HideUncommittedChanges: knockmapi.Bool(true),
		},
	)
	if err != nil {
		var apierr *knockmapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReusableRequestListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knockmapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithServiceToken("My Service Token"),
	)
	_, err := client.ReusableRequests.List(context.TODO(), knockmapi.ReusableRequestListParams{
		After:                  knockmapi.String("after"),
		Annotate:               knockmapi.Bool(true),
		Before:                 knockmapi.String("before"),
		Branch:                 knockmapi.String("feature-branch"),
		Environment:            knockmapi.String("development"),
		HideUncommittedChanges: knockmapi.Bool(true),
		Limit:                  knockmapi.Int(0),
	})
	if err != nil {
		var apierr *knockmapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReusableRequestArchiveWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knockmapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithServiceToken("My Service Token"),
	)
	_, err := client.ReusableRequests.Archive(
		context.TODO(),
		"fetch-user",
		knockmapi.ReusableRequestArchiveParams{
			Environment: knockmapi.String("development"),
		},
	)
	if err != nil {
		var apierr *knockmapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReusableRequestUpdateUnarchiveWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knockmapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithServiceToken("My Service Token"),
	)
	_, err := client.ReusableRequests.UpdateUnarchive(
		context.TODO(),
		"fetch-user",
		knockmapi.ReusableRequestUpdateUnarchiveParams{
			Environment: knockmapi.String("development"),
		},
	)
	if err != nil {
		var apierr *knockmapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReusableRequestUpsertWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knockmapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithServiceToken("My Service Token"),
	)
	_, err := client.ReusableRequests.Upsert(
		context.TODO(),
		"fetch-user",
		knockmapi.ReusableRequestUpsertParams{
			ReusableRequest: knockmapi.ReusableRequestRequestParam{
				Settings: knockmapi.ReusableRequestSettingsParam{
					Method: knockmapi.ReusableRequestSettingsMethodGet,
					Body:   param.Null[string](),
					DataInputSchema: map[string]any{
						"foo": "bar",
					},
					DataInputValuesStr: knockmapi.String("data_input_values_str"),
					Headers: knockmapi.ReusableRequestSettingsHeadersUnionParam{
						OfRequestTemplateHeadersArray: []knockmapi.ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItemParam{{
							Key:   "X-API-Key",
							Value: "1234567890",
						}},
					},
					QueryParams: knockmapi.ReusableRequestSettingsQueryParamsUnionParam{
						OfRequestTemplateQuerysArray: []knockmapi.ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItemParam{{
							Key:   "key",
							Value: "value",
						}},
					},
					ResponsePath: param.Null[string](),
					SigningKey:   param.Null[string](),
					URL:          knockmapi.String("https://example.com"),
				},
				Type: knockmapi.ReusableRequestRequestTypeHTTPFetch,
				Conditions: knockmapi.ConditionGroupUnionParam{
					OfConditionGroupAllMatch: &knockmapi.ConditionGroupAllMatchParam{
						All: []knockmapi.ConditionParam{{
							Operator: knockmapi.ConditionOperatorEqualTo,
							Variable: "recipient.property",
							Argument: knockmapi.String("some_property"),
						}},
					},
				},
				Description: knockmapi.String("Loads a user from the API."),
				Name:        knockmapi.String("Fetch user"),
			},
			AllowEmpty:    knockmapi.Bool(true),
			Annotate:      knockmapi.Bool(true),
			Branch:        knockmapi.String("feature-branch"),
			Commit:        knockmapi.Bool(true),
			CommitMessage: knockmapi.String("commit_message"),
			Environment:   knockmapi.String("development"),
			Force:         knockmapi.Bool(true),
		},
	)
	if err != nil {
		var apierr *knockmapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReusableRequestValidateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := knockmapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithServiceToken("My Service Token"),
	)
	_, err := client.ReusableRequests.Validate(
		context.TODO(),
		"fetch-user",
		knockmapi.ReusableRequestValidateParams{
			ReusableRequest: knockmapi.ReusableRequestRequestParam{
				Settings: knockmapi.ReusableRequestSettingsParam{
					Method: knockmapi.ReusableRequestSettingsMethodGet,
					Body:   param.Null[string](),
					DataInputSchema: map[string]any{
						"foo": "bar",
					},
					DataInputValuesStr: knockmapi.String("data_input_values_str"),
					Headers: knockmapi.ReusableRequestSettingsHeadersUnionParam{
						OfRequestTemplateHeadersArray: []knockmapi.ReusableRequestSettingsHeadersRequestTemplateHeadersArrayItemParam{{
							Key:   "X-API-Key",
							Value: "1234567890",
						}},
					},
					QueryParams: knockmapi.ReusableRequestSettingsQueryParamsUnionParam{
						OfRequestTemplateQuerysArray: []knockmapi.ReusableRequestSettingsQueryParamsRequestTemplateQueryParamsArrayItemParam{{
							Key:   "key",
							Value: "value",
						}},
					},
					ResponsePath: param.Null[string](),
					SigningKey:   param.Null[string](),
					URL:          knockmapi.String("https://example.com"),
				},
				Type: knockmapi.ReusableRequestRequestTypeHTTPFetch,
				Conditions: knockmapi.ConditionGroupUnionParam{
					OfConditionGroupAllMatch: &knockmapi.ConditionGroupAllMatchParam{
						All: []knockmapi.ConditionParam{{
							Operator: knockmapi.ConditionOperatorEqualTo,
							Variable: "recipient.property",
							Argument: knockmapi.String("some_property"),
						}},
					},
				},
				Description: knockmapi.String("Loads a user from the API."),
				Name:        knockmapi.String("Fetch user"),
			},
			Branch:      knockmapi.String("feature-branch"),
			Environment: knockmapi.String("development"),
		},
	)
	if err != nil {
		var apierr *knockmapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
