//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package netapp

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/netapp/mgmt/2021-10-01/netapp"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActiveDirectoryStatus = original.ActiveDirectoryStatus

const (
	ActiveDirectoryStatusCreated  ActiveDirectoryStatus = original.ActiveDirectoryStatusCreated
	ActiveDirectoryStatusDeleted  ActiveDirectoryStatus = original.ActiveDirectoryStatusDeleted
	ActiveDirectoryStatusError    ActiveDirectoryStatus = original.ActiveDirectoryStatusError
	ActiveDirectoryStatusInUse    ActiveDirectoryStatus = original.ActiveDirectoryStatusInUse
	ActiveDirectoryStatusUpdating ActiveDirectoryStatus = original.ActiveDirectoryStatusUpdating
)

type ApplicationType = original.ApplicationType

const (
	ApplicationTypeSAPHANA ApplicationType = original.ApplicationTypeSAPHANA
)

type AvsDataStore = original.AvsDataStore

const (
	AvsDataStoreDisabled AvsDataStore = original.AvsDataStoreDisabled
	AvsDataStoreEnabled  AvsDataStore = original.AvsDataStoreEnabled
)

type BackupType = original.BackupType

const (
	BackupTypeManual    BackupType = original.BackupTypeManual
	BackupTypeScheduled BackupType = original.BackupTypeScheduled
)

type CheckNameResourceTypes = original.CheckNameResourceTypes

const (
	CheckNameResourceTypesMicrosoftNetAppnetAppAccounts                              CheckNameResourceTypes = original.CheckNameResourceTypesMicrosoftNetAppnetAppAccounts
	CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools                 CheckNameResourceTypes = original.CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools
	CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes          CheckNameResourceTypes = original.CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes
	CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots CheckNameResourceTypes = original.CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots
)

type CheckQuotaNameResourceTypes = original.CheckQuotaNameResourceTypes

const (
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccounts                              CheckQuotaNameResourceTypes = original.CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccounts
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools                 CheckQuotaNameResourceTypes = original.CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes          CheckQuotaNameResourceTypes = original.CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots CheckQuotaNameResourceTypes = original.CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots
)

type ChownMode = original.ChownMode

