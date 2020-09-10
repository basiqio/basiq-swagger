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

// NewDeleteAuthLinkParams creates a new DeleteAuthLinkParams object
// with the default values initialized.
func NewDeleteAuthLinkParams() *DeleteAuthLinkParams {
	var ()
	return &DeleteAuthLinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAuthLinkParamsWithTimeout creates a new DeleteAuthLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAuthLinkParamsWithTimeout(timeout time.Duration) *DeleteAuthLinkParams {
	var ()
	return &DeleteAuthLinkParams{

		timeout: timeout,
	}
}

// NewDeleteAuthLinkParamsWithContext creates a new DeleteAuthLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAuthLinkParamsWithContext(ctx context.Context) *DeleteAuthLinkParams {
	var ()
	return &DeleteAuthLinkParams{

		Context: ctx,
	}
}

// NewDeleteAuthLinkParamsWithHTTPClient creates a new DeleteAuthLinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAuthLinkParamsWithHTTPClient(client *http.Client) *DeleteAuthLinkParams {
	var ()
	return &DeleteAuthLinkParams{
		HTTPClient: client,
	}
}

/*DeleteAuthLinkParams contains all the parameters to send to the API endpoint
for the delete auth link operation typically these are written to a http.Request
*/
type DeleteAuthLinkParams struct {

	/*UserID
	  The identifier of the user.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete auth link params
func (o *DeleteAuthLinkParams) WithTimeout(timeout time.Duration) *DeleteAuthLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete auth link params
func (o *DeleteAuthLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete auth link params
func (o *DeleteAuthLinkParams) WithContext(ctx context.Context) *DeleteAuthLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete auth link params
func (o *DeleteAuthLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete auth link params
func (o *DeleteAuthLinkParams) WithHTTPClient(client *http.Client) *DeleteAuthLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete auth link params
func (o *DeleteAuthLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the delete auth link params
func (o *DeleteAuthLinkParams) WithUserID(userID string) *DeleteAuthLinkParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete auth link params
func (o *DeleteAuthLinkParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAuthLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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