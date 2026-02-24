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
)

func TestPartialGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Partials.Get(
		context.TODO(),
		"partial_key",
		knockmapi.PartialGetParams{
			Environment:            "development",
			Annotate:               knockmapi.Bool(true),
			Branch:                 knockmapi.String("feature-branch"),
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

func TestPartialListWithOptionalParams(t *testing.T) {
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
	_, err := client.Partials.List(context.TODO(), knockmapi.PartialListParams{
		Environment:            "development",
		After:                  knockmapi.String("after"),
		Annotate:               knockmapi.Bool(true),
		Before:                 knockmapi.String("before"),
		Branch:                 knockmapi.String("feature-branch"),
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

func TestPartialUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Partials.Upsert(
		context.TODO(),
		"partial_key",
		knockmapi.PartialUpsertParams{
			Environment: "development",
			Partial: knockmapi.PartialUpsertParamsPartial{
				Content:     "<p>Hello, world!</p>",
				Name:        "My Partial",
				Type:        "html",
				Description: knockmapi.String("This is a test partial"),
				IconName:    knockmapi.String("icon_name"),
				InputSchema: []knockmapi.PartialUpsertParamsPartialInputSchemaUnion{{
					OfMessageTypeTextField: &knockmapi.MessageTypeTextFieldParam{
						Key:   "text_field",
						Label: knockmapi.String("My text field"),
						Type:  knockmapi.MessageTypeTextFieldTypeText,
						Settings: knockmapi.MessageTypeTextFieldSettingsParam{
							Default:     knockmapi.String("A placeholder"),
							Description: knockmapi.String("A description of the text field"),
							MaxLength:   knockmapi.Int(100),
							MinLength:   knockmapi.Int(10),
							Placeholder: knockmapi.String("A placeholder for the field."),
							Required:    knockmapi.Bool(true),
						},
					},
				}},
				VisualBlockEnabled: knockmapi.Bool(true),
			},
			Annotate:      knockmapi.Bool(true),
			Branch:        knockmapi.String("feature-branch"),
			Commit:        knockmapi.Bool(true),
			CommitMessage: knockmapi.String("commit_message"),
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

func TestPartialValidateWithOptionalParams(t *testing.T) {
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
	_, err := client.Partials.Validate(
		context.TODO(),
		"partial_key",
		knockmapi.PartialValidateParams{
			Environment: "development",
			Partial: knockmapi.PartialValidateParamsPartial{
				Content:     "<p>Hello, world!</p>",
				Name:        "My Partial",
				Type:        "html",
				Description: knockmapi.String("This is a test partial"),
				IconName:    knockmapi.String("icon_name"),
				InputSchema: []knockmapi.PartialValidateParamsPartialInputSchemaUnion{{
					OfMessageTypeTextField: &knockmapi.MessageTypeTextFieldParam{
						Key:   "text_field",
						Label: knockmapi.String("My text field"),
						Type:  knockmapi.MessageTypeTextFieldTypeText,
						Settings: knockmapi.MessageTypeTextFieldSettingsParam{
							Default:     knockmapi.String("A placeholder"),
							Description: knockmapi.String("A description of the text field"),
							MaxLength:   knockmapi.Int(100),
							MinLength:   knockmapi.Int(10),
							Placeholder: knockmapi.String("A placeholder for the field."),
							Required:    knockmapi.Bool(true),
						},
					},
				}},
				VisualBlockEnabled: knockmapi.Bool(true),
			},
			Branch: knockmapi.String("feature-branch"),
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
