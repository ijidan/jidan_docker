FROM nginx:1.12.2

#配置文件
COPY ./file/consul-templates.ctmpl /usr/local/src

#保证服务器与容器有相同的时间
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

#备份配置文件
RUN cp /etc/nginx/nginx.conf /etc/nginx/nginx.conf.back
RUN cp -r /etc/nginx/conf.d /etc/nginx/conf.d.back



ADD ./file/consul-template_0.20.0_linux_amd64.tgz /usr/local/bin/
RUN chmod a+x /usr/local/bin/consul-template


RUN ls /usr/local/src

#导出http/https端口
EXPOSE 80 443

#启动nginx
CMD ["nginx", "-g", "daemon off;"]
#

RUN /usr/local/bin/consul-template --version
RUN /usr/local/bin/consul-template \
	-consul-addr=consul:8500 \
	-template="/usr/local/src/consul-templates.ctmpl:/etc/nginx/conf.d/www.test.com.conf:nginx -s reload"