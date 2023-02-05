package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "rr.com/go-usermgmt-grpc/usermgmt"
)

const (
	address = "localhost:50051"
)

func main(){
	conn,err := grpc.Dial(address,grpc.WithTimeout(5*time.Second), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var new_users = make(map[string]int32)

	new_users["Alice"] = 44
	new_users["Bob"] = 66

	for name, age :=range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf(`user Details:
Name: %s
Age: %d
Id: %d`, r.GetName(), r.GetAge(), r.GetId())
	}

	params := &pb.GetUsersParams{}
	r, err := c.GetUsers(ctx, params)
	if err != nil{
		log.Fatalf("could not retrieve users: %v", err)
	}
	log.Print("\nUSER LIST: \n" )
	fmt.Printf("r.GetUsers(): %v\n", r.GetUsers())

}