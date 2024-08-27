
FROM golang:1.22

WORKDIR /app
COPY . .
WORKDIR /app/cmd/api

RUN go build -o /app/myapp

EXPOSE 8080
CMD ["/app/myapp"]