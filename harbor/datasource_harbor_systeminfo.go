package harbor

import (
	"context"

	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"
	"github.com/sandhose/terraform-provider-harbor/api/client/products"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceHarborSystemInfo() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHarborSystemInfoRead,
		Schema: map[string]*schema.Schema{
			"with_notary": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"with_clair": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"with_admiral": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"admiral_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"registry_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_creation_restriction": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"self_registration": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"has_ca_root": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"harbor_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceHarborSystemInfoRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)
	resp, err := client.Products.GetSysteminfo(products.NewGetSysteminfoParamsWithContext(context.TODO()), nil)

	if err != nil {
		return err
	}

	i := resp.Payload
	// TODO: unique ID?
	d.SetId("systeminfo")
	d.Set("admiral_endpoint", i.AdmiralEndpoint)
	d.Set("auth_mode", i.AuthMode)
	d.Set("external_url", i.ExternalURL)
	d.Set("harbor_version", i.HarborVersion)
	d.Set("has_ca_root", i.HasCaRoot)
	d.Set("project_creation_restriction", i.ProjectCreationRestriction)
	d.Set("registry_url", i.RegistryURL)
	d.Set("self_registration", i.SelfRegistration)
	d.Set("with_admiral", i.WithAdmiral)
	d.Set("with_clair", i.WithClair)
	d.Set("with_notary", i.WithNotary)
	return nil
}
