// Code generated by go-swagger; DO NOT EDIT.

package expenses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/basiqio/basiq-swagger/dist/models"
)

// PostExpensesReader is a Reader for the PostExpenses structure.
type PostExpensesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExpensesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExpensesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPostExpensesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExpensesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExpensesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExpensesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExpensesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostExpensesOK creates a PostExpensesOK with default headers values
func NewPostExpensesOK() *PostExpensesOK {
	return &PostExpensesOK{}
}

/*PostExpensesOK handles this case with default header values.

Returns a created expenses resource, if the operation succeeded.
*/
type PostExpensesOK struct {
	Payload *models.ExpensesResponse
}

func (o *PostExpensesOK) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/expenses][%d] postExpensesOK  %+v", 200, o.Payload)
}

func (o *PostExpensesOK) GetPayload() *models.ExpensesResponse {
	return o.Payload
}

func (o *PostExpensesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExpensesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExpensesNoContent creates a PostExpensesNoContent with default headers values
func NewPostExpensesNoContent() *PostExpensesNoContent {
	return &PostExpensesNoContent{}
}

/*PostExpensesNoContent handles this case with default header values.

Returns no content if there are none transactions for the requested period.
*/
type PostExpensesNoContent struct {
}

func (o *PostExpensesNoContent) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/expenses][%d] postExpensesNoContent ", 204)
}

func (o *PostExpensesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostExpensesBadRequest creates a PostExpensesBadRequest with default headers values
func NewPostExpensesBadRequest() *PostExpensesBadRequest {
	return &PostExpensesBadRequest{}
}

/*PostExpensesBadRequest handles this case with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error.
*/
type PostExpensesBadRequest struct {
	Payload *models.BadRequestError
}

func (o *PostExpensesBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/expenses][%d] postExpensesBadRequest  %+v", 400, o.Payload)
}

func (o *PostExpensesBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *PostExpensesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExpensesForbidden creates a PostExpensesForbidden with default headers values
func NewPostExpensesForbidden() *PostExpensesForbidden {
	return &PostExpensesForbidden{}
}

/*PostExpensesForbidden handles this case with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type PostExpensesForbidden struct {
	Payload *models.ForbiddenAccessError
}

func (o *PostExpensesForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/expenses][%d] postExpensesForbidden  %+v", 403, o.Payload)
}

func (o *PostExpensesForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *PostExpensesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExpensesNotFound creates a PostExpensesNotFound with default headers values
func NewPostExpensesNotFound() *PostExpensesNotFound {
	return &PostExpensesNotFound{}
}

/*PostExpensesNotFound handles this case with default header values.

Returns error indicating that server can't find requested resource.
*/
type PostExpensesNotFound struct {
	Payload *models.NotFoundError
}

func (o *PostExpensesNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/expenses][%d] postExpensesNotFound  %+v", 404, o.Payload)
}

func (o *PostExpensesNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *PostExpensesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExpensesInternalServerError creates a PostExpensesInternalServerError with default headers values
func NewPostExpensesInternalServerError() *PostExpensesInternalServerError {
	return &PostExpensesInternalServerError{}
}

/*PostExpensesInternalServerError handles this case with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type PostExpensesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostExpensesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/expenses][%d] postExpensesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExpensesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostExpensesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
