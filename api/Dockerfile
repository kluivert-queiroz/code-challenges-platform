FROM docker:dind

RUN apk add --no-cache \
    go \
    bash

WORKDIR /app

COPY . .

RUN go build -o api cmd/api.go

CMD ["./api"]