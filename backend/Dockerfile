FROM alpine:latest

ENV HOST "0.0.0.0:9999"

ADD main /

EXPOSE 9999

ENTRYPOINT /main -host ${HOST}