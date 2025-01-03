package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/unicom-jf/learning-protobuf/backend/learningpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	book := &pb.AddressBook{}
	fname := "address.book"
	in, err := os.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found.  Creating new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
		person := pb.Person{
			Id:    1234,
			Name:  "John Doe",
			Email: "jdoe@example.com",
			Phones: []*pb.Person_PhoneNumber{
					{Number: "555-4321", Type: pb.PhoneType_PHONE_TYPE_HOME},
			},
		}
		book.People = append(book.People, &person)
	} else {
		if err := proto.Unmarshal(in, book); err != nil {
			log.Fatalln("Failed to parse address book:", err)
		}
		if len(book.People) == 2 {
			fmt.Println("done with 2.")
			return
		}
		person := pb.Person{
			Id:    12342,
			Name:  "John Doe2",
			Email: "jdoe2@example.com",
			Phones: []*pb.Person_PhoneNumber{
					{Number: "555-43212", Type: pb.PhoneType_PHONE_TYPE_HOME},
			},
		}
		book.People = append(book.People, &person)
	}
// Write the new address book back to disk.
out, err := proto.Marshal(book)
if err != nil {
	log.Fatalln("Failed to encode address book:", err)
}
if err := os.WriteFile(fname, out, 0644); err != nil {
	log.Fatalln("Failed to write address book:", err)
}
	//fmt.Printf("ok: %v\n", person)
	fmt.Println("done")
}