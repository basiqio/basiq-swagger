// Code generated by go-swagger; DO NOT EDIT.

package identities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/basiqio/basiq-swagger/dist/models"
)

// GetUserIdentitiesReader is a Reader for the GetUserIdentities structure.
type GetUserIdentitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserIdentitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserIdentitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserIdentitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserIdentitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserIdentitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserIdentitiesOK creates a GetUserIdentitiesOK with default headers values
func NewGetUserIdentitiesOK() *GetUserIdentitiesOK {
	return &GetUserIdentitiesOK{}
}

/*
GetUserIdentitiesOK describes a response with status code 200, with default header values.

Returns list of identities.
*/
type GetUserIdentitiesOK struct {
	Payload *models.IdentitiesGetResponseResource
}

// IsSuccess returns true when this get user identities o k response has a 2xx status code
func (o *GetUserIdentitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user identities o k response has a 3xx status code
func (o *GetUserIdentitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user identities o k response has a 4xx status code
func (o *GetUserIdentitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user identities o k response has a 5xx status code
func (o *GetUserIdentitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user identities o k response a status code equal to that given
func (o *GetUserIdentitiesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserIdentitiesOK) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/identitites][%d] getUserIdentitiesOK  %+v", 200, o.Payload)
}

func (o *GetUserIdentitiesOK) String() string {
	return fmt.Sprintf("[GET /users/{userId}/identitites][%d] getUserIdentitiesOK  %+v", 200, o.Payload)
}

func (o *GetUserIdentitiesOK) GetPayload() *models.IdentitiesGetResponseResource {
	return o.Payload
}

func (o *GetUserIdentitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdentitiesGetResponseResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserIdentitiesBadRequest creates a GetUserIdentitiesBadRequest with default headers values
func NewGetUserIdentitiesBadRequest() *GetUserIdentitiesBadRequest {
	return &GetUserIdentitiesBadRequest{}
}

/*
GetUserIdentitiesBadRequest describes a response with status code 400, with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error.
*/
type GetUserIdentitiesBadRequest struct {
	Payload *models.BadRequestError
}

// IsSuccess returns true when this get user identities bad request response has a 2xx status code
func (o *GetUserIdentitiesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user identities bad request response has a 3xx status code
func (o *GetUserIdentitiesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user identities bad request response has a 4xx status code
func (o *GetUserIdentitiesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user identities bad request response has a 5xx status code
func (o *GetUserIdentitiesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user identities bad request response a status code equal to that given
func (o *GetUserIdentitiesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUserIdentitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/identitites][%d] getUserIdentitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserIdentitiesBadRequest) String() string {
	return fmt.Sprintf("[GET /users/{userId}/identitites][%d] getUserIdentitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserIdentitiesBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetUserIdentitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserIdentitiesForbidden creates a GetUserIdentitiesForbidden with default headers values
func NewGetUserIdentitiesForbidden() *GetUserIdentitiesForbidden {
	return &GetUserIdentitiesForbidden{}
}

/*
GetUserIdentitiesForbidden describes a response with status code 403, with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type GetUserIdentitiesForbidden struct {
	Payload *models.ForbiddenAccessError
}

// IsSuccess returns true when this get user identities forbidden response has a 2xx status code
func (o *GetUserIdentitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user identities forbidden response has a 3xx status code
func (o *GetUserIdentitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user identities forbidden response has a 4xx status code
func (o *GetUserIdentitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user identities forbidden response has a 5xx status code
func (o *GetUserIdentitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user identities forbidden response a status code equal to that given
func (o *GetUserIdentitiesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUserIdentitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/identitites][%d] getUserIdentitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetUserIdentitiesForbidden) String() string {
	return fmt.Sprintf("[GET /users/{userId}/identitites][%d] getUserIdentitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetUserIdentitiesForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *GetUserIdentitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserIdentitiesInternalServerError creates a GetUserIdentitiesInternalServerError with default headers values
func NewGetUserIdentitiesInternalServerError() *GetUserIdentitiesInternalServerError {
	return &GetUserIdentitiesInternalServerError{}
}

/*
GetUserIdentitiesInternalServerError describes a response with status code 500, with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type GetUserIdentitiesInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get user identities internal server error response has a 2xx status code
func (o *GetUserIdentitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user identities internal server error response has a 3xx status code
func (o *GetUserIdentitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user identities internal server error response has a 4xx status code
func (o *GetUserIdentitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user identities internal server error response has a 5xx status code
func (o *GetUserIdentitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user identities internal server error response a status code equal to that given
func (o *GetUserIdentitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserIdentitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/identitites][%d] getUserIdentitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserIdentitiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /users/{userId}/identitites][%d] getUserIdentitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserIdentitiesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUserIdentitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}