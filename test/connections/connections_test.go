package connections

import (
	"context"
	"encoding/json"
	"github.com/basiqio/basiq-swagger/dist/client/connections"
	"github.com/basiqio/basiq-swagger/dist/client/users"
	"github.com/basiqio/basiq-swagger/dist/models"
	"github.com/basiqio/basiq-swagger/test"
	httptransport "github.com/go-openapi/runtime/client"
	"strings"
	"testing"
)

func TestGetConnections(t *testing.T) {
	filter := "status.eq('active')"
	getConnParams := &connections.GetConnectionsParams{UserID: "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
		Filter:  &filter,
		Context: context.TODO(),
	}

	connectionsRsp, err := test.Client.Connections.GetConnections(getConnParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))

	if err != nil {
		t.Fatalf("Error getting connections: %v", err)
	}

	e, err := json.Marshal(connectionsRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getConnections.json", t)

	test.AssertJson(t, s, string(e))
}

func TestGetConnectionsEmpty(t *testing.T) {
	filter := "status.eq('invalid')"
	getConnParams := &connections.GetConnectionsParams{UserID: "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
		Filter:  &filter,
		Context: context.TODO(),
	}

	connectionsRsp, err := test.Client.Connections.GetConnections(getConnParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))

	if err != nil {
		t.Fatalf("Error getting connections: %v", err)
	}

	e, err := json.Marshal(connectionsRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getConnectionsEmpty.json", t)

	test.AssertJson(t, s, string(e))
}

func TestGetConnection(t *testing.T) {
	getConnParams := &connections.GetConnectionParams{UserID: "8cda72db-b11f-4b8e-a4ca-3c5b1de4e4b5",
		ConnectionID: "64011b56-988a-4c2c-bfca-16c836cec23f",
		Context:      context.TODO(),
	}

	connectionRsp, err := test.Client.Connections.GetConnection(getConnParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))

	if err != nil {
		t.Fatalf("Error getting connection: %v", err)
	}

	e, err := json.Marshal(connectionRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := test.GetJsonResponse("./responses/getConnection.json", t)

	test.AssertJson(t, s, string(e))
}

func TestCreateConnection(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))

	userID := setupUser(t)
	defer cleanupUser(t, userID)

	loginID := "Wentworth-Smith"
	password := "whislter"
	institutionID := "AU00000"
	institution := models.InstitutionModel{ID: &institutionID}

	connCreateParams := &connections.PostConnectionParams{
		UserID:                  userID,
		UserConnectionsPostData: &models.UserConnectionsPostData{Password: &password, LoginID: &loginID, Institution: &institution},
		Context:                 context.TODO(),
	}

	connCreateRsp, err := test.Client.Connections.PostConnection(connCreateParams, token)

	if err != nil {
		t.Fatalf("Error posting connection, Error: %v", err)
	}

	e, err := json.Marshal(connCreateRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := strings.Replace(test.GetJsonResponse("./responses/createConnection.json", t), "fa521b1b-0031-4644-8885-7ed125e659c1", *connCreateRsp.GetPayload().ID, 2)
	test.AssertJson(t, s, string(e))
}

