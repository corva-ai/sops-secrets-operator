package mariadbapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2018-06-01/mariadb"
)

// ServersClientAPI contains the set of methods on the ServersClient type.
type ServersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, serverName string, parameters mariadb.ServerForCreate) (result mariadb.ServersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ServersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.Server, err error)
	List(ctx context.Context) (result mariadb.ServerListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result mariadb.ServerListResult, err error)
	Restart(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ServersRestartFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, parameters mariadb.ServerUpdateParameters) (result mariadb.ServersUpdateFuture, err error)
}

var _ ServersClientAPI = (*mariadb.ServersClient)(nil)

// ReplicasClientAPI contains the set of methods on the ReplicasClient type.
type ReplicasClientAPI interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ServerListResult, err error)
}

var _ ReplicasClientAPI = (*mariadb.ReplicasClient)(nil)

// FirewallRulesClientAPI contains the set of methods on the FirewallRulesClient type.
type FirewallRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters mariadb.FirewallRule) (result mariadb.FirewallRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result mariadb.FirewallRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result mariadb.FirewallRule, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.FirewallRuleListResult, err error)
}

var _ FirewallRulesClientAPI = (*mariadb.FirewallRulesClient)(nil)

// VirtualNetworkRulesClientAPI contains the set of methods on the VirtualNetworkRulesClient type.
type VirtualNetworkRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters mariadb.VirtualNetworkRule) (result mariadb.VirtualNetworkRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string) (result mariadb.VirtualNetworkRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string) (result mariadb.VirtualNetworkRule, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.VirtualNetworkRuleListResultPage, err error)
}

var _ VirtualNetworkRulesClientAPI = (*mariadb.VirtualNetworkRulesClient)(nil)

// DatabasesClientAPI contains the set of methods on the DatabasesClient type.
type DatabasesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters mariadb.Database) (result mariadb.DatabasesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result mariadb.DatabasesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result mariadb.Database, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.DatabaseListResult, err error)
}

var _ DatabasesClientAPI = (*mariadb.DatabasesClient)(nil)

// ConfigurationsClientAPI contains the set of methods on the ConfigurationsClient type.
type ConfigurationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, configurationName string, parameters mariadb.Configuration) (result mariadb.ConfigurationsCreateOrUpdateFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, configurationName string) (result mariadb.Configuration, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ConfigurationListResult, err error)
}

var _ ConfigurationsClientAPI = (*mariadb.ConfigurationsClient)(nil)

// LogFilesClientAPI contains the set of methods on the LogFilesClient type.
type LogFilesClientAPI interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.LogFileListResult, err error)
}

var _ LogFilesClientAPI = (*mariadb.LogFilesClient)(nil)

// LocationBasedPerformanceTierClientAPI contains the set of methods on the LocationBasedPerformanceTierClient type.
type LocationBasedPerformanceTierClientAPI interface {
	List(ctx context.Context, locationName string) (result mariadb.PerformanceTierListResult, err error)
}

var _ LocationBasedPerformanceTierClientAPI = (*mariadb.LocationBasedPerformanceTierClient)(nil)

// CheckNameAvailabilityClientAPI contains the set of methods on the CheckNameAvailabilityClient type.
type CheckNameAvailabilityClientAPI interface {
	Execute(ctx context.Context, nameAvailabilityRequest mariadb.NameAvailabilityRequest) (result mariadb.NameAvailability, err error)
}

var _ CheckNameAvailabilityClientAPI = (*mariadb.CheckNameAvailabilityClient)(nil)

// ServerSecurityAlertPoliciesClientAPI contains the set of methods on the ServerSecurityAlertPoliciesClient type.
type ServerSecurityAlertPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters mariadb.ServerSecurityAlertPolicy) (result mariadb.ServerSecurityAlertPoliciesCreateOrUpdateFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result mariadb.ServerSecurityAlertPolicy, err error)
}

var _ ServerSecurityAlertPoliciesClientAPI = (*mariadb.ServerSecurityAlertPoliciesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result mariadb.OperationListResult, err error)
}

var _ OperationsClientAPI = (*mariadb.OperationsClient)(nil)
