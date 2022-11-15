// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/basiqio/basiq-swagger/dist/models"
)

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteUserServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserNoContent creates a DeleteUserNoContent with default headers values
func NewDeleteUserNoContent() *DeleteUserNoContent {
	return &DeleteUserNoContent{}
}

/*
DeleteUserNoContent describes a response with status code 204, with default header values.

Deletion succeeded.
*/
type DeleteUserNoContent struct {
}

// IsSuccess returns true when this delete user no content response has a 2xx status code
func (o *DeleteUserNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user no content response has a 3xx status code
func (o *DeleteUserNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user no content response has a 4xx status code
func (o *DeleteUserNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user no content response has a 5xx status code
func (o *DeleteUserNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user no content response a status code equal to that given
func (o *DeleteUserNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserNoContent ", 204)
}

func (o *DeleteUserNoContent) String() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserNoContent ", 204)
}

func (o *DeleteUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserBadRequest creates a DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() *DeleteUserBadRequest {
	return &DeleteUserBadRequest{}
}

/*
DeleteUserBadRequest describes a response with status code 400, with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error.
*/
type DeleteUserBadRequest struct {
	Payload *models.BadRequestError
}

// IsSuccess returns true when this delete user bad request response has a 2xx status code
func (o *DeleteUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user bad request response has a 3xx status code
func (o *DeleteUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user bad request response has a 4xx status code
func (o *DeleteUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user bad request response has a 5xx status code
func (o *DeleteUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user bad request response a status code equal to that given
func (o *DeleteUserBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserBadRequest) String() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *DeleteUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserForbidden creates a DeleteUserForbidden with default headers values
func NewDeleteUserForbidden() *DeleteUserForbidden {
	return &DeleteUserForbidden{}
}

/*
DeleteUserForbidden describes a response with status code 403, with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type DeleteUserForbidden struct {
	Payload *models.ForbiddenAccessError
}

// IsSuccess returns true when this delete user forbidden response has a 2xx status code
func (o *DeleteUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user forbidden response has a 3xx status code
func (o *DeleteUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user forbidden response has a 4xx status code
func (o *DeleteUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user forbidden response has a 5xx status code
func (o *DeleteUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user forbidden response a status code equal to that given
func (o *DeleteUserForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserForbidden) String() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *DeleteUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserNotFound creates a DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {
	return &DeleteUserNotFound{}
}

/*
DeleteUserNotFound describes a response with status code 404, with default header values.

Returns error indicating that server can't find requested resource.
*/
type DeleteUserNotFound struct {
	Payload *models.NotFoundError
}

// IsSuccess returns true when this delete user not found response has a 2xx status code
func (o *DeleteUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user not found response has a 3xx status code
func (o *DeleteUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user not found response has a 4xx status code
func (o *DeleteUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user not found response has a 5xx status code
func (o *DeleteUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user not found response a status code equal to that given
func (o *DeleteUserNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserNotFound) String() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *DeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserInternalServerError creates a DeleteUserInternalServerError with default headers values
func NewDeleteUserInternalServerError() *DeleteUserInternalServerError {
	return &DeleteUserInternalServerError{}
}

/*
DeleteUserInternalServerError describes a response with status code 500, with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type DeleteUserInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this delete user internal server error response has a 2xx status code
func (o *DeleteUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user internal server error response has a 3xx status code
func (o *DeleteUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user internal server error response has a 4xx status code
func (o *DeleteUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user internal server error response has a 5xx status code
func (o *DeleteUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete user internal server error response a status code equal to that given
func (o *DeleteUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *DeleteUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserServiceUnavailable creates a DeleteUserServiceUnavailable with default headers values
func NewDeleteUserServiceUnavailable() *DeleteUserServiceUnavailable {
	return &DeleteUserServiceUnavailable{}
}

/*
DeleteUserServiceUnavailable describes a response with status code 503, with default header values.

Returns error response code indicates that the server is not ready to handle the request.
*/
type DeleteUserServiceUnavailable struct {
	Payload *models.StatusServiceUnavailableError
}

// IsSuccess returns true when this delete user service unavailable response has a 2xx status code
func (o *DeleteUserServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user service unavailable response has a 3xx status code
func (o *DeleteUserServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user service unavailable response has a 4xx status code
func (o *DeleteUserServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user service unavailable response has a 5xx status code
func (o *DeleteUserServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete user service unavailable response a status code equal to that given
func (o *DeleteUserServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteUserServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteUserServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /users/{userId}][%d] deleteUserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteUserServiceUnavailable) GetPayload() *models.StatusServiceUnavailableError {
	return o.Payload
}

func (o *DeleteUserServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusServiceUnavailableError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
