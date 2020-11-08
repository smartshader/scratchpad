FROM alpine:3.12.1

RUN apk update && apk add flite

ADD app /app

ENTRYPOINT ["/app"]