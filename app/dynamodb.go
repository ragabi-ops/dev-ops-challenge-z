package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (hc *HealthCheck) ListTable() (HealthCheck, error) {
	sess, err := getSession()
	genericErrorHandler(err)

	dbSvc := dynamodb.New(sess)

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	health := HealthCheck{}
	health.DockerRepo = "ghcr.io/ragabi-ops/go-app"
	health.Status = "Healthy"

	tables, err := dbSvc.ListTablesWithContext(timeoutCtx, &dynamodb.ListTablesInput{})
	if err != nil {
		log.Println(err.Error())
		health.Status = "Unhealty"
		return health, err
	}

	log.Println("Tables:")
	for _, table := range tables.TableNames {
		log.Println(*table)
	}

	return health, nil
}

func (it *Item) getItem() (Item, error) {
	sess, err := getSession()
	genericErrorHandler(err)

	dbSvc := dynamodb.New(sess)

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := dbSvc.GetItemWithContext(timeoutCtx, &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"codeName": {
				S: aws.String("theDoctor"),
			},
		},
		TableName: aws.String("devops-challenge"),
	})

	if err != nil {
		log.Printf("Got error calling GetItem: %s", err)
	}

	if result.Item == nil {
		log.Println("No Item Was found")
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	genericErrorHandler(err)

	return item, nil

}

func genericErrorHandler(err error) error {
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func getSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(true),
		Region:                        aws.String("us-east-1"),
		Endpoint:                      aws.String("http://local-dynamo-service:8000")})
	if err != nil {
		log.Println(err)
		return sess, err
	}
	return sess, nil
}
