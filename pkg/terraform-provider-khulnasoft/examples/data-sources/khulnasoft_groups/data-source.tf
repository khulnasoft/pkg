data "khulnasoft_groups" "groups" {}

output "first_group_name" {
  value = data.khulnasoft_groups.groups.groups.0.name
}