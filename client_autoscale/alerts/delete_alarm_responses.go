package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeleteAlarmReader is a Reader for the DeleteAlarm structure.
type DeleteAlarmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteAlarmReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewDeleteAlarmDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteAlarmDefault creates a DeleteAlarmDefault with default headers values
func NewDeleteAlarmDefault(code int) *DeleteAlarmDefault {
	return &DeleteAlarmDefault{
		_statusCode: code,
	}
}

/*DeleteAlarmDefault handles this case with default header values.

successful operation
*/
type DeleteAlarmDefault struct {
	_statusCode int
}

// Code gets the status code for the delete alarm default response
func (o *DeleteAlarmDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAlarmDefault) Error() string {
	return fmt.Sprintf("[DELETE /clusters/{clusterId}/alerts/metric/{alertId}][%d] deleteAlarm default ", o._statusCode)
}

func (o *DeleteAlarmDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}