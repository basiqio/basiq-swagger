package test

import (
	"context"
	"fmt"
	client2 "github.com/basiqio/basiq-swagger/dist/client"
	token2 "github.com/basiqio/basiq-swagger/dist/client/token"
	"github.com/basiqio/basiq-swagger/dist/client/users"
	"github.com/basiqio/basiq-swagger/dist/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

// InMemoryToken holds auth details
type InMemoryToken struct {
	token       string
	expiresIn   int64
	timeElapsed *time.Time
}

type Token interface {
	getToken(t *testing.T) string
}

var Client *client2.Basiq200
var Cfg Config
var TokenHolder InMemoryToken

func init() {
	TokenHolder = InMemoryToken{}
	Cfg = Config{}
	err := ReadFile(&Cfg)
	if err != nil {
		fmt.Printf("Error initializing tests. Error: %s", err)
		os.Exit(2)
	}
	Client = client2.New(httptransport.New(client2.DefaultHost, client2.DefaultBasePath, nil), strfmt.Default)
}

func (token *InMemoryToken) GetToken(t *testing.T) string {
	if token.timeElapsed == nil || int64(time.Now().Sub(*token.timeElapsed).Seconds()) >= token.expiresIn {
		tokenReq := token2.NewPostTokenParamsWithContext(context.TODO())
		tokenReq.SetBasiqVersion("2.0")
		tokenOK, err := Client.Token.PostToken(tokenReq, httptransport.APIKeyAuth("Authorization", "header", "Basic "+Cfg.ApiKey))

		if err != nil {
			t.Fatalf("Error getting token : %v", err)
		}

		token.token = *tokenOK.Payload.AccessToken
		token.expiresIn = *tokenOK.Payload.ExpiresIn
		time := time.Now()
		token.timeElapsed = &time
		return token.token
	}

	return token.token
}

func GetJsonResponse(file string, t *testing.T) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatalf("Cannot read file, Error: %v", err)
	}

	return string(content)
}

// AssertJson compares if two json strings are the same
func AssertJson(t *testing.T, expected, actual string) {
	t.Helper()
	match, err := gomega.MatchJSON(expected).Match(actual)
	if err != nil {
		t.Fatalf("Error matching json. Error: %v", err)
	}

	if !match {
		fmt.Printf("%v\n%v\n", expected, actual)
		t.Fatalf("Responses are not the same.")
	}
}

func CleanupUser(t *testing.T, userId string) {
	userDeleteRequest := &users.DeleteUserParams{
		UserID:  userId,
		Context: context.TODO(),
	}

	_, err := Client.Users.DeleteUser(userDeleteRequest, httptransport.BearerToken(TokenHolder.GetToken(t)))
	if err != nil {
		t.Fatalf("Error deleting user: %v", err)
	}
}

func SetupUser(t *testing.T) string {
	userEmail := "gavin1224@hooli.com"
	userMobile := "+61410888999"

	userPostRequest := &users.CreateUserParams{
		User: &models.CreateUser{
			Email:  userEmail,
			Mobile: userMobile,
		},
		Context: context.TODO(),
	}

	userResponse, err := Client.Users.CreateUser(userPostRequest, httptransport.BearerToken(TokenHolder.GetToken(t)))

	if err != nil {
		t.Fatal("Error creating user")
	}
	return *userResponse.GetPayload().ID
}
