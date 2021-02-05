package test

import (
	"context"
	jwt "github.com/smh2274/Felstorm/internal/api"
	"github.com/smh2274/Felstorm/internal/services"
	"github.com/smh2274/Felstorm/internal/util"
	"github.com/spf13/viper"
	"log"
	"testing"
	"time"
)

var v *viper.Viper

func init() {
	////初始化证书
	//creds, err := credentials.NewClientTLSFromFile("../ssl/domain.crt", "domain.com")
	//if err != nil {
	//	log.Fatalf("failed to load credentials: %v", err)
	//}
	//
	//// setup call options
	//opts := []grpc.DialOption{
	//	//grpc.WithTransportCredentials(creds),
	//	grpc.WithInsecure(),
	//}
	//opts = append(opts, grpc.WithBlock())
	//
	//ctx, _ := context.WithTimeout(context.Background(), time.Second*15)
	//
	//connect, err := grpc.DialContext(ctx, "127.0.0.1:8800", grpc.WithTransportCredentials(creds))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if connect == nil {
	//	log.Fatal("connect init nil")
	//}
	//
	//tokenClient = jwt.NewTokenClient(connect)

	viper, err := util.LoadConfig()
	if err != nil {
		log.Print(err)
	}

	v = viper
}

func TestFelStormServer(t *testing.T) {
	server := new(services.GRPCTokenServer)
	server.V = v

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	res, err := server.GetToken(ctx, &jwt.GetTokenRequest{
		Audience: "test",
		Exp:      int64(time.Minute),
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("get token: %s", res.GetToken())
	}
}
