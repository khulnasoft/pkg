package aws

import (
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/apigateway"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/athena"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/cloudfront"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/cloudtrail"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/cloudwatch"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/codebuild"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/config"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/documentdb"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/dynamodb"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/ec2"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/ecr"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/ecs"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/efs"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/eks"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/elasticache"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/elasticsearch"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/elb"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/iam"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/kinesis"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/lambda"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/mq"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/msk"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/neptune"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/rds"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/redshift"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/s3"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/sam"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/sns"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/sqs"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/ssm"
	"go.khulnasoft.com/pkg/iac/adapters/cloudformation/aws/workspaces"
	"go.khulnasoft.com/pkg/iac/providers/aws"
	"go.khulnasoft.com/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a Cloudformation AWS instance
func Adapt(cfFile parser.FileContext) aws.AWS {
	return aws.AWS{
		APIGateway:    apigateway.Adapt(cfFile),
		Athena:        athena.Adapt(cfFile),
		Cloudfront:    cloudfront.Adapt(cfFile),
		CloudTrail:    cloudtrail.Adapt(cfFile),
		CloudWatch:    cloudwatch.Adapt(cfFile),
		CodeBuild:     codebuild.Adapt(cfFile),
		Config:        config.Adapt(cfFile),
		DocumentDB:    documentdb.Adapt(cfFile),
		DynamoDB:      dynamodb.Adapt(cfFile),
		EC2:           ec2.Adapt(cfFile),
		ECR:           ecr.Adapt(cfFile),
		ECS:           ecs.Adapt(cfFile),
		EFS:           efs.Adapt(cfFile),
		IAM:           iam.Adapt(cfFile),
		EKS:           eks.Adapt(cfFile),
		ElastiCache:   elasticache.Adapt(cfFile),
		Elasticsearch: elasticsearch.Adapt(cfFile),
		ELB:           elb.Adapt(cfFile),
		MSK:           msk.Adapt(cfFile),
		MQ:            mq.Adapt(cfFile),
		Kinesis:       kinesis.Adapt(cfFile),
		Lambda:        lambda.Adapt(cfFile),
		Neptune:       neptune.Adapt(cfFile),
		RDS:           rds.Adapt(cfFile),
		Redshift:      redshift.Adapt(cfFile),
		S3:            s3.Adapt(cfFile),
		SAM:           sam.Adapt(cfFile),
		SNS:           sns.Adapt(cfFile),
		SQS:           sqs.Adapt(cfFile),
		SSM:           ssm.Adapt(cfFile),
		WorkSpaces:    workspaces.Adapt(cfFile),
	}
}
