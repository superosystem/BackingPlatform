# BUILDER
FROM golang:1.20.0-alpine3.17 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

RUN make engine

# DISTRIBUTION
FROM alpine:3.17

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 8000

COPY --from=builder /app/engine /app
COPY --from=builder /app/migrater /app
COPY --from=builder /app/migrations /app/migrations

CMD ["/app/engine"]