func TestRefreshConnectionsAndConnection(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	userId := "8071a28a-51a6-4fc1-8b53-3c6d0d335a83"
	connRefParams := &connections.RefreshConnectionsParams{
		UserID:  userId,
		Context: context.TODO(),
	}

	connRefreshRsp, err := test.Client.Connections.RefreshConnections(connRefParams, token)

	if err != nil {
		t.Fatalf("error refreshing connections. Error: %v", err)
	}

	e, err := json.Marshal(connRefreshRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := strings.Replace(test.GetJsonResponse("./responses/refreshConnections.json", t),
		"92791915-6f34-4895-80e5-9ea1a2273ec6",
		*connRefreshRsp.GetPayload().ConnectionsRefreshJobs[0].ID, 2)
	test.AssertJson(t, s, string(e))

	connectionRefParams := &connections.RefreshConnectionParams{
		UserID:       userId,
		ConnectionID: "0a08e2cf-3c1f-4c62-a128-c04c82a5851a",
		Context:      context.TODO(),
	}

	refConnectionRsp, err := test.Client.Connections.RefreshConnection(connectionRefParams, token)

	if err != nil {
		t.Fatalf("Error refreshing connection. Error: %v", err)
	}

	e, err = json.Marshal(refConnectionRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s = strings.Replace(test.GetJsonResponse("./responses/refreshConnection.json", t),
		"df9e9086-0354-42ea-afb2-6df2e45c4eae",
		*refConnectionRsp.GetPayload().ID, 2)
	test.AssertJson(t, s, string(e))
}

func TestUpdateConnection(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))
	password := "whislter"
	connUpdateeParams := &connections.UpdateConnectionParams{
		ConnectionID:           "0a08e2cf-3c1f-4c62-a128-c04c82a5851a",
		UserID:                 "8071a28a-51a6-4fc1-8b53-3c6d0d335a83",
		UserConnectionPostData: &models.UserConnectionPostData{Password: &password},
		Context:                context.TODO(),
	}

	connUpdateRsp, err := test.Client.Connections.UpdateConnection(connUpdateeParams, token)

	if err != nil {
		t.Fatalf("Error updating connection, Error: %v", err)
	}

	e, err := json.Marshal(connUpdateRsp.GetPayload())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	s := strings.Replace(test.GetJsonResponse("./responses/updateConnection.json", t),
		"aff7df1f-20ba-4c9a-8af3-f3c62280047c",
		*connUpdateRsp.GetPayload().ID, 2)
	test.AssertJson(t, s, string(e))
}

func TestDeleteConnection(t *testing.T) {
	token := httptransport.BearerToken(test.TokenHolder.GetToken(t))

	userID := setupUser(t)
	defer cleanupUser(t, userID)

	loginID := "Wentworth-Smith"
	password := "whislter"
	institutionID := "AU00000"
	institution := models.InstitutionModel{ID: &institutionID}

	connCreateParams := &connections.PostConnectionParams{
		UserID:                  userID,
		UserConnectionsPostData: &models.UserConnectionsPostData{Password: &password, LoginID: &loginID, Institution: &institution},
		Context:                 context.TODO(),
	}

	_, err := test.Client.Connections.PostConnection(connCreateParams, token)

	if err != nil {
		t.Fatalf("Error posting connection, Error: %v", err)
	}
	//get connection
	getConnParams := &connections.GetConnectionsParams{UserID: userID,
		Context: context.TODO(),
	}

	connectionsRsp, err := test.Client.Connections.GetConnections(getConnParams, httptransport.BearerToken(test.TokenHolder.GetToken(t)))

	connDeleteParams := &connections.DeleteConnectionParams{
		UserID:       userID,
		ConnectionID: *connectionsRsp.GetPayload().ConnectionsData[0].ID,
		Context:      context.TODO(),
	}

	_, err = test.Client.Connections.DeleteConnection(connDeleteParams, token)
	if err != nil {
		t.Fatalf("Error deleting connection, Error: %v", err)
	}
}

func cleanupUser(t *testing.T, userId string) {
	userDeleteRequest := &users.DeleteUserParams{
		UserID:  userId,
		Context: context.TODO(),
	}

	_, err := test.Client.Users.DeleteUser(userDeleteRequest, httptransport.BearerToken(test.TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}

func setupUser(t *testing.T) string {
	userEmail := "gavin1224@hooli.com"
	userMobile := "+61410888999"

	userPostRequest := &users.CreateUserParams{
		User: &models.UserPost{
			Email:  userEmail,
			Mobile: userMobile,
		},
		Context: context.TODO(),
	}

	userResponse, err := test.Client.Users.CreateUser(userPostRequest, httptransport.BearerToken(test.TokenHolder.GetToken(t)))

	if err != nil {
		t.Fatal("Error creating user")
	}
	return *userResponse.GetPayload().ID
}
