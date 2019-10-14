# terraform-provider-harbor

This is a [Terraform](https://terraform.io) provider for the [Harbor](https://goharbor.io) registry.

Right now it only implements the two things I needed: projects and robot accounts.

This is also one of my first Go project, feel free to tell me if I'm doing things wrong.

The API client is generated using [`go-swagger`](https://github.com/go-swagger/go-swagger).

# Installation

1. `make build`
2. provider should now be available as `$GOPATH/bin/terraform-provider-harbor`

# License

MIT
