FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o bin .

EXPOSE 8080

ENTRYPOINT ["/app/bin"]
