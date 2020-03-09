---
layout: "harbor"
page_title: "Harbor: harbor_usergroup
sidebar_current: "docs-harbor-resource-usergroup"
description: |-
  Create and manage user groups in Harbor.
---

# harbor\_usergroup

Use this resource to create and manage Harbor usergroups.

## Example Usage

### Basic Auth

```hcl
resource "harbor_usergroup" "foo" {
  name = "foo"
  type = 2
}
```

### LDAP

```hcl
resource "harbor_usergroup" "bar" {
  name = "bar"
  type = 1
  ldap_group_dn = "cn=example,cn=groups,cn=accounts,dc=domain,dc=com"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the user group.
* `type` - (Required) The type of the user group. `1` for `ldap`, `2` for `http`.
* `ldap_group_dn` - (Required if `type = 1`) The LDAP Group DN of the user group.

## Attributes Reference

The following attributes are exported:

* `group_id` - The ID of the user group.

## Import

User groups can be imported using the `group_id`, e.g.

```shell
$ terraform import harbor_usergroup.main 12345
```
