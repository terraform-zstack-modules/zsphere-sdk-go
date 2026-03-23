// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// RemoveSNSWeComAtPersonParamDetail RemoveSNSWeComAtPerson detail param
type RemoveSNSWeComAtPersonParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveSNSWeComAtPersonParam RemoveSNSWeComAtPerson request param
type RemoveSNSWeComAtPersonParam struct {
	BaseParam
	Params RemoveSNSWeComAtPersonParamDetail `json:"removeSNSWeComAtPerson"`
}
// AddSNSWeComAtPersonParamDetail AddSNSWeComAtPerson detail param
type AddSNSWeComAtPersonParamDetail struct {
	UserId string `json:"userId" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	Remark *string `json:"remark,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSNSWeComAtPersonParam AddSNSWeComAtPerson request param
type AddSNSWeComAtPersonParam struct {
	BaseParam
	Params AddSNSWeComAtPersonParamDetail `json:"params"`
}
