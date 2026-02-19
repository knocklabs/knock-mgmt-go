// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/knocklabs/knock-mgmt-go"
	"github.com/knocklabs/knock-mgmt-go/internal/testutil"
	"github.com/knocklabs/knock-mgmt-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Mock server doesn't support callbacks yet")
	page, err := client.Workflows.List(context.TODO(), knockmapi.WorkflowListParams{
		Environment: "development",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", page)
}
