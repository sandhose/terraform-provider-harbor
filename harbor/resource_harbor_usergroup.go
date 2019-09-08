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

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceHarborUserGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceHarborUserGroupCreate,
		Read:   resourceHarborUserGroupRead,
		Update: resourceHarborUserGroupUpdate,
		Delete: resourceHarborUserGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntInSlice([]int{1, 2}),
			},
			"ldap_group_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},

			// computed
			"group_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceHarborUserGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)
	userGroupParams := &apimodels.UserGroup{
		GroupName: d.Get("name").(string),
	}

	groupType := int64(d.Get("type").(int))

	switch groupType {
	case 1:
		if attr, ok := d.GetOk("ldap_group_dn"); ok {
			userGroupParams.LdapGroupDn = attr.(string)
		}
	}

	_, err := client.Products.PostUsergroups(
		products.NewPostUsergroupsParamsWithContext(context.TODO()).
			WithUsergroup(userGroupParams),
		nil,
	)
	if err != nil {
		log.Printf("[DEBUG] User group creation failed")
		return err
	}

	resp, err := client.Products.GetUsergroups(
		products.NewGetUsergroupsParamsWithContext(context.TODO()),
		nil,
	)
	if err != nil {
		log.Printf("[DEBUG] User group loading failed")
		return err
	}

	var usergroup *apimodels.UserGroup
	for _, p := range resp.Payload {
		if strings.ToLower(p.GroupName) == strings.ToLower(userGroupParams.GroupName) {
			usergroup = p
			break
		}
	}

	if usergroup == nil {
		return fmt.Errorf("User group not found")
	}

	harborUserGroupUpdate(d, usergroup)

	return nil
}

func resourceHarborUserGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	usergroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	resp, err := client.Products.GetUsergroupsGroupID(
		products.NewGetUsergroupsGroupIDParamsWithContext(context.TODO()).
			WithGroupID(usergroupID),
		nil,
	)
	if err != nil {
		log.Printf("[DEBUG] User group loading failed")
		return err
	}

	harborUserGroupUpdate(d, resp.Payload)

	return nil
}

func resourceHarborUserGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)
	usergroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	// Only the name can be mutated
	usergroupParams := &apimodels.UserGroup{
		GroupName: d.Get("name").(string),
	}

	_, err = client.Products.PutUsergroupsGroupID(
		products.NewPutUsergroupsGroupIDParamsWithContext(context.TODO()).
			WithGroupID(usergroupID).
			WithUsergroup(usergroupParams),
		nil,
	)
	if err != nil {
		log.Printf("[DEBUG] User group update failed")
		return err
	}

	return nil
}

func resourceHarborUserGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	usergroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	_, err = client.Products.DeleteUsergroupsGroupID(
		products.NewDeleteUsergroupsGroupIDParamsWithContext(context.TODO()).
			WithGroupID(usergroupID),
		nil,
	)
	if err != nil {
		log.Printf("[DEBUG] User group deletion failed")
		return err
	}

	return nil
}

func harborUserGroupUpdate(d *schema.ResourceData, u *apimodels.UserGroup) {
	d.SetId(fmt.Sprint(u.ID))

	d.Set("group_id", u.ID)
	d.Set("name", u.GroupName)
	d.Set("type", u.GroupType)
	d.Set("ldap_group_dn", u.LdapGroupDn)
}
