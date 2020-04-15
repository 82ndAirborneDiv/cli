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

// GetOrderReader is a Reader for the GetOrder structure.
type GetOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetOrderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrderOK creates a GetOrderOK with default headers values
func NewGetOrderOK() *GetOrderOK {
	return &GetOrderOK{}
}

/*GetOrderOK handles this case with default header values.

Get the solver order for the given commit
*/
type GetOrderOK struct {
	Payload *mono_models.Order
}

func (o *GetOrderOK) Error() string {
	return fmt.Sprintf("[GET /vcs/commits/{commitID}/order][%d] getOrderOK  %+v", 200, o.Payload)
}

func (o *GetOrderOK) GetPayload() *mono_models.Order {
	return o.Payload
}

func (o *GetOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderNotFound creates a GetOrderNotFound with default headers values
func NewGetOrderNotFound() *GetOrderNotFound {
	return &GetOrderNotFound{}
}

/*GetOrderNotFound handles this case with default header values.

order was not found
*/
type GetOrderNotFound struct {
	Payload *mono_models.Message
}

func (o *GetOrderNotFound) Error() string {
	return fmt.Sprintf("[GET /vcs/commits/{commitID}/order][%d] getOrderNotFound  %+v", 404, o.Payload)
}

func (o *GetOrderNotFound) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetOrderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderInternalServerError creates a GetOrderInternalServerError with default headers values
func NewGetOrderInternalServerError() *GetOrderInternalServerError {
	return &GetOrderInternalServerError{}
}

/*GetOrderInternalServerError handles this case with default header values.

error retrieving order
*/
type GetOrderInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetOrderInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vcs/commits/{commitID}/order][%d] getOrderInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrderInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetOrderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
