package khulnasoft

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestKhulnasoftRoleManagement(t *testing.T) {

	if isSaasEnv() {
		t.Skip("Skipping prem role test because its on saas env")
	}
	t.Parallel()
	roleName := acctest.RandomWithPrefix("roleTest")
	description := "roleTest1"
	newDescription := "roleTest2"
	permission := "Administrator"
	scope := "Global"
	//roleNewName := roleName + "new"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccRoleDestroy,
		Steps: []resource.TestStep{
			{
				// Config returns the test resource
				Config: testAccCheckKhulnasoftRole(roleName, description, permission, scope),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKhulnasoftRolesExists("khulnasoft_role.new"),
				),
			},
			{
				// Config returns the test resource
				Config: testAccCheckKhulnasoftRole(roleName, newDescription, permission, scope),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKhulnasoftRolesExists("khulnasoft_role.new"),
				),
			},
			{
				ResourceName:      "khulnasoft_role.new",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckKhulnasoftRole(roleName, description, permission, scope string) string {
	return fmt.Sprintf(`
	resource "khulnasoft_role" "new" {
		role_name   = "%s"
		description = "%s"
		permission = "%s"
		scopes = ["%s"]
    }`, roleName, description, permission, scope)
}

func testAccCheckKhulnasoftRolesExists(n string) resource.TestCheckFunc {
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

func testAccRoleDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "khulnasoft_role.new" {
			continue
		}

		if rs.Primary.ID != "" {
			return fmt.Errorf("Object %q still exists", rs.Primary.ID)
		}
		return nil
	}
	return nil
}
