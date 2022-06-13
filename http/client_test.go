package http_test

import (
	"context"
	"encoding/json"
	http2 "net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	http3_go_lib "form3"
	"form3/http"
	"form3/model"
)

func TestClient_Create_success(t *testing.T) {
	want := model.AccountData{
		ID:             "111",
		OrganisationID: "abc-123",
		Type:           "account",
		Attributes: model.AccountAttributes{
			Name:    []string{"andres"},
			BankID:  "123-321",
			Country: "AR",
		},
	}

	ctx := context.Background()
	server := httptest.NewServer(http2.HandlerFunc(func(response http2.ResponseWriter, request *http2.Request) {
		response.WriteHeader(http2.StatusOK)

		json.NewEncoder(response).Encode(want)
	}))
	defer server.Close()

	client := http.NewClient("111",
		"account",
		"abc-123",
		http.BaseURL(server.URL),
		http.Timeout(3*time.Second))

	clientOptions := []http3_go_lib.AccountOption{
		http3_go_lib.BankID("123-321"),
	}
	got, err := client.Create(ctx, []string{"andres"}, "AR", clientOptions...)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
