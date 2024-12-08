package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	dynamodbApi "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodbTypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	awsAdapter "github.com/khulnasoft/tunnel-aws/internal/adapters/cloud/aws"
	"github.com/khulnasoft/tunnel-aws/pkg/concurrency"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/dynamodb"
	"github.com/khulnasoft/tunnel/pkg/iac/state"
	tunnelTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type adapter struct {
	*awsAdapter.RootAdapter
	client *dynamodbApi.Client
}

func init() {
	awsAdapter.RegisterServiceAdapter(&adapter{})
}

func (a *adapter) Name() string {
	return "dynamodb"
}

func (a *adapter) Provider() string {
	return "aws"
}

func (a *adapter) Adapt(root *awsAdapter.RootAdapter, state *state.State) error {
	a.RootAdapter = root
	a.client = dynamodbApi.NewFromConfig(root.SessionConfig())
	var err error

	state.AWS.DynamoDB.Tables, err = a.getTables()
	if err != nil {
		return err
	}

	return nil
}

func (a *adapter) getTables() (tables []dynamodb.Table, err error) {

	a.Tracker().SetServiceLabel("Discovering DynamoDB tables...")

	var apiTables []string
	var input dynamodbApi.ListTablesInput
	for {
		output, err := a.client.ListTables(a.Context(), &input)
		if err != nil {
			return nil, err
		}
		apiTables = append(apiTables, output.TableNames...)
		a.Tracker().SetTotalResources(len(apiTables))
		if output.LastEvaluatedTableName == nil {
			break
		}
		input.ExclusiveStartTableName = output.LastEvaluatedTableName
	}

	a.Tracker().SetServiceLabel("Adapting DynamoDB tables...")
	return concurrency.Adapt(apiTables, a.RootAdapter, a.adaptTable), nil

}

func (a *adapter) adaptTable(tableName string) (*dynamodb.Table, error) {

	tableMetadata := a.CreateMetadata(tableName)

	table, err := a.client.DescribeTable(a.Context(), &dynamodbApi.DescribeTableInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		return nil, err
	}
	encryption := dynamodb.ServerSideEncryption{
		Metadata: tableMetadata,
		Enabled:  tunnelTypes.BoolDefault(false, tableMetadata),
		KMSKeyID: tunnelTypes.StringDefault("", tableMetadata),
	}
	if table.Table.SSEDescription != nil {

		if table.Table.SSEDescription.Status == dynamodbTypes.SSEStatusEnabled {
			encryption.Enabled = tunnelTypes.BoolDefault(true, tableMetadata)
		}

		if table.Table.SSEDescription.KMSMasterKeyArn != nil {
			encryption.KMSKeyID = tunnelTypes.StringDefault(*table.Table.SSEDescription.KMSMasterKeyArn, tableMetadata)
		}
	}
	pitRecovery := tunnelTypes.Bool(false, tableMetadata)
	continuousBackup, err := a.client.DescribeContinuousBackups(a.Context(), &dynamodbApi.DescribeContinuousBackupsInput{
		TableName: aws.String(tableName),
	})

	if err != nil && continuousBackup != nil && continuousBackup.ContinuousBackupsDescription != nil &&
		continuousBackup.ContinuousBackupsDescription.PointInTimeRecoveryDescription != nil {
		if continuousBackup.ContinuousBackupsDescription.PointInTimeRecoveryDescription.PointInTimeRecoveryStatus == dynamodbTypes.PointInTimeRecoveryStatusEnabled {
			pitRecovery = tunnelTypes.BoolDefault(true, tableMetadata)
		}

	}
	return &dynamodb.Table{
		Metadata:             tableMetadata,
		ServerSideEncryption: encryption,
		PointInTimeRecovery:  pitRecovery,
	}, nil
}