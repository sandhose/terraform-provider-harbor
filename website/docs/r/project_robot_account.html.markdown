---
layout: "harbor"
page_title: "Harbor: harbor_project_robot_account
sidebar_current: "docs-harbor-resource-project-robot-account"
description: |-
  Create and manage robot accounts in Harbor.
---

# harbor\_project\_robot\_account

Use this resource to create and manage Harbor robot accounts.

## Example Usage

```hcl
resource "harbor_project" "foo" {
  name = "foo"
}

resource "harbor_project_robot_account" "foo" {
  project_id = harbor_project.foo.project_id
  name = "foo"
  description = "foo"

  access {
    action = "pull"
    resource = "repository"
  }
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The ID of the project.
* `name` - (Required) The name of the robot account.
* `description` - (Optional) Description of the robot account.
* `access` - (Optional) The access granted to the robot account. See [Access](#access)
* `disabled` - (Optional) Set to whether to enable the robot account. Defaults to `false`

## Access

* `action` - (Required) The action granted to the robot accounts.
* `resource` - (Required) The resource the action is granted on.

## Attributes Reference

The following attributes are exported:

* `username` - The username of the robot account.
* `token` - The token of the robot account.

## Import

Robot accounts can be imported using the `project_id/robot_id`, e.g.

```
$ terraform import harbor_project_robot_account.main 12345/12345
```
