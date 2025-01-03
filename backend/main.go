package main

import "fmt"
import pb "github.com/unicom-jf/learning-protobuf/backend/learningpb"

func main() {
	pb.Person{
    Id:    1234,
    Name:  "John Doe",
    Email: "jdoe@example.com",
    Phones: []*pb.Person_PhoneNumber{
        {Number: "555-4321", Type: pb.PhoneType_PHONE_TYPE_HOME},
    },
	}
	fmt.Println("ok")
}