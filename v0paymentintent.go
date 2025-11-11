// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/devdraft-go/internal/apijson"
	"github.com/stainless-sdks/devdraft-go/internal/requestconfig"
	"github.com/stainless-sdks/devdraft-go/option"
	"github.com/stainless-sdks/devdraft-go/packages/param"
)

// V0PaymentIntentService contains methods and other services that help with
// interacting with the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0PaymentIntentService] method instead.
type V0PaymentIntentService struct {
	Options []option.RequestOption
}

// NewV0PaymentIntentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV0PaymentIntentService(opts ...option.RequestOption) (r V0PaymentIntentService) {
	r = V0PaymentIntentService{}
	r.Options = opts
	return
}

// Creates a new bank payment intent for fiat-to-stablecoin transfers.
//
// This endpoint allows you to create payment intents for bank transfers (ACH,
// Wire, SEPA) that convert to stablecoins. Perfect for onboarding users from
// traditional banking to crypto.
//
// ## Supported Payment Rails
//
// - **ACH_PUSH**: US bank transfers (same-day or standard)
// - **WIRE**: International wire transfers
// - **SEPA**: European bank transfers
//
// ## Use Cases
//
// - USD bank account to USDC conversion
// - EUR bank account to EURC conversion
// - MXN bank account to stablecoin conversion
// - Flexible amount payment intents for variable pricing
//
// ## Supported Source Currencies
//
// - **USD**: US Dollar
// - **EUR**: Euro
// - **MXN**: Mexican Peso
//
// ## Example: USD Bank to USDC
//
// ```json
//
//	{
//	  "sourcePaymentRail": "ach_push",
//	  "sourceCurrency": "usd",
//	  "destinationCurrency": "usdc",
//	  "destinationNetwork": "ethereum",
//	  "destinationAddress": "0x742d35Cc6634C0532925a3b8D4C9db96c4b4d8e1",
//	  "amount": "1000.00",
//	  "customer_first_name": "John",
//	  "customer_last_name": "Doe",
//	  "customer_email": "john.doe@example.com",
//	  "ach_reference": "INV12345"
//	}
//
// ```
//
// ## Reference Fields
//
// Use appropriate reference fields based on the payment rail:
//
// - `ach_reference`: For ACH transfers (max 10 chars, alphanumeric + spaces)
// - `wire_message`: For wire transfers (max 256 chars)
// - `sepa_reference`: For SEPA transfers (6-140 chars, specific character set)
//
// ## Idempotency
//
// Include an `idempotency-key` header with a unique UUID v4 to prevent duplicate
// payments. Subsequent requests with the same key will return the original
// response.
func (r *V0PaymentIntentService) NewBank(ctx context.Context, body V0PaymentIntentNewBankParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/payment-intents/bank"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Creates a new stable payment intent for stablecoin-to-stablecoin transfers.
//
// This endpoint allows you to create payment intents for transfers between
// different stablecoins and networks. Perfect for cross-chain stablecoin swaps and
// conversions.
//
// ## Use Cases
//
// - USDC to EURC conversions
// - Cross-chain stablecoin transfers (e.g., Ethereum USDC to Polygon USDC)
// - Flexible amount payment intents for dynamic pricing
//
// ## Example: USDC to EURC Conversion
//
// ```json
//
//	{
//	  "sourceCurrency": "usdc",
//	  "sourceNetwork": "ethereum",
//	  "destinationCurrency": "eurc",
//	  "destinationNetwork": "polygon",
//	  "destinationAddress": "0x742d35Cc6634C0532925a3b8D4C9db96c4b4d8e1",
//	  "amount": "100.00",
//	  "customer_first_name": "John",
//	  "customer_last_name": "Doe",
//	  "customer_email": "john.doe@example.com"
//	}
//
// ```
//
// ## Flexible Amount Payments
//
// Omit the `amount` field to create a flexible payment intent where users can
// specify the amount during payment.
//
// ## Idempotency
//
// Include an `idempotency-key` header with a unique UUID v4 to prevent duplicate
// payments. Subsequent requests with the same key will return the original
// response.
func (r *V0PaymentIntentService) NewStable(ctx context.Context, body V0PaymentIntentNewStableParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/payment-intents/stablecoin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// The blockchain network where the source currency resides. Determines gas fees
// and transaction speed.
type BridgePaymentRail string

const (
	BridgePaymentRailEthereum        BridgePaymentRail = "ethereum"
	BridgePaymentRailSolana          BridgePaymentRail = "solana"
	BridgePaymentRailPolygon         BridgePaymentRail = "polygon"
	BridgePaymentRailAvalancheCChain BridgePaymentRail = "avalanche_c_chain"
	BridgePaymentRailBase            BridgePaymentRail = "base"
	BridgePaymentRailArbitrum        BridgePaymentRail = "arbitrum"
	BridgePaymentRailOptimism        BridgePaymentRail = "optimism"
	BridgePaymentRailStellar         BridgePaymentRail = "stellar"
	BridgePaymentRailTron            BridgePaymentRail = "tron"
	BridgePaymentRailBridgeWallet    BridgePaymentRail = "bridge_wallet"
	BridgePaymentRailWire            BridgePaymentRail = "wire"
	BridgePaymentRailACH             BridgePaymentRail = "ach"
	BridgePaymentRailACHPush         BridgePaymentRail = "ach_push"
	BridgePaymentRailACHSameDay      BridgePaymentRail = "ach_same_day"
	BridgePaymentRailSepa            BridgePaymentRail = "sepa"
	BridgePaymentRailSwift           BridgePaymentRail = "swift"
	BridgePaymentRailSpei            BridgePaymentRail = "spei"
)

// The stablecoin currency to convert FROM. This is the currency the customer will
// pay with.
type StableCoinCurrency string

const (
	StableCoinCurrencyUsdc StableCoinCurrency = "usdc"
	StableCoinCurrencyEurc StableCoinCurrency = "eurc"
)

type V0PaymentIntentNewBankParams struct {
	// The stablecoin currency to convert FROM. This is the currency the customer will
	// pay with.
	//
	// Any of "usdc", "eurc".
	DestinationCurrency StableCoinCurrency `json:"destinationCurrency,omitzero,required"`
	// The blockchain network where the source currency resides. Determines gas fees
	// and transaction speed.
	//
	// Any of "ethereum", "solana", "polygon", "avalanche_c_chain", "base", "arbitrum",
	// "optimism", "stellar", "tron", "bridge_wallet", "wire", "ach", "ach_push",
	// "ach_same_day", "sepa", "swift", "spei".
	DestinationNetwork BridgePaymentRail `json:"destinationNetwork,omitzero,required"`
	// The fiat currency to convert FROM. Must match the currency of the source payment
	// rail.
	//
	// Any of "usd", "eur", "mxn".
	SourceCurrency V0PaymentIntentNewBankParamsSourceCurrency `json:"sourceCurrency,omitzero,required"`
	// The blockchain network where the source currency resides. Determines gas fees
	// and transaction speed.
	//
	// Any of "ethereum", "solana", "polygon", "avalanche_c_chain", "base", "arbitrum",
	// "optimism", "stellar", "tron", "bridge_wallet", "wire", "ach", "ach_push",
	// "ach_same_day", "sepa", "swift", "spei".
	SourcePaymentRail BridgePaymentRail `json:"sourcePaymentRail,omitzero,required"`
	// ACH reference (for ACH transfers)
	ACHReference param.Opt[string] `json:"ach_reference,omitzero"`
	// Payment amount (optional for flexible amount)
	Amount param.Opt[string] `json:"amount,omitzero"`
	// Customer address
	CustomerAddress param.Opt[string] `json:"customer_address,omitzero"`
	// Customer country
	CustomerCountry param.Opt[string] `json:"customer_country,omitzero"`
	// Customer country ISO code
	CustomerCountryISO param.Opt[string] `json:"customer_countryISO,omitzero"`
	// Customer email address
	CustomerEmail param.Opt[string] `json:"customer_email,omitzero"`
	// Customer first name
	CustomerFirstName param.Opt[string] `json:"customer_first_name,omitzero"`
	// Customer last name
	CustomerLastName param.Opt[string] `json:"customer_last_name,omitzero"`
	// Customer province/state
	CustomerProvince param.Opt[string] `json:"customer_province,omitzero"`
	// Customer province/state ISO code
	CustomerProvinceISO param.Opt[string] `json:"customer_provinceISO,omitzero"`
	// Destination wallet address. Supports Ethereum (0x...) and Solana address
	// formats.
	DestinationAddress param.Opt[string] `json:"destinationAddress,omitzero"`
	// Customer phone number
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	// SEPA reference (for SEPA transfers)
	SepaReference param.Opt[string] `json:"sepa_reference,omitzero"`
	// Wire transfer message (for WIRE transfers)
	WireMessage param.Opt[string] `json:"wire_message,omitzero"`
	paramObj
}

func (r V0PaymentIntentNewBankParams) MarshalJSON() (data []byte, err error) {
	type shadow V0PaymentIntentNewBankParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0PaymentIntentNewBankParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The fiat currency to convert FROM. Must match the currency of the source payment
// rail.
type V0PaymentIntentNewBankParamsSourceCurrency string

const (
	V0PaymentIntentNewBankParamsSourceCurrencyUsd V0PaymentIntentNewBankParamsSourceCurrency = "usd"
	V0PaymentIntentNewBankParamsSourceCurrencyEur V0PaymentIntentNewBankParamsSourceCurrency = "eur"
	V0PaymentIntentNewBankParamsSourceCurrencyMxn V0PaymentIntentNewBankParamsSourceCurrency = "mxn"
)

type V0PaymentIntentNewStableParams struct {
	// The blockchain network where the source currency resides. Determines gas fees
	// and transaction speed.
	//
	// Any of "ethereum", "solana", "polygon", "avalanche_c_chain", "base", "arbitrum",
	// "optimism", "stellar", "tron", "bridge_wallet", "wire", "ach", "ach_push",
	// "ach_same_day", "sepa", "swift", "spei".
	DestinationNetwork BridgePaymentRail `json:"destinationNetwork,omitzero,required"`
	// The stablecoin currency to convert FROM. This is the currency the customer will
	// pay with.
	//
	// Any of "usdc", "eurc".
	SourceCurrency StableCoinCurrency `json:"sourceCurrency,omitzero,required"`
	// The blockchain network where the source currency resides. Determines gas fees
	// and transaction speed.
	//
	// Any of "ethereum", "solana", "polygon", "avalanche_c_chain", "base", "arbitrum",
	// "optimism", "stellar", "tron", "bridge_wallet", "wire", "ach", "ach_push",
	// "ach_same_day", "sepa", "swift", "spei".
	SourceNetwork BridgePaymentRail `json:"sourceNetwork,omitzero,required"`
	// Payment amount in the source currency. Omit for flexible amount payments where
	// users specify the amount during checkout.
	Amount param.Opt[string] `json:"amount,omitzero"`
	// Customer's full address. Required for compliance in certain jurisdictions and
	// high-value transactions.
	CustomerAddress param.Opt[string] `json:"customer_address,omitzero"`
	// Customer's country of residence. Used for compliance and tax reporting.
	CustomerCountry param.Opt[string] `json:"customer_country,omitzero"`
	// Customer's country ISO 3166-1 alpha-2 code. Used for automated compliance
	// checks.
	CustomerCountryISO param.Opt[string] `json:"customer_countryISO,omitzero"`
	// Customer's email address. Used for transaction notifications and receipts.
	// Highly recommended for all transactions.
	CustomerEmail param.Opt[string] `json:"customer_email,omitzero" format:"email"`
	// Customer's first name. Used for transaction records and compliance. Required for
	// amounts over $1000.
	CustomerFirstName param.Opt[string] `json:"customer_first_name,omitzero"`
	// Customer's last name. Used for transaction records and compliance. Required for
	// amounts over $1000.
	CustomerLastName param.Opt[string] `json:"customer_last_name,omitzero"`
	// Customer's state or province. Required for US and Canadian customers.
	CustomerProvince param.Opt[string] `json:"customer_province,omitzero"`
	// Customer's state or province ISO code. Used for automated tax calculations.
	CustomerProvinceISO param.Opt[string] `json:"customer_provinceISO,omitzero"`
	// The wallet address where converted funds will be sent. Supports Ethereum (0x...)
	// and Solana address formats.
	DestinationAddress param.Opt[string] `json:"destinationAddress,omitzero"`
	// Customer's phone number with country code. Used for SMS notifications and
	// verification.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	// The stablecoin currency to convert FROM. This is the currency the customer will
	// pay with.
	//
	// Any of "usdc", "eurc".
	DestinationCurrency StableCoinCurrency `json:"destinationCurrency,omitzero"`
	paramObj
}

func (r V0PaymentIntentNewStableParams) MarshalJSON() (data []byte, err error) {
	type shadow V0PaymentIntentNewStableParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0PaymentIntentNewStableParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
