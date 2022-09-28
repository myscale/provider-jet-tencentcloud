# Terrajet TencentCloud Provider

`provider-jet-tencentcloud` is a [Crossplane](https://crossplane.io/) provider that
is built using [Terrajet](https://github.com/crossplane/terrajet) code
generation tools and exposes XRM-conformant managed resources for the
TencentCloud API. And currently, tencentcloud_eks_cluster, tencentcloud_subnet and tencentcloud_vpc are supported.

## Deploying
Add a submodule to your project first:
```console
make submodules
```

Generate CRDs:
```console
make generate
```

Run provider locally:
```console
make run
```

Build images:
```console
make build PLATFORM=linux_amd64
```

Build, push, and install:
```console
make all
```

## Governance and Owners

provider-jet-tencentcloud is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-jet-tencentcloud adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-jet-tencentcloud is under the Apache 2.0 license.
