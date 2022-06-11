


help:
	echo 'help message'


SRC_DIR=./protos
DST_DIR=./gen

build_proto:
	#	protoc -I=$(SRC_DIR) --go_out=$(DST_DIR) *.proto
#	protoc -I=./protos --go_out=./gen test.proto
#	protoc -I=./protos --go_out=./gen tutorial.person.proto
#
#	protoc -I=./protos --go_out=./gen info.proto

	protoc --go_out=. protos/*.proto
	#protoc --go_out=plugins=grpc:./gen protos/*.proto
	protoc --go-grpc_out=. protos/*.proto

run:
	go run main.go



build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./development/bin/server cmd/server/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./development/bin/client cmd/client/main.go


