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

func dataSourceHarborUserGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHarborUserGroupRead,
		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"name"},
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"group_id"},
			},

			// computed
			"type": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ldap_group_dn": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceHarborUserGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*apiclient.Harbor)

	set := 0
	var usergroupID int64
	if id, ok := d.GetOk("group_id"); ok {
		usergroupID = int64(id.(int))
		set++
	}

	var usergroupName string
	if name, ok := d.GetOk("name"); ok {
		usergroupName = name.(string)
		set++
	}

	if set != 1 {
		return fmt.Errorf("One of %q or %q has to be provided", "group_id", "name")
	}

	var usergroup *apimodels.UserGroup

	if usergroupName != "" {
		resp, err := client.Products.GetUsergroups(
			products.NewGetUsergroupsParamsWithContext(context.TODO()),
			nil,
		)
		if err != nil {
			return err
		}

		for _, p := range resp.Payload {
			if strings.ToLower(p.GroupName) == strings.ToLower(usergroupName) {
				usergroup = p
				break
			}
		}
	} else {
		resp, err := client.Products.GetUsergroupsGroupID(
			products.NewGetUsergroupsGroupIDParamsWithContext(context.TODO()).
				WithGroupID(usergroupID),
			nil,
		)
		if err != nil {
			return err
		}

		usergroup = resp.Payload
	}

	if usergroup == nil {
		return fmt.Errorf("Harbor User Group not found")
	}

	harborUserGroupUpdate(d, usergroup)

	return nil
}
