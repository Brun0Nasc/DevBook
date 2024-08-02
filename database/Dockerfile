FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=devbook
ENV MYSQL_USER=golang
ENV MYSQL_PASSWORD=dev123456

EXPOSE 3306

COPY ./init.sql /docker-entrypoint-initdb.d/
