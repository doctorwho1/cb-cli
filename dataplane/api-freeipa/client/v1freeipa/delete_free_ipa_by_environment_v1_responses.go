// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteFreeIpaByEnvironmentV1Reader is a Reader for the DeleteFreeIpaByEnvironmentV1 structure.
type DeleteFreeIpaByEnvironmentV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFreeIpaByEnvironmentV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteFreeIpaByEnvironmentV1Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteFreeIpaByEnvironmentV1Default creates a DeleteFreeIpaByEnvironmentV1Default with default headers values
func NewDeleteFreeIpaByEnvironmentV1Default(code int) *DeleteFreeIpaByEnvironmentV1Default {
	return &DeleteFreeIpaByEnvironmentV1Default{
		_statusCode: code,
	}
}

/*DeleteFreeIpaByEnvironmentV1Default handles this case with default header values.

successful operation
*/
type DeleteFreeIpaByEnvironmentV1Default struct {
	_statusCode int
}

// Code gets the status code for the delete free ipa by environment v1 default response
func (o *DeleteFreeIpaByEnvironmentV1Default) Code() int {
	return o._statusCode
}

func (o *DeleteFreeIpaByEnvironmentV1Default) Error() string {
	return fmt.Sprintf("[DELETE /v1/freeipa][%d] deleteFreeIpaByEnvironmentV1 default ", o._statusCode)
}

func (o *DeleteFreeIpaByEnvironmentV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}