// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/knocklabs/knock-mgmt-go"
	"github.com/knocklabs/knock-mgmt-go/internal/testutil"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/param"
)

func TestGuideGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
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
	_, err := client.Guides.Get(
		context.TODO(),
		"guide_key",
		knockmapi.GuideGetParams{
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

func TestGuideListWithOptionalParams(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
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
	_, err := client.Guides.List(context.TODO(), knockmapi.GuideListParams{
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

func TestGuideActivateWithOptionalParams(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
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
	_, err := client.Guides.Activate(
		context.TODO(),
		"guide_key",
		knockmapi.GuideActivateParams{
			Environment: "development",
			Branch:      knockmapi.String("feature-branch"),
			OfGuideScheduledActivations: &knockmapi.GuideActivateParamsBodyGuideScheduledActivationParams{
				From:  knockmapi.Time(time.Now()),
				Until: knockmapi.Time(time.Now()),
			},
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

func TestGuideArchive(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
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
	_, err := client.Guides.Archive(context.TODO(), "guide_key")
	if err != nil {
		var apierr *knockmapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGuideUpsertWithOptionalParams(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
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
	_, err := client.Guides.Upsert(
		context.TODO(),
		"guide_key",
		knockmapi.GuideUpsertParams{
			Environment: "development",
			Guide: knockmapi.GuideUpsertParamsGuide{
				ChannelKey: "in-app-guide",
				Name:       "Getting Started Guide",
				Steps: []knockmapi.GuideStepParam{{
					Ref:              "welcome-step",
					SchemaKey:        "tooltip",
					SchemaSemver:     "1.0.0",
					SchemaVariantKey: "default",
					Name:             knockmapi.String("Welcome to the App"),
					Values: map[string]any{
						"text_field": "bar",
					},
				}},
				ActivationURLPatterns: []knockmapi.GuideActivationURLPatternParam{{
					Directive: knockmapi.GuideActivationURLPatternDirectiveAllow,
					Pathname:  knockmapi.String("/dashboard/*"),
					Search:    knockmapi.String("tab=settings"),
				}},
				ArchivedAt:       knockmapi.Time(time.Now()),
				DeletedAt:        knockmapi.Time(time.Now()),
				Description:      knockmapi.String("A guide to help users get started with the application"),
				TargetAudienceID: param.Null[string](),
				TargetPropertyConditions: knockmapi.ConditionGroupUnionParam{
					OfConditionGroupAllMatch: &knockmapi.ConditionGroupConditionGroupAllMatchParam{
						All: []knockmapi.ConditionParam{{
							Operator: knockmapi.ConditionOperatorEqualTo,
							Variable: "recipient.property",
							Argument: knockmapi.String("some_property"),
						}},
					},
				},
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

func TestGuideValidateWithOptionalParams(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
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
	_, err := client.Guides.Validate(
		context.TODO(),
		"guide_key",
		knockmapi.GuideValidateParams{
			Environment: "development",
			Guide: knockmapi.GuideValidateParamsGuide{
				ChannelKey: "in-app-guide",
				Name:       "Getting Started Guide",
				Steps: []knockmapi.GuideStepParam{{
					Ref:              "welcome-step",
					SchemaKey:        "tooltip",
					SchemaSemver:     "1.0.0",
					SchemaVariantKey: "default",
					Name:             knockmapi.String("Welcome to the App"),
					Values: map[string]any{
						"text_field": "bar",
					},
				}},
				ActivationURLPatterns: []knockmapi.GuideActivationURLPatternParam{{
					Directive: knockmapi.GuideActivationURLPatternDirectiveAllow,
					Pathname:  knockmapi.String("/dashboard/*"),
					Search:    knockmapi.String("tab=settings"),
				}},
				ArchivedAt:       knockmapi.Time(time.Now()),
				DeletedAt:        knockmapi.Time(time.Now()),
				Description:      knockmapi.String("A guide to help users get started with the application"),
				TargetAudienceID: param.Null[string](),
				TargetPropertyConditions: knockmapi.ConditionGroupUnionParam{
					OfConditionGroupAllMatch: &knockmapi.ConditionGroupConditionGroupAllMatchParam{
						All: []knockmapi.ConditionParam{{
							Operator: knockmapi.ConditionOperatorEqualTo,
							Variable: "recipient.property",
							Argument: knockmapi.String("some_property"),
						}},
					},
				},
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
