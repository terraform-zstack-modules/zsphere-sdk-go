// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// RemoveSNSDingTalkAtPersonParamDetail RemoveSNSDingTalkAtPerson detail param
type RemoveSNSDingTalkAtPersonParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveSNSDingTalkAtPersonParam RemoveSNSDingTalkAtPerson request param
type RemoveSNSDingTalkAtPersonParam struct {
	BaseParam
	Params RemoveSNSDingTalkAtPersonParamDetail `json:"removeSNSDingTalkAtPerson"`
}
// AddSNSDingTalkAtPersonParamDetail AddSNSDingTalkAtPerson detail param
type AddSNSDingTalkAtPersonParamDetail struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	EndpointUuid string `json:"endpointUuid" validate:"required"`
	Remark *string `json:"remark,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddSNSDingTalkAtPersonParam AddSNSDingTalkAtPerson request param
type AddSNSDingTalkAtPersonParam struct {
	BaseParam
	Params AddSNSDingTalkAtPersonParamDetail `json:"params"`
}
