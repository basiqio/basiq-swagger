// Code generated by go-swagger; DO NOT EDIT.

package affordability

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new affordability API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for affordability API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAffordability(params *GetAffordabilityParams, authInfo runtime.ClientAuthInfoWriter) (*GetAffordabilityOK, error)

	PostAffordability(params *PostAffordabilityParams, authInfo runtime.ClientAuthInfoWriter) (*PostAffordabilityOK, *PostAffordabilityCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAffordability retrieves the details of an affordability summary you need only supply the unique affordability identifier
*/
func (a *Client) GetAffordability(params *GetAffordabilityParams, authInfo runtime.ClientAuthInfoWriter) (*GetAffordabilityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAffordabilityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAffordability",
		Method:             "GET",
		PathPattern:        "/users/{userId}/affordability/{snapshotId}",
		ProducesMediaTypes: []string{"application/json", "application/pdf"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAffordabilityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAffordabilityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAffordability: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAffordability uses this to create a new affordability report
*/
func (a *Client) PostAffordability(params *PostAffordabilityParams, authInfo runtime.ClientAuthInfoWriter) (*PostAffordabilityOK, *PostAffordabilityCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAffordabilityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postAffordability",
		Method:             "POST",
		PathPattern:        "/users/{userId}/affordability",
		ProducesMediaTypes: []string{"application/json", "application/pdf"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAffordabilityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostAffordabilityOK:
		return value, nil, nil
	case *PostAffordabilityCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for affordability: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
