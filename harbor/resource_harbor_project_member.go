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

func resourceHarborProjectMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceHarborProjectMemberCreate,
		Read:   resourceHarborProjectMemberRead,
		Delete: resourceHarborProjectMemberDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"group_type": {
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"group_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceHarborProjectMemberCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	groupName := d.Get("group_name").(string)

	projectMember := &apimodels.ProjectMember{
		RoleID: int64(d.Get("role_id").(int)),
		MemberGroup: &apimodels.UserGroup{
			GroupType: int64(d.Get("group_type").(int)),
			GroupName: groupName,
		},
	}

	projectID := int64(d.Get("project_id").(int))

	_, err := client.Products.PostProjectsProjectIDMembers(
		products.NewPostProjectsProjectIDMembersParamsWithContext(context.TODO()).
			WithProjectID(projectID).
			WithProjectMember(projectMember),
		nil,
	)

	if err != nil {
		log.Printf("[DEBUG] Project member creation failed")
		return err
	}

	listResp, err := client.Products.GetProjectsProjectIDMembers(
		products.NewGetProjectsProjectIDMembersParamsWithContext(context.TODO()).
			WithProjectID(projectID),
		nil,
	)

	if err != nil {
		log.Printf("[DEBUG] Project member loading failed")
		return err
	}

	var foundMember *apimodels.ProjectMemberEntity
	for _, member := range listResp.Payload {
		if member.EntityType == "g" && member.EntityName == groupName {
			foundMember = member
			break
		}
	}

	if foundMember == nil {
		return fmt.Errorf("could not find member %s", groupName)
	}

	resourceHarborProjectMemberRefresh(d, foundMember)

	return nil
}

func resourceHarborProjectMemberRefresh(d *schema.ResourceData, r *apimodels.ProjectMemberEntity) {
	d.SetId(fmt.Sprintf("%d/%d", r.ProjectID, r.ID))
	d.Set("role_id", r.RoleID)
	d.Set("group_name", r.EntityName)
}

func resourceHarborProjectMemberParseID(id string) (int64, int64, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid id %s", id)
	}

	projectID, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	memberID, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return int64(projectID), int64(memberID), nil
}

func resourceHarborProjectMemberRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	projectID, memberID, err := resourceHarborProjectMemberParseID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Products.GetProjectsProjectIDMembersMid(
		products.NewGetProjectsProjectIDMembersMidParamsWithContext(context.TODO()).
			WithProjectID(projectID).
			WithMid(memberID),
		nil,
	)

	if err != nil {
		return err
	}

	resourceHarborProjectMemberRefresh(d, resp.Payload)

	return nil
}

func resourceHarborProjectMemberDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	projectID, memberID, err := resourceHarborProjectMemberParseID(d.Id())
	if err != nil {
		return err
	}

	_, err = client.Products.DeleteProjectsProjectIDMembersMid(
		products.NewDeleteProjectsProjectIDMembersMidParamsWithContext(context.TODO()).
			WithProjectID(projectID).
			WithMid(memberID),
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}
