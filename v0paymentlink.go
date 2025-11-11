// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/devdraftengineer/devdraft-go/internal/apijson"
	"github.com/devdraftengineer/devdraft-go/internal/apiquery"
	"github.com/devdraftengineer/devdraft-go/internal/requestconfig"
	"github.com/devdraftengineer/devdraft-go/option"
	"github.com/devdraftengineer/devdraft-go/packages/param"
)

// V0PaymentLinkService contains methods and other services that help with
// interacting with the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0PaymentLinkService] method instead.
type V0PaymentLinkService struct {
	Options []option.RequestOption
}

// NewV0PaymentLinkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV0PaymentLinkService(opts ...option.RequestOption) (r V0PaymentLinkService) {
	r = V0PaymentLinkService{}
	r.Options = opts
	return
}

// Creates a new payment link with the provided details. Supports both simple
// one-time payments and complex product bundles.
func (r *V0PaymentLinkService) New(ctx context.Context, body V0PaymentLinkNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/payment-links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get a payment link by ID
func (r *V0PaymentLinkService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/payment-links/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Update a payment link
func (r *V0PaymentLinkService) Update(ctx context.Context, id string, body V0PaymentLinkUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/payment-links/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Get all payment links
func (r *V0PaymentLinkService) List(ctx context.Context, query V0PaymentLinkListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/payment-links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type V0PaymentLinkNewParams struct {
	// Whether to allow mobile payment
	AllowMobilePayment bool `json:"allowMobilePayment,required"`
	// Whether to allow quantity adjustment
	AllowQuantityAdjustment bool `json:"allowQuantityAdjustment,required"`
	// Whether to collect address
	CollectAddress bool `json:"collectAddress,required"`
	// Whether to collect tax
	CollectTax bool `json:"collectTax,required"`
	// Currency
	//
	// Any of "usdc", "eurc".
	Currency V0PaymentLinkNewParamsCurrency `json:"currency,omitzero,required"`
	// Type of the payment link
	//
	// Any of "INVOICE", "PRODUCT", "COLLECTION", "DONATION".
	LinkType V0PaymentLinkNewParamsLinkType `json:"linkType,omitzero,required"`
	// Display title for the payment link. This appears on the checkout page and in
	// customer communications.
	Title string `json:"title,required"`
	// Unique URL slug for the payment link. Can be a full URL or just the path
	// segment. Must be unique within your account.
	URL string `json:"url,required"`
	// Amount for the payment link
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Cover image URL
	CoverImage param.Opt[string] `json:"coverImage,omitzero"`
	// Customer ID
	CustomerID param.Opt[string] `json:"customerId,omitzero"`
	// Detailed description of what the customer is purchasing. Supports markdown
	// formatting.
	Description param.Opt[string] `json:"description,omitzero"`
	// Expiration date
	ExpirationDate param.Opt[time.Time] `json:"expiration_date,omitzero" format:"date-time"`
	// Whether the payment link is for all products
	IsForAllProduct param.Opt[bool] `json:"isForAllProduct,omitzero"`
	// Whether to limit payments
	LimitPayments param.Opt[bool] `json:"limitPayments,omitzero"`
	// Maximum number of payments
	MaxPayments param.Opt[float64] `json:"maxPayments,omitzero"`
	// Payment for ID
	PaymentForID param.Opt[string] `json:"paymentForId,omitzero"`
	// Tax ID
	TaxID param.Opt[string] `json:"taxId,omitzero"`
	// Custom fields
	CustomFields any `json:"customFields,omitzero"`
	// Array of products in the payment link
	PaymentLinkProducts []V0PaymentLinkNewParamsPaymentLinkProduct `json:"paymentLinkProducts,omitzero"`
	paramObj
}

func (r V0PaymentLinkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V0PaymentLinkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0PaymentLinkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Currency
type V0PaymentLinkNewParamsCurrency string

const (
	V0PaymentLinkNewParamsCurrencyUsdc V0PaymentLinkNewParamsCurrency = "usdc"
	V0PaymentLinkNewParamsCurrencyEurc V0PaymentLinkNewParamsCurrency = "eurc"
)

// Type of the payment link
type V0PaymentLinkNewParamsLinkType string

const (
	V0PaymentLinkNewParamsLinkTypeInvoice    V0PaymentLinkNewParamsLinkType = "INVOICE"
	V0PaymentLinkNewParamsLinkTypeProduct    V0PaymentLinkNewParamsLinkType = "PRODUCT"
	V0PaymentLinkNewParamsLinkTypeCollection V0PaymentLinkNewParamsLinkType = "COLLECTION"
	V0PaymentLinkNewParamsLinkTypeDonation   V0PaymentLinkNewParamsLinkType = "DONATION"
)

// The properties ProductID, Quantity are required.
type V0PaymentLinkNewParamsPaymentLinkProduct struct {
	// UUID of the product to include in this payment link. Must be a valid product
	// from your catalog.
	ProductID string `json:"productId,required" format:"uuid"`
	// Quantity of this product to include. Must be at least 1.
	Quantity int64 `json:"quantity,required"`
	paramObj
}

func (r V0PaymentLinkNewParamsPaymentLinkProduct) MarshalJSON() (data []byte, err error) {
	type shadow V0PaymentLinkNewParamsPaymentLinkProduct
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0PaymentLinkNewParamsPaymentLinkProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0PaymentLinkUpdateParams struct {
	paramObj
}

func (r V0PaymentLinkUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V0PaymentLinkUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0PaymentLinkUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0PaymentLinkListParams struct {
	// Number of records to skip (must be non-negative)
	Skip param.Opt[string] `query:"skip,omitzero" json:"-"`
	// Number of records to take (must be positive)
	Take param.Opt[string] `query:"take,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V0PaymentLinkListParams]'s query parameters as
// `url.Values`.
func (r V0PaymentLinkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
