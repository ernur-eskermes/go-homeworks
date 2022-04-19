package main

import (
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"test/proto/notification"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := notification.NewNotificationServiceClient(conn)

	response, err := c.Notify(context.Background(), &notification.NotificationRequest{Message: "Я на пути к становлению ниндзей!"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("STATUS:", response.Status)
}
