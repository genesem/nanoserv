package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//const lPort = ":3000" // tcp4 port to serve
//const lDir = "."      // root directory, ex: "./web" for subdir with name 'web'

var (
	addr = flag.String("addr", ":3000", "tcp4 host and port to listen")
	root = flag.String("root", ".", "root directory")
)

const msg404 = `
<html lang="en">
 <head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width">
  <title>404 - not found</title>
<style>
* { padding: 0; margin: 0; }
html {
    -webkit-font-smoothing: antialiased;
    background-color: #fafafa;}
body {
    font-family: 'Helvetica Neue', Helvetica, Arial, 'lucida grande', tahoma, verdana, arial, sans-serif;
    color: #222222;
    margin: 0;
    line-height: 1.5em;}</style>
</head><body>
  <div style="padding:40px;"> 
  <h2>404 - not found</h2><hr size="1" width="30%"/>
  <br/><a href="/">go back</a></div>
</body></html>`

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, msg404)
}

type logServer struct {
	hdl http.Handler
}

func (l *logServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "nanoServ")
	fi, err := os.Stat(*root + r.RequestURI)
	if fi == nil && err != nil {
		log.Printf("404: %s\n", r.RequestURI) //"stat /does/not/exist: no such file or directory"
		notFoundHandler(w, r)
		return
	}
	log.Println(r.Method, r.RequestURI)
	l.hdl.ServeHTTP(w, r)
}

func main() {

	flag.Parse()
	srv := &http.Server{
		Addr:           *addr,
		Handler:        &logServer{hdl: http.FileServer(http.Dir(*root))},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 15,
	}

	log.Printf("webServer started at %v\n\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
