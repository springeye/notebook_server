FROM golang:1.17.2
ADD . /app
WORKDIR /app

RUN go mod tidy

RUN go build -ldflags '-w -s' -o server
CMD ["/app/server"]