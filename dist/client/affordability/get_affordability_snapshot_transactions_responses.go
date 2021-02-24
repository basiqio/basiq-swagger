// Code generated by go-swagger; DO NOT EDIT.

package affordability

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/basiqio/basiq-swagger/dist/models"
)

// GetAffordabilitySnapshotTransactionsReader is a Reader for the GetAffordabilitySnapshotTransactions structure.
type GetAffordabilitySnapshotTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAffordabilitySnapshotTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAffordabilitySnapshotTransactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAffordabilitySnapshotTransactionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAffordabilitySnapshotTransactionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAffordabilitySnapshotTransactionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAffordabilitySnapshotTransactionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAffordabilitySnapshotTransactionsOK creates a GetAffordabilitySnapshotTransactionsOK with default headers values
func NewGetAffordabilitySnapshotTransactionsOK() *GetAffordabilitySnapshotTransactionsOK {
	return &GetAffordabilitySnapshotTransactionsOK{}
}

/*GetAffordabilitySnapshotTransactionsOK handles this case with default header values.

Returns a affordability resource transactions JSON or CSV file, if the operation succeeded.
*/
type GetAffordabilitySnapshotTransactionsOK struct {
	Payload *models.AffordabilityTransactionsResponse
}

func (o *GetAffordabilitySnapshotTransactionsOK) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/affordability/{snapshotId}/transactions][%d] getAffordabilitySnapshotTransactionsOK  %+v", 200, o.Payload)
}

func (o *GetAffordabilitySnapshotTransactionsOK) GetPayload() *models.AffordabilityTransactionsResponse {
	return o.Payload
}

func (o *GetAffordabilitySnapshotTransactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AffordabilityTransactionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAffordabilitySnapshotTransactionsBadRequest creates a GetAffordabilitySnapshotTransactionsBadRequest with default headers values
func NewGetAffordabilitySnapshotTransactionsBadRequest() *GetAffordabilitySnapshotTransactionsBadRequest {
	return &GetAffordabilitySnapshotTransactionsBadRequest{}
}

/*GetAffordabilitySnapshotTransactionsBadRequest handles this case with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error.
*/
type GetAffordabilitySnapshotTransactionsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetAffordabilitySnapshotTransactionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/affordability/{snapshotId}/transactions][%d] getAffordabilitySnapshotTransactionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAffordabilitySnapshotTransactionsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetAffordabilitySnapshotTransactionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAffordabilitySnapshotTransactionsForbidden creates a GetAffordabilitySnapshotTransactionsForbidden with default headers values
func NewGetAffordabilitySnapshotTransactionsForbidden() *GetAffordabilitySnapshotTransactionsForbidden {
	return &GetAffordabilitySnapshotTransactionsForbidden{}
}

/*GetAffordabilitySnapshotTransactionsForbidden handles this case with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type GetAffordabilitySnapshotTransactionsForbidden struct {
	Payload *models.ForbiddenAccessError
}

func (o *GetAffordabilitySnapshotTransactionsForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/affordability/{snapshotId}/transactions][%d] getAffordabilitySnapshotTransactionsForbidden  %+v", 403, o.Payload)
}

func (o *GetAffordabilitySnapshotTransactionsForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *GetAffordabilitySnapshotTransactionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAffordabilitySnapshotTransactionsNotFound creates a GetAffordabilitySnapshotTransactionsNotFound with default headers values
func NewGetAffordabilitySnapshotTransactionsNotFound() *GetAffordabilitySnapshotTransactionsNotFound {
	return &GetAffordabilitySnapshotTransactionsNotFound{}
}

/*GetAffordabilitySnapshotTransactionsNotFound handles this case with default header values.

Returns error indicating that server can't find requested resource.
*/
type GetAffordabilitySnapshotTransactionsNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetAffordabilitySnapshotTransactionsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/affordability/{snapshotId}/transactions][%d] getAffordabilitySnapshotTransactionsNotFound  %+v", 404, o.Payload)
}

func (o *GetAffordabilitySnapshotTransactionsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetAffordabilitySnapshotTransactionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAffordabilitySnapshotTransactionsInternalServerError creates a GetAffordabilitySnapshotTransactionsInternalServerError with default headers values
func NewGetAffordabilitySnapshotTransactionsInternalServerError() *GetAffordabilitySnapshotTransactionsInternalServerError {
	return &GetAffordabilitySnapshotTransactionsInternalServerError{}
}

/*GetAffordabilitySnapshotTransactionsInternalServerError handles this case with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type GetAffordabilitySnapshotTransactionsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetAffordabilitySnapshotTransactionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/affordability/{snapshotId}/transactions][%d] getAffordabilitySnapshotTransactionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAffordabilitySnapshotTransactionsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetAffordabilitySnapshotTransactionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}