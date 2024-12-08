package khulnasoft

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftPermissionsSetDatasource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKhulnasoftPermissionsSetDataSource(),
				Check:  testAccCheckKhulnasoftPermissionsSetDataSourceExists("data.khulnasoft_permissions_sets.testpermissionsset"),
			},
		},
	})
}

func testAccCheckKhulnasoftPermissionsSetDataSource() string {
	return `
	data "khulnasoft_permissions_sets" "testpermissionsset" {}
	`
}

func testAccCheckKhulnasoftPermissionsSetDataSourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return NewNotFoundErrorf("%s in state", n)
		}

		if rs.Primary.ID == "" {
			return NewNotFoundErrorf("Id for %s in state", n)
		}

		return nil
	}
}
