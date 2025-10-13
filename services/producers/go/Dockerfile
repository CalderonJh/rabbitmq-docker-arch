FROM golang:1.22-alpine
WORKDIR /app
RUN apk add --no-cache netcat-openbsd
COPY . .
RUN go mod init app && go get github.com/streadway/amqp && go build -o main .
CMD ["./main"]