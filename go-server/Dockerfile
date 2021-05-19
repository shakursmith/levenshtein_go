FROM golang:1.12.0-alpine3.9

RUN mkdir /app

RUN apk add git

ADD . /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]