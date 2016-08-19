package recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new recipes API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for recipes API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
DeleteRecipesAccountName deletes public owned or private recipe by name

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) DeleteRecipesAccountName(params *DeleteRecipesAccountNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRecipesAccountNameParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "DeleteRecipesAccountName",
		Method:             "DELETE",
		PathPattern:        "/recipes/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRecipesAccountNameReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
DeleteRecipesID deletes recipe by id

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) DeleteRecipesID(params *DeleteRecipesIDParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRecipesIDParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "DeleteRecipesID",
		Method:             "DELETE",
		PathPattern:        "/recipes/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRecipesIDReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
DeleteRecipesUserName deletes private recipe by name

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) DeleteRecipesUserName(params *DeleteRecipesUserNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRecipesUserNameParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "DeleteRecipesUserName",
		Method:             "DELETE",
		PathPattern:        "/recipes/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRecipesUserNameReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
GetRecipesAccount retrieves public and private owned recipes

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) GetRecipesAccount(params *GetRecipesAccountParams) (*GetRecipesAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecipesAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetRecipesAccount",
		Method:             "GET",
		PathPattern:        "/recipes/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRecipesAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRecipesAccountOK), nil
}

/*
GetRecipesAccountName retrieves a public or private owned recipe by name

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) GetRecipesAccountName(params *GetRecipesAccountNameParams) (*GetRecipesAccountNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecipesAccountNameParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetRecipesAccountName",
		Method:             "GET",
		PathPattern:        "/recipes/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRecipesAccountNameReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRecipesAccountNameOK), nil
}

/*
GetRecipesID retrieves recipe by id

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) GetRecipesID(params *GetRecipesIDParams) (*GetRecipesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecipesIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetRecipesID",
		Method:             "GET",
		PathPattern:        "/recipes/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRecipesIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRecipesIDOK), nil
}

/*
GetRecipesUser retrieves private recipes

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) GetRecipesUser(params *GetRecipesUserParams) (*GetRecipesUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecipesUserParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetRecipesUser",
		Method:             "GET",
		PathPattern:        "/recipes/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRecipesUserReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRecipesUserOK), nil
}

/*
GetRecipesUserName retrieves a private recipe by name

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) GetRecipesUserName(params *GetRecipesUserNameParams) (*GetRecipesUserNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRecipesUserNameParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetRecipesUserName",
		Method:             "GET",
		PathPattern:        "/recipes/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRecipesUserNameReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRecipesUserNameOK), nil
}

/*
PostRecipesAccount creates recipe as public resource

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) PostRecipesAccount(params *PostRecipesAccountParams) (*PostRecipesAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRecipesAccountParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "PostRecipesAccount",
		Method:             "POST",
		PathPattern:        "/recipes/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRecipesAccountReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRecipesAccountOK), nil
}

/*
PostRecipesUser creates recipe as private resource

Recipes are basically script extensions to a cluster that run on a set of nodes before or after the Ambari cluster installation.
*/
func (a *Client) PostRecipesUser(params *PostRecipesUserParams) (*PostRecipesUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRecipesUserParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "PostRecipesUser",
		Method:             "POST",
		PathPattern:        "/recipes/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRecipesUserReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRecipesUserOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}
