// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateNasFileSystemParamDetail UpdateNasFileSystem detail param
type UpdateNasFileSystemParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateNasFileSystemParam UpdateNasFileSystem request param
type UpdateNasFileSystemParam struct {
	BaseParam
	Params UpdateNasFileSystemParamDetail `json:"updateNasFileSystem"`
}
// DeleteNasFileSystemParamDetail DeleteNasFileSystem detail param
type DeleteNasFileSystemParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteNasFileSystemParam DeleteNasFileSystem request param
type DeleteNasFileSystemParam struct {
	BaseParam
	Params DeleteNasFileSystemParamDetail `json:"deleteNasFileSystem"`
}
