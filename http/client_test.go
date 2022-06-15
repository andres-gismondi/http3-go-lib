package http_test

import (
	"context"
	"encoding/json"
	"fmt"
	http2 "net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	http3_go_lib "form3"
	"form3/http"
	"form3/model"
)

func TestClient_Create(t *testing.T) {
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

	type fields struct {
		http2.HandlerFunc
	}

	tests := []struct {
		name string
		fields
		want model.AccountData
	}{
		{
			name: "success",
			fields: fields{
				http2.HandlerFunc(func(response http2.ResponseWriter, request *http2.Request) {
					response.WriteHeader(http2.StatusOK)
					json.NewEncoder(response).Encode(want)
				}),
			},
			want: want,
		},
		{
			name: "error decoding response",
			fields: fields{
				http2.HandlerFunc(func(response http2.ResponseWriter, request *http2.Request) {
					response.WriteHeader(http2.StatusOK)
					json.NewEncoder(response).Encode("{")
				}),
			},
			want: model.AccountData{},
		},
		{
			name: "response api error",
			fields: fields{
				http2.HandlerFunc(func(response http2.ResponseWriter, request *http2.Request) {
					response.WriteHeader(http2.StatusBadRequest)
				}),
			},
			want: model.AccountData{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			server := httptest.NewServer(tt.fields.HandlerFunc)
			defer server.Close()

			client := http.NewClient("111",
				"account",
				"abc-123",
				http.BaseURL(server.URL),
				http.Timeout(3*time.Second))

			clientOptions := []http3_go_lib.AccountOption{
				http3_go_lib.BankID("123-321"),
			}
			got, _ := client.Create(ctx, []string{"andres"}, "AR", clientOptions...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_Fetch(t *testing.T) {
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

	type fields struct {
		http2.HandlerFunc
	}

	tests := []struct {
		name string
		fields
		want model.AccountData
	}{
		{
			name: "success",
			fields: fields{
				http2.HandlerFunc(func(response http2.ResponseWriter, request *http2.Request) {
					response.WriteHeader(http2.StatusOK)
					json.NewEncoder(response).Encode(want)
				}),
			},
			want: want,
		},
		{
			name: "error decoding response",
			fields: fields{
				http2.HandlerFunc(func(response http2.ResponseWriter, request *http2.Request) {
					response.WriteHeader(http2.StatusOK)
					json.NewEncoder(response).Encode("{")
				}),
			},
			want: model.AccountData{},
		},
		{
			name: "response api error",
			fields: fields{
				http2.HandlerFunc(func(response http2.ResponseWriter, request *http2.Request) {
					response.WriteHeader(http2.StatusBadRequest)
				}),
			},
			want: model.AccountData{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			server := httptest.NewServer(tt.fields.HandlerFunc)
			defer server.Close()

			client := http.NewClient("111",
				"account",
				"abc-123",
				http.BaseURL(server.URL),
				http.Timeout(3*time.Second))

			got, _ := client.Fetch(ctx, "123-321")
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_Delete(t *testing.T) {
	type fields struct {
		http2.HandlerFunc
	}

	tests := []struct {
		name string
		fields
		want error
	}{
		{
			name: "success",
			fields: fields{
				http2.HandlerFunc(func(response http2.ResponseWriter, request *http2.Request) {
					response.WriteHeader(http2.StatusOK)
				}),
			},
			want: nil,
		},
		{
			name: "error deleting account",
			fields: fields{
				http2.HandlerFunc(func(response http2.ResponseWriter, request *http2.Request) {
					response.WriteHeader(http2.StatusBadRequest)
				}),
			},
			want: fmt.Errorf("account api error: %s", "400 Bad Request"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			server := httptest.NewServer(tt.fields.HandlerFunc)
			defer server.Close()

			client := http.NewClient("111",
				"account",
				"abc-123",
				http.BaseURL(server.URL),
				http.Timeout(3*time.Second))

			err := client.Delete(ctx, "123-321", 0)
			assert.Equal(t, tt.want, err)
		})
	}
}
