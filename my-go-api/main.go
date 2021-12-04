package main

import (
	"log"

	"exam6/mydb"
	"exam6/SQL/get"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	_ "github.com/go-sql-driver/mysql"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// まずDBと繋げる
	db, err := mydb.Connection()
	// ハンドリングは直後に
	if err != nil {
		log.Print("データベースのオープンに失敗")

		return events.APIGatewayProxyResponse {
			Body: "",
			StatusCode: 500,
		},err
	}

	get.SelectAll(db, request)

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
