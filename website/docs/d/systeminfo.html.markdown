---
layout: "harbor"
page_title: "Harbor: harbor_systeminfo
sidebar_current: "docs-harbor-datasource-systeminfo"
description: |-
  Looks up the information about a Harbor instance.
---

# harbor\_project\_robot\_account

Use this data source to get information about a specific Harbor instance.

## Example Usage

```hcl
data "harbor_systeminfo" "foo" {
}
```

## Attributes Reference

The following attributes are exported:

* `admiral_endpoint` - The [Admiral](https://vmware.github.io/admiral/) endpoint.
* `auth_mode` - The [auth mode](https://github.com/goharbor/harbor/blob/master/docs/installation_guide.md#configuring-harbor) that is enabled.
* `external_url` - The [External URL](https://github.com/goharbor/harbor/blob/master/docs/installation_guide.md#configuring-harbor) of Harbor.
* `harbor_version` - The version of Harbor.
* `has_ca_root` - If Harbor has a CA root cert file ready for download.
* `project_creation_restriction` - If [project creation](https://github.com/goharbor/harbor/blob/master/docs/user_guide.md#managing-project-creation) is restricted.
* `registry_url` - The [Registry URL](https://docs.docker.com/registry/) of Harbor.
* `self_registration` - If [self registration](https://github.com/goharbor/harbor/blob/master/docs/user_guide.md#managing-self-registration) is enabled.
* `with_admiral` - If [Admiral](https://vmware.github.io/admiral/) is enabled.
* `with_clair` - If [Clair](https://github.com/coreos/clair) is enabled.
* `with_notary` - If [Notary](https://docs.docker.com/notary/getting_started/) is enabled.
