// Code generated by go-swagger; DO NOT EDIT.

package connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// GetImagesByTypeReader is a Reader for the GetImagesByType structure.
type GetImagesByTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImagesByTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetImagesByTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetImagesByTypeOK creates a GetImagesByTypeOK with default headers values
func NewGetImagesByTypeOK() *GetImagesByTypeOK {
	return &GetImagesByTypeOK{}
}

/*GetImagesByTypeOK handles this case with default header values.

successful operation
*/
type GetImagesByTypeOK struct {
	Payload GetImagesByTypeOKBody
}

func (o *GetImagesByTypeOK) Error() string {
	return fmt.Sprintf("[GET /connectors/images/{type}][%d] getImagesByTypeOK  %+v", 200, o.Payload)
}

func (o *GetImagesByTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetImagesByTypeOKBody get images by type o k body
swagger:model GetImagesByTypeOKBody
*/

type GetImagesByTypeOKBody map[string]string

// Validate validates this get images by type o k body
func (o GetImagesByTypeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if swag.IsZero(o) { // not required
		return nil
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
