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

// GetJobsReader is a Reader for the GetJobs structure.
type GetJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetJobsOK creates a GetJobsOK with default headers values
func NewGetJobsOK() *GetJobsOK {
	return &GetJobsOK{}
}

/*GetJobsOK handles this case with default header values.

Returns a job details if a valid job ID was provided.
*/
type GetJobsOK struct {
	Payload *models.JobData
}

func (o *GetJobsOK) Error() string {
	return fmt.Sprintf("[GET /jobs/{jobId}][%d] getJobsOK  %+v", 200, o.Payload)
}

func (o *GetJobsOK) GetPayload() *models.JobData {
	return o.Payload
}

func (o *GetJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobsBadRequest creates a GetJobsBadRequest with default headers values
func NewGetJobsBadRequest() *GetJobsBadRequest {
	return &GetJobsBadRequest{}
}

/*GetJobsBadRequest handles this case with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error
*/
type GetJobsBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetJobsBadRequest) Error() string {
	return fmt.Sprintf("[GET /jobs/{jobId}][%d] getJobsBadRequest  %+v", 400, o.Payload)
}

func (o *GetJobsBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobsForbidden creates a GetJobsForbidden with default headers values
func NewGetJobsForbidden() *GetJobsForbidden {
	return &GetJobsForbidden{}
}

/*GetJobsForbidden handles this case with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type GetJobsForbidden struct {
	Payload *models.ForbiddenAccessError
}

func (o *GetJobsForbidden) Error() string {
	return fmt.Sprintf("[GET /jobs/{jobId}][%d] getJobsForbidden  %+v", 403, o.Payload)
}

func (o *GetJobsForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *GetJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobsNotFound creates a GetJobsNotFound with default headers values
func NewGetJobsNotFound() *GetJobsNotFound {
	return &GetJobsNotFound{}
}

/*GetJobsNotFound handles this case with default header values.

Returns error indicating that server can't find requested resource.
*/
type GetJobsNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetJobsNotFound) Error() string {
	return fmt.Sprintf("[GET /jobs/{jobId}][%d] getJobsNotFound  %+v", 404, o.Payload)
}

func (o *GetJobsNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobsInternalServerError creates a GetJobsInternalServerError with default headers values
func NewGetJobsInternalServerError() *GetJobsInternalServerError {
	return &GetJobsInternalServerError{}
}

/*GetJobsInternalServerError handles this case with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type GetJobsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetJobsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /jobs/{jobId}][%d] getJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJobsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
