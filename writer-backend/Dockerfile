# Building the binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

## caching dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . . 
RUN go build -v -o writer .

# Running the binary
FROM alpine:latest

WORKDIR /app

## copy the binary from the builder stage
COPY --from=builder /app/writer . 

RUN chmod +x writer

EXPOSE 8080

CMD ["./writer"]