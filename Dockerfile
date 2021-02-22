# Builder
FROM golang:1.15-alpine AS builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

RUN make engine

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app 

EXPOSE 9000

COPY --from=builder /app/engine /app

CMD /app/engine
