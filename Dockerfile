FROM golang:1.16.8-alpine3.14 AS builder
WORKDIR /workspace
ARG SERVICE

RUN apk add --no-cache gcc libc-dev curl

COPY . .
RUN mkdir -p /protobuf/google/protobuf && \
    for f in any duration descriptor empty struct timestamp wrappers; do \
    curl -L -o /protobuf/google/protobuf/${f}.proto https://raw.githubusercontent.com/google/protobuf/master/src/google/protobuf/${f}.proto; \
    done
RUN mkdir pb && for listProto in $(ls proto/); do protoc proto/$listProto/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=/protobuf -I=proto; done
RUN /usr/local/go/bin/go mod init ${SERVICE} && GOPRIVATE=github.com/leonardochristofer/*  /usr/local/go/bin/go mod tidy && mkdir /tmp/result && /usr/local/go/bin/go get github.com/gin-gonic/gin@v1.7.7
RUN CGO_ENABLED=0 GOOS=linux /usr/local/go/bin/go build -o /tmp/result/${SERVICE} main.go

FROM alpine:3.13
WORKDIR /app
ARG SERVICE
ARG PORT

RUN mkdir -p /app/log && touch /app/log/${SERVICE}.log
COPY --from=builder /tmp/result/* /app/

EXPOSE ${PORT}
