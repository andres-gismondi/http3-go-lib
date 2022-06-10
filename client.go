package http3_go_lib

import (
	"context"

	"form3/model"
)

type Client interface {

	// Register an existing bank or Create a new one
	Create(ctx context.Context, name []string, country string, options ...AccountOption) (model.AccountData, error)

	Fetch(ctx context.Context, id string) (model.AccountData, error)

	// Delete an existing bank
	Delete(ctx context.Context, accountID string, version int) error
}
