package aws

import (
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/apigateway"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/athena"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/cloudfront"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/cloudtrail"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/cloudwatch"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/codebuild"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/config"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/documentdb"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/dynamodb"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/ec2"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/ecr"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/ecs"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/efs"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/eks"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/elasticache"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/elasticsearch"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/elb"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/emr"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/iam"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/kinesis"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/kms"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/lambda"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/mq"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/msk"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/neptune"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/provider"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/rds"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/redshift"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/s3"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/sns"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/sqs"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/ssm"
	"go.khulnasoft.com/pkg/iac/adapters/terraform/aws/workspaces"
	"go.khulnasoft.com/pkg/iac/providers/aws"
	"go.khulnasoft.com/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) aws.AWS {
	return aws.AWS{
		Meta: aws.Meta{
			TFProviders: provider.Adapt(modules),
		},
		APIGateway:    apigateway.Adapt(modules),
		Athena:        athena.Adapt(modules),
		Cloudfront:    cloudfront.Adapt(modules),
		CloudTrail:    cloudtrail.Adapt(modules),
		CloudWatch:    cloudwatch.Adapt(modules),
		CodeBuild:     codebuild.Adapt(modules),
		Config:        config.Adapt(modules),
		DocumentDB:    documentdb.Adapt(modules),
		DynamoDB:      dynamodb.Adapt(modules),
		EC2:           ec2.Adapt(modules),
		ECR:           ecr.Adapt(modules),
		ECS:           ecs.Adapt(modules),
		EFS:           efs.Adapt(modules),
		EKS:           eks.Adapt(modules),
		ElastiCache:   elasticache.Adapt(modules),
		Elasticsearch: elasticsearch.Adapt(modules),
		ELB:           elb.Adapt(modules),
		EMR:           emr.Adapt(modules),
		IAM:           iam.Adapt(modules),
		Kinesis:       kinesis.Adapt(modules),
		KMS:           kms.Adapt(modules),
		Lambda:        lambda.Adapt(modules),
		MQ:            mq.Adapt(modules),
		MSK:           msk.Adapt(modules),
		Neptune:       neptune.Adapt(modules),
		RDS:           rds.Adapt(modules),
		Redshift:      redshift.Adapt(modules),
		S3:            s3.Adapt(modules),
		SNS:           sns.Adapt(modules),
		SQS:           sqs.Adapt(modules),
		SSM:           ssm.Adapt(modules),
		WorkSpaces:    workspaces.Adapt(modules),
	}
}
