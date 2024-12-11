FROM golang:1.22-alpine

RUN mkdir -p /app

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./moniteur

EXPOSE 443

CMD ["./moniteur"]