data "khulnasoft_roles_mapping" "roles_mapping" {}

output "role_mapping_all" {
  value = data.khulnasoft_roles_mapping.roles_mapping
}

output "role_mapping_saml" {
  value = data.khulnasoft_roles_mapping.roles_mapping.saml
}
