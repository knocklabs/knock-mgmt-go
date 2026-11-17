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

func TestFeedbackNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Feedback.New(context.TODO(), knockmapi.FeedbackNewParams{
		Message: "Please name the missing variable in workflow validation errors.",
		Context: knockmapi.FeedbackNewParamsContext{
			Actual:   knockmapi.String("actual"),
			Expected: knockmapi.String("An error that identifies the missing variable."),
			Impact:   knockmapi.String("impact"),
			Model:    knockmapi.String("claude-opus-5-5"),
			Product:  knockmapi.String("product"),
			Task:     knockmapi.String("task"),
			Tool:     knockmapi.String("upsertWorkflow"),
			Type:     knockmapi.String("friction"),
		},
	})
	if err != nil {
		var apierr *knockmapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
