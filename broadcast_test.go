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

func TestBroadcastGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Broadcasts.Get(
		context.TODO(),
		"broadcast_key",
		knockmapi.BroadcastGetParams{
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

func TestBroadcastListWithOptionalParams(t *testing.T) {
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
	_, err := client.Broadcasts.List(context.TODO(), knockmapi.BroadcastListParams{
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

func TestBroadcastCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Broadcasts.Cancel(
		context.TODO(),
		"broadcast_key",
		knockmapi.BroadcastCancelParams{
			Environment: "development",
			Branch:      knockmapi.String("feature-branch"),
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

func TestBroadcastSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Broadcasts.Send(
		context.TODO(),
		"broadcast_key",
		knockmapi.BroadcastSendParams{
			Environment: "development",
			Branch:      knockmapi.String("feature-branch"),
			SendAt:      knockmapi.Time(time.Now()),
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

func TestBroadcastUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Broadcasts.Upsert(
		context.TODO(),
		"broadcast_key",
		knockmapi.BroadcastUpsertParams{
			Environment: "development",
			Broadcast: knockmapi.BroadcastRequestParam{
				Name: "My Broadcast",
				Steps: []knockmapi.BroadcastRequestStepUnionParam{{
					OfWorkflowInAppFeedStep: &knockmapi.WorkflowInAppFeedStepParam{
						Ref: "channel_1",
						Template: knockmapi.InAppFeedTemplateParam{
							MarkdownBody: "Hello **{{ recipient.name }}**",
							ActionButtons: []knockmapi.InAppFeedTemplateActionButtonParam{{
								Action: "https://example.com",
								Label:  "Button 1",
							}},
							ActionURL: knockmapi.String("{{ vars.app_url }}"),
						},
						Type:            knockmapi.WorkflowInAppFeedStepTypeChannel,
						ChannelGroupKey: param.Null[string](),
						ChannelKey:      knockmapi.String("in-app-feed"),
						ChannelOverrides: knockmapi.InAppFeedChannelSettingsParam{
							LinkTracking: knockmapi.Bool(true),
						},
						ChannelType: knockmapi.WorkflowInAppFeedStepChannelTypeInAppFeed,
						Conditions: knockmapi.ConditionGroupUnionParam{
							OfConditionGroupAllMatch: &knockmapi.ConditionGroupConditionGroupAllMatchParam{
								All: []knockmapi.ConditionParam{{
									Operator: knockmapi.ConditionOperatorEqualTo,
									Variable: "recipient.property",
									Argument: knockmapi.String("some_property"),
								}},
							},
						},
						Description: knockmapi.String("This is a description of the channel step"),
						Name:        knockmapi.String("Channel 1"),
						SendWindows: []knockmapi.SendWindowParam{{
							Day:   knockmapi.SendWindowDayMonday,
							Type:  knockmapi.SendWindowTypeSend,
							From:  knockmapi.String("18:11:19.117Z"),
							Until: knockmapi.String("18:11:19.117Z"),
						}},
					},
				}},
				Categories:  []string{"announcement"},
				Description: knockmapi.String("A broadcast to all users"),
				ScheduledAt: knockmapi.Time(time.Now()),
				Settings: knockmapi.BroadcastRequestSettingsParam{
					IsCommercial:        knockmapi.Bool(true),
					OverridePreferences: knockmapi.Bool(false),
				},
				TargetAudienceKey: knockmapi.String("all-users"),
			},
			Annotate: knockmapi.Bool(true),
			Branch:   knockmapi.String("feature-branch"),
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

func TestBroadcastValidateWithOptionalParams(t *testing.T) {
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
	_, err := client.Broadcasts.Validate(
		context.TODO(),
		"broadcast_key",
		knockmapi.BroadcastValidateParams{
			Environment: "development",
			Broadcast: knockmapi.BroadcastRequestParam{
				Name: "My Broadcast",
				Steps: []knockmapi.BroadcastRequestStepUnionParam{{
					OfWorkflowInAppFeedStep: &knockmapi.WorkflowInAppFeedStepParam{
						Ref: "channel_1",
						Template: knockmapi.InAppFeedTemplateParam{
							MarkdownBody: "Hello **{{ recipient.name }}**",
							ActionButtons: []knockmapi.InAppFeedTemplateActionButtonParam{{
								Action: "https://example.com",
								Label:  "Button 1",
							}},
							ActionURL: knockmapi.String("{{ vars.app_url }}"),
						},
						Type:            knockmapi.WorkflowInAppFeedStepTypeChannel,
						ChannelGroupKey: param.Null[string](),
						ChannelKey:      knockmapi.String("in-app-feed"),
						ChannelOverrides: knockmapi.InAppFeedChannelSettingsParam{
							LinkTracking: knockmapi.Bool(true),
						},
						ChannelType: knockmapi.WorkflowInAppFeedStepChannelTypeInAppFeed,
						Conditions: knockmapi.ConditionGroupUnionParam{
							OfConditionGroupAllMatch: &knockmapi.ConditionGroupConditionGroupAllMatchParam{
								All: []knockmapi.ConditionParam{{
									Operator: knockmapi.ConditionOperatorEqualTo,
									Variable: "recipient.property",
									Argument: knockmapi.String("some_property"),
								}},
							},
						},
						Description: knockmapi.String("This is a description of the channel step"),
						Name:        knockmapi.String("Channel 1"),
						SendWindows: []knockmapi.SendWindowParam{{
							Day:   knockmapi.SendWindowDayMonday,
							Type:  knockmapi.SendWindowTypeSend,
							From:  knockmapi.String("18:11:19.117Z"),
							Until: knockmapi.String("18:11:19.117Z"),
						}},
					},
				}},
				Categories:  []string{"announcement"},
				Description: knockmapi.String("A broadcast to all users"),
				ScheduledAt: knockmapi.Time(time.Now()),
				Settings: knockmapi.BroadcastRequestSettingsParam{
					IsCommercial:        knockmapi.Bool(true),
					OverridePreferences: knockmapi.Bool(false),
				},
				TargetAudienceKey: knockmapi.String("all-users"),
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
