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

func listTable() {
	sess, err := getSession()
	genericErrorHandler(err)

	dbSvc := dynamodb.New(sess)

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tables, err := dbSvc.ListTablesWithContext(timeoutCtx, &dynamodb.ListTablesInput{})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Tables:")
	for _, table := range tables.TableNames {
		log.Println(*table)
	}
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
		log.Fatalf("Got error calling GetItem: %s", err)
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
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:8000")})
	if err != nil {
		log.Println(err)
		return sess, err
	}

	return sess, nil
}
