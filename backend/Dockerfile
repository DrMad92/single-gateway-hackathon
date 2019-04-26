FROM alpine:latest

ENV DBFILE "dbfile.csv"
ENV HOST "0.0.0.0:9999"

ADD main /
ADD dbfile.csv /

ENTRYPOINT /main -dbfile ${DBFILE} -host ${HOST}