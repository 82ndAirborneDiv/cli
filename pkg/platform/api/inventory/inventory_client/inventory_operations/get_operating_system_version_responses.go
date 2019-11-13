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

// GetOperatingSystemVersionReader is a Reader for the GetOperatingSystemVersion structure.
type GetOperatingSystemVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOperatingSystemVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOperatingSystemVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetOperatingSystemVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOperatingSystemVersionOK creates a GetOperatingSystemVersionOK with default headers values
func NewGetOperatingSystemVersionOK() *GetOperatingSystemVersionOK {
	return &GetOperatingSystemVersionOK{}
}

/*GetOperatingSystemVersionOK handles this case with default header values.

The retrieved operating system version
*/
type GetOperatingSystemVersionOK struct {
	Payload *inventory_models.V1OperatingSystemVersion
}

func (o *GetOperatingSystemVersionOK) Error() string {
	return fmt.Sprintf("[GET /v1/operating-systems/{operating_system_id}/versions/{operating_system_version_id}][%d] getOperatingSystemVersionOK  %+v", 200, o.Payload)
}

func (o *GetOperatingSystemVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1OperatingSystemVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOperatingSystemVersionDefault creates a GetOperatingSystemVersionDefault with default headers values
func NewGetOperatingSystemVersionDefault(code int) *GetOperatingSystemVersionDefault {
	return &GetOperatingSystemVersionDefault{
		_statusCode: code,
	}
}

/*GetOperatingSystemVersionDefault handles this case with default header values.

generic error response
*/
type GetOperatingSystemVersionDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get operating system version default response
func (o *GetOperatingSystemVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetOperatingSystemVersionDefault) Error() string {
	return fmt.Sprintf("[GET /v1/operating-systems/{operating_system_id}/versions/{operating_system_version_id}][%d] getOperatingSystemVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetOperatingSystemVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}