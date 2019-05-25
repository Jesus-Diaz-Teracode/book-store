FROM golang:stretch AS build

WORKDIR /go/src/github.com/Jesus-Diaz-Teracode/book-store

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o book_store .



FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=build /go/src/github.com/Jesus-Diaz-Teracode/book-store/book_store .

CMD ["./book_store"]