// Code generated by go-swagger; DO NOT EDIT.

package identities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetUserIdentitiesParams creates a new GetUserIdentitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserIdentitiesParams() *GetUserIdentitiesParams {
	return &GetUserIdentitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserIdentitiesParamsWithTimeout creates a new GetUserIdentitiesParams object
// with the ability to set a timeout on a request.
func NewGetUserIdentitiesParamsWithTimeout(timeout time.Duration) *GetUserIdentitiesParams {
	return &GetUserIdentitiesParams{
		timeout: timeout,
	}
}

// NewGetUserIdentitiesParamsWithContext creates a new GetUserIdentitiesParams object
// with the ability to set a context for a request.
func NewGetUserIdentitiesParamsWithContext(ctx context.Context) *GetUserIdentitiesParams {
	return &GetUserIdentitiesParams{
		Context: ctx,
	}
}

// NewGetUserIdentitiesParamsWithHTTPClient creates a new GetUserIdentitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserIdentitiesParamsWithHTTPClient(client *http.Client) *GetUserIdentitiesParams {
	return &GetUserIdentitiesParams{
		HTTPClient: client,
	}
}

/*
GetUserIdentitiesParams contains all the parameters to send to the API endpoint

	for the get user identities operation.

	Typically these are written to a http.Request.
*/
type GetUserIdentitiesParams struct {

	/* UserID.

	   User identifier.
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user identities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserIdentitiesParams) WithDefaults() *GetUserIdentitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user identities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserIdentitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user identities params
func (o *GetUserIdentitiesParams) WithTimeout(timeout time.Duration) *GetUserIdentitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user identities params
func (o *GetUserIdentitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user identities params
func (o *GetUserIdentitiesParams) WithContext(ctx context.Context) *GetUserIdentitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user identities params
func (o *GetUserIdentitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user identities params
func (o *GetUserIdentitiesParams) WithHTTPClient(client *http.Client) *GetUserIdentitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user identities params
func (o *GetUserIdentitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get user identities params
func (o *GetUserIdentitiesParams) WithUserID(userID string) *GetUserIdentitiesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user identities params
func (o *GetUserIdentitiesParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserIdentitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