const (
	ChownModeRestricted   ChownMode = original.ChownModeRestricted
	ChownModeUnrestricted ChownMode = original.ChownModeUnrestricted
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type EnableSubvolumes = original.EnableSubvolumes

const (
	EnableSubvolumesDisabled EnableSubvolumes = original.EnableSubvolumesDisabled
	EnableSubvolumesEnabled  EnableSubvolumes = original.EnableSubvolumesEnabled
)

type EncryptionType = original.EncryptionType

const (
	EncryptionTypeDouble EncryptionType = original.EncryptionTypeDouble
	EncryptionTypeSingle EncryptionType = original.EncryptionTypeSingle
)

type EndpointType = original.EndpointType

const (
	EndpointTypeDst EndpointType = original.EndpointTypeDst
	EndpointTypeSrc EndpointType = original.EndpointTypeSrc
)

type InAvailabilityReasonType = original.InAvailabilityReasonType

const (
	InAvailabilityReasonTypeAlreadyExists InAvailabilityReasonType = original.InAvailabilityReasonTypeAlreadyExists
	InAvailabilityReasonTypeInvalid       InAvailabilityReasonType = original.InAvailabilityReasonTypeInvalid
)

type MetricAggregationType = original.MetricAggregationType

const (
	MetricAggregationTypeAverage MetricAggregationType = original.MetricAggregationTypeAverage
)

type MirrorState = original.MirrorState

const (
	MirrorStateBroken        MirrorState = original.MirrorStateBroken
	MirrorStateMirrored      MirrorState = original.MirrorStateMirrored
	MirrorStateUninitialized MirrorState = original.MirrorStateUninitialized
)

type NetworkFeatures = original.NetworkFeatures

const (
	NetworkFeaturesBasic    NetworkFeatures = original.NetworkFeaturesBasic
	NetworkFeaturesStandard NetworkFeatures = original.NetworkFeaturesStandard
)

type QosType = original.QosType

const (
	QosTypeAuto   QosType = original.QosTypeAuto
	QosTypeManual QosType = original.QosTypeManual
)

type RelationshipStatus = original.RelationshipStatus

const (
	RelationshipStatusIdle         RelationshipStatus = original.RelationshipStatusIdle
	RelationshipStatusTransferring RelationshipStatus = original.RelationshipStatusTransferring
)

type ReplicationSchedule = original.ReplicationSchedule

const (
	ReplicationSchedule10minutely ReplicationSchedule = original.ReplicationSchedule10minutely
	ReplicationScheduleDaily      ReplicationSchedule = original.ReplicationScheduleDaily
	ReplicationScheduleHourly     ReplicationSchedule = original.ReplicationScheduleHourly
)

type SecurityStyle = original.SecurityStyle

const (
	SecurityStyleNtfs SecurityStyle = original.SecurityStyleNtfs
	SecurityStyleUnix SecurityStyle = original.SecurityStyleUnix
)

type ServiceLevel = original.ServiceLevel

const (
	ServiceLevelPremium     ServiceLevel = original.ServiceLevelPremium
	ServiceLevelStandard    ServiceLevel = original.ServiceLevelStandard
	ServiceLevelStandardZRS ServiceLevel = original.ServiceLevelStandardZRS
	ServiceLevelUltra       ServiceLevel = original.ServiceLevelUltra
)

type VolumeStorageToNetworkProximity = original.VolumeStorageToNetworkProximity

const (
	VolumeStorageToNetworkProximityDefault VolumeStorageToNetworkProximity = original.VolumeStorageToNetworkProximityDefault
	VolumeStorageToNetworkProximityT1      VolumeStorageToNetworkProximity = original.VolumeStorageToNetworkProximityT1
	VolumeStorageToNetworkProximityT2      VolumeStorageToNetworkProximity = original.VolumeStorageToNetworkProximityT2
)

type Account = original.Account
type AccountBackupsClient = original.AccountBackupsClient
type AccountBackupsDeleteFuture = original.AccountBackupsDeleteFuture
type AccountEncryption = original.AccountEncryption
type AccountList = original.AccountList
type AccountListIterator = original.AccountListIterator
type AccountListPage = original.AccountListPage
type AccountPatch = original.AccountPatch
type AccountProperties = original.AccountProperties
type AccountsClient = original.AccountsClient
type AccountsCreateOrUpdateFuture = original.AccountsCreateOrUpdateFuture
type AccountsDeleteFuture = original.AccountsDeleteFuture
type AccountsUpdateFuture = original.AccountsUpdateFuture
type ActiveDirectory = original.ActiveDirectory
type AuthorizeRequest = original.AuthorizeRequest
type AzureEntityResource = original.AzureEntityResource
type Backup = original.Backup
type BackupPatch = original.BackupPatch
type BackupPoliciesClient = original.BackupPoliciesClient
type BackupPoliciesCreateFuture = original.BackupPoliciesCreateFuture
type BackupPoliciesDeleteFuture = original.BackupPoliciesDeleteFuture
type BackupPoliciesList = original.BackupPoliciesList
type BackupPoliciesUpdateFuture = original.BackupPoliciesUpdateFuture
type BackupPolicy = original.BackupPolicy
type BackupPolicyDetails = original.BackupPolicyDetails
type BackupPolicyPatch = original.BackupPolicyPatch
type BackupPolicyProperties = original.BackupPolicyProperties
type BackupProperties = original.BackupProperties
type BackupStatus = original.BackupStatus
type BackupsClient = original.BackupsClient
type BackupsCreateFuture = original.BackupsCreateFuture
type BackupsDeleteFuture = original.BackupsDeleteFuture
type BackupsList = original.BackupsList
type BackupsUpdateFuture = original.BackupsUpdateFuture
type BaseClient = original.BaseClient
type BreakReplicationRequest = original.BreakReplicationRequest
type CapacityPool = original.CapacityPool
type CapacityPoolList = original.CapacityPoolList
type CapacityPoolListIterator = original.CapacityPoolListIterator
type CapacityPoolListPage = original.CapacityPoolListPage
type CapacityPoolPatch = original.CapacityPoolPatch
type CheckAvailabilityResponse = original.CheckAvailabilityResponse
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type DailySchedule = original.DailySchedule
type Dimension = original.Dimension
type ExportPolicyRule = original.ExportPolicyRule
type FilePathAvailabilityRequest = original.FilePathAvailabilityRequest
type HourlySchedule = original.HourlySchedule
type LdapSearchScopeOpt = original.LdapSearchScopeOpt
type LogSpecification = original.LogSpecification
type MetricSpecification = original.MetricSpecification
type MonthlySchedule = original.MonthlySchedule
type MountTarget = original.MountTarget
type MountTargetProperties = original.MountTargetProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type PlacementKeyValuePairs = original.PlacementKeyValuePairs
type PoolChangeRequest = original.PoolChangeRequest
type PoolPatchProperties = original.PoolPatchProperties
type PoolProperties = original.PoolProperties
type PoolsClient = original.PoolsClient
type PoolsCreateOrUpdateFuture = original.PoolsCreateOrUpdateFuture
type PoolsDeleteFuture = original.PoolsDeleteFuture
type PoolsUpdateFuture = original.PoolsUpdateFuture
type ProxyResource = original.ProxyResource
type QuotaAvailabilityRequest = original.QuotaAvailabilityRequest
type ReplicationObject = original.ReplicationObject
type ReplicationStatus = original.ReplicationStatus
type Resource = original.Resource
type ResourceClient = original.ResourceClient
type ResourceIdentity = original.ResourceIdentity
type ResourceNameAvailabilityRequest = original.ResourceNameAvailabilityRequest
type ResourceQuotaLimitsClient = original.ResourceQuotaLimitsClient
type RestoreStatus = original.RestoreStatus
type ServiceSpecification = original.ServiceSpecification
type Snapshot = original.Snapshot
type SnapshotPoliciesClient = original.SnapshotPoliciesClient
type SnapshotPoliciesDeleteFuture = original.SnapshotPoliciesDeleteFuture
type SnapshotPoliciesList = original.SnapshotPoliciesList
type SnapshotPoliciesUpdateFuture = original.SnapshotPoliciesUpdateFuture
type SnapshotPolicy = original.SnapshotPolicy
type SnapshotPolicyDetails = original.SnapshotPolicyDetails
type SnapshotPolicyPatch = original.SnapshotPolicyPatch
type SnapshotPolicyProperties = original.SnapshotPolicyProperties
type SnapshotPolicyVolumeList = original.SnapshotPolicyVolumeList
type SnapshotProperties = original.SnapshotProperties
type SnapshotRestoreFiles = original.SnapshotRestoreFiles
type SnapshotsClient = original.SnapshotsClient
type SnapshotsCreateFuture = original.SnapshotsCreateFuture
type SnapshotsDeleteFuture = original.SnapshotsDeleteFuture
type SnapshotsList = original.SnapshotsList
type SnapshotsRestoreFilesFuture = original.SnapshotsRestoreFilesFuture
type SnapshotsUpdateFuture = original.SnapshotsUpdateFuture
type SubscriptionQuotaItem = original.SubscriptionQuotaItem
type SubscriptionQuotaItemList = original.SubscriptionQuotaItemList
type SubscriptionQuotaItemProperties = original.SubscriptionQuotaItemProperties
type SubvolumeInfo = original.SubvolumeInfo
type SubvolumeModel = original.SubvolumeModel
type SubvolumeModelProperties = original.SubvolumeModelProperties
type SubvolumePatchParams = original.SubvolumePatchParams
type SubvolumePatchRequest = original.SubvolumePatchRequest
type SubvolumeProperties = original.SubvolumeProperties
type SubvolumesClient = original.SubvolumesClient
type SubvolumesCreateFuture = original.SubvolumesCreateFuture
type SubvolumesDeleteFuture = original.SubvolumesDeleteFuture
type SubvolumesGetMetadataFuture = original.SubvolumesGetMetadataFuture
type SubvolumesList = original.SubvolumesList
type SubvolumesListIterator = original.SubvolumesListIterator
type SubvolumesListPage = original.SubvolumesListPage
type SubvolumesUpdateFuture = original.SubvolumesUpdateFuture
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource
type Vault = original.Vault
type VaultList = original.VaultList
type VaultProperties = original.VaultProperties
type VaultsClient = original.VaultsClient
type Volume = original.Volume
type VolumeBackupProperties = original.VolumeBackupProperties
type VolumeBackups = original.VolumeBackups
type VolumeGroup = original.VolumeGroup
type VolumeGroupDetails = original.VolumeGroupDetails
type VolumeGroupList = original.VolumeGroupList
type VolumeGroupListProperties = original.VolumeGroupListProperties
type VolumeGroupMetaData = original.VolumeGroupMetaData
type VolumeGroupProperties = original.VolumeGroupProperties
type VolumeGroupVolumeProperties = original.VolumeGroupVolumeProperties
type VolumeGroupsClient = original.VolumeGroupsClient
type VolumeGroupsCreateFuture = original.VolumeGroupsCreateFuture
type VolumeGroupsDeleteFuture = original.VolumeGroupsDeleteFuture
type VolumeList = original.VolumeList
type VolumeListIterator = original.VolumeListIterator
type VolumeListPage = original.VolumeListPage
type VolumePatch = original.VolumePatch
type VolumePatchProperties = original.VolumePatchProperties
type VolumePatchPropertiesDataProtection = original.VolumePatchPropertiesDataProtection
type VolumePatchPropertiesExportPolicy = original.VolumePatchPropertiesExportPolicy
type VolumeProperties = original.VolumeProperties
type VolumePropertiesDataProtection = original.VolumePropertiesDataProtection
type VolumePropertiesExportPolicy = original.VolumePropertiesExportPolicy
type VolumeRevert = original.VolumeRevert
type VolumeSnapshotProperties = original.VolumeSnapshotProperties
type VolumesAuthorizeReplicationFuture = original.VolumesAuthorizeReplicationFuture
type VolumesBreakReplicationFuture = original.VolumesBreakReplicationFuture
type VolumesClient = original.VolumesClient
type VolumesCreateOrUpdateFuture = original.VolumesCreateOrUpdateFuture
type VolumesDeleteFuture = original.VolumesDeleteFuture
type VolumesDeleteReplicationFuture = original.VolumesDeleteReplicationFuture
type VolumesPoolChangeFuture = original.VolumesPoolChangeFuture
type VolumesReInitializeReplicationFuture = original.VolumesReInitializeReplicationFuture
type VolumesResyncReplicationFuture = original.VolumesResyncReplicationFuture
type VolumesRevertFuture = original.VolumesRevertFuture
type VolumesUpdateFuture = original.VolumesUpdateFuture
type WeeklySchedule = original.WeeklySchedule

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountBackupsClient(subscriptionID string) AccountBackupsClient {
	return original.NewAccountBackupsClient(subscriptionID)
}
func NewAccountBackupsClientWithBaseURI(baseURI string, subscriptionID string) AccountBackupsClient {
	return original.NewAccountBackupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAccountListIterator(page AccountListPage) AccountListIterator {
	return original.NewAccountListIterator(page)
}
func NewAccountListPage(cur AccountList, getNextPage func(context.Context, AccountList) (AccountList, error)) AccountListPage {
	return original.NewAccountListPage(cur, getNextPage)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewBackupPoliciesClient(subscriptionID string) BackupPoliciesClient {
	return original.NewBackupPoliciesClient(subscriptionID)
}
func NewBackupPoliciesClientWithBaseURI(baseURI string, subscriptionID string) BackupPoliciesClient {
	return original.NewBackupPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewBackupsClient(subscriptionID string) BackupsClient {
	return original.NewBackupsClient(subscriptionID)
}
func NewBackupsClientWithBaseURI(baseURI string, subscriptionID string) BackupsClient {
	return original.NewBackupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewCapacityPoolListIterator(page CapacityPoolListPage) CapacityPoolListIterator {
	return original.NewCapacityPoolListIterator(page)
}
func NewCapacityPoolListPage(cur CapacityPoolList, getNextPage func(context.Context, CapacityPoolList) (CapacityPoolList, error)) CapacityPoolListPage {
	return original.NewCapacityPoolListPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPoolsClient(subscriptionID string) PoolsClient {
	return original.NewPoolsClient(subscriptionID)
}
func NewPoolsClientWithBaseURI(baseURI string, subscriptionID string) PoolsClient {
	return original.NewPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceClient(subscriptionID string) ResourceClient {
	return original.NewResourceClient(subscriptionID)
}
func NewResourceClientWithBaseURI(baseURI string, subscriptionID string) ResourceClient {
	return original.NewResourceClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceQuotaLimitsClient(subscriptionID string) ResourceQuotaLimitsClient {
	return original.NewResourceQuotaLimitsClient(subscriptionID)
}
func NewResourceQuotaLimitsClientWithBaseURI(baseURI string, subscriptionID string) ResourceQuotaLimitsClient {
	return original.NewResourceQuotaLimitsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSnapshotPoliciesClient(subscriptionID string) SnapshotPoliciesClient {
	return original.NewSnapshotPoliciesClient(subscriptionID)
}
func NewSnapshotPoliciesClientWithBaseURI(baseURI string, subscriptionID string) SnapshotPoliciesClient {
	return original.NewSnapshotPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSnapshotsClient(subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClient(subscriptionID)
}
func NewSnapshotsClientWithBaseURI(baseURI string, subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubvolumesClient(subscriptionID string) SubvolumesClient {
	return original.NewSubvolumesClient(subscriptionID)
}
func NewSubvolumesClientWithBaseURI(baseURI string, subscriptionID string) SubvolumesClient {
	return original.NewSubvolumesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubvolumesListIterator(page SubvolumesListPage) SubvolumesListIterator {
	return original.NewSubvolumesListIterator(page)
}
func NewSubvolumesListPage(cur SubvolumesList, getNextPage func(context.Context, SubvolumesList) (SubvolumesList, error)) SubvolumesListPage {
	return original.NewSubvolumesListPage(cur, getNextPage)
}
func NewVaultsClient(subscriptionID string) VaultsClient {
	return original.NewVaultsClient(subscriptionID)
}
func NewVaultsClientWithBaseURI(baseURI string, subscriptionID string) VaultsClient {
	return original.NewVaultsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVolumeGroupsClient(subscriptionID string) VolumeGroupsClient {
	return original.NewVolumeGroupsClient(subscriptionID)
}
func NewVolumeGroupsClientWithBaseURI(baseURI string, subscriptionID string) VolumeGroupsClient {
	return original.NewVolumeGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVolumeListIterator(page VolumeListPage) VolumeListIterator {
	return original.NewVolumeListIterator(page)
}
func NewVolumeListPage(cur VolumeList, getNextPage func(context.Context, VolumeList) (VolumeList, error)) VolumeListPage {
	return original.NewVolumeListPage(cur, getNextPage)
}
func NewVolumesClient(subscriptionID string) VolumesClient {
	return original.NewVolumesClient(subscriptionID)
}
func NewVolumesClientWithBaseURI(baseURI string, subscriptionID string) VolumesClient {
	return original.NewVolumesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActiveDirectoryStatusValues() []ActiveDirectoryStatus {
	return original.PossibleActiveDirectoryStatusValues()
}
func PossibleApplicationTypeValues() []ApplicationType {
	return original.PossibleApplicationTypeValues()
}
func PossibleAvsDataStoreValues() []AvsDataStore {
	return original.PossibleAvsDataStoreValues()
}
func PossibleBackupTypeValues() []BackupType {
	return original.PossibleBackupTypeValues()
}
func PossibleCheckNameResourceTypesValues() []CheckNameResourceTypes {
	return original.PossibleCheckNameResourceTypesValues()
}
func PossibleCheckQuotaNameResourceTypesValues() []CheckQuotaNameResourceTypes {
	return original.PossibleCheckQuotaNameResourceTypesValues()
}
func PossibleChownModeValues() []ChownMode {
	return original.PossibleChownModeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleEnableSubvolumesValues() []EnableSubvolumes {
	return original.PossibleEnableSubvolumesValues()
}
func PossibleEncryptionTypeValues() []EncryptionType {
	return original.PossibleEncryptionTypeValues()
}
func PossibleEndpointTypeValues() []EndpointType {
	return original.PossibleEndpointTypeValues()
}
func PossibleInAvailabilityReasonTypeValues() []InAvailabilityReasonType {
	return original.PossibleInAvailabilityReasonTypeValues()
}
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return original.PossibleMetricAggregationTypeValues()
}
func PossibleMirrorStateValues() []MirrorState {
	return original.PossibleMirrorStateValues()
}
func PossibleNetworkFeaturesValues() []NetworkFeatures {
	return original.PossibleNetworkFeaturesValues()
}
func PossibleQosTypeValues() []QosType {
	return original.PossibleQosTypeValues()
}
func PossibleRelationshipStatusValues() []RelationshipStatus {
	return original.PossibleRelationshipStatusValues()
}
func PossibleReplicationScheduleValues() []ReplicationSchedule {
	return original.PossibleReplicationScheduleValues()
}
func PossibleSecurityStyleValues() []SecurityStyle {
	return original.PossibleSecurityStyleValues()
}
func PossibleServiceLevelValues() []ServiceLevel {
	return original.PossibleServiceLevelValues()
}
func PossibleVolumeStorageToNetworkProximityValues() []VolumeStorageToNetworkProximity {
	return original.PossibleVolumeStorageToNetworkProximityValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
