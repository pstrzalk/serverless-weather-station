package main

import (
	"net/http"

	"github.com/pstrzalk/serverless-weather-station/weatherstation"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var initialized = false
var ginLambda *ginadapter.GinLambda

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if !initialized {
		ginEngine := weatherstation.MountAuthorizedRoute("/sensors/humidity", "get", processRequest)
		ginLambda = ginadapter.New(ginEngine)
		initialized = true
	}
	return ginLambda.Proxy(req)
}

func processRequest(c *gin.Context) {
	state := weatherstation.HumiditySensorState()
	c.JSON(http.StatusCreated, state)
}

func main() {
	lambda.Start(Handler)
}
