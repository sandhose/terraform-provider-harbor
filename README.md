# terraform-provider-harbor

This is a [Terraform](https://terraform.io) provider for the [Harbor](https://goharbor.io) registry.

Right now it only implements the two things I needed: projects and robot accounts.

This is also one of my first Go project, feel free to tell me if I'm doing things wrong.

The API client is generated using [`go-swagger`](https://github.com/go-swagger/go-swagger).

# Installation

- Fetch the binary from the [releases](https://github.com/sandhose/terraform-provider-harbor/releases).
- Install it in your Terraform plugin directory (`~/.terraform.d/plugins/`), see <https://www.terraform.io/docs/configuration/providers.html#third-party-plugins>

# License

MIT
