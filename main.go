package main

import (
	"log"
	"net/http"

	"github.com/yosssi/go-fileserver"
)

func main() {
	fs := fileserver.New(fileserver.Options{})
	http.Handle("/", fs.Serve(http.Dir("/static")))
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}