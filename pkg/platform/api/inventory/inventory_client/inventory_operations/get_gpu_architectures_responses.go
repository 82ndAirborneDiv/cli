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

// GetGpuArchitecturesReader is a Reader for the GetGpuArchitectures structure.
type GetGpuArchitecturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGpuArchitecturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetGpuArchitecturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetGpuArchitecturesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetGpuArchitecturesOK creates a GetGpuArchitecturesOK with default headers values
func NewGetGpuArchitecturesOK() *GetGpuArchitecturesOK {
	return &GetGpuArchitecturesOK{}
}

/*GetGpuArchitecturesOK handles this case with default header values.

A paginated list of GPU architectures
*/
type GetGpuArchitecturesOK struct {
	Payload *inventory_models.V1GpuArchitecturePagedList
}

func (o *GetGpuArchitecturesOK) Error() string {
	return fmt.Sprintf("[GET /v1/gpu-architectures][%d] getGpuArchitecturesOK  %+v", 200, o.Payload)
}

func (o *GetGpuArchitecturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1GpuArchitecturePagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGpuArchitecturesDefault creates a GetGpuArchitecturesDefault with default headers values
func NewGetGpuArchitecturesDefault(code int) *GetGpuArchitecturesDefault {
	return &GetGpuArchitecturesDefault{
		_statusCode: code,
	}
}

/*GetGpuArchitecturesDefault handles this case with default header values.

generic error response
*/
type GetGpuArchitecturesDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get gpu architectures default response
func (o *GetGpuArchitecturesDefault) Code() int {
	return o._statusCode
}

func (o *GetGpuArchitecturesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/gpu-architectures][%d] getGpuArchitectures default  %+v", o._statusCode, o.Payload)
}

func (o *GetGpuArchitecturesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
