package main

// based on https://gist.github.com/hauxe/f2ea1901216177ccf9550a1b8bd59178

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

// FileSystem custom file system handler
type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func main() {
	port := flag.String("p", "80", "port to serve on")
	directory := flag.String("d", "/static", "the directory of static file to host")
	flag.Parse()

	fileServer := http.FileServer(FileSystem{http.Dir(*directory)})
	http.Handle("/", fileServer)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}