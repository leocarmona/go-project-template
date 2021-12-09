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

