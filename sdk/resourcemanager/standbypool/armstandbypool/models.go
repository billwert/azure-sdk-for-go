//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstandbypool

import "time"

// ContainerGroupProfile - Details of the ContainerGroupProfile.
type ContainerGroupProfile struct {
	// REQUIRED; Specifies container group profile id of standby container groups.
	ID *string

	// Specifies revision of container group profile.
	Revision *int64
}

// ContainerGroupProfileUpdate - Details of the ContainerGroupProfile.
type ContainerGroupProfileUpdate struct {
	// Specifies container group profile id of standby container groups.
	ID *string

	// Specifies revision of container group profile.
	Revision *int64
}

// ContainerGroupProperties - Details of the ContainerGroupProperties.
type ContainerGroupProperties struct {
	// REQUIRED; Specifies container group profile of standby container groups.
	ContainerGroupProfile *ContainerGroupProfile

	// Specifies subnet Ids for container group.
	SubnetIDs []*Subnet
}

// ContainerGroupPropertiesUpdate - Details of the ContainerGroupProperties.
type ContainerGroupPropertiesUpdate struct {
	// Specifies container group profile of standby container groups.
	ContainerGroupProfile *ContainerGroupProfileUpdate

	// Specifies subnet Ids for container group.
	SubnetIDs []*Subnet
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// StandbyContainerGroupPoolElasticityProfile - Specifies the elasticity profile of the standby container group pools.
type StandbyContainerGroupPoolElasticityProfile struct {
	// REQUIRED; Specifies maximum number of standby container groups in the standby pool.
	MaxReadyCapacity *int64

	// Specifies refill policy of the pool.
	RefillPolicy *RefillPolicy
}

// StandbyContainerGroupPoolElasticityProfileUpdate - Specifies the elasticity profile of the standby container group pools.
type StandbyContainerGroupPoolElasticityProfileUpdate struct {
	// Specifies maximum number of standby container groups in the standby pool.
	MaxReadyCapacity *int64

	// Specifies refill policy of the pool.
	RefillPolicy *RefillPolicy
}

// StandbyContainerGroupPoolResource - A StandbyContainerGroupPoolResource.
type StandbyContainerGroupPoolResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *StandbyContainerGroupPoolResourceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// StandbyContainerGroupPoolResourceListResult - The response of a StandbyContainerGroupPoolResource list operation.
type StandbyContainerGroupPoolResourceListResult struct {
	// REQUIRED; The StandbyContainerGroupPoolResource items on this page
	Value []*StandbyContainerGroupPoolResource

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// StandbyContainerGroupPoolResourceProperties - Details of the StandbyContainerGroupPool.
type StandbyContainerGroupPoolResourceProperties struct {
	// REQUIRED; Specifies container group properties of standby container group pools.
	ContainerGroupProperties *ContainerGroupProperties

	// REQUIRED; Specifies elasticity profile of standby container group pools.
	ElasticityProfile *StandbyContainerGroupPoolElasticityProfile

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState
}

// StandbyContainerGroupPoolResourceUpdate - The type used for update operations of the StandbyContainerGroupPoolResource.
type StandbyContainerGroupPoolResourceUpdate struct {
	// The updatable properties of the StandbyContainerGroupPoolResource.
	Properties *StandbyContainerGroupPoolResourceUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// StandbyContainerGroupPoolResourceUpdateProperties - The updatable properties of the StandbyContainerGroupPoolResource.
type StandbyContainerGroupPoolResourceUpdateProperties struct {
	// Specifies container group properties of standby container group pools.
	ContainerGroupProperties *ContainerGroupPropertiesUpdate

	// Specifies elasticity profile of standby container group pools.
	ElasticityProfile *StandbyContainerGroupPoolElasticityProfileUpdate
}

// StandbyVirtualMachinePoolElasticityProfile - Details of the elasticity profile.
type StandbyVirtualMachinePoolElasticityProfile struct {
	// REQUIRED; Specifies the maximum number of virtual machines in the standby virtual machine pool.
	MaxReadyCapacity *int64
}

// StandbyVirtualMachinePoolElasticityProfileUpdate - Details of the elasticity profile.
type StandbyVirtualMachinePoolElasticityProfileUpdate struct {
	// Specifies the maximum number of virtual machines in the standby virtual machine pool.
	MaxReadyCapacity *int64
}

// StandbyVirtualMachinePoolResource - A StandbyVirtualMachinePoolResource.
type StandbyVirtualMachinePoolResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *StandbyVirtualMachinePoolResourceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// StandbyVirtualMachinePoolResourceListResult - The response of a StandbyVirtualMachinePoolResource list operation.
type StandbyVirtualMachinePoolResourceListResult struct {
	// REQUIRED; The StandbyVirtualMachinePoolResource items on this page
	Value []*StandbyVirtualMachinePoolResource

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// StandbyVirtualMachinePoolResourceProperties - Details of the StandbyVirtualMachinePool.
type StandbyVirtualMachinePoolResourceProperties struct {
	// REQUIRED; Specifies the desired state of virtual machines in the pool.
	VirtualMachineState *VirtualMachineState

	// Specifies the fully qualified resource ID of a virtual machine scale set the pool is attached to.
	AttachedVirtualMachineScaleSetID *string

	// Specifies the elasticity profile of the standby virtual machine pools.
	ElasticityProfile *StandbyVirtualMachinePoolElasticityProfile

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState
}

// StandbyVirtualMachinePoolResourceUpdate - The type used for update operations of the StandbyVirtualMachinePoolResource.
type StandbyVirtualMachinePoolResourceUpdate struct {
	// The updatable properties of the StandbyVirtualMachinePoolResource.
	Properties *StandbyVirtualMachinePoolResourceUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// StandbyVirtualMachinePoolResourceUpdateProperties - The updatable properties of the StandbyVirtualMachinePoolResource.
type StandbyVirtualMachinePoolResourceUpdateProperties struct {
	// Specifies the fully qualified resource ID of a virtual machine scale set the pool is attached to.
	AttachedVirtualMachineScaleSetID *string

	// Specifies the elasticity profile of the standby virtual machine pools.
	ElasticityProfile *StandbyVirtualMachinePoolElasticityProfileUpdate

	// Specifies the desired state of virtual machines in the pool.
	VirtualMachineState *VirtualMachineState
}

// StandbyVirtualMachineResource - Concrete proxy resource types can be created by aliasing this type using a specific property
// type.
type StandbyVirtualMachineResource struct {
	// The resource-specific properties for this resource.
	Properties *StandbyVirtualMachineResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// StandbyVirtualMachineResourceListResult - The response of a StandbyVirtualMachineResource list operation.
type StandbyVirtualMachineResourceListResult struct {
	// REQUIRED; The StandbyVirtualMachineResource items on this page
	Value []*StandbyVirtualMachineResource

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// StandbyVirtualMachineResourceProperties - Details of the StandbyVirtualMachine.
type StandbyVirtualMachineResourceProperties struct {
	// REQUIRED; Resource id of the virtual machine.
	VirtualMachineResourceID *string

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState
}

// Subnet of container group
type Subnet struct {
	// REQUIRED; Specifies ARM resource id of the subnet.
	ID *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}
