FROM centos:7

RUN yum -y install gcc gcc-c++ make perl zip unzip wget vim java \
	&& cd /usr/local/src \
	&& wget https://mirrors.tuna.tsinghua.edu.cn/apache/lucene/solr/6.6.3/solr-6.6.3.tgz \
	&& tar -zxvf solr-6.6.3.tgz \
	&& mv solr-6.6.3 /usr/local/
#EXPOSE 8983
CMD ["/usr/local/solr-6.6.3/bin/solr","start","-force"]