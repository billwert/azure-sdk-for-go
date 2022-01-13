//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabric_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterGetOperation_example.json
func ExampleClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClustersClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClustersClientGetResult)
}

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterPutOperation_example_max.json
func ExampleClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armservicefabric.Cluster{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armservicefabric.ClusterProperties{
				AddOnFeatures: []*armservicefabric.AddOnFeatures{
					armservicefabric.AddOnFeatures("RepairManager").ToPtr(),
					armservicefabric.AddOnFeatures("DnsService").ToPtr(),
					armservicefabric.AddOnFeatures("BackupRestoreService").ToPtr(),
					armservicefabric.AddOnFeatures("ResourceMonitorService").ToPtr()},
				ApplicationTypeVersionsCleanupPolicy: &armservicefabric.ApplicationTypeVersionsCleanupPolicy{
					MaxUnusedVersionsToKeep: to.Int64Ptr(2),
				},
				AzureActiveDirectory: &armservicefabric.AzureActiveDirectory{
					ClientApplication:  to.StringPtr("<client-application>"),
					ClusterApplication: to.StringPtr("<cluster-application>"),
					TenantID:           to.StringPtr("<tenant-id>"),
				},
				CertificateCommonNames: &armservicefabric.ServerCertificateCommonNames{
					CommonNames: []*armservicefabric.ServerCertificateCommonName{
						{
							CertificateCommonName:       to.StringPtr("<certificate-common-name>"),
							CertificateIssuerThumbprint: to.StringPtr("<certificate-issuer-thumbprint>"),
						}},
					X509StoreName: armservicefabric.StoreName("My").ToPtr(),
				},
				ClientCertificateCommonNames: []*armservicefabric.ClientCertificateCommonName{
					{
						CertificateCommonName:       to.StringPtr("<certificate-common-name>"),
						CertificateIssuerThumbprint: to.StringPtr("<certificate-issuer-thumbprint>"),
						IsAdmin:                     to.BoolPtr(true),
					}},
				ClientCertificateThumbprints: []*armservicefabric.ClientCertificateThumbprint{
					{
						CertificateThumbprint: to.StringPtr("<certificate-thumbprint>"),
						IsAdmin:               to.BoolPtr(true),
					}},
				ClusterCodeVersion: to.StringPtr("<cluster-code-version>"),
				DiagnosticsStorageAccountConfig: &armservicefabric.DiagnosticsStorageAccountConfig{
					BlobEndpoint:            to.StringPtr("<blob-endpoint>"),
					ProtectedAccountKeyName: to.StringPtr("<protected-account-key-name>"),
					QueueEndpoint:           to.StringPtr("<queue-endpoint>"),
					StorageAccountName:      to.StringPtr("<storage-account-name>"),
					TableEndpoint:           to.StringPtr("<table-endpoint>"),
				},
				EventStoreServiceEnabled: to.BoolPtr(true),
				FabricSettings: []*armservicefabric.SettingsSectionDescription{
					{
						Name: to.StringPtr("<name>"),
						Parameters: []*armservicefabric.SettingsParameterDescription{
							{
								Name:  to.StringPtr("<name>"),
								Value: to.StringPtr("<value>"),
							}},
					}},
				InfrastructureServiceManager: to.BoolPtr(true),
				ManagementEndpoint:           to.StringPtr("<management-endpoint>"),
				NodeTypes: []*armservicefabric.NodeTypeDescription{
					{
						Name: to.StringPtr("<name>"),
						ApplicationPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Int32Ptr(30000),
							StartPort: to.Int32Ptr(20000),
						},
						ClientConnectionEndpointPort: to.Int32Ptr(19000),
						DurabilityLevel:              armservicefabric.DurabilityLevel("Silver").ToPtr(),
						EphemeralPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Int32Ptr(64000),
							StartPort: to.Int32Ptr(49000),
						},
						HTTPGatewayEndpointPort:   to.Int32Ptr(19007),
						IsPrimary:                 to.BoolPtr(true),
						IsStateless:               to.BoolPtr(false),
						MultipleAvailabilityZones: to.BoolPtr(true),
						VMInstanceCount:           to.Int32Ptr(5),
					}},
				Notifications: []*armservicefabric.Notification{
					{
						IsEnabled:            to.BoolPtr(true),
						NotificationCategory: armservicefabric.NotificationCategory("WaveProgress").ToPtr(),
						NotificationLevel:    armservicefabric.NotificationLevel("Critical").ToPtr(),
						NotificationTargets: []*armservicefabric.NotificationTarget{
							{
								NotificationChannel: armservicefabric.NotificationChannel("EmailUser").ToPtr(),
								Receivers: []*string{
									to.StringPtr("****@microsoft.com"),
									to.StringPtr("****@microsoft.com")},
							},
							{
								NotificationChannel: armservicefabric.NotificationChannel("EmailSubscription").ToPtr(),
								Receivers: []*string{
									to.StringPtr("Owner"),
									to.StringPtr("AccountAdmin")},
							}},
					},
					{
						IsEnabled:            to.BoolPtr(true),
						NotificationCategory: armservicefabric.NotificationCategory("WaveProgress").ToPtr(),
						NotificationLevel:    armservicefabric.NotificationLevel("All").ToPtr(),
						NotificationTargets: []*armservicefabric.NotificationTarget{
							{
								NotificationChannel: armservicefabric.NotificationChannel("EmailUser").ToPtr(),
								Receivers: []*string{
									to.StringPtr("****@microsoft.com"),
									to.StringPtr("****@microsoft.com")},
							},
							{
								NotificationChannel: armservicefabric.NotificationChannel("EmailSubscription").ToPtr(),
								Receivers: []*string{
									to.StringPtr("Owner"),
									to.StringPtr("AccountAdmin")},
							}},
					}},
				ReliabilityLevel: armservicefabric.ReliabilityLevel("Platinum").ToPtr(),
				ReverseProxyCertificateCommonNames: &armservicefabric.ServerCertificateCommonNames{
					CommonNames: []*armservicefabric.ServerCertificateCommonName{
						{
							CertificateCommonName:       to.StringPtr("<certificate-common-name>"),
							CertificateIssuerThumbprint: to.StringPtr("<certificate-issuer-thumbprint>"),
						}},
					X509StoreName: armservicefabric.StoreName("My").ToPtr(),
				},
				SfZonalUpgradeMode: armservicefabric.SfZonalUpgradeMode("Hierarchical").ToPtr(),
				UpgradeDescription: &armservicefabric.ClusterUpgradePolicy{
					DeltaHealthPolicy: &armservicefabric.ClusterUpgradeDeltaHealthPolicy{
						ApplicationDeltaHealthPolicies: map[string]*armservicefabric.ApplicationDeltaHealthPolicy{
							"fabric:/myApp1": {
								DefaultServiceTypeDeltaHealthPolicy: &armservicefabric.ServiceTypeDeltaHealthPolicy{
									MaxPercentDeltaUnhealthyServices: to.Int32Ptr(0),
								},
								ServiceTypeDeltaHealthPolicies: map[string]*armservicefabric.ServiceTypeDeltaHealthPolicy{
									"myServiceType1": {
										MaxPercentDeltaUnhealthyServices: to.Int32Ptr(0),
									},
								},
							},
						},
						MaxPercentDeltaUnhealthyApplications:       to.Int32Ptr(0),
						MaxPercentDeltaUnhealthyNodes:              to.Int32Ptr(0),
						MaxPercentUpgradeDomainDeltaUnhealthyNodes: to.Int32Ptr(0),
					},
					ForceRestart:              to.BoolPtr(false),
					HealthCheckRetryTimeout:   to.StringPtr("<health-check-retry-timeout>"),
					HealthCheckStableDuration: to.StringPtr("<health-check-stable-duration>"),
					HealthCheckWaitDuration:   to.StringPtr("<health-check-wait-duration>"),
					HealthPolicy: &armservicefabric.ClusterHealthPolicy{
						ApplicationHealthPolicies: map[string]*armservicefabric.ApplicationHealthPolicy{
							"fabric:/myApp1": {
								DefaultServiceTypeHealthPolicy: &armservicefabric.ServiceTypeHealthPolicy{
									MaxPercentUnhealthyServices: to.Int32Ptr(0),
								},
								ServiceTypeHealthPolicies: map[string]*armservicefabric.ServiceTypeHealthPolicy{
									"myServiceType1": {
										MaxPercentUnhealthyServices: to.Int32Ptr(100),
									},
								},
							},
						},
						MaxPercentUnhealthyApplications: to.Int32Ptr(0),
						MaxPercentUnhealthyNodes:        to.Int32Ptr(0),
					},
					UpgradeDomainTimeout:          to.StringPtr("<upgrade-domain-timeout>"),
					UpgradeReplicaSetCheckTimeout: to.StringPtr("<upgrade-replica-set-check-timeout>"),
					UpgradeTimeout:                to.StringPtr("<upgrade-timeout>"),
				},
				UpgradeMode:                   armservicefabric.UpgradeMode("Manual").ToPtr(),
				UpgradePauseEndTimestampUTC:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-25T22:00:00Z"); return t }()),
				UpgradePauseStartTimestampUTC: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-21T22:00:00Z"); return t }()),
				UpgradeWave:                   armservicefabric.ClusterUpgradeCadence("Wave1").ToPtr(),
				VMImage:                       to.StringPtr("<vmimage>"),
				VmssZonalUpgradeMode:          armservicefabric.VmssZonalUpgradeMode("Parallel").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.ClustersClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterPatchOperation_example.json
func ExampleClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armservicefabric.ClusterUpdateParameters{
			Properties: &armservicefabric.ClusterPropertiesUpdateParameters{
				EventStoreServiceEnabled: to.BoolPtr(true),
				NodeTypes: []*armservicefabric.NodeTypeDescription{
					{
						Name: to.StringPtr("<name>"),
						ApplicationPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Int32Ptr(30000),
							StartPort: to.Int32Ptr(20000),
						},
						ClientConnectionEndpointPort: to.Int32Ptr(19000),
						DurabilityLevel:              armservicefabric.DurabilityLevel("Bronze").ToPtr(),
						EphemeralPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Int32Ptr(64000),
							StartPort: to.Int32Ptr(49000),
						},
						HTTPGatewayEndpointPort: to.Int32Ptr(19007),
						IsPrimary:               to.BoolPtr(true),
						VMInstanceCount:         to.Int32Ptr(5),
					},
					{
						Name: to.StringPtr("<name>"),
						ApplicationPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Int32Ptr(2000),
							StartPort: to.Int32Ptr(1000),
						},
						ClientConnectionEndpointPort: to.Int32Ptr(0),
						DurabilityLevel:              armservicefabric.DurabilityLevel("Bronze").ToPtr(),
						EphemeralPorts: &armservicefabric.EndpointRangeDescription{
							EndPort:   to.Int32Ptr(4000),
							StartPort: to.Int32Ptr(3000),
						},
						HTTPGatewayEndpointPort: to.Int32Ptr(0),
						IsPrimary:               to.BoolPtr(false),
						VMInstanceCount:         to.Int32Ptr(3),
					}},
				ReliabilityLevel:              armservicefabric.ReliabilityLevel("Bronze").ToPtr(),
				UpgradeMode:                   armservicefabric.UpgradeMode("Automatic").ToPtr(),
				UpgradePauseEndTimestampUTC:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-25T22:00:00Z"); return t }()),
				UpgradePauseStartTimestampUTC: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-21T22:00:00Z"); return t }()),
				UpgradeWave:                   armservicefabric.ClusterUpgradeCadence("Wave").ToPtr(),
			},
			Tags: map[string]*string{
				"a": to.StringPtr("b"),
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
	log.Printf("Response result: %#v\n", res.ClustersClientUpdateResult)
}

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterDeleteOperation_example.json
func ExampleClustersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClustersClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterListByResourceGroupOperation_example.json
func ExampleClustersClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClustersClient("<subscription-id>", cred, nil)
	res, err := client.ListByResourceGroup(ctx,
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClustersClientListByResourceGroupResult)
}

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ClusterListOperation_example.json
func ExampleClustersClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClustersClient("<subscription-id>", cred, nil)
	res, err := client.List(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClustersClientListResult)
}

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ListUpgradableVersionsMinMax_example.json
func ExampleClustersClient_ListUpgradableVersions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClustersClient("<subscription-id>", cred, nil)
	res, err := client.ListUpgradableVersions(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		&armservicefabric.ClustersClientListUpgradableVersionsOptions{VersionsDescription: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClustersClientListUpgradableVersionsResult)
}