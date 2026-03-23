// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CleanSoftwarePackageParamDetail CleanSoftwarePackage detail param
type CleanSoftwarePackageParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// CleanSoftwarePackageParam CleanSoftwarePackage request param
type CleanSoftwarePackageParam struct {
	BaseParam
	Params CleanSoftwarePackageParamDetail `json:"cleanSoftwarePackage"`
}
