package flexsubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetPublicFlexSubscriptionByNameReader is a Reader for the GetPublicFlexSubscriptionByName structure.
type GetPublicFlexSubscriptionByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPublicFlexSubscriptionByNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicFlexSubscriptionByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicFlexSubscriptionByNameOK creates a GetPublicFlexSubscriptionByNameOK with default headers values
func NewGetPublicFlexSubscriptionByNameOK() *GetPublicFlexSubscriptionByNameOK {
	return &GetPublicFlexSubscriptionByNameOK{}
}

/*GetPublicFlexSubscriptionByNameOK handles this case with default header values.

successful operation
*/
type GetPublicFlexSubscriptionByNameOK struct {
	Payload *models_cloudbreak.FlexSubscriptionResponse
}

func (o *GetPublicFlexSubscriptionByNameOK) Error() string {
	return fmt.Sprintf("[GET /flexsubscriptions/account/{name}][%d] getPublicFlexSubscriptionByNameOK  %+v", 200, o.Payload)
}

func (o *GetPublicFlexSubscriptionByNameOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.FlexSubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}