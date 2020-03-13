// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	inventory_models "github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// SolveOrderReader is a Reader for the SolveOrder structure.
type SolveOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SolveOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewSolveOrderCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSolveOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSolveOrderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSolveOrderCreated creates a SolveOrderCreated with default headers values
func NewSolveOrderCreated() *SolveOrderCreated {
	return &SolveOrderCreated{
		CacheControl: "private, max-age 3600",
	}
}

/*SolveOrderCreated handles this case with default header values.

Returns the ids of and links to the created recipes
*/
type SolveOrderCreated struct {
	CacheControl string

	Payload inventory_models.V1SolutionResponse
}

func (o *SolveOrderCreated) Error() string {
	return fmt.Sprintf("[POST /v1/solutions][%d] solveOrderCreated  %+v", 201, o.Payload)
}

func (o *SolveOrderCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSolveOrderBadRequest creates a SolveOrderBadRequest with default headers values
func NewSolveOrderBadRequest() *SolveOrderBadRequest {
	return &SolveOrderBadRequest{}
}

/*SolveOrderBadRequest handles this case with default header values.

If the order is invalid
*/
type SolveOrderBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *SolveOrderBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/solutions][%d] solveOrderBadRequest  %+v", 400, o.Payload)
}

func (o *SolveOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSolveOrderDefault creates a SolveOrderDefault with default headers values
func NewSolveOrderDefault(code int) *SolveOrderDefault {
	return &SolveOrderDefault{
		_statusCode: code,
	}
}

/*SolveOrderDefault handles this case with default header values.

If there is an error processing the order
*/
type SolveOrderDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the solve order default response
func (o *SolveOrderDefault) Code() int {
	return o._statusCode
}

func (o *SolveOrderDefault) Error() string {
	return fmt.Sprintf("[POST /v1/solutions][%d] solveOrder default  %+v", o._statusCode, o.Payload)
}

func (o *SolveOrderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
