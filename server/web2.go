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
// 这个我们就输出了hello word,然后我们从源码来解析这个东西，我们看到最后的main函数执行的是HandleFunc这个函数我们从源代码中找到这段的源代码来看如下
 
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}
// pattern是解析的路径的字符串，然后执行一个handler的函数方法，如上例子我们传入的hello,他会执行DefaultServeMux,我们在查看源代码的时候会看到var DefaultServeMux = NewServeMux()我们再查看NewServeMux这个源代码
 
func NewServeMux() *ServeMux {
        return &ServeMux{m: make(map[string]muxEntry)} 
}
//而里边的返回一个新的ServeMux
type ServeMux struct {
    // contains filtered or unexported fields
}
// 所以我们就可以这样字
 
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