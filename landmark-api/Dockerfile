# Building the binary of the App
FROM golang:latest AS builder
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/landmark-api /app/cmd/landmark-api/main.go

# Moving the binary to the 'final Image' to make it smaller
FROM alpine
WORKDIR /app
COPY --from=builder bin/landmark-api landmark-api

CMD ["./landmark-api"]

EXPOSE 80
