FROM alpine:3.18.3

RUN apk update &&\
    apk upgrade &&\
    apk add git

COPY ./build/go-diff /go-diff
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]