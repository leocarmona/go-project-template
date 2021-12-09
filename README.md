# go-project-template

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/leocarmona/go-project-template)

**TL;DR:**
```shell
make up
```

## Prerequisites
[![Docker](https://img.shields.io/badge/Docker-19.03.9-blue)](https://www.docker.com/)
[![Docker-compose](https://img.shields.io/badge/Docker--compose-1.29.2-blue)](https://github.com/docker/compose/releases)
[![GNU Make](https://img.shields.io/badge/GNU%20Make-4.2.1-lightgrey)](https://www.gnu.org/software/make/)
[![GNU Bash](https://img.shields.io/badge/GNU%20Bash-4.2.1-lightgrey)](https://www.gnu.org/software/bash/)

## Table of Contents
* [TL;DR](#go-project-template)
* [Prerequisites](#prerequisites)
* [About the Project](#about-the-project)
* [Environment Variables](#environment-variables)
* [Main Commands](#main-commands)
  * [Up](#up)
  * [Down](#down)
  * [Testing](#testing)
  * [Documentation](#documentation)

## About The Project

This is a project template for golang API projects based on [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html),  [Hexagonal Architecture](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749) and [DDD](https://martinfowler.com/tags/domain%20driven%20design.html).

Aiming to be simple as possible, this project follows [Golang Standard Project Layout](https://github.com/golang-standards/project-layout) structure applying code patterns found in [Uber Style Guide](https://github.com/uber-go/guide/blob/master/style.md) being a guide with useful ideas on writing Go code in general.

We use these following technologies in this project:
* [echo](https://github.com/labstack/echo);
* [aws-lambda](https://aws.amazon.com/lambda/);
* [postgres](https://www.postgresql.org/);
* [redis](https://redis.io/);
* [docker](https://www.docker.com/);
* [docker-compose](https://github.com/docker/compose/);
* and other technologies.

## Environment Variables

Configure the following environment variables for local project execution:

| Variable                          | Default Value                                            |
| --------------------------------- | -------------------------------------------------------- |
| SERVICE_NAME                      | go-project-template                                      |
| SERVICE_VERSION                   | 0.0.1                                                    |
| ENVIRONMENT                       | local                                                    |
| LAMBDA                            | false                                                    |
| LOG_LEVEL                         | debug                                                    |
| SERVER_HOST                       | 0.0.0.0                                                  |
| SERVER_PORT                       | 5000                                                     |
| SERVER_TIMEOUT                    | 30                                                       |
| DB_READ_HOST                      | localhost                                                |
| DB_READ_PORT                      | 5432                                                     |
| DB_READ_NAME                      | go-project-template                                      |
| DB_READ_USERNAME                  | postgres                                                 |
| DB_READ_PASSWORD                  | postgres123                                              |
| DB_READ_LAZY_CONNECTION           | true                                                     |
| DB_READ_MIN_CONNECTIONS           | 2                                                        |
| DB_READ_MAX_CONNECTIONS           | 10                                                       |
| DB_READ_CONNECTION_MAX_LIFE_TIME  | 900                                                      |
| DB_READ_CONNECTION_MAX_IDLE_TIME  | 60                                                       |
| DB_WRITE_HOST                     | localhost                                                |
| DB_WRITE_PORT                     | 5432                                                     |
| DB_WRITE_NAME                     | go-project-template                                      |
| DB_WRITE_USERNAME                 | postgres                                                 |
| DB_WRITE_PASSWORD                 | postgres123                                              |
| DB_WRITE_LAZY_CONNECTION          | true                                                     |
| DB_WRITE_MIN_CONNECTIONS          | 2                                                        |
| DB_WRITE_MAX_CONNECTIONS          | 10                                                       |
| DB_WRITE_CONNECTION_MAX_LIFE_TIME | 900                                                      |
| DB_WRITE_CONNECTION_MAX_IDLE_TIME | 60                                                       |
| REDIS_HOST                        | localhost                                                |
| REDIS_PORT                        | 6379                                                     |
| REDIS_PASSWORD                    |                                                          |
| REDIS_DB                          | 1                                                        |
| REDIS_LAZY_CONNECTION             | true                                                     |

Environment variables separated by semicolon:

```
SERVICE_NAME=go-project-template;SERVICE_VERSION=0.0.1;ENVIRONMENT=local;LAMBDA=false;LOG_LEVEL=debug;SERVER_HOST=0.0.0.0;SERVER_PORT=5000;SERVER_TIMEOUT=30;DB_READ_HOST=localhost;DB_READ_PORT=5432;DB_READ_NAME=go-project-template;DB_READ_USERNAME=postgres;DB_READ_PASSWORD=postgres123;DB_READ_LAZY_CONNECTION=true;DB_READ_MIN_CONNECTIONS=2;DB_READ_MAX_CONNECTIONS=10;DB_READ_CONNECTION_MAX_LIFE_TIME=900;DB_READ_CONNECTION_MAX_IDLE_TIME=60;DB_WRITE_HOST=localhost;DB_WRITE_PORT=5432;DB_WRITE_NAME=go-project-template;DB_WRITE_USERNAME=postgres;DB_WRITE_PASSWORD=postgres123;DB_WRITE_LAZY_CONNECTION=true;DB_WRITE_MIN_CONNECTIONS=2;DB_WRITE_MAX_CONNECTIONS=10;DB_WRITE_CONNECTION_MAX_LIFE_TIME=900;DB_WRITE_CONNECTION_MAX_IDLE_TIME=60;REDIS_HOST=localhost;REDIS_PORT=6379;REDIS_PASSWORD=;REDIS_DB=1;REDIS_LAZY_CONNECTION=true;
```

## Main Commands

### Up
To start application and all dependenciove in this project, run:
```shell
make up
```

### Down
To remove all docker containers downloaded and installed by this project, run:
```shell
make down
```

### Testing
In progress ...

### Documentation
In progress ...

