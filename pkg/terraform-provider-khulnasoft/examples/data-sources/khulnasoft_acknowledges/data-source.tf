data "khulnasoft_acknowledges" "acknowledges" {}

output "acknowledges" {
  value = data.khulnasoft_acknowledges.acknowledges
}
