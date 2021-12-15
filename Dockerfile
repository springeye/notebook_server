FROM golang:1.17.5-alpine3.15  as builder
RUN apk update && apk add --virtual build-dependencies build-base gcc wget git
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go build -a -tags netgo -ldflags '-w -s' -o server


FROM alpine:latest as prod
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder "/app/server" "/app/server"
VOLUME /app/config.yaml
WORKDIR /app
CMD ["/app/server"]