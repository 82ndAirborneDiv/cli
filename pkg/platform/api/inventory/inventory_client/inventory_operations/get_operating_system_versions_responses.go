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

// GetOperatingSystemVersionsReader is a Reader for the GetOperatingSystemVersions structure.
type GetOperatingSystemVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOperatingSystemVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOperatingSystemVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetOperatingSystemVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOperatingSystemVersionsOK creates a GetOperatingSystemVersionsOK with default headers values
func NewGetOperatingSystemVersionsOK() *GetOperatingSystemVersionsOK {
	return &GetOperatingSystemVersionsOK{}
}

/*GetOperatingSystemVersionsOK handles this case with default header values.

A paginated list of operating system versions
*/
type GetOperatingSystemVersionsOK struct {
	Payload *inventory_models.OperatingSystemVersionPagedList
}

func (o *GetOperatingSystemVersionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/operating-systems/{operating_system_id}/versions][%d] getOperatingSystemVersionsOK  %+v", 200, o.Payload)
}

func (o *GetOperatingSystemVersionsOK) GetPayload() *inventory_models.OperatingSystemVersionPagedList {
	return o.Payload
}

func (o *GetOperatingSystemVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.OperatingSystemVersionPagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOperatingSystemVersionsDefault creates a GetOperatingSystemVersionsDefault with default headers values
func NewGetOperatingSystemVersionsDefault(code int) *GetOperatingSystemVersionsDefault {
	return &GetOperatingSystemVersionsDefault{
		_statusCode: code,
	}
}

/*GetOperatingSystemVersionsDefault handles this case with default header values.

generic error response
*/
type GetOperatingSystemVersionsDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get operating system versions default response
func (o *GetOperatingSystemVersionsDefault) Code() int {
	return o._statusCode
}

func (o *GetOperatingSystemVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/operating-systems/{operating_system_id}/versions][%d] getOperatingSystemVersions default  %+v", o._statusCode, o.Payload)
}

func (o *GetOperatingSystemVersionsDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetOperatingSystemVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
