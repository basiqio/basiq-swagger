// Code generated by go-swagger; DO NOT EDIT.

package connections

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

// NewRefreshConnectionsParams creates a new RefreshConnectionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRefreshConnectionsParams() *RefreshConnectionsParams {
	return &RefreshConnectionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRefreshConnectionsParamsWithTimeout creates a new RefreshConnectionsParams object
// with the ability to set a timeout on a request.
func NewRefreshConnectionsParamsWithTimeout(timeout time.Duration) *RefreshConnectionsParams {
	return &RefreshConnectionsParams{
		timeout: timeout,
	}
}

// NewRefreshConnectionsParamsWithContext creates a new RefreshConnectionsParams object
// with the ability to set a context for a request.
func NewRefreshConnectionsParamsWithContext(ctx context.Context) *RefreshConnectionsParams {
	return &RefreshConnectionsParams{
		Context: ctx,
	}
}

// NewRefreshConnectionsParamsWithHTTPClient creates a new RefreshConnectionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRefreshConnectionsParamsWithHTTPClient(client *http.Client) *RefreshConnectionsParams {
	return &RefreshConnectionsParams{
		HTTPClient: client,
	}
}

/* RefreshConnectionsParams contains all the parameters to send to the API endpoint
   for the refresh connections operation.

   Typically these are written to a http.Request.
*/
type RefreshConnectionsParams struct {

	/* UserID.

	   The identifier of the user.
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the refresh connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RefreshConnectionsParams) WithDefaults() *RefreshConnectionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the refresh connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RefreshConnectionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the refresh connections params
func (o *RefreshConnectionsParams) WithTimeout(timeout time.Duration) *RefreshConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refresh connections params
func (o *RefreshConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refresh connections params
func (o *RefreshConnectionsParams) WithContext(ctx context.Context) *RefreshConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refresh connections params
func (o *RefreshConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refresh connections params
func (o *RefreshConnectionsParams) WithHTTPClient(client *http.Client) *RefreshConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refresh connections params
func (o *RefreshConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the refresh connections params
func (o *RefreshConnectionsParams) WithUserID(userID string) *RefreshConnectionsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the refresh connections params
func (o *RefreshConnectionsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *RefreshConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
