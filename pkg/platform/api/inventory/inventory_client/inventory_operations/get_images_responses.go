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

// GetImagesReader is a Reader for the GetImages structure.
type GetImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetImagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetImagesOK creates a GetImagesOK with default headers values
func NewGetImagesOK() *GetImagesOK {
	return &GetImagesOK{}
}

/*GetImagesOK handles this case with default header values.

A paginated list of images
*/
type GetImagesOK struct {
	Payload *inventory_models.V1ImagePagedList
}

func (o *GetImagesOK) Error() string {
	return fmt.Sprintf("[GET /v1/images][%d] getImagesOK  %+v", 200, o.Payload)
}

func (o *GetImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1ImagePagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImagesDefault creates a GetImagesDefault with default headers values
func NewGetImagesDefault(code int) *GetImagesDefault {
	return &GetImagesDefault{
		_statusCode: code,
	}
}

/*GetImagesDefault handles this case with default header values.

generic error response
*/
type GetImagesDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get images default response
func (o *GetImagesDefault) Code() int {
	return o._statusCode
}

func (o *GetImagesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/images][%d] getImages default  %+v", o._statusCode, o.Payload)
}

func (o *GetImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
