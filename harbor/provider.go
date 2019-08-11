package harbor

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("HARBOR_HOST", ""),
				Description: "The hostname (in form of URI) of the Harbor server.",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("HARBOR_USERNAME", ""),
				Description: "The username to use for HTTP basic authentication when accessing the Harbor server.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("HARBOR_PASSWORD", ""),
				Description: "The password to use for HTTP basic authentication when accessing the Harbor server.",
			},
		},
		ResourcesMap:   map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc:  providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return nil, nil
}
