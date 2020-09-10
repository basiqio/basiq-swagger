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

// RefreshConnectionsReader is a Reader for the RefreshConnections structure.
type RefreshConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRefreshConnectionsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRefreshConnectionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRefreshConnectionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRefreshConnectionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRefreshConnectionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRefreshConnectionsAccepted creates a RefreshConnectionsAccepted with default headers values
func NewRefreshConnectionsAccepted() *RefreshConnectionsAccepted {
	return &RefreshConnectionsAccepted{}
}

/*RefreshConnectionsAccepted handles this case with default header values.

Returns a created jobs resource, if the operation succeeded.
*/
type RefreshConnectionsAccepted struct {
	Payload *models.ConnectionsRefreshResource
}

func (o *RefreshConnectionsAccepted) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/connections/refresh][%d] refreshConnectionsAccepted  %+v", 202, o.Payload)
}

func (o *RefreshConnectionsAccepted) GetPayload() *models.ConnectionsRefreshResource {
	return o.Payload
}

func (o *RefreshConnectionsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConnectionsRefreshResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshConnectionsBadRequest creates a RefreshConnectionsBadRequest with default headers values
func NewRefreshConnectionsBadRequest() *RefreshConnectionsBadRequest {
	return &RefreshConnectionsBadRequest{}
}

/*RefreshConnectionsBadRequest handles this case with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error
*/
type RefreshConnectionsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *RefreshConnectionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/connections/refresh][%d] refreshConnectionsBadRequest  %+v", 400, o.Payload)
}

func (o *RefreshConnectionsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *RefreshConnectionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshConnectionsForbidden creates a RefreshConnectionsForbidden with default headers values
func NewRefreshConnectionsForbidden() *RefreshConnectionsForbidden {
	return &RefreshConnectionsForbidden{}
}

/*RefreshConnectionsForbidden handles this case with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type RefreshConnectionsForbidden struct {
	Payload *models.ForbiddenAccessError
}

func (o *RefreshConnectionsForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/connections/refresh][%d] refreshConnectionsForbidden  %+v", 403, o.Payload)
}

func (o *RefreshConnectionsForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *RefreshConnectionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshConnectionsNotFound creates a RefreshConnectionsNotFound with default headers values
func NewRefreshConnectionsNotFound() *RefreshConnectionsNotFound {
	return &RefreshConnectionsNotFound{}
}

/*RefreshConnectionsNotFound handles this case with default header values.

Returns error indicating that server can't find requested resource.
*/
type RefreshConnectionsNotFound struct {
	Payload *models.NotFoundError
}

func (o *RefreshConnectionsNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/connections/refresh][%d] refreshConnectionsNotFound  %+v", 404, o.Payload)
}

func (o *RefreshConnectionsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *RefreshConnectionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshConnectionsInternalServerError creates a RefreshConnectionsInternalServerError with default headers values
func NewRefreshConnectionsInternalServerError() *RefreshConnectionsInternalServerError {
	return &RefreshConnectionsInternalServerError{}
}

/*RefreshConnectionsInternalServerError handles this case with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type RefreshConnectionsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *RefreshConnectionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/connections/refresh][%d] refreshConnectionsInternalServerError  %+v", 500, o.Payload)
}

func (o *RefreshConnectionsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *RefreshConnectionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}