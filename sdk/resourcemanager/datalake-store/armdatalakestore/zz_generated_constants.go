//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakestore

const (
	moduleName    = "armdatalakestore"
	moduleVersion = "v0.3.1"
)

// CheckNameAvailabilityParametersType - The resource type. Note: This should not be set by the user, as the constant value
// is Microsoft.DataLakeStore/accounts
type CheckNameAvailabilityParametersType string

const (
	CheckNameAvailabilityParametersTypeMicrosoftDataLakeStoreAccounts CheckNameAvailabilityParametersType = "Microsoft.DataLakeStore/accounts"
)

// PossibleCheckNameAvailabilityParametersTypeValues returns the possible values for the CheckNameAvailabilityParametersType const type.
func PossibleCheckNameAvailabilityParametersTypeValues() []CheckNameAvailabilityParametersType {
	return []CheckNameAvailabilityParametersType{
		CheckNameAvailabilityParametersTypeMicrosoftDataLakeStoreAccounts,
	}
}

// ToPtr returns a *CheckNameAvailabilityParametersType pointing to the current value.
func (c CheckNameAvailabilityParametersType) ToPtr() *CheckNameAvailabilityParametersType {
	return &c
}

// DataLakeStoreAccountState - The state of the Data Lake Store account.
type DataLakeStoreAccountState string

const (
	DataLakeStoreAccountStateActive    DataLakeStoreAccountState = "Active"
	DataLakeStoreAccountStateSuspended DataLakeStoreAccountState = "Suspended"
)

// PossibleDataLakeStoreAccountStateValues returns the possible values for the DataLakeStoreAccountState const type.
func PossibleDataLakeStoreAccountStateValues() []DataLakeStoreAccountState {
	return []DataLakeStoreAccountState{
		DataLakeStoreAccountStateActive,
		DataLakeStoreAccountStateSuspended,
	}
}

// ToPtr returns a *DataLakeStoreAccountState pointing to the current value.
func (c DataLakeStoreAccountState) ToPtr() *DataLakeStoreAccountState {
	return &c
}

// DataLakeStoreAccountStatus - The provisioning status of the Data Lake Store account.
type DataLakeStoreAccountStatus string

const (
	DataLakeStoreAccountStatusFailed     DataLakeStoreAccountStatus = "Failed"
	DataLakeStoreAccountStatusCreating   DataLakeStoreAccountStatus = "Creating"
	DataLakeStoreAccountStatusRunning    DataLakeStoreAccountStatus = "Running"
	DataLakeStoreAccountStatusSucceeded  DataLakeStoreAccountStatus = "Succeeded"
	DataLakeStoreAccountStatusPatching   DataLakeStoreAccountStatus = "Patching"
	DataLakeStoreAccountStatusSuspending DataLakeStoreAccountStatus = "Suspending"
	DataLakeStoreAccountStatusResuming   DataLakeStoreAccountStatus = "Resuming"
	DataLakeStoreAccountStatusDeleting   DataLakeStoreAccountStatus = "Deleting"
	DataLakeStoreAccountStatusDeleted    DataLakeStoreAccountStatus = "Deleted"
	DataLakeStoreAccountStatusUndeleting DataLakeStoreAccountStatus = "Undeleting"
	DataLakeStoreAccountStatusCanceled   DataLakeStoreAccountStatus = "Canceled"
)

// PossibleDataLakeStoreAccountStatusValues returns the possible values for the DataLakeStoreAccountStatus const type.
func PossibleDataLakeStoreAccountStatusValues() []DataLakeStoreAccountStatus {
	return []DataLakeStoreAccountStatus{
		DataLakeStoreAccountStatusFailed,
		DataLakeStoreAccountStatusCreating,
		DataLakeStoreAccountStatusRunning,
		DataLakeStoreAccountStatusSucceeded,
		DataLakeStoreAccountStatusPatching,
		DataLakeStoreAccountStatusSuspending,
		DataLakeStoreAccountStatusResuming,
		DataLakeStoreAccountStatusDeleting,
		DataLakeStoreAccountStatusDeleted,
		DataLakeStoreAccountStatusUndeleting,
		DataLakeStoreAccountStatusCanceled,
	}
}

