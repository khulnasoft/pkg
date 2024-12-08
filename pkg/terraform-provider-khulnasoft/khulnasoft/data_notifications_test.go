package khulnasoft

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftNotificationsDatasource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKhulnasoftNotificationsDataSource(),
				Check:  testAccCheckKhulnasoftNotificationsDataSourceExists("data.khulnasoft_notifications.testnotifications"),
			},
		},
	})
}

func testAccCheckKhulnasoftNotificationsDataSource() string {
	return `
	data "khulnasoft_notifications" "testnotifications" {}
	`
}

func testAccCheckKhulnasoftNotificationsDataSourceExists(n string) resource.TestCheckFunc {
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
