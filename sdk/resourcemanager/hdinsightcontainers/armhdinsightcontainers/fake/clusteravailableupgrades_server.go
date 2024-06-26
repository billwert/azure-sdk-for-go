//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
	"net/http"
	"net/url"
	"regexp"
)

// ClusterAvailableUpgradesServer is a fake server for instances of the armhdinsightcontainers.ClusterAvailableUpgradesClient type.
type ClusterAvailableUpgradesServer struct {
	// NewListPager is the fake for method ClusterAvailableUpgradesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, clusterPoolName string, clusterName string, options *armhdinsightcontainers.ClusterAvailableUpgradesClientListOptions) (resp azfake.PagerResponder[armhdinsightcontainers.ClusterAvailableUpgradesClientListResponse])
}

// NewClusterAvailableUpgradesServerTransport creates a new instance of ClusterAvailableUpgradesServerTransport with the provided implementation.
// The returned ClusterAvailableUpgradesServerTransport instance is connected to an instance of armhdinsightcontainers.ClusterAvailableUpgradesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewClusterAvailableUpgradesServerTransport(srv *ClusterAvailableUpgradesServer) *ClusterAvailableUpgradesServerTransport {
	return &ClusterAvailableUpgradesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armhdinsightcontainers.ClusterAvailableUpgradesClientListResponse]](),
	}
}

// ClusterAvailableUpgradesServerTransport connects instances of armhdinsightcontainers.ClusterAvailableUpgradesClient to instances of ClusterAvailableUpgradesServer.
// Don't use this type directly, use NewClusterAvailableUpgradesServerTransport instead.
type ClusterAvailableUpgradesServerTransport struct {
	srv          *ClusterAvailableUpgradesServer
	newListPager *tracker[azfake.PagerResponder[armhdinsightcontainers.ClusterAvailableUpgradesClientListResponse]]
}

// Do implements the policy.Transporter interface for ClusterAvailableUpgradesServerTransport.
func (c *ClusterAvailableUpgradesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ClusterAvailableUpgradesClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ClusterAvailableUpgradesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HDInsight/clusterpools/(?P<clusterPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/availableUpgrades`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterPoolName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListPager(resourceGroupNameParam, clusterPoolNameParam, clusterNameParam, nil)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armhdinsightcontainers.ClusterAvailableUpgradesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}
