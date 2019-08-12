package harbor

import (
	"github.com/go-openapi/strfmt"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARBOR_HOST", nil),
				Description: "The hostname (in form of URI) of the Harbor server.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARBOR_USERNAME", nil),
				Description: "The username to use for HTTP basic authentication when accessing the Harbor server.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HARBOR_PASSWORD", nil),
				Description: "The password to use for HTTP basic authentication when accessing the Harbor server.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"harbor_project": dataSourceHarborProject(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	cfg := apiclient.DefaultTransportConfig().
		WithHost(d.Get("host").(string))
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	transport.DefaultAuthentication = httptransport.BasicAuth(
		d.Get("username").(string),
		d.Get("password").(string),
	)
	client := apiclient.New(transport, strfmt.Default)
	return client, nil
}
