package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftresourceKhulnasoftLabel(t *testing.T) {
	t.Parallel()
	name := acctest.RandomWithPrefix("terraform-test")
	description := "terraform-test"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: CheckDestroy("khulnasoft_khulnasoft_label.new"),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKhulnasoftKhulnasoftLabel(name, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKhulnasoftKhulnasoftLabelExists("khulnasoft_khulnasoft_label.new"),
				),
			},
			{
				ResourceName:      "khulnasoft_khulnasoft_label.new",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckKhulnasoftKhulnasoftLabel(name, description string) string {
	return fmt.Sprintf(`
	resource "khulnasoft_khulnasoft_label" "new" {
		name = "%s"
		description = "%s"
	}`, name, description)

}

func testAccCheckKhulnasoftKhulnasoftLabelExists(n string) resource.TestCheckFunc {
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
