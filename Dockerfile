FROM golang:1.12
MAINTAINER "lhlyu"
ADD main /go/main
ADD config.yaml /go/conf/config.yaml
RUN chmod 777 -R /go
ENV LANG en_US.UTF-8
CMD ["/bin/bash","-c","/go/main"]

# 创建镜像
# docker build -t iyu .
# 创建容器运行
# docker run -itd --name=iyu -p 8080:8080 iyu
