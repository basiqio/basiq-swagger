// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new jobs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for jobs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetJobs(params *GetJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetJobsOK, error)

	GetUserJobs(params *GetUserJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserJobsOK, error)

	PostJobMfa(params *PostJobMfaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostJobMfaAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetJobs retrieves the details of an existing job you need only supply the unique job identifier that was returned upon job creation

	<b>Tracking the status of a job</b><br/>

Every step of the job has a status property that depicts its current state.<br/>
<b>Find out what steps have been completed</b><br/>
Depending on the job being executed, some jobs will have multiple steps which need to be executed, for e.g. refreshing a connection requires the following steps to be completed:
<ol><li>Establish successful authentication with institution</li>
<li>Optional: mfa-challenge only appears in MFA challenge flow.</li>
<li>Fetch latest list of accounts</li>
<li>Fetch latest list of transactions</li></ol>
You can keep track of the steps that have been completed by observing the results array property. As each step is successfully completed, its status will be updated and a result object with the link to the affected resource will be present. In the event that a step has failed, the result object will contain an embedded error object.
*/
func (a *Client) GetJobs(params *GetJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getJobs",
		Method:             "GET",
		PathPattern:        "/jobs/{jobId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJobsReader{formats: a.formats},
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
	success, ok := result.(*GetJobsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getJobs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetUserJobs retrieves the details of an existing jobs you need only supply the unique user identifier

	<b>Tracking the status of a job</b><br/>

Every step of the job has a status property that depicts its current state.<br/>
<b>Find out what steps have been completed</b><br/>
Depending on the job being executed, some jobs will have multiple steps which need to be executed, for e.g. refreshing a connection requires the following steps to be completed:
<ol><li>Establish successful authentication with institution</li>
<li>Fetch latest list of accounts</li>
<li>Fetch latest list of transactions</li></ol>
You can keep track of the steps that have been completed by observing the results array property. As each step is successfully completed, its status will be updated and a result object with the link to the affected resource will be present. In the event that a step has failed, the result object will contain an embedded error object.
*/
func (a *Client) GetUserJobs(params *GetUserJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserJobsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserJobs",
		Method:             "GET",
		PathPattern:        "/users/{userId}/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserJobsReader{formats: a.formats},
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
	success, ok := result.(*GetUserJobsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserJobs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PostJobMfa uses this to create a new m f a challenge response to job step mfa challenge

	Ensure that you generate an authentication token with

scope = CLIENT_ACCESS and basiq-version = 2.1 to create this resource
*/
func (a *Client) PostJobMfa(params *PostJobMfaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostJobMfaAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostJobMfaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postJobMfa",
		Method:             "POST",
		PathPattern:        "/jobs/{jobId}/mfa",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostJobMfaReader{formats: a.formats},
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
	success, ok := result.(*PostJobMfaAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postJobMfa: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
