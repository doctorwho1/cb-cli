package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new cluster API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
DeleteStacksIDCluster deletes cluster on a specific stack

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) DeleteStacksIDCluster(params *DeleteStacksIDClusterParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStacksIDClusterParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "DeleteStacksIDCluster",
		Method:             "DELETE",
		PathPattern:        "/stacks/{id}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &DeleteStacksIDClusterReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
GetStacksAccountNameCluster retrieves cluster by stack name public

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetStacksAccountNameCluster(params *GetStacksAccountNameClusterParams) (*GetStacksAccountNameClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStacksAccountNameClusterParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetStacksAccountNameCluster",
		Method:             "GET",
		PathPattern:        "/stacks/account/{name}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &GetStacksAccountNameClusterReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStacksAccountNameClusterOK), nil
}

/*
GetStacksIDCluster retrieves cluster by stack id

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetStacksIDCluster(params *GetStacksIDClusterParams) (*GetStacksIDClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStacksIDClusterParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetStacksIDCluster",
		Method:             "GET",
		PathPattern:        "/stacks/{id}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &GetStacksIDClusterReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStacksIDClusterOK), nil
}

/*
GetStacksUserNameCluster retrieves cluster by stack name private

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetStacksUserNameCluster(params *GetStacksUserNameClusterParams) (*GetStacksUserNameClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStacksUserNameClusterParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "GetStacksUserNameCluster",
		Method:             "GET",
		PathPattern:        "/stacks/user/{name}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &GetStacksUserNameClusterReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStacksUserNameClusterOK), nil
}

/*
PostStacksIDCluster creates cluster for stack

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) PostStacksIDCluster(params *PostStacksIDClusterParams) (*PostStacksIDClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStacksIDClusterParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "PostStacksIDCluster",
		Method:             "POST",
		PathPattern:        "/stacks/{id}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &PostStacksIDClusterReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostStacksIDClusterOK), nil
}

/*
PutStacksIDCluster updates cluster by stack id

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) PutStacksIDCluster(params *PutStacksIDClusterParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutStacksIDClusterParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "PutStacksIDCluster",
		Method:             "PUT",
		PathPattern:        "/stacks/{id}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &PutStacksIDClusterReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
GetConfigs gets cluster properties with blueprint outputs

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetConfigs(params *GetConfigsParams) (*GetConfigsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getConfigs",
		Method:             "POST",
		PathPattern:        "/stacks/{id}/cluster/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"", "http", "https"},
		Params:             params,
		Reader:             &GetConfigsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConfigsOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}
