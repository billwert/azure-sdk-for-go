//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualHubBgpConnectionServer is a fake server for instances of the armnetwork.VirtualHubBgpConnectionClient type.
type VirtualHubBgpConnectionServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualHubBgpConnectionClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, parameters armnetwork.BgpConnection, options *armnetwork.VirtualHubBgpConnectionClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualHubBgpConnectionClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *armnetwork.VirtualHubBgpConnectionClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualHubBgpConnectionClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *armnetwork.VirtualHubBgpConnectionClientGetOptions) (resp azfake.Responder[armnetwork.VirtualHubBgpConnectionClientGetResponse], errResp azfake.ErrorResponder)
}

// NewVirtualHubBgpConnectionServerTransport creates a new instance of VirtualHubBgpConnectionServerTransport with the provided implementation.
// The returned VirtualHubBgpConnectionServerTransport instance is connected to an instance of armnetwork.VirtualHubBgpConnectionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualHubBgpConnectionServerTransport(srv *VirtualHubBgpConnectionServer) *VirtualHubBgpConnectionServerTransport {
	return &VirtualHubBgpConnectionServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionClientDeleteResponse]](),
	}
}

// VirtualHubBgpConnectionServerTransport connects instances of armnetwork.VirtualHubBgpConnectionClient to instances of VirtualHubBgpConnectionServer.
// Don't use this type directly, use NewVirtualHubBgpConnectionServerTransport instead.
type VirtualHubBgpConnectionServerTransport struct {
	srv                 *VirtualHubBgpConnectionServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.VirtualHubBgpConnectionClientDeleteResponse]]
}

// Do implements the policy.Transporter interface for VirtualHubBgpConnectionServerTransport.
func (v *VirtualHubBgpConnectionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualHubBgpConnectionClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualHubBgpConnectionClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualHubBgpConnectionClient.Get":
		resp, err = v.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualHubBgpConnectionServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/bgpConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.BgpConnection](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualHubNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualHubName")])
		if err != nil {
			return nil, err
		}
		connectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, virtualHubNameUnescaped, connectionNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualHubBgpConnectionServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/bgpConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualHubNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualHubName")])
		if err != nil {
			return nil, err
		}
		connectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, virtualHubNameUnescaped, connectionNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualHubBgpConnectionServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/virtualHubs/(?P<virtualHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/bgpConnections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualHubNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("virtualHubName")])
	if err != nil {
		return nil, err
	}
	connectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameUnescaped, virtualHubNameUnescaped, connectionNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BgpConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
