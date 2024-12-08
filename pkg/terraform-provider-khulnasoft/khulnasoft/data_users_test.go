package khulnasoft

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftUserManagementDatasource(t *testing.T) {

	if isSaasEnv() {
		t.Skip("Skipping user test because its saas env")
	}
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKhulnasoftUserDataSource(),
				Check:  testAccCheckKhulnasoftUsersDataSourceExists("data.khulnasoft_users.testusers"),
			},
		},
	})
}

func testAccCheckKhulnasoftUserDataSource() string {
	return `
	data "khulnasoft_users" "testusers" {}
	`
}

func testAccCheckKhulnasoftUsersDataSourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return NewNotFoundErrorf("%s in state", n)
		}

		if rs.Primary.ID == "" {
			return NewNotFoundErrorf("ID for %s in state", n)
		}

		return nil
	}
}
