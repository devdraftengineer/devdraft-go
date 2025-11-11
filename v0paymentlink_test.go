// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/devdraft-go"
	"github.com/stainless-sdks/devdraft-go/internal/testutil"
	"github.com/stainless-sdks/devdraft-go/option"
)

func TestV0PaymentLinkNewWithOptionalParams(t *testing.T) {
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
	err := client.V0.PaymentLinks.New(context.TODO(), devdraft.V0PaymentLinkNewParams{
		AllowMobilePayment:      true,
		AllowQuantityAdjustment: true,
		CollectAddress:          true,
		CollectTax:              true,
		Currency:                devdraft.V0PaymentLinkNewParamsCurrencyUsdc,
		LinkType:                devdraft.V0PaymentLinkNewParamsLinkTypeProduct,
		Title:                   "Premium Subscription",
		URL:                     "https://checkout.example.com/pay/123",
		Amount:                  devdraft.Float(29.99),
		CoverImage:              devdraft.String("https://example.com/images/premium-subscription.jpg"),
		CustomerID:              devdraft.String("123e4567-e89b-12d3-a456-426614174002"),
		CustomFields: map[string]interface{}{
			"customField1": "value1",
			"customField2": "value2",
		},
		Description:     devdraft.String("Get access to all premium features with our monthly subscription plan."),
		ExpirationDate:  devdraft.Time(time.Now()),
		IsForAllProduct: devdraft.Bool(false),
		LimitPayments:   devdraft.Bool(true),
		MaxPayments:     devdraft.Float(100),
		PaymentForID:    devdraft.String("sub_123456789"),
		PaymentLinkProducts: []devdraft.V0PaymentLinkNewParamsPaymentLinkProduct{{
			ProductID: "123e4567-e89b-12d3-a456-426614174003",
			Quantity:  1,
		}, {
			ProductID: "123e4567-e89b-12d3-a456-426614174004",
			Quantity:  2,
		}},
		TaxID: devdraft.String("123e4567-e89b-12d3-a456-426614174005"),
	})
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0PaymentLinkGet(t *testing.T) {
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
	err := client.V0.PaymentLinks.Get(context.TODO(), "id")
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0PaymentLinkUpdate(t *testing.T) {
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
	err := client.V0.PaymentLinks.Update(
		context.TODO(),
		"id",
		devdraft.V0PaymentLinkUpdateParams{},
	)
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0PaymentLinkListWithOptionalParams(t *testing.T) {
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
	err := client.V0.PaymentLinks.List(context.TODO(), devdraft.V0PaymentLinkListParams{
		Skip: devdraft.String("skip"),
		Take: devdraft.String("take"),
	})
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
