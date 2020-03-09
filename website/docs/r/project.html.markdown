---
layout: "harbor"
page_title: "Harbor: harbor_project
sidebar_current: "docs-harbor-resource-project"
description: |-
  Create and manage projects in Harbor.
---

# harbor\_project

Use this resource to create and manage Harbor projects.

## Example Usage

```hcl
resource "harbor_project" "foo" {
  name = "foo"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the project.

## Attributes Reference

The following attributes are exported:

* `project_id` - The ID of the project.

## Import

Projects can be imported using the `project_id`, e.g.

```
$ terraform import harbor_project.main 12345
```
