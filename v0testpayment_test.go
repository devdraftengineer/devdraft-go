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

func TestV0TestPaymentGet(t *testing.T) {
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
	_, err := client.V0.TestPayment.Get(context.TODO(), "id")
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0TestPaymentProcessWithOptionalParams(t *testing.T) {
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
	_, err := client.V0.TestPayment.Process(context.TODO(), devdraft.V0TestPaymentProcessParams{
		Amount:      100.5,
		Currency:    "USD",
		Description: "Test payment for API",
		CustomerID:  devdraft.String("cus_12345"),
	})
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0TestPaymentRefund(t *testing.T) {
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
	_, err := client.V0.TestPayment.Refund(context.TODO(), "id")
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
