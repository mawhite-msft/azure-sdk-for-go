//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwebpubsub

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// HubsClient contains the methods for the WebPubSubHubs group.
// Don't use this type directly, use NewHubsClient() instead.
type HubsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHubsClient creates a new instance of HubsClient with the specified values.
// subscriptionID - Gets subscription Id which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHubsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HubsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &HubsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a hub setting.
// If the operation fails it returns an *azcore.ResponseError type.
// hubName - The hub name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// resourceName - The name of the resource.
// parameters - The resource of WebPubSubHub and its properties
// options - HubsClientBeginCreateOrUpdateOptions contains the optional parameters for the HubsClient.BeginCreateOrUpdate
// method.
func (client *HubsClient) BeginCreateOrUpdate(ctx context.Context, hubName string, resourceGroupName string, resourceName string, parameters Hub, options *HubsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[HubsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, hubName, resourceGroupName, resourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[HubsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[HubsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a hub setting.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *HubsClient) createOrUpdate(ctx context.Context, hubName string, resourceGroupName string, resourceName string, parameters Hub, options *HubsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, hubName, resourceGroupName, resourceName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HubsClient) createOrUpdateCreateRequest(ctx context.Context, hubName string, resourceGroupName string, resourceName string, parameters Hub, options *HubsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/hubs/{hubName}"
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete a hub setting.
// If the operation fails it returns an *azcore.ResponseError type.
// hubName - The hub name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// resourceName - The name of the resource.
// options - HubsClientBeginDeleteOptions contains the optional parameters for the HubsClient.BeginDelete method.
func (client *HubsClient) BeginDelete(ctx context.Context, hubName string, resourceGroupName string, resourceName string, options *HubsClientBeginDeleteOptions) (*armruntime.Poller[HubsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, hubName, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[HubsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[HubsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a hub setting.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *HubsClient) deleteOperation(ctx context.Context, hubName string, resourceGroupName string, resourceName string, options *HubsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, hubName, resourceGroupName, resourceName, options)
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
func (client *HubsClient) deleteCreateRequest(ctx context.Context, hubName string, resourceGroupName string, resourceName string, options *HubsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/hubs/{hubName}"
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get a hub setting.
// If the operation fails it returns an *azcore.ResponseError type.
// hubName - The hub name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// resourceName - The name of the resource.
// options - HubsClientGetOptions contains the optional parameters for the HubsClient.Get method.
func (client *HubsClient) Get(ctx context.Context, hubName string, resourceGroupName string, resourceName string, options *HubsClientGetOptions) (HubsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, hubName, resourceGroupName, resourceName, options)
	if err != nil {
		return HubsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HubsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HubsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HubsClient) getCreateRequest(ctx context.Context, hubName string, resourceGroupName string, resourceName string, options *HubsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/hubs/{hubName}"
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HubsClient) getHandleResponse(resp *http.Response) (HubsClientGetResponse, error) {
	result := HubsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Hub); err != nil {
		return HubsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List hub settings.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// resourceName - The name of the resource.
// options - HubsClientListOptions contains the optional parameters for the HubsClient.List method.
func (client *HubsClient) NewListPager(resourceGroupName string, resourceName string, options *HubsClientListOptions) *runtime.Pager[HubsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[HubsClientListResponse]{
		More: func(page HubsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HubsClientListResponse) (HubsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HubsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HubsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HubsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *HubsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *HubsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/hubs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HubsClient) listHandleResponse(resp *http.Response) (HubsClientListResponse, error) {
	result := HubsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HubList); err != nil {
		return HubsClientListResponse{}, err
	}
	return result, nil
}
