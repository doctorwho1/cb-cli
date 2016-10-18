package templates

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

// PostTemplatesUserReader is a Reader for the PostTemplatesUser structure.
type PostTemplatesUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostTemplatesUserReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostTemplatesUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTemplatesUserOK creates a PostTemplatesUserOK with default headers values
func NewPostTemplatesUserOK() *PostTemplatesUserOK {
	return &PostTemplatesUserOK{}
}

/*PostTemplatesUserOK handles this case with default header values.

successful operation
*/
type PostTemplatesUserOK struct {
	Payload *models.TemplateResponse
}

func (o *PostTemplatesUserOK) Error() string {
	return fmt.Sprintf("[POST /templates/user][%d] postTemplatesUserOK  %+v", 200, o.Payload)
}

func (o *PostTemplatesUserOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
