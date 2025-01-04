package main

import (
	"fmt"
	"log"
	pb "github.com/unicom-jf/learning-protobuf/backend/learningpb"
	"google.golang.org/protobuf/proto"
	"github.com/nats-io/nats.go"
)

func doSubscribe(natsConn *nats.Conn, exit chan int, ready chan int) {
	subj := "q.login.jf10"
	_, err := natsConn.Subscribe(subj, func(msg *nats.Msg) {
		book := &pb.AddressBook{}
		proto.Unmarshal(msg.Data, book)
		log.Printf("msg.data : %v, len:%d\n", book, len(msg.Data))
		exit <- 1
	})
	if err == nil {
		log.Printf("Listening on [%s]", subj)
		ready <- 1
	} else {
		log.Fatal(err)
	}
}
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
	exit := make(chan int)
	ready := make(chan int)
	go doSubscribe(natsConn, exit, ready)
	<-ready
	data, _ := proto.Marshal(book)
	natsConn.Publish("q.login.jf10", data)
	<-exit
	fmt.Println("done")
}