FROM golang:alpine AS builder
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update && apk --no-cache add build-base
RUN apk update && apk --no-cache add build-base
WORKDIR /go/src/github.com/zjyl1994/livetv/
COPY . . 
#RUN GOPROXY="https://goproxy.io" GO111MODULE=on go build -o livetv .
RUN GO111MODULE=on go build -o livetv .

FROM python:3.9-alpine
RUN apk --no-cache add ca-certificates tzdata libc6-compat libgcc libstdc++
RUN wget https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -O /usr/local/bin/yt-dlp && chmod a+rx /usr/local/bin/yt-dlp
WORKDIR /root
COPY --from=builder /go/src/github.com/zjyl1994/livetv/view ./view
COPY --from=builder /go/src/github.com/zjyl1994/livetv/assert ./assert
COPY --from=builder /go/src/github.com/zjyl1994/livetv/.env .
COPY --from=builder /go/src/github.com/zjyl1994/livetv/livetv .
EXPOSE 9000
VOLUME ["/root/data"]
CMD ["./livetv"]