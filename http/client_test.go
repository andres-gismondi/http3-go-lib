package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	http2 "net/http"
	"net/http/httptest"
	"testing"
	"time"

	"form3/http"
	"form3/model"
)

func TestClient_success(t *testing.T) {
	client := http.NewClient("111",
		"account",
		"abc-123",
		http.Timeout(3*time.Second))

	ctx := context.Background()
	server := httptest.NewServer(http2.HandlerFunc(func(response http2.ResponseWriter, request *http2.Request) {
		response.WriteHeader(http2.StatusOK)

		account := model.AccountData{}
		response.Write([]byte(fmt.Sprintf("%v", account)))
	}))
	defer server.Close()

	client.
}
