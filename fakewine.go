package FakeWine

import (
	"net/http"
	"sync"
)

/**
*
* @author Liu Weiyi
* @date 2019-03-14 15:58
 */

type Handle func(ctx *Context)

type H map[string]interface{}

type HandlersChain []Handle

type Fengine struct {
	GRouter
	pool  sync.Pool
	trees methodTrees
}

func (engine *Fengine) handleFunc(ctx *Context) {
	path := ctx.request.URL.Path
	root := engine.trees.get(ctx.request.Method)
	handles, params, _ := root.getValue(path, Params{}, false)

	ctx.handles = handles
	ctx.params = params
	ctx.Next()

	ctx.writer.writeHeaderNow()

}

func (engine *Fengine) addRoute(path, method string, handles ...Handle) {
	root := engine.trees.get(method)
	if root == nil {
		root = new(node)
		engine.trees = append(engine.trees, methodTree{method: method, root: root})
	}
	root.addRoute(path, handles)
}

func (engine *Fengine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := engine.pool.Get().(*Context)

	c.writer.reset(w)
	c.request = r
	c.reset()

	engine.handleFunc(c)
	engine.pool.Put(c)
}

func (engine *Fengine) Run(address string) error {
	return http.ListenAndServe(address, engine)
}

func New() *Fengine {
	fake := &Fengine{
		GRouter: GRouter{
			path: "/",
			root: true,
		},
		trees: make(methodTrees, 0, 9),
	}
	fake.GRouter.engine = fake
	fake.pool.New = func() interface{} {
		return &Context{engine: fake}
	}
	return fake
}
