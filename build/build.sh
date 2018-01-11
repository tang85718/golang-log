#!/usr/bin/env bash
#go build -o p-vip-gateway ../main.go
echo "构建 go 应用..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o logtest ../test.go
# 登录 build hub

echo "构建 build 本地镜像..."
# 构建 build 镜像并上传到 build hub
export VERSION=1.1
docker build -t logtest:${VERSION} .
docker tag logtest:${VERSION} 10.0.1.200:5000/logtest:${VERSION}
docker push 10.0.1.200:5000/logtest:${VERSION}

#build tag logtest:${VERSION} 10.0.1.200:5000/logtest:latest
#build push 10.0.1.200:5000/logtest:latest

echo "构建完成"