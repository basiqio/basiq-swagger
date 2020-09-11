package accounts

import (
	"context"
	"encoding/json"
	"github.com/basiqio/basiq-swagger/dist/client/accounts"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
	"testing"
)

func TestGetAccounts(t *testing.T) {
	accParams := &accounts.GetAccountsParams{
		UserID:  "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
		Context: context.TODO(),
	}

	accRsp, err := test.Client.Accounts.GetAccounts(accParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error getting accounts: %v", err)
	}

	e, err := json.Marshal(accRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getAccounts.json", t)

	test.AssertJson(t, s, string(e))
}

func TestGetAccount(t *testing.T) {
	accParams := &accounts.GetAccountParams{
		UserID:    "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
		AccountID: "f58d4d6a-61c6-4893-96da-249fe09eb2ad",
		Context:   context.TODO(),
	}

	connRsp, err := test.Client.Accounts.GetAccount(accParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error getting account: %v", err)
	}

	e, err := json.Marshal(connRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getAccount.json", t)

	test.AssertJson(t, s, string(e))
}

func TestGetAccountsEmpty(t *testing.T) {
	accParams := &accounts.GetAccountsParams{
		UserID:  "dd41fa66-32fc-4e01-8a1f-3b9402414071",
		Context: context.TODO(),
	}

	accRsp, err := test.Client.Accounts.GetAccounts(accParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error getting accounts: %v", err)
	}

	e, err := json.Marshal(accRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getAccountsEmpty.json", t)

	test.AssertJson(t, s, string(e))
}
