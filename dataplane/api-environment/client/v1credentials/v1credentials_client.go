// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1credentials API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1credentials API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AuthorizeCodeGrantFlowBasedCredentialV1 authorizes oauth2 authorization code grant flow

Authorize code grant flow based credential creation.
*/
func (a *Client) AuthorizeCodeGrantFlowBasedCredentialV1(params *AuthorizeCodeGrantFlowBasedCredentialV1Params) (*AuthorizeCodeGrantFlowBasedCredentialV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthorizeCodeGrantFlowBasedCredentialV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "authorizeCodeGrantFlowBasedCredentialV1",
		Method:             "GET",
		PathPattern:        "/v1/credentials/code_grant_flow/authorization/{cloudPlatform}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AuthorizeCodeGrantFlowBasedCredentialV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthorizeCodeGrantFlowBasedCredentialV1OK), nil

}

/*
CodeGrantFlowBasedCredentialV1 starts a credential creation with oauth2 authorization code grant flow

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) CodeGrantFlowBasedCredentialV1(params *CodeGrantFlowBasedCredentialV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCodeGrantFlowBasedCredentialV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "codeGrantFlowBasedCredentialV1",
		Method:             "POST",
		PathPattern:        "/v1/credentials/code_grant_flow/init",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CodeGrantFlowBasedCredentialV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
CodeGrantFlowOnExistingCredentialV1 reinitializes oauth2 authorization code grant flow on an existing credential

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) CodeGrantFlowOnExistingCredentialV1(params *CodeGrantFlowOnExistingCredentialV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCodeGrantFlowOnExistingCredentialV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "codeGrantFlowOnExistingCredentialV1",
		Method:             "GET",
		PathPattern:        "/v1/credentials/code_grant_flow/init/{name}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CodeGrantFlowOnExistingCredentialV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
CreateCredentialV1 creates credential

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) CreateCredentialV1(params *CreateCredentialV1Params) (*CreateCredentialV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCredentialV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCredentialV1",
		Method:             "POST",
		PathPattern:        "/v1/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateCredentialV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateCredentialV1OK), nil

}

/*
DeleteCredentialByNameV1 deletes credential by name

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) DeleteCredentialByNameV1(params *DeleteCredentialByNameV1Params) (*DeleteCredentialByNameV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCredentialByNameV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCredentialByNameV1",
		Method:             "DELETE",
		PathPattern:        "/v1/credentials/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCredentialByNameV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCredentialByNameV1OK), nil

}

/*
DeleteCredentialByResourceCrnV1 deletes credential by crn

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) DeleteCredentialByResourceCrnV1(params *DeleteCredentialByResourceCrnV1Params) (*DeleteCredentialByResourceCrnV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCredentialByResourceCrnV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCredentialByResourceCrnV1",
		Method:             "DELETE",
		PathPattern:        "/v1/credentials/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCredentialByResourceCrnV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCredentialByResourceCrnV1OK), nil

}

/*
DeleteCredentialsV1 deletes multiple credentials by name

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) DeleteCredentialsV1(params *DeleteCredentialsV1Params) (*DeleteCredentialsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCredentialsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCredentialsV1",
		Method:             "DELETE",
		PathPattern:        "/v1/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCredentialsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCredentialsV1OK), nil

}

/*
GetCreateCredentialForCli produces cli command input for credential creation

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) GetCreateCredentialForCli(params *GetCreateCredentialForCliParams) (*GetCreateCredentialForCliOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCreateCredentialForCliParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCreateCredentialForCli",
		Method:             "POST",
		PathPattern:        "/v1/credentials/cli_create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCreateCredentialForCliReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCreateCredentialForCliOK), nil

}

/*
GetCredentialByEnvironmentCrnV1 gets credential by environment crn

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) GetCredentialByEnvironmentCrnV1(params *GetCredentialByEnvironmentCrnV1Params) (*GetCredentialByEnvironmentCrnV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialByEnvironmentCrnV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCredentialByEnvironmentCrnV1",
		Method:             "GET",
		PathPattern:        "/v1/credentials/environment/crn/{environmentCrn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCredentialByEnvironmentCrnV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCredentialByEnvironmentCrnV1OK), nil

}

/*
GetCredentialByEnvironmentNameV1 gets credential by environment name

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) GetCredentialByEnvironmentNameV1(params *GetCredentialByEnvironmentNameV1Params) (*GetCredentialByEnvironmentNameV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialByEnvironmentNameV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCredentialByEnvironmentNameV1",
		Method:             "GET",
		PathPattern:        "/v1/credentials/environment/name/{environmentName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCredentialByEnvironmentNameV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCredentialByEnvironmentNameV1OK), nil

}

/*
GetCredentialByNameV1 gets credential by name

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) GetCredentialByNameV1(params *GetCredentialByNameV1Params) (*GetCredentialByNameV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialByNameV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCredentialByNameV1",
		Method:             "GET",
		PathPattern:        "/v1/credentials/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCredentialByNameV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCredentialByNameV1OK), nil

}

/*
GetCredentialByResourceCrnV1 gets credential by crn

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) GetCredentialByResourceCrnV1(params *GetCredentialByResourceCrnV1Params) (*GetCredentialByResourceCrnV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCredentialByResourceCrnV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCredentialByResourceCrnV1",
		Method:             "GET",
		PathPattern:        "/v1/credentials/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCredentialByResourceCrnV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCredentialByResourceCrnV1OK), nil

}

/*
GetPrerequisitesForCloudPlatform gets credential prerequisites for cloud platform

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) GetPrerequisitesForCloudPlatform(params *GetPrerequisitesForCloudPlatformParams) (*GetPrerequisitesForCloudPlatformOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrerequisitesForCloudPlatformParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrerequisitesForCloudPlatform",
		Method:             "GET",
		PathPattern:        "/v1/credentials/prerequisites/{cloudPlatform}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrerequisitesForCloudPlatformReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrerequisitesForCloudPlatformOK), nil

}

/*
InteractiveLoginCredentialV1 interactives login

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) InteractiveLoginCredentialV1(params *InteractiveLoginCredentialV1Params) (*InteractiveLoginCredentialV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInteractiveLoginCredentialV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "interactiveLoginCredentialV1",
		Method:             "POST",
		PathPattern:        "/v1/credentials/interactive_login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InteractiveLoginCredentialV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InteractiveLoginCredentialV1OK), nil

}

/*
ListCredentialsV1 lists credentials

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) ListCredentialsV1(params *ListCredentialsV1Params) (*ListCredentialsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCredentialsV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listCredentialsV1",
		Method:             "GET",
		PathPattern:        "/v1/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListCredentialsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListCredentialsV1OK), nil

}

/*
PutCredentialV1 modifies public credential resource

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) PutCredentialV1(params *PutCredentialV1Params) (*PutCredentialV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCredentialV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putCredentialV1",
		Method:             "PUT",
		PathPattern:        "/v1/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCredentialV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutCredentialV1OK), nil

}

/*
VerifyCredentialByCrn verifies permissions by crn

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) VerifyCredentialByCrn(params *VerifyCredentialByCrnParams) (*VerifyCredentialByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVerifyCredentialByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "verifyCredentialByCrn",
		Method:             "GET",
		PathPattern:        "/v1/credentials/crn/{crn}/verify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VerifyCredentialByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VerifyCredentialByCrnOK), nil

}

/*
VerifyCredentialByName verifies permissions by name

Cloudbreak is launching Hadoop clusters on the user's behalf - on different cloud providers. One key point is that Cloudbreak does not store your Cloud provider account details (such as username, password, keys, private SSL certificates, etc). We work around the concept that Identity and Access Management is fully controlled by you - the end user. The Cloudbreak deployer is purely acting on behalf of the end user - without having access to the user's account.
*/
func (a *Client) VerifyCredentialByName(params *VerifyCredentialByNameParams) (*VerifyCredentialByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVerifyCredentialByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "verifyCredentialByName",
		Method:             "GET",
		PathPattern:        "/v1/credentials/name/{name}/verify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VerifyCredentialByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VerifyCredentialByNameOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
