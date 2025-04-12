// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/knock-mapi-go"
	"github.com/stainless-sdks/knock-mapi-go/internal/testutil"
	"github.com/stainless-sdks/knock-mapi-go/option"
)

func TestWorkflowStepPreviewTemplateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: currently no good way to test endpoints defining callbacks, Prism mock server will fail trying to reach the provided callback url")
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
		"workflow_key",
		"step_ref",
		knockmapi.WorkflowStepPreviewTemplateParams{
			Environment: "development",
			Recipient: knockmapi.WorkflowStepPreviewTemplateParamsRecipientUnion{
				OfString: knockmapi.String("dnedry"),
			},
			Actor: knockmapi.WorkflowStepPreviewTemplateParamsActorUnion{
				OfString: knockmapi.String("dnedry"),
			},
			Data: map[string]interface{}{
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
