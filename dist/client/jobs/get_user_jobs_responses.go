// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/basiqio/basiq-swagger/dist/models"
)

// GetUserJobsReader is a Reader for the GetUserJobs structure.
type GetUserJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserJobsOK creates a GetUserJobsOK with default headers values
func NewGetUserJobsOK() *GetUserJobsOK {
	return &GetUserJobsOK{}
}

/*GetUserJobsOK handles this case with default header values.

Returns a list of jobs with details.
*/
type GetUserJobsOK struct {
	Payload *models.JobsResponseResource
}

func (o *GetUserJobsOK) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/jobs][%d] getUserJobsOK  %+v", 200, o.Payload)
}

func (o *GetUserJobsOK) GetPayload() *models.JobsResponseResource {
	return o.Payload
}

func (o *GetUserJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobsResponseResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserJobsBadRequest creates a GetUserJobsBadRequest with default headers values
func NewGetUserJobsBadRequest() *GetUserJobsBadRequest {
	return &GetUserJobsBadRequest{}
}

/*GetUserJobsBadRequest handles this case with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error
*/
type GetUserJobsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetUserJobsBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/jobs][%d] getUserJobsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserJobsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetUserJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserJobsForbidden creates a GetUserJobsForbidden with default headers values
func NewGetUserJobsForbidden() *GetUserJobsForbidden {
	return &GetUserJobsForbidden{}
}

/*GetUserJobsForbidden handles this case with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type GetUserJobsForbidden struct {
	Payload *models.ForbiddenAccessError
}

func (o *GetUserJobsForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/jobs][%d] getUserJobsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserJobsForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *GetUserJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserJobsNotFound creates a GetUserJobsNotFound with default headers values
func NewGetUserJobsNotFound() *GetUserJobsNotFound {
	return &GetUserJobsNotFound{}
}

/*GetUserJobsNotFound handles this case with default header values.

Returns error indicating that server can't find requested resource.
*/
type GetUserJobsNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetUserJobsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/jobs][%d] getUserJobsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserJobsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetUserJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserJobsInternalServerError creates a GetUserJobsInternalServerError with default headers values
func NewGetUserJobsInternalServerError() *GetUserJobsInternalServerError {
	return &GetUserJobsInternalServerError{}
}

/*GetUserJobsInternalServerError handles this case with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type GetUserJobsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUserJobsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/jobs][%d] getUserJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserJobsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUserJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
