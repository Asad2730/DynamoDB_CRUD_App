package product

import (
	"encoding/json"
	"time"

	"github.com/Asad2730/DynamoDB_CRUD_App/internal/entities"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Product struct {
	entities.Base
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

func InterfaceToModel(data interface{}) (instance *Product, err error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return instance, err
	}

	return instance, json.Unmarshal(bytes, &instance)
}

func (p *Product) GetFilterId() map[string]interface{} {
	return nil
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) Bytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Product) GetMap() map[string]interface{} {
	return nil
}

func (p *Product) ParseDynamoAttributeToStruct(res *dynamodb.GetItemOutput) map[string]interface{} {
	return nil
}
