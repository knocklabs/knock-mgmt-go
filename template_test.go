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

func TestTemplatePreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Templates.Preview(context.TODO(), knockmapi.TemplatePreviewParams{
		Environment: "development",
		ChannelType: knockmapi.TemplatePreviewParamsChannelTypeEmail,
		Recipient: shared.RecipientReferenceUnionParam{
			OfString: knockmapi.String("user_123"),
		},
		Template: knockmapi.TemplatePreviewParamsTemplateUnion{
			OfEmailTemplate: &knockmapi.EmailTemplateParam{
				Settings: knockmapi.EmailTemplateSettingsParam{
					AttachmentKey: knockmapi.String("attachments"),
					LayoutKey:     knockmapi.String("default"),
					PreContent:    knockmapi.String("pre_content"),
				},
				Subject:  "Hello {{ recipient.name }}",
				HTMLBody: knockmapi.String("<p>Welcome!</p>"),
				IsMjml:   knockmapi.Bool(true),
				TextBody: knockmapi.String("Hello, {{ recipient.name }}! Welcome to {{ vars.app_name }} Get started here: {{ data.sign_in_url }}."),
				VisualBlocks: []knockmapi.EmailTemplateVisualBlockUnionParam{{
					OfEmailMarkdownBlock: &knockmapi.EmailTemplateVisualBlockEmailMarkdownBlockParam{
						Content: "# Hello, {{ recipient.name }}!\n\nWelcome to **{{ vars.app_name }}**. [Get started here]({{ data.sign_in_url }}).",
						Type:    "markdown",
						ID:      knockmapi.String("123e4567-e89b-12d3-a456-426614174000"),
						LayoutAttrs: knockmapi.EmailTemplateVisualBlockEmailMarkdownBlockLayoutAttrsParam{
							PaddingBottom: 8,
							PaddingLeft:   4,
							PaddingRight:  4,
							PaddingTop:    8,
						},
						Variant: "default",
						Version: knockmapi.Int(1),
					},
				}, {
					OfEmailHTMLBlock: &knockmapi.EmailTemplateVisualBlockEmailHTMLBlockParam{
						Content: "<p>Hello, {{ recipient.name }}!</p><p>Welcome to <strong>{{ vars.app_name }}</strong>. <a href='{{ data.sign_in_url }}'>Get started here</a>.</p>",
						Type:    "html",
						ID:      knockmapi.String("123e4567-e89b-12d3-a456-426614174000"),
						LayoutAttrs: knockmapi.EmailTemplateVisualBlockEmailHTMLBlockLayoutAttrsParam{
							PaddingBottom: 8,
							PaddingLeft:   4,
							PaddingRight:  4,
							PaddingTop:    8,
						},
						Version: knockmapi.Int(1),
					},
				}, {
					OfEmailDividerBlock: &knockmapi.EmailTemplateVisualBlockEmailDividerBlockParam{
						Type: "divider",
						ID:   knockmapi.String("123e4567-e89b-12d3-a456-426614174000"),
						LayoutAttrs: knockmapi.EmailTemplateVisualBlockEmailDividerBlockLayoutAttrsParam{
							PaddingBottom: 8,
							PaddingLeft:   4,
							PaddingRight:  4,
							PaddingTop:    8,
						},
						Version: knockmapi.Int(1),
					},
				}, {
					OfEmailPartialBlock: &knockmapi.EmailTemplateVisualBlockEmailPartialBlockParam{
						Attrs: map[string]any{
							"foo": "bar",
						},
						Key:  "my-partial",
						Name: "My partial",
						Type: "partial",
						ID:   knockmapi.String("123e4567-e89b-12d3-a456-426614174000"),
						LayoutAttrs: knockmapi.EmailTemplateVisualBlockEmailPartialBlockLayoutAttrsParam{
							PaddingBottom: 8,
							PaddingLeft:   4,
							PaddingRight:  4,
							PaddingTop:    8,
						},
						Version: knockmapi.Int(1),
					},
				}, {
					OfEmailImageBlock: &knockmapi.EmailTemplateVisualBlockEmailImageBlockParam{
						Type:   "image",
						URL:    "https://example.com/image.png",
						ID:     knockmapi.String("123e4567-e89b-12d3-a456-426614174000"),
						Action: knockmapi.String("action"),
						Alt:    knockmapi.String("Example image"),
						LayoutAttrs: knockmapi.EmailTemplateVisualBlockEmailImageBlockLayoutAttrsParam{
							HorizontalAlign: "center",
							PaddingBottom:   4,
							PaddingLeft:     0,
							PaddingRight:    0,
							PaddingTop:      4,
						},
						StyleAttrs: knockmapi.EmailTemplateVisualBlockEmailImageBlockStyleAttrsParam{
							Width: knockmapi.String("25%"),
						},
						Version: knockmapi.Int(1),
					},
				}, {
					OfEmailButtonSetBlock: &knockmapi.EmailTemplateVisualBlockEmailButtonSetBlockParam{
						Buttons: []knockmapi.EmailTemplateVisualBlockEmailButtonSetBlockButtonParam{{
							Action:  "https://example.com/button-action",
							Label:   "Click me",
							Variant: "solid",
							SizeAttrs: knockmapi.EmailTemplateVisualBlockEmailButtonSetBlockButtonSizeAttrsParam{
								IsFullwidth: knockmapi.Bool(false),
								Size:        "sm",
							},
							StyleAttrs: knockmapi.EmailTemplateVisualBlockEmailButtonSetBlockButtonStyleAttrsParam{
								BackgroundColor: knockmapi.String("#000000"),
								BorderColor:     knockmapi.String("#000000"),
								BorderRadius:    knockmapi.Int(6),
								BorderWidth:     knockmapi.Int(1),
								TextColor:       knockmapi.String("#FFFFFF"),
							},
						}},
						Type: "button_set",
						ID:   knockmapi.String("123e4567-e89b-12d3-a456-426614174000"),
						LayoutAttrs: knockmapi.EmailTemplateVisualBlockEmailButtonSetBlockLayoutAttrsParam{
							ColumnGap:       8,
							HorizontalAlign: "left",
							PaddingBottom:   8,
							PaddingLeft:     4,
							PaddingRight:    4,
							PaddingTop:      8,
						},
						Version: knockmapi.Int(1),
					},
				}},
			},
		},
		Branch: knockmapi.String("feature-branch"),
		Actor: shared.RecipientReferenceUnionParam{
			OfObjectRecipientReference: &shared.RecipientReferenceObjectRecipientReferenceParam{
				ID:         "project_1",
				Collection: "projects",
			},
		},
		Data: map[string]any{
			"order_id": "bar",
		},
		Layout: knockmapi.TemplatePreviewParamsLayout{
			HTMLContent: knockmapi.String("html_content"),
			Key:         knockmapi.String("key"),
			TextContent: knockmapi.String("text_content"),
		},
		Tenant: knockmapi.String("tenant"),
		Workflow: knockmapi.TemplatePreviewParamsWorkflow{
			Key:                 "key",
			Categories:          []string{"string"},
			Commercial:          knockmapi.Bool(true),
			OverridePreferences: knockmapi.Bool(true),
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
