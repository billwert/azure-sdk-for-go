//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiothub

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

// Client contains the methods for the IotHub group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// subscriptionID - The subscription identifier.
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

// BeginManualFailover - Manually initiate a failover for the IoT Hub to its secondary region. To learn more, see https://aka.ms/manualfailover
// If the operation fails it returns an *azcore.ResponseError type.
// iotHubName - Name of the IoT hub to failover
// resourceGroupName - Name of the resource group containing the IoT hub resource
// failoverInput - Region to failover to. Must be the Azure paired region. Get the value from the secondary location in the
// locations property. To learn more, see https://aka.ms/manualfailover/region
// options - ClientBeginManualFailoverOptions contains the optional parameters for the Client.BeginManualFailover method.
func (client *Client) BeginManualFailover(ctx context.Context, iotHubName string, resourceGroupName string, failoverInput FailoverInput, options *ClientBeginManualFailoverOptions) (ClientManualFailoverPollerResponse, error) {
	resp, err := client.manualFailover(ctx, iotHubName, resourceGroupName, failoverInput, options)
	if err != nil {
		return ClientManualFailoverPollerResponse{}, err
	}
	result := ClientManualFailoverPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("Client.ManualFailover", "", resp, client.pl)
	if err != nil {
		return ClientManualFailoverPollerResponse{}, err
	}
	result.Poller = &ClientManualFailoverPoller{
		pt: pt,
	}
	return result, nil
}

// ManualFailover - Manually initiate a failover for the IoT Hub to its secondary region. To learn more, see https://aka.ms/manualfailover
// If the operation fails it returns an *azcore.ResponseError type.
func (client *Client) manualFailover(ctx context.Context, iotHubName string, resourceGroupName string, failoverInput FailoverInput, options *ClientBeginManualFailoverOptions) (*http.Response, error) {
	req, err := client.manualFailoverCreateRequest(ctx, iotHubName, resourceGroupName, failoverInput, options)
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

// manualFailoverCreateRequest creates the ManualFailover request.
func (client *Client) manualFailoverCreateRequest(ctx context.Context, iotHubName string, resourceGroupName string, failoverInput FailoverInput, options *ClientBeginManualFailoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{iotHubName}/failover"
	if iotHubName == "" {
		return nil, errors.New("parameter iotHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotHubName}", url.PathEscape(iotHubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, failoverInput)
}
