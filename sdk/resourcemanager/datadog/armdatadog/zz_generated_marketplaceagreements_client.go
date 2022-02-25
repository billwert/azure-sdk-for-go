//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatadog

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

// MarketplaceAgreementsClient contains the methods for the MarketplaceAgreements group.
// Don't use this type directly, use NewMarketplaceAgreementsClient() instead.
type MarketplaceAgreementsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMarketplaceAgreementsClient creates a new instance of MarketplaceAgreementsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMarketplaceAgreementsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *MarketplaceAgreementsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &MarketplaceAgreementsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Create Datadog marketplace agreement in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - MarketplaceAgreementsClientCreateOrUpdateOptions contains the optional parameters for the MarketplaceAgreementsClient.CreateOrUpdate
// method.
func (client *MarketplaceAgreementsClient) CreateOrUpdate(ctx context.Context, options *MarketplaceAgreementsClientCreateOrUpdateOptions) (MarketplaceAgreementsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, options)
	if err != nil {
		return MarketplaceAgreementsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MarketplaceAgreementsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MarketplaceAgreementsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MarketplaceAgreementsClient) createOrUpdateCreateRequest(ctx context.Context, options *MarketplaceAgreementsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Datadog/agreements/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *MarketplaceAgreementsClient) createOrUpdateHandleResponse(resp *http.Response) (MarketplaceAgreementsClientCreateOrUpdateResponse, error) {
	result := MarketplaceAgreementsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgreementResource); err != nil {
		return MarketplaceAgreementsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// List - List Datadog marketplace agreements in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - MarketplaceAgreementsClientListOptions contains the optional parameters for the MarketplaceAgreementsClient.List
// method.
func (client *MarketplaceAgreementsClient) List(options *MarketplaceAgreementsClientListOptions) *MarketplaceAgreementsClientListPager {
	return &MarketplaceAgreementsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp MarketplaceAgreementsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AgreementResourceListResponse.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *MarketplaceAgreementsClient) listCreateRequest(ctx context.Context, options *MarketplaceAgreementsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Datadog/agreements"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MarketplaceAgreementsClient) listHandleResponse(resp *http.Response) (MarketplaceAgreementsClientListResponse, error) {
	result := MarketplaceAgreementsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgreementResourceListResponse); err != nil {
		return MarketplaceAgreementsClientListResponse{}, err
	}
	return result, nil
}
