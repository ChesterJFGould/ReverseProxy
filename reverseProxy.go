package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const (
	certPath = "/etc/letsencrypt/live/chestergould.xyz/fullchain.pem"
	keyPath = "/etc/letsencrypt/live/chestergould.xyz/privkey.pem"
)

func printErrorAndExit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func redirectHTTPS(w http.ResponseWriter, r *http.Request) {
	var httpsUrl url.URL

	httpsUrl = *r.URL
	httpsUrl.Scheme = "https"

	http.Redirect(w, r, httpsUrl.String(), http.StatusPermanentRedirect)
}

func main () {
	go http.ListenAndServe(":80", http.HandlerFunc(redirectHTTPS))

	websiteUrl, err := url.Parse("http://localhost:3141")
	printErrorAndExit(err)

	searchUrl, err := url.Parse("http://localhost:2718")
	printErrorAndExit(err)

	http.Handle("chestergould.xyz/", httputil.NewSingleHostReverseProxy(websiteUrl))
	http.Handle("www.chestergould.xyz/", httputil.NewSingleHostReverseProxy(websiteUrl))
	http.Handle("search.chestergould.xyz/", httputil.NewSingleHostReverseProxy(searchUrl))
	err = http.ListenAndServeTLS(":443", certPath, keyPath, nil)
	printErrorAndExit(err)
}
