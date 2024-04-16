# 使用官方的Go基础镜像作为构建环境
FROM golang:latest AS builder

# 设置工作目录
WORKDIR /app

# 将当前目录下的所有源代码复制到容器内的/app目录
COPY . .

# 设置环境变量（这里使用您提供的环境变量）
ENV APP_MODE=debug
ENV HTTP_PORT=:3000
ENV JWT_KEY=89js82js72
ENV GOPROXY=https://goproxy.cn
ENV MYSQL_DBIP=127.0.0.1
ENV MYSQL_DBPORT=3306
ENV MYSQL_DBUSER=root
ENV MYSQL_DBPASSWORD=123456
ENV MYSQL_DBNAME=test
ENV REDIS_RDSIP=127.0.0.1
ENV REDIS_RDSPORT=6379
ENV REDIS_RDSNO=0
ENV REDIS_RDSUSER=""
ENV REDIS_RDSPASSWORD=
ENV MYSQL_ROOT_PASSWORD=${MYSQL_DBPASSWORD}
# 运行go mod下载依赖
RUN go mod tidy

# 设置时区和语言环境
ENV TZ=Asia/Shanghai
ENV LANG C.UTF-8

# 安装必要的依赖（如：tzdata以支持时区设置）
RUN apt-key adv --recv-keys --keyserver keyserver.ubuntu.com 3B4FE6ACC0B21F32

ADD sources.list /etc/apt/

RUN apt-get update && apt-get install -y tzdata

# 安装MySQL和Redis
RUN apt-get install -y mysql-server redis-server
# RUN service mysql start
# RUN echo 'update mysql.user set plugin="mysql_native_password" where User="root";update mysql.user set authentication_string=password('${MYSQL_ROOT_PASSWORD}') where User="root" and Host = "localhost";flush privileges;'|mysql  -u root --password=""
# RUN service mysql start
# RUN mysql -h${MYSQL_DBIP} -P${MYSQL_DBPORT} -u${MYSQL_DBUSER} -p${MYSQL_DBPASSWORD} -e \
#     "UPDATE mysql.user SET authentication_string=PASSWORD('${MYSQL_DBPASSWORD}'), plugin='mysql_native_password' WHERE User='root';"
# # 初始化MySQL（设置root用户密码等）
# RUN bash -c 'debconf-set-selections <<< "mysql-server mysql-server/root_password password 123456"' && \
#     bash -c 'debconf-set-selections <<< "mysql-server mysql-server/root_password_again password 123456"' 

# RUN service mysql stop
# 启动Redis并设置其配置（如设置密码）

# 添加启动脚本
COPY start.sh /start.sh
RUN chmod +x /start.sh
# 修改CMD指令
CMD ["/start.sh"]

