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

// GetAuthLinkReader is a Reader for the GetAuthLink structure.
type GetAuthLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthLinkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetAuthLinkGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthLinkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthLinkServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthLinkOK creates a GetAuthLinkOK with default headers values
func NewGetAuthLinkOK() *GetAuthLinkOK {
	return &GetAuthLinkOK{}
}

/*GetAuthLinkOK handles this case with default header values.

Returns details of a connection.
*/
type GetAuthLinkOK struct {
	Payload *models.AuthLinksResponseResource
}

func (o *GetAuthLinkOK) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/auth_link][%d] getAuthLinkOK  %+v", 200, o.Payload)
}

func (o *GetAuthLinkOK) GetPayload() *models.AuthLinksResponseResource {
	return o.Payload
}

func (o *GetAuthLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthLinksResponseResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthLinkBadRequest creates a GetAuthLinkBadRequest with default headers values
func NewGetAuthLinkBadRequest() *GetAuthLinkBadRequest {
	return &GetAuthLinkBadRequest{}
}

/*GetAuthLinkBadRequest handles this case with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error
*/
type GetAuthLinkBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetAuthLinkBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/auth_link][%d] getAuthLinkBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthLinkBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetAuthLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthLinkForbidden creates a GetAuthLinkForbidden with default headers values
func NewGetAuthLinkForbidden() *GetAuthLinkForbidden {
	return &GetAuthLinkForbidden{}
}

/*GetAuthLinkForbidden handles this case with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type GetAuthLinkForbidden struct {
	Payload *models.ForbiddenAccessError
}

func (o *GetAuthLinkForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/auth_link][%d] getAuthLinkForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthLinkForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *GetAuthLinkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthLinkNotFound creates a GetAuthLinkNotFound with default headers values
func NewGetAuthLinkNotFound() *GetAuthLinkNotFound {
	return &GetAuthLinkNotFound{}
}

/*GetAuthLinkNotFound handles this case with default header values.

Returns error indicating that server can't find requested resource.
*/
type GetAuthLinkNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetAuthLinkNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/auth_link][%d] getAuthLinkNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthLinkNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetAuthLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthLinkGone creates a GetAuthLinkGone with default headers values
func NewGetAuthLinkGone() *GetAuthLinkGone {
	return &GetAuthLinkGone{}
}

/*GetAuthLinkGone handles this case with default header values.

Returns error indicating that access to the target resource is no longer available at the origin server and that this condition is likely to be permanent.
*/
type GetAuthLinkGone struct {
	Payload *models.GoneError
}

func (o *GetAuthLinkGone) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/auth_link][%d] getAuthLinkGone  %+v", 410, o.Payload)
}

func (o *GetAuthLinkGone) GetPayload() *models.GoneError {
	return o.Payload
}

func (o *GetAuthLinkGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GoneError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthLinkInternalServerError creates a GetAuthLinkInternalServerError with default headers values
func NewGetAuthLinkInternalServerError() *GetAuthLinkInternalServerError {
	return &GetAuthLinkInternalServerError{}
}

/*GetAuthLinkInternalServerError handles this case with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type GetAuthLinkInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetAuthLinkInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/auth_link][%d] getAuthLinkInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthLinkInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetAuthLinkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthLinkServiceUnavailable creates a GetAuthLinkServiceUnavailable with default headers values
func NewGetAuthLinkServiceUnavailable() *GetAuthLinkServiceUnavailable {
	return &GetAuthLinkServiceUnavailable{}
}

/*GetAuthLinkServiceUnavailable handles this case with default header values.

Returns error response code indicates that the server is not ready to handle the request.
*/
type GetAuthLinkServiceUnavailable struct {
	Payload *models.StatusServiceUnavailableError
}

func (o *GetAuthLinkServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/auth_link][%d] getAuthLinkServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthLinkServiceUnavailable) GetPayload() *models.StatusServiceUnavailableError {
	return o.Payload
}

func (o *GetAuthLinkServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusServiceUnavailableError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
