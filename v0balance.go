// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/devdraft-go/internal/apijson"
	"github.com/stainless-sdks/devdraft-go/internal/requestconfig"
	"github.com/stainless-sdks/devdraft-go/option"
	"github.com/stainless-sdks/devdraft-go/packages/respjson"
)

// V0BalanceService contains methods and other services that help with interacting
// with the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0BalanceService] method instead.
type V0BalanceService struct {
	Options []option.RequestOption
}

// NewV0BalanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV0BalanceService(opts ...option.RequestOption) (r V0BalanceService) {
	r = V0BalanceService{}
	r.Options = opts
	return
}

// Retrieves both USDC and EURC balances across all wallets for the specified app.
//
// This comprehensive endpoint provides:
//
// - Complete USDC balance aggregation with detailed breakdown
// - Complete EURC balance aggregation with detailed breakdown
// - Total USD equivalent value across both currencies
// - Individual wallet and chain-specific balance details
//
// ## Use Cases
//
// - Complete financial dashboard overview
// - Multi-currency balance reporting
// - Comprehensive wallet management
// - Cross-currency analytics and reporting
//
// ## Response Format
//
// The response includes separate aggregations for each currency plus a combined
// USD value estimate, providing complete visibility into stablecoin holdings.
func (r *V0BalanceService) GetAllStablecoinBalances(ctx context.Context, opts ...option.RequestOption) (res *V0BalanceGetAllStablecoinBalancesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v0/balance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves the total EURC balance across all wallets for the specified app.
//
// This endpoint aggregates EURC balances from all associated wallets and provides:
//
// - Total EURC balance across all chains and wallets
// - Detailed breakdown by individual wallet and blockchain network
// - Contract addresses and wallet identifiers for each balance
//
// ## Use Cases
//
// - Dashboard balance display
// - European market operations
// - Euro-denominated financial reporting
// - Cross-currency balance tracking
//
// ## Response Format
//
// The response includes both the aggregated total and detailed breakdown, enabling
// comprehensive euro stablecoin balance management.
func (r *V0BalanceService) GetEurc(ctx context.Context, opts ...option.RequestOption) (res *AggregatedBalance, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v0/balance/eurc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves the total USDC balance across all wallets for the specified app.
//
// This endpoint aggregates USDC balances from all associated wallets and provides:
//
// - Total USDC balance across all chains and wallets
// - Detailed breakdown by individual wallet and blockchain network
// - Contract addresses and wallet identifiers for each balance
//
// ## Use Cases
//
// - Dashboard balance display
// - Financial reporting and reconciliation
// - Real-time balance monitoring
// - Multi-chain balance aggregation
//
// ## Response Format
//
// The response includes both the aggregated total and detailed breakdown, allowing
// for comprehensive balance tracking and wallet-specific analysis.
func (r *V0BalanceService) GetUsdc(ctx context.Context, opts ...option.RequestOption) (res *AggregatedBalance, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v0/balance/usdc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AggregatedBalance struct {
	// Detailed breakdown of balances by wallet and chain
	Balances [][]any `json:"balances,required"`
	// The stablecoin currency
	//
	// Any of "usdc", "eurc".
	Currency AggregatedBalanceCurrency `json:"currency,required"`
	// The total aggregated balance across all wallets and chains
	TotalBalance string `json:"total_balance,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Balances     respjson.Field
		Currency     respjson.Field
		TotalBalance respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AggregatedBalance) RawJSON() string { return r.JSON.raw }
func (r *AggregatedBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The stablecoin currency
type AggregatedBalanceCurrency string

const (
	AggregatedBalanceCurrencyUsdc AggregatedBalanceCurrency = "usdc"
	AggregatedBalanceCurrencyEurc AggregatedBalanceCurrency = "eurc"
)

type V0BalanceGetAllStablecoinBalancesResponse struct {
	// EURC balance aggregation
	Eurc AggregatedBalance `json:"eurc,required"`
	// Total value in USD equivalent
	TotalUsdValue string `json:"total_usd_value,required"`
	// USDC balance aggregation
	Usdc AggregatedBalance `json:"usdc,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Eurc          respjson.Field
		TotalUsdValue respjson.Field
		Usdc          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V0BalanceGetAllStablecoinBalancesResponse) RawJSON() string { return r.JSON.raw }
func (r *V0BalanceGetAllStablecoinBalancesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
