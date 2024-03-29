// Code generated by go-swagger; DO NOT EDIT.

package expenses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new expenses API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for expenses API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetExpenses(params *GetExpensesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExpensesOK, error)

	PostExpenses(params *PostExpensesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostExpensesOK, *PostExpensesNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetExpenses retrieves the details of an expenses summary you need only supply the unique expenses identifier
*/
func (a *Client) GetExpenses(params *GetExpensesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExpensesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExpensesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExpenses",
		Method:             "GET",
		PathPattern:        "/users/{userId}/expenses/{snapshotId}",
		ProducesMediaTypes: []string{"application/json", "text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExpensesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetExpensesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExpenses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostExpenses uses this to create a new expenses report
*/
func (a *Client) PostExpenses(params *PostExpensesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostExpensesOK, *PostExpensesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostExpensesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postExpenses",
		Method:             "POST",
		PathPattern:        "/users/{userId}/expenses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostExpensesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostExpensesOK:
		return value, nil, nil
	case *PostExpensesNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for expenses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
