package main

import (
	"echelon/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	netListener := getNetListener(7000)
	grpcServer := grpc.NewServer()

	filterServiceImpl := service.NewFilterServiceGrpcImpl()
	service.RegisterFilterServiceServer(grpcServer, filterServiceImpl)

	// start the server
	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
