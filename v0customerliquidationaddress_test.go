// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/devdraft-go"
	"github.com/stainless-sdks/devdraft-go/internal/testutil"
	"github.com/stainless-sdks/devdraft-go/option"
)

func TestV0CustomerLiquidationAddressNewWithOptionalParams(t *testing.T) {
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
	_, err := client.V0.Customers.LiquidationAddresses.New(
		context.TODO(),
		"cust_123",
		devdraft.V0CustomerLiquidationAddressNewParams{
			Address:                   "0x4d0280da2f2fDA5103914bCc5aad114743152A9c",
			Chain:                     devdraft.V0CustomerLiquidationAddressNewParamsChainEthereum,
			Currency:                  devdraft.V0CustomerLiquidationAddressNewParamsCurrencyUsdc,
			BridgeWalletID:            devdraft.String("bw_123"),
			CustomDeveloperFeePercent: devdraft.String("2.5"),
			DestinationACHReference:   devdraft.String("ACH123"),
			DestinationAddress:        devdraft.String("0x1234567890abcdef1234567890abcdef12345678"),
			DestinationBlockchainMemo: devdraft.String("liquidation-memo-123"),
			DestinationCurrency:       devdraft.V0CustomerLiquidationAddressNewParamsDestinationCurrencyEur,
			DestinationPaymentRail:    devdraft.BridgePaymentRailSepa,
			DestinationSepaReference:  devdraft.String("SEPA-REF-123456"),
			DestinationWireMessage:    devdraft.String("Liquidation payment for customer"),
			ExternalAccountID:         devdraft.String("ext_123"),
			PrefundedAccountID:        devdraft.String("pf_acc_123"),
			ReturnAddress:             devdraft.String("0xabcdefabcdefabcdefabcdefabcdefabcdef"),
		},
	)
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0CustomerLiquidationAddressGet(t *testing.T) {
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
	_, err := client.V0.Customers.LiquidationAddresses.Get(
		context.TODO(),
		"la_generated_id_123",
		devdraft.V0CustomerLiquidationAddressGetParams{
			CustomerID: "cust_123",
		},
	)
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0CustomerLiquidationAddressList(t *testing.T) {
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
	_, err := client.V0.Customers.LiquidationAddresses.List(context.TODO(), "cust_123")
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
