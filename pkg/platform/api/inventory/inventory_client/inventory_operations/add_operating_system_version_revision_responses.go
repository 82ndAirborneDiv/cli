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

// AddOperatingSystemVersionRevisionReader is a Reader for the AddOperatingSystemVersionRevision structure.
type AddOperatingSystemVersionRevisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOperatingSystemVersionRevisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddOperatingSystemVersionRevisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddOperatingSystemVersionRevisionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddOperatingSystemVersionRevisionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddOperatingSystemVersionRevisionOK creates a AddOperatingSystemVersionRevisionOK with default headers values
func NewAddOperatingSystemVersionRevisionOK() *AddOperatingSystemVersionRevisionOK {
	return &AddOperatingSystemVersionRevisionOK{}
}

/*AddOperatingSystemVersionRevisionOK handles this case with default header values.

The updated state of the operating system version
*/
type AddOperatingSystemVersionRevisionOK struct {
	Payload *inventory_models.OperatingSystemVersion
}

func (o *AddOperatingSystemVersionRevisionOK) Error() string {
	return fmt.Sprintf("[POST /v1/operating-systems/{operating_system_id}/versions/{operating_system_version_id}/revisions][%d] addOperatingSystemVersionRevisionOK  %+v", 200, o.Payload)
}

func (o *AddOperatingSystemVersionRevisionOK) GetPayload() *inventory_models.OperatingSystemVersion {
	return o.Payload
}

func (o *AddOperatingSystemVersionRevisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.OperatingSystemVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOperatingSystemVersionRevisionBadRequest creates a AddOperatingSystemVersionRevisionBadRequest with default headers values
func NewAddOperatingSystemVersionRevisionBadRequest() *AddOperatingSystemVersionRevisionBadRequest {
	return &AddOperatingSystemVersionRevisionBadRequest{}
}

/*AddOperatingSystemVersionRevisionBadRequest handles this case with default header values.

If the operating system version revision is invalid
*/
type AddOperatingSystemVersionRevisionBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddOperatingSystemVersionRevisionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/operating-systems/{operating_system_id}/versions/{operating_system_version_id}/revisions][%d] addOperatingSystemVersionRevisionBadRequest  %+v", 400, o.Payload)
}

func (o *AddOperatingSystemVersionRevisionBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddOperatingSystemVersionRevisionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOperatingSystemVersionRevisionDefault creates a AddOperatingSystemVersionRevisionDefault with default headers values
func NewAddOperatingSystemVersionRevisionDefault(code int) *AddOperatingSystemVersionRevisionDefault {
	return &AddOperatingSystemVersionRevisionDefault{
		_statusCode: code,
	}
}

/*AddOperatingSystemVersionRevisionDefault handles this case with default header values.

If there is an error processing the request
*/
type AddOperatingSystemVersionRevisionDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add operating system version revision default response
func (o *AddOperatingSystemVersionRevisionDefault) Code() int {
	return o._statusCode
}

func (o *AddOperatingSystemVersionRevisionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/operating-systems/{operating_system_id}/versions/{operating_system_version_id}/revisions][%d] addOperatingSystemVersionRevision default  %+v", o._statusCode, o.Payload)
}

func (o *AddOperatingSystemVersionRevisionDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddOperatingSystemVersionRevisionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
