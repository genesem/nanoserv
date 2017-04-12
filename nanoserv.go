package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

var root string // root directory

const (
	sign   = "nanoServ"
	msg404 = `
<html lang="en">
 <head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width">
  <title>404 - not found</title>
   <style>
	* { padding:0; margin:0; }
	html {
	    -webkit-font-smoothing: antialiased;
	    background-color:#fafaf0;}
	body {
	    font-family: Helvetica, Arial, Verdana;
	    color:#222222; line-height:1.5em;}
   </style>
  </head>
<body>
  <div style="padding:40px;"> 
  <h2>404 - not found</h2><hr size="1" width="30%"/>
  <br/><a href="/">go back</a></div>
</body></html>`
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, msg404)
}

type logServer struct {
	hdl http.Handler
}

func (l *logServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", sign)

	rq := r.URL.Path

	fi, err := os.Stat(root + rq)
	if fi == nil && err != nil {

		log.Printf("404: %s\n", rq) //does not exist
		notFoundHandler(w, r)
		return
	}
	log.Println(r.Method, r.RequestURI)
	l.hdl.ServeHTTP(w, r)
}

func init() {
	const (
		defRoot = "."
		usage   = "root directory"
	)
	flag.StringVar(&root, "root", defRoot, usage)
	flag.StringVar(&root, "r", defRoot, usage+" (shorthand)")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = ":3000"

	} else {
		port = ":" + port
	}

	addr := flag.String("addr", port, "tcp4 host and port to listen")
	flag.Parse()
	root = path.Clean(root)

	srv := &http.Server{
		Addr:           *addr,
		Handler:        &logServer{hdl: http.FileServer(http.Dir(root))},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 15,
	}

	log.Printf("webServer started at %v with root=%v\n\n", srv.Addr, root)
	log.Fatal(srv.ListenAndServe())
}
