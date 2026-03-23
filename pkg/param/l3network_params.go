// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateL3NetworkParamDetail CreateL3Network detail param
type CreateL3NetworkParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	L2NetworkUuid string `json:"l2NetworkUuid" validate:"required"`
	Category *string `json:"category,omitempty"`
	IpVersion *int `json:"ipVersion,omitempty"`
	System bool `json:"system,omitempty"`
	DnsDomain *string `json:"dnsDomain,omitempty"`
	EnableIPAM *bool `json:"enableIPAM,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateL3NetworkParam CreateL3Network request param
type CreateL3NetworkParam struct {
	BaseParam
	Params CreateL3NetworkParamDetail `json:"params"`
}
// DeleteL3NetworkParamDetail DeleteL3Network detail param
type DeleteL3NetworkParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteL3NetworkParam DeleteL3Network request param
type DeleteL3NetworkParam struct {
	BaseParam
	Params DeleteL3NetworkParamDetail `json:"deleteL3Network"`
}
// UpdateL3NetworkParamDetail UpdateL3Network detail param
type UpdateL3NetworkParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	DnsDomain *string `json:"dnsDomain,omitempty"`
	Category *string `json:"category,omitempty"`
	System *bool `json:"system,omitempty"`
}

// UpdateL3NetworkParam UpdateL3Network request param
type UpdateL3NetworkParam struct {
	BaseParam
	Params UpdateL3NetworkParamDetail `json:"updateL3Network"`
}
