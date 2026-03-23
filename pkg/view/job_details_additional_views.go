// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// JobDetailsView JobDetails
type JobDetailsView struct {
	LongJobUuid string `json:"longJobUuid,omitempty"`
	LongJobState string `json:"longJobState,omitempty"`
	SoftwarePackageUuid string `json:"softwarePackageUuid,omitempty"`
	SoftwarePackageUploadUrl string `json:"softwarePackageUploadUrl,omitempty"`
	Offset int64 `json:"offset,omitempty"`
}

