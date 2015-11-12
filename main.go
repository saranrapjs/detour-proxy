package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/saranrapjs/detour-proxy/proxy"
)

func main() {
	const (
		defaultPort                 = ":80"
		defaultPortUsage            = "default listening port, ':80', ':8080'..."
		defaultTarget               = "http://127.0.0.1:8080"
		defaultTargetUsage          = "default redirect url, 'http://127.0.0.1:8080'"
		defaultWhiteRoutes          = `^/svc/community/personas/*`
		defaultWhiteRoutesUsage     = "list of routes to detour, as regexp, '/path1*,/path2*...."
		defaultDestinationHost      = "http://www.stg.gtm.nytimes.com"
		defaultDestinationHostUsage = "destination for non-detoured routes"
	)

	// flags
	port := flag.String("port", defaultPort, defaultPortUsage)

	url := flag.String("url", defaultTarget, defaultTargetUsage)
	destinationHost := flag.String("destination", defaultDestinationHost, defaultDestinationHostUsage)

	routesRegexp := flag.String("routes", defaultWhiteRoutes, defaultWhiteRoutesUsage)

	flag.Parse()

	fmt.Printf("🚧  detour proxy will run on %s 🚧\n", *port)
	fmt.Printf("...detouring requests to %s on %s\n", *url, *routesRegexp)
	fmt.Printf("......all other requests go to %s\n", *destinationHost)

	reg, _ := regexp.Compile(*routesRegexp)
	routes := []*regexp.Regexp{reg}

	// much of this owes to http://www.darul.io/post/2015-07-22_go-lang-simple-reverse-proxy
	px := proxy.New(*url, *destinationHost)
	px.RoutePatterns = routes

	// server
	http.HandleFunc("/", px.Handle)
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
