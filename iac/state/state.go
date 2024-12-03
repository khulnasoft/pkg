package state

import (
	"reflect"

	"go.khulnasoft.com/pkg/iac/providers/aws"
	"go.khulnasoft.com/pkg/iac/providers/azure"
	"go.khulnasoft.com/pkg/iac/providers/cloudstack"
	"go.khulnasoft.com/pkg/iac/providers/digitalocean"
	"go.khulnasoft.com/pkg/iac/providers/github"
	"go.khulnasoft.com/pkg/iac/providers/google"
	"go.khulnasoft.com/pkg/iac/providers/kubernetes"
	"go.khulnasoft.com/pkg/iac/providers/nifcloud"
	"go.khulnasoft.com/pkg/iac/providers/openstack"
	"go.khulnasoft.com/pkg/iac/providers/oracle"
	"go.khulnasoft.com/pkg/iac/rego/convert"
)

type State struct {
	AWS          aws.AWS
	Azure        azure.Azure
	CloudStack   cloudstack.CloudStack
	DigitalOcean digitalocean.DigitalOcean
	GitHub       github.GitHub
	Google       google.Google
	Kubernetes   kubernetes.Kubernetes
	OpenStack    openstack.OpenStack
	Oracle       oracle.Oracle
	Nifcloud     nifcloud.Nifcloud
}

func (a *State) ToRego() any {
	return convert.StructToRego(reflect.ValueOf(a))
}
