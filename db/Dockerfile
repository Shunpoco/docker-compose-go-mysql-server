FROM mysql:8.0

ENV MYSQL_DATABASE="app"
ENV MYSQL_USER="username"
ENV MYSQL_PASSWORD="password"
ENV MYSQL_ROOT_PASSWORD="rootpassword"

ADD init.sql /docker-entrypoint-initdb.d/
