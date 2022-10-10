/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-jet-tencentcloud/config/eks"
	"github.com/crossplane-contrib/provider-jet-tencentcloud/config/kubernetes"
	"github.com/crossplane-contrib/provider-jet-tencentcloud/config/nodepool"
	"github.com/crossplane-contrib/provider-jet-tencentcloud/config/role"
	"github.com/crossplane-contrib/provider-jet-tencentcloud/config/rolepolicyattachment"
	"github.com/crossplane-contrib/provider-jet-tencentcloud/config/routetable"
	"github.com/crossplane-contrib/provider-jet-tencentcloud/config/routetableentry"
	"github.com/crossplane-contrib/provider-jet-tencentcloud/config/securitygroup"
	"github.com/crossplane-contrib/provider-jet-tencentcloud/config/securitygrouprule"
	"github.com/crossplane-contrib/provider-jet-tencentcloud/config/subnet"
	"github.com/crossplane-contrib/provider-jet-tencentcloud/config/vpc"
)

const (
	resourcePrefix = "tencentcloud"
	modulePath     = "github.com/crossplane-contrib/provider-jet-tencentcloud"
)

//go:embed schema.json
var providerSchema string

// IncludedResources include resource
var IncludedResources = []string{
	"tencentcloud_vpc$",
	"tencentcloud_subnet$",
	"tencentcloud_eks_cluster$",
	"tencentcloud_kubernetes_cluster$",
	"tencentcloud_kubernetes_node_pool$",
	"tencentcloud_cam_role$",
	"tencentcloud_cam_role_policy_attachment$",
	"tencentcloud_security_group$",
	"tencentcloud_security_group_rule$",
	"tencentcloud_route_table$",
	"tencentcloud_route_table_entry$",
}

// skipList
var skipList []string

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList(IncludedResources),
		tjconfig.WithSkipList(skipList),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		vpc.Configure,
		subnet.Configure,
		eks.Configure,
		kubernetes.Configure,
		nodepool.Configure,
		role.Configure,
		rolepolicyattachment.Configure,
		securitygroup.Configure,
		securitygrouprule.Configure,
		routetable.Configure,
		routetableentry.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
