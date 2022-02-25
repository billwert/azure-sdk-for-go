//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogz

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

// MonitorsClient contains the methods for the Monitors group.
// Don't use this type directly, use NewMonitorsClient() instead.
type MonitorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMonitorsClient creates a new instance of MonitorsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMonitorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *MonitorsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &MonitorsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreate - Create a monitor resource. This create operation can take upto 10 minutes to complete.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - MonitorsClientBeginCreateOptions contains the optional parameters for the MonitorsClient.BeginCreate method.
func (client *MonitorsClient) BeginCreate(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsClientBeginCreateOptions) (MonitorsClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return MonitorsClientCreatePollerResponse{}, err
	}
	result := MonitorsClientCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MonitorsClient.Create", "azure-async-operation", resp, client.pl)
	if err != nil {
		return MonitorsClientCreatePollerResponse{}, err
	}
	result.Poller = &MonitorsClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Create a monitor resource. This create operation can take upto 10 minutes to complete.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *MonitorsClient) create(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *MonitorsClient) createCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// BeginDelete - Delete a monitor resource. This delete operation can take upto 10 minutes to complete.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - MonitorsClientBeginDeleteOptions contains the optional parameters for the MonitorsClient.BeginDelete method.
func (client *MonitorsClient) BeginDelete(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsClientBeginDeleteOptions) (MonitorsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return MonitorsClientDeletePollerResponse{}, err
	}
	result := MonitorsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MonitorsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return MonitorsClientDeletePollerResponse{}, err
	}
	result.Poller = &MonitorsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a monitor resource. This delete operation can take upto 10 minutes to complete.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *MonitorsClient) deleteOperation(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MonitorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the properties of a specific monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - MonitorsClientGetOptions contains the optional parameters for the MonitorsClient.Get method.
func (client *MonitorsClient) Get(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsClientGetOptions) (MonitorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return MonitorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MonitorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MonitorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MonitorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MonitorsClient) getHandleResponse(resp *http.Response) (MonitorsClientGetResponse, error) {
	result := MonitorsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitorResource); err != nil {
		return MonitorsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - List all monitors under the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - MonitorsClientListByResourceGroupOptions contains the optional parameters for the MonitorsClient.ListByResourceGroup
// method.
func (client *MonitorsClient) ListByResourceGroup(resourceGroupName string, options *MonitorsClientListByResourceGroupOptions) *MonitorsClientListByResourceGroupPager {
	return &MonitorsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp MonitorsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MonitorResourceListResponse.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *MonitorsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *MonitorsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors"
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
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *MonitorsClient) listByResourceGroupHandleResponse(resp *http.Response) (MonitorsClientListByResourceGroupResponse, error) {
	result := MonitorsClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitorResourceListResponse); err != nil {
		return MonitorsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - List all monitors under the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - MonitorsClientListBySubscriptionOptions contains the optional parameters for the MonitorsClient.ListBySubscription
// method.
func (client *MonitorsClient) ListBySubscription(options *MonitorsClientListBySubscriptionOptions) *MonitorsClientListBySubscriptionPager {
	return &MonitorsClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp MonitorsClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MonitorResourceListResponse.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *MonitorsClient) listBySubscriptionCreateRequest(ctx context.Context, options *MonitorsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Logz/monitors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *MonitorsClient) listBySubscriptionHandleResponse(resp *http.Response) (MonitorsClientListBySubscriptionResponse, error) {
	result := MonitorsClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitorResourceListResponse); err != nil {
		return MonitorsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListMonitoredResources - List the resources currently being monitored by the Logz monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - MonitorsClientListMonitoredResourcesOptions contains the optional parameters for the MonitorsClient.ListMonitoredResources
// method.
func (client *MonitorsClient) ListMonitoredResources(resourceGroupName string, monitorName string, options *MonitorsClientListMonitoredResourcesOptions) *MonitorsClientListMonitoredResourcesPager {
	return &MonitorsClientListMonitoredResourcesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listMonitoredResourcesCreateRequest(ctx, resourceGroupName, monitorName, options)
		},
		advancer: func(ctx context.Context, resp MonitorsClientListMonitoredResourcesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MonitoredResourceListResponse.NextLink)
		},
	}
}

// listMonitoredResourcesCreateRequest creates the ListMonitoredResources request.
func (client *MonitorsClient) listMonitoredResourcesCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsClientListMonitoredResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/listMonitoredResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMonitoredResourcesHandleResponse handles the ListMonitoredResources response.
func (client *MonitorsClient) listMonitoredResourcesHandleResponse(resp *http.Response) (MonitorsClientListMonitoredResourcesResponse, error) {
	result := MonitorsClientListMonitoredResourcesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoredResourceListResponse); err != nil {
		return MonitorsClientListMonitoredResourcesResponse{}, err
	}
	return result, nil
}

// ListUserRoles - List the user's roles configured on Logz.io side for the account corresponding to the monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - MonitorsClientListUserRolesOptions contains the optional parameters for the MonitorsClient.ListUserRoles method.
func (client *MonitorsClient) ListUserRoles(resourceGroupName string, monitorName string, options *MonitorsClientListUserRolesOptions) *MonitorsClientListUserRolesPager {
	return &MonitorsClientListUserRolesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listUserRolesCreateRequest(ctx, resourceGroupName, monitorName, options)
		},
		advancer: func(ctx context.Context, resp MonitorsClientListUserRolesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.UserRoleListResponse.NextLink)
		},
	}
}

// listUserRolesCreateRequest creates the ListUserRoles request.
func (client *MonitorsClient) listUserRolesCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsClientListUserRolesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/listUserRoles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// listUserRolesHandleResponse handles the ListUserRoles response.
func (client *MonitorsClient) listUserRolesHandleResponse(resp *http.Response) (MonitorsClientListUserRolesResponse, error) {
	result := MonitorsClientListUserRolesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.UserRoleListResponse); err != nil {
		return MonitorsClientListUserRolesResponse{}, err
	}
	return result, nil
}

// Update - Update a monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// monitorName - Monitor resource name
// options - MonitorsClientUpdateOptions contains the optional parameters for the MonitorsClient.Update method.
func (client *MonitorsClient) Update(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsClientUpdateOptions) (MonitorsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return MonitorsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MonitorsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MonitorsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *MonitorsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *MonitorsClient) updateHandleResponse(resp *http.Response) (MonitorsClientUpdateResponse, error) {
	result := MonitorsClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitorResource); err != nil {
		return MonitorsClientUpdateResponse{}, err
	}
	return result, nil
}
