package connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// GetOchestratorsByTypeReader is a Reader for the GetOchestratorsByType structure.
type GetOchestratorsByTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetOchestratorsByTypeReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOchestratorsByTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOchestratorsByTypeOK creates a GetOchestratorsByTypeOK with default headers values
func NewGetOchestratorsByTypeOK() *GetOchestratorsByTypeOK {
	return &GetOchestratorsByTypeOK{}
}

/*GetOchestratorsByTypeOK handles this case with default header values.

successful operation
*/
type GetOchestratorsByTypeOK struct {
	Payload []string
}

func (o *GetOchestratorsByTypeOK) Error() string {
	return fmt.Sprintf("[GET /connectors/ochestrators/{type}][%d] getOchestratorsByTypeOK  %+v", 200, o.Payload)
}

func (o *GetOchestratorsByTypeOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
