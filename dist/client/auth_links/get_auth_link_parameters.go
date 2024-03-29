// Code generated by go-swagger; DO NOT EDIT.

package auth_links

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

// NewGetAuthLinkParams creates a new GetAuthLinkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAuthLinkParams() *GetAuthLinkParams {
	return &GetAuthLinkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthLinkParamsWithTimeout creates a new GetAuthLinkParams object
// with the ability to set a timeout on a request.
func NewGetAuthLinkParamsWithTimeout(timeout time.Duration) *GetAuthLinkParams {
	return &GetAuthLinkParams{
		timeout: timeout,
	}
}

// NewGetAuthLinkParamsWithContext creates a new GetAuthLinkParams object
// with the ability to set a context for a request.
func NewGetAuthLinkParamsWithContext(ctx context.Context) *GetAuthLinkParams {
	return &GetAuthLinkParams{
		Context: ctx,
	}
}

// NewGetAuthLinkParamsWithHTTPClient creates a new GetAuthLinkParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAuthLinkParamsWithHTTPClient(client *http.Client) *GetAuthLinkParams {
	return &GetAuthLinkParams{
		HTTPClient: client,
	}
}

/* GetAuthLinkParams contains all the parameters to send to the API endpoint
   for the get auth link operation.

   Typically these are written to a http.Request.
*/
type GetAuthLinkParams struct {

	/* UserID.

	   The identifier of the user.
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get auth link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthLinkParams) WithDefaults() *GetAuthLinkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get auth link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthLinkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get auth link params
func (o *GetAuthLinkParams) WithTimeout(timeout time.Duration) *GetAuthLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auth link params
func (o *GetAuthLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auth link params
func (o *GetAuthLinkParams) WithContext(ctx context.Context) *GetAuthLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auth link params
func (o *GetAuthLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auth link params
func (o *GetAuthLinkParams) WithHTTPClient(client *http.Client) *GetAuthLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auth link params
func (o *GetAuthLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get auth link params
func (o *GetAuthLinkParams) WithUserID(userID string) *GetAuthLinkParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get auth link params
func (o *GetAuthLinkParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
