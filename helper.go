package tfe

import (
	"context"
	"fmt"
)

// TestAccountDetails represents the basic account information
// of a TFE/TFC user.
//
// See FetcheTestAccountDetails for more information.
type TestAccountDetails struct {
	ID       string `json:"id" jsonapi:"primary,users"`
	Username string `jsonapi:"attr,username"`
	Email    string `jsonapi:"attr,email"`
}

// FetchTestAccountDetails returns TestAccountDetails
// of the user running the tests.
//
// Use this helper to fetch the username and email
// address associated with the token used to run the tests.
func FetchTestAccountDetails(client *Client) (*TestAccountDetails, error) {
	tad := &TestAccountDetails{}
	req, err := client.newRequest("GET", "account/details", nil)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	err = client.do(ctx, req, tad)
	if err != nil {
		return nil, fmt.Errorf("could not fetch test user details: %v", err)
	}
	return tad, nil
}
