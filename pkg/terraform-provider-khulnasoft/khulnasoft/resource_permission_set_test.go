package khulnasoft

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftPermissionSetManagement(t *testing.T) {
	t.Parallel()
	name := acctest.RandomWithPrefix("terraform")
	description := "created from terraform "
	author := "system"
	ui_access := true
	is_super := false
	actions := "risks.vulnerabilities.read,images.read"

	if isSaasEnv() {
		// todo: remove this after solving the following issue: https://scalock.atlassian.net/browse/SLK-62403
		t.Skip("Skipping user test because its saas env")
		author = os.Getenv("KHULNASOFT_USER")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccPermissionSetDestroy,
		Steps: []resource.TestStep{
			{
				// Config returns the test resource
				Config: testAccCheckKhulnasoftPermissionSet(name, description, author, ui_access, is_super, actions),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKhulnasoftPermissionSetExists("khulnasoft_permissions_sets.new"),
				),
			},
			{
				ResourceName:      "khulnasoft_permissions_sets.new",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckKhulnasoftPermissionSet(name string, description string, author string, ui_access bool, is_super bool, actions string) string {
	return fmt.Sprintf(`
	resource "khulnasoft_permissions_sets" "new" {
		name = "%s"
		description     = "%s"
		author = "%s"
		ui_access = "%v"
		is_super = "%v"
		actions = [
		  "%s"
		]
	  }`, name, description, author, ui_access, is_super, actions)
}

func testAccCheckKhulnasoftPermissionSetExists(n string) resource.TestCheckFunc {
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

func testAccPermissionSetDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "khulnasoft_permissions_sets.new" {
			continue
		}

		if rs.Primary.ID != "" {
			return fmt.Errorf("Object %q still exists", rs.Primary.ID)
		}
		return nil
	}
	return nil
}
