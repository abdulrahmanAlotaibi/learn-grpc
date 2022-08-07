package routes

import (
	
	"fmt"
	"net/http"

	"github.com/abdulrahmanAlotaibi/learn-rpc/personsvc/pb"
	"github.com/gin-gonic/gin"
)

type CreatePersonRequestBody struct{
	Name string `form:"name"`
	Club string `form:"club"`
}

func CreatePerson(ctx *gin.Context, c pb.PersonServiceClient){
	body := CreatePersonRequestBody{}
	fmt.Print(body)
	name := ctx.PostForm("name")
	club := ctx.PostForm("club")
	if  &club == nil || &name == nil  {
		return
	}

	fmt.Println(name, club)

	// DB Create


	ctx.JSON(http.StatusCreated, "Order has been created")
}

