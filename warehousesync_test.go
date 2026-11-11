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

func TestWarehouseSyncGetWithOptionalParams(t *testing.T) {
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
	_, err := client.WarehouseSyncs.Get(
		context.TODO(),
		"warehouse_sync_key",
		knockmapi.WarehouseSyncGetParams{
			Environment:            knockmapi.String("development"),
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

func TestWarehouseSyncListWithOptionalParams(t *testing.T) {
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
	_, err := client.WarehouseSyncs.List(context.TODO(), knockmapi.WarehouseSyncListParams{
		After:                  knockmapi.String("after"),
		Before:                 knockmapi.String("before"),
		Environment:            knockmapi.String("development"),
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

func TestWarehouseSyncListRunsWithOptionalParams(t *testing.T) {
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
	_, err := client.WarehouseSyncs.ListRuns(
		context.TODO(),
		"warehouse_sync_key",
		knockmapi.WarehouseSyncListRunsParams{
			Environment: knockmapi.String("development"),
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

func TestWarehouseSyncRunWithOptionalParams(t *testing.T) {
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
	_, err := client.WarehouseSyncs.Run(
		context.TODO(),
		"warehouse_sync_key",
		knockmapi.WarehouseSyncRunParams{
			Environment: knockmapi.String("development"),
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

func TestWarehouseSyncUpdateSettingsWithOptionalParams(t *testing.T) {
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
	_, err := client.WarehouseSyncs.UpdateSettings(
		context.TODO(),
		"warehouse_sync_key",
		knockmapi.WarehouseSyncUpdateSettingsParams{
			Environment:      knockmapi.String("development"),
			Enabled:          knockmapi.Bool(false),
			FrequencyMinutes: 1440,
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

func TestWarehouseSyncUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.WarehouseSyncs.Upsert(
		context.TODO(),
		"warehouse_sync_key",
		knockmapi.WarehouseSyncUpsertParams{
			WarehouseSync: knockmapi.WarehouseSyncRequestParam{
				ChangeTracking: knockmapi.WarehouseSyncChangeTrackingParam{
					IsDeletedColumn:    knockmapi.String("is_deleted"),
					LastModifiedColumn: knockmapi.String("updated_at"),
				},
				Description: knockmapi.String("description"),
				Mappings: []knockmapi.WarehouseSyncMappingParam{{
					SourceColumn: "id",
					TargetField:  "id",
					OnNull:       knockmapi.WarehouseSyncMappingOnNullSetNull,
				}, {
					SourceColumn: "email",
					TargetField:  "email",
					OnNull:       knockmapi.WarehouseSyncMappingOnNullSetNull,
				}},
				Name:             knockmapi.String("Users"),
				PrimaryKeyColumn: knockmapi.String("id"),
				Source: knockmapi.WarehouseSyncSourceParam{
					Mode:   knockmapi.WarehouseSyncSourceModeSql,
					Name:   knockmapi.String("name"),
					Query:  knockmapi.String("SELECT id, email, name, updated_at FROM analytics.users"),
					Schema: knockmapi.String("schema"),
				},
				Target: knockmapi.WarehouseSyncTargetParam{
					AudienceKey: knockmapi.String("audience_key"),
					Collection:  knockmapi.String("projects"),
				},
				TargetType:   knockmapi.WarehouseSyncRequestTargetTypeUser,
				VersionID:    knockmapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				WarehouseKey: knockmapi.String("analytics"),
			},
			Commit:        knockmapi.Bool(true),
			CommitMessage: knockmapi.String("commit_message"),
			Environment:   knockmapi.String("development"),
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

func TestWarehouseSyncValidateWithOptionalParams(t *testing.T) {
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
	_, err := client.WarehouseSyncs.Validate(
		context.TODO(),
		"warehouse_sync_key",
		knockmapi.WarehouseSyncValidateParams{
			WarehouseSync: knockmapi.WarehouseSyncRequestParam{
				ChangeTracking: knockmapi.WarehouseSyncChangeTrackingParam{
					IsDeletedColumn:    knockmapi.String("is_deleted"),
					LastModifiedColumn: knockmapi.String("updated_at"),
				},
				Description: knockmapi.String("description"),
				Mappings: []knockmapi.WarehouseSyncMappingParam{{
					SourceColumn: "id",
					TargetField:  "id",
					OnNull:       knockmapi.WarehouseSyncMappingOnNullSetNull,
				}, {
					SourceColumn: "email",
					TargetField:  "email",
					OnNull:       knockmapi.WarehouseSyncMappingOnNullSetNull,
				}},
				Name:             knockmapi.String("Users"),
				PrimaryKeyColumn: knockmapi.String("id"),
				Source: knockmapi.WarehouseSyncSourceParam{
					Mode:   knockmapi.WarehouseSyncSourceModeSql,
					Name:   knockmapi.String("name"),
					Query:  knockmapi.String("SELECT id, email, name, updated_at FROM analytics.users"),
					Schema: knockmapi.String("schema"),
				},
				Target: knockmapi.WarehouseSyncTargetParam{
					AudienceKey: knockmapi.String("audience_key"),
					Collection:  knockmapi.String("projects"),
				},
				TargetType:   knockmapi.WarehouseSyncRequestTargetTypeUser,
				VersionID:    knockmapi.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				WarehouseKey: knockmapi.String("analytics"),
			},
			Environment: knockmapi.String("development"),
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
