FROM golang:1.18-alpine3.17 AS builder
WORKDIR /app
COPY . .
ENV GOPROXY=https://goproxy.cn
RUN go build -o main server/server.go

FROM alpine:3.17
#ENV GIN_MODE=release
ENV DbSource="root:root@tcp(compose-mysql:3306)/bank?parseTime=true&loc=Local"
WORKDIR /app
##不要忘了为alpine版本创建软连接
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .
COPY --from=builder /app/migrate .
COPY db/migration ./migration
COPY start.sh .
COPY wait-for-it.sh .
EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]