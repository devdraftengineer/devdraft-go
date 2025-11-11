// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/devdraft-go/internal/requestconfig"
	"github.com/stainless-sdks/devdraft-go/option"
)

// V0Service contains methods and other services that help with interacting with
// the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0Service] method instead.
type V0Service struct {
	Options        []option.RequestOption
	Health         V0HealthService
	TestPayment    V0TestPaymentService
	Customers      V0CustomerService
	PaymentLinks   V0PaymentLinkService
	PaymentIntents V0PaymentIntentService
	Webhooks       V0WebhookService
	Transfers      V0TransferService
	Balance        V0BalanceService
	ExchangeRate   V0ExchangeRateService
	Products       V0ProductService
	Invoices       V0InvoiceService
	Taxes          V0TaxService
}

// NewV0Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV0Service(opts ...option.RequestOption) (r V0Service) {
	r = V0Service{}
	r.Options = opts
	r.Health = NewV0HealthService(opts...)
	r.TestPayment = NewV0TestPaymentService(opts...)
	r.Customers = NewV0CustomerService(opts...)
	r.PaymentLinks = NewV0PaymentLinkService(opts...)
	r.PaymentIntents = NewV0PaymentIntentService(opts...)
	r.Webhooks = NewV0WebhookService(opts...)
	r.Transfers = NewV0TransferService(opts...)
	r.Balance = NewV0BalanceService(opts...)
	r.ExchangeRate = NewV0ExchangeRateService(opts...)
	r.Products = NewV0ProductService(opts...)
	r.Invoices = NewV0InvoiceService(opts...)
	r.Taxes = NewV0TaxService(opts...)
	return
}

// Get wallets for an app
func (r *V0Service) GetWallets(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/wallets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}
