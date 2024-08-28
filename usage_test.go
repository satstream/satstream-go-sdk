// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package satstream_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/satstream-go"
	"github.com/stainless-sdks/satstream-go/internal/testutil"
	"github.com/stainless-sdks/satstream-go/option"
	"github.com/stainless-sdks/satstream-go/shared"
)

func TestUsage(t *testing.T) {
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
	order, err := client.Store.NewOrder(context.TODO(), satstream.StoreNewOrderParams{
		Order: shared.OrderParam{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", order.ID)
}
