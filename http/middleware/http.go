package middleware

// 模拟http Handler处理流程
type Handler interface {
	ServeHTTP(i int)
}
//实现了ServeHTTP接口，可将同签名函数转化成Handler
type HandlerFunc func(i int)

//HandlerFunc方法实现接口
func (h HandlerFunc)ServeHTTP(i int)  {
	h(i)
}

//实际上只要你的handler函数签名是：
//
//func (ResponseWriter, *Request)
//那么这个handler和HandlerFunc()就有了一致的函数签名，可以将该handler()函数进行类型转换，转为HandlerFunc。
// 而HandlerFunc实现了Handler这个接口。
// 在http库需要调用你的handler函数来处理http请求时，会调用HandlerFunc()的ServeHTTP()函数，可见一个请求的基本调用链是这样的：
//
//h = getHandler() => h.ServeHTTP(w, r) => h(w, r)