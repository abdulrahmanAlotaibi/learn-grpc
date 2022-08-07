package api

import (
	"github.com/gin-gonic/gin"
	"github.com/abdulrahmanAlotaibi/learn-rpc/personsvc"
	"github.com/abdulrahmanAlotaibi/learn-rpc/personsvc/routes"
)

func Init(){
	// Create http server
	r := gin.Default()

	// Run a rpc client that request data from rpc server at 50053
	orderSVC := &personsvc.ServiceClient{
		Client:personsvc.InitServiceClient("localhost:50053"),
	}

	// redirect /orders request to order microservice
	r.POST("/orders", func(ctx *gin.Context) {
		routes.CreatePerson(ctx, orderSVC.Client) // Redirect to RPC function/service
	})
	
	r.Run(":3000")
}