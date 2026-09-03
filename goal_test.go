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

func TestGoalGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Goals.Get(
		context.TODO(),
		"goal_key",
		knockmapi.GoalGetParams{
			Environment: "development",
			Annotate:    knockmapi.Bool(true),
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

func TestGoalListWithOptionalParams(t *testing.T) {
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
	_, err := client.Goals.List(context.TODO(), knockmapi.GoalListParams{
		Environment: "development",
		After:       knockmapi.String("after"),
		Annotate:    knockmapi.Bool(true),
		Before:      knockmapi.String("before"),
		Branch:      knockmapi.String("feature-branch"),
		Limit:       knockmapi.Int(0),
	})
	if err != nil {
		var apierr *knockmapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGoalArchive(t *testing.T) {
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
	_, err := client.Goals.Archive(
		context.TODO(),
		"goal_key",
		knockmapi.GoalArchiveParams{
			Environment: "development",
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

func TestGoalClone(t *testing.T) {
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
	_, err := client.Goals.Clone(
		context.TODO(),
		"goal_key",
		knockmapi.GoalCloneParams{
			Environment: "development",
			Clone: knockmapi.GoalCloneParamsClone{
				Environment: "production",
				Key:         "trial-conversion-copy",
				Name:        "Trial Conversion Copy",
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

func TestGoalUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Goals.Upsert(
		context.TODO(),
		"goal_key",
		knockmapi.GoalUpsertParams{
			Environment: "development",
			Goal: knockmapi.GoalRequestParam{
				Condition: knockmapi.GoalConditionParam{
					Event: knockmapi.GoalConditionEventUnionParam{
						OfWorkflowWaitForEventRecipientEvent: &knockmapi.GoalConditionEventWorkflowWaitForEventRecipientEventParam{
							EventType: "recipient",
							EventKey:  "updated",
						},
					},
					MatchConditions: []knockmapi.ConditionGroupUnionParam{{
						OfConditionGroupAllMatch: &knockmapi.ConditionGroupAllMatchParam{
							All: []knockmapi.ConditionParam{{
								Operator: knockmapi.ConditionOperatorEqualTo,
								Variable: "recipient.property",
								Argument: knockmapi.String("some_property"),
							}},
						},
					}},
				},
				Name:        "Trial Conversion",
				Description: knockmapi.String("Tracks when a trial user converts to paid"),
			},
			Annotate: knockmapi.Bool(true),
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

func TestGoalValidateWithOptionalParams(t *testing.T) {
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
	_, err := client.Goals.Validate(
		context.TODO(),
		"goal_key",
		knockmapi.GoalValidateParams{
			Environment: "development",
			Goal: knockmapi.GoalRequestParam{
				Condition: knockmapi.GoalConditionParam{
					Event: knockmapi.GoalConditionEventUnionParam{
						OfWorkflowWaitForEventRecipientEvent: &knockmapi.GoalConditionEventWorkflowWaitForEventRecipientEventParam{
							EventType: "recipient",
							EventKey:  "updated",
						},
					},
					MatchConditions: []knockmapi.ConditionGroupUnionParam{{
						OfConditionGroupAllMatch: &knockmapi.ConditionGroupAllMatchParam{
							All: []knockmapi.ConditionParam{{
								Operator: knockmapi.ConditionOperatorEqualTo,
								Variable: "recipient.property",
								Argument: knockmapi.String("some_property"),
							}},
						},
					}},
				},
				Name:        "Trial Conversion",
				Description: knockmapi.String("Tracks when a trial user converts to paid"),
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
