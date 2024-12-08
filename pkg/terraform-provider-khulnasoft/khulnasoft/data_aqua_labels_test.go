package khulnasoft

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftKhulnasoftLabelsDatasource(t *testing.T) {
	t.Parallel()
	name := acctest.RandomWithPrefix("terraform-test")
	description := "terraform-test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKhulnasoftKhulnasoftLabelsDataSource(name, description),
				Check:  testAccCheckKhulnasoftKhulnasoftLabelsDataSourceExists("data.khulnasoft_khulnasoft_labels.test_khulnasoft_labels"),
			},
		},
	})
}

func testAccCheckKhulnasoftKhulnasoftLabelsDataSource(name, description string) string {
	return fmt.Sprintf(`
	resource "khulnasoft_khulnasoft_label" "new" {
		name = "%s"
		description = "%s"
	}

	data "khulnasoft_khulnasoft_labels" "test_khulnasoft_labels" {
	depends_on = [khulnasoft_khulnasoft_label.new]
	}
	`, name, description)

}

func testAccCheckKhulnasoftKhulnasoftLabelsDataSourceExists(n string) resource.TestCheckFunc {
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
