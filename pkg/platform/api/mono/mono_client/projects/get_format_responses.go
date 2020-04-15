// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// GetFormatReader is a Reader for the GetFormat structure.
type GetFormatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFormatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFormatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetFormatNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFormatInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFormatOK creates a GetFormatOK with default headers values
func NewGetFormatOK() *GetFormatOK {
	return &GetFormatOK{}
}

/*GetFormatOK handles this case with default header values.

Success
*/
type GetFormatOK struct {
	Payload *mono_models.Format
}

func (o *GetFormatOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros/{distroID}/formats/{formatID}][%d] getFormatOK  %+v", 200, o.Payload)
}

func (o *GetFormatOK) GetPayload() *mono_models.Format {
	return o.Payload
}

func (o *GetFormatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Format)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFormatNotFound creates a GetFormatNotFound with default headers values
func NewGetFormatNotFound() *GetFormatNotFound {
	return &GetFormatNotFound{}
}

/*GetFormatNotFound handles this case with default header values.

Not Found
*/
type GetFormatNotFound struct {
	Payload *mono_models.Message
}

func (o *GetFormatNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros/{distroID}/formats/{formatID}][%d] getFormatNotFound  %+v", 404, o.Payload)
}

func (o *GetFormatNotFound) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetFormatNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFormatInternalServerError creates a GetFormatInternalServerError with default headers values
func NewGetFormatInternalServerError() *GetFormatInternalServerError {
	return &GetFormatInternalServerError{}
}

/*GetFormatInternalServerError handles this case with default header values.

Server Error
*/
type GetFormatInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetFormatInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros/{distroID}/formats/{formatID}][%d] getFormatInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFormatInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetFormatInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
