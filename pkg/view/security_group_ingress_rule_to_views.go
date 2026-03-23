// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SecurityGroupIngressRuleTOView SecurityGroupIngressRuleTO
type SecurityGroupIngressRuleTOView struct {
	FriendCidrs []string `json:"friendCidrs,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	SecurityGroupUuid string `json:"securityGroupUuid,omitempty"`
	Type string `json:"type,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	State string `json:"state,omitempty"`
	Priority int `json:"priority,omitempty"`
	Description string `json:"description,omitempty"`
	SrcIpRange string `json:"srcIpRange,omitempty"`
	DstIpRange string `json:"dstIpRange,omitempty"`
	SrcPortRange string `json:"srcPortRange,omitempty"`
	DstPortRange string `json:"dstPortRange,omitempty"`
	Action string `json:"action,omitempty"`
	RemoteSecurityGroupUuid string `json:"remoteSecurityGroupUuid,omitempty"`
	AllowedCidr string `json:"allowedCidr,omitempty"`
	StartPort int `json:"startPort,omitempty"`
	EndPort int `json:"endPort,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

