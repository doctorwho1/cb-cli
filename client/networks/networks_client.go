package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new networks API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for networks API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
DeleteNetworksAccountName deletes public owned or private network by name

Provider specific network settings could be configured by using Network resources.
*/
func (a *Client) DeleteNetworksAccountName(params *DeleteNetworksAccountNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworksAccountNameParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "DeleteNetworksAccountName",
		Method:             "DELETE",
		PathPattern:        "/networks/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &DeleteNetworksAccountNameReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
DeleteNetworksID deletes network by id

Provider specific network settings could be configured by using Network resources.
*/
func (a *Client) DeleteNetworksID(params *DeleteNetworksIDParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworksIDParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "DeleteNetworksID",
		Method:             "DELETE",
		PathPattern:        "/networks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &DeleteNetworksIDReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
DeleteNetworksUserName deletes private network by name

Provider specific network settings could be configured by using Network resources.
*/
func (a *Client) DeleteNetworksUserName(params *DeleteNetworksUserNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworksUserNameParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "DeleteNetworksUserName",
		Method:             "DELETE",
		PathPattern:        "/networks/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &DeleteNetworksUserNameReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
GetNetworksAccount retrieves public and private owned networks

Provider specific network settings could be configured by using Network resources.
*/
func (a *Client) GetNetworksAccount(params *GetNetworksAccountParams) (*GetNetworksAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetNetworksAccount",
		Method:             "GET",
		PathPattern:        "/networks/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &GetNetworksAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNetworksAccountOK), nil
}

/*
GetNetworksAccountName retrieves a public or private owned network by name

Provider specific network settings could be configured by using Network resources.
*/
func (a *Client) GetNetworksAccountName(params *GetNetworksAccountNameParams) (*GetNetworksAccountNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksAccountNameParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetNetworksAccountName",
		Method:             "GET",
		PathPattern:        "/networks/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &GetNetworksAccountNameReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNetworksAccountNameOK), nil
}

/*
GetNetworksID retrieves network by id

Provider specific network settings could be configured by using Network resources.
*/
func (a *Client) GetNetworksID(params *GetNetworksIDParams) (*GetNetworksIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetNetworksID",
		Method:             "GET",
		PathPattern:        "/networks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &GetNetworksIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNetworksIDOK), nil
}

/*
GetNetworksUser retrieves private networks

Provider specific network settings could be configured by using Network resources.
*/
func (a *Client) GetNetworksUser(params *GetNetworksUserParams) (*GetNetworksUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksUserParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetNetworksUser",
		Method:             "GET",
		PathPattern:        "/networks/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &GetNetworksUserReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNetworksUserOK), nil
}

/*
GetNetworksUserName retrieves a private network by name

Provider specific network settings could be configured by using Network resources.
*/
func (a *Client) GetNetworksUserName(params *GetNetworksUserNameParams) (*GetNetworksUserNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworksUserNameParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetNetworksUserName",
		Method:             "GET",
		PathPattern:        "/networks/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &GetNetworksUserNameReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNetworksUserNameOK), nil
}

/*
PostNetworksAccount creates network as public resource

Provider specific network settings could be configured by using Network resources.
*/
func (a *Client) PostNetworksAccount(params *PostNetworksAccountParams) (*PostNetworksAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNetworksAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "PostNetworksAccount",
		Method:             "POST",
		PathPattern:        "/networks/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &PostNetworksAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNetworksAccountOK), nil
}

/*
PostNetworksUser creates network as private resource

Provider specific network settings could be configured by using Network resources.
*/
func (a *Client) PostNetworksUser(params *PostNetworksUserParams) (*PostNetworksUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNetworksUserParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "PostNetworksUser",
		Method:             "POST",
		PathPattern:        "/networks/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &PostNetworksUserReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNetworksUserOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}
