nanoserv
========

Nano http web server written in go lang.

###Features:

* custom 404 error message
* screen requests log
* serve static files from the run directory


###Install:

* `wget https://raw.githubusercontent.com/genesem/nanoserv/master/nanoserv.go`
* `go build nanoserv.go && ./nanoserv`

-----

By default web server listen on :3000 tcp port.

This can be changed altering the line #5: const lPort = ":3000" to desired port.
