FROM golang:1.16.3-alpine as Build
RUN apk add --no-cache git build-base
WORKDIR /app
COPY . /app/
RUN go build -mod=vendor -ldflags "-s -w" -o cobweb

FROM alpine:3.9.4
WORKDIR /app
COPY --from=Build /app/cobweb /app/cobweb
# 要加上这个 golang 才会处理额外添加的 dns 记录(如: --add-host)
RUN echo "hosts: files dns" > /etc/nsswitch.conf
VOLUME [ "/app/data" ]
ENV \
  DB="/app/data/cobweb.db"
CMD ["./cobweb"]
