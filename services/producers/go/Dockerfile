FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go mod init app && go get github.com/streadway/amqp && go build -o main .
CMD ["./main"]