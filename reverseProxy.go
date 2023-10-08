package main

import (
        "fmt"
        "net/http"
        "net/http/httputil"
        "net/url"
        "os"
)

const (
        certPath = "/etc/letsencrypt/live/chestergould.ca/fullchain.pem"
        keyPath = "/etc/letsencrypt/live/chestergould.ca/privkey.pem"
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
        httpsUrl.Host = r.Host
        httpsUrl.Scheme = "https"

        http.Redirect(w, r, httpsUrl.String(), http.StatusPermanentRedirect)
}

func main () {
        go http.ListenAndServe(":80", http.HandlerFunc(redirectHTTPS))

        websiteUrl, err := url.Parse("http://localhost:3141")
        printErrorAndExit(err)
        pleromaUrl, err := url.Parse("http://localhost:4000")
        printErrorAndExit(err)

        /*
        searchUrl, err := url.Parse("http://localhost:2718")
        printErrorAndExit(err)
        */

        http.Handle("chestergould.ca/", httputil.NewSingleHostReverseProxy(websiteUrl))
        http.Handle("www.chestergould.ca/", httputil.NewSingleHostReverseProxy(websiteUrl))
        http.Handle("pleroma.chestergould.ca/", httputil.NewSingleHostReverseProxy(pleromaUrl))
        /*
        http.Handle("search.chestergould.xyz/", httputil.NewSingleHostReverseProxy(searchUrl))
        */
        err = http.ListenAndServeTLS(":443", certPath, keyPath, nil)
        printErrorAndExit(err)
}
