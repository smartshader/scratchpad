FROM golang:1.15.3-alpine AS build-env

# Allow Go to retrieve the dependencies for the build step
RUN apk add --no-cache git

# Secure against running as root
RUN adduser -D -u 10000 imad
RUN mkdir /gophercon/ && chown imad /gophercon/
USER imad

WORKDIR /gophercon/
ADD . /gophercon/

# Compile the binary, we don't want to run the cgo resolver
RUN CGO_ENABLED=0 go build -o /gophercon/gc .

# final stage
FROM alpine:3.12.1

# Secure against running as root
RUN adduser -D -u 10000 imad
USER imad

WORKDIR /
COPY --from=build-env /gophercon/certs/* /
COPY --from=build-env /gophercon/gc /

EXPOSE 8080

CMD ["gc"]
