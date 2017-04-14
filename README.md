nanoserv
========

Nano http web server written in go lang.

[![Build Status](https://travis-ci.org/genesem/nanoserv.svg?branch=master)](https://travis-ci.org/genesem/nanoserv)

#### Features:

* custom 404 error message
* screen requests log
* serve static files from the run directory but it is changeable
* no third party dependencies only core go lang packages used


#### Install:

  `wget https://raw.githubusercontent.com/genesem/nanoserv/master/nanoserv.go`

  `go build nanoserv.go && ./nanoserv`

#### Usage:

By default web server listen on :3000 tcp port.

How to use custom port and web root: `./nanoserv -addr=:3001 -root=/var/www`.

Please don't use trailing slash in the `root` value, i.e. `-root=/var/www/` is incorrect.

Use `./nanoserv -h` for help with flags.

#### Note:

Nanoserv is now using `os.LookupEnv("PORT")` instead of os.GetEnv("PORT"),

so the minimum version of the Go compiler required is Go 1.5.



