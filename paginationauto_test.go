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

func TestAutoPagination(t *testing.T) {
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
	iter := client.Workflows.ListAutoPaging(context.TODO(), knockmapi.WorkflowListParams{
		Environment: "development",
	})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		workflow := iter.Current()
		t.Logf("%+v\n", workflow.Valid)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
