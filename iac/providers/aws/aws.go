package aws

import (
	"go.khulnasoft.com/pkg/iac/providers/aws/accessanalyzer"
	"go.khulnasoft.com/pkg/iac/providers/aws/apigateway"
	"go.khulnasoft.com/pkg/iac/providers/aws/athena"
	"go.khulnasoft.com/pkg/iac/providers/aws/cloudfront"
	"go.khulnasoft.com/pkg/iac/providers/aws/cloudtrail"
	"go.khulnasoft.com/pkg/iac/providers/aws/cloudwatch"
	"go.khulnasoft.com/pkg/iac/providers/aws/codebuild"
	"go.khulnasoft.com/pkg/iac/providers/aws/config"
	"go.khulnasoft.com/pkg/iac/providers/aws/documentdb"
	"go.khulnasoft.com/pkg/iac/providers/aws/dynamodb"
	"go.khulnasoft.com/pkg/iac/providers/aws/ec2"
	"go.khulnasoft.com/pkg/iac/providers/aws/ecr"
	"go.khulnasoft.com/pkg/iac/providers/aws/ecs"
	"go.khulnasoft.com/pkg/iac/providers/aws/efs"
	"go.khulnasoft.com/pkg/iac/providers/aws/eks"
	"go.khulnasoft.com/pkg/iac/providers/aws/elasticache"
	"go.khulnasoft.com/pkg/iac/providers/aws/elasticsearch"
	"go.khulnasoft.com/pkg/iac/providers/aws/elb"
	"go.khulnasoft.com/pkg/iac/providers/aws/emr"
	"go.khulnasoft.com/pkg/iac/providers/aws/iam"
	"go.khulnasoft.com/pkg/iac/providers/aws/kinesis"
	"go.khulnasoft.com/pkg/iac/providers/aws/kms"
	"go.khulnasoft.com/pkg/iac/providers/aws/lambda"
	"go.khulnasoft.com/pkg/iac/providers/aws/mq"
	"go.khulnasoft.com/pkg/iac/providers/aws/msk"
	"go.khulnasoft.com/pkg/iac/providers/aws/neptune"
	"go.khulnasoft.com/pkg/iac/providers/aws/rds"
	"go.khulnasoft.com/pkg/iac/providers/aws/redshift"
	"go.khulnasoft.com/pkg/iac/providers/aws/s3"
	"go.khulnasoft.com/pkg/iac/providers/aws/sam"
	"go.khulnasoft.com/pkg/iac/providers/aws/sns"
	"go.khulnasoft.com/pkg/iac/providers/aws/sqs"
	"go.khulnasoft.com/pkg/iac/providers/aws/ssm"
	"go.khulnasoft.com/pkg/iac/providers/aws/workspaces"
)

type AWS struct {
	Meta           Meta
	AccessAnalyzer accessanalyzer.AccessAnalyzer
	APIGateway     apigateway.APIGateway
	Athena         athena.Athena
	Cloudfront     cloudfront.Cloudfront
	CloudTrail     cloudtrail.CloudTrail
	CloudWatch     cloudwatch.CloudWatch
	CodeBuild      codebuild.CodeBuild
	Config         config.Config
	DocumentDB     documentdb.DocumentDB
	DynamoDB       dynamodb.DynamoDB
	EC2            ec2.EC2
	ECR            ecr.ECR
	ECS            ecs.ECS
	EFS            efs.EFS
	EKS            eks.EKS
	ElastiCache    elasticache.ElastiCache
	Elasticsearch  elasticsearch.Elasticsearch
	ELB            elb.ELB
	EMR            emr.EMR
	IAM            iam.IAM
	Kinesis        kinesis.Kinesis
	KMS            kms.KMS
	Lambda         lambda.Lambda
	MQ             mq.MQ
	MSK            msk.MSK
	Neptune        neptune.Neptune
	RDS            rds.RDS
	Redshift       redshift.Redshift
	SAM            sam.SAM
	S3             s3.S3
	SNS            sns.SNS
	SQS            sqs.SQS
	SSM            ssm.SSM
	WorkSpaces     workspaces.WorkSpaces
}

type Meta struct {
	TFProviders []TerraformProvider
}
