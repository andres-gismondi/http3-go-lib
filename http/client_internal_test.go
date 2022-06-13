package http_test

import (
	"context"
	"fmt"
	"testing"

	"form3/http"
)

func TestClientInternal_Create_success(t *testing.T) {
	client := http.NewClient("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", "accounts", "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	res, err := client.Create(context.Background(), []string{"andres"}, "AR")
	if err != nil {
		fmt.Printf("error: [%v]", err)
	}

	fmt.Printf("Data: [%v]", res)
}
