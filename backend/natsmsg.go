package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
	pb "github.com/unicom-jf/learning-protobuf/backend/learningpb"
	"google.golang.org/protobuf/proto"
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

func doSubjSaveBook(natsConn *nats.Conn) {
	subj := "savebookRequest"
	_, err := natsConn.Subscribe(subj, func(msg *nats.Msg) {
		book := &pb.AddressBook{}
		proto.Unmarshal(msg.Data, book)
		log.Printf("msg.data : %v, len:%d\n", book, len(msg.Data))
		saveBook(book)
		//exit <- 1
	})
	if err == nil {
		log.Printf("Listening on [%s]", subj)
		//ready <- 1
	} else {
		log.Fatal(err)
	}
}
func doSubjGetBook(natsConn *nats.Conn) {
	subj := "getbookRequest"
	_, err := natsConn.Subscribe(subj, func(msg *nats.Msg) {
		book := &pb.AddressBook{}
		//proto.Unmarshal(msg.Data, book)
		log.Printf("msg.data : %v, len:%d\n", book, len(msg.Data))
		book, err := getBook()
		if err != nil {
			log.Printf("err: %v\n", err)
			return
		}
		data, _ := proto.Marshal(book)
		natsConn.Publish("getBookResponse", data)
		//exit <- 1
	})
	if err == nil {
		log.Printf("Listening on [%s]", subj)

	} else {
		log.Fatal(err)
	}
}

func main() {
	/*
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
	*/
	natsConn, err := nats.Connect("localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	exit := make(chan int)
	//ready := make(chan int)
	go doSubjGetBook(natsConn)
	go doSubjSaveBook(natsConn)
	//go doSubscribe(natsConn, exit, ready)
	//<-ready
	//data, _ := proto.Marshal(book)
	//natsConn.Publish("q.login.jf10", data)
	<-exit
	fmt.Println("done")
}


func saveBook(book *pb.AddressBook) {
	fname := "address.book"

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := os.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

}

func getBook()(*pb.AddressBook, error) {
	book := &pb.AddressBook{}
	fname := "address.book"
	in, err := os.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			//fmt.Printf("%s: File not found.  Creating new file.\n", fname)
			return book, nil
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}
	err = proto.Unmarshal(in, book)
	return book, err
}