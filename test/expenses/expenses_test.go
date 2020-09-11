package expenses

import (
	"context"
	"encoding/json"
	"github.com/basiqio/basiq-swagger/dist/client/expenses"
	"github.com/basiqio/basiq-swagger/dist/models"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
	"strings"
	"testing"
)

func TestGetExpenses(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5"
	snapshotID := "55275323-6dec-4458-938e-9f7d21edba00"

	expensesGetParams := &expenses.GetExpensesParams{UserID: userID,
		SnapshotID: snapshotID,
		Context:    context.TODO(),
	}

	expensesRsp, err := test.Client.Expenses.GetExpenses(expensesGetParams, token)

	if err != nil {
		t.Fatalf("Cannot get expenses. Error: %v", err)
	}

	e, err := json.Marshal(expensesRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getExpenses.json", t)

	test.AssertJson(t, s, string(e))
}

func TestPostExpenses(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5"
	accountID := "24c742dc-bba6-4bc2-8298-fe731a17be68"
	expensesPostParams := &expenses.PostExpensesParams{UserID: userID,
		ExpensesPostRequest: &models.ExpensesPost{Accounts: []string{accountID}, FromMonth: "2019-05", ToMonth: "2019-09"},
		Context:             context.TODO(),
	}

	expensesRsp, err := test.Client.Expenses.PostExpenses(expensesPostParams, token)

	if err != nil {
		t.Fatalf("Cannot post expenses. Error: %v", err)
	}

	e, err := json.Marshal(expensesRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/postExpenses.json", t)

	test.AssertJson(t, s, string(e))
}

func TestPostExpensesBadRequest(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5"
	accountID := "14aa946c-6f87-44b7-bab9-d192d6261471"
	expensesPostParams := &expenses.PostExpensesParams{UserID: userID,
		ExpensesPostRequest: &models.ExpensesPost{FromMonth: "2015-11", ToMonth: "2020-10", Accounts: []string{accountID}},
		Context:             context.TODO(),
	}

	_, err := test.Client.Expenses.PostExpenses(expensesPostParams, token)

	if err != nil {
		badRequest, ok := err.(*expenses.PostExpensesBadRequest)
		if !ok {
			t.Fatal("Bad request error not returned.")
		}

		s := strings.Replace(
			test.GetJsonResponse("./responses/badRequest.json", t), "fbf82ac8-be3d-463f-b502-5f0749637ae1",
			*badRequest.GetPayload().CorrelationID,
			1)
		e, err := json.Marshal(badRequest.GetPayload())
		if err != nil {
			t.Fatalf("Cannot marshal error model. Error: %v", err)
		}

		test.AssertJson(t, s, string(e))
	} else {
		t.Fatalf("Error should be returned.")
	}
}