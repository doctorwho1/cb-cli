package client_autoscale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/client_autoscale/alerts"
	"github.com/hortonworks/hdc-cli/client_autoscale/clusters"
	"github.com/hortonworks/hdc-cli/client_autoscale/configurations"
	"github.com/hortonworks/hdc-cli/client_autoscale/history"
	"github.com/hortonworks/hdc-cli/client_autoscale/policies"
)

// Default auto scaling HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new auto scaling HTTP client.
func NewHTTPClient(formats strfmt.Registry) *AutoScaling {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost", "/", []string{"http", "https"})
	return New(transport, formats)
}

// New creates a new auto scaling client
func New(transport client.Transport, formats strfmt.Registry) *AutoScaling {
	cli := new(AutoScaling)
	cli.Transport = transport

	cli.Alerts = alerts.New(transport, formats)

	cli.Clusters = clusters.New(transport, formats)

	cli.Configurations = configurations.New(transport, formats)

	cli.History = history.New(transport, formats)

	cli.Policies = policies.New(transport, formats)

	return cli
}

// AutoScaling is a client for auto scaling
type AutoScaling struct {
	Alerts *alerts.Client

	Clusters *clusters.Client

	Configurations *configurations.Client

	History *history.Client

	Policies *policies.Client

	Transport client.Transport
}

// SetTransport changes the transport on the client and all its subresources
func (c *AutoScaling) SetTransport(transport client.Transport) {
	c.Transport = transport

	c.Alerts.SetTransport(transport)

	c.Clusters.SetTransport(transport)

	c.Configurations.SetTransport(transport)

	c.History.SetTransport(transport)

	c.Policies.SetTransport(transport)

}