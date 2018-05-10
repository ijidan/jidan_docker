# docker-image
docker镜像整合的lnmp环境
<br/>
<br/>

#### 官方生成的镜像组合：nginx镜像+php镜像+mysql镜像+phpMyAdmin镜像 构成

	以下目录根据实际情况，进行修改，包括修改docker-compose.yml文件中的路径
	
	如果用户为 vagrant
	则docker数据目录为      /home/vagrant/docker
	nginx的配置文件目录     /home/vagrant/docker/nginx/conf.d
	nginx的默认配置文件     /home/vagrant/docker/nginx/conf.d/default.conf
	php的项目目录           /home/vagrant/www
	phpMyAdmin的地址        http://localhost:8080/
	mysql的数据保存目录     /home/vagrant/docker/mysql/data
	mysql的日志目录         /home/vagrant/docker/mysql/logs
	mysql的配置文件目录     /home/vagrant/docker/mysql/conf
	mysql账号、密码         root/123456

<br/>

1、创建目录：nginx的配置文件目录，php项目目录，mysql的数据目录、日志目录、配置文件目录
```bash
sudo mkdir -p /home/vagrant/docker/nginx/conf.d /home/vagrant/www /home/vagrant/docker/mysql/data /home/vagrant/docker/mysql/logs /home/vagrant/docker/mysql/conf
```  

2、将目录下的nginx的默认配置文件default.conf，复制到nginx的配置文件目录下
```bash
cp ./default.conf /home/vagrant/docker/nginx/conf.d/
```  

3、执行docker-compose命令，创建、构建镜像，并启动、运行容器
```bash
docker-compose up -d
```  
其他命令
```bash
docker-compose start    启动容器
docker-compose stop     停止容器
docker-compose ps       查看容器运行状态
docker-compose rm       删除容器
```