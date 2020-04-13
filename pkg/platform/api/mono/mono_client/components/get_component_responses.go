// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// GetComponentReader is a Reader for the GetComponent structure.
type GetComponentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComponentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComponentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetComponentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetComponentOK creates a GetComponentOK with default headers values
func NewGetComponentOK() *GetComponentOK {
	return &GetComponentOK{}
}

/*GetComponentOK handles this case with default header values.

Component Record
*/
type GetComponentOK struct {
	Payload *mono_models.Component
}

func (o *GetComponentOK) Error() string {
	return fmt.Sprintf("[GET /components/{componentID}][%d] getComponentOK  %+v", 200, o.Payload)
}

func (o *GetComponentOK) GetPayload() *mono_models.Component {
	return o.Payload
}

func (o *GetComponentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Component)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentNotFound creates a GetComponentNotFound with default headers values
func NewGetComponentNotFound() *GetComponentNotFound {
	return &GetComponentNotFound{}
}

/*GetComponentNotFound handles this case with default header values.

Not Found
*/
type GetComponentNotFound struct {
	Payload *mono_models.Message
}

func (o *GetComponentNotFound) Error() string {
	return fmt.Sprintf("[GET /components/{componentID}][%d] getComponentNotFound  %+v", 404, o.Payload)
}

func (o *GetComponentNotFound) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetComponentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
