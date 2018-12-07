// Code generated by go-swagger; DO NOT EDIT.

package v1networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPrivateNetworkReader is a Reader for the GetPrivateNetwork structure.
type GetPrivateNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateNetworkOK creates a GetPrivateNetworkOK with default headers values
func NewGetPrivateNetworkOK() *GetPrivateNetworkOK {
	return &GetPrivateNetworkOK{}
}

/*GetPrivateNetworkOK handles this case with default header values.

successful operation
*/
type GetPrivateNetworkOK struct {
	Payload *models_cloudbreak.NetworkResponse
}

func (o *GetPrivateNetworkOK) Error() string {
	return fmt.Sprintf("[GET /v1/networks/user/{name}][%d] getPrivateNetworkOK  %+v", 200, o.Payload)
}

func (o *GetPrivateNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.NetworkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}