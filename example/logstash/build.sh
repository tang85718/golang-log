#!/usr/bin/env bash

# 构建 build 镜像并上传到 build hub
export VERSION=1.0
docker build -t logstash:${VERSION} .
docker tag logstash:${VERSION} 10.0.1.200:5000/logstash:${VERSION}
docker push 10.0.1.200:5000/logstash:${VERSION}

#build tag logtest:${VERSION} 10.0.1.200:5000/logtest:latest
#build push 10.0.1.200:5000/logtest:latest

echo "构建完成"