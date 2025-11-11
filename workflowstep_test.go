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

func TestWorkflowStepPreviewTemplateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Steps.PreviewTemplate(
		context.TODO(),
		"step_ref",
		knockmapi.WorkflowStepPreviewTemplateParams{
			WorkflowKey: "workflow_key",
			Environment: "development",
			Recipient: knockmapi.WorkflowStepPreviewTemplateParamsRecipientUnion{
				OfString: knockmapi.String("dnedry"),
			},
			Branch: knockmapi.String("feature-branch"),
			Actor: knockmapi.WorkflowStepPreviewTemplateParamsActorUnion{
				OfString: knockmapi.String("dnedry"),
			},
			Data: map[string]any{
				"park_id": "bar",
			},
			Tenant: knockmapi.String("acme-corp"),
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
