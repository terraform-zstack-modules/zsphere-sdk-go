// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ResourceEnsembleInventoryView ResourceEnsemble
type ResourceEnsembleInventoryView struct {
	BaseInfoView
	BaseTimeView
	MasterUuid string `json:"masterUuid,omitempty"`
	MasterResourceName string `json:"masterResourceName,omitempty"`
	MasterResourceType string `json:"masterResourceType,omitempty"`
	Members []ResourceInventoryView `json:"members,omitempty"`
}

