# learning-protobuf

typescript talks to go with protobuf message

Frontend provides UI to maintain and list addressbook's records.
Backend stores the records in a file.
Frontend talks to backend using protobuf message with websocket.

Running in backend directory:
protoc -I=../proto --go_out=./learningpb ../proto/addressbook.proto
protoc --proto_path=../proto --go_out=./learningpb ../proto/addressbook.proto
