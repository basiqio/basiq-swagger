// Code generated by go-swagger; DO NOT EDIT.

package income

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new income API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for income API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetIncome(params *GetIncomeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIncomeOK, error)

	PostIncome(params *PostIncomeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostIncomeOK, *PostIncomeNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetIncome retrieves the details of an income summary you need only supply the unique income identifier
*/
func (a *Client) GetIncome(params *GetIncomeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetIncomeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIncomeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getIncome",
		Method:             "GET",
		PathPattern:        "/users/{userId}/income/{snapshotId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIncomeReader{formats: a.formats},
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
	success, ok := result.(*GetIncomeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getIncome: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostIncome uses this to create a new income report
*/
func (a *Client) PostIncome(params *PostIncomeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostIncomeOK, *PostIncomeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIncomeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postIncome",
		Method:             "POST",
		PathPattern:        "/users/{userId}/income",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostIncomeReader{formats: a.formats},
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
	case *PostIncomeOK:
		return value, nil, nil
	case *PostIncomeNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for income: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
