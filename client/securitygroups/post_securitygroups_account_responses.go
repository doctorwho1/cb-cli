package securitygroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models"
)

// PostSecuritygroupsAccountReader is a Reader for the PostSecuritygroupsAccount structure.
type PostSecuritygroupsAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostSecuritygroupsAccountReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostSecuritygroupsAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSecuritygroupsAccountOK creates a PostSecuritygroupsAccountOK with default headers values
func NewPostSecuritygroupsAccountOK() *PostSecuritygroupsAccountOK {
	return &PostSecuritygroupsAccountOK{}
}

/*PostSecuritygroupsAccountOK handles this case with default header values.

successful operation
*/
type PostSecuritygroupsAccountOK struct {
	Payload *models.SecurityGroupResponse
}

func (o *PostSecuritygroupsAccountOK) Error() string {
	return fmt.Sprintf("[POST /securitygroups/account][%d] postSecuritygroupsAccountOK  %+v", 200, o.Payload)
}

func (o *PostSecuritygroupsAccountOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
