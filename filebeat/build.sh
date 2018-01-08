#!/usr/bin/env bash

# 构建 docker 镜像并上传到 docker hub
export VERSION=1.0
docker build -t filebeat:${VERSION} .

docker tag filebeat:${VERSION} 10.0.1.200:5000/filebeat:${VERSION}
docker tag filebeat:${VERSION} 10.0.1.200:5000/filebeat:latest

docker push 10.0.1.200:5000/filebeat:${VERSION}
docker push 10.0.1.200:5000/filebeat:latest

echo "构建完成"