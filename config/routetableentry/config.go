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

package routetableentry

import (
	tjconfig "github.com/crossplane/terrajet/pkg/config"
)

// Configure configures the null group
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("tencentcloud_route_table_entry", func(r *tjconfig.Resource) {
		r.ExternalName = tjconfig.IdentifierFromProvider
		r.ShortGroup = "routetableentry"
		r.References["route_table_id"] = tjconfig.Reference{
			Type:              "github.com/crossplane-contrib/provider-jet-tencentcloud/apis/routetable/v1alpha1.RouteTable",
			RefFieldName:      "RouteTableIdRefs",
			SelectorFieldName: "RouteTableIdSelector",
		}
	})
}
