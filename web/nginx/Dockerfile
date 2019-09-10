FROM nginx:1.12.2

#保证服务器与容器有相同的时间
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

#备份配置文件
RUN cp /etc/nginx/nginx.conf /etc/nginx/nginx.conf.back
RUN cp -r /etc/nginx/conf.d /etc/nginx/conf.d.back

#导出http/https端口
EXPOSE 80 443

#启动nginx
CMD ["nginx", "-g", "daemon off;"]