// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"context"
	"net/http"
	"slices"

	"github.com/devdraftengineer/devdraft-go/internal/apijson"
	"github.com/devdraftengineer/devdraft-go/internal/requestconfig"
	"github.com/devdraftengineer/devdraft-go/option"
	"github.com/devdraftengineer/devdraft-go/packages/param"
)

// V0TransferService contains methods and other services that help with interacting
// with the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0TransferService] method instead.
type V0TransferService struct {
	Options []option.RequestOption
}

// NewV0TransferService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV0TransferService(opts ...option.RequestOption) (r V0TransferService) {
	r = V0TransferService{}
	r.Options = opts
	return
}

// Create a direct bank transfer
func (r *V0TransferService) NewDirectBank(ctx context.Context, body V0TransferNewDirectBankParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/transfers/direct-bank"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Create a direct wallet transfer
func (r *V0TransferService) NewDirectWallet(ctx context.Context, body V0TransferNewDirectWalletParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/transfers/direct-wallet"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Create an external bank transfer
func (r *V0TransferService) NewExternalBankTransfer(ctx context.Context, body V0TransferNewExternalBankTransferParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/transfers/external-bank-transfer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Create an external stablecoin transfer
func (r *V0TransferService) NewExternalStablecoinTransfer(ctx context.Context, body V0TransferNewExternalStablecoinTransferParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/transfers/external-stablecoin-transfer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Create a stablecoin conversion
func (r *V0TransferService) NewStablecoinConversion(ctx context.Context, body V0TransferNewStablecoinConversionParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/transfers/stablecoin-conversion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type V0TransferNewDirectBankParams struct {
	// The amount to transfer
	Amount float64 `json:"amount,required"`
	// The destination currency
	DestinationCurrency string `json:"destinationCurrency,required"`
	// The payment rail to use
	PaymentRail string `json:"paymentRail,required"`
	// The source currency
	SourceCurrency string `json:"sourceCurrency,required"`
	// The ID of the bridge wallet to transfer from
	WalletID string `json:"walletId,required"`
	// ACH transfer reference
	ACHReference param.Opt[string] `json:"ach_reference,omitzero"`
	// SEPA transfer reference
	SepaReference param.Opt[string] `json:"sepa_reference,omitzero"`
	// Wire transfer message
	WireMessage param.Opt[string] `json:"wire_message,omitzero"`
	paramObj
}

func (r V0TransferNewDirectBankParams) MarshalJSON() (data []byte, err error) {
	type shadow V0TransferNewDirectBankParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0TransferNewDirectBankParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0TransferNewDirectWalletParams struct {
	// The amount to transfer
	Amount float64 `json:"amount,required"`
	// The network to use for the transfer
	Network string `json:"network,required"`
	// The stablecoin currency to use
	StableCoinCurrency string `json:"stableCoinCurrency,required"`
	// The ID of the bridge wallet to transfer from
	WalletID string `json:"walletId,required"`
	paramObj
}

func (r V0TransferNewDirectWalletParams) MarshalJSON() (data []byte, err error) {
	type shadow V0TransferNewDirectWalletParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0TransferNewDirectWalletParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0TransferNewExternalBankTransferParams struct {
	// The destination currency
	DestinationCurrency string `json:"destinationCurrency,required"`
	// The destination payment rail (fiat payment method)
	//
	// Any of "ach", "ach_push", "ach_same_day", "wire", "sepa", "swift", "spei".
	DestinationPaymentRail V0TransferNewExternalBankTransferParamsDestinationPaymentRail `json:"destinationPaymentRail,omitzero,required"`
	// The external account ID for the bank transfer
	ExternalAccountID string `json:"external_account_id,required"`
	// The source currency
	SourceCurrency string `json:"sourceCurrency,required"`
	// The ID of the source bridge wallet
	SourceWalletID string `json:"sourceWalletId,required"`
	// ACH reference message (1-10 characters, only for ACH transfers)
	ACHReference param.Opt[string] `json:"ach_reference,omitzero"`
	// The amount to transfer
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// SEPA reference message (6-140 characters, only for SEPA transfers)
	SepaReference param.Opt[string] `json:"sepa_reference,omitzero"`
	// SPEI reference message (1-40 characters, only for SPEI transfers)
	SpeiReference param.Opt[string] `json:"spei_reference,omitzero"`
	// SWIFT charges bearer (only for SWIFT transfers)
	SwiftCharges param.Opt[string] `json:"swift_charges,omitzero"`
	// SWIFT reference message (1-190 characters, only for SWIFT transfers)
	SwiftReference param.Opt[string] `json:"swift_reference,omitzero"`
	// Wire transfer message (1-256 characters, only for wire transfers)
	WireMessage param.Opt[string] `json:"wire_message,omitzero"`
	paramObj
}

func (r V0TransferNewExternalBankTransferParams) MarshalJSON() (data []byte, err error) {
	type shadow V0TransferNewExternalBankTransferParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0TransferNewExternalBankTransferParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The destination payment rail (fiat payment method)
type V0TransferNewExternalBankTransferParamsDestinationPaymentRail string

const (
	V0TransferNewExternalBankTransferParamsDestinationPaymentRailACH        V0TransferNewExternalBankTransferParamsDestinationPaymentRail = "ach"
	V0TransferNewExternalBankTransferParamsDestinationPaymentRailACHPush    V0TransferNewExternalBankTransferParamsDestinationPaymentRail = "ach_push"
	V0TransferNewExternalBankTransferParamsDestinationPaymentRailACHSameDay V0TransferNewExternalBankTransferParamsDestinationPaymentRail = "ach_same_day"
	V0TransferNewExternalBankTransferParamsDestinationPaymentRailWire       V0TransferNewExternalBankTransferParamsDestinationPaymentRail = "wire"
	V0TransferNewExternalBankTransferParamsDestinationPaymentRailSepa       V0TransferNewExternalBankTransferParamsDestinationPaymentRail = "sepa"
	V0TransferNewExternalBankTransferParamsDestinationPaymentRailSwift      V0TransferNewExternalBankTransferParamsDestinationPaymentRail = "swift"
	V0TransferNewExternalBankTransferParamsDestinationPaymentRailSpei       V0TransferNewExternalBankTransferParamsDestinationPaymentRail = "spei"
)

type V0TransferNewExternalStablecoinTransferParams struct {
	// Beneficiary ID for the stablecoin transfer
	BeneficiaryID string `json:"beneficiaryId,required"`
	// The destination currency
	DestinationCurrency string `json:"destinationCurrency,required"`
	// The source currency
	SourceCurrency string `json:"sourceCurrency,required"`
	// The ID of the source bridge wallet
	SourceWalletID string `json:"sourceWalletId,required"`
	// The amount to transfer
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// Blockchain memo for the transfer
	BlockchainMemo param.Opt[string] `json:"blockchain_memo,omitzero"`
	paramObj
}

func (r V0TransferNewExternalStablecoinTransferParams) MarshalJSON() (data []byte, err error) {
	type shadow V0TransferNewExternalStablecoinTransferParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0TransferNewExternalStablecoinTransferParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0TransferNewStablecoinConversionParams struct {
	// The amount to convert
	Amount float64 `json:"amount,required"`
	// The destination currency
	DestinationCurrency string `json:"destinationCurrency,required"`
	// The source currency
	SourceCurrency string `json:"sourceCurrency,required"`
	// The source network
	SourceNetwork string `json:"sourceNetwork,required"`
	// The ID of the bridge wallet to use
	WalletID string `json:"walletId,required"`
	paramObj
}

func (r V0TransferNewStablecoinConversionParams) MarshalJSON() (data []byte, err error) {
	type shadow V0TransferNewStablecoinConversionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0TransferNewStablecoinConversionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
