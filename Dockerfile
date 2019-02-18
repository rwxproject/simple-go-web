FROM alpine:3.9

ARG user=rwx
ARG group=rwx
ARG uid=1000
ARG gid=1000

WORKDIR /app

COPY ./dist/* /app

RUN addgroup -g ${gid} ${group} \
    && adduser -u ${uid} -G ${group} -s /bin/sh -D ${user} \
    && chown -R ${user}:${group} /app

USER ${user}

CMD [ "./main" ]