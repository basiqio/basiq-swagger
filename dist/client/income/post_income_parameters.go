// Code generated by go-swagger; DO NOT EDIT.

package income

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

// NewPostIncomeParams creates a new PostIncomeParams object
// with the default values initialized.
func NewPostIncomeParams() *PostIncomeParams {
	var ()
	return &PostIncomeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIncomeParamsWithTimeout creates a new PostIncomeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIncomeParamsWithTimeout(timeout time.Duration) *PostIncomeParams {
	var ()
	return &PostIncomeParams{

		timeout: timeout,
	}
}

// NewPostIncomeParamsWithContext creates a new PostIncomeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIncomeParamsWithContext(ctx context.Context) *PostIncomeParams {
	var ()
	return &PostIncomeParams{

		Context: ctx,
	}
}

// NewPostIncomeParamsWithHTTPClient creates a new PostIncomeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIncomeParamsWithHTTPClient(client *http.Client) *PostIncomeParams {
	var ()
	return &PostIncomeParams{
		HTTPClient: client,
	}
}

/*PostIncomeParams contains all the parameters to send to the API endpoint
for the post income operation typically these are written to a http.Request
*/
type PostIncomeParams struct {

	/*IncomePostRequest*/
	IncomePostRequest *models.IncomePostRequest
	/*UserID
	  The identifier of the user.

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post income params
func (o *PostIncomeParams) WithTimeout(timeout time.Duration) *PostIncomeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post income params
func (o *PostIncomeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post income params
func (o *PostIncomeParams) WithContext(ctx context.Context) *PostIncomeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post income params
func (o *PostIncomeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post income params
func (o *PostIncomeParams) WithHTTPClient(client *http.Client) *PostIncomeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post income params
func (o *PostIncomeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncomePostRequest adds the incomePostRequest to the post income params
func (o *PostIncomeParams) WithIncomePostRequest(incomePostRequest *models.IncomePostRequest) *PostIncomeParams {
	o.SetIncomePostRequest(incomePostRequest)
	return o
}

// SetIncomePostRequest adds the incomePostRequest to the post income params
func (o *PostIncomeParams) SetIncomePostRequest(incomePostRequest *models.IncomePostRequest) {
	o.IncomePostRequest = incomePostRequest
}

// WithUserID adds the userID to the post income params
func (o *PostIncomeParams) WithUserID(userID string) *PostIncomeParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the post income params
func (o *PostIncomeParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PostIncomeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncomePostRequest != nil {
		if err := r.SetBodyParam(o.IncomePostRequest); err != nil {
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
