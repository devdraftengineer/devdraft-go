// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/devdraftengineer/devdraft-go"
	"github.com/devdraftengineer/devdraft-go/internal/testutil"
	"github.com/devdraftengineer/devdraft-go/option"
)

func TestV0TransferNewDirectBankWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := devdraft.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSecret("My Secret"),
		option.WithIdempotencyKey("My Idempotency Key"),
	)
	err := client.V0.Transfers.NewDirectBank(context.TODO(), devdraft.V0TransferNewDirectBankParams{
		Amount:              0,
		DestinationCurrency: "destinationCurrency",
		PaymentRail:         "paymentRail",
		SourceCurrency:      "sourceCurrency",
		WalletID:            "walletId",
		ACHReference:        devdraft.String("ach_reference"),
		SepaReference:       devdraft.String("sepa_reference"),
		WireMessage:         devdraft.String("wire_message"),
	})
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0TransferNewDirectWallet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := devdraft.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSecret("My Secret"),
		option.WithIdempotencyKey("My Idempotency Key"),
	)
	err := client.V0.Transfers.NewDirectWallet(context.TODO(), devdraft.V0TransferNewDirectWalletParams{
		Amount:             0,
		Network:            "network",
		StableCoinCurrency: "stableCoinCurrency",
		WalletID:           "walletId",
	})
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0TransferNewExternalBankTransferWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := devdraft.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSecret("My Secret"),
		option.WithIdempotencyKey("My Idempotency Key"),
	)
	err := client.V0.Transfers.NewExternalBankTransfer(context.TODO(), devdraft.V0TransferNewExternalBankTransferParams{
		DestinationCurrency:    "destinationCurrency",
		DestinationPaymentRail: devdraft.V0TransferNewExternalBankTransferParamsDestinationPaymentRailACH,
		ExternalAccountID:      "external_account_id",
		SourceCurrency:         "sourceCurrency",
		SourceWalletID:         "sourceWalletId",
		ACHReference:           devdraft.String("x"),
		Amount:                 devdraft.Float(0),
		SepaReference:          devdraft.String("xxxxxx"),
		SpeiReference:          devdraft.String("x"),
		SwiftCharges:           devdraft.String("swift_charges"),
		SwiftReference:         devdraft.String("x"),
		WireMessage:            devdraft.String("x"),
	})
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0TransferNewExternalStablecoinTransferWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := devdraft.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSecret("My Secret"),
		option.WithIdempotencyKey("My Idempotency Key"),
	)
	err := client.V0.Transfers.NewExternalStablecoinTransfer(context.TODO(), devdraft.V0TransferNewExternalStablecoinTransferParams{
		BeneficiaryID:       "beneficiary_12345",
		DestinationCurrency: "destinationCurrency",
		SourceCurrency:      "sourceCurrency",
		SourceWalletID:      "sourceWalletId",
		Amount:              devdraft.Float(0),
		BlockchainMemo:      devdraft.String("blockchain_memo"),
	})
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0TransferNewStablecoinConversion(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := devdraft.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSecret("My Secret"),
		option.WithIdempotencyKey("My Idempotency Key"),
	)
	err := client.V0.Transfers.NewStablecoinConversion(context.TODO(), devdraft.V0TransferNewStablecoinConversionParams{
		Amount:              0,
		DestinationCurrency: "destinationCurrency",
		SourceCurrency:      "sourceCurrency",
		SourceNetwork:       "sourceNetwork",
		WalletID:            "walletId",
	})
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
