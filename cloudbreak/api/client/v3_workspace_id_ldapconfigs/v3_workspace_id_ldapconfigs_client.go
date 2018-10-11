// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_ldapconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v3 workspace id ldapconfigs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v3 workspace id ldapconfigs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AttachLdapResourceToEnvironments attaches ldap resource to environemnts

LDAP server integration enables the user to provide a central place to store usernames and passwords for the users of his/her clusters.
*/
func (a *Client) AttachLdapResourceToEnvironments(params *AttachLdapResourceToEnvironmentsParams) (*AttachLdapResourceToEnvironmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachLdapResourceToEnvironmentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "attachLdapResourceToEnvironments",
		Method:             "PUT",
		PathPattern:        "/v3/{workspaceId}/ldapconfigs/{name}/attach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AttachLdapResourceToEnvironmentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AttachLdapResourceToEnvironmentsOK), nil

}

/*
CreateLdapConfigsInWorkspace creates l d a p config in workspace

LDAP server integration enables the user to provide a central place to store usernames and passwords for the users of his/her clusters.
*/
func (a *Client) CreateLdapConfigsInWorkspace(params *CreateLdapConfigsInWorkspaceParams) (*CreateLdapConfigsInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLdapConfigsInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createLdapConfigsInWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/ldapconfigs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateLdapConfigsInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateLdapConfigsInWorkspaceOK), nil

}

/*
DeleteLdapConfigsInWorkspace deletes l d a p config by name in workspace

LDAP server integration enables the user to provide a central place to store usernames and passwords for the users of his/her clusters.
*/
func (a *Client) DeleteLdapConfigsInWorkspace(params *DeleteLdapConfigsInWorkspaceParams) (*DeleteLdapConfigsInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLdapConfigsInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteLdapConfigsInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v3/{workspaceId}/ldapconfigs/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteLdapConfigsInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLdapConfigsInWorkspaceOK), nil

}

/*
DetachLdapResourceFromEnvironments detaches ldap resource from environemnts

LDAP server integration enables the user to provide a central place to store usernames and passwords for the users of his/her clusters.
*/
func (a *Client) DetachLdapResourceFromEnvironments(params *DetachLdapResourceFromEnvironmentsParams) (*DetachLdapResourceFromEnvironmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetachLdapResourceFromEnvironmentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "detachLdapResourceFromEnvironments",
		Method:             "PUT",
		PathPattern:        "/v3/{workspaceId}/ldapconfigs/{name}/detach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DetachLdapResourceFromEnvironmentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DetachLdapResourceFromEnvironmentsOK), nil

}

/*
GetLdapConfigInWorkspace gets l d a p config by name in workspace

LDAP server integration enables the user to provide a central place to store usernames and passwords for the users of his/her clusters.
*/
func (a *Client) GetLdapConfigInWorkspace(params *GetLdapConfigInWorkspaceParams) (*GetLdapConfigInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLdapConfigInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLdapConfigInWorkspace",
		Method:             "GET",
		PathPattern:        "/v3/{workspaceId}/ldapconfigs/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLdapConfigInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLdapConfigInWorkspaceOK), nil

}

/*
GetLdapRequestByNameAndWorkspaceID gets request

LDAP server integration enables the user to provide a central place to store usernames and passwords for the users of his/her clusters.
*/
func (a *Client) GetLdapRequestByNameAndWorkspaceID(params *GetLdapRequestByNameAndWorkspaceIDParams) (*GetLdapRequestByNameAndWorkspaceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLdapRequestByNameAndWorkspaceIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLdapRequestByNameAndWorkspaceId",
		Method:             "GET",
		PathPattern:        "/v3/{workspaceId}/ldapconfigs/{name}/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLdapRequestByNameAndWorkspaceIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLdapRequestByNameAndWorkspaceIDOK), nil

}

/*
ListLdapsByWorkspace lists l d a p configs for the given workspace

LDAP server integration enables the user to provide a central place to store usernames and passwords for the users of his/her clusters.
*/
func (a *Client) ListLdapsByWorkspace(params *ListLdapsByWorkspaceParams) (*ListLdapsByWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListLdapsByWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listLdapsByWorkspace",
		Method:             "GET",
		PathPattern:        "/v3/{workspaceId}/ldapconfigs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListLdapsByWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListLdapsByWorkspaceOK), nil

}

/*
PostLdapConnectionTestInWorkspace tests that the connection could be established of an existing or new l d a p config
*/
func (a *Client) PostLdapConnectionTestInWorkspace(params *PostLdapConnectionTestInWorkspaceParams) (*PostLdapConnectionTestInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLdapConnectionTestInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLdapConnectionTestInWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/ldapconfigs/testconnect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostLdapConnectionTestInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLdapConnectionTestInWorkspaceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
