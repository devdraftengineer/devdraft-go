// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/devdraftengineer/devdraft-go"
	"github.com/devdraftengineer/devdraft-go/internal/testutil"
	"github.com/devdraftengineer/devdraft-go/option"
)

func TestV0InvoiceNewWithOptionalParams(t *testing.T) {
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
	err := client.V0.Invoices.New(context.TODO(), devdraft.V0InvoiceNewParams{
		CreateInvoice: devdraft.CreateInvoiceParam{
			Currency:   devdraft.CreateInvoiceCurrencyUsdc,
			CustomerID: "customer_id",
			Delivery:   devdraft.CreateInvoiceDeliveryEmail,
			DueDate:    time.Now(),
			Email:      "email",
			Items: []devdraft.CreateInvoiceItemParam{{
				ProductID: "product_id",
				Quantity:  0,
			}},
			Name:           "name",
			PartialPayment: true,
			PaymentLink:    true,
			PaymentMethods: []string{"ACH"},
			Status:         devdraft.CreateInvoiceStatusDraft,
			Address:        devdraft.String("address"),
			Logo:           devdraft.String("logo"),
			PhoneNumber:    devdraft.String("phone_number"),
			SendDate:       devdraft.Time(time.Now()),
			TaxID:          devdraft.String("taxId"),
		},
	})
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0InvoiceGet(t *testing.T) {
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
	err := client.V0.Invoices.Get(context.TODO(), "id")
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV0InvoiceUpdateWithOptionalParams(t *testing.T) {
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
	err := client.V0.Invoices.Update(
		context.TODO(),
		"id",
		devdraft.V0InvoiceUpdateParams{
			CreateInvoice: devdraft.CreateInvoiceParam{
				Currency:   devdraft.CreateInvoiceCurrencyUsdc,
				CustomerID: "customer_id",
				Delivery:   devdraft.CreateInvoiceDeliveryEmail,
				DueDate:    time.Now(),
				Email:      "email",
				Items: []devdraft.CreateInvoiceItemParam{{
					ProductID: "product_id",
					Quantity:  0,
				}},
				Name:           "name",
				PartialPayment: true,
				PaymentLink:    true,
				PaymentMethods: []string{"ACH"},
				Status:         devdraft.CreateInvoiceStatusDraft,
				Address:        devdraft.String("address"),
				Logo:           devdraft.String("logo"),
				PhoneNumber:    devdraft.String("phone_number"),
				SendDate:       devdraft.Time(time.Now()),
				TaxID:          devdraft.String("taxId"),
			},
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

func TestV0InvoiceListWithOptionalParams(t *testing.T) {
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
	err := client.V0.Invoices.List(context.TODO(), devdraft.V0InvoiceListParams{
		Skip: devdraft.Float(0),
		Take: devdraft.Float(0),
	})
	if err != nil {
		var apierr *devdraft.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
