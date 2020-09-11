// Code generated by go-swagger; DO NOT EDIT.

package statements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/basiqio/basiq-swagger/dist/models"
)

// CreateStatementReader is a Reader for the CreateStatement structure.
type CreateStatementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStatementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateStatementAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateStatementBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateStatementForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateStatementNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateStatementInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateStatementServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateStatementAccepted creates a CreateStatementAccepted with default headers values
func NewCreateStatementAccepted() *CreateStatementAccepted {
	return &CreateStatementAccepted{}
}

/*CreateStatementAccepted handles this case with default header values.

Returns a job details if a valid job ID was provided.
*/
type CreateStatementAccepted struct {
	Payload *models.StatementUploadResource
}

func (o *CreateStatementAccepted) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/statements][%d] createStatementAccepted  %+v", 202, o.Payload)
}

func (o *CreateStatementAccepted) GetPayload() *models.StatementUploadResource {
	return o.Payload
}

func (o *CreateStatementAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatementUploadResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStatementBadRequest creates a CreateStatementBadRequest with default headers values
func NewCreateStatementBadRequest() *CreateStatementBadRequest {
	return &CreateStatementBadRequest{}
}

/*CreateStatementBadRequest handles this case with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error.
*/
type CreateStatementBadRequest struct {
	Payload *models.BadRequestError
}

func (o *CreateStatementBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/statements][%d] createStatementBadRequest  %+v", 400, o.Payload)
}

func (o *CreateStatementBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *CreateStatementBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStatementForbidden creates a CreateStatementForbidden with default headers values
func NewCreateStatementForbidden() *CreateStatementForbidden {
	return &CreateStatementForbidden{}
}

/*CreateStatementForbidden handles this case with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type CreateStatementForbidden struct {
	Payload *models.ForbiddenAccessError
}

func (o *CreateStatementForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/statements][%d] createStatementForbidden  %+v", 403, o.Payload)
}

func (o *CreateStatementForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *CreateStatementForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStatementNotFound creates a CreateStatementNotFound with default headers values
func NewCreateStatementNotFound() *CreateStatementNotFound {
	return &CreateStatementNotFound{}
}

/*CreateStatementNotFound handles this case with default header values.

Returns error indicating that server can't find requested resource.
*/
type CreateStatementNotFound struct {
	Payload *models.NotFoundError
}

func (o *CreateStatementNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/statements][%d] createStatementNotFound  %+v", 404, o.Payload)
}

func (o *CreateStatementNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *CreateStatementNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStatementInternalServerError creates a CreateStatementInternalServerError with default headers values
func NewCreateStatementInternalServerError() *CreateStatementInternalServerError {
	return &CreateStatementInternalServerError{}
}

/*CreateStatementInternalServerError handles this case with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type CreateStatementInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *CreateStatementInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/statements][%d] createStatementInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateStatementInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *CreateStatementInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStatementServiceUnavailable creates a CreateStatementServiceUnavailable with default headers values
func NewCreateStatementServiceUnavailable() *CreateStatementServiceUnavailable {
	return &CreateStatementServiceUnavailable{}
}

/*CreateStatementServiceUnavailable handles this case with default header values.

Returns error response code indicates that the server is not ready to handle the request.
*/
type CreateStatementServiceUnavailable struct {
	Payload *models.StatusServiceUnavailableError
}

func (o *CreateStatementServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/statements][%d] createStatementServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateStatementServiceUnavailable) GetPayload() *models.StatusServiceUnavailableError {
	return o.Payload
}

func (o *CreateStatementServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusServiceUnavailableError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