// ToPtr returns a *DataLakeStoreAccountStatus pointing to the current value.
func (c DataLakeStoreAccountStatus) ToPtr() *DataLakeStoreAccountStatus {
	return &c
}

// EncryptionConfigType - The type of encryption configuration being used. Currently the only supported types are 'UserManaged'
// and 'ServiceManaged'.
type EncryptionConfigType string

const (
	EncryptionConfigTypeUserManaged    EncryptionConfigType = "UserManaged"
	EncryptionConfigTypeServiceManaged EncryptionConfigType = "ServiceManaged"
)

// PossibleEncryptionConfigTypeValues returns the possible values for the EncryptionConfigType const type.
func PossibleEncryptionConfigTypeValues() []EncryptionConfigType {
	return []EncryptionConfigType{
		EncryptionConfigTypeUserManaged,
		EncryptionConfigTypeServiceManaged,
	}
}

// ToPtr returns a *EncryptionConfigType pointing to the current value.
func (c EncryptionConfigType) ToPtr() *EncryptionConfigType {
	return &c
}

// EncryptionProvisioningState - The current state of encryption provisioning for this Data Lake Store account.
type EncryptionProvisioningState string

const (
	EncryptionProvisioningStateCreating  EncryptionProvisioningState = "Creating"
	EncryptionProvisioningStateSucceeded EncryptionProvisioningState = "Succeeded"
)

// PossibleEncryptionProvisioningStateValues returns the possible values for the EncryptionProvisioningState const type.
func PossibleEncryptionProvisioningStateValues() []EncryptionProvisioningState {
	return []EncryptionProvisioningState{
		EncryptionProvisioningStateCreating,
		EncryptionProvisioningStateSucceeded,
	}
}

// ToPtr returns a *EncryptionProvisioningState pointing to the current value.
func (c EncryptionProvisioningState) ToPtr() *EncryptionProvisioningState {
	return &c
}

// EncryptionState - The current state of encryption for this Data Lake Store account.
type EncryptionState string

const (
	EncryptionStateEnabled  EncryptionState = "Enabled"
	EncryptionStateDisabled EncryptionState = "Disabled"
)

// PossibleEncryptionStateValues returns the possible values for the EncryptionState const type.
func PossibleEncryptionStateValues() []EncryptionState {
	return []EncryptionState{
		EncryptionStateEnabled,
		EncryptionStateDisabled,
	}
}

// ToPtr returns a *EncryptionState pointing to the current value.
func (c EncryptionState) ToPtr() *EncryptionState {
	return &c
}

// FirewallAllowAzureIPsState - The current state of allowing or disallowing IPs originating within Azure through the firewall.
// If the firewall is disabled, this is not enforced.
type FirewallAllowAzureIPsState string

const (
	FirewallAllowAzureIPsStateEnabled  FirewallAllowAzureIPsState = "Enabled"
	FirewallAllowAzureIPsStateDisabled FirewallAllowAzureIPsState = "Disabled"
)

// PossibleFirewallAllowAzureIPsStateValues returns the possible values for the FirewallAllowAzureIPsState const type.
func PossibleFirewallAllowAzureIPsStateValues() []FirewallAllowAzureIPsState {
	return []FirewallAllowAzureIPsState{
		FirewallAllowAzureIPsStateEnabled,
		FirewallAllowAzureIPsStateDisabled,
	}
}

// ToPtr returns a *FirewallAllowAzureIPsState pointing to the current value.
func (c FirewallAllowAzureIPsState) ToPtr() *FirewallAllowAzureIPsState {
	return &c
}

// FirewallState - The current state of the IP address firewall for this Data Lake Store account.
type FirewallState string

const (
	FirewallStateEnabled  FirewallState = "Enabled"
	FirewallStateDisabled FirewallState = "Disabled"
)

// PossibleFirewallStateValues returns the possible values for the FirewallState const type.
func PossibleFirewallStateValues() []FirewallState {
	return []FirewallState{
		FirewallStateEnabled,
		FirewallStateDisabled,
	}
}

// ToPtr returns a *FirewallState pointing to the current value.
func (c FirewallState) ToPtr() *FirewallState {
	return &c
}

// OperationOrigin - The intended executor of the operation.
type OperationOrigin string

