package affordabilityList

import (
	"context"
	"encoding/json"
	"github.com/basiqio/basiq-swagger/dist/client/affordability"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
	"strings"
	"testing"
)

func TestGetAffordabilityList(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := "53078a89-8267-4883-9169-9a652805a2c8"

	affordabilityParams := &affordability.GetAffordabilityListParams{
		UserID:  userID,
		Context: context.TODO(),
	}

	affordabilityListRsp, err := test.Client.Affordability.GetAffordabilityList(affordabilityParams, token)

	if err != nil {
		t.Fatalf("Cannot get affordability list. Error: %v", err)
	}

	e, err := json.Marshal(affordabilityListRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getAffordabilityList.json", t)

	test.AssertJson(t, s, string(e))
}

func TestGetAffordabilityListBadRequest(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	notValidUserID := "fabd34fd-0f3c-4678-8db6-4ff547c4f5"

	affordabilityParams := &affordability.GetAffordabilityListParams{
		UserID:  notValidUserID,
		Context: context.TODO(),
	}

	_, err := test.Client.Affordability.GetAffordabilityList(affordabilityParams, token)

	if err != nil {
		badRequest, ok := err.(*affordability.GetAffordabilityListBadRequest)
		if !ok {
			t.Fatal("Bad request error not returned.")
		}

		s := strings.Replace(
			test.GetJsonResponse("./responses/badRequest.json", t), "e6b937bb-4815-44ec-a6b3-2bcfa7496245",
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
