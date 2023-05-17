FROM golang:1.19-alpine as builder

WORKDIR /app

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
ENV EMAIL_FROM default@email_from
ENV EMAIL_TO "default@email_1 default2@email_2 default3@email_3"
ENV EMAIL_PASSWORD default

ADD . /app

RUN go build -o hamster-media

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/hamster-media /app/hamster-media

EXPOSE 8888

CMD ["/app/hamster-media"]
