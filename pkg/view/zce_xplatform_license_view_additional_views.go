// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZceXPlatformLicenseViewView ZceXPlatformLicenseView
type ZceXPlatformLicenseViewView struct {
	IssuedTime string `json:"issuedTime,omitempty"`
	ExpiredTime string `json:"expiredTime,omitempty"`
	FsId string `json:"fsId,omitempty"`
	RelatedClusterId string `json:"relatedClusterId,omitempty"`
	Expired bool `json:"expired,omitempty"`
}

