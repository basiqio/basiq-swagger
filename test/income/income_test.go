package income

import (
	"context"
	"encoding/json"
	"github.com/basiqio/basiq-swagger/dist/client/income"
	"github.com/basiqio/basiq-swagger/dist/models"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
	"strings"
	"testing"
)

func TestGetIncome(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5"
	snapshotID := "46137f25-af96-40d0-883f-16fd871b7f23"

	incomeGetParams := &income.GetIncomeParams{UserID: userID,
		SnapshotID: snapshotID,
		Context:    context.TODO(),
	}

	incomeRsp, err := test.Client.Income.GetIncome(incomeGetParams, token)

	if err != nil {
		t.Fatalf("Cannot get income. Error: %v", err)
	}

	e, err := json.Marshal(incomeRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getIncome.json", t)
	s = strings.Replace(s, "46137f25-af96-40d0-883f-16fd871b7f23", *incomeRsp.GetPayload().ID, 1)

	test.AssertJson(t, s, string(e))
}

func TestPostIncome(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5"
	accountID := "14aa946c-6f87-44b7-bab9-d192d6261471"
	incomePostParams := &income.PostIncomeParams{UserID: userID,
		IncomePostRequest: &models.IncomePost{Accounts: []string{accountID}, FromMonth: "2018-11", ToMonth: "2019-10"},
		Context:           context.TODO(),
	}

	incomeRsp, err := test.Client.Income.PostIncome(incomePostParams, token)

	if err != nil {
		t.Fatalf("Cannot post income. Error: %v", err)
	}

	e, err := json.Marshal(incomeRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/postIncome.json", t)
	s = strings.Replace(s, "c4c4466d-55e1-42b5-88cf-7b04f32af485", *incomeRsp.GetPayload().ID, 2)

	test.AssertJson(t, s, string(e))
}

func TestPostIncomeBadRequest(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5"
	accountID := "14aa946c-6f87-44b7-bab9-d192d6261471"
	incomePostParams := &income.PostIncomeParams{UserID: userID,
		IncomePostRequest: &models.IncomePost{FromMonth: "2015-11", ToMonth: "2020-10", Accounts: []string{accountID}},
		Context:           context.TODO(),
	}

	_, err := test.Client.Income.PostIncome(incomePostParams, token)

	if err != nil {
		badRequest, ok := err.(*income.PostIncomeBadRequest)
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
