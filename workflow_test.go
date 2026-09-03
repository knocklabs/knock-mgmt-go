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
	"github.com/knocklabs/knock-mgmt-go/shared"
)

func TestWorkflowGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Get(
		context.TODO(),
		"workflow_key",
		knockmapi.WorkflowGetParams{
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

func TestWorkflowListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.List(context.TODO(), knockmapi.WorkflowListParams{
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

func TestWorkflowActivateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Activate(
		context.TODO(),
		"workflow_key",
		knockmapi.WorkflowActivateParams{
			Environment: "development",
			Status:      true,
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

func TestWorkflowRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Run(
		context.TODO(),
		"workflow_key",
		knockmapi.WorkflowRunParams{
			Environment: "development",
			Recipients: []knockmapi.WorkflowRunParamsRecipientUnion{{
				OfInlineIdentifyUserRequest: &knockmapi.InlineIdentifyUserRequestParam{
					ID:    "user_1",
					Email: knockmapi.String("jane@example.com"),
					Name:  knockmapi.String("Jane Doe"),
				},
			}},
			Branch: knockmapi.String("feature-branch"),
			Actor: knockmapi.WorkflowRunParamsActorUnion{
				OfString: knockmapi.String("user_1"),
			},
			CancellationKey: knockmapi.String("cancellation_key"),
			Data: map[string]any{
				"park_id": "bar",
			},
			Tenant: knockmapi.String("tenant"),
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

func TestWorkflowUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Upsert(
		context.TODO(),
		"workflow_key",
		knockmapi.WorkflowUpsertParams{
			Environment: "development",
			Workflow: knockmapi.WorkflowRequestParam{
				Name: "My Workflow",
				Steps: []knockmapi.WorkflowStepUnionParam{{
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
							OfConditionGroupAllMatch: &knockmapi.ConditionGroupAllMatchParam{
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
							From:  knockmapi.String("09:00"),
							Until: knockmapi.String("17:00"),
						}},
					},
				}},
				Categories: []string{"string"},
				Conditions: knockmapi.ConditionGroupUnionParam{
					OfConditionGroupAllMatch: &knockmapi.ConditionGroupAllMatchParam{
						All: []knockmapi.ConditionParam{{
							Operator: knockmapi.ConditionOperatorEqualTo,
							Variable: "recipient.property",
							Argument: knockmapi.String("some_property"),
						}},
					},
				},
				Description: knockmapi.String("description"),
				GoalAttachment: shared.GoalAttachmentParam{
					GoalKey:               "trial-conversion",
					AttributionWindowDays: knockmapi.Int(7),
				},
				Settings: knockmapi.WorkflowRequestSettingsParam{
					IsCommercial:        knockmapi.Bool(false),
					OverridePreferences: knockmapi.Bool(false),
				},
				Tags: []string{"string"},
				TriggerDataJsonSchema: map[string]any{
					"foo": "bar",
				},
				TriggerFrequency: knockmapi.WorkflowRequestTriggerFrequencyEveryTrigger,
			},
			AllowEmpty:    knockmapi.Bool(true),
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

func TestWorkflowValidateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Validate(
		context.TODO(),
		"workflow_key",
		knockmapi.WorkflowValidateParams{
			Environment: "development",
			Workflow: knockmapi.WorkflowRequestParam{
				Name: "My Workflow",
				Steps: []knockmapi.WorkflowStepUnionParam{{
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
							OfConditionGroupAllMatch: &knockmapi.ConditionGroupAllMatchParam{
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
							From:  knockmapi.String("09:00"),
							Until: knockmapi.String("17:00"),
						}},
					},
				}},
				Categories: []string{"string"},
				Conditions: knockmapi.ConditionGroupUnionParam{
					OfConditionGroupAllMatch: &knockmapi.ConditionGroupAllMatchParam{
						All: []knockmapi.ConditionParam{{
							Operator: knockmapi.ConditionOperatorEqualTo,
							Variable: "recipient.property",
							Argument: knockmapi.String("some_property"),
						}},
					},
				},
				Description: knockmapi.String("description"),
				GoalAttachment: shared.GoalAttachmentParam{
					GoalKey:               "trial-conversion",
					AttributionWindowDays: knockmapi.Int(7),
				},
				Settings: knockmapi.WorkflowRequestSettingsParam{
					IsCommercial:        knockmapi.Bool(false),
					OverridePreferences: knockmapi.Bool(false),
				},
				Tags: []string{"string"},
				TriggerDataJsonSchema: map[string]any{
					"foo": "bar",
				},
				TriggerFrequency: knockmapi.WorkflowRequestTriggerFrequencyEveryTrigger,
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
