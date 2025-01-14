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

// GetReleaseReader is a Reader for the GetRelease structure.
type GetReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetReleaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReleaseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReleaseOK creates a GetReleaseOK with default headers values
func NewGetReleaseOK() *GetReleaseOK {
	return &GetReleaseOK{}
}

/*GetReleaseOK handles this case with default header values.

Success
*/
type GetReleaseOK struct {
	Payload *mono_models.Release
}

func (o *GetReleaseOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}][%d] getReleaseOK  %+v", 200, o.Payload)
}

func (o *GetReleaseOK) GetPayload() *mono_models.Release {
	return o.Payload
}

func (o *GetReleaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Release)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseNotFound creates a GetReleaseNotFound with default headers values
func NewGetReleaseNotFound() *GetReleaseNotFound {
	return &GetReleaseNotFound{}
}

/*GetReleaseNotFound handles this case with default header values.

Not Found
*/
type GetReleaseNotFound struct {
	Payload *mono_models.Message
}

func (o *GetReleaseNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}][%d] getReleaseNotFound  %+v", 404, o.Payload)
}

func (o *GetReleaseNotFound) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetReleaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseInternalServerError creates a GetReleaseInternalServerError with default headers values
func NewGetReleaseInternalServerError() *GetReleaseInternalServerError {
	return &GetReleaseInternalServerError{}
}

/*GetReleaseInternalServerError handles this case with default header values.

Server Error
*/
type GetReleaseInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetReleaseInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}][%d] getReleaseInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReleaseInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetReleaseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
