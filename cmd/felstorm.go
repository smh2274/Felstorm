package main

import (
	"fmt"
	jwt "github.com/smh2274/Felstorm/api"
	"github.com/smh2274/Felstorm/internal/logger"
	"github.com/smh2274/Felstorm/internal/services"
	"github.com/smh2274/Felstorm/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main() {
	// 捕获panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	viper, err := util.LoadConfig()
	if err != nil {
		log.Fatalf("load config fail: %v", err)
	}

	// init log
	err = logger.InitLogger(viper)
	if err != nil {
		log.Fatalf("init log fail: %v", err)
	}

	address := viper.GetString("server.address")
	port := viper.GetString("server.port")

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// tls
	opts := make([]grpc.ServerOption, 0)

	creds, err := credentials.NewServerTLSFromFile(
		viper.GetString("ssl.cert"),
		viper.GetString("ssl.key"))
	opts = append(opts, grpc.Creds(creds))

	server := grpc.NewServer(opts...)

	tokenServ := &services.GRPCTokenServer{
		V: viper,
	}

	jwt.RegisterTokenServer(server, tokenServ)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to start felstorm server: %v", err)
	} else {
		log.Printf("felstorm server start success: %s:%s", address, port)
	}
}
