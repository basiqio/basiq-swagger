// Code generated by go-swagger; DO NOT EDIT.

package users

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

	"github.com/basiqio/basiq-swagger/dist/models"
)

// NewUpdateUserParams creates a new UpdateUserParams object
// with the default values initialized.
func NewUpdateUserParams() *UpdateUserParams {
	var ()
	return &UpdateUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserParamsWithTimeout creates a new UpdateUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserParamsWithTimeout(timeout time.Duration) *UpdateUserParams {
	var ()
	return &UpdateUserParams{

		timeout: timeout,
	}
}

// NewUpdateUserParamsWithContext creates a new UpdateUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserParamsWithContext(ctx context.Context) *UpdateUserParams {
	var ()
	return &UpdateUserParams{

		Context: ctx,
	}
}

// NewUpdateUserParamsWithHTTPClient creates a new UpdateUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserParamsWithHTTPClient(client *http.Client) *UpdateUserParams {
	var ()
	return &UpdateUserParams{
		HTTPClient: client,
	}
}

/*UpdateUserParams contains all the parameters to send to the API endpoint
for the update user operation typically these are written to a http.Request
*/
type UpdateUserParams struct {

	/*User*/
	User *models.UpdateUser
	/*UserID
	  The identifier of the user to be retrieved.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update user params
func (o *UpdateUserParams) WithTimeout(timeout time.Duration) *UpdateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user params
func (o *UpdateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user params
func (o *UpdateUserParams) WithContext(ctx context.Context) *UpdateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user params
func (o *UpdateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user params
func (o *UpdateUserParams) WithHTTPClient(client *http.Client) *UpdateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user params
func (o *UpdateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the update user params
func (o *UpdateUserParams) WithUser(user *models.UpdateUser) *UpdateUserParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the update user params
func (o *UpdateUserParams) SetUser(user *models.UpdateUser) {
	o.User = user
}

// WithUserID adds the userID to the update user params
func (o *UpdateUserParams) WithUserID(userID string) *UpdateUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update user params
func (o *UpdateUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.User != nil {
		if err := r.SetBodyParam(o.User); err != nil {
			return err
		}
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
