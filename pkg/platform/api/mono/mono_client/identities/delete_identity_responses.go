// Code generated by go-swagger; DO NOT EDIT.

package identities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// DeleteIdentityReader is a Reader for the DeleteIdentity structure.
type DeleteIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIdentityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteIdentityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteIdentityOK creates a DeleteIdentityOK with default headers values
func NewDeleteIdentityOK() *DeleteIdentityOK {
	return &DeleteIdentityOK{}
}

/*DeleteIdentityOK handles this case with default header values.

Identity deleted
*/
type DeleteIdentityOK struct {
	Payload *mono_models.Message
}

func (o *DeleteIdentityOK) Error() string {
	return fmt.Sprintf("[DELETE /identities/{identityID}][%d] deleteIdentityOK  %+v", 200, o.Payload)
}

func (o *DeleteIdentityOK) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *DeleteIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityBadRequest creates a DeleteIdentityBadRequest with default headers values
func NewDeleteIdentityBadRequest() *DeleteIdentityBadRequest {
	return &DeleteIdentityBadRequest{}
}

/*DeleteIdentityBadRequest handles this case with default header values.

Bad Request
*/
type DeleteIdentityBadRequest struct {
	Payload *mono_models.Message
}

func (o *DeleteIdentityBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /identities/{identityID}][%d] deleteIdentityBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIdentityBadRequest) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *DeleteIdentityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityForbidden creates a DeleteIdentityForbidden with default headers values
func NewDeleteIdentityForbidden() *DeleteIdentityForbidden {
	return &DeleteIdentityForbidden{}
}

/*DeleteIdentityForbidden handles this case with default header values.

Forbidden
*/
type DeleteIdentityForbidden struct {
}

func (o *DeleteIdentityForbidden) Error() string {
	return fmt.Sprintf("[DELETE /identities/{identityID}][%d] deleteIdentityForbidden ", 403)
}

func (o *DeleteIdentityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
