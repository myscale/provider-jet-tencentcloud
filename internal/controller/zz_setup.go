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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	clusterendpoint "github.com/crossplane-contrib/provider-jet-tencentcloud/internal/controller/clusterendpoint/clusterendpoint"
	cluster "github.com/crossplane-contrib/provider-jet-tencentcloud/internal/controller/eks/cluster"
	clusterkubernetes "github.com/crossplane-contrib/provider-jet-tencentcloud/internal/controller/kubernetes/cluster"
	nodepool "github.com/crossplane-contrib/provider-jet-tencentcloud/internal/controller/nodepool/nodepool"
	providerconfig "github.com/crossplane-contrib/provider-jet-tencentcloud/internal/controller/providerconfig"
	role "github.com/crossplane-contrib/provider-jet-tencentcloud/internal/controller/role/role"
	rolepolicyattachment "github.com/crossplane-contrib/provider-jet-tencentcloud/internal/controller/rolepolicyattachment/rolepolicyattachment"
	table "github.com/crossplane-contrib/provider-jet-tencentcloud/internal/controller/routetable/table"
	tableentry "github.com/crossplane-contrib/provider-jet-tencentcloud/internal/controller/routetableentry/tableentry"
	group "github.com/crossplane-contrib/provider-jet-tencentcloud/internal/controller/securitygroup/group"
	grouprule "github.com/crossplane-contrib/provider-jet-tencentcloud/internal/controller/securitygrouprule/grouprule"
	subnet "github.com/crossplane-contrib/provider-jet-tencentcloud/internal/controller/subnet/subnet"
	vpc "github.com/crossplane-contrib/provider-jet-tencentcloud/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		clusterendpoint.Setup,
		cluster.Setup,
		clusterkubernetes.Setup,
		nodepool.Setup,
		providerconfig.Setup,
		role.Setup,
		rolepolicyattachment.Setup,
		table.Setup,
		tableentry.Setup,
		group.Setup,
		grouprule.Setup,
		subnet.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
