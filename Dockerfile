FROM alpine:3.18.3

RUN apk update &&\
    apk upgrade &&\
    apk add git jq

COPY ./build/go-diff /usr/bin/go-diff
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]