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

func TestManualPagination(t *testing.T) {
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
	page, err := client.Workflows.List(context.TODO(), knockmapi.WorkflowListParams{
		Environment: "development",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, workflow := range page.Entries {
		t.Logf("%+v\n", workflow.Valid)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, workflow := range page.Entries {
			t.Logf("%+v\n", workflow.Valid)
		}
	}
}
