FROM alpine:3.22

LABEL org.opencontainers.image.source="https://github.com/dcjulian29/docker-bind"
LABEL org.opencontainers.image.description="A Docker Container and execution tool for BIND"

ARG VERSION

COPY docker-entrypoint.sh /docker-entrypoint.sh

RUN apk -U upgrade && \
  apk add bind=${VERSION} && \
  apk add bind-tools=${VERSION} && \
  apk add bind-dnssec-tools=${VERSION} && \
  apk add bind-dnssec-root=${VERSION} && \
  rm -rf /var/cache/apk/* && \
  rm -f /etc/bind/rndc.key && \
  chmod 755 /docker-entrypoint.sh

VOLUME /data

WORKDIR /data

ENTRYPOINT [ "/docker-entrypoint.sh" ]
