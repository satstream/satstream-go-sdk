// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package satstream_test

import (
	"context"
	"os"
	"testing"

	"github.com/satstream/satstream-go-sdk"
	"github.com/satstream/satstream-go-sdk/internal/testutil"
	"github.com/satstream/satstream-go-sdk/option"
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
	)
	_, err := client.Fees.List(context.TODO())
	if err != nil {
		t.Error(err)
	}
}
