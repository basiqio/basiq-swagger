package affordability

import (
	"context"
	"encoding/json"
	"github.com/basiqio/basiq-swagger/dist/client/affordability"
	"github.com/basiqio/basiq-swagger/dist/models"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
	"strings"
	"testing"
)

func TestGetAffordability(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5"
	snapshotID := "91b42a5b-2f9a-479d-9014-cdc211155f74"

	affordabilityParams := &affordability.GetAffordabilityParams{UserID: userID,
		SnapshotID: snapshotID,
		Context:    context.TODO(),
	}

	affordabilityRsp, err := test.Client.Affordability.GetAffordability(affordabilityParams, token)

	if err != nil {
		t.Fatalf("Cannot get affordability. Error: %v", err)
	}

	e, err := json.Marshal(affordabilityRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getAffordability.json", t)
	s = strings.Replace(s, "91b42a5b-2f9a-479d-9014-cdc211155f74", *affordabilityRsp.GetPayload().ID, 2)

	test.AssertJson(t, s, string(e))
}

func TestPostAffordability(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5"
	accountID := "14aa946c-6f87-44b7-bab9-d192d6261471"
	affordabilityPostParams := &affordability.PostAffordabilityParams{UserID: userID,
		AffordabilityPostRequest: &models.AffordabilityPost{FromMonth: "2018-11", ToMonth: "2019-10", Accounts: []string{accountID}},
		Context:                  context.TODO(),
	}

	affordabilityPostRsp, _, err := test.Client.Affordability.PostAffordability(affordabilityPostParams, token)

	if err != nil {
		t.Fatalf("Cannot post affordability. Error: %v", err)
	}

	e, err := json.Marshal(affordabilityPostRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/postAffordability.json", t)
	s = strings.Replace(s, "expensesLink", *affordabilityPostRsp.GetPayload().Links.Expenses, 1)
	s = strings.Replace(s, "selfLink", *affordabilityPostRsp.GetPayload().Links.Self, 1)
	s = strings.Replace(s, "incomeLink", *affordabilityPostRsp.GetPayload().Links.Income, 1)
	s = strings.Replace(s, "2020-09-11T18:45:56", *affordabilityPostRsp.GetPayload().GeneratedDate, 1)
	s = strings.Replace(s, "17245135-ccd2-4b66-8b34-bef1a8342827", *affordabilityPostRsp.GetPayload().ID, 1)

	test.AssertJson(t, s, string(e))
}

func TestPostAffordabilityEmptyResponse(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5"
	accountID := "14aa946c-6f87-44b7-bab9-d192d6261471"
	affordabilityPostParams := &affordability.PostAffordabilityParams{UserID: userID,
		AffordabilityPostRequest: &models.AffordabilityPost{FromMonth: "2015-11", ToMonth: "2016-10", Accounts: []string{accountID}},
		Context:                  context.TODO(),
	}

	_, _, err := test.Client.Affordability.PostAffordability(affordabilityPostParams, token)

	if err == nil {
		t.Fatalf("Response 204 should be returned . Error: %v", err)
	}
	// todo check to 204 error rsp
}

func TestPostAffordabilityBadRequest(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5"
	accountID := "14aa946c-6f87-44b7-bab9-d192d6261471"
	affordabilityPostParams := &affordability.PostAffordabilityParams{UserID: userID,
		AffordabilityPostRequest: &models.AffordabilityPost{FromMonth: "2015-11", ToMonth: "2020-10", Accounts: []string{accountID}},
		Context:                  context.TODO(),
	}

	_, _, err := test.Client.Affordability.PostAffordability(affordabilityPostParams, token)

	if err != nil {
		badRequest, ok := err.(*affordability.PostAffordabilityBadRequest)
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

func TestPostAffordabilityPdf(t *testing.T) {
	_ = httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5"
	accountID := "14aa946c-6f87-44b7-bab9-d192d6261471"
	_ = &affordability.PostAffordabilityParams{UserID: userID,
		AffordabilityPostRequest: &models.AffordabilityPost{FromMonth: "2018-11", ToMonth: "2019-10", Accounts: []string{accountID}},
		Context:                  context.TODO(),
	}

}
