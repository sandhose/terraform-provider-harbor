---
layout: "harbor"
page_title: "Harbor: harbor_usergroup
sidebar_current: "docs-harbor-datasource-usergroup"
description: |-
  Looks up the information about a Harbor user group.
---

# harbor\_usergroup

Use this data source to get information about a specific Harbor
usergroup which already exists.

## Example Usage

```hcl
# Using the usergroup name
data "harbor_usergroup" "foo" {
  name = "foo"
}

# Using the usergroup ID
data "harbor_usergroup" "foo" {
  group_id = 1
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The name of the user group.
* `group_id` - (Optional) The ID of the user group.

~> **NOTE**: Both `name` and `group_id` cannot be specified.

## Attributes Reference

The following attributes are exported:

* `name` - The name of the user group.
* `type` - The type of the user group. `1` for `ldap`, `2` for `http`
* `ldap_group_dn` - The LDAP Group DN of the user group.
* `group_id` - The ID of the user group.