const (
	OperationOriginSystem     OperationOrigin = "system"
	OperationOriginUser       OperationOrigin = "user"
	OperationOriginUserSystem OperationOrigin = "user,system"
)

// PossibleOperationOriginValues returns the possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{
		OperationOriginSystem,
		OperationOriginUser,
		OperationOriginUserSystem,
	}
}

// ToPtr returns a *OperationOrigin pointing to the current value.
func (c OperationOrigin) ToPtr() *OperationOrigin {
	return &c
}

// SubscriptionState - The subscription state.
type SubscriptionState string

const (
	SubscriptionStateDeleted      SubscriptionState = "Deleted"
	SubscriptionStateRegistered   SubscriptionState = "Registered"
	SubscriptionStateSuspended    SubscriptionState = "Suspended"
	SubscriptionStateUnregistered SubscriptionState = "Unregistered"
	SubscriptionStateWarned       SubscriptionState = "Warned"
)

// PossibleSubscriptionStateValues returns the possible values for the SubscriptionState const type.
func PossibleSubscriptionStateValues() []SubscriptionState {
	return []SubscriptionState{
		SubscriptionStateDeleted,
		SubscriptionStateRegistered,
		SubscriptionStateSuspended,
		SubscriptionStateUnregistered,
		SubscriptionStateWarned,
	}
}

// ToPtr returns a *SubscriptionState pointing to the current value.
func (c SubscriptionState) ToPtr() *SubscriptionState {
	return &c
}

// TierType - The commitment tier to use for next month.
type TierType string

const (
	TierTypeConsumption     TierType = "Consumption"
	TierTypeCommitment1TB   TierType = "Commitment_1TB"
	TierTypeCommitment10TB  TierType = "Commitment_10TB"
	TierTypeCommitment100TB TierType = "Commitment_100TB"
	TierTypeCommitment500TB TierType = "Commitment_500TB"
	TierTypeCommitment1PB   TierType = "Commitment_1PB"
	TierTypeCommitment5PB   TierType = "Commitment_5PB"
)

// PossibleTierTypeValues returns the possible values for the TierType const type.
func PossibleTierTypeValues() []TierType {
	return []TierType{
		TierTypeConsumption,
		TierTypeCommitment1TB,
		TierTypeCommitment10TB,
		TierTypeCommitment100TB,
		TierTypeCommitment500TB,
		TierTypeCommitment1PB,
		TierTypeCommitment5PB,
	}
}

// ToPtr returns a *TierType pointing to the current value.
func (c TierType) ToPtr() *TierType {
	return &c
}

// TrustedIDProviderState - The current state of the trusted identity provider feature for this Data Lake Store account.
type TrustedIDProviderState string

const (
	TrustedIDProviderStateEnabled  TrustedIDProviderState = "Enabled"
	TrustedIDProviderStateDisabled TrustedIDProviderState = "Disabled"
)

// PossibleTrustedIDProviderStateValues returns the possible values for the TrustedIDProviderState const type.
func PossibleTrustedIDProviderStateValues() []TrustedIDProviderState {
	return []TrustedIDProviderState{
		TrustedIDProviderStateEnabled,
		TrustedIDProviderStateDisabled,
	}
}

// ToPtr returns a *TrustedIDProviderState pointing to the current value.
func (c TrustedIDProviderState) ToPtr() *TrustedIDProviderState {
	return &c
}

// UsageUnit - Gets the unit of measurement.
type UsageUnit string

const (
	UsageUnitCount           UsageUnit = "Count"
	UsageUnitBytes           UsageUnit = "Bytes"
	UsageUnitSeconds         UsageUnit = "Seconds"
	UsageUnitPercent         UsageUnit = "Percent"
	UsageUnitCountsPerSecond UsageUnit = "CountsPerSecond"
	UsageUnitBytesPerSecond  UsageUnit = "BytesPerSecond"
)

// PossibleUsageUnitValues returns the possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{
		UsageUnitCount,
		UsageUnitBytes,
		UsageUnitSeconds,
		UsageUnitPercent,
		UsageUnitCountsPerSecond,
		UsageUnitBytesPerSecond,
	}
}

// ToPtr returns a *UsageUnit pointing to the current value.
func (c UsageUnit) ToPtr() *UsageUnit {
	return &c
}
