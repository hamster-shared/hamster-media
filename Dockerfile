FROM golang:1.19-alpine as builder

WORKDIR /app

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
ENV EMAIL_FROM hamster@hamsternet.io
ENV EMAIL_TO 335247945@qq.com

ADD . /app

RUN go build -o hamster-media


FROM alpine:3

WORKDIR /app

COPY --from=builder /app/hamster-media /app/hamster-media

EXPOSE 8888

CMD ["/app/hamster-media"]
