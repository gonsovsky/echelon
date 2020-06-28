package main

import (
	"context"
	"echelon/service"
	"fmt"
	"google.golang.org/grpc"
	"strconv"
)

func main() {
	serverAddress := "localhost:7000"

	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := service.NewFilterServiceClient(conn)

	for i := range [10]int{} {
		model := service.Record{
			F: map[string]string{
				"ValueA": strconv.Itoa(i),
				"ValueB": strconv.Itoa(i),
				"ValueC": strconv.Itoa(i),
			},
			S: &service.Input{
				ValueA: strconv.Itoa(i),
				ValueB: strconv.Itoa(i),
				ValueC: strconv.Itoa(i),
			},
		}

		if responseMessage, e := client.Add(context.Background(), &model); e != nil {
			panic(fmt.Sprintf("Error: %v", e))
		} else {
			fmt.Println(responseMessage)
		}
	}
}
