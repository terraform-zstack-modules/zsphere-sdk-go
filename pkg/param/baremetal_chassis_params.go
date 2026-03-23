// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateBaremetalChassisParamDetail UpdateBaremetalChassis detail param
type UpdateBaremetalChassisParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IpmiAddress *string `json:"ipmiAddress,omitempty"`
	IpmiPort *int `json:"ipmiPort,omitempty"`
	IpmiUsername *string `json:"ipmiUsername,omitempty"`
	IpmiPassword *string `json:"ipmiPassword,omitempty"`
}

// UpdateBaremetalChassisParam UpdateBaremetalChassis request param
type UpdateBaremetalChassisParam struct {
	BaseParam
	Params UpdateBaremetalChassisParamDetail `json:"updateBaremetalChassis"`
}
// DeleteBaremetalChassisParamDetail DeleteBaremetalChassis detail param
type DeleteBaremetalChassisParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteBaremetalChassisParam DeleteBaremetalChassis request param
type DeleteBaremetalChassisParam struct {
	BaseParam
	Params DeleteBaremetalChassisParamDetail `json:"deleteBaremetalChassis"`
}
// CreateBaremetalChassisParamDetail CreateBaremetalChassis detail param
type CreateBaremetalChassisParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	IpmiAddress string `json:"ipmiAddress" validate:"required"`
	IpmiPort *int `json:"ipmiPort,omitempty"`
	IpmiUsername string `json:"ipmiUsername" validate:"required"`
	IpmiPassword string `json:"ipmiPassword" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateBaremetalChassisParam CreateBaremetalChassis request param
type CreateBaremetalChassisParam struct {
	BaseParam
	Params CreateBaremetalChassisParamDetail `json:"params"`
}
// InspectBaremetalChassisParamDetail InspectBaremetalChassis detail param
type InspectBaremetalChassisParamDetail struct {
}

// InspectBaremetalChassisParam InspectBaremetalChassis request param
type InspectBaremetalChassisParam struct {
	BaseParam
	Params InspectBaremetalChassisParamDetail `json:"inspectBaremetalChassis"`
}
