//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// Client contains the methods for the DomainServices group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *Client {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &Client{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - The Create Domain Service operation creates a new domain service with the specified parameters. If
// the specific service already exists, then any patchable properties will be updated and any immutable
// properties will remain unchanged.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// domainServiceName - The name of the domain service.
// domainService - Properties supplied to the Create or Update a Domain Service operation.
// options - ClientBeginCreateOrUpdateOptions contains the optional parameters for the Client.BeginCreateOrUpdate method.
func (client *Client) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService, options *ClientBeginCreateOrUpdateOptions) (ClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, domainServiceName, domainService, options)
	if err != nil {
		return ClientCreateOrUpdatePollerResponse{}, err
	}
	result := ClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("Client.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return ClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - The Create Domain Service operation creates a new domain service with the specified parameters. If the
// specific service already exists, then any patchable properties will be updated and any immutable
// properties will remain unchanged.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *Client) createOrUpdate(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService, options *ClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, domainServiceName, domainService, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *Client) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService, options *ClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, domainService)
}

// BeginDelete - The Delete Domain Service operation deletes an existing Domain Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// domainServiceName - The name of the domain service.
// options - ClientBeginDeleteOptions contains the optional parameters for the Client.BeginDelete method.
func (client *Client) BeginDelete(ctx context.Context, resourceGroupName string, domainServiceName string, options *ClientBeginDeleteOptions) (ClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, domainServiceName, options)
	if err != nil {
		return ClientDeletePollerResponse{}, err
	}
	result := ClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("Client.Delete", "", resp, client.pl)
	if err != nil {
		return ClientDeletePollerResponse{}, err
	}
	result.Poller = &ClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - The Delete Domain Service operation deletes an existing Domain Service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *Client) deleteOperation(ctx context.Context, resourceGroupName string, domainServiceName string, options *ClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, domainServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *Client) deleteCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, options *ClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - The Get Domain Service operation retrieves a json representation of the Domain Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// domainServiceName - The name of the domain service.
// options - ClientGetOptions contains the optional parameters for the Client.Get method.
func (client *Client) Get(ctx context.Context, resourceGroupName string, domainServiceName string, options *ClientGetOptions) (ClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, domainServiceName, options)
	if err != nil {
		return ClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *Client) getCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, options *ClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *Client) getHandleResponse(resp *http.Response) (ClientGetResponse, error) {
	result := ClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainService); err != nil {
		return ClientGetResponse{}, err
	}
	return result, nil
}

// List - The List Domain Services in Subscription operation lists all the domain services available under the given subscription
// (and across all resource groups within that subscription).
// If the operation fails it returns an *azcore.ResponseError type.
// options - ClientListOptions contains the optional parameters for the Client.List method.
func (client *Client) List(options *ClientListOptions) *ClientListPager {
	return &ClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DomainServiceListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *Client) listCreateRequest(ctx context.Context, options *ClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AAD/domainServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *Client) listHandleResponse(resp *http.Response) (ClientListResponse, error) {
	result := ClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainServiceListResult); err != nil {
		return ClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - The List Domain Services in Resource Group operation lists all the domain services available under
// the given resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// options - ClientListByResourceGroupOptions contains the optional parameters for the Client.ListByResourceGroup method.
func (client *Client) ListByResourceGroup(resourceGroupName string, options *ClientListByResourceGroupOptions) *ClientListByResourceGroupPager {
	return &ClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DomainServiceListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *Client) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *Client) listByResourceGroupHandleResponse(resp *http.Response) (ClientListByResourceGroupResponse, error) {
	result := ClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainServiceListResult); err != nil {
		return ClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The Update Domain Service operation can be used to update the existing deployment. The update call only supports
// the properties listed in the PATCH body.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// domainServiceName - The name of the domain service.
// domainService - Properties supplied to the Update a Domain Service operation.
// options - ClientBeginUpdateOptions contains the optional parameters for the Client.BeginUpdate method.
func (client *Client) BeginUpdate(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService, options *ClientBeginUpdateOptions) (ClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, domainServiceName, domainService, options)
	if err != nil {
		return ClientUpdatePollerResponse{}, err
	}
	result := ClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("Client.Update", "", resp, client.pl)
	if err != nil {
		return ClientUpdatePollerResponse{}, err
	}
	result.Poller = &ClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - The Update Domain Service operation can be used to update the existing deployment. The update call only supports
// the properties listed in the PATCH body.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *Client) update(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService, options *ClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, domainServiceName, domainService, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *Client) updateCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, domainService DomainService, options *ClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, domainService)
}