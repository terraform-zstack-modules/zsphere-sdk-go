// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PortMirrorSessionMirrorNetworkRefInventoryView PortMirrorSessionMirrorNetworkRef
type PortMirrorSessionMirrorNetworkRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	SrcTunnelUuid string `json:"srcTunnelUuid,omitempty"`
	DstTunnelUuid string `json:"dstTunnelUuid,omitempty"`
	Type string `json:"type,omitempty"`
	SessionUuid string `json:"sessionUuid,omitempty"`
	SrcTunnel MirrorNetworkUsedIpInventoryView `json:"srcTunnel,omitempty"`
	DstTunnel MirrorNetworkUsedIpInventoryView `json:"dstTunnel,omitempty"`
}

