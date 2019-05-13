// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// CreateProxyConfigInWorkspaceReader is a Reader for the CreateProxyConfigInWorkspace structure.
type CreateProxyConfigInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProxyConfigInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateProxyConfigInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateProxyConfigInWorkspaceOK creates a CreateProxyConfigInWorkspaceOK with default headers values
func NewCreateProxyConfigInWorkspaceOK() *CreateProxyConfigInWorkspaceOK {
	return &CreateProxyConfigInWorkspaceOK{}
}

/*CreateProxyConfigInWorkspaceOK handles this case with default header values.

successful operation
*/
type CreateProxyConfigInWorkspaceOK struct {
	Payload *model.ProxyV4Response
}

func (o *CreateProxyConfigInWorkspaceOK) Error() string {
	return fmt.Sprintf("[POST /v4/{workspaceId}/proxies][%d] createProxyConfigInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *CreateProxyConfigInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ProxyV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}