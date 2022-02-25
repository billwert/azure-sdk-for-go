//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabric

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

// ClusterVersionsClient contains the methods for the ClusterVersions group.
// Don't use this type directly, use NewClusterVersionsClient() instead.
type ClusterVersionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewClusterVersionsClient creates a new instance of ClusterVersionsClient with the specified values.
// subscriptionID - The customer subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClusterVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ClusterVersionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ClusterVersionsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Gets information about an available Service Fabric cluster code version.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location for the cluster code versions. This is different from cluster location.
// clusterVersion - The cluster code version.
// options - ClusterVersionsClientGetOptions contains the optional parameters for the ClusterVersionsClient.Get method.
func (client *ClusterVersionsClient) Get(ctx context.Context, location string, clusterVersion string, options *ClusterVersionsClientGetOptions) (ClusterVersionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, clusterVersion, options)
	if err != nil {
		return ClusterVersionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClusterVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClusterVersionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ClusterVersionsClient) getCreateRequest(ctx context.Context, location string, clusterVersion string, options *ClusterVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions/{clusterVersion}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if clusterVersion == "" {
		return nil, errors.New("parameter clusterVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterVersion}", url.PathEscape(clusterVersion))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ClusterVersionsClient) getHandleResponse(resp *http.Response) (ClusterVersionsClientGetResponse, error) {
	result := ClusterVersionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterCodeVersionsListResult); err != nil {
		return ClusterVersionsClientGetResponse{}, err
	}
	return result, nil
}

// GetByEnvironment - Gets information about an available Service Fabric cluster code version by environment.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location for the cluster code versions. This is different from cluster location.
// environment - The operating system of the cluster. The default means all.
// clusterVersion - The cluster code version.
// options - ClusterVersionsClientGetByEnvironmentOptions contains the optional parameters for the ClusterVersionsClient.GetByEnvironment
// method.
func (client *ClusterVersionsClient) GetByEnvironment(ctx context.Context, location string, environment ClusterVersionsEnvironment, clusterVersion string, options *ClusterVersionsClientGetByEnvironmentOptions) (ClusterVersionsClientGetByEnvironmentResponse, error) {
	req, err := client.getByEnvironmentCreateRequest(ctx, location, environment, clusterVersion, options)
	if err != nil {
		return ClusterVersionsClientGetByEnvironmentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClusterVersionsClientGetByEnvironmentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClusterVersionsClientGetByEnvironmentResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByEnvironmentHandleResponse(resp)
}

// getByEnvironmentCreateRequest creates the GetByEnvironment request.
func (client *ClusterVersionsClient) getByEnvironmentCreateRequest(ctx context.Context, location string, environment ClusterVersionsEnvironment, clusterVersion string, options *ClusterVersionsClientGetByEnvironmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions/{clusterVersion}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if environment == "" {
		return nil, errors.New("parameter environment cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environment}", url.PathEscape(string(environment)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if clusterVersion == "" {
		return nil, errors.New("parameter clusterVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterVersion}", url.PathEscape(clusterVersion))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByEnvironmentHandleResponse handles the GetByEnvironment response.
func (client *ClusterVersionsClient) getByEnvironmentHandleResponse(resp *http.Response) (ClusterVersionsClientGetByEnvironmentResponse, error) {
	result := ClusterVersionsClientGetByEnvironmentResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterCodeVersionsListResult); err != nil {
		return ClusterVersionsClientGetByEnvironmentResponse{}, err
	}
	return result, nil
}

// List - Gets all available code versions for Service Fabric cluster resources by location.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location for the cluster code versions. This is different from cluster location.
// options - ClusterVersionsClientListOptions contains the optional parameters for the ClusterVersionsClient.List method.
func (client *ClusterVersionsClient) List(ctx context.Context, location string, options *ClusterVersionsClientListOptions) (ClusterVersionsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, options)
	if err != nil {
		return ClusterVersionsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClusterVersionsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClusterVersionsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ClusterVersionsClient) listCreateRequest(ctx context.Context, location string, options *ClusterVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ClusterVersionsClient) listHandleResponse(resp *http.Response) (ClusterVersionsClientListResponse, error) {
	result := ClusterVersionsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterCodeVersionsListResult); err != nil {
		return ClusterVersionsClientListResponse{}, err
	}
	return result, nil
}

// ListByEnvironment - Gets all available code versions for Service Fabric cluster resources by environment.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location for the cluster code versions. This is different from cluster location.
// environment - The operating system of the cluster. The default means all.
// options - ClusterVersionsClientListByEnvironmentOptions contains the optional parameters for the ClusterVersionsClient.ListByEnvironment
// method.
func (client *ClusterVersionsClient) ListByEnvironment(ctx context.Context, location string, environment ClusterVersionsEnvironment, options *ClusterVersionsClientListByEnvironmentOptions) (ClusterVersionsClientListByEnvironmentResponse, error) {
	req, err := client.listByEnvironmentCreateRequest(ctx, location, environment, options)
	if err != nil {
		return ClusterVersionsClientListByEnvironmentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClusterVersionsClientListByEnvironmentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClusterVersionsClientListByEnvironmentResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByEnvironmentHandleResponse(resp)
}

// listByEnvironmentCreateRequest creates the ListByEnvironment request.
func (client *ClusterVersionsClient) listByEnvironmentCreateRequest(ctx context.Context, location string, environment ClusterVersionsEnvironment, options *ClusterVersionsClientListByEnvironmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if environment == "" {
		return nil, errors.New("parameter environment cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environment}", url.PathEscape(string(environment)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByEnvironmentHandleResponse handles the ListByEnvironment response.
func (client *ClusterVersionsClient) listByEnvironmentHandleResponse(resp *http.Response) (ClusterVersionsClientListByEnvironmentResponse, error) {
	result := ClusterVersionsClientListByEnvironmentResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterCodeVersionsListResult); err != nil {
		return ClusterVersionsClientListByEnvironmentResponse{}, err
	}
	return result, nil
}
