FROM golang:1.19.3-alpine

WORKDIR /app
COPY . .

RUN go get ./...
RUN go mod tidy
RUN apk add build-base

ENV URL ""

CMD ["go", "run", "main.go"]