FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 
RUN go build -v -o writer .
RUN chmod +x writer

EXPOSE 8080

CMD ["./writer"]
