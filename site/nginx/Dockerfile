FROM centos:7

ENV NGINX_VERSION 1.10.3
ENV OPENSSL_VERSION 1.0.2
ENV PCRE_VERSION 8.40
ENV ZLIB_VERSION 1.2.11
ENV BUILD_ROOT /usr/local/src/nginx

RUN yum -y update \
    && yum list \
    && yum -y install vim lrzsz curl wget zip unzip make perl zlib zlib-devel gcc gcc-c++ libtool  openssl openssl-devel pcre pcre-devel libtermcap-devel ncurses-devel libevent-devel readline-devel

RUN mkdir -p $BUILD_ROOT \
	&& cd $BUILD_ROOT \
	&& curl https://nginx.org/download/nginx-$NGINX_VERSION.tar.gz -o $BUILD_ROOT/nginx-$NGINX_VERSION.tar.gz \
	&& tar vxzf nginx-$NGINX_VERSION.tar.gz \
	&& cd nginx-$NGINX_VERSION \
	&& BUILD_CONFIG="\
	    --prefix=/usr/local/nginx \
	    --with-http_stub_status_module \
	    --with-pcre \
	    --with-http_ssl_module \
	    --with-http_v2_module \
	    --with-threads \
	    " \
	&& ./configure $BUILD_CONFIG \
	&& make && make install \
	&& rm -rf $BUILD_ROOT

EXPOSE 80 443
CMD ["/usr/local/nginx/sbin/nginx", "-g", "daemon off;"]