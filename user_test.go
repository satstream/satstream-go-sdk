// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package satstream_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/satstream/satstream-go-sdk"
	"github.com/satstream/satstream-go-sdk/internal/testutil"
	"github.com/satstream/satstream-go-sdk/option"
)

func TestUserNewWithOptionalParams(t *testing.T) {
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
	err := client.User.New(context.TODO(), satstream.UserNewParams{
		User: satstream.UserParam{
			ID:         satstream.F(int64(10)),
			Email:      satstream.F("john@email.com"),
			FirstName:  satstream.F("John"),
			LastName:   satstream.F("James"),
			Password:   satstream.F("12345"),
			Phone:      satstream.F("12345"),
			Username:   satstream.F("theUser"),
			UserStatus: satstream.F(int64(1)),
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

func TestUserGet(t *testing.T) {
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
	_, err := client.User.Get(context.TODO(), "username")
	if err != nil {
		var apierr *satstream.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	err := client.User.Update(
		context.TODO(),
		"username",
		satstream.UserUpdateParams{
			User: satstream.UserParam{
				ID:         satstream.F(int64(10)),
				Email:      satstream.F("john@email.com"),
				FirstName:  satstream.F("John"),
				LastName:   satstream.F("James"),
				Password:   satstream.F("12345"),
				Phone:      satstream.F("12345"),
				Username:   satstream.F("theUser"),
				UserStatus: satstream.F(int64(1)),
			},
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

func TestUserDelete(t *testing.T) {
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
	err := client.User.Delete(context.TODO(), "username")
	if err != nil {
		var apierr *satstream.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserNewWithList(t *testing.T) {
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
	_, err := client.User.NewWithList(context.TODO(), satstream.UserNewWithListParams{
		Items: []satstream.UserParam{{
			ID:         satstream.F(int64(10)),
			Email:      satstream.F("john@email.com"),
			FirstName:  satstream.F("John"),
			LastName:   satstream.F("James"),
			Password:   satstream.F("12345"),
			Phone:      satstream.F("12345"),
			Username:   satstream.F("theUser"),
			UserStatus: satstream.F(int64(1)),
		}, {
			ID:         satstream.F(int64(10)),
			Email:      satstream.F("john@email.com"),
			FirstName:  satstream.F("John"),
			LastName:   satstream.F("James"),
			Password:   satstream.F("12345"),
			Phone:      satstream.F("12345"),
			Username:   satstream.F("theUser"),
			UserStatus: satstream.F(int64(1)),
		}, {
			ID:         satstream.F(int64(10)),
			Email:      satstream.F("john@email.com"),
			FirstName:  satstream.F("John"),
			LastName:   satstream.F("James"),
			Password:   satstream.F("12345"),
			Phone:      satstream.F("12345"),
			Username:   satstream.F("theUser"),
			UserStatus: satstream.F(int64(1)),
		}},
	})
	if err != nil {
		var apierr *satstream.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoginWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Login(context.TODO(), satstream.UserLoginParams{
		Password: satstream.F("password"),
		Username: satstream.F("username"),
	})
	if err != nil {
		var apierr *satstream.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLogout(t *testing.T) {
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
	err := client.User.Logout(context.TODO())
	if err != nil {
		var apierr *satstream.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
