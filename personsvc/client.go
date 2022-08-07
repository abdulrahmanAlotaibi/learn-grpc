package personsvc

import (
	"fmt"

	"github.com/abdulrahmanAlotaibi/learn-rpc/personsvc/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.PersonServiceClient
}

func InitServiceClient (port string) pb.PersonServiceClient{
	cc,err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect: ", err)
	}


	return pb.NewPersonServiceClient(cc)
}