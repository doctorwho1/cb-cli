// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeImageStackInWorkspaceV4Reader is a Reader for the ChangeImageStackInWorkspaceV4 structure.
type ChangeImageStackInWorkspaceV4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeImageStackInWorkspaceV4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewChangeImageStackInWorkspaceV4Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewChangeImageStackInWorkspaceV4Default creates a ChangeImageStackInWorkspaceV4Default with default headers values
func NewChangeImageStackInWorkspaceV4Default(code int) *ChangeImageStackInWorkspaceV4Default {
	return &ChangeImageStackInWorkspaceV4Default{
		_statusCode: code,
	}
}

/*ChangeImageStackInWorkspaceV4Default handles this case with default header values.

successful operation
*/
type ChangeImageStackInWorkspaceV4Default struct {
	_statusCode int
}

// Code gets the status code for the change image stack in workspace v4 default response
func (o *ChangeImageStackInWorkspaceV4Default) Code() int {
	return o._statusCode
}

func (o *ChangeImageStackInWorkspaceV4Default) Error() string {
	return fmt.Sprintf("[PUT /v4/{workspaceId}/stacks/{name}/change_image][%d] changeImageStackInWorkspaceV4 default ", o._statusCode)
}

func (o *ChangeImageStackInWorkspaceV4Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}