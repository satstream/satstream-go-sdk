// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package satstream_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/satstream-go"
	"github.com/stainless-sdks/satstream-go/internal/testutil"
	"github.com/stainless-sdks/satstream-go/option"
)

func TestPetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.New(context.TODO(), satstream.PetNewParams{
		Pet: satstream.PetParam{
			Name:      satstream.F("doggie"),
			PhotoURLs: satstream.F([]string{"string", "string", "string"}),
			ID:        satstream.F(int64(10)),
			Category: satstream.F(satstream.PetCategoryParam{
				ID:   satstream.F(int64(1)),
				Name: satstream.F("Dogs"),
			}),
			Status: satstream.F(satstream.PetStatusAvailable),
			Tags: satstream.F([]satstream.PetTagParam{{
				ID:   satstream.F(int64(0)),
				Name: satstream.F("name"),
			}, {
				ID:   satstream.F(int64(0)),
				Name: satstream.F("name"),
			}, {
				ID:   satstream.F(int64(0)),
				Name: satstream.F("name"),
			}}),
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

func TestPetGet(t *testing.T) {
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
	_, err := client.Pets.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *satstream.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.Update(context.TODO(), satstream.PetUpdateParams{
		Pet: satstream.PetParam{
			Name:      satstream.F("doggie"),
			PhotoURLs: satstream.F([]string{"string", "string", "string"}),
			ID:        satstream.F(int64(10)),
			Category: satstream.F(satstream.PetCategoryParam{
				ID:   satstream.F(int64(1)),
				Name: satstream.F("Dogs"),
			}),
			Status: satstream.F(satstream.PetStatusAvailable),
			Tags: satstream.F([]satstream.PetTagParam{{
				ID:   satstream.F(int64(0)),
				Name: satstream.F("name"),
			}, {
				ID:   satstream.F(int64(0)),
				Name: satstream.F("name"),
			}, {
				ID:   satstream.F(int64(0)),
				Name: satstream.F("name"),
			}}),
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

func TestPetDelete(t *testing.T) {
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
	err := client.Pets.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *satstream.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByStatus(context.TODO(), satstream.PetFindByStatusParams{
		Status: satstream.F(satstream.PetFindByStatusParamsStatusAvailable),
	})
	if err != nil {
		var apierr *satstream.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByTagsWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByTags(context.TODO(), satstream.PetFindByTagsParams{
		Tags: satstream.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *satstream.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateByIDWithOptionalParams(t *testing.T) {
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
	err := client.Pets.UpdateByID(
		context.TODO(),
		int64(0),
		satstream.PetUpdateByIDParams{
			Name:   satstream.F("name"),
			Status: satstream.F("status"),
		},
	)
	if err != nil {
		var apierr *satstream.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUploadImageWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.UploadImage(
		context.TODO(),
		int64(0),
		satstream.PetUploadImageParams{
			Image:              io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			AdditionalMetadata: satstream.F("additionalMetadata"),
		},
	)
	if err != nil {
		var apierr *satstream.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
