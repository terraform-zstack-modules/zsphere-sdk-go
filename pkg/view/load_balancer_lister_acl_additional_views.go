// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LoadBalancerListerAclView LoadBalancerListerAcl
type LoadBalancerListerAclView struct {
	AclUuid string `json:"aclUuid,omitempty"`
	ListenerUuid string `json:"listenerUuid,omitempty"`
	ServerGroupUuids []string `json:"serverGroupUuids,omitempty"`
}

