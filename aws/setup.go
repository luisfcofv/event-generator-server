package db

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Config struct {
	Table  string // required
	Region string // optional
	Limit  int64  // optional
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

var DynamodbClient *dynamodb.DynamoDB
var AppConfig Config

func Setup() {
	fmt.Println("Setting up AWS session...")
	AppConfig = Config{"World", "eu-west-2", 0}

	awscfg := &aws.Config{}
	awscfg.WithRegion(AppConfig.Region)

	// Create the session that the DynamoDB service will use.
	session := session.Must(session.NewSession(awscfg))

	// Create the DynamoDB service client to make the query request with.
	DynamodbClient = dynamodb.New(session)
}
