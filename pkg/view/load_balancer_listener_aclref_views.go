// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LoadBalancerListenerACLRefInventoryView LoadBalancerListenerACLRef
type LoadBalancerListenerACLRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	ListenerUuid string `json:"listenerUuid,omitempty"`
	ServerGroupUuid string `json:"serverGroupUuid,omitempty"`
	AclUuid string `json:"aclUuid,omitempty"`
	Type string `json:"type,omitempty"`
}

