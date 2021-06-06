FROM golang:1.16.4
WORKDIR /app
COPY . .
RUN go build -o main main.go && rm main.go
CMD ./main
