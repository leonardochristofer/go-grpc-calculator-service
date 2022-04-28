# Golang Stuff
GOCMD=go

GORUN=$(GOCMD) run

ENV=local

GOPRIVATE=github.com/aethereal333/*

SERVICE=go-grpc-calculator-service

init:
	$(GOCMD) mod init $(SERVICE)

reinit:
	rm go.mod go.sum && $(GOCMD) mod init $(SERVICE)

tidy:
	ENV=local GOPRIVATE=$(GOPRIVATE) $(GOCMD) mod tidy

run:
	ENV=$(ENV) $(GORUN) main.go

# Proto Generator
proto-gen:
	protoc proto/*/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=proto --experimental_allow_proto3_optional