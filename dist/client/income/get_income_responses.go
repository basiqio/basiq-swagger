// Code generated by go-swagger; DO NOT EDIT.

package income

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/basiqio/basiq-swagger/dist/models"
)

// GetIncomeReader is a Reader for the GetIncome structure.
type GetIncomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIncomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIncomeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIncomeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIncomeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIncomeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIncomeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIncomeOK creates a GetIncomeOK with default headers values
func NewGetIncomeOK() *GetIncomeOK {
	return &GetIncomeOK{}
}

/*GetIncomeOK handles this case with default header values.

Returns a saved income resource, if the operation succeeded.
*/
type GetIncomeOK struct {
	Payload *models.IncomeResponse
}

func (o *GetIncomeOK) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/income/{snapshotId}][%d] getIncomeOK  %+v", 200, o.Payload)
}

func (o *GetIncomeOK) GetPayload() *models.IncomeResponse {
	return o.Payload
}

func (o *GetIncomeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncomeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncomeBadRequest creates a GetIncomeBadRequest with default headers values
func NewGetIncomeBadRequest() *GetIncomeBadRequest {
	return &GetIncomeBadRequest{}
}

/*GetIncomeBadRequest handles this case with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error.
*/
type GetIncomeBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetIncomeBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/income/{snapshotId}][%d] getIncomeBadRequest  %+v", 400, o.Payload)
}

func (o *GetIncomeBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetIncomeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncomeForbidden creates a GetIncomeForbidden with default headers values
func NewGetIncomeForbidden() *GetIncomeForbidden {
	return &GetIncomeForbidden{}
}

/*GetIncomeForbidden handles this case with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type GetIncomeForbidden struct {
	Payload *models.ForbiddenAccessError
}

func (o *GetIncomeForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/income/{snapshotId}][%d] getIncomeForbidden  %+v", 403, o.Payload)
}

func (o *GetIncomeForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *GetIncomeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncomeNotFound creates a GetIncomeNotFound with default headers values
func NewGetIncomeNotFound() *GetIncomeNotFound {
	return &GetIncomeNotFound{}
}

/*GetIncomeNotFound handles this case with default header values.

Returns error indicating that server can't find requested resource.
*/
type GetIncomeNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetIncomeNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/income/{snapshotId}][%d] getIncomeNotFound  %+v", 404, o.Payload)
}

func (o *GetIncomeNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetIncomeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIncomeInternalServerError creates a GetIncomeInternalServerError with default headers values
func NewGetIncomeInternalServerError() *GetIncomeInternalServerError {
	return &GetIncomeInternalServerError{}
}

/*GetIncomeInternalServerError handles this case with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type GetIncomeInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetIncomeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/income/{snapshotId}][%d] getIncomeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIncomeInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetIncomeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
