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

// GetPlatformsReader is a Reader for the GetPlatforms structure.
type GetPlatformsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlatformsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPlatformsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPlatformsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPlatformsOK creates a GetPlatformsOK with default headers values
func NewGetPlatformsOK() *GetPlatformsOK {
	return &GetPlatformsOK{}
}

/*GetPlatformsOK handles this case with default header values.

A paginated list of platforms
*/
type GetPlatformsOK struct {
	Payload *inventory_models.PlatformPagedList
}

func (o *GetPlatformsOK) Error() string {
	return fmt.Sprintf("[GET /v1/platforms][%d] getPlatformsOK  %+v", 200, o.Payload)
}

func (o *GetPlatformsOK) GetPayload() *inventory_models.PlatformPagedList {
	return o.Payload
}

func (o *GetPlatformsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.PlatformPagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformsDefault creates a GetPlatformsDefault with default headers values
func NewGetPlatformsDefault(code int) *GetPlatformsDefault {
	return &GetPlatformsDefault{
		_statusCode: code,
	}
}

/*GetPlatformsDefault handles this case with default header values.

generic error response
*/
type GetPlatformsDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get platforms default response
func (o *GetPlatformsDefault) Code() int {
	return o._statusCode
}

func (o *GetPlatformsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/platforms][%d] getPlatforms default  %+v", o._statusCode, o.Payload)
}

func (o *GetPlatformsDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetPlatformsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
