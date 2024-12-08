data "khulnasoft_integration_state" "integration_state" {}

output "khulnasoft_integration_state" {
  value = data.khulnasoft_integration_state.integration_state
}
