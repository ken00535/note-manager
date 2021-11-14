FROM alpine
COPY ./bin /app/bin
COPY ./log /app/log
COPY ./config /app/config
COPY ./config/config.aws.yaml /app/config/config.yaml
COPY ./dist /app/dist
COPY ./certs /app/certs

RUN apk add --no-cache tzdata
ENV TZ Asia/Taipei

WORKDIR /app
EXPOSE 9300
ENTRYPOINT [ "./bin/main" ]