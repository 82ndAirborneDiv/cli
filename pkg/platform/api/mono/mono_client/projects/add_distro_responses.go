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

// AddDistroReader is a Reader for the AddDistro structure.
type AddDistroReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDistroReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddDistroOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddDistroBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddDistroForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddDistroConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddDistroInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddDistroOK creates a AddDistroOK with default headers values
func NewAddDistroOK() *AddDistroOK {
	return &AddDistroOK{}
}

/*AddDistroOK handles this case with default header values.

Distro Added
*/
type AddDistroOK struct {
	Payload *mono_models.Distro
}

func (o *AddDistroOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros][%d] addDistroOK  %+v", 200, o.Payload)
}

func (o *AddDistroOK) GetPayload() *mono_models.Distro {
	return o.Payload
}

func (o *AddDistroOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Distro)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDistroBadRequest creates a AddDistroBadRequest with default headers values
func NewAddDistroBadRequest() *AddDistroBadRequest {
	return &AddDistroBadRequest{}
}

/*AddDistroBadRequest handles this case with default header values.

Bad Request
*/
type AddDistroBadRequest struct {
	Payload *mono_models.Message
}

func (o *AddDistroBadRequest) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros][%d] addDistroBadRequest  %+v", 400, o.Payload)
}

func (o *AddDistroBadRequest) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddDistroBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDistroForbidden creates a AddDistroForbidden with default headers values
func NewAddDistroForbidden() *AddDistroForbidden {
	return &AddDistroForbidden{}
}

/*AddDistroForbidden handles this case with default header values.

Unauthorized
*/
type AddDistroForbidden struct {
	Payload *mono_models.Message
}

func (o *AddDistroForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros][%d] addDistroForbidden  %+v", 403, o.Payload)
}

func (o *AddDistroForbidden) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddDistroForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDistroConflict creates a AddDistroConflict with default headers values
func NewAddDistroConflict() *AddDistroConflict {
	return &AddDistroConflict{}
}

/*AddDistroConflict handles this case with default header values.

Conflict
*/
type AddDistroConflict struct {
	Payload *mono_models.Message
}

func (o *AddDistroConflict) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros][%d] addDistroConflict  %+v", 409, o.Payload)
}

func (o *AddDistroConflict) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddDistroConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDistroInternalServerError creates a AddDistroInternalServerError with default headers values
func NewAddDistroInternalServerError() *AddDistroInternalServerError {
	return &AddDistroInternalServerError{}
}

/*AddDistroInternalServerError handles this case with default header values.

Server Error
*/
type AddDistroInternalServerError struct {
	Payload *mono_models.Message
}

func (o *AddDistroInternalServerError) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros][%d] addDistroInternalServerError  %+v", 500, o.Payload)
}

func (o *AddDistroInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *AddDistroInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
