// Code generated by go-swagger; DO NOT EDIT.

package expenses

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

// NewPostExpensesParams creates a new PostExpensesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostExpensesParams() *PostExpensesParams {
	return &PostExpensesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostExpensesParamsWithTimeout creates a new PostExpensesParams object
// with the ability to set a timeout on a request.
func NewPostExpensesParamsWithTimeout(timeout time.Duration) *PostExpensesParams {
	return &PostExpensesParams{
		timeout: timeout,
	}
}

// NewPostExpensesParamsWithContext creates a new PostExpensesParams object
// with the ability to set a context for a request.
func NewPostExpensesParamsWithContext(ctx context.Context) *PostExpensesParams {
	return &PostExpensesParams{
		Context: ctx,
	}
}

// NewPostExpensesParamsWithHTTPClient creates a new PostExpensesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostExpensesParamsWithHTTPClient(client *http.Client) *PostExpensesParams {
	return &PostExpensesParams{
		HTTPClient: client,
	}
}

/*
PostExpensesParams contains all the parameters to send to the API endpoint

	for the post expenses operation.

	Typically these are written to a http.Request.
*/
type PostExpensesParams struct {

	// ExpensesPostRequest.
	ExpensesPostRequest *models.ExpensesPostRequest

	/* UserID.

	   The identifier of the user.
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post expenses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExpensesParams) WithDefaults() *PostExpensesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post expenses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostExpensesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post expenses params
func (o *PostExpensesParams) WithTimeout(timeout time.Duration) *PostExpensesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post expenses params
func (o *PostExpensesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post expenses params
func (o *PostExpensesParams) WithContext(ctx context.Context) *PostExpensesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post expenses params
func (o *PostExpensesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post expenses params
func (o *PostExpensesParams) WithHTTPClient(client *http.Client) *PostExpensesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post expenses params
func (o *PostExpensesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpensesPostRequest adds the expensesPostRequest to the post expenses params
func (o *PostExpensesParams) WithExpensesPostRequest(expensesPostRequest *models.ExpensesPostRequest) *PostExpensesParams {
	o.SetExpensesPostRequest(expensesPostRequest)
	return o
}

// SetExpensesPostRequest adds the expensesPostRequest to the post expenses params
func (o *PostExpensesParams) SetExpensesPostRequest(expensesPostRequest *models.ExpensesPostRequest) {
	o.ExpensesPostRequest = expensesPostRequest
}

// WithUserID adds the userID to the post expenses params
func (o *PostExpensesParams) WithUserID(userID string) *PostExpensesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the post expenses params
func (o *PostExpensesParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PostExpensesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ExpensesPostRequest != nil {
		if err := r.SetBodyParam(o.ExpensesPostRequest); err != nil {
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
