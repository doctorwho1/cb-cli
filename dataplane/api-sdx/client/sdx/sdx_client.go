// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sdx API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sdx API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateSdx creates s d x cluster
*/
func (a *Client) CreateSdx(params *CreateSdxParams) (*CreateSdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSdx",
		Method:             "POST",
		PathPattern:        "/sdx/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSdxOK), nil

}

/*
DeleteSdx deletes s d x cluster
*/
func (a *Client) DeleteSdx(params *DeleteSdxParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSdxParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSdx",
		Method:             "DELETE",
		PathPattern:        "/sdx/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteSdxByCrn deletes s d x cluster by crn
*/
func (a *Client) DeleteSdxByCrn(params *DeleteSdxByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSdxByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSdxByCrn",
		Method:             "DELETE",
		PathPattern:        "/sdx/crn/{clusterCrn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetSdx gets s d x cluster
*/
func (a *Client) GetSdx(params *GetSdxParams) (*GetSdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdx",
		Method:             "GET",
		PathPattern:        "/sdx/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxOK), nil

}

/*
GetSdxByCrn gets s d x cluster by crn
*/
func (a *Client) GetSdxByCrn(params *GetSdxByCrnParams) (*GetSdxByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdxByCrn",
		Method:             "GET",
		PathPattern:        "/sdx/crn/{clusterCrn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxByCrnOK), nil

}

/*
GetSdxByEnvCrn gets s d x cluster by environment crn
*/
func (a *Client) GetSdxByEnvCrn(params *GetSdxByEnvCrnParams) (*GetSdxByEnvCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxByEnvCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdxByEnvCrn",
		Method:             "GET",
		PathPattern:        "/sdx/envcrn/{envCrn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxByEnvCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxByEnvCrnOK), nil

}

/*
GetSdxDetail gets s d x cluster detail
*/
func (a *Client) GetSdxDetail(params *GetSdxDetailParams) (*GetSdxDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxDetailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdxDetail",
		Method:             "GET",
		PathPattern:        "/sdx/{name}/detail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxDetailReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxDetailOK), nil

}

/*
GetSdxDetailByCrn gets s d x cluster detail by crn
*/
func (a *Client) GetSdxDetailByCrn(params *GetSdxDetailByCrnParams) (*GetSdxDetailByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxDetailByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdxDetailByCrn",
		Method:             "GET",
		PathPattern:        "/sdx/crn/{clusterCrn}/detail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxDetailByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxDetailByCrnOK), nil

}

/*
ListSdx lists s d x clusters
*/
func (a *Client) ListSdx(params *ListSdxParams) (*ListSdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSdx",
		Method:             "GET",
		PathPattern:        "/sdx/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSdxOK), nil

}

/*
RedeploySdx redeploys s d x cluster
*/
func (a *Client) RedeploySdx(params *RedeploySdxParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRedeploySdxParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "redeploySdx",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/redeploy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RedeploySdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RedeploySdxByCrn redeploys s d x cluster by crn
*/
func (a *Client) RedeploySdxByCrn(params *RedeploySdxByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRedeploySdxByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "redeploySdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{clusterCrn}/redeploy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RedeploySdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RepairSdxNode repairs an sdx node in the specified hostgroup
*/
func (a *Client) RepairSdxNode(params *RepairSdxNodeParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepairSdxNodeParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "repairSdxNode",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/manual_repair",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RepairSdxNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RepairSdxNodeByCrn repairs an sdx node in the specified hostgroup by crn
*/
func (a *Client) RepairSdxNodeByCrn(params *RepairSdxNodeByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepairSdxNodeByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "repairSdxNodeByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/manual_repair",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RepairSdxNodeByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RetrySdx retries sdx
*/
func (a *Client) RetrySdx(params *RetrySdxParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrySdxParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrySdx",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RetrySdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RetrySdxByCrn retries sdx by crn
*/
func (a *Client) RetrySdxByCrn(params *RetrySdxByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrySdxByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrySdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RetrySdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StartSdxByCrn starts sdx by crn
*/
func (a *Client) StartSdxByCrn(params *StartSdxByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartSdxByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startSdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StartSdxByName starts sdx
*/
func (a *Client) StartSdxByName(params *StartSdxByNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartSdxByNameParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startSdxByName",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartSdxByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StopSdxByCrn stops sdx by crn
*/
func (a *Client) StopSdxByCrn(params *StopSdxByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopSdxByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopSdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StopSdxByName stops sdx
*/
func (a *Client) StopSdxByName(params *StopSdxByNameParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopSdxByNameParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopSdxByName",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopSdxByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
SyncSdx syncs s d x cluster by name
*/
func (a *Client) SyncSdx(params *SyncSdxParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncSdxParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "syncSdx",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SyncSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
SyncSdxByCrn syncs s d x cluster by crn
*/
func (a *Client) SyncSdxByCrn(params *SyncSdxByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncSdxByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "syncSdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SyncSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
