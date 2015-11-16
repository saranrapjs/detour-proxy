package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/saranrapjs/detour-proxy/proxy"
	"github.com/saranrapjs/goodhosts"
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

	fromHost := flag.String("from", "", "(optional) manage the /etc/hosts entry for the intended detour 'from' host")

	flag.Parse()

	if *fromHost != "" {
		hosts, err := goodhosts.NewHosts()
		if err != nil {
			log.Fatal("Hosts editor: ", err)
		}
		if hosts.Has("127.0.0.1", *fromHost) {
			log.Fatal("The detour 'from' host already has an entry in '/etc/hosts'; please remove this before proceeding, or disable the 'from' flag")
		}
		hosts.Add("127.0.0.1", *fromHost)
		if err := hosts.Flush(); err != nil {
			log.Fatal("An error editing hosts: ", err)
		}
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, syscall.SIGTERM)
		go func() {
			<-c
			fmt.Println("Cleaning up /etc/hosts...")
			hosts.Remove("127.0.0.1", *fromHost)
			hosts.Flush()
			os.Exit(1)
		}()

	}

	fmt.Printf("🚧  detour proxy will run on %s 🚧\n", *port)
	fmt.Printf("...detouring requests to %s on %s\n", *url, *routesRegexp)
	fmt.Printf("......all other requests go to %s\n", *destinationHost)
	if *fromHost != "" {
		fmt.Printf(".........and mapping requests from %s to %s\n", *fromHost, *url)
	}

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
