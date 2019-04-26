FROM alpine:latest

ENV HOST "0.0.0.0:9999"

ADD main /
ADD dbfile.csv /

ENTRYPOINT /main -host ${HOST}