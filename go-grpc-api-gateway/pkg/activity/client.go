package auth

import (
	"fmt"

	// "github.com/YOUR_USERNAME/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/notaduck/go-grpc-api-gateway/pkg/activity/pb"
	"github.com/notaduck/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ActivityServiceClient
}

func InitServiceClient(c *config.Config) pb.ActivityServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AcivitySvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewActivityServiceClient(cc)
}
