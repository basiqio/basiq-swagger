package test

import (
	"context"
	"encoding/json"
	"github.com/basiqio/basiq-swagger/dist/client/users"
	"github.com/basiqio/basiq-swagger/dist/models"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
	"strings"
	"testing"
)

func TestGetUser(t *testing.T) {
	userGetRequest := &users.GetUserParams{
		UserID:  "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
		Context: context.TODO(),
	}

	userResponse, err := test.Client.Users.GetUser(userGetRequest, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error getting user: %v", err)
	}

	e, err := json.Marshal(userResponse.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getUser.json", t)

	test.AssertJson(t, s, string(e))
}

func TestGetUserForbidden(t *testing.T) {
	userGetRequest := &users.GetUserParams{
		UserID:  "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
		Context: context.TODO(),
	}

	_, err := test.Client.Users.GetUser(userGetRequest, httptransport.BearerToken(test.Cfg.ExpiredToken))
	if err != nil {

		forbidden, ok := err.(*users.GetUserForbidden)
		if !ok {
			t.Fatal("Forbidden error not returned.")
		}

		errorModel, err := json.Marshal(forbidden.GetPayload())
		if err != nil {
			t.Fatalf("Cannot marshal error model. Error: %v", err)
		}
		s := strings.Replace(
			test.GetJsonResponse("./responses/getUserForbidden.json", t), "bf4ac637-4e5c-4c87-aa1c-6158beb129a3",
			*forbidden.GetPayload().CorrelationID,
			2)

		test.AssertJson(t, s, string(errorModel))
	} else {
		t.Fatal("Error model should be returned")
	}
}

func TestUpdateUser(t *testing.T) {

	userUpdateRequest := &users.UpdateUserParams{
		UserID: "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
		User: &models.UpdateUser{
			Email:  "gavin@hooli.com",
			Mobile: "+61410888666",
		},
		Context: context.TODO(),
	}
	userResponse, err := test.Client.Users.UpdateUser(userUpdateRequest, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error getting user: %v", err)
	}

	e, err := json.Marshal(userResponse.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/updateUser.json", t)

	test.AssertJson(t, s, string(e))
}

func TestCreateAndDeleteUser(t *testing.T) {
	userEmail := "gavin1224@hooli.com"
	userMobile := "+61410888999"

	userPostRequest := &users.CreateUserParams{
		User: &models.CreateUser{
			Email:  userEmail,
			Mobile: userMobile,
		},
		Context: context.TODO(),
	}

	userResponse, err := test.Client.Users.CreateUser(userPostRequest, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error creating user. Error: %v", err)
	}

	e, err := json.Marshal(userResponse.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	// replace userId to the created one
	s := strings.Replace(
		test.GetJsonResponse("./responses/postUser.json", t), "86c79d8c-5033-46b3-b16e-d6cee4200fb3",
		*userResponse.GetPayload().ID,
		2)

	test.AssertJson(t, s, string(e))

	// deleting a user
	userDeleteRequest := &users.DeleteUserParams{
		UserID:  *userResponse.GetPayload().ID,
		Context: context.TODO(),
	}

	_, err = test.Client.Users.DeleteUser(userDeleteRequest, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}

func TestCreateBadRequest(t *testing.T) {
	userEmail := ""
	userMobile := ""

	userPostRequest := &users.CreateUserParams{
		User: &models.CreateUser{
			Email:  userEmail,
			Mobile: userMobile,
		},
		Context: context.TODO(),
	}

	_, err := test.Client.Users.CreateUser(userPostRequest, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {

		badRequest, ok := err.(*users.CreateUserBadRequest)
		if !ok {
			t.Fatal("Bad request error not returned.")
		}

		s := strings.Replace(
			test.GetJsonResponse("./responses/createUserBadRequest.json", t), "ef351da5-12b1-4ba0-be8c-59d9b2e388e8",
			*badRequest.GetPayload().CorrelationID,
			1)
		e, err := json.Marshal(badRequest.GetPayload())
		if err != nil {
			t.Fatalf("Cannot marshal error model. Error: %v", err)
		}

		test.AssertJson(t, s, string(e))
	} else {
		t.Fatal("Bad request error not returned.")
	}
}

func TestGetUserNotFound(t *testing.T) {
	userGetRequest := &users.GetUserParams{
		UserID:  "ed70cdc8-8314-4ca7-adc3-a7b34d299832",
		Context: context.TODO(),
	}

	_, err := test.Client.Users.GetUser(userGetRequest, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		notFoundError, ok := err.(*users.GetUserNotFound)
		if !ok {
			t.Fatal("Bad request error not returned.")
		}

		s := strings.Replace(
			test.GetJsonResponse("./responses/userNotFound.json", t), "886fee38-a7a7-472f-82d1-18f2e00f07b6",
			*notFoundError.GetPayload().CorrelationID,
			1)
		e, err := json.Marshal(notFoundError.GetPayload())
		if err != nil {
			t.Fatalf("Cannot marshal error model. Error: %v", err)
		}

		test.AssertJson(t, s, string(e))
	} else {
		t.Fatal("Not found error not returned.")
	}
}
