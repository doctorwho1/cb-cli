package connectors

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

// GetDisktypesReader is a Reader for the GetDisktypes structure.
type GetDisktypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetDisktypesReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDisktypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDisktypesOK creates a GetDisktypesOK with default headers values
func NewGetDisktypesOK() *GetDisktypesOK {
	return &GetDisktypesOK{}
}

/*GetDisktypesOK handles this case with default header values.

successful operation
*/
type GetDisktypesOK struct {
	Payload *models.PlatformDisksJSON
}

func (o *GetDisktypesOK) Error() string {
	return fmt.Sprintf("[GET /connectors/disktypes][%d] getDisktypesOK  %+v", 200, o.Payload)
}

func (o *GetDisktypesOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlatformDisksJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
