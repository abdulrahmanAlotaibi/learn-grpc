package routes

import (
	
	"fmt"
	"net/http"

	"github.com/abdulrahmanAlotaibi/learn-rpc/personsvc/pb"
	"github.com/gin-gonic/gin"
)

type CreatePersonRequestBody struct{
	Name string `json:"name"`
	Club string `json:"club"`
}

func CreatePerson(ctx *gin.Context, c pb.PersonServiceClient){
	body := CreatePersonRequestBody{}

	if  err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	name , _ := ctx.Get("name")
	
	club , _ :=  ctx.Get("club")
	fmt.Println(name, club)

	// DB Create


	ctx.JSON(http.StatusCreated, "Order has been created")
}

