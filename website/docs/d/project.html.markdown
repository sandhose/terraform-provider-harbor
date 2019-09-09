---
layout: "harbor"
page_title: "Harbor: harbor_project
sidebar_current: "docs-harbor-datasource-project"
description: |-
  Looks up the information about a Harbor project.
---

# harbor\_project

Use this data source to get information about a specific Harbor project which already exists.

## Example Usage

```hcl
# Using the project name
data "harbor_project" "foo" {
  name = "foo"
}

# Using the project ID
data "harbor_project" "foo" {
  project_id = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the project.

## Attributes Reference

The following attributes are exported:

* `project_id` - The ID of the project.
