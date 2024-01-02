FROM golang:1.21.5-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY . .

RUN go build -o menuupdationsystembackend

EXPOSE 8080

CMD ["./menuupdationsystembackend"]
