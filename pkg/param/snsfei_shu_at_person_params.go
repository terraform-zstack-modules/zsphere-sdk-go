// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// RemoveSNSFeiShuAtPersonParamDetail RemoveSNSFeiShuAtPerson detail param
type RemoveSNSFeiShuAtPersonParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveSNSFeiShuAtPersonParam RemoveSNSFeiShuAtPerson request param
type RemoveSNSFeiShuAtPersonParam struct {
	BaseParam
	Params RemoveSNSFeiShuAtPersonParamDetail `json:"removeSNSFeiShuAtPerson"`
}
// AddSNSFeiShuAtPersonParamDetail AddSNSFeiShuAtPerson detail param
type AddSNSFeiShuAtPersonParamDetail struct {
	UserId string `json:"userId" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	Remark *string `json:"remark,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSNSFeiShuAtPersonParam AddSNSFeiShuAtPerson request param
type AddSNSFeiShuAtPersonParam struct {
	BaseParam
	Params AddSNSFeiShuAtPersonParamDetail `json:"params"`
}
