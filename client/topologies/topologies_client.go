package topologies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new topologies API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for topologies API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
DeleteTopologiesAccountID deletes topology by id

A topology gives system administrators an easy way to associate compute nodes with data centers and racks.
*/
func (a *Client) DeleteTopologiesAccountID(params *DeleteTopologiesAccountIDParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTopologiesAccountIDParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "DeleteTopologiesAccountID",
		Method:             "DELETE",
		PathPattern:        "/topologies/account/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTopologiesAccountIDReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
GetTopologiesAccount retrieves topoligies

A topology gives system administrators an easy way to associate compute nodes with data centers and racks.
*/
func (a *Client) GetTopologiesAccount(params *GetTopologiesAccountParams) (*GetTopologiesAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTopologiesAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetTopologiesAccount",
		Method:             "GET",
		PathPattern:        "/topologies/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTopologiesAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTopologiesAccountOK), nil
}

/*
GetTopologiesAccountID retrieves topology by id

A template gives developers and systems administrators an easy way to create and manage a collection of cloud infrastructure related resources, maintaining and updating them in an orderly and predictable fashion. Templates are cloud specific - and on top of the infrastructural setup they collect the information such as the used machine images, the datacenter location, instance types, and can capture and control region-specific infrastructure variations. We support heterogenous clusters - this one Hadoop cluster can be built by combining different templates.
*/
func (a *Client) GetTopologiesAccountID(params *GetTopologiesAccountIDParams) (*GetTopologiesAccountIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTopologiesAccountIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetTopologiesAccountID",
		Method:             "GET",
		PathPattern:        "/topologies/account/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTopologiesAccountIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTopologiesAccountIDOK), nil
}

/*
PostTopologiesAccount creates topology as public resource

A topology gives system administrators an easy way to associate compute nodes with data centers and racks.
*/
func (a *Client) PostTopologiesAccount(params *PostTopologiesAccountParams) (*PostTopologiesAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTopologiesAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "PostTopologiesAccount",
		Method:             "POST",
		PathPattern:        "/topologies/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostTopologiesAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTopologiesAccountOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}
