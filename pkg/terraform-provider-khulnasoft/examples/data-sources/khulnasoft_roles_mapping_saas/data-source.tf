data "khulnasoft_roles_mapping_saas" "roles_mapping_saas" {}

output "role_mapping" {
  value = data.khulnasoft_roles_mapping_saas.roles_mapping_saas.roles_mapping
}
