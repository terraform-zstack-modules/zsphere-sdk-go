// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// IscsiServerClusterRefInventoryView IscsiServerClusterRef
type IscsiServerClusterRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	IscsiServerUuid string `json:"iscsiServerUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
}

