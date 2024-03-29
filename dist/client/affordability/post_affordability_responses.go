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

// PostAffordabilityReader is a Reader for the PostAffordability structure.
type PostAffordabilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAffordabilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAffordabilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostAffordabilityCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPostAffordabilityNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAffordabilityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAffordabilityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAffordabilityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAffordabilityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAffordabilityOK creates a PostAffordabilityOK with default headers values
func NewPostAffordabilityOK() *PostAffordabilityOK {
	return &PostAffordabilityOK{}
}

/* PostAffordabilityOK describes a response with status code 200, with default header values.

Returns a created affordability resource, if the operation succeeded.
*/
type PostAffordabilityOK struct {
	Payload *models.AffordabilityResponse
}

func (o *PostAffordabilityOK) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/affordability][%d] postAffordabilityOK  %+v", 200, o.Payload)
}
func (o *PostAffordabilityOK) GetPayload() *models.AffordabilityResponse {
	return o.Payload
}

func (o *PostAffordabilityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AffordabilityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAffordabilityCreated creates a PostAffordabilityCreated with default headers values
func NewPostAffordabilityCreated() *PostAffordabilityCreated {
	return &PostAffordabilityCreated{}
}

/* PostAffordabilityCreated describes a response with status code 201, with default header values.

Returns a created affordability PDF Report, if the operation succeeded.
*/
type PostAffordabilityCreated struct {
}

func (o *PostAffordabilityCreated) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/affordability][%d] postAffordabilityCreated ", 201)
}

func (o *PostAffordabilityCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAffordabilityNoContent creates a PostAffordabilityNoContent with default headers values
func NewPostAffordabilityNoContent() *PostAffordabilityNoContent {
	return &PostAffordabilityNoContent{}
}

/* PostAffordabilityNoContent describes a response with status code 204, with default header values.

Returns no content if there are none transactions for the requested period.
*/
type PostAffordabilityNoContent struct {
}

func (o *PostAffordabilityNoContent) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/affordability][%d] postAffordabilityNoContent ", 204)
}

func (o *PostAffordabilityNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAffordabilityBadRequest creates a PostAffordabilityBadRequest with default headers values
func NewPostAffordabilityBadRequest() *PostAffordabilityBadRequest {
	return &PostAffordabilityBadRequest{}
}

/* PostAffordabilityBadRequest describes a response with status code 400, with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error.
*/
type PostAffordabilityBadRequest struct {
	Payload *models.BadRequestError
}

func (o *PostAffordabilityBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/affordability][%d] postAffordabilityBadRequest  %+v", 400, o.Payload)
}
func (o *PostAffordabilityBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *PostAffordabilityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAffordabilityForbidden creates a PostAffordabilityForbidden with default headers values
func NewPostAffordabilityForbidden() *PostAffordabilityForbidden {
	return &PostAffordabilityForbidden{}
}

/* PostAffordabilityForbidden describes a response with status code 403, with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type PostAffordabilityForbidden struct {
	Payload *models.ForbiddenAccessError
}

func (o *PostAffordabilityForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/affordability][%d] postAffordabilityForbidden  %+v", 403, o.Payload)
}
func (o *PostAffordabilityForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *PostAffordabilityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAffordabilityNotFound creates a PostAffordabilityNotFound with default headers values
func NewPostAffordabilityNotFound() *PostAffordabilityNotFound {
	return &PostAffordabilityNotFound{}
}

/* PostAffordabilityNotFound describes a response with status code 404, with default header values.

Returns error indicating that server can't find requested resource.
*/
type PostAffordabilityNotFound struct {
	Payload *models.NotFoundError
}

func (o *PostAffordabilityNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/affordability][%d] postAffordabilityNotFound  %+v", 404, o.Payload)
}
func (o *PostAffordabilityNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *PostAffordabilityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAffordabilityInternalServerError creates a PostAffordabilityInternalServerError with default headers values
func NewPostAffordabilityInternalServerError() *PostAffordabilityInternalServerError {
	return &PostAffordabilityInternalServerError{}
}

/* PostAffordabilityInternalServerError describes a response with status code 500, with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type PostAffordabilityInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostAffordabilityInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/{userId}/affordability][%d] postAffordabilityInternalServerError  %+v", 500, o.Payload)
}
func (o *PostAffordabilityInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostAffordabilityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
