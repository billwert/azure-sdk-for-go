//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/liveevent-list-all.json
func ExampleLiveEventsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewLiveEventsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<account-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/liveevent-list-by-name.json
func ExampleLiveEventsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewLiveEventsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<live-event-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LiveEventsClientGetResult)
}

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/liveevent-create.json
func ExampleLiveEventsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewLiveEventsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<live-event-name>",
		armmediaservices.LiveEvent{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
			Properties: &armmediaservices.LiveEventProperties{
				Description: to.StringPtr("<description>"),
				Input: &armmediaservices.LiveEventInput{
					AccessControl: &armmediaservices.LiveEventInputAccessControl{
						IP: &armmediaservices.IPAccessControl{
							Allow: []*armmediaservices.IPRange{
								{
									Name:               to.StringPtr("<name>"),
									Address:            to.StringPtr("<address>"),
									SubnetPrefixLength: to.Int32Ptr(0),
								}},
						},
					},
					KeyFrameIntervalDuration: to.StringPtr("<key-frame-interval-duration>"),
					StreamingProtocol:        armmediaservices.LiveEventInputProtocol("RTMP").ToPtr(),
				},
				Preview: &armmediaservices.LiveEventPreview{
					AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
						IP: &armmediaservices.IPAccessControl{
							Allow: []*armmediaservices.IPRange{
								{
									Name:               to.StringPtr("<name>"),
									Address:            to.StringPtr("<address>"),
									SubnetPrefixLength: to.Int32Ptr(0),
								}},
						},
					},
				},
			},
		},
		&armmediaservices.LiveEventsClientBeginCreateOptions{AutoStart: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LiveEventsClientCreateResult)
}

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/liveevent-update.json
func ExampleLiveEventsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewLiveEventsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<live-event-name>",
		armmediaservices.LiveEvent{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
				"tag3": to.StringPtr("value3"),
			},
			Properties: &armmediaservices.LiveEventProperties{
				Description: to.StringPtr("<description>"),
				Input: &armmediaservices.LiveEventInput{
					AccessControl: &armmediaservices.LiveEventInputAccessControl{
						IP: &armmediaservices.IPAccessControl{
							Allow: []*armmediaservices.IPRange{
								{
									Name:    to.StringPtr("<name>"),
									Address: to.StringPtr("<address>"),
								}},
						},
					},
					KeyFrameIntervalDuration: to.StringPtr("<key-frame-interval-duration>"),
					StreamingProtocol:        armmediaservices.LiveEventInputProtocol("FragmentedMP4").ToPtr(),
				},
				Preview: &armmediaservices.LiveEventPreview{
					AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
						IP: &armmediaservices.IPAccessControl{
							Allow: []*armmediaservices.IPRange{
								{
									Name:    to.StringPtr("<name>"),
									Address: to.StringPtr("<address>"),
								}},
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LiveEventsClientUpdateResult)
}

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/liveevent-delete.json
func ExampleLiveEventsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewLiveEventsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<live-event-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/liveevent-allocate.json
func ExampleLiveEventsClient_BeginAllocate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewLiveEventsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginAllocate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<live-event-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/liveevent-start.json
func ExampleLiveEventsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewLiveEventsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStart(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<live-event-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/liveevent-stop.json
func ExampleLiveEventsClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewLiveEventsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginStop(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<live-event-name>",
		armmediaservices.LiveEventActionInput{
			RemoveOutputsOnStop: to.BoolPtr(false),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/liveevent-reset.json
func ExampleLiveEventsClient_BeginReset() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewLiveEventsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginReset(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<live-event-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
