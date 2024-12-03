package ec2

import (
	"go.khulnasoft.com/pkg/iac/providers/aws/ec2"
	"go.khulnasoft.com/pkg/iac/scanners/cloudformation/parser"
)

func getVolumes(ctx parser.FileContext) (volumes []ec2.Volume) {

	volumeResources := ctx.GetResourcesByType("AWS::EC2::Volume")
	for _, r := range volumeResources {

		volume := ec2.Volume{
			Metadata: r.Metadata(),
			Encryption: ec2.Encryption{
				Metadata: r.Metadata(),
				Enabled:  r.GetBoolProperty("Encrypted"),
				KMSKeyID: r.GetStringProperty("KmsKeyId"),
			},
		}

		volumes = append(volumes, volume)
	}
	return volumes
}
