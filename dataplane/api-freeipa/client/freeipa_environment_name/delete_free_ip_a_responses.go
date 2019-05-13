// Code generated by go-swagger; DO NOT EDIT.

package freeipa_environment_name

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteFreeIPAReader is a Reader for the DeleteFreeIPA structure.
type DeleteFreeIPAReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFreeIPAReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteFreeIPADefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteFreeIPADefault creates a DeleteFreeIPADefault with default headers values
func NewDeleteFreeIPADefault(code int) *DeleteFreeIPADefault {
	return &DeleteFreeIPADefault{
		_statusCode: code,
	}
}

/*DeleteFreeIPADefault handles this case with default header values.

successful operation
*/
type DeleteFreeIPADefault struct {
	_statusCode int
}

// Code gets the status code for the delete free IP a default response
func (o *DeleteFreeIPADefault) Code() int {
	return o._statusCode
}

func (o *DeleteFreeIPADefault) Error() string {
	return fmt.Sprintf("[DELETE /freeipa/{environmentName}/{name}][%d] deleteFreeIPA default ", o._statusCode)
}

func (o *DeleteFreeIPADefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}