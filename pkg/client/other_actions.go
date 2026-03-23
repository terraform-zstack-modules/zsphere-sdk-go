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

// CreateAutoScalingGroupAddingNewInstanceRule creates AutoScalingGroupAddingNewInstanceRule
func (cli *ZSClient) CreateAutoScalingGroupAddingNewInstanceRule(ctx context.Context, params param.CreateAutoScalingGroupAddingNewInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/autoscaling/rules/adding-new-instance"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL3NetworksForChangeVmNicNetwork gets CandidateL3NetworksForChangeVmNicNetwork by uuid
func (cli *ZSClient) GetCandidateL3NetworksForChangeVmNicNetwork(ctx context.Context, uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/nics", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetServiceTypeOnHostNetworkBonding operates on ServiceTypeOnHostNetworkBonding
func (cli *ZSClient) SetServiceTypeOnHostNetworkBonding(ctx context.Context, params param.SetServiceTypeOnHostNetworkBondingParam) (*view.HostNetworkBondingServiceRefInventoryView, error) {
	resp := view.HostNetworkBondingServiceRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hosts/bondings/service-types"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidatePrimaryStoragesForCreatingVm gets CandidatePrimaryStoragesForCreatingVm by uuid
func (cli *ZSClient) GetCandidatePrimaryStoragesForCreatingVm(ctx context.Context) (*view.GetCandidatePrimaryStoragesForCreatingVmView, error) {
	var resp view.GetCandidatePrimaryStoragesForCreatingVmView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/candidate-storages", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmConsolePassword gets VmConsolePassword by uuid
func (cli *ZSClient) GetVmConsolePassword(ctx context.Context, uuid string) (*view.GetVmConsolePasswordView, error) {
	var resp view.GetVmConsolePasswordView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceBindableConfig gets ResourceBindableConfig by uuid
func (cli *ZSClient) GetResourceBindableConfig(ctx context.Context) (*view.GetResourceBindableConfigView, error) {
	var resp view.GetResourceBindableConfigView
	if err := cli.GetWithRespKey(ctx, "v1/resource-configurations/bindable", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryPciDevicePciDeviceOffering queries PciDevicePciDeviceOffering list
func (cli *ZSClient) QueryPciDevicePciDeviceOffering(ctx context.Context, params *param.QueryParam) ([]view.PciDevicePciDeviceOfferingRefInventoryView, error) {
	var resp []view.PciDevicePciDeviceOfferingRefInventoryView
	return resp, cli.List(ctx, "v1/pci-devices/pci-devices/pci-device-offerings", params, &resp)
}

// PagePciDevicePciDeviceOffering Pagination
func (cli *ZSClient) PagePciDevicePciDeviceOffering(ctx context.Context, params *param.QueryParam) ([]view.PciDevicePciDeviceOfferingRefInventoryView, int, error) {
	var pciDevicePciDeviceOfferings []view.PciDevicePciDeviceOfferingRefInventoryView
	total, err := cli.Page(ctx, "v1/pci-devices/pci-devices/pci-device-offerings", params, &pciDevicePciDeviceOfferings)
	return pciDevicePciDeviceOfferings, total, err
}

// GetVmInstanceHaLevel gets VmInstanceHaLevel by uuid
func (cli *ZSClient) GetVmInstanceHaLevel(ctx context.Context, uuid string) (*view.GetVmInstanceHaLevelView, error) {
	var resp view.GetVmInstanceHaLevelView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAccessControlListToLoadBalancer adds AccessControlListToLoadBalancer
func (cli *ZSClient) AddAccessControlListToLoadBalancer(ctx context.Context, listenerUuid string, params param.AddAccessControlListToLoadBalancerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/load-balancers/listeners/%s/access-control-lists", listenerUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LogOut operates on LogOut
func (cli *ZSClient) LogOut(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/accounts/sessions", uuid, string(deleteMode))
}

// GetVmXmlHookScript gets VmXmlHookScript by uuid
func (cli *ZSClient) GetVmXmlHookScript(ctx context.Context, uuid string) (*view.GetVmXmlHookScriptView, error) {
	var resp view.GetVmXmlHookScriptView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAudit queries Audit list
func (cli *ZSClient) QueryAudit(ctx context.Context, params *param.QueryParam) ([]view.AuditsInventoryView, error) {
	var resp []view.AuditsInventoryView
	return resp, cli.List(ctx, "v1/zwatch/audit-records", params, &resp)
}

// PageAudit Pagination
func (cli *ZSClient) PageAudit(ctx context.Context, params *param.QueryParam) ([]view.AuditsInventoryView, int, error) {
	var audits []view.AuditsInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/audit-records", params, &audits)
	return audits, total, err
}

// RemoveAccountFromGroup removes AccountFromGroup
func (cli *ZSClient) RemoveAccountFromGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/account-groups", uuid, string(deleteMode))
}

// RemoveResourcesFromDirectory removes ResourcesFromDirectory
func (cli *ZSClient) RemoveResourcesFromDirectory(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/remove/resources/directory", uuid, string(deleteMode))
}

// CreateVmFromVmBackup creates VmFromVmBackup
func (cli *ZSClient) CreateVmFromVmBackup(ctx context.Context, groupUuid string, params param.CreateVmFromVmBackupParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/from/vm-backups/%s", groupUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImageQga gets ImageQga by uuid
func (cli *ZSClient) GetImageQga(ctx context.Context, uuid string) (*view.GetImageQgaView, error) {
	var resp view.GetImageQgaView
	if err := cli.GetWithRespKey(ctx, "v1/images", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInterdependentL3NetworksBackupStorages gets InterdependentL3NetworksBackupStorages by uuid
func (cli *ZSClient) GetInterdependentL3NetworksBackupStorages(ctx context.Context) (*view.GetInterdependentL3NetworksBackupStoragesView, error) {
	var resp view.GetInterdependentL3NetworksBackupStoragesView
	if err := cli.GetWithRespKey(ctx, "v1/backupStorage-l3networks/dependencies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteExportedDatabaseBackupFromBackupStorage deletes ExportedDatabaseBackupFromBackupStorage
func (cli *ZSClient) DeleteExportedDatabaseBackupFromBackupStorage(ctx context.Context, databaseBackupUuid string, backupStorageUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/exported-database-backup", databaseBackupUuid, fmt.Sprintf("backup-storage/%s", backupStorageUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// UnexportNbdVolumes operates on UnexportNbdVolumes
func (cli *ZSClient) UnexportNbdVolumes(ctx context.Context, params param.UnexportNbdVolumesParam) (*view.UnexportNbdVolumesEventView, error) {
	resp := view.UnexportNbdVolumesEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cbt-task/unexportvolume"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachNetworkServiceToL3Network operates on NetworkServiceToL3Network
func (cli *ZSClient) AttachNetworkServiceToL3Network(ctx context.Context, l3NetworkUuid string, params param.AttachNetworkServiceToL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks/%s/network-services", l3NetworkUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVpcRouter queries VpcRouter list
func (cli *ZSClient) QueryVpcRouter(ctx context.Context, params *param.QueryParam) ([]view.VpcRouterVmInventoryView, error) {
	var resp []view.VpcRouterVmInventoryView
	return resp, cli.List(ctx, "v1/vpc/virtual-routers", params, &resp)
}

func (cli *ZSClient) GetVpcRouter(ctx context.Context, uuid string) (*view.VpcRouterVmInventoryView, error) {
	var resp view.VpcRouterVmInventoryView
	if err := cli.Get(ctx, "v1/vpc/virtual-routers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcRouter Pagination
func (cli *ZSClient) PageVpcRouter(ctx context.Context, params *param.QueryParam) ([]view.VpcRouterVmInventoryView, int, error) {
	var vpcRouters []view.VpcRouterVmInventoryView
	total, err := cli.Page(ctx, "v1/vpc/virtual-routers", params, &vpcRouters)
	return vpcRouters, total, err
}

// SetVmClockTrack operates on VmClockTrack
func (cli *ZSClient) SetVmClockTrack(ctx context.Context, uuid string, params param.SetVmClockTrackParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "inventory", map[string]interface{}{
		"setVmClockTrack": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateEmailMonitorTriggerAction updates EmailMonitorTrigger
func (cli *ZSClient) UpdateEmailMonitorTriggerAction(ctx context.Context, uuid string, params param.UpdateEmailMonitorTriggerActionParam) (*view.MonitorTriggerActionInventoryView, error) {
	resp := view.MonitorTriggerActionInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/monitoring/trigger-actions/emails", uuid, "", "inventory", map[string]interface{}{
		"updateEmailMonitorTriggerAction": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryLocalRaidController queries LocalRaidController list
func (cli *ZSClient) QueryLocalRaidController(ctx context.Context, params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, error) {
	var resp []view.RaidPhysicalDriveInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/local-raid/controllers", params, &resp)
}

func (cli *ZSClient) GetLocalRaidController(ctx context.Context, uuid string) (*view.RaidPhysicalDriveInventoryView, error) {
	var resp view.RaidPhysicalDriveInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/local-raid/controllers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLocalRaidController Pagination
func (cli *ZSClient) PageLocalRaidController(ctx context.Context, params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, int, error) {
	var localRaidControllers []view.RaidPhysicalDriveInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/local-raid/controllers", params, &localRaidControllers)
	return localRaidControllers, total, err
}

// SNSHttpTestConnection operates on HttpTestConnection
func (cli *ZSClient) SNSHttpTestConnection(ctx context.Context, params param.SNSHttpTestConnectionParam) (*view.SNSHttpTestConnectionEventView, error) {
	resp := view.SNSHttpTestConnectionEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/http/test-connection"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExecuteAutoScalingRule operates on ExecuteAutoScalingRule
func (cli *ZSClient) ExecuteAutoScalingRule(ctx context.Context, uuid string, params param.ExecuteAutoScalingRuleParam) (*view.ExecuteAutoScalingRuleEventView, error) {
	resp := view.ExecuteAutoScalingRuleEventView{}
	if err := cli.PutWithSpec(ctx, "v1/autoscaling/rules", uuid, "actions", "", map[string]interface{}{
		"executeAutoScalingRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageSecurityLevel operates on ImageSecurityLevel
func (cli *ZSClient) SetImageSecurityLevel(ctx context.Context, uuid string, params param.SetImageSecurityLevelParam) (*view.SetImageSecurityLevelEventView, error) {
	resp := view.SetImageSecurityLevelEventView{}
	if err := cli.PutWithSpec(ctx, "v1/images", uuid, "actions", "", map[string]interface{}{
		"setImageSecurityLevel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBackupStorageState changes BackupStorageState
func (cli *ZSClient) ChangeBackupStorageState(ctx context.Context, uuid string, params param.ChangeBackupStorageStateParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage", uuid, "actions", "inventory", map[string]interface{}{
		"changeBackupStorageState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachVmFromVmSchedulingRuleGroup operates on VmFromVmSchedulingRuleGroup
func (cli *ZSClient) DetachVmFromVmSchedulingRuleGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vmSchedulingRuleGroup", uuid, string(deleteMode))
}

// GetCandidateIsoForAttachingVm gets CandidateIsoForAttachingVm by uuid
func (cli *ZSClient) GetCandidateIsoForAttachingVm(ctx context.Context, uuid string) (*view.ImageInventoryView, error) {
	var resp view.ImageInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SecurityMachineDetectSync operates on MachineDetectSync
func (cli *ZSClient) SecurityMachineDetectSync(ctx context.Context, uuid string, params param.SecurityMachineDetectSyncParam) (*view.SecurityMachineDetectSyncEventView, error) {
	resp := view.SecurityMachineDetectSyncEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/security-machine/%s/detect/sync/actions", uuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSecurityGroupState changes SecurityGroupState
func (cli *ZSClient) ChangeSecurityGroupState(ctx context.Context, uuid string, params param.ChangeSecurityGroupStateParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/security-groups", uuid, "actions", "inventory", map[string]interface{}{
		"changeSecurityGroupState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVRouterNetworksToOspfArea adds VRouterNetworksToOspfArea
func (cli *ZSClient) AddVRouterNetworksToOspfArea(ctx context.Context, routerAreaUuid, vRouterUuid string, params param.AddVRouterNetworksToOspfAreaParam) (*view.NetworkRouterAreaRefInventoryView, error) {
	resp := view.NetworkRouterAreaRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/routerArea/%s/router/%s/addnetworks", routerAreaUuid, vRouterUuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrometheusMetricLabelValue gets PrometheusMetricLabelValue by uuid
func (cli *ZSClient) GetPrometheusMetricLabelValue(ctx context.Context) (*view.GetPrometheusMetricLabelValueView, error) {
	var resp view.GetPrometheusMetricLabelValueView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/metrics/prometheus/label-values", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAlarmData updates AlarmData
func (cli *ZSClient) UpdateAlarmData(ctx context.Context, params param.UpdateAlarmDataParam) (*view.UpdateAlarmDataEventView, error) {
	resp := view.UpdateAlarmDataEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/alarm-histories/actions", "", "", map[string]interface{}{
		"updateAlarmData": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchCreateHostKernelInterface operates on CreateHostKernelInterface
func (cli *ZSClient) BatchCreateHostKernelInterface(ctx context.Context, l3NetworkUuid string, params param.BatchCreateHostKernelInterfaceParam) (*view.BatchCreateHostKernelInterfaceEventView, error) {
	resp := view.BatchCreateHostKernelInterfaceEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks/%s/kernel-interfaces", l3NetworkUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNSEmailTestConnection operates on EmailTestConnection
func (cli *ZSClient) SNSEmailTestConnection(ctx context.Context, params param.SNSEmailTestConnectionParam) (*view.SNSEmailTestConnectionEventView, error) {
	resp := view.SNSEmailTestConnectionEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/email/test-connection"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAutoScalingGroupState changes AutoScalingGroupState
func (cli *ZSClient) ChangeAutoScalingGroupState(ctx context.Context, uuid string, params param.ChangeAutoScalingGroupStateParam) (*view.AutoScalingGroupInventoryView, error) {
	resp := view.AutoScalingGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/autoscaling/groups", uuid, "actions", "inventory", map[string]interface{}{
		"changeAutoScalingGroupState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAutoScalingGroupRemovalInstanceRule creates AutoScalingGroupRemovalInstanceRule
func (cli *ZSClient) CreateAutoScalingGroupRemovalInstanceRule(ctx context.Context, params param.CreateAutoScalingGroupRemovalInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/autoscaling/rules/removal-instance"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeEventSubscriptionState changes EventSubscriptionState
func (cli *ZSClient) ChangeEventSubscriptionState(ctx context.Context, uuid string, params param.ChangeEventSubscriptionStateParam) (*view.EventSubscriptionInventoryView, error) {
	resp := view.EventSubscriptionInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/change/eventSubscription", uuid, "state", "inventory", map[string]interface{}{
		"changeEventSubscriptionState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachL3NetworkToVm operates on L3NetworkToVm
func (cli *ZSClient) AttachL3NetworkToVm(ctx context.Context, vmInstanceUuid, l3NetworkUuid string, params param.AttachL3NetworkToVmParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/%s/l3-networks/%s", vmInstanceUuid, l3NetworkUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPrimaryStorageToCluster operates on PrimaryStorageToCluster
func (cli *ZSClient) AttachPrimaryStorageToCluster(ctx context.Context, clusterUuid, primaryStorageUuid string, params param.AttachPrimaryStorageToClusterParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/clusters/%s/primary-storage/%s", clusterUuid, primaryStorageUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachL2NetworkToCluster operates on L2NetworkToCluster
func (cli *ZSClient) AttachL2NetworkToCluster(ctx context.Context, l2NetworkUuid, clusterUuid string, params param.AttachL2NetworkToClusterParam) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l2-networks/%s/clusters/%s", l2NetworkUuid, clusterUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmNicType changes VmNicType
func (cli *ZSClient) ChangeVmNicType(ctx context.Context, vmNicUuid string, params param.ChangeVmNicTypeParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances/nics", vmNicUuid, "actions", "inventory", map[string]interface{}{
		"changeVmNicType": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeFirewallRuleState changes FirewallRuleState
func (cli *ZSClient) ChangeFirewallRuleState(ctx context.Context, uuid string, params param.ChangeFirewallRuleStateParam) (*view.VpcFirewallRuleInventoryView, error) {
	resp := view.VpcFirewallRuleInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vpcfirewalls/rules", uuid, "actions", "inventory", map[string]interface{}{
		"changeFirewallRuleState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMdevDeviceCandidates gets MdevDeviceCandidates by uuid
func (cli *ZSClient) GetMdevDeviceCandidates(ctx context.Context) (*view.MdevDeviceInventoryView, error) {
	var resp view.MdevDeviceInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/mdev-devices/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachCCSCertificateFromAccount operates on CCSCertificateFromAccount
func (cli *ZSClient) DetachCCSCertificateFromAccount(ctx context.Context, accountUuid string, params param.DetachCCSCertificateFromAccountParam) (*view.DetachCCSCertificateFromAccountEventView, error) {
	resp := view.DetachCCSCertificateFromAccountEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/crypto/ccs-certificate/detach-account/%s", accountUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddHostRouteToL3Network adds HostRouteToL3Network
func (cli *ZSClient) AddHostRouteToL3Network(ctx context.Context, l3NetworkUuid string, params param.AddHostRouteToL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks/%s/hostroute", l3NetworkUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddInstanceToMonitorGroup adds InstanceToMonitorGroup
func (cli *ZSClient) AddInstanceToMonitorGroup(ctx context.Context, groupUuid string, params param.AddInstanceToMonitorGroupParam) (*view.MonitorGroupInstanceInventoryView, error) {
	resp := view.MonitorGroupInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/monitorgroups/%s/actions", groupUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTwoFactorAuthenticationState gets TwoFactorAuthenticationState by uuid
func (cli *ZSClient) GetTwoFactorAuthenticationState(ctx context.Context) (*view.GetTwoFactorAuthenticationStateView, error) {
	var resp view.GetTwoFactorAuthenticationStateView
	if err := cli.GetWithRespKey(ctx, "v1/twofactorauthentication/state", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachMdevDeviceToVm operates on MdevDeviceToVm
func (cli *ZSClient) AttachMdevDeviceToVm(ctx context.Context, mdevDeviceUuid, vmInstanceUuid string, params param.AttachMdevDeviceToVmParam) (*view.MdevDeviceInventoryView, error) {
	resp := view.MdevDeviceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/mdev-devices/%s/vm-instances/%s", mdevDeviceUuid, vmInstanceUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BootstrapMiniHost operates on MiniHost
func (cli *ZSClient) BootstrapMiniHost(ctx context.Context, params param.BootstrapMiniHostParam) (*view.BootstrapMiniHostEventView, error) {
	resp := view.BootstrapMiniHostEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/mini-clusters/hosts"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveActionFromAlarm removes ActionFromAlarm
func (cli *ZSClient) RemoveActionFromAlarm(ctx context.Context, alarmUuid string, actionUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/zwatch/alarms", alarmUuid, fmt.Sprintf("actions/%s", actionUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// ChangeEipState changes EipState
func (cli *ZSClient) ChangeEipState(ctx context.Context, uuid string, params param.ChangeEipStateParam) (*view.EipInventoryView, error) {
	resp := view.EipInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/eips", uuid, "actions", "inventory", map[string]interface{}{
		"changeEipState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachSshKeyPairFromVmInstance operates on SshKeyPairFromVmInstance
func (cli *ZSClient) DetachSshKeyPairFromVmInstance(ctx context.Context, sshKeyPairUuid string, vmInstanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/ssh-key-pair", sshKeyPairUuid, fmt.Sprintf("vm-instance/%s", vmInstanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetPrimaryStorageCandidatesForVmMigration gets PrimaryStorageCandidatesForVmMigration by uuid
func (cli *ZSClient) GetPrimaryStorageCandidatesForVmMigration(ctx context.Context, uuid string) (*view.PrimaryStorageInventoryView, error) {
	var resp view.PrimaryStorageInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DecodeStackTemplate operates on DecodeStackTemplate
func (cli *ZSClient) DecodeStackTemplate(ctx context.Context, params param.DecodeStackTemplateParam) (*view.DecodeStackTemplateView, error) {
	resp := view.DecodeStackTemplateView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cloudformation/stack/preview/resource"), "resources", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVirtualRouter updates VirtualRouter
func (cli *ZSClient) UpdateVirtualRouter(ctx context.Context, vmInstanceUuid string, params param.UpdateVirtualRouterParam) (*view.VirtualRouterVmInventoryView, error) {
	resp := view.VirtualRouterVmInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances/appliances/virtual-routers", vmInstanceUuid, "actions", "inventory", map[string]interface{}{
		"updateVirtualRouter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrimaryStorageMigrateVolume operates on PrimaryStorageMigrateVolume
func (cli *ZSClient) PrimaryStorageMigrateVolume(ctx context.Context, volumeUuid string, params param.PrimaryStorageMigrateVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage/volumes", volumeUuid, "actions", "inventory", map[string]interface{}{
		"primaryStorageMigrateVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVSwitchTypes gets VSwitchTypes by uuid
func (cli *ZSClient) GetVSwitchTypes(ctx context.Context) (*view.GetVSwitchTypesView, error) {
	var resp view.GetVSwitchTypesView
	if err := cli.GetWithRespKey(ctx, "v1/l2-networks/vSwitchTypes", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateL2HardwareVxlanNetworkPool creates L2HardwareVxlanNetworkPool
func (cli *ZSClient) CreateL2HardwareVxlanNetworkPool(ctx context.Context, params param.CreateL2HardwareVxlanNetworkPoolParam) (*view.CreateL2HardwareVxlanNetworkPoolEventView, error) {
	resp := view.CreateL2HardwareVxlanNetworkPoolEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l2-networks/hardware-vxlan-pool"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmBootOrder gets VmBootOrder by uuid
func (cli *ZSClient) GetVmBootOrder(ctx context.Context, uuid string) (*view.GetVmBootOrderView, error) {
	var resp view.GetVmBootOrderView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryThirdpartyAlert queries ThirdpartyAlert list
func (cli *ZSClient) QueryThirdpartyAlert(ctx context.Context, params *param.QueryParam) ([]view.ThirdpartyOriginalAlertInventoryView, error) {
	var resp []view.ThirdpartyOriginalAlertInventoryView
	return resp, cli.List(ctx, "v1/zwatch/third-party/alerts", params, &resp)
}

func (cli *ZSClient) GetThirdpartyAlert(ctx context.Context, uuid string) (*view.ThirdpartyOriginalAlertInventoryView, error) {
	var resp view.ThirdpartyOriginalAlertInventoryView
	if err := cli.Get(ctx, "v1/zwatch/third-party/alerts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageThirdpartyAlert Pagination
func (cli *ZSClient) PageThirdpartyAlert(ctx context.Context, params *param.QueryParam) ([]view.ThirdpartyOriginalAlertInventoryView, int, error) {
	var thirdpartyAlerts []view.ThirdpartyOriginalAlertInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/third-party/alerts", params, &thirdpartyAlerts)
	return thirdpartyAlerts, total, err
}

// SetVmBootOrder operates on VmBootOrder
func (cli *ZSClient) SetVmBootOrder(ctx context.Context, uuid string, params param.SetVmBootOrderParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "inventory", map[string]interface{}{
		"setVmBootOrder": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDatabaseBackupFromImageStore gets DatabaseBackupFromImageStore by uuid
func (cli *ZSClient) GetDatabaseBackupFromImageStore(ctx context.Context) (*view.GetDatabaseBackupFromImageStoreView, error) {
	var resp view.GetDatabaseBackupFromImageStoreView
	if err := cli.GetWithRespKey(ctx, "v1/database-backups/image-store", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResizeDataVolume operates on DataVolume
func (cli *ZSClient) ResizeDataVolume(ctx context.Context, uuid string, params param.ResizeDataVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volumes/data/resize", uuid, "actions", "inventory", map[string]interface{}{
		"resizeDataVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEipAttachableVmNics gets EipAttachableVmNics by uuid
func (cli *ZSClient) GetEipAttachableVmNics(ctx context.Context, uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/eips", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIpv6RangeByNetworkCidr adds Ipv6RangeByNetworkCidr
func (cli *ZSClient) AddIpv6RangeByNetworkCidr(ctx context.Context, l3NetworkUuid string, params param.AddIpv6RangeByNetworkCidrParam) (*view.IpRangeInventoryView, error) {
	resp := view.IpRangeInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks/%s/ipv6-ranges/by-cidr", l3NetworkUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceSharing gets ResourceSharing by uuid
func (cli *ZSClient) GetResourceSharing(ctx context.Context) (*view.GetResourceSharingView, error) {
	var resp view.GetResourceSharingView
	if err := cli.GetWithRespKey(ctx, "v1/iam1/resource-ensemble/view-sharing", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LocateLocalRaidPhysicalDrive operates on LocalRaidPhysicalDrive
func (cli *ZSClient) LocateLocalRaidPhysicalDrive(ctx context.Context, uuid string, params param.LocateLocalRaidPhysicalDriveParam) (*view.RaidPhysicalDriveInventoryView, error) {
	resp := view.RaidPhysicalDriveInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/storage-devices/local-raid/physical-drives", uuid, "actions", "", map[string]interface{}{
		"locateLocalRaidPhysicalDrive": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchQuery operates on Query
func (cli *ZSClient) BatchQuery(ctx context.Context) (*view.BatchQueryView, error) {
	var resp view.BatchQueryView
	if err := cli.GetWithRespKey(ctx, "v1/batch-queries", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpBaremetalChassisBonding operates on UpBaremetalChassisBonding
func (cli *ZSClient) CleanUpBaremetalChassisBonding(ctx context.Context, chassisUuid string, params param.CleanUpBaremetalChassisBondingParam) (*view.CleanUpBaremetalChassisBondingEventView, error) {
	resp := view.CleanUpBaremetalChassisBondingEventView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/chassis", chassisUuid, "actions", "", map[string]interface{}{
		"cleanUpBaremetalChassisBonding": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReloadExternalService operates on ReloadExternalService
func (cli *ZSClient) ReloadExternalService(ctx context.Context, params param.ReloadExternalServiceParam) (*view.ReloadExternalServiceEventView, error) {
	resp := view.ReloadExternalServiceEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/external/services", "", "", map[string]interface{}{
		"reloadExternalService": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemovePciDeviceSpecFromVmInstance removes PciDeviceSpecFromVmInstance
func (cli *ZSClient) RemovePciDeviceSpecFromVmInstance(ctx context.Context, pciSpecUuid string, vmInstanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/pci-device-specs", pciSpecUuid, fmt.Sprintf("vm-instances/%s", vmInstanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// DetachGuestToolsIsoFromVm operates on GuestToolsIsoFromVm
func (cli *ZSClient) DetachGuestToolsIsoFromVm(ctx context.Context, uuid string, params param.DetachGuestToolsIsoFromVmParam) (*view.DetachGuestToolsIsoFromVmEventView, error) {
	resp := view.DetachGuestToolsIsoFromVmEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"detachGuestToolsIsoFromVm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveServerGroupFromLoadBalancerListener removes ServerGroupFromLoadBalancerListener
func (cli *ZSClient) RemoveServerGroupFromLoadBalancerListener(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/load-balancers/listeners", uuid, string(deleteMode))
}

// SetVmNicSecurityGroup operates on VmNicSecurityGroup
func (cli *ZSClient) SetVmNicSecurityGroup(ctx context.Context, vmNicUuid string, params param.SetVmNicSecurityGroupParam) (*view.VmNicSecurityGroupRefInventoryView, error) {
	resp := view.VmNicSecurityGroupRefInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/security-groups/nics", vmNicUuid, "actions", "inventory", map[string]interface{}{
		"setVmNicSecurityGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSharedBlockToSharedBlockGroup adds SharedBlockToSharedBlockGroup
func (cli *ZSClient) AddSharedBlockToSharedBlockGroup(ctx context.Context, uuid string, params param.AddSharedBlockToSharedBlockGroupParam) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	resp := view.SharedBlockGroupPrimaryStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/primary-storage/sharedblockgroup/%s/sharedblocks", uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryResourcePrice queries ResourcePrice list
func (cli *ZSClient) QueryResourcePrice(ctx context.Context, params *param.QueryParam) ([]view.PriceInventoryView, error) {
	var resp []view.PriceInventoryView
	return resp, cli.List(ctx, "v1/billings/prices", params, &resp)
}

func (cli *ZSClient) GetResourcePrice(ctx context.Context, uuid string) (*view.PriceInventoryView, error) {
	var resp view.PriceInventoryView
	if err := cli.Get(ctx, "v1/billing/prices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageResourcePrice Pagination
func (cli *ZSClient) PageResourcePrice(ctx context.Context, params *param.QueryParam) ([]view.PriceInventoryView, int, error) {
	var resourcePrices []view.PriceInventoryView
	total, err := cli.Page(ctx, "v1/billings/prices", params, &resourcePrices)
	return resourcePrices, total, err
}

// GetVolumeSnapshotSize gets VolumeSnapshotSize by uuid
func (cli *ZSClient) GetVolumeSnapshotSize(ctx context.Context, uuid string, params param.GetVolumeSnapshotSizeParam) (*view.GetVolumeSnapshotSizeEventView, error) {
	resp := view.GetVolumeSnapshotSizeEventView{}
	if err := cli.PutWithSpec(ctx, "v1/volume-snapshots", uuid, "actions", "", map[string]interface{}{
		"getVolumeSnapshotSize": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshCaptcha operates on Captcha
func (cli *ZSClient) RefreshCaptcha(ctx context.Context) (*view.RefreshCaptchaView, error) {
	var resp view.RefreshCaptchaView
	if err := cli.GetWithRespKey(ctx, "v1/captcha/refresh", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteTag deletes Tag
func (cli *ZSClient) DeleteTag(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/tags", uuid, string(deleteMode))
}

// BatchSyncVolumeSize operates on SyncVolumeSize
func (cli *ZSClient) BatchSyncVolumeSize(ctx context.Context, params param.BatchSyncVolumeSizeParam) (*view.BatchSyncVolumeSizeView, error) {
	resp := view.BatchSyncVolumeSizeView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/volumes/batch-sync-volumes"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportNbdVolumes operates on NbdVolumes
func (cli *ZSClient) ExportNbdVolumes(ctx context.Context, params param.ExportNbdVolumesParam) (*view.ExportNbdVolumesEventView, error) {
	resp := view.ExportNbdVolumesEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cbt-task/exportvolume"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHypervisorTypes gets HypervisorTypes by uuid
func (cli *ZSClient) GetHypervisorTypes(ctx context.Context) (*view.GetHypervisorTypesView, error) {
	var resp view.GetHypervisorTypesView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/hypervisor-types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SelfTestLocalRaid operates on LocalRaid
func (cli *ZSClient) SelfTestLocalRaid(ctx context.Context, uuid string, params param.SelfTestLocalRaidParam) (*view.SelfTestLocalRaidEventView, error) {
	resp := view.SelfTestLocalRaidEventView{}
	if err := cli.PutWithSpec(ctx, "v1/storage-devices/local-raid/physical-drives", uuid, "actions", "", map[string]interface{}{
		"selfTestLocalRaid": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmAttachableDataVolume gets VmAttachableDataVolume by uuid
func (cli *ZSClient) GetVmAttachableDataVolume(ctx context.Context, uuid string) (*view.VolumeInventoryView, error) {
	var resp view.VolumeInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmMonitorNumber gets VmMonitorNumber by uuid
func (cli *ZSClient) GetVmMonitorNumber(ctx context.Context, uuid string) (*view.GetVmMonitorNumberView, error) {
	var resp view.GetVmMonitorNumberView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSNSApplicationPlatformState changes SNSApplicationPlatformState
func (cli *ZSClient) ChangeSNSApplicationPlatformState(ctx context.Context, uuid string, params param.ChangeSNSApplicationPlatformStateParam) (*view.SNSApplicationPlatformInventoryView, error) {
	resp := view.SNSApplicationPlatformInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/sns/application-platforms", uuid, "actions", "inventory", map[string]interface{}{
		"changeSNSApplicationPlatformState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidatePriceUserConfig operates on PriceUserConfig
func (cli *ZSClient) ValidatePriceUserConfig(ctx context.Context, params param.ValidatePriceUserConfigParam) (*view.ValidatePriceUserConfigEventView, error) {
	resp := view.ValidatePriceUserConfigEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/billings/accounts/actions", "", "", map[string]interface{}{
		"validatePriceUserConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateResourcePrice updates ResourcePrice
func (cli *ZSClient) UpdateResourcePrice(ctx context.Context, uuid string, params param.UpdateResourcePriceParam) (*view.PriceInventoryView, error) {
	resp := view.PriceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/billings/prices", uuid, "actions", "inventory", map[string]interface{}{
		"updateResourcePrice": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveActionFromEventSubscription removes ActionFromEventSubscription
func (cli *ZSClient) RemoveActionFromEventSubscription(ctx context.Context, subscriptionUuid string, actionUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/zwatch/events/subscriptions", subscriptionUuid, fmt.Sprintf("actions/%s", actionUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// CheckKVMHostConfigFile operates on KVMHostConfigFile
func (cli *ZSClient) CheckKVMHostConfigFile(ctx context.Context) (*view.CheckHostConfigFileView, error) {
	resp := view.CheckHostConfigFileView{}
	if err := cli.Post(ctx, "v1/hosts/kvm/from-file/check", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachTagFromResources operates on TagFromResources
func (cli *ZSClient) DetachTagFromResources(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/tags", uuid, string(deleteMode))
}

// GetCandidateLdapEntryForBinding gets CandidateLdapEntryForBinding by uuid
func (cli *ZSClient) GetCandidateLdapEntryForBinding(ctx context.Context) (*view.LdapEntryInventoryView, error) {
	var resp view.LdapEntryInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/ldap/entries/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNSSnmpTestConnection operates on SnmpTestConnection
func (cli *ZSClient) SNSSnmpTestConnection(ctx context.Context, params param.SNSSnmpTestConnectionParam) (*view.SNSSnmpTestConnectionEventView, error) {
	resp := view.SNSSnmpTestConnectionEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/snmp/test-connection"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryGCJob queries GCJob list
func (cli *ZSClient) QueryGCJob(ctx context.Context, params *param.QueryParam) ([]view.GarbageCollectorInventoryView, error) {
	var resp []view.GarbageCollectorInventoryView
	return resp, cli.List(ctx, "v1/gc-jobs", params, &resp)
}

func (cli *ZSClient) GetGCJob(ctx context.Context, uuid string) (*view.GarbageCollectorInventoryView, error) {
	var resp view.GarbageCollectorInventoryView
	if err := cli.Get(ctx, "v1/gc-jobs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageGCJob Pagination
func (cli *ZSClient) PageGCJob(ctx context.Context, params *param.QueryParam) ([]view.GarbageCollectorInventoryView, int, error) {
	var gCJobs []view.GarbageCollectorInventoryView
	total, err := cli.Page(ctx, "v1/gc-jobs", params, &gCJobs)
	return gCJobs, total, err
}

// ChangeHostState changes HostState
func (cli *ZSClient) ChangeHostState(ctx context.Context, uuid string, params param.ChangeHostStateParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts", uuid, "actions", "inventory", map[string]interface{}{
		"changeHostState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVmNicMac updates VmNicMac
func (cli *ZSClient) UpdateVmNicMac(ctx context.Context, vmNicUuid string, params param.UpdateVmNicMacParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances/nics", vmNicUuid, "actions", "inventory", map[string]interface{}{
		"updateVmNicMac": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmInstanceHaLevel deletes VmInstanceHaLevel
func (cli *ZSClient) DeleteVmInstanceHaLevel(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// DetachNvmeServerFromCluster operates on NvmeServerFromCluster
func (cli *ZSClient) DetachNvmeServerFromCluster(ctx context.Context, clusterUuid string, uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/clusters", clusterUuid, fmt.Sprintf("storage-devices/nvme/servers/%s", uuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetBackupStorageTypes gets BackupStorageTypes by uuid
func (cli *ZSClient) GetBackupStorageTypes(ctx context.Context) (*view.GetBackupStorageTypesView, error) {
	var resp view.GetBackupStorageTypesView
	if err := cli.GetWithRespKey(ctx, "v1/backup-storage/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeQos gets VolumeQos by uuid
func (cli *ZSClient) GetVolumeQos(ctx context.Context, uuid string) (*view.GetVolumeQosView, error) {
	var resp view.GetVolumeQosView
	if err := cli.GetWithRespKey(ctx, "v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteResourcePrice deletes ResourcePrice
func (cli *ZSClient) DeleteResourcePrice(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/billings/prices", uuid, string(deleteMode))
}

// DeleteMetricData deletes MetricData
func (cli *ZSClient) DeleteMetricData(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/metrics", uuid, string(deleteMode))
}

// AddRemoteCidrsToIPsecConnection adds RemoteCidrsToIPsecConnection
func (cli *ZSClient) AddRemoteCidrsToIPsecConnection(ctx context.Context, uuid string, params param.AddRemoteCidrsToIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/ipsec/%s/remote-cidrs", uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOnBaremetalChassis operates on PowerOnBaremetalChassis
func (cli *ZSClient) PowerOnBaremetalChassis(ctx context.Context, chassisUuid string, params param.PowerOnBaremetalChassisParam) (*view.PowerOnBaremetalChassisEventView, error) {
	resp := view.PowerOnBaremetalChassisEventView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/chassis", chassisUuid, "actions", "", map[string]interface{}{
		"powerOnBaremetalChassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddLabelToAlarm adds LabelToAlarm
func (cli *ZSClient) AddLabelToAlarm(ctx context.Context, alarmUuid string, params param.AddLabelToAlarmParam) (*view.AlarmLabelInventoryView, error) {
	resp := view.AlarmLabelInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/alarms/%s/labels", alarmUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeFromVolumeSnapshot creates DataVolumeFromVolumeSnapshot
func (cli *ZSClient) CreateDataVolumeFromVolumeSnapshot(ctx context.Context, volumeSnapshotUuid string, params param.CreateDataVolumeFromVolumeSnapshotParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/volumes/data/from/volume-snapshots/%s", volumeSnapshotUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachIsoFromVmInstance operates on IsoFromVmInstance
func (cli *ZSClient) DetachIsoFromVmInstance(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// DetachSecurityGroupFromL3Network operates on SecurityGroupFromL3Network
func (cli *ZSClient) DetachSecurityGroupFromL3Network(ctx context.Context, securityGroupUuid string, l3NetworkUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/security-groups", securityGroupUuid, fmt.Sprintf("l3-networks/%s", l3NetworkUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// ExportVmOvaPackage operates on VmOvaPackage
func (cli *ZSClient) ExportVmOvaPackage(ctx context.Context, params param.ExportVmOvaPackageParam) (*view.ImagePackageInventoryView, error) {
	resp := view.ImagePackageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/ovf/ova-packages"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportVmOvaPackageAsync Async
func (cli *ZSClient) ExportVmOvaPackageAsync(ctx context.Context, params param.ExportVmOvaPackageParam) (string, error) {

	resource := "ovf/ova-packages"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// RevertVmFromCdpBackup operates on VmFromCdpBackup
func (cli *ZSClient) RevertVmFromCdpBackup(ctx context.Context, vmInstanceUuid string, params param.RevertVmFromCdpBackupParam) (*view.RevertVmFromCdpBackupEventView, error) {
	resp := view.RevertVmFromCdpBackupEventView{}
	if err := cli.PutWithSpec(ctx, "v1/cdp-backups", vmInstanceUuid, "actions", "", map[string]interface{}{
		"revertVmFromCdpBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNSFeiShuTestConnection operates on FeiShuTestConnection
func (cli *ZSClient) SNSFeiShuTestConnection(ctx context.Context, params param.SNSFeiShuTestConnectionParam) (*view.SNSFeiShuTestConnectionEventView, error) {
	resp := view.SNSFeiShuTestConnectionEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/feishu/test-connection"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVirtualizerInfo gets VirtualizerInfo by uuid
func (cli *ZSClient) GetVirtualizerInfo(ctx context.Context) (*view.VirtualizerInfoInventoryView, error) {
	var resp view.VirtualizerInfoInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/virtualizer-info", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkIpStatistic gets L3NetworkIpStatistic by uuid
func (cli *ZSClient) GetL3NetworkIpStatistic(ctx context.Context, uuid string) (*view.GetL3NetworkIpStatisticView, error) {
	var resp view.GetL3NetworkIpStatisticView
	if err := cli.GetWithRespKey(ctx, "v1/l3-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncAccountsFromLdapServer operates on AccountsFromLdapServer
func (cli *ZSClient) SyncAccountsFromLdapServer(ctx context.Context, uuid string, params param.SyncAccountsFromLdapServerParam) (*view.SyncAccountsFromLdapServerEventView, error) {
	resp := view.SyncAccountsFromLdapServerEventView{}
	if err := cli.PutWithSpec(ctx, "v1/ldap/servers", uuid, "actions", "", map[string]interface{}{
		"syncAccountsFromLdapServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSchedulerExecutionReport gets SchedulerExecutionReport by uuid
func (cli *ZSClient) GetSchedulerExecutionReport(ctx context.Context) (*view.GetSchedulerExecutionReportView, error) {
	var resp view.GetSchedulerExecutionReportView
	if err := cli.GetWithRespKey(ctx, "v1/scheduler/report", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRuleFromConfigFile creates FirewallRuleFromConfigFile
func (cli *ZSClient) CreateFirewallRuleFromConfigFile(ctx context.Context, params param.CreateFirewallRuleFromConfigFileParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	resp := view.VpcFirewallRuleSetInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpcfirewalls/rules/from-file"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostSensors gets HostSensors by uuid
func (cli *ZSClient) GetHostSensors(ctx context.Context, uuid string) (*view.GetHostSensorsView, error) {
	var resp view.GetHostSensorsView
	if err := cli.GetWithRespKey(ctx, "v1/hosts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImageCandidatesForVmToChange gets ImageCandidatesForVmToChange by uuid
func (cli *ZSClient) GetImageCandidatesForVmToChange(ctx context.Context, uuid string) (*view.ImageInventoryView, error) {
	var resp view.ImageInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeImageState changes ImageState
func (cli *ZSClient) ChangeImageState(ctx context.Context, uuid string, params param.ChangeImageStateParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/images", uuid, "actions", "inventory", map[string]interface{}{
		"changeImageState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// KvmRunShell operates on KvmRunShell
func (cli *ZSClient) KvmRunShell(ctx context.Context, params param.KvmRunShellParam) (*view.KvmRunShellEventView, error) {
	resp := view.KvmRunShellEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/kvm/actions", "", "inventory", map[string]interface{}{
		"kvmRunShell": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRolePolicyActions gets RolePolicyActions by uuid
func (cli *ZSClient) GetRolePolicyActions(ctx context.Context) (*view.GetRolePolicyActionsView, error) {
	var resp view.GetRolePolicyActionsView
	if err := cli.GetWithRespKey(ctx, "v1/identities/role/policy-actions", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVRouterOspfArea updates VRouterOspfArea
func (cli *ZSClient) UpdateVRouterOspfArea(ctx context.Context, uuid string, params param.UpdateVRouterOspfAreaParam) (*view.RouterAreaInventoryView, error) {
	resp := view.RouterAreaInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/routerArea", uuid, "actions", "inventory", map[string]interface{}{
		"updateVRouterOspfArea": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoverBackupFromImageStoreBackupStorage operates on BackupFromImageStoreBackupStorage
func (cli *ZSClient) RecoverBackupFromImageStoreBackupStorage(ctx context.Context, uuid string, params param.RecoverBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volume-backups", uuid, "actions", "inventory", map[string]interface{}{
		"recoverBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageTypes gets PrimaryStorageTypes by uuid
func (cli *ZSClient) GetPrimaryStorageTypes(ctx context.Context) (*view.GetPrimaryStorageTypesView, error) {
	var resp view.GetPrimaryStorageTypesView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateZStoneHostConfig updates ZStoneHostConfig
func (cli *ZSClient) UpdateZStoneHostConfig(ctx context.Context, params param.UpdateZStoneHostConfigParam) (*view.UpdateZStoneHostConfigEventView, error) {
	resp := view.UpdateZStoneHostConfigEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/zstone-plugin/config/host", "", "", map[string]interface{}{
		"updateZStoneHostConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryIPSecConnection queries IPSecConnection list
func (cli *ZSClient) QueryIPSecConnection(ctx context.Context, params *param.QueryParam) ([]view.IPsecConnectionInventoryView, error) {
	var resp []view.IPsecConnectionInventoryView
	return resp, cli.List(ctx, "v1/ipsec", params, &resp)
}

func (cli *ZSClient) GetIPSecConnection(ctx context.Context, uuid string) (*view.IPsecConnectionInventoryView, error) {
	var resp view.IPsecConnectionInventoryView
	if err := cli.Get(ctx, "v1/ipsec", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIPSecConnection Pagination
func (cli *ZSClient) PageIPSecConnection(ctx context.Context, params *param.QueryParam) ([]view.IPsecConnectionInventoryView, int, error) {
	var iPSecConnections []view.IPsecConnectionInventoryView
	total, err := cli.Page(ctx, "v1/ipsec", params, &iPSecConnections)
	return iPSecConnections, total, err
}

// BatchDeleteVolumeSnapshot operates on DeleteVolumeSnapshot
func (cli *ZSClient) BatchDeleteVolumeSnapshot(ctx context.Context, params param.BatchDeleteVolumeSnapshotParam) (*view.BatchDeleteVolumeSnapshotEventView, error) {
	resp := view.BatchDeleteVolumeSnapshotEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/volume-snapshots/batch-delete", "", "", map[string]interface{}{
		"batchDeleteVolumeSnapshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryThirdPartyAccountSourceBinding queries ThirdPartyAccountSourceBinding list
func (cli *ZSClient) QueryThirdPartyAccountSourceBinding(ctx context.Context, params *param.QueryParam) ([]view.AccountThirdPartyAccountSourceRefInventoryView, error) {
	var resp []view.AccountThirdPartyAccountSourceRefInventoryView
	return resp, cli.List(ctx, "v1/account-import/bindings", params, &resp)
}

// PageThirdPartyAccountSourceBinding Pagination
func (cli *ZSClient) PageThirdPartyAccountSourceBinding(ctx context.Context, params *param.QueryParam) ([]view.AccountThirdPartyAccountSourceRefInventoryView, int, error) {
	var thirdPartyAccountSourceBindings []view.AccountThirdPartyAccountSourceRefInventoryView
	total, err := cli.Page(ctx, "v1/account-import/bindings", params, &thirdPartyAccountSourceBindings)
	return thirdPartyAccountSourceBindings, total, err
}

// ReloadLicense operates on ReloadLicense
func (cli *ZSClient) ReloadLicense(ctx context.Context, params param.ReloadLicenseParam) (*view.LicenseInventoryView, error) {
	resp := view.LicenseInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/licenses/actions", "", "inventory", map[string]interface{}{
		"reloadLicense": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteNicQos deletes NicQos
func (cli *ZSClient) DeleteNicQos(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// GetResourceStackVmStatus gets ResourceStackVmStatus by uuid
func (cli *ZSClient) GetResourceStackVmStatus(ctx context.Context) (*view.GetResourceStackVmStatusView, error) {
	var resp view.GetResourceStackVmStatusView
	if err := cli.GetWithRespKey(ctx, "v1/cloudformation/stack/monitor/vmstatus", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExpungeDataVolume operates on DataVolume
func (cli *ZSClient) ExpungeDataVolume(ctx context.Context, uuid string, params param.ExpungeDataVolumeParam) error {
	return cli.Put(ctx, "v1/volumes", uuid, map[string]interface{}{
		"expungeDataVolume": params.Params,
	}, nil)
}

// AddActionToEventSubscription adds ActionToEventSubscription
func (cli *ZSClient) AddActionToEventSubscription(ctx context.Context, subscriptionUuid string, params param.AddActionToEventSubscriptionParam) (*view.EventSubscriptionInventoryView, error) {
	resp := view.EventSubscriptionInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/events/subscriptions/%s/actions", subscriptionUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVRouterRouterId gets VRouterRouterId by uuid
func (cli *ZSClient) GetVRouterRouterId(ctx context.Context, uuid string) (*view.GetVRouterRouterIdView, error) {
	var resp view.GetVRouterRouterIdView
	if err := cli.GetWithRespKey(ctx, "v1/routerArea", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetZBoxBackupDetails gets ZBoxBackupDetails by uuid
func (cli *ZSClient) GetZBoxBackupDetails(ctx context.Context, uuid string) (*view.GetZBoxBackupDetailsView, error) {
	var resp view.GetZBoxBackupDetailsView
	if err := cli.GetWithRespKey(ctx, "v1/externalbackup/zbox", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetExternalServices gets ExternalServices by uuid
func (cli *ZSClient) GetExternalServices(ctx context.Context) (*view.ExternalServiceInventoryView, error) {
	var resp view.ExternalServiceInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/external/services", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveDnsFromVpcRouter removes DnsFromVpcRouter
func (cli *ZSClient) RemoveDnsFromVpcRouter(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vpc/virtual-routers", uuid, string(deleteMode))
}

// GetCandidateNetworkInterfaces gets CandidateNetworkInterfaces by uuid
func (cli *ZSClient) GetCandidateNetworkInterfaces(ctx context.Context) (*view.GetCandidateNetworkInterfacesView, error) {
	var resp view.GetCandidateNetworkInterfacesView
	if err := cli.GetWithRespKey(ctx, "v1/cluster/hosts-network-interfaces", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckMemorySnapshotGroupConflict operates on MemorySnapshotGroupConflict
func (cli *ZSClient) CheckMemorySnapshotGroupConflict(ctx context.Context, uuid string) (*view.CheckMemorySnapshotGroupConflictView, error) {
	var resp view.CheckMemorySnapshotGroupConflictView
	if err := cli.GetWithRespKey(ctx, "v1/memory-snapshots/groups", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAccessControlListServerGroup changes AccessControlListServerGroup
func (cli *ZSClient) ChangeAccessControlListServerGroup(ctx context.Context, aclUuid string, params param.ChangeAccessControlListServerGroupParam) (*view.LoadBalancerListerAclView, error) {
	resp := view.LoadBalancerListerAclView{}
	if err := cli.PutWithSpec(ctx, "v1/load-balancers/listener/acl", aclUuid, "servergroup/actions", "inventory", map[string]interface{}{
		"changeAccessControlListServerGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAvailableTriggers gets AvailableTriggers by uuid
func (cli *ZSClient) GetAvailableTriggers(ctx context.Context) (*view.SchedulerTriggerInventoryView, error) {
	var resp view.SchedulerTriggerInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/scheduler/triggers/available", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReimageVmInstance operates on ReimageVmInstance
func (cli *ZSClient) ReimageVmInstance(ctx context.Context, vmInstanceUuid string, params param.ReimageVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "actions", "inventory", map[string]interface{}{
		"reimageVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MoveAccountGroup operates on MoveAccountGroup
func (cli *ZSClient) MoveAccountGroup(ctx context.Context, uuid string, params param.MoveAccountGroupParam) (*view.AccountGroupInventoryView, error) {
	resp := view.AccountGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/account-groups", uuid, "actions", "inventory", map[string]interface{}{
		"moveAccountGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAtPersonOfAtFeiShuEndpoint updates AtPersonOfAtFeiShuEndpoint
func (cli *ZSClient) UpdateAtPersonOfAtFeiShuEndpoint(ctx context.Context, uuid string, params param.UpdateAtPersonOfAtFeiShuEndpointParam) (*view.SNSFeiShuAtPersonInventoryView, error) {
	resp := view.SNSFeiShuAtPersonInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/sns/application-endpoints/feishu/at-persons", uuid, "actions", "inventory", map[string]interface{}{
		"updateAtPersonOfAtFeiShuEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallRuleTemplate queries FirewallRuleTemplate list
func (cli *ZSClient) QueryFirewallRuleTemplate(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleTemplateInventoryView, error) {
	var resp []view.VpcFirewallRuleTemplateInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls/rules/templates", params, &resp)
}

func (cli *ZSClient) GetFirewallRuleTemplate(ctx context.Context, uuid string) (*view.VpcFirewallRuleTemplateInventoryView, error) {
	var resp view.VpcFirewallRuleTemplateInventoryView
	if err := cli.Get(ctx, "v1/vpcfirewalls/rules/templates/uuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallRuleTemplate Pagination
func (cli *ZSClient) PageFirewallRuleTemplate(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleTemplateInventoryView, int, error) {
	var firewallRuleTemplates []view.VpcFirewallRuleTemplateInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls/rules/templates", params, &firewallRuleTemplates)
	return firewallRuleTemplates, total, err
}

// CreateL2HardwareVxlanNetwork creates L2HardwareVxlanNetwork
func (cli *ZSClient) CreateL2HardwareVxlanNetwork(ctx context.Context, params param.CreateL2HardwareVxlanNetworkParam) (*view.CreateL2HardwareVxlanNetworkEventView, error) {
	resp := view.CreateL2HardwareVxlanNetworkEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l2-networks/hardware-vxlan"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageStoreBackupStorageQuota operates on ImageStoreBackupStorageQuota
func (cli *ZSClient) SetImageStoreBackupStorageQuota(ctx context.Context, params param.SetImageStoreBackupStorageQuotaParam) (*view.SetImageStoreBackupStorageQuotaEventView, error) {
	resp := view.SetImageStoreBackupStorageQuotaEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage/image-store/actions", "", "", map[string]interface{}{
		"setImageStoreBackupStorageQuota": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeClusterState changes ClusterState
func (cli *ZSClient) ChangeClusterState(ctx context.Context, uuid string, params param.ChangeClusterStateParam) (*view.ClusterInventoryView, error) {
	resp := view.ClusterInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/clusters", uuid, "actions", "inventory", map[string]interface{}{
		"changeClusterState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVfNicHaState changes VfNicHaState
func (cli *ZSClient) ChangeVfNicHaState(ctx context.Context, vfNicUuid string, params param.ChangeVfNicHaStateParam) (*view.VmVfNicInventoryView, error) {
	resp := view.VmVfNicInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances/nics", vfNicUuid, "actions", "inventory", map[string]interface{}{
		"changeVfNicHaState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGlobalConfigOptions gets GlobalConfigOptions by uuid
func (cli *ZSClient) GetGlobalConfigOptions(ctx context.Context, category string, name string) (*view.GetGlobalConfigOptionsView, error) {
	var resp view.GetGlobalConfigOptionsView
	err := cli.GetWithSpec(ctx, "v1/global-configurations", category, fmt.Sprintf("%s", name), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ApplyMonitorTemplateToMonitorGroup operates on MonitorTemplateToMonitorGroup
func (cli *ZSClient) ApplyMonitorTemplateToMonitorGroup(ctx context.Context, templateUuid, groupUuid string, params param.ApplyMonitorTemplateToMonitorGroupParam) (*view.MonitorGroupTemplateRefInventoryView, error) {
	resp := view.MonitorGroupTemplateRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/monitortemplates/%s/monitorgroups/%s", templateUuid, groupUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PutMetricData operates on PutMetricData
func (cli *ZSClient) PutMetricData(ctx context.Context, params param.PutMetricDataParam) (*view.PutMetricDataEventView, error) {
	resp := view.PutMetricDataEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/metrics"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAttachablePublicL3ForVRouter gets AttachablePublicL3ForVRouter by uuid
func (cli *ZSClient) GetAttachablePublicL3ForVRouter(ctx context.Context, uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/appliances/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RerunLongJob operates on RerunLongJob
func (cli *ZSClient) RerunLongJob(ctx context.Context, uuid string, params param.RerunLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/longjobs", uuid, "actions", "inventory", map[string]interface{}{
		"rerunLongJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangePortMirrorState changes PortMirrorState
func (cli *ZSClient) ChangePortMirrorState(ctx context.Context, uuid string, params param.ChangePortMirrorStateParam) (*view.PortMirrorInventoryView, error) {
	resp := view.PortMirrorInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/port-mirrors", uuid, "actions", "inventory", map[string]interface{}{
		"changePortMirrorState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetNicQos operates on NicQos
func (cli *ZSClient) SetNicQos(ctx context.Context, uuid string, params param.SetNicQosParam) (*view.SetNicQosEventView, error) {
	resp := view.SetNicQosEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"setNicQos": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsubscribeSNSTopic operates on UnsubscribeSNSTopic
func (cli *ZSClient) UnsubscribeSNSTopic(ctx context.Context, topicUuid string, endpointUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/sns/topics", topicUuid, fmt.Sprintf("endpoints/%s", endpointUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// CancelLongJob operates on CancelLongJob
func (cli *ZSClient) CancelLongJob(ctx context.Context, uuid string, params param.CancelLongJobParam) (*view.CancelLongJobEventView, error) {
	resp := view.CancelLongJobEventView{}
	if err := cli.PutWithSpec(ctx, "v1/longjobs", uuid, "actions", "", map[string]interface{}{
		"cancelLongJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRouteTableVpcVRouterCandidate gets RouteTableVpcVRouterCandidate by uuid
func (cli *ZSClient) GetRouteTableVpcVRouterCandidate(ctx context.Context) (*view.VpcRouterVmInventoryView, error) {
	var resp view.VpcRouterVmInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vpc/virtual-routers/get-vpc-candidate", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteExportedImageFromBackupStorage deletes ExportedImageFromBackupStorage
func (cli *ZSClient) DeleteExportedImageFromBackupStorage(ctx context.Context, backupStorageUuid string, imageUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/backup-storage", backupStorageUuid, fmt.Sprintf("exported-images/%s", imageUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// QueryFirewallRuleSet queries FirewallRuleSet list
func (cli *ZSClient) QueryFirewallRuleSet(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleSetInventoryView, error) {
	var resp []view.VpcFirewallRuleSetInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls/ruleSets", params, &resp)
}

func (cli *ZSClient) GetFirewallRuleSet(ctx context.Context, uuid string) (*view.VpcFirewallRuleSetInventoryView, error) {
	var resp view.VpcFirewallRuleSetInventoryView
	if err := cli.Get(ctx, "v1/vpcfirewalls/ruleSets", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallRuleSet Pagination
func (cli *ZSClient) PageFirewallRuleSet(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleSetInventoryView, int, error) {
	var firewallRuleSets []view.VpcFirewallRuleSetInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls/ruleSets", params, &firewallRuleSets)
	return firewallRuleSets, total, err
}

// GetVmUsbRedirect gets VmUsbRedirect by uuid
func (cli *ZSClient) GetVmUsbRedirect(ctx context.Context, uuid string) (*view.GetVmUsbRedirectView, error) {
	var resp view.GetVmUsbRedirectView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateClusterOS updates ClusterOS
func (cli *ZSClient) UpdateClusterOS(ctx context.Context, uuid string, params param.UpdateClusterOSParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/clusters", uuid, "actions", "inventory", map[string]interface{}{
		"updateClusterOS": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateAccountBilling operates on AccountBilling
func (cli *ZSClient) GenerateAccountBilling(ctx context.Context, accountUuid string, params param.GenerateAccountBillingParam) (*view.GenerateAccountBillingEventView, error) {
	resp := view.GenerateAccountBillingEventView{}
	if err := cli.PutWithSpec(ctx, "v1/billings/accounts", accountUuid, "actions", "", map[string]interface{}{
		"generateAccountBilling": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEventData gets EventData by uuid
func (cli *ZSClient) GetEventData(ctx context.Context) (*view.GetEventDataView, error) {
	var resp view.GetEventDataView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/events", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckIpAvailability operates on IpAvailability
func (cli *ZSClient) CheckIpAvailability(ctx context.Context, l3NetworkUuid string, ip string) (*view.CheckIpAvailabilityView, error) {
	var resp view.CheckIpAvailabilityView
	err := cli.GetWithSpec(ctx, "v1/l3-networks", l3NetworkUuid, fmt.Sprintf("ip/%s/availability", ip), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ZStoneTestConnection operates on ZStoneTestConnection
func (cli *ZSClient) ZStoneTestConnection(ctx context.Context, params param.ZStoneTestConnectionParam) (*view.ZStoneTestConnectionView, error) {
	resp := view.ZStoneTestConnectionView{}
	if err := cli.PutWithRespKey(ctx, "v1/zstone-plugin/test-connection", "", "", map[string]interface{}{
		"zStoneTestConnection": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateHostKernelInterfaces gets CandidateHostKernelInterfaces by uuid
func (cli *ZSClient) GetCandidateHostKernelInterfaces(ctx context.Context) (*view.GetCandidateHostKernelInterfacesView, error) {
	var resp view.GetCandidateHostKernelInterfacesView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/kernel-interfaces", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveVmNicFromLoadBalancer removes VmNicFromLoadBalancer
func (cli *ZSClient) RemoveVmNicFromLoadBalancer(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/load-balancers/listeners", uuid, string(deleteMode))
}

// CalculateResourceSpending operates on ResourceSpending
func (cli *ZSClient) CalculateResourceSpending(ctx context.Context, params param.CalculateResourceSpendingParam) (*view.CalculateResourceSpendingView, error) {
	resp := view.CalculateResourceSpendingView{}
	if err := cli.PutWithRespKey(ctx, "v1/billings/resources/actions", "", "", map[string]interface{}{
		"calculateResourceSpending": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAlarmRecord queries AlarmRecord list
func (cli *ZSClient) QueryAlarmRecord(ctx context.Context, params *param.QueryParam) ([]view.AlarmRecordsInventoryView, error) {
	var resp []view.AlarmRecordsInventoryView
	return resp, cli.List(ctx, "v1/zwatch/alarm-records", params, &resp)
}

// PageAlarmRecord Pagination
func (cli *ZSClient) PageAlarmRecord(ctx context.Context, params *param.QueryParam) ([]view.AlarmRecordsInventoryView, int, error) {
	var alarmRecords []view.AlarmRecordsInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/alarm-records", params, &alarmRecords)
	return alarmRecords, total, err
}

// GetVRouterFlowCounter gets VRouterFlowCounter by uuid
func (cli *ZSClient) GetVRouterFlowCounter(ctx context.Context, uuid string) (*view.GetVRouterFlowCounterView, error) {
	var resp view.GetVRouterFlowCounterView
	if err := cli.GetWithRespKey(ctx, "v1/flowmeters", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachBackupStorageFromZone operates on BackupStorageFromZone
func (cli *ZSClient) DetachBackupStorageFromZone(ctx context.Context, zoneUuid string, backupStorageUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/zones", zoneUuid, fmt.Sprintf("backup-storage/%s", backupStorageUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// PowerResetBaremetalChassis operates on PowerResetBaremetalChassis
func (cli *ZSClient) PowerResetBaremetalChassis(ctx context.Context, chassisUuid string, params param.PowerResetBaremetalChassisParam) (*view.PowerResetBaremetalChassisEventView, error) {
	resp := view.PowerResetBaremetalChassisEventView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/chassis", chassisUuid, "actions", "", map[string]interface{}{
		"powerResetBaremetalChassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostPowerStatus gets HostPowerStatus by uuid
func (cli *ZSClient) GetHostPowerStatus(ctx context.Context, uuid string, params param.GetHostPowerStatusParam) (*view.HostIpmiInventoryView, error) {
	resp := view.HostIpmiInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/power", uuid, "actions", "", map[string]interface{}{
		"getHostPowerStatus": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetChainTask gets ChainTask by uuid
func (cli *ZSClient) GetChainTask(ctx context.Context) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.GetWithRespKey(ctx, "v1/core/task-details", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpTrashOnPrimaryStorage operates on UpTrashOnPrimaryStorage
func (cli *ZSClient) CleanUpTrashOnPrimaryStorage(ctx context.Context, uuid string, params param.CleanUpTrashOnPrimaryStorageParam) (*view.CleanUpTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpTrashOnPrimaryStorageEventView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage", uuid, "trash/actions", "", map[string]interface{}{
		"cleanUpTrashOnPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddDisasterImageStoreBackupStorage adds DisasterImageStoreBackupStorage
func (cli *ZSClient) AddDisasterImageStoreBackupStorage(ctx context.Context, params param.AddDisasterImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	resp := view.ImageStoreBackupStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/backup-storage/image-store/disaster"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmSchedulingRulesExecuteState gets VmSchedulingRulesExecuteState by uuid
func (cli *ZSClient) GetVmSchedulingRulesExecuteState(ctx context.Context, params param.GetVmSchedulingRulesExecuteStateParam) (*view.GetVmSchedulingRulesExecuteStateView, error) {
	resp := view.GetVmSchedulingRulesExecuteStateView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/get/vmSchedulingRules/conflict/state"), "ruleMapState", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVolumesSnapshot creates VolumesSnapshot
func (cli *ZSClient) CreateVolumesSnapshot(ctx context.Context, params param.CreateVolumesSnapshotParam) (*view.VolumeSnapshotInventoryView, error) {
	resp := view.VolumeSnapshotInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/volumes/volume-snapshots"), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIpAddressCapacity gets IpAddressCapacity by uuid
func (cli *ZSClient) GetIpAddressCapacity(ctx context.Context) (*view.GetIpAddressCapacityView, error) {
	var resp view.GetIpAddressCapacityView
	if err := cli.GetWithRespKey(ctx, "v1/ip-capacity", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeHostPassword changes HostPassword
func (cli *ZSClient) ChangeHostPassword(ctx context.Context, hostUuid string, params param.ChangeHostPasswordParam) (*view.ChangeHostPasswordEventView, error) {
	resp := view.ChangeHostPasswordEventView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/kvm", hostUuid, "actions", "", map[string]interface{}{
		"changeHostPassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSlbInstance creates SlbInstance
func (cli *ZSClient) CreateSlbInstance(ctx context.Context, params param.CreateSlbInstanceParam) (*view.SlbVmInstanceInventoryView, error) {
	resp := view.SlbVmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/load-balancers/slb/instances"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangePortForwardingRuleState changes PortForwardingRuleState
func (cli *ZSClient) ChangePortForwardingRuleState(ctx context.Context, uuid string, params param.ChangePortForwardingRuleStateParam) (*view.PortForwardingRuleInventoryView, error) {
	resp := view.PortForwardingRuleInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/port-forwarding", uuid, "actions", "inventory", map[string]interface{}{
		"changePortForwardingRuleState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetZceXCapability gets ZceXCapability by uuid
func (cli *ZSClient) GetZceXCapability(ctx context.Context) (*view.GetZceXCapabilityView, error) {
	var resp view.GetZceXCapabilityView
	if err := cli.GetWithRespKey(ctx, "v1/zce-x-plugin/capability", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrometheusQueryLabelValues operates on PrometheusQueryLabelValues
func (cli *ZSClient) PrometheusQueryLabelValues(ctx context.Context) (*view.PrometheusQueryLabelValuesView, error) {
	var resp view.PrometheusQueryLabelValuesView
	if err := cli.GetWithRespKey(ctx, "v1/prometheus/labels", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PauseVmInstance operates on PauseVmInstance
func (cli *ZSClient) PauseVmInstance(ctx context.Context, uuid string, params param.PauseVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "inventory", map[string]interface{}{
		"pauseVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSignatureServerEncryptPublicKey gets SignatureServerEncryptPublicKey by uuid
func (cli *ZSClient) GetSignatureServerEncryptPublicKey(ctx context.Context) (*view.GetSignatureServerEncryptPublicKeyView, error) {
	var resp view.GetSignatureServerEncryptPublicKeyView
	if err := cli.GetWithRespKey(ctx, "v1/secret-resource-pool-token/signature-server-encrypt-public-key", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateClusterSupportDRS operates on ClusterSupportDRS
func (cli *ZSClient) ValidateClusterSupportDRS(ctx context.Context, uuid string) (*view.ValidateClusterSupportDRSView, error) {
	var resp view.ValidateClusterSupportDRSView
	if err := cli.GetWithRespKey(ctx, "v1/clusters", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShrinkVolumeSnapshot operates on ShrinkVolumeSnapshot
func (cli *ZSClient) ShrinkVolumeSnapshot(ctx context.Context, uuid string, params param.ShrinkVolumeSnapshotParam) (*view.ShrinkVolumeSnapshotEventView, error) {
	resp := view.ShrinkVolumeSnapshotEventView{}
	if err := cli.PutWithSpec(ctx, "v1/volume-snapshots/shrink", uuid, "actions", "shrinkResult", map[string]interface{}{
		"shrinkVolumeSnapshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachRoleFromAccountGroup operates on RoleFromAccountGroup
func (cli *ZSClient) DetachRoleFromAccountGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/account-groups", uuid, string(deleteMode))
}

// AddBackupStoragesToReplicationGroup adds BackupStoragesToReplicationGroup
func (cli *ZSClient) AddBackupStoragesToReplicationGroup(ctx context.Context, replicationGroupUuid string, params param.AddBackupStoragesToReplicationGroupParam) (*view.ImageReplicationGroupBackupStorageRefInventoryView, error) {
	resp := view.ImageReplicationGroupBackupStorageRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/image-replication-groups/%s", replicationGroupUuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddHostToHostSchedulingRuleGroup adds HostToHostSchedulingRuleGroup
func (cli *ZSClient) AddHostToHostSchedulingRuleGroup(ctx context.Context, hostGroupUuid, hostUuid string, params param.AddHostToHostSchedulingRuleGroupParam) (*view.AddHostToHostSchedulingRuleGroupEventView, error) {
	resp := view.AddHostToHostSchedulingRuleGroupEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hostSchedulingRuleGroup/%s/host/%s", hostGroupUuid, hostUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmNicAttachedNetworkService gets VmNicAttachedNetworkService by uuid
func (cli *ZSClient) GetVmNicAttachedNetworkService(ctx context.Context, uuid string) (*view.GetVmNicAttachedNetworkServiceView, error) {
	var resp view.GetVmNicAttachedNetworkServiceView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/nics", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFirewallRuleSet deletes FirewallRuleSet
func (cli *ZSClient) DeleteFirewallRuleSet(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vpcfirewalls/ruleSets", uuid, string(deleteMode))
}

// GetVmHostname gets VmHostname by uuid
func (cli *ZSClient) GetVmHostname(ctx context.Context, uuid string) (*view.GetVmHostnameView, error) {
	var resp view.GetVmHostnameView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSchedulerJobsToSchedulerJobGroup adds SchedulerJobsToSchedulerJobGroup
func (cli *ZSClient) AddSchedulerJobsToSchedulerJobGroup(ctx context.Context, schedulerJobGroupUuid string, params param.AddSchedulerJobsToSchedulerJobGroupParam) (*view.SchedulerJobGroupJobRefInventoryView, error) {
	resp := view.SchedulerJobGroupJobRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/scheduler/jobgroups/%s/job", schedulerJobGroupUuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LocalStorageMigrateVolume operates on LocalStorageMigrateVolume
func (cli *ZSClient) LocalStorageMigrateVolume(ctx context.Context, volumeUuid string, params param.LocalStorageMigrateVolumeParam) (*view.LocalStorageResourceRefInventoryView, error) {
	resp := view.LocalStorageResourceRefInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage/local-storage/volumes", volumeUuid, "actions", "inventory", map[string]interface{}{
		"localStorageMigrateVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachL3NetworkFromVm operates on L3NetworkFromVm
func (cli *ZSClient) DetachL3NetworkFromVm(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances/nics", uuid, string(deleteMode))
}

// AttachNicToBonding operates on NicToBonding
func (cli *ZSClient) AttachNicToBonding(ctx context.Context, uuid string, params param.AttachNicToBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/bondings", uuid, "attach", "inventory", map[string]interface{}{
		"attachNicToBonding": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVRouterOspfArea creates VRouterOspfArea
func (cli *ZSClient) CreateVRouterOspfArea(ctx context.Context, params param.CreateVRouterOspfAreaParam) (*view.RouterAreaInventoryView, error) {
	resp := view.RouterAreaInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/routerArea"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetSecurityMachineKey operates on SecurityMachineKey
func (cli *ZSClient) SetSecurityMachineKey(ctx context.Context, uuid string, params param.SetSecurityMachineKeyParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/secret-resource-pool-token/set/%s/actions", uuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeTemplateFromVolume creates DataVolumeTemplateFromVolume
func (cli *ZSClient) CreateDataVolumeTemplateFromVolume(ctx context.Context, volumeUuid string, params param.CreateDataVolumeTemplateFromVolumeParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/images/data-volume-templates/from/volumes/%s", volumeUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeTemplateFromVolumeAsync Async
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeAsync(ctx context.Context, params param.CreateDataVolumeTemplateFromVolumeParam) (string, error) {

	resource := "images/data-volume-templates/from/volumes/{volumeUuid}"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// CreateOAuthClient creates OAuthClient
func (cli *ZSClient) CreateOAuthClient(ctx context.Context, params param.CreateOAuthClientParam) (*view.OAuth2ClientInventoryView, error) {
	resp := view.OAuth2ClientInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/create/oauth2/client"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedEip gets VpcAttachedEip by uuid
func (cli *ZSClient) GetVpcAttachedEip(ctx context.Context, uuid string, params param.GetVpcAttachedEipParam) (*view.EipInventoryView, error) {
	resp := view.EipInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/virtual-routers/%s/attached-eip", uuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveSchedulerJobFromSchedulerTrigger removes SchedulerJobFromSchedulerTrigger
func (cli *ZSClient) RemoveSchedulerJobFromSchedulerTrigger(ctx context.Context, schedulerJobUuid string, schedulerTriggerUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/scheduler/jobs", schedulerJobUuid, fmt.Sprintf("scheduler/triggers/%s", schedulerTriggerUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// ExportDatabaseBackupFromBackupStorage operates on DatabaseBackupFromBackupStorage
func (cli *ZSClient) ExportDatabaseBackupFromBackupStorage(ctx context.Context, databaseBackupUuid string, backupStorageUuid string, params param.ExportDatabaseBackupFromBackupStorageParam) (*view.ExportDatabaseBackupFromBackupStorageEventView, error) {
	resp := view.ExportDatabaseBackupFromBackupStorageEventView{}
	err := cli.PutWithSpec(ctx, "v1/database-backups", databaseBackupUuid, fmt.Sprintf("backup-storage/%s/actions", backupStorageUuid), "", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeIPSecConnectionState changes IPSecConnectionState
func (cli *ZSClient) ChangeIPSecConnectionState(ctx context.Context, uuid string, params param.ChangeIPSecConnectionStateParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/ipsec", uuid, "actions", "inventory", map[string]interface{}{
		"changeIPSecConnectionState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeMediaState changes MediaState
func (cli *ZSClient) ChangeMediaState(ctx context.Context, uuid string, params param.ChangeMediaStateParam) (*view.MediaInventoryView, error) {
	resp := view.MediaInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/media", uuid, "actions", "inventory", map[string]interface{}{
		"changeMediaState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSSOClient gets SSOClient by uuid
func (cli *ZSClient) GetSSOClient(ctx context.Context) (*view.ThirdPartyAccountSourceInventoryView, error) {
	var resp view.ThirdPartyAccountSourceInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/get/sso/client", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEmailMonitorTriggerAction creates EmailMonitorTrigger
func (cli *ZSClient) CreateEmailMonitorTriggerAction(ctx context.Context, params param.CreateEmailMonitorTriggerActionParam) (*view.MonitorTriggerActionInventoryView, error) {
	resp := view.MonitorTriggerActionInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/monitoring/trigger-actions/emails"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVpcVRouterDistributedRoutingEnabled operates on VpcVRouterDistributedRoutingEnabled
func (cli *ZSClient) SetVpcVRouterDistributedRoutingEnabled(ctx context.Context, uuid string, params param.SetVpcVRouterDistributedRoutingEnabledParam) (*view.SetVpcVRouterDistributedRoutingEnabledEventView, error) {
	resp := view.SetVpcVRouterDistributedRoutingEnabledEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/virtual-routers/%s/distributed-routing", uuid), "enabled", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDirectoryUsage gets DirectoryUsage by uuid
func (cli *ZSClient) GetDirectoryUsage(ctx context.Context) (*view.GetDirectoryUsageView, error) {
	var resp view.GetDirectoryUsageView
	if err := cli.GetWithRespKey(ctx, "v1/software-package/directory/usage", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLocalRaidPhysicalDriveSmart gets LocalRaidPhysicalDriveSmart by uuid
func (cli *ZSClient) GetLocalRaidPhysicalDriveSmart(ctx context.Context, uuid string) (*view.GetLocalRaidPhysicalDriveSmartView, error) {
	var resp view.GetLocalRaidPhysicalDriveSmartView
	if err := cli.GetWithRespKey(ctx, "v1/storage-devices/local-raid/physical-drives", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVmNetworkConfig updates VmNetworkConfig
func (cli *ZSClient) UpdateVmNetworkConfig(ctx context.Context, vmInstanceUuid string, params param.UpdateVmNetworkConfigParam) (*view.UpdateVmNetworkConfigEventView, error) {
	resp := view.UpdateVmNetworkConfigEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "update-nic-config", "", map[string]interface{}{
		"updateVmNetworkConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHostname updates Hostname
func (cli *ZSClient) UpdateHostname(ctx context.Context, uuid string, params param.UpdateHostnameParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/hostname", uuid, "actions", "", map[string]interface{}{
		"updateHostname": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmStaticIp operates on VmStaticIp
func (cli *ZSClient) SetVmStaticIp(ctx context.Context, vmInstanceUuid string, params param.SetVmStaticIpParam) (*view.SetVmStaticIpEventView, error) {
	resp := view.SetVmStaticIpEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "actions", "", map[string]interface{}{
		"setVmStaticIp": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmSshKey gets VmSshKey by uuid
func (cli *ZSClient) GetVmSshKey(ctx context.Context, uuid string) (*view.GetVmSshKeyView, error) {
	var resp view.GetVmSshKeyView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateL2NetworkVirtualNetworkId updates L2NetworkVirtualNetworkId
func (cli *ZSClient) UpdateL2NetworkVirtualNetworkId(ctx context.Context, uuid string, params param.UpdateL2NetworkVirtualNetworkIdParam) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/l2-networks", uuid, "actions", "inventory", map[string]interface{}{
		"updateL2NetworkVirtualNetworkId": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmGuestToolsInfo gets VmGuestToolsInfo by uuid
func (cli *ZSClient) GetVmGuestToolsInfo(ctx context.Context, uuid string) (*view.GetVmGuestToolsInfoView, error) {
	var resp view.GetVmGuestToolsInfoView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateDiskOfferingUserConfig operates on DiskOfferingUserConfig
func (cli *ZSClient) ValidateDiskOfferingUserConfig(ctx context.Context, params param.ValidateDiskOfferingUserConfigParam) (*view.ValidateDiskOfferingUserConfigEventView, error) {
	resp := view.ValidateDiskOfferingUserConfigEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/billings/accounts/actions", "", "", map[string]interface{}{
		"validateDiskOfferingUserConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmRDP operates on VmRDP
func (cli *ZSClient) SetVmRDP(ctx context.Context, uuid string, params param.SetVmRDPParam) (*view.SetVmRDPEventView, error) {
	resp := view.SetVmRDPEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"setVmRDP": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RunSchedulerTrigger operates on RunSchedulerTrigger
func (cli *ZSClient) RunSchedulerTrigger(ctx context.Context, uuid string, params param.RunSchedulerTriggerParam) (*view.RunSchedulerTriggerEventView, error) {
	resp := view.RunSchedulerTriggerEventView{}
	if err := cli.PutWithSpec(ctx, "v1/scheduler/triggers", uuid, "actions", "", map[string]interface{}{
		"runSchedulerTrigger": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachL2NetworkToHost operates on L2NetworkToHost
func (cli *ZSClient) AttachL2NetworkToHost(ctx context.Context, l2NetworkUuid, hostUuid string, params param.AttachL2NetworkToHostParam) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l2-networks/%s/hosts/%s", l2NetworkUuid, hostUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOnHost operates on PowerOnHost
func (cli *ZSClient) PowerOnHost(ctx context.Context, uuid string, params param.PowerOnHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/power", uuid, "actions", "", map[string]interface{}{
		"powerOnHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AckAlarmData operates on AlarmData
func (cli *ZSClient) AckAlarmData(ctx context.Context, params param.AckAlarmDataParam) (*view.AlertDataAckInventoryView, error) {
	resp := view.AlertDataAckInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/alarm-histories/acknowledgments"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveDnsFromL3Network removes DnsFromL3Network
func (cli *ZSClient) RemoveDnsFromL3Network(ctx context.Context, l3NetworkUuid string, dns string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/l3-networks", l3NetworkUuid, fmt.Sprintf("dns/%s", dns), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// SNSWeComTestConnection operates on WeComTestConnection
func (cli *ZSClient) SNSWeComTestConnection(ctx context.Context, params param.SNSWeComTestConnectionParam) (*view.SNSWeComTestConnectionEventView, error) {
	resp := view.SNSWeComTestConnectionEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/we-com/test-connection"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveCertificateFromLoadBalancerListener removes CertificateFromLoadBalancerListener
func (cli *ZSClient) RemoveCertificateFromLoadBalancerListener(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/load-balancers/listeners", uuid, string(deleteMode))
}

// ProvisionVirtualRouterConfig operates on ProvisionVirtualRouterConfig
func (cli *ZSClient) ProvisionVirtualRouterConfig(ctx context.Context, vmInstanceUuid string, params param.ProvisionVirtualRouterConfigParam) (*view.ApplianceVmInventoryView, error) {
	resp := view.ApplianceVmInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances/appliances/virtual-routers", vmInstanceUuid, "provision", "inventory", map[string]interface{}{
		"provisionVirtualRouterConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmQga operates on VmQga
func (cli *ZSClient) SetVmQga(ctx context.Context, uuid string, params param.SetVmQgaParam) (*view.SetVmQgaEventView, error) {
	resp := view.SetVmQgaEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"setVmQga": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidatePassword operates on Password
func (cli *ZSClient) ValidatePassword(ctx context.Context, params param.ValidatePasswordParam) (*view.ValidatePasswordView, error) {
	resp := view.ValidatePasswordView{}
	if err := cli.PutWithRespKey(ctx, "v1/password/verify", "", "", map[string]interface{}{
		"validatePassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPortForwardingAttachableVmNics gets PortForwardingAttachableVmNics by uuid
func (cli *ZSClient) GetPortForwardingAttachableVmNics(ctx context.Context, uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/port-forwarding", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveRendezvousPointFromMulticastRouter removes RendezvousPointFromMulticastRouter
func (cli *ZSClient) RemoveRendezvousPointFromMulticastRouter(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/multicast/virtual-routers", uuid, string(deleteMode))
}

// GetChronyServers gets ChronyServers by uuid
func (cli *ZSClient) GetChronyServers(ctx context.Context) (*view.GetChronyServersView, error) {
	var resp view.GetChronyServersView
	if err := cli.GetWithRespKey(ctx, "v1/zops/chrony/servers", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachL3NetworkToVmNic operates on L3NetworkToVmNic
func (cli *ZSClient) AttachL3NetworkToVmNic(ctx context.Context, vmNicUuid, l3NetworkUuid string, params param.AttachL3NetworkToVmNicParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/nics/%s/l3-networks/%s", vmNicUuid, l3NetworkUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSecurityMachineState changes SecurityMachineState
func (cli *ZSClient) ChangeSecurityMachineState(ctx context.Context, uuid string, params param.ChangeSecurityMachineStateParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/security-machines", uuid, "actions", "inventory", map[string]interface{}{
		"changeSecurityMachineState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmQxlMemory operates on VmQxlMemory
func (cli *ZSClient) SetVmQxlMemory(ctx context.Context, uuid string, params param.SetVmQxlMemoryParam) (*view.SetVmQxlMemoryEventView, error) {
	resp := view.SetVmQxlMemoryEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"setVmQxlMemory": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SubscribeEvent operates on SubscribeEvent
func (cli *ZSClient) SubscribeEvent(ctx context.Context, params param.SubscribeEventParam) (*view.EventSubscriptionInventoryView, error) {
	resp := view.EventSubscriptionInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/events/subscriptions"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageCandidatesForVolumeMigration gets PrimaryStorageCandidatesForVolumeMigration by uuid
func (cli *ZSClient) GetPrimaryStorageCandidatesForVolumeMigration(ctx context.Context, uuid string) (*view.PrimaryStorageInventoryView, error) {
	var resp view.PrimaryStorageInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddLocalPrimaryStorage adds LocalPrimaryStorage
func (cli *ZSClient) AddLocalPrimaryStorage(ctx context.Context) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/local-storage", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeFormat gets VolumeFormat by uuid
func (cli *ZSClient) GetVolumeFormat(ctx context.Context) (*view.GetVolumeFormatView, error) {
	var resp view.GetVolumeFormatView
	if err := cli.GetWithRespKey(ctx, "v1/volumes/formats", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAtPersonOfAtDingTalkEndpoint updates AtPersonOfAtDingTalkEndpoint
func (cli *ZSClient) UpdateAtPersonOfAtDingTalkEndpoint(ctx context.Context, uuid string, params param.UpdateAtPersonOfAtDingTalkEndpointParam) (*view.SNSDingTalkAtPersonInventoryView, error) {
	resp := view.SNSDingTalkAtPersonInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/sns/application-endpoints/ding-talk/at-persons", uuid, "actions", "inventory", map[string]interface{}{
		"updateAtPersonOfAtDingTalkEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceAccount gets ResourceAccount by uuid
func (cli *ZSClient) GetResourceAccount(ctx context.Context) (*view.StringView, error) {
	var resp view.StringView
	if err := cli.GetWithRespKey(ctx, "v1/resources/accounts", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSimulatorBackupStorage adds SimulatorBackupStorage
func (cli *ZSClient) AddSimulatorBackupStorage(ctx context.Context, params param.AddSimulatorBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/backup-storage/simulators"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceEnsembleMembers gets ResourceEnsembleMembers by uuid
func (cli *ZSClient) GetResourceEnsembleMembers(ctx context.Context) (*view.ResourceEnsembleInventoryView, error) {
	var resp view.GetResourceEnsembleMembersView
	if err := cli.GetWithRespKey(ctx, "v1/iam1/resource-ensemble", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeSecretResourcePoolState changes SecretResourcePoolState
func (cli *ZSClient) ChangeSecretResourcePoolState(ctx context.Context, uuid string, params param.ChangeSecretResourcePoolStateParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/secret-resource-pools", uuid, "actions", "inventory", map[string]interface{}{
		"changeSecretResourcePoolState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVxlanL2Network deletes VxlanL2Network
func (cli *ZSClient) DeleteVxlanL2Network(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l2-networks/vxlan", uuid, string(deleteMode))
}

// RemoveVmFromAffinityGroup removes VmFromAffinityGroup
func (cli *ZSClient) RemoveVmFromAffinityGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/affinity-groups", uuid, string(deleteMode))
}

// SetVolumeIoThreadPin operates on VolumeIoThreadPin
func (cli *ZSClient) SetVolumeIoThreadPin(ctx context.Context, uuid string, params param.SetVolumeIoThreadPinParam) (*view.SetVolumeIoThreadPinEventView, error) {
	resp := view.SetVolumeIoThreadPinEventView{}
	if err := cli.PutWithSpec(ctx, "v1/volumes", uuid, "actions", "", map[string]interface{}{
		"setVolumeIoThreadPin": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePriorityConfig updates PriorityConfig
func (cli *ZSClient) UpdatePriorityConfig(ctx context.Context, uuid string, params param.UpdatePriorityConfigParam) (*view.UpdatePriorityConfigEventView, error) {
	resp := view.UpdatePriorityConfigEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-priority-config", uuid, "actions", "", map[string]interface{}{
		"updatePriorityConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IdentifyHost operates on IdentifyHost
func (cli *ZSClient) IdentifyHost(ctx context.Context, uuid string, params param.IdentifyHostParam) (*view.IdentifyHostEventView, error) {
	resp := view.IdentifyHostEventView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/kvm", uuid, "actions", "", map[string]interface{}{
		"identifyHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRootVolumeTemplateFromVolumeBackup creates RootVolumeTemplateFromVolumeBackup
func (cli *ZSClient) CreateRootVolumeTemplateFromVolumeBackup(ctx context.Context, backupUuid string, params param.CreateRootVolumeTemplateFromVolumeBackupParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/images/root-volume-templates/from/volume-template/%s", backupUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckFirewallRuleConfigFile operates on FirewallRuleConfigFile
func (cli *ZSClient) CheckFirewallRuleConfigFile(ctx context.Context, params param.CheckFirewallRuleConfigFileParam) (*view.CheckFirewallRuleConfigFileView, error) {
	resp := view.CheckFirewallRuleConfigFileView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpcfirewalls/rules/from-file/check"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateAffinityGroupForCreatingVm gets CandidateAffinityGroupForCreatingVm by uuid
func (cli *ZSClient) GetCandidateAffinityGroupForCreatingVm(ctx context.Context) (*view.AffinityGroupInventoryView, error) {
	var resp view.AffinityGroupInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/candidate-affinityGroup", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmConsoleAddress gets VmConsoleAddress by uuid
func (cli *ZSClient) GetVmConsoleAddress(ctx context.Context, uuid string) (*view.GetVmConsoleAddressView, error) {
	var resp view.GetVmConsoleAddressView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetFlowMeterRouterId operates on FlowMeterRouterId
func (cli *ZSClient) SetFlowMeterRouterId(ctx context.Context, vRouterUuid string, params param.SetFlowMeterRouterIdParam) (*view.SetFlowMeterRouterIdEventView, error) {
	resp := view.SetFlowMeterRouterIdEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/flowmeters/%s/routerid", vRouterUuid), "routerId", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckNetworkReachable operates on NetworkReachable
func (cli *ZSClient) CheckNetworkReachable(ctx context.Context) (*view.CheckNetworkReachableView, error) {
	var resp view.CheckNetworkReachableView
	if err := cli.GetWithRespKey(ctx, "v1/zops/check/network", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddStorageProtocol adds StorageProtocol
func (cli *ZSClient) AddStorageProtocol(ctx context.Context, params param.AddStorageProtocolParam) (*view.AddStorageProtocolEventView, error) {
	resp := view.AddStorageProtocolEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/primary-storage/protocols"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLoadBalancerListenerACLEntries gets LoadBalancerListenerACLEntries by uuid
func (cli *ZSClient) GetLoadBalancerListenerACLEntries(ctx context.Context) (*view.StringView, error) {
	var resp view.StringView
	if err := cli.GetWithRespKey(ctx, "v1/load-balancers/listeners/access-control-lists/entries", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHostIommuState updates HostIommuState
func (cli *ZSClient) UpdateHostIommuState(ctx context.Context, uuid string, params param.UpdateHostIommuStateParam) (*view.UpdateHostIommuStateEventView, error) {
	resp := view.UpdateHostIommuStateEventView{}
	if err := cli.PutWithSpec(ctx, "v1/pci-device/hosts", uuid, "actions", "state", map[string]interface{}{
		"updateHostIommuState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsubscribeEvent operates on UnsubscribeEvent
func (cli *ZSClient) UnsubscribeEvent(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/events/subscriptions", uuid, string(deleteMode))
}

// GetMonitorItem gets MonitorItem by uuid
func (cli *ZSClient) GetMonitorItem(ctx context.Context) (*view.ItemInventoryView, error) {
	var resp view.ItemInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/monitoring/items", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseRecords gets LicenseRecords by uuid
func (cli *ZSClient) GetLicenseRecords(ctx context.Context) (*view.LicenseHistoryInventoryView, error) {
	var resp view.LicenseHistoryInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/licenses/records", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachL2NetworkFromHost operates on L2NetworkFromHost
func (cli *ZSClient) DetachL2NetworkFromHost(ctx context.Context, l2NetworkUuid string, hostUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/l2-networks", l2NetworkUuid, fmt.Sprintf("hosts/%s", hostUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// RemoveMonFromCephPrimaryStorage removes MonFromCephPrimaryStorage
func (cli *ZSClient) RemoveMonFromCephPrimaryStorage(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/primary-storage/ceph", uuid, string(deleteMode))
}

// GetVmsSchedulingStateFromSchedulingRule gets VmsSchedulingStateFromSchedulingRule by uuid
func (cli *ZSClient) GetVmsSchedulingStateFromSchedulingRule(ctx context.Context, params param.GetVmsSchedulingStateFromSchedulingRuleParam) (*view.GetVmsSchedulingStateFromSchedulingRuleView, error) {
	resp := view.GetVmsSchedulingStateFromSchedulingRuleView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/get/vms/schedulingState/from/SchedulingRule"), "ruleMapState", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachSecurityGroupToL3Network operates on SecurityGroupToL3Network
func (cli *ZSClient) AttachSecurityGroupToL3Network(ctx context.Context, securityGroupUuid, l3NetworkUuid string, params param.AttachSecurityGroupToL3NetworkParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/security-groups/%s/l3-networks/%s", securityGroupUuid, l3NetworkUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAlarmState changes AlarmState
func (cli *ZSClient) ChangeAlarmState(ctx context.Context, uuid string, params param.ChangeAlarmStateParam) (*view.AlarmInventoryView, error) {
	resp := view.AlarmInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/alarms", uuid, "actions", "inventory", map[string]interface{}{
		"changeAlarmState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLocalStorageHostDiskCapacity gets LocalStorageHostDiskCapacity by uuid
func (cli *ZSClient) GetLocalStorageHostDiskCapacity(ctx context.Context, uuid string) (*view.HostDiskCapacityView, error) {
	var resp view.HostDiskCapacityView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/local-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVmNicDriver updates VmNicDriver
func (cli *ZSClient) UpdateVmNicDriver(ctx context.Context, vmInstanceUuid string, params param.UpdateVmNicDriverParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "actions", "inventory", map[string]interface{}{
		"updateVmNicDriver": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetIpOnHostNetworkInterface operates on IpOnHostNetworkInterface
func (cli *ZSClient) SetIpOnHostNetworkInterface(ctx context.Context, interfaceUuid string, params param.SetIpOnHostNetworkInterfaceParam) (*view.HostNetworkInterfaceInventoryView, error) {
	resp := view.HostNetworkInterfaceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hosts/nics/%s/ip", interfaceUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmSshKey deletes VmSshKey
func (cli *ZSClient) DeleteVmSshKey(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// DetachNicFromBonding operates on NicFromBonding
func (cli *ZSClient) DetachNicFromBonding(ctx context.Context, uuid string, params param.DetachNicFromBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/bondings", uuid, "detach", "inventory", map[string]interface{}{
		"detachNicFromBonding": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPolicyRouteRuleSetFromVirtualRouter gets PolicyRouteRuleSetFromVirtualRouter by uuid
func (cli *ZSClient) GetPolicyRouteRuleSetFromVirtualRouter(ctx context.Context, uuid string) (*view.PolicyRouteRuleSetInventoryView, error) {
	var resp view.PolicyRouteRuleSetInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/policy-routes/rulesets/virtualrouter", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVxlanPoolRemoteVtep deletes VxlanPoolRemoteVtep
func (cli *ZSClient) DeleteVxlanPoolRemoteVtep(ctx context.Context, l2NetworkUuid string, clusterUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/l2-networks", l2NetworkUuid, fmt.Sprintf("clusters/%s/delete/remote-vtep-ip", clusterUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// ChangeMonitorTriggerActionState changes MonitorTriggerActionState
func (cli *ZSClient) ChangeMonitorTriggerActionState(ctx context.Context, uuid string, params param.ChangeMonitorTriggerActionStateParam) (*view.MonitorTriggerActionInventoryView, error) {
	resp := view.MonitorTriggerActionInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/monitoring/trigger-actions", uuid, "actions", "inventory", map[string]interface{}{
		"changeMonitorTriggerActionState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoverDataVolume operates on DataVolume
func (cli *ZSClient) RecoverDataVolume(ctx context.Context, uuid string, params param.RecoverDataVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volumes", uuid, "actions", "inventory", map[string]interface{}{
		"recoverDataVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MigrateVm operates on Vm
func (cli *ZSClient) MigrateVm(ctx context.Context, vmInstanceUuid string, params param.MigrateVmParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "actions", "inventory", map[string]interface{}{
		"migrateVm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmPassword changes VmPassword
func (cli *ZSClient) ChangeVmPassword(ctx context.Context, uuid string, params param.ChangeVmPasswordParam) (*view.ChangeVmPasswordEventView, error) {
	resp := view.ChangeVmPasswordEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"changeVmPassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FlattenVmInstance operates on FlattenVmInstance
func (cli *ZSClient) FlattenVmInstance(ctx context.Context, uuid string, params param.FlattenVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "inventory", map[string]interface{}{
		"flattenVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcMulticastRoute gets VpcMulticastRoute by uuid
func (cli *ZSClient) GetVpcMulticastRoute(ctx context.Context, uuid string) (*view.MulticastRouteInventoryView, error) {
	var resp view.MulticastRouteInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/multicast/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEventRecord queries EventRecord list
func (cli *ZSClient) QueryEventRecord(ctx context.Context, params *param.QueryParam) ([]view.EventRecordsInventoryView, error) {
	var resp []view.EventRecordsInventoryView
	return resp, cli.List(ctx, "v1/zwatch/event-records", params, &resp)
}

// PageEventRecord Pagination
func (cli *ZSClient) PageEventRecord(ctx context.Context, params *param.QueryParam) ([]view.EventRecordsInventoryView, int, error) {
	var eventRecords []view.EventRecordsInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/event-records", params, &eventRecords)
	return eventRecords, total, err
}

// DeleteVmUserDefinedXmlHookScript deletes VmUserDefinedXmlHookScript
func (cli *ZSClient) DeleteVmUserDefinedXmlHookScript(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// SetVmUserDefinedXmlHookScript operates on VmUserDefinedXmlHookScript
func (cli *ZSClient) SetVmUserDefinedXmlHookScript(ctx context.Context, vmInstanceUuid string, params param.SetVmUserDefinedXmlHookScriptParam) (*view.SetVmUserDefinedXmlHookScriptEventView, error) {
	resp := view.SetVmUserDefinedXmlHookScriptEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "actions", "", map[string]interface{}{
		"setVmUserDefinedXmlHookScript": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostAllocatorStrategies gets HostAllocatorStrategies by uuid
func (cli *ZSClient) GetHostAllocatorStrategies(ctx context.Context) (*view.GetHostAllocatorStrategiesView, error) {
	var resp view.GetHostAllocatorStrategiesView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/allocators/strategies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCCSCertificateAccountState updates CCSCertificateAccountState
func (cli *ZSClient) UpdateCCSCertificateAccountState(ctx context.Context, accountUuid string, params param.UpdateCCSCertificateAccountStateParam) (*view.CCSCertificateInventoryView, error) {
	resp := view.CCSCertificateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/crypto/ccs-certificate/update-state/%s", accountUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncZBoxCapacity operates on ZBoxCapacity
func (cli *ZSClient) SyncZBoxCapacity(ctx context.Context, uuid string, params param.SyncZBoxCapacityParam) (*view.ZBoxInventoryView, error) {
	resp := view.ZBoxInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zbox", uuid, "actions", "", map[string]interface{}{
		"syncZBoxCapacity": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInterfaceServiceTypeStatistic gets InterfaceServiceTypeStatistic by uuid
func (cli *ZSClient) GetInterfaceServiceTypeStatistic(ctx context.Context) (*view.GetInterfaceServiceTypeStatisticView, error) {
	var resp view.GetInterfaceServiceTypeStatisticView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/hosts-network-interfaces/service-type-statistic", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AckEventData operates on EventData
func (cli *ZSClient) AckEventData(ctx context.Context, params param.AckEventDataParam) (*view.AlertDataAckInventoryView, error) {
	resp := view.AlertDataAckInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/event-histories/acknowledgments"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AllocateHostResource operates on HostResource
func (cli *ZSClient) AllocateHostResource(ctx context.Context, uuid string, params param.AllocateHostResourceParam) (*view.AllocateHostResourceEventView, error) {
	resp := view.AllocateHostResourceEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hosts/%s/allocate-resource", uuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVmsFromSchedulingState operates on ListVmsFromSchedulingState
func (cli *ZSClient) ListVmsFromSchedulingState(ctx context.Context, params param.ListVmsFromSchedulingStateParam) (*view.ListVmsFromSchedulingStateView, error) {
	resp := view.ListVmsFromSchedulingStateView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/list/vms/from/executeState"), "uuids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRootVolumeTemplateFromVolumeSnapshot creates RootVolumeTemplateFromVolumeSnapshot
func (cli *ZSClient) CreateRootVolumeTemplateFromVolumeSnapshot(ctx context.Context, snapshotUuid string, params param.CreateRootVolumeTemplateFromVolumeSnapshotParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/images/root-volume-templates/from/volume-snapshots/%s", snapshotUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateMiniHosts gets CandidateMiniHosts by uuid
func (cli *ZSClient) GetCandidateMiniHosts(ctx context.Context) (*view.GetCandidateMiniHostsView, error) {
	var resp view.GetCandidateMiniHostsView
	if err := cli.GetWithRespKey(ctx, "v1/mini-clusters/candidate-hosts", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckElaborationContent operates on ElaborationContent
func (cli *ZSClient) CheckElaborationContent(ctx context.Context, params param.CheckElaborationContentParam) (*view.CheckElaborationContentView, error) {
	resp := view.CheckElaborationContentView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/errorcode/elaborations/check"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmConsolePassword deletes VmConsolePassword
func (cli *ZSClient) DeleteVmConsolePassword(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// RevokeResourceSharing operates on RevokeResourceSharing
func (cli *ZSClient) RevokeResourceSharing(ctx context.Context, params param.RevokeResourceSharingParam) (*view.RevokeResourceSharingEventView, error) {
	resp := view.RevokeResourceSharingEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/accounts/resources/actions", "", "", map[string]interface{}{
		"revokeResourceSharing": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmBackup creates VmBackup
func (cli *ZSClient) CreateVmBackup(ctx context.Context, rootVolumeUuid string, params param.CreateVmBackupParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/volumes/%s/vm-backups", rootVolumeUuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmBackupAsync Async
func (cli *ZSClient) CreateVmBackupAsync(ctx context.Context, params param.CreateVmBackupParam) (string, error) {

	resource := "volumes/{rootVolumeUuid}/vm-backups"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetPrimaryStorageLicenseInfo gets PrimaryStorageLicenseInfo by uuid
func (cli *ZSClient) GetPrimaryStorageLicenseInfo(ctx context.Context, uuid string) (*view.GetPrimaryStorageLicenseInfoView, error) {
	var resp view.GetPrimaryStorageLicenseInfoView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeL3NetworkState changes L3NetworkState
func (cli *ZSClient) ChangeL3NetworkState(ctx context.Context, uuid string, params param.ChangeL3NetworkStateParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/l3-networks", uuid, "actions", "inventory", map[string]interface{}{
		"changeL3NetworkState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostNUMATopology gets HostNUMATopology by uuid
func (cli *ZSClient) GetHostNUMATopology(ctx context.Context, uuid string, params param.GetHostNUMATopologyParam) (*view.GetHostNUMATopologyEventView, error) {
	resp := view.GetHostNUMATopologyEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hosts/%s/numa", uuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateL2VirtualSwitch creates L2VirtualSwitch
func (cli *ZSClient) CreateL2VirtualSwitch(ctx context.Context, params param.CreateL2VirtualSwitchParam) (*view.CreateL2VirtualSwitchEventView, error) {
	resp := view.CreateL2VirtualSwitchEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l2-networks/virtual-switch"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVmNicToLoadBalancer adds VmNicToLoadBalancer
func (cli *ZSClient) AddVmNicToLoadBalancer(ctx context.Context, listenerUuid string, params param.AddVmNicToLoadBalancerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/load-balancers/listeners/%s/vm-instances/nics", listenerUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEncryptedField gets EncryptedField by uuid
func (cli *ZSClient) GetEncryptedField(ctx context.Context) (*view.GetEncryptedFieldView, error) {
	var resp view.GetEncryptedFieldView
	if err := cli.GetWithRespKey(ctx, "v1/encrypted/fields", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachBaremetalPxeServerToCluster operates on BaremetalPxeServerToCluster
func (cli *ZSClient) AttachBaremetalPxeServerToCluster(ctx context.Context, clusterUuid, pxeServerUuid string, params param.AttachBaremetalPxeServerToClusterParam) (*view.BaremetalPxeServerInventoryView, error) {
	resp := view.BaremetalPxeServerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/clusters/%s/pxeservers/%s", clusterUuid, pxeServerUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetClusterDRSStatus gets ClusterDRSStatus by uuid
func (cli *ZSClient) GetClusterDRSStatus(ctx context.Context) (*view.GetClusterDRSStatusView, error) {
	var resp view.GetClusterDRSStatusView
	if err := cli.GetWithRespKey(ctx, "v1/clusters/drs/status", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmStartingCandidateClustersHosts gets VmStartingCandidateClustersHosts by uuid
func (cli *ZSClient) GetVmStartingCandidateClustersHosts(ctx context.Context, uuid string) (*view.GetVmStartingCandidateClustersHostsView, error) {
	var resp view.GetVmStartingCandidateClustersHostsView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoverVmBackupFromImageStoreBackupStorage operates on VmBackupFromImageStoreBackupStorage
func (cli *ZSClient) RecoverVmBackupFromImageStoreBackupStorage(ctx context.Context, groupUuid string, params param.RecoverVmBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-backups", groupUuid, "actions", "inventories", map[string]interface{}{
		"recoverVmBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmNuma gets VmNuma by uuid
func (cli *ZSClient) GetVmNuma(ctx context.Context, uuid string) (*view.GetVmNumaView, error) {
	var resp view.GetVmNumaView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevokeResourceSharingToGroup operates on RevokeResourceSharingToGroup
func (cli *ZSClient) RevokeResourceSharingToGroup(ctx context.Context, params param.RevokeResourceSharingToGroupParam) (*view.RevokeResourceSharingToGroupEventView, error) {
	resp := view.RevokeResourceSharingToGroupEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/account-groups/resources/actions", "", "", map[string]interface{}{
		"revokeResourceSharingToGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DiscoverExternalPrimaryStorage operates on DiscoverExternalPrimaryStorage
func (cli *ZSClient) DiscoverExternalPrimaryStorage(ctx context.Context, params param.DiscoverExternalPrimaryStorageParam) (*view.ExternalPrimaryStorageInventoryView, error) {
	resp := view.ExternalPrimaryStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/primary-storage/addon/discover"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeZoneState changes ZoneState
func (cli *ZSClient) ChangeZoneState(ctx context.Context, uuid string, params param.ChangeZoneStateParam) (*view.ZoneInventoryView, error) {
	resp := view.ZoneInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zones", uuid, "actions", "inventory", map[string]interface{}{
		"changeZoneState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeIoThreadPin gets VolumeIoThreadPin by uuid
func (cli *ZSClient) GetVolumeIoThreadPin(ctx context.Context, uuid string) (*view.GetVolumeIoThreadPinView, error) {
	var resp view.GetVolumeIoThreadPinView
	if err := cli.GetWithRespKey(ctx, "v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateLdapBinding creates LdapBinding
func (cli *ZSClient) CreateLdapBinding(ctx context.Context, params param.CreateLdapBindingParam) (*view.AccountThirdPartyAccountSourceRefInventoryView, error) {
	resp := view.AccountThirdPartyAccountSourceRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/ldap/bindings"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolume creates DataVolume
func (cli *ZSClient) CreateDataVolume(ctx context.Context, params param.CreateDataVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/volumes/data"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchCreateBaremetalChassis operates on CreateBaremetalChassis
func (cli *ZSClient) BatchCreateBaremetalChassis(ctx context.Context, params param.BatchCreateBaremetalChassisParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/baremetal/chassis/from-file"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchCreateBaremetalChassisAsync Async
func (cli *ZSClient) BatchCreateBaremetalChassisAsync(ctx context.Context, params param.BatchCreateBaremetalChassisParam) (string, error) {

	resource := "baremetal/chassis/from-file"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// AddSchedulerJobToSchedulerTrigger adds SchedulerJobToSchedulerTrigger
func (cli *ZSClient) AddSchedulerJobToSchedulerTrigger(ctx context.Context, schedulerJobUuid, schedulerTriggerUuid string, params param.AddSchedulerJobToSchedulerTriggerParam) (*view.SchedulerJobSchedulerTriggerInventoryView, error) {
	resp := view.SchedulerJobSchedulerTriggerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/scheduler/jobs/%s/scheduler/triggers/%s", schedulerJobUuid, schedulerTriggerUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetZSha2Status gets ZSha2Status by uuid
func (cli *ZSClient) GetZSha2Status(ctx context.Context) (*view.ZSha2StatusViewView, error) {
	var resp view.GetZSha2StatusView
	if err := cli.GetWithRespKey(ctx, "v1/management-nodes/zsha2/status", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetNicQos gets NicQos by uuid
func (cli *ZSClient) GetNicQos(ctx context.Context, uuid string) (*view.GetNicQosView, error) {
	var resp view.GetNicQosView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedOspf gets VpcAttachedOspf by uuid
func (cli *ZSClient) GetVpcAttachedOspf(ctx context.Context, uuid string, params param.GetVpcAttachedOspfParam) (*view.NetworkRouterAreaRefInventoryView, error) {
	resp := view.NetworkRouterAreaRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/virtual-routers/%s/attached-ospf", uuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmNicNetwork changes VmNicNetwork
func (cli *ZSClient) ChangeVmNicNetwork(ctx context.Context, vmNicUuid, destL3NetworkUuid string, params param.ChangeVmNicNetworkParam) (*view.VmNicInventoryView, error) {
	resp := view.VmNicInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/nics/%s/l3-networks/%s", vmNicUuid, destL3NetworkUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAccountToGroup adds AccountToGroup
func (cli *ZSClient) AddAccountToGroup(ctx context.Context, groupUuid string, params param.AddAccountToGroupParam) (*view.AddAccountToGroupEventView, error) {
	resp := view.AddAccountToGroupEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/account-groups/%s/accounts", groupUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOffHost operates on PowerOffHost
func (cli *ZSClient) PowerOffHost(ctx context.Context, params param.PowerOffHostParam) (*view.PowerOffHostEventView, error) {
	resp := view.PowerOffHostEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/hosts/power-off/actions", "", "", map[string]interface{}{
		"powerOffHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveLabelFromAlarm removes LabelFromAlarm
func (cli *ZSClient) RemoveLabelFromAlarm(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/alarms/labels", uuid, string(deleteMode))
}

// UpdateVmPriority updates VmPriority
func (cli *ZSClient) UpdateVmPriority(ctx context.Context, uuid string, params param.UpdateVmPriorityParam) (*view.UpdateVmPriorityEventView, error) {
	resp := view.UpdateVmPriorityEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"updateVmPriority": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVersion gets Version by uuid
func (cli *ZSClient) GetVersion(ctx context.Context) (*view.GetVersionView, error) {
	resp := view.GetVersionView{}
	if err := cli.PutWithRespKey(ctx, "v1/management-nodes/actions", "", "", map[string]interface{}{
		"getVersion": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseCapabilities gets LicenseCapabilities by uuid
func (cli *ZSClient) GetLicenseCapabilities(ctx context.Context) (*view.GetLicenseCapabilitiesView, error) {
	var resp view.GetLicenseCapabilitiesView
	if err := cli.GetWithRespKey(ctx, "v1/licenses/capabilities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachMdevDeviceFromVm operates on MdevDeviceFromVm
func (cli *ZSClient) DetachMdevDeviceFromVm(ctx context.Context, mdevDeviceUuid string, vmInstanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/mdev-devices", mdevDeviceUuid, fmt.Sprintf("vm-instances/%s", vmInstanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// DeleteVmHostname deletes VmHostname
func (cli *ZSClient) DeleteVmHostname(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// GetCandidateBackupStorageForCreatingImage gets CandidateBackupStorageForCreatingImage by uuid
func (cli *ZSClient) GetCandidateBackupStorageForCreatingImage(ctx context.Context) (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/images/candidate-backup-storage", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) GetBackupStorageForCreatingImageFromVolume(ctx context.Context, uuid string) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	if err := cli.GetWithSpec(ctx, "v1/images/volumes", uuid, "candidate-backup-storage", "inventories", nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *ZSClient) GetBackupStorageForCreatingImageFromVolumeSnapshot(ctx context.Context, uuid string) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	if err := cli.GetWithSpec(ctx, "v1/images/volume-snapshots", uuid, "candidate-backup-storage", "inventories", nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// AttachAutoScalingTemplateToGroup operates on AutoScalingTemplateToGroup
func (cli *ZSClient) AttachAutoScalingTemplateToGroup(ctx context.Context, uuid, groupUuid string, params param.AttachAutoScalingTemplateToGroupParam) (*view.AutoScalingGroupInventoryView, error) {
	resp := view.AutoScalingGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/autoscaling/template/%s/groups/%s", uuid, groupUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCpuMemoryCapacity gets CpuMemoryCapacity by uuid
func (cli *ZSClient) GetCpuMemoryCapacity(ctx context.Context) (*view.GetCpuMemoryCapacityView, error) {
	var resp view.GetCpuMemoryCapacityView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/capacities/cpu-memory", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIntegrityResource adds IntegrityResource
func (cli *ZSClient) AddIntegrityResource(ctx context.Context, params param.AddIntegrityResourceParam) (*view.AddIntegrityResourceEventView, error) {
	resp := view.AddIntegrityResourceEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/integrity/resource/actions"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRuleTemplate creates FirewallRuleTemplate
func (cli *ZSClient) CreateFirewallRuleTemplate(ctx context.Context, params param.CreateFirewallRuleTemplateParam) (*view.VpcFirewallRuleTemplateInventoryView, error) {
	resp := view.VpcFirewallRuleTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpcfirewalls/rules/template"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckVipPortAvailability operates on VipPortAvailability
func (cli *ZSClient) CheckVipPortAvailability(ctx context.Context, uuid string) (*view.CheckVipPortAvailabilityView, error) {
	var resp view.CheckVipPortAvailabilityView
	if err := cli.GetWithRespKey(ctx, "v1/vips", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateClustersForAttachingL2Network gets CandidateClustersForAttachingL2Network by uuid
func (cli *ZSClient) GetCandidateClustersForAttachingL2Network(ctx context.Context, uuid string) (*view.ClusterInventoryView, error) {
	var resp view.ClusterInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/l2-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckScsiLunClusterStatus operates on ScsiLunClusterStatus
func (cli *ZSClient) CheckScsiLunClusterStatus(ctx context.Context, uuid string, clusterUuid string, params param.CheckScsiLunClusterStatusParam) (*view.ScsiLunClusterStatusInventoryView, error) {
	resp := view.ScsiLunClusterStatusInventoryView{}
	err := cli.PutWithSpec(ctx, "v1/storage-devices/scsi-lun", uuid, fmt.Sprintf("cluster/%s", clusterUuid), "inventory", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckBatchDataIntegrity operates on BatchDataIntegrity
func (cli *ZSClient) CheckBatchDataIntegrity(ctx context.Context) (*view.CheckBatchDataIntegrityView, error) {
	var resp view.CheckBatchDataIntegrityView
	if err := cli.GetWithRespKey(ctx, "v1/check/batch/data/integrity/", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAutoScalingGroupRemovalInstanceRule updates AutoScalingGroupRemovalInstanceRule
func (cli *ZSClient) UpdateAutoScalingGroupRemovalInstanceRule(ctx context.Context, uuid string, params param.UpdateAutoScalingGroupRemovalInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/autoscaling/rules/removal-instance", uuid, "actions", "inventory", map[string]interface{}{
		"updateAutoScalingGroupRemovalInstanceRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmSoundType operates on VmSoundType
func (cli *ZSClient) SetVmSoundType(ctx context.Context, uuid string, params param.SetVmSoundTypeParam) (*view.SetVmSoundTypeEventView, error) {
	resp := view.SetVmSoundTypeEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"setVmSoundType": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeL3NetworkDhcpIpAddress changes L3NetworkDhcpIpAddress
func (cli *ZSClient) ChangeL3NetworkDhcpIpAddress(ctx context.Context, l3NetworkUuid string, params param.ChangeL3NetworkDhcpIpAddressParam) (*view.ChangeL3NetworkDhcpIpAddressEventView, error) {
	resp := view.ChangeL3NetworkDhcpIpAddressEventView{}
	if err := cli.PutWithSpec(ctx, "v1/l3-networks", l3NetworkUuid, "dhcp-ip", "", map[string]interface{}{
		"changeL3NetworkDhcpIpAddress": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckVolumeSnapshotGroupAvailability operates on VolumeSnapshotGroupAvailability
func (cli *ZSClient) CheckVolumeSnapshotGroupAvailability(ctx context.Context) (*view.CheckVolumeSnapshotGroupAvailabilityView, error) {
	var resp view.CheckVolumeSnapshotGroupAvailabilityView
	if err := cli.GetWithRespKey(ctx, "v1/volume-snapshots/groups/availabilities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MergeDataOnBackupStorage operates on MergeDataOnBackupStorage
func (cli *ZSClient) MergeDataOnBackupStorage(ctx context.Context, backupStorageUuid string, params param.MergeDataOnBackupStorageParam) (*view.MergeDataOnBackupStorageEventView, error) {
	resp := view.MergeDataOnBackupStorageEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cdp-task/mergedata/%s", backupStorageUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddEmailAddressToSNSEmailEndpoint adds EmailAddressToSNSEmailEndpoint
func (cli *ZSClient) AddEmailAddressToSNSEmailEndpoint(ctx context.Context, params param.AddEmailAddressToSNSEmailEndpointParam) (*view.SNSEmailAddressInventoryView, error) {
	resp := view.SNSEmailAddressInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/emails/email-addresses"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAffinityGroupState changes AffinityGroupState
func (cli *ZSClient) ChangeAffinityGroupState(ctx context.Context, uuid string, params param.ChangeAffinityGroupStateParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/affinity-groups", uuid, "actions", "inventory", map[string]interface{}{
		"changeAffinityGroupState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVmNicInSecurityGroup queries VmNicInSecurityGroup list
func (cli *ZSClient) QueryVmNicInSecurityGroup(ctx context.Context, params *param.QueryParam) ([]view.VmNicSecurityGroupRefInventoryView, error) {
	var resp []view.VmNicSecurityGroupRefInventoryView
	return resp, cli.List(ctx, "v1/security-groups/vm-instances/nics", params, &resp)
}

// PageVmNicInSecurityGroup Pagination
func (cli *ZSClient) PageVmNicInSecurityGroup(ctx context.Context, params *param.QueryParam) ([]view.VmNicSecurityGroupRefInventoryView, int, error) {
	var vmNicInSecurityGroups []view.VmNicSecurityGroupRefInventoryView
	total, err := cli.Page(ctx, "v1/security-groups/vm-instances/nics", params, &vmNicInSecurityGroups)
	return vmNicInSecurityGroups, total, err
}

// ChangeSecurityGroupRuleState changes SecurityGroupRuleState
func (cli *ZSClient) ChangeSecurityGroupRuleState(ctx context.Context, securityGroupUuid string, params param.ChangeSecurityGroupRuleStateParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/security-groups", securityGroupUuid, "rules/state/actions", "inventory", map[string]interface{}{
		"changeSecurityGroupRuleState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoveryImageFromImageStoreBackupStorage operates on yImageFromImageStoreBackupStorage
func (cli *ZSClient) RecoveryImageFromImageStoreBackupStorage(ctx context.Context, uuid string, params param.RecoveryImageFromImageStoreBackupStorageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage", uuid, "actions", "inventory", map[string]interface{}{
		"recoveryImageFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVmNicToSecurityGroup adds VmNicToSecurityGroup
func (cli *ZSClient) AddVmNicToSecurityGroup(ctx context.Context, securityGroupUuid string, params param.AddVmNicToSecurityGroupParam) (*view.AddVmNicToSecurityGroupEventView, error) {
	resp := view.AddVmNicToSecurityGroupEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/security-groups/%s/vm-instances/nics", securityGroupUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEventFromResourceStack queries EventFromResourceStack list
func (cli *ZSClient) QueryEventFromResourceStack(ctx context.Context, params *param.QueryParam) ([]view.CloudFormationStackEventInventoryView, error) {
	var resp []view.CloudFormationStackEventInventoryView
	return resp, cli.List(ctx, "v1/cloudformation/event", params, &resp)
}

func (cli *ZSClient) GetEventFromResourceStack(ctx context.Context, uuid string) (*view.CloudFormationStackEventInventoryView, error) {
	var resp view.CloudFormationStackEventInventoryView
	if err := cli.Get(ctx, "v1/cloudformation/event", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEventFromResourceStack Pagination
func (cli *ZSClient) PageEventFromResourceStack(ctx context.Context, params *param.QueryParam) ([]view.CloudFormationStackEventInventoryView, int, error) {
	var eventFromResourceStacks []view.CloudFormationStackEventInventoryView
	total, err := cli.Page(ctx, "v1/cloudformation/event", params, &eventFromResourceStacks)
	return eventFromResourceStacks, total, err
}

// GetNodeRoles gets NodeRoles by uuid
func (cli *ZSClient) GetNodeRoles(ctx context.Context) (*view.NodeRolesViewView, error) {
	var resp view.NodeRolesViewView
	if err := cli.GetWithRespKey(ctx, "v1/zsv/nodes/roles", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevertVmFromSnapshotGroup operates on VmFromSnapshotGroup
func (cli *ZSClient) RevertVmFromSnapshotGroup(ctx context.Context, uuid string, params param.RevertVmFromSnapshotGroupParam) (*view.RevertVmFromSnapshotGroupEventView, error) {
	resp := view.RevertVmFromSnapshotGroupEventView{}
	if err := cli.PutWithSpec(ctx, "v1/volume-snapshots/group", uuid, "actions", "results", map[string]interface{}{
		"revertVmFromSnapshotGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateEmailAddressOfSNSEmailEndpoint updates EmailAddressOfSNSEmailEndpoint
func (cli *ZSClient) UpdateEmailAddressOfSNSEmailEndpoint(ctx context.Context, params param.UpdateEmailAddressOfSNSEmailEndpointParam) (*view.SNSEmailAddressInventoryView, error) {
	resp := view.SNSEmailAddressInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sns/application-endpoints/emails/email-addresses", "", "inventory", map[string]interface{}{
		"updateEmailAddressOfSNSEmailEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachCCSCertificateToAccount operates on CCSCertificateToAccount
func (cli *ZSClient) AttachCCSCertificateToAccount(ctx context.Context, accountUuid string, params param.AttachCCSCertificateToAccountParam) (*view.CCSCertificateInventoryView, error) {
	resp := view.CCSCertificateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/crypto/ccs-certificate/attach-account/%s", accountUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachFirewallRuleSetFromL3 operates on FirewallRuleSetFromL3
func (cli *ZSClient) DetachFirewallRuleSetFromL3(ctx context.Context, l3Uuid, ruleSetUuid string, params param.DetachFirewallRuleSetFromL3Param) (*view.DetachFirewallRuleSetFromL3EventView, error) {
	resp := view.DetachFirewallRuleSetFromL3EventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpcfirewalls/l3networks/%s/ruleSets/%s", l3Uuid, ruleSetUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryTwoFactorAuthentication queries TwoFactorAuthentication list
func (cli *ZSClient) QueryTwoFactorAuthentication(ctx context.Context, params *param.QueryParam) ([]view.TwoFactorAuthenticationSecretInventoryView, error) {
	var resp []view.TwoFactorAuthenticationSecretInventoryView
	return resp, cli.List(ctx, "v1/twofactorauthentication/secrets", params, &resp)
}

func (cli *ZSClient) GetTwoFactorAuthentication(ctx context.Context, uuid string) (*view.TwoFactorAuthenticationSecretInventoryView, error) {
	var resp view.TwoFactorAuthenticationSecretInventoryView
	if err := cli.Get(ctx, "v1/twofactorauthentication/secrets", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTwoFactorAuthentication Pagination
func (cli *ZSClient) PageTwoFactorAuthentication(ctx context.Context, params *param.QueryParam) ([]view.TwoFactorAuthenticationSecretInventoryView, int, error) {
	var twoFactorAuthentications []view.TwoFactorAuthenticationSecretInventoryView
	total, err := cli.Page(ctx, "v1/twofactorauthentication/secrets", params, &twoFactorAuthentications)
	return twoFactorAuthentications, total, err
}

// ListVmSchedulingRulesFromExecuteState operates on ListVmSchedulingRulesFromExecuteState
func (cli *ZSClient) ListVmSchedulingRulesFromExecuteState(ctx context.Context, params param.ListVmSchedulingRulesFromExecuteStateParam) (*view.ListVmSchedulingRulesFromExecuteStateView, error) {
	resp := view.ListVmSchedulingRulesFromExecuteStateView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/list/vmSchedulingRules/from/conflict/state"), "uuids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateImagesForCreatingVm gets CandidateImagesForCreatingVm by uuid
func (cli *ZSClient) GetCandidateImagesForCreatingVm(ctx context.Context, uuid string) (*view.ImageInventoryView, error) {
	var resp view.ImageInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/images/primaryStorage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmUserDefinedXml operates on VmUserDefinedXml
func (cli *ZSClient) SetVmUserDefinedXml(ctx context.Context, vmInstanceUuid string, params param.SetVmUserDefinedXmlParam) (*view.SetVmUserDefinedXmlEventView, error) {
	resp := view.SetVmUserDefinedXmlEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "actions", "", map[string]interface{}{
		"setVmUserDefinedXml": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageQga operates on ImageQga
func (cli *ZSClient) SetImageQga(ctx context.Context, uuid string, params param.SetImageQgaParam) (*view.SetImageQgaEventView, error) {
	resp := view.SetImageQgaEventView{}
	if err := cli.PutWithSpec(ctx, "v1/images", uuid, "actions", "", map[string]interface{}{
		"setImageQga": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVirtualSwitchUplinkGroup updates VirtualSwitchUplinkGroup
func (cli *ZSClient) UpdateVirtualSwitchUplinkGroup(ctx context.Context, uuid string, params param.UpdateVirtualSwitchUplinkGroupParam) (*view.UplinkGroupInventoryView, error) {
	resp := view.UplinkGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/l2-networks/virtual-switch", uuid, "uplink-group", "inventory", map[string]interface{}{
		"updateVirtualSwitchUplinkGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMetricLabelValue gets MetricLabelValue by uuid
func (cli *ZSClient) GetMetricLabelValue(ctx context.Context) (*view.GetMetricLabelValueView, error) {
	var resp view.GetMetricLabelValueView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/metrics/label-values", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateZonesClustersHostsForCreatingVm gets CandidateZonesClustersHostsForCreatingVm by uuid
func (cli *ZSClient) GetCandidateZonesClustersHostsForCreatingVm(ctx context.Context) (*view.GetCandidateZonesClustersHostsForCreatingVmView, error) {
	var resp view.GetCandidateZonesClustersHostsForCreatingVmView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/candidate-destinations", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// TakeVmConsoleScreenshot operates on TakeVmConsoleScreenshot
func (cli *ZSClient) TakeVmConsoleScreenshot(ctx context.Context, uuid string, params param.TakeVmConsoleScreenshotParam) (*view.TakeVmConsoleScreenshotEventView, error) {
	resp := view.TakeVmConsoleScreenshotEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"takeVmConsoleScreenshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveVRouterNetworksFromOspfArea removes VRouterNetworksFromOspfArea
func (cli *ZSClient) RemoveVRouterNetworksFromOspfArea(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/routerArea/networks", uuid, string(deleteMode))
}

// CreateResourcePrice creates ResourcePrice
func (cli *ZSClient) CreateResourcePrice(ctx context.Context, params param.CreateResourcePriceParam) (*view.PriceInventoryView, error) {
	resp := view.PriceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/billings/prices"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveSchedulerJobGroupFromSchedulerTrigger removes SchedulerJobGroupFromSchedulerTrigger
func (cli *ZSClient) RemoveSchedulerJobGroupFromSchedulerTrigger(ctx context.Context, schedulerJobGroupUuid string, schedulerTriggerUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/scheduler/jobgroups", schedulerJobGroupUuid, fmt.Sprintf("scheduler/triggers/%s", schedulerTriggerUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// UpdateHostNqn updates HostNqn
func (cli *ZSClient) UpdateHostNqn(ctx context.Context, uuid string, params param.UpdateHostNqnParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/nqn", uuid, "actions", "", map[string]interface{}{
		"updateHostNqn": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAccountPriceTableBinding changes AccountPriceTableBinding
func (cli *ZSClient) ChangeAccountPriceTableBinding(ctx context.Context, tableUuid string, accountUuid string, params param.ChangeAccountPriceTableBindingParam) (*view.PriceTableInventoryView, error) {
	resp := view.PriceTableInventoryView{}
	err := cli.PutWithSpec(ctx, "v1/billings/price-tables", tableUuid, fmt.Sprintf("accounts/%s", accountUuid), "inventory", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MountBlockDevice operates on MountBlockDevice
func (cli *ZSClient) MountBlockDevice(ctx context.Context, params param.MountBlockDeviceParam) (*view.MountBlockDeviceEventView, error) {
	resp := view.MountBlockDeviceEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/host/mount-block-device"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVolumeQos deletes VolumeQos
func (cli *ZSClient) DeleteVolumeQos(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/volumes", uuid, string(deleteMode))
}

// DeleteVmBackup deletes VmBackup
func (cli *ZSClient) DeleteVmBackup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-backups", uuid, string(deleteMode))
}

// SetVmSecurityLevel operates on VmSecurityLevel
func (cli *ZSClient) SetVmSecurityLevel(ctx context.Context, uuid string, params param.SetVmSecurityLevelParam) (*view.SetVmSecurityLevelEventView, error) {
	resp := view.SetVmSecurityLevelEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"setVmSecurityLevel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveMdevDeviceSpecFromVmInstance removes MdevDeviceSpecFromVmInstance
func (cli *ZSClient) RemoveMdevDeviceSpecFromVmInstance(ctx context.Context, mdevSpecUuid string, vmInstanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/mdev-device-specs", mdevSpecUuid, fmt.Sprintf("vm-instances/%s", vmInstanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// SyncVolumeSize operates on VolumeSize
func (cli *ZSClient) SyncVolumeSize(ctx context.Context, uuid string, params param.SyncVolumeSizeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volumes", uuid, "actions", "inventory", map[string]interface{}{
		"syncVolumeSize": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTrashOnBackupStorage gets TrashOnBackupStorage by uuid
func (cli *ZSClient) GetTrashOnBackupStorage(ctx context.Context) (*view.InstallPathRecycleInventoryView, error) {
	var resp view.InstallPathRecycleInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/backup-storage/trash", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkDhcpIpAddress gets L3NetworkDhcpIpAddress by uuid
func (cli *ZSClient) GetL3NetworkDhcpIpAddress(ctx context.Context, uuid string) (*view.GetL3NetworkDhcpIpAddressView, error) {
	var resp view.GetL3NetworkDhcpIpAddressView
	if err := cli.GetWithRespKey(ctx, "v1/l3-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeDiskOfferingState changes DiskOfferingState
func (cli *ZSClient) ChangeDiskOfferingState(ctx context.Context, uuid string, params param.ChangeDiskOfferingStateParam) (*view.DiskOfferingInventoryView, error) {
	resp := view.DiskOfferingInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/disk-offerings", uuid, "actions", "inventory", map[string]interface{}{
		"changeDiskOfferingState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRuleSet creates FirewallRuleSet
func (cli *ZSClient) CreateFirewallRuleSet(ctx context.Context, params param.CreateFirewallRuleSetParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	resp := view.VpcFirewallRuleSetInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpcfirewalls/ruleSets"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RequestConsoleAccess operates on RequestConsoleAccess
func (cli *ZSClient) RequestConsoleAccess(ctx context.Context, params param.RequestConsoleAccessParam) (*view.ConsoleInventoryView, error) {
	resp := view.ConsoleInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/consoles"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBaremetalChassisPowerStatus gets BaremetalChassisPowerStatus by uuid
func (cli *ZSClient) GetBaremetalChassisPowerStatus(ctx context.Context, uuid string) (*view.GetBaremetalChassisPowerStatusView, error) {
	var resp view.GetBaremetalChassisPowerStatusView
	if err := cli.GetWithRespKey(ctx, "v1/baremetal/chassis", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateEventData updates EventData
func (cli *ZSClient) UpdateEventData(ctx context.Context, params param.UpdateEventDataParam) (*view.UpdateEventDataEventView, error) {
	resp := view.UpdateEventDataEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/events/actions", "", "", map[string]interface{}{
		"updateEventData": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UngenerateSriovPciDevices operates on UngenerateSriovPciDevices
func (cli *ZSClient) UngenerateSriovPciDevices(ctx context.Context, pciDeviceUuid string, params param.UngenerateSriovPciDevicesParam) (*view.UngenerateVirtualPciDevicesEventView, error) {
	resp := view.UngenerateVirtualPciDevicesEventView{}
	if err := cli.PutWithSpec(ctx, "v1/pci-devices", pciDeviceUuid, "actions", "", map[string]interface{}{
		"ungenerateSriovPciDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshFirewall operates on Firewall
func (cli *ZSClient) RefreshFirewall(ctx context.Context, uuid string, params param.RefreshFirewallParam) (*view.VpcFirewallInventoryView, error) {
	resp := view.VpcFirewallInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vpcfirewalls/refresh", uuid, "actions", "inventory", map[string]interface{}{
		"refreshFirewall": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachL3NetworksFromIPsecConnection operates on L3NetworksFromIPsecConnection
func (cli *ZSClient) DetachL3NetworksFromIPsecConnection(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ipsec", uuid, string(deleteMode))
}

// UpdateAutoScalingGroupAddingNewInstanceRule updates AutoScalingGroupAddingNewInstanceRule
func (cli *ZSClient) UpdateAutoScalingGroupAddingNewInstanceRule(ctx context.Context, uuid string, params param.UpdateAutoScalingGroupAddingNewInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/autoscaling/rules/adding-new-instance", uuid, "actions", "inventory", map[string]interface{}{
		"updateAutoScalingGroupAddingNewInstanceRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmStaticIp deletes VmStaticIp
func (cli *ZSClient) DeleteVmStaticIp(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// AttachMonitorTriggerActionToTrigger operates on MonitorTriggerActionToTrigger
func (cli *ZSClient) AttachMonitorTriggerActionToTrigger(ctx context.Context, triggerUuid, actionUuid string, params param.AttachMonitorTriggerActionToTriggerParam) (*view.AttachMonitorTriggerActionToTriggerEventView, error) {
	resp := view.AttachMonitorTriggerActionToTriggerEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/monitoring/triggers/%s/trigger-actions/%s", triggerUuid, actionUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFaultToleranceVms gets FaultToleranceVms by uuid
func (cli *ZSClient) GetFaultToleranceVms(ctx context.Context) (*view.GetFaultToleranceVmsView, error) {
	var resp view.GetFaultToleranceVmsView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/fault-tolerance/sub-vms", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangePreconfigurationTemplateState changes PreconfigurationTemplateState
func (cli *ZSClient) ChangePreconfigurationTemplateState(ctx context.Context, uuid string, params param.ChangePreconfigurationTemplateStateParam) (*view.PreconfigurationTemplateInventoryView, error) {
	resp := view.PreconfigurationTemplateInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/preconfigurations", uuid, "actions", "inventory", map[string]interface{}{
		"changePreconfigurationTemplateState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromVolumeSnapshot creates VmInstanceFromVolumeSnapshot
func (cli *ZSClient) CreateVmInstanceFromVolumeSnapshot(ctx context.Context, volumeSnapshotUuid string, params param.CreateVmInstanceFromVolumeSnapshotParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/from/volume-snapshots/%s", volumeSnapshotUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckCephPlugin operates on CephPlugin
func (cli *ZSClient) CheckCephPlugin(ctx context.Context, params param.CheckCephPluginParam) (*view.CephPluginConnectionViewView, error) {
	resp := view.CephPluginConnectionViewView{}
	if err := cli.PutWithRespKey(ctx, "v1/ceph-plugin/check", "", "", map[string]interface{}{
		"checkCephPlugin": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHostIscsiInitiatorName updates HostIscsiInitiatorName
func (cli *ZSClient) UpdateHostIscsiInitiatorName(ctx context.Context, uuid string, params param.UpdateHostIscsiInitiatorNameParam) (*view.KVMHostInventoryView, error) {
	resp := view.KVMHostInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/kvm/iscsiInitiatorName", uuid, "actions", "", map[string]interface{}{
		"updateHostIscsiInitiatorName": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachL3NetworksToIPsecConnection operates on L3NetworksToIPsecConnection
func (cli *ZSClient) AttachL3NetworksToIPsecConnection(ctx context.Context, uuid string, params param.AttachL3NetworksToIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/ipsec/%s/l3networks", uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrometheusQueryVmMonitoringData operates on PrometheusQueryVmMonitoringData
func (cli *ZSClient) PrometheusQueryVmMonitoringData(ctx context.Context) (*view.PrometheusQueryVmMonitoringDataView, error) {
	var resp view.PrometheusQueryVmMonitoringDataView
	if err := cli.GetWithRespKey(ctx, "v1/prometheus/vm-instances", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateResourceConfigs updates ResourceConfigs
func (cli *ZSClient) UpdateResourceConfigs(ctx context.Context, resourceUuid string, params param.UpdateResourceConfigsParam) (*view.ResourceConfigStructView, error) {
	resp := view.ResourceConfigStructView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/resource-configurations/%s/resource-configs/actions", resourceUuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateZceXClusterConfig updates ZceXClusterConfig
func (cli *ZSClient) UpdateZceXClusterConfig(ctx context.Context, params param.UpdateZceXClusterConfigParam) (*view.UpdateZceXClusterConfigEventView, error) {
	resp := view.UpdateZceXClusterConfigEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/zce-x-plugin/config/cluster", "", "", map[string]interface{}{
		"updateZceXClusterConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevertVolumeFromSnapshot operates on VolumeFromSnapshot
func (cli *ZSClient) RevertVolumeFromSnapshot(ctx context.Context, uuid string, params param.RevertVolumeFromSnapshotParam) (*view.RevertVolumeFromSnapshotEventView, error) {
	resp := view.RevertVolumeFromSnapshotEventView{}
	if err := cli.PutWithSpec(ctx, "v1/volume-snapshots", uuid, "actions", "", map[string]interface{}{
		"revertVolumeFromSnapshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddNfsPrimaryStorage adds NfsPrimaryStorage
func (cli *ZSClient) AddNfsPrimaryStorage(ctx context.Context) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/nfs", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBlockPrimaryStorageMetadata gets BlockPrimaryStorageMetadata by uuid
func (cli *ZSClient) GetBlockPrimaryStorageMetadata(ctx context.Context, params param.GetBlockPrimaryStorageMetadataParam) (*view.BlockPrimaryStorageInventoryView, error) {
	resp := view.BlockPrimaryStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/primary-storage/block/metadata"), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBonding updates Bonding
func (cli *ZSClient) UpdateBonding(ctx context.Context, uuid string, params param.UpdateBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/bondings", uuid, "actions", "inventory", map[string]interface{}{
		"updateBonding": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetManagementNodeArch gets ManagementNodeArch by uuid
func (cli *ZSClient) GetManagementNodeArch(ctx context.Context) (*view.GetManagementNodeArchView, error) {
	resp := view.GetManagementNodeArchView{}
	if err := cli.PutWithRespKey(ctx, "v1/management-nodes/actions", "", "", map[string]interface{}{
		"getManagementNodeArch": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachScsiLunFromHost operates on ScsiLunFromHost
func (cli *ZSClient) DetachScsiLunFromHost(ctx context.Context, uuid string, params param.DetachScsiLunFromHostParam) (*view.ScsiLunInventoryView, error) {
	resp := view.ScsiLunInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/storage-devices/scsi-lun", uuid, "actions", "inventory", map[string]interface{}{
		"detachScsiLunFromHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableCbtTask operates on DisableCbtTask
func (cli *ZSClient) DisableCbtTask(ctx context.Context, uuid string, params param.DisableCbtTaskParam) (*view.CbtTaskInventoryView, error) {
	resp := view.CbtTaskInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cbt-task/disable/%s", uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshLocalRaid operates on LocalRaid
func (cli *ZSClient) RefreshLocalRaid(ctx context.Context, params param.RefreshLocalRaidParam) (*view.RaidControllerInventoryView, error) {
	resp := view.RaidControllerInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/storage-devices/local-raid/actions", "", "inventories", map[string]interface{}{
		"refreshLocalRaid": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachTagToResources operates on TagToResources
func (cli *ZSClient) AttachTagToResources(ctx context.Context, tagUuid string, params param.AttachTagToResourcesParam) (*view.AttachTagToResourcesEventView, error) {
	resp := view.AttachTagToResourcesEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/tags/%s/resources", tagUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteLdapBinding deletes LdapBinding
func (cli *ZSClient) DeleteLdapBinding(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ldap/bindings", uuid, string(deleteMode))
}

// UpdateSubscribeEvent updates SubscribeEvent
func (cli *ZSClient) UpdateSubscribeEvent(ctx context.Context, uuid string, params param.UpdateSubscribeEventParam) (*view.EventSubscriptionInventoryView, error) {
	resp := view.EventSubscriptionInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/events/subscriptions", uuid, "actions", "inventory", map[string]interface{}{
		"updateSubscribeEvent": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangePrimaryStorageState changes PrimaryStorageState
func (cli *ZSClient) ChangePrimaryStorageState(ctx context.Context, uuid string, params param.ChangePrimaryStorageStateParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage", uuid, "actions", "inventory", map[string]interface{}{
		"changePrimaryStorageState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedNetflow gets VpcAttachedNetflow by uuid
func (cli *ZSClient) GetVpcAttachedNetflow(ctx context.Context, uuid string, params param.GetVpcAttachedNetflowParam) (*view.FlowMeterInventoryView, error) {
	resp := view.FlowMeterInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/virtual-routers/%s/attached-netflow", uuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAuditData gets AuditData by uuid
func (cli *ZSClient) GetAuditData(ctx context.Context) (*view.GetAuditDataView, error) {
	var resp view.GetAuditDataView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/audits", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmSshKey operates on VmSshKey
func (cli *ZSClient) SetVmSshKey(ctx context.Context, uuid string, params param.SetVmSshKeyParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "inventory", map[string]interface{}{
		"setVmSshKey": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteLicense deletes License
func (cli *ZSClient) DeleteLicense(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/licenses/mn", uuid, string(deleteMode))
}

// GetSpiceCertificates gets SpiceCertificates by uuid
func (cli *ZSClient) GetSpiceCertificates(ctx context.Context) (*view.GetSpiceCertificatesView, error) {
	var resp view.GetSpiceCertificatesView
	if err := cli.GetWithRespKey(ctx, "v1/spice/certificates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncDatabaseBackupFromImageStoreBackupStorage operates on DatabaseBackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncDatabaseBackupFromImageStoreBackupStorage(ctx context.Context, uuid string, params param.SyncDatabaseBackupFromImageStoreBackupStorageParam) (*view.DatabaseBackupInventoryView, error) {
	resp := view.DatabaseBackupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/database-backups", uuid, "actions", "inventory", map[string]interface{}{
		"syncDatabaseBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFirewallIpSetTemplate deletes FirewallIpSetTemplate
func (cli *ZSClient) DeleteFirewallIpSetTemplate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vpcfirewalls/ipset/templates", uuid, string(deleteMode))
}

// SNSDingTalkTestConnection operates on DingTalkTestConnection
func (cli *ZSClient) SNSDingTalkTestConnection(ctx context.Context, params param.SNSDingTalkTestConnectionParam) (*view.SNSDingTalkTestConnectionEventView, error) {
	resp := view.SNSDingTalkTestConnectionEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/ding-talk/test-connection"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportImageFromBackupStorage operates on ImageFromBackupStorage
func (cli *ZSClient) ExportImageFromBackupStorage(ctx context.Context, backupStorageUuid string, params param.ExportImageFromBackupStorageParam) (*view.ExportImageFromBackupStorageEventView, error) {
	resp := view.ExportImageFromBackupStorageEventView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage", backupStorageUuid, "actions", "", map[string]interface{}{
		"exportImageFromBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FailoverFaultToleranceVm operates on FailoverFaultToleranceVm
func (cli *ZSClient) FailoverFaultToleranceVm(ctx context.Context, params param.FailoverFaultToleranceVmParam) (*view.FailoverFaultToleranceVmEventView, error) {
	resp := view.FailoverFaultToleranceVmEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/vm-instances/fault-tolerance", "", "", map[string]interface{}{
		"failoverFaultToleranceVm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EjectZBox operates on EjectZBox
func (cli *ZSClient) EjectZBox(ctx context.Context, uuid string, params param.EjectZBoxParam) (*view.ZBoxInventoryView, error) {
	resp := view.ZBoxInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zbox", uuid, "actions", "", map[string]interface{}{
		"ejectZBox": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrometheusQueryMetadata operates on PrometheusQueryMetadata
func (cli *ZSClient) PrometheusQueryMetadata(ctx context.Context) (*view.PrometheusQueryMetadataView, error) {
	var resp view.PrometheusQueryMetadataView
	if err := cli.GetWithRespKey(ctx, "v1/prometheus/meta-data", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallIpSetTemplate creates FirewallIpSetTemplate
func (cli *ZSClient) CreateFirewallIpSetTemplate(ctx context.Context, params param.CreateFirewallIpSetTemplateParam) (*view.VpcFirewallIpSetTemplateInventoryView, error) {
	resp := view.VpcFirewallIpSetTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpcfirewalls/ipset/templates"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachMonitorTriggerActionFromTrigger operates on MonitorTriggerActionFromTrigger
func (cli *ZSClient) DetachMonitorTriggerActionFromTrigger(ctx context.Context, triggerUuid string, actionUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/monitoring/triggers", triggerUuid, fmt.Sprintf("trigger-actions/%s", actionUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// DetachPolicyRouteRuleSetFromL3 operates on PolicyRouteRuleSetFromL3
func (cli *ZSClient) DetachPolicyRouteRuleSetFromL3(ctx context.Context, ruleSetUuid string, l3Uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/policy-routes/rulesets", ruleSetUuid, fmt.Sprintf("l3networks/%s", l3Uuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// QueryIpAddress queries IpAddress list
func (cli *ZSClient) QueryIpAddress(ctx context.Context, params *param.QueryParam) ([]view.UsedIpInventoryView, error) {
	var resp []view.UsedIpInventoryView
	return resp, cli.List(ctx, "v1/l3-networks/ip-address", params, &resp)
}

func (cli *ZSClient) GetIpAddress(ctx context.Context, uuid string) (*view.UsedIpInventoryView, error) {
	var resp view.UsedIpInventoryView
	if err := cli.Get(ctx, "v1/l3-networks/ip-address", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIpAddress Pagination
func (cli *ZSClient) PageIpAddress(ctx context.Context, params *param.QueryParam) ([]view.UsedIpInventoryView, int, error) {
	var ipAddress []view.UsedIpInventoryView
	total, err := cli.Page(ctx, "v1/l3-networks/ip-address", params, &ipAddress)
	return ipAddress, total, err
}

// CreateL2TfNetwork creates L2TfNetwork
func (cli *ZSClient) CreateL2TfNetwork(ctx context.Context, params param.CreateL2TfNetworkParam) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l2-networks/tf"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFirewallRuleTemplate deletes FirewallRuleTemplate
func (cli *ZSClient) DeleteFirewallRuleTemplate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vpcfirewalls/rules/templates", uuid, string(deleteMode))
}

// GetInterdependentL3NetworksImages gets InterdependentL3NetworksImages by uuid
func (cli *ZSClient) GetInterdependentL3NetworksImages(ctx context.Context) (*view.GetInterdependentL3NetworkImageView, error) {
	var resp view.GetInterdependentL3NetworkImageView
	if err := cli.GetWithRespKey(ctx, "v1/images-l3networks/dependencies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPciDeviceFromVm operates on PciDeviceFromVm
func (cli *ZSClient) DetachPciDeviceFromVm(ctx context.Context, pciDeviceUuid string, params param.DetachPciDeviceFromVmParam) (*view.PciDeviceInventoryView, error) {
	resp := view.PciDeviceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/pci-device/pci-devices/%s/detach", pciDeviceUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateVolumeSnapshotChain operates on VolumeSnapshotChain
func (cli *ZSClient) ValidateVolumeSnapshotChain(ctx context.Context, uuid string, params param.ValidateVolumeSnapshotChainParam) (*view.ValidateVolumeSnapshotChainEventView, error) {
	resp := view.ValidateVolumeSnapshotChainEventView{}
	if err := cli.PutWithSpec(ctx, "v1/volumes", uuid, "actions", "", map[string]interface{}{
		"validateVolumeSnapshotChain": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeHostNetworkInterfaceLldpMode changes HostNetworkInterfaceLldpMode
func (cli *ZSClient) ChangeHostNetworkInterfaceLldpMode(ctx context.Context, params param.ChangeHostNetworkInterfaceLldpModeParam) (*view.HostNetworkInterfaceLldpInventoryView, error) {
	resp := view.HostNetworkInterfaceLldpInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hostNetworkInterface/lldp/actions", "", "inventories", map[string]interface{}{
		"changeHostNetworkInterfaceLldpMode": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmNicsForLoadBalancerServerGroup gets CandidateVmNicsForLoadBalancerServerGroup by uuid
func (cli *ZSClient) GetCandidateVmNicsForLoadBalancerServerGroup(ctx context.Context) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/load-balancers/servergroups/candidate-nics", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachIscsiServerToCluster operates on IscsiServerToCluster
func (cli *ZSClient) AttachIscsiServerToCluster(ctx context.Context, clusterUuid, uuid string, params param.AttachIscsiServerToClusterParam) (*view.IscsiServerInventoryView, error) {
	resp := view.IscsiServerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/clusters/%s/storage-devices/iscsi/servers/%s", clusterUuid, uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachRoleToAccount operates on RoleToAccount
func (cli *ZSClient) AttachRoleToAccount(ctx context.Context, accountUuid, roleUuid string, params param.AttachRoleToAccountParam) (*view.AttachRoleToAccountEventView, error) {
	resp := view.AttachRoleToAccountEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/identities/accounts/%s/roles/%s", accountUuid, roleUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FlattenVolume operates on FlattenVolume
func (cli *ZSClient) FlattenVolume(ctx context.Context, uuid string, params param.FlattenVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volumes", uuid, "actions", "inventory", map[string]interface{}{
		"flattenVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachIsoToVmInstance operates on IsoToVmInstance
func (cli *ZSClient) AttachIsoToVmInstance(ctx context.Context, vmInstanceUuid, isoUuid string, params param.AttachIsoToVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/%s/iso/%s", vmInstanceUuid, isoUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVRouterRouterId operates on VRouterRouterId
func (cli *ZSClient) SetVRouterRouterId(ctx context.Context, vRouterUuid string, params param.SetVRouterRouterIdParam) (*view.SetVRouterRouterIdEventView, error) {
	resp := view.SetVRouterRouterIdEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/routerArea/%s/routerid", vRouterUuid), "routerId", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateAffinityGroupForAttachingVm gets CandidateAffinityGroupForAttachingVm by uuid
func (cli *ZSClient) GetCandidateAffinityGroupForAttachingVm(ctx context.Context) (*view.AffinityGroupInventoryView, error) {
	var resp view.AffinityGroupInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/affinityGroup/attachingVm", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFirewallIpSetTemplate updates FirewallIpSetTemplate
func (cli *ZSClient) UpdateFirewallIpSetTemplate(ctx context.Context, uuid string, params param.UpdateFirewallIpSetTemplateParam) (*view.VpcFirewallIpSetTemplateInventoryView, error) {
	resp := view.VpcFirewallIpSetTemplateInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vpcfirewalls/ipset/templates", uuid, "actions", "inventory", map[string]interface{}{
		"updateFirewallIpSetTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateLicense updates License
func (cli *ZSClient) UpdateLicense(ctx context.Context, managementNodeUuid string, params param.UpdateLicenseParam) (*view.UpdateLicenseEventView, error) {
	resp := view.UpdateLicenseEventView{}
	if err := cli.PutWithSpec(ctx, "v1/licenses/mn", managementNodeUuid, "actions", "", map[string]interface{}{
		"updateLicense": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAccessControlListRedirectRule adds AccessControlListRedirectRule
func (cli *ZSClient) AddAccessControlListRedirectRule(ctx context.Context, aclUuid string, params param.AddAccessControlListRedirectRuleParam) (*view.AccessControlListEntryInventoryView, error) {
	resp := view.AccessControlListEntryInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/access-control-lists/%s/redirectRules", aclUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCdpTaskData deletes CdpTaskData
func (cli *ZSClient) DeleteCdpTaskData(ctx context.Context, uuid string, params param.DeleteCdpTaskDataParam) (*view.DeleteCdpTaskDataEventView, error) {
	resp := view.DeleteCdpTaskDataEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cdp-task/%s/data", uuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachHostFromHostSchedulingRuleGroup operates on HostFromHostSchedulingRuleGroup
func (cli *ZSClient) DetachHostFromHostSchedulingRuleGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hostSchedulingRuleGroup", uuid, string(deleteMode))
}

// UninstallSoftwarePackage operates on UninstallSoftwarePackage
func (cli *ZSClient) UninstallSoftwarePackage(ctx context.Context, uuid string, params param.UninstallSoftwarePackageParam) (*view.UninstallSoftwarePackageEventView, error) {
	resp := view.UninstallSoftwarePackageEventView{}
	if err := cli.PutWithSpec(ctx, "v1/software-package", uuid, "actions", "", map[string]interface{}{
		"uninstallSoftwarePackage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTextTemplateArg gets TextTemplateArg by uuid
func (cli *ZSClient) GetTextTemplateArg(ctx context.Context) (*view.GetTextTemplateArgView, error) {
	var resp view.GetTextTemplateArgView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/textTemplateArg", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFirewall deletes Firewall
func (cli *ZSClient) DeleteFirewall(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vpcfirewalls", uuid, string(deleteMode))
}

// GetPciDeviceSpecCandidates gets PciDeviceSpecCandidates by uuid
func (cli *ZSClient) GetPciDeviceSpecCandidates(ctx context.Context) (*view.PciDeviceSpecInventoryView, error) {
	var resp view.PciDeviceSpecInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/pci-device-specs/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmCapabilities gets VmCapabilities by uuid
func (cli *ZSClient) GetVmCapabilities(ctx context.Context, uuid string) (*view.GetVmCapabilitiesView, error) {
	var resp view.GetVmCapabilitiesView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAccessKeyState changes AccessKeyState
func (cli *ZSClient) ChangeAccessKeyState(ctx context.Context, uuid string, params param.ChangeAccessKeyStateParam) (*view.AccessKeyInventoryView, error) {
	resp := view.AccessKeyInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/accesskeys", uuid, "actions", "inventory", map[string]interface{}{
		"changeAccessKeyState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachVmNicToVm operates on VmNicToVm
func (cli *ZSClient) AttachVmNicToVm(ctx context.Context, vmInstanceUuid, vmNicUuid string, params param.AttachVmNicToVmParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/%s/nices/%s", vmInstanceUuid, vmNicUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveMonFromCephBackupStorage removes MonFromCephBackupStorage
func (cli *ZSClient) RemoveMonFromCephBackupStorage(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/backup-storage/ceph", uuid, string(deleteMode))
}

// PrometheusQueryPassThrough operates on PrometheusQueryPassThrough
func (cli *ZSClient) PrometheusQueryPassThrough(ctx context.Context) (*view.PrometheusQueryPassThroughView, error) {
	var resp view.PrometheusQueryPassThroughView
	if err := cli.GetWithRespKey(ctx, "v1/prometheus/all", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmDeviceAddress gets VmDeviceAddress by uuid
func (cli *ZSClient) GetVmDeviceAddress(ctx context.Context) (*view.GetVmDeviceAddressView, error) {
	var resp view.GetVmDeviceAddressView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/devices", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveInstanceFromMonitorGroup removes InstanceFromMonitorGroup
func (cli *ZSClient) RemoveInstanceFromMonitorGroup(ctx context.Context, groupUuid string, instanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/zwatch/monitorgroups", groupUuid, fmt.Sprintf("actions/%s", instanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// CleanQueue operates on Queue
func (cli *ZSClient) CleanQueue(ctx context.Context, params param.CleanQueueParam) (*view.CleanQueueEventView, error) {
	resp := view.CleanQueueEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/clean/queue", "", "", map[string]interface{}{
		"cleanQueue": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveAccessControlListFromLoadBalancer removes AccessControlListFromLoadBalancer
func (cli *ZSClient) RemoveAccessControlListFromLoadBalancer(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/load-balancers/listeners", uuid, string(deleteMode))
}

// RemoveLabelFromEventSubscription removes LabelFromEventSubscription
func (cli *ZSClient) RemoveLabelFromEventSubscription(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/events/subscriptions/labels", uuid, string(deleteMode))
}

// ChangeInstanceOfferingState changes InstanceOfferingState
func (cli *ZSClient) ChangeInstanceOfferingState(ctx context.Context, uuid string, params param.ChangeInstanceOfferingStateParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/instance-offerings", uuid, "actions", "inventory", map[string]interface{}{
		"changeInstanceOfferingState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAccountGroupTree gets AccountGroupTree by uuid
func (cli *ZSClient) GetAccountGroupTree(ctx context.Context) (*view.AccountGroupViewView, error) {
	var resp view.GetAccountGroupTreeView
	if err := cli.GetWithRespKey(ctx, "v1/account-groups/tree", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetBackupStorageCapacity gets BackupStorageCapacity by uuid
func (cli *ZSClient) GetBackupStorageCapacity(ctx context.Context) (*view.GetBackupStorageCapacityView, error) {
	var resp view.GetBackupStorageCapacityView
	if err := cli.GetWithRespKey(ctx, "v1/backup-storage/capacities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateSeMdevDevices operates on SeMdevDevices
func (cli *ZSClient) GenerateSeMdevDevices(ctx context.Context, mttyDeviceUuid string, params param.GenerateSeMdevDevicesParam) (*view.GenerateSeMdevDevicesEventView, error) {
	resp := view.GenerateSeMdevDevicesEventView{}
	if err := cli.PutWithSpec(ctx, "v1/mtty-devices", mttyDeviceUuid, "actions", "", map[string]interface{}{
		"generateSeMdevDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetManagementNodeOS gets ManagementNodeOS by uuid
func (cli *ZSClient) GetManagementNodeOS(ctx context.Context) (*view.GetManagementNodeOSView, error) {
	resp := view.GetManagementNodeOSView{}
	if err := cli.PutWithRespKey(ctx, "v1/management/actions", "", "", map[string]interface{}{
		"getManagementNodeOS": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateMiniCluster creates MiniCluster
func (cli *ZSClient) CreateMiniCluster(ctx context.Context, params param.CreateMiniClusterParam) (*view.ClusterInventoryView, error) {
	resp := view.ClusterInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/mini-clusters"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncImageFromImageStoreBackupStorage operates on ImageFromImageStoreBackupStorage
func (cli *ZSClient) SyncImageFromImageStoreBackupStorage(ctx context.Context, uuid string, params param.SyncImageFromImageStoreBackupStorageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/images", uuid, "actions", "inventory", map[string]interface{}{
		"syncImageFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExecuteDRSScheduling operates on ExecuteDRSScheduling
func (cli *ZSClient) ExecuteDRSScheduling(ctx context.Context, uuid string, params param.ExecuteDRSSchedulingParam) (*view.ExecuteDRSSchedulingEventView, error) {
	resp := view.ExecuteDRSSchedulingEventView{}
	if err := cli.PutWithSpec(ctx, "v1/clusters/drs", uuid, "actions", "", map[string]interface{}{
		"executeDRSScheduling": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVipState changes VipState
func (cli *ZSClient) ChangeVipState(ctx context.Context, uuid string, params param.ChangeVipStateParam) (*view.VipInventoryView, error) {
	resp := view.VipInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vips", uuid, "actions", "inventory", map[string]interface{}{
		"changeVipState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UndoSnapshotCreation operates on UndoSnapshotCreation
func (cli *ZSClient) UndoSnapshotCreation(ctx context.Context, uuid string, params param.UndoSnapshotCreationParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volumes", uuid, "actions", "inventory", map[string]interface{}{
		"undoSnapshotCreation": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmQga gets VmQga by uuid
func (cli *ZSClient) GetVmQga(ctx context.Context, uuid string) (*view.GetVmQgaView, error) {
	var resp view.GetVmQgaView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmFromVolumeBackup creates VmFromVolumeBackup
func (cli *ZSClient) CreateVmFromVolumeBackup(ctx context.Context, backupUuid string, params param.CreateVmFromVolumeBackupParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/from/vm-backup/%s", backupUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PreviewResourceStack operates on PreviewResourceStack
func (cli *ZSClient) PreviewResourceStack(ctx context.Context, params param.PreviewResourceStackParam) (*view.PreviewResourceStackView, error) {
	resp := view.PreviewResourceStackView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cloudformation/stack/preview"), "preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmvNUMATopology gets VmvNUMATopology by uuid
func (cli *ZSClient) GetVmvNUMATopology(ctx context.Context, uuid string) (*view.GetVmvNUMATopologyView, error) {
	var resp view.GetVmvNUMATopologyView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveSchedulerJobsFromSchedulerJobGroup removes SchedulerJobsFromSchedulerJobGroup
func (cli *ZSClient) RemoveSchedulerJobsFromSchedulerJobGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/scheduler/jobgroups", uuid, string(deleteMode))
}

// GetManagementNodesStatus gets ManagementNodesStatus by uuid
func (cli *ZSClient) GetManagementNodesStatus(ctx context.Context) (*view.ManagementsStatusViewView, error) {
	var resp view.GetManagementNodesStatusView
	if err := cli.GetWithRespKey(ctx, "v1/management-nodes/status", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// GetHostPhysicalMemoryFacts gets HostPhysicalMemoryFacts by uuid
func (cli *ZSClient) GetHostPhysicalMemoryFacts(ctx context.Context, uuid string) (*view.HostPhysicalMemoryInventoryView, error) {
	var resp view.HostPhysicalMemoryInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/physical-memory-facts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseInfo gets LicenseInfo by uuid
func (cli *ZSClient) GetLicenseInfo(ctx context.Context) (*view.LicenseInventoryView, error) {
	var resp view.GetLicenseInfoView
	if err := cli.GetWithRespKey(ctx, "v1/licenses", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// ChangeSchedulerState changes SchedulerState
func (cli *ZSClient) ChangeSchedulerState(ctx context.Context, uuid string, params param.ChangeSchedulerStateParam) (*view.SchedulerJobInventoryView, error) {
	resp := view.SchedulerJobInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/schedulers", uuid, "", "inventory", map[string]interface{}{
		"changeSchedulerState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateMdevDevices operates on MdevDevices
func (cli *ZSClient) GenerateMdevDevices(ctx context.Context, pciDeviceUuid string, params param.GenerateMdevDevicesParam) (*view.GenerateVirtualPciDevicesEventView, error) {
	resp := view.GenerateVirtualPciDevicesEventView{}
	if err := cli.PutWithSpec(ctx, "v1/pci-devices", pciDeviceUuid, "actions", "", map[string]interface{}{
		"generateMdevDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPriceTableToAccount operates on PriceTableToAccount
func (cli *ZSClient) AttachPriceTableToAccount(ctx context.Context, tableUuid, accountUuid string, params param.AttachPriceTableToAccountParam) (*view.PriceTableInventoryView, error) {
	resp := view.PriceTableInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/billings/price-tables/%s/accounts/%s", tableUuid, accountUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUploadSoftwarePackageJobDetails gets UploadSoftwarePackageJobDetails by uuid
func (cli *ZSClient) GetUploadSoftwarePackageJobDetails(ctx context.Context, uuid string) (*view.GetUploadSoftwarePackageJobDetailsView, error) {
	var resp view.GetUploadSoftwarePackageJobDetailsView
	if err := cli.GetWithRespKey(ctx, "v1/software-package/upload-jobs/details", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSNSTopicState changes SNSTopicState
func (cli *ZSClient) ChangeSNSTopicState(ctx context.Context, uuid string, params param.ChangeSNSTopicStateParam) (*view.SNSTopicInventoryView, error) {
	resp := view.SNSTopicInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/topics", uuid, "actions", "inventory", map[string]interface{}{
		"changeSNSTopicState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachScsiLunToVmInstance operates on ScsiLunToVmInstance
func (cli *ZSClient) AttachScsiLunToVmInstance(ctx context.Context, vmInstanceUuid, uuid string, params param.AttachScsiLunToVmInstanceParam) (*view.ScsiLunInventoryView, error) {
	resp := view.ScsiLunInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/%s/scsi-lun/%s", vmInstanceUuid, uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmUptime gets VmUptime by uuid
func (cli *ZSClient) GetVmUptime(ctx context.Context, uuid string) (*view.GetVmUptimeView, error) {
	var resp view.GetVmUptimeView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveRemoteCidrsFromIPsecConnection removes RemoteCidrsFromIPsecConnection
func (cli *ZSClient) RemoveRemoteCidrsFromIPsecConnection(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/ipsec", uuid, string(deleteMode))
}

// AddMonToCephPrimaryStorage adds MonToCephPrimaryStorage
func (cli *ZSClient) AddMonToCephPrimaryStorage(ctx context.Context, uuid string, params param.AddMonToCephPrimaryStorageParam) (*view.CephPrimaryStorageInventoryView, error) {
	resp := view.CephPrimaryStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/primary-storage/ceph/%s/mons", uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryLocalRaidPhysicalDrive queries LocalRaidPhysicalDrive list
func (cli *ZSClient) QueryLocalRaidPhysicalDrive(ctx context.Context, params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, error) {
	var resp []view.RaidPhysicalDriveInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/local-raid/physical-drives", params, &resp)
}

func (cli *ZSClient) GetLocalRaidPhysicalDrive(ctx context.Context, uuid string) (*view.RaidPhysicalDriveInventoryView, error) {
	var resp view.RaidPhysicalDriveInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/local-raid/physical-drives", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLocalRaidPhysicalDrive Pagination
func (cli *ZSClient) PageLocalRaidPhysicalDrive(ctx context.Context, params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, int, error) {
	var localRaidPhysicalDrives []view.RaidPhysicalDriveInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/local-raid/physical-drives", params, &localRaidPhysicalDrives)
	return localRaidPhysicalDrives, total, err
}

// RemoveHostRouteFromL3Network removes HostRouteFromL3Network
func (cli *ZSClient) RemoveHostRouteFromL3Network(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l3-networks", uuid, string(deleteMode))
}

// BackupStorageMigrateImage operates on StorageMigrateImage
func (cli *ZSClient) BackupStorageMigrateImage(ctx context.Context, imageUuid string, params param.BackupStorageMigrateImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage/images", imageUuid, "actions", "inventory", map[string]interface{}{
		"backupStorageMigrateImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVmToAffinityGroup adds VmToAffinityGroup
func (cli *ZSClient) AddVmToAffinityGroup(ctx context.Context, affinityGroupUuid, uuid string, params param.AddVmToAffinityGroupParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/affinity-groups/%s/vm-instances/%s", affinityGroupUuid, uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageAllocatorStrategies gets PrimaryStorageAllocatorStrategies by uuid
func (cli *ZSClient) GetPrimaryStorageAllocatorStrategies(ctx context.Context) (*view.GetPrimaryStorageAllocatorStrategiesView, error) {
	var resp view.GetPrimaryStorageAllocatorStrategiesView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/allocators/strategies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveBackendServerFromServerGroup removes BackendServerFromServerGroup
func (cli *ZSClient) RemoveBackendServerFromServerGroup(ctx context.Context, serverGroupUuid string, params param.RemoveBackendServerFromServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/load-balancers/servergroups", serverGroupUuid, "backendservers/actions", "inventory", map[string]interface{}{
		"removeBackendServerFromServerGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVirtualSwitchUplinkBondings updates VirtualSwitchUplinkBondings
func (cli *ZSClient) UpdateVirtualSwitchUplinkBondings(ctx context.Context, uuid string, params param.UpdateVirtualSwitchUplinkBondingsParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/l2-networks/virtual-switch", uuid, "uplink-bondings", "inventories", map[string]interface{}{
		"updateVirtualSwitchUplinkBondings": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPlatformTimeZone gets PlatformTimeZone by uuid
func (cli *ZSClient) GetPlatformTimeZone(ctx context.Context) (*view.GetPlatformTimeZoneView, error) {
	var resp view.GetPlatformTimeZoneView
	if err := cli.GetWithRespKey(ctx, "v1/management-nodes/platform-timezone", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedVip gets VpcAttachedVip by uuid
func (cli *ZSClient) GetVpcAttachedVip(ctx context.Context, uuid string, params param.GetVpcAttachedVipParam) (*view.VipInventoryView, error) {
	resp := view.VipInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/virtual-routers/%s/attached-vip", uuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIpv6Range adds Ipv6Range
func (cli *ZSClient) AddIpv6Range(ctx context.Context, l3NetworkUuid string, params param.AddIpv6RangeParam) (*view.IpRangeInventoryView, error) {
	resp := view.IpRangeInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks/%s/ipv6-ranges", l3NetworkUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmInstanceDefaultCdRom operates on VmInstanceDefaultCdRom
func (cli *ZSClient) SetVmInstanceDefaultCdRom(ctx context.Context, vmInstanceUuid string, uuid string, params param.SetVmInstanceDefaultCdRomParam) (*view.VmCdRomInventoryView, error) {
	resp := view.VmCdRomInventoryView{}
	err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, fmt.Sprintf("cdroms/%s/actions", uuid), "inventory", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshSharedblockDeviceCapacity operates on SharedblockDeviceCapacity
func (cli *ZSClient) RefreshSharedblockDeviceCapacity(ctx context.Context, sharedBlockGroupUuid, uuid string, params param.RefreshSharedblockDeviceCapacityParam) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	resp := view.SharedBlockGroupPrimaryStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/primary-storage/sharedblockgroup/%s/sharedblocks/%s", sharedBlockGroupUuid, uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FstrimVm operates on FstrimVm
func (cli *ZSClient) FstrimVm(ctx context.Context, uuid string, params param.FstrimVmParam) (*view.FstrimVmEventView, error) {
	resp := view.FstrimVmEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/%s/actions", uuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckBaremetalChassisConfigFile operates on BaremetalChassisConfigFile
func (cli *ZSClient) CheckBaremetalChassisConfigFile(ctx context.Context, params param.CheckBaremetalChassisConfigFileParam) (*view.CheckBaremetalChassisConfigFileView, error) {
	resp := view.CheckBaremetalChassisConfigFileView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/baremetal/chassis/from-file/check"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachL2NetworkFromCluster operates on L2NetworkFromCluster
func (cli *ZSClient) DetachL2NetworkFromCluster(ctx context.Context, l2NetworkUuid string, clusterUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/l2-networks", l2NetworkUuid, fmt.Sprintf("clusters/%s", clusterUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// ChangeMulticastRouterState changes MulticastRouterState
func (cli *ZSClient) ChangeMulticastRouterState(ctx context.Context, uuid string, params param.ChangeMulticastRouterStateParam) (*view.MulticastRouterInventoryView, error) {
	resp := view.MulticastRouterInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/multicast/virtual-routers", uuid, "actions", "inventory", map[string]interface{}{
		"changeMulticastRouterState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFreeIp gets FreeIp by uuid
func (cli *ZSClient) GetFreeIp(ctx context.Context) (*view.FreeIpInventoryView, error) {
	var resp view.FreeIpInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/l3-networks/ip/free", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (cli *ZSClient) GetFreeIpOfL3Network(ctx context.Context, uuid string) ([]view.FreeIpInventoryView, error) {
	var resp []view.FreeIpInventoryView
	if err := cli.GetWithSpec(ctx, "v1/l3-networks", uuid, "ip/free", "inventories", nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (cli *ZSClient) GetFreeIpOfIpRange(ctx context.Context, uuid string) ([]view.FreeIpInventoryView, error) {
	var resp []view.FreeIpInventoryView
	if err := cli.GetWithSpec(ctx, "v1/l3-networks/ip-ranges", uuid, "ip/free", "inventories", nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// CheckCephHealthStatus operates on CephHealthStatus
func (cli *ZSClient) CheckCephHealthStatus(ctx context.Context) (*view.CheckCephHealthStatusView, error) {
	resp := view.CheckCephHealthStatusView{}
	if err := cli.Post(ctx, "v1/zops/check-ceph-health", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmNicState changes VmNicState
func (cli *ZSClient) ChangeVmNicState(ctx context.Context, vmNicUuid string, params param.ChangeVmNicStateParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances/nics", vmNicUuid, "actions", "inventory", map[string]interface{}{
		"changeVmNicState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateL2PortGroup creates L2PortGroup
func (cli *ZSClient) CreateL2PortGroup(ctx context.Context, params param.CreateL2PortGroupParam) (*view.CreateL2PortGroupEventView, error) {
	resp := view.CreateL2PortGroupEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l2-networks/port-group"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateInstanceOfferingUserConfig operates on InstanceOfferingUserConfig
func (cli *ZSClient) ValidateInstanceOfferingUserConfig(ctx context.Context, params param.ValidateInstanceOfferingUserConfigParam) (*view.ValidateInstanceOfferingUserConfigEventView, error) {
	resp := view.ValidateInstanceOfferingUserConfigEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/billings/accounts/actions", "", "", map[string]interface{}{
		"validateInstanceOfferingUserConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryTag queries Tag list
func (cli *ZSClient) QueryTag(ctx context.Context, params *param.QueryParam) ([]view.TagPatternInventoryView, error) {
	var resp []view.TagPatternInventoryView
	return resp, cli.List(ctx, "v1/tags", params, &resp)
}

func (cli *ZSClient) GetTag(ctx context.Context, uuid string) (*view.TagPatternInventoryView, error) {
	var resp view.TagPatternInventoryView
	if err := cli.Get(ctx, "v1/tags", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTag Pagination
func (cli *ZSClient) PageTag(ctx context.Context, params *param.QueryParam) ([]view.TagPatternInventoryView, int, error) {
	var tags []view.TagPatternInventoryView
	total, err := cli.Page(ctx, "v1/tags", params, &tags)
	return tags, total, err
}

// UnprotectVmInstanceRecoveryPoint operates on UnprotectVmInstanceRecoveryPoint
func (cli *ZSClient) UnprotectVmInstanceRecoveryPoint(ctx context.Context, vmInstanceUuid string, params param.UnprotectVmInstanceRecoveryPointParam) (*view.UnprotectVmInstanceRecoveryPointEventView, error) {
	resp := view.UnprotectVmInstanceRecoveryPointEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "unprotect-recovery-point", "", map[string]interface{}{
		"unprotectVmInstanceRecoveryPoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// TriggerGCJob operates on TriggerGCJob
func (cli *ZSClient) TriggerGCJob(ctx context.Context, uuid string, params param.TriggerGCJobParam) (*view.TriggerGCJobEventView, error) {
	resp := view.TriggerGCJobEventView{}
	if err := cli.PutWithSpec(ctx, "v1/gc-jobs", uuid, "actions", "", map[string]interface{}{
		"triggerGCJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmHostname operates on VmHostname
func (cli *ZSClient) SetVmHostname(ctx context.Context, uuid string, params param.SetVmHostnameParam) (*view.SetVmHostnameEventView, error) {
	resp := view.SetVmHostnameEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"setVmHostname": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ApplyRuleSetChanges operates on RuleSetChanges
func (cli *ZSClient) ApplyRuleSetChanges(ctx context.Context, uuid string, params param.ApplyRuleSetChangesParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	resp := view.VpcFirewallRuleSetInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vpcfirewalls/ruleSets/apply", uuid, "actions", "inventory", map[string]interface{}{
		"applyRuleSetChanges": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrimaryStorageMigrateVm operates on PrimaryStorageMigrateVm
func (cli *ZSClient) PrimaryStorageMigrateVm(ctx context.Context, vmInstanceUuid string, params param.PrimaryStorageMigrateVmParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "actions", "inventory", map[string]interface{}{
		"primaryStorageMigrateVm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoverDatabaseFromBackup operates on DatabaseFromBackup
func (cli *ZSClient) RecoverDatabaseFromBackup(ctx context.Context, params param.RecoverDatabaseFromBackupParam) (*view.RecoverDatabaseFromBackupEventView, error) {
	resp := view.RecoverDatabaseFromBackupEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/database-backups/actions", "", "", map[string]interface{}{
		"recoverDatabaseFromBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UngenerateMdevDevices operates on UngenerateMdevDevices
func (cli *ZSClient) UngenerateMdevDevices(ctx context.Context, pciDeviceUuid string, params param.UngenerateMdevDevicesParam) (*view.UngenerateVirtualPciDevicesEventView, error) {
	resp := view.UngenerateVirtualPciDevicesEventView{}
	if err := cli.PutWithSpec(ctx, "v1/pci-devices", pciDeviceUuid, "actions", "", map[string]interface{}{
		"ungenerateMdevDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSimulatorPrimaryStorage adds SimulatorPrimaryStorage
func (cli *ZSClient) AddSimulatorPrimaryStorage(ctx context.Context, params param.AddSimulatorPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/primary-storage/simulators"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MoveDirectory operates on MoveDirectory
func (cli *ZSClient) MoveDirectory(ctx context.Context, params param.MoveDirectoryParam) (*view.MoveDirectoryEventView, error) {
	resp := view.MoveDirectoryEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/move/directory", "", "", map[string]interface{}{
		"moveDirectory": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachVRouterRouteTableFromVRouter operates on VRouterRouteTableFromVRouter
func (cli *ZSClient) DetachVRouterRouteTableFromVRouter(ctx context.Context, routeTableUuid string, virtualRouterVmUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/vrouter-route-tables", routeTableUuid, fmt.Sprintf("detach/%s", virtualRouterVmUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetVRouterOspfNeighbor gets VRouterOspfNeighbor by uuid
func (cli *ZSClient) GetVRouterOspfNeighbor(ctx context.Context, uuid string) (*view.GetVRouterOspfNeighborView, error) {
	var resp view.GetVRouterOspfNeighborView
	if err := cli.GetWithRespKey(ctx, "v1/routerArea", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVipUsedPorts gets VipUsedPorts by uuid
func (cli *ZSClient) GetVipUsedPorts(ctx context.Context, uuid string) (*view.VipPortRangeInventoryView, error) {
	var resp view.VipPortRangeInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vips", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmConsolePassword operates on VmConsolePassword
func (cli *ZSClient) SetVmConsolePassword(ctx context.Context, uuid string, params param.SetVmConsolePasswordParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "inventory", map[string]interface{}{
		"setVmConsolePassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVpcVRouter creates VpcVRouter
func (cli *ZSClient) CreateVpcVRouter(ctx context.Context, params param.CreateVpcVRouterParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/virtual-routers"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachFirewallRuleSetToL3 operates on FirewallRuleSetToL3
func (cli *ZSClient) AttachFirewallRuleSetToL3(ctx context.Context, ruleSetUuid, l3Uuid string, params param.AttachFirewallRuleSetToL3Param) (*view.VpcFirewallRuleSetL3RefInventoryView, error) {
	resp := view.VpcFirewallRuleSetL3RefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpcfirewalls/ruleSets/%s/l3networks/%s", ruleSetUuid, l3Uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpStorageTrashOnPrimaryStorage operates on UpStorageTrashOnPrimaryStorage
func (cli *ZSClient) CleanUpStorageTrashOnPrimaryStorage(ctx context.Context, uuid string, params param.CleanUpStorageTrashOnPrimaryStorageParam) (*view.CleanUpStorageTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpStorageTrashOnPrimaryStorageEventView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage", uuid, "storagetrash/actions", "", map[string]interface{}{
		"cleanUpStorageTrashOnPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMdevDeviceSpecCandidates gets MdevDeviceSpecCandidates by uuid
func (cli *ZSClient) GetMdevDeviceSpecCandidates(ctx context.Context) (*view.MdevDeviceSpecInventoryView, error) {
	var resp view.MdevDeviceSpecInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/mdev-device-specs/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFlowMeterRouterId gets FlowMeterRouterId by uuid
func (cli *ZSClient) GetFlowMeterRouterId(ctx context.Context, uuid string) (*view.GetFlowMeterRouterIdView, error) {
	var resp view.GetFlowMeterRouterIdView
	if err := cli.GetWithRespKey(ctx, "v1/flowmeters", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPciDeviceCandidatesForNewCreateVm gets PciDeviceCandidatesForNewCreateVm by uuid
func (cli *ZSClient) GetPciDeviceCandidatesForNewCreateVm(ctx context.Context) (*view.PciDeviceInventoryView, error) {
	var resp view.PciDeviceInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/pci-device/candidate-pci-devices-for-new-create-vm", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVRouterFlowMeterNetwork queries VRouterFlowMeterNetwork list
func (cli *ZSClient) QueryVRouterFlowMeterNetwork(ctx context.Context, params *param.QueryParam) ([]view.NetworkRouterFlowMeterRefInventoryView, error) {
	var resp []view.NetworkRouterFlowMeterRefInventoryView
	return resp, cli.List(ctx, "v1/flowmeters/networks", params, &resp)
}

func (cli *ZSClient) GetVRouterFlowMeterNetwork(ctx context.Context, uuid string) (*view.NetworkRouterFlowMeterRefInventoryView, error) {
	var resp view.NetworkRouterFlowMeterRefInventoryView
	if err := cli.Get(ctx, "v1/flowmeters/networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVRouterFlowMeterNetwork Pagination
func (cli *ZSClient) PageVRouterFlowMeterNetwork(ctx context.Context, params *param.QueryParam) ([]view.NetworkRouterFlowMeterRefInventoryView, int, error) {
	var vRouterFlowMeterNetworks []view.NetworkRouterFlowMeterRefInventoryView
	total, err := cli.Page(ctx, "v1/flowmeters/networks", params, &vRouterFlowMeterNetworks)
	return vRouterFlowMeterNetworks, total, err
}

// GetHostTask gets HostTask by uuid
func (cli *ZSClient) GetHostTask(ctx context.Context) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/task-details", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetManagementNodeDirCapacity gets ManagementNodeDirCapacity by uuid
func (cli *ZSClient) GetManagementNodeDirCapacity(ctx context.Context) (*view.GetManagementNodeDirCapacityView, error) {
	var resp view.GetManagementNodeDirCapacityView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/mn", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallRuleSetL3Ref queries FirewallRuleSetL3Ref list
func (cli *ZSClient) QueryFirewallRuleSetL3Ref(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleSetL3RefInventoryView, error) {
	var resp []view.VpcFirewallRuleSetL3RefInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls/l3networks/rulesets/refs", params, &resp)
}

// PageFirewallRuleSetL3Ref Pagination
func (cli *ZSClient) PageFirewallRuleSetL3Ref(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleSetL3RefInventoryView, int, error) {
	var firewallRuleSetL3Refs []view.VpcFirewallRuleSetL3RefInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls/l3networks/rulesets/refs", params, &firewallRuleSetL3Refs)
	return firewallRuleSetL3Refs, total, err
}

// GetAlarmData gets AlarmData by uuid
func (cli *ZSClient) GetAlarmData(ctx context.Context) (*view.GetAlarmDataView, error) {
	var resp view.GetAlarmDataView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/alarm-histories", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UngroupVolumeSnapshotGroup operates on UngroupVolumeSnapshotGroup
func (cli *ZSClient) UngroupVolumeSnapshotGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/volume-snapshots/ungroup", uuid, string(deleteMode))
}

// SubscribeSNSTopic operates on SubscribeSNSTopic
func (cli *ZSClient) SubscribeSNSTopic(ctx context.Context, topicUuid, endpointUuid string, params param.SubscribeSNSTopicParam) (*view.SubscribeSNSTopicEventView, error) {
	resp := view.SubscribeSNSTopicEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/topics/%s/endpoints/%s", topicUuid, endpointUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmNicForSecurityGroup gets CandidateVmNicForSecurityGroup by uuid
func (cli *ZSClient) GetCandidateVmNicForSecurityGroup(ctx context.Context, uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/security-groups", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmRDP gets VmRDP by uuid
func (cli *ZSClient) GetVmRDP(ctx context.Context, uuid string) (*view.GetVmRDPView, error) {
	var resp view.GetVmRDPView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPciDeviceToVm operates on PciDeviceToVm
func (cli *ZSClient) AttachPciDeviceToVm(ctx context.Context, pciDeviceUuid string, params param.AttachPciDeviceToVmParam) (*view.PciDeviceInventoryView, error) {
	resp := view.PciDeviceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/pci-device/pci-devices/%s/attach", pciDeviceUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanupBillingUsage operates on upBillingUsage
func (cli *ZSClient) CleanupBillingUsage(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/billings/usage", uuid, string(deleteMode))
}

// GetCandidateL2NetworksForAttachingCluster gets CandidateL2NetworksForAttachingCluster by uuid
func (cli *ZSClient) GetCandidateL2NetworksForAttachingCluster(ctx context.Context, uuid string) (*view.L2NetworkDataView, error) {
	var resp view.L2NetworkDataView
	if err := cli.GetWithRespKey(ctx, "v1/cluster", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsOpensourceVersion operates on IsOpensourceVersion
func (cli *ZSClient) IsOpensourceVersion(ctx context.Context) (*view.IsOpensourceVersionView, error) {
	var resp view.IsOpensourceVersionView
	if err := cli.GetWithRespKey(ctx, "v1/meta-data/opensource", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ConvertVmInstanceToTemplatedVmInstance operates on ConvertVmInstanceToTemplatedVmInstance
func (cli *ZSClient) ConvertVmInstanceToTemplatedVmInstance(ctx context.Context, vmInstanceUuid string, params param.ConvertVmInstanceToTemplatedVmInstanceParam) (*view.TemplatedVmInstanceInventoryView, error) {
	resp := view.TemplatedVmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/%s/convert-to-templatedVmInstance", vmInstanceUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsVfNicAvailableInL3Network operates on IsVfNicAvailableInL3Network
func (cli *ZSClient) IsVfNicAvailableInL3Network(ctx context.Context, l3NetworkUuid string, hostUuid string) (*view.IsVfNicAvailableInL3NetworkView, error) {
	var resp view.IsVfNicAvailableInL3NetworkView
	err := cli.GetWithSpec(ctx, "v1/l3-networks", l3NetworkUuid, fmt.Sprintf("hosts/%s/vfnicavailable", hostUuid), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAllMetricMetadata gets AllMetricMetadata by uuid
func (cli *ZSClient) GetAllMetricMetadata(ctx context.Context) (*view.GetAllMetricMetadataView, error) {
	var resp view.GetAllMetricMetadataView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/metrics/meta-data", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceFromResourceStack gets ResourceFromResourceStack by uuid
func (cli *ZSClient) GetResourceFromResourceStack(ctx context.Context) (*view.GetResourceFromResourceStackView, error) {
	var resp view.GetResourceFromResourceStackView
	if err := cli.GetWithRespKey(ctx, "v1/cloudformation/stack/resources", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MoveResourcesToDirectory operates on MoveResourcesToDirectory
func (cli *ZSClient) MoveResourcesToDirectory(ctx context.Context, params param.MoveResourcesToDirectoryParam) (*view.MoveResourcesToDirectoryEventView, error) {
	resp := view.MoveResourcesToDirectoryEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/move/resources/directory", "", "", map[string]interface{}{
		"moveResourcesToDirectory": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSupportedCloudFormationResources gets SupportedCloudFormationResources by uuid
func (cli *ZSClient) GetSupportedCloudFormationResources(ctx context.Context) (*view.GetSupportedCloudFormationResourcesView, error) {
	var resp view.GetSupportedCloudFormationResourcesView
	if err := cli.GetWithRespKey(ctx, "v1/cloudformation/resources", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVmBackup operates on VmBackup
func (cli *ZSClient) SyncVmBackup(ctx context.Context, imageStoreUuid string, params param.SyncVmBackupParam) (*view.SyncVmBackupEventView, error) {
	resp := view.SyncVmBackupEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-backups/imageStore", imageStoreUuid, "actions", "result", map[string]interface{}{
		"syncVmBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// InstallSoftwarePackage operates on InstallSoftwarePackage
func (cli *ZSClient) InstallSoftwarePackage(ctx context.Context, uuid string, params param.InstallSoftwarePackageParam) (*view.InstallSoftwarePackageEventView, error) {
	resp := view.InstallSoftwarePackageEventView{}
	if err := cli.PutWithSpec(ctx, "v1/software-package/install", uuid, "actions", "", map[string]interface{}{
		"installSoftwarePackage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVRouterOspfNetwork queries VRouterOspfNetwork list
func (cli *ZSClient) QueryVRouterOspfNetwork(ctx context.Context, params *param.QueryParam) ([]view.NetworkRouterAreaRefInventoryView, error) {
	var resp []view.NetworkRouterAreaRefInventoryView
	return resp, cli.List(ctx, "v1/routerArea/network", params, &resp)
}

func (cli *ZSClient) GetVRouterOspfNetwork(ctx context.Context, uuid string) (*view.NetworkRouterAreaRefInventoryView, error) {
	var resp view.NetworkRouterAreaRefInventoryView
	if err := cli.Get(ctx, "v1/routerArea/networkR", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVRouterOspfNetwork Pagination
func (cli *ZSClient) PageVRouterOspfNetwork(ctx context.Context, params *param.QueryParam) ([]view.NetworkRouterAreaRefInventoryView, int, error) {
	var vRouterOspfNetworks []view.NetworkRouterAreaRefInventoryView
	total, err := cli.Page(ctx, "v1/routerArea/network", params, &vRouterOspfNetworks)
	return vRouterOspfNetworks, total, err
}

// CreateTag creates Tag
func (cli *ZSClient) CreateTag(ctx context.Context, params param.CreateTagParam) (*view.TagPatternInventoryView, error) {
	resp := view.TagPatternInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/tags"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromVolumeSnapshotGroup creates VmInstanceFromVolumeSnapshotGroup
func (cli *ZSClient) CreateVmInstanceFromVolumeSnapshotGroup(ctx context.Context, volumeSnapshotGroupUuid string, params param.CreateVmInstanceFromVolumeSnapshotGroupParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/from/volume-snapshots/group/%s", volumeSnapshotGroupUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ZceXTestConnection operates on ZceXTestConnection
func (cli *ZSClient) ZceXTestConnection(ctx context.Context, params param.ZceXTestConnectionParam) (*view.ZceXTestConnectionView, error) {
	resp := view.ZceXTestConnectionView{}
	if err := cli.PutWithRespKey(ctx, "v1/zce-x-plugin/test-connection", "", "", map[string]interface{}{
		"zceXTestConnection": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBaremetalChassisState changes BaremetalChassisState
func (cli *ZSClient) ChangeBaremetalChassisState(ctx context.Context, uuid string, params param.ChangeBaremetalChassisStateParam) (*view.BaremetalChassisInventoryView, error) {
	resp := view.BaremetalChassisInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/chassis", uuid, "actions", "inventory", map[string]interface{}{
		"changeBaremetalChassisState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAttachableVpcL3Network gets AttachableVpcL3Network by uuid
func (cli *ZSClient) GetAttachableVpcL3Network(ctx context.Context, uuid string, params param.GetAttachableVpcL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/virtual-routers/%s/attachable-vpc-l3s", uuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachSshKeyPairToVmInstance operates on SshKeyPairToVmInstance
func (cli *ZSClient) AttachSshKeyPairToVmInstance(ctx context.Context, sshKeyPairUuid, vmInstanceUuid string, params param.AttachSshKeyPairToVmInstanceParam) (*view.SshKeyPairInventoryView, error) {
	resp := view.SshKeyPairInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/ssh-key-pair/%s/vm-instance/%s", sshKeyPairUuid, vmInstanceUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEmailTriggerAction queries EmailTrigger list
func (cli *ZSClient) QueryEmailTriggerAction(ctx context.Context, params *param.QueryParam) ([]view.MonitorTriggerActionInventoryView, error) {
	var resp []view.MonitorTriggerActionInventoryView
	return resp, cli.List(ctx, "v1/monitoring/trigger-actions/emails", params, &resp)
}

func (cli *ZSClient) GetEmailTriggerAction(ctx context.Context, uuid string) (*view.MonitorTriggerActionInventoryView, error) {
	var resp view.MonitorTriggerActionInventoryView
	if err := cli.Get(ctx, "v1/monitoring/trigger-actions/emails", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEmailTriggerAction Pagination
func (cli *ZSClient) PageEmailTriggerAction(ctx context.Context, params *param.QueryParam) ([]view.MonitorTriggerActionInventoryView, int, error) {
	var emailTriggers []view.MonitorTriggerActionInventoryView
	total, err := cli.Page(ctx, "v1/monitoring/trigger-actions/emails", params, &emailTriggers)
	return emailTriggers, total, err
}

// ReloadElaboration operates on ReloadElaboration
func (cli *ZSClient) ReloadElaboration(ctx context.Context) (*view.ReloadElaborationEventView, error) {
	resp := view.ReloadElaborationEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/errorcode/actions", "", "", map[string]interface{}{
		"reloadElaboration": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkMtu gets L3NetworkMtu by uuid
func (cli *ZSClient) GetL3NetworkMtu(ctx context.Context, uuid string) (*view.GetL3NetworkMtuView, error) {
	var resp view.GetL3NetworkMtuView
	if err := cli.GetWithRespKey(ctx, "v1/l3-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReconnectVirtualRouter operates on VirtualRouter
func (cli *ZSClient) ReconnectVirtualRouter(ctx context.Context, vmInstanceUuid string, params param.ReconnectVirtualRouterParam) (*view.ApplianceVmInventoryView, error) {
	resp := view.ApplianceVmInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances/appliances/virtual-routers", vmInstanceUuid, "actions", "inventory", map[string]interface{}{
		"reconnectVirtualRouter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSecurityGroupRulePriority updates SecurityGroupRulePriority
func (cli *ZSClient) UpdateSecurityGroupRulePriority(ctx context.Context, securityGroupUuid string, params param.UpdateSecurityGroupRulePriorityParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/security-groups", securityGroupUuid, "rules/priority/actions", "inventory", map[string]interface{}{
		"updateSecurityGroupRulePriority": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceInAccountGroup gets ResourceInAccountGroup by uuid
func (cli *ZSClient) GetResourceInAccountGroup(ctx context.Context, uuid string) (*view.GetResourceInAccountGroupView, error) {
	var resp view.GetResourceInAccountGroupView
	if err := cli.GetWithRespKey(ctx, "v1/account-groups", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddDnsToL3Network adds DnsToL3Network
func (cli *ZSClient) AddDnsToL3Network(ctx context.Context, l3NetworkUuid string, params param.AddDnsToL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks/%s/dns", l3NetworkUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryPortMirrorNetworkUsedIp queries PortMirrorNetworkUsedIp list
func (cli *ZSClient) QueryPortMirrorNetworkUsedIp(ctx context.Context, params *param.QueryParam) ([]view.MirrorNetworkUsedIpInventoryView, error) {
	var resp []view.MirrorNetworkUsedIpInventoryView
	return resp, cli.List(ctx, "v1/port-mirrors/networks/usedIps", params, &resp)
}

func (cli *ZSClient) GetPortMirrorNetworkUsedIp(ctx context.Context, uuid string) (*view.MirrorNetworkUsedIpInventoryView, error) {
	var resp view.MirrorNetworkUsedIpInventoryView
	if err := cli.Get(ctx, "v1/port-mirrors/networks/usedIps", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePortMirrorNetworkUsedIp Pagination
func (cli *ZSClient) PagePortMirrorNetworkUsedIp(ctx context.Context, params *param.QueryParam) ([]view.MirrorNetworkUsedIpInventoryView, int, error) {
	var portMirrorNetworkUsedIps []view.MirrorNetworkUsedIpInventoryView
	total, err := cli.Page(ctx, "v1/port-mirrors/networks/usedIps", params, &portMirrorNetworkUsedIps)
	return portMirrorNetworkUsedIps, total, err
}

// SetVmMonitorNumber operates on VmMonitorNumber
func (cli *ZSClient) SetVmMonitorNumber(ctx context.Context, uuid string, params param.SetVmMonitorNumberParam) (*view.SetVmMonitorNumberEventView, error) {
	resp := view.SetVmMonitorNumberEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"setVmMonitorNumber": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeLoadBalancerBackendServer changes LoadBalancerBackendServer
func (cli *ZSClient) ChangeLoadBalancerBackendServer(ctx context.Context, serverGroupUuid string, params param.ChangeLoadBalancerBackendServerParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/load-balancers/servergroups", serverGroupUuid, "backendserver/actions", "inventory", map[string]interface{}{
		"changeLoadBalancerBackendServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RestartResourceStack operates on RestartResourceStack
func (cli *ZSClient) RestartResourceStack(ctx context.Context, uuid string, params param.RestartResourceStackParam) (*view.ResourceStackInventoryView, error) {
	resp := view.ResourceStackInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/cloudformation/stack", uuid, "actions", "inventory", map[string]interface{}{
		"restartResourceStack": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmMigrationCandidateHosts gets VmMigrationCandidateHosts by uuid
func (cli *ZSClient) GetVmMigrationCandidateHosts(ctx context.Context, uuid string) (*view.HostInventoryView, error) {
	var resp view.HostInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL3NetworksForIpSecConnection gets CandidateL3NetworksForIpSecConnection by uuid
func (cli *ZSClient) GetCandidateL3NetworksForIpSecConnection(ctx context.Context) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/ipsec/candidatesL3Networks", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachBackupStorageToZone operates on BackupStorageToZone
func (cli *ZSClient) AttachBackupStorageToZone(ctx context.Context, zoneUuid, backupStorageUuid string, params param.AttachBackupStorageToZoneParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zones/%s/backup-storage/%s", zoneUuid, backupStorageUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddPciDeviceSpecToVmInstance adds PciDeviceSpecToVmInstance
func (cli *ZSClient) AddPciDeviceSpecToVmInstance(ctx context.Context, pciSpecUuid, vmInstanceUuid string, params param.AddPciDeviceSpecToVmInstanceParam) (*view.VmInstancePciDeviceSpecRefInventoryView, error) {
	resp := view.VmInstancePciDeviceSpecRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/pci-device-specs/%s/vm-instances/%s", pciSpecUuid, vmInstanceUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResizeRootVolume operates on RootVolume
func (cli *ZSClient) ResizeRootVolume(ctx context.Context, uuid string, params param.ResizeRootVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volumes/resize", uuid, "actions", "inventory", map[string]interface{}{
		"resizeRootVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SNSMicrosoftTeamsTestConnection operates on MicrosoftTeamsTestConnection
func (cli *ZSClient) SNSMicrosoftTeamsTestConnection(ctx context.Context, params param.SNSMicrosoftTeamsTestConnectionParam) (*view.SNSMicrosoftTeamsTestConnectionEventView, error) {
	resp := view.SNSMicrosoftTeamsTestConnectionEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/microsoft-teams/test-connection"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromTemplatedVmInstance creates VmInstanceFromTemplatedVmInstance
func (cli *ZSClient) CreateVmInstanceFromTemplatedVmInstance(ctx context.Context, templatedVmInstanceUuid string, params param.CreateVmInstanceFromTemplatedVmInstanceParam) (*view.CreateVmInstanceFromTemplatedVmInstanceEventView, error) {
	resp := view.CreateVmInstanceFromTemplatedVmInstanceEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/%s/create-vmInstance-from-templated-vmInstance", templatedVmInstanceUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetZStoneCapability gets ZStoneCapability by uuid
func (cli *ZSClient) GetZStoneCapability(ctx context.Context) (*view.GetZStoneCapabilityView, error) {
	var resp view.GetZStoneCapabilityView
	if err := cli.GetWithRespKey(ctx, "v1/zstone-plugin/capability", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLatestGuestToolsForVm gets LatestGuestToolsForVm by uuid
func (cli *ZSClient) GetLatestGuestToolsForVm(ctx context.Context, uuid string) (*view.GuestToolsInventoryView, error) {
	var resp view.GetLatestGuestToolsForVmView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SyncVmBackupFromImageStoreBackupStorage operates on VmBackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncVmBackupFromImageStoreBackupStorage(ctx context.Context, groupUuid string, params param.SyncVmBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-backups", groupUuid, "actions", "inventories", map[string]interface{}{
		"syncVmBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOffBaremetalChassis operates on PowerOffBaremetalChassis
func (cli *ZSClient) PowerOffBaremetalChassis(ctx context.Context, chassisUuid string, params param.PowerOffBaremetalChassisParam) (*view.PowerOffBaremetalChassisEventView, error) {
	resp := view.PowerOffBaremetalChassisEventView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/chassis", chassisUuid, "actions", "", map[string]interface{}{
		"powerOffBaremetalChassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateInterfaceVlanIds gets CandidateInterfaceVlanIds by uuid
func (cli *ZSClient) GetCandidateInterfaceVlanIds(ctx context.Context) (*view.GetCandidateInterfaceVlanIdsView, error) {
	var resp view.GetCandidateInterfaceVlanIdsView
	if err := cli.GetWithRespKey(ctx, "v1/host/network-interface-vlan-ids", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddCertificateToLoadBalancerListener adds CertificateToLoadBalancerListener
func (cli *ZSClient) AddCertificateToLoadBalancerListener(ctx context.Context, listenerUuid string, params param.AddCertificateToLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/load-balancers/listeners/%s/certificate", listenerUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFaultToleranceVmInstance creates FaultToleranceVmInstance
func (cli *ZSClient) CreateFaultToleranceVmInstance(ctx context.Context, params param.CreateFaultToleranceVmInstanceParam) (*view.CreateFaultToleranceVmInstanceEventView, error) {
	resp := view.CreateFaultToleranceVmInstanceEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/fault-tolerance"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteResourceStackVmPortMonitor deletes ResourceStackVmPortMonitor
func (cli *ZSClient) DeleteResourceStackVmPortMonitor(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/cloudformation/stack/monitor/delvm", uuid, string(deleteMode))
}

// GetNetworkServiceTypes gets NetworkServiceTypes by uuid
func (cli *ZSClient) GetNetworkServiceTypes(ctx context.Context) (*view.GetNetworkServiceTypesView, error) {
	var resp view.GetNetworkServiceTypesView
	if err := cli.GetWithRespKey(ctx, "v1/network-services/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmUserDefinedXml deletes VmUserDefinedXml
func (cli *ZSClient) DeleteVmUserDefinedXml(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// DeleteGCJob deletes GCJob
func (cli *ZSClient) DeleteGCJob(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/gc-jobs", uuid, string(deleteMode))
}

// DeleteEmailAddressOfSNSEmailEndpoint deletes EmailAddressOfSNSEmailEndpoint
func (cli *ZSClient) DeleteEmailAddressOfSNSEmailEndpoint(ctx context.Context, endpointUuid string, emailAddressUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/sns/application-endpoints/emails", endpointUuid, fmt.Sprintf("email-addresses/%s", emailAddressUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetCurrentTime gets CurrentTime by uuid
func (cli *ZSClient) GetCurrentTime(ctx context.Context) (*view.GetCurrentTimeView, error) {
	resp := view.GetCurrentTimeView{}
	if err := cli.PutWithRespKey(ctx, "v1/management-nodes/actions", "", "", map[string]interface{}{
		"getCurrentTime": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CalculateAccountSpending operates on AccountSpending
func (cli *ZSClient) CalculateAccountSpending(ctx context.Context, accountUuid string, params param.CalculateAccountSpendingParam) (*view.CalculateAccountSpendingView, error) {
	resp := view.CalculateAccountSpendingView{}
	if err := cli.PutWithSpec(ctx, "v1/billings/accounts", accountUuid, "actions", "", map[string]interface{}{
		"calculateAccountSpending": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSSOClient deletes SSOClient
func (cli *ZSClient) DeleteSSOClient(ctx context.Context, params param.DeleteSSOClientParam) (*view.DeleteSSOClientEventView, error) {
	resp := view.DeleteSSOClientEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/delete/sso/client"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedIpsec gets VpcAttachedIpsec by uuid
func (cli *ZSClient) GetVpcAttachedIpsec(ctx context.Context, uuid string, params param.GetVpcAttachedIpsecParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/virtual-routers/%s/attached-ipsec", uuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmAttachableL3Network gets VmAttachableL3Network by uuid
func (cli *ZSClient) GetVmAttachableL3Network(ctx context.Context, uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImagesFromImageStoreBackupStorage gets ImagesFromImageStoreBackupStorage by uuid
func (cli *ZSClient) GetImagesFromImageStoreBackupStorage(ctx context.Context, uuid string, params param.GetImagesFromImageStoreBackupStorageParam) (*view.GetImagesFromImageStoreBackupStorageView, error) {
	resp := view.GetImagesFromImageStoreBackupStorageView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage", uuid, "image-store", "infos", map[string]interface{}{
		"getImagesFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncChronyServers operates on ChronyServers
func (cli *ZSClient) SyncChronyServers(ctx context.Context) (*view.SyncChronyServersEventView, error) {
	resp := view.SyncChronyServersEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/zops/chrony/actions", "", "", map[string]interface{}{
		"syncChronyServers": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetElaborationCategories gets ElaborationCategories by uuid
func (cli *ZSClient) GetElaborationCategories(ctx context.Context) (*view.GetElaborationCategoriesView, error) {
	var resp view.GetElaborationCategoriesView
	if err := cli.GetWithRespKey(ctx, "v1/errorcode/elaborations/categories", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetScsiLunCandidatesForAttachingVm gets ScsiLunCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetScsiLunCandidatesForAttachingVm(ctx context.Context, uuid string) (*view.ScsiLunInventoryView, error) {
	var resp view.ScsiLunInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmInstanceProtectedRecoveryPoints gets VmInstanceProtectedRecoveryPoints by uuid
func (cli *ZSClient) GetVmInstanceProtectedRecoveryPoints(ctx context.Context, uuid string) (*view.GetVmInstanceProtectedRecoveryPointsView, error) {
	var resp view.GetVmInstanceProtectedRecoveryPointsView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVmToVmSchedulingRuleGroup adds VmToVmSchedulingRuleGroup
func (cli *ZSClient) AddVmToVmSchedulingRuleGroup(ctx context.Context, vmGroupUuid, vmUuid string, params param.AddVmToVmSchedulingRuleGroupParam) (*view.AddVmToVmSchedulingRuleGroupEventView, error) {
	resp := view.AddVmToVmSchedulingRuleGroupEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vmSchedulingRuleGroup/%s/vmInstance/%s", vmGroupUuid, vmUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostMultipathTopology gets HostMultipathTopology by uuid
func (cli *ZSClient) GetHostMultipathTopology(ctx context.Context) (*view.GetHostMultipathTopologyView, error) {
	var resp view.GetHostMultipathTopologyView
	if err := cli.GetWithRespKey(ctx, "v1/storage-devices/multipath/topology", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostWebSshUrl gets HostWebSshUrl by uuid
func (cli *ZSClient) GetHostWebSshUrl(ctx context.Context, params param.GetHostWebSshUrlParam) (*view.GetHostWebSshUrlEventView, error) {
	resp := view.GetHostWebSshUrlEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hosts/webssh"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncBackupFromImageStoreBackupStorage operates on BackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncBackupFromImageStoreBackupStorage(ctx context.Context, uuid string, params param.SyncBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volume-backups", uuid, "actions", "inventory", map[string]interface{}{
		"syncBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetL3NetworkMtu operates on L3NetworkMtu
func (cli *ZSClient) SetL3NetworkMtu(ctx context.Context, l3NetworkUuid string, params param.SetL3NetworkMtuParam) (*view.SetL3NetworkMtuEventView, error) {
	resp := view.SetL3NetworkMtuEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks/%s/mtu", l3NetworkUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkRouterInterfaceIp gets L3NetworkRouterInterfaceIp by uuid
func (cli *ZSClient) GetL3NetworkRouterInterfaceIp(ctx context.Context, uuid string) (*view.GetL3NetworkRouterInterfaceIpView, error) {
	var resp view.GetL3NetworkRouterInterfaceIpView
	if err := cli.GetWithRespKey(ctx, "v1/l3-networks", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVmClock operates on VmClock
func (cli *ZSClient) SyncVmClock(ctx context.Context, uuid string, params param.SyncVmClockParam) (*view.SyncVmClockEventView, error) {
	resp := view.SyncVmClockEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"syncVmClock": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallIpSetTemplate queries FirewallIpSetTemplate list
func (cli *ZSClient) QueryFirewallIpSetTemplate(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallIpSetTemplateInventoryView, error) {
	var resp []view.VpcFirewallIpSetTemplateInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls/ipset/templates", params, &resp)
}

func (cli *ZSClient) GetFirewallIpSetTemplate(ctx context.Context, uuid string) (*view.VpcFirewallIpSetTemplateInventoryView, error) {
	var resp view.VpcFirewallIpSetTemplateInventoryView
	if err := cli.Get(ctx, "v1/vpcfirewalls/ipset/templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallIpSetTemplate Pagination
func (cli *ZSClient) PageFirewallIpSetTemplate(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallIpSetTemplateInventoryView, int, error) {
	var firewallIpSetTemplates []view.VpcFirewallIpSetTemplateInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls/ipset/templates", params, &firewallIpSetTemplates)
	return firewallIpSetTemplates, total, err
}

// CreateSNSSnmpEndpoint creates SNSSnmpEndpoint
func (cli *ZSClient) CreateSNSSnmpEndpoint(ctx context.Context) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.Post(ctx, "v1/sns/application-endpoints/snmp", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostNetworkFacts gets HostNetworkFacts by uuid
func (cli *ZSClient) GetHostNetworkFacts(ctx context.Context, uuid string) (*view.GetHostNetworkFactsView, error) {
	var resp view.GetHostNetworkFactsView
	if err := cli.GetWithRespKey(ctx, "v1/hosts/network-facts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpTrashOnBackupStorage operates on UpTrashOnBackupStorage
func (cli *ZSClient) CleanUpTrashOnBackupStorage(ctx context.Context, uuid string, params param.CleanUpTrashOnBackupStorageParam) (*view.CleanUpTrashOnBackupStorageEventView, error) {
	resp := view.CleanUpTrashOnBackupStorageEventView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage", uuid, "trash/actions", "", map[string]interface{}{
		"cleanUpTrashOnBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVRouterNetworksToFlowMeter adds VRouterNetworksToFlowMeter
func (cli *ZSClient) AddVRouterNetworksToFlowMeter(ctx context.Context, flowMeterUuid, vRouterUuid string, params param.AddVRouterNetworksToFlowMeterParam) (*view.NetworkRouterFlowMeterRefInventoryView, error) {
	resp := view.NetworkRouterFlowMeterRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/flowmeters/%s/router/%s/addnetworks", flowMeterUuid, vRouterUuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPriceTableFromAccount operates on PriceTableFromAccount
func (cli *ZSClient) DetachPriceTableFromAccount(ctx context.Context, tableUuid string, accountUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/billings/price-tables", tableUuid, fmt.Sprintf("accounts/%s", accountUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// SetVmBootVolume operates on VmBootVolume
func (cli *ZSClient) SetVmBootVolume(ctx context.Context, vmInstanceUuid string, params param.SetVmBootVolumeParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "actions", "inventory", map[string]interface{}{
		"setVmBootVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnlockIdentity operates on UnlockIdentity
func (cli *ZSClient) UnlockIdentity(ctx context.Context) (*view.UnlockIdentityView, error) {
	var resp view.UnlockIdentityView
	if err := cli.GetWithRespKey(ctx, "v1/login/control/unlock", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmNicsForPortMirror gets CandidateVmNicsForPortMirror by uuid
func (cli *ZSClient) GetCandidateVmNicsForPortMirror(ctx context.Context, portMirrorUuid string, typeParam string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	err := cli.GetWithSpec(ctx, "v1/port-mirrors", portMirrorUuid, fmt.Sprintf("vm-instances/candidate-nics/%s", typeParam), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmSchedulingRuleState changes VmSchedulingRuleState
func (cli *ZSClient) ChangeVmSchedulingRuleState(ctx context.Context, uuid string, params param.ChangeVmSchedulingRuleStateParam) (*view.VmSchedulingRuleInventoryView, error) {
	resp := view.VmSchedulingRuleInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vmSchedulingRule", uuid, "actions", "inventory", map[string]interface{}{
		"changeVmSchedulingRuleState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVpcHaGroupMonitorIps changes VpcHaGroupMonitorIps
func (cli *ZSClient) ChangeVpcHaGroupMonitorIps(ctx context.Context, uuid string, params param.ChangeVpcHaGroupMonitorIpsParam) (*view.VpcHaGroupInventoryView, error) {
	resp := view.VpcHaGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vpc/hagroups", uuid, "monitorIps", "inventory", map[string]interface{}{
		"changeVpcHaGroupMonitorIps": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRule creates FirewallRule
func (cli *ZSClient) CreateFirewallRule(ctx context.Context, params param.CreateFirewallRuleParam) (*view.VpcFirewallRuleInventoryView, error) {
	resp := view.VpcFirewallRuleInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpcfirewalls/rules"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RenewSession operates on RenewSession
func (cli *ZSClient) RenewSession(ctx context.Context, sessionUuid string, params param.RenewSessionParam) (*view.SessionInventoryView, error) {
	resp := view.SessionInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/accounts/sessions", sessionUuid, "renew", "inventory", map[string]interface{}{
		"renewSession": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ConvertTemplatedVmInstanceToVmInstance operates on ConvertTemplatedVmInstanceToVmInstance
func (cli *ZSClient) ConvertTemplatedVmInstanceToVmInstance(ctx context.Context, templatedVmInstanceUuid string, params param.ConvertTemplatedVmInstanceToVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/%s/convert-to-vmInstance", templatedVmInstanceUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmConsoleMode operates on VmConsoleMode
func (cli *ZSClient) SetVmConsoleMode(ctx context.Context, uuid string, params param.SetVmConsoleModeParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "inventory", map[string]interface{}{
		"setVmConsoleMode": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmEmulatorPinning gets VmEmulatorPinning by uuid
func (cli *ZSClient) GetVmEmulatorPinning(ctx context.Context, uuid string) (*view.GetVmEmulatorPinningView, error) {
	var resp view.GetVmEmulatorPinningView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDataVolumeAttachableVm gets DataVolumeAttachableVm by uuid
func (cli *ZSClient) GetDataVolumeAttachableVm(ctx context.Context, uuid string) (*view.VmInstanceInventoryView, error) {
	var resp view.VmInstanceInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIpRangeByNetworkCidr adds IpRangeByNetworkCidr
func (cli *ZSClient) AddIpRangeByNetworkCidr(ctx context.Context, l3NetworkUuid string, params param.AddIpRangeByNetworkCidrParam) (*view.IpRangeInventoryView, error) {
	resp := view.IpRangeInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks/%s/ip-ranges/by-cidr", l3NetworkUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLdapEntry gets LdapEntry by uuid
func (cli *ZSClient) GetLdapEntry(ctx context.Context) (*view.LdapEntryInventoryView, error) {
	var resp view.LdapEntryInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/ldap/entry", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateL2NoVlanNetwork creates L2NoVlanNetwork
func (cli *ZSClient) CreateL2NoVlanNetwork(ctx context.Context) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.Post(ctx, "v1/l2-networks/no-vlan", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UngenerateSeMdevDevices operates on UngenerateSeMdevDevices
func (cli *ZSClient) UngenerateSeMdevDevices(ctx context.Context, mttyDeviceUuid string, params param.UngenerateSeMdevDevicesParam) (*view.UngenerateSeMdevDevicesEventView, error) {
	resp := view.UngenerateSeMdevDevicesEventView{}
	if err := cli.PutWithSpec(ctx, "v1/mtty-devices", mttyDeviceUuid, "actions", "", map[string]interface{}{
		"ungenerateSeMdevDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddMonToCephBackupStorage adds MonToCephBackupStorage
func (cli *ZSClient) AddMonToCephBackupStorage(ctx context.Context, uuid string, params param.AddMonToCephBackupStorageParam) (*view.CephBackupStorageInventoryView, error) {
	resp := view.CephBackupStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/backup-storage/ceph/%s/mons", uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmEmulatorPinning operates on VmEmulatorPinning
func (cli *ZSClient) SetVmEmulatorPinning(ctx context.Context, uuid string, params param.SetVmEmulatorPinningParam) (*view.SetVmEmulatorPinningEventView, error) {
	resp := view.SetVmEmulatorPinningEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"setVmEmulatorPinning": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseUKeyStatus gets LicenseUKeyStatus by uuid
func (cli *ZSClient) GetLicenseUKeyStatus(ctx context.Context) (*view.UKeyInventoryView, error) {
	resp := view.UKeyInventoryView{}
	if err := cli.Post(ctx, "v1/licenses/actions", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceNames gets ResourceNames by uuid
func (cli *ZSClient) GetResourceNames(ctx context.Context) (*view.ResourceInventoryView, error) {
	var resp view.ResourceInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/resources/names", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetL3NetworkRouterInterfaceIp operates on L3NetworkRouterInterfaceIp
func (cli *ZSClient) SetL3NetworkRouterInterfaceIp(ctx context.Context, l3NetworkUuid string, params param.SetL3NetworkRouterInterfaceIpParam) (*view.SetL3NetworkRouterInterfaceIpEventView, error) {
	resp := view.SetL3NetworkRouterInterfaceIpEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks/%s/router-interface-ip", l3NetworkUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceConfigs gets ResourceConfigs by uuid
func (cli *ZSClient) GetResourceConfigs(ctx context.Context, resourceUuid string, category string) (*view.GetResourceConfigsView, error) {
	var resp view.GetResourceConfigsView
	err := cli.GetWithSpec(ctx, "v1/resource-configurations", resourceUuid, fmt.Sprintf("%s", category), "", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryEthernetVF queries EthernetVF list
func (cli *ZSClient) QueryEthernetVF(ctx context.Context, params *param.QueryParam) ([]view.EthernetVfPciDeviceInventoryView, error) {
	var resp []view.EthernetVfPciDeviceInventoryView
	return resp, cli.List(ctx, "v1/pci-device/ethernet-vfs", params, &resp)
}

func (cli *ZSClient) GetEthernetVF(ctx context.Context, uuid string) (*view.EthernetVfPciDeviceInventoryView, error) {
	var resp view.EthernetVfPciDeviceInventoryView
	if err := cli.Get(ctx, "v1/pci-device/ethernet-vfs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEthernetVF Pagination
func (cli *ZSClient) PageEthernetVF(ctx context.Context, params *param.QueryParam) ([]view.EthernetVfPciDeviceInventoryView, int, error) {
	var ethernetVFs []view.EthernetVfPciDeviceInventoryView
	total, err := cli.Page(ctx, "v1/pci-device/ethernet-vfs", params, &ethernetVFs)
	return ethernetVFs, total, err
}

// DetachPrimaryStorageFromCluster operates on PrimaryStorageFromCluster
func (cli *ZSClient) DetachPrimaryStorageFromCluster(ctx context.Context, clusterUuid string, primaryStorageUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/clusters", clusterUuid, fmt.Sprintf("primary-storage/%s", primaryStorageUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// UpdateFirewallRuleTemplate updates FirewallRuleTemplate
func (cli *ZSClient) UpdateFirewallRuleTemplate(ctx context.Context, uuid string, params param.UpdateFirewallRuleTemplateParam) (*view.VpcFirewallRuleTemplateInventoryView, error) {
	resp := view.VpcFirewallRuleTemplateInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vpcfirewalls/rules/template", uuid, "actions", "inventory", map[string]interface{}{
		"updateFirewallRuleTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUsbDeviceCandidatesForAttachingVm gets UsbDeviceCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetUsbDeviceCandidatesForAttachingVm(ctx context.Context, uuid string) (*view.UsbDeviceInventoryView, error) {
	var resp view.UsbDeviceInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFactoryModeState gets FactoryModeState by uuid
func (cli *ZSClient) GetFactoryModeState(ctx context.Context) (*view.GetFactoryModeStateView, error) {
	var resp view.GetFactoryModeStateView
	if err := cli.GetWithRespKey(ctx, "v1/management-nodes/factory-mode-state", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckStackTemplateParameters operates on StackTemplateParameters
func (cli *ZSClient) CheckStackTemplateParameters(ctx context.Context, params param.CheckStackTemplateParametersParam) (*view.CheckStackTemplateParametersView, error) {
	resp := view.CheckStackTemplateParametersView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cloudformation/stack/check"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL3NetworksForLoadBalancer gets CandidateL3NetworksForLoadBalancer by uuid
func (cli *ZSClient) GetCandidateL3NetworksForLoadBalancer(ctx context.Context, uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/load-balancers/listeners", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddServerGroupToLoadBalancerListener adds ServerGroupToLoadBalancerListener
func (cli *ZSClient) AddServerGroupToLoadBalancerListener(ctx context.Context, listenerUuid string, params param.AddServerGroupToLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/load-balancers/listeners/%s/servergroups", listenerUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetActiveAlarmStatus gets ActiveAlarmStatus by uuid
func (cli *ZSClient) GetActiveAlarmStatus(ctx context.Context) (*view.GetActiveAlarmStatusView, error) {
	var resp view.GetActiveAlarmStatusView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/activealarms/status", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerResetHost operates on PowerResetHost
func (cli *ZSClient) PowerResetHost(ctx context.Context, uuid string, params param.PowerResetHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/power", uuid, "actions", "", map[string]interface{}{
		"powerResetHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallRule queries FirewallRule list
func (cli *ZSClient) QueryFirewallRule(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleInventoryView, error) {
	var resp []view.VpcFirewallRuleInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls/rules", params, &resp)
}

func (cli *ZSClient) GetFirewallRule(ctx context.Context, uuid string) (*view.VpcFirewallRuleInventoryView, error) {
	var resp view.VpcFirewallRuleInventoryView
	if err := cli.Get(ctx, "v1/vpcfirewalls/rules", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallRule Pagination
func (cli *ZSClient) PageFirewallRule(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallRuleInventoryView, int, error) {
	var firewallRules []view.VpcFirewallRuleInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls/rules", params, &firewallRules)
	return firewallRules, total, err
}

// RevertVmFromVmBackup operates on VmFromVmBackup
func (cli *ZSClient) RevertVmFromVmBackup(ctx context.Context, groupUuid string, params param.RevertVmFromVmBackupParam) (*view.RevertVmFromVmBackupEventView, error) {
	resp := view.RevertVmFromVmBackupEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-backups", groupUuid, "actions", "", map[string]interface{}{
		"revertVmFromVmBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteZceXAlertPlatform deletes ZceXAlertPlatform
func (cli *ZSClient) DeleteZceXAlertPlatform(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zce-x-plugin", uuid, string(deleteMode))
}

// AttachNvmeServerToCluster operates on NvmeServerToCluster
func (cli *ZSClient) AttachNvmeServerToCluster(ctx context.Context, clusterUuid, uuid string, params param.AttachNvmeServerToClusterParam) (*view.NvmeServerInventoryView, error) {
	resp := view.NvmeServerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/clusters/%s/storage-devices/nvme/servers/%s", clusterUuid, uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmNuma operates on VmNuma
func (cli *ZSClient) SetVmNuma(ctx context.Context, uuid string, params param.SetVmNumaParam) (*view.SetVmNumaEventView, error) {
	resp := view.SetVmNumaEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"setVmNuma": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostResourceAllocation gets HostResourceAllocation by uuid
func (cli *ZSClient) GetHostResourceAllocation(ctx context.Context, uuid string, params param.GetHostResourceAllocationParam) (*view.GetHostResourceAllocationEventView, error) {
	resp := view.GetHostResourceAllocationEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hosts/%s/resource-allocation", uuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachUsbDeviceToVm operates on UsbDeviceToVm
func (cli *ZSClient) AttachUsbDeviceToVm(ctx context.Context, usbDeviceUuid string, params param.AttachUsbDeviceToVmParam) (*view.UsbDeviceInventoryView, error) {
	resp := view.UsbDeviceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/usb-device/usb-devices/%s/attach", usbDeviceUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVRouterOspfArea queries VRouterOspfArea list
func (cli *ZSClient) QueryVRouterOspfArea(ctx context.Context, params *param.QueryParam) ([]view.RouterAreaInventoryView, error) {
	var resp []view.RouterAreaInventoryView
	return resp, cli.List(ctx, "v1/routerArea", params, &resp)
}

func (cli *ZSClient) GetVRouterOspfArea(ctx context.Context, uuid string) (*view.RouterAreaInventoryView, error) {
	var resp view.RouterAreaInventoryView
	if err := cli.Get(ctx, "v1/routerArea", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVRouterOspfArea Pagination
func (cli *ZSClient) PageVRouterOspfArea(ctx context.Context, params *param.QueryParam) ([]view.RouterAreaInventoryView, int, error) {
	var vRouterOspfAreas []view.RouterAreaInventoryView
	total, err := cli.Page(ctx, "v1/routerArea", params, &vRouterOspfAreas)
	return vRouterOspfAreas, total, err
}

// GetLicenseAddOns gets LicenseAddOns by uuid
func (cli *ZSClient) GetLicenseAddOns(ctx context.Context) (*view.GetLicenseAddOnsView, error) {
	var resp view.GetLicenseAddOnsView
	if err := cli.GetWithRespKey(ctx, "v1/licenses/addons", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFirewallRuleSet updates FirewallRuleSet
func (cli *ZSClient) UpdateFirewallRuleSet(ctx context.Context, uuid string, params param.UpdateFirewallRuleSetParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	resp := view.VpcFirewallRuleSetInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vpcfirewalls/ruleSets", uuid, "actions", "inventory", map[string]interface{}{
		"updateFirewallRuleSet": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshSearchIndexes operates on SearchIndexes
func (cli *ZSClient) RefreshSearchIndexes(ctx context.Context) (*view.RefreshSearchIndexesView, error) {
	var resp view.RefreshSearchIndexesView
	if err := cli.GetWithRespKey(ctx, "v1/search/indexes/refresh", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CalculateImageHash operates on ImageHash
func (cli *ZSClient) CalculateImageHash(ctx context.Context, uuid string, params param.CalculateImageHashParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/images", uuid, "actions", "inventory", map[string]interface{}{
		"calculateImageHash": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcIPsecLog gets VpcIPsecLog by uuid
func (cli *ZSClient) GetVpcIPsecLog(ctx context.Context) (*view.GetVpcIPsecLogView, error) {
	var resp view.GetVpcIPsecLogView
	if err := cli.GetWithRespKey(ctx, "v1/vpc/virtual-routers/ipseclog", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromOvf creates VmInstanceFromOvf
func (cli *ZSClient) CreateVmInstanceFromOvf(ctx context.Context, params param.CreateVmInstanceFromOvfParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/ovf/create-vm-instance"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromOvfAsync Async
func (cli *ZSClient) CreateVmInstanceFromOvfAsync(ctx context.Context, params param.CreateVmInstanceFromOvfParam) (string, error) {

	resource := "ovf/create-vm-instance"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetL2NetworkTypes gets L2NetworkTypes by uuid
func (cli *ZSClient) GetL2NetworkTypes(ctx context.Context) (*view.GetL2NetworkTypesView, error) {
	var resp view.GetL2NetworkTypesView
	if err := cli.GetWithRespKey(ctx, "v1/l2-networks/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShutdownHost operates on ShutdownHost
func (cli *ZSClient) ShutdownHost(ctx context.Context, uuid string, params param.ShutdownHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/power", uuid, "actions", "", map[string]interface{}{
		"shutdownHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVmImage changes VmImage
func (cli *ZSClient) ChangeVmImage(ctx context.Context, vmInstanceUuid string, params param.ChangeVmImageParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "actions", "inventory", map[string]interface{}{
		"changeVmImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddResourcesToDirectory adds ResourcesToDirectory
func (cli *ZSClient) AddResourcesToDirectory(ctx context.Context, params param.AddResourcesToDirectoryParam) (*view.AddResourcesToDirectoryEventView, error) {
	resp := view.AddResourcesToDirectoryEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/add/resources/directory"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachGuestToolsIsoToVm operates on GuestToolsIsoToVm
func (cli *ZSClient) AttachGuestToolsIsoToVm(ctx context.Context, uuid string, params param.AttachGuestToolsIsoToVmParam) (*view.AttachGuestToolsIsoToVmEventView, error) {
	resp := view.AttachGuestToolsIsoToVmEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"attachGuestToolsIsoToVm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmTask gets VmTask by uuid
func (cli *ZSClient) GetVmTask(ctx context.Context) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/task-details", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableCdpTask operates on DisableCdpTask
func (cli *ZSClient) DisableCdpTask(ctx context.Context, uuid string, params param.DisableCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	resp := view.CdpTaskInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cdp-task/disable/%s", uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetIpOnHostNetworkBonding operates on IpOnHostNetworkBonding
func (cli *ZSClient) SetIpOnHostNetworkBonding(ctx context.Context, bondingUuid string, params param.SetIpOnHostNetworkBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hosts/bondings/%s/ip", bondingUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ZSha2Demote operates on ZSha2Demote
func (cli *ZSClient) ZSha2Demote(ctx context.Context) (*view.ZSha2DemoteEventView, error) {
	resp := view.ZSha2DemoteEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/management-nodes/zsha2/demote", "", "", map[string]interface{}{
		"zSha2Demote": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBonding creates Bonding
func (cli *ZSClient) CreateBonding(ctx context.Context, params param.CreateBondingParam) (*view.HostNetworkBondingInventoryView, error) {
	resp := view.HostNetworkBondingInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hosts/bondings"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeResourceOwner changes ResourceOwner
func (cli *ZSClient) ChangeResourceOwner(ctx context.Context, accountUuid string, params param.ChangeResourceOwnerParam) (*view.AccountResourceRefInventoryView, error) {
	resp := view.AccountResourceRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/account/%s/resources", accountUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostIommuState gets HostIommuState by uuid
func (cli *ZSClient) GetHostIommuState(ctx context.Context, uuid string) (*view.GetHostIommuStateView, error) {
	var resp view.GetHostIommuStateView
	if err := cli.GetWithRespKey(ctx, "v1/pci-device/hosts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachUsbDeviceFromVm operates on UsbDeviceFromVm
func (cli *ZSClient) DetachUsbDeviceFromVm(ctx context.Context, usbDeviceUuid string, params param.DetachUsbDeviceFromVmParam) (*view.UsbDeviceInventoryView, error) {
	resp := view.UsbDeviceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/usb-device/usb-devices/%s/detach", usbDeviceUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMetricData gets MetricData by uuid
func (cli *ZSClient) GetMetricData(ctx context.Context) (*view.GetMetricDataView, error) {
	var resp view.GetMetricDataView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/metrics", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableCbtTask operates on EnableCbtTask
func (cli *ZSClient) EnableCbtTask(ctx context.Context, uuid string, params param.EnableCbtTaskParam) (*view.EnableCbtTaskEventView, error) {
	resp := view.EnableCbtTaskEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cbt-task/enable/%s", uuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeTemplateFromVolumeSnapshot creates DataVolumeTemplateFromVolumeSnapshot
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeSnapshot(ctx context.Context, snapshotUuid string, params param.CreateDataVolumeTemplateFromVolumeSnapshotParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/images/data-volume-templates/from/volume-snapshots/%s", snapshotUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachRoleFromAccount operates on RoleFromAccount
func (cli *ZSClient) DetachRoleFromAccount(ctx context.Context, accountUuid string, roleUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/identities/accounts", accountUuid, fmt.Sprintf("roles/%s", roleUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AddLabelToEventSubscription adds LabelToEventSubscription
func (cli *ZSClient) AddLabelToEventSubscription(ctx context.Context, subscriptionUuid string, params param.AddLabelToEventSubscriptionParam) (*view.EventSubscriptionLabelInventoryView, error) {
	resp := view.EventSubscriptionLabelInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/events/subscriptions/%s/labels", subscriptionUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddRendezvousPointToMulticastRouter adds RendezvousPointToMulticastRouter
func (cli *ZSClient) AddRendezvousPointToMulticastRouter(ctx context.Context, uuid string, params param.AddRendezvousPointToMulticastRouterParam) (*view.MulticastRouterInventoryView, error) {
	resp := view.MulticastRouterInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/multicast/virtual-routers/%s/RendezvousPoint", uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcVRouterDistributedRoutingConnections gets VpcVRouterDistributedRoutingConnections by uuid
func (cli *ZSClient) GetVpcVRouterDistributedRoutingConnections(ctx context.Context, uuid string) (*view.StringView, error) {
	var resp view.StringView
	if err := cli.GetWithRespKey(ctx, "v1/vpc/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QuerySNSTopicSubscriber queries SNSTopicSubscriber list
func (cli *ZSClient) QuerySNSTopicSubscriber(ctx context.Context, params *param.QueryParam) ([]view.SNSSubscriberInventoryView, error) {
	var resp []view.SNSSubscriberInventoryView
	return resp, cli.List(ctx, "v1/sns/topics/subscribers", params, &resp)
}

func (cli *ZSClient) GetSNSTopicSubscriber(ctx context.Context, uuid string) (*view.SNSSubscriberInventoryView, error) {
	var resp view.SNSSubscriberInventoryView
	if err := cli.Get(ctx, "v1/sns/topics/subscribers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSTopicSubscriber Pagination
func (cli *ZSClient) PageSNSTopicSubscriber(ctx context.Context, params *param.QueryParam) ([]view.SNSSubscriberInventoryView, int, error) {
	var sNSTopicSubscribers []view.SNSSubscriberInventoryView
	total, err := cli.Page(ctx, "v1/sns/topics/subscribers", params, &sNSTopicSubscribers)
	return sNSTopicSubscribers, total, err
}

// UpdateThirdpartyAlerts updates ThirdpartyAlerts
func (cli *ZSClient) UpdateThirdpartyAlerts(ctx context.Context, params param.UpdateThirdpartyAlertsParam) (*view.UpdateThirdpartyAlertsEventView, error) {
	resp := view.UpdateThirdpartyAlertsEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/third-party/alerts/actions", "", "", map[string]interface{}{
		"updateThirdpartyAlerts": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmUsbRedirect operates on VmUsbRedirect
func (cli *ZSClient) SetVmUsbRedirect(ctx context.Context, uuid string, params param.SetVmUsbRedirectParam) (*view.SetVmUsbRedirectEventView, error) {
	resp := view.SetVmUsbRedirectEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"setVmUsbRedirect": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostCandidatesForVmMigration gets HostCandidatesForVmMigration by uuid
func (cli *ZSClient) GetHostCandidatesForVmMigration(ctx context.Context, uuid string) (*view.HostInventoryView, error) {
	var resp view.HostInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/hosts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmNicAttachableEips gets VmNicAttachableEips by uuid
func (cli *ZSClient) GetVmNicAttachableEips(ctx context.Context, uuid string) (*view.EipInventoryView, error) {
	var resp view.EipInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances/nics", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFactoryModeState updates FactoryModeState
func (cli *ZSClient) UpdateFactoryModeState(ctx context.Context, params param.UpdateFactoryModeStateParam) (*view.UpdateFactoryModeStateEventView, error) {
	resp := view.UpdateFactoryModeStateEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/management-nodes/actions", "", "", map[string]interface{}{
		"updateFactoryModeState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateChronyServers updates ChronyServers
func (cli *ZSClient) UpdateChronyServers(ctx context.Context, params param.UpdateChronyServersParam) (*view.UpdateChronyServersEventView, error) {
	resp := view.UpdateChronyServersEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/zops/chrony/actions", "", "", map[string]interface{}{
		"updateChronyServers": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DebugSignal operates on DebugSignal
func (cli *ZSClient) DebugSignal(ctx context.Context, params param.DebugSignalParam) (*view.DebugSignalEventView, error) {
	resp := view.DebugSignalEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/debug"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPhysicalMachineBlockDevices gets PhysicalMachineBlockDevices by uuid
func (cli *ZSClient) GetPhysicalMachineBlockDevices(ctx context.Context) (*view.GetPhysicalMachineBlockDevicesView, error) {
	var resp view.GetPhysicalMachineBlockDevicesView
	if err := cli.GetWithRespKey(ctx, "v1/host/get-block-devices", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPolicyRouteRuleSetToL3 operates on PolicyRouteRuleSetToL3
func (cli *ZSClient) AttachPolicyRouteRuleSetToL3(ctx context.Context, ruleSetUuid, l3Uuid string, params param.AttachPolicyRouteRuleSetToL3Param) (*view.AttachPolicyRouteRuleSetToL3EventView, error) {
	resp := view.AttachPolicyRouteRuleSetToL3EventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/policy-routes/rulesets/%s/l3networks/%s", ruleSetUuid, l3Uuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateNetworkBondings gets CandidateNetworkBondings by uuid
func (cli *ZSClient) GetCandidateNetworkBondings(ctx context.Context) (*view.HostNetworkBondingInventoryView, error) {
	var resp view.HostNetworkBondingInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/cluster/hosts-network-bondings", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOAuthClient updates OAuthClient
func (cli *ZSClient) UpdateOAuthClient(ctx context.Context, params param.UpdateOAuthClientParam) (*view.OAuth2ClientInventoryView, error) {
	resp := view.OAuth2ClientInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/update/oauth2/client"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromVolume creates VmInstanceFromVolume
func (cli *ZSClient) CreateVmInstanceFromVolume(ctx context.Context, params param.CreateVmInstanceFromVolumeParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/from/volume"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRolesForAccountGroup gets RolesForAccountGroup by uuid
func (cli *ZSClient) GetRolesForAccountGroup(ctx context.Context, uuid string) (*view.GetRolesForAccountGroupView, error) {
	var resp view.GetRolesForAccountGroupView
	if err := cli.GetWithRespKey(ctx, "v1/account-groups", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcVRouterDistributedRoutingEnabled gets VpcVRouterDistributedRoutingEnabled by uuid
func (cli *ZSClient) GetVpcVRouterDistributedRoutingEnabled(ctx context.Context, uuid string) (*view.GetVpcVRouterDistributedRoutingEnabledView, error) {
	var resp view.GetVpcVRouterDistributedRoutingEnabledView
	if err := cli.GetWithRespKey(ctx, "v1/vpc/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetZWatchAlertHistogram gets ZWatchAlertHistogram by uuid
func (cli *ZSClient) GetZWatchAlertHistogram(ctx context.Context) (*view.GetZWatchAlertHistogramView, error) {
	var resp view.GetZWatchAlertHistogramView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/alert-histories/histogram", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageBootMode operates on ImageBootMode
func (cli *ZSClient) SetImageBootMode(ctx context.Context, uuid string, params param.SetImageBootModeParam) (*view.SetImageBootModeEventView, error) {
	resp := view.SetImageBootModeEventView{}
	if err := cli.PutWithSpec(ctx, "v1/images", uuid, "actions", "", map[string]interface{}{
		"setImageBootMode": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachDataVolumeFromVm operates on DataVolumeFromVm
func (cli *ZSClient) DetachDataVolumeFromVm(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/volumes", uuid, string(deleteMode))
}

// DetachAutoScalingTemplateFromGroup operates on AutoScalingTemplateFromGroup
func (cli *ZSClient) DetachAutoScalingTemplateFromGroup(ctx context.Context, templateUuid string, groupUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/autoscaling/template", templateUuid, fmt.Sprintf("groups/%s", groupUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// CreateRootVolumeTemplateFromRootVolume creates RootVolumeTemplateFromRootVolume
func (cli *ZSClient) CreateRootVolumeTemplateFromRootVolume(ctx context.Context, rootVolumeUuid string, params param.CreateRootVolumeTemplateFromRootVolumeParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/images/root-volume-templates/from/volumes/%s", rootVolumeUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRootVolumeTemplateFromRootVolumeAsync Async
func (cli *ZSClient) CreateRootVolumeTemplateFromRootVolumeAsync(ctx context.Context, params param.CreateRootVolumeTemplateFromRootVolumeParam) (string, error) {

	resource := "images/root-volume-templates/from/volumes/{rootVolumeUuid}"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetVmsCapabilities gets VmsCapabilities by uuid
func (cli *ZSClient) GetVmsCapabilities(ctx context.Context, params param.GetVmsCapabilitiesParam) (*view.GetVmsCapabilitiesEventView, error) {
	resp := view.GetVmsCapabilitiesEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/capabilities"), "vmsCaps", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevokeMonitorTemplateFromMonitorGroup operates on RevokeMonitorTemplateFromMonitorGroup
func (cli *ZSClient) RevokeMonitorTemplateFromMonitorGroup(ctx context.Context, templateUuid string, groupUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/zwatch/monitortemplates", templateUuid, fmt.Sprintf("monitorgroups/%s", groupUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// DeleteFirewallRule deletes FirewallRule
func (cli *ZSClient) DeleteFirewallRule(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vpcfirewalls/rules", uuid, string(deleteMode))
}

// ShareResource operates on ShareResource
func (cli *ZSClient) ShareResource(ctx context.Context, params param.ShareResourceParam) (*view.ShareResourceEventView, error) {
	resp := view.ShareResourceEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/accounts/resources/actions", "", "", map[string]interface{}{
		"shareResource": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAccountQuotaUsage gets AccountQuotaUsage by uuid
func (cli *ZSClient) GetAccountQuotaUsage(ctx context.Context, uuid string) (*view.GetAccountQuotaUsageView, error) {
	var resp view.GetAccountQuotaUsageView
	if err := cli.GetWithRespKey(ctx, "v1/accounts/quota", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateL3NetworksForServerGroup gets CandidateL3NetworksForServerGroup by uuid
func (cli *ZSClient) GetCandidateL3NetworksForServerGroup(ctx context.Context) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/load-balancers/servergroups/candidate-l3network", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmFromCdpBackup creates VmFromCdpBackup
func (cli *ZSClient) CreateVmFromCdpBackup(ctx context.Context, params param.CreateVmFromCdpBackupParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/cdp-backups/actions", "", "inventory", map[string]interface{}{
		"createVmFromCdpBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAtPersonOfAtWeComEndpoint updates AtPersonOfAtWeComEndpoint
func (cli *ZSClient) UpdateAtPersonOfAtWeComEndpoint(ctx context.Context, uuid string, params param.UpdateAtPersonOfAtWeComEndpointParam) (*view.SNSWeComAtPersonInventoryView, error) {
	resp := view.SNSWeComAtPersonInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/sns/application-endpoints/we-com/at-persons", uuid, "actions", "inventory", map[string]interface{}{
		"updateAtPersonOfAtWeComEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SubmitLongJob operates on SubmitLongJob
func (cli *ZSClient) SubmitLongJob(ctx context.Context, params param.SubmitLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/longjobs"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateZceXAlertPlatform creates ZceXAlertPlatform
func (cli *ZSClient) CreateZceXAlertPlatform(ctx context.Context, uuid string, params param.CreateZceXAlertPlatformParam) (*view.ZceXThirdPartyPlatformAlertRefInventoryView, error) {
	resp := view.ZceXThirdPartyPlatformAlertRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zce-x-plugin/%s/alert-platform", uuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeTemplateFromVolumeBackup creates DataVolumeTemplateFromVolumeBackup
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeBackup(ctx context.Context, backupUuid string, params param.CreateDataVolumeTemplateFromVolumeBackupParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/images/data-volume-templates/from/volume-template/%s", backupUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachDataVolumeFromHost operates on DataVolumeFromHost
func (cli *ZSClient) DetachDataVolumeFromHost(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/volumes", uuid, string(deleteMode))
}

// GetDebugSignal gets DebugSignal by uuid
func (cli *ZSClient) GetDebugSignal(ctx context.Context) (*view.GetDebugSignalView, error) {
	var resp view.GetDebugSignalView
	if err := cli.GetWithRespKey(ctx, "v1/debug", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmInstanceRecoveryPoints gets VmInstanceRecoveryPoints by uuid
func (cli *ZSClient) GetVmInstanceRecoveryPoints(ctx context.Context, uuid string) (*view.GetVmInstanceRecoveryPointsView, error) {
	var resp view.GetVmInstanceRecoveryPointsView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSystemTags creates SystemTags
func (cli *ZSClient) CreateSystemTags(ctx context.Context, resourceUuid string, params param.CreateSystemTagsParam) (*view.SystemTagInventoryView, error) {
	resp := view.SystemTagInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/system-tags/%s/tags", resourceUuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachDataVolumeToHost operates on DataVolumeToHost
func (cli *ZSClient) AttachDataVolumeToHost(ctx context.Context, volumeUuid, hostUuid string, params param.AttachDataVolumeToHostParam) (*view.AttachDataVolumeToHostEventView, error) {
	resp := view.AttachDataVolumeToHostEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/volumes/%s/hosts/%s", volumeUuid, hostUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SecurityMachineEncrypt operates on MachineEncrypt
func (cli *ZSClient) SecurityMachineEncrypt(ctx context.Context, params param.SecurityMachineEncryptParam) (*view.SecurityMachineEncryptEventView, error) {
	resp := view.SecurityMachineEncryptEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/security-machine/encrypt/actions"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShareResourceToGroup operates on ShareResourceToGroup
func (cli *ZSClient) ShareResourceToGroup(ctx context.Context, params param.ShareResourceToGroupParam) (*view.ShareResourceToGroupEventView, error) {
	resp := view.ShareResourceToGroupEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/account-groups/resources/actions", "", "", map[string]interface{}{
		"shareResourceToGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetL3NetworkTypes gets L3NetworkTypes by uuid
func (cli *ZSClient) GetL3NetworkTypes(ctx context.Context) (*view.GetL3NetworkTypesView, error) {
	var resp view.GetL3NetworkTypesView
	if err := cli.GetWithRespKey(ctx, "v1/l3-networks/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMemorySnapshotGroupReference gets MemorySnapshotGroupReference by uuid
func (cli *ZSClient) GetMemorySnapshotGroupReference(ctx context.Context) (*view.VolumeSnapshotGroupInventoryView, error) {
	var resp view.VolumeSnapshotGroupInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/memory-snapshots/group/reference", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpImageCacheOnPrimaryStorage operates on UpImageCacheOnPrimaryStorage
func (cli *ZSClient) CleanUpImageCacheOnPrimaryStorage(ctx context.Context, uuid string, params param.CleanUpImageCacheOnPrimaryStorageParam) (*view.CleanUpImageCacheOnPrimaryStorageEventView, error) {
	resp := view.CleanUpImageCacheOnPrimaryStorageEventView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage", uuid, "actions", "", map[string]interface{}{
		"cleanUpImageCacheOnPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddKVMHostFromConfigFile adds KVMHostFromConfigFile
func (cli *ZSClient) AddKVMHostFromConfigFile(ctx context.Context) (*view.AddHostFromConfigFileEventView, error) {
	resp := view.AddHostFromConfigFileEventView{}
	if err := cli.Post(ctx, "v1/hosts/kvm/from-file", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddKVMHostFromConfigFileAsync Async
func (cli *ZSClient) AddKVMHostFromConfigFileAsync(ctx context.Context, params param.AddKVMHostFromConfigFileParam) (string, error) {

	resource := "hosts/kvm/from-file"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetVpcVRouterNetworkServiceState gets VpcVRouterNetworkServiceState by uuid
func (cli *ZSClient) GetVpcVRouterNetworkServiceState(ctx context.Context, uuid string) (*view.GetVpcVRouterNetworkServiceStateView, error) {
	var resp view.GetVpcVRouterNetworkServiceStateView
	if err := cli.GetWithRespKey(ctx, "v1/vpc/virtual-routers", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachNetworkServiceFromL3Network operates on NetworkServiceFromL3Network
func (cli *ZSClient) DetachNetworkServiceFromL3Network(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l3-networks", uuid, string(deleteMode))
}

// DeleteVmBootMode deletes VmBootMode
func (cli *ZSClient) DeleteVmBootMode(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-instances", uuid, string(deleteMode))
}

// CreateDataVolumeFromVolumeBackup creates DataVolumeFromVolumeBackup
func (cli *ZSClient) CreateDataVolumeFromVolumeBackup(ctx context.Context, backupUuid string, params param.CreateDataVolumeFromVolumeBackupParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/volumes/data-volume/from/volume-template/%s", backupUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVMForAttachingAffinityGroup gets CandidateVMForAttachingAffinityGroup by uuid
func (cli *ZSClient) GetCandidateVMForAttachingAffinityGroup(ctx context.Context) (*view.VmInstanceInventoryView, error) {
	var resp view.VmInstanceInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/VM/attachingGroup", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddActionToAlarm adds ActionToAlarm
func (cli *ZSClient) AddActionToAlarm(ctx context.Context, alarmUuid string, params param.AddActionToAlarmParam) (*view.AlarmInventoryView, error) {
	resp := view.AlarmInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/alarms/%s/actions", alarmUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFirewallRule updates FirewallRule
func (cli *ZSClient) UpdateFirewallRule(ctx context.Context, uuid string, params param.UpdateFirewallRuleParam) (*view.VpcFirewallRuleInventoryView, error) {
	resp := view.VpcFirewallRuleInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vpcfirewalls/rules", uuid, "actions", "inventory", map[string]interface{}{
		"updateFirewallRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ZQLQuery operates on ZQLQuery
func (cli *ZSClient) ZQLQuery(ctx context.Context) (*view.ZQLQueryView, error) {
	var resp view.ZQLQueryView
	if err := cli.GetWithRespKey(ctx, "v1/zql", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSharedMountPointPrimaryStorage adds SharedMountPointPrimaryStorage
func (cli *ZSClient) AddSharedMountPointPrimaryStorage(ctx context.Context) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/smp", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSupportAPIs gets SupportAPIs by uuid
func (cli *ZSClient) GetSupportAPIs(ctx context.Context) (*view.GetSupportAPIsView, error) {
	resp := view.GetSupportAPIsView{}
	if err := cli.PutWithRespKey(ctx, "v1/management-nodes/actions", "", "", map[string]interface{}{
		"getSupportAPIs": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetElaborations gets Elaborations by uuid
func (cli *ZSClient) GetElaborations(ctx context.Context) (*view.GetElaborationsView, error) {
	var resp view.GetElaborationsView
	if err := cli.GetWithRespKey(ctx, "v1/errorcode/elaborations", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTrashOnPrimaryStorage gets TrashOnPrimaryStorage by uuid
func (cli *ZSClient) GetTrashOnPrimaryStorage(ctx context.Context) (*view.InstallPathRecycleInventoryView, error) {
	var resp view.InstallPathRecycleInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/trash", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAccessPath gets AccessPath by uuid
func (cli *ZSClient) GetAccessPath(ctx context.Context) (*view.GetAccessPathView, error) {
	var resp view.GetAccessPathView
	if err := cli.GetWithRespKey(ctx, "v1/block-volumes/access/path", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmNicsForLoadBalancer gets CandidateVmNicsForLoadBalancer by uuid
func (cli *ZSClient) GetCandidateVmNicsForLoadBalancer(ctx context.Context, uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/load-balancers/listeners", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageUsageReport gets PrimaryStorageUsageReport by uuid
func (cli *ZSClient) GetPrimaryStorageUsageReport(ctx context.Context, uuid string) (*view.GetPrimaryStorageUsageReportView, error) {
	var resp view.GetPrimaryStorageUsageReportView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachBaremetalPxeServerFromCluster operates on BaremetalPxeServerFromCluster
func (cli *ZSClient) DetachBaremetalPxeServerFromCluster(ctx context.Context, clusterUuid string, pxeServerUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/clusters", clusterUuid, fmt.Sprintf("pxeservers/%s", pxeServerUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// RevertVolumeFromVolumeBackup operates on VolumeFromVolumeBackup
func (cli *ZSClient) RevertVolumeFromVolumeBackup(ctx context.Context, uuid string, params param.RevertVolumeFromVolumeBackupParam) (*view.RevertVolumeFromVolumeBackupEventView, error) {
	resp := view.RevertVolumeFromVolumeBackupEventView{}
	if err := cli.PutWithSpec(ctx, "v1/volume-backups", uuid, "actions", "", map[string]interface{}{
		"revertVolumeFromVolumeBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAccessControlListRedirectRule changes AccessControlListRedirectRule
func (cli *ZSClient) ChangeAccessControlListRedirectRule(ctx context.Context, uuid string, params param.ChangeAccessControlListRedirectRuleParam) (*view.AccessControlListEntryInventoryView, error) {
	resp := view.AccessControlListEntryInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/access-control-lists/redirectRules", uuid, "actions", "inventory", map[string]interface{}{
		"changeAccessControlListRedirectRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeFromVolumeTemplate creates DataVolumeFromVolumeTemplate
func (cli *ZSClient) CreateDataVolumeFromVolumeTemplate(ctx context.Context, imageUuid string, params param.CreateDataVolumeFromVolumeTemplateParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/volumes/data/from/data-volume-templates/%s", imageUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddResourceStackVmPortMonitor adds ResourceStackVmPortMonitor
func (cli *ZSClient) AddResourceStackVmPortMonitor(ctx context.Context, params param.AddResourceStackVmPortMonitorParam) (*view.AddResourceStackVmPortMonitorEventView, error) {
	resp := view.AddResourceStackVmPortMonitorEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cloudformation/stack/monitor/addvm"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LocalStorageGetVolumeMigratableHosts operates on LocalStorageGetVolumeMigratableHosts
func (cli *ZSClient) LocalStorageGetVolumeMigratableHosts(ctx context.Context, uuid string) (*view.HostInventoryView, error) {
	var resp view.HostInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSNSApplicationEndpointState changes SNSApplicationEndpointState
func (cli *ZSClient) ChangeSNSApplicationEndpointState(ctx context.Context, uuid string, params param.ChangeSNSApplicationEndpointStateParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/sns/application-endpoints", uuid, "actions", "inventory", map[string]interface{}{
		"changeSNSApplicationEndpointState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedLoadBalancer gets VpcAttachedLoadBalancer by uuid
func (cli *ZSClient) GetVpcAttachedLoadBalancer(ctx context.Context, uuid string, params param.GetVpcAttachedLoadBalancerParam) (*view.LoadBalancerInventoryView, error) {
	resp := view.LoadBalancerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/virtual-routers/%s/attached-lb", uuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateZStoneClusterConfig updates ZStoneClusterConfig
func (cli *ZSClient) UpdateZStoneClusterConfig(ctx context.Context, params param.UpdateZStoneClusterConfigParam) (*view.UpdateZStoneClusterConfigEventView, error) {
	resp := view.UpdateZStoneClusterConfigEventView{}
	if err := cli.PutWithRespKey(ctx, "v1/zstone-plugin/config/cluster", "", "", map[string]interface{}{
		"updateZStoneClusterConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedPortForwardingRules gets VpcAttachedPortForwardingRules by uuid
func (cli *ZSClient) GetVpcAttachedPortForwardingRules(ctx context.Context, uuid string, params param.GetVpcAttachedPortForwardingRulesParam) (*view.PortForwardingRuleInventoryView, error) {
	resp := view.PortForwardingRuleInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/virtual-routers/%s/attached-pf", uuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVpcVRouterNetworkServiceState operates on VpcVRouterNetworkServiceState
func (cli *ZSClient) SetVpcVRouterNetworkServiceState(ctx context.Context, uuid string, params param.SetVpcVRouterNetworkServiceStateParam) (*view.SetVpcVRouterNetworkServiceStateEventView, error) {
	resp := view.SetVpcVRouterNetworkServiceStateEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/virtual-routers/%s/networkservicestate", uuid), "state", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAccountBilling queries AccountBilling list
func (cli *ZSClient) QueryAccountBilling(ctx context.Context, params *param.QueryParam) ([]view.BillingInventoryView, error) {
	var resp []view.BillingInventoryView
	return resp, cli.List(ctx, "v1/billing/billings", params, &resp)
}

func (cli *ZSClient) GetAccountBilling(ctx context.Context, uuid string) (*view.BillingInventoryView, error) {
	var resp view.BillingInventoryView
	if err := cli.Get(ctx, "v1/billing/billings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAccountBilling Pagination
func (cli *ZSClient) PageAccountBilling(ctx context.Context, params *param.QueryParam) ([]view.BillingInventoryView, int, error) {
	var accountBillings []view.BillingInventoryView
	total, err := cli.Page(ctx, "v1/billing/billings", params, &accountBillings)
	return accountBillings, total, err
}

// AddDnsToVpcRouter adds DnsToVpcRouter
func (cli *ZSClient) AddDnsToVpcRouter(ctx context.Context, uuid string, params param.AddDnsToVpcRouterParam) (*view.VpcRouterVmInventoryView, error) {
	resp := view.VpcRouterVmInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/virtual-routers/%s/dns", uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmXml gets VmXml by uuid
func (cli *ZSClient) GetVmXml(ctx context.Context, uuid string) (*view.GetVmXmlView, error) {
	var resp view.GetVmXmlView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmInstanceFirstBootDevice gets VmInstanceFirstBootDevice by uuid
func (cli *ZSClient) GetVmInstanceFirstBootDevice(ctx context.Context, uuid string) (*view.GetVmInstanceFirstBootDeviceView, error) {
	var resp view.GetVmInstanceFirstBootDeviceView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetServiceTypeOnHostNetworkInterface operates on ServiceTypeOnHostNetworkInterface
func (cli *ZSClient) SetServiceTypeOnHostNetworkInterface(ctx context.Context, params param.SetServiceTypeOnHostNetworkInterfaceParam) (*view.HostNetworkInterfaceServiceRefInventoryView, error) {
	resp := view.HostNetworkInterfaceServiceRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hosts/nics/service-types"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteIpAddress deletes IpAddress
func (cli *ZSClient) DeleteIpAddress(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l3-networks", uuid, string(deleteMode))
}

// AttachRoleToAccountGroup operates on RoleToAccountGroup
func (cli *ZSClient) AttachRoleToAccountGroup(ctx context.Context, groupUuid string, params param.AttachRoleToAccountGroupParam) (*view.AttachRoleToAccountGroupEventView, error) {
	resp := view.AttachRoleToAccountGroupEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/account-groups/%s/roles", groupUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddBackendServerToServerGroup adds BackendServerToServerGroup
func (cli *ZSClient) AddBackendServerToServerGroup(ctx context.Context, serverGroupUuid string, params param.AddBackendServerToServerGroupParam) (*view.LoadBalancerServerGroupInventoryView, error) {
	resp := view.LoadBalancerServerGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/load-balancers/servergroups/%s/backendservers", serverGroupUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnmountVmInstanceRecoveryPoint operates on UnmountVmInstanceRecoveryPoint
func (cli *ZSClient) UnmountVmInstanceRecoveryPoint(ctx context.Context, params param.UnmountVmInstanceRecoveryPointParam) (*view.UnmountVmInstanceRecoveryPointEventView, error) {
	resp := view.UnmountVmInstanceRecoveryPointEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cdp-backup-storage/unmount-recovery-point"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsReadyToGo operates on IsReadyToGo
func (cli *ZSClient) IsReadyToGo(ctx context.Context) (*view.IsReadyToGoView, error) {
	var resp view.IsReadyToGoView
	if err := cli.GetWithRespKey(ctx, "v1/management-nodes/ready", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostIommuStatus gets HostIommuStatus by uuid
func (cli *ZSClient) GetHostIommuStatus(ctx context.Context, uuid string) (*view.GetHostIommuStatusView, error) {
	var resp view.GetHostIommuStatusView
	if err := cli.GetWithRespKey(ctx, "v1/pci-device/hosts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBackupStorageCandidatesForImageMigration gets BackupStorageCandidatesForImageMigration by uuid
func (cli *ZSClient) GetBackupStorageCandidatesForImageMigration(ctx context.Context, uuid string) (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/backup-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DescribeVmInstanceRecoveryPoint operates on DescribeVmInstanceRecoveryPoint
func (cli *ZSClient) DescribeVmInstanceRecoveryPoint(ctx context.Context, uuid string) (*view.DescribeVmInstanceRecoveryPointView, error) {
	var resp view.DescribeVmInstanceRecoveryPointView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateSriovPciDevices operates on SriovPciDevices
func (cli *ZSClient) GenerateSriovPciDevices(ctx context.Context, pciDeviceUuid string, params param.GenerateSriovPciDevicesParam) (*view.GenerateVirtualPciDevicesEventView, error) {
	resp := view.GenerateVirtualPciDevicesEventView{}
	if err := cli.PutWithSpec(ctx, "v1/pci-devices", pciDeviceUuid, "actions", "", map[string]interface{}{
		"generateSriovPciDevices": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPciDeviceCandidatesForAttachingVm gets PciDeviceCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetPciDeviceCandidatesForAttachingVm(ctx context.Context, uuid string) (*view.PciDeviceInventoryView, error) {
	var resp view.PciDeviceInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVRouterOspfArea deletes VRouterOspfArea
func (cli *ZSClient) DeleteVRouterOspfArea(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/routerArea", uuid, string(deleteMode))
}

// CalculateAccountBillingSpending operates on AccountBillingSpending
func (cli *ZSClient) CalculateAccountBillingSpending(ctx context.Context, accountUuid string, params param.CalculateAccountBillingSpendingParam) (*view.CalculateAccountBillingSpendingView, error) {
	resp := view.CalculateAccountBillingSpendingView{}
	if err := cli.PutWithSpec(ctx, "v1/billings/accounts", accountUuid, "actions", "", map[string]interface{}{
		"calculateAccountBillingSpending": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeMonitorTriggerState changes MonitorTriggerState
func (cli *ZSClient) ChangeMonitorTriggerState(ctx context.Context, uuid string, params param.ChangeMonitorTriggerStateParam) (*view.MonitorTriggerInventoryView, error) {
	resp := view.MonitorTriggerInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/monitoring/triggers", uuid, "actions", "inventory", map[string]interface{}{
		"changeMonitorTriggerState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostBlockDevices gets HostBlockDevices by uuid
func (cli *ZSClient) GetHostBlockDevices(ctx context.Context, uuid string) (*view.GetHostBlockDevicesView, error) {
	var resp view.GetHostBlockDevicesView
	if err := cli.GetWithRespKey(ctx, "v1/hosts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTaskProgress gets TaskProgress by uuid
func (cli *ZSClient) GetTaskProgress(ctx context.Context, uuid string) (*view.TaskProgressInventoryView, error) {
	var resp view.TaskProgressInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/task-progresses", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartDataProtection starts DataProtection
func (cli *ZSClient) StartDataProtection(ctx context.Context, params param.StartDataProtectionParam) (*view.StartDataProtectionEventView, error) {
	resp := view.StartDataProtectionEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/start/data/protection/"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartDataProtectionAsync Async
func (cli *ZSClient) StartDataProtectionAsync(ctx context.Context, params param.StartDataProtectionParam) (string, error) {

	resource := "start/data/protection/"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetVipAvailablePort gets VipAvailablePort by uuid
func (cli *ZSClient) GetVipAvailablePort(ctx context.Context, uuid string) (*view.GetVipAvailablePortView, error) {
	var resp view.GetVipAvailablePortView
	if err := cli.GetWithRespKey(ctx, "v1/vips", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeActiveAlarmState changes ActiveAlarmState
func (cli *ZSClient) ChangeActiveAlarmState(ctx context.Context, params param.ChangeActiveAlarmStateParam) (*view.ChangeActiveAlarmStateEventView, error) {
	resp := view.ChangeActiveAlarmStateEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/activealarms/actions"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVolumeState changes VolumeState
func (cli *ZSClient) ChangeVolumeState(ctx context.Context, uuid string, params param.ChangeVolumeStateParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volumes", uuid, "actions", "inventory", map[string]interface{}{
		"changeVolumeState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmCleanTraffic operates on VmCleanTraffic
func (cli *ZSClient) SetVmCleanTraffic(ctx context.Context, uuid string, params param.SetVmCleanTrafficParam) (*view.SetVmCleanTrafficEventView, error) {
	resp := view.SetVmCleanTrafficEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"setVmCleanTraffic": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmBootMode operates on VmBootMode
func (cli *ZSClient) SetVmBootMode(ctx context.Context, uuid string, params param.SetVmBootModeParam) (*view.SetVmBootModeEventView, error) {
	resp := view.SetVmBootModeEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", uuid, "actions", "", map[string]interface{}{
		"setVmBootMode": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MountVmInstanceRecoveryPoint operates on MountVmInstanceRecoveryPoint
func (cli *ZSClient) MountVmInstanceRecoveryPoint(ctx context.Context, params param.MountVmInstanceRecoveryPointParam) (*view.MountVmInstanceRecoveryPointEventView, error) {
	resp := view.MountVmInstanceRecoveryPointEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cdp-backup-storage/mount-recovery-point"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncImageSize operates on ImageSize
func (cli *ZSClient) SyncImageSize(ctx context.Context, uuid string, params param.SyncImageSizeParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/images", uuid, "actions", "inventory", map[string]interface{}{
		"syncImageSize": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVxlanPoolRemoteVtep creates VxlanPoolRemoteVtep
func (cli *ZSClient) CreateVxlanPoolRemoteVtep(ctx context.Context, l2NetworkUuid, clusterUuid string, params param.CreateVxlanPoolRemoteVtepParam) (*view.RemoteVtepInventoryView, error) {
	resp := view.RemoteVtepInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l2-networks/%s/clusters/%s/remote-vtep-ip", l2NetworkUuid, clusterUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNoTriggerSchedulerJobs gets NoTriggerSchedulerJobs by uuid
func (cli *ZSClient) GetNoTriggerSchedulerJobs(ctx context.Context) (*view.SchedulerJobInventoryView, error) {
	var resp view.SchedulerJobInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/scheduler/jobs/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ProtectVmInstanceRecoveryPoint operates on ProtectVmInstanceRecoveryPoint
func (cli *ZSClient) ProtectVmInstanceRecoveryPoint(ctx context.Context, vmInstanceUuid string, params param.ProtectVmInstanceRecoveryPointParam) (*view.ProtectVmInstanceRecoveryPointEventView, error) {
	resp := view.ProtectVmInstanceRecoveryPointEventView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "protect-recovery-point", "", map[string]interface{}{
		"protectVmInstanceRecoveryPoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceStackFromResource gets ResourceStackFromResource by uuid
func (cli *ZSClient) GetResourceStackFromResource(ctx context.Context) (*view.GetResourceStackFromResourceView, error) {
	var resp view.GetResourceStackFromResourceView
	if err := cli.GetWithRespKey(ctx, "v1/cloudformation/resources/stack", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryPhysicalDriveSelfTestHistory queries PhysicalDriveSelfTestHistory list
func (cli *ZSClient) QueryPhysicalDriveSelfTestHistory(ctx context.Context, params *param.QueryParam) ([]view.PhysicalDriveSmartSelfTestHistoryInventoryView, error) {
	var resp []view.PhysicalDriveSmartSelfTestHistoryInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/local-raid/physical-drives/self-test", params, &resp)
}

func (cli *ZSClient) GetPhysicalDriveSelfTestHistory(ctx context.Context, uuid string) (*view.PhysicalDriveSmartSelfTestHistoryInventoryView, error) {
	var resp view.PhysicalDriveSmartSelfTestHistoryInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/local-raid/physical-drives", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePhysicalDriveSelfTestHistory Pagination
func (cli *ZSClient) PagePhysicalDriveSelfTestHistory(ctx context.Context, params *param.QueryParam) ([]view.PhysicalDriveSmartSelfTestHistoryInventoryView, int, error) {
	var physicalDriveSelfTestHistories []view.PhysicalDriveSmartSelfTestHistoryInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/local-raid/physical-drives/self-test", params, &physicalDriveSelfTestHistories)
	return physicalDriveSelfTestHistories, total, err
}

// GetClusterHostNetworkFacts gets ClusterHostNetworkFacts by uuid
func (cli *ZSClient) GetClusterHostNetworkFacts(ctx context.Context, uuid string) (*view.GetClusterHostNetworkFactsView, error) {
	var resp view.GetClusterHostNetworkFactsView
	if err := cli.GetWithRespKey(ctx, "v1/cluster/hosts-network-facts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ParseOvf operates on ParseOvf
func (cli *ZSClient) ParseOvf(ctx context.Context, params param.ParseOvfParam) (*view.ParseOvfView, error) {
	resp := view.ParseOvfView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/ovf/parse"), "ovfInfo", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFaultToleranceVm queries FaultToleranceVm list
func (cli *ZSClient) QueryFaultToleranceVm(ctx context.Context, params *param.QueryParam) ([]view.FaultToleranceVmGroupInventoryView, error) {
	var resp []view.FaultToleranceVmGroupInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/fault-tolerance", params, &resp)
}

func (cli *ZSClient) GetFaultToleranceVm(ctx context.Context, uuid string) (*view.FaultToleranceVmGroupInventoryView, error) {
	var resp view.FaultToleranceVmGroupInventoryView
	if err := cli.Get(ctx, "v1/vm-instances/fault-tolerance", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFaultToleranceVm Pagination
func (cli *ZSClient) PageFaultToleranceVm(ctx context.Context, params *param.QueryParam) ([]view.FaultToleranceVmGroupInventoryView, int, error) {
	var faultToleranceVms []view.FaultToleranceVmGroupInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/fault-tolerance", params, &faultToleranceVms)
	return faultToleranceVms, total, err
}

// AddSchedulerJobGroupToSchedulerTrigger adds SchedulerJobGroupToSchedulerTrigger
func (cli *ZSClient) AddSchedulerJobGroupToSchedulerTrigger(ctx context.Context, schedulerJobGroupUuid, schedulerTriggerUuid string, params param.AddSchedulerJobGroupToSchedulerTriggerParam) (*view.SchedulerJobGroupSchedulerTriggerRefInventoryView, error) {
	resp := view.SchedulerJobGroupSchedulerTriggerRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/scheduler/jobgroups/%s/scheduler/triggers/%s", schedulerJobGroupUuid, schedulerTriggerUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSharedBlockCandidate gets SharedBlockCandidate by uuid
func (cli *ZSClient) GetSharedBlockCandidate(ctx context.Context) (*view.GetSharedBlockCandidateView, error) {
	var resp view.GetSharedBlockCandidateView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/sharedblockgroup/sharedblock-candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReclaimSpaceFromImageStore operates on ReclaimSpaceFromImageStore
func (cli *ZSClient) ReclaimSpaceFromImageStore(ctx context.Context, uuid string, params param.ReclaimSpaceFromImageStoreParam) (*view.ReclaimSpaceFromImageStoreEventView, error) {
	resp := view.ReclaimSpaceFromImageStoreEventView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage/image-store", uuid, "actions", "gcResult", map[string]interface{}{
		"reclaimSpaceFromImageStore": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UploadSoftwarePackage operates on UploadSoftwarePackage
func (cli *ZSClient) UploadSoftwarePackage(ctx context.Context, params param.UploadSoftwarePackageParam) (*view.SoftwarePackageInventoryView, error) {
	resp := view.SoftwarePackageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/software-packages/upload"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UploadSoftwarePackageAsync Async
func (cli *ZSClient) UploadSoftwarePackageAsync(ctx context.Context, params param.UploadSoftwarePackageParam) (string, error) {

	resource := "software-packages/upload"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetAllEventMetadata gets AllEventMetadata by uuid
func (cli *ZSClient) GetAllEventMetadata(ctx context.Context) (*view.GetAllEventMetadataView, error) {
	var resp view.GetAllEventMetadataView
	if err := cli.GetWithRespKey(ctx, "v1/zwatch/events/meta-data", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmForAttachingIso gets CandidateVmForAttachingIso by uuid
func (cli *ZSClient) GetCandidateVmForAttachingIso(ctx context.Context, uuid string) (*view.VmInstanceInventoryView, error) {
	var resp view.VmInstanceInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/images/iso", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBonding deletes Bonding
func (cli *ZSClient) DeleteBonding(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/hosts/bondings", uuid, string(deleteMode))
}

// AttachDataVolumeToVm operates on DataVolumeToVm
func (cli *ZSClient) AttachDataVolumeToVm(ctx context.Context, volumeUuid, vmInstanceUuid string, params param.AttachDataVolumeToVmParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/volumes/%s/vm-instances/%s", volumeUuid, vmInstanceUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDataVolume deletes DataVolume
func (cli *ZSClient) DeleteDataVolume(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/volumes", uuid, string(deleteMode))
}

// DeleteVmNicFromSecurityGroup deletes VmNicFromSecurityGroup
func (cli *ZSClient) DeleteVmNicFromSecurityGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/security-groups", uuid, string(deleteMode))
}

// UpdateTag updates Tag
func (cli *ZSClient) UpdateTag(ctx context.Context, uuid string, params param.UpdateTagParam) (*view.TagPatternInventoryView, error) {
	resp := view.TagPatternInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/tags", uuid, "actions", "inventory", map[string]interface{}{
		"updateTag": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUploadImageJobDetails gets UploadImageJobDetails by uuid
func (cli *ZSClient) GetUploadImageJobDetails(ctx context.Context, uuid string) (*view.GetUploadImageJobDetailsView, error) {
	var resp view.GetUploadImageJobDetailsView
	if err := cli.GetWithRespKey(ctx, "v1/images/upload-job/details", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachIscsiServerFromCluster operates on IscsiServerFromCluster
func (cli *ZSClient) DetachIscsiServerFromCluster(ctx context.Context, clusterUuid string, uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/clusters", clusterUuid, fmt.Sprintf("storage-devices/iscsi/servers/%s", uuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// SetVolumeQos operates on VolumeQos
func (cli *ZSClient) SetVolumeQos(ctx context.Context, uuid string, params param.SetVolumeQosParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volumes", uuid, "actions", "inventory", map[string]interface{}{
		"setVolumeQos": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTemplatedVmInstanceFromVmInstance creates TemplatedVmInstanceFromVmInstance
func (cli *ZSClient) CreateTemplatedVmInstanceFromVmInstance(ctx context.Context, vmInstanceUuid string, params param.CreateTemplatedVmInstanceFromVmInstanceParam) (*view.CreateTemplatedVmInstanceFromVmInstanceEventView, error) {
	resp := view.CreateTemplatedVmInstanceFromVmInstanceEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/%s/create-templated-vmInstance", vmInstanceUuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeCapabilities gets VolumeCapabilities by uuid
func (cli *ZSClient) GetVolumeCapabilities(ctx context.Context, uuid string) (*view.GetVolumeCapabilitiesView, error) {
	var resp view.GetVolumeCapabilitiesView
	if err := cli.GetWithRespKey(ctx, "v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachVRouterRouteTableToVRouter operates on VRouterRouteTableToVRouter
func (cli *ZSClient) AttachVRouterRouteTableToVRouter(ctx context.Context, routeTableUuid string, params param.AttachVRouterRouteTableToVRouterParam) (*view.VRouterRouteTableInventoryView, error) {
	resp := view.VRouterRouteTableInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vrouter-route-tables/%s/attach", routeTableUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVxlanVtep creates VxlanVtep
func (cli *ZSClient) CreateVxlanVtep(ctx context.Context, params param.CreateVxlanVtepParam) (*view.VtepInventoryView, error) {
	resp := view.VtepInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l2-networks/vxlan/vteps"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddMdevDeviceSpecToVmInstance adds MdevDeviceSpecToVmInstance
func (cli *ZSClient) AddMdevDeviceSpecToVmInstance(ctx context.Context, mdevSpecUuid, vmInstanceUuid string, params param.AddMdevDeviceSpecToVmInstanceParam) (*view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	resp := view.VmInstanceMdevDeviceSpecRefInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/mdev-device-specs/%s/vm-instances/%s", mdevSpecUuid, vmInstanceUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachScsiLunFromVmInstance operates on ScsiLunFromVmInstance
func (cli *ZSClient) DetachScsiLunFromVmInstance(ctx context.Context, vmInstanceUuid string, uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, fmt.Sprintf("scsi-lun/%s", uuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// EnableCdpTask operates on EnableCdpTask
func (cli *ZSClient) EnableCdpTask(ctx context.Context, uuid string, params param.EnableCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	resp := view.CdpTaskInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cdp-task/enable/%s", uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableCdpTaskAsync Async
func (cli *ZSClient) EnableCdpTaskAsync(ctx context.Context, params param.EnableCdpTaskParam) (string, error) {

	resource := "cdp-task/enable/{uuid}"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// RegisterLicenseRequestedApplication operates on LicenseRequestedApplication
func (cli *ZSClient) RegisterLicenseRequestedApplication(ctx context.Context, params param.RegisterLicenseRequestedApplicationParam) (*view.RegisterLicenseRequestedApplicationEventView, error) {
	resp := view.RegisterLicenseRequestedApplicationEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/licenses/applications"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmInstanceHaLevel operates on VmInstanceHaLevel
func (cli *ZSClient) SetVmInstanceHaLevel(ctx context.Context, uuid string, params param.SetVmInstanceHaLevelParam) (*view.SetVmInstanceHaLevelEventView, error) {
	resp := view.SetVmInstanceHaLevelEventView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-instances/%s/ha-levels", uuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveVRouterNetworksFromFlowMeter removes VRouterNetworksFromFlowMeter
func (cli *ZSClient) RemoveVRouterNetworksFromFlowMeter(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/flowmeters/networks", uuid, string(deleteMode))
}

