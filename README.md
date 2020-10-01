# Basiq-swagger
Basiq API documentation.

### Installation
The main step is to create a client using Swagger, and with a client you call all Basiq endpoints.

### Install packages
To build sdk use ./build.sh script. Before building the sdk make sure you have goswagger installed (https://goswagger.io).

Installation of the goswagger tool:

go get -u github.com/go-swagger/go-swagger/cmd/swagger 

For this generation to compile you need to have some packages in your GOPATH:

        * github.com/go-openapi/errors
        * github.com/go-openapi/runtime
        * github.com/go-openapi/runtime/client
        * github.com/go-openapi/strfmt

#### Quick start
To run these examples just create main.go in the root of the project. 

#### Example GET Token:
```
import (
	"context"
	client2 "github.com/basiqio/basiq-swagger/dist/client"
	token2 "github.com/basiqio/basiq-swagger/dist/client/token"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"log"
)


func main() {

	client := client2.New(httptransport.New(client2.DefaultHost, client2.DefaultBasePath, nil), strfmt.Default)

	tokenReq := token2.NewPostTokenParamsWithContext(context.TODO())
	tokenReq.SetAuthorization("Basic " + apiKey)
	tokenReq.SetBasiqVersion("2.0")
	tokenResponse, err := client.Token.PostToken(tokenReq, nil)

	if err != nil {
		log.Printf("Error getting token: %v", err)
	}

	token := tokenResponse.GetPayload()

	log.Printf("AccessToken: %v, Expires: %v, Token Type: %v", *token.AccessToken, *token.ExpiresIn, *token.TokenType)

}
```

#### Example POST User:
```
import (
	"context"
	client2 "github.com/basiqio/basiq-swagger/dist/client"
    "github.com/basiqio/basiq-swagger/dist/client/users"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"log"
)


func main() {

   	client := client2.New(httptransport.New(client2.DefaultHost, client2.DefaultBasePath, nil), strfmt.Default)

	userPostRequest := &users.CreateUserParams{
		User: &models.UserPost{
			Email:  userEmail,
			Mobile: userMobile,
		},
		Context: context.TODO(),
	}

	userResponse, err := client.Users.CreateUser(userPostRequest, httptransport.BearerToken(accessToken))

	if err != nil {
		log.Printf("%+v", err)
	}

	user := userResponse.GetPayload()
}
```

#### Example POST AuthLinks: 
```
import (
	"context"
	client2 "github.com/basiqio/basiq-swagger/dist/client"
	"github.com/basiqio/basiq-swagger/dist/client/auth_links"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"log"
)

func main() {

   	client := client2.New(httptransport.New(client2.DefaultHost, client2.DefaultBasePath, nil), strfmt.Default)

	authLinkPostParams := &auth_links.PostAuthLinkParams{
		AuthLinksPostData: auth_links.PostAuthLinkBody{Mobile: userMobile},
		UserID:            userID,
		Context:           context.TODO(),
	}

	authLinkRsp, err := client.AuthLinks.PostAuthLink(authLinkPostParams, httptransport.BearerToken(accessToken))
	if err != nil {
		log.Printf("%+v", err)
	}

    authLink := authLinkRsp.GetPayload()
}
```

#### Example GET AuthLinks: 

```
import (
	"context"
	client2 "github.com/basiqio/basiq-swagger/dist/client"
	"github.com/basiqio/basiq-swagger/dist/client/auth_links"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"log"
)

func main() {

   	client := client2.New(httptransport.New(client2.DefaultHost, client2.DefaultBasePath, nil), strfmt.Default)

	authLinkGetParams := &auth_links.GetAuthLinkParams{
		UserID:  userID,
		Context: context.TODO(),
	}
	authLinkGetRsp, err := client.AuthLinks.GetAuthLink(authLinkGetParams, httptransport.BearerToken(accessToken))

	if err != nil {
		log.Printf("%+v", err)
	}

    authLinks := authLinkGetRsp.GetPayload()
}
```

#### Example GET transactions:
```
import (
	"context"
	client2 "github.com/basiqio/basiq-swagger/dist/client"
    "github.com/basiqio/basiq-swagger/dist/client/transactions"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"log"
)

func main() {

   	client := client2.New(httptransport.New(client2.DefaultHost, client2.DefaultBasePath, nil), strfmt.Default)

	txParams := &transactions.GetTransactionsParams{
		UserID:  userID,
		Context: context.TODO(),
	}

	txRsp, err := client.Transactions.GetTransactions(txParams, httptransport.BearerToken(accessToken))
	if err != nil {
		log.Printf("Error getting transactions: %v", err)
	}

	txs := txRsp.GetPayload()
}
```

#### Example POST affordability:
```
import (
	"context"
	client2 "github.com/basiqio/basiq-swagger/dist/client"
    "github.com/basiqio/basiq-swagger/dist/client/affordability"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"log"
)

func main() {

   	client := client2.New(httptransport.New(client2.DefaultHost, client2.DefaultBasePath, nil), strfmt.Default)

	affordabilityPostParams := &affordability.PostAffordabilityParams{
        UserID: userID,
    	Context: context.TODO(),
    }

    affordabilityPostRsp, _, _, err := client.Affordability.PostAffordability(affordabilityPostParams, httptransport.BearerToken(*token.AccessToken))
	
    if err != nil {
		log.Printf("Error getting affordability: %v", err)
	}

	affordabilityReport := affordabilityPostRsp.GetPayload()
}
```
Adding filter params for POST affordability example:
```
   affordabilityPostParams := &affordability.PostAffordabilityParams{
        UserID: userID,
   		AffordabilityPostRequest: &models.AffordabilityPost{FromMonth: "2018-11", ToMonth: "2019-10", Accounts: []string{accountID}},
   		Context:                  context.TODO(),
   	}
```

#### Example GET affordability
```
import (
	"context"
	client2 "github.com/basiqio/basiq-swagger/dist/client"
    "github.com/basiqio/basiq-swagger/dist/client/affordability"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"log"
)

func main() {

   	client := client2.New(httptransport.New(client2.DefaultHost, client2.DefaultBasePath, nil), strfmt.Default)

	affordabilityParams := &affordability.GetAffordabilityParams{
        UserID: userID,
		SnapshotID: snapshotID,
		Context:    context.TODO(),
	}

	affordabilityRsp, err := client.Affordability.GetAffordability(affordabilityParams, httptransport.BearerToken(accessToken))

	if err != nil {
		log.Printf("Error: %v", err)
	}

	affordabilityReport := affordabilityRsp.GetPayload()
}
```


### Changelog
Please check the changelog for more details on API changes


### How to Contribute
Report an issue: https://github.com/basiqio/basiq-swagger/issues/new/choose

### Usage
Link to Swagger file https://github.com/basiqio/basiq-swagger

Links to API docs: https://api.basiq.io/reference