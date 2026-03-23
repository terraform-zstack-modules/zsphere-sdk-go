// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ImageStoreImageStructView ImageStoreImageStruct
type ImageStoreImageStructView struct {
	Id string `json:"id,omitempty"`
	Parent string `json:"parent,omitempty"`
	Blobsum string `json:"blobsum,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Author string `json:"author,omitempty"`
	Arch string `json:"arch,omitempty"`
	Desc string `json:"desc,omitempty"`
	Size int64 `json:"size,omitempty"`
	Virtualsize int64 `json:"virtualsize,omitempty"`
	Name string `json:"name,omitempty"`
}

