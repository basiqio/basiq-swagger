package token

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/basiqio/basiq-swagger/dist/client/token"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
)

func TestPostToken(t *testing.T) {
	tokenParams := &token.PostTokenParams{
		Authorization: "Basic " + test.Cfg.ApiKey,
		BasiqVersion:  "2.0",
		Context:       context.TODO(),
	}

	tokenResponse, err := test.Client.Token.PostToken(tokenParams, httptransport.PassThroughAuth)
	if err != nil {
		t.Fatalf("Error getting token: %v", err)
	}

	e, err := json.Marshal(tokenResponse.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := strings.Replace(
		test.GetJsonResponse("./responses/postToken.json", t), "eyJh...vH4g",
		*tokenResponse.GetPayload().AccessToken,
		2)

	test.AssertJson(t, s, string(e))
}

func TestPostTokenBadRequest(t *testing.T) {
	tokenParams := &token.PostTokenParams{
		Authorization: "Basic " + test.Cfg.ApiKey,
		BasiqVersion:  "",
		Context:       context.TODO(),
	}

	_, err := test.Client.Token.PostToken(tokenParams, httptransport.PassThroughAuth)
	if err != nil {
		badRequest, ok := err.(*token.PostTokenBadRequest)
		if !ok {
			t.Fatal("Bad request error not returned.")
		}

		errorModel, err := json.Marshal(badRequest.GetPayload())
		if err != nil {
			t.Fatalf("Cannot marshal error model. Error: %v", err)
		}
		s := strings.Replace(
			test.GetJsonResponse("./responses/postTokenBadRequest.json", t), "1c16f75c-36fb-421f-98a5-bf18a0b6583c",
			*badRequest.GetPayload().CorrelationID,
			2)

		test.AssertJson(t, s, string(errorModel))
	} else {
		t.Fatal("Error model should be returned")
	}
}

func TestPostTokenNotFound(t *testing.T) {
	tokenParams := &token.PostTokenParams{
		Authorization: "Basic " + test.Cfg.DeletedApiKey,
		BasiqVersion:  "2.0",
		Context:       context.TODO(),
	}

	_, err := test.Client.Token.PostToken(tokenParams, httptransport.PassThroughAuth)
	if err != nil {
		fmt.Println(test.Cfg.DeletedApiKey)
		fmt.Println(err.Error())
		notFound, ok := err.(*token.PostTokenNotFound)
		if !ok {
			t.Fatal("Not found error not returned.")
		}

		errorModel, err := json.Marshal(notFound.GetPayload())
		if err != nil {
			t.Fatalf("Cannot marshal error model. Error: %v", err)
		}
		s := strings.Replace(
			test.GetJsonResponse("./responses/postTokenNotFound.json", t), "1c16f75c-36fb-421f-98a5-bf18a0b6583c",
			*notFound.GetPayload().CorrelationID,
			2)

		test.AssertJson(t, s, string(errorModel))
	} else {
		t.Fatal("Error model should be returned")
	}
}
