package harbor

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"
	"github.com/sandhose/terraform-provider-harbor/api/client/products"
	apimodels "github.com/sandhose/terraform-provider-harbor/api/models"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceHarborProjectRobotAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceHarborProjectRobotAccountCreate,
		Read:   resourceHarborProjectRobotAccountRead,
		Update: resourceHarborProjectRobotAccountUpdate,
		Delete: resourceHarborProjectRobotAccountDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
				ForceNew: true,
			},
			"access": {
				Type:     schema.TypeSet,
				ForceNew: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:     schema.TypeString,
							Required: true,
						},
						"resource": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// computed
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expires_at": {
				Type:     schema.TypeInt,
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
		},
	}
}

func resourceHarborProjectRobotAccountCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	var access []*apimodels.RobotAccountAccess

	for _, rawValue := range d.Get("access").(*schema.Set).List() {
		value := rawValue.(map[string]interface{})
		access = append(access, &apimodels.RobotAccountAccess{
			Action:   value["action"].(string),
			Resource: value["resource"].(string),
		})
	}

	robotAccountCreate := apimodels.RobotAccountCreate{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Access:      access,
	}

	projectID := int64(d.Get("project_id").(int))

	resp, err := client.Products.PostProjectsProjectIDRobots(
		products.NewPostProjectsProjectIDRobotsParamsWithContext(context.TODO()).
			WithProjectID(projectID).
			WithRobot(&robotAccountCreate),
		nil,
	)

	if err != nil {
		log.Printf("[DEBUG] Robot account creation failed")
		return err
	}

	d.Set("username", resp.Payload.Name)
	d.Set("token", resp.Payload.Token)

	listResp, err := client.Products.GetProjectsProjectIDRobots(
		products.NewGetProjectsProjectIDRobotsParamsWithContext(context.TODO()).
			WithProjectID(projectID),
		nil,
	)

	if err != nil {
		log.Printf("[DEBUG] Robot account loading failed")
		return err
	}

	var foundRobot *apimodels.RobotAccount
	for _, robot := range listResp.Payload {
		if robot.Name == resp.Payload.Name {
			foundRobot = robot
			break
		}
	}

	if foundRobot == nil {
		return fmt.Errorf("could not found robot %s", resp.Payload.Name)
	}

	resourceHarborProjectRobotAccountRefresh(d, foundRobot)

	return nil
}

func resourceHarborProjectRobotAccountRefresh(d *schema.ResourceData, r *apimodels.RobotAccount) {
	d.SetId(fmt.Sprintf("%d/%d", r.ProjectID, r.ID))
	d.Set("description", r.Description)
	d.Set("username", r.Name)
	d.Set("creation_time", r.CreationTime)
	d.Set("update_time", r.UpdateTime)
	d.Set("project_id", r.ProjectID)
	d.Set("disabled", r.Disabled)
}

func resourceHarborProjectRobotAccountParseID(id string) (int64, int64, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid id %s", id)
	}

	projectID, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	robotID, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return int64(projectID), int64(robotID), nil
}

func resourceHarborProjectRobotAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	projectID, robotID, err := resourceHarborProjectRobotAccountParseID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Products.GetProjectsProjectIDRobotsRobotID(
		products.NewGetProjectsProjectIDRobotsRobotIDParamsWithContext(context.TODO()).
			WithProjectID(projectID).
			WithRobotID(robotID),
		nil,
	)

	if err != nil {
		return err
	}

	resourceHarborProjectRobotAccountRefresh(d, resp.Payload)

	return nil
}

func resourceHarborProjectRobotAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	projectID, robotID, err := resourceHarborProjectRobotAccountParseID(d.Id())
	if err != nil {
		return err
	}

	disabled := d.Get("disabled").(bool)

	_, err = client.Products.PutProjectsProjectIDRobotsRobotID(
		products.NewPutProjectsProjectIDRobotsRobotIDParamsWithContext(context.TODO()).
			WithProjectID(projectID).
			WithRobotID(robotID).
			WithRobot(&apimodels.RobotAccountUpdate{
				Disabled: disabled,
			}),
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func resourceHarborProjectRobotAccountDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	projectID, robotID, err := resourceHarborProjectRobotAccountParseID(d.Id())
	if err != nil {
		return err
	}

	_, err = client.Products.DeleteProjectsProjectIDRobotsRobotID(
		products.NewDeleteProjectsProjectIDRobotsRobotIDParamsWithContext(context.TODO()).
			WithProjectID(projectID).
			WithRobotID(robotID),
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}
