// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/devdraftengineer/devdraft-go/internal/apijson"
	"github.com/devdraftengineer/devdraft-go/internal/apiquery"
	"github.com/devdraftengineer/devdraft-go/internal/requestconfig"
	"github.com/devdraftengineer/devdraft-go/option"
	"github.com/devdraftengineer/devdraft-go/packages/respjson"
)

// V0ExchangeRateService contains methods and other services that help with
// interacting with the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0ExchangeRateService] method instead.
type V0ExchangeRateService struct {
	Options []option.RequestOption
}

// NewV0ExchangeRateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV0ExchangeRateService(opts ...option.RequestOption) (r V0ExchangeRateService) {
	r = V0ExchangeRateService{}
	r.Options = opts
	return
}

// Retrieves the current exchange rate for converting EUR to USD (EURC to USDC).
//
// This endpoint provides real-time exchange rate information for stablecoin
// conversions:
//
// - Mid-market rate for reference pricing
// - Buy rate for actual conversion calculations
// - Sell rate for reverse conversion scenarios
//
// ## Use Cases
//
// - Display current exchange rates in dashboards
// - Calculate conversion amounts for EURC to USDC transfers
// - Financial reporting and analytics
// - Real-time pricing for multi-currency operations
//
// ## Rate Information
//
// - **Mid-market rate**: The theoretical middle rate between buy and sell
// - **Buy rate**: Rate used when converting FROM EUR TO USD (what you get)
// - **Sell rate**: Rate used when converting FROM USD TO EUR (what you pay)
//
// The rates are updated in real-time and reflect current market conditions.
func (r *V0ExchangeRateService) GetEurToUsd(ctx context.Context, opts ...option.RequestOption) (res *ExchangeRateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v0/exchange-rate/eur-to-usd"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves the current exchange rate between two specified currencies.
//
// This flexible endpoint allows you to get exchange rates for any supported
// currency pair:
//
// - Supports USD and EUR currency codes
// - Provides comprehensive rate information
// - Real-time market data
//
// ## Supported Currency Pairs
//
// - USD to EUR (USDC to EURC)
// - EUR to USD (EURC to USDC)
//
// ## Query Parameters
//
// - **from**: Source currency code (usd, eur)
// - **to**: Target currency code (usd, eur)
//
// ## Use Cases
//
// - Flexible exchange rate queries
// - Multi-currency application support
// - Dynamic currency conversion calculations
// - Financial analytics and reporting
//
// ## Rate Information
//
// All rates are provided with full market context including mid-market, buy, and
// sell rates.
func (r *V0ExchangeRateService) GetExchangeRate(ctx context.Context, query V0ExchangeRateGetExchangeRateParams, opts ...option.RequestOption) (res *ExchangeRateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v0/exchange-rate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the current exchange rate for converting USD to EUR (USDC to EURC).
//
// This endpoint provides real-time exchange rate information for stablecoin
// conversions:
//
// - Mid-market rate for reference pricing
// - Buy rate for actual conversion calculations
// - Sell rate for reverse conversion scenarios
//
// ## Use Cases
//
// - Display current exchange rates in dashboards
// - Calculate conversion amounts for USDC to EURC transfers
// - Financial reporting and analytics
// - Real-time pricing for multi-currency operations
//
// ## Rate Information
//
// - **Mid-market rate**: The theoretical middle rate between buy and sell
// - **Buy rate**: Rate used when converting FROM USD TO EUR (what you get)
// - **Sell rate**: Rate used when converting FROM EUR TO USD (what you pay)
//
// The rates are updated in real-time and reflect current market conditions.
func (r *V0ExchangeRateService) GetUsdToEur(ctx context.Context, opts ...option.RequestOption) (res *ExchangeRateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v0/exchange-rate/usd-to-eur"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ExchangeRateResponse struct {
	// Rate for buying target currency (what you get when converting from source to
	// target)
	BuyRate string `json:"buy_rate,required"`
	// Source currency code (USD for USDC)
	From string `json:"from,required"`
	// Mid-market exchange rate from source to target currency
	MidmarketRate string `json:"midmarket_rate,required"`
	// Rate for selling target currency (what you pay when converting from target to
	// source)
	SellRate string `json:"sell_rate,required"`
	// Target currency code (EUR for EURC)
	To string `json:"to,required"`
	// Timestamp when the exchange rate was last updated
	Timestamp string `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BuyRate       respjson.Field
		From          respjson.Field
		MidmarketRate respjson.Field
		SellRate      respjson.Field
		To            respjson.Field
		Timestamp     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExchangeRateResponse) RawJSON() string { return r.JSON.raw }
func (r *ExchangeRateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0ExchangeRateGetExchangeRateParams struct {
	// Source currency code (e.g., usd)
	From string `query:"from,required" json:"-"`
	// Target currency code (e.g., eur)
	To string `query:"to,required" json:"-"`
	paramObj
}

// URLQuery serializes [V0ExchangeRateGetExchangeRateParams]'s query parameters as
// `url.Values`.
func (r V0ExchangeRateGetExchangeRateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
