// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteClusterParamDetail DeleteCluster detail param
type DeleteClusterParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteClusterParam DeleteCluster request param
type DeleteClusterParam struct {
	BaseParam
	Params DeleteClusterParamDetail `json:"deleteCluster"`
}
// UpdateClusterParamDetail UpdateCluster detail param
type UpdateClusterParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateClusterParam UpdateCluster request param
type UpdateClusterParam struct {
	BaseParam
	Params UpdateClusterParamDetail `json:"updateCluster"`
}
// CreateClusterParamDetail CreateCluster detail param
type CreateClusterParamDetail struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	HypervisorType string `json:"hypervisorType" validate:"required"`
	Type *string `json:"type,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateClusterParam CreateCluster request param
type CreateClusterParam struct {
	BaseParam
	Params CreateClusterParamDetail `json:"params"`
}
