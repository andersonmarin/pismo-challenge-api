FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN apk add --update git && \
  go get && \
  go build -o api

CMD ["/app/api"]