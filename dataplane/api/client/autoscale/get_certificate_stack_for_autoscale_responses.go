// Code generated by go-swagger; DO NOT EDIT.

package autoscale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetCertificateStackForAutoscaleReader is a Reader for the GetCertificateStackForAutoscale structure.
type GetCertificateStackForAutoscaleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCertificateStackForAutoscaleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCertificateStackForAutoscaleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCertificateStackForAutoscaleOK creates a GetCertificateStackForAutoscaleOK with default headers values
func NewGetCertificateStackForAutoscaleOK() *GetCertificateStackForAutoscaleOK {
	return &GetCertificateStackForAutoscaleOK{}
}

/*GetCertificateStackForAutoscaleOK handles this case with default header values.

successful operation
*/
type GetCertificateStackForAutoscaleOK struct {
	Payload *model.CertificateV4Response
}

func (o *GetCertificateStackForAutoscaleOK) Error() string {
	return fmt.Sprintf("[GET /autoscale/stack/{id}/certificate][%d] getCertificateStackForAutoscaleOK  %+v", 200, o.Payload)
}

func (o *GetCertificateStackForAutoscaleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CertificateV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
