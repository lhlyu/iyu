FROM golang:1.12
MAINTAINER "lhlyu"
ADD iyu /go/iyu
ADD config.yaml /go/conf/config.yaml
RUN chmod 777 -R /go
ENV LANG en_US.UTF-8
CMD ["/bin/bash","-c","/go/iyu"]

# 创建镜像
# docker build -t iyu .
# 创建容器运行
# docker run -itd --network iyu-bridge -p 9876:8080 iyu
