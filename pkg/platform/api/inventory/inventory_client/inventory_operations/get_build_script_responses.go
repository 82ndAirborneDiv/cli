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

// GetBuildScriptReader is a Reader for the GetBuildScript structure.
type GetBuildScriptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildScriptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildScriptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBuildScriptDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBuildScriptOK creates a GetBuildScriptOK with default headers values
func NewGetBuildScriptOK() *GetBuildScriptOK {
	return &GetBuildScriptOK{}
}

/*GetBuildScriptOK handles this case with default header values.

The retrieved build script
*/
type GetBuildScriptOK struct {
	Payload *inventory_models.BuildScript
}

func (o *GetBuildScriptOK) Error() string {
	return fmt.Sprintf("[GET /v1/build-scripts/{build_script_id}][%d] getBuildScriptOK  %+v", 200, o.Payload)
}

func (o *GetBuildScriptOK) GetPayload() *inventory_models.BuildScript {
	return o.Payload
}

func (o *GetBuildScriptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.BuildScript)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildScriptDefault creates a GetBuildScriptDefault with default headers values
func NewGetBuildScriptDefault(code int) *GetBuildScriptDefault {
	return &GetBuildScriptDefault{
		_statusCode: code,
	}
}

/*GetBuildScriptDefault handles this case with default header values.

generic error response
*/
type GetBuildScriptDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get build script default response
func (o *GetBuildScriptDefault) Code() int {
	return o._statusCode
}

func (o *GetBuildScriptDefault) Error() string {
	return fmt.Sprintf("[GET /v1/build-scripts/{build_script_id}][%d] getBuildScript default  %+v", o._statusCode, o.Payload)
}

func (o *GetBuildScriptDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetBuildScriptDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
