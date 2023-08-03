package adapter

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Database struct {
	connection *dynamodb.Client
	logMode    bool
}

type Interface interface {
	Health() bool
	FindAll(condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error)
	FindOne(condtion map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error)
	CreateOrUpdate(entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error)
	Delete(condtion map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error)
}

func NewAdapter(con *dynamodb.Client) Interface {
	return &Database{
		connection: con,
		logMode:    false,
	}
}

func (db *Database) Health() bool {
	ctx := context.Background()
	_, err := db.connection.ListTables(ctx, &dynamodb.ListTablesInput{})
	return err == nil
}

func (db *Database) FindAll(condition expression.Expression, tableName string) (*dynamodb.ScanOutput, error) {

	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  condition.Names(),
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		TableName:                 aws.String(tableName),
	}

	rs, err := db.connection.Scan(context.Background(), input)

	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (db *Database) FindOne(condtion map[string]interface{}, tableName string) (*dynamodb.GetItemOutput, error) {

	condtionParsed, err := attributevalue.MarshalMap(condtion)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       condtionParsed,
	}

	rs, err := db.connection.GetItem(context.Background(), input)

	if err != nil {
		return nil, err
	}

	return rs, nil
}

func (db *Database) CreateOrUpdate(entity interface{}, tableName string) (*dynamodb.PutItemOutput, error) {

	entityParsed, err := attributevalue.MarshalMap(entity)

	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      entityParsed,
	}

	rs, err := db.connection.PutItem(context.Background(), input)

	if err != nil {
		return nil, err
	}

	return rs, nil

}

func (db *Database) Delete(condtion map[string]interface{}, tableName string) (*dynamodb.DeleteItemOutput, error) {

	condtionParsed, err := attributevalue.MarshalMap(condtion)

	if err != nil {
		return nil, err
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key:       condtionParsed,
	}

	rs, err := db.connection.DeleteItem(context.Background(), input)

	if err != nil {
		return nil, err
	}

	return rs, nil
}
