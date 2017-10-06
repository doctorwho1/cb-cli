package gcp

import "github.com/hortonworks/hdc-cli/cli/cloud"

func (p *GcpProvider) GetNetworkParamatersTemplate(mode cloud.NetworkMode) map[string]interface{} {
	switch mode {
	case cloud.EXISTING_NETWORK_NEW_SUBNET:
		return map[string]interface{}{"networkId": "____"}
	case cloud.EXISTING_NETWORK_EXISTING_SUBNET:
		return map[string]interface{}{"networkId": "____", "subnetId": "____", "noPublicIp": false, "noFirewallRules": false}
	case cloud.LEGACY_NETWORK:
		return map[string]interface{}{"networkId": ""}
	default:
		return nil
	}
}
