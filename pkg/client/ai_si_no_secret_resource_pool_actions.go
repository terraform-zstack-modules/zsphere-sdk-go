// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"fmt"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// CreateAiSiNoSecretResourcePool creates AiSiNoSecretResourcePool
func (cli *ZSClient) CreateAiSiNoSecretResourcePool(ctx context.Context, params param.CreateAiSiNoSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/secret-resource-pool/aisino"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
