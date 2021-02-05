# --- 构建编译环境 --
FROM golang:1.15 AS builder

# 设置环境变量
ENV GOPROXY https://goproxy.cn,direct
ENV GOPRIVATE github.com/smh2274/

WORKDIR /go/src/github.com/smh2274/Felstorm

# 拷贝需要编译的文件
COPY . /go/src/github.com/smh2274/Felstorm

# 设置git的url，使其可以访问smh2274的私有仓库
RUN git config --global url."https://058b9a7bee77a0e72d4d10dd270c21f03ba4d5ae:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# 编译
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOROOT_FINAL=$(pwd) go build -a -ldflags '-w -extldflags "-static"'  -gcflags=-trimpath=$(pwd) -asmflags=-trimpath=$(pwd) cmd/felstorm.go

# --- 构建运行环境 ---
FROM envoyproxy/envoy-alpine:v1.17.0 AS prod

RUN mkdir -p /Azeroth/Felstorm/config \
    && mkdir -p /Azeroth/Felstorm/ssl \
    && mkdir -p /Azeroth/Felstorm/log \
    && touch /Azeroth/Felstorm/log/envoy_access.log \
    && chmod 777 /Azeroth/Felstorm/log/envoy_access.log

# 拷贝编译环境的二进制文件
COPY --from=builder /go/src/github.com/smh2274/Felstorm/felstorm /Azeroth/Felstorm/felstorm
COPY docker_prepare/* /Azeroth/Felstorm/
COPY internal/ssl/* /Azeroth/Felstorm/

RUN mv /Azeroth/Felstorm/envoy.yaml /Azeroth/Felstorm/config/ \
   && mv /Azeroth/Felstorm/felstorm_conf.yaml /Azeroth/Felstorm/config/ \
   && mv /Azeroth/Felstorm/domain.* /Azeroth/Felstorm/ssl/

# 设置时区
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
 && apk add --no-cache tzdata \
 && echo "Asia/Shanghai" > /etc/timezone \
 && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

EXPOSE 8800

CMD /Azeroth/Felstorm/run.sh