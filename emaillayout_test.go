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

func TestEmailLayoutGetWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailLayouts.Get(
		context.TODO(),
		"email_layout_key",
		knockmapi.EmailLayoutGetParams{
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

func TestEmailLayoutListWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailLayouts.List(context.TODO(), knockmapi.EmailLayoutListParams{
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

func TestEmailLayoutUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailLayouts.Upsert(
		context.TODO(),
		"email_layout_key",
		knockmapi.EmailLayoutUpsertParams{
			Environment: "development",
			EmailLayout: knockmapi.EmailLayoutUpsertParamsEmailLayout{
				HTMLLayout: "<html><body>Hello, world!</body></html>",
				Name:       "Transactional",
				TextLayout: "Hello, world!",
				FooterLinks: []knockmapi.EmailLayoutUpsertParamsEmailLayoutFooterLink{{
					Text: "Example",
					URL:  "http://example.com",
				}},
				IsMjml: knockmapi.Bool(true),
			},
			Annotate:      knockmapi.Bool(true),
			Branch:        knockmapi.String("feature-branch"),
			Commit:        knockmapi.Bool(true),
			CommitMessage: knockmapi.String("commit_message"),
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

func TestEmailLayoutValidateWithOptionalParams(t *testing.T) {
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
	_, err := client.EmailLayouts.Validate(
		context.TODO(),
		"email_layout_key",
		knockmapi.EmailLayoutValidateParams{
			Environment: "development",
			EmailLayout: knockmapi.EmailLayoutValidateParamsEmailLayout{
				HTMLLayout: "<html><body>Hello, world!</body></html>",
				Name:       "Transactional",
				TextLayout: "Hello, world!",
				FooterLinks: []knockmapi.EmailLayoutValidateParamsEmailLayoutFooterLink{{
					Text: "Example",
					URL:  "http://example.com",
				}},
				IsMjml: knockmapi.Bool(true),
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
