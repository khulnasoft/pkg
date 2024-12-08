data "khulnasoft_gateways" "testgateway" {}

output "gateway_data" {
  value = data.khulnasoft_gateways.testgateway
}

output "gateway_name" {
  value = data.khulnasoft_gateways.testgateway.gateways[0].id
}
output "gateway_status" {
  value = data.khulnasoft_gateways.testgateway.gateways[0].status
}
output "gateway_description" {
  value = data.khulnasoft_gateways.testgateway.gateways[0].description
}

output "gateway_version" {
  value = data.khulnasoft_gateways.testgateway.gateways[0].version
}

output "gateway_hostname" {
  value = data.khulnasoft_gateways.testgateway.gateways[0].hostname
}
output "gateway_grpc_address" {
  value = data.khulnasoft_gateways.testgateway.gateways[0].grpc_address
}