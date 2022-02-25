//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

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

// ResourceQuotaLimitsClient contains the methods for the NetAppResourceQuotaLimits group.
// Don't use this type directly, use NewResourceQuotaLimitsClient() instead.
type ResourceQuotaLimitsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewResourceQuotaLimitsClient creates a new instance of ResourceQuotaLimitsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourceQuotaLimitsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ResourceQuotaLimitsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ResourceQuotaLimitsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Get the default and current subscription quota limit
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location
// quotaLimitName - The name of the Quota Limit
// options - ResourceQuotaLimitsClientGetOptions contains the optional parameters for the ResourceQuotaLimitsClient.Get method.
func (client *ResourceQuotaLimitsClient) Get(ctx context.Context, location string, quotaLimitName string, options *ResourceQuotaLimitsClientGetOptions) (ResourceQuotaLimitsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, quotaLimitName, options)
	if err != nil {
		return ResourceQuotaLimitsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceQuotaLimitsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceQuotaLimitsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ResourceQuotaLimitsClient) getCreateRequest(ctx context.Context, location string, quotaLimitName string, options *ResourceQuotaLimitsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetApp/locations/{location}/quotaLimits/{quotaLimitName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if quotaLimitName == "" {
		return nil, errors.New("parameter quotaLimitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{quotaLimitName}", url.PathEscape(quotaLimitName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ResourceQuotaLimitsClient) getHandleResponse(resp *http.Response) (ResourceQuotaLimitsClientGetResponse, error) {
	result := ResourceQuotaLimitsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionQuotaItem); err != nil {
		return ResourceQuotaLimitsClientGetResponse{}, err
	}
	return result, nil
}

// List - Get the default and current limits for quotas
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location
// options - ResourceQuotaLimitsClientListOptions contains the optional parameters for the ResourceQuotaLimitsClient.List
// method.
func (client *ResourceQuotaLimitsClient) List(ctx context.Context, location string, options *ResourceQuotaLimitsClientListOptions) (ResourceQuotaLimitsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, options)
	if err != nil {
		return ResourceQuotaLimitsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceQuotaLimitsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceQuotaLimitsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ResourceQuotaLimitsClient) listCreateRequest(ctx context.Context, location string, options *ResourceQuotaLimitsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetApp/locations/{location}/quotaLimits"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourceQuotaLimitsClient) listHandleResponse(resp *http.Response) (ResourceQuotaLimitsClientListResponse, error) {
	result := ResourceQuotaLimitsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionQuotaItemList); err != nil {
		return ResourceQuotaLimitsClientListResponse{}, err
	}
	return result, nil
}
