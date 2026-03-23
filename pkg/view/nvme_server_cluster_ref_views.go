// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NvmeServerClusterRefInventoryView NvmeServerClusterRef
type NvmeServerClusterRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	NvmeServerUuid string `json:"nvmeServerUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
}

