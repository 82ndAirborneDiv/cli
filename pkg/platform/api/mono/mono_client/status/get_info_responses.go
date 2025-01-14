// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// GetInfoReader is a Reader for the GetInfo structure.
type GetInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInfoOK creates a GetInfoOK with default headers values
func NewGetInfoOK() *GetInfoOK {
	return &GetInfoOK{}
}

/*GetInfoOK handles this case with default header values.

Success
*/
type GetInfoOK struct {
	Payload *mono_models.SysInfo
}

func (o *GetInfoOK) Error() string {
	return fmt.Sprintf("[GET /info][%d] getInfoOK  %+v", 200, o.Payload)
}

func (o *GetInfoOK) GetPayload() *mono_models.SysInfo {
	return o.Payload
}

func (o *GetInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.SysInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfoUnauthorized creates a GetInfoUnauthorized with default headers values
func NewGetInfoUnauthorized() *GetInfoUnauthorized {
	return &GetInfoUnauthorized{}
}

/*GetInfoUnauthorized handles this case with default header values.

Unauthenticated
*/
type GetInfoUnauthorized struct {
	Payload *mono_models.Message
}

func (o *GetInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /info][%d] getInfoUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInfoUnauthorized) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfoForbidden creates a GetInfoForbidden with default headers values
func NewGetInfoForbidden() *GetInfoForbidden {
	return &GetInfoForbidden{}
}

/*GetInfoForbidden handles this case with default header values.

Unauthorized
*/
type GetInfoForbidden struct {
	Payload *mono_models.Message
}

func (o *GetInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /info][%d] getInfoForbidden  %+v", 403, o.Payload)
}

func (o *GetInfoForbidden) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
