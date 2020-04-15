// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// AddUserReader is a Reader for the AddUser structure.
type AddUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddUserOK creates a AddUserOK with default headers values
func NewAddUserOK() *AddUserOK {
	return &AddUserOK{}
}

/*AddUserOK handles this case with default header values.

Successfully Created
*/
type AddUserOK struct {
	Payload *mono_models.JWT
}

func (o *AddUserOK) Error() string {
	return fmt.Sprintf("[POST /users][%d] addUserOK  %+v", 200, o.Payload)
}

func (o *AddUserOK) GetPayload() *mono_models.JWT {
	return o.Payload
}

func (o *AddUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.JWT)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserBadRequest creates a AddUserBadRequest with default headers values
func NewAddUserBadRequest() *AddUserBadRequest {
	return &AddUserBadRequest{}
}

/*AddUserBadRequest handles this case with default header values.

Bad Request
*/
type AddUserBadRequest struct {
	Payload *mono_models.Message
}

func (o *AddUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /users][%d] addUserBadRequest  %+v", 400, o.Payload)
}

func (o *AddUserBadRequest) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserConflict creates a AddUserConflict with default headers values
func NewAddUserConflict() *AddUserConflict {
	return &AddUserConflict{}
}

/*AddUserConflict handles this case with default header values.

Conflict
*/
type AddUserConflict struct {
	Payload *mono_models.Message
}

func (o *AddUserConflict) Error() string {
	return fmt.Sprintf("[POST /users][%d] addUserConflict  %+v", 409, o.Payload)
}

func (o *AddUserConflict) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserInternalServerError creates a AddUserInternalServerError with default headers values
func NewAddUserInternalServerError() *AddUserInternalServerError {
	return &AddUserInternalServerError{}
}

/*AddUserInternalServerError handles this case with default header values.

Server Error
*/
type AddUserInternalServerError struct {
	Payload *mono_models.Message
}

func (o *AddUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users][%d] addUserInternalServerError  %+v", 500, o.Payload)
}

func (o *AddUserInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
