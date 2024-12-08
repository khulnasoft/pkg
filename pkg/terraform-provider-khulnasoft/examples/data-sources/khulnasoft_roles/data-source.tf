data "khulnasoft_roles" "roles" {}

//Output the first role
output "first_user_name" {
  value = data.khulnasoft_roles.roles.roles[0]
}
