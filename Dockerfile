FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $PWD/.build/godo-service
COPY . .

RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/godo

FROM scratch
COPY --from=builder /go/bin/godo /go/bin/godo

ENTRYPOINT ["/go/bin/godo"]