#!/bin/bash

# Start MySQL
service mysql start
echo 'update mysql.user set plugin="mysql_native_password" where User="root";update mysql.user set authentication_string=password('${MYSQL_ROOT_PASSWORD}') where User="root" and Host = "localhost";flush privileges;'|mysql  -u root --password=""


# Start Redis
service redis-server start

# Start your Go application

go run main.go

