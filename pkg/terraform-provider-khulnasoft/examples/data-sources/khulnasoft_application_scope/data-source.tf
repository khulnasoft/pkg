data "khulnasoft_application_scope" "default" {
  name = "Global"
}

output "scopes" {
  value = data.khulnasoft_application_scope.default
}