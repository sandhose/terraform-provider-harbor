package harbor

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"
	"github.com/sandhose/terraform-provider-harbor/api/client/products"
	apimodels "github.com/sandhose/terraform-provider-harbor/api/models"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceHarborProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceHarborProjectCreate,
		Read:   resourceHarborProjectRead,
		// Update: resourceHarborProjectUpdate,
		Delete: resourceHarborProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// computed
			"project_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
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

func resourceHarborProjectCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)
	name := d.Get("name").(string)

	_, err := client.Products.PostProjects(
		products.NewPostProjectsParamsWithContext(context.TODO()).
			WithProject(&apimodels.ProjectReq{
				ProjectName: name,
			}),
		nil,
	)

	if err != nil {
		return err
	}

	resp, err := client.Products.GetProjects(
		products.NewGetProjectsParamsWithContext(context.TODO()).
			WithName(&name),
		nil,
	)

	if err != nil {
		return err
	}

	var project *apimodels.Project
	for _, p := range resp.Payload {
		if strings.ToLower(p.Name) == strings.ToLower(name) {
			project = p
			break
		}
	}

	if project == nil {
		return fmt.Errorf("Project not found")
	}

	harborProjectUpdate(d, project)

	return nil
}

func resourceHarborProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)

	if err != nil {
		return err
	}

	resp, err := client.Products.GetProjectsProjectID(
		products.NewGetProjectsProjectIDParamsWithContext(context.TODO()).
			WithProjectID(projectID),
		nil,
	)

	if err != nil {
		return err
	}

	harborProjectUpdate(d, resp.Payload)

	return nil
}

// func resourceHarborProjectUpdate(d *schema.ResourceData, meta interface{}) error {
// 	return nil
// }

func resourceHarborProjectDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)

	if err != nil {
		return err
	}

	_, err = client.Products.DeleteProjectsProjectID(
		products.NewDeleteProjectsProjectIDParamsWithContext(context.TODO()).
			WithProjectID(projectID),
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func harborProjectUpdate(d *schema.ResourceData, p *apimodels.Project) {
	d.SetId(fmt.Sprint(p.ProjectID))

	d.Set("project_id", p.ProjectID)
	d.Set("name", p.Name)
	d.Set("owner_id", p.OwnerID)
	d.Set("owner_name", p.OwnerName)
	d.Set("creation_time", p.CreationTime)
	d.Set("update_time", p.UpdateTime)
	d.Set("deleted", p.Deleted)
	d.Set("repo_count", p.RepoCount)
	d.Set("chart_count", p.ChartCount)
}
