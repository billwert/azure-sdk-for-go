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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ManagementGroupNetworkManagerConnectionsServer is a fake server for instances of the armnetwork.ManagementGroupNetworkManagerConnectionsClient type.
type ManagementGroupNetworkManagerConnectionsServer struct {
	// CreateOrUpdate is the fake for method ManagementGroupNetworkManagerConnectionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, managementGroupID string, networkManagerConnectionName string, parameters armnetwork.ManagerConnection, options *armnetwork.ManagementGroupNetworkManagerConnectionsClientCreateOrUpdateOptions) (resp azfake.Responder[armnetwork.ManagementGroupNetworkManagerConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ManagementGroupNetworkManagerConnectionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, managementGroupID string, networkManagerConnectionName string, options *armnetwork.ManagementGroupNetworkManagerConnectionsClientDeleteOptions) (resp azfake.Responder[armnetwork.ManagementGroupNetworkManagerConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagementGroupNetworkManagerConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, managementGroupID string, networkManagerConnectionName string, options *armnetwork.ManagementGroupNetworkManagerConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.ManagementGroupNetworkManagerConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ManagementGroupNetworkManagerConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(managementGroupID string, options *armnetwork.ManagementGroupNetworkManagerConnectionsClientListOptions) (resp azfake.PagerResponder[armnetwork.ManagementGroupNetworkManagerConnectionsClientListResponse])
}

// NewManagementGroupNetworkManagerConnectionsServerTransport creates a new instance of ManagementGroupNetworkManagerConnectionsServerTransport with the provided implementation.
// The returned ManagementGroupNetworkManagerConnectionsServerTransport instance is connected to an instance of armnetwork.ManagementGroupNetworkManagerConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagementGroupNetworkManagerConnectionsServerTransport(srv *ManagementGroupNetworkManagerConnectionsServer) *ManagementGroupNetworkManagerConnectionsServerTransport {
	return &ManagementGroupNetworkManagerConnectionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.ManagementGroupNetworkManagerConnectionsClientListResponse]](),
	}
}

// ManagementGroupNetworkManagerConnectionsServerTransport connects instances of armnetwork.ManagementGroupNetworkManagerConnectionsClient to instances of ManagementGroupNetworkManagerConnectionsServer.
// Don't use this type directly, use NewManagementGroupNetworkManagerConnectionsServerTransport instead.
type ManagementGroupNetworkManagerConnectionsServerTransport struct {
	srv          *ManagementGroupNetworkManagerConnectionsServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.ManagementGroupNetworkManagerConnectionsClientListResponse]]
}

// Do implements the policy.Transporter interface for ManagementGroupNetworkManagerConnectionsServerTransport.
func (m *ManagementGroupNetworkManagerConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagementGroupNetworkManagerConnectionsClient.CreateOrUpdate":
		resp, err = m.dispatchCreateOrUpdate(req)
	case "ManagementGroupNetworkManagerConnectionsClient.Delete":
		resp, err = m.dispatchDelete(req)
	case "ManagementGroupNetworkManagerConnectionsClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagementGroupNetworkManagerConnectionsClient.NewListPager":
		resp, err = m.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagementGroupNetworkManagerConnectionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/providers/Microsoft.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkManagerConnections/(?P<networkManagerConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.ManagerConnection](req)
	if err != nil {
		return nil, err
	}
	managementGroupIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	networkManagerConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.CreateOrUpdate(req.Context(), managementGroupIDUnescaped, networkManagerConnectionNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagerConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagementGroupNetworkManagerConnectionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if m.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/providers/Microsoft.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkManagerConnections/(?P<networkManagerConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	networkManagerConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Delete(req.Context(), managementGroupIDUnescaped, networkManagerConnectionNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagementGroupNetworkManagerConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkManagerConnections/(?P<networkManagerConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	networkManagerConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), managementGroupIDUnescaped, networkManagerConnectionNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagerConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagementGroupNetworkManagerConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/providers/Microsoft.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkManagerConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		managementGroupIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armnetwork.ManagementGroupNetworkManagerConnectionsClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.ManagementGroupNetworkManagerConnectionsClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := m.srv.NewListPager(managementGroupIDUnescaped, options)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.ManagementGroupNetworkManagerConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}
