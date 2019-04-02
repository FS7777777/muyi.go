package middleware

//中间件
type middleware func(Handler) Handler

type Router struct {
	middlewareChain [] middleware
	Mux map[string] Handler
}

func NewRouter() *Router{
	return &Router{
		Mux:make(map[string]Handler),
	}
}

func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(route string, h Handler) {
	var mergedHandler = h

	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}

	r.Mux[route] = mergedHandler
}