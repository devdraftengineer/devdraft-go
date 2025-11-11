// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/devdraft-go/internal/apijson"
	"github.com/stainless-sdks/devdraft-go/internal/requestconfig"
	"github.com/stainless-sdks/devdraft-go/option"
	"github.com/stainless-sdks/devdraft-go/packages/param"
	"github.com/stainless-sdks/devdraft-go/packages/respjson"
)

// V0CustomerLiquidationAddressService contains methods and other services that
// help with interacting with the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0CustomerLiquidationAddressService] method instead.
type V0CustomerLiquidationAddressService struct {
	Options []option.RequestOption
}

// NewV0CustomerLiquidationAddressService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV0CustomerLiquidationAddressService(opts ...option.RequestOption) (r V0CustomerLiquidationAddressService) {
	r = V0CustomerLiquidationAddressService{}
	r.Options = opts
	return
}

// Create a new liquidation address for a customer. Liquidation addresses allow
//
//	customers to automatically liquidate cryptocurrency holdings to fiat or other
//	stablecoins based on configured parameters.
//
//	**Required fields:**
//	- chain: Blockchain network (ethereum, polygon, arbitrum, base)
//	- currency: Stablecoin currency (usdc, eurc, dai, pyusd, usdt)
//	- address: Valid blockchain address
//
//	**At least one destination must be specified:**
//	- external_account_id: External bank account
//	- prefunded_account_id: Developer's prefunded account
//	- bridge_wallet_id: Bridge wallet
//	- destination_address: Crypto wallet address
//
//	**Payment Rails:**
//	Different payment rails have different requirements:
//	- ACH: Requires external_account_id, supports destination_ach_reference
//	- SEPA: Requires external_account_id, supports destination_sepa_reference
//	- Wire: Requires external_account_id, supports destination_wire_message
//	- Crypto: Requires destination_address, supports destination_blockchain_memo
func (r *V0CustomerLiquidationAddressService) New(ctx context.Context, customerID string, body V0CustomerLiquidationAddressNewParams, opts ...option.RequestOption) (res *LiquidationAddressResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("api/v0/customers/%s/liquidation_addresses", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific liquidation address by its ID for a given customer.
func (r *V0CustomerLiquidationAddressService) Get(ctx context.Context, liquidationAddressID string, query V0CustomerLiquidationAddressGetParams, opts ...option.RequestOption) (res *LiquidationAddressResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.CustomerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	if liquidationAddressID == "" {
		err = errors.New("missing required liquidationAddressId parameter")
		return
	}
	path := fmt.Sprintf("api/v0/customers/%s/liquidation_addresses/%s", query.CustomerID, liquidationAddressID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve all liquidation addresses associated with a specific customer.
func (r *V0CustomerLiquidationAddressService) List(ctx context.Context, customerID string, opts ...option.RequestOption) (res *[]LiquidationAddressResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("api/v0/customers/%s/liquidation_addresses", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LiquidationAddressResponse struct {
	// Unique identifier for the liquidation address
	ID string `json:"id,required"`
	// Liquidation address
	Address string `json:"address,required"`
	// Blockchain chain
	Chain string `json:"chain,required"`
	// Creation timestamp
	CreatedAt string `json:"created_at,required"`
	// Currency
	Currency string `json:"currency,required"`
	// Customer ID this liquidation address belongs to
	CustomerID string `json:"customer_id,required"`
	// Current state of the liquidation address
	State string `json:"state,required"`
	// Last update timestamp
	UpdatedAt string `json:"updated_at,required"`
	// Bridge wallet ID
	BridgeWalletID string `json:"bridge_wallet_id"`
	// Custom developer fee percent
	CustomDeveloperFeePercent string `json:"custom_developer_fee_percent"`
	// Destination currency
	DestinationCurrency string `json:"destination_currency"`
	// Destination payment rail
	DestinationPaymentRail string `json:"destination_payment_rail"`
	// External account ID
	ExternalAccountID string `json:"external_account_id"`
	// Prefunded account ID
	PrefundedAccountID string `json:"prefunded_account_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		Address                   respjson.Field
		Chain                     respjson.Field
		CreatedAt                 respjson.Field
		Currency                  respjson.Field
		CustomerID                respjson.Field
		State                     respjson.Field
		UpdatedAt                 respjson.Field
		BridgeWalletID            respjson.Field
		CustomDeveloperFeePercent respjson.Field
		DestinationCurrency       respjson.Field
		DestinationPaymentRail    respjson.Field
		ExternalAccountID         respjson.Field
		PrefundedAccountID        respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LiquidationAddressResponse) RawJSON() string { return r.JSON.raw }
func (r *LiquidationAddressResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0CustomerLiquidationAddressNewParams struct {
	// The liquidation address on the blockchain
	Address string `json:"address,required"`
	// The blockchain chain for the liquidation address
	//
	// Any of "ethereum", "solana", "polygon", "avalanche_c_chain", "base", "arbitrum",
	// "optimism", "stellar", "tron".
	Chain V0CustomerLiquidationAddressNewParamsChain `json:"chain,omitzero,required"`
	// The currency for the liquidation address
	//
	// Any of "usdc", "eurc", "dai", "pyusd", "usdt".
	Currency V0CustomerLiquidationAddressNewParamsCurrency `json:"currency,omitzero,required"`
	// Bridge Wallet to send funds to
	BridgeWalletID param.Opt[string] `json:"bridge_wallet_id,omitzero"`
	// Custom developer fee percentage (Base 100 percentage: 10.2% = "10.2")
	CustomDeveloperFeePercent param.Opt[string] `json:"custom_developer_fee_percent,omitzero"`
	// Reference for ACH transactions
	DestinationACHReference param.Opt[string] `json:"destination_ach_reference,omitzero"`
	// Crypto wallet address for crypto transfers
	DestinationAddress param.Opt[string] `json:"destination_address,omitzero"`
	// Memo for blockchain transactions
	DestinationBlockchainMemo param.Opt[string] `json:"destination_blockchain_memo,omitzero"`
	// Reference for SEPA transactions
	DestinationSepaReference param.Opt[string] `json:"destination_sepa_reference,omitzero"`
	// Message for wire transfers
	DestinationWireMessage param.Opt[string] `json:"destination_wire_message,omitzero"`
	// External bank account to send funds to
	ExternalAccountID param.Opt[string] `json:"external_account_id,omitzero"`
	// Developer's prefunded account id
	PrefundedAccountID param.Opt[string] `json:"prefunded_account_id,omitzero"`
	// Address to return funds on failed transactions (Not supported on Stellar)
	ReturnAddress param.Opt[string] `json:"return_address,omitzero"`
	// Currency for sending funds
	//
	// Any of "usd", "eur", "mxn", "usdc", "eurc", "dai", "pyusd", "usdt".
	DestinationCurrency V0CustomerLiquidationAddressNewParamsDestinationCurrency `json:"destination_currency,omitzero"`
	// The blockchain network where the source currency resides. Determines gas fees
	// and transaction speed.
	//
	// Any of "ethereum", "solana", "polygon", "avalanche_c_chain", "base", "arbitrum",
	// "optimism", "stellar", "tron", "bridge_wallet", "wire", "ach", "ach_push",
	// "ach_same_day", "sepa", "swift", "spei".
	DestinationPaymentRail BridgePaymentRail `json:"destination_payment_rail,omitzero"`
	paramObj
}

func (r V0CustomerLiquidationAddressNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V0CustomerLiquidationAddressNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0CustomerLiquidationAddressNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The blockchain chain for the liquidation address
type V0CustomerLiquidationAddressNewParamsChain string

const (
	V0CustomerLiquidationAddressNewParamsChainEthereum        V0CustomerLiquidationAddressNewParamsChain = "ethereum"
	V0CustomerLiquidationAddressNewParamsChainSolana          V0CustomerLiquidationAddressNewParamsChain = "solana"
	V0CustomerLiquidationAddressNewParamsChainPolygon         V0CustomerLiquidationAddressNewParamsChain = "polygon"
	V0CustomerLiquidationAddressNewParamsChainAvalancheCChain V0CustomerLiquidationAddressNewParamsChain = "avalanche_c_chain"
	V0CustomerLiquidationAddressNewParamsChainBase            V0CustomerLiquidationAddressNewParamsChain = "base"
	V0CustomerLiquidationAddressNewParamsChainArbitrum        V0CustomerLiquidationAddressNewParamsChain = "arbitrum"
	V0CustomerLiquidationAddressNewParamsChainOptimism        V0CustomerLiquidationAddressNewParamsChain = "optimism"
	V0CustomerLiquidationAddressNewParamsChainStellar         V0CustomerLiquidationAddressNewParamsChain = "stellar"
	V0CustomerLiquidationAddressNewParamsChainTron            V0CustomerLiquidationAddressNewParamsChain = "tron"
)

// The currency for the liquidation address
type V0CustomerLiquidationAddressNewParamsCurrency string

const (
	V0CustomerLiquidationAddressNewParamsCurrencyUsdc  V0CustomerLiquidationAddressNewParamsCurrency = "usdc"
	V0CustomerLiquidationAddressNewParamsCurrencyEurc  V0CustomerLiquidationAddressNewParamsCurrency = "eurc"
	V0CustomerLiquidationAddressNewParamsCurrencyDai   V0CustomerLiquidationAddressNewParamsCurrency = "dai"
	V0CustomerLiquidationAddressNewParamsCurrencyPyusd V0CustomerLiquidationAddressNewParamsCurrency = "pyusd"
	V0CustomerLiquidationAddressNewParamsCurrencyUsdt  V0CustomerLiquidationAddressNewParamsCurrency = "usdt"
)

// Currency for sending funds
type V0CustomerLiquidationAddressNewParamsDestinationCurrency string

const (
	V0CustomerLiquidationAddressNewParamsDestinationCurrencyUsd   V0CustomerLiquidationAddressNewParamsDestinationCurrency = "usd"
	V0CustomerLiquidationAddressNewParamsDestinationCurrencyEur   V0CustomerLiquidationAddressNewParamsDestinationCurrency = "eur"
	V0CustomerLiquidationAddressNewParamsDestinationCurrencyMxn   V0CustomerLiquidationAddressNewParamsDestinationCurrency = "mxn"
	V0CustomerLiquidationAddressNewParamsDestinationCurrencyUsdc  V0CustomerLiquidationAddressNewParamsDestinationCurrency = "usdc"
	V0CustomerLiquidationAddressNewParamsDestinationCurrencyEurc  V0CustomerLiquidationAddressNewParamsDestinationCurrency = "eurc"
	V0CustomerLiquidationAddressNewParamsDestinationCurrencyDai   V0CustomerLiquidationAddressNewParamsDestinationCurrency = "dai"
	V0CustomerLiquidationAddressNewParamsDestinationCurrencyPyusd V0CustomerLiquidationAddressNewParamsDestinationCurrency = "pyusd"
	V0CustomerLiquidationAddressNewParamsDestinationCurrencyUsdt  V0CustomerLiquidationAddressNewParamsDestinationCurrency = "usdt"
)

type V0CustomerLiquidationAddressGetParams struct {
	CustomerID string `path:"customerId,required" json:"-"`
	paramObj
}
