// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	inventory_models "github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// AddOperatingSystemReader is a Reader for the AddOperatingSystem structure.
type AddOperatingSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOperatingSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddOperatingSystemCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddOperatingSystemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddOperatingSystemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddOperatingSystemCreated creates a AddOperatingSystemCreated with default headers values
func NewAddOperatingSystemCreated() *AddOperatingSystemCreated {
	return &AddOperatingSystemCreated{}
}

/*AddOperatingSystemCreated handles this case with default header values.

The added operating system
*/
type AddOperatingSystemCreated struct {
	Payload *inventory_models.V1OperatingSystem
}

func (o *AddOperatingSystemCreated) Error() string {
	return fmt.Sprintf("[POST /v1/operating-systems][%d] addOperatingSystemCreated  %+v", 201, o.Payload)
}

func (o *AddOperatingSystemCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1OperatingSystem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOperatingSystemBadRequest creates a AddOperatingSystemBadRequest with default headers values
func NewAddOperatingSystemBadRequest() *AddOperatingSystemBadRequest {
	return &AddOperatingSystemBadRequest{}
}

/*AddOperatingSystemBadRequest handles this case with default header values.

If the operating system is invalid
*/
type AddOperatingSystemBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddOperatingSystemBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/operating-systems][%d] addOperatingSystemBadRequest  %+v", 400, o.Payload)
}

func (o *AddOperatingSystemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOperatingSystemDefault creates a AddOperatingSystemDefault with default headers values
func NewAddOperatingSystemDefault(code int) *AddOperatingSystemDefault {
	return &AddOperatingSystemDefault{
		_statusCode: code,
	}
}

/*AddOperatingSystemDefault handles this case with default header values.

If there is an error processing the request
*/
type AddOperatingSystemDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add operating system default response
func (o *AddOperatingSystemDefault) Code() int {
	return o._statusCode
}

func (o *AddOperatingSystemDefault) Error() string {
	return fmt.Sprintf("[POST /v1/operating-systems][%d] addOperatingSystem default  %+v", o._statusCode, o.Payload)
}

func (o *AddOperatingSystemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
