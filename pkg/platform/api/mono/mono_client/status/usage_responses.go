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

// UsageReader is a Reader for the Usage structure.
type UsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsageOK creates a UsageOK with default headers values
func NewUsageOK() *UsageOK {
	return &UsageOK{}
}

/*UsageOK handles this case with default header values.

Success
*/
type UsageOK struct {
	Payload *mono_models.UsageInfo
}

func (o *UsageOK) Error() string {
	return fmt.Sprintf("[GET /usage][%d] usageOK  %+v", 200, o.Payload)
}

func (o *UsageOK) GetPayload() *mono_models.UsageInfo {
	return o.Payload
}

func (o *UsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.UsageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsageBadRequest creates a UsageBadRequest with default headers values
func NewUsageBadRequest() *UsageBadRequest {
	return &UsageBadRequest{}
}

/*UsageBadRequest handles this case with default header values.

Bad Request
*/
type UsageBadRequest struct {
	Payload *mono_models.Message
}

func (o *UsageBadRequest) Error() string {
	return fmt.Sprintf("[GET /usage][%d] usageBadRequest  %+v", 400, o.Payload)
}

func (o *UsageBadRequest) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *UsageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsageForbidden creates a UsageForbidden with default headers values
func NewUsageForbidden() *UsageForbidden {
	return &UsageForbidden{}
}

/*UsageForbidden handles this case with default header values.

Unauthorized
*/
type UsageForbidden struct {
	Payload *mono_models.Message
}

func (o *UsageForbidden) Error() string {
	return fmt.Sprintf("[GET /usage][%d] usageForbidden  %+v", 403, o.Payload)
}

func (o *UsageForbidden) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *UsageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
