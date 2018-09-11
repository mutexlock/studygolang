# http包
- http请求最后都是调用transport包里的RoundTrip方法

```
    type RoundTripper interface {
      RoundTrip(*Request) (*Response, error)
    }
```

# Server
```
http.Handle("/foo", fooHandler)
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})
log.Fatal(http.ListenAndServe(":8080", nil))
```
一般情况，我们使用以上的代码进行httpserver的监听，其中主要调用的方法是 ListenAndServe,如下
```
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```
首先，会根据地址和handler来创建Server，之后调用server.ListenAndServe()
