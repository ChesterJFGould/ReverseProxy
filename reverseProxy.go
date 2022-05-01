package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func printErrorAndExit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func main () {
	websiteUrl, err := url.Parse("http://localhost:3141")
	printErrorAndExit(err)

	searchUrl, err := url.Parse("http://localhost:2718")
	printErrorAndExit(err)

	http.Handle("localhost/", httputil.NewSingleHostReverseProxy(websiteUrl))
	http.Handle("search.localhost/", httputil.NewSingleHostReverseProxy(searchUrl))
	http.ListenAndServeTLS(":8080", "test.server.crt", "test.server.key", nil)
}
