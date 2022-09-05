ARG APP_NAME

FROM   golang:1.19.0-alpine3.16 AS go

RUN mkdir -p /home/app/

COPY $APP_NAME /home/app

CMD [$APP_NAME, "serve"]

