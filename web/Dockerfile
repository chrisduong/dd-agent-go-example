# NOTE: this is not image for Production. For Production, we should copy the binary file only
FROM golang:1.13-alpine

WORKDIR /app

RUN apk add --no-cache ca-certificates

ADD main.go .

# Build the Go app
RUN go build -o helloworld .

EXPOSE 8080

ENTRYPOINT ["./helloworld"]
