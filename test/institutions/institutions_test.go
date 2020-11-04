package institutions

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/basiqio/basiq-swagger/dist/client/institutions"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
)

func TestGetInstitutions(t *testing.T) {
	filter := "institution.country.in('Yugoslavia')" // Filter will return OK response with empty list. Avoids large json.
	institutionsParams := &institutions.GetInstitutionsParams{
		Filter:  &filter,
		Context: context.TODO(),
	}

	institutionsResponse, err := test.Client.Institutions.GetInstitutions(institutionsParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error getting institutions: %v", err)
	}

	e, err := json.Marshal(institutionsResponse.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getInstitutions.json", t)

	test.AssertJson(t, s, string(e))
}

func TestGetInstitutionsUnauthorized(t *testing.T) {
	filter := "institution.country.in('Australia')"
	institutionsParams := &institutions.GetInstitutionsParams{
		Filter:  &filter,
		Context: context.TODO(),
	}

	_, err := test.Client.Institutions.GetInstitutions(institutionsParams, httptransport.PassThroughAuth)
	if err != nil {
		unauthorized, ok := err.(*institutions.GetInstitutionsUnauthorized)
		if !ok {
			t.Fatal("Unauthorized error not returned.")
		}

		errorModel, err := json.Marshal(unauthorized.GetPayload())
		if err != nil {
			t.Fatalf("Cannot marshal error model. Error: %v", err)
		}
		s := strings.Replace(
			test.GetJsonResponse("./responses/getInstitutionsUnauthorized.json", t), "1c16f75c-36fb-421f-98a5-bf18a0b6583c",
			*unauthorized.GetPayload().CorrelationID,
			2)

		test.AssertJson(t, s, string(errorModel))
	} else {
		t.Fatal("Error model should be returned")
	}
}

func TestGetInstitution(t *testing.T) {
	institutionParams := &institutions.GetInstitutionParams{
		InstitutionID: "AU00000",
		Context:       context.TODO(),
	}

	institutionResponse, err := test.Client.Institutions.GetInstitution(institutionParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error getting institution: %v", err)
	}

	e, err := json.Marshal(institutionResponse.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getInstitution.json", t)
	s = strings.Replace(
		s, "3820",
		strconv.Itoa(int(institutionResponse.GetPayload().Stats.AverageDurationMs.RetrieveAccounts)),
		1)
	s = strings.Replace(
		s, "12487",
		strconv.Itoa(int(institutionResponse.GetPayload().Stats.AverageDurationMs.RetrieveMeta)),
		1)
	s = strings.Replace(
		s, "7127",
		strconv.Itoa(int(institutionResponse.GetPayload().Stats.AverageDurationMs.RetrieveTransactions)),
		1)
	s = strings.Replace(
		s, "31588",
		strconv.Itoa(int(institutionResponse.GetPayload().Stats.AverageDurationMs.Total)),
		1)
	s = strings.Replace(
		s, "8154",
		strconv.Itoa(int(institutionResponse.GetPayload().Stats.AverageDurationMs.VerifyCredentials)),
		1)

	test.AssertJson(t, s, string(e))
}

func TestGetInstitutionNotFound(t *testing.T) {
	institutionParams := &institutions.GetInstitutionParams{
		InstitutionID: "YU00000",
		Context:       context.TODO(),
	}

	_, err := test.Client.Institutions.GetInstitution(institutionParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		notFound, ok := err.(*institutions.GetInstitutionNotFound)
		if !ok {
			t.Fatal("Not found error not returned.")
		}

		errorModel, err := json.Marshal(notFound.GetPayload())
		if err != nil {
			t.Fatalf("Cannot marshal error model. Error: %v", err)
		}
		s := strings.Replace(
			test.GetJsonResponse("./responses/getInstitutionNotFound.json", t), "1c16f75c-36fb-421f-98a5-bf18a0b6583c",
			*notFound.GetPayload().CorrelationID,
			2)

		test.AssertJson(t, s, string(errorModel))
	} else {
		t.Fatal("Error model should be returned")
	}
}

func TestGetPublicInstitutions(t *testing.T) {
	filter := "institution.country.in('Yugoslavia')"
	institutionsParams := &institutions.GetPublicInstitutionsParams{
		Filter:  &filter,
		Context: context.TODO(),
	}

	institutionsResponse, err := test.Client.Institutions.GetPublicInstitutions(institutionsParams)
	if err != nil {
		t.Fatalf("Error getting institutions: %v", err)
	}

	e, err := json.Marshal(institutionsResponse.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getPublicInstitutions.json", t)

	test.AssertJson(t, s, string(e))
}
