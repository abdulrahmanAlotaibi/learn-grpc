package routes 

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/abdulrahamnAlotaibi/learn-rpc/person-svc/pb"

)

type CreatePersonRequestBody struct{
	name string
	club string
}

func CreatePerson(ctx *gin.Context, c pb.){}