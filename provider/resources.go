// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grafana

import (
	"fmt"
	"path/filepath"

	"github.com/grafana/terraform-provider-grafana/grafana"
	"github.com/nij4t/pulumi-grafana/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "grafana"
	// modules:
	mainMod = "" // the grafana module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(grafana.Provider("v1.17.0")())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "grafana",
		Description: "A Pulumi package for creating and managing grafana cloud resources.",
		Keywords:    []string{"pulumi", "grafana"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/nij4t/pulumi-grafana",
		Config: map[string]*tfbridge.SchemaInfo{
			"auth": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_AUTH"},
				},
			},
			"url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GRAFANA_URL"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type.
			"grafana_api_key": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ApiKey")},
			"grafana_builtin_role_assignment": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "BuiltinRoleAssignment"),
			},
			"grafana_dashboard":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Dashboard")},
			"grafana_dashboard_permission": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DashboardPermission")},
			"grafana_data_source":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DataSource")},
			"grafana_data_source_permission": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DataSourcePermission"),
			},
			"grafana_folder":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Folder")},
			"grafana_folder_permission": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FolderPermission")},
			"grafana_organization":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Organization")},
			"grafana_playlist":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Playlist")},
			"grafana_role":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Role")},
			"grafana_synthetic_monitoring_check": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "SynteticMonitoringCheck"),
			},
			"grafana_synthetic_monitoring_probe": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "SynteticMonitoringProbe"),
			},
			"grafana_team":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Team")},
			"grafana_team_external_group": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TeamExternalGroup")},
			"grafana_team_preferences":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TemPreferences")},
			"grafana_user":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "User")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainMod, "getAmi")},
			"grafana_folder": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFolder")},
			"grafana_synthetic_monitoring_probe": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSynteticMonitoringProbe"),
			},
			"grafana_synthetic_monitoring_probes": {
				Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSynteticMonitoringProbes"),
			},
			"grafana_user": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUser")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			// Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
