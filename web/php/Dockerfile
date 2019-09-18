FROM php:7.1-fpm

#COPY ./file/sources.list.stretch /etc/apt/sources.list
COPY ./file/libmemcached-1.0.18.tar.gz /usr/local/src/
COPY ./file/php7.zip /usr/local/src/
COPY ./file/phpredis-master.zip /usr/local/src/
COPY ./file/mongodb-1.1.9.tgz /usr/local/src/


# Extensions: ctype, dom, fileinfo, ftp, hash, iconv, json, pdo, pdo_sqlite, session,
# tokenizer, simplexml, xml, xmlreader, xmlwriter and phar are bundled and compiled into
# PHP by default. If missing, install them directly by `docker-php-ext-install extension_name`

# Notice:
# 1. Mcrypt was DEPRECATED in PHP 7.1.0, and REMOVED in PHP 7.2.0.
# 2. opcache requires PHP version >= 7.0.0.
# 3. soap requires libxml2-dev.
# 4. xml, xmlrpc, wddx require libxml2-dev and libxslt-dev.
# 5. Line `&& :\` is just for better reading and do nothing.
RUN apt-get update \
	&& apt-get install -y wget gcc make perl zip unzip openssl vim \
    && apt-get install -y libfreetype6-dev libjpeg62-turbo-dev libpng-dev libmemcached-dev zlib1g-dev \
    && docker-php-ext-configure gd --with-freetype-dir=/usr/include/ --with-jpeg-dir=/usr/include/ \
    && docker-php-ext-install gd \
#    && :\
#    && apt-get install -y libicu-dev \
#    && docker-php-ext-install intl \
#    && :\
#    && apt-get install -y libxml2-dev \
#    && apt-get install -y libxslt-dev \
#    && docker-php-ext-install soap \
#    && docker-php-ext-install xsl \
#    && docker-php-ext-install xmlrpc \
#    && docker-php-ext-install wddx \
#    && :\
#    && apt-get install -y libbz2-dev \
#    && docker-php-ext-install bz2 \
    && :\
    && docker-php-ext-install zip \
    && docker-php-ext-install pcntl \
    && docker-php-ext-install pdo \
    && docker-php-ext-install pdo_mysql \
    && docker-php-ext-install mysqli \
    && docker-php-ext-install mbstring
    #&& docker-php-ext-install exif \
    #&& docker-php-ext-install bcmath \
    #&& docker-php-ext-install calendar \
    #&& docker-php-ext-install sockets \
    #&& docker-php-ext-install gettext \
    #&& docker-php-ext-install shmop \
    #&& docker-php-ext-install sysvmsg \
    #&& docker-php-ext-install sysvsem \
    #&& docker-php-ext-install sysvshm \
    #&& docker-php-ext-install opcache \
    #&& docker-php-ext-install pdo_firebird \
    #&& docker-php-ext-install pdo_dblib \
    #&& docker-php-ext-install pdo_oci \
    #&& docker-php-ext-install pdo_odbc \
    #&& docker-php-ext-install pdo_pgsql \
    #&& docker-php-ext-install pgsql \
    #&& docker-php-ext-install oci8 \
    #&& docker-php-ext-install odbc \
    #&& docker-php-ext-install dba \
    #&& docker-php-ext-install interbase \
    #&& :\
	#&& apt-get install -y libmcrypt-dev \
	#&& docker-php-ext-install mcrypt \
    #&& :\
	#&& apt-get install -y curl \
	#&& apt-get install -y libcurl3 \
	#&& apt-get install -y libcurl4-openssl-dev \
	#&& docker-php-ext-install curl \
	#&& :\
	#&& apt-get install -y libreadline-dev \
	#&& docker-php-ext-install readline \
    #&& :\
    #&& apt-get install -y libsnmp-dev \
    #&& apt-get install -y snmp \
    #&& docker-php-ext-install snmp \
    #&& :\
    #&& apt-get install -y libpspell-dev \
    #&& apt-get install -y aspell-en \
    #&& docker-php-ext-install pspell \
    #&& :\
    #&& apt-get install -y librecode0 \
    #&& apt-get install -y librecode-dev \
    #&& docker-php-ext-install recode \
    #&& :\
    #&& apt-get install -y libtidy-dev \
    #&& docker-php-ext-install tidy \
    #&& :\
    #&& apt-get install -y libgmp-dev \
    #&& ln -s /usr/include/x86_64-linux-gnu/gmp.h /usr/include/gmp.h \
    #&& docker-php-ext-install gmp \
    #&& :\
    #&& apt-get install -y postgresql-client \
    #&& apt-get install -y mysql-client \
    #&& :\
    #&& apt-get install -y libc-client-dev \
    #&& docker-php-ext-configure imap --with-kerberos --with-imap-ssl \
    #&& docker-php-ext-install imap \
    #&& :\
    #&& apt-get install -y libldb-dev \
    #&& apt-get install -y libldap2-dev \
    #&& docker-php-ext-configure ldap --with-libdir=lib/x86_64-linux-gnu \
    #&& docker-php-ext-install ldap \

# Composer
#RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/bin/ --filename=composer \

# Install extension using pecl
# Notice: if pecl install get error
#    `No releases available for package "pecl.php.net/xxx"`
# or
#    `Package "xxx" does not have REST xml available`
# Please turn on proxy (The proxy IP may be docker host IP or others):
# RUN pear config-set http_proxy http://10.0.75.1:1080
#memcached
RUN pecl install memcached-3.0.4 \
    && docker-php-ext-enable memcached
#swoole
RUN pecl install swoole-4.4.5 \
    && docker-php-ext-enable swoole
#solr
RUN pecl install solr-2.5.0 \
    && docker-php-ext-enable solr


#source install
WORKDIR /usr/local/src
RUN unzip phpredis-master.zip \
    && rm phpredis-master.zip \
    && ( \
        cd phpredis-master \
        && phpize \
        && ./configure \
        && make -j$(nproc) \
        && make install \
    ) \
    && rm -rf phpredis-master \
    && docker-php-ext-enable redis

#RUN wget http://launchpad.net/libmemcached/1.0/1.0.18/+download/libmemcached-1.0.18.tar.gz \
#RUN tar zxvf libmemcached-1.0.18.tar.gz \
#	&& rm libmemcached-1.0.18.tar.gz \
#    && cd libmemcached-1.0.18 \
#    && ./configure --prefix=/usr/local/libmemcached \
#    && make -j$(nproc) \
#    && make install \
#    && rm -rf libmemcached-1.0.18
#
#


#
##RUN  wget http://github.com/php-memcached-dev/php-memcached/archive/php7.zip \
#RUN	 unzip php7.zip \
#	 && rm php7.zip \
#	 && (\
#	     cd php-memcached-php7 \
#	     && phpize \
#	     && ./configure --with-libmemcached-dir=/usr/local/libmemcached --disable-memcached-sasl \
#	     && make -j$(nproc) \
#	     && make install \
#	 )\
#	 && rm -rf php-memcached-php7 \
#	 && docker-php-ext-enable memcached

#RUN	 tar zxvf mongodb-1.1.9.tgz \
#	 && rm mongodb-1.1.9.tgz \
#	 && (\
#	     cd mongodb-1.1.9  \
#	     && phpize \
#	     && ./configure \
#	     && make -j$(nproc) \
#	     && make install \
#	 )\
#	 && rm -rf mongodb-1.1.9 \
#	 && docker-php-ext-enable mongodb

#Composer
COPY ./file/composer /usr/local/bin/
RUN cd /usr/local/bin \
	&& chmod +x composer

EXPOSE 9000

CMD ["php-fpm", "-F"]

