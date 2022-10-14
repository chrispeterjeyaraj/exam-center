FROM mongo:4
COPY /db/users.json /users.json
CMD mongoimport --host mongo --db examcenter --collection users --drop --file /users.json --jsonArray

FROM golang:1.19.2 AS development
WORKDIR /examcenterserver
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex@latest
EXPOSE 4000
CMD reflex -g '*.go' go run main.go --start-service

FROM golang:1.19.2 AS builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /examcenterserver
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o examcenterserver

FROM alpine:latest AS production
RUN apk add --no-cache ca-certificates
COPY --from=builder examcenterserver .
EXPOSE 4000
CMD ./examcenterserver



