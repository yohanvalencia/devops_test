FROM golang:1.16.4
WORKDIR /app
COPY . .
RUN go build -o main server.go && rm server.go
CMD ./main
