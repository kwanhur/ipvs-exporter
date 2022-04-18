# ipvs-exporter

[![Build Status](https://travis-ci.org/kwanhur/ipvs-exporter.svg?branch=master)](https://travis-ci.org/kwanhur/ipvs-exporter)
[![Docker Pulls](https://img.shields.io/docker/pulls/kwanhur/ipvs-exporter.svg)](https://hub.docker.com/r/kwanhur/ipvs-exporter)
[![Github All Releases](https://img.shields.io/github/downloads/kwanhur/ipvs-exporter/total.svg)](https://github.com/kwanhur/ipvs-exporter)
[![GitHub release](https://img.shields.io/github/release/kwanhur/ipvs-exporter.svg)](https://github.com/kwanhur/ipvs-exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/kwanhur/ipvs-exporter)](https://goreportcard.com/report/github.com/kwanhur/ipvs-exporter)

Simple server that scrapes Linux kernel module ip_vs stats through [ipvs](https://github.com/moby/ipvs) and exports them via HTTP for Prometheus consumption

This repo starts with [nginx-vts-exporter](https://github.com/hnlq715/nginx-vts-exporter), many thanks to [hnlq715](https://github.com/hnlq715).

## Table of Contents
* [Dependency](#dependency)
* [Download](#download)
* [Compile](#compile)
  * [build binary](#build-binary)
  * [build docker image](#build-docker-image)
* [Run](#run)
  * [run binary](#run-binary)
  * [run docker image](#run-docker)
* [Environment variables](#environment-variables)
* [Authors](#authors)
* [Copyright and License](#copyright-and-license)

## Dependency

* [ipvs](https://github.com/moby/ipvs)
* [Prometheus](https://prometheus.io/)
* [Golang](https://golang.org/)

[Back to TOC](#table-of-contents)

## Download

Binary can be downloaded from [Releases](https://github.com/kwanhur/ipvs-exporter/releases) page.

[Back to TOC](#table-of-contents)

## Compile

### build binary

``` shell
make
```

### build RPM package
``` shell
make rpm
```

### build docker image
``` shell
make docker
```

## Docker Hub Image
``` shell
docker pull kwanhur/ipvs-exporter:latest
```
It can be used directly instead of having to build the image yourself.
([Docker Hub kwanhur/ipvs-exporter](https://hub.docker.com/r/kwanhur/ipvs-exporter/))

[Back to TOC](#table-of-contents)

## Run

### run binary
``` shell
nohup /usr/bin/ipvs-exporter
```

### run docker
```
docker run  -ti --rm kwanhur/ipvs-exporter
```

[Back to TOC](#table-of-contents)

## Environment variables

This image is configurable using different env variables

| Variable name    | Default  | Description                      |
|------------------|----------|----------------------------------|
| METRICS_ENDPOINT | /metrics | Metrics endpoint exportation URI |
| METRICS_ADDR     | :9911    | Metrics exportation address:port |
| METRICS_NS       | ipvs     | Prometheus metrics Namespaces    |

[Back to TOC](#table-of-contents)

## Authors

kwanhur <huang_hua2012@163.com>

[Back to TOC](#table-of-contents)

## Copyright and License

This module is licensed under the Apache License 2.0 .

Copyright (C) 2020, by kwanhur <huang_hua2012@163.com>

All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

[Back to TOC](#table-of-contents)