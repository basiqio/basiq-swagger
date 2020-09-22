package statements

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/basiqio/basiq-swagger/dist/client/statements"
	"github.com/basiqio/basiq-swagger/test"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"io/ioutil"
	"strings"
	"testing"
)

func TestStatements(t *testing.T) {

	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := test.SetupUser(t)
	defer test.CleanupUser(t, userID)

	statementFile, err := ioutil.ReadFile("statement.csv")
	if err != nil {
		t.Fatalf("Cannot read statement file. Error: %v", err)
	}

	statement := runtime.NamedReader("statement", bytes.NewBuffer(statementFile))

	statementsParams := &statements.CreateStatementParams{
		UserID:        userID,
		Statement:     statement,
		InstitutionID: "AU00201",
		Context:       context.TODO(),
	}

	statementsRsp, err := test.Client.Statements.CreateStatement(statementsParams, token)

	if err != nil {
		t.Fatalf("Cannot upload statement: %v", err)
	}

	e, err := json.Marshal(statementsRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := strings.Replace(test.GetJsonResponse("./responses/statementUpload.json", t),
		"ede24fef-22fc-4bb9-ae31-2fe6cfd43c5e",
		*statementsRsp.GetPayload().ID, 2)

	test.AssertJson(t, s, string(e))
}

func TestStatementsNoInstitution(t *testing.T) {

	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userID := test.SetupUser(t)
	defer test.CleanupUser(t, userID)

	statementFile, err := ioutil.ReadFile("statement.csv")
	if err != nil {
		t.Fatalf("Cannot read statement file. Error: %v", err)
	}

	statement := runtime.NamedReader("statement", bytes.NewBuffer(statementFile))

	statementsParams := &statements.CreateStatementParams{
		UserID:    userID,
		Statement: statement,
		Context:   context.TODO(),
	}

	_, err = test.Client.Statements.CreateStatement(statementsParams, token)

	if err != nil {
		badRequest, ok := err.(*statements.CreateStatementBadRequest)
		if !ok {
			t.Fatalf("Bad request should be returned")
		}
		e, err := json.Marshal(badRequest.GetPayload())
		if err != nil {
			t.Fatalf("Error: %v", err)
		}

		s := strings.Replace(test.GetJsonResponse("./responses/badRequest.json", t),
			"0d68bb88-d40c-4c5f-8416-c07021f07123",
			*badRequest.GetPayload().CorrelationID, 1)

		test.AssertJson(t, s, string(e))

	} else {
		t.Fatalf("Error should be returned")
	}
}
