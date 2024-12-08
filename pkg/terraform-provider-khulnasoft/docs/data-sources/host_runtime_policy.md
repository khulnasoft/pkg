---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "khulnasoft_host_runtime_policy Data Source - terraform-provider-khulnasoft"
subcategory: ""
description: |-
  
---

# khulnasoft_host_runtime_policy (Data Source)



## Example Usage

```terraform
data "khulnasoft_host_runtime_policy" "host_runtime_policy" {
  name = "hostRuntimePolicyName"
}

output "host_runtime_policy_details" {
  value = data.khulnasoft_host_runtime_policy.host_runtime_policy
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the host runtime policy

### Optional

- `auditing` (Block List, Max: 1) (see [below for nested schema](#nestedblock--auditing))
- `file_integrity_monitoring` (Block List) Configuration for file integrity monitoring. (see [below for nested schema](#nestedblock--file_integrity_monitoring))
- `malware_scan_options` (Block List, Max: 1) Configuration for Real-Time Malware Protection. (see [below for nested schema](#nestedblock--malware_scan_options))
- `package_block` (Block List, Max: 1) (see [below for nested schema](#nestedblock--package_block))

### Read-Only

- `application_scopes` (List of String) Indicates the application scope of the service.
- `audit_all_os_user_activity` (Boolean) If true, all process activity will be audited.
- `audit_brute_force_login` (Boolean) Detects brute force login attempts
- `audit_full_command_arguments` (Boolean) If true, full command arguments will be audited.
- `audit_host_failed_login_events` (Boolean) If true, host failed logins will be audited.
- `audit_host_successful_login_events` (Boolean) If true, host successful logins will be audited.
- `audit_user_account_management` (Boolean) If true, account management will be audited.
- `author` (String) Username of the account that created the service.
- `block_cryptocurrency_mining` (Boolean) Detect and prevent communication to DNS/IP addresses known to be used for Cryptocurrency Mining
- `blocked_files` (List of String) List of files that are prevented from being read, modified and executed in the containers.
- `description` (String) The description of the host runtime policy
- `enable_ip_reputation` (Boolean) If true, detect and prevent communication from containers to IP addresses known to have a bad reputation.
- `enabled` (Boolean) Indicates if the runtime policy is enabled or not.
- `enforce` (Boolean) Indicates that policy should effect container execution (not just for audit).
- `enforce_after_days` (Number) Indicates the number of days after which the runtime policy will be changed to enforce mode.
- `id` (String) The ID of this resource.
- `monitor_system_log_integrity` (Boolean) If true, system log will be monitored.
- `monitor_system_time_changes` (Boolean) If true, system time changes will be monitored.
- `monitor_windows_services` (Boolean) If true, windows service operations will be monitored.
- `os_groups_allowed` (List of String) List of OS (Linux or Windows) groups that are allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
- `os_groups_blocked` (List of String) List of OS (Linux or Windows) groups that are not allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
- `os_users_allowed` (List of String) List of OS (Linux or Windows) users that are allowed to authenticate to the host, and block authentication requests from all others.
- `os_users_blocked` (List of String) List of OS (Linux or Windows) users that are not allowed to authenticate to the host, and block authentication requests from all others.
- `port_scanning_detection` (Boolean) If true, port scanning behaviors will be audited.
- `scope_expression` (String) Logical expression of how to compute the dependency of the scope variables.
- `scope_variables` (List of Object) List of scope attributes. (see [below for nested schema](#nestedatt--scope_variables))
- `windows_registry_monitoring` (List of Object) Configuration for windows registry monitoring. (see [below for nested schema](#nestedatt--windows_registry_monitoring))
- `windows_registry_protection` (List of Object) Configuration for windows registry protection. (see [below for nested schema](#nestedatt--windows_registry_protection))

<a id="nestedblock--auditing"></a>
### Nested Schema for `auditing`

Optional:

- `audit_all_network` (Boolean)
- `audit_all_processes` (Boolean)
- `audit_failed_login` (Boolean)
- `audit_os_user_activity` (Boolean)
- `audit_process_cmdline` (Boolean)
- `audit_success_login` (Boolean)
- `audit_user_account_management` (Boolean)
- `enabled` (Boolean)


<a id="nestedblock--file_integrity_monitoring"></a>
### Nested Schema for `file_integrity_monitoring`

Optional:

- `enabled` (Boolean) If true, file integrity monitoring is enabled.
- `exceptional_monitored_files` (List of String) List of paths to be excluded from monitoring.
- `exceptional_monitored_files_processes` (List of String) List of processes to be excluded from monitoring.
- `exceptional_monitored_files_users` (List of String) List of users to be excluded from monitoring.
- `monitored_files` (List of String) List of paths to be monitored.
- `monitored_files_attributes` (Boolean) Whether to monitor file attribute operations.
- `monitored_files_create` (Boolean) Whether to monitor file create operations.
- `monitored_files_delete` (Boolean) Whether to monitor file delete operations.
- `monitored_files_modify` (Boolean) Whether to monitor file modify operations.
- `monitored_files_processes` (List of String) List of processes associated with monitored files.
- `monitored_files_read` (Boolean) Whether to monitor file read operations.
- `monitored_files_users` (List of String) List of users associated with monitored files.


<a id="nestedblock--malware_scan_options"></a>
### Nested Schema for `malware_scan_options`

Optional:

- `action` (String) Set Action, Defaults to 'Alert' when empty
- `enabled` (Boolean) Defines if enabled or not
- `exclude_directories` (List of String) List of registry paths to be excluded from being protected.
- `exclude_processes` (List of String) List of registry processes to be excluded from being protected.
- `include_directories` (List of String) List of registry paths to be excluded from being protected.


<a id="nestedblock--package_block"></a>
### Nested Schema for `package_block`

Optional:

- `block_packages_processes` (List of String)
- `block_packages_users` (List of String)
- `enabled` (Boolean)
- `exceptional_block_packages_files` (List of String)
- `exceptional_block_packages_processes` (List of String)
- `exceptional_block_packages_users` (List of String)
- `packages_black_list` (List of String)


<a id="nestedatt--scope_variables"></a>
### Nested Schema for `scope_variables`

Read-Only:

- `attribute` (String)
- `name` (String)
- `value` (String)


<a id="nestedatt--windows_registry_monitoring"></a>
### Nested Schema for `windows_registry_monitoring`

Read-Only:

- `excluded_paths` (List of String)
- `excluded_processes` (List of String)
- `excluded_users` (List of String)
- `monitor_attributes` (Boolean)
- `monitor_create` (Boolean)
- `monitor_delete` (Boolean)
- `monitor_modify` (Boolean)
- `monitor_read` (Boolean)
- `monitored_paths` (List of String)
- `monitored_processes` (List of String)
- `monitored_users` (List of String)


<a id="nestedatt--windows_registry_protection"></a>
### Nested Schema for `windows_registry_protection`

Read-Only:

- `excluded_paths` (List of String)
- `excluded_processes` (List of String)
- `excluded_users` (List of String)
- `protected_paths` (List of String)
- `protected_processes` (List of String)
- `protected_users` (List of String)

