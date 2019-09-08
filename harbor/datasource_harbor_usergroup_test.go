package harbor

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccHarborUserGroupDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccHarborUserGroupDataSourceConfig(expectedDataSourceUserGroupName),
				Check: resource.ComposeTestCheckFunc(
					testAccHarborUserGroupDataSource("data.harbor_usergroup.bar"),
				),
			},
		},
	})
}

func testAccHarborUserGroupDataSource(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["id"] == "" {
			return fmt.Errorf("Expected to read user group data from Harbor")
		}

		if a["name"] != expectedDataSourceUserGroupName {
			return fmt.Errorf("Expected the user group name to be: %s, but got: %s", expectedDataSourceUserGroupName, a["name"])
		}
		return nil
	}
}

func testAccHarborUserGroupDataSourceConfig(rName string) string {
	return fmt.Sprintf(`
resource "harbor_usergroup" "foo" {
	name = "%[1]s"
	type = 2
}
data "harbor_usergroup" "bar" {
	name = "${harbor_usergroup.foo.name}"
}
`, rName)
}
