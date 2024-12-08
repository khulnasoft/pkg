data "khulnasoft_khulnasoft_labels" "khulnasoft_labels" {}

# Print all Khulnasoft labels
output "scopes" {
  value = data.khulnasoft_khulnasoft_labels.khulnasoft_labels
}