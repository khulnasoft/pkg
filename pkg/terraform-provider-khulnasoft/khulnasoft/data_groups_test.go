package khulnasoft

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftGroupsDatasource(t *testing.T) {

	if !isSaasEnv() {
		t.Skip("Skipping saas groups test because its on prem env")
	}
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKhulnasoftGroupsDataSource(),
				Check:  testAccCheckKhulnasoftGroupsDataSourceExists("data.khulnasoft_groups.testgroups"),
			},
		},
	})
}

func testAccCheckKhulnasoftGroupsDataSource() string {
	return `
	
	data "khulnasoft_groups" "testgroups" {}
	`

}

func testAccCheckKhulnasoftGroupsDataSourceExists(n string) resource.TestCheckFunc {
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
