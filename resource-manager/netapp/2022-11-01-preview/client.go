package v2022_11_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/backuppolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/backups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/backupvaults"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/capacitypools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/filelocks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/groupidlistforldapuser"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/netappaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/netappresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/poolchange"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/resetcifspassword"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/restore"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/snapshotpolicy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/snapshotpolicylistvolumes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/snapshots"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/subvolumes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/volumegroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/volumequotarules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/volumes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/volumesrelocation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/volumesreplication"
	"github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/volumesrevert"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	BackupPolicy              *backuppolicy.BackupPolicyClient
	BackupVaults              *backupvaults.BackupVaultsClient
	Backups                   *backups.BackupsClient
	CapacityPools             *capacitypools.CapacityPoolsClient
	FileLocks                 *filelocks.FileLocksClient
	GroupIdListForLDAPUser    *groupidlistforldapuser.GroupIdListForLDAPUserClient
	NetAppAccounts            *netappaccounts.NetAppAccountsClient
	NetAppResource            *netappresource.NetAppResourceClient
	PoolChange                *poolchange.PoolChangeClient
	ResetCifsPassword         *resetcifspassword.ResetCifsPasswordClient
	Restore                   *restore.RestoreClient
	SnapshotPolicy            *snapshotpolicy.SnapshotPolicyClient
	SnapshotPolicyListVolumes *snapshotpolicylistvolumes.SnapshotPolicyListVolumesClient
	Snapshots                 *snapshots.SnapshotsClient
	SubVolumes                *subvolumes.SubVolumesClient
	VolumeGroups              *volumegroups.VolumeGroupsClient
	VolumeQuotaRules          *volumequotarules.VolumeQuotaRulesClient
	Volumes                   *volumes.VolumesClient
	VolumesRelocation         *volumesrelocation.VolumesRelocationClient
	VolumesReplication        *volumesreplication.VolumesReplicationClient
	VolumesRevert             *volumesrevert.VolumesRevertClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	backupPolicyClient, err := backuppolicy.NewBackupPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupPolicy client: %+v", err)
	}
	configureFunc(backupPolicyClient.Client)

	backupVaultsClient, err := backupvaults.NewBackupVaultsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupVaults client: %+v", err)
	}
	configureFunc(backupVaultsClient.Client)

	backupsClient, err := backups.NewBackupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Backups client: %+v", err)
	}
	configureFunc(backupsClient.Client)

	capacityPoolsClient, err := capacitypools.NewCapacityPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapacityPools client: %+v", err)
	}
	configureFunc(capacityPoolsClient.Client)

	fileLocksClient, err := filelocks.NewFileLocksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FileLocks client: %+v", err)
	}
	configureFunc(fileLocksClient.Client)

	groupIdListForLDAPUserClient, err := groupidlistforldapuser.NewGroupIdListForLDAPUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupIdListForLDAPUser client: %+v", err)
	}
	configureFunc(groupIdListForLDAPUserClient.Client)

	netAppAccountsClient, err := netappaccounts.NewNetAppAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetAppAccounts client: %+v", err)
	}
	configureFunc(netAppAccountsClient.Client)

	netAppResourceClient, err := netappresource.NewNetAppResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetAppResource client: %+v", err)
	}
	configureFunc(netAppResourceClient.Client)

	poolChangeClient, err := poolchange.NewPoolChangeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PoolChange client: %+v", err)
	}
	configureFunc(poolChangeClient.Client)

	resetCifsPasswordClient, err := resetcifspassword.NewResetCifsPasswordClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResetCifsPassword client: %+v", err)
	}
	configureFunc(resetCifsPasswordClient.Client)

	restoreClient, err := restore.NewRestoreClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Restore client: %+v", err)
	}
	configureFunc(restoreClient.Client)

	snapshotPolicyClient, err := snapshotpolicy.NewSnapshotPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SnapshotPolicy client: %+v", err)
	}
	configureFunc(snapshotPolicyClient.Client)

	snapshotPolicyListVolumesClient, err := snapshotpolicylistvolumes.NewSnapshotPolicyListVolumesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SnapshotPolicyListVolumes client: %+v", err)
	}
	configureFunc(snapshotPolicyListVolumesClient.Client)

	snapshotsClient, err := snapshots.NewSnapshotsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Snapshots client: %+v", err)
	}
	configureFunc(snapshotsClient.Client)

	subVolumesClient, err := subvolumes.NewSubVolumesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubVolumes client: %+v", err)
	}
	configureFunc(subVolumesClient.Client)

	volumeGroupsClient, err := volumegroups.NewVolumeGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VolumeGroups client: %+v", err)
	}
	configureFunc(volumeGroupsClient.Client)

	volumeQuotaRulesClient, err := volumequotarules.NewVolumeQuotaRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VolumeQuotaRules client: %+v", err)
	}
	configureFunc(volumeQuotaRulesClient.Client)

	volumesClient, err := volumes.NewVolumesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Volumes client: %+v", err)
	}
	configureFunc(volumesClient.Client)

	volumesRelocationClient, err := volumesrelocation.NewVolumesRelocationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VolumesRelocation client: %+v", err)
	}
	configureFunc(volumesRelocationClient.Client)

	volumesReplicationClient, err := volumesreplication.NewVolumesReplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VolumesReplication client: %+v", err)
	}
	configureFunc(volumesReplicationClient.Client)

	volumesRevertClient, err := volumesrevert.NewVolumesRevertClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VolumesRevert client: %+v", err)
	}
	configureFunc(volumesRevertClient.Client)

	return &Client{
		BackupPolicy:              backupPolicyClient,
		BackupVaults:              backupVaultsClient,
		Backups:                   backupsClient,
		CapacityPools:             capacityPoolsClient,
		FileLocks:                 fileLocksClient,
		GroupIdListForLDAPUser:    groupIdListForLDAPUserClient,
		NetAppAccounts:            netAppAccountsClient,
		NetAppResource:            netAppResourceClient,
		PoolChange:                poolChangeClient,
		ResetCifsPassword:         resetCifsPasswordClient,
		Restore:                   restoreClient,
		SnapshotPolicy:            snapshotPolicyClient,
		SnapshotPolicyListVolumes: snapshotPolicyListVolumesClient,
		Snapshots:                 snapshotsClient,
		SubVolumes:                subVolumesClient,
		VolumeGroups:              volumeGroupsClient,
		VolumeQuotaRules:          volumeQuotaRulesClient,
		Volumes:                   volumesClient,
		VolumesRelocation:         volumesRelocationClient,
		VolumesReplication:        volumesReplicationClient,
		VolumesRevert:             volumesRevertClient,
	}, nil
}
