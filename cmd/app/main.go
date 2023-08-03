package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Asad2730/DynamoDB_CRUD_App/config"
	"github.com/Asad2730/DynamoDB_CRUD_App/internal/repositories/adapter"
	"github.com/Asad2730/DynamoDB_CRUD_App/internal/repositories/instance"
	"github.com/Asad2730/DynamoDB_CRUD_App/internal/routes"
	"github.com/Asad2730/DynamoDB_CRUD_App/utils/logger"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	configs := config.GetConfig()
	connection, err := instance.GetConnection()

	if err != nil {
		fmt.Println("Failed to create connection", err.Error())
	}

	repository := adapter.NewAdapter(connection)
	logger.INFO("waiting for service to start ...", nil)
	errors := Migrate(connection)

	if len(errors) > 0 {
		for _, er := range errors {
			logger.PANIC("Error on migration ...", er)
		}
	}

	logger.PANIC("", checkTables(connection))

	port := fmt.Sprintf(":v%", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	logger.INFO("service is running on port", port)
	http.ListenAndServe(port, router)
}

func Migrate(connection *dynamodb.Client) []error {
	var errors []error
	callMigrateAndAppendError(&errors, connection, &RulesProduct.Rules{})
	return errors
}

func callMigrateAndAppendError(errors *[]error, connection *dynamodb.Client, rule rules.Interface) {
	err := rule.Migrate(connection)
	if err != nil {
		*errors = append(*errors, err)
	}
}

func checkTables(connection *dynamodb.Client) error {
	result, err := connection.ListTables(context.Background(), &dynamodb.ListTablesInput{})

	if result != nil {
		if len(result.TableNames) == 0 {
			logger.INFO("Tables not found", nil)
		}

		for _, tableName := range result.TableNames {
			logger.INFO("Table Found:", tableName)
		}
	}

	return err
}
