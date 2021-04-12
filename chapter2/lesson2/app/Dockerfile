FROM golang:1.16.3-alpine

RUN apk update && apk add --no-cache git gcc musl-dev

RUN mkdir -p /app/src
WORKDIR /app/src

COPY src/go.mod .
COPY src/go.sum .
RUN go mod download

COPY ./src .

CMD [ "go", "run", "./main.go" ]
