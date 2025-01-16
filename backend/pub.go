package main

import (
	"fmt"
	"log"
	pb "github.com/unicom-jf/learning-protobuf/backend/learningpb"
	"google.golang.org/protobuf/proto"
	"github.com/nats-io/nats.go"
)

func main() {
	book := &pb.AddressBook{}
	person := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
				{Number: "555-4321", Type: pb.PhoneType_PHONE_TYPE_HOME},
		},
	}
	book.People = append(book.People, &person)

	natsConn, err := nats.Connect("localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	data, _ := proto.Marshal(book)
	fmt.Printf("data: %v\n", data)
	//
	natsConn.Publish("hello", data)
	natsConn.Publish("hello", []byte("hello from go"))
	fmt.Println("done")
}