# detour-proxy

![under construction](https://cloud.githubusercontent.com/assets/707098/11125216/b5554b94-8936-11e5-921d-c03b0ef82c3c.png)

a tiny proxy for dirverting pattern-matched http requests to a local server

#### Usage

```bash
$ detour-proxy -h
Usage of detour-proxy:
  -destination string
    	destination for non-detoured routes (default "http://www.stg.gtm.nytimes.com")
  -port string
    	default listening port, ':80', ':8080'... (default ":80")
  -routes string
    	list of routes to detour, as regexp, '/path1*,/path2*.... (default "^/svc/community/personas/*")
  -url string
    	default redirect url, 'http://127.0.0.1:8080' (default "http://127.0.0.1:8080")
```

#### Daps

much of the code owes itself to http://www.darul.io/post/2015-07-22_go-lang-simple-reverse-proxy