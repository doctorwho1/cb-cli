// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StopSdxByCrnReader is a Reader for the StopSdxByCrn structure.
type StopSdxByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopSdxByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewStopSdxByCrnDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewStopSdxByCrnDefault creates a StopSdxByCrnDefault with default headers values
func NewStopSdxByCrnDefault(code int) *StopSdxByCrnDefault {
	return &StopSdxByCrnDefault{
		_statusCode: code,
	}
}

/*StopSdxByCrnDefault handles this case with default header values.

successful operation
*/
type StopSdxByCrnDefault struct {
	_statusCode int
}

// Code gets the status code for the stop sdx by crn default response
func (o *StopSdxByCrnDefault) Code() int {
	return o._statusCode
}

func (o *StopSdxByCrnDefault) Error() string {
	return fmt.Sprintf("[POST /sdx/crn/{crn}/stop][%d] stopSdxByCrn default ", o._statusCode)
}

func (o *StopSdxByCrnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}