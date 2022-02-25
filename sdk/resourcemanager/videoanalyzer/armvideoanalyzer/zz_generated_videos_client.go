//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvideoanalyzer

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
	"strconv"
	"strings"
)

// VideosClient contains the methods for the Videos group.
// Don't use this type directly, use NewVideosClient() instead.
type VideosClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVideosClient creates a new instance of VideosClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVideosClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VideosClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &VideosClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// CreateOrUpdate - Creates a new video resource or updates an existing video resource with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// videoName - The Video name.
// parameters - The request parameters
// options - VideosClientCreateOrUpdateOptions contains the optional parameters for the VideosClient.CreateOrUpdate method.
func (client *VideosClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, videoName string, parameters VideoEntity, options *VideosClientCreateOrUpdateOptions) (VideosClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, videoName, parameters, options)
	if err != nil {
		return VideosClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VideosClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return VideosClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VideosClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, videoName string, parameters VideoEntity, options *VideosClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/videos/{videoName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if videoName == "" {
		return nil, errors.New("parameter videoName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{videoName}", url.PathEscape(videoName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VideosClient) createOrUpdateHandleResponse(resp *http.Response) (VideosClientCreateOrUpdateResponse, error) {
	result := VideosClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VideoEntity); err != nil {
		return VideosClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing video resource and its underlying data. This operation is irreversible.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// videoName - The Video name.
// options - VideosClientDeleteOptions contains the optional parameters for the VideosClient.Delete method.
func (client *VideosClient) Delete(ctx context.Context, resourceGroupName string, accountName string, videoName string, options *VideosClientDeleteOptions) (VideosClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, videoName, options)
	if err != nil {
		return VideosClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VideosClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return VideosClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return VideosClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VideosClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, videoName string, options *VideosClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/videos/{videoName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if videoName == "" {
		return nil, errors.New("parameter videoName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{videoName}", url.PathEscape(videoName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves an existing video resource with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// videoName - The Video name.
// options - VideosClientGetOptions contains the optional parameters for the VideosClient.Get method.
func (client *VideosClient) Get(ctx context.Context, resourceGroupName string, accountName string, videoName string, options *VideosClientGetOptions) (VideosClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, videoName, options)
	if err != nil {
		return VideosClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VideosClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VideosClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VideosClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, videoName string, options *VideosClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/videos/{videoName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if videoName == "" {
		return nil, errors.New("parameter videoName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{videoName}", url.PathEscape(videoName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VideosClient) getHandleResponse(resp *http.Response) (VideosClientGetResponse, error) {
	result := VideosClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VideoEntity); err != nil {
		return VideosClientGetResponse{}, err
	}
	return result, nil
}

// List - Retrieves a list of video resources that have been created, along with their JSON representations.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// options - VideosClientListOptions contains the optional parameters for the VideosClient.List method.
func (client *VideosClient) List(resourceGroupName string, accountName string, options *VideosClientListOptions) *VideosClientListPager {
	return &VideosClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp VideosClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VideoEntityCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VideosClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *VideosClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/videos"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VideosClient) listHandleResponse(resp *http.Response) (VideosClientListResponse, error) {
	result := VideosClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VideoEntityCollection); err != nil {
		return VideosClientListResponse{}, err
	}
	return result, nil
}

// ListContentToken - Generates a streaming token which can be used for accessing content from video content URLs, for a video
// resource with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// videoName - The Video name.
// options - VideosClientListContentTokenOptions contains the optional parameters for the VideosClient.ListContentToken method.
func (client *VideosClient) ListContentToken(ctx context.Context, resourceGroupName string, accountName string, videoName string, options *VideosClientListContentTokenOptions) (VideosClientListContentTokenResponse, error) {
	req, err := client.listContentTokenCreateRequest(ctx, resourceGroupName, accountName, videoName, options)
	if err != nil {
		return VideosClientListContentTokenResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VideosClientListContentTokenResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VideosClientListContentTokenResponse{}, runtime.NewResponseError(resp)
	}
	return client.listContentTokenHandleResponse(resp)
}

// listContentTokenCreateRequest creates the ListContentToken request.
func (client *VideosClient) listContentTokenCreateRequest(ctx context.Context, resourceGroupName string, accountName string, videoName string, options *VideosClientListContentTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/videos/{videoName}/listContentToken"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if videoName == "" {
		return nil, errors.New("parameter videoName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{videoName}", url.PathEscape(videoName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listContentTokenHandleResponse handles the ListContentToken response.
func (client *VideosClient) listContentTokenHandleResponse(resp *http.Response) (VideosClientListContentTokenResponse, error) {
	result := VideosClientListContentTokenResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VideoContentToken); err != nil {
		return VideosClientListContentTokenResponse{}, err
	}
	return result, nil
}

// Update - Updates individual properties of an existing video resource with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The Azure Video Analyzer account name.
// videoName - The Video name.
// parameters - The request parameters
// options - VideosClientUpdateOptions contains the optional parameters for the VideosClient.Update method.
func (client *VideosClient) Update(ctx context.Context, resourceGroupName string, accountName string, videoName string, parameters VideoEntity, options *VideosClientUpdateOptions) (VideosClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, videoName, parameters, options)
	if err != nil {
		return VideosClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VideosClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VideosClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *VideosClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, videoName string, parameters VideoEntity, options *VideosClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/videos/{videoName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if videoName == "" {
		return nil, errors.New("parameter videoName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{videoName}", url.PathEscape(videoName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *VideosClient) updateHandleResponse(resp *http.Response) (VideosClientUpdateResponse, error) {
	result := VideosClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VideoEntity); err != nil {
		return VideosClientUpdateResponse{}, err
	}
	return result, nil
}
