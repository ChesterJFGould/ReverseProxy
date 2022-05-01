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

	http.Handle("chestergould.xyz/", httputil.NewSingleHostReverseProxy(websiteUrl))
	http.Handle("www.chestergould.xyz/", httputil.NewSingleHostReverseProxy(websiteUrl))
	http.Handle("search.chestergould.xyz/", httputil.NewSingleHostReverseProxy(searchUrl))
	http.ListenAndServeTLS(":80", "test.server.crt", "test.server.key", nil)
}
