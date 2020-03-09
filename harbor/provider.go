package harbor

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
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
		ResourcesMap: map[string]*schema.Resource{
			"harbor_project":               resourceHarborProject(),
			"harbor_project_robot_account": resourceHarborProjectRobotAccount(),
			"harbor_usergroup":             resourceHarborUserGroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"harbor_system_info": dataSourceHarborSystemInfo(),
			"harbor_project":     dataSourceHarborProject(),
			"harbor_usergroup":   dataSourceHarborUserGroup(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	cfg := &Config{
		Host:     d.Get("host").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
	}
	return cfg.Config()
}
