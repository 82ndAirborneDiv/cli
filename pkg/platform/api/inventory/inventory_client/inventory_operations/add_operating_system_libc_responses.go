// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// AddOperatingSystemLibcReader is a Reader for the AddOperatingSystemLibc structure.
type AddOperatingSystemLibcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOperatingSystemLibcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddOperatingSystemLibcOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddOperatingSystemLibcBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddOperatingSystemLibcDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddOperatingSystemLibcOK creates a AddOperatingSystemLibcOK with default headers values
func NewAddOperatingSystemLibcOK() *AddOperatingSystemLibcOK {
	return &AddOperatingSystemLibcOK{}
}

/*AddOperatingSystemLibcOK handles this case with default header values.

The libc added to the operating system
*/
type AddOperatingSystemLibcOK struct {
	Payload *inventory_models.Libc
}

func (o *AddOperatingSystemLibcOK) Error() string {
	return fmt.Sprintf("[POST /v1/operating-systems/{operating_system_id}/libcs][%d] addOperatingSystemLibcOK  %+v", 200, o.Payload)
}

func (o *AddOperatingSystemLibcOK) GetPayload() *inventory_models.Libc {
	return o.Payload
}

func (o *AddOperatingSystemLibcOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.Libc)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOperatingSystemLibcBadRequest creates a AddOperatingSystemLibcBadRequest with default headers values
func NewAddOperatingSystemLibcBadRequest() *AddOperatingSystemLibcBadRequest {
	return &AddOperatingSystemLibcBadRequest{}
}

/*AddOperatingSystemLibcBadRequest handles this case with default header values.

If the libc ID doesn't exist
*/
type AddOperatingSystemLibcBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddOperatingSystemLibcBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/operating-systems/{operating_system_id}/libcs][%d] addOperatingSystemLibcBadRequest  %+v", 400, o.Payload)
}

func (o *AddOperatingSystemLibcBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddOperatingSystemLibcBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOperatingSystemLibcDefault creates a AddOperatingSystemLibcDefault with default headers values
func NewAddOperatingSystemLibcDefault(code int) *AddOperatingSystemLibcDefault {
	return &AddOperatingSystemLibcDefault{
		_statusCode: code,
	}
}

/*AddOperatingSystemLibcDefault handles this case with default header values.

generic error response
*/
type AddOperatingSystemLibcDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add operating system libc default response
func (o *AddOperatingSystemLibcDefault) Code() int {
	return o._statusCode
}

func (o *AddOperatingSystemLibcDefault) Error() string {
	return fmt.Sprintf("[POST /v1/operating-systems/{operating_system_id}/libcs][%d] addOperatingSystemLibc default  %+v", o._statusCode, o.Payload)
}

func (o *AddOperatingSystemLibcDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddOperatingSystemLibcDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
