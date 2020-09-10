// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/basiqio/basiq-swagger/dist/models"
)

// GetAccountReader is a Reader for the GetAccount structure.
type GetAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountOK creates a GetAccountOK with default headers values
func NewGetAccountOK() *GetAccountOK {
	return &GetAccountOK{}
}

/*GetAccountOK handles this case with default header values.

Returns a account with details.
*/
type GetAccountOK struct {
	Payload *models.AccountResponseResource
}

func (o *GetAccountOK) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/accounts/{accountId}][%d] getAccountOK  %+v", 200, o.Payload)
}

func (o *GetAccountOK) GetPayload() *models.AccountResponseResource {
	return o.Payload
}

func (o *GetAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountResponseResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountBadRequest creates a GetAccountBadRequest with default headers values
func NewGetAccountBadRequest() *GetAccountBadRequest {
	return &GetAccountBadRequest{}
}

/*GetAccountBadRequest handles this case with default header values.

Returns error that server cannot or will not process the request due to something that is perceived to be a client error
*/
type GetAccountBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/accounts/{accountId}][%d] getAccountBadRequest  %+v", 400, o.Payload)
}

func (o *GetAccountBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountForbidden creates a GetAccountForbidden with default headers values
func NewGetAccountForbidden() *GetAccountForbidden {
	return &GetAccountForbidden{}
}

/*GetAccountForbidden handles this case with default header values.

Error that access is forbidden and tied to the application logic, such as insufficient rights to a resource.
*/
type GetAccountForbidden struct {
	Payload *models.ForbiddenAccessError
}

func (o *GetAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/accounts/{accountId}][%d] getAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetAccountForbidden) GetPayload() *models.ForbiddenAccessError {
	return o.Payload
}

func (o *GetAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForbiddenAccessError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountNotFound creates a GetAccountNotFound with default headers values
func NewGetAccountNotFound() *GetAccountNotFound {
	return &GetAccountNotFound{}
}

/*GetAccountNotFound handles this case with default header values.

Returns error indicating that server can't find requested resource.
*/
type GetAccountNotFound struct {
	Payload *models.NotFoundError
}

func (o *GetAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/accounts/{accountId}][%d] getAccountNotFound  %+v", 404, o.Payload)
}

func (o *GetAccountNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *GetAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountInternalServerError creates a GetAccountInternalServerError with default headers values
func NewGetAccountInternalServerError() *GetAccountInternalServerError {
	return &GetAccountInternalServerError{}
}

/*GetAccountInternalServerError handles this case with default header values.

Returns error response code indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
*/
type GetAccountInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetAccountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{userId}/accounts/{accountId}][%d] getAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccountInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}