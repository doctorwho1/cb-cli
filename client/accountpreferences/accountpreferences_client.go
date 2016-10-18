package accountpreferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new accountpreferences API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for accountpreferences API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
Get retrieves account preferences for admin user

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) Get(params *GetParams) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "get",
		Method:             "GET",
		PathPattern:        "/accountpreferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOK), nil
}

/*
Post posts account preferences of admin user

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) Post(params *PostParams) (*PostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "post",
		Method:             "POST",
		PathPattern:        "/accountpreferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &PostReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostOK), nil
}

/*
Put updates account preferences of admin user

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) Put(params *PutParams) (*PutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "put",
		Method:             "PUT",
		PathPattern:        "/accountpreferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &PutReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutOK), nil
}

/*
Validate validates account preferences of all stacks

Account related preferences that could be managed by the account admins and different restrictions could be added to Cloudbreak resources.
*/
func (a *Client) Validate(params *ValidateParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "validate",
		Method:             "GET",
		PathPattern:        "/accountpreferences/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &ValidateReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}
