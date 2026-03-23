// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateVmNicParamDetail CreateVmNic detail param
type CreateVmNicParamDetail struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	Ip *string `json:"ip,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmNicParam CreateVmNic request param
type CreateVmNicParam struct {
	BaseParam
	Params CreateVmNicParamDetail `json:"params"`
}
// DeleteVmNicParamDetail DeleteVmNic detail param
type DeleteVmNicParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVmNicParam DeleteVmNic request param
type DeleteVmNicParam struct {
	BaseParam
	Params DeleteVmNicParamDetail `json:"deleteVmNic"`
}
