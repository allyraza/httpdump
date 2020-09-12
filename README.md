### httpdump

http handler to dump `http.Request` to stdout.

### How to use it?

```go
package main

import (
    // ...
    "http://github.com/allyraza/httpdump"
)

func main() {
    helloHandler := http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello"))
    })

    http.Handle("/dump", httpdump.RequestDumpHandler(helloHandler))

    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
```

### License

MIT
