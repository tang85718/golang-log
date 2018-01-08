#!/usr/bin/env bash

echo "构建 docker 本地镜像..."
# 构建 docker 镜像并上传到 docker hub
export VERSION=1.9
docker build --force-rm -t confd:${VERSION} .
docker tag confd:${VERSION} 10.0.1.200:5000/confd:${VERSION}
docker push 10.0.1.200:5000/confd:${VERSION}

docker tag confd:${VERSION} 10.0.1.200:5000/confd:latest
docker push 10.0.1.200:5000/confd:latest

echo "构建完成"