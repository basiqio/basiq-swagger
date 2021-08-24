// Code generated by go-swagger; DO NOT EDIT.

package affordability

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/basiqio/basiq-swagger/dist/models"
)

// GetAffordabilityListReader is a Reader for the GetAffordabilityList structure.
type GetAffordabilityListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAffordabilityListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAffordabilityListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAffordabilityListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAffordabilityListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAffordabilityListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAffordabilityListOK creates a GetAffordabilityListOK with default headers values
func NewGetAffordabilityListOK() *GetAffordabilityListOK {
	return &GetAffordabilityListOK{}
}

/* GetAffordabilityListOK describes a response with status code 200, with default header values.

Returns a list with a data property that contains an array of affordability resources.
*/
type GetAffordabilityListOK struct {
	Payload *models.AffordabilityListResponse
}

func (o *GetAffordabilityListOK) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/affordability][%d] getAffordabilityListOK  %+v", 200, o.Payload)
}
func (o *GetAffordabilityListOK) GetPayload() *models.AffordabilityListResponse {
	return o.Payload
}

func (o *GetAffordabilityListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AffordabilityListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAffordabilityListBadRequest creates a GetAffordabilityListBadRequest with default headers values
func NewGetAffordabilityListBadRequest() *GetAffordabilityListBadRequest {
	return &GetAffordabilityListBadRequest{}
}

/* GetAffordabilityListBadRequest describes a response with status code 400, with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error.
*/
type GetAffordabilityListBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetAffordabilityListBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/affordability][%d] getAffordabilityListBadRequest  %+v", 400, o.Payload)
}
func (o *GetAffordabilityListBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetAffordabilityListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAffordabilityListNotFound creates a GetAffordabilityListNotFound with default headers values
func NewGetAffordabilityListNotFound() *GetAffordabilityListNotFound {
	return &GetAffordabilityListNotFound{}
}

/* GetAffordabilityListNotFound describes a response with status code 404, with default header values.

Returns error indicating that server can't find requested resource.
*/
type GetAffordabilityListNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetAffordabilityListNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/affordability][%d] getAffordabilityListNotFound  %+v", 404, o.Payload)
}
func (o *GetAffordabilityListNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetAffordabilityListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAffordabilityListInternalServerError creates a GetAffordabilityListInternalServerError with default headers values
func NewGetAffordabilityListInternalServerError() *GetAffordabilityListInternalServerError {
	return &GetAffordabilityListInternalServerError{}
}

/* GetAffordabilityListInternalServerError describes a response with status code 500, with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type GetAffordabilityListInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetAffordabilityListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/affordability][%d] getAffordabilityListInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAffordabilityListInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetAffordabilityListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}