package http3_go_lib

import (
	"context"

	"form3/model"
)

// Client implement a wrapper to Form3 to manage different actions from accounts
type Client interface {

	// Create register an existing bank or create a new one
	// This method can be provided by account options depending on required fields.
	// see link below:
	//https://api-docs.form3.tech/api.html#organisation-accounts
	Create(ctx context.Context, name []string, country string, options ...AccountOption) (model.AccountData, error)

	// Fetch Get a specific account by ID
	Fetch(ctx context.Context, id string) (model.AccountData, error)

	// Delete an existing account by ID
	Delete(ctx context.Context, accountID string, version int) error
}
