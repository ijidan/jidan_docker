#!/bin/bash

docker run -d \
	-v /vagrant/docker-micro-services/nginx/www:/usr/share/nginx/html \
	-v /vagrant/docker-micro-services/nginx/conf/nginx.conf:/etc/nginx/nginx.conf \
	-v /vagrant/docker-micro-services/nginx/conf/conf.d:/etc/nginx/conf.d  \
	-v /vagrant/docker-micro-services/nginx/log:/var/log/nginx \
	nginx-consul-template-image
