// Code generated by go-swagger; DO NOT EDIT.

package auth_links

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/basiqio/basiq-swagger/dist/models"
)

// DeleteAuthLinkReader is a Reader for the DeleteAuthLink structure.
type DeleteAuthLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAuthLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAuthLinkNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAuthLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAuthLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAuthLinkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteAuthLinkServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAuthLinkNoContent creates a DeleteAuthLinkNoContent with default headers values
func NewDeleteAuthLinkNoContent() *DeleteAuthLinkNoContent {
	return &DeleteAuthLinkNoContent{}
}

/*DeleteAuthLinkNoContent handles this case with default header values.

Returns an empty body if the delete succeeded.
*/
type DeleteAuthLinkNoContent struct {
}

func (o *DeleteAuthLinkNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}/auth_link][%d] deleteAuthLinkNoContent ", 204)
}

func (o *DeleteAuthLinkNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAuthLinkBadRequest creates a DeleteAuthLinkBadRequest with default headers values
func NewDeleteAuthLinkBadRequest() *DeleteAuthLinkBadRequest {
	return &DeleteAuthLinkBadRequest{}
}

/*DeleteAuthLinkBadRequest handles this case with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error
*/
type DeleteAuthLinkBadRequest struct {
	Payload *models.BadRequestError
}

func (o *DeleteAuthLinkBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}/auth_link][%d] deleteAuthLinkBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAuthLinkBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *DeleteAuthLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthLinkNotFound creates a DeleteAuthLinkNotFound with default headers values
func NewDeleteAuthLinkNotFound() *DeleteAuthLinkNotFound {
	return &DeleteAuthLinkNotFound{}
}

/*DeleteAuthLinkNotFound handles this case with default header values.

Returns error indicating that server can't find requested resource.
*/
type DeleteAuthLinkNotFound struct {
	Payload *models.NotFoundError
}

func (o *DeleteAuthLinkNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}/auth_link][%d] deleteAuthLinkNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAuthLinkNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *DeleteAuthLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthLinkInternalServerError creates a DeleteAuthLinkInternalServerError with default headers values
func NewDeleteAuthLinkInternalServerError() *DeleteAuthLinkInternalServerError {
	return &DeleteAuthLinkInternalServerError{}
}

/*DeleteAuthLinkInternalServerError handles this case with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type DeleteAuthLinkInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *DeleteAuthLinkInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}/auth_link][%d] deleteAuthLinkInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAuthLinkInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *DeleteAuthLinkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthLinkServiceUnavailable creates a DeleteAuthLinkServiceUnavailable with default headers values
func NewDeleteAuthLinkServiceUnavailable() *DeleteAuthLinkServiceUnavailable {
	return &DeleteAuthLinkServiceUnavailable{}
}

/*DeleteAuthLinkServiceUnavailable handles this case with default header values.

Returns error response code indicates that the server is not ready to handle the request.
*/
type DeleteAuthLinkServiceUnavailable struct {
	Payload *models.StatusServiceUnavailableError
}

func (o *DeleteAuthLinkServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}/auth_link][%d] deleteAuthLinkServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteAuthLinkServiceUnavailable) GetPayload() *models.StatusServiceUnavailableError {
	return o.Payload
}

func (o *DeleteAuthLinkServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusServiceUnavailableError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
