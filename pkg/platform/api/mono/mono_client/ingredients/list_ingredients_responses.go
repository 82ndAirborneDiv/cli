// Code generated by go-swagger; DO NOT EDIT.

package ingredients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// ListIngredientsReader is a Reader for the ListIngredients structure.
type ListIngredientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIngredientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListIngredientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListIngredientsOK creates a ListIngredientsOK with default headers values
func NewListIngredientsOK() *ListIngredientsOK {
	return &ListIngredientsOK{}
}

/*ListIngredientsOK handles this case with default header values.

Success
*/
type ListIngredientsOK struct {
	Payload []*mono_models.Ingredient
}

func (o *ListIngredientsOK) Error() string {
	return fmt.Sprintf("[GET /ingredients][%d] listIngredientsOK  %+v", 200, o.Payload)
}

func (o *ListIngredientsOK) GetPayload() []*mono_models.Ingredient {
	return o.Payload
}

func (o *ListIngredientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
