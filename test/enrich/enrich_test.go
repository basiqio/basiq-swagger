package enrich

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/basiqio/basiq-swagger/dist/client/enrich"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
)

func TestGetEnrich(t *testing.T) {
	accountType := "credit-card"
	amount := float64(-124)
	enrichParams := &enrich.GetEnrichParams{
		AccountType: &accountType,
		Amount:      &amount,
		Institution: "AU00000",
		Q:           "Starbucks",
		Context:     context.TODO(),
	}

	enrichResponse, err := test.Client.Enrich.GetEnrich(enrichParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error getting enrich: %v", err)
	}

	e, err := json.Marshal(enrichResponse.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getEnrich.json", t)

	test.AssertJson(t, s, string(e))
}

func TestGetEnrichBadRequest(t *testing.T) {
	accountType := "transaction"
	amount := float64(-12.95)
	enrichParams := &enrich.GetEnrichParams{
		AccountType: &accountType,
		Amount:      &amount,
		Institution: "",
		Q:           "Starbucks",
		Context:     context.TODO(),
	}

	_, err := test.Client.Enrich.GetEnrich(enrichParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		badRequest, ok := err.(*enrich.GetEnrichBadRequest)
		if !ok {
			t.Fatal("Bad request error not returned.")
		}

		errorModel, err := json.Marshal(badRequest.GetPayload())
		if err != nil {
			t.Fatalf("Cannot marshal error model. Error: %v", err)
		}
		s := strings.Replace(
			test.GetJsonResponse("./responses/getEnrichBadRequest.json", t), "1c16f75c-36fb-421f-98a5-bf18a0b6583c",
			*badRequest.GetPayload().CorrelationID,
			2)

		test.AssertJson(t, s, string(errorModel))
	} else {
		t.Fatal("Error model should be returned")
	}
}
