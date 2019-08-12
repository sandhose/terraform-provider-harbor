package harbor

import (
	"context"
	"fmt"
	"strings"

	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"
	"github.com/sandhose/terraform-provider-harbor/api/client/products"
	apimodels "github.com/sandhose/terraform-provider-harbor/api/models"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceHarborProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHarborProjectRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"name"},
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"project_id"},
			},

			// computed
			"owner_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"owner_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deleted": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"repo_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"chart_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceHarborProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	set := 0
	var projectID int64
	if id, ok := d.GetOk("project_id"); ok {
		projectID = int64(id.(int))
		set++
	}

	var projectName string
	if name, ok := d.GetOk("name"); ok {
		projectName = name.(string)
		set++
	}

	if set != 1 {
		return fmt.Errorf("One of %q or %q has to be provided", "project_id", "name")
	}

	var project *apimodels.Project
	if projectName != "" {
		resp, err := client.Products.GetProjects(
			products.NewGetProjectsParamsWithContext(context.TODO()).
				WithName(&projectName),
			nil,
		)

		if err != nil {
			return err
		}

		for _, p := range resp.Payload {
			if strings.ToLower(p.Name) == strings.ToLower(projectName) {
				project = p
				break
			}
		}
	} else {
		resp, err := client.Products.GetProjectsProjectID(
			products.NewGetProjectsProjectIDParamsWithContext(context.TODO()).
				WithProjectID(projectID),
			nil,
		)

		if err != nil {
			return err
		}

		project = resp.Payload
	}

	if project == nil {
		return fmt.Errorf("Project not found")
	}

	harborProjectUpdate(d, project)

	return nil
}
