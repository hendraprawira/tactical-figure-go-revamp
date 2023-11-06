FROM golang:1.18.4-alpine


RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /tactical-figure-1

EXPOSE 8008
EXPOSE 8009

ENTRYPOINT ["/tactical-figure-1"]