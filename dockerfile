FROM alpine
COPY ./bin /app/bin
COPY ./log /app/log
COPY ./config /app/config
COPY ./config/config.yaml /app/config/config.yaml

RUN apk add --no-cache tzdata
ENV TZ Asia/Taipei

WORKDIR /app
EXPOSE 9300
ENTRYPOINT [ "./bin/main" ]