package all

import (
	_ "go.khulnasoft.com/pkg/fanal/analyzer/config/azurearm"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/config/cloudformation"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/config/dockerfile"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/config/helm"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/config/json"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/config/k8s"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/config/terraform"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/config/terraformplan/json"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/config/terraformplan/snapshot"
	_ "go.khulnasoft.com/pkg/fanal/analyzer/config/yaml"
)
