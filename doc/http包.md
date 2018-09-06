# http包
- http请求最后都是调用transport包里的RoundTrip方法

```
    type RoundTripper interface {
      RoundTrip(*Request) (*Response, error)
    }
```
