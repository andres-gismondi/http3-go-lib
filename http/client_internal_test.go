//go:build integration
// +build integration

package http_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"form3/http"
)

func TestClientInternal_Create_success(t *testing.T) {
	client := http.NewClient(
		"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		"accounts",
		"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		http.BaseURL("http://localhost:8080"))
	got, err := client.Create(context.Background(), []string{"andres"}, "AR")

	assert.Nil(t, err)
	assert.NotNil(t, got)

	want := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	assert.Equal(t, want, got.ID)
}

func TestClientInternal_Fetch_success(t *testing.T) {
	ctx := context.Background()
	client := http.NewClient(
		"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		"accounts",
		"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		http.BaseURL("http://localhost:8080"))

	wanted, err := client.Create(ctx, []string{"andres"}, "AR")
	assert.Nil(t, err)

	bankId := wanted.Attributes.BankID
	got, err := client.Fetch(ctx, bankId)
	assert.Nil(t, err)

	want := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	assert.Equal(t, want, got.ID)
}

func TestClientInternal_Delete_success(t *testing.T) {
	ctx := context.Background()
	client := http.NewClient(
		"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		"accounts",
		"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		http.BaseURL("http://localhost:8080"))

	wanted, err := client.Create(ctx, []string{"andres"}, "AR")
	assert.Nil(t, err)

	bankId := wanted.Attributes.BankID
	got, err := client.Fetch(ctx, bankId)
	assert.Nil(t, err)
	id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	assert.Equal(t, id, got.ID)

	version := wanted.Version
	err = client.Delete(ctx, bankId, int(version))
	assert.Nil(t, err)

	got, err = client.Fetch(ctx, bankId)
	assert.Contains(t, err.Error(), "not found")
}
