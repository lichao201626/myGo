package main
import (
    "io"
    "net/http"
    //"strings"
    "time"
)
var (
    server = &http.Server{
        Addr:           ":9090",
        Handler:        &ppserver{},
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    handlersMap = make(map[string]HandlersFunc)
)
type ppserver struct {
}
 
func (*ppserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if h, ok := handlersMap[r.URL.String()]; ok {
        h(w, r)
    }
    //io.WriteString(w, "URL"+r.URL.String())
}
 
func f1(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "111111111111")
}
 
func f2(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "2222222222222")
}
 
type HandlersFunc func(http.ResponseWriter, *http.Request)
 
func Hello(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello"))
}
 
func main() {
 
    handlersMap["/hello"] = Hello
    handlersMap["/f1"] = f1
    handlersMap["/f2"] = f2
 
    server.ListenAndServe()
 
}


web1.go
 
package main
 
import (
    "io"
    "net/http"
)
 
func hello(rw http.ResponseWriter, req *http.Request) {
    io.WriteString(rw, "hello widuu")
}
 
func main() {
    http.HandleFunc("/", hello)  //设定访问的路径
    http.ListenAndServe(":8080", nil) //设定端口和handler
}
这个我们就输出了hello word,然后我们从源码来解析这个东西，我们看到最后的main函数执行的是HandleFunc这个函数我们从源代码中找到这段的源代码来看如下
 
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}
pattern是解析的路径的字符串，然后执行一个handler的函数方法，如上例子我们传入的hello,他会执行DefaultServeMux,我们在查看源代码的时候会看到var DefaultServeMux = NewServeMux()我们再查看NewServeMux这个源代码
 
func NewServeMux() *ServeMux {
        return &ServeMux{m: make(map[string]muxEntry)} 
}
//而里边的返回一个新的ServeMux
type ServeMux struct {
    // contains filtered or unexported fields
}
所以我们就可以这样字
 
//申明一个ServeMux结构
type MyHandler struct{}
mux := http.NewServeMux()
//我们可以通过一下 http提供的
func Handle(pattern string, handler Handler)  //第一个是我们的路径字符串 第二个是这样的 是个接口
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
//实现这个接口我们就要继承ServeHTTP这个方法 所以代码
func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "URL"+r.URL.String())
}
//我们查看源代码
func (mux *ServeMux) Handle(pattern string, handler Handler) //这个新的ServeMux低下的Handle来设置 这里的Handler也是Handler interface所以我们将这个
mux.Handle("/", &MyHandler{})
mux.HandleFunc("/hello", sayHello)// 我们一样可以通过handleFunc设置
//源代码func ListenAndServe(addr string, handler Handler) error 
http.ListenAndServe(":8080",mux) //所以我们把mux传进去
web2.go
 
//完整代码
package main
 
import (
    "io"
    "net/http"
)
 
type MyHandle struct{}
 
func main() {
    mux := http.NewServeMux()
    mux.Handle("/", &MyHandle{})
    http.ListenAndServe(":8080", mux)
}
 
func (*MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "URL"+r.URL.String())
}
然后我们继续深入点
 
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}
//返回的serve我们查看它的结构
type Server struct {
    Addr           string        // TCP address to listen on, ":http" if empty
    Handler        Handler       // handler to invoke, http.DefaultServeMux if nil
    ReadTimeout    time.Duration // maximum duration before timing out read of the request
    WriteTimeout   time.Duration // maximum duration before timing out write of the response
    MaxHeaderBytes int           // maximum size of request headers, DefaultMaxHeaderBytes if 0
    TLSConfig      *tls.Config   // optional TLS config, used by ListenAndServeTLS
 
    // TLSNextProto optionally specifies a function to take over
    // ownership of the provided TLS connection when an NPN
    // protocol upgrade has occurred.  The map key is the protocol
    // name negotiated. The Handler argument should be used to
    // handle HTTP requests and will initialize the Request's TLS
    // and RemoteAddr if not already set.  The connection is
    // automatically closed when the function returns.
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
}
//我们自己设置
type MyHandle struct{}
server := http.Server{
        Addr:        ":8080",
        Handler:     &MyHandle{},
        ReadTimeout: 6 * time.Second,
    }
//我们查看过了 我们要实现路由分发映射就待这样 我们看到了下边的f 是一个HandlerFunc类型
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
//所以我们申明一下
var mux map[string]func(http.ResponseWriter, *http.Request)
mux = make(map[string]func(http.ResponseWriter, *http.Request))
mux["/hello"] = hello
mux["/bye"] = bye
err := server.ListenAndServe()
if err != nil {
    log.Fatal(err)
}
//这样我们就可以做到了路径的映射 
func (*MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if h, ok := mux[r.URL.String()]; ok {
        h(w, r)
    }
    io.WriteString(w, "URL"+r.URL.String())
}
 
func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello 模块")
}
 
func bye(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "bye 模块")
}
//可能有人不懂mux["/hello"] = hello 然后低下的h(w,r) 我简单的解释一下 看个例子 go里边都可以是类型
type test func(int) bool //定一个test的func(int) bool 类型
func isAdd(i int) bool {
    if i%2 == 0 {
        return false
    }
    return true
}
 
func filter(s []int, f test) []int {
    var result []int
    for _, v := range s {
        if f(v) {
            result = append(result, v)
        }
    }
    return result
}
 
func main() {
    slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
    b := filter(slice, isAdd)
    fmt.Println(b)
}
//是不是懂点了 其实就类似于
f:=func(x int){
  fmt.Println("hello")
}
f();
web3.go
 
package main
 
import (
    "io"
    "log"
    "net/http"
    "time"
)
 
var mux map[string]func(http.ResponseWriter, *http.Request)
 
func main() {
    server := http.Server{
        Addr:        ":8080",
        Handler:     &MyHandle{},
        ReadTimeout: 6 * time.Second,
    }
    mux = make(map[string]func(http.ResponseWriter, *http.Request))
    mux["/hello"] = hello
    mux["/bye"] = bye
    err := server.ListenAndServe()
 
    if err != nil {
        log.Fatal(err)
    }
}
 
type MyHandle struct{}
 
func (*MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if h, ok := mux[r.URL.String()]; ok {
        h(w, r)
    }
    io.WriteString(w, "URL"+r.URL.String())
}
 
func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello 模块")
}
 
func bye(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "bye 模块")
}