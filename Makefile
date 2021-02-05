init:
	go get github.com/golang/protobuf@v1.4.1
	go get github.com/golang/protobuf/protoc-gen-go@v1.3.0
	go get -u google.golang.org/grpc@v1.26.0

gen-grpc:
	protoc -I . $(file) --go_out=plugins=grpc:.