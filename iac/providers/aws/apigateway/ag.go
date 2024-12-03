package apigateway

import (
	v1 "go.khulnasoft.com/pkg/iac/providers/aws/apigateway/v1"
	v2 "go.khulnasoft.com/pkg/iac/providers/aws/apigateway/v2"
)

type APIGateway struct {
	V1 v1.APIGateway
	V2 v2.APIGateway
}
