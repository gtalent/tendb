FROM alpine:latest

ADD 10db /usr/local/bin/10db

CMD 10db serve
