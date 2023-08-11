package rulesproduct

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/Asad2730/DynamoDB_CRUD_App/internal/entities"
	"github.com/Asad2730/DynamoDB_CRUD_App/internal/entities/product"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type Interface interface{}

type Rules struct{}

func NewRules() Interface {
	return &Rules{}
}

func (r *Rules) ConvertIOReaderToStruct(data io.Reader, model interface{}) (body interface{}, err error) {
	if data == nil {
		return nil, errors.New("body is invalid")
	}

	return model, json.NewDecoder(data).Decode(model)
}

func (r *Rules) GetMock() interface{} {
	return product.Product{
		Base: entities.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: uuid.NewString(),
	}
}

func (r *Rules) Migrate(connection *dynamodb.Client) error {
	return r.CreateTable(connection)
}

func (r *Rules) Validate(model interface{}) error {

	model, err := product.InterfaceToModel(model)

	if err != nil {
		return err
	}

	return validation.ValidateStruct(&model, validation.Field(&model, validation.Required))
}

func (r *Rules) CreateTable(connection *dynamodb.Client) error {

	table := &product.Product{}

	key := []types.KeySchemaElement{
		{
			AttributeName: aws.String("_id"),
			KeyType:       types.KeyTypeHash,
		},
	}

	pro := &types.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(10),
		WriteCapacityUnits: aws.Int64(10),
	}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("_id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema:             key,
		ProvisionedThroughput: pro,
		TableName:             aws.String(table.TableName()),
	}

	_, err := connection.CreateTable(context.TODO(), input)

	return err
}
