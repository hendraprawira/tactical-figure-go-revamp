FROM golang:1.18.4-alpine


RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /tactical-figure

EXPOSE 8009
EXPOSE 8008

ENTRYPOINT ["/tactical-figure"]