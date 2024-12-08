data "khulnasoft_host_runtime_policy" "host_runtime_policy" {
  name = "hostRuntimePolicyName"
}

output "host_runtime_policy_details" {
  value = data.khulnasoft_host_runtime_policy.host_runtime_policy
}