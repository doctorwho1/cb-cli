// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// GetFlowLogsByResourceNameReader is a Reader for the GetFlowLogsByResourceName structure.
type GetFlowLogsByResourceNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowLogsByResourceNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFlowLogsByResourceNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFlowLogsByResourceNameOK creates a GetFlowLogsByResourceNameOK with default headers values
func NewGetFlowLogsByResourceNameOK() *GetFlowLogsByResourceNameOK {
	return &GetFlowLogsByResourceNameOK{}
}

/*GetFlowLogsByResourceNameOK handles this case with default header values.

successful operation
*/
type GetFlowLogsByResourceNameOK struct {
	Payload []*model.FlowLogResponse
}

func (o *GetFlowLogsByResourceNameOK) Error() string {
	return fmt.Sprintf("[GET /flow/logs/resource/name/{resourceName}][%d] getFlowLogsByResourceNameOK  %+v", 200, o.Payload)
}

func (o *GetFlowLogsByResourceNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}