FROM alpine:3.16

RUN apk -U add bind-tools bind-dnssec-tools bind-dnssec-root && rm -rf /var/cache/apk/*

VOLUME /home/root
WORKDIR /home/root

ENTRYPOINT [ "/bin/sh" ]
