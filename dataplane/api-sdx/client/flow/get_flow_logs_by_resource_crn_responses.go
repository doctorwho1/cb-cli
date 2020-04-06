// Code generated by go-swagger; DO NOT EDIT.

package flow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// GetFlowLogsByResourceCrnReader is a Reader for the GetFlowLogsByResourceCrn structure.
type GetFlowLogsByResourceCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowLogsByResourceCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFlowLogsByResourceCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFlowLogsByResourceCrnOK creates a GetFlowLogsByResourceCrnOK with default headers values
func NewGetFlowLogsByResourceCrnOK() *GetFlowLogsByResourceCrnOK {
	return &GetFlowLogsByResourceCrnOK{}
}

/*GetFlowLogsByResourceCrnOK handles this case with default header values.

successful operation
*/
type GetFlowLogsByResourceCrnOK struct {
	Payload []*model.FlowLogResponse
}

func (o *GetFlowLogsByResourceCrnOK) Error() string {
	return fmt.Sprintf("[GET /flow/logs/resource/crn/{resourceCrn}][%d] getFlowLogsByResourceCrnOK  %+v", 200, o.Payload)
}

func (o *GetFlowLogsByResourceCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}