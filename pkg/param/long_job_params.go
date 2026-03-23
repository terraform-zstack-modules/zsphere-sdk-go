// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CleanLongJobParamDetail CleanLongJob detail param
type CleanLongJobParamDetail struct {
}

// CleanLongJobParam CleanLongJob request param
type CleanLongJobParam struct {
	BaseParam
	Params CleanLongJobParamDetail `json:"cleanLongJob"`
}
// ResumeLongJobParamDetail ResumeLongJob detail param
type ResumeLongJobParamDetail struct {
}

// ResumeLongJobParam ResumeLongJob request param
type ResumeLongJobParam struct {
	BaseParam
	Params ResumeLongJobParamDetail `json:"resumeLongJob"`
}
// DeleteLongJobParamDetail DeleteLongJob detail param
type DeleteLongJobParamDetail struct {
}

// DeleteLongJobParam DeleteLongJob request param
type DeleteLongJobParam struct {
	BaseParam
	Params DeleteLongJobParamDetail `json:"deleteLongJob"`
}
// UpdateLongJobParamDetail UpdateLongJob detail param
type UpdateLongJobParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateLongJobParam UpdateLongJob request param
type UpdateLongJobParam struct {
	BaseParam
	Params UpdateLongJobParamDetail `json:"updateLongJob"`
}
