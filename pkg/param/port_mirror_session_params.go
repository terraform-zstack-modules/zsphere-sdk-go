// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreatePortMirrorSessionParamDetail CreatePortMirrorSession detail param
type CreatePortMirrorSessionParamDetail struct {
	PortMirrorUuid string `json:"portMirrorUuid" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
	SrcEndPoint string `json:"srcEndPoint" validate:"required"`
	DstEndPoint string `json:"dstEndPoint" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePortMirrorSessionParam CreatePortMirrorSession request param
type CreatePortMirrorSessionParam struct {
	BaseParam
	Params CreatePortMirrorSessionParamDetail `json:"params"`
}
// DeletePortMirrorSessionParamDetail DeletePortMirrorSession detail param
type DeletePortMirrorSessionParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeletePortMirrorSessionParam DeletePortMirrorSession request param
type DeletePortMirrorSessionParam struct {
	BaseParam
	Params DeletePortMirrorSessionParamDetail `json:"deletePortMirrorSession"`
}
