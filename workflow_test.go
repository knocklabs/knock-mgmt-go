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

func TestWorkflowGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Get(
		context.TODO(),
		"workflow_key",
		knockmapi.WorkflowGetParams{
			Environment:            "development",
			Annotate:               knockmapi.Bool(true),
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
	_, err := client.Workflows.List(context.TODO(), knockmapi.WorkflowListParams{
		Environment:            "development",
		After:                  knockmapi.String("after"),
		Annotate:               knockmapi.Bool(true),
		Before:                 knockmapi.String("before"),
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

func TestWorkflowActivate(t *testing.T) {
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
	_, err := client.Workflows.Activate(
		context.TODO(),
		"workflow_key",
		knockmapi.WorkflowActivateParams{
			Environment: "development",
			Status:      true,
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
	_, err := client.Workflows.Run(
		context.TODO(),
		"workflow_key",
		knockmapi.WorkflowRunParams{
			Environment: "development",
			Recipients: []knockmapi.WorkflowRunParamsRecipientUnion{{
				OfString: knockmapi.String("dnedry"),
			}},
			Actor: knockmapi.WorkflowRunParamsActorUnion{
				OfObjectRecipientReference: &knockmapi.WorkflowRunParamsActorObjectRecipientReference{
					ID:         "project_1",
					Collection: "projects",
				},
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
	_, err := client.Workflows.Upsert(
		context.TODO(),
		"workflow_key",
		knockmapi.WorkflowUpsertParams{
			Environment: "development",
			Workflow: knockmapi.WorkflowUpsertParamsWorkflow{
				Name: "My Workflow",
				Steps: []knockmapi.WorkflowStepUnionParam{{
					OfWorkflowInAppFeedStep: &knockmapi.WorkflowInAppFeedStepParam{
						Name: "Channel 1",
						Ref:  "channel_1",
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
						SendWindows: []knockmapi.SendWindowParam{{
							Day:   knockmapi.SendWindowDayMonday,
							Type:  knockmapi.SendWindowTypeSend,
							From:  knockmapi.Time("18:11:19.117Z"),
							Until: knockmapi.Time("18:11:19.117Z"),
						}},
					},
				}},
				Categories: []string{"string"},
				Conditions: knockmapi.ConditionGroupUnionParam{
					OfConditionGroupAllMatch: &knockmapi.ConditionGroupConditionGroupAllMatchParam{
						All: []knockmapi.ConditionParam{{
							Operator: knockmapi.ConditionOperatorEqualTo,
							Variable: "recipient.property",
							Argument: knockmapi.String("some_property"),
						}},
					},
				},
				Description: knockmapi.String("description"),
				Settings: knockmapi.WorkflowUpsertParamsWorkflowSettings{
					IsCommercial:        knockmapi.Bool(false),
					OverridePreferences: knockmapi.Bool(false),
				},
				TriggerDataJsonSchema: map[string]any{
					"foo": "bar",
				},
				TriggerFrequency: "every_trigger",
			},
			Annotate:      knockmapi.Bool(true),
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

func TestWorkflowValidateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Validate(
		context.TODO(),
		"workflow_key",
		knockmapi.WorkflowValidateParams{
			Environment: "development",
			Workflow: knockmapi.WorkflowValidateParamsWorkflow{
				Name: "My Workflow",
				Steps: []knockmapi.WorkflowStepUnionParam{{
					OfWorkflowInAppFeedStep: &knockmapi.WorkflowInAppFeedStepParam{
						Name: "Channel 1",
						Ref:  "channel_1",
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
						SendWindows: []knockmapi.SendWindowParam{{
							Day:   knockmapi.SendWindowDayMonday,
							Type:  knockmapi.SendWindowTypeSend,
							From:  knockmapi.Time("18:11:19.117Z"),
							Until: knockmapi.Time("18:11:19.117Z"),
						}},
					},
				}},
				Categories: []string{"string"},
				Conditions: knockmapi.ConditionGroupUnionParam{
					OfConditionGroupAllMatch: &knockmapi.ConditionGroupConditionGroupAllMatchParam{
						All: []knockmapi.ConditionParam{{
							Operator: knockmapi.ConditionOperatorEqualTo,
							Variable: "recipient.property",
							Argument: knockmapi.String("some_property"),
						}},
					},
				},
				Description: knockmapi.String("description"),
				Settings: knockmapi.WorkflowValidateParamsWorkflowSettings{
					IsCommercial:        knockmapi.Bool(false),
					OverridePreferences: knockmapi.Bool(false),
				},
				TriggerDataJsonSchema: map[string]any{
					"foo": "bar",
				},
				TriggerFrequency: "every_trigger",
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
