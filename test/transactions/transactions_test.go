package transactions

import (
	"context"
	"encoding/json"
	"github.com/basiqio/basiq-swagger/dist/client/transactions"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
	"testing"
)

func TestGetTransactions(t *testing.T) {
	limit := int64(1)
	txParams := &transactions.GetTransactionsParams{
		UserID:  "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
		Limit:   &limit,
		Context: context.TODO(),
	}

	txRsp, err := test.Client.Transactions.GetTransactions(txParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error getting transactions: %v", err)
	}

	e, err := json.Marshal(txRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getTransactions.json", t)

	test.AssertJson(t, s, string(e))
}

func TestGetTransaction(t *testing.T) {
	txParams := &transactions.GetTransactionParams{
		UserID:        "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
		TransactionID: "bf3cb262-ccf0-4ad3-aaf3-20395fa1d75b",
		Context:       context.TODO(),
	}

	txRsp, err := test.Client.Transactions.GetTransaction(txParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error getting transaction: %v", err)
	}

	e, err := json.Marshal(txRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getTransaction.json", t)

	test.AssertJson(t, s, string(e))
}
func TestGetTransactionsEmpty(t *testing.T) {
	txParams := &transactions.GetTransactionsParams{
		UserID:  "dd41fa66-32fc-4e01-8a1f-3b9402414071",
		Context: context.TODO(),
	}

	txRsp, err := test.Client.Transactions.GetTransactions(txParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error getting transactions: %v", err)
	}

	e, err := json.Marshal(txRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getTransactionsEmpty.json", t)

	test.AssertJson(t, s, string(e))
}
