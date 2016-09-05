FROM alpine:3.4
MAINTAINER Mike Mackintosh <m@zyp.io>

RUN apk add --update --no-cache bash curl ca-certificates openssl \
        && rm -rf /var/cache/apk/*

RUN update-ca-certificates

COPY bin/kessha-amd64 /usr/sbin/kesshad

RUN mkdir -p /etc/kessha/
COPY contrib/id_* /etc/kessha/

COPY contrib/entrypoint.sh /entrypoint.sh
COPY .env /env
RUN chmod +x /entrypoint.sh

EXPOSE 2022

CMD ["/entrypoint.sh"]
