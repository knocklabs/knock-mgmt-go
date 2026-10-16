// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/knocklabs/knock-mgmt-go"
	"github.com/knocklabs/knock-mgmt-go/internal/testutil"
	"github.com/knocklabs/knock-mgmt-go/option"
)

func TestAuditLogListWithOptionalParams(t *testing.T) {
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
	_, err := client.AuditLogs.List(context.TODO(), knockmapi.AuditLogListParams{
		Action: knockmapi.AuditLogListParamsActionUnion{
			OfString: knockmapi.String("string"),
		},
		ActorID: knockmapi.AuditLogListParamsActorIDUnion{
			OfString: knockmapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
		After:       knockmapi.String("after"),
		Before:      knockmapi.String("before"),
		Environment: knockmapi.String("development"),
		Limit:       knockmapi.Int(0),
		ResourceType: knockmapi.AuditLogListParamsResourceTypeUnion{
			OfAuditLogListsResourceTypeString: knockmapi.String("account"),
		},
		Since: knockmapi.Time(time.Now()),
		Until: knockmapi.Time(time.Now()),
	})
	if err != nil {
		var apierr *knockmapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
