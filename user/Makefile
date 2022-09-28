
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/micro/micro/v3/cmd/protoc-gen-micro@latest
	go install github.com/micro/micro/v3/cmd/protoc-gen-openapi@latest

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/user.proto

.PHONY: proto
proto:
#	protoc --proto_path=. --micro_out=. --go_out=:. proto/user.proto
	docker run --rm -v $(pwd):$(pwd) -w $(pwd) -e ICODE=4E4B49DC715827DA cap1573/cap-protoc -I ./ --go_out=./ --micro_out=./ ./proto/user/user.proto
.PHONY: build
build:
	go build -o user *.go

.PHONY: dockerBuild
dockerBuild:
	docker build -t user:latest .

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user:latest
