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

func TestMessageTypeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.MessageTypes.Get(
		context.TODO(),
		"email",
		knockmapi.MessageTypeGetParams{
			Environment:            "development",
			Annotate:               knockmapi.Bool(true),
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

func TestMessageTypeListWithOptionalParams(t *testing.T) {
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
	_, err := client.MessageTypes.List(context.TODO(), knockmapi.MessageTypeListParams{
		Environment:            "development",
		After:                  knockmapi.String("after"),
		Annotate:               knockmapi.Bool(true),
		Before:                 knockmapi.String("before"),
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

func TestMessageTypeUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.MessageTypes.Upsert(
		context.TODO(),
		"email",
		knockmapi.MessageTypeUpsertParams{
			Environment: "development",
			MessageType: knockmapi.MessageTypeUpsertParamsMessageType{
				Description: knockmapi.String("This is a message type"),
				Name:        "My Message Type",
				Preview:     "<div>Hello, world!</div>",
				IconName:    knockmapi.String("icon_name"),
				Semver:      knockmapi.String("1.0.0"),
				Variants: []knockmapi.MessageTypeVariantParam{{
					Fields: []knockmapi.MessageTypeVariantFieldUnionParam{{
						OfMessageTypeTextField: &knockmapi.MessageTypeTextFieldParam{
							Key:   "text_field",
							Label: knockmapi.String("My text field"),
							Type:  knockmapi.MessageTypeTextFieldTypeText,
							Settings: knockmapi.MessageTypeTextFieldSettingsParam{
								Default:     knockmapi.String("A placeholder"),
								Description: knockmapi.String("A description of the text field"),
								MaxLength:   knockmapi.Int(100),
								MinLength:   knockmapi.Int(10),
								Required:    knockmapi.Bool(true),
							},
						},
					}},
					Key:  "default",
					Name: "Default",
				}},
			},
			Annotate:      knockmapi.Bool(true),
			Commit:        knockmapi.Bool(true),
			CommitMessage: knockmapi.String("commit_message"),
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

func TestMessageTypeValidateWithOptionalParams(t *testing.T) {
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
	_, err := client.MessageTypes.Validate(
		context.TODO(),
		"email",
		knockmapi.MessageTypeValidateParams{
			Environment: "development",
			MessageType: knockmapi.MessageTypeValidateParamsMessageType{
				Description: knockmapi.String("This is a message type"),
				Name:        "My Message Type",
				Preview:     "<div>Hello, world!</div>",
				IconName:    knockmapi.String("icon_name"),
				Semver:      knockmapi.String("1.0.0"),
				Variants: []knockmapi.MessageTypeVariantParam{{
					Fields: []knockmapi.MessageTypeVariantFieldUnionParam{{
						OfMessageTypeTextField: &knockmapi.MessageTypeTextFieldParam{
							Key:   "text_field",
							Label: knockmapi.String("My text field"),
							Type:  knockmapi.MessageTypeTextFieldTypeText,
							Settings: knockmapi.MessageTypeTextFieldSettingsParam{
								Default:     knockmapi.String("A placeholder"),
								Description: knockmapi.String("A description of the text field"),
								MaxLength:   knockmapi.Int(100),
								MinLength:   knockmapi.Int(10),
								Required:    knockmapi.Bool(true),
							},
						},
					}},
					Key:  "default",
					Name: "Default",
				}},
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
