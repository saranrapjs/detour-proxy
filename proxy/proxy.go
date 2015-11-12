package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

// Proxy is a proxy ;)
type Proxy struct {
	detourProxy   *httputil.ReverseProxy
	defaultProxy  *httputil.ReverseProxy
	RoutePatterns []*regexp.Regexp // add some route patterns with regexp
}

// New creates a new Proxy
func New(detourURL string, defaultURL string) *Proxy {
	dt, _ := url.Parse(detourURL)
	df, _ := url.Parse(defaultURL)

	return &Proxy{
		detourProxy:  httputil.NewSingleHostReverseProxy(dt),
		defaultProxy: httputil.NewSingleHostReverseProxy(df),
	}
}

// Handle deals w/ proxy requests
func (p *Proxy) Handle(w http.ResponseWriter, r *http.Request) {
	if p.RoutePatterns != nil && p.parseWhiteList(r) {
		log.Println("🚧 => ", r.URL)
		p.detourProxy.ServeHTTP(w, r)
	} else {
		p.defaultProxy.ServeHTTP(w, r)
	}
}

func (p *Proxy) parseWhiteList(r *http.Request) bool {
	for _, regexp := range p.RoutePatterns {
		if regexp.MatchString(r.URL.Path) {
			// let's forward it
			return true
		}
	}
	return false
}
