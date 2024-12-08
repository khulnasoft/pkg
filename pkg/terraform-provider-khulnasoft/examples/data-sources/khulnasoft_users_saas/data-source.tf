data "khulnasoft_users" "users" {}

output "first_user_email" {
  value = data.khulnasoft_users_saas.users.users[0].email
}