#!/usr/bin/env bash

# 构建 docker 镜像并上传到 docker hub
export VERSION=1.0
docker build -t logstash:${VERSION} .
docker tag logstash:${VERSION} 10.0.1.200:5000/logstash:${VERSION}
docker push 10.0.1.200:5000/logstash:${VERSION}

#docker tag logtest:${VERSION} 10.0.1.200:5000/logtest:latest
#docker push 10.0.1.200:5000/logtest:latest

echo "构建完成"