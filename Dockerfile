FROM alpine:3.4
MAINTAINER Mike Mackintosh <m@zyp.io>


COPY bin/kessha-amd64 /usr/bin/kesshad
COPY contrib/entrypoint.sh /entrypoint.sh

RUN mkdir -p /etc/kessha/
COPY contrib/id_* /etc/kessha/
RUN chmod +x /entrypoint.sh

CMD ["/entrypoint.sh"]
