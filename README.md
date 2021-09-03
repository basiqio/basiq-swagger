# Basiq-swagger
Use this repository to generate a client using the Basiq Swagger/Open API spec and `go-swagger`. `go-swagger` brings to the `go` community a complete suite of fully-featured, high-performance, API components to work with a Swagger API: server, client and data model.

## Usage
### Install dependencies
Firstly make sure you have go-swagger installed (https://goswagger.io).

Installation of the `go-swagger` tool:

`go get -u github.com/go-swagger/go-swagger/cmd/swagger`

### Testing
The build script runs tests against the newly generated SDK. 
In the `/test` directory there is a `config.json` file with three properties; 
```json
{
  "ApiKey": "",
  "DeletedApiKey": "",
  "ExpiredToken": ""
}
```
You need to populate these with their relevant values in order for the tests to pass. The sdk will still be generated even if the tests fail, however we strongly encourage you to ensure all tests are passing before proceeding. 

### Build
To generate the sdk, run the build script:  `./build.sh`
The output will be the `dist` directory. 

### Quick start
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
	tokenReq.SetBasiqVersion("2.0")
	auth := httptransport.APIKeyAuth("Authorization", "header", "Basic "+ apiKey)  
	tokenResponse, err := client.Token.PostToken(tokenReq, auth)

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