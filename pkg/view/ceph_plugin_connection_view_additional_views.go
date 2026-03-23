// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CephPluginConnectionViewView CephPluginConnectionView
type CephPluginConnectionViewView struct {
	Ip string `json:"ip,omitempty"`
	PluginType string `json:"pluginType,omitempty"`
	PluginProperties map[string]interface{} `json:"pluginProperties,omitempty"`
	ManagementNodeUuid string `json:"managementNodeUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
}

