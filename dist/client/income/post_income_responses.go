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

// PostIncomeReader is a Reader for the PostIncome structure.
type PostIncomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIncomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIncomeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPostIncomeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIncomeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIncomeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIncomeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIncomeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostIncomeOK creates a PostIncomeOK with default headers values
func NewPostIncomeOK() *PostIncomeOK {
	return &PostIncomeOK{}
}

/*PostIncomeOK handles this case with default header values.

Returns a created income resource, if the operation succeeded.
*/
type PostIncomeOK struct {
	Payload *models.IncomeResponse
}

func (o *PostIncomeOK) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/income][%d] postIncomeOK  %+v", 200, o.Payload)
}

func (o *PostIncomeOK) GetPayload() *models.IncomeResponse {
	return o.Payload
}

func (o *PostIncomeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IncomeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIncomeNoContent creates a PostIncomeNoContent with default headers values
func NewPostIncomeNoContent() *PostIncomeNoContent {
	return &PostIncomeNoContent{}
}

/*PostIncomeNoContent handles this case with default header values.

Returns no content if there are none transactions for the requested period.
*/
type PostIncomeNoContent struct {
}

func (o *PostIncomeNoContent) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/income][%d] postIncomeNoContent ", 204)
}

func (o *PostIncomeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIncomeBadRequest creates a PostIncomeBadRequest with default headers values
func NewPostIncomeBadRequest() *PostIncomeBadRequest {
	return &PostIncomeBadRequest{}
}

/*PostIncomeBadRequest handles this case with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error.
*/
type PostIncomeBadRequest struct {
	Payload *models.BadRequestError
}

func (o *PostIncomeBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/income][%d] postIncomeBadRequest  %+v", 400, o.Payload)
}

func (o *PostIncomeBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *PostIncomeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIncomeForbidden creates a PostIncomeForbidden with default headers values
func NewPostIncomeForbidden() *PostIncomeForbidden {
	return &PostIncomeForbidden{}
}

/*PostIncomeForbidden handles this case with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type PostIncomeForbidden struct {
	Payload *models.ForbiddenAccessError
}

func (o *PostIncomeForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/income][%d] postIncomeForbidden  %+v", 403, o.Payload)
}

func (o *PostIncomeForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *PostIncomeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIncomeNotFound creates a PostIncomeNotFound with default headers values
func NewPostIncomeNotFound() *PostIncomeNotFound {
	return &PostIncomeNotFound{}
}

/*PostIncomeNotFound handles this case with default header values.

Returns error indicating that server can't find requested resource.
*/
type PostIncomeNotFound struct {
	Payload *models.NotFoundError
}

func (o *PostIncomeNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/income][%d] postIncomeNotFound  %+v", 404, o.Payload)
}

func (o *PostIncomeNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *PostIncomeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIncomeInternalServerError creates a PostIncomeInternalServerError with default headers values
func NewPostIncomeInternalServerError() *PostIncomeInternalServerError {
	return &PostIncomeInternalServerError{}
}

/*PostIncomeInternalServerError handles this case with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type PostIncomeInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostIncomeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/income][%d] postIncomeInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIncomeInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostIncomeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
