// Code generated by go-swagger; DO NOT EDIT.

package enrich

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/basiqio/basiq-swagger/dist/models"
)

// GetEnrichReader is a Reader for the GetEnrich structure.
type GetEnrichReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnrichReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEnrichOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEnrichBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEnrichOK creates a GetEnrichOK with default headers values
func NewGetEnrichOK() *GetEnrichOK {
	return &GetEnrichOK{}
}

/*
GetEnrichOK describes a response with status code 200, with default header values.

The service enriches transaction data with multiple attributes, returning a transaction classification and three metadata components. The transaction classification will first determine if the transaction is of type payment, transfer, cash-withdrawal, bank-fee etc. The engine then derives merchant information, purchase location and prescribes an industry standard categorisation for each payment transaction.
*/
type GetEnrichOK struct {
	Payload *models.GetEnrichResponse
}

// IsSuccess returns true when this get enrich o k response has a 2xx status code
func (o *GetEnrichOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get enrich o k response has a 3xx status code
func (o *GetEnrichOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get enrich o k response has a 4xx status code
func (o *GetEnrichOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get enrich o k response has a 5xx status code
func (o *GetEnrichOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get enrich o k response a status code equal to that given
func (o *GetEnrichOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEnrichOK) Error() string {
	return fmt.Sprintf("[GET /enrich][%d] getEnrichOK  %+v", 200, o.Payload)
}

func (o *GetEnrichOK) String() string {
	return fmt.Sprintf("[GET /enrich][%d] getEnrichOK  %+v", 200, o.Payload)
}

func (o *GetEnrichOK) GetPayload() *models.GetEnrichResponse {
	return o.Payload
}

func (o *GetEnrichOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetEnrichResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnrichBadRequest creates a GetEnrichBadRequest with default headers values
func NewGetEnrichBadRequest() *GetEnrichBadRequest {
	return &GetEnrichBadRequest{}
}

/*
GetEnrichBadRequest describes a response with status code 400, with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error.
*/
type GetEnrichBadRequest struct {
	Payload *models.BadRequestError
}

// IsSuccess returns true when this get enrich bad request response has a 2xx status code
func (o *GetEnrichBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get enrich bad request response has a 3xx status code
func (o *GetEnrichBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get enrich bad request response has a 4xx status code
func (o *GetEnrichBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get enrich bad request response has a 5xx status code
func (o *GetEnrichBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get enrich bad request response a status code equal to that given
func (o *GetEnrichBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetEnrichBadRequest) Error() string {
	return fmt.Sprintf("[GET /enrich][%d] getEnrichBadRequest  %+v", 400, o.Payload)
}

func (o *GetEnrichBadRequest) String() string {
	return fmt.Sprintf("[GET /enrich][%d] getEnrichBadRequest  %+v", 400, o.Payload)
}

func (o *GetEnrichBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetEnrichBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
