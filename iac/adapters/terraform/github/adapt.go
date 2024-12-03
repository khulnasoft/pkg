package github

import (
	"go.khulnasoft.com/pkg/iac/adapters/terraform/github/branch_protections"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/github/repositories"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/github/secrets"
	"go.khulnasoft.com/pkg/iac/providers/github"
	"go.khulnasoft.com/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) github.GitHub {
	return github.GitHub{
		Repositories:       repositories.Adapt(modules),
		EnvironmentSecrets: secrets.Adapt(modules),
		BranchProtections:  branch_protections.Adapt(modules),
	}
}
