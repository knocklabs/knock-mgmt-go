// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package knockmapi

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/knocklabs/knock-mgmt-go/internal/apijson"
	"github.com/knocklabs/knock-mgmt-go/internal/requestconfig"
	"github.com/knocklabs/knock-mgmt-go/option"
	"github.com/knocklabs/knock-mgmt-go/packages/respjson"
)

// Resources for managing your Knock account.
//
// BillingService contains methods and other services that help with interacting
// with the knock mgmt API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingService] method instead.
type BillingService struct {
	Options []option.RequestOption
}

// NewBillingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBillingService(opts ...option.RequestOption) (r BillingService) {
	r = BillingService{}
	r.Options = opts
	return
}

// Returns a snapshot of the current draft invoice for the account, including
// amount due, plan, per-metric usage, and remaining credit blocks.
func (r *BillingService) GetSummary(ctx context.Context, opts ...option.RequestOption) (res *BillingSummary, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/billing/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A snapshot of the current draft Orb invoice for the account, including plan,
// per-metric usage, and remaining credit blocks.
type BillingSummary struct {
	// The amount due on the draft invoice, as a decimal string.
	AmountDue string `json:"amount_due" api:"required"`
	// Unexpired Orb credit blocks for custom pricing units on the draft invoice,
	// including scheduled and depleted allotments.
	Credits []BillingSummaryCredit `json:"credits" api:"required"`
	// The invoice currency code, such as USD.
	Currency string `json:"currency" api:"required"`
	// The Orb draft invoice id, if present.
	InvoiceID string `json:"invoice_id" api:"required"`
	// The end of the current usage period, hoisted from the first usage line item.
	PeriodEnd time.Time `json:"period_end" api:"required" format:"date-time"`
	// The start of the current usage period, hoisted from the first usage line item.
	PeriodStart time.Time `json:"period_start" api:"required" format:"date-time"`
	// The Knock plan currently assigned to the account.
	Plan BillingSummaryPlan `json:"plan" api:"required"`
	// The date the draft invoice is scheduled to be issued.
	TargetDate time.Time `json:"target_date" api:"required" format:"date"`
	// Per-metric usage line items from the draft invoice.
	Usage []BillingSummaryUsage `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AmountDue   respjson.Field
		Credits     respjson.Field
		Currency    respjson.Field
		InvoiceID   respjson.Field
		PeriodEnd   respjson.Field
		PeriodStart respjson.Field
		Plan        respjson.Field
		TargetDate  respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingSummary) RawJSON() string { return r.JSON.raw }
func (r *BillingSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An unexpired Orb credit block, including scheduled and depleted allotments.
type BillingSummaryCredit struct {
	// The Orb credit block id.
	ID string `json:"id" api:"required"`
	// The Orb custom pricing-unit identifier, such as mnr_credits or ai_credits.
	Currency string `json:"currency" api:"required"`
	// When the block becomes effective.
	EffectiveAt time.Time `json:"effective_at" api:"required" format:"date-time"`
	// When the block expires.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// The display name for the credit currency.
	Name string `json:"name" api:"required"`
	// The block's original allocation.
	Quantity float64 `json:"quantity" api:"required"`
	// The block's remaining balance.
	Remaining float64 `json:"remaining" api:"required"`
	// Derived block status: scheduled when not yet effective, depleted when remaining
	// is zero or less, otherwise active.
	//
	// Any of "active", "depleted", "scheduled".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Currency    respjson.Field
		EffectiveAt respjson.Field
		ExpiresAt   respjson.Field
		Name        respjson.Field
		Quantity    respjson.Field
		Remaining   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingSummaryCredit) RawJSON() string { return r.JSON.raw }
func (r *BillingSummaryCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Knock plan currently assigned to the account.
type BillingSummaryPlan struct {
	// The unique identifier of the plan.
	ID string `json:"id" api:"required"`
	// The human-readable plan name.
	DisplayName string `json:"display_name" api:"required"`
	// The plan type.
	//
	// Any of "free", "starter", "growth", "enterprise".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DisplayName respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingSummaryPlan) RawJSON() string { return r.JSON.raw }
func (r *BillingSummaryPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Usage for a single Orb billable metric on the current draft invoice.
type BillingSummaryUsage struct {
	// The billed dollar amount for this metric after credits and adjustments.
	Amount string `json:"amount" api:"required"`
	// Prepaid credits applied to this metric during the current service period.
	CreditsApplied int64 `json:"credits_applied" api:"required"`
	// The Orb billable metric id.
	MetricID string `json:"metric_id" api:"required"`
	// The Orb billable metric name.
	Name string `json:"name" api:"required"`
	// The Orb price currency for this line item, used to correlate usage with credit
	// blocks.
	PricingCurrency string `json:"pricing_currency" api:"required"`
	// Orb line-item usage quantity for this period. This is not replaced by
	// credits_applied.
	Quantity int64 `json:"quantity" api:"required"`
	// Whether the GraphQL usage quantity is raw usage or credits applied. Prefer
	// quantity and credits_applied for new clients.
	//
	// Any of "usage", "credits".
	QuantityType string `json:"quantity_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount          respjson.Field
		CreditsApplied  respjson.Field
		MetricID        respjson.Field
		Name            respjson.Field
		PricingCurrency respjson.Field
		Quantity        respjson.Field
		QuantityType    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingSummaryUsage) RawJSON() string { return r.JSON.raw }
func (r *BillingSummaryUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
