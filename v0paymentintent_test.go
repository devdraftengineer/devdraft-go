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

func TestV0PaymentIntentNewBankWithOptionalParams(t *testing.T) {
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
	err := client.V0.PaymentIntents.NewBank(context.TODO(), devdraft.V0PaymentIntentNewBankParams{
		DestinationCurrency: devdraft.StableCoinCurrencyUsdc,
		DestinationNetwork:  devdraft.BridgePaymentRailEthereum,
		SourceCurrency:      devdraft.V0PaymentIntentNewBankParamsSourceCurrencyUsd,
		SourcePaymentRail:   devdraft.BridgePaymentRailACHPush,
		ACHReference:        devdraft.String("INV12345"),
		Amount:              devdraft.String("1000.00"),
		CustomerAddress:     devdraft.String("123 Main St, New York, NY 10001"),
		CustomerCountry:     devdraft.String("United States"),
		CustomerCountryISO:  devdraft.String("US"),
		CustomerEmail:       devdraft.String("john.doe@example.com"),
		CustomerFirstName:   devdraft.String("John"),
		CustomerLastName:    devdraft.String("Doe"),
		CustomerProvince:    devdraft.String("New York"),
		CustomerProvinceISO: devdraft.String("NY"),
		DestinationAddress:  devdraft.String("0x742d35Cc6634C0532925a3b8D4C9db96c4b4d8e1"),
		PhoneNumber:         devdraft.String("+1-555-123-4567"),
		SepaReference:       devdraft.String("REF-123456789"),
		WireMessage:         devdraft.String("Payment for invoice #12345"),
	})
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0PaymentIntentNewStableWithOptionalParams(t *testing.T) {
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
	err := client.V0.PaymentIntents.NewStable(context.TODO(), devdraft.V0PaymentIntentNewStableParams{
		DestinationNetwork:  devdraft.BridgePaymentRailPolygon,
		SourceCurrency:      devdraft.StableCoinCurrencyUsdc,
		SourceNetwork:       devdraft.BridgePaymentRailEthereum,
		Amount:              devdraft.String("100.00"),
		CustomerAddress:     devdraft.String("123 Main St, New York, NY 10001"),
		CustomerCountry:     devdraft.String("United States"),
		CustomerCountryISO:  devdraft.String("US"),
		CustomerEmail:       devdraft.String("john.doe@example.com"),
		CustomerFirstName:   devdraft.String("John"),
		CustomerLastName:    devdraft.String("Doe"),
		CustomerProvince:    devdraft.String("New York"),
		CustomerProvinceISO: devdraft.String("NY"),
		DestinationAddress:  devdraft.String("0x742d35Cc6634C0532925a3b8D4C9db96c4b4d8e1"),
		DestinationCurrency: devdraft.StableCoinCurrencyEurc,
		PhoneNumber:         devdraft.String("+1-555-123-4567"),
	})
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
