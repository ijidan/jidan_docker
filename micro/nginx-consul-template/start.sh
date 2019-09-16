#!/bin/bash

nohup nginx -g 'daemon off;' &

/bin/consul-template --version
/bin/consul-template \
	-consul-addr=consul:8500 \
	-template="/etc/nginx/conf.d/consul-templates.conf.ctmpl:/etc/nginx/conf.d/consul-templates.conf:nginx -s reload"
