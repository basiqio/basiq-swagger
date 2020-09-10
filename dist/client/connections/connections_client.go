// Code generated by go-swagger; DO NOT EDIT.

package connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new connections API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for connections API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteConnection(params *DeleteConnectionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteConnectionNoContent, error)

	GetConnection(params *GetConnectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetConnectionOK, error)

	GetConnections(params *GetConnectionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetConnectionsOK, error)

	PostConnection(params *PostConnectionParams, authInfo runtime.ClientAuthInfoWriter) (*PostConnectionAccepted, error)

	RefreshConnection(params *RefreshConnectionParams, authInfo runtime.ClientAuthInfoWriter) (*RefreshConnectionAccepted, error)

	RefreshConnections(params *RefreshConnectionsParams, authInfo runtime.ClientAuthInfoWriter) (*RefreshConnectionsAccepted, error)

	UpdateConnection(params *UpdateConnectionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateConnectionAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteConnection permanentlies deletes a connection

  Once the connection has been deleted, all of the associated financial data e.g. accounts and transactions can still be accessed via the users end-point.
Note that this action cannot be undone.
*/
func (a *Client) DeleteConnection(params *DeleteConnectionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteConnectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteConnectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteConnection",
		Method:             "DELETE",
		PathPattern:        "/users/{userId}/connections/{connectionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteConnectionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetConnection uses this to retrieve details of a specific connection

  This request will return back a connection object with most of the fields that were submitted when the connection was first created. The connection object will also return a list of URLs to the associated account, transaction and institution objects.
The status property of the connection object identifies the state of the connection. Use this to work out if the connection is still valid, or whether to take further action (e.g. if the connection credentials are no longer valid you may ask the user to re-submit their details).
<br/>
Note that due to security the loginId, password, securityCode are never returned.
Profile data represents the name, phone, email and address of the logged in user or data sharer. Only data made available by institution can be returned. An institution may offer the option for a customer to hide all personal data or add 2FA to access the data - in this case no data would be returned for all data points. Note, that when a Connection is deleted - the profile data will also be deleted. If a phone number or email address is masked by the institution - the string will be shown exactly as it is provided by the institution.
*/
func (a *Client) GetConnection(params *GetConnectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConnectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getConnection",
		Method:             "GET",
		PathPattern:        "/users/{userId}/connections/{connectionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetConnections uses this to retrieve details of the connections

  Returns a list with a data property that contains an array of connections. Each entry in the array is a separate object. If no data is returned, the resulting array will be empty.
*/
func (a *Client) GetConnections(params *GetConnectionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetConnectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConnectionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getConnections",
		Method:             "GET",
		PathPattern:        "/users/{userId}/connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetConnectionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConnectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getConnections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostConnection uses this to create a new connection

  When a new connection request is made the server will create a job that will process the following steps:
<table>
<thead>
<tr><td>#</td><td>Step</td><td>Description</td></tr>
</thead>
<tr>
<td>1</td>
<td>
verify-credentials
</td>
<td>
The server will attempt to authenticate with the target institution using the supplied credentials.
</td>
<tr>
<td>2</td>
<td>
retrieve-accounts
</td>
<td>
The server will retrieve the complete list of accounts and their details e.g. account number, name and balances.
</td>
<tr>
<td>3</td>
<td>
retrieve-transactions
</td>
<td>
The server will fetch the associated transactions for each of the accounts.
</td>
</tr>
</table>
<br/>
Note that the time it takes to complete the processes above will vary depending on the volume of data along with the general latency between our servers and the financial institution. As a rough guide this entire process could take anywhere between 3 - 30 secs.
*/
func (a *Client) PostConnection(params *PostConnectionParams, authInfo runtime.ClientAuthInfoWriter) (*PostConnectionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostConnectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postConnection",
		Method:             "POST",
		PathPattern:        "/users/{userId}/connections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostConnectionAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RefreshConnection uses this to retrieve the latest financial data

  Similar to when a connection is first created, the refresh resource will initiate the following series of steps to retrieve the latest financial data from the target institution:
<table>
<thead>
<tr><td>#</td><td>Step</td><td>Description</td></tr>
</thead>
<tr><td>1</td><td>verify-credentials</td><td>The server will attempt to authenticate with the target institution using the supplied credentials.</td></tr>
<tr><td>2</td><td>retrieve-accounts</td><td>
The server will retrieve the complete list of accounts and their details e.g. account number, name and balances.</td></tr>
<tr><td>3</td><td>retrieve-transactions</td><td>The server will fetch the associated transactions for each of the accounts.</td></tr>
</table>
*/
func (a *Client) RefreshConnection(params *RefreshConnectionParams, authInfo runtime.ClientAuthInfoWriter) (*RefreshConnectionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRefreshConnectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "refreshConnection",
		Method:             "POST",
		PathPattern:        "/users/{userId}/connections/{connectionId}/refresh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RefreshConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RefreshConnectionAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for refreshConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RefreshConnections uses this to refresh of all connections
*/
func (a *Client) RefreshConnections(params *RefreshConnectionsParams, authInfo runtime.ClientAuthInfoWriter) (*RefreshConnectionsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRefreshConnectionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "refreshConnections",
		Method:             "POST",
		PathPattern:        "/users/{userId}/connections/refresh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RefreshConnectionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RefreshConnectionsAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for refreshConnections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateConnection uses this to update the details of a specific connection
*/
func (a *Client) UpdateConnection(params *UpdateConnectionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateConnectionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateConnectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateConnection",
		Method:             "POST",
		PathPattern:        "/users/{userId}/connections/{connectionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateConnectionAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}