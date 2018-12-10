// Code generated by go-swagger; DO NOT EDIT.

package v1subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// SubscribeSubscriptionReader is a Reader for the SubscribeSubscription structure.
type SubscribeSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscribeSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSubscribeSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubscribeSubscriptionOK creates a SubscribeSubscriptionOK with default headers values
func NewSubscribeSubscriptionOK() *SubscribeSubscriptionOK {
	return &SubscribeSubscriptionOK{}
}

/*SubscribeSubscriptionOK handles this case with default header values.

successful operation
*/
type SubscribeSubscriptionOK struct {
	Payload *model.ID
}

func (o *SubscribeSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /v1/subscriptions][%d] subscribeSubscriptionOK  %+v", 200, o.Payload)
}

func (o *SubscribeSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}