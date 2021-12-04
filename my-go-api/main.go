package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	// ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	// ErrNon200Response = errors.New("Non 200 Response found")
)

type Task struct {
	Id int
	Content string
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// まずDBと繋げる
	db, err := mydb.Connection()

	if err != nil {
		log.Print("データベースのオープンに失敗")

		// エラー処理
		return events.APIGatewayProxyResponse {
			Body: "",
			StatusCode: 500,
		},err
	}

	defer db.Close()

	// デフォルトの返信？
	return events.APIGatewayProxyResponse{
		Body:       "OK",
		StatusCode: 200,
	}, nil
}



func main() {
	lambda.Start(handler)
}
