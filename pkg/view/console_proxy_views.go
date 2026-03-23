// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ConsoleProxyInventoryView ConsoleProxy
type ConsoleProxyInventoryView struct {
	BaseInfoView
	BaseTimeView
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	AgentIp string `json:"agentIp,omitempty"`
	Token string `json:"token,omitempty"`
	AgentType string `json:"agentType,omitempty"`
	ProxyHostname string `json:"proxyHostname,omitempty"`
	ProxyPort int `json:"proxyPort,omitempty"`
	TargetSchema string `json:"targetSchema,omitempty"`
	TargetHostname string `json:"targetHostname,omitempty"`
	TargetPort int `json:"targetPort,omitempty"`
	Scheme string `json:"scheme,omitempty"`
	ProxyIdentity string `json:"proxyIdentity,omitempty"`
	Status string `json:"status,omitempty"`
	Version string `json:"version,omitempty"`
	ExpiredDate time.Time `json:"expiredDate,omitempty"`
}

