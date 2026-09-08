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
	"github.com/knocklabs/knock-mgmt-go/shared"
)

func TestWorkflowStepPreviewTemplateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Steps.PreviewTemplate(
		context.TODO(),
		"step_ref",
		knockmapi.WorkflowStepPreviewTemplateParams{
			WorkflowKey: "workflow_key",
			Recipient: shared.RecipientReferenceUnionParam{
				OfString: knockmapi.String("dnedry"),
			},
			Branch:      knockmapi.String("feature-branch"),
			Environment: knockmapi.String("development"),
			Actor: shared.RecipientReferenceUnionParam{
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
