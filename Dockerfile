FROM golang:1.13-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates

ADD . .
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o helloworld main.go


###
FROM alpine:3.9

RUN apk --no-cache add ca-certificates
COPY --from=builder /app /app

EXPOSE 8080

ENTRYPOINT ["./app/helloworld"]
