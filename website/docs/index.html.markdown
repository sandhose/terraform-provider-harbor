---
layout: "harbor"
page_title: "Provider: Harbor"
sidebar_current: "docs-harbor-index"
description: |-
  Harbor is an open source cloud native registry that stores, signs,
  and scans container images for vulnerabilities.
---

# Harbor Provider

[Harbor](https://goharbor.io/) is an open source cloud native registry that
stores, signs, and scans container images for vulnerabilities.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Harbor provider
provider "harbor" {
  host     = var.harbor_host
  username = var.harbor_username
  password = var.harbor_password
}

# Discover Harbor system information
data "harbor_systeminfo" "foo" {
}

# Create a project
resource "harbor_project" "foo" {
  name = "foo"
}

# Add a robot account
resource "harbor_project_robot_account" "foo" {
  project_id = harbor_project.foo.project_id

  name        = "foo"
  description = "Harbor website"

  access = {
    action  = "pull"
    resource = "repository"
  }

  disabled = false
}

# Add a usergroup
resource "harbor_usergroup" "foo" {
  name = "foo"
  type = 1
}
```

## Argument Reference

The following arguments are supported:

* `host` - (Required) Your Harbor endpoint. Can also use `HARBOR_HOST` environment variable.
* `username` - (Required) Your Harbor username. Can also use `HARBOR_USERNAME` environment variable.
* `password` - (Required) Your Harbor password. Can also use `HARBOR_PASSWORD` environment variable.
