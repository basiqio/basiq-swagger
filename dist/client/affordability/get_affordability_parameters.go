// Code generated by go-swagger; DO NOT EDIT.

package affordability

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

// NewGetAffordabilityParams creates a new GetAffordabilityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAffordabilityParams() *GetAffordabilityParams {
	return &GetAffordabilityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAffordabilityParamsWithTimeout creates a new GetAffordabilityParams object
// with the ability to set a timeout on a request.
func NewGetAffordabilityParamsWithTimeout(timeout time.Duration) *GetAffordabilityParams {
	return &GetAffordabilityParams{
		timeout: timeout,
	}
}

// NewGetAffordabilityParamsWithContext creates a new GetAffordabilityParams object
// with the ability to set a context for a request.
func NewGetAffordabilityParamsWithContext(ctx context.Context) *GetAffordabilityParams {
	return &GetAffordabilityParams{
		Context: ctx,
	}
}

// NewGetAffordabilityParamsWithHTTPClient creates a new GetAffordabilityParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAffordabilityParamsWithHTTPClient(client *http.Client) *GetAffordabilityParams {
	return &GetAffordabilityParams{
		HTTPClient: client,
	}
}

/* GetAffordabilityParams contains all the parameters to send to the API endpoint
   for the get affordability operation.

   Typically these are written to a http.Request.
*/
type GetAffordabilityParams struct {

	/* SnapshotID.

	   The identifier of the affordability report to be retrieved.
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

// WithDefaults hydrates default values in the get affordability params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAffordabilityParams) WithDefaults() *GetAffordabilityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get affordability params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAffordabilityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get affordability params
func (o *GetAffordabilityParams) WithTimeout(timeout time.Duration) *GetAffordabilityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get affordability params
func (o *GetAffordabilityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get affordability params
func (o *GetAffordabilityParams) WithContext(ctx context.Context) *GetAffordabilityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get affordability params
func (o *GetAffordabilityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get affordability params
func (o *GetAffordabilityParams) WithHTTPClient(client *http.Client) *GetAffordabilityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get affordability params
func (o *GetAffordabilityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSnapshotID adds the snapshotID to the get affordability params
func (o *GetAffordabilityParams) WithSnapshotID(snapshotID string) *GetAffordabilityParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the get affordability params
func (o *GetAffordabilityParams) SetSnapshotID(snapshotID string) {
	o.SnapshotID = snapshotID
}

// WithUserID adds the userID to the get affordability params
func (o *GetAffordabilityParams) WithUserID(userID string) *GetAffordabilityParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get affordability params
func (o *GetAffordabilityParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAffordabilityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
