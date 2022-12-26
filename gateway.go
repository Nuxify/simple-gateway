package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/go-chi/chi"
)

// GatewayController handles http request to the target service
func GatewayController(w http.ResponseWriter, r *http.Request) {
	service := strings.ToUpper(chi.URLParam(r, "service"))

	// check if service is in plural form (e.g USERS instead of USER)
	if service[len(service)-1:] == "S" {
		service = CastPluralToSingularForm(service)
	}

	// replace '-' with '_'
	service = strings.ReplaceAll(service, "-", "_")

	ServeReverseProxy(os.Getenv(fmt.Sprintf("%s_API_URL", service)), w, r)
}

// ServeReverseProxy serve a reverse proxy for a given url
func ServeReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}
