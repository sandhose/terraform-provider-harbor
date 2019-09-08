package harbor

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"
	"github.com/sandhose/terraform-provider-harbor/api/client/products"
)

func TestAccHarborUserGroup_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHarborUserGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHarborUserGroupConfig(expectedResourceUserGroupName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHarborUserGroupExists("harbor_usergroup.foo"),
					resource.TestCheckResourceAttr(
						"harbor_usergroup.foo", "name", expectedResourceUserGroupName),
					resource.TestCheckResourceAttr(
						"harbor_usergroup.foo", "type", "2"),
				),
			},
			{
				Config: testAccHarborUserGroupConfigUpdated(expectedResourceUpdatedUserGroupName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHarborUserGroupExists("harbor_usergroup.foo"),
					resource.TestCheckResourceAttr(
						"harbor_usergroup.foo", "name", expectedResourceUpdatedUserGroupName),
					resource.TestCheckResourceAttr(
						"harbor_usergroup.foo", "type", "2"),
				),
			},
		},
	})
}

func testAccCheckHarborUserGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*apiclient.Harbor)
	for _, r := range s.RootModule().Resources {
		if r.Type != "harbor_usergroup" {
			continue
		}

		id, err := strconv.ParseInt(r.Primary.ID, 10, 32)
		if err != nil {
			return err
		}

		_, err = client.Products.GetUsergroupsGroupID(
			products.NewGetUsergroupsGroupIDParamsWithContext(context.TODO()).
				WithGroupID(int64(id)),
			nil,
		)
		if err == nil {
			return fmt.Errorf("Harbor user group still exists")
		}

	}
	return nil
}

func testAccCheckHarborUserGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		a := rs.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to read user group data from Harbor")
		}

		client := testAccProvider.Meta().(*apiclient.Harbor)

		id, err := strconv.ParseInt(a["id"], 10, 32)
		if err != nil {
			return err
		}

		resp, err := client.Products.GetUsergroupsGroupID(
			products.NewGetUsergroupsGroupIDParamsWithContext(context.TODO()).
				WithGroupID(int64(id)),
			nil,
		)
		if err != nil {
			return err
		}

		found := resp.Payload

		if strconv.Itoa(int(found.ID)) != a["id"] {
			return fmt.Errorf("Harbor user group not found: %v - %v", a["id"], found)
		}

		return nil
	}
}

func testAccHarborUserGroupConfig(rName string) string {
	return fmt.Sprintf(`
resource "harbor_usergroup" "foo" {
	name = "%s"
	type = 2
}
`, rName)
}

func testAccHarborUserGroupConfigUpdated(rName string) string {
	return fmt.Sprintf(`
resource "harbor_usergroup" "foo" {
	name = "%s"
	type = 2
}
`, rName)
}
