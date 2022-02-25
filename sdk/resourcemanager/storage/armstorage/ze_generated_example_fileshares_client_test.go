//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/DeletedFileSharesList.json
func ExampleFileSharesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewFileSharesClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<account-name>",
		&armstorage.FileSharesClientListOptions{Maxpagesize: nil,
			Filter: nil,
			Expand: to.StringPtr("<expand>"),
		})
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

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/FileSharesPut_NFS.json
func ExampleFileSharesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewFileSharesClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-name>",
		armstorage.FileShare{
			FileShareProperties: &armstorage.FileShareProperties{
				EnabledProtocols: armstorage.EnabledProtocols("NFS").ToPtr(),
			},
		},
		&armstorage.FileSharesClientCreateOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FileSharesClientCreateResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/FileShareAclsPatch.json
func ExampleFileSharesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewFileSharesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-name>",
		armstorage.FileShare{
			FileShareProperties: &armstorage.FileShareProperties{
				SignedIdentifiers: []*armstorage.SignedIdentifier{
					{
						AccessPolicy: &armstorage.AccessPolicy{
							ExpiryTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-05-01T08:49:37.0000000Z"); return t }()),
							Permission: to.StringPtr("<permission>"),
							StartTime:  to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T08:49:37.0000000Z"); return t }()),
						},
						ID: to.StringPtr("<id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FileSharesClientUpdateResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/FileSharesGet_Stats.json
func ExampleFileSharesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewFileSharesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-name>",
		&armstorage.FileSharesClientGetOptions{Expand: to.StringPtr("<expand>"),
			XMSSnapshot: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FileSharesClientGetResult)
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/FileSharesDelete.json
func ExampleFileSharesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewFileSharesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-name>",
		&armstorage.FileSharesClientDeleteOptions{XMSSnapshot: nil,
			Include: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/FileSharesRestore.json
func ExampleFileSharesClient_Restore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewFileSharesClient("<subscription-id>", cred, nil)
	_, err = client.Restore(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-name>",
		armstorage.DeletedShare{
			DeletedShareName:    to.StringPtr("<deleted-share-name>"),
			DeletedShareVersion: to.StringPtr("<deleted-share-version>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/FileSharesLease_Acquire.json
func ExampleFileSharesClient_Lease() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewFileSharesClient("<subscription-id>", cred, nil)
	res, err := client.Lease(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-name>",
		&armstorage.FileSharesClientLeaseOptions{XMSSnapshot: nil,
			Parameters: &armstorage.LeaseShareRequest{
				Action:        armstorage.LeaseShareAction("Acquire").ToPtr(),
				LeaseDuration: to.Int32Ptr(-1),
			},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FileSharesClientLeaseResult)
}
