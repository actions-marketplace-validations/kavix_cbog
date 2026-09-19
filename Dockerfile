FROM alpine:3.20

RUN apk add --no-cache ca-certificates git

COPY cbog /usr/local/bin/cbog

ENTRYPOINT ["/usr/local/bin/cbog"]
