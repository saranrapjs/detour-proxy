# detour-proxy

![under construction](https://cloud.githubusercontent.com/assets/707098/11125298/2cc30770-8937-11e5-9bec-97c5bb4da0f1.gif)

a tiny proxy for dirverting pattern-matched http requests to a local server

#### Usage

```bash
$ detour-proxy -h
Usage of detour-proxy:
  -destination string
    	destination for non-detoured routes (default "http://www.stg.gtm.nytimes.com")
  -from string
    	(optional) manage the /etc/hosts entry for the intended detour 'from' host    	
  -port string
    	default listening port, ':80', ':8080'... (default ":80")
  -routes string
    	list of routes to detour, as regexp, '/path1*,/path2*.... (default "^/svc/community/personas/*")
  -url string
    	default redirect url, 'http://127.0.0.1:8080' (default "http://127.0.0.1:8080")
```

##### Example

This will direct requests from `http://www.stg.nytimes.com`:
- to `http://127.0.0.1:8080` IF they match `^/svc/community/personas/*`
- ELSE they will be directed to `http://www.stg.gtm.nytimes.com`

```
detour-proxy \
	-from www.stg.nytimes.com \
	-routes "^/svc/community/personas/*" \
	-url "http://127.0.0.1:8080" \
	-destination "http://www.stg.gtm.nytimes.com"
````

#### Daps

much of the code owes itself to http://www.darul.io/post/2015-07-22_go-lang-simple-reverse-proxy