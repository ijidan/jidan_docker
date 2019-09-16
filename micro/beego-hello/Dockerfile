# beego-hello

# base image
#FROM alpine:latest
FROM centos:latest
#FROM busybox:latest

#alpine、busybox不支持curl命令
#centos不支持ifconfig命令

# maintainer info

#获取网络地址
#RUN yum install -y net-tools.x86_64

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# workspace
WORKDIR "/projects/bin"

# files
ADD beego-hello/beego-hello.tar.gz .
RUN chmod +x beego-hello

# port
# api
EXPOSE 80/tcp

# entrypoint
ENTRYPOINT ["/projects/bin/beego-hello"]
