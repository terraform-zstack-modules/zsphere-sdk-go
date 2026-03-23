// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateImagePackageParamDetail UpdateImagePackage detail param
type UpdateImagePackageParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateImagePackageParam UpdateImagePackage request param
type UpdateImagePackageParam struct {
	BaseParam
	Params UpdateImagePackageParamDetail `json:"updateImagePackage"`
}
// DeleteImagePackageParamDetail DeleteImagePackage detail param
type DeleteImagePackageParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteImagePackageParam DeleteImagePackage request param
type DeleteImagePackageParam struct {
	BaseParam
	Params DeleteImagePackageParamDetail `json:"deleteImagePackage"`
}
