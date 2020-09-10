// Code generated by go-swagger; DO NOT EDIT.

package statements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new statements API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for statements API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateStatement(params *CreateStatementParams, authInfo runtime.ClientAuthInfoWriter) (*CreateStatementAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateStatement as user can choose to share their financial data by uploading official pdf bank statements instead of creating a bank connection once the statement object is successfully created you can use it to obtain the user s latest financial data extracted from the bank statement i e accounts and transactions

  <blockquote>The endpoint also accepts csv files conforming to our file specification. Contact us directly for more details.</blockquote>

Create a new statement by uploading an official pdf bank statement or csv file statement. When a new statement request is made, the server will create a job that will process the following steps:
<table>
<thead><tr><td>#</td><td>Step</td><td>Description</td></tr></thead>
<tbody>
<tr><td>1</td><td>verify-credentials</td><td>The server will verify the file, validate the statement layout and attempt to parse the target statement</td></tr>
<tr><td>2</td><td>retrieve-accounts</td><td>The server will retrieve the complete list of accounts and their details e.g. account number, name and balances</td></tr>
<tr><td>3</td><td>retrieve-transactions</td><td>The server will fetch the associated transactions for each of the accounts</td></tr>
</tbody>
</table>

You can check the status of each step by querying the job resource (returned when the statement is created).
*/
func (a *Client) CreateStatement(params *CreateStatementParams, authInfo runtime.ClientAuthInfoWriter) (*CreateStatementAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateStatementParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createStatement",
		Method:             "POST",
		PathPattern:        "/users/{userId}/statements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateStatementReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateStatementAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createStatement: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}