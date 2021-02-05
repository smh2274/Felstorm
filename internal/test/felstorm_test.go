package test

import (
	"context"
	jwt "github.com/smh2274/Felstorm/internal/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"testing"
	"time"
)

var tokenClient jwt.TokenClient

func init() {
	//初始化证书
	creds, err := credentials.NewClientTLSFromFile("../ssl/domain.crt", "domain.com")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	// setup call options
	opts := []grpc.DialOption{
		//grpc.WithTransportCredentials(creds),
		grpc.WithInsecure(),
	}
	opts = append(opts, grpc.WithBlock())

	ctx, _ := context.WithTimeout(context.Background(), time.Second*15)

	connect, err := grpc.DialContext(ctx, "127.0.0.1:8800", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	if connect == nil {
		log.Fatal("connect init nil")
	}

	tokenClient = jwt.NewTokenClient(connect)
}

func TestFelStormServer(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*15)

	token, err := tokenClient.GetToken(ctx , &jwt.GetTokenRequest{
		Audience:             "envoy_test",
		Exp:                  int64(time.Minute),
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(token)
	}

}
