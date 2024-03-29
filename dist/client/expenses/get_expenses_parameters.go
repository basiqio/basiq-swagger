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
)

// NewGetExpensesParams creates a new GetExpensesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExpensesParams() *GetExpensesParams {
	return &GetExpensesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExpensesParamsWithTimeout creates a new GetExpensesParams object
// with the ability to set a timeout on a request.
func NewGetExpensesParamsWithTimeout(timeout time.Duration) *GetExpensesParams {
	return &GetExpensesParams{
		timeout: timeout,
	}
}

// NewGetExpensesParamsWithContext creates a new GetExpensesParams object
// with the ability to set a context for a request.
func NewGetExpensesParamsWithContext(ctx context.Context) *GetExpensesParams {
	return &GetExpensesParams{
		Context: ctx,
	}
}

// NewGetExpensesParamsWithHTTPClient creates a new GetExpensesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExpensesParamsWithHTTPClient(client *http.Client) *GetExpensesParams {
	return &GetExpensesParams{
		HTTPClient: client,
	}
}

/* GetExpensesParams contains all the parameters to send to the API endpoint
   for the get expenses operation.

   Typically these are written to a http.Request.
*/
type GetExpensesParams struct {

	/* SnapshotID.

	   The identifier of the expenses report to be retrieved.
	*/
	SnapshotID string

	/* UserID.

	   The identifier of the user.
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get expenses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExpensesParams) WithDefaults() *GetExpensesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get expenses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExpensesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get expenses params
func (o *GetExpensesParams) WithTimeout(timeout time.Duration) *GetExpensesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get expenses params
func (o *GetExpensesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get expenses params
func (o *GetExpensesParams) WithContext(ctx context.Context) *GetExpensesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get expenses params
func (o *GetExpensesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get expenses params
func (o *GetExpensesParams) WithHTTPClient(client *http.Client) *GetExpensesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get expenses params
func (o *GetExpensesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotID adds the snapshotID to the get expenses params
func (o *GetExpensesParams) WithSnapshotID(snapshotID string) *GetExpensesParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the get expenses params
func (o *GetExpensesParams) SetSnapshotID(snapshotID string) {
	o.SnapshotID = snapshotID
}

// WithUserID adds the userID to the get expenses params
func (o *GetExpensesParams) WithUserID(userID string) *GetExpensesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get expenses params
func (o *GetExpensesParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetExpensesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param snapshotId
	if err := r.SetPathParam("snapshotId", o.SnapshotID); err != nil {
		return err
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
