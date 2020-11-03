// Code generated by go-swagger; DO NOT EDIT.

package version_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// GetRevertCommitReader is a Reader for the GetRevertCommit structure.
type GetRevertCommitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRevertCommitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRevertCommitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRevertCommitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRevertCommitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRevertCommitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRevertCommitOK creates a GetRevertCommitOK with default headers values
func NewGetRevertCommitOK() *GetRevertCommitOK {
	return &GetRevertCommitOK{}
}

/*GetRevertCommitOK handles this case with default header values.

Calculate a revert commit
*/
type GetRevertCommitOK struct {
	Payload *mono_models.Commit
}

func (o *GetRevertCommitOK) Error() string {
	return fmt.Sprintf("[GET /vcs/commits/{commitFromID}/revert/{commitToID}][%d] getRevertCommitOK  %+v", 200, o.Payload)
}

func (o *GetRevertCommitOK) GetPayload() *mono_models.Commit {
	return o.Payload
}

func (o *GetRevertCommitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Commit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRevertCommitForbidden creates a GetRevertCommitForbidden with default headers values
func NewGetRevertCommitForbidden() *GetRevertCommitForbidden {
	return &GetRevertCommitForbidden{}
}

/*GetRevertCommitForbidden handles this case with default header values.

Forbidden
*/
type GetRevertCommitForbidden struct {
	Payload *mono_models.Message
}

func (o *GetRevertCommitForbidden) Error() string {
	return fmt.Sprintf("[GET /vcs/commits/{commitFromID}/revert/{commitToID}][%d] getRevertCommitForbidden  %+v", 403, o.Payload)
}

func (o *GetRevertCommitForbidden) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetRevertCommitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRevertCommitNotFound creates a GetRevertCommitNotFound with default headers values
func NewGetRevertCommitNotFound() *GetRevertCommitNotFound {
	return &GetRevertCommitNotFound{}
}

/*GetRevertCommitNotFound handles this case with default header values.

commit was not found
*/
type GetRevertCommitNotFound struct {
	Payload *mono_models.Message
}

func (o *GetRevertCommitNotFound) Error() string {
	return fmt.Sprintf("[GET /vcs/commits/{commitFromID}/revert/{commitToID}][%d] getRevertCommitNotFound  %+v", 404, o.Payload)
}

func (o *GetRevertCommitNotFound) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetRevertCommitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRevertCommitInternalServerError creates a GetRevertCommitInternalServerError with default headers values
func NewGetRevertCommitInternalServerError() *GetRevertCommitInternalServerError {
	return &GetRevertCommitInternalServerError{}
}

/*GetRevertCommitInternalServerError handles this case with default header values.

error calculating commit
*/
type GetRevertCommitInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetRevertCommitInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vcs/commits/{commitFromID}/revert/{commitToID}][%d] getRevertCommitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRevertCommitInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetRevertCommitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}