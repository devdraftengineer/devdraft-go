// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/devdraft-go/internal/apijson"
	"github.com/stainless-sdks/devdraft-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/devdraft-go/internal/encoding/json"
	"github.com/stainless-sdks/devdraft-go/internal/requestconfig"
	"github.com/stainless-sdks/devdraft-go/option"
	"github.com/stainless-sdks/devdraft-go/packages/param"
)

// V0InvoiceService contains methods and other services that help with interacting
// with the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0InvoiceService] method instead.
type V0InvoiceService struct {
	Options []option.RequestOption
}

// NewV0InvoiceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV0InvoiceService(opts ...option.RequestOption) (r V0InvoiceService) {
	r = V0InvoiceService{}
	r.Options = opts
	return
}

// Create a new invoice
func (r *V0InvoiceService) New(ctx context.Context, body V0InvoiceNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Get an invoice by ID
func (r *V0InvoiceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/invoices/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Update an invoice
func (r *V0InvoiceService) Update(ctx context.Context, id string, body V0InvoiceUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/invoices/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Get all invoices
func (r *V0InvoiceService) List(ctx context.Context, query V0InvoiceListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// The properties Currency, CustomerID, Delivery, DueDate, Email, Items, Name,
// PartialPayment, PaymentLink, PaymentMethods, Status are required.
type CreateInvoiceParam struct {
	// Currency for the invoice
	//
	// Any of "usdc", "eurc".
	Currency CreateInvoiceCurrency `json:"currency,omitzero,required"`
	// Customer ID
	CustomerID string `json:"customer_id,required"`
	// Delivery method
	//
	// Any of "EMAIL", "MANUALLY".
	Delivery CreateInvoiceDelivery `json:"delivery,omitzero,required"`
	// Due date of the invoice
	DueDate time.Time `json:"due_date,required" format:"date-time"`
	// Email address
	Email string `json:"email,required"`
	// Array of products in the invoice
	Items []CreateInvoiceItemParam `json:"items,omitzero,required"`
	// Name of the invoice
	Name string `json:"name,required"`
	// Allow partial payments
	PartialPayment bool `json:"partial_payment,required"`
	// Whether to generate a payment link
	PaymentLink bool `json:"payment_link,required"`
	// Array of accepted payment methods
	//
	// Any of "ACH", "BANK_TRANSFER", "CREDIT_CARD", "CASH", "MOBILE_MONEY", "CRYPTO".
	PaymentMethods []string `json:"payment_methods,omitzero,required"`
	// Invoice status
	//
	// Any of "DRAFT", "OPEN", "PASTDUE", "PAID", "PARTIALLYPAID".
	Status CreateInvoiceStatus `json:"status,omitzero,required"`
	// Address
	Address param.Opt[string] `json:"address,omitzero"`
	// Logo URL
	Logo param.Opt[string] `json:"logo,omitzero"`
	// Phone number
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// Send date
	SendDate param.Opt[time.Time] `json:"send_date,omitzero" format:"date-time"`
	// Tax ID
	TaxID param.Opt[string] `json:"taxId,omitzero"`
	paramObj
}

func (r CreateInvoiceParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateInvoiceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateInvoiceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Currency for the invoice
type CreateInvoiceCurrency string

const (
	CreateInvoiceCurrencyUsdc CreateInvoiceCurrency = "usdc"
	CreateInvoiceCurrencyEurc CreateInvoiceCurrency = "eurc"
)

// Delivery method
type CreateInvoiceDelivery string

const (
	CreateInvoiceDeliveryEmail    CreateInvoiceDelivery = "EMAIL"
	CreateInvoiceDeliveryManually CreateInvoiceDelivery = "MANUALLY"
)

// The properties ProductID, Quantity are required.
type CreateInvoiceItemParam struct {
	// Product ID
	ProductID string `json:"product_id,required"`
	// Quantity of the product
	Quantity float64 `json:"quantity,required"`
	paramObj
}

func (r CreateInvoiceItemParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateInvoiceItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateInvoiceItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Invoice status
type CreateInvoiceStatus string

const (
	CreateInvoiceStatusDraft         CreateInvoiceStatus = "DRAFT"
	CreateInvoiceStatusOpen          CreateInvoiceStatus = "OPEN"
	CreateInvoiceStatusPastdue       CreateInvoiceStatus = "PASTDUE"
	CreateInvoiceStatusPaid          CreateInvoiceStatus = "PAID"
	CreateInvoiceStatusPartiallypaid CreateInvoiceStatus = "PARTIALLYPAID"
)

type V0InvoiceNewParams struct {
	CreateInvoice CreateInvoiceParam
	paramObj
}

func (r V0InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateInvoice)
}
func (r *V0InvoiceNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateInvoice)
}

type V0InvoiceUpdateParams struct {
	CreateInvoice CreateInvoiceParam
	paramObj
}

func (r V0InvoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateInvoice)
}
func (r *V0InvoiceUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateInvoice)
}

type V0InvoiceListParams struct {
	// Number of records to skip
	Skip param.Opt[float64] `query:"skip,omitzero" json:"-"`
	// Number of records to take
	Take param.Opt[float64] `query:"take,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V0InvoiceListParams]'s query parameters as `url.Values`.
func (r V0InvoiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
