// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/devdraftengineer/devdraft-go/internal/apijson"
	"github.com/devdraftengineer/devdraft-go/internal/requestconfig"
	"github.com/devdraftengineer/devdraft-go/option"
	"github.com/devdraftengineer/devdraft-go/packages/param"
	"github.com/devdraftengineer/devdraft-go/packages/respjson"
)

// V0TestPaymentService contains methods and other services that help with
// interacting with the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0TestPaymentService] method instead.
type V0TestPaymentService struct {
	Options []option.RequestOption
}

// NewV0TestPaymentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV0TestPaymentService(opts ...option.RequestOption) (r V0TestPaymentService) {
	r = V0TestPaymentService{}
	r.Options = opts
	return
}

// Get payment details by ID
func (r *V0TestPaymentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PaymentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/test-payment/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Creates a new payment. Requires an idempotency key to prevent duplicate payments
// on retry.
//
// ## Idempotency Key Best Practices
//
//  1. **Generate unique keys**: Use UUIDs or similar unique identifiers, prefixed
//     with a descriptive operation name
//  2. **Store keys client-side**: Save the key with the original request so you can
//     retry with the same key
//  3. **Key format**: Between 6-64 alphanumeric characters
//  4. **Expiration**: Keys expire after 24 hours by default
//  5. **Use case**: Perfect for ensuring payment operations are never processed
//     twice, even during network failures
//
// ## Example Request (curl)
//
// ```bash
//
//	curl -X POST \
//	  https://api.example.com/rest-api/v0/test-payment \
//	  -H 'Content-Type: application/json' \
//	  -H 'Client-Key: your-api-key' \
//	  -H 'Client-Secret: your-api-secret' \
//	  -H 'Idempotency-Key: payment_123456_unique_key' \
//	  -d '{
//	    "amount": 100.00,
//	    "currency": "USD",
//	    "description": "Test payment for order #12345",
//	    "customerId": "cus_12345"
//	  }'
//
// ```
//
// ## Example Response (First Request)
//
// ```json
//
//	{
//	  "id": "pay_abc123xyz456",
//	  "amount": 100.0,
//	  "currency": "USD",
//	  "status": "succeeded",
//	  "timestamp": "2023-07-01T12:00:00.000Z"
//	}
//
// ```
//
// ## Example Response (Duplicate Request)
//
// The exact same response will be returned for any duplicate request with the same
// idempotency key, without creating a new payment.
//
// ## Retry Scenario Example
//
// Network failure during payment submission:
//
//  1. Client creates payment request with idempotency key:
//     "payment_123456_unique_key"
//  2. Request begins processing, but network connection fails before response
//     received
//  3. Client retries the exact same request with the same idempotency key
//  4. Server detects duplicate idempotency key and returns the result of the first
//     request
//  5. No duplicate payment is created
//
// If you retry with same key but different parameters (e.g., different amount),
// you'll receive a 409 Conflict error.
func (r *V0TestPaymentService) Process(ctx context.Context, body V0TestPaymentProcessParams, opts ...option.RequestOption) (res *PaymentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v0/test-payment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Refunds an existing payment. Requires an idempotency key to prevent duplicate
// refunds on retry.
//
// ## Idempotency Key Benefits for Refunds
//
// Refunds are critical financial operations where duplicates can cause serious
// issues. Using idempotency keys ensures:
//
//  1. **Prevent duplicate refunds**: Even if a request times out or fails, retrying
//     with the same key won't issue multiple refunds
//  2. **Safe retries**: Network failures or timeouts won't risk creating multiple
//     refunds
//  3. **Consistent response**: Always get the same response for the same operation
//
// ## Example Request (curl)
//
// ```bash
//
//	curl -X POST \
//	  https://api.example.com/rest-api/v0/test-payment/pay_abc123xyz456/refund \
//	  -H 'Content-Type: application/json' \
//	  -H 'Client-Key: your-api-key' \
//	  -H 'Client-Secret: your-api-secret' \
//	  -H 'Idempotency-Key: refund_123456_unique_key'
//
// ```
//
// ## Example Response (First Request)
//
// ```json
//
//	{
//	  "id": "ref_xyz789",
//	  "paymentId": "pay_abc123xyz456",
//	  "amount": 100.0,
//	  "status": "succeeded",
//	  "timestamp": "2023-07-01T14:30:00.000Z"
//	}
//
// ```
//
// ## Example Response (Duplicate Request)
//
// The exact same response will be returned for any duplicate request with the same
// idempotency key, without creating a new refund.
//
// ## Idempotency Key Guidelines
//
// - Use a unique key for each distinct refund operation
// - Store keys client-side to ensure you can retry with the same key if needed
// - Keys expire after 24 hours by default
func (r *V0TestPaymentService) Refund(ctx context.Context, id string, opts ...option.RequestOption) (res *V0TestPaymentRefundResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/test-payment/%s/refund", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type PaymentResponse struct {
	// Payment ID
	ID string `json:"id,required"`
	// The amount charged
	Amount float64 `json:"amount,required"`
	// The currency code
	Currency string `json:"currency,required"`
	// Payment status
	Status string `json:"status,required"`
	// Timestamp of the payment
	Timestamp string `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		Currency    respjson.Field
		Status      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0TestPaymentRefundResponse struct {
	// Refund ID
	ID string `json:"id,required"`
	// The amount refunded
	Amount float64 `json:"amount,required"`
	// Original payment ID
	PaymentID string `json:"paymentId,required"`
	// Refund status
	Status string `json:"status,required"`
	// Timestamp of the refund
	Timestamp string `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		PaymentID   respjson.Field
		Status      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V0TestPaymentRefundResponse) RawJSON() string { return r.JSON.raw }
func (r *V0TestPaymentRefundResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0TestPaymentProcessParams struct {
	// The amount to charge
	Amount float64 `json:"amount,required"`
	// The currency code
	Currency string `json:"currency,required"`
	// Description of the payment
	Description string `json:"description,required"`
	// Customer reference ID
	CustomerID param.Opt[string] `json:"customerId,omitzero"`
	paramObj
}

func (r V0TestPaymentProcessParams) MarshalJSON() (data []byte, err error) {
	type shadow V0TestPaymentProcessParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0TestPaymentProcessParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
