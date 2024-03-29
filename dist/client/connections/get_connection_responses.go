// Code generated by go-swagger; DO NOT EDIT.

package connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/basiqio/basiq-swagger/dist/models"
)

// GetConnectionReader is a Reader for the GetConnection structure.
type GetConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConnectionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConnectionOK creates a GetConnectionOK with default headers values
func NewGetConnectionOK() *GetConnectionOK {
	return &GetConnectionOK{}
}

/* GetConnectionOK describes a response with status code 200, with default header values.

Returns details of a connection.
*/
type GetConnectionOK struct {
	Payload *models.ConnectionGetResponseResource
}

func (o *GetConnectionOK) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/connections/{connectionId}][%d] getConnectionOK  %+v", 200, o.Payload)
}
func (o *GetConnectionOK) GetPayload() *models.ConnectionGetResponseResource {
	return o.Payload
}

func (o *GetConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConnectionGetResponseResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionBadRequest creates a GetConnectionBadRequest with default headers values
func NewGetConnectionBadRequest() *GetConnectionBadRequest {
	return &GetConnectionBadRequest{}
}

/* GetConnectionBadRequest describes a response with status code 400, with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error
*/
type GetConnectionBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetConnectionBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/connections/{connectionId}][%d] getConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetConnectionBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionUnauthorized creates a GetConnectionUnauthorized with default headers values
func NewGetConnectionUnauthorized() *GetConnectionUnauthorized {
	return &GetConnectionUnauthorized{}
}

/* GetConnectionUnauthorized describes a response with status code 401, with default header values.

Error status response code indicates that the request has not been applied because it lacks valid authentication credentials for the target resource.
*/
type GetConnectionUnauthorized struct {
	Payload *models.UnauthorizedError
}

func (o *GetConnectionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/connections/{connectionId}][%d] getConnectionUnauthorized  %+v", 401, o.Payload)
}
func (o *GetConnectionUnauthorized) GetPayload() *models.UnauthorizedError {
	return o.Payload
}

func (o *GetConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionForbidden creates a GetConnectionForbidden with default headers values
func NewGetConnectionForbidden() *GetConnectionForbidden {
	return &GetConnectionForbidden{}
}

/* GetConnectionForbidden describes a response with status code 403, with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type GetConnectionForbidden struct {
	Payload *models.ForbiddenAccessError
}

func (o *GetConnectionForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/connections/{connectionId}][%d] getConnectionForbidden  %+v", 403, o.Payload)
}
func (o *GetConnectionForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *GetConnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionNotFound creates a GetConnectionNotFound with default headers values
func NewGetConnectionNotFound() *GetConnectionNotFound {
	return &GetConnectionNotFound{}
}

/* GetConnectionNotFound describes a response with status code 404, with default header values.

Returns error indicating that server can't find requested resource.
*/
type GetConnectionNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetConnectionNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/connections/{connectionId}][%d] getConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetConnectionNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionInternalServerError creates a GetConnectionInternalServerError with default headers values
func NewGetConnectionInternalServerError() *GetConnectionInternalServerError {
	return &GetConnectionInternalServerError{}
}

/* GetConnectionInternalServerError describes a response with status code 500, with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type GetConnectionInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/connections/{connectionId}][%d] getConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetConnectionInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConnectionServiceUnavailable creates a GetConnectionServiceUnavailable with default headers values
func NewGetConnectionServiceUnavailable() *GetConnectionServiceUnavailable {
	return &GetConnectionServiceUnavailable{}
}

/* GetConnectionServiceUnavailable describes a response with status code 503, with default header values.

Returns error response code indicates that the server is not ready to handle the request.
*/
type GetConnectionServiceUnavailable struct {
	Payload *models.StatusServiceUnavailableError
}

func (o *GetConnectionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/connections/{connectionId}][%d] getConnectionServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetConnectionServiceUnavailable) GetPayload() *models.StatusServiceUnavailableError {
	return o.Payload
}

func (o *GetConnectionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusServiceUnavailableError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
