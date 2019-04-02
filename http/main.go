package main

import
(   "fmt"
	"../http/middleware"
)


func log(next middleware.Handler) middleware.Handler{
	return middleware.HandlerFunc(func(i int) {
		fmt.Println("log in")
		//next 方法实现了ServeHTTP接口，调用ServeHTTP实际会触发自身代码
		next.ServeHTTP(i)
		fmt.Println("log out")
	})
}

func time(next middleware.Handler) middleware.Handler{
	return middleware.HandlerFunc(func(i int) {
		fmt.Println("time in")
		//next 方法实现了ServeHTTP接口，调用ServeHTTP实际会触发自身代码
		next.ServeHTTP(i)
		fmt.Println("time out")
	})
}


func main() {
	router := middleware.NewRouter()
	router.Use(log)
	router.Use(time)
	router.Add("/",middleware.HandlerFunc(func(i int) {
		fmt.Println("login")
	}))
	router.Mux["/"].ServeHTTP(8)
}
