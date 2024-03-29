// Code generated by go-swagger; DO NOT EDIT.

package auth_links

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new auth links API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for auth links API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAuthLink(params *DeleteAuthLinkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAuthLinkNoContent, error)

	GetAuthLink(params *GetAuthLinkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAuthLinkOK, error)

	PostAuthLink(params *PostAuthLinkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAuthLinkCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAuthLink permanentlies deletes an auth link resource once deleted the URL associated with the deleted object will no longer be valid

  <blockquote>Note that this action cannot be undone.</blockquote>

<blockquote>The auth_link is a URL that directs a User to Basiq's hosted consent workflow to link banks and securely share data. When the user selects 'I have disclosed all my accounts' the auth_link is automatically deleted.</blockquote>

Returns an empty body if the delete succeeded. Otherwise, this call returns an error in the event of a failure.
*/
func (a *Client) DeleteAuthLink(params *DeleteAuthLinkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAuthLinkNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAuthLinkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAuthLink",
		Method:             "DELETE",
		PathPattern:        "/users/{userId}/auth_link",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAuthLinkReader{formats: a.formats},
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
	success, ok := result.(*DeleteAuthLinkNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAuthLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAuthLink uses this to retrieve the latest last auth link generated for the specified user

  Returns the latest/last auth_link generated for the specified user. Returns an error otherwise.
*/
func (a *Client) GetAuthLink(params *GetAuthLinkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAuthLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuthLinkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAuthLink",
		Method:             "GET",
		PathPattern:        "/users/{userId}/auth_link",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAuthLinkReader{formats: a.formats},
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
	success, ok := result.(*GetAuthLinkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAuthLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAuthLink ans auth link object can be generated to securely capture data from a user using the URL allows data to be captured via the hosted basiq connect consent workflow for a given user

  Create a new auth_link object by making a POST request to the auth_link endpoint. The new auth_link will effectively delete previous auth-links for that User/applicant, rendering the previous URL(s) invalid. The 'mobile' attribute is optional. If it is specified this number will take preference over the User object mobile number for 2FA SMS verification.

Returns a created auth_link resource, if the operation succeeded. Returns an error if the post failed (e.g. not supplying required properties).
*/
func (a *Client) PostAuthLink(params *PostAuthLinkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAuthLinkCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAuthLinkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postAuthLink",
		Method:             "POST",
		PathPattern:        "/users/{userId}/auth_link",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAuthLinkReader{formats: a.formats},
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
	success, ok := result.(*PostAuthLinkCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postAuthLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
