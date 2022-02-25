//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservices

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// ClientCheckNameAvailabilityResponse contains the response from method Client.CheckNameAvailability.
type ClientCheckNameAvailabilityResponse struct {
	ClientCheckNameAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientCheckNameAvailabilityResult contains the result from method Client.CheckNameAvailability.
type ClientCheckNameAvailabilityResult struct {
	CheckNameAvailabilityResult
}

// OperationsClientGetOperationResultResponse contains the response from method OperationsClient.GetOperationResult.
type OperationsClientGetOperationResultResponse struct {
	OperationsClientGetOperationResultResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientGetOperationResultResult contains the result from method OperationsClient.GetOperationResult.
type OperationsClientGetOperationResultResult struct {
	Vault
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	ClientDiscoveryResponse
}

// OperationsClientOperationStatusGetResponse contains the response from method OperationsClient.OperationStatusGet.
type OperationsClientOperationStatusGetResponse struct {
	OperationsClientOperationStatusGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientOperationStatusGetResult contains the result from method OperationsClient.OperationStatusGet.
type OperationsClientOperationStatusGetResult struct {
	OperationResource
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResourcesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesClientGetResult contains the result from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResult struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	PrivateLinkResourcesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PrivateLinkResourcesClientListResult contains the result from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResult struct {
	PrivateLinkResources
}

// RegisteredIdentitiesClientDeleteResponse contains the response from method RegisteredIdentitiesClient.Delete.
type RegisteredIdentitiesClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReplicationUsagesClientListResponse contains the response from method ReplicationUsagesClient.List.
type ReplicationUsagesClientListResponse struct {
	ReplicationUsagesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ReplicationUsagesClientListResult contains the result from method ReplicationUsagesClient.List.
type ReplicationUsagesClientListResult struct {
	ReplicationUsageList
}

// UsagesClientListByVaultsResponse contains the response from method UsagesClient.ListByVaults.
type UsagesClientListByVaultsResponse struct {
	UsagesClientListByVaultsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// UsagesClientListByVaultsResult contains the result from method UsagesClient.ListByVaults.
type UsagesClientListByVaultsResult struct {
	VaultUsageList
}

// VaultCertificatesClientCreateResponse contains the response from method VaultCertificatesClient.Create.
type VaultCertificatesClientCreateResponse struct {
	VaultCertificatesClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VaultCertificatesClientCreateResult contains the result from method VaultCertificatesClient.Create.
type VaultCertificatesClientCreateResult struct {
	VaultCertificateResponse
}

// VaultExtendedInfoClientCreateOrUpdateResponse contains the response from method VaultExtendedInfoClient.CreateOrUpdate.
type VaultExtendedInfoClientCreateOrUpdateResponse struct {
	VaultExtendedInfoClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VaultExtendedInfoClientCreateOrUpdateResult contains the result from method VaultExtendedInfoClient.CreateOrUpdate.
type VaultExtendedInfoClientCreateOrUpdateResult struct {
	VaultExtendedInfoResource
}

// VaultExtendedInfoClientGetResponse contains the response from method VaultExtendedInfoClient.Get.
type VaultExtendedInfoClientGetResponse struct {
	VaultExtendedInfoClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VaultExtendedInfoClientGetResult contains the result from method VaultExtendedInfoClient.Get.
type VaultExtendedInfoClientGetResult struct {
	VaultExtendedInfoResource
}

// VaultExtendedInfoClientUpdateResponse contains the response from method VaultExtendedInfoClient.Update.
type VaultExtendedInfoClientUpdateResponse struct {
	VaultExtendedInfoClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VaultExtendedInfoClientUpdateResult contains the result from method VaultExtendedInfoClient.Update.
type VaultExtendedInfoClientUpdateResult struct {
	VaultExtendedInfoResource
}

// VaultsClientCreateOrUpdatePollerResponse contains the response from method VaultsClient.CreateOrUpdate.
type VaultsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *VaultsClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l VaultsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (VaultsClientCreateOrUpdateResponse, error) {
	respType := VaultsClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Vault)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a VaultsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *VaultsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *VaultsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("VaultsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &VaultsClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// VaultsClientCreateOrUpdateResponse contains the response from method VaultsClient.CreateOrUpdate.
type VaultsClientCreateOrUpdateResponse struct {
	VaultsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VaultsClientCreateOrUpdateResult contains the result from method VaultsClient.CreateOrUpdate.
type VaultsClientCreateOrUpdateResult struct {
	Vault
}

// VaultsClientDeleteResponse contains the response from method VaultsClient.Delete.
type VaultsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VaultsClientGetResponse contains the response from method VaultsClient.Get.
type VaultsClientGetResponse struct {
	VaultsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VaultsClientGetResult contains the result from method VaultsClient.Get.
type VaultsClientGetResult struct {
	Vault
}

// VaultsClientListByResourceGroupResponse contains the response from method VaultsClient.ListByResourceGroup.
type VaultsClientListByResourceGroupResponse struct {
	VaultsClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VaultsClientListByResourceGroupResult contains the result from method VaultsClient.ListByResourceGroup.
type VaultsClientListByResourceGroupResult struct {
	VaultList
}

// VaultsClientListBySubscriptionIDResponse contains the response from method VaultsClient.ListBySubscriptionID.
type VaultsClientListBySubscriptionIDResponse struct {
	VaultsClientListBySubscriptionIDResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VaultsClientListBySubscriptionIDResult contains the result from method VaultsClient.ListBySubscriptionID.
type VaultsClientListBySubscriptionIDResult struct {
	VaultList
}

// VaultsClientUpdatePollerResponse contains the response from method VaultsClient.Update.
type VaultsClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *VaultsClientUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l VaultsClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (VaultsClientUpdateResponse, error) {
	respType := VaultsClientUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Vault)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a VaultsClientUpdatePollerResponse from the provided client and resume token.
func (l *VaultsClientUpdatePollerResponse) Resume(ctx context.Context, client *VaultsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("VaultsClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &VaultsClientUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// VaultsClientUpdateResponse contains the response from method VaultsClient.Update.
type VaultsClientUpdateResponse struct {
	VaultsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// VaultsClientUpdateResult contains the result from method VaultsClient.Update.
type VaultsClientUpdateResult struct {
	Vault
}
