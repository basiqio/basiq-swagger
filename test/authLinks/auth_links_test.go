package authLinks

import (
	"context"
	"encoding/json"
	"github.com/basiqio/basiq-swagger/dist/client/auth_links"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
	"regexp"
	"testing"
)

func TestPostGetDeleteAuthLink(t *testing.T) {
	authLinkPostParams := &auth_links.PostAuthLinkParams{
		AuthLinksPostData: auth_links.PostAuthLinkBody{Mobile: "+61410888666"},
		UserID:            "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
		Context:           context.TODO(),
	}

	authLinkRsp, err := test.Client.AuthLinks.PostAuthLink(authLinkPostParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error getting auth link: %v", err)
	}

	e, err := json.Marshal(authLinkRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	authLinkGetParams := &auth_links.GetAuthLinkParams{
		UserID:  "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
		Context: context.TODO(),
	}
	authLinkGetRsp, err := test.Client.AuthLinks.GetAuthLink(authLinkGetParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))

	s, err := json.Marshal(authLinkGetRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	// remove id from the response
	re := regexp.MustCompile(`(\"id\":\"[\w]{8}(-[\w]{4}){3}-[\w]{12}\",)`)
	repl := re.ReplaceAllString(string(s), ``)
	test.AssertJson(t, repl, string(e))

	authLinkDeleteParams := &auth_links.DeleteAuthLinkParams{
		UserID:  "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
		Context: context.TODO(),
	}
	_, err = test.Client.AuthLinks.DeleteAuthLink(authLinkDeleteParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatal("Delete of auth link failed.")
	}

	_, err = test.Client.AuthLinks.GetAuthLink(authLinkGetParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))

	if err != nil {
		_, ok := err.(*auth_links.GetAuthLinkNotFound)
		if !ok {
			t.Fatal("Returned response should be not found.")
		}
	} else {
		t.Fatal("Auth link should be not found.")
	}
}
