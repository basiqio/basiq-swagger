// Code generated by go-swagger; DO NOT EDIT.

package transactions

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
	"github.com/go-openapi/swag"
)

// NewGetTransactionsParams creates a new GetTransactionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTransactionsParams() *GetTransactionsParams {
	return &GetTransactionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTransactionsParamsWithTimeout creates a new GetTransactionsParams object
// with the ability to set a timeout on a request.
func NewGetTransactionsParamsWithTimeout(timeout time.Duration) *GetTransactionsParams {
	return &GetTransactionsParams{
		timeout: timeout,
	}
}

// NewGetTransactionsParamsWithContext creates a new GetTransactionsParams object
// with the ability to set a context for a request.
func NewGetTransactionsParamsWithContext(ctx context.Context) *GetTransactionsParams {
	return &GetTransactionsParams{
		Context: ctx,
	}
}

// NewGetTransactionsParamsWithHTTPClient creates a new GetTransactionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTransactionsParamsWithHTTPClient(client *http.Client) *GetTransactionsParams {
	return &GetTransactionsParams{
		HTTPClient: client,
	}
}

/* GetTransactionsParams contains all the parameters to send to the API endpoint
   for the get transactions operation.

   Typically these are written to a http.Request.
*/
type GetTransactionsParams struct {

	/* Filter.

	   Transaction filters.
	*/
	Filter *string

	/* Limit.

	   This represents the maximum number of items that may be included in the response (maximum of 500). Note that by default 500 items are returned if this value is not specified.

	   Format: int64
	   Default: 500
	*/
	Limit *int64

	/* UserID.

	   User identifier.
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTransactionsParams) WithDefaults() *GetTransactionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTransactionsParams) SetDefaults() {
	var (
		limitDefault = int64(500)
	)

	val := GetTransactionsParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get transactions params
func (o *GetTransactionsParams) WithTimeout(timeout time.Duration) *GetTransactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get transactions params
func (o *GetTransactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get transactions params
func (o *GetTransactionsParams) WithContext(ctx context.Context) *GetTransactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get transactions params
func (o *GetTransactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get transactions params
func (o *GetTransactionsParams) WithHTTPClient(client *http.Client) *GetTransactionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get transactions params
func (o *GetTransactionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the get transactions params
func (o *GetTransactionsParams) WithFilter(filter *string) *GetTransactionsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get transactions params
func (o *GetTransactionsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the get transactions params
func (o *GetTransactionsParams) WithLimit(limit *int64) *GetTransactionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get transactions params
func (o *GetTransactionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithUserID adds the userID to the get transactions params
func (o *GetTransactionsParams) WithUserID(userID string) *GetTransactionsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get transactions params
func (o *GetTransactionsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTransactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
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
