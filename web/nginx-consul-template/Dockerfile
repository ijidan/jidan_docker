FROM nginx:latest
#MAINTAINER 

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# add consul-template
ADD consul-template_0.20.0_linux_amd64.tgz /bin/
RUN chmod a+x /bin/consul-template

# start shell
ADD start.sh /bin/start.sh
RUN chmod a+x /bin/start.sh

CMD ["/bin/start.sh"]
#CMD nohup sh -c "nginx -g 'daemon off;' && /bin/consul-template --version && /bin/consul-template -consul-addr=$CONSUL_ADDR -template='/etc/nginx/conf.d/consul-templates.conf.ctmpl:/etc/nginx/conf.d/consul-templates.conf:nginx -s reload'"

#build
#docker build -t nginx-consul-template-image:latest .


#run centos
#docker run -d -p 8100:80 -p 8443:443 \
#	-v /home/nginx/www:/usr/share/nginx/html \
#	-v /home/nginx/conf/nginx.conf:/etc/nginx/nginx.conf \
#	-v /home/nginx/conf/conf.d:/etc/nginx/conf.d \
#	-v /home/nginx/log:/var/log/nginx \
#	--link consul:consul \
#	--name nginx-consul-template-docker nginx-consul-template-image:latest


#run mac
#docker run -d -p 8100:80 -p 8443:443 \
#	-v /Users/mac/docker-micro-services/nginx/www:/usr/share/nginx/html \
#	-v /Users/mac/docker-micro-services/nginx/conf/nginx.conf:/etc/nginx/nginx.conf \
#	-v /Users/mac/docker-micro-services/nginx/conf/conf.d:/etc/nginx/conf.d \
#	-v /Users/mac/docker-micro-services/nginx/log:/var/log/nginx \
#	--link consul:consul \
#	--name nginx-consul-template-docker nginx-consul-template-image:latest

