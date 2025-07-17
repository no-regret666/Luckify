FROM golang:1.24

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
ENV TZ=Asia/Shanghai

# 安装 tzdata，并设置时区
RUN apt-get update && \
    apt-get install -y tzdata && \
    ln -snf /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo "${TZ}" > /etc/timezone

# 安装 modd 工具
RUN go install github.com/cortesi/modd/cmd/modd@latest

RUN chmod +x /go/data/server/*


WORKDIR /go

CMD ["modd"]