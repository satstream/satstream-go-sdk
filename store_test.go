// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package satstream_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/satstream/satstream-go-sdk"
	"github.com/satstream/satstream-go-sdk/internal/testutil"
	"github.com/satstream/satstream-go-sdk/option"
	"github.com/satstream/satstream-go-sdk/shared"
)

func TestStoreNewOrderWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := satstream.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Store.NewOrder(context.TODO(), satstream.StoreNewOrderParams{
		Order: shared.OrderParam{
			ID:       satstream.F(int64(10)),
			Complete: satstream.F(true),
			PetID:    satstream.F(int64(198772)),
			Quantity: satstream.F(int64(7)),
			ShipDate: satstream.F(time.Now()),
			Status:   satstream.F(shared.OrderStatusPlaced),
		},
	})
	if err != nil {
		var apierr *satstream.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoreInventory(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := satstream.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Store.Inventory(context.TODO())
	if err != nil {
		var apierr *satstream.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
