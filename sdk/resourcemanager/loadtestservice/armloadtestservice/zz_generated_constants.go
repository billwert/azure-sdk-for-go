//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armloadtestservice

const (
	moduleName    = "armloadtestservice"
	moduleVersion = "v0.2.1"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// ToPtr returns a *ActionType pointing to the current value.
func (c ActionType) ToPtr() *ActionType {
	return &c
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ToPtr returns a *Origin pointing to the current value.
func (c Origin) ToPtr() *Origin {
	return &c
}

// ResourceState - Load Test resources provisioning states.
type ResourceState string

const (
	ResourceStateCanceled  ResourceState = "Canceled"
	ResourceStateDeleted   ResourceState = "Deleted"
	ResourceStateFailed    ResourceState = "Failed"
	ResourceStateSucceeded ResourceState = "Succeeded"
)

// PossibleResourceStateValues returns the possible values for the ResourceState const type.
func PossibleResourceStateValues() []ResourceState {
	return []ResourceState{
		ResourceStateCanceled,
		ResourceStateDeleted,
		ResourceStateFailed,
		ResourceStateSucceeded,
	}
}

// ToPtr returns a *ResourceState pointing to the current value.
func (c ResourceState) ToPtr() *ResourceState {
	return &c
}

// SystemAssignedServiceIdentityType - Type of managed service identity (either system assigned, or none).
type SystemAssignedServiceIdentityType string

const (
	SystemAssignedServiceIdentityTypeNone           SystemAssignedServiceIdentityType = "None"
	SystemAssignedServiceIdentityTypeSystemAssigned SystemAssignedServiceIdentityType = "SystemAssigned"
)

// PossibleSystemAssignedServiceIdentityTypeValues returns the possible values for the SystemAssignedServiceIdentityType const type.
func PossibleSystemAssignedServiceIdentityTypeValues() []SystemAssignedServiceIdentityType {
	return []SystemAssignedServiceIdentityType{
		SystemAssignedServiceIdentityTypeNone,
		SystemAssignedServiceIdentityTypeSystemAssigned,
	}
}

// ToPtr returns a *SystemAssignedServiceIdentityType pointing to the current value.
func (c SystemAssignedServiceIdentityType) ToPtr() *SystemAssignedServiceIdentityType {
	return &c
}
