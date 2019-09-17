#!/bin/bash

nohup nginx -g 'daemon off;' &

/bin/consul-template --version
/bin/consul-template \
	-consul-addr=consul_server:8500 \
	-template="/etc/nginx/conf.d/consul-templates.http.ctmpl:/etc/nginx/conf.d/service_http.conf:nginx -s reload"
