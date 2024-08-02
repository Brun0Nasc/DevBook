# Dockerfile to local database (MySQL)

Here you find a guide to execute the mysql container on your local machine.

## Prerequisites

- Docker installed on your machine. For installation instructions, refer to the [official Docker documentation](https://docs.docker.com/get-docker/).

## Instructions

In the terminal, within the project directory, run the following command to build the Docker image:

```bash
docker build -t devbook-image .
```

After building the image, run a container from it:

```bash
docker run --name devbook -d -p 3306:3306 devbook-image
```

To access the MySQL client within the container, run:

```bash
docker exec -it devbook mysql -u root -p
```

You can also execute MySQL commands directly from the host without entering the container terminal. For example, to list all databases:

```bash
docker exec -it mysql-container mysql -u root -p -e "SHOW DATABASES;"
```